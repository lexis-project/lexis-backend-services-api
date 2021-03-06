// Code generated by go-swagger; DO NOT EDIT.

package staging

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

// NewCloudNFSExportRemoveParams creates a new CloudNFSExportRemoveParams object
// with the default values initialized.
func NewCloudNFSExportRemoveParams() *CloudNFSExportRemoveParams {
	var ()
	return &CloudNFSExportRemoveParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCloudNFSExportRemoveParamsWithTimeout creates a new CloudNFSExportRemoveParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCloudNFSExportRemoveParamsWithTimeout(timeout time.Duration) *CloudNFSExportRemoveParams {
	var ()
	return &CloudNFSExportRemoveParams{

		timeout: timeout,
	}
}

// NewCloudNFSExportRemoveParamsWithContext creates a new CloudNFSExportRemoveParams object
// with the default values initialized, and the ability to set a context for a request
func NewCloudNFSExportRemoveParamsWithContext(ctx context.Context) *CloudNFSExportRemoveParams {
	var ()
	return &CloudNFSExportRemoveParams{

		Context: ctx,
	}
}

// NewCloudNFSExportRemoveParamsWithHTTPClient creates a new CloudNFSExportRemoveParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCloudNFSExportRemoveParamsWithHTTPClient(client *http.Client) *CloudNFSExportRemoveParams {
	var ()
	return &CloudNFSExportRemoveParams{
		HTTPClient: client,
	}
}

/*CloudNFSExportRemoveParams contains all the parameters to send to the API endpoint
for the cloud n f s export remove operation typically these are written to a http.Request
*/
type CloudNFSExportRemoveParams struct {

	/*Param
	  IP of the cloud machine

	*/
	Param string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cloud n f s export remove params
func (o *CloudNFSExportRemoveParams) WithTimeout(timeout time.Duration) *CloudNFSExportRemoveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cloud n f s export remove params
func (o *CloudNFSExportRemoveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cloud n f s export remove params
func (o *CloudNFSExportRemoveParams) WithContext(ctx context.Context) *CloudNFSExportRemoveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cloud n f s export remove params
func (o *CloudNFSExportRemoveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cloud n f s export remove params
func (o *CloudNFSExportRemoveParams) WithHTTPClient(client *http.Client) *CloudNFSExportRemoveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cloud n f s export remove params
func (o *CloudNFSExportRemoveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithParam adds the param to the cloud n f s export remove params
func (o *CloudNFSExportRemoveParams) WithParam(param string) *CloudNFSExportRemoveParams {
	o.SetParam(param)
	return o
}

// SetParam adds the param to the cloud n f s export remove params
func (o *CloudNFSExportRemoveParams) SetParam(param string) {
	o.Param = param
}

// WriteToRequest writes these params to a swagger request
func (o *CloudNFSExportRemoveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param param
	if err := r.SetPathParam("param", o.Param); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
