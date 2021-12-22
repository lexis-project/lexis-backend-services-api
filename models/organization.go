// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Organization organization
//
// swagger:model Organization
type Organization struct {

	// created by
	// Format: uuid
	CreatedBy strfmt.UUID `json:"CreatedBy,omitempty"`

	// creation date
	// Format: date-time
	CreationDate strfmt.DateTime `json:"CreationDate,omitempty"`

	// formal name
	FormalName string `json:"FormalName,omitempty"`

	// ID
	// Format: uuid
	ID strfmt.UUID `json:"ID,omitempty"`

	// organization email address
	// Format: email
	OrganizationEmailAddress strfmt.Email `json:"OrganizationEmailAddress,omitempty"`

	// organization status
	// Enum: [PENDING_APPROVAL APPROVED DISABLED TERMINATED]
	OrganizationStatus string `json:"OrganizationStatus,omitempty"`

	// primary telephone number
	PrimaryTelephoneNumber string `json:"PrimaryTelephoneNumber,omitempty"`

	// registered address1
	RegisteredAddress1 string `json:"RegisteredAddress1,omitempty"`

	// registered address2
	RegisteredAddress2 string `json:"RegisteredAddress2,omitempty"`

	// registered address3
	RegisteredAddress3 string `json:"RegisteredAddress3,omitempty"`

	// registered country
	RegisteredCountry string `json:"RegisteredCountry,omitempty"`

	// v a t registration number
	VATRegistrationNumber string `json:"VATRegistrationNumber,omitempty"`

	// website
	Website string `json:"Website,omitempty"`
}

// Validate validates this organization
func (m *Organization) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganizationEmailAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganizationStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Organization) validateCreatedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedBy) { // not required
		return nil
	}

	if err := validate.FormatOf("CreatedBy", "body", "uuid", m.CreatedBy.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Organization) validateCreationDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CreationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("CreationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Organization) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("ID", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Organization) validateOrganizationEmailAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.OrganizationEmailAddress) { // not required
		return nil
	}

	if err := validate.FormatOf("OrganizationEmailAddress", "body", "email", m.OrganizationEmailAddress.String(), formats); err != nil {
		return err
	}

	return nil
}

var organizationTypeOrganizationStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PENDING_APPROVAL","APPROVED","DISABLED","TERMINATED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		organizationTypeOrganizationStatusPropEnum = append(organizationTypeOrganizationStatusPropEnum, v)
	}
}

const (

	// OrganizationOrganizationStatusPENDINGAPPROVAL captures enum value "PENDING_APPROVAL"
	OrganizationOrganizationStatusPENDINGAPPROVAL string = "PENDING_APPROVAL"

	// OrganizationOrganizationStatusAPPROVED captures enum value "APPROVED"
	OrganizationOrganizationStatusAPPROVED string = "APPROVED"

	// OrganizationOrganizationStatusDISABLED captures enum value "DISABLED"
	OrganizationOrganizationStatusDISABLED string = "DISABLED"

	// OrganizationOrganizationStatusTERMINATED captures enum value "TERMINATED"
	OrganizationOrganizationStatusTERMINATED string = "TERMINATED"
)

// prop value enum
func (m *Organization) validateOrganizationStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, organizationTypeOrganizationStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Organization) validateOrganizationStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.OrganizationStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateOrganizationStatusEnum("OrganizationStatus", "body", m.OrganizationStatus); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Organization) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Organization) UnmarshalBinary(b []byte) error {
	var res Organization
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
