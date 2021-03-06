// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DatasetFlags dataset flags
//
// swagger:model DatasetFlags
type DatasetFlags struct {

	// compression
	Compression string `json:"compression,omitempty"`

	// encryption
	Encryption string `json:"encryption,omitempty"`
}

// Validate validates this dataset flags
func (m *DatasetFlags) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DatasetFlags) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatasetFlags) UnmarshalBinary(b []byte) error {
	var res DatasetFlags
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
