// Code generated by go-swagger; DO NOT EDIT.

package account_export

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

// NewGetAccountExportIDParams creates a new GetAccountExportIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAccountExportIDParams() *GetAccountExportIDParams {
	return &GetAccountExportIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccountExportIDParamsWithTimeout creates a new GetAccountExportIDParams object
// with the ability to set a timeout on a request.
func NewGetAccountExportIDParamsWithTimeout(timeout time.Duration) *GetAccountExportIDParams {
	return &GetAccountExportIDParams{
		timeout: timeout,
	}
}

// NewGetAccountExportIDParamsWithContext creates a new GetAccountExportIDParams object
// with the ability to set a context for a request.
func NewGetAccountExportIDParamsWithContext(ctx context.Context) *GetAccountExportIDParams {
	return &GetAccountExportIDParams{
		Context: ctx,
	}
}

// NewGetAccountExportIDParamsWithHTTPClient creates a new GetAccountExportIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAccountExportIDParamsWithHTTPClient(client *http.Client) *GetAccountExportIDParams {
	return &GetAccountExportIDParams{
		HTTPClient: client,
	}
}

/* GetAccountExportIDParams contains all the parameters to send to the API endpoint
   for the get account export Id operation.

   Typically these are written to a http.Request.
*/
type GetAccountExportIDParams struct {

	/* ExcludeFields.

	   A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
	*/
	ExcludeFields []string

	/* ExportID.

	   The unique id for the account export.
	*/
	ExportID string

	/* Fields.

	   A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
	*/
	Fields []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get account export Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAccountExportIDParams) WithDefaults() *GetAccountExportIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get account export Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAccountExportIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get account export Id params
func (o *GetAccountExportIDParams) WithTimeout(timeout time.Duration) *GetAccountExportIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get account export Id params
func (o *GetAccountExportIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get account export Id params
func (o *GetAccountExportIDParams) WithContext(ctx context.Context) *GetAccountExportIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get account export Id params
func (o *GetAccountExportIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get account export Id params
func (o *GetAccountExportIDParams) WithHTTPClient(client *http.Client) *GetAccountExportIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get account export Id params
func (o *GetAccountExportIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExcludeFields adds the excludeFields to the get account export Id params
func (o *GetAccountExportIDParams) WithExcludeFields(excludeFields []string) *GetAccountExportIDParams {
	o.SetExcludeFields(excludeFields)
	return o
}

// SetExcludeFields adds the excludeFields to the get account export Id params
func (o *GetAccountExportIDParams) SetExcludeFields(excludeFields []string) {
	o.ExcludeFields = excludeFields
}

// WithExportID adds the exportID to the get account export Id params
func (o *GetAccountExportIDParams) WithExportID(exportID string) *GetAccountExportIDParams {
	o.SetExportID(exportID)
	return o
}

// SetExportID adds the exportId to the get account export Id params
func (o *GetAccountExportIDParams) SetExportID(exportID string) {
	o.ExportID = exportID
}

// WithFields adds the fields to the get account export Id params
func (o *GetAccountExportIDParams) WithFields(fields []string) *GetAccountExportIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get account export Id params
func (o *GetAccountExportIDParams) SetFields(fields []string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccountExportIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ExcludeFields != nil {

		// binding items for exclude_fields
		joinedExcludeFields := o.bindParamExcludeFields(reg)

		// query array param exclude_fields
		if err := r.SetQueryParam("exclude_fields", joinedExcludeFields...); err != nil {
			return err
		}
	}

	// path param export_id
	if err := r.SetPathParam("export_id", o.ExportID); err != nil {
		return err
	}

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetAccountExportID binds the parameter exclude_fields
func (o *GetAccountExportIDParams) bindParamExcludeFields(formats strfmt.Registry) []string {
	excludeFieldsIR := o.ExcludeFields

	var excludeFieldsIC []string
	for _, excludeFieldsIIR := range excludeFieldsIR { // explode []string

		excludeFieldsIIV := excludeFieldsIIR // string as string
		excludeFieldsIC = append(excludeFieldsIC, excludeFieldsIIV)
	}

	// items.CollectionFormat: "csv"
	excludeFieldsIS := swag.JoinByFormat(excludeFieldsIC, "csv")

	return excludeFieldsIS
}

// bindParamGetAccountExportID binds the parameter fields
func (o *GetAccountExportIDParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.Fields

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsIS := swag.JoinByFormat(fieldsIC, "csv")

	return fieldsIS
}