// Code generated by go-swagger; DO NOT EDIT.

package workflow_management

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

// NewGetWorkflowsParams creates a new GetWorkflowsParams object
// with the default values initialized.
func NewGetWorkflowsParams() *GetWorkflowsParams {

	return &GetWorkflowsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkflowsParamsWithTimeout creates a new GetWorkflowsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWorkflowsParamsWithTimeout(timeout time.Duration) *GetWorkflowsParams {

	return &GetWorkflowsParams{

		timeout: timeout,
	}
}

// NewGetWorkflowsParamsWithContext creates a new GetWorkflowsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWorkflowsParamsWithContext(ctx context.Context) *GetWorkflowsParams {

	return &GetWorkflowsParams{

		Context: ctx,
	}
}

// NewGetWorkflowsParamsWithHTTPClient creates a new GetWorkflowsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWorkflowsParamsWithHTTPClient(client *http.Client) *GetWorkflowsParams {

	return &GetWorkflowsParams{
		HTTPClient: client,
	}
}

/*GetWorkflowsParams contains all the parameters to send to the API endpoint
for the get workflows operation typically these are written to a http.Request
*/
type GetWorkflowsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get workflows params
func (o *GetWorkflowsParams) WithTimeout(timeout time.Duration) *GetWorkflowsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workflows params
func (o *GetWorkflowsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workflows params
func (o *GetWorkflowsParams) WithContext(ctx context.Context) *GetWorkflowsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workflows params
func (o *GetWorkflowsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workflows params
func (o *GetWorkflowsParams) WithHTTPClient(client *http.Client) *GetWorkflowsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workflows params
func (o *GetWorkflowsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkflowsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
