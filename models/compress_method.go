// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// CompressMethod compress method
//
// swagger:model CompressMethod
type CompressMethod string

const (

	// CompressMethodZip captures enum value "zip"
	CompressMethodZip CompressMethod = "zip"

	// CompressMethodFile captures enum value "file"
	CompressMethodFile CompressMethod = "file"
)

// for schema
var compressMethodEnum []interface{}

func init() {
	var res []CompressMethod
	if err := json.Unmarshal([]byte(`["zip","file"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		compressMethodEnum = append(compressMethodEnum, v)
	}
}

func (m CompressMethod) validateCompressMethodEnum(path, location string, value CompressMethod) error {
	if err := validate.EnumCase(path, location, value, compressMethodEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this compress method
func (m CompressMethod) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCompressMethodEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
