// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HeappeTemplateParameter heappe template parameter
//
// swagger:model HeappeTemplateParameter
type HeappeTemplateParameter struct {

	// description
	Description *string `json:"Description,omitempty"`

	// identifier
	Identifier *string `json:"Identifier,omitempty"`
}

// Validate validates this heappe template parameter
func (m *HeappeTemplateParameter) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HeappeTemplateParameter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HeappeTemplateParameter) UnmarshalBinary(b []byte) error {
	var res HeappeTemplateParameter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
