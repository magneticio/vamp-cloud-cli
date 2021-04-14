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

// PutApplicationsApplicationIDServicesServiceIDReader is a Reader for the PutApplicationsApplicationIDServicesServiceID structure.
type PutApplicationsApplicationIDServicesServiceIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutApplicationsApplicationIDServicesServiceIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutApplicationsApplicationIDServicesServiceIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPutApplicationsApplicationIDServicesServiceIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutApplicationsApplicationIDServicesServiceIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutApplicationsApplicationIDServicesServiceIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPutApplicationsApplicationIDServicesServiceIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutApplicationsApplicationIDServicesServiceIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPutApplicationsApplicationIDServicesServiceIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutApplicationsApplicationIDServicesServiceIDOK creates a PutApplicationsApplicationIDServicesServiceIDOK with default headers values
func NewPutApplicationsApplicationIDServicesServiceIDOK() *PutApplicationsApplicationIDServicesServiceIDOK {
	return &PutApplicationsApplicationIDServicesServiceIDOK{}
}

/*PutApplicationsApplicationIDServicesServiceIDOK handles this case with default header values.

OK
*/
type PutApplicationsApplicationIDServicesServiceIDOK struct {
}

func (o *PutApplicationsApplicationIDServicesServiceIDOK) Error() string {
	return fmt.Sprintf("[PUT /applications/{applicationID}/services/{serviceID}][%d] putApplicationsApplicationIdServicesServiceIdOK ", 200)
}

func (o *PutApplicationsApplicationIDServicesServiceIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutApplicationsApplicationIDServicesServiceIDUnauthorized creates a PutApplicationsApplicationIDServicesServiceIDUnauthorized with default headers values
func NewPutApplicationsApplicationIDServicesServiceIDUnauthorized() *PutApplicationsApplicationIDServicesServiceIDUnauthorized {
	return &PutApplicationsApplicationIDServicesServiceIDUnauthorized{}
}

/*PutApplicationsApplicationIDServicesServiceIDUnauthorized handles this case with default header values.

The requester is not authorized.
*/
type PutApplicationsApplicationIDServicesServiceIDUnauthorized struct {
	Payload *models.Error
}

func (o *PutApplicationsApplicationIDServicesServiceIDUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /applications/{applicationID}/services/{serviceID}][%d] putApplicationsApplicationIdServicesServiceIdUnauthorized  %+v", 401, o.Payload)
}

func (o *PutApplicationsApplicationIDServicesServiceIDUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutApplicationsApplicationIDServicesServiceIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutApplicationsApplicationIDServicesServiceIDForbidden creates a PutApplicationsApplicationIDServicesServiceIDForbidden with default headers values
func NewPutApplicationsApplicationIDServicesServiceIDForbidden() *PutApplicationsApplicationIDServicesServiceIDForbidden {
	return &PutApplicationsApplicationIDServicesServiceIDForbidden{}
}

/*PutApplicationsApplicationIDServicesServiceIDForbidden handles this case with default header values.

The requester does not have access rights to the resource.
*/
type PutApplicationsApplicationIDServicesServiceIDForbidden struct {
	Payload *models.Error
}

func (o *PutApplicationsApplicationIDServicesServiceIDForbidden) Error() string {
	return fmt.Sprintf("[PUT /applications/{applicationID}/services/{serviceID}][%d] putApplicationsApplicationIdServicesServiceIdForbidden  %+v", 403, o.Payload)
}

func (o *PutApplicationsApplicationIDServicesServiceIDForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutApplicationsApplicationIDServicesServiceIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutApplicationsApplicationIDServicesServiceIDNotFound creates a PutApplicationsApplicationIDServicesServiceIDNotFound with default headers values
func NewPutApplicationsApplicationIDServicesServiceIDNotFound() *PutApplicationsApplicationIDServicesServiceIDNotFound {
	return &PutApplicationsApplicationIDServicesServiceIDNotFound{}
}

/*PutApplicationsApplicationIDServicesServiceIDNotFound handles this case with default header values.

The server can not find the requested resource.
*/
type PutApplicationsApplicationIDServicesServiceIDNotFound struct {
	Payload *models.Error
}

func (o *PutApplicationsApplicationIDServicesServiceIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /applications/{applicationID}/services/{serviceID}][%d] putApplicationsApplicationIdServicesServiceIdNotFound  %+v", 404, o.Payload)
}

func (o *PutApplicationsApplicationIDServicesServiceIDNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutApplicationsApplicationIDServicesServiceIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutApplicationsApplicationIDServicesServiceIDConflict creates a PutApplicationsApplicationIDServicesServiceIDConflict with default headers values
func NewPutApplicationsApplicationIDServicesServiceIDConflict() *PutApplicationsApplicationIDServicesServiceIDConflict {
	return &PutApplicationsApplicationIDServicesServiceIDConflict{}
}

/*PutApplicationsApplicationIDServicesServiceIDConflict handles this case with default header values.

The request cannot be completed due to a conflict with the current state of the resource.
*/
type PutApplicationsApplicationIDServicesServiceIDConflict struct {
	Payload *models.Error
}

func (o *PutApplicationsApplicationIDServicesServiceIDConflict) Error() string {
	return fmt.Sprintf("[PUT /applications/{applicationID}/services/{serviceID}][%d] putApplicationsApplicationIdServicesServiceIdConflict  %+v", 409, o.Payload)
}

func (o *PutApplicationsApplicationIDServicesServiceIDConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutApplicationsApplicationIDServicesServiceIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutApplicationsApplicationIDServicesServiceIDInternalServerError creates a PutApplicationsApplicationIDServicesServiceIDInternalServerError with default headers values
func NewPutApplicationsApplicationIDServicesServiceIDInternalServerError() *PutApplicationsApplicationIDServicesServiceIDInternalServerError {
	return &PutApplicationsApplicationIDServicesServiceIDInternalServerError{}
}

/*PutApplicationsApplicationIDServicesServiceIDInternalServerError handles this case with default header values.

The server has encountered a situation it does not know how to handle.
*/
type PutApplicationsApplicationIDServicesServiceIDInternalServerError struct {
	Payload *models.Error
}

func (o *PutApplicationsApplicationIDServicesServiceIDInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /applications/{applicationID}/services/{serviceID}][%d] putApplicationsApplicationIdServicesServiceIdInternalServerError  %+v", 500, o.Payload)
}

func (o *PutApplicationsApplicationIDServicesServiceIDInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutApplicationsApplicationIDServicesServiceIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutApplicationsApplicationIDServicesServiceIDDefault creates a PutApplicationsApplicationIDServicesServiceIDDefault with default headers values
func NewPutApplicationsApplicationIDServicesServiceIDDefault(code int) *PutApplicationsApplicationIDServicesServiceIDDefault {
	return &PutApplicationsApplicationIDServicesServiceIDDefault{
		_statusCode: code,
	}
}

/*PutApplicationsApplicationIDServicesServiceIDDefault handles this case with default header values.

Generic error.
*/
type PutApplicationsApplicationIDServicesServiceIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put applications application ID services service ID default response
func (o *PutApplicationsApplicationIDServicesServiceIDDefault) Code() int {
	return o._statusCode
}

func (o *PutApplicationsApplicationIDServicesServiceIDDefault) Error() string {
	return fmt.Sprintf("[PUT /applications/{applicationID}/services/{serviceID}][%d] PutApplicationsApplicationIDServicesServiceID default  %+v", o._statusCode, o.Payload)
}

func (o *PutApplicationsApplicationIDServicesServiceIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutApplicationsApplicationIDServicesServiceIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
