// Code generated by go-swagger; DO NOT EDIT.

package approval_system_management

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

// NewListProjectHPCResourceRequestParams creates a new ListProjectHPCResourceRequestParams object
// with the default values initialized.
func NewListProjectHPCResourceRequestParams() *ListProjectHPCResourceRequestParams {
	var ()
	return &ListProjectHPCResourceRequestParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListProjectHPCResourceRequestParamsWithTimeout creates a new ListProjectHPCResourceRequestParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListProjectHPCResourceRequestParamsWithTimeout(timeout time.Duration) *ListProjectHPCResourceRequestParams {
	var ()
	return &ListProjectHPCResourceRequestParams{

		timeout: timeout,
	}
}

// NewListProjectHPCResourceRequestParamsWithContext creates a new ListProjectHPCResourceRequestParams object
// with the default values initialized, and the ability to set a context for a request
func NewListProjectHPCResourceRequestParamsWithContext(ctx context.Context) *ListProjectHPCResourceRequestParams {
	var ()
	return &ListProjectHPCResourceRequestParams{

		Context: ctx,
	}
}

// NewListProjectHPCResourceRequestParamsWithHTTPClient creates a new ListProjectHPCResourceRequestParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListProjectHPCResourceRequestParamsWithHTTPClient(client *http.Client) *ListProjectHPCResourceRequestParams {
	var ()
	return &ListProjectHPCResourceRequestParams{
		HTTPClient: client,
	}
}

/*ListProjectHPCResourceRequestParams contains all the parameters to send to the API endpoint
for the list project h p c resource request operation typically these are written to a http.Request
*/
type ListProjectHPCResourceRequestParams struct {

	/*AssociatedLEXISProject
	  LEXIS project ID

	*/
	AssociatedLEXISProject string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list project h p c resource request params
func (o *ListProjectHPCResourceRequestParams) WithTimeout(timeout time.Duration) *ListProjectHPCResourceRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list project h p c resource request params
func (o *ListProjectHPCResourceRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list project h p c resource request params
func (o *ListProjectHPCResourceRequestParams) WithContext(ctx context.Context) *ListProjectHPCResourceRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list project h p c resource request params
func (o *ListProjectHPCResourceRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list project h p c resource request params
func (o *ListProjectHPCResourceRequestParams) WithHTTPClient(client *http.Client) *ListProjectHPCResourceRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list project h p c resource request params
func (o *ListProjectHPCResourceRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAssociatedLEXISProject adds the associatedLEXISProject to the list project h p c resource request params
func (o *ListProjectHPCResourceRequestParams) WithAssociatedLEXISProject(associatedLEXISProject string) *ListProjectHPCResourceRequestParams {
	o.SetAssociatedLEXISProject(associatedLEXISProject)
	return o
}

// SetAssociatedLEXISProject adds the associatedLEXISProject to the list project h p c resource request params
func (o *ListProjectHPCResourceRequestParams) SetAssociatedLEXISProject(associatedLEXISProject string) {
	o.AssociatedLEXISProject = associatedLEXISProject
}

// WriteToRequest writes these params to a swagger request
func (o *ListProjectHPCResourceRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param AssociatedLEXISProject
	if err := r.SetPathParam("AssociatedLEXISProject", o.AssociatedLEXISProject); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
