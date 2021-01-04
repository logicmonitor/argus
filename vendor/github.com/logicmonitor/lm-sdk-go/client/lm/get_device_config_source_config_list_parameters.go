// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetDeviceConfigSourceConfigListParams creates a new GetDeviceConfigSourceConfigListParams object
// with the default values initialized.
func NewGetDeviceConfigSourceConfigListParams() *GetDeviceConfigSourceConfigListParams {
	var (
		offsetDefault = int32(0)
		sizeDefault   = int32(50)
	)
	return &GetDeviceConfigSourceConfigListParams{
		Offset: &offsetDefault,
		Size:   &sizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceConfigSourceConfigListParamsWithTimeout creates a new GetDeviceConfigSourceConfigListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDeviceConfigSourceConfigListParamsWithTimeout(timeout time.Duration) *GetDeviceConfigSourceConfigListParams {
	var (
		offsetDefault = int32(0)
		sizeDefault   = int32(50)
	)
	return &GetDeviceConfigSourceConfigListParams{
		Offset: &offsetDefault,
		Size:   &sizeDefault,

		timeout: timeout,
	}
}

// NewGetDeviceConfigSourceConfigListParamsWithContext creates a new GetDeviceConfigSourceConfigListParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDeviceConfigSourceConfigListParamsWithContext(ctx context.Context) *GetDeviceConfigSourceConfigListParams {
	var (
		offsetDefault = int32(0)
		sizeDefault   = int32(50)
	)
	return &GetDeviceConfigSourceConfigListParams{
		Offset: &offsetDefault,
		Size:   &sizeDefault,

		Context: ctx,
	}
}

// NewGetDeviceConfigSourceConfigListParamsWithHTTPClient creates a new GetDeviceConfigSourceConfigListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDeviceConfigSourceConfigListParamsWithHTTPClient(client *http.Client) *GetDeviceConfigSourceConfigListParams {
	var (
		offsetDefault = int32(0)
		sizeDefault   = int32(50)
	)
	return &GetDeviceConfigSourceConfigListParams{
		Offset:     &offsetDefault,
		Size:       &sizeDefault,
		HTTPClient: client,
	}
}

/*GetDeviceConfigSourceConfigListParams contains all the parameters to send to the API endpoint
for the get device config source config list operation typically these are written to a http.Request
*/
type GetDeviceConfigSourceConfigListParams struct {

	/*DeviceID*/
	DeviceID int32
	/*Fields*/
	Fields *string
	/*Filter*/
	Filter *string
	/*HdsID*/
	HdsID int32
	/*InstanceID*/
	InstanceID int32
	/*Offset*/
	Offset *int32
	/*Size*/
	Size *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get device config source config list params
func (o *GetDeviceConfigSourceConfigListParams) WithTimeout(timeout time.Duration) *GetDeviceConfigSourceConfigListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device config source config list params
func (o *GetDeviceConfigSourceConfigListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device config source config list params
func (o *GetDeviceConfigSourceConfigListParams) WithContext(ctx context.Context) *GetDeviceConfigSourceConfigListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device config source config list params
func (o *GetDeviceConfigSourceConfigListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device config source config list params
func (o *GetDeviceConfigSourceConfigListParams) WithHTTPClient(client *http.Client) *GetDeviceConfigSourceConfigListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device config source config list params
func (o *GetDeviceConfigSourceConfigListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceID adds the deviceID to the get device config source config list params
func (o *GetDeviceConfigSourceConfigListParams) WithDeviceID(deviceID int32) *GetDeviceConfigSourceConfigListParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the get device config source config list params
func (o *GetDeviceConfigSourceConfigListParams) SetDeviceID(deviceID int32) {
	o.DeviceID = deviceID
}

// WithFields adds the fields to the get device config source config list params
func (o *GetDeviceConfigSourceConfigListParams) WithFields(fields *string) *GetDeviceConfigSourceConfigListParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get device config source config list params
func (o *GetDeviceConfigSourceConfigListParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFilter adds the filter to the get device config source config list params
func (o *GetDeviceConfigSourceConfigListParams) WithFilter(filter *string) *GetDeviceConfigSourceConfigListParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get device config source config list params
func (o *GetDeviceConfigSourceConfigListParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithHdsID adds the hdsID to the get device config source config list params
func (o *GetDeviceConfigSourceConfigListParams) WithHdsID(hdsID int32) *GetDeviceConfigSourceConfigListParams {
	o.SetHdsID(hdsID)
	return o
}

// SetHdsID adds the hdsId to the get device config source config list params
func (o *GetDeviceConfigSourceConfigListParams) SetHdsID(hdsID int32) {
	o.HdsID = hdsID
}

// WithInstanceID adds the instanceID to the get device config source config list params
func (o *GetDeviceConfigSourceConfigListParams) WithInstanceID(instanceID int32) *GetDeviceConfigSourceConfigListParams {
	o.SetInstanceID(instanceID)
	return o
}

// SetInstanceID adds the instanceId to the get device config source config list params
func (o *GetDeviceConfigSourceConfigListParams) SetInstanceID(instanceID int32) {
	o.InstanceID = instanceID
}

// WithOffset adds the offset to the get device config source config list params
func (o *GetDeviceConfigSourceConfigListParams) WithOffset(offset *int32) *GetDeviceConfigSourceConfigListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get device config source config list params
func (o *GetDeviceConfigSourceConfigListParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithSize adds the size to the get device config source config list params
func (o *GetDeviceConfigSourceConfigListParams) WithSize(size *int32) *GetDeviceConfigSourceConfigListParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get device config source config list params
func (o *GetDeviceConfigSourceConfigListParams) SetSize(size *int32) {
	o.Size = size
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceConfigSourceConfigListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param deviceId
	if err := r.SetPathParam("deviceId", swag.FormatInt32(o.DeviceID)); err != nil {
		return err
	}

	if o.Fields != nil {

		// query param fields
		var qrFields string
		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {
			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}

	}

	if o.Filter != nil {

		// query param filter
		var qrFilter string
		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {
			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}

	}

	// path param hdsId
	if err := r.SetPathParam("hdsId", swag.FormatInt32(o.HdsID)); err != nil {
		return err
	}

	// path param instanceId
	if err := r.SetPathParam("instanceId", swag.FormatInt32(o.InstanceID)); err != nil {
		return err
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int32
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt32(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.Size != nil {

		// query param size
		var qrSize int32
		if o.Size != nil {
			qrSize = *o.Size
		}
		qSize := swag.FormatInt32(qrSize)
		if qSize != "" {
			if err := r.SetQueryParam("size", qSize); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}