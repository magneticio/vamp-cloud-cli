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

// PatchApplicationsApplicationIDIngressesIngressIDReader is a Reader for the PatchApplicationsApplicationIDIngressesIngressID structure.
type PatchApplicationsApplicationIDIngressesIngressIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchApplicationsApplicationIDIngressesIngressIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchApplicationsApplicationIDIngressesIngressIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchApplicationsApplicationIDIngressesIngressIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchApplicationsApplicationIDIngressesIngressIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchApplicationsApplicationIDIngressesIngressIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchApplicationsApplicationIDIngressesIngressIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchApplicationsApplicationIDIngressesIngressIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchApplicationsApplicationIDIngressesIngressIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchApplicationsApplicationIDIngressesIngressIDOK creates a PatchApplicationsApplicationIDIngressesIngressIDOK with default headers values
func NewPatchApplicationsApplicationIDIngressesIngressIDOK() *PatchApplicationsApplicationIDIngressesIngressIDOK {
	return &PatchApplicationsApplicationIDIngressesIngressIDOK{}
}

/*PatchApplicationsApplicationIDIngressesIngressIDOK handles this case with default header values.

OK
*/
type PatchApplicationsApplicationIDIngressesIngressIDOK struct {
}

func (o *PatchApplicationsApplicationIDIngressesIngressIDOK) Error() string {
	return fmt.Sprintf("[PATCH /applications/{applicationID}/ingresses/{ingressID}][%d] patchApplicationsApplicationIdIngressesIngressIdOK ", 200)
}

func (o *PatchApplicationsApplicationIDIngressesIngressIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchApplicationsApplicationIDIngressesIngressIDBadRequest creates a PatchApplicationsApplicationIDIngressesIngressIDBadRequest with default headers values
func NewPatchApplicationsApplicationIDIngressesIngressIDBadRequest() *PatchApplicationsApplicationIDIngressesIngressIDBadRequest {
	return &PatchApplicationsApplicationIDIngressesIngressIDBadRequest{}
}

/*PatchApplicationsApplicationIDIngressesIngressIDBadRequest handles this case with default header values.

The request is invalid.
*/
type PatchApplicationsApplicationIDIngressesIngressIDBadRequest struct {
	Payload *models.Error
}

func (o *PatchApplicationsApplicationIDIngressesIngressIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /applications/{applicationID}/ingresses/{ingressID}][%d] patchApplicationsApplicationIdIngressesIngressIdBadRequest  %+v", 400, o.Payload)
}

func (o *PatchApplicationsApplicationIDIngressesIngressIDBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchApplicationsApplicationIDIngressesIngressIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchApplicationsApplicationIDIngressesIngressIDUnauthorized creates a PatchApplicationsApplicationIDIngressesIngressIDUnauthorized with default headers values
func NewPatchApplicationsApplicationIDIngressesIngressIDUnauthorized() *PatchApplicationsApplicationIDIngressesIngressIDUnauthorized {
	return &PatchApplicationsApplicationIDIngressesIngressIDUnauthorized{}
}

/*PatchApplicationsApplicationIDIngressesIngressIDUnauthorized handles this case with default header values.

The requester is not authorized.
*/
type PatchApplicationsApplicationIDIngressesIngressIDUnauthorized struct {
	Payload *models.Error
}

func (o *PatchApplicationsApplicationIDIngressesIngressIDUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /applications/{applicationID}/ingresses/{ingressID}][%d] patchApplicationsApplicationIdIngressesIngressIdUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchApplicationsApplicationIDIngressesIngressIDUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchApplicationsApplicationIDIngressesIngressIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchApplicationsApplicationIDIngressesIngressIDForbidden creates a PatchApplicationsApplicationIDIngressesIngressIDForbidden with default headers values
func NewPatchApplicationsApplicationIDIngressesIngressIDForbidden() *PatchApplicationsApplicationIDIngressesIngressIDForbidden {
	return &PatchApplicationsApplicationIDIngressesIngressIDForbidden{}
}

/*PatchApplicationsApplicationIDIngressesIngressIDForbidden handles this case with default header values.

The requester does not have access rights to the resource.
*/
type PatchApplicationsApplicationIDIngressesIngressIDForbidden struct {
	Payload *models.Error
}

func (o *PatchApplicationsApplicationIDIngressesIngressIDForbidden) Error() string {
	return fmt.Sprintf("[PATCH /applications/{applicationID}/ingresses/{ingressID}][%d] patchApplicationsApplicationIdIngressesIngressIdForbidden  %+v", 403, o.Payload)
}

func (o *PatchApplicationsApplicationIDIngressesIngressIDForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchApplicationsApplicationIDIngressesIngressIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchApplicationsApplicationIDIngressesIngressIDNotFound creates a PatchApplicationsApplicationIDIngressesIngressIDNotFound with default headers values
func NewPatchApplicationsApplicationIDIngressesIngressIDNotFound() *PatchApplicationsApplicationIDIngressesIngressIDNotFound {
	return &PatchApplicationsApplicationIDIngressesIngressIDNotFound{}
}

/*PatchApplicationsApplicationIDIngressesIngressIDNotFound handles this case with default header values.

The server can not find the requested resource.
*/
type PatchApplicationsApplicationIDIngressesIngressIDNotFound struct {
	Payload *models.Error
}

func (o *PatchApplicationsApplicationIDIngressesIngressIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /applications/{applicationID}/ingresses/{ingressID}][%d] patchApplicationsApplicationIdIngressesIngressIdNotFound  %+v", 404, o.Payload)
}

func (o *PatchApplicationsApplicationIDIngressesIngressIDNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchApplicationsApplicationIDIngressesIngressIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchApplicationsApplicationIDIngressesIngressIDInternalServerError creates a PatchApplicationsApplicationIDIngressesIngressIDInternalServerError with default headers values
func NewPatchApplicationsApplicationIDIngressesIngressIDInternalServerError() *PatchApplicationsApplicationIDIngressesIngressIDInternalServerError {
	return &PatchApplicationsApplicationIDIngressesIngressIDInternalServerError{}
}

/*PatchApplicationsApplicationIDIngressesIngressIDInternalServerError handles this case with default header values.

The server has encountered a situation it does not know how to handle.
*/
type PatchApplicationsApplicationIDIngressesIngressIDInternalServerError struct {
	Payload *models.Error
}

func (o *PatchApplicationsApplicationIDIngressesIngressIDInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /applications/{applicationID}/ingresses/{ingressID}][%d] patchApplicationsApplicationIdIngressesIngressIdInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchApplicationsApplicationIDIngressesIngressIDInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchApplicationsApplicationIDIngressesIngressIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchApplicationsApplicationIDIngressesIngressIDDefault creates a PatchApplicationsApplicationIDIngressesIngressIDDefault with default headers values
func NewPatchApplicationsApplicationIDIngressesIngressIDDefault(code int) *PatchApplicationsApplicationIDIngressesIngressIDDefault {
	return &PatchApplicationsApplicationIDIngressesIngressIDDefault{
		_statusCode: code,
	}
}

/*PatchApplicationsApplicationIDIngressesIngressIDDefault handles this case with default header values.

Generic error.
*/
type PatchApplicationsApplicationIDIngressesIngressIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch applications application ID ingresses ingress ID default response
func (o *PatchApplicationsApplicationIDIngressesIngressIDDefault) Code() int {
	return o._statusCode
}

func (o *PatchApplicationsApplicationIDIngressesIngressIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /applications/{applicationID}/ingresses/{ingressID}][%d] PatchApplicationsApplicationIDIngressesIngressID default  %+v", o._statusCode, o.Payload)
}

func (o *PatchApplicationsApplicationIDIngressesIngressIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchApplicationsApplicationIDIngressesIngressIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
