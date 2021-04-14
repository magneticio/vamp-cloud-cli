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

	"github.com/magneticio/vamp-cloud-cli/models"
)

// NewPutApplicationsApplicationIDServicesServiceIDParams creates a new PutApplicationsApplicationIDServicesServiceIDParams object
// with the default values initialized.
func NewPutApplicationsApplicationIDServicesServiceIDParams() *PutApplicationsApplicationIDServicesServiceIDParams {
	var ()
	return &PutApplicationsApplicationIDServicesServiceIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutApplicationsApplicationIDServicesServiceIDParamsWithTimeout creates a new PutApplicationsApplicationIDServicesServiceIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutApplicationsApplicationIDServicesServiceIDParamsWithTimeout(timeout time.Duration) *PutApplicationsApplicationIDServicesServiceIDParams {
	var ()
	return &PutApplicationsApplicationIDServicesServiceIDParams{

		timeout: timeout,
	}
}

// NewPutApplicationsApplicationIDServicesServiceIDParamsWithContext creates a new PutApplicationsApplicationIDServicesServiceIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutApplicationsApplicationIDServicesServiceIDParamsWithContext(ctx context.Context) *PutApplicationsApplicationIDServicesServiceIDParams {
	var ()
	return &PutApplicationsApplicationIDServicesServiceIDParams{

		Context: ctx,
	}
}

// NewPutApplicationsApplicationIDServicesServiceIDParamsWithHTTPClient creates a new PutApplicationsApplicationIDServicesServiceIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutApplicationsApplicationIDServicesServiceIDParamsWithHTTPClient(client *http.Client) *PutApplicationsApplicationIDServicesServiceIDParams {
	var ()
	return &PutApplicationsApplicationIDServicesServiceIDParams{
		HTTPClient: client,
	}
}

/*PutApplicationsApplicationIDServicesServiceIDParams contains all the parameters to send to the API endpoint
for the put applications application ID services service ID operation typically these are written to a http.Request
*/
type PutApplicationsApplicationIDServicesServiceIDParams struct {

	/*PolicySelectionStrategyInput
	  Policy selection strategy for service

	*/
	PolicySelectionStrategyInput *models.PolicySelectionStrategyInput
	/*ApplicationID
	  application ID

	*/
	ApplicationID int64
	/*ServiceID
	  application ID

	*/
	ServiceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put applications application ID services service ID params
func (o *PutApplicationsApplicationIDServicesServiceIDParams) WithTimeout(timeout time.Duration) *PutApplicationsApplicationIDServicesServiceIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put applications application ID services service ID params
func (o *PutApplicationsApplicationIDServicesServiceIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put applications application ID services service ID params
func (o *PutApplicationsApplicationIDServicesServiceIDParams) WithContext(ctx context.Context) *PutApplicationsApplicationIDServicesServiceIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put applications application ID services service ID params
func (o *PutApplicationsApplicationIDServicesServiceIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put applications application ID services service ID params
func (o *PutApplicationsApplicationIDServicesServiceIDParams) WithHTTPClient(client *http.Client) *PutApplicationsApplicationIDServicesServiceIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put applications application ID services service ID params
func (o *PutApplicationsApplicationIDServicesServiceIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPolicySelectionStrategyInput adds the policySelectionStrategyInput to the put applications application ID services service ID params
func (o *PutApplicationsApplicationIDServicesServiceIDParams) WithPolicySelectionStrategyInput(policySelectionStrategyInput *models.PolicySelectionStrategyInput) *PutApplicationsApplicationIDServicesServiceIDParams {
	o.SetPolicySelectionStrategyInput(policySelectionStrategyInput)
	return o
}

// SetPolicySelectionStrategyInput adds the policySelectionStrategyInput to the put applications application ID services service ID params
func (o *PutApplicationsApplicationIDServicesServiceIDParams) SetPolicySelectionStrategyInput(policySelectionStrategyInput *models.PolicySelectionStrategyInput) {
	o.PolicySelectionStrategyInput = policySelectionStrategyInput
}

// WithApplicationID adds the applicationID to the put applications application ID services service ID params
func (o *PutApplicationsApplicationIDServicesServiceIDParams) WithApplicationID(applicationID int64) *PutApplicationsApplicationIDServicesServiceIDParams {
	o.SetApplicationID(applicationID)
	return o
}

// SetApplicationID adds the applicationId to the put applications application ID services service ID params
func (o *PutApplicationsApplicationIDServicesServiceIDParams) SetApplicationID(applicationID int64) {
	o.ApplicationID = applicationID
}

// WithServiceID adds the serviceID to the put applications application ID services service ID params
func (o *PutApplicationsApplicationIDServicesServiceIDParams) WithServiceID(serviceID int64) *PutApplicationsApplicationIDServicesServiceIDParams {
	o.SetServiceID(serviceID)
	return o
}

// SetServiceID adds the serviceId to the put applications application ID services service ID params
func (o *PutApplicationsApplicationIDServicesServiceIDParams) SetServiceID(serviceID int64) {
	o.ServiceID = serviceID
}

// WriteToRequest writes these params to a swagger request
func (o *PutApplicationsApplicationIDServicesServiceIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PolicySelectionStrategyInput != nil {
		if err := r.SetBodyParam(o.PolicySelectionStrategyInput); err != nil {
			return err
		}
	}

	// path param applicationID
	if err := r.SetPathParam("applicationID", swag.FormatInt64(o.ApplicationID)); err != nil {
		return err
	}

	// path param serviceID
	if err := r.SetPathParam("serviceID", swag.FormatInt64(o.ServiceID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
