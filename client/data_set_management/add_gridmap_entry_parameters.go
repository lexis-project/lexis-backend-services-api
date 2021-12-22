// Code generated by go-swagger; DO NOT EDIT.

package data_set_management

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

// NewAddGridmapEntryParams creates a new AddGridmapEntryParams object
// with the default values initialized.
func NewAddGridmapEntryParams() *AddGridmapEntryParams {
	var ()
	return &AddGridmapEntryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddGridmapEntryParamsWithTimeout creates a new AddGridmapEntryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddGridmapEntryParamsWithTimeout(timeout time.Duration) *AddGridmapEntryParams {
	var ()
	return &AddGridmapEntryParams{

		timeout: timeout,
	}
}

// NewAddGridmapEntryParamsWithContext creates a new AddGridmapEntryParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddGridmapEntryParamsWithContext(ctx context.Context) *AddGridmapEntryParams {
	var ()
	return &AddGridmapEntryParams{

		Context: ctx,
	}
}

// NewAddGridmapEntryParamsWithHTTPClient creates a new AddGridmapEntryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddGridmapEntryParamsWithHTTPClient(client *http.Client) *AddGridmapEntryParams {
	var ()
	return &AddGridmapEntryParams{
		HTTPClient: client,
	}
}

/*AddGridmapEntryParams contains all the parameters to send to the API endpoint
for the add gridmap entry operation typically these are written to a http.Request
*/
type AddGridmapEntryParams struct {

	/*Parameters
	  parameters

	*/
	Parameters AddGridmapEntryBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add gridmap entry params
func (o *AddGridmapEntryParams) WithTimeout(timeout time.Duration) *AddGridmapEntryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add gridmap entry params
func (o *AddGridmapEntryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add gridmap entry params
func (o *AddGridmapEntryParams) WithContext(ctx context.Context) *AddGridmapEntryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add gridmap entry params
func (o *AddGridmapEntryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add gridmap entry params
func (o *AddGridmapEntryParams) WithHTTPClient(client *http.Client) *AddGridmapEntryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add gridmap entry params
func (o *AddGridmapEntryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithParameters adds the parameters to the add gridmap entry params
func (o *AddGridmapEntryParams) WithParameters(parameters AddGridmapEntryBody) *AddGridmapEntryParams {
	o.SetParameters(parameters)
	return o
}

// SetParameters adds the parameters to the add gridmap entry params
func (o *AddGridmapEntryParams) SetParameters(parameters AddGridmapEntryBody) {
	o.Parameters = parameters
}

// WriteToRequest writes these params to a swagger request
func (o *AddGridmapEntryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Parameters); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
