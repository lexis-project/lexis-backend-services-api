// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DataSize data size
//
// swagger:model DataSize
type DataSize struct {

	// Status of the datasize get process
	// Required: true
	Result *string `json:"result"`

	// dataset size
	Size string `json:"size,omitempty"`

	// dataset files smaller than 32 MB
	Smallfiles string `json:"smallfiles,omitempty"`

	// dataset files
	Totalfiles string `json:"totalfiles,omitempty"`
}

// Validate validates this data size
func (m *DataSize) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataSize) validateResult(formats strfmt.Registry) error {

	if err := validate.Required("result", "body", m.Result); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DataSize) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataSize) UnmarshalBinary(b []byte) error {
	var res DataSize
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
