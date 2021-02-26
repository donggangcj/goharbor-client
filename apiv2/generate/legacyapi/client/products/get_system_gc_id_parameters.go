// Code generated by go-swagger; DO NOT EDIT.

package products

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

// NewGetSystemGcIDParams creates a new GetSystemGcIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSystemGcIDParams() *GetSystemGcIDParams {
	return &GetSystemGcIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSystemGcIDParamsWithTimeout creates a new GetSystemGcIDParams object
// with the ability to set a timeout on a request.
func NewGetSystemGcIDParamsWithTimeout(timeout time.Duration) *GetSystemGcIDParams {
	return &GetSystemGcIDParams{
		timeout: timeout,
	}
}

// NewGetSystemGcIDParamsWithContext creates a new GetSystemGcIDParams object
// with the ability to set a context for a request.
func NewGetSystemGcIDParamsWithContext(ctx context.Context) *GetSystemGcIDParams {
	return &GetSystemGcIDParams{
		Context: ctx,
	}
}

// NewGetSystemGcIDParamsWithHTTPClient creates a new GetSystemGcIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSystemGcIDParamsWithHTTPClient(client *http.Client) *GetSystemGcIDParams {
	return &GetSystemGcIDParams{
		HTTPClient: client,
	}
}

/* GetSystemGcIDParams contains all the parameters to send to the API endpoint
   for the get system gc ID operation.

   Typically these are written to a http.Request.
*/
type GetSystemGcIDParams struct {

	/* ID.

	   Relevant job ID

	   Format: int64
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get system gc ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSystemGcIDParams) WithDefaults() *GetSystemGcIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get system gc ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSystemGcIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get system gc ID params
func (o *GetSystemGcIDParams) WithTimeout(timeout time.Duration) *GetSystemGcIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get system gc ID params
func (o *GetSystemGcIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get system gc ID params
func (o *GetSystemGcIDParams) WithContext(ctx context.Context) *GetSystemGcIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get system gc ID params
func (o *GetSystemGcIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get system gc ID params
func (o *GetSystemGcIDParams) WithHTTPClient(client *http.Client) *GetSystemGcIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get system gc ID params
func (o *GetSystemGcIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get system gc ID params
func (o *GetSystemGcIDParams) WithID(id int64) *GetSystemGcIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get system gc ID params
func (o *GetSystemGcIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetSystemGcIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}