// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/donggangcj/goharbor-client/v3/apiv2/model"
)

// DeleteProjectReader is a Reader for the DeleteProject structure.
type DeleteProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteProjectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewDeleteProjectPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteProjectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteProjectOK creates a DeleteProjectOK with default headers values
func NewDeleteProjectOK() *DeleteProjectOK {
	return &DeleteProjectOK{}
}

/* DeleteProjectOK describes a response with status code 200, with default header values.

Success
*/
type DeleteProjectOK struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *DeleteProjectOK) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_id}][%d] deleteProjectOK ", 200)
}

func (o *DeleteProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewDeleteProjectBadRequest creates a DeleteProjectBadRequest with default headers values
func NewDeleteProjectBadRequest() *DeleteProjectBadRequest {
	return &DeleteProjectBadRequest{}
}

/* DeleteProjectBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type DeleteProjectBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteProjectBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_id}][%d] deleteProjectBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteProjectBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteProjectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteProjectForbidden creates a DeleteProjectForbidden with default headers values
func NewDeleteProjectForbidden() *DeleteProjectForbidden {
	return &DeleteProjectForbidden{}
}

/* DeleteProjectForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteProjectForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteProjectForbidden) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_id}][%d] deleteProjectForbidden  %+v", 403, o.Payload)
}
func (o *DeleteProjectForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteProjectNotFound creates a DeleteProjectNotFound with default headers values
func NewDeleteProjectNotFound() *DeleteProjectNotFound {
	return &DeleteProjectNotFound{}
}

/* DeleteProjectNotFound describes a response with status code 404, with default header values.

Not found
*/
type DeleteProjectNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteProjectNotFound) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_id}][%d] deleteProjectNotFound  %+v", 404, o.Payload)
}
func (o *DeleteProjectNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteProjectPreconditionFailed creates a DeleteProjectPreconditionFailed with default headers values
func NewDeleteProjectPreconditionFailed() *DeleteProjectPreconditionFailed {
	return &DeleteProjectPreconditionFailed{}
}

/* DeleteProjectPreconditionFailed describes a response with status code 412, with default header values.

Precondition failed
*/
type DeleteProjectPreconditionFailed struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteProjectPreconditionFailed) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_id}][%d] deleteProjectPreconditionFailed  %+v", 412, o.Payload)
}
func (o *DeleteProjectPreconditionFailed) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteProjectPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteProjectInternalServerError creates a DeleteProjectInternalServerError with default headers values
func NewDeleteProjectInternalServerError() *DeleteProjectInternalServerError {
	return &DeleteProjectInternalServerError{}
}

/* DeleteProjectInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type DeleteProjectInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteProjectInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_id}][%d] deleteProjectInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteProjectInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteProjectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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