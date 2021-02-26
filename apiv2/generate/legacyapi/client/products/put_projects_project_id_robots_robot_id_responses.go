// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PutProjectsProjectIDRobotsRobotIDReader is a Reader for the PutProjectsProjectIDRobotsRobotID structure.
type PutProjectsProjectIDRobotsRobotIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutProjectsProjectIDRobotsRobotIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutProjectsProjectIDRobotsRobotIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewPutProjectsProjectIDRobotsRobotIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutProjectsProjectIDRobotsRobotIDOK creates a PutProjectsProjectIDRobotsRobotIDOK with default headers values
func NewPutProjectsProjectIDRobotsRobotIDOK() *PutProjectsProjectIDRobotsRobotIDOK {
	return &PutProjectsProjectIDRobotsRobotIDOK{}
}

/* PutProjectsProjectIDRobotsRobotIDOK describes a response with status code 200, with default header values.

Robot account has been modified success.
*/
type PutProjectsProjectIDRobotsRobotIDOK struct {
}

func (o *PutProjectsProjectIDRobotsRobotIDOK) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_id}/robots/{robot_id}][%d] putProjectsProjectIdRobotsRobotIdOK ", 200)
}

func (o *PutProjectsProjectIDRobotsRobotIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutProjectsProjectIDRobotsRobotIDInternalServerError creates a PutProjectsProjectIDRobotsRobotIDInternalServerError with default headers values
func NewPutProjectsProjectIDRobotsRobotIDInternalServerError() *PutProjectsProjectIDRobotsRobotIDInternalServerError {
	return &PutProjectsProjectIDRobotsRobotIDInternalServerError{}
}

/* PutProjectsProjectIDRobotsRobotIDInternalServerError describes a response with status code 500, with default header values.

Unexpected generate errors.
*/
type PutProjectsProjectIDRobotsRobotIDInternalServerError struct {
}

func (o *PutProjectsProjectIDRobotsRobotIDInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_id}/robots/{robot_id}][%d] putProjectsProjectIdRobotsRobotIdInternalServerError ", 500)
}

func (o *PutProjectsProjectIDRobotsRobotIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}