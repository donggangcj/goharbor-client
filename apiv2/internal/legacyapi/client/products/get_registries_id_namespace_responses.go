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

// GetRegistriesIDNamespaceReader is a Reader for the GetRegistriesIDNamespace structure.
type GetRegistriesIDNamespaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRegistriesIDNamespaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRegistriesIDNamespaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetRegistriesIDNamespaceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRegistriesIDNamespaceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRegistriesIDNamespaceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRegistriesIDNamespaceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRegistriesIDNamespaceOK creates a GetRegistriesIDNamespaceOK with default headers values
func NewGetRegistriesIDNamespaceOK() *GetRegistriesIDNamespaceOK {
	return &GetRegistriesIDNamespaceOK{}
}

/* GetRegistriesIDNamespaceOK describes a response with status code 200, with default header values.

Success
*/
type GetRegistriesIDNamespaceOK struct {
	Payload []*legacy.Namespace
}

func (o *GetRegistriesIDNamespaceOK) Error() string {
	return fmt.Sprintf("[GET /registries/{id}/namespace][%d] getRegistriesIdNamespaceOK  %+v", 200, o.Payload)
}
func (o *GetRegistriesIDNamespaceOK) GetPayload() []*legacy.Namespace {
	return o.Payload
}

func (o *GetRegistriesIDNamespaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRegistriesIDNamespaceUnauthorized creates a GetRegistriesIDNamespaceUnauthorized with default headers values
func NewGetRegistriesIDNamespaceUnauthorized() *GetRegistriesIDNamespaceUnauthorized {
	return &GetRegistriesIDNamespaceUnauthorized{}
}

/* GetRegistriesIDNamespaceUnauthorized describes a response with status code 401, with default header values.

User need to login first.
*/
type GetRegistriesIDNamespaceUnauthorized struct {
}

func (o *GetRegistriesIDNamespaceUnauthorized) Error() string {
	return fmt.Sprintf("[GET /registries/{id}/namespace][%d] getRegistriesIdNamespaceUnauthorized ", 401)
}

func (o *GetRegistriesIDNamespaceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRegistriesIDNamespaceForbidden creates a GetRegistriesIDNamespaceForbidden with default headers values
func NewGetRegistriesIDNamespaceForbidden() *GetRegistriesIDNamespaceForbidden {
	return &GetRegistriesIDNamespaceForbidden{}
}

/* GetRegistriesIDNamespaceForbidden describes a response with status code 403, with default header values.

User has no privilege for the operation.
*/
type GetRegistriesIDNamespaceForbidden struct {
}

func (o *GetRegistriesIDNamespaceForbidden) Error() string {
	return fmt.Sprintf("[GET /registries/{id}/namespace][%d] getRegistriesIdNamespaceForbidden ", 403)
}

func (o *GetRegistriesIDNamespaceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRegistriesIDNamespaceNotFound creates a GetRegistriesIDNamespaceNotFound with default headers values
func NewGetRegistriesIDNamespaceNotFound() *GetRegistriesIDNamespaceNotFound {
	return &GetRegistriesIDNamespaceNotFound{}
}

/* GetRegistriesIDNamespaceNotFound describes a response with status code 404, with default header values.

No registry found.
*/
type GetRegistriesIDNamespaceNotFound struct {
}

func (o *GetRegistriesIDNamespaceNotFound) Error() string {
	return fmt.Sprintf("[GET /registries/{id}/namespace][%d] getRegistriesIdNamespaceNotFound ", 404)
}

func (o *GetRegistriesIDNamespaceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRegistriesIDNamespaceInternalServerError creates a GetRegistriesIDNamespaceInternalServerError with default headers values
func NewGetRegistriesIDNamespaceInternalServerError() *GetRegistriesIDNamespaceInternalServerError {
	return &GetRegistriesIDNamespaceInternalServerError{}
}

/* GetRegistriesIDNamespaceInternalServerError describes a response with status code 500, with default header values.

Unexpected internal errors.
*/
type GetRegistriesIDNamespaceInternalServerError struct {
}

func (o *GetRegistriesIDNamespaceInternalServerError) Error() string {
	return fmt.Sprintf("[GET /registries/{id}/namespace][%d] getRegistriesIdNamespaceInternalServerError ", 500)
}

func (o *GetRegistriesIDNamespaceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
