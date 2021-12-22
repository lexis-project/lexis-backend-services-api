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

// IncreaseCreditReader is a Reader for the IncreaseCredit structure.
type IncreaseCreditReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IncreaseCreditReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIncreaseCreditOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewIncreaseCreditUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewIncreaseCreditForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewIncreaseCreditNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewIncreaseCreditInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIncreaseCreditOK creates a IncreaseCreditOK with default headers values
func NewIncreaseCreditOK() *IncreaseCreditOK {
	return &IncreaseCreditOK{}
}

/*IncreaseCreditOK handles this case with default header values.

Credit status of the account with the provided id
*/
type IncreaseCreditOK struct {
	Payload *models.CreditStatus
}

func (o *IncreaseCreditOK) Error() string {
	return fmt.Sprintf("[POST /accounting/{id}/manage/{medium}/increase/{amount}][%d] increaseCreditOK  %+v", 200, o.Payload)
}

func (o *IncreaseCreditOK) GetPayload() *models.CreditStatus {
	return o.Payload
}

func (o *IncreaseCreditOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreditStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIncreaseCreditUnauthorized creates a IncreaseCreditUnauthorized with default headers values
func NewIncreaseCreditUnauthorized() *IncreaseCreditUnauthorized {
	return &IncreaseCreditUnauthorized{}
}

/*IncreaseCreditUnauthorized handles this case with default header values.

Unauthorized
*/
type IncreaseCreditUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *IncreaseCreditUnauthorized) Error() string {
	return fmt.Sprintf("[POST /accounting/{id}/manage/{medium}/increase/{amount}][%d] increaseCreditUnauthorized  %+v", 401, o.Payload)
}

func (o *IncreaseCreditUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *IncreaseCreditUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIncreaseCreditForbidden creates a IncreaseCreditForbidden with default headers values
func NewIncreaseCreditForbidden() *IncreaseCreditForbidden {
	return &IncreaseCreditForbidden{}
}

/*IncreaseCreditForbidden handles this case with default header values.

Forbidden
*/
type IncreaseCreditForbidden struct {
	Payload *models.ErrorResponse
}

func (o *IncreaseCreditForbidden) Error() string {
	return fmt.Sprintf("[POST /accounting/{id}/manage/{medium}/increase/{amount}][%d] increaseCreditForbidden  %+v", 403, o.Payload)
}

func (o *IncreaseCreditForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *IncreaseCreditForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIncreaseCreditNotFound creates a IncreaseCreditNotFound with default headers values
func NewIncreaseCreditNotFound() *IncreaseCreditNotFound {
	return &IncreaseCreditNotFound{}
}

/*IncreaseCreditNotFound handles this case with default header values.

The account with the id provided doesn't exist
*/
type IncreaseCreditNotFound struct {
	Payload *models.ErrorResponse
}

func (o *IncreaseCreditNotFound) Error() string {
	return fmt.Sprintf("[POST /accounting/{id}/manage/{medium}/increase/{amount}][%d] increaseCreditNotFound  %+v", 404, o.Payload)
}

func (o *IncreaseCreditNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *IncreaseCreditNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIncreaseCreditInternalServerError creates a IncreaseCreditInternalServerError with default headers values
func NewIncreaseCreditInternalServerError() *IncreaseCreditInternalServerError {
	return &IncreaseCreditInternalServerError{}
}

/*IncreaseCreditInternalServerError handles this case with default header values.

Something unexpected happend, error raised
*/
type IncreaseCreditInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *IncreaseCreditInternalServerError) Error() string {
	return fmt.Sprintf("[POST /accounting/{id}/manage/{medium}/increase/{amount}][%d] increaseCreditInternalServerError  %+v", 500, o.Payload)
}

func (o *IncreaseCreditInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *IncreaseCreditInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}