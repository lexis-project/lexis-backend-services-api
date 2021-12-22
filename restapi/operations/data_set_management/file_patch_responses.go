// Code generated by go-swagger; DO NOT EDIT.

package data_set_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"github.com/lexis-project/lexis-backend-services-api.git/models"
)

// FilePatchNoContentCode is the HTTP code returned for type FilePatchNoContent
const FilePatchNoContentCode int = 204

/*FilePatchNoContent Upload offset was updated

swagger:response filePatchNoContent
*/
type FilePatchNoContent struct {
	/*Needed to make browsers accept the additional headers used by
	the tus protocol.


	*/
	AccessControlExposeHeaders string `json:"Access-Control-Expose-Headers"`
	/*Protocol version

	 */
	TusResumable string `json:"Tus-Resumable"`
	/*Added by the expiration extension. The Upload-Expires response header indicates the time after which the unfinished upload expires. A Server MAY wish to remove incomplete uploads after a given period of time to prevent abandoned uploads from taking up extra storage. The Client SHOULD use this header to determine if an upload is still valid before attempting to resume the upload. This header MUST be included in every PATCH response if the upload is going to expire. If the expiration is known at the creation, the Upload-Expires header MUST be included in the response to the initial POST request. Its value MAY change over time. If a Client does attempt to resume an upload which has since been removed by the Server, the Server SHOULD respond with the 404 Not Found or 410 Gone status. The latter one SHOULD be used if the Server is keeping track of expired uploads. In both cases the Client SHOULD start a new upload. The value of the Upload-Expires header MUST be in RFC 7231 datetime format.

	 */
	UploadExpires string `json:"Upload-Expires"`
	/*The Upload-Offset request and response header indicates a byte offset within a resource. The value MUST be a non-negative integer.

	 */
	UploadOffset int64 `json:"Upload-Offset"`
}

// NewFilePatchNoContent creates FilePatchNoContent with default headers values
func NewFilePatchNoContent() *FilePatchNoContent {

	return &FilePatchNoContent{}
}

// WithAccessControlExposeHeaders adds the accessControlExposeHeaders to the file patch no content response
func (o *FilePatchNoContent) WithAccessControlExposeHeaders(accessControlExposeHeaders string) *FilePatchNoContent {
	o.AccessControlExposeHeaders = accessControlExposeHeaders
	return o
}

// SetAccessControlExposeHeaders sets the accessControlExposeHeaders to the file patch no content response
func (o *FilePatchNoContent) SetAccessControlExposeHeaders(accessControlExposeHeaders string) {
	o.AccessControlExposeHeaders = accessControlExposeHeaders
}

// WithTusResumable adds the tusResumable to the file patch no content response
func (o *FilePatchNoContent) WithTusResumable(tusResumable string) *FilePatchNoContent {
	o.TusResumable = tusResumable
	return o
}

// SetTusResumable sets the tusResumable to the file patch no content response
func (o *FilePatchNoContent) SetTusResumable(tusResumable string) {
	o.TusResumable = tusResumable
}

// WithUploadExpires adds the uploadExpires to the file patch no content response
func (o *FilePatchNoContent) WithUploadExpires(uploadExpires string) *FilePatchNoContent {
	o.UploadExpires = uploadExpires
	return o
}

// SetUploadExpires sets the uploadExpires to the file patch no content response
func (o *FilePatchNoContent) SetUploadExpires(uploadExpires string) {
	o.UploadExpires = uploadExpires
}

// WithUploadOffset adds the uploadOffset to the file patch no content response
func (o *FilePatchNoContent) WithUploadOffset(uploadOffset int64) *FilePatchNoContent {
	o.UploadOffset = uploadOffset
	return o
}

// SetUploadOffset sets the uploadOffset to the file patch no content response
func (o *FilePatchNoContent) SetUploadOffset(uploadOffset int64) {
	o.UploadOffset = uploadOffset
}

// WriteResponse to the client
func (o *FilePatchNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

	// response header Upload-Expires

	uploadExpires := o.UploadExpires
	if uploadExpires != "" {
		rw.Header().Set("Upload-Expires", uploadExpires)
	}

	// response header Upload-Offset

	uploadOffset := swag.FormatInt64(o.UploadOffset)
	if uploadOffset != "" {
		rw.Header().Set("Upload-Offset", uploadOffset)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// FilePatchBadRequestCode is the HTTP code returned for type FilePatchBadRequest
const FilePatchBadRequestCode int = 400

/*FilePatchBadRequest Added by the checksum extension. The checksum algorithm is not supported by the server

swagger:response filePatchBadRequest
*/
type FilePatchBadRequest struct {
	/*Needed to make browsers accept the additional headers used by
	the tus protocol.


	*/
	AccessControlExposeHeaders string `json:"Access-Control-Expose-Headers"`
	/*Protocol version

	 */
	TusResumable string `json:"Tus-Resumable"`
}

// NewFilePatchBadRequest creates FilePatchBadRequest with default headers values
func NewFilePatchBadRequest() *FilePatchBadRequest {

	return &FilePatchBadRequest{}
}

// WithAccessControlExposeHeaders adds the accessControlExposeHeaders to the file patch bad request response
func (o *FilePatchBadRequest) WithAccessControlExposeHeaders(accessControlExposeHeaders string) *FilePatchBadRequest {
	o.AccessControlExposeHeaders = accessControlExposeHeaders
	return o
}

// SetAccessControlExposeHeaders sets the accessControlExposeHeaders to the file patch bad request response
func (o *FilePatchBadRequest) SetAccessControlExposeHeaders(accessControlExposeHeaders string) {
	o.AccessControlExposeHeaders = accessControlExposeHeaders
}

// WithTusResumable adds the tusResumable to the file patch bad request response
func (o *FilePatchBadRequest) WithTusResumable(tusResumable string) *FilePatchBadRequest {
	o.TusResumable = tusResumable
	return o
}

// SetTusResumable sets the tusResumable to the file patch bad request response
func (o *FilePatchBadRequest) SetTusResumable(tusResumable string) {
	o.TusResumable = tusResumable
}

// WriteResponse to the client
func (o *FilePatchBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

	rw.WriteHeader(400)
}

// FilePatchUnauthorizedCode is the HTTP code returned for type FilePatchUnauthorized
const FilePatchUnauthorizedCode int = 401

/*FilePatchUnauthorized Authorization failed

swagger:response filePatchUnauthorized
*/
type FilePatchUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewFilePatchUnauthorized creates FilePatchUnauthorized with default headers values
func NewFilePatchUnauthorized() *FilePatchUnauthorized {

	return &FilePatchUnauthorized{}
}

// WithPayload adds the payload to the file patch unauthorized response
func (o *FilePatchUnauthorized) WithPayload(payload *models.ErrorResponse) *FilePatchUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the file patch unauthorized response
func (o *FilePatchUnauthorized) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FilePatchUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// FilePatchForbiddenCode is the HTTP code returned for type FilePatchForbidden
const FilePatchForbiddenCode int = 403

/*FilePatchForbidden In the concatenation extension, the Server MUST respond with the 403 Forbidden status to PATCH requests against a final upload URL and MUST NOT modify the final or its partial uploads.

swagger:response filePatchForbidden
*/
type FilePatchForbidden struct {
	/*Needed to make browsers accept the additional headers used by
	the tus protocol.


	*/
	AccessControlExposeHeaders string `json:"Access-Control-Expose-Headers"`
	/*Protocol version

	 */
	TusResumable string `json:"Tus-Resumable"`
}

// NewFilePatchForbidden creates FilePatchForbidden with default headers values
func NewFilePatchForbidden() *FilePatchForbidden {

	return &FilePatchForbidden{}
}

// WithAccessControlExposeHeaders adds the accessControlExposeHeaders to the file patch forbidden response
func (o *FilePatchForbidden) WithAccessControlExposeHeaders(accessControlExposeHeaders string) *FilePatchForbidden {
	o.AccessControlExposeHeaders = accessControlExposeHeaders
	return o
}

// SetAccessControlExposeHeaders sets the accessControlExposeHeaders to the file patch forbidden response
func (o *FilePatchForbidden) SetAccessControlExposeHeaders(accessControlExposeHeaders string) {
	o.AccessControlExposeHeaders = accessControlExposeHeaders
}

// WithTusResumable adds the tusResumable to the file patch forbidden response
func (o *FilePatchForbidden) WithTusResumable(tusResumable string) *FilePatchForbidden {
	o.TusResumable = tusResumable
	return o
}

// SetTusResumable sets the tusResumable to the file patch forbidden response
func (o *FilePatchForbidden) SetTusResumable(tusResumable string) {
	o.TusResumable = tusResumable
}

// WriteResponse to the client
func (o *FilePatchForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

	rw.WriteHeader(403)
}

// FilePatchNotFoundCode is the HTTP code returned for type FilePatchNotFound
const FilePatchNotFoundCode int = 404

/*FilePatchNotFound PATCH request against a non-existent resource

swagger:response filePatchNotFound
*/
type FilePatchNotFound struct {
	/*Needed to make browsers accept the additional headers used by
	the tus protocol.


	*/
	AccessControlExposeHeaders string `json:"Access-Control-Expose-Headers"`
	/*Protocol version

	 */
	TusResumable string `json:"Tus-Resumable"`
}

// NewFilePatchNotFound creates FilePatchNotFound with default headers values
func NewFilePatchNotFound() *FilePatchNotFound {

	return &FilePatchNotFound{}
}

// WithAccessControlExposeHeaders adds the accessControlExposeHeaders to the file patch not found response
func (o *FilePatchNotFound) WithAccessControlExposeHeaders(accessControlExposeHeaders string) *FilePatchNotFound {
	o.AccessControlExposeHeaders = accessControlExposeHeaders
	return o
}

// SetAccessControlExposeHeaders sets the accessControlExposeHeaders to the file patch not found response
func (o *FilePatchNotFound) SetAccessControlExposeHeaders(accessControlExposeHeaders string) {
	o.AccessControlExposeHeaders = accessControlExposeHeaders
}

// WithTusResumable adds the tusResumable to the file patch not found response
func (o *FilePatchNotFound) WithTusResumable(tusResumable string) *FilePatchNotFound {
	o.TusResumable = tusResumable
	return o
}

// SetTusResumable sets the tusResumable to the file patch not found response
func (o *FilePatchNotFound) SetTusResumable(tusResumable string) {
	o.TusResumable = tusResumable
}

// WriteResponse to the client
func (o *FilePatchNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

	rw.WriteHeader(404)
}

// FilePatchGoneCode is the HTTP code returned for type FilePatchGone
const FilePatchGoneCode int = 410

/*FilePatchGone PATCH request against a non-existent resource

swagger:response filePatchGone
*/
type FilePatchGone struct {
	/*Needed to make browsers accept the additional headers used by
	the tus protocol.


	*/
	AccessControlExposeHeaders string `json:"Access-Control-Expose-Headers"`
	/*Protocol version

	 */
	TusResumable string `json:"Tus-Resumable"`
}

// NewFilePatchGone creates FilePatchGone with default headers values
func NewFilePatchGone() *FilePatchGone {

	return &FilePatchGone{}
}

// WithAccessControlExposeHeaders adds the accessControlExposeHeaders to the file patch gone response
func (o *FilePatchGone) WithAccessControlExposeHeaders(accessControlExposeHeaders string) *FilePatchGone {
	o.AccessControlExposeHeaders = accessControlExposeHeaders
	return o
}

// SetAccessControlExposeHeaders sets the accessControlExposeHeaders to the file patch gone response
func (o *FilePatchGone) SetAccessControlExposeHeaders(accessControlExposeHeaders string) {
	o.AccessControlExposeHeaders = accessControlExposeHeaders
}

// WithTusResumable adds the tusResumable to the file patch gone response
func (o *FilePatchGone) WithTusResumable(tusResumable string) *FilePatchGone {
	o.TusResumable = tusResumable
	return o
}

// SetTusResumable sets the tusResumable to the file patch gone response
func (o *FilePatchGone) SetTusResumable(tusResumable string) {
	o.TusResumable = tusResumable
}

// WriteResponse to the client
func (o *FilePatchGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

	rw.WriteHeader(410)
}

// FilePatchPreconditionFailedCode is the HTTP code returned for type FilePatchPreconditionFailed
const FilePatchPreconditionFailedCode int = 412

/*FilePatchPreconditionFailed Precondition Failed

swagger:response filePatchPreconditionFailed
*/
type FilePatchPreconditionFailed struct {
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

// NewFilePatchPreconditionFailed creates FilePatchPreconditionFailed with default headers values
func NewFilePatchPreconditionFailed() *FilePatchPreconditionFailed {

	return &FilePatchPreconditionFailed{}
}

// WithAccessControlExposeHeaders adds the accessControlExposeHeaders to the file patch precondition failed response
func (o *FilePatchPreconditionFailed) WithAccessControlExposeHeaders(accessControlExposeHeaders string) *FilePatchPreconditionFailed {
	o.AccessControlExposeHeaders = accessControlExposeHeaders
	return o
}

// SetAccessControlExposeHeaders sets the accessControlExposeHeaders to the file patch precondition failed response
func (o *FilePatchPreconditionFailed) SetAccessControlExposeHeaders(accessControlExposeHeaders string) {
	o.AccessControlExposeHeaders = accessControlExposeHeaders
}

// WithTusResumable adds the tusResumable to the file patch precondition failed response
func (o *FilePatchPreconditionFailed) WithTusResumable(tusResumable string) *FilePatchPreconditionFailed {
	o.TusResumable = tusResumable
	return o
}

// SetTusResumable sets the tusResumable to the file patch precondition failed response
func (o *FilePatchPreconditionFailed) SetTusResumable(tusResumable string) {
	o.TusResumable = tusResumable
}

// WithTusVersion adds the tusVersion to the file patch precondition failed response
func (o *FilePatchPreconditionFailed) WithTusVersion(tusVersion string) *FilePatchPreconditionFailed {
	o.TusVersion = tusVersion
	return o
}

// SetTusVersion sets the tusVersion to the file patch precondition failed response
func (o *FilePatchPreconditionFailed) SetTusVersion(tusVersion string) {
	o.TusVersion = tusVersion
}

// WriteResponse to the client
func (o *FilePatchPreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// FilePatchUnsupportedMediaTypeCode is the HTTP code returned for type FilePatchUnsupportedMediaType
const FilePatchUnsupportedMediaTypeCode int = 415

/*FilePatchUnsupportedMediaType Content-Type was not application/offset+octet-stream

swagger:response filePatchUnsupportedMediaType
*/
type FilePatchUnsupportedMediaType struct {
	/*Needed to make browsers accept the additional headers used by
	the tus protocol.


	*/
	AccessControlExposeHeaders string `json:"Access-Control-Expose-Headers"`
	/*Protocol version

	 */
	TusResumable string `json:"Tus-Resumable"`
}

// NewFilePatchUnsupportedMediaType creates FilePatchUnsupportedMediaType with default headers values
func NewFilePatchUnsupportedMediaType() *FilePatchUnsupportedMediaType {

	return &FilePatchUnsupportedMediaType{}
}

// WithAccessControlExposeHeaders adds the accessControlExposeHeaders to the file patch unsupported media type response
func (o *FilePatchUnsupportedMediaType) WithAccessControlExposeHeaders(accessControlExposeHeaders string) *FilePatchUnsupportedMediaType {
	o.AccessControlExposeHeaders = accessControlExposeHeaders
	return o
}

// SetAccessControlExposeHeaders sets the accessControlExposeHeaders to the file patch unsupported media type response
func (o *FilePatchUnsupportedMediaType) SetAccessControlExposeHeaders(accessControlExposeHeaders string) {
	o.AccessControlExposeHeaders = accessControlExposeHeaders
}

// WithTusResumable adds the tusResumable to the file patch unsupported media type response
func (o *FilePatchUnsupportedMediaType) WithTusResumable(tusResumable string) *FilePatchUnsupportedMediaType {
	o.TusResumable = tusResumable
	return o
}

// SetTusResumable sets the tusResumable to the file patch unsupported media type response
func (o *FilePatchUnsupportedMediaType) SetTusResumable(tusResumable string) {
	o.TusResumable = tusResumable
}

// WriteResponse to the client
func (o *FilePatchUnsupportedMediaType) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

	rw.WriteHeader(415)
}

// FilePatchStatus460Code is the HTTP code returned for type FilePatchStatus460
const FilePatchStatus460Code int = 460

/*FilePatchStatus460 Added by the checksum extension. Checksums mismatch

swagger:response filePatchStatus460
*/
type FilePatchStatus460 struct {
	/*Needed to make browsers accept the additional headers used by
	the tus protocol.


	*/
	AccessControlExposeHeaders string `json:"Access-Control-Expose-Headers"`
	/*Protocol version

	 */
	TusResumable string `json:"Tus-Resumable"`
}

// NewFilePatchStatus460 creates FilePatchStatus460 with default headers values
func NewFilePatchStatus460() *FilePatchStatus460 {

	return &FilePatchStatus460{}
}

// WithAccessControlExposeHeaders adds the accessControlExposeHeaders to the file patch status460 response
func (o *FilePatchStatus460) WithAccessControlExposeHeaders(accessControlExposeHeaders string) *FilePatchStatus460 {
	o.AccessControlExposeHeaders = accessControlExposeHeaders
	return o
}

// SetAccessControlExposeHeaders sets the accessControlExposeHeaders to the file patch status460 response
func (o *FilePatchStatus460) SetAccessControlExposeHeaders(accessControlExposeHeaders string) {
	o.AccessControlExposeHeaders = accessControlExposeHeaders
}

// WithTusResumable adds the tusResumable to the file patch status460 response
func (o *FilePatchStatus460) WithTusResumable(tusResumable string) *FilePatchStatus460 {
	o.TusResumable = tusResumable
	return o
}

// SetTusResumable sets the tusResumable to the file patch status460 response
func (o *FilePatchStatus460) SetTusResumable(tusResumable string) {
	o.TusResumable = tusResumable
}

// WriteResponse to the client
func (o *FilePatchStatus460) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

	rw.WriteHeader(460)
}

// FilePatchServiceUnavailableCode is the HTTP code returned for type FilePatchServiceUnavailable
const FilePatchServiceUnavailableCode int = 503

/*FilePatchServiceUnavailable Service Unavailable

swagger:response filePatchServiceUnavailable
*/
type FilePatchServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewFilePatchServiceUnavailable creates FilePatchServiceUnavailable with default headers values
func NewFilePatchServiceUnavailable() *FilePatchServiceUnavailable {

	return &FilePatchServiceUnavailable{}
}

// WithPayload adds the payload to the file patch service unavailable response
func (o *FilePatchServiceUnavailable) WithPayload(payload *models.ErrorResponse) *FilePatchServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the file patch service unavailable response
func (o *FilePatchServiceUnavailable) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FilePatchServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
