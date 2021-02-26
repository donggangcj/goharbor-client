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

// GetProjectsProjectIDRobotsReader is a Reader for the GetProjectsProjectIDRobots structure.
type GetProjectsProjectIDRobotsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectsProjectIDRobotsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectsProjectIDRobotsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetProjectsProjectIDRobotsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetProjectsProjectIDRobotsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetProjectsProjectIDRobotsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetProjectsProjectIDRobotsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetProjectsProjectIDRobotsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetProjectsProjectIDRobotsOK creates a GetProjectsProjectIDRobotsOK with default headers values
func NewGetProjectsProjectIDRobotsOK() *GetProjectsProjectIDRobotsOK {
	return &GetProjectsProjectIDRobotsOK{}
}

/* GetProjectsProjectIDRobotsOK describes a response with status code 200, with default header values.

Get project robot accounts successfully.
*/
type GetProjectsProjectIDRobotsOK struct {
	Payload []*legacy.RobotAccount
}

func (o *GetProjectsProjectIDRobotsOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/robots][%d] getProjectsProjectIdRobotsOK  %+v", 200, o.Payload)
}
func (o *GetProjectsProjectIDRobotsOK) GetPayload() []*legacy.RobotAccount {
	return o.Payload
}

func (o *GetProjectsProjectIDRobotsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectsProjectIDRobotsBadRequest creates a GetProjectsProjectIDRobotsBadRequest with default headers values
func NewGetProjectsProjectIDRobotsBadRequest() *GetProjectsProjectIDRobotsBadRequest {
	return &GetProjectsProjectIDRobotsBadRequest{}
}

/* GetProjectsProjectIDRobotsBadRequest describes a response with status code 400, with default header values.

The project id is invalid.
*/
type GetProjectsProjectIDRobotsBadRequest struct {
}

func (o *GetProjectsProjectIDRobotsBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/robots][%d] getProjectsProjectIdRobotsBadRequest ", 400)
}

func (o *GetProjectsProjectIDRobotsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectsProjectIDRobotsUnauthorized creates a GetProjectsProjectIDRobotsUnauthorized with default headers values
func NewGetProjectsProjectIDRobotsUnauthorized() *GetProjectsProjectIDRobotsUnauthorized {
	return &GetProjectsProjectIDRobotsUnauthorized{}
}

/* GetProjectsProjectIDRobotsUnauthorized describes a response with status code 401, with default header values.

User need to log in first.
*/
type GetProjectsProjectIDRobotsUnauthorized struct {
}

func (o *GetProjectsProjectIDRobotsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/robots][%d] getProjectsProjectIdRobotsUnauthorized ", 401)
}

func (o *GetProjectsProjectIDRobotsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectsProjectIDRobotsForbidden creates a GetProjectsProjectIDRobotsForbidden with default headers values
func NewGetProjectsProjectIDRobotsForbidden() *GetProjectsProjectIDRobotsForbidden {
	return &GetProjectsProjectIDRobotsForbidden{}
}

/* GetProjectsProjectIDRobotsForbidden describes a response with status code 403, with default header values.

User in session does not have permission to the project.
*/
type GetProjectsProjectIDRobotsForbidden struct {
}

func (o *GetProjectsProjectIDRobotsForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/robots][%d] getProjectsProjectIdRobotsForbidden ", 403)
}

func (o *GetProjectsProjectIDRobotsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectsProjectIDRobotsNotFound creates a GetProjectsProjectIDRobotsNotFound with default headers values
func NewGetProjectsProjectIDRobotsNotFound() *GetProjectsProjectIDRobotsNotFound {
	return &GetProjectsProjectIDRobotsNotFound{}
}

/* GetProjectsProjectIDRobotsNotFound describes a response with status code 404, with default header values.

Project ID does not exist.
*/
type GetProjectsProjectIDRobotsNotFound struct {
}

func (o *GetProjectsProjectIDRobotsNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/robots][%d] getProjectsProjectIdRobotsNotFound ", 404)
}

func (o *GetProjectsProjectIDRobotsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectsProjectIDRobotsInternalServerError creates a GetProjectsProjectIDRobotsInternalServerError with default headers values
func NewGetProjectsProjectIDRobotsInternalServerError() *GetProjectsProjectIDRobotsInternalServerError {
	return &GetProjectsProjectIDRobotsInternalServerError{}
}

/* GetProjectsProjectIDRobotsInternalServerError describes a response with status code 500, with default header values.

Unexpected internal errors.
*/
type GetProjectsProjectIDRobotsInternalServerError struct {
}

func (o *GetProjectsProjectIDRobotsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/robots][%d] getProjectsProjectIdRobotsInternalServerError ", 500)
}

func (o *GetProjectsProjectIDRobotsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
