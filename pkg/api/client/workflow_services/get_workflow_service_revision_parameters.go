// Code generated by go-swagger; DO NOT EDIT.

package workflow_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetWorkflowServiceRevisionParams creates a new GetWorkflowServiceRevisionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWorkflowServiceRevisionParams() *GetWorkflowServiceRevisionParams {
	return &GetWorkflowServiceRevisionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkflowServiceRevisionParamsWithTimeout creates a new GetWorkflowServiceRevisionParams object
// with the ability to set a timeout on a request.
func NewGetWorkflowServiceRevisionParamsWithTimeout(timeout time.Duration) *GetWorkflowServiceRevisionParams {
	return &GetWorkflowServiceRevisionParams{
		timeout: timeout,
	}
}

// NewGetWorkflowServiceRevisionParamsWithContext creates a new GetWorkflowServiceRevisionParams object
// with the ability to set a context for a request.
func NewGetWorkflowServiceRevisionParamsWithContext(ctx context.Context) *GetWorkflowServiceRevisionParams {
	return &GetWorkflowServiceRevisionParams{
		Context: ctx,
	}
}

// NewGetWorkflowServiceRevisionParamsWithHTTPClient creates a new GetWorkflowServiceRevisionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWorkflowServiceRevisionParamsWithHTTPClient(client *http.Client) *GetWorkflowServiceRevisionParams {
	return &GetWorkflowServiceRevisionParams{
		HTTPClient: client,
	}
}

/* GetWorkflowServiceRevisionParams contains all the parameters to send to the API endpoint
   for the get workflow service revision operation.

   Typically these are written to a http.Request.
*/
type GetWorkflowServiceRevisionParams struct {

	/* Namespace.

	   target namespace
	*/
	Namespace string

	/* Rev.

	   target service revison
	*/
	Rev string

	/* Svn.

	   target service name
	*/
	Svn string

	/* Workflow.

	   path to target workflow
	*/
	Workflow string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get workflow service revision params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWorkflowServiceRevisionParams) WithDefaults() *GetWorkflowServiceRevisionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get workflow service revision params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWorkflowServiceRevisionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get workflow service revision params
func (o *GetWorkflowServiceRevisionParams) WithTimeout(timeout time.Duration) *GetWorkflowServiceRevisionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workflow service revision params
func (o *GetWorkflowServiceRevisionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workflow service revision params
func (o *GetWorkflowServiceRevisionParams) WithContext(ctx context.Context) *GetWorkflowServiceRevisionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workflow service revision params
func (o *GetWorkflowServiceRevisionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workflow service revision params
func (o *GetWorkflowServiceRevisionParams) WithHTTPClient(client *http.Client) *GetWorkflowServiceRevisionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workflow service revision params
func (o *GetWorkflowServiceRevisionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the get workflow service revision params
func (o *GetWorkflowServiceRevisionParams) WithNamespace(namespace string) *GetWorkflowServiceRevisionParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get workflow service revision params
func (o *GetWorkflowServiceRevisionParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithRev adds the rev to the get workflow service revision params
func (o *GetWorkflowServiceRevisionParams) WithRev(rev string) *GetWorkflowServiceRevisionParams {
	o.SetRev(rev)
	return o
}

// SetRev adds the rev to the get workflow service revision params
func (o *GetWorkflowServiceRevisionParams) SetRev(rev string) {
	o.Rev = rev
}

// WithSvn adds the svn to the get workflow service revision params
func (o *GetWorkflowServiceRevisionParams) WithSvn(svn string) *GetWorkflowServiceRevisionParams {
	o.SetSvn(svn)
	return o
}

// SetSvn adds the svn to the get workflow service revision params
func (o *GetWorkflowServiceRevisionParams) SetSvn(svn string) {
	o.Svn = svn
}

// WithWorkflow adds the workflow to the get workflow service revision params
func (o *GetWorkflowServiceRevisionParams) WithWorkflow(workflow string) *GetWorkflowServiceRevisionParams {
	o.SetWorkflow(workflow)
	return o
}

// SetWorkflow adds the workflow to the get workflow service revision params
func (o *GetWorkflowServiceRevisionParams) SetWorkflow(workflow string) {
	o.Workflow = workflow
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkflowServiceRevisionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// query param rev
	qrRev := o.Rev
	qRev := qrRev
	if qRev != "" {

		if err := r.SetQueryParam("rev", qRev); err != nil {
			return err
		}
	}

	// query param svn
	qrSvn := o.Svn
	qSvn := qrSvn
	if qSvn != "" {

		if err := r.SetQueryParam("svn", qSvn); err != nil {
			return err
		}
	}

	// path param workflow
	if err := r.SetPathParam("workflow", o.Workflow); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}