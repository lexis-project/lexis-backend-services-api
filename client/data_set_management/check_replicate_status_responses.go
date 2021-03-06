// Code generated by go-swagger; DO NOT EDIT.

package data_set_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/lexis-project/lexis-backend-services-api.git/models"
)

// CheckReplicateStatusReader is a Reader for the CheckReplicateStatus structure.
type CheckReplicateStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckReplicateStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCheckReplicateStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCheckReplicateStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCheckReplicateStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCheckReplicateStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCheckReplicateStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 414:
		result := NewCheckReplicateStatusRequestURITooLong()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCheckReplicateStatusTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCheckReplicateStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCheckReplicateStatusOK creates a CheckReplicateStatusOK with default headers values
func NewCheckReplicateStatusOK() *CheckReplicateStatusOK {
	return &CheckReplicateStatusOK{}
}

/*CheckReplicateStatusOK handles this case with default header values.

This means that the status has been returned to the user in the response body.
*/
type CheckReplicateStatusOK struct {
	Payload *models.DataReplication
}

func (o *CheckReplicateStatusOK) Error() string {
	return fmt.Sprintf("[GET /dataset/replicate/{request_id}][%d] checkReplicateStatusOK  %+v", 200, o.Payload)
}

func (o *CheckReplicateStatusOK) GetPayload() *models.DataReplication {
	return o.Payload
}

func (o *CheckReplicateStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DataReplication)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckReplicateStatusBadRequest creates a CheckReplicateStatusBadRequest with default headers values
func NewCheckReplicateStatusBadRequest() *CheckReplicateStatusBadRequest {
	return &CheckReplicateStatusBadRequest{}
}

/*CheckReplicateStatusBadRequest handles this case with default header values.

This means that there's something wrong in the input parameters and the server couldn't understand the request.
*/
type CheckReplicateStatusBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *CheckReplicateStatusBadRequest) Error() string {
	return fmt.Sprintf("[GET /dataset/replicate/{request_id}][%d] checkReplicateStatusBadRequest  %+v", 400, o.Payload)
}

func (o *CheckReplicateStatusBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CheckReplicateStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckReplicateStatusUnauthorized creates a CheckReplicateStatusUnauthorized with default headers values
func NewCheckReplicateStatusUnauthorized() *CheckReplicateStatusUnauthorized {
	return &CheckReplicateStatusUnauthorized{}
}

/*CheckReplicateStatusUnauthorized handles this case with default header values.

This means that the user is not authenticated with keycloak and data transfer can't be triggered unless the user first log in with a valid user
*/
type CheckReplicateStatusUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CheckReplicateStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /dataset/replicate/{request_id}][%d] checkReplicateStatusUnauthorized  %+v", 401, o.Payload)
}

func (o *CheckReplicateStatusUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CheckReplicateStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckReplicateStatusForbidden creates a CheckReplicateStatusForbidden with default headers values
func NewCheckReplicateStatusForbidden() *CheckReplicateStatusForbidden {
	return &CheckReplicateStatusForbidden{}
}

/*CheckReplicateStatusForbidden handles this case with default header values.

This means that the resource the user is trying to transfer from or to is not readable or writable by the user. User doesn't have the correct rights to either read the source file or write on the target system location.
*/
type CheckReplicateStatusForbidden struct {
	Payload *models.ErrorResponse
}

func (o *CheckReplicateStatusForbidden) Error() string {
	return fmt.Sprintf("[GET /dataset/replicate/{request_id}][%d] checkReplicateStatusForbidden  %+v", 403, o.Payload)
}

func (o *CheckReplicateStatusForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CheckReplicateStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckReplicateStatusNotFound creates a CheckReplicateStatusNotFound with default headers values
func NewCheckReplicateStatusNotFound() *CheckReplicateStatusNotFound {
	return &CheckReplicateStatusNotFound{}
}

/*CheckReplicateStatusNotFound handles this case with default header values.

This means that either the source path or the the target path on the system doesn't exist.
*/
type CheckReplicateStatusNotFound struct {
	Payload *models.ErrorResponse
}

func (o *CheckReplicateStatusNotFound) Error() string {
	return fmt.Sprintf("[GET /dataset/replicate/{request_id}][%d] checkReplicateStatusNotFound  %+v", 404, o.Payload)
}

func (o *CheckReplicateStatusNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CheckReplicateStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckReplicateStatusRequestURITooLong creates a CheckReplicateStatusRequestURITooLong with default headers values
func NewCheckReplicateStatusRequestURITooLong() *CheckReplicateStatusRequestURITooLong {
	return &CheckReplicateStatusRequestURITooLong{}
}

/*CheckReplicateStatusRequestURITooLong handles this case with default header values.

This means that the either the source path or the target path is longer than the server is willing to interpret.
*/
type CheckReplicateStatusRequestURITooLong struct {
	Payload *models.ErrorResponse
}

func (o *CheckReplicateStatusRequestURITooLong) Error() string {
	return fmt.Sprintf("[GET /dataset/replicate/{request_id}][%d] checkReplicateStatusRequestUriTooLong  %+v", 414, o.Payload)
}

func (o *CheckReplicateStatusRequestURITooLong) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CheckReplicateStatusRequestURITooLong) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckReplicateStatusTooManyRequests creates a CheckReplicateStatusTooManyRequests with default headers values
func NewCheckReplicateStatusTooManyRequests() *CheckReplicateStatusTooManyRequests {
	return &CheckReplicateStatusTooManyRequests{}
}

/*CheckReplicateStatusTooManyRequests handles this case with default header values.

This means that the user has sent too many requests in a given amount of time. The user is advised to wait and send the request again.
*/
type CheckReplicateStatusTooManyRequests struct {
	Payload *models.ErrorResponse
}

func (o *CheckReplicateStatusTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /dataset/replicate/{request_id}][%d] checkReplicateStatusTooManyRequests  %+v", 429, o.Payload)
}

func (o *CheckReplicateStatusTooManyRequests) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CheckReplicateStatusTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckReplicateStatusInternalServerError creates a CheckReplicateStatusInternalServerError with default headers values
func NewCheckReplicateStatusInternalServerError() *CheckReplicateStatusInternalServerError {
	return &CheckReplicateStatusInternalServerError{}
}

/*CheckReplicateStatusInternalServerError handles this case with default header values.

This means that something has gone wrong on the data transfer steering server. The user is advised to wait and send the request again.
*/
type CheckReplicateStatusInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *CheckReplicateStatusInternalServerError) Error() string {
	return fmt.Sprintf("[GET /dataset/replicate/{request_id}][%d] checkReplicateStatusInternalServerError  %+v", 500, o.Payload)
}

func (o *CheckReplicateStatusInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CheckReplicateStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
