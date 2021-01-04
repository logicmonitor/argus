// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// EC2NetscanPolicyCredential e c2 netscan policy credential
// swagger:model EC2NetscanPolicyCredential
type EC2NetscanPolicyCredential struct {

	// Custom credentials that should be used for this scan
	Custom map[string]string `json:"custom,omitempty"`

	// The ID of the device group that credentials should be inherited from, for this scan
	DeviceGroupID int32 `json:"deviceGroupId,omitempty"`

	// The name of the device group that credentials should be inherited from, for this scan
	DeviceGroupName string `json:"deviceGroupName,omitempty"`
}

// Validate validates this e c2 netscan policy credential
func (m *EC2NetscanPolicyCredential) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EC2NetscanPolicyCredential) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EC2NetscanPolicyCredential) UnmarshalBinary(b []byte) error {
	var res EC2NetscanPolicyCredential
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}