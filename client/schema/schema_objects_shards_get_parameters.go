//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2024 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by go-swagger; DO NOT EDIT.

package schema

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

// NewSchemaObjectsShardsGetParams creates a new SchemaObjectsShardsGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSchemaObjectsShardsGetParams() *SchemaObjectsShardsGetParams {
	return &SchemaObjectsShardsGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSchemaObjectsShardsGetParamsWithTimeout creates a new SchemaObjectsShardsGetParams object
// with the ability to set a timeout on a request.
func NewSchemaObjectsShardsGetParamsWithTimeout(timeout time.Duration) *SchemaObjectsShardsGetParams {
	return &SchemaObjectsShardsGetParams{
		timeout: timeout,
	}
}

// NewSchemaObjectsShardsGetParamsWithContext creates a new SchemaObjectsShardsGetParams object
// with the ability to set a context for a request.
func NewSchemaObjectsShardsGetParamsWithContext(ctx context.Context) *SchemaObjectsShardsGetParams {
	return &SchemaObjectsShardsGetParams{
		Context: ctx,
	}
}

// NewSchemaObjectsShardsGetParamsWithHTTPClient creates a new SchemaObjectsShardsGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSchemaObjectsShardsGetParamsWithHTTPClient(client *http.Client) *SchemaObjectsShardsGetParams {
	return &SchemaObjectsShardsGetParams{
		HTTPClient: client,
	}
}

/*
SchemaObjectsShardsGetParams contains all the parameters to send to the API endpoint

	for the schema objects shards get operation.

	Typically these are written to a http.Request.
*/
type SchemaObjectsShardsGetParams struct {

	// ClassName.
	ClassName string

	// Tenant.
	Tenant *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the schema objects shards get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SchemaObjectsShardsGetParams) WithDefaults() *SchemaObjectsShardsGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the schema objects shards get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SchemaObjectsShardsGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the schema objects shards get params
func (o *SchemaObjectsShardsGetParams) WithTimeout(timeout time.Duration) *SchemaObjectsShardsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the schema objects shards get params
func (o *SchemaObjectsShardsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the schema objects shards get params
func (o *SchemaObjectsShardsGetParams) WithContext(ctx context.Context) *SchemaObjectsShardsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the schema objects shards get params
func (o *SchemaObjectsShardsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the schema objects shards get params
func (o *SchemaObjectsShardsGetParams) WithHTTPClient(client *http.Client) *SchemaObjectsShardsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the schema objects shards get params
func (o *SchemaObjectsShardsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClassName adds the className to the schema objects shards get params
func (o *SchemaObjectsShardsGetParams) WithClassName(className string) *SchemaObjectsShardsGetParams {
	o.SetClassName(className)
	return o
}

// SetClassName adds the className to the schema objects shards get params
func (o *SchemaObjectsShardsGetParams) SetClassName(className string) {
	o.ClassName = className
}

// WithTenant adds the tenant to the schema objects shards get params
func (o *SchemaObjectsShardsGetParams) WithTenant(tenant *string) *SchemaObjectsShardsGetParams {
	o.SetTenant(tenant)
	return o
}

// SetTenant adds the tenant to the schema objects shards get params
func (o *SchemaObjectsShardsGetParams) SetTenant(tenant *string) {
	o.Tenant = tenant
}

// WriteToRequest writes these params to a swagger request
func (o *SchemaObjectsShardsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param className
	if err := r.SetPathParam("className", o.ClassName); err != nil {
		return err
	}

	if o.Tenant != nil {

		// query param tenant
		var qrTenant string

		if o.Tenant != nil {
			qrTenant = *o.Tenant
		}
		qTenant := qrTenant
		if qTenant != "" {

			if err := r.SetQueryParam("tenant", qTenant); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
