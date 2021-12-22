// Code generated by go-swagger; DO NOT EDIT.

package approval_system_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/lexis-project/lexis-backend-services-api.git/models"
)

// ListProjectHPCApprovedResourceRequestOKCode is the HTTP code returned for type ListProjectHPCApprovedResourceRequestOK
const ListProjectHPCApprovedResourceRequestOKCode int = 200

/*ListProjectHPCApprovedResourceRequestOK list of HPC approved resource requests returned for particular LEXIS project

swagger:response listProjectHPCApprovedResourceRequestOK
*/
type ListProjectHPCApprovedResourceRequestOK struct {

	/*
	  In: Body
	*/
	Payload []*models.ApprovalSystemApprovedResourceRequest `json:"body,omitempty"`
}

// NewListProjectHPCApprovedResourceRequestOK creates ListProjectHPCApprovedResourceRequestOK with default headers values
func NewListProjectHPCApprovedResourceRequestOK() *ListProjectHPCApprovedResourceRequestOK {

	return &ListProjectHPCApprovedResourceRequestOK{}
}

// WithPayload adds the payload to the list project h p c approved resource request o k response
func (o *ListProjectHPCApprovedResourceRequestOK) WithPayload(payload []*models.ApprovalSystemApprovedResourceRequest) *ListProjectHPCApprovedResourceRequestOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list project h p c approved resource request o k response
func (o *ListProjectHPCApprovedResourceRequestOK) SetPayload(payload []*models.ApprovalSystemApprovedResourceRequest) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListProjectHPCApprovedResourceRequestOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.ApprovalSystemApprovedResourceRequest, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListProjectHPCApprovedResourceRequestNotFoundCode is the HTTP code returned for type ListProjectHPCApprovedResourceRequestNotFound
const ListProjectHPCApprovedResourceRequestNotFoundCode int = 404

/*ListProjectHPCApprovedResourceRequestNotFound The AssociatedLEXISProject ID provided does not exist.

swagger:response listProjectHPCApprovedResourceRequestNotFound
*/
type ListProjectHPCApprovedResourceRequestNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ApprovalSystemMissingResponse `json:"body,omitempty"`
}

// NewListProjectHPCApprovedResourceRequestNotFound creates ListProjectHPCApprovedResourceRequestNotFound with default headers values
func NewListProjectHPCApprovedResourceRequestNotFound() *ListProjectHPCApprovedResourceRequestNotFound {

	return &ListProjectHPCApprovedResourceRequestNotFound{}
}

// WithPayload adds the payload to the list project h p c approved resource request not found response
func (o *ListProjectHPCApprovedResourceRequestNotFound) WithPayload(payload *models.ApprovalSystemMissingResponse) *ListProjectHPCApprovedResourceRequestNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list project h p c approved resource request not found response
func (o *ListProjectHPCApprovedResourceRequestNotFound) SetPayload(payload *models.ApprovalSystemMissingResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListProjectHPCApprovedResourceRequestNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListProjectHPCApprovedResourceRequestInternalServerErrorCode is the HTTP code returned for type ListProjectHPCApprovedResourceRequestInternalServerError
const ListProjectHPCApprovedResourceRequestInternalServerErrorCode int = 500

/*ListProjectHPCApprovedResourceRequestInternalServerError unexpected error

swagger:response listProjectHPCApprovedResourceRequestInternalServerError
*/
type ListProjectHPCApprovedResourceRequestInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ApprovalSystemErrorResponse `json:"body,omitempty"`
}

// NewListProjectHPCApprovedResourceRequestInternalServerError creates ListProjectHPCApprovedResourceRequestInternalServerError with default headers values
func NewListProjectHPCApprovedResourceRequestInternalServerError() *ListProjectHPCApprovedResourceRequestInternalServerError {

	return &ListProjectHPCApprovedResourceRequestInternalServerError{}
}

// WithPayload adds the payload to the list project h p c approved resource request internal server error response
func (o *ListProjectHPCApprovedResourceRequestInternalServerError) WithPayload(payload *models.ApprovalSystemErrorResponse) *ListProjectHPCApprovedResourceRequestInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list project h p c approved resource request internal server error response
func (o *ListProjectHPCApprovedResourceRequestInternalServerError) SetPayload(payload *models.ApprovalSystemErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListProjectHPCApprovedResourceRequestInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}