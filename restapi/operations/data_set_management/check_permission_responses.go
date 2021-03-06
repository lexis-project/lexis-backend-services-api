// Code generated by go-swagger; DO NOT EDIT.

package data_set_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/lexis-project/lexis-backend-services-api.git/models"
)

// CheckPermissionOKCode is the HTTP code returned for type CheckPermissionOK
const CheckPermissionOKCode int = 200

/*CheckPermissionOK User has permission to write

swagger:response checkPermissionOK
*/
type CheckPermissionOK struct {

	/*
	  In: Body
	*/
	Payload *CheckPermissionOKBody `json:"body,omitempty"`
}

// NewCheckPermissionOK creates CheckPermissionOK with default headers values
func NewCheckPermissionOK() *CheckPermissionOK {

	return &CheckPermissionOK{}
}

// WithPayload adds the payload to the check permission o k response
func (o *CheckPermissionOK) WithPayload(payload *CheckPermissionOKBody) *CheckPermissionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check permission o k response
func (o *CheckPermissionOK) SetPayload(payload *CheckPermissionOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckPermissionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CheckPermissionBadRequestCode is the HTTP code returned for type CheckPermissionBadRequest
const CheckPermissionBadRequestCode int = 400

/*CheckPermissionBadRequest Malformed request

swagger:response checkPermissionBadRequest
*/
type CheckPermissionBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCheckPermissionBadRequest creates CheckPermissionBadRequest with default headers values
func NewCheckPermissionBadRequest() *CheckPermissionBadRequest {

	return &CheckPermissionBadRequest{}
}

// WithPayload adds the payload to the check permission bad request response
func (o *CheckPermissionBadRequest) WithPayload(payload *models.ErrorResponse) *CheckPermissionBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check permission bad request response
func (o *CheckPermissionBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckPermissionBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CheckPermissionUnauthorizedCode is the HTTP code returned for type CheckPermissionUnauthorized
const CheckPermissionUnauthorizedCode int = 401

/*CheckPermissionUnauthorized Authorization failed

swagger:response checkPermissionUnauthorized
*/
type CheckPermissionUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCheckPermissionUnauthorized creates CheckPermissionUnauthorized with default headers values
func NewCheckPermissionUnauthorized() *CheckPermissionUnauthorized {

	return &CheckPermissionUnauthorized{}
}

// WithPayload adds the payload to the check permission unauthorized response
func (o *CheckPermissionUnauthorized) WithPayload(payload *models.ErrorResponse) *CheckPermissionUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check permission unauthorized response
func (o *CheckPermissionUnauthorized) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckPermissionUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CheckPermissionForbiddenCode is the HTTP code returned for type CheckPermissionForbidden
const CheckPermissionForbiddenCode int = 403

/*CheckPermissionForbidden User does not have permission

swagger:response checkPermissionForbidden
*/
type CheckPermissionForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCheckPermissionForbidden creates CheckPermissionForbidden with default headers values
func NewCheckPermissionForbidden() *CheckPermissionForbidden {

	return &CheckPermissionForbidden{}
}

// WithPayload adds the payload to the check permission forbidden response
func (o *CheckPermissionForbidden) WithPayload(payload *models.ErrorResponse) *CheckPermissionForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check permission forbidden response
func (o *CheckPermissionForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckPermissionForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CheckPermissionBadGatewayCode is the HTTP code returned for type CheckPermissionBadGateway
const CheckPermissionBadGatewayCode int = 502

/*CheckPermissionBadGateway Bad Gateway

swagger:response checkPermissionBadGateway
*/
type CheckPermissionBadGateway struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCheckPermissionBadGateway creates CheckPermissionBadGateway with default headers values
func NewCheckPermissionBadGateway() *CheckPermissionBadGateway {

	return &CheckPermissionBadGateway{}
}

// WithPayload adds the payload to the check permission bad gateway response
func (o *CheckPermissionBadGateway) WithPayload(payload *models.ErrorResponse) *CheckPermissionBadGateway {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check permission bad gateway response
func (o *CheckPermissionBadGateway) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckPermissionBadGateway) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(502)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CheckPermissionServiceUnavailableCode is the HTTP code returned for type CheckPermissionServiceUnavailable
const CheckPermissionServiceUnavailableCode int = 503

/*CheckPermissionServiceUnavailable Error accessing backend service

swagger:response checkPermissionServiceUnavailable
*/
type CheckPermissionServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewCheckPermissionServiceUnavailable creates CheckPermissionServiceUnavailable with default headers values
func NewCheckPermissionServiceUnavailable() *CheckPermissionServiceUnavailable {

	return &CheckPermissionServiceUnavailable{}
}

// WithPayload adds the payload to the check permission service unavailable response
func (o *CheckPermissionServiceUnavailable) WithPayload(payload *models.ErrorResponse) *CheckPermissionServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check permission service unavailable response
func (o *CheckPermissionServiceUnavailable) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckPermissionServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
