// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// WebsiteCheckPoint website check point
// swagger:model WebsiteCheckPoint
type WebsiteCheckPoint struct {

	// geo info
	// Read Only: true
	GeoInfo string `json:"geoInfo,omitempty"`

	// id
	// Read Only: true
	ID int32 `json:"id,omitempty"`

	// smg Id
	// Read Only: true
	SmgID int32 `json:"smgId,omitempty"`
}

// Validate validates this website check point
func (m *WebsiteCheckPoint) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WebsiteCheckPoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebsiteCheckPoint) UnmarshalBinary(b []byte) error {
	var res WebsiteCheckPoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}