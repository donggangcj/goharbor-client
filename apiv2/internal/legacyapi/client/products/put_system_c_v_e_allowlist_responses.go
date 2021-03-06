// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PutSystemCVEAllowlistReader is a Reader for the PutSystemCVEAllowlist structure.
type PutSystemCVEAllowlistReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutSystemCVEAllowlistReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutSystemCVEAllowlistOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPutSystemCVEAllowlistUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutSystemCVEAllowlistForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutSystemCVEAllowlistInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutSystemCVEAllowlistOK creates a PutSystemCVEAllowlistOK with default headers values
func NewPutSystemCVEAllowlistOK() *PutSystemCVEAllowlistOK {
	return &PutSystemCVEAllowlistOK{}
}

/* PutSystemCVEAllowlistOK describes a response with status code 200, with default header values.

Successfully updated the CVE allowlist.
*/
type PutSystemCVEAllowlistOK struct {
}

func (o *PutSystemCVEAllowlistOK) Error() string {
	return fmt.Sprintf("[PUT /system/CVEAllowlist][%d] putSystemCVEAllowlistOK ", 200)
}

func (o *PutSystemCVEAllowlistOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutSystemCVEAllowlistUnauthorized creates a PutSystemCVEAllowlistUnauthorized with default headers values
func NewPutSystemCVEAllowlistUnauthorized() *PutSystemCVEAllowlistUnauthorized {
	return &PutSystemCVEAllowlistUnauthorized{}
}

/* PutSystemCVEAllowlistUnauthorized describes a response with status code 401, with default header values.

User is not authenticated.
*/
type PutSystemCVEAllowlistUnauthorized struct {
}

func (o *PutSystemCVEAllowlistUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /system/CVEAllowlist][%d] putSystemCVEAllowlistUnauthorized ", 401)
}

func (o *PutSystemCVEAllowlistUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutSystemCVEAllowlistForbidden creates a PutSystemCVEAllowlistForbidden with default headers values
func NewPutSystemCVEAllowlistForbidden() *PutSystemCVEAllowlistForbidden {
	return &PutSystemCVEAllowlistForbidden{}
}

/* PutSystemCVEAllowlistForbidden describes a response with status code 403, with default header values.

User does not have permission to call this API.
*/
type PutSystemCVEAllowlistForbidden struct {
}

func (o *PutSystemCVEAllowlistForbidden) Error() string {
	return fmt.Sprintf("[PUT /system/CVEAllowlist][%d] putSystemCVEAllowlistForbidden ", 403)
}

func (o *PutSystemCVEAllowlistForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutSystemCVEAllowlistInternalServerError creates a PutSystemCVEAllowlistInternalServerError with default headers values
func NewPutSystemCVEAllowlistInternalServerError() *PutSystemCVEAllowlistInternalServerError {
	return &PutSystemCVEAllowlistInternalServerError{}
}

/* PutSystemCVEAllowlistInternalServerError describes a response with status code 500, with default header values.

Unexpected internal errors.
*/
type PutSystemCVEAllowlistInternalServerError struct {
}

func (o *PutSystemCVEAllowlistInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /system/CVEAllowlist][%d] putSystemCVEAllowlistInternalServerError ", 500)
}

func (o *PutSystemCVEAllowlistInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
