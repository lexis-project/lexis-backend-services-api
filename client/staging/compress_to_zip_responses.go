// Code generated by go-swagger; DO NOT EDIT.

package staging

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/lexis-project/lexis-backend-services-api.git/models"
)

// CompressToZipReader is a Reader for the CompressToZip structure.
type CompressToZipReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CompressToZipReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCompressToZipCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCompressToZipBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCompressToZipUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCompressToZipForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 502:
		result := NewCompressToZipBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewCompressToZipServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCompressToZipCreated creates a CompressToZipCreated with default headers values
func NewCompressToZipCreated() *CompressToZipCreated {
	return &CompressToZipCreated{}
}

/*CompressToZipCreated handles this case with default header values.

Request was added to the queue
*/
type CompressToZipCreated struct {
	Payload *models.SteeringRequestID
}

func (o *CompressToZipCreated) Error() string {
	return fmt.Sprintf("[POST /dataset/compress/zip][%d] compressToZipCreated  %+v", 201, o.Payload)
}

func (o *CompressToZipCreated) GetPayload() *models.SteeringRequestID {
	return o.Payload
}

func (o *CompressToZipCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SteeringRequestID)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompressToZipBadRequest creates a CompressToZipBadRequest with default headers values
func NewCompressToZipBadRequest() *CompressToZipBadRequest {
	return &CompressToZipBadRequest{}
}

/*CompressToZipBadRequest handles this case with default header values.

Malformed Request
*/
type CompressToZipBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *CompressToZipBadRequest) Error() string {
	return fmt.Sprintf("[POST /dataset/compress/zip][%d] compressToZipBadRequest  %+v", 400, o.Payload)
}

func (o *CompressToZipBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CompressToZipBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompressToZipUnauthorized creates a CompressToZipUnauthorized with default headers values
func NewCompressToZipUnauthorized() *CompressToZipUnauthorized {
	return &CompressToZipUnauthorized{}
}

/*CompressToZipUnauthorized handles this case with default header values.

Unauthorized
*/
type CompressToZipUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CompressToZipUnauthorized) Error() string {
	return fmt.Sprintf("[POST /dataset/compress/zip][%d] compressToZipUnauthorized  %+v", 401, o.Payload)
}

func (o *CompressToZipUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CompressToZipUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompressToZipForbidden creates a CompressToZipForbidden with default headers values
func NewCompressToZipForbidden() *CompressToZipForbidden {
	return &CompressToZipForbidden{}
}

/*CompressToZipForbidden handles this case with default header values.

Forbidden
*/
type CompressToZipForbidden struct {
	Payload *models.ErrorResponse
}

func (o *CompressToZipForbidden) Error() string {
	return fmt.Sprintf("[POST /dataset/compress/zip][%d] compressToZipForbidden  %+v", 403, o.Payload)
}

func (o *CompressToZipForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CompressToZipForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompressToZipBadGateway creates a CompressToZipBadGateway with default headers values
func NewCompressToZipBadGateway() *CompressToZipBadGateway {
	return &CompressToZipBadGateway{}
}

/*CompressToZipBadGateway handles this case with default header values.

Bad Gateway
*/
type CompressToZipBadGateway struct {
	Payload *models.ErrorResponse
}

func (o *CompressToZipBadGateway) Error() string {
	return fmt.Sprintf("[POST /dataset/compress/zip][%d] compressToZipBadGateway  %+v", 502, o.Payload)
}

func (o *CompressToZipBadGateway) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CompressToZipBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompressToZipServiceUnavailable creates a CompressToZipServiceUnavailable with default headers values
func NewCompressToZipServiceUnavailable() *CompressToZipServiceUnavailable {
	return &CompressToZipServiceUnavailable{}
}

/*CompressToZipServiceUnavailable handles this case with default header values.

Server error
*/
type CompressToZipServiceUnavailable struct {
	Payload *models.ErrorResponse
}

func (o *CompressToZipServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /dataset/compress/zip][%d] compressToZipServiceUnavailable  %+v", 503, o.Payload)
}

func (o *CompressToZipServiceUnavailable) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CompressToZipServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CompressToZipBody compress to zip body
swagger:model CompressToZipBody
*/
type CompressToZipBody struct {

	// size
	// Required: true
	Size *int64 `json:"size"`

	// source path
	// Required: true
	SourcePath *string `json:"source_path"`

	// source system
	// Required: true
	// Enum: [lrz_iRODS]
	SourceSystem *string `json:"source_system"`
}

// Validate validates this compress to zip body
func (o *CompressToZipBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSize(formats); err != nil {
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

func (o *CompressToZipBody) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("parameters"+"."+"size", "body", o.Size); err != nil {
		return err
	}

	return nil
}

func (o *CompressToZipBody) validateSourcePath(formats strfmt.Registry) error {

	if err := validate.Required("parameters"+"."+"source_path", "body", o.SourcePath); err != nil {
		return err
	}

	return nil
}

var compressToZipBodyTypeSourceSystemPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["lrz_iRODS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		compressToZipBodyTypeSourceSystemPropEnum = append(compressToZipBodyTypeSourceSystemPropEnum, v)
	}
}

const (

	// CompressToZipBodySourceSystemLrziRODS captures enum value "lrz_iRODS"
	CompressToZipBodySourceSystemLrziRODS string = "lrz_iRODS"
)

// prop value enum
func (o *CompressToZipBody) validateSourceSystemEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, compressToZipBodyTypeSourceSystemPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *CompressToZipBody) validateSourceSystem(formats strfmt.Registry) error {

	if err := validate.Required("parameters"+"."+"source_system", "body", o.SourceSystem); err != nil {
		return err
	}

	// value enum
	if err := o.validateSourceSystemEnum("parameters"+"."+"source_system", "body", *o.SourceSystem); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CompressToZipBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CompressToZipBody) UnmarshalBinary(b []byte) error {
	var res CompressToZipBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
