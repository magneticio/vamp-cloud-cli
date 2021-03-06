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

	"github.com/magneticio/vamp-cloud-cli/models"
)

// NewPostApplicationsParams creates a new PostApplicationsParams object
// with the default values initialized.
func NewPostApplicationsParams() *PostApplicationsParams {
	var ()
	return &PostApplicationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostApplicationsParamsWithTimeout creates a new PostApplicationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostApplicationsParamsWithTimeout(timeout time.Duration) *PostApplicationsParams {
	var ()
	return &PostApplicationsParams{

		timeout: timeout,
	}
}

// NewPostApplicationsParamsWithContext creates a new PostApplicationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostApplicationsParamsWithContext(ctx context.Context) *PostApplicationsParams {
	var ()
	return &PostApplicationsParams{

		Context: ctx,
	}
}

// NewPostApplicationsParamsWithHTTPClient creates a new PostApplicationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostApplicationsParamsWithHTTPClient(client *http.Client) *PostApplicationsParams {
	var ()
	return &PostApplicationsParams{
		HTTPClient: client,
	}
}

/*PostApplicationsParams contains all the parameters to send to the API endpoint
for the post applications operation typically these are written to a http.Request
*/
type PostApplicationsParams struct {

	/*Application
	  The application to create

	*/
	Application *models.ApplicationInput

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post applications params
func (o *PostApplicationsParams) WithTimeout(timeout time.Duration) *PostApplicationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post applications params
func (o *PostApplicationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post applications params
func (o *PostApplicationsParams) WithContext(ctx context.Context) *PostApplicationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post applications params
func (o *PostApplicationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post applications params
func (o *PostApplicationsParams) WithHTTPClient(client *http.Client) *PostApplicationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post applications params
func (o *PostApplicationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplication adds the application to the post applications params
func (o *PostApplicationsParams) WithApplication(application *models.ApplicationInput) *PostApplicationsParams {
	o.SetApplication(application)
	return o
}

// SetApplication adds the application to the post applications params
func (o *PostApplicationsParams) SetApplication(application *models.ApplicationInput) {
	o.Application = application
}

// WriteToRequest writes these params to a swagger request
func (o *PostApplicationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Application != nil {
		if err := r.SetBodyParam(o.Application); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
