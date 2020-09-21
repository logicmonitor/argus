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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vkumbhar94/lm-sdk-go/models"
)

// NewAddCollectorParams creates a new AddCollectorParams object
// with the default values initialized.
func NewAddCollectorParams() *AddCollectorParams {
	var ()
	return &AddCollectorParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddCollectorParamsWithTimeout creates a new AddCollectorParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddCollectorParamsWithTimeout(timeout time.Duration) *AddCollectorParams {
	var ()
	return &AddCollectorParams{

		timeout: timeout,
	}
}

// NewAddCollectorParamsWithContext creates a new AddCollectorParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddCollectorParamsWithContext(ctx context.Context) *AddCollectorParams {
	var ()
	return &AddCollectorParams{

		Context: ctx,
	}
}

// NewAddCollectorParamsWithHTTPClient creates a new AddCollectorParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddCollectorParamsWithHTTPClient(client *http.Client) *AddCollectorParams {
	var ()
	return &AddCollectorParams{
		HTTPClient: client,
	}
}

/*AddCollectorParams contains all the parameters to send to the API endpoint
for the add collector operation typically these are written to a http.Request
*/
type AddCollectorParams struct {

	/*Body*/
	Body *models.Collector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add collector params
func (o *AddCollectorParams) WithTimeout(timeout time.Duration) *AddCollectorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add collector params
func (o *AddCollectorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add collector params
func (o *AddCollectorParams) WithContext(ctx context.Context) *AddCollectorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add collector params
func (o *AddCollectorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add collector params
func (o *AddCollectorParams) WithHTTPClient(client *http.Client) *AddCollectorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add collector params
func (o *AddCollectorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add collector params
func (o *AddCollectorParams) WithBody(body *models.Collector) *AddCollectorParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add collector params
func (o *AddCollectorParams) SetBody(body *models.Collector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddCollectorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}