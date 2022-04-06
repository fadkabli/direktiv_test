package flow

import (
	"bytes"
	"context"
	"errors"
	"io"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/direktiv/direktiv/pkg/flow/ent"
	"github.com/direktiv/direktiv/pkg/flow/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (flow *flow) InstanceAnnotation(ctx context.Context, req *grpc.InstanceAnnotationRequest) (*grpc.InstanceAnnotationResponse, error) {

	flow.sugar.Debugf("Handling gRPC request: %s", this())

	nsc := flow.db.Namespace

	d, err := flow.traverseToInstanceAnnotation(ctx, nsc, req.GetNamespace(), req.GetInstance(), req.GetKey())
	if err != nil {
		return nil, err
	}

	var resp grpc.InstanceAnnotationResponse

	resp.Namespace = d.ns().Name
	resp.Instance = d.in.ID.String()
	resp.Key = d.annotation.Name
	resp.CreatedAt = timestamppb.New(d.annotation.CreatedAt)
	resp.UpdatedAt = timestamppb.New(d.annotation.UpdatedAt)
	resp.Checksum = d.annotation.Hash
	resp.TotalSize = int64(d.annotation.Size)

	if resp.TotalSize > parcelSize {
		return nil, status.Error(codes.ResourceExhausted, "annotation too large to return without using the parcelling API")
	}

	resp.Data = d.annotation.Data

	return &resp, nil

}

func (flow *flow) InstanceAnnotationParcels(req *grpc.InstanceAnnotationRequest, srv grpc.Flow_InstanceAnnotationParcelsServer) error {

	flow.sugar.Debugf("Handling gRPC request: %s", this())

	ctx := srv.Context()

	nsc := flow.db.Namespace

	d, err := flow.traverseToInstanceAnnotation(ctx, nsc, req.GetNamespace(), req.GetInstance(), req.GetKey())
	if err != nil {
		return err
	}

	rdr := bytes.NewReader(d.annotation.Data)

	for {

		resp := new(grpc.InstanceAnnotationResponse)

		resp.Namespace = d.ns().Name
		resp.Instance = d.in.ID.String()
		resp.Key = d.annotation.Name
		resp.CreatedAt = timestamppb.New(d.annotation.CreatedAt)
		resp.UpdatedAt = timestamppb.New(d.annotation.UpdatedAt)
		resp.Checksum = d.annotation.Hash
		resp.TotalSize = int64(d.annotation.Size)

		buf := new(bytes.Buffer)
		k, err := io.CopyN(buf, rdr, parcelSize)
		if err != nil {

			if errors.Is(err, io.EOF) {
				err = nil
			}

			if err == nil && k == 0 {
				return nil
			}

			if err != nil {
				return err
			}

		}

		resp.Data = buf.Bytes()

		err = srv.Send(resp)
		if err != nil {
			return err
		}

	}

}

func (flow *flow) InstanceAnnotations(ctx context.Context, req *grpc.InstanceAnnotationsRequest) (*grpc.InstanceAnnotationsResponse, error) {

	flow.sugar.Debugf("Handling gRPC request: %s", this())

	p, err := getPagination(req.Pagination)
	if err != nil {
		return nil, err
	}

	opts := []ent.AnnotationPaginateOption{}
	opts = append(opts, annotationsOrder(p)...)
	opts = append(opts, annotationsFilter(p)...)

	nsc := flow.db.Namespace

	d, err := flow.getInstance(ctx, nsc, req.GetNamespace(), req.GetInstance(), false)
	if err != nil {
		return nil, err
	}

	query := d.in.QueryAnnotations()
	cx, err := query.Paginate(ctx, p.After(), p.First(), p.Before(), p.Last(), opts...)
	if err != nil {
		return nil, err
	}

	var resp grpc.InstanceAnnotationsResponse

	resp.Namespace = d.ns().Name
	resp.Instance = d.in.ID.String()

	err = atob(cx, &resp.Annotations)
	if err != nil {
		return nil, err
	}

	for i := range cx.Edges {

		edge := cx.Edges[i]
		annotation := edge.Node

		v := resp.Annotations.Edges[i].Node
		v.Checksum = annotation.Hash
		v.CreatedAt = timestamppb.New(annotation.CreatedAt)
		v.Size = int64(annotation.Size)
		v.UpdatedAt = timestamppb.New(annotation.UpdatedAt)

	}

	return &resp, nil

}

func (flow *flow) InstanceAnnotationsStream(req *grpc.InstanceAnnotationsRequest, srv grpc.Flow_InstanceAnnotationsStreamServer) error {

	flow.sugar.Debugf("Handling gRPC request: %s", this())

	ctx := srv.Context()
	phash := ""
	nhash := ""

	p, err := getPagination(req.Pagination)
	if err != nil {
		return err
	}

	opts := []ent.AnnotationPaginateOption{}
	opts = append(opts, annotationsOrder(p)...)
	opts = append(opts, annotationsFilter(p)...)

	nsc := flow.db.Namespace
	d, err := flow.getInstance(ctx, nsc, req.GetNamespace(), req.GetInstance(), false)
	if err != nil {
		return err
	}

	sub := flow.pubsub.SubscribeInstanceAnnotations(d.in)
	defer flow.cleanup(sub.Close)

resend:

	query := d.in.QueryAnnotations()
	cx, err := query.Paginate(ctx, p.After(), p.First(), p.Before(), p.Last(), opts...)
	if err != nil {
		return err
	}

	resp := new(grpc.InstanceAnnotationsResponse)

	resp.Namespace = d.ns().Name
	resp.Instance = d.in.ID.String()

	err = atob(cx, &resp.Annotations)
	if err != nil {
		return err
	}

	for i := range cx.Edges {

		edge := cx.Edges[i]
		annotation := edge.Node

		v := resp.Annotations.Edges[i].Node
		v.Checksum = annotation.Hash
		v.CreatedAt = timestamppb.New(annotation.CreatedAt)
		v.Size = int64(annotation.Size)
		v.UpdatedAt = timestamppb.New(annotation.UpdatedAt)

	}

	nhash = checksum(resp)
	if nhash != phash {
		err = srv.Send(resp)
		if err != nil {
			return err
		}
	}
	phash = nhash

	more := sub.Wait(ctx)
	if !more {
		return nil
	}

	goto resend

}

func (flow *flow) SetInstanceAnnotation(ctx context.Context, req *grpc.SetInstanceAnnotationRequest) (*grpc.SetInstanceAnnotationResponse, error) {

	flow.sugar.Debugf("Handling gRPC request: %s", this())

	tx, err := flow.db.Tx(ctx)
	if err != nil {
		return nil, err
	}
	defer rollback(tx)

	nsc := tx.Namespace
	annotationc := tx.Annotation

	key := req.GetKey()

	d, err := flow.getInstance(ctx, nsc, req.GetNamespace(), req.GetInstance(), false)
	if err != nil {
		return nil, err
	}

	var annotation *ent.Annotation
	var newVar bool

	annotation, newVar, err = flow.SetAnnotation(ctx, annotationc, d.in, key, req.GetData())
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	if newVar {
		flow.logToInstance(ctx, time.Now(), d.in, "Created instance annotation '%s'.", key)
	} else {
		flow.logToInstance(ctx, time.Now(), d.in, "Updated instance annotation '%s'.", key)

	}
	flow.pubsub.NotifyInstanceAnnotations(d.in)

	var resp grpc.SetInstanceAnnotationResponse

	resp.Namespace = d.ns().Name
	resp.Instance = d.in.ID.String()
	resp.Key = key
	resp.CreatedAt = timestamppb.New(annotation.CreatedAt)
	resp.UpdatedAt = timestamppb.New(annotation.UpdatedAt)
	resp.Checksum = annotation.Hash
	resp.TotalSize = int64(annotation.Size)

	return &resp, nil

}

func (flow *flow) SetInstanceAnnotationParcels(srv grpc.Flow_SetInstanceAnnotationParcelsServer) error {

	flow.sugar.Debugf("Handling gRPC request: %s", this())

	ctx := srv.Context()

	req, err := srv.Recv()
	if err != nil {
		return err
	}

	namespace := req.GetNamespace()
	instance := req.GetInstance()
	key := req.GetKey()

	totalSize := int(req.GetTotalSize())

	buf := new(bytes.Buffer)

	for {

		_, err = io.Copy(buf, bytes.NewReader(req.Data))
		if err != nil {
			return err
		}

		if req.TotalSize <= 0 {
			if buf.Len() >= totalSize {
				break
			}
		}

		req, err = srv.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		if req.TotalSize <= 0 {
			if buf.Len() >= totalSize {
				break
			}
		} else {
			if req == nil {
				break
			}
		}

		if int(req.GetTotalSize()) != totalSize {
			return errors.New("totalSize changed mid stream")
		}

	}

	if buf.Len() > totalSize {
		return errors.New("received more data than expected")
	}

	tx, err := flow.db.Tx(ctx)
	if err != nil {
		return err
	}
	defer rollback(tx)

	nsc := tx.Namespace
	annotationc := tx.Annotation

	d, err := flow.getInstance(ctx, nsc, namespace, instance, false)
	if err != nil {
		return err
	}

	var annotation *ent.Annotation
	var newVar bool

	annotation, newVar, err = flow.SetAnnotation(ctx, annotationc, d.in, key, req.GetData())
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	if newVar {
		flow.logToInstance(ctx, time.Now(), d.in, "Created instance annotation '%s'.", key)
	} else {
		flow.logToInstance(ctx, time.Now(), d.in, "Updated instance annotation '%s'.", key)
	}

	flow.pubsub.NotifyInstanceAnnotations(d.in)

	var resp grpc.SetInstanceAnnotationResponse

	resp.Namespace = d.ns().Name
	resp.Instance = d.in.ID.String()
	resp.Key = key
	resp.CreatedAt = timestamppb.New(annotation.CreatedAt)
	resp.UpdatedAt = timestamppb.New(annotation.UpdatedAt)
	resp.Checksum = annotation.Hash
	resp.TotalSize = int64(annotation.Size)

	err = srv.SendAndClose(&resp)
	if err != nil {
		return err
	}

	return nil

}

func (flow *flow) DeleteInstanceAnnotation(ctx context.Context, req *grpc.DeleteInstanceAnnotationRequest) (*emptypb.Empty, error) {

	flow.sugar.Debugf("Handling gRPC request: %s", this())

	tx, err := flow.db.Tx(ctx)
	if err != nil {
		return nil, err
	}
	defer rollback(tx)

	nsc := tx.Namespace

	d, err := flow.traverseToInstanceAnnotation(ctx, nsc, req.GetNamespace(), req.GetInstance(), req.GetKey())
	if err != nil {
		return nil, err
	}

	annotationc := tx.Annotation

	err = annotationc.DeleteOne(d.annotation).Exec(ctx)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	flow.logToInstance(ctx, time.Now(), d.in, "Deleted instance annotation '%s'.", d.annotation.Name)
	flow.pubsub.NotifyInstanceAnnotations(d.in)

	var resp emptypb.Empty

	return &resp, nil

}

func (flow *flow) RenameInstanceAnnotation(ctx context.Context, req *grpc.RenameInstanceAnnotationRequest) (*grpc.RenameInstanceAnnotationResponse, error) {

	flow.sugar.Debugf("Handling gRPC request: %s", this())

	tx, err := flow.db.Tx(ctx)
	if err != nil {
		return nil, err
	}
	defer rollback(tx)

	nsc := tx.Namespace
	d, err := flow.traverseToInstanceAnnotation(ctx, nsc, req.GetNamespace(), req.GetInstance(), req.GetOld())
	if err != nil {
		return nil, err
	}

	annotation, err := d.annotation.Update().SetName(req.GetNew()).Save(ctx)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	flow.logToInstance(ctx, time.Now(), d.in, "Renamed instance annotation from '%s' to '%s'.", req.GetOld(), req.GetNew())
	flow.pubsub.NotifyInstanceAnnotations(d.in)

	var resp grpc.RenameInstanceAnnotationResponse

	resp.Checksum = d.annotation.Hash
	resp.CreatedAt = timestamppb.New(d.annotation.CreatedAt)
	resp.Key = annotation.Name
	resp.Namespace = d.ns().Name
	resp.TotalSize = int64(d.annotation.Size)
	resp.UpdatedAt = timestamppb.New(d.annotation.UpdatedAt)

	return &resp, nil

}
