package hooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewHooksGetAllParams creates a new HooksGetAllParams object
// with the default values initialized.
func NewHooksGetAllParams() *HooksGetAllParams {

	return &HooksGetAllParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewHooksGetAllParamsWithTimeout creates a new HooksGetAllParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewHooksGetAllParamsWithTimeout(timeout time.Duration) *HooksGetAllParams {

	return &HooksGetAllParams{

		timeout: timeout,
	}
}

// NewHooksGetAllParamsWithContext creates a new HooksGetAllParams object
// with the default values initialized, and the ability to set a context for a request
func NewHooksGetAllParamsWithContext(ctx context.Context) *HooksGetAllParams {

	return &HooksGetAllParams{

		Context: ctx,
	}
}

// NewHooksGetAllParamsWithHTTPClient creates a new HooksGetAllParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewHooksGetAllParamsWithHTTPClient(client *http.Client) *HooksGetAllParams {

	return &HooksGetAllParams{
		HTTPClient: client,
	}
}

/*HooksGetAllParams contains all the parameters to send to the API endpoint
for the hooks get all operation typically these are written to a http.Request
*/
type HooksGetAllParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the hooks get all params
func (o *HooksGetAllParams) WithTimeout(timeout time.Duration) *HooksGetAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the hooks get all params
func (o *HooksGetAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the hooks get all params
func (o *HooksGetAllParams) WithContext(ctx context.Context) *HooksGetAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the hooks get all params
func (o *HooksGetAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the hooks get all params
func (o *HooksGetAllParams) WithHTTPClient(client *http.Client) *HooksGetAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the hooks get all params
func (o *HooksGetAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *HooksGetAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
