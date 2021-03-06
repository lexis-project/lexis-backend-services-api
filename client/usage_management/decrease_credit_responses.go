// Code generated by go-swagger; DO NOT EDIT.

package usage_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/lexis-project/lexis-backend-services-api.git/models"
)

// DecreaseCreditReader is a Reader for the DecreaseCredit structure.
type DecreaseCreditReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DecreaseCreditReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDecreaseCreditOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDecreaseCreditUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDecreaseCreditForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDecreaseCreditNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDecreaseCreditInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDecreaseCreditOK creates a DecreaseCreditOK with default headers values
func NewDecreaseCreditOK() *DecreaseCreditOK {
	return &DecreaseCreditOK{}
}

/*DecreaseCreditOK handles this case with default header values.

Credit status of the account with the provided id
*/
type DecreaseCreditOK struct {
	Payload *models.CreditStatus
}

func (o *DecreaseCreditOK) Error() string {
	return fmt.Sprintf("[POST /accounting/{id}/manage/{medium}/decrease/{amount}][%d] decreaseCreditOK  %+v", 200, o.Payload)
}

func (o *DecreaseCreditOK) GetPayload() *models.CreditStatus {
	return o.Payload
}

func (o *DecreaseCreditOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreditStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDecreaseCreditUnauthorized creates a DecreaseCreditUnauthorized with default headers values
func NewDecreaseCreditUnauthorized() *DecreaseCreditUnauthorized {
	return &DecreaseCreditUnauthorized{}
}

/*DecreaseCreditUnauthorized handles this case with default header values.

Unauthorized
*/
type DecreaseCreditUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *DecreaseCreditUnauthorized) Error() string {
	return fmt.Sprintf("[POST /accounting/{id}/manage/{medium}/decrease/{amount}][%d] decreaseCreditUnauthorized  %+v", 401, o.Payload)
}

func (o *DecreaseCreditUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DecreaseCreditUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDecreaseCreditForbidden creates a DecreaseCreditForbidden with default headers values
func NewDecreaseCreditForbidden() *DecreaseCreditForbidden {
	return &DecreaseCreditForbidden{}
}

/*DecreaseCreditForbidden handles this case with default header values.

Forbidden
*/
type DecreaseCreditForbidden struct {
	Payload *models.ErrorResponse
}

func (o *DecreaseCreditForbidden) Error() string {
	return fmt.Sprintf("[POST /accounting/{id}/manage/{medium}/decrease/{amount}][%d] decreaseCreditForbidden  %+v", 403, o.Payload)
}

func (o *DecreaseCreditForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DecreaseCreditForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDecreaseCreditNotFound creates a DecreaseCreditNotFound with default headers values
func NewDecreaseCreditNotFound() *DecreaseCreditNotFound {
	return &DecreaseCreditNotFound{}
}

/*DecreaseCreditNotFound handles this case with default header values.

The account with the id provided doesn't exist
*/
type DecreaseCreditNotFound struct {
	Payload *models.ErrorResponse
}

func (o *DecreaseCreditNotFound) Error() string {
	return fmt.Sprintf("[POST /accounting/{id}/manage/{medium}/decrease/{amount}][%d] decreaseCreditNotFound  %+v", 404, o.Payload)
}

func (o *DecreaseCreditNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DecreaseCreditNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDecreaseCreditInternalServerError creates a DecreaseCreditInternalServerError with default headers values
func NewDecreaseCreditInternalServerError() *DecreaseCreditInternalServerError {
	return &DecreaseCreditInternalServerError{}
}

/*DecreaseCreditInternalServerError handles this case with default header values.

Something unexpected happend, error raised
*/
type DecreaseCreditInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *DecreaseCreditInternalServerError) Error() string {
	return fmt.Sprintf("[POST /accounting/{id}/manage/{medium}/decrease/{amount}][%d] decreaseCreditInternalServerError  %+v", 500, o.Payload)
}

func (o *DecreaseCreditInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DecreaseCreditInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
