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

// GetReleasesIDReader is a Reader for the GetReleasesID structure.
type GetReleasesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReleasesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReleasesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetReleasesIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetReleasesIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetReleasesIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetReleasesIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetReleasesIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetReleasesIDOK creates a GetReleasesIDOK with default headers values
func NewGetReleasesIDOK() *GetReleasesIDOK {
	return &GetReleasesIDOK{}
}

/*GetReleasesIDOK handles this case with default header values.

Release status object
*/
type GetReleasesIDOK struct {
	Payload *models.Release
}

func (o *GetReleasesIDOK) Error() string {
	return fmt.Sprintf("[GET /releases/{id}][%d] getReleasesIdOK  %+v", 200, o.Payload)
}

func (o *GetReleasesIDOK) GetPayload() *models.Release {
	return o.Payload
}

func (o *GetReleasesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Release)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReleasesIDUnauthorized creates a GetReleasesIDUnauthorized with default headers values
func NewGetReleasesIDUnauthorized() *GetReleasesIDUnauthorized {
	return &GetReleasesIDUnauthorized{}
}

/*GetReleasesIDUnauthorized handles this case with default header values.

The requester is not authorized.
*/
type GetReleasesIDUnauthorized struct {
	Payload *models.Error
}

func (o *GetReleasesIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /releases/{id}][%d] getReleasesIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetReleasesIDUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetReleasesIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReleasesIDForbidden creates a GetReleasesIDForbidden with default headers values
func NewGetReleasesIDForbidden() *GetReleasesIDForbidden {
	return &GetReleasesIDForbidden{}
}

/*GetReleasesIDForbidden handles this case with default header values.

The requester does not have access rights to the resource.
*/
type GetReleasesIDForbidden struct {
	Payload *models.Error
}

func (o *GetReleasesIDForbidden) Error() string {
	return fmt.Sprintf("[GET /releases/{id}][%d] getReleasesIdForbidden  %+v", 403, o.Payload)
}

func (o *GetReleasesIDForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetReleasesIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReleasesIDNotFound creates a GetReleasesIDNotFound with default headers values
func NewGetReleasesIDNotFound() *GetReleasesIDNotFound {
	return &GetReleasesIDNotFound{}
}

/*GetReleasesIDNotFound handles this case with default header values.

The server can not find the requested resource.
*/
type GetReleasesIDNotFound struct {
	Payload *models.Error
}

func (o *GetReleasesIDNotFound) Error() string {
	return fmt.Sprintf("[GET /releases/{id}][%d] getReleasesIdNotFound  %+v", 404, o.Payload)
}

func (o *GetReleasesIDNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetReleasesIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReleasesIDInternalServerError creates a GetReleasesIDInternalServerError with default headers values
func NewGetReleasesIDInternalServerError() *GetReleasesIDInternalServerError {
	return &GetReleasesIDInternalServerError{}
}

/*GetReleasesIDInternalServerError handles this case with default header values.

The server has encountered a situation it does not know how to handle.
*/
type GetReleasesIDInternalServerError struct {
	Payload *models.Error
}

func (o *GetReleasesIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /releases/{id}][%d] getReleasesIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetReleasesIDInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetReleasesIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReleasesIDDefault creates a GetReleasesIDDefault with default headers values
func NewGetReleasesIDDefault(code int) *GetReleasesIDDefault {
	return &GetReleasesIDDefault{
		_statusCode: code,
	}
}

/*GetReleasesIDDefault handles this case with default header values.

Generic error.
*/
type GetReleasesIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get releases ID default response
func (o *GetReleasesIDDefault) Code() int {
	return o._statusCode
}

func (o *GetReleasesIDDefault) Error() string {
	return fmt.Sprintf("[GET /releases/{id}][%d] GetReleasesID default  %+v", o._statusCode, o.Payload)
}

func (o *GetReleasesIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetReleasesIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
