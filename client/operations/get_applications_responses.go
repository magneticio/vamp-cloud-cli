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

// GetApplicationsReader is a Reader for the GetApplications structure.
type GetApplicationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApplicationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApplicationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetApplicationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetApplicationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetApplicationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetApplicationsOK creates a GetApplicationsOK with default headers values
func NewGetApplicationsOK() *GetApplicationsOK {
	return &GetApplicationsOK{}
}

/*GetApplicationsOK handles this case with default header values.

List of application objects
*/
type GetApplicationsOK struct {
	Payload *models.Applications
}

func (o *GetApplicationsOK) Error() string {
	return fmt.Sprintf("[GET /applications][%d] getApplicationsOK  %+v", 200, o.Payload)
}

func (o *GetApplicationsOK) GetPayload() *models.Applications {
	return o.Payload
}

func (o *GetApplicationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Applications)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApplicationsUnauthorized creates a GetApplicationsUnauthorized with default headers values
func NewGetApplicationsUnauthorized() *GetApplicationsUnauthorized {
	return &GetApplicationsUnauthorized{}
}

/*GetApplicationsUnauthorized handles this case with default header values.

The requester is not authorized.
*/
type GetApplicationsUnauthorized struct {
	Payload *models.Error
}

func (o *GetApplicationsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /applications][%d] getApplicationsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetApplicationsUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetApplicationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApplicationsInternalServerError creates a GetApplicationsInternalServerError with default headers values
func NewGetApplicationsInternalServerError() *GetApplicationsInternalServerError {
	return &GetApplicationsInternalServerError{}
}

/*GetApplicationsInternalServerError handles this case with default header values.

The server has encountered a situation it does not know how to handle.
*/
type GetApplicationsInternalServerError struct {
	Payload *models.Error
}

func (o *GetApplicationsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /applications][%d] getApplicationsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetApplicationsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetApplicationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApplicationsDefault creates a GetApplicationsDefault with default headers values
func NewGetApplicationsDefault(code int) *GetApplicationsDefault {
	return &GetApplicationsDefault{
		_statusCode: code,
	}
}

/*GetApplicationsDefault handles this case with default header values.

Generic error.
*/
type GetApplicationsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get applications default response
func (o *GetApplicationsDefault) Code() int {
	return o._statusCode
}

func (o *GetApplicationsDefault) Error() string {
	return fmt.Sprintf("[GET /applications][%d] GetApplications default  %+v", o._statusCode, o.Payload)
}

func (o *GetApplicationsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetApplicationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}