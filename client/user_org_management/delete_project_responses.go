// Code generated by go-swagger; DO NOT EDIT.

package user_org_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/lexis-project/lexis-backend-services-api.git/models"
)

// DeleteProjectReader is a Reader for the DeleteProject structure.
type DeleteProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteProjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteProjectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteProjectOK creates a DeleteProjectOK with default headers values
func NewDeleteProjectOK() *DeleteProjectOK {
	return &DeleteProjectOK{}
}

/*DeleteProjectOK handles this case with default header values.

deleted project
*/
type DeleteProjectOK struct {
	Payload *models.DeletedResponse
}

func (o *DeleteProjectOK) Error() string {
	return fmt.Sprintf("[DELETE /project/{id}][%d] deleteProjectOK  %+v", 200, o.Payload)
}

func (o *DeleteProjectOK) GetPayload() *models.DeletedResponse {
	return o.Payload
}

func (o *DeleteProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeletedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProjectUnauthorized creates a DeleteProjectUnauthorized with default headers values
func NewDeleteProjectUnauthorized() *DeleteProjectUnauthorized {
	return &DeleteProjectUnauthorized{}
}

/*DeleteProjectUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteProjectUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *DeleteProjectUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /project/{id}][%d] deleteProjectUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteProjectUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteProjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProjectForbidden creates a DeleteProjectForbidden with default headers values
func NewDeleteProjectForbidden() *DeleteProjectForbidden {
	return &DeleteProjectForbidden{}
}

/*DeleteProjectForbidden handles this case with default header values.

Forbidden
*/
type DeleteProjectForbidden struct {
	Payload *models.ErrorResponse
}

func (o *DeleteProjectForbidden) Error() string {
	return fmt.Sprintf("[DELETE /project/{id}][%d] deleteProjectForbidden  %+v", 403, o.Payload)
}

func (o *DeleteProjectForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProjectNotFound creates a DeleteProjectNotFound with default headers values
func NewDeleteProjectNotFound() *DeleteProjectNotFound {
	return &DeleteProjectNotFound{}
}

/*DeleteProjectNotFound handles this case with default header values.

project with not found
*/
type DeleteProjectNotFound struct {
	Payload *models.ErrorResponse
}

func (o *DeleteProjectNotFound) Error() string {
	return fmt.Sprintf("[DELETE /project/{id}][%d] deleteProjectNotFound  %+v", 404, o.Payload)
}

func (o *DeleteProjectNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProjectInternalServerError creates a DeleteProjectInternalServerError with default headers values
func NewDeleteProjectInternalServerError() *DeleteProjectInternalServerError {
	return &DeleteProjectInternalServerError{}
}

/*DeleteProjectInternalServerError handles this case with default header values.

unexpected error
*/
type DeleteProjectInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *DeleteProjectInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /project/{id}][%d] deleteProjectInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteProjectInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteProjectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
