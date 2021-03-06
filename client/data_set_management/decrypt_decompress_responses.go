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

// DecryptDecompressReader is a Reader for the DecryptDecompress structure.
type DecryptDecompressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DecryptDecompressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDecryptDecompressCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDecryptDecompressBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDecryptDecompressUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDecryptDecompressForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDecryptDecompressNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 414:
		result := NewDecryptDecompressRequestURITooLong()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDecryptDecompressTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDecryptDecompressInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDecryptDecompressCreated creates a DecryptDecompressCreated with default headers values
func NewDecryptDecompressCreated() *DecryptDecompressCreated {
	return &DecryptDecompressCreated{}
}

/*DecryptDecompressCreated handles this case with default header values.

The response code means that the decryption and decompression has been initiated. Status of the operation can be checked by querying the status.
*/
type DecryptDecompressCreated struct {
	Payload *models.SteeringRequestID
}

func (o *DecryptDecompressCreated) Error() string {
	return fmt.Sprintf("[POST /dataset/encryption/decrypt_decompress][%d] decryptDecompressCreated  %+v", 201, o.Payload)
}

func (o *DecryptDecompressCreated) GetPayload() *models.SteeringRequestID {
	return o.Payload
}

func (o *DecryptDecompressCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SteeringRequestID)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDecryptDecompressBadRequest creates a DecryptDecompressBadRequest with default headers values
func NewDecryptDecompressBadRequest() *DecryptDecompressBadRequest {
	return &DecryptDecompressBadRequest{}
}

/*DecryptDecompressBadRequest handles this case with default header values.

This means that there's something wrong in the input parameters and the server couldn't understand the request.
*/
type DecryptDecompressBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *DecryptDecompressBadRequest) Error() string {
	return fmt.Sprintf("[POST /dataset/encryption/decrypt_decompress][%d] decryptDecompressBadRequest  %+v", 400, o.Payload)
}

func (o *DecryptDecompressBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DecryptDecompressBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDecryptDecompressUnauthorized creates a DecryptDecompressUnauthorized with default headers values
func NewDecryptDecompressUnauthorized() *DecryptDecompressUnauthorized {
	return &DecryptDecompressUnauthorized{}
}

/*DecryptDecompressUnauthorized handles this case with default header values.

This means that the user is not authenticated with keycloak decryption with decompression can't be triggered unless the user first log in with a valid user
*/
type DecryptDecompressUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *DecryptDecompressUnauthorized) Error() string {
	return fmt.Sprintf("[POST /dataset/encryption/decrypt_decompress][%d] decryptDecompressUnauthorized  %+v", 401, o.Payload)
}

func (o *DecryptDecompressUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DecryptDecompressUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDecryptDecompressForbidden creates a DecryptDecompressForbidden with default headers values
func NewDecryptDecompressForbidden() *DecryptDecompressForbidden {
	return &DecryptDecompressForbidden{}
}

/*DecryptDecompressForbidden handles this case with default header values.

This means that the resource the user is trying to decrypt and decompress from or to is not readable or writable by the user. User doesn't have the correct rights to read the source file.
*/
type DecryptDecompressForbidden struct {
	Payload *models.ErrorResponse
}

func (o *DecryptDecompressForbidden) Error() string {
	return fmt.Sprintf("[POST /dataset/encryption/decrypt_decompress][%d] decryptDecompressForbidden  %+v", 403, o.Payload)
}

func (o *DecryptDecompressForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DecryptDecompressForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDecryptDecompressNotFound creates a DecryptDecompressNotFound with default headers values
func NewDecryptDecompressNotFound() *DecryptDecompressNotFound {
	return &DecryptDecompressNotFound{}
}

/*DecryptDecompressNotFound handles this case with default header values.

This means that either the source path on the system doesn't exist.
*/
type DecryptDecompressNotFound struct {
	Payload *models.ErrorResponse
}

func (o *DecryptDecompressNotFound) Error() string {
	return fmt.Sprintf("[POST /dataset/encryption/decrypt_decompress][%d] decryptDecompressNotFound  %+v", 404, o.Payload)
}

func (o *DecryptDecompressNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DecryptDecompressNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDecryptDecompressRequestURITooLong creates a DecryptDecompressRequestURITooLong with default headers values
func NewDecryptDecompressRequestURITooLong() *DecryptDecompressRequestURITooLong {
	return &DecryptDecompressRequestURITooLong{}
}

/*DecryptDecompressRequestURITooLong handles this case with default header values.

This means that the either the source path is longer than the server is willing to interpret.
*/
type DecryptDecompressRequestURITooLong struct {
	Payload *models.ErrorResponse
}

func (o *DecryptDecompressRequestURITooLong) Error() string {
	return fmt.Sprintf("[POST /dataset/encryption/decrypt_decompress][%d] decryptDecompressRequestUriTooLong  %+v", 414, o.Payload)
}

func (o *DecryptDecompressRequestURITooLong) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DecryptDecompressRequestURITooLong) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDecryptDecompressTooManyRequests creates a DecryptDecompressTooManyRequests with default headers values
func NewDecryptDecompressTooManyRequests() *DecryptDecompressTooManyRequests {
	return &DecryptDecompressTooManyRequests{}
}

/*DecryptDecompressTooManyRequests handles this case with default header values.

This means that the user has sent too many requests in a given amount of time. The user is advised to wait and send the request again.
*/
type DecryptDecompressTooManyRequests struct {
	Payload *models.ErrorResponse
}

func (o *DecryptDecompressTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /dataset/encryption/decrypt_decompress][%d] decryptDecompressTooManyRequests  %+v", 429, o.Payload)
}

func (o *DecryptDecompressTooManyRequests) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DecryptDecompressTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDecryptDecompressInternalServerError creates a DecryptDecompressInternalServerError with default headers values
func NewDecryptDecompressInternalServerError() *DecryptDecompressInternalServerError {
	return &DecryptDecompressInternalServerError{}
}

/*DecryptDecompressInternalServerError handles this case with default header values.

This means that something has gone wrong on the burst buffer. The user is advised to wait and send the request again.
*/
type DecryptDecompressInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *DecryptDecompressInternalServerError) Error() string {
	return fmt.Sprintf("[POST /dataset/encryption/decrypt_decompress][%d] decryptDecompressInternalServerError  %+v", 500, o.Payload)
}

func (o *DecryptDecompressInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DecryptDecompressInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DecryptDecompressBody decrypt decompress body
swagger:model DecryptDecompressBody
*/
type DecryptDecompressBody struct {

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

// Validate validates this decrypt decompress body
func (o *DecryptDecompressBody) Validate(formats strfmt.Registry) error {
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

func (o *DecryptDecompressBody) validateProject(formats strfmt.Registry) error {

	if err := validate.Required("parameters"+"."+"project", "body", o.Project); err != nil {
		return err
	}

	return nil
}

func (o *DecryptDecompressBody) validateSourcePath(formats strfmt.Registry) error {

	if err := validate.Required("parameters"+"."+"source_path", "body", o.SourcePath); err != nil {
		return err
	}

	return nil
}

func (o *DecryptDecompressBody) validateSourceSystem(formats strfmt.Registry) error {

	if err := validate.Required("parameters"+"."+"source_system", "body", o.SourceSystem); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DecryptDecompressBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DecryptDecompressBody) UnmarshalBinary(b []byte) error {
	var res DecryptDecompressBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
