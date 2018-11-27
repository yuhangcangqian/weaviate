/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2018 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * AUTHOR: Bob van Luijt (bob@kub.design)
 * See www.creativesoftwarefdn.org for details
 * Contact: @CreativeSofwFdn / bob@kub.design
 */
// Code generated by go-swagger; DO NOT EDIT.

package things

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// NewWeaviateThingsPatchParams creates a new WeaviateThingsPatchParams object
// no default values defined in spec.
func NewWeaviateThingsPatchParams() WeaviateThingsPatchParams {

	return WeaviateThingsPatchParams{}
}

// WeaviateThingsPatchParams contains all the bound params for the weaviate things patch operation
// typically these are obtained from a http.Request
//
// swagger:parameters weaviate.things.patch
type WeaviateThingsPatchParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*If `async` is true, return a 202 if the patch is accepted. You will receive this response before the data is persisted. If `async` is false, you will receive confirmation after the update is persisted. The value of `async` defaults to false.
	  In: query
	*/
	Async *bool
	/*JSONPatch document as defined by RFC 6902.
	  Required: true
	  In: body
	*/
	Body []*models.PatchDocument
	/*Unique ID of the thing.
	  Required: true
	  In: path
	*/
	ThingID strfmt.UUID
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewWeaviateThingsPatchParams() beforehand.
func (o *WeaviateThingsPatchParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qAsync, qhkAsync, _ := qs.GetOK("async")
	if err := o.bindAsync(qAsync, qhkAsync, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body []*models.PatchDocument
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("body", "body"))
			} else {
				res = append(res, errors.NewParseError("body", "body", "", err))
			}
		} else {
			// validate array of body objects
			for i := range body {
				if body[i] == nil {
					continue
				}
				if err := body[i].Validate(route.Formats); err != nil {
					res = append(res, err)
					break
				}
			}

			if len(res) == 0 {
				o.Body = body
			}
		}
	} else {
		res = append(res, errors.Required("body", "body"))
	}
	rThingID, rhkThingID, _ := route.Params.GetOK("thingId")
	if err := o.bindThingID(rThingID, rhkThingID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAsync binds and validates parameter Async from query.
func (o *WeaviateThingsPatchParams) bindAsync(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("async", "query", "bool", raw)
	}
	o.Async = &value

	return nil
}

// bindThingID binds and validates parameter ThingID from path.
func (o *WeaviateThingsPatchParams) bindThingID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("thingId", "path", "strfmt.UUID", raw)
	}
	o.ThingID = *(value.(*strfmt.UUID))

	if err := o.validateThingID(formats); err != nil {
		return err
	}

	return nil
}

// validateThingID carries on validations for parameter ThingID
func (o *WeaviateThingsPatchParams) validateThingID(formats strfmt.Registry) error {

	if err := validate.FormatOf("thingId", "path", "uuid", o.ThingID.String(), formats); err != nil {
		return err
	}
	return nil
}
