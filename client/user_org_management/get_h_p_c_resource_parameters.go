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

// NewGetHPCResourceParams creates a new GetHPCResourceParams object
// with the default values initialized.
func NewGetHPCResourceParams() *GetHPCResourceParams {
	var ()
	return &GetHPCResourceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetHPCResourceParamsWithTimeout creates a new GetHPCResourceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetHPCResourceParamsWithTimeout(timeout time.Duration) *GetHPCResourceParams {
	var ()
	return &GetHPCResourceParams{

		timeout: timeout,
	}
}

// NewGetHPCResourceParamsWithContext creates a new GetHPCResourceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetHPCResourceParamsWithContext(ctx context.Context) *GetHPCResourceParams {
	var ()
	return &GetHPCResourceParams{

		Context: ctx,
	}
}

// NewGetHPCResourceParamsWithHTTPClient creates a new GetHPCResourceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetHPCResourceParamsWithHTTPClient(client *http.Client) *GetHPCResourceParams {
	var ()
	return &GetHPCResourceParams{
		HTTPClient: client,
	}
}

/*GetHPCResourceParams contains all the parameters to send to the API endpoint
for the get h p c resource operation typically these are written to a http.Request
*/
type GetHPCResourceParams struct {

	/*ID
	  Id of HPCResource to be obtained

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get h p c resource params
func (o *GetHPCResourceParams) WithTimeout(timeout time.Duration) *GetHPCResourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get h p c resource params
func (o *GetHPCResourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get h p c resource params
func (o *GetHPCResourceParams) WithContext(ctx context.Context) *GetHPCResourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get h p c resource params
func (o *GetHPCResourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get h p c resource params
func (o *GetHPCResourceParams) WithHTTPClient(client *http.Client) *GetHPCResourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get h p c resource params
func (o *GetHPCResourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get h p c resource params
func (o *GetHPCResourceParams) WithID(id string) *GetHPCResourceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get h p c resource params
func (o *GetHPCResourceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetHPCResourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
