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

// ApplicationInput application input
//
// swagger:model ApplicationInput
type ApplicationInput struct {

	// cluster ID
	ClusterID int64 `json:"clusterID,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// ingress type
	// Enum: [NGINX CONTOUR NONE]
	IngressType string `json:"ingressType,omitempty"`

	// name
	// Min Length: 1
	Name string `json:"name,omitempty"`

	// namespace
	// Min Length: 1
	Namespace string `json:"namespace,omitempty"`
}

// Validate validates this application input
func (m *ApplicationInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIngressType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var applicationInputTypeIngressTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NGINX","CONTOUR","NONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		applicationInputTypeIngressTypePropEnum = append(applicationInputTypeIngressTypePropEnum, v)
	}
}

const (

	// ApplicationInputIngressTypeNGINX captures enum value "NGINX"
	ApplicationInputIngressTypeNGINX string = "NGINX"

	// ApplicationInputIngressTypeCONTOUR captures enum value "CONTOUR"
	ApplicationInputIngressTypeCONTOUR string = "CONTOUR"

	// ApplicationInputIngressTypeNONE captures enum value "NONE"
	ApplicationInputIngressTypeNONE string = "NONE"
)

// prop value enum
func (m *ApplicationInput) validateIngressTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, applicationInputTypeIngressTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ApplicationInput) validateIngressType(formats strfmt.Registry) error {

	if swag.IsZero(m.IngressType) { // not required
		return nil
	}

	// value enum
	if err := m.validateIngressTypeEnum("ingressType", "body", m.IngressType); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationInput) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", string(m.Name), 1); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationInput) validateNamespace(formats strfmt.Registry) error {

	if swag.IsZero(m.Namespace) { // not required
		return nil
	}

	if err := validate.MinLength("namespace", "body", string(m.Namespace), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationInput) UnmarshalBinary(b []byte) error {
	var res ApplicationInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
