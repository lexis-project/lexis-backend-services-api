// Code generated by go-swagger; DO NOT EDIT.

package data_set_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/lexis-project/lexis-backend-services-api.git/models"
)

// ReplicateReader is a Reader for the Replicate structure.
type ReplicateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplicateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewReplicateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReplicateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewReplicateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewReplicateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReplicateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 414:
		result := NewReplicateRequestURITooLong()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewReplicateTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReplicateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReplicateCreated creates a ReplicateCreated with default headers values
func NewReplicateCreated() *ReplicateCreated {
	return &ReplicateCreated{}
}

/*ReplicateCreated handles this case with default header values.

The response code means that the data transfer has been initiated. Status of the transfer can be checked by querying the status.
*/
type ReplicateCreated struct {
	Payload *models.SteeringRequestID
}

func (o *ReplicateCreated) Error() string {
	return fmt.Sprintf("[POST /dataset/replicate][%d] replicateCreated  %+v", 201, o.Payload)
}

func (o *ReplicateCreated) GetPayload() *models.SteeringRequestID {
	return o.Payload
}

func (o *ReplicateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SteeringRequestID)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplicateBadRequest creates a ReplicateBadRequest with default headers values
func NewReplicateBadRequest() *ReplicateBadRequest {
	return &ReplicateBadRequest{}
}

/*ReplicateBadRequest handles this case with default header values.

This means that there's something wrong in the input parameters and the server couldn't understand the request.
*/
type ReplicateBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *ReplicateBadRequest) Error() string {
	return fmt.Sprintf("[POST /dataset/replicate][%d] replicateBadRequest  %+v", 400, o.Payload)
}

func (o *ReplicateBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ReplicateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplicateUnauthorized creates a ReplicateUnauthorized with default headers values
func NewReplicateUnauthorized() *ReplicateUnauthorized {
	return &ReplicateUnauthorized{}
}

/*ReplicateUnauthorized handles this case with default header values.

This means that the user is not authenticated with keycloak and data transfer can't be triggered unless the user first log in with a valid user
*/
type ReplicateUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *ReplicateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /dataset/replicate][%d] replicateUnauthorized  %+v", 401, o.Payload)
}

func (o *ReplicateUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ReplicateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplicateForbidden creates a ReplicateForbidden with default headers values
func NewReplicateForbidden() *ReplicateForbidden {
	return &ReplicateForbidden{}
}

/*ReplicateForbidden handles this case with default header values.

This means that the resource the user is trying to transfer from or to is not readable or writable by the user. User doesn't have the correct rights to either read the source file or write on the target system location.
*/
type ReplicateForbidden struct {
	Payload *models.ErrorResponse
}

func (o *ReplicateForbidden) Error() string {
	return fmt.Sprintf("[POST /dataset/replicate][%d] replicateForbidden  %+v", 403, o.Payload)
}

func (o *ReplicateForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ReplicateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplicateNotFound creates a ReplicateNotFound with default headers values
func NewReplicateNotFound() *ReplicateNotFound {
	return &ReplicateNotFound{}
}

/*ReplicateNotFound handles this case with default header values.

This means that either the source path or the the target path on the system doesn't exist.
*/
type ReplicateNotFound struct {
	Payload *models.ErrorResponse
}

func (o *ReplicateNotFound) Error() string {
	return fmt.Sprintf("[POST /dataset/replicate][%d] replicateNotFound  %+v", 404, o.Payload)
}

func (o *ReplicateNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ReplicateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplicateRequestURITooLong creates a ReplicateRequestURITooLong with default headers values
func NewReplicateRequestURITooLong() *ReplicateRequestURITooLong {
	return &ReplicateRequestURITooLong{}
}

/*ReplicateRequestURITooLong handles this case with default header values.

This means that the either the source path or the target path is longer than the server is willing to interpret.
*/
type ReplicateRequestURITooLong struct {
	Payload *models.ErrorResponse
}

func (o *ReplicateRequestURITooLong) Error() string {
	return fmt.Sprintf("[POST /dataset/replicate][%d] replicateRequestUriTooLong  %+v", 414, o.Payload)
}

func (o *ReplicateRequestURITooLong) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ReplicateRequestURITooLong) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplicateTooManyRequests creates a ReplicateTooManyRequests with default headers values
func NewReplicateTooManyRequests() *ReplicateTooManyRequests {
	return &ReplicateTooManyRequests{}
}

/*ReplicateTooManyRequests handles this case with default header values.

This means that the user has sent too many requests in a given amount of time. The user is advised to wait and send the request again.
*/
type ReplicateTooManyRequests struct {
	Payload *models.ErrorResponse
}

func (o *ReplicateTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /dataset/replicate][%d] replicateTooManyRequests  %+v", 429, o.Payload)
}

func (o *ReplicateTooManyRequests) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ReplicateTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplicateInternalServerError creates a ReplicateInternalServerError with default headers values
func NewReplicateInternalServerError() *ReplicateInternalServerError {
	return &ReplicateInternalServerError{}
}

/*ReplicateInternalServerError handles this case with default header values.

This means that something has gone wrong on the data transfer steering server. The user is advised to wait and send the request again.
*/
type ReplicateInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ReplicateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /dataset/replicate][%d] replicateInternalServerError  %+v", 500, o.Payload)
}

func (o *ReplicateInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ReplicateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ReplicateBody replicate body
swagger:model ReplicateBody
*/
type ReplicateBody struct {

	// source path
	// Required: true
	SourcePath *string `json:"source_path"`

	// Possible values are: "lrz_iRODS", "it4i_iRODS"
	//
	// Required: true
	SourceSystem *string `json:"source_system"`

	// target path
	TargetPath string `json:"target_path,omitempty"`

	// Possible values are: "lrz_iRODS", "it4i_iRODS"
	//
	// Required: true
	TargetSystem *string `json:"target_system"`
}

// Validate validates this replicate body
func (o *ReplicateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSourcePath(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSourceSystem(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTargetSystem(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ReplicateBody) validateSourcePath(formats strfmt.Registry) error {

	if err := validate.Required("parameters"+"."+"source_path", "body", o.SourcePath); err != nil {
		return err
	}

	return nil
}

func (o *ReplicateBody) validateSourceSystem(formats strfmt.Registry) error {

	if err := validate.Required("parameters"+"."+"source_system", "body", o.SourceSystem); err != nil {
		return err
	}

	return nil
}

func (o *ReplicateBody) validateTargetSystem(formats strfmt.Registry) error {

	if err := validate.Required("parameters"+"."+"target_system", "body", o.TargetSystem); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ReplicateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ReplicateBody) UnmarshalBinary(b []byte) error {
	var res ReplicateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
