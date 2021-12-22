// Code generated by go-swagger; DO NOT EDIT.

package staging

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/lexis-project/lexis-backend-services-api.git/models"
)

// StageCreatedCode is the HTTP code returned for type StageCreated
const StageCreatedCode int = 201

/*StageCreated The response code means that the data transfer has been initiated. Status of the transfer can be checked by querying the status.

swagger:response stageCreated
*/
type StageCreated struct {

	/*
	  In: Body
	*/
	Payload *models.SteeringRequestID `json:"body,omitempty"`
}

// NewStageCreated creates StageCreated with default headers values
func NewStageCreated() *StageCreated {

	return &StageCreated{}
}

// WithPayload adds the payload to the stage created response
func (o *StageCreated) WithPayload(payload *models.SteeringRequestID) *StageCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the stage created response
func (o *StageCreated) SetPayload(payload *models.SteeringRequestID) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *StageCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// StageBadRequestCode is the HTTP code returned for type StageBadRequest
const StageBadRequestCode int = 400

/*StageBadRequest This means that there's something wrong in the input parameters and the server couldn't understand the request.

swagger:response stageBadRequest
*/
type StageBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewStageBadRequest creates StageBadRequest with default headers values
func NewStageBadRequest() *StageBadRequest {

	return &StageBadRequest{}
}

// WithPayload adds the payload to the stage bad request response
func (o *StageBadRequest) WithPayload(payload *models.ErrorResponse) *StageBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the stage bad request response
func (o *StageBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *StageBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// StageUnauthorizedCode is the HTTP code returned for type StageUnauthorized
const StageUnauthorizedCode int = 401

/*StageUnauthorized This means that the user is not authenticated with keycloak and data transfer can't be triggered unless the user first log in with a valid user

swagger:response stageUnauthorized
*/
type StageUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewStageUnauthorized creates StageUnauthorized with default headers values
func NewStageUnauthorized() *StageUnauthorized {

	return &StageUnauthorized{}
}

// WithPayload adds the payload to the stage unauthorized response
func (o *StageUnauthorized) WithPayload(payload *models.ErrorResponse) *StageUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the stage unauthorized response
func (o *StageUnauthorized) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *StageUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// StageForbiddenCode is the HTTP code returned for type StageForbidden
const StageForbiddenCode int = 403

/*StageForbidden This means that the resource the user is trying to transfer from or to is not readable or writable by the user. User doesn't have the correct rights to either read the source file or write on the target system location.

swagger:response stageForbidden
*/
type StageForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewStageForbidden creates StageForbidden with default headers values
func NewStageForbidden() *StageForbidden {

	return &StageForbidden{}
}

// WithPayload adds the payload to the stage forbidden response
func (o *StageForbidden) WithPayload(payload *models.ErrorResponse) *StageForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the stage forbidden response
func (o *StageForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *StageForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// StageNotFoundCode is the HTTP code returned for type StageNotFound
const StageNotFoundCode int = 404

/*StageNotFound This means that either the source path or the the target path on the system doesn't exist.

swagger:response stageNotFound
*/
type StageNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewStageNotFound creates StageNotFound with default headers values
func NewStageNotFound() *StageNotFound {

	return &StageNotFound{}
}

// WithPayload adds the payload to the stage not found response
func (o *StageNotFound) WithPayload(payload *models.ErrorResponse) *StageNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the stage not found response
func (o *StageNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *StageNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// StageRequestURITooLongCode is the HTTP code returned for type StageRequestURITooLong
const StageRequestURITooLongCode int = 414

/*StageRequestURITooLong This means that the either the source path or the target path is longer than the server is willing to interpret.

swagger:response stageRequestUriTooLong
*/
type StageRequestURITooLong struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewStageRequestURITooLong creates StageRequestURITooLong with default headers values
func NewStageRequestURITooLong() *StageRequestURITooLong {

	return &StageRequestURITooLong{}
}

// WithPayload adds the payload to the stage request Uri too long response
func (o *StageRequestURITooLong) WithPayload(payload *models.ErrorResponse) *StageRequestURITooLong {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the stage request Uri too long response
func (o *StageRequestURITooLong) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *StageRequestURITooLong) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(414)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// StageTooManyRequestsCode is the HTTP code returned for type StageTooManyRequests
const StageTooManyRequestsCode int = 429

/*StageTooManyRequests This means that the user has sent too many requests in a given amount of time. The user is advised to wait and send the request again.

swagger:response stageTooManyRequests
*/
type StageTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewStageTooManyRequests creates StageTooManyRequests with default headers values
func NewStageTooManyRequests() *StageTooManyRequests {

	return &StageTooManyRequests{}
}

// WithPayload adds the payload to the stage too many requests response
func (o *StageTooManyRequests) WithPayload(payload *models.ErrorResponse) *StageTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the stage too many requests response
func (o *StageTooManyRequests) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *StageTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// StageInternalServerErrorCode is the HTTP code returned for type StageInternalServerError
const StageInternalServerErrorCode int = 500

/*StageInternalServerError This means that something has gone wrong on the data transfer steering server. The user is advised to wait and send the request again.

swagger:response stageInternalServerError
*/
type StageInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewStageInternalServerError creates StageInternalServerError with default headers values
func NewStageInternalServerError() *StageInternalServerError {

	return &StageInternalServerError{}
}

// WithPayload adds the payload to the stage internal server error response
func (o *StageInternalServerError) WithPayload(payload *models.ErrorResponse) *StageInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the stage internal server error response
func (o *StageInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *StageInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
