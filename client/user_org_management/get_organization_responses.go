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

// GetOrganizationReader is a Reader for the GetOrganization structure.
type GetOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetOrganizationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOrganizationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOrganizationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOrganizationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrganizationOK creates a GetOrganizationOK with default headers values
func NewGetOrganizationOK() *GetOrganizationOK {
	return &GetOrganizationOK{}
}

/*GetOrganizationOK handles this case with default header values.

organization returned
*/
type GetOrganizationOK struct {
	Payload *models.Organization
}

func (o *GetOrganizationOK) Error() string {
	return fmt.Sprintf("[GET /organization/{id}][%d] getOrganizationOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationOK) GetPayload() *models.Organization {
	return o.Payload
}

func (o *GetOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Organization)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationUnauthorized creates a GetOrganizationUnauthorized with default headers values
func NewGetOrganizationUnauthorized() *GetOrganizationUnauthorized {
	return &GetOrganizationUnauthorized{}
}

/*GetOrganizationUnauthorized handles this case with default header values.

Unauthorized
*/
type GetOrganizationUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *GetOrganizationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /organization/{id}][%d] getOrganizationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOrganizationUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetOrganizationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationForbidden creates a GetOrganizationForbidden with default headers values
func NewGetOrganizationForbidden() *GetOrganizationForbidden {
	return &GetOrganizationForbidden{}
}

/*GetOrganizationForbidden handles this case with default header values.

Forbidden
*/
type GetOrganizationForbidden struct {
	Payload *models.ErrorResponse
}

func (o *GetOrganizationForbidden) Error() string {
	return fmt.Sprintf("[GET /organization/{id}][%d] getOrganizationForbidden  %+v", 403, o.Payload)
}

func (o *GetOrganizationForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetOrganizationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationNotFound creates a GetOrganizationNotFound with default headers values
func NewGetOrganizationNotFound() *GetOrganizationNotFound {
	return &GetOrganizationNotFound{}
}

/*GetOrganizationNotFound handles this case with default header values.

organization with organizationId not found
*/
type GetOrganizationNotFound struct {
	Payload *models.ErrorResponse
}

func (o *GetOrganizationNotFound) Error() string {
	return fmt.Sprintf("[GET /organization/{id}][%d] getOrganizationNotFound  %+v", 404, o.Payload)
}

func (o *GetOrganizationNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetOrganizationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationInternalServerError creates a GetOrganizationInternalServerError with default headers values
func NewGetOrganizationInternalServerError() *GetOrganizationInternalServerError {
	return &GetOrganizationInternalServerError{}
}

/*GetOrganizationInternalServerError handles this case with default header values.

unexpected error
*/
type GetOrganizationInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *GetOrganizationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /organization/{id}][%d] getOrganizationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOrganizationInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetOrganizationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}