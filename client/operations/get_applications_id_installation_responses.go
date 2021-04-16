// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/magneticio/vamp-cloud-cli/models"
)

// GetApplicationsIDInstallationReader is a Reader for the GetApplicationsIDInstallation structure.
type GetApplicationsIDInstallationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApplicationsIDInstallationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApplicationsIDInstallationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetApplicationsIDInstallationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetApplicationsIDInstallationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetApplicationsIDInstallationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetApplicationsIDInstallationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetApplicationsIDInstallationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetApplicationsIDInstallationOK creates a GetApplicationsIDInstallationOK with default headers values
func NewGetApplicationsIDInstallationOK() *GetApplicationsIDInstallationOK {
	return &GetApplicationsIDInstallationOK{}
}

/*GetApplicationsIDInstallationOK handles this case with default header values.

Installation command
*/
type GetApplicationsIDInstallationOK struct {
	Payload *models.Installation
}

func (o *GetApplicationsIDInstallationOK) Error() string {
	return fmt.Sprintf("[GET /applications/{id}/installation][%d] getApplicationsIdInstallationOK  %+v", 200, o.Payload)
}

func (o *GetApplicationsIDInstallationOK) GetPayload() *models.Installation {
	return o.Payload
}

func (o *GetApplicationsIDInstallationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Installation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApplicationsIDInstallationUnauthorized creates a GetApplicationsIDInstallationUnauthorized with default headers values
func NewGetApplicationsIDInstallationUnauthorized() *GetApplicationsIDInstallationUnauthorized {
	return &GetApplicationsIDInstallationUnauthorized{}
}

/*GetApplicationsIDInstallationUnauthorized handles this case with default header values.

The requester is not authorized.
*/
type GetApplicationsIDInstallationUnauthorized struct {
	Payload *models.Error
}

func (o *GetApplicationsIDInstallationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /applications/{id}/installation][%d] getApplicationsIdInstallationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetApplicationsIDInstallationUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetApplicationsIDInstallationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApplicationsIDInstallationForbidden creates a GetApplicationsIDInstallationForbidden with default headers values
func NewGetApplicationsIDInstallationForbidden() *GetApplicationsIDInstallationForbidden {
	return &GetApplicationsIDInstallationForbidden{}
}

/*GetApplicationsIDInstallationForbidden handles this case with default header values.

The requester does not have access rights to the resource.
*/
type GetApplicationsIDInstallationForbidden struct {
	Payload *models.Error
}

func (o *GetApplicationsIDInstallationForbidden) Error() string {
	return fmt.Sprintf("[GET /applications/{id}/installation][%d] getApplicationsIdInstallationForbidden  %+v", 403, o.Payload)
}

func (o *GetApplicationsIDInstallationForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetApplicationsIDInstallationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApplicationsIDInstallationNotFound creates a GetApplicationsIDInstallationNotFound with default headers values
func NewGetApplicationsIDInstallationNotFound() *GetApplicationsIDInstallationNotFound {
	return &GetApplicationsIDInstallationNotFound{}
}

/*GetApplicationsIDInstallationNotFound handles this case with default header values.

The server can not find the requested resource.
*/
type GetApplicationsIDInstallationNotFound struct {
	Payload *models.Error
}

func (o *GetApplicationsIDInstallationNotFound) Error() string {
	return fmt.Sprintf("[GET /applications/{id}/installation][%d] getApplicationsIdInstallationNotFound  %+v", 404, o.Payload)
}

func (o *GetApplicationsIDInstallationNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetApplicationsIDInstallationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApplicationsIDInstallationInternalServerError creates a GetApplicationsIDInstallationInternalServerError with default headers values
func NewGetApplicationsIDInstallationInternalServerError() *GetApplicationsIDInstallationInternalServerError {
	return &GetApplicationsIDInstallationInternalServerError{}
}

/*GetApplicationsIDInstallationInternalServerError handles this case with default header values.

The server has encountered a situation it does not know how to handle.
*/
type GetApplicationsIDInstallationInternalServerError struct {
	Payload *models.Error
}

func (o *GetApplicationsIDInstallationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /applications/{id}/installation][%d] getApplicationsIdInstallationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetApplicationsIDInstallationInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetApplicationsIDInstallationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApplicationsIDInstallationDefault creates a GetApplicationsIDInstallationDefault with default headers values
func NewGetApplicationsIDInstallationDefault(code int) *GetApplicationsIDInstallationDefault {
	return &GetApplicationsIDInstallationDefault{
		_statusCode: code,
	}
}

/*GetApplicationsIDInstallationDefault handles this case with default header values.

Generic error.
*/
type GetApplicationsIDInstallationDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get applications ID installation default response
func (o *GetApplicationsIDInstallationDefault) Code() int {
	return o._statusCode
}

func (o *GetApplicationsIDInstallationDefault) Error() string {
	return fmt.Sprintf("[GET /applications/{id}/installation][%d] GetApplicationsIDInstallation default  %+v", o._statusCode, o.Payload)
}

func (o *GetApplicationsIDInstallationDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetApplicationsIDInstallationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
