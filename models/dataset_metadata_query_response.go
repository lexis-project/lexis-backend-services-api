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

// DatasetMetadataQueryResponse dataset metadata query response
//
// swagger:model DatasetMetadataQueryResponse
type DatasetMetadataQueryResponse struct {

	// eudat
	Eudat *Eudat `json:"eudat,omitempty"`

	// flags
	// Required: true
	Flags *DatasetFlags `json:"flags"`

	// location
	// Required: true
	Location *Location `json:"location"`

	// metadata
	// Required: true
	Metadata *DatasetMetadata `json:"metadata"`
}

// Validate validates this dataset metadata query response
func (m *DatasetMetadataQueryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEudat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatasetMetadataQueryResponse) validateEudat(formats strfmt.Registry) error {

	if swag.IsZero(m.Eudat) { // not required
		return nil
	}

	if m.Eudat != nil {
		if err := m.Eudat.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eudat")
			}
			return err
		}
	}

	return nil
}

func (m *DatasetMetadataQueryResponse) validateFlags(formats strfmt.Registry) error {

	if err := validate.Required("flags", "body", m.Flags); err != nil {
		return err
	}

	if m.Flags != nil {
		if err := m.Flags.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flags")
			}
			return err
		}
	}

	return nil
}

func (m *DatasetMetadataQueryResponse) validateLocation(formats strfmt.Registry) error {

	if err := validate.Required("location", "body", m.Location); err != nil {
		return err
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

func (m *DatasetMetadataQueryResponse) validateMetadata(formats strfmt.Registry) error {

	if err := validate.Required("metadata", "body", m.Metadata); err != nil {
		return err
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DatasetMetadataQueryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatasetMetadataQueryResponse) UnmarshalBinary(b []byte) error {
	var res DatasetMetadataQueryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
