// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GatewayRanConfigs RAN configuration for LTE gateway
// swagger:model gatewayRanConfigs
type GatewayRanConfigs struct {

	// pci
	// Required: true
	// Maximum: 503
	// Minimum: > 0
	Pci uint32 `json:"pci"`

	// transmit enabled
	// Required: true
	TransmitEnabled *bool `json:"transmit_enabled"`
}

// Validate validates this gateway ran configs
func (m *GatewayRanConfigs) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePci(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransmitEnabled(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GatewayRanConfigs) validatePci(formats strfmt.Registry) error {

	if err := validate.Required("pci", "body", uint32(m.Pci)); err != nil {
		return err
	}

	if err := validate.MinimumInt("pci", "body", int64(m.Pci), 0, true); err != nil {
		return err
	}

	if err := validate.MaximumInt("pci", "body", int64(m.Pci), 503, false); err != nil {
		return err
	}

	return nil
}

func (m *GatewayRanConfigs) validateTransmitEnabled(formats strfmt.Registry) error {

	if err := validate.Required("transmit_enabled", "body", m.TransmitEnabled); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GatewayRanConfigs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GatewayRanConfigs) UnmarshalBinary(b []byte) error {
	var res GatewayRanConfigs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
