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

// GetClustersReader is a Reader for the GetClusters structure.
type GetClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClustersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetClustersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetClustersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetClustersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetClustersOK creates a GetClustersOK with default headers values
func NewGetClustersOK() *GetClustersOK {
	return &GetClustersOK{}
}

/*GetClustersOK handles this case with default header values.

List of cluster objects
*/
type GetClustersOK struct {
	Payload *models.Clusters
}

func (o *GetClustersOK) Error() string {
	return fmt.Sprintf("[GET /clusters][%d] getClustersOK  %+v", 200, o.Payload)
}

func (o *GetClustersOK) GetPayload() *models.Clusters {
	return o.Payload
}

func (o *GetClustersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Clusters)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClustersUnauthorized creates a GetClustersUnauthorized with default headers values
func NewGetClustersUnauthorized() *GetClustersUnauthorized {
	return &GetClustersUnauthorized{}
}

/*GetClustersUnauthorized handles this case with default header values.

The requester is not authorized.
*/
type GetClustersUnauthorized struct {
	Payload *models.Error
}

func (o *GetClustersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /clusters][%d] getClustersUnauthorized  %+v", 401, o.Payload)
}

func (o *GetClustersUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetClustersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClustersInternalServerError creates a GetClustersInternalServerError with default headers values
func NewGetClustersInternalServerError() *GetClustersInternalServerError {
	return &GetClustersInternalServerError{}
}

/*GetClustersInternalServerError handles this case with default header values.

The server has encountered a situation it does not know how to handle.
*/
type GetClustersInternalServerError struct {
	Payload *models.Error
}

func (o *GetClustersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /clusters][%d] getClustersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetClustersInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetClustersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClustersDefault creates a GetClustersDefault with default headers values
func NewGetClustersDefault(code int) *GetClustersDefault {
	return &GetClustersDefault{
		_statusCode: code,
	}
}

/*GetClustersDefault handles this case with default header values.

Generic error.
*/
type GetClustersDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get clusters default response
func (o *GetClustersDefault) Code() int {
	return o._statusCode
}

func (o *GetClustersDefault) Error() string {
	return fmt.Sprintf("[GET /clusters][%d] GetClusters default  %+v", o._statusCode, o.Payload)
}

func (o *GetClustersDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetClustersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
