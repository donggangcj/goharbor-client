// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/donggangcj/goharbor-client/v3/apiv2/model/legacy"
)

// GetScannersRegistrationIDReader is a Reader for the GetScannersRegistrationID structure.
type GetScannersRegistrationIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScannersRegistrationIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScannersRegistrationIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetScannersRegistrationIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetScannersRegistrationIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetScannersRegistrationIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetScannersRegistrationIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetScannersRegistrationIDOK creates a GetScannersRegistrationIDOK with default headers values
func NewGetScannersRegistrationIDOK() *GetScannersRegistrationIDOK {
	return &GetScannersRegistrationIDOK{}
}

/* GetScannersRegistrationIDOK describes a response with status code 200, with default header values.

The details of the scanner registration.
*/
type GetScannersRegistrationIDOK struct {
	Payload *legacy.ScannerRegistration
}

func (o *GetScannersRegistrationIDOK) Error() string {
	return fmt.Sprintf("[GET /scanners/{registration_id}][%d] getScannersRegistrationIdOK  %+v", 200, o.Payload)
}
func (o *GetScannersRegistrationIDOK) GetPayload() *legacy.ScannerRegistration {
	return o.Payload
}

func (o *GetScannersRegistrationIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(legacy.ScannerRegistration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScannersRegistrationIDUnauthorized creates a GetScannersRegistrationIDUnauthorized with default headers values
func NewGetScannersRegistrationIDUnauthorized() *GetScannersRegistrationIDUnauthorized {
	return &GetScannersRegistrationIDUnauthorized{}
}

/* GetScannersRegistrationIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized request
*/
type GetScannersRegistrationIDUnauthorized struct {
}

func (o *GetScannersRegistrationIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /scanners/{registration_id}][%d] getScannersRegistrationIdUnauthorized ", 401)
}

func (o *GetScannersRegistrationIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetScannersRegistrationIDForbidden creates a GetScannersRegistrationIDForbidden with default headers values
func NewGetScannersRegistrationIDForbidden() *GetScannersRegistrationIDForbidden {
	return &GetScannersRegistrationIDForbidden{}
}

/* GetScannersRegistrationIDForbidden describes a response with status code 403, with default header values.

Request is not allowed, system role required
*/
type GetScannersRegistrationIDForbidden struct {
}

func (o *GetScannersRegistrationIDForbidden) Error() string {
	return fmt.Sprintf("[GET /scanners/{registration_id}][%d] getScannersRegistrationIdForbidden ", 403)
}

func (o *GetScannersRegistrationIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetScannersRegistrationIDNotFound creates a GetScannersRegistrationIDNotFound with default headers values
func NewGetScannersRegistrationIDNotFound() *GetScannersRegistrationIDNotFound {
	return &GetScannersRegistrationIDNotFound{}
}

/* GetScannersRegistrationIDNotFound describes a response with status code 404, with default header values.

The requested object is not found
*/
type GetScannersRegistrationIDNotFound struct {
}

func (o *GetScannersRegistrationIDNotFound) Error() string {
	return fmt.Sprintf("[GET /scanners/{registration_id}][%d] getScannersRegistrationIdNotFound ", 404)
}

func (o *GetScannersRegistrationIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetScannersRegistrationIDInternalServerError creates a GetScannersRegistrationIDInternalServerError with default headers values
func NewGetScannersRegistrationIDInternalServerError() *GetScannersRegistrationIDInternalServerError {
	return &GetScannersRegistrationIDInternalServerError{}
}

/* GetScannersRegistrationIDInternalServerError describes a response with status code 500, with default header values.

Internal server error happened
*/
type GetScannersRegistrationIDInternalServerError struct {
}

func (o *GetScannersRegistrationIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /scanners/{registration_id}][%d] getScannersRegistrationIdInternalServerError ", 500)
}

func (o *GetScannersRegistrationIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}