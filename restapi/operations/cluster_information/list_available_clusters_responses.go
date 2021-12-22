// Code generated by go-swagger; DO NOT EDIT.

package cluster_information

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/lexis-project/lexis-backend-services-api.git/models"
)

// ListAvailableClustersOKCode is the HTTP code returned for type ListAvailableClustersOK
const ListAvailableClustersOKCode int = 200

/*ListAvailableClustersOK list of clusters retruend

swagger:response listAvailableClustersOK
*/
type ListAvailableClustersOK struct {

	/*
	  In: Body
	*/
	Payload []*models.HeappeCluster `json:"body,omitempty"`
}

// NewListAvailableClustersOK creates ListAvailableClustersOK with default headers values
func NewListAvailableClustersOK() *ListAvailableClustersOK {

	return &ListAvailableClustersOK{}
}

// WithPayload adds the payload to the list available clusters o k response
func (o *ListAvailableClustersOK) WithPayload(payload []*models.HeappeCluster) *ListAvailableClustersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list available clusters o k response
func (o *ListAvailableClustersOK) SetPayload(payload []*models.HeappeCluster) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAvailableClustersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.HeappeCluster, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListAvailableClustersBadRequestCode is the HTTP code returned for type ListAvailableClustersBadRequest
const ListAvailableClustersBadRequestCode int = 400

/*ListAvailableClustersBadRequest Bad Request

swagger:response listAvailableClustersBadRequest
*/
type ListAvailableClustersBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.HeappeBadRequest `json:"body,omitempty"`
}

// NewListAvailableClustersBadRequest creates ListAvailableClustersBadRequest with default headers values
func NewListAvailableClustersBadRequest() *ListAvailableClustersBadRequest {

	return &ListAvailableClustersBadRequest{}
}

// WithPayload adds the payload to the list available clusters bad request response
func (o *ListAvailableClustersBadRequest) WithPayload(payload *models.HeappeBadRequest) *ListAvailableClustersBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list available clusters bad request response
func (o *ListAvailableClustersBadRequest) SetPayload(payload *models.HeappeBadRequest) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAvailableClustersBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListAvailableClustersRequestEntityTooLargeCode is the HTTP code returned for type ListAvailableClustersRequestEntityTooLarge
const ListAvailableClustersRequestEntityTooLargeCode int = 413

/*ListAvailableClustersRequestEntityTooLarge Client Error

swagger:response listAvailableClustersRequestEntityTooLarge
*/
type ListAvailableClustersRequestEntityTooLarge struct {

	/*
	  In: Body
	*/
	Payload *models.HeappeError `json:"body,omitempty"`
}

// NewListAvailableClustersRequestEntityTooLarge creates ListAvailableClustersRequestEntityTooLarge with default headers values
func NewListAvailableClustersRequestEntityTooLarge() *ListAvailableClustersRequestEntityTooLarge {

	return &ListAvailableClustersRequestEntityTooLarge{}
}

// WithPayload adds the payload to the list available clusters request entity too large response
func (o *ListAvailableClustersRequestEntityTooLarge) WithPayload(payload *models.HeappeError) *ListAvailableClustersRequestEntityTooLarge {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list available clusters request entity too large response
func (o *ListAvailableClustersRequestEntityTooLarge) SetPayload(payload *models.HeappeError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAvailableClustersRequestEntityTooLarge) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(413)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListAvailableClustersTooManyRequestsCode is the HTTP code returned for type ListAvailableClustersTooManyRequests
const ListAvailableClustersTooManyRequestsCode int = 429

/*ListAvailableClustersTooManyRequests Client Error

swagger:response listAvailableClustersTooManyRequests
*/
type ListAvailableClustersTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *models.HeappeError `json:"body,omitempty"`
}

// NewListAvailableClustersTooManyRequests creates ListAvailableClustersTooManyRequests with default headers values
func NewListAvailableClustersTooManyRequests() *ListAvailableClustersTooManyRequests {

	return &ListAvailableClustersTooManyRequests{}
}

// WithPayload adds the payload to the list available clusters too many requests response
func (o *ListAvailableClustersTooManyRequests) WithPayload(payload *models.HeappeError) *ListAvailableClustersTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list available clusters too many requests response
func (o *ListAvailableClustersTooManyRequests) SetPayload(payload *models.HeappeError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAvailableClustersTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListAvailableClustersInternalServerErrorCode is the HTTP code returned for type ListAvailableClustersInternalServerError
const ListAvailableClustersInternalServerErrorCode int = 500

/*ListAvailableClustersInternalServerError Server Error

swagger:response listAvailableClustersInternalServerError
*/
type ListAvailableClustersInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewListAvailableClustersInternalServerError creates ListAvailableClustersInternalServerError with default headers values
func NewListAvailableClustersInternalServerError() *ListAvailableClustersInternalServerError {

	return &ListAvailableClustersInternalServerError{}
}

// WithPayload adds the payload to the list available clusters internal server error response
func (o *ListAvailableClustersInternalServerError) WithPayload(payload *models.ErrorResponse) *ListAvailableClustersInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list available clusters internal server error response
func (o *ListAvailableClustersInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAvailableClustersInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}