// Code generated by go-swagger; DO NOT EDIT.

package icon

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/donggangcj/goharbor-client/v3/apiv2/model"
)

// GetIconReader is a Reader for the GetIcon structure.
type GetIconReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIconReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIconOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIconBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetIconNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIconInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetIconOK creates a GetIconOK with default headers values
func NewGetIconOK() *GetIconOK {
	return &GetIconOK{}
}

/* GetIconOK describes a response with status code 200, with default header values.

Success
*/
type GetIconOK struct {
	Payload *model.Icon
}

func (o *GetIconOK) Error() string {
	return fmt.Sprintf("[GET /icons/{digest}][%d] getIconOK  %+v", 200, o.Payload)
}
func (o *GetIconOK) GetPayload() *model.Icon {
	return o.Payload
}

func (o *GetIconOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.Icon)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIconBadRequest creates a GetIconBadRequest with default headers values
func NewGetIconBadRequest() *GetIconBadRequest {
	return &GetIconBadRequest{}
}

/* GetIconBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetIconBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetIconBadRequest) Error() string {
	return fmt.Sprintf("[GET /icons/{digest}][%d] getIconBadRequest  %+v", 400, o.Payload)
}
func (o *GetIconBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetIconBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIconNotFound creates a GetIconNotFound with default headers values
func NewGetIconNotFound() *GetIconNotFound {
	return &GetIconNotFound{}
}

/* GetIconNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetIconNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetIconNotFound) Error() string {
	return fmt.Sprintf("[GET /icons/{digest}][%d] getIconNotFound  %+v", 404, o.Payload)
}
func (o *GetIconNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetIconNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIconInternalServerError creates a GetIconInternalServerError with default headers values
func NewGetIconInternalServerError() *GetIconInternalServerError {
	return &GetIconInternalServerError{}
}

/* GetIconInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetIconInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetIconInternalServerError) Error() string {
	return fmt.Sprintf("[GET /icons/{digest}][%d] getIconInternalServerError  %+v", 500, o.Payload)
}
func (o *GetIconInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetIconInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
