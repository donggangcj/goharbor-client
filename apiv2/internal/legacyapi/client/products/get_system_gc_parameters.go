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
)

// NewGetSystemGcParams creates a new GetSystemGcParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSystemGcParams() *GetSystemGcParams {
	return &GetSystemGcParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSystemGcParamsWithTimeout creates a new GetSystemGcParams object
// with the ability to set a timeout on a request.
func NewGetSystemGcParamsWithTimeout(timeout time.Duration) *GetSystemGcParams {
	return &GetSystemGcParams{
		timeout: timeout,
	}
}

// NewGetSystemGcParamsWithContext creates a new GetSystemGcParams object
// with the ability to set a context for a request.
func NewGetSystemGcParamsWithContext(ctx context.Context) *GetSystemGcParams {
	return &GetSystemGcParams{
		Context: ctx,
	}
}

// NewGetSystemGcParamsWithHTTPClient creates a new GetSystemGcParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSystemGcParamsWithHTTPClient(client *http.Client) *GetSystemGcParams {
	return &GetSystemGcParams{
		HTTPClient: client,
	}
}

/* GetSystemGcParams contains all the parameters to send to the API endpoint
   for the get system gc operation.

   Typically these are written to a http.Request.
*/
type GetSystemGcParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get system gc params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSystemGcParams) WithDefaults() *GetSystemGcParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get system gc params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSystemGcParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get system gc params
func (o *GetSystemGcParams) WithTimeout(timeout time.Duration) *GetSystemGcParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get system gc params
func (o *GetSystemGcParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get system gc params
func (o *GetSystemGcParams) WithContext(ctx context.Context) *GetSystemGcParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get system gc params
func (o *GetSystemGcParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get system gc params
func (o *GetSystemGcParams) WithHTTPClient(client *http.Client) *GetSystemGcParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get system gc params
func (o *GetSystemGcParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetSystemGcParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
