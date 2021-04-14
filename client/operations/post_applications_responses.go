// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/magneticio/vamp-cloud-cli/models"
)

// PostApplicationsReader is a Reader for the PostApplications structure.
type PostApplicationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostApplicationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostApplicationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostApplicationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostApplicationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPostApplicationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostApplicationsOK creates a PostApplicationsOK with default headers values
func NewPostApplicationsOK() *PostApplicationsOK {
	return &PostApplicationsOK{}
}

/*PostApplicationsOK handles this case with default header values.

ID of created application
*/
type PostApplicationsOK struct {
	Payload *PostApplicationsOKBody
}

func (o *PostApplicationsOK) Error() string {
	return fmt.Sprintf("[POST /applications][%d] postApplicationsOK  %+v", 200, o.Payload)
}

func (o *PostApplicationsOK) GetPayload() *PostApplicationsOKBody {
	return o.Payload
}

func (o *PostApplicationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostApplicationsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostApplicationsUnauthorized creates a PostApplicationsUnauthorized with default headers values
func NewPostApplicationsUnauthorized() *PostApplicationsUnauthorized {
	return &PostApplicationsUnauthorized{}
}

/*PostApplicationsUnauthorized handles this case with default header values.

The requester is not authorized.
*/
type PostApplicationsUnauthorized struct {
	Payload *models.Error
}

func (o *PostApplicationsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /applications][%d] postApplicationsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostApplicationsUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostApplicationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostApplicationsInternalServerError creates a PostApplicationsInternalServerError with default headers values
func NewPostApplicationsInternalServerError() *PostApplicationsInternalServerError {
	return &PostApplicationsInternalServerError{}
}

/*PostApplicationsInternalServerError handles this case with default header values.

The server has encountered a situation it does not know how to handle.
*/
type PostApplicationsInternalServerError struct {
	Payload *models.Error
}

func (o *PostApplicationsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /applications][%d] postApplicationsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostApplicationsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostApplicationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostApplicationsDefault creates a PostApplicationsDefault with default headers values
func NewPostApplicationsDefault(code int) *PostApplicationsDefault {
	return &PostApplicationsDefault{
		_statusCode: code,
	}
}

/*PostApplicationsDefault handles this case with default header values.

Generic error.
*/
type PostApplicationsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post applications default response
func (o *PostApplicationsDefault) Code() int {
	return o._statusCode
}

func (o *PostApplicationsDefault) Error() string {
	return fmt.Sprintf("[POST /applications][%d] PostApplications default  %+v", o._statusCode, o.Payload)
}

func (o *PostApplicationsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostApplicationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostApplicationsOKBody post applications o k body
swagger:model PostApplicationsOKBody
*/
type PostApplicationsOKBody struct {

	// The application ID
	ID int64 `json:"id,omitempty"`
}

// Validate validates this post applications o k body
func (o *PostApplicationsOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostApplicationsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostApplicationsOKBody) UnmarshalBinary(b []byte) error {
	var res PostApplicationsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
