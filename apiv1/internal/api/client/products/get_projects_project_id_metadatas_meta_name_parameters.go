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

// NewGetProjectsProjectIDMetadatasMetaNameParams creates a new GetProjectsProjectIDMetadatasMetaNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProjectsProjectIDMetadatasMetaNameParams() *GetProjectsProjectIDMetadatasMetaNameParams {
	return &GetProjectsProjectIDMetadatasMetaNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProjectsProjectIDMetadatasMetaNameParamsWithTimeout creates a new GetProjectsProjectIDMetadatasMetaNameParams object
// with the ability to set a timeout on a request.
func NewGetProjectsProjectIDMetadatasMetaNameParamsWithTimeout(timeout time.Duration) *GetProjectsProjectIDMetadatasMetaNameParams {
	return &GetProjectsProjectIDMetadatasMetaNameParams{
		timeout: timeout,
	}
}

// NewGetProjectsProjectIDMetadatasMetaNameParamsWithContext creates a new GetProjectsProjectIDMetadatasMetaNameParams object
// with the ability to set a context for a request.
func NewGetProjectsProjectIDMetadatasMetaNameParamsWithContext(ctx context.Context) *GetProjectsProjectIDMetadatasMetaNameParams {
	return &GetProjectsProjectIDMetadatasMetaNameParams{
		Context: ctx,
	}
}

// NewGetProjectsProjectIDMetadatasMetaNameParamsWithHTTPClient creates a new GetProjectsProjectIDMetadatasMetaNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProjectsProjectIDMetadatasMetaNameParamsWithHTTPClient(client *http.Client) *GetProjectsProjectIDMetadatasMetaNameParams {
	return &GetProjectsProjectIDMetadatasMetaNameParams{
		HTTPClient: client,
	}
}

/* GetProjectsProjectIDMetadatasMetaNameParams contains all the parameters to send to the API endpoint
   for the get projects project ID metadatas meta name operation.

   Typically these are written to a http.Request.
*/
type GetProjectsProjectIDMetadatasMetaNameParams struct {

	/* MetaName.

	   The name of metadat.
	*/
	MetaName string

	/* ProjectID.

	   Project ID for filtering results.

	   Format: int64
	*/
	ProjectID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get projects project ID metadatas meta name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectsProjectIDMetadatasMetaNameParams) WithDefaults() *GetProjectsProjectIDMetadatasMetaNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get projects project ID metadatas meta name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectsProjectIDMetadatasMetaNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get projects project ID metadatas meta name params
func (o *GetProjectsProjectIDMetadatasMetaNameParams) WithTimeout(timeout time.Duration) *GetProjectsProjectIDMetadatasMetaNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get projects project ID metadatas meta name params
func (o *GetProjectsProjectIDMetadatasMetaNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get projects project ID metadatas meta name params
func (o *GetProjectsProjectIDMetadatasMetaNameParams) WithContext(ctx context.Context) *GetProjectsProjectIDMetadatasMetaNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get projects project ID metadatas meta name params
func (o *GetProjectsProjectIDMetadatasMetaNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get projects project ID metadatas meta name params
func (o *GetProjectsProjectIDMetadatasMetaNameParams) WithHTTPClient(client *http.Client) *GetProjectsProjectIDMetadatasMetaNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get projects project ID metadatas meta name params
func (o *GetProjectsProjectIDMetadatasMetaNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMetaName adds the metaName to the get projects project ID metadatas meta name params
func (o *GetProjectsProjectIDMetadatasMetaNameParams) WithMetaName(metaName string) *GetProjectsProjectIDMetadatasMetaNameParams {
	o.SetMetaName(metaName)
	return o
}

// SetMetaName adds the metaName to the get projects project ID metadatas meta name params
func (o *GetProjectsProjectIDMetadatasMetaNameParams) SetMetaName(metaName string) {
	o.MetaName = metaName
}

// WithProjectID adds the projectID to the get projects project ID metadatas meta name params
func (o *GetProjectsProjectIDMetadatasMetaNameParams) WithProjectID(projectID int64) *GetProjectsProjectIDMetadatasMetaNameParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get projects project ID metadatas meta name params
func (o *GetProjectsProjectIDMetadatasMetaNameParams) SetProjectID(projectID int64) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *GetProjectsProjectIDMetadatasMetaNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param meta_name
	if err := r.SetPathParam("meta_name", o.MetaName); err != nil {
		return err
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
