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

// NewListOrganizationsParams creates a new ListOrganizationsParams object
// with the default values initialized.
func NewListOrganizationsParams() *ListOrganizationsParams {
	var (
		scopeDefault = string("OWN")
	)
	return &ListOrganizationsParams{
		Scope: &scopeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewListOrganizationsParamsWithTimeout creates a new ListOrganizationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListOrganizationsParamsWithTimeout(timeout time.Duration) *ListOrganizationsParams {
	var (
		scopeDefault = string("OWN")
	)
	return &ListOrganizationsParams{
		Scope: &scopeDefault,

		timeout: timeout,
	}
}

// NewListOrganizationsParamsWithContext creates a new ListOrganizationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListOrganizationsParamsWithContext(ctx context.Context) *ListOrganizationsParams {
	var (
		scopeDefault = string("OWN")
	)
	return &ListOrganizationsParams{
		Scope: &scopeDefault,

		Context: ctx,
	}
}

// NewListOrganizationsParamsWithHTTPClient creates a new ListOrganizationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListOrganizationsParamsWithHTTPClient(client *http.Client) *ListOrganizationsParams {
	var (
		scopeDefault = string("OWN")
	)
	return &ListOrganizationsParams{
		Scope:      &scopeDefault,
		HTTPClient: client,
	}
}

/*ListOrganizationsParams contains all the parameters to send to the API endpoint
for the list organizations operation typically these are written to a http.Request
*/
type ListOrganizationsParams struct {

	/*Scope
	  organization scope switch

	*/
	Scope *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list organizations params
func (o *ListOrganizationsParams) WithTimeout(timeout time.Duration) *ListOrganizationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list organizations params
func (o *ListOrganizationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list organizations params
func (o *ListOrganizationsParams) WithContext(ctx context.Context) *ListOrganizationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list organizations params
func (o *ListOrganizationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list organizations params
func (o *ListOrganizationsParams) WithHTTPClient(client *http.Client) *ListOrganizationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list organizations params
func (o *ListOrganizationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithScope adds the scope to the list organizations params
func (o *ListOrganizationsParams) WithScope(scope *string) *ListOrganizationsParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the list organizations params
func (o *ListOrganizationsParams) SetScope(scope *string) {
	o.Scope = scope
}

// WriteToRequest writes these params to a swagger request
func (o *ListOrganizationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Scope != nil {

		// query param scope
		var qrScope string
		if o.Scope != nil {
			qrScope = *o.Scope
		}
		qScope := qrScope
		if qScope != "" {
			if err := r.SetQueryParam("scope", qScope); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
