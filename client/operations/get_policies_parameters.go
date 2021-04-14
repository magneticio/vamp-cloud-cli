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
)

// NewGetPoliciesParams creates a new GetPoliciesParams object
// with the default values initialized.
func NewGetPoliciesParams() *GetPoliciesParams {

	return &GetPoliciesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPoliciesParamsWithTimeout creates a new GetPoliciesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPoliciesParamsWithTimeout(timeout time.Duration) *GetPoliciesParams {

	return &GetPoliciesParams{

		timeout: timeout,
	}
}

// NewGetPoliciesParamsWithContext creates a new GetPoliciesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPoliciesParamsWithContext(ctx context.Context) *GetPoliciesParams {

	return &GetPoliciesParams{

		Context: ctx,
	}
}

// NewGetPoliciesParamsWithHTTPClient creates a new GetPoliciesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPoliciesParamsWithHTTPClient(client *http.Client) *GetPoliciesParams {

	return &GetPoliciesParams{
		HTTPClient: client,
	}
}

/*GetPoliciesParams contains all the parameters to send to the API endpoint
for the get policies operation typically these are written to a http.Request
*/
type GetPoliciesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get policies params
func (o *GetPoliciesParams) WithTimeout(timeout time.Duration) *GetPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get policies params
func (o *GetPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get policies params
func (o *GetPoliciesParams) WithContext(ctx context.Context) *GetPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get policies params
func (o *GetPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get policies params
func (o *GetPoliciesParams) WithHTTPClient(client *http.Client) *GetPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get policies params
func (o *GetPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
