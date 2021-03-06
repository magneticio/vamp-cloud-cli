// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IngressInput ingress input
//
// swagger:model IngressInput
type IngressInput struct {

	// domain name
	// Min Length: 1
	DomainName string `json:"domainName,omitempty"`

	// routes
	Routes []*Route `json:"routes"`

	// tls secret name
	TLSSecretName string `json:"tlsSecretName,omitempty"`

	// tls type
	// Enum: [NO_TLS TLS_EDGE]
	TLSType string `json:"tlsType,omitempty"`
}

// Validate validates this ingress input
func (m *IngressInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDomainName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoutes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTLSType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IngressInput) validateDomainName(formats strfmt.Registry) error {

	if swag.IsZero(m.DomainName) { // not required
		return nil
	}

	if err := validate.MinLength("domainName", "body", string(m.DomainName), 1); err != nil {
		return err
	}

	return nil
}

func (m *IngressInput) validateRoutes(formats strfmt.Registry) error {

	if swag.IsZero(m.Routes) { // not required
		return nil
	}

	for i := 0; i < len(m.Routes); i++ {
		if swag.IsZero(m.Routes[i]) { // not required
			continue
		}

		if m.Routes[i] != nil {
			if err := m.Routes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("routes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var ingressInputTypeTLSTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NO_TLS","TLS_EDGE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ingressInputTypeTLSTypePropEnum = append(ingressInputTypeTLSTypePropEnum, v)
	}
}

const (

	// IngressInputTLSTypeNOTLS captures enum value "NO_TLS"
	IngressInputTLSTypeNOTLS string = "NO_TLS"

	// IngressInputTLSTypeTLSEDGE captures enum value "TLS_EDGE"
	IngressInputTLSTypeTLSEDGE string = "TLS_EDGE"
)

// prop value enum
func (m *IngressInput) validateTLSTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ingressInputTypeTLSTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IngressInput) validateTLSType(formats strfmt.Registry) error {

	if swag.IsZero(m.TLSType) { // not required
		return nil
	}

	// value enum
	if err := m.validateTLSTypeEnum("tlsType", "body", m.TLSType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IngressInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IngressInput) UnmarshalBinary(b []byte) error {
	var res IngressInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
