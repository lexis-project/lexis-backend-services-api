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

// CompressEncryptReader is a Reader for the CompressEncrypt structure.
type CompressEncryptReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CompressEncryptReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCompressEncryptCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCompressEncryptBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCompressEncryptUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCompressEncryptForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCompressEncryptNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 414:
		result := NewCompressEncryptRequestURITooLong()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCompressEncryptTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCompressEncryptInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCompressEncryptCreated creates a CompressEncryptCreated with default headers values
func NewCompressEncryptCreated() *CompressEncryptCreated {
	return &CompressEncryptCreated{}
}

/*CompressEncryptCreated handles this case with default header values.

The response code means that the compression and encryption has been initiated. Status of the operation can be checked by querying the status.
*/
type CompressEncryptCreated struct {
	Payload *models.SteeringRequestID
}

func (o *CompressEncryptCreated) Error() string {
	return fmt.Sprintf("[POST /dataset/encryption/compress_encrypt][%d] compressEncryptCreated  %+v", 201, o.Payload)
}

func (o *CompressEncryptCreated) GetPayload() *models.SteeringRequestID {
	return o.Payload
}

func (o *CompressEncryptCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SteeringRequestID)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompressEncryptBadRequest creates a CompressEncryptBadRequest with default headers values
func NewCompressEncryptBadRequest() *CompressEncryptBadRequest {
	return &CompressEncryptBadRequest{}
}

/*CompressEncryptBadRequest handles this case with default header values.

This means that there's something wrong in the input parameters and the server couldn't understand the request.
*/
type CompressEncryptBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *CompressEncryptBadRequest) Error() string {
	return fmt.Sprintf("[POST /dataset/encryption/compress_encrypt][%d] compressEncryptBadRequest  %+v", 400, o.Payload)
}

func (o *CompressEncryptBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CompressEncryptBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompressEncryptUnauthorized creates a CompressEncryptUnauthorized with default headers values
func NewCompressEncryptUnauthorized() *CompressEncryptUnauthorized {
	return &CompressEncryptUnauthorized{}
}

/*CompressEncryptUnauthorized handles this case with default header values.

This means that the user is not authenticated with keycloak and compression with encryption can't be triggered unless the user first log in with a valid user
*/
type CompressEncryptUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CompressEncryptUnauthorized) Error() string {
	return fmt.Sprintf("[POST /dataset/encryption/compress_encrypt][%d] compressEncryptUnauthorized  %+v", 401, o.Payload)
}

func (o *CompressEncryptUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CompressEncryptUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompressEncryptForbidden creates a CompressEncryptForbidden with default headers values
func NewCompressEncryptForbidden() *CompressEncryptForbidden {
	return &CompressEncryptForbidden{}
}

/*CompressEncryptForbidden handles this case with default header values.

This means that the resource the user is trying to compress and encrypt is not readable by the user. User doesn't have the correct rights to read the source file.
*/
type CompressEncryptForbidden struct {
	Payload *models.ErrorResponse
}

func (o *CompressEncryptForbidden) Error() string {
	return fmt.Sprintf("[POST /dataset/encryption/compress_encrypt][%d] compressEncryptForbidden  %+v", 403, o.Payload)
}

func (o *CompressEncryptForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CompressEncryptForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompressEncryptNotFound creates a CompressEncryptNotFound with default headers values
func NewCompressEncryptNotFound() *CompressEncryptNotFound {
	return &CompressEncryptNotFound{}
}

/*CompressEncryptNotFound handles this case with default header values.

This means that either the source path on the system doesn't exist.
*/
type CompressEncryptNotFound struct {
	Payload *models.ErrorResponse
}

func (o *CompressEncryptNotFound) Error() string {
	return fmt.Sprintf("[POST /dataset/encryption/compress_encrypt][%d] compressEncryptNotFound  %+v", 404, o.Payload)
}

func (o *CompressEncryptNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CompressEncryptNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompressEncryptRequestURITooLong creates a CompressEncryptRequestURITooLong with default headers values
func NewCompressEncryptRequestURITooLong() *CompressEncryptRequestURITooLong {
	return &CompressEncryptRequestURITooLong{}
}

/*CompressEncryptRequestURITooLong handles this case with default header values.

This means that the source path is longer than the server is willing to interpret.
*/
type CompressEncryptRequestURITooLong struct {
	Payload *models.ErrorResponse
}

func (o *CompressEncryptRequestURITooLong) Error() string {
	return fmt.Sprintf("[POST /dataset/encryption/compress_encrypt][%d] compressEncryptRequestUriTooLong  %+v", 414, o.Payload)
}

func (o *CompressEncryptRequestURITooLong) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CompressEncryptRequestURITooLong) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompressEncryptTooManyRequests creates a CompressEncryptTooManyRequests with default headers values
func NewCompressEncryptTooManyRequests() *CompressEncryptTooManyRequests {
	return &CompressEncryptTooManyRequests{}
}

/*CompressEncryptTooManyRequests handles this case with default header values.

This means that the user has sent too many requests in a given amount of time. The user is advised to wait and send the request again.
*/
type CompressEncryptTooManyRequests struct {
	Payload *models.ErrorResponse
}

func (o *CompressEncryptTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /dataset/encryption/compress_encrypt][%d] compressEncryptTooManyRequests  %+v", 429, o.Payload)
}

func (o *CompressEncryptTooManyRequests) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CompressEncryptTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompressEncryptInternalServerError creates a CompressEncryptInternalServerError with default headers values
func NewCompressEncryptInternalServerError() *CompressEncryptInternalServerError {
	return &CompressEncryptInternalServerError{}
}

/*CompressEncryptInternalServerError handles this case with default header values.

This means that something has gone wrong on the burst buffer. The user is advised to wait and send the request again.
*/
type CompressEncryptInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *CompressEncryptInternalServerError) Error() string {
	return fmt.Sprintf("[POST /dataset/encryption/compress_encrypt][%d] compressEncryptInternalServerError  %+v", 500, o.Payload)
}

func (o *CompressEncryptInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CompressEncryptInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CompressEncryptBody compress encrypt body
swagger:model CompressEncryptBody
*/
type CompressEncryptBody struct {

	// project
	// Required: true
	Project *string `json:"project"`

	// source path
	// Required: true
	SourcePath *string `json:"source_path"`

	// source system
	// Required: true
	SourceSystem *string `json:"source_system"`
}

// Validate validates this compress encrypt body
func (o *CompressEncryptBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSourcePath(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSourceSystem(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CompressEncryptBody) validateProject(formats strfmt.Registry) error {

	if err := validate.Required("parameters"+"."+"project", "body", o.Project); err != nil {
		return err
	}

	return nil
}

func (o *CompressEncryptBody) validateSourcePath(formats strfmt.Registry) error {

	if err := validate.Required("parameters"+"."+"source_path", "body", o.SourcePath); err != nil {
		return err
	}

	return nil
}

func (o *CompressEncryptBody) validateSourceSystem(formats strfmt.Registry) error {

	if err := validate.Required("parameters"+"."+"source_system", "body", o.SourceSystem); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CompressEncryptBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CompressEncryptBody) UnmarshalBinary(b []byte) error {
	var res CompressEncryptBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
