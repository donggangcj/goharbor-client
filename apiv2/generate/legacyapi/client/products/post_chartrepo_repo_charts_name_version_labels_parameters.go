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

	"github.com/donggangcj/goharbor-client/v3/apiv2/model/legacy"
)

// NewPostChartrepoRepoChartsNameVersionLabelsParams creates a new PostChartrepoRepoChartsNameVersionLabelsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostChartrepoRepoChartsNameVersionLabelsParams() *PostChartrepoRepoChartsNameVersionLabelsParams {
	return &PostChartrepoRepoChartsNameVersionLabelsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostChartrepoRepoChartsNameVersionLabelsParamsWithTimeout creates a new PostChartrepoRepoChartsNameVersionLabelsParams object
// with the ability to set a timeout on a request.
func NewPostChartrepoRepoChartsNameVersionLabelsParamsWithTimeout(timeout time.Duration) *PostChartrepoRepoChartsNameVersionLabelsParams {
	return &PostChartrepoRepoChartsNameVersionLabelsParams{
		timeout: timeout,
	}
}

// NewPostChartrepoRepoChartsNameVersionLabelsParamsWithContext creates a new PostChartrepoRepoChartsNameVersionLabelsParams object
// with the ability to set a context for a request.
func NewPostChartrepoRepoChartsNameVersionLabelsParamsWithContext(ctx context.Context) *PostChartrepoRepoChartsNameVersionLabelsParams {
	return &PostChartrepoRepoChartsNameVersionLabelsParams{
		Context: ctx,
	}
}

// NewPostChartrepoRepoChartsNameVersionLabelsParamsWithHTTPClient creates a new PostChartrepoRepoChartsNameVersionLabelsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostChartrepoRepoChartsNameVersionLabelsParamsWithHTTPClient(client *http.Client) *PostChartrepoRepoChartsNameVersionLabelsParams {
	return &PostChartrepoRepoChartsNameVersionLabelsParams{
		HTTPClient: client,
	}
}

/* PostChartrepoRepoChartsNameVersionLabelsParams contains all the parameters to send to the API endpoint
   for the post chartrepo repo charts name version labels operation.

   Typically these are written to a http.Request.
*/
type PostChartrepoRepoChartsNameVersionLabelsParams struct {

	/* Label.

	   The label being marked to the chart version
	*/
	Label *legacy.Label

	/* Name.

	   The chart name
	*/
	Name string

	/* Repo.

	   The project name
	*/
	Repo string

	/* Version.

	   The chart version
	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post chartrepo repo charts name version labels params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostChartrepoRepoChartsNameVersionLabelsParams) WithDefaults() *PostChartrepoRepoChartsNameVersionLabelsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post chartrepo repo charts name version labels params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostChartrepoRepoChartsNameVersionLabelsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post chartrepo repo charts name version labels params
func (o *PostChartrepoRepoChartsNameVersionLabelsParams) WithTimeout(timeout time.Duration) *PostChartrepoRepoChartsNameVersionLabelsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post chartrepo repo charts name version labels params
func (o *PostChartrepoRepoChartsNameVersionLabelsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post chartrepo repo charts name version labels params
func (o *PostChartrepoRepoChartsNameVersionLabelsParams) WithContext(ctx context.Context) *PostChartrepoRepoChartsNameVersionLabelsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post chartrepo repo charts name version labels params
func (o *PostChartrepoRepoChartsNameVersionLabelsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post chartrepo repo charts name version labels params
func (o *PostChartrepoRepoChartsNameVersionLabelsParams) WithHTTPClient(client *http.Client) *PostChartrepoRepoChartsNameVersionLabelsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post chartrepo repo charts name version labels params
func (o *PostChartrepoRepoChartsNameVersionLabelsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLabel adds the label to the post chartrepo repo charts name version labels params
func (o *PostChartrepoRepoChartsNameVersionLabelsParams) WithLabel(label *legacy.Label) *PostChartrepoRepoChartsNameVersionLabelsParams {
	o.SetLabel(label)
	return o
}

// SetLabel adds the label to the post chartrepo repo charts name version labels params
func (o *PostChartrepoRepoChartsNameVersionLabelsParams) SetLabel(label *legacy.Label) {
	o.Label = label
}

// WithName adds the name to the post chartrepo repo charts name version labels params
func (o *PostChartrepoRepoChartsNameVersionLabelsParams) WithName(name string) *PostChartrepoRepoChartsNameVersionLabelsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the post chartrepo repo charts name version labels params
func (o *PostChartrepoRepoChartsNameVersionLabelsParams) SetName(name string) {
	o.Name = name
}

// WithRepo adds the repo to the post chartrepo repo charts name version labels params
func (o *PostChartrepoRepoChartsNameVersionLabelsParams) WithRepo(repo string) *PostChartrepoRepoChartsNameVersionLabelsParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the post chartrepo repo charts name version labels params
func (o *PostChartrepoRepoChartsNameVersionLabelsParams) SetRepo(repo string) {
	o.Repo = repo
}

// WithVersion adds the version to the post chartrepo repo charts name version labels params
func (o *PostChartrepoRepoChartsNameVersionLabelsParams) WithVersion(version string) *PostChartrepoRepoChartsNameVersionLabelsParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the post chartrepo repo charts name version labels params
func (o *PostChartrepoRepoChartsNameVersionLabelsParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *PostChartrepoRepoChartsNameVersionLabelsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Label != nil {
		if err := r.SetBodyParam(o.Label); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param repo
	if err := r.SetPathParam("repo", o.Repo); err != nil {
		return err
	}

	// path param version
	if err := r.SetPathParam("version", o.Version); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}