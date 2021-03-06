// Code generated by go-swagger; DO NOT EDIT.

package user_org_management

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

// NewDeleteHPCResourceParams creates a new DeleteHPCResourceParams object
// with the default values initialized.
func NewDeleteHPCResourceParams() *DeleteHPCResourceParams {
	var ()
	return &DeleteHPCResourceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteHPCResourceParamsWithTimeout creates a new DeleteHPCResourceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteHPCResourceParamsWithTimeout(timeout time.Duration) *DeleteHPCResourceParams {
	var ()
	return &DeleteHPCResourceParams{

		timeout: timeout,
	}
}

// NewDeleteHPCResourceParamsWithContext creates a new DeleteHPCResourceParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteHPCResourceParamsWithContext(ctx context.Context) *DeleteHPCResourceParams {
	var ()
	return &DeleteHPCResourceParams{

		Context: ctx,
	}
}

// NewDeleteHPCResourceParamsWithHTTPClient creates a new DeleteHPCResourceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteHPCResourceParamsWithHTTPClient(client *http.Client) *DeleteHPCResourceParams {
	var ()
	return &DeleteHPCResourceParams{
		HTTPClient: client,
	}
}

/*DeleteHPCResourceParams contains all the parameters to send to the API endpoint
for the delete h p c resource operation typically these are written to a http.Request
*/
type DeleteHPCResourceParams struct {

	/*ID
	  Id of HPCResource to be obtained

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete h p c resource params
func (o *DeleteHPCResourceParams) WithTimeout(timeout time.Duration) *DeleteHPCResourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete h p c resource params
func (o *DeleteHPCResourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete h p c resource params
func (o *DeleteHPCResourceParams) WithContext(ctx context.Context) *DeleteHPCResourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete h p c resource params
func (o *DeleteHPCResourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete h p c resource params
func (o *DeleteHPCResourceParams) WithHTTPClient(client *http.Client) *DeleteHPCResourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete h p c resource params
func (o *DeleteHPCResourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete h p c resource params
func (o *DeleteHPCResourceParams) WithID(id string) *DeleteHPCResourceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete h p c resource params
func (o *DeleteHPCResourceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteHPCResourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
