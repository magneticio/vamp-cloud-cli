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

// NewGetReleasesParams creates a new GetReleasesParams object
// with the default values initialized.
func NewGetReleasesParams() *GetReleasesParams {
	var ()
	return &GetReleasesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetReleasesParamsWithTimeout creates a new GetReleasesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetReleasesParamsWithTimeout(timeout time.Duration) *GetReleasesParams {
	var ()
	return &GetReleasesParams{

		timeout: timeout,
	}
}

// NewGetReleasesParamsWithContext creates a new GetReleasesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetReleasesParamsWithContext(ctx context.Context) *GetReleasesParams {
	var ()
	return &GetReleasesParams{

		Context: ctx,
	}
}

// NewGetReleasesParamsWithHTTPClient creates a new GetReleasesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetReleasesParamsWithHTTPClient(client *http.Client) *GetReleasesParams {
	var ()
	return &GetReleasesParams{
		HTTPClient: client,
	}
}

/*GetReleasesParams contains all the parameters to send to the API endpoint
for the get releases operation typically these are written to a http.Request
*/
type GetReleasesParams struct {

	/*ApplicationID
	  application ID

	*/
	ApplicationID *int64
	/*Count
	  desirable number of items in response

	*/
	Count *int64
	/*ServiceID
	  service ID

	*/
	ServiceID *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get releases params
func (o *GetReleasesParams) WithTimeout(timeout time.Duration) *GetReleasesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get releases params
func (o *GetReleasesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get releases params
func (o *GetReleasesParams) WithContext(ctx context.Context) *GetReleasesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get releases params
func (o *GetReleasesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get releases params
func (o *GetReleasesParams) WithHTTPClient(client *http.Client) *GetReleasesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get releases params
func (o *GetReleasesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationID adds the applicationID to the get releases params
func (o *GetReleasesParams) WithApplicationID(applicationID *int64) *GetReleasesParams {
	o.SetApplicationID(applicationID)
	return o
}

// SetApplicationID adds the applicationId to the get releases params
func (o *GetReleasesParams) SetApplicationID(applicationID *int64) {
	o.ApplicationID = applicationID
}

// WithCount adds the count to the get releases params
func (o *GetReleasesParams) WithCount(count *int64) *GetReleasesParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get releases params
func (o *GetReleasesParams) SetCount(count *int64) {
	o.Count = count
}

// WithServiceID adds the serviceID to the get releases params
func (o *GetReleasesParams) WithServiceID(serviceID *int64) *GetReleasesParams {
	o.SetServiceID(serviceID)
	return o
}

// SetServiceID adds the serviceId to the get releases params
func (o *GetReleasesParams) SetServiceID(serviceID *int64) {
	o.ServiceID = serviceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetReleasesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ApplicationID != nil {

		// query param applicationID
		var qrApplicationID int64
		if o.ApplicationID != nil {
			qrApplicationID = *o.ApplicationID
		}
		qApplicationID := swag.FormatInt64(qrApplicationID)
		if qApplicationID != "" {
			if err := r.SetQueryParam("applicationID", qApplicationID); err != nil {
				return err
			}
		}

	}

	if o.Count != nil {

		// query param count
		var qrCount int64
		if o.Count != nil {
			qrCount = *o.Count
		}
		qCount := swag.FormatInt64(qrCount)
		if qCount != "" {
			if err := r.SetQueryParam("count", qCount); err != nil {
				return err
			}
		}

	}

	if o.ServiceID != nil {

		// query param serviceID
		var qrServiceID int64
		if o.ServiceID != nil {
			qrServiceID = *o.ServiceID
		}
		qServiceID := swag.FormatInt64(qrServiceID)
		if qServiceID != "" {
			if err := r.SetQueryParam("serviceID", qServiceID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
