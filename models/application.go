// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Application application
//
// swagger:model Application
type Application struct {

	// cluster ID
	ClusterID int64 `json:"clusterID,omitempty"`

	// id
	// Required: true
	// Read Only: true
	ID int64 `json:"id"`

	// is owner
	IsOwner bool `json:"isOwner,omitempty"`

	// metadata
	Metadata []*Metadata `json:"metadata"`

	// name
	// Min Length: 1
	Name string `json:"name,omitempty"`
}

// Validate validates this application
func (m *Application) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Application) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Application) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	for i := 0; i < len(m.Metadata); i++ {
		if swag.IsZero(m.Metadata[i]) { // not required
			continue
		}

		if m.Metadata[i] != nil {
			if err := m.Metadata[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("metadata" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Application) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", string(m.Name), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Application) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Application) UnmarshalBinary(b []byte) error {
	var res Application
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}