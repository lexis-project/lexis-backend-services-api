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

// CheckSizeStatusReader is a Reader for the CheckSizeStatus structure.
type CheckSizeStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckSizeStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCheckSizeStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCheckSizeStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCheckSizeStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCheckSizeStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCheckSizeStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 414:
		result := NewCheckSizeStatusRequestURITooLong()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCheckSizeStatusTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCheckSizeStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCheckSizeStatusOK creates a CheckSizeStatusOK with default headers values
func NewCheckSizeStatusOK() *CheckSizeStatusOK {
	return &CheckSizeStatusOK{}
}

/*CheckSizeStatusOK handles this case with default header values.

This means that the status has been returned to the user in the response body.
*/
type CheckSizeStatusOK struct {
	Payload *models.DataSize
}

func (o *CheckSizeStatusOK) Error() string {
	return fmt.Sprintf("[GET /dataset/data/size/{request_id}][%d] checkSizeStatusOK  %+v", 200, o.Payload)
}

func (o *CheckSizeStatusOK) GetPayload() *models.DataSize {
	return o.Payload
}

func (o *CheckSizeStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DataSize)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckSizeStatusBadRequest creates a CheckSizeStatusBadRequest with default headers values
func NewCheckSizeStatusBadRequest() *CheckSizeStatusBadRequest {
	return &CheckSizeStatusBadRequest{}
}

/*CheckSizeStatusBadRequest handles this case with default header values.

This means that there's something wrong in the input parameters and the server couldn't understand the request.
*/
type CheckSizeStatusBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *CheckSizeStatusBadRequest) Error() string {
	return fmt.Sprintf("[GET /dataset/data/size/{request_id}][%d] checkSizeStatusBadRequest  %+v", 400, o.Payload)
}

func (o *CheckSizeStatusBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CheckSizeStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckSizeStatusUnauthorized creates a CheckSizeStatusUnauthorized with default headers values
func NewCheckSizeStatusUnauthorized() *CheckSizeStatusUnauthorized {
	return &CheckSizeStatusUnauthorized{}
}

/*CheckSizeStatusUnauthorized handles this case with default header values.

This means that the user is not authenticated with keycloak and data transfer can't be triggered unless the user first log in with a valid user
*/
type CheckSizeStatusUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CheckSizeStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /dataset/data/size/{request_id}][%d] checkSizeStatusUnauthorized  %+v", 401, o.Payload)
}

func (o *CheckSizeStatusUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CheckSizeStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckSizeStatusForbidden creates a CheckSizeStatusForbidden with default headers values
func NewCheckSizeStatusForbidden() *CheckSizeStatusForbidden {
	return &CheckSizeStatusForbidden{}
}

/*CheckSizeStatusForbidden handles this case with default header values.

This means that the resource the user is trying to transfer from or to is not readable or writable by the user. User doesn't have the correct rights to either read the source file or write on the target system location.
*/
type CheckSizeStatusForbidden struct {
	Payload *models.ErrorResponse
}

func (o *CheckSizeStatusForbidden) Error() string {
	return fmt.Sprintf("[GET /dataset/data/size/{request_id}][%d] checkSizeStatusForbidden  %+v", 403, o.Payload)
}

func (o *CheckSizeStatusForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CheckSizeStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckSizeStatusNotFound creates a CheckSizeStatusNotFound with default headers values
func NewCheckSizeStatusNotFound() *CheckSizeStatusNotFound {
	return &CheckSizeStatusNotFound{}
}

/*CheckSizeStatusNotFound handles this case with default header values.

This means that either the source path or the the target path on the system doesn't exist.
*/
type CheckSizeStatusNotFound struct {
	Payload *models.ErrorResponse
}

func (o *CheckSizeStatusNotFound) Error() string {
	return fmt.Sprintf("[GET /dataset/data/size/{request_id}][%d] checkSizeStatusNotFound  %+v", 404, o.Payload)
}

func (o *CheckSizeStatusNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CheckSizeStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckSizeStatusRequestURITooLong creates a CheckSizeStatusRequestURITooLong with default headers values
func NewCheckSizeStatusRequestURITooLong() *CheckSizeStatusRequestURITooLong {
	return &CheckSizeStatusRequestURITooLong{}
}

/*CheckSizeStatusRequestURITooLong handles this case with default header values.

This means that the either the source path or the target path is longer than the server is willing to interpret.
*/
type CheckSizeStatusRequestURITooLong struct {
	Payload *models.ErrorResponse
}

func (o *CheckSizeStatusRequestURITooLong) Error() string {
	return fmt.Sprintf("[GET /dataset/data/size/{request_id}][%d] checkSizeStatusRequestUriTooLong  %+v", 414, o.Payload)
}

func (o *CheckSizeStatusRequestURITooLong) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CheckSizeStatusRequestURITooLong) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckSizeStatusTooManyRequests creates a CheckSizeStatusTooManyRequests with default headers values
func NewCheckSizeStatusTooManyRequests() *CheckSizeStatusTooManyRequests {
	return &CheckSizeStatusTooManyRequests{}
}

/*CheckSizeStatusTooManyRequests handles this case with default header values.

This means that the user has sent too many requests in a given amount of time. The user is advised to wait and send the request again.
*/
type CheckSizeStatusTooManyRequests struct {
	Payload *models.ErrorResponse
}

func (o *CheckSizeStatusTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /dataset/data/size/{request_id}][%d] checkSizeStatusTooManyRequests  %+v", 429, o.Payload)
}

func (o *CheckSizeStatusTooManyRequests) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CheckSizeStatusTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckSizeStatusInternalServerError creates a CheckSizeStatusInternalServerError with default headers values
func NewCheckSizeStatusInternalServerError() *CheckSizeStatusInternalServerError {
	return &CheckSizeStatusInternalServerError{}
}

/*CheckSizeStatusInternalServerError handles this case with default header values.

This means that something has gone wrong on the data transfer steering server. The user is advised to wait and send the request again.
*/
type CheckSizeStatusInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *CheckSizeStatusInternalServerError) Error() string {
	return fmt.Sprintf("[GET /dataset/data/size/{request_id}][%d] checkSizeStatusInternalServerError  %+v", 500, o.Payload)
}

func (o *CheckSizeStatusInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CheckSizeStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}