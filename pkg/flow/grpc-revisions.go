package flow

import (
	"context"
	"errors"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/direktiv/direktiv/pkg/flow/ent"
	entref "github.com/direktiv/direktiv/pkg/flow/ent/ref"
	entrev "github.com/direktiv/direktiv/pkg/flow/ent/revision"
	entwf "github.com/direktiv/direktiv/pkg/flow/ent/workflow"
	"github.com/direktiv/direktiv/pkg/flow/grpc"
)

func (flow *flow) Revisions(ctx context.Context, req *grpc.RevisionsRequest) (*grpc.RevisionsResponse, error) {

	flow.sugar.Debugf("Handling gRPC request: %s", this())

	cached, err := flow.traverseToWorkflow(ctx, nil, req.GetNamespace(), req.GetPath())
	if err != nil {
		return nil, err
	}

	clients := flow.edb.Clients(nil)

	query := clients.Ref.Query().Where(entref.HasWorkflowWith(entwf.ID(cached.Workflow.ID)), entref.Immutable(true))

	results, pi, err := paginate[*ent.RefQuery, *ent.Ref](ctx, req.Pagination, query, refsOrderings, refsFilters)
	if err != nil {
		return nil, err
	}

	resp := new(grpc.RevisionsResponse)
	resp.Namespace = cached.Namespace.Name
	resp.PageInfo = pi

	err = atob(results, &resp.Results)
	if err != nil {
		return nil, err
	}

	err = atob(cached.Inode(), &resp.Node)
	if err != nil {
		return nil, err
	}

	resp.Node.Path = cached.Path()
	resp.Node.Parent = cached.Dir()

	return resp, nil

}

func (flow *flow) RevisionsStream(req *grpc.RevisionsRequest, srv grpc.Flow_RevisionsStreamServer) error {

	flow.sugar.Debugf("Handling gRPC request: %s", this())

	ctx := srv.Context()
	phash := ""
	nhash := ""

	cached, err := flow.traverseToWorkflow(ctx, nil, req.GetNamespace(), req.GetPath())
	if err != nil {
		return err
	}

	sub := flow.pubsub.SubscribeWorkflow(cached)
	defer flow.cleanup(sub.Close)

resend:

	clients := flow.edb.Clients(nil)

	query := clients.Ref.Query().Where(entref.HasWorkflowWith(entwf.ID(cached.Workflow.ID)), entref.Immutable(true))

	results, pi, err := paginate[*ent.RefQuery, *ent.Ref](ctx, req.Pagination, query, refsOrderings, refsFilters)
	if err != nil {
		return err
	}

	resp := new(grpc.RevisionsResponse)
	resp.Namespace = cached.Namespace.Name
	resp.PageInfo = pi

	err = atob(results, &resp.Results)
	if err != nil {
		return err
	}

	err = atob(cached.Inode(), &resp.Node)
	if err != nil {
		return err
	}

	resp.Node.Path = cached.Path()
	resp.Node.Parent = cached.Dir()

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

func (flow *flow) DeleteRevision(ctx context.Context, req *grpc.DeleteRevisionRequest) (*emptypb.Empty, error) {

	flow.sugar.Debugf("Handling gRPC request: %s", this())

	tx, err := flow.database.Tx(ctx)
	if err != nil {
		return nil, err
	}
	defer rollback(tx)

	cached, err := flow.traverseToRef(ctx, tx, req.GetNamespace(), req.GetPath(), req.GetRevision())
	if err != nil {
		return nil, err
	}

	if !cached.Ref.Immutable {
		return nil, errors.New("not a revision")
	}

	clients := flow.edb.Clients(nil)

	query := clients.Ref.Query().Where(entref.HasRevisionWith(entrev.ID(cached.Revision.ID)), entref.Immutable(false))

	xrefs, err := query.All(ctx)
	if err != nil {
		return nil, err
	}

	if len(xrefs) > 1 || (len(xrefs) == 1 && xrefs[0].Name != "latest") {
		return nil, errors.New("cannot delete revision while refs to it exist")
	}

	if len(xrefs) == 1 && xrefs[0].Name == "latest" {
		err = flow.configureRouter(ctx, tx, cached, rcfBreaking,
			func() error {

				err := clients.Ref.DeleteOneID(cached.Ref.ID).Exec(ctx)
				if err != nil {
					return err
				}

				return nil

			},
			tx.Commit,
		)
		if err != nil {
			return nil, err
		}
	} else {
		err = flow.configureRouter(ctx, tx, cached, rcfBreaking,
			func() error {

				err := clients.Revision.DeleteOneID(cached.Revision.ID).Exec(ctx)
				if err != nil {
					return err
				}

				return nil

			},
			tx.Commit,
		)
		if err != nil {
			return nil, err
		}
	}

	flow.logToWorkflow(ctx, time.Now(), cached, "Deleted workflow revision: %s.", cached.Revision.ID.String())
	flow.pubsub.NotifyWorkflow(cached.Workflow)

	var resp emptypb.Empty

	return &resp, nil

}
