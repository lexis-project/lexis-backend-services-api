// Code generated by go-swagger; DO NOT EDIT.

package approval_system_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/lexis-project/lexis-backend-services-api.git/models"
)

// ListProjectHPCResourceRequestOKCode is the HTTP code returned for type ListProjectHPCResourceRequestOK
const ListProjectHPCResourceRequestOKCode int = 200

/*ListProjectHPCResourceRequestOK list of HPC resource requests returned for particular LEXIS project

swagger:response listProjectHPCResourceRequestOK
*/
type ListProjectHPCResourceRequestOK struct {

	/*
	  In: Body
	*/
	Payload []*models.ApprovalSystemResourceRequest `json:"body,omitempty"`
}

// NewListProjectHPCResourceRequestOK creates ListProjectHPCResourceRequestOK with default headers values
func NewListProjectHPCResourceRequestOK() *ListProjectHPCResourceRequestOK {

	return &ListProjectHPCResourceRequestOK{}
}

// WithPayload adds the payload to the list project h p c resource request o k response
func (o *ListProjectHPCResourceRequestOK) WithPayload(payload []*models.ApprovalSystemResourceRequest) *ListProjectHPCResourceRequestOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list project h p c resource request o k response
func (o *ListProjectHPCResourceRequestOK) SetPayload(payload []*models.ApprovalSystemResourceRequest) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListProjectHPCResourceRequestOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.ApprovalSystemResourceRequest, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListProjectHPCResourceRequestNotFoundCode is the HTTP code returned for type ListProjectHPCResourceRequestNotFound
const ListProjectHPCResourceRequestNotFoundCode int = 404

/*ListProjectHPCResourceRequestNotFound The AssociatedLEXISProject ID provided does not exist.

swagger:response listProjectHPCResourceRequestNotFound
*/
type ListProjectHPCResourceRequestNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ApprovalSystemMissingResponse `json:"body,omitempty"`
}

// NewListProjectHPCResourceRequestNotFound creates ListProjectHPCResourceRequestNotFound with default headers values
func NewListProjectHPCResourceRequestNotFound() *ListProjectHPCResourceRequestNotFound {

	return &ListProjectHPCResourceRequestNotFound{}
}

// WithPayload adds the payload to the list project h p c resource request not found response
func (o *ListProjectHPCResourceRequestNotFound) WithPayload(payload *models.ApprovalSystemMissingResponse) *ListProjectHPCResourceRequestNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list project h p c resource request not found response
func (o *ListProjectHPCResourceRequestNotFound) SetPayload(payload *models.ApprovalSystemMissingResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListProjectHPCResourceRequestNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListProjectHPCResourceRequestInternalServerErrorCode is the HTTP code returned for type ListProjectHPCResourceRequestInternalServerError
const ListProjectHPCResourceRequestInternalServerErrorCode int = 500

/*ListProjectHPCResourceRequestInternalServerError unexpected error

swagger:response listProjectHPCResourceRequestInternalServerError
*/
type ListProjectHPCResourceRequestInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ApprovalSystemErrorResponse `json:"body,omitempty"`
}

// NewListProjectHPCResourceRequestInternalServerError creates ListProjectHPCResourceRequestInternalServerError with default headers values
func NewListProjectHPCResourceRequestInternalServerError() *ListProjectHPCResourceRequestInternalServerError {

	return &ListProjectHPCResourceRequestInternalServerError{}
}

// WithPayload adds the payload to the list project h p c resource request internal server error response
func (o *ListProjectHPCResourceRequestInternalServerError) WithPayload(payload *models.ApprovalSystemErrorResponse) *ListProjectHPCResourceRequestInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list project h p c resource request internal server error response
func (o *ListProjectHPCResourceRequestInternalServerError) SetPayload(payload *models.ApprovalSystemErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListProjectHPCResourceRequestInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}