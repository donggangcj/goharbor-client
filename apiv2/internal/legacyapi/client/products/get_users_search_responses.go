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

// GetUsersSearchReader is a Reader for the GetUsersSearch structure.
type GetUsersSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsersSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsersSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetUsersSearchInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUsersSearchOK creates a GetUsersSearchOK with default headers values
func NewGetUsersSearchOK() *GetUsersSearchOK {
	return &GetUsersSearchOK{}
}

/* GetUsersSearchOK describes a response with status code 200, with default header values.

Search users by username, email successfully.
*/
type GetUsersSearchOK struct {
	Payload []*legacy.UserSearch
}

func (o *GetUsersSearchOK) Error() string {
	return fmt.Sprintf("[GET /users/search][%d] getUsersSearchOK  %+v", 200, o.Payload)
}
func (o *GetUsersSearchOK) GetPayload() []*legacy.UserSearch {
	return o.Payload
}

func (o *GetUsersSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersSearchInternalServerError creates a GetUsersSearchInternalServerError with default headers values
func NewGetUsersSearchInternalServerError() *GetUsersSearchInternalServerError {
	return &GetUsersSearchInternalServerError{}
}

/* GetUsersSearchInternalServerError describes a response with status code 500, with default header values.

Unexpected internal errors.
*/
type GetUsersSearchInternalServerError struct {
}

func (o *GetUsersSearchInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users/search][%d] getUsersSearchInternalServerError ", 500)
}

func (o *GetUsersSearchInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
