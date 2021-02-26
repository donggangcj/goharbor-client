// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetProjectDeletableParams creates a new GetProjectDeletableParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProjectDeletableParams() *GetProjectDeletableParams {
	return &GetProjectDeletableParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProjectDeletableParamsWithTimeout creates a new GetProjectDeletableParams object
// with the ability to set a timeout on a request.
func NewGetProjectDeletableParamsWithTimeout(timeout time.Duration) *GetProjectDeletableParams {
	return &GetProjectDeletableParams{
		timeout: timeout,
	}
}

// NewGetProjectDeletableParamsWithContext creates a new GetProjectDeletableParams object
// with the ability to set a context for a request.
func NewGetProjectDeletableParamsWithContext(ctx context.Context) *GetProjectDeletableParams {
	return &GetProjectDeletableParams{
		Context: ctx,
	}
}

// NewGetProjectDeletableParamsWithHTTPClient creates a new GetProjectDeletableParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProjectDeletableParamsWithHTTPClient(client *http.Client) *GetProjectDeletableParams {
	return &GetProjectDeletableParams{
		HTTPClient: client,
	}
}

/* GetProjectDeletableParams contains all the parameters to send to the API endpoint
   for the get project deletable operation.

   Typically these are written to a http.Request.
*/
type GetProjectDeletableParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string

	/* ProjectID.

	   The ID of the project

	   Format: int64
	*/
	ProjectID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get project deletable params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectDeletableParams) WithDefaults() *GetProjectDeletableParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get project deletable params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectDeletableParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get project deletable params
func (o *GetProjectDeletableParams) WithTimeout(timeout time.Duration) *GetProjectDeletableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get project deletable params
func (o *GetProjectDeletableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get project deletable params
func (o *GetProjectDeletableParams) WithContext(ctx context.Context) *GetProjectDeletableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get project deletable params
func (o *GetProjectDeletableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get project deletable params
func (o *GetProjectDeletableParams) WithHTTPClient(client *http.Client) *GetProjectDeletableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get project deletable params
func (o *GetProjectDeletableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the get project deletable params
func (o *GetProjectDeletableParams) WithXRequestID(xRequestID *string) *GetProjectDeletableParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get project deletable params
func (o *GetProjectDeletableParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithProjectID adds the projectID to the get project deletable params
func (o *GetProjectDeletableParams) WithProjectID(projectID int64) *GetProjectDeletableParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get project deletable params
func (o *GetProjectDeletableParams) SetProjectID(projectID int64) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *GetProjectDeletableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}
	}

	// path param project_id
	if err := r.SetPathParam("project_id", swag.FormatInt64(o.ProjectID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}