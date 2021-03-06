// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetServiceVersionsIDParams creates a new GetServiceVersionsIDParams object
// with the default values initialized.
func NewGetServiceVersionsIDParams() *GetServiceVersionsIDParams {
	var ()
	return &GetServiceVersionsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetServiceVersionsIDParamsWithTimeout creates a new GetServiceVersionsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetServiceVersionsIDParamsWithTimeout(timeout time.Duration) *GetServiceVersionsIDParams {
	var ()
	return &GetServiceVersionsIDParams{

		timeout: timeout,
	}
}

// NewGetServiceVersionsIDParamsWithContext creates a new GetServiceVersionsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetServiceVersionsIDParamsWithContext(ctx context.Context) *GetServiceVersionsIDParams {
	var ()
	return &GetServiceVersionsIDParams{

		Context: ctx,
	}
}

// NewGetServiceVersionsIDParamsWithHTTPClient creates a new GetServiceVersionsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetServiceVersionsIDParamsWithHTTPClient(client *http.Client) *GetServiceVersionsIDParams {
	var ()
	return &GetServiceVersionsIDParams{
		HTTPClient: client,
	}
}

/*GetServiceVersionsIDParams contains all the parameters to send to the API endpoint
for the get service versions ID operation typically these are written to a http.Request
*/
type GetServiceVersionsIDParams struct {

	/*ID*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get service versions ID params
func (o *GetServiceVersionsIDParams) WithTimeout(timeout time.Duration) *GetServiceVersionsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get service versions ID params
func (o *GetServiceVersionsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get service versions ID params
func (o *GetServiceVersionsIDParams) WithContext(ctx context.Context) *GetServiceVersionsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get service versions ID params
func (o *GetServiceVersionsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get service versions ID params
func (o *GetServiceVersionsIDParams) WithHTTPClient(client *http.Client) *GetServiceVersionsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get service versions ID params
func (o *GetServiceVersionsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get service versions ID params
func (o *GetServiceVersionsIDParams) WithID(id int64) *GetServiceVersionsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get service versions ID params
func (o *GetServiceVersionsIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetServiceVersionsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
