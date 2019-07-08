package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewFindParams creates a new FindParams object
// with the default values initialized.
func NewFindParams() *FindParams {
	var ()
	return &FindParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindParamsWithTimeout creates a new FindParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindParamsWithTimeout(timeout time.Duration) *FindParams {
	var ()
	return &FindParams{

		timeout: timeout,
	}
}

// NewFindParamsWithContext creates a new FindParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindParamsWithContext(ctx context.Context) *FindParams {
	var ()
	return &FindParams{

		Context: ctx,
	}
}

/*FindParams contains all the parameters to send to the API endpoint
for the find operation typically these are written to a http.Request
*/
type FindParams struct {

	/*XRateLimit*/
	XRateLimit int32
	/*Limit*/
	Limit int32
	/*Tags*/
	Tags []int32

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the find params
func (o *FindParams) WithTimeout(timeout time.Duration) *FindParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find params
func (o *FindParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find params
func (o *FindParams) WithContext(ctx context.Context) *FindParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find params
func (o *FindParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithXRateLimit adds the xRateLimit to the find params
func (o *FindParams) WithXRateLimit(xRateLimit int32) *FindParams {
	o.SetXRateLimit(xRateLimit)
	return o
}

// SetXRateLimit adds the xRateLimit to the find params
func (o *FindParams) SetXRateLimit(xRateLimit int32) {
	o.XRateLimit = xRateLimit
}

// WithLimit adds the limit to the find params
func (o *FindParams) WithLimit(limit int32) *FindParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the find params
func (o *FindParams) SetLimit(limit int32) {
	o.Limit = limit
}

// WithTags adds the tags to the find params
func (o *FindParams) WithTags(tags []int32) *FindParams {
	o.SetTags(tags)
	return o
}

// SetTags adds the tags to the find params
func (o *FindParams) SetTags(tags []int32) {
	o.Tags = tags
}

// WriteToRequest writes these params to a swagger request
func (o *FindParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// header param X-Rate-Limit
	if err := r.SetHeaderParam("X-Rate-Limit", swag.FormatInt32(o.XRateLimit)); err != nil {
		return err
	}

	// form param limit
	frLimit := o.Limit
	fLimit := swag.FormatInt32(frLimit)
	if err := r.SetFormParam("limit", fLimit); err != nil {
		return err
	}

	var valuesTags []string
	for _, v := range o.Tags {
		valuesTags = append(valuesTags, swag.FormatInt32(v))
	}

	joinedTags := swag.JoinByFormat(valuesTags, "multi")
	// form array param tags
	if err := r.SetFormParam("tags", joinedTags...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
