// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostRegistriesReader is a Reader for the PostRegistries structure.
type PostRegistriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRegistriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostRegistriesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostRegistriesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostRegistriesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostRegistriesConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostRegistriesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostRegistriesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostRegistriesCreated creates a PostRegistriesCreated with default headers values
func NewPostRegistriesCreated() *PostRegistriesCreated {
	return &PostRegistriesCreated{}
}

/* PostRegistriesCreated describes a response with status code 201, with default header values.

Registry created successfully.
*/
type PostRegistriesCreated struct {
}

func (o *PostRegistriesCreated) Error() string {
	return fmt.Sprintf("[POST /registries][%d] postRegistriesCreated ", 201)
}

func (o *PostRegistriesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRegistriesBadRequest creates a PostRegistriesBadRequest with default headers values
func NewPostRegistriesBadRequest() *PostRegistriesBadRequest {
	return &PostRegistriesBadRequest{}
}

/* PostRegistriesBadRequest describes a response with status code 400, with default header values.

Unsatisfied with constraints of the registry creation.
*/
type PostRegistriesBadRequest struct {
}

func (o *PostRegistriesBadRequest) Error() string {
	return fmt.Sprintf("[POST /registries][%d] postRegistriesBadRequest ", 400)
}

func (o *PostRegistriesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRegistriesUnauthorized creates a PostRegistriesUnauthorized with default headers values
func NewPostRegistriesUnauthorized() *PostRegistriesUnauthorized {
	return &PostRegistriesUnauthorized{}
}

/* PostRegistriesUnauthorized describes a response with status code 401, with default header values.

User need to log in first.
*/
type PostRegistriesUnauthorized struct {
}

func (o *PostRegistriesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /registries][%d] postRegistriesUnauthorized ", 401)
}

func (o *PostRegistriesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRegistriesConflict creates a PostRegistriesConflict with default headers values
func NewPostRegistriesConflict() *PostRegistriesConflict {
	return &PostRegistriesConflict{}
}

/* PostRegistriesConflict describes a response with status code 409, with default header values.

Registry name already exists.
*/
type PostRegistriesConflict struct {
}

func (o *PostRegistriesConflict) Error() string {
	return fmt.Sprintf("[POST /registries][%d] postRegistriesConflict ", 409)
}

func (o *PostRegistriesConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRegistriesUnsupportedMediaType creates a PostRegistriesUnsupportedMediaType with default headers values
func NewPostRegistriesUnsupportedMediaType() *PostRegistriesUnsupportedMediaType {
	return &PostRegistriesUnsupportedMediaType{}
}

/* PostRegistriesUnsupportedMediaType describes a response with status code 415, with default header values.

The Media Type of the request is not supported, it has to be "application/json"
*/
type PostRegistriesUnsupportedMediaType struct {
}

func (o *PostRegistriesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /registries][%d] postRegistriesUnsupportedMediaType ", 415)
}

func (o *PostRegistriesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRegistriesInternalServerError creates a PostRegistriesInternalServerError with default headers values
func NewPostRegistriesInternalServerError() *PostRegistriesInternalServerError {
	return &PostRegistriesInternalServerError{}
}

/* PostRegistriesInternalServerError describes a response with status code 500, with default header values.

Unexpected generate errors.
*/
type PostRegistriesInternalServerError struct {
}

func (o *PostRegistriesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /registries][%d] postRegistriesInternalServerError ", 500)
}

func (o *PostRegistriesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
