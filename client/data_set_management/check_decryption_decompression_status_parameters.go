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

// NewCheckDecryptionDecompressionStatusParams creates a new CheckDecryptionDecompressionStatusParams object
// with the default values initialized.
func NewCheckDecryptionDecompressionStatusParams() *CheckDecryptionDecompressionStatusParams {
	var ()
	return &CheckDecryptionDecompressionStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCheckDecryptionDecompressionStatusParamsWithTimeout creates a new CheckDecryptionDecompressionStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCheckDecryptionDecompressionStatusParamsWithTimeout(timeout time.Duration) *CheckDecryptionDecompressionStatusParams {
	var ()
	return &CheckDecryptionDecompressionStatusParams{

		timeout: timeout,
	}
}

// NewCheckDecryptionDecompressionStatusParamsWithContext creates a new CheckDecryptionDecompressionStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewCheckDecryptionDecompressionStatusParamsWithContext(ctx context.Context) *CheckDecryptionDecompressionStatusParams {
	var ()
	return &CheckDecryptionDecompressionStatusParams{

		Context: ctx,
	}
}

// NewCheckDecryptionDecompressionStatusParamsWithHTTPClient creates a new CheckDecryptionDecompressionStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCheckDecryptionDecompressionStatusParamsWithHTTPClient(client *http.Client) *CheckDecryptionDecompressionStatusParams {
	var ()
	return &CheckDecryptionDecompressionStatusParams{
		HTTPClient: client,
	}
}

/*CheckDecryptionDecompressionStatusParams contains all the parameters to send to the API endpoint
for the check decryption decompression status operation typically these are written to a http.Request
*/
type CheckDecryptionDecompressionStatusParams struct {

	/*RequestID*/
	RequestID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the check decryption decompression status params
func (o *CheckDecryptionDecompressionStatusParams) WithTimeout(timeout time.Duration) *CheckDecryptionDecompressionStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the check decryption decompression status params
func (o *CheckDecryptionDecompressionStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the check decryption decompression status params
func (o *CheckDecryptionDecompressionStatusParams) WithContext(ctx context.Context) *CheckDecryptionDecompressionStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the check decryption decompression status params
func (o *CheckDecryptionDecompressionStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the check decryption decompression status params
func (o *CheckDecryptionDecompressionStatusParams) WithHTTPClient(client *http.Client) *CheckDecryptionDecompressionStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the check decryption decompression status params
func (o *CheckDecryptionDecompressionStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestID adds the requestID to the check decryption decompression status params
func (o *CheckDecryptionDecompressionStatusParams) WithRequestID(requestID strfmt.UUID) *CheckDecryptionDecompressionStatusParams {
	o.SetRequestID(requestID)
	return o
}

// SetRequestID adds the requestId to the check decryption decompression status params
func (o *CheckDecryptionDecompressionStatusParams) SetRequestID(requestID strfmt.UUID) {
	o.RequestID = requestID
}

// WriteToRequest writes these params to a swagger request
func (o *CheckDecryptionDecompressionStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param request_id
	if err := r.SetPathParam("request_id", o.RequestID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}