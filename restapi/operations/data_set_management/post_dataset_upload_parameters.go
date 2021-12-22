// Code generated by go-swagger; DO NOT EDIT.

package data_set_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewPostDatasetUploadParams creates a new PostDatasetUploadParams object
// no default values defined in spec.
func NewPostDatasetUploadParams() PostDatasetUploadParams {

	return PostDatasetUploadParams{}
}

// PostDatasetUploadParams contains all the bound params for the post dataset upload operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostDatasetUpload
type PostDatasetUploadParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Must be 0 for creation extension. May be a positive number for creation-with-upload extension.
	  In: header
	*/
	ContentLength *int64
	/*Protocol version
	  Required: true
	  In: header
	*/
	TusResumable string
	/*Added by the checksum extension. The Upload-Checksum request header contains information about the checksum of the current body payload. The header MUST consist of the name of the used checksum algorithm and the Base64 encoded checksum separated by a space.
	  In: header
	*/
	UploadChecksum *string
	/*Added by the Concatenation extension. The Upload-Concat request and response header MUST be set in both partial and final upload creation requests. It indicates whether the upload is either a partial or final upload. If the upload is a partial one, the header value MUST be partial. In the case of a final upload, its value MUST be final followed by a semicolon and a space-separated list of partial upload URLs that will be concatenated. The partial uploads URLs MAY be absolute or relative and MUST NOT contain spaces as defined in RFC 3986.
	  In: header
	*/
	UploadConcat *string
	/*Added by the creation-defer-length extension. The Upload-Defer-Length request and response header indicates that the size of the upload is not known currently and will be transferred later. Its value MUST be 1. If the length of an upload is not deferred, this header MUST be omitted.
	  In: header
	*/
	UploadDeferLength *int64
	/*The Upload-Length request and response header indicates the size of the entire upload in bytes. The value MUST be a non-negative integer. In the concatenation extension, the Client MUST NOT include the Upload-Length header in the final upload creation
	  In: header
	*/
	UploadLength *int64
	/*Added by the Creation extension. The Upload-Metadata request and response header MUST consist of one or more comma-separated key-value pairs. The key and value MUST be separated by a space. The key MUST NOT contain spaces and commas and MUST NOT be empty. The key SHOULD be ASCII encoded and the value MUST be Base64 encoded. All keys MUST be unique. The value MAY be empty. In these cases, the space, which would normally separate the key and the value, MAY be left out. Since metadata can contain arbitrary binary values, Servers SHOULD carefully validate metadata values or sanitize them before using them as header values to avoid header smuggling.
	  In: header
	*/
	UploadMetadata *string
	/*
	  In: header
	*/
	UploadOffset *int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostDatasetUploadParams() beforehand.
func (o *PostDatasetUploadParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := o.bindContentLength(r.Header[http.CanonicalHeaderKey("Content-Length")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindTusResumable(r.Header[http.CanonicalHeaderKey("Tus-Resumable")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindUploadChecksum(r.Header[http.CanonicalHeaderKey("Upload-Checksum")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindUploadConcat(r.Header[http.CanonicalHeaderKey("Upload-Concat")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindUploadDeferLength(r.Header[http.CanonicalHeaderKey("Upload-Defer-Length")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindUploadLength(r.Header[http.CanonicalHeaderKey("Upload-Length")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindUploadMetadata(r.Header[http.CanonicalHeaderKey("Upload-Metadata")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindUploadOffset(r.Header[http.CanonicalHeaderKey("Upload-offset")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindContentLength binds and validates parameter ContentLength from header.
func (o *PostDatasetUploadParams) bindContentLength(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("Content-Length", "header", "int64", raw)
	}
	o.ContentLength = &value

	return nil
}

// bindTusResumable binds and validates parameter TusResumable from header.
func (o *PostDatasetUploadParams) bindTusResumable(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("Tus-Resumable", "header", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("Tus-Resumable", "header", raw); err != nil {
		return err
	}

	o.TusResumable = raw

	if err := o.validateTusResumable(formats); err != nil {
		return err
	}

	return nil
}

// validateTusResumable carries on validations for parameter TusResumable
func (o *PostDatasetUploadParams) validateTusResumable(formats strfmt.Registry) error {

	if err := validate.EnumCase("Tus-Resumable", "header", o.TusResumable, []interface{}{"1.0.0"}, true); err != nil {
		return err
	}

	return nil
}

// bindUploadChecksum binds and validates parameter UploadChecksum from header.
func (o *PostDatasetUploadParams) bindUploadChecksum(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.UploadChecksum = &raw

	return nil
}

// bindUploadConcat binds and validates parameter UploadConcat from header.
func (o *PostDatasetUploadParams) bindUploadConcat(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.UploadConcat = &raw

	return nil
}

// bindUploadDeferLength binds and validates parameter UploadDeferLength from header.
func (o *PostDatasetUploadParams) bindUploadDeferLength(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("Upload-Defer-Length", "header", "int64", raw)
	}
	o.UploadDeferLength = &value

	if err := o.validateUploadDeferLength(formats); err != nil {
		return err
	}

	return nil
}

// validateUploadDeferLength carries on validations for parameter UploadDeferLength
func (o *PostDatasetUploadParams) validateUploadDeferLength(formats strfmt.Registry) error {

	if err := validate.EnumCase("Upload-Defer-Length", "header", *o.UploadDeferLength, []interface{}{1}, true); err != nil {
		return err
	}

	return nil
}

// bindUploadLength binds and validates parameter UploadLength from header.
func (o *PostDatasetUploadParams) bindUploadLength(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("Upload-Length", "header", "int64", raw)
	}
	o.UploadLength = &value

	return nil
}

// bindUploadMetadata binds and validates parameter UploadMetadata from header.
func (o *PostDatasetUploadParams) bindUploadMetadata(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.UploadMetadata = &raw

	return nil
}

// bindUploadOffset binds and validates parameter UploadOffset from header.
func (o *PostDatasetUploadParams) bindUploadOffset(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("Upload-offset", "header", "int64", raw)
	}
	o.UploadOffset = &value

	return nil
}