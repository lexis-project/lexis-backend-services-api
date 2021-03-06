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

// GetProjectReader is a Reader for the GetProject structure.
type GetProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetProjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetProjectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetProjectOK creates a GetProjectOK with default headers values
func NewGetProjectOK() *GetProjectOK {
	return &GetProjectOK{}
}

/*GetProjectOK handles this case with default header values.

project returned
*/
type GetProjectOK struct {
	Payload *models.Project
}

func (o *GetProjectOK) Error() string {
	return fmt.Sprintf("[GET /project/{id}][%d] getProjectOK  %+v", 200, o.Payload)
}

func (o *GetProjectOK) GetPayload() *models.Project {
	return o.Payload
}

func (o *GetProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Project)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectUnauthorized creates a GetProjectUnauthorized with default headers values
func NewGetProjectUnauthorized() *GetProjectUnauthorized {
	return &GetProjectUnauthorized{}
}

/*GetProjectUnauthorized handles this case with default header values.

Unauthorized
*/
type GetProjectUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *GetProjectUnauthorized) Error() string {
	return fmt.Sprintf("[GET /project/{id}][%d] getProjectUnauthorized  %+v", 401, o.Payload)
}

func (o *GetProjectUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetProjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectForbidden creates a GetProjectForbidden with default headers values
func NewGetProjectForbidden() *GetProjectForbidden {
	return &GetProjectForbidden{}
}

/*GetProjectForbidden handles this case with default header values.

Forbidden
*/
type GetProjectForbidden struct {
	Payload *models.ErrorResponse
}

func (o *GetProjectForbidden) Error() string {
	return fmt.Sprintf("[GET /project/{id}][%d] getProjectForbidden  %+v", 403, o.Payload)
}

func (o *GetProjectForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectNotFound creates a GetProjectNotFound with default headers values
func NewGetProjectNotFound() *GetProjectNotFound {
	return &GetProjectNotFound{}
}

/*GetProjectNotFound handles this case with default header values.

project with not found
*/
type GetProjectNotFound struct {
	Payload *models.ErrorResponse
}

func (o *GetProjectNotFound) Error() string {
	return fmt.Sprintf("[GET /project/{id}][%d] getProjectNotFound  %+v", 404, o.Payload)
}

func (o *GetProjectNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectInternalServerError creates a GetProjectInternalServerError with default headers values
func NewGetProjectInternalServerError() *GetProjectInternalServerError {
	return &GetProjectInternalServerError{}
}

/*GetProjectInternalServerError handles this case with default header values.

unexpected error
*/
type GetProjectInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *GetProjectInternalServerError) Error() string {
	return fmt.Sprintf("[GET /project/{id}][%d] getProjectInternalServerError  %+v", 500, o.Payload)
}

func (o *GetProjectInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetProjectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
