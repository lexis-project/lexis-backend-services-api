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

// NewCertificateParams creates a new CertificateParams object
// with the default values initialized.
func NewCertificateParams() *CertificateParams {

	return &CertificateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCertificateParamsWithTimeout creates a new CertificateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCertificateParamsWithTimeout(timeout time.Duration) *CertificateParams {

	return &CertificateParams{

		timeout: timeout,
	}
}

// NewCertificateParamsWithContext creates a new CertificateParams object
// with the default values initialized, and the ability to set a context for a request
func NewCertificateParamsWithContext(ctx context.Context) *CertificateParams {

	return &CertificateParams{

		Context: ctx,
	}
}

// NewCertificateParamsWithHTTPClient creates a new CertificateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCertificateParamsWithHTTPClient(client *http.Client) *CertificateParams {

	return &CertificateParams{
		HTTPClient: client,
	}
}

/*CertificateParams contains all the parameters to send to the API endpoint
for the certificate operation typically these are written to a http.Request
*/
type CertificateParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the certificate params
func (o *CertificateParams) WithTimeout(timeout time.Duration) *CertificateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the certificate params
func (o *CertificateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the certificate params
func (o *CertificateParams) WithContext(ctx context.Context) *CertificateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the certificate params
func (o *CertificateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the certificate params
func (o *CertificateParams) WithHTTPClient(client *http.Client) *CertificateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the certificate params
func (o *CertificateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CertificateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}