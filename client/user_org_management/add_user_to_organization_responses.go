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

// AddUserToOrganizationReader is a Reader for the AddUserToOrganization structure.
type AddUserToOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddUserToOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddUserToOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAddUserToOrganizationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddUserToOrganizationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddUserToOrganizationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddUserToOrganizationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddUserToOrganizationOK creates a AddUserToOrganizationOK with default headers values
func NewAddUserToOrganizationOK() *AddUserToOrganizationOK {
	return &AddUserToOrganizationOK{}
}

/*AddUserToOrganizationOK handles this case with default header values.

user updated
*/
type AddUserToOrganizationOK struct {
	Payload *models.OKResponse
}

func (o *AddUserToOrganizationOK) Error() string {
	return fmt.Sprintf("[PUT /organization/{id}/user/{userID}][%d] addUserToOrganizationOK  %+v", 200, o.Payload)
}

func (o *AddUserToOrganizationOK) GetPayload() *models.OKResponse {
	return o.Payload
}

func (o *AddUserToOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OKResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddUserToOrganizationUnauthorized creates a AddUserToOrganizationUnauthorized with default headers values
func NewAddUserToOrganizationUnauthorized() *AddUserToOrganizationUnauthorized {
	return &AddUserToOrganizationUnauthorized{}
}

/*AddUserToOrganizationUnauthorized handles this case with default header values.

Authorization error
*/
type AddUserToOrganizationUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *AddUserToOrganizationUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /organization/{id}/user/{userID}][%d] addUserToOrganizationUnauthorized  %+v", 401, o.Payload)
}

func (o *AddUserToOrganizationUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddUserToOrganizationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddUserToOrganizationForbidden creates a AddUserToOrganizationForbidden with default headers values
func NewAddUserToOrganizationForbidden() *AddUserToOrganizationForbidden {
	return &AddUserToOrganizationForbidden{}
}

/*AddUserToOrganizationForbidden handles this case with default header values.

Authorization error
*/
type AddUserToOrganizationForbidden struct {
	Payload *models.ErrorResponse
}

func (o *AddUserToOrganizationForbidden) Error() string {
	return fmt.Sprintf("[PUT /organization/{id}/user/{userID}][%d] addUserToOrganizationForbidden  %+v", 403, o.Payload)
}

func (o *AddUserToOrganizationForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddUserToOrganizationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddUserToOrganizationNotFound creates a AddUserToOrganizationNotFound with default headers values
func NewAddUserToOrganizationNotFound() *AddUserToOrganizationNotFound {
	return &AddUserToOrganizationNotFound{}
}

/*AddUserToOrganizationNotFound handles this case with default header values.

organization with organizationId not found
*/
type AddUserToOrganizationNotFound struct {
	Payload *models.ErrorResponse
}

func (o *AddUserToOrganizationNotFound) Error() string {
	return fmt.Sprintf("[PUT /organization/{id}/user/{userID}][%d] addUserToOrganizationNotFound  %+v", 404, o.Payload)
}

func (o *AddUserToOrganizationNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddUserToOrganizationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddUserToOrganizationInternalServerError creates a AddUserToOrganizationInternalServerError with default headers values
func NewAddUserToOrganizationInternalServerError() *AddUserToOrganizationInternalServerError {
	return &AddUserToOrganizationInternalServerError{}
}

/*AddUserToOrganizationInternalServerError handles this case with default header values.

unexpected error
*/
type AddUserToOrganizationInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *AddUserToOrganizationInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /organization/{id}/user/{userID}][%d] addUserToOrganizationInternalServerError  %+v", 500, o.Payload)
}

func (o *AddUserToOrganizationInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddUserToOrganizationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
