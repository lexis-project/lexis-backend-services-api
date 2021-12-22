// Code generated by go-swagger; DO NOT EDIT.

package data_set_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/lexis-project/lexis-backend-services-api.git/models"
)

// FilesDeleteNoContentCode is the HTTP code returned for type FilesDeleteNoContent
const FilesDeleteNoContentCode int = 204

/*FilesDeleteNoContent Upload was terminated

swagger:response filesDeleteNoContent
*/
type FilesDeleteNoContent struct {
	/*Needed to make browsers accept the additional headers used by
	the tus protocol.


	*/
	AccessControlExposeHeaders string `json:"Access-Control-Expose-Headers"`
	/*Protocol version

	 */
	TusResumable string `json:"Tus-Resumable"`
}

// NewFilesDeleteNoContent creates FilesDeleteNoContent with default headers values
func NewFilesDeleteNoContent() *FilesDeleteNoContent {

	return &FilesDeleteNoContent{}
}

// WithAccessControlExposeHeaders adds the accessControlExposeHeaders to the files delete no content response
func (o *FilesDeleteNoContent) WithAccessControlExposeHeaders(accessControlExposeHeaders string) *FilesDeleteNoContent {
	o.AccessControlExposeHeaders = accessControlExposeHeaders
	return o
}

// SetAccessControlExposeHeaders sets the accessControlExposeHeaders to the files delete no content response
func (o *FilesDeleteNoContent) SetAccessControlExposeHeaders(accessControlExposeHeaders string) {
	o.AccessControlExposeHeaders = accessControlExposeHeaders
}

// WithTusResumable adds the tusResumable to the files delete no content response
func (o *FilesDeleteNoContent) WithTusResumable(tusResumable string) *FilesDeleteNoContent {
	o.TusResumable = tusResumable
	return o
}

// SetTusResumable sets the tusResumable to the files delete no content response
func (o *FilesDeleteNoContent) SetTusResumable(tusResumable string) {
	o.TusResumable = tusResumable
}

// WriteResponse to the client
func (o *FilesDeleteNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Access-Control-Expose-Headers

	accessControlExposeHeaders := o.AccessControlExposeHeaders
	if accessControlExposeHeaders != "" {
		rw.Header().Set("Access-Control-Expose-Headers", accessControlExposeHeaders)
	}

	// response header Tus-Resumable

	tusResumable := o.TusResumable
	if tusResumable != "" {
		rw.Header().Set("Tus-Resumable", tusResumable)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// FilesDeleteUnauthorizedCode is the HTTP code returned for type FilesDeleteUnauthorized
const FilesDeleteUnauthorizedCode int = 401

/*FilesDeleteUnauthorized Authorization failed

swagger:response filesDeleteUnauthorized
*/
type FilesDeleteUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewFilesDeleteUnauthorized creates FilesDeleteUnauthorized with default headers values
func NewFilesDeleteUnauthorized() *FilesDeleteUnauthorized {

	return &FilesDeleteUnauthorized{}
}

// WithPayload adds the payload to the files delete unauthorized response
func (o *FilesDeleteUnauthorized) WithPayload(payload *models.ErrorResponse) *FilesDeleteUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the files delete unauthorized response
func (o *FilesDeleteUnauthorized) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FilesDeleteUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// FilesDeletePreconditionFailedCode is the HTTP code returned for type FilesDeletePreconditionFailed
const FilesDeletePreconditionFailedCode int = 412

/*FilesDeletePreconditionFailed Precondition Failed

swagger:response filesDeletePreconditionFailed
*/
type FilesDeletePreconditionFailed struct {
	/*Needed to make browsers accept the additional headers used by
	the tus protocol.


	*/
	AccessControlExposeHeaders string `json:"Access-Control-Expose-Headers"`
	/*Protocol version

	 */
	TusResumable string `json:"Tus-Resumable"`
	/*The Tus-Version response header MUST be a comma-separated list of protocol versions supported by the Server. The list MUST be sorted by Server's preference where the first one is the most preferred one.

	 */
	TusVersion string `json:"Tus-Version"`
}

// NewFilesDeletePreconditionFailed creates FilesDeletePreconditionFailed with default headers values
func NewFilesDeletePreconditionFailed() *FilesDeletePreconditionFailed {

	return &FilesDeletePreconditionFailed{}
}

// WithAccessControlExposeHeaders adds the accessControlExposeHeaders to the files delete precondition failed response
func (o *FilesDeletePreconditionFailed) WithAccessControlExposeHeaders(accessControlExposeHeaders string) *FilesDeletePreconditionFailed {
	o.AccessControlExposeHeaders = accessControlExposeHeaders
	return o
}

// SetAccessControlExposeHeaders sets the accessControlExposeHeaders to the files delete precondition failed response
func (o *FilesDeletePreconditionFailed) SetAccessControlExposeHeaders(accessControlExposeHeaders string) {
	o.AccessControlExposeHeaders = accessControlExposeHeaders
}

// WithTusResumable adds the tusResumable to the files delete precondition failed response
func (o *FilesDeletePreconditionFailed) WithTusResumable(tusResumable string) *FilesDeletePreconditionFailed {
	o.TusResumable = tusResumable
	return o
}

// SetTusResumable sets the tusResumable to the files delete precondition failed response
func (o *FilesDeletePreconditionFailed) SetTusResumable(tusResumable string) {
	o.TusResumable = tusResumable
}

// WithTusVersion adds the tusVersion to the files delete precondition failed response
func (o *FilesDeletePreconditionFailed) WithTusVersion(tusVersion string) *FilesDeletePreconditionFailed {
	o.TusVersion = tusVersion
	return o
}

// SetTusVersion sets the tusVersion to the files delete precondition failed response
func (o *FilesDeletePreconditionFailed) SetTusVersion(tusVersion string) {
	o.TusVersion = tusVersion
}

// WriteResponse to the client
func (o *FilesDeletePreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Access-Control-Expose-Headers

	accessControlExposeHeaders := o.AccessControlExposeHeaders
	if accessControlExposeHeaders != "" {
		rw.Header().Set("Access-Control-Expose-Headers", accessControlExposeHeaders)
	}

	// response header Tus-Resumable

	tusResumable := o.TusResumable
	if tusResumable != "" {
		rw.Header().Set("Tus-Resumable", tusResumable)
	}

	// response header Tus-Version

	tusVersion := o.TusVersion
	if tusVersion != "" {
		rw.Header().Set("Tus-Version", tusVersion)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(412)
}

// FilesDeleteServiceUnavailableCode is the HTTP code returned for type FilesDeleteServiceUnavailable
const FilesDeleteServiceUnavailableCode int = 503

/*FilesDeleteServiceUnavailable Service Unavailable

swagger:response filesDeleteServiceUnavailable
*/
type FilesDeleteServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewFilesDeleteServiceUnavailable creates FilesDeleteServiceUnavailable with default headers values
func NewFilesDeleteServiceUnavailable() *FilesDeleteServiceUnavailable {

	return &FilesDeleteServiceUnavailable{}
}

// WithPayload adds the payload to the files delete service unavailable response
func (o *FilesDeleteServiceUnavailable) WithPayload(payload *models.ErrorResponse) *FilesDeleteServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the files delete service unavailable response
func (o *FilesDeleteServiceUnavailable) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FilesDeleteServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
