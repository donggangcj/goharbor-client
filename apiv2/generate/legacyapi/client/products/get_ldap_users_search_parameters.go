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

// NewGetLdapUsersSearchParams creates a new GetLdapUsersSearchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLdapUsersSearchParams() *GetLdapUsersSearchParams {
	return &GetLdapUsersSearchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLdapUsersSearchParamsWithTimeout creates a new GetLdapUsersSearchParams object
// with the ability to set a timeout on a request.
func NewGetLdapUsersSearchParamsWithTimeout(timeout time.Duration) *GetLdapUsersSearchParams {
	return &GetLdapUsersSearchParams{
		timeout: timeout,
	}
}

// NewGetLdapUsersSearchParamsWithContext creates a new GetLdapUsersSearchParams object
// with the ability to set a context for a request.
func NewGetLdapUsersSearchParamsWithContext(ctx context.Context) *GetLdapUsersSearchParams {
	return &GetLdapUsersSearchParams{
		Context: ctx,
	}
}

// NewGetLdapUsersSearchParamsWithHTTPClient creates a new GetLdapUsersSearchParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLdapUsersSearchParamsWithHTTPClient(client *http.Client) *GetLdapUsersSearchParams {
	return &GetLdapUsersSearchParams{
		HTTPClient: client,
	}
}

/* GetLdapUsersSearchParams contains all the parameters to send to the API endpoint
   for the get ldap users search operation.

   Typically these are written to a http.Request.
*/
type GetLdapUsersSearchParams struct {

	/* Username.

	   Registered user ID
	*/
	Username *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get ldap users search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLdapUsersSearchParams) WithDefaults() *GetLdapUsersSearchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get ldap users search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLdapUsersSearchParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get ldap users search params
func (o *GetLdapUsersSearchParams) WithTimeout(timeout time.Duration) *GetLdapUsersSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get ldap users search params
func (o *GetLdapUsersSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get ldap users search params
func (o *GetLdapUsersSearchParams) WithContext(ctx context.Context) *GetLdapUsersSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get ldap users search params
func (o *GetLdapUsersSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get ldap users search params
func (o *GetLdapUsersSearchParams) WithHTTPClient(client *http.Client) *GetLdapUsersSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get ldap users search params
func (o *GetLdapUsersSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUsername adds the username to the get ldap users search params
func (o *GetLdapUsersSearchParams) WithUsername(username *string) *GetLdapUsersSearchParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the get ldap users search params
func (o *GetLdapUsersSearchParams) SetUsername(username *string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *GetLdapUsersSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Username != nil {

		// query param username
		var qrUsername string

		if o.Username != nil {
			qrUsername = *o.Username
		}
		qUsername := qrUsername
		if qUsername != "" {

			if err := r.SetQueryParam("username", qUsername); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}