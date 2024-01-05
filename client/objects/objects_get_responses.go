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

package objects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/weaviate/weaviate/entities/models"
)

// ObjectsGetReader is a Reader for the ObjectsGet structure.
type ObjectsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ObjectsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewObjectsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewObjectsGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewObjectsGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewObjectsGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewObjectsGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewObjectsGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewObjectsGetOK creates a ObjectsGetOK with default headers values
func NewObjectsGetOK() *ObjectsGetOK {
	return &ObjectsGetOK{}
}

/*
ObjectsGetOK describes a response with status code 200, with default header values.

Successful response.
*/
type ObjectsGetOK struct {
	Payload *models.Object
}

// IsSuccess returns true when this objects get o k response has a 2xx status code
func (o *ObjectsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this objects get o k response has a 3xx status code
func (o *ObjectsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this objects get o k response has a 4xx status code
func (o *ObjectsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this objects get o k response has a 5xx status code
func (o *ObjectsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this objects get o k response a status code equal to that given
func (o *ObjectsGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the objects get o k response
func (o *ObjectsGetOK) Code() int {
	return 200
}

func (o *ObjectsGetOK) Error() string {
	return fmt.Sprintf("[GET /objects/{id}][%d] objectsGetOK  %+v", 200, o.Payload)
}

func (o *ObjectsGetOK) String() string {
	return fmt.Sprintf("[GET /objects/{id}][%d] objectsGetOK  %+v", 200, o.Payload)
}

func (o *ObjectsGetOK) GetPayload() *models.Object {
	return o.Payload
}

func (o *ObjectsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Object)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewObjectsGetBadRequest creates a ObjectsGetBadRequest with default headers values
func NewObjectsGetBadRequest() *ObjectsGetBadRequest {
	return &ObjectsGetBadRequest{}
}

/*
ObjectsGetBadRequest describes a response with status code 400, with default header values.

Malformed request.
*/
type ObjectsGetBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this objects get bad request response has a 2xx status code
func (o *ObjectsGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this objects get bad request response has a 3xx status code
func (o *ObjectsGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this objects get bad request response has a 4xx status code
func (o *ObjectsGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this objects get bad request response has a 5xx status code
func (o *ObjectsGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this objects get bad request response a status code equal to that given
func (o *ObjectsGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the objects get bad request response
func (o *ObjectsGetBadRequest) Code() int {
	return 400
}

func (o *ObjectsGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /objects/{id}][%d] objectsGetBadRequest  %+v", 400, o.Payload)
}

func (o *ObjectsGetBadRequest) String() string {
	return fmt.Sprintf("[GET /objects/{id}][%d] objectsGetBadRequest  %+v", 400, o.Payload)
}

func (o *ObjectsGetBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ObjectsGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewObjectsGetUnauthorized creates a ObjectsGetUnauthorized with default headers values
func NewObjectsGetUnauthorized() *ObjectsGetUnauthorized {
	return &ObjectsGetUnauthorized{}
}

/*
ObjectsGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized or invalid credentials.
*/
type ObjectsGetUnauthorized struct {
}

// IsSuccess returns true when this objects get unauthorized response has a 2xx status code
func (o *ObjectsGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this objects get unauthorized response has a 3xx status code
func (o *ObjectsGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this objects get unauthorized response has a 4xx status code
func (o *ObjectsGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this objects get unauthorized response has a 5xx status code
func (o *ObjectsGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this objects get unauthorized response a status code equal to that given
func (o *ObjectsGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the objects get unauthorized response
func (o *ObjectsGetUnauthorized) Code() int {
	return 401
}

func (o *ObjectsGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /objects/{id}][%d] objectsGetUnauthorized ", 401)
}

func (o *ObjectsGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /objects/{id}][%d] objectsGetUnauthorized ", 401)
}

func (o *ObjectsGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewObjectsGetForbidden creates a ObjectsGetForbidden with default headers values
func NewObjectsGetForbidden() *ObjectsGetForbidden {
	return &ObjectsGetForbidden{}
}

/*
ObjectsGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ObjectsGetForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this objects get forbidden response has a 2xx status code
func (o *ObjectsGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this objects get forbidden response has a 3xx status code
func (o *ObjectsGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this objects get forbidden response has a 4xx status code
func (o *ObjectsGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this objects get forbidden response has a 5xx status code
func (o *ObjectsGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this objects get forbidden response a status code equal to that given
func (o *ObjectsGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the objects get forbidden response
func (o *ObjectsGetForbidden) Code() int {
	return 403
}

func (o *ObjectsGetForbidden) Error() string {
	return fmt.Sprintf("[GET /objects/{id}][%d] objectsGetForbidden  %+v", 403, o.Payload)
}

func (o *ObjectsGetForbidden) String() string {
	return fmt.Sprintf("[GET /objects/{id}][%d] objectsGetForbidden  %+v", 403, o.Payload)
}

func (o *ObjectsGetForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ObjectsGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewObjectsGetNotFound creates a ObjectsGetNotFound with default headers values
func NewObjectsGetNotFound() *ObjectsGetNotFound {
	return &ObjectsGetNotFound{}
}

/*
ObjectsGetNotFound describes a response with status code 404, with default header values.

Successful query result but no resource was found.
*/
type ObjectsGetNotFound struct {
}

// IsSuccess returns true when this objects get not found response has a 2xx status code
func (o *ObjectsGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this objects get not found response has a 3xx status code
func (o *ObjectsGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this objects get not found response has a 4xx status code
func (o *ObjectsGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this objects get not found response has a 5xx status code
func (o *ObjectsGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this objects get not found response a status code equal to that given
func (o *ObjectsGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the objects get not found response
func (o *ObjectsGetNotFound) Code() int {
	return 404
}

func (o *ObjectsGetNotFound) Error() string {
	return fmt.Sprintf("[GET /objects/{id}][%d] objectsGetNotFound ", 404)
}

func (o *ObjectsGetNotFound) String() string {
	return fmt.Sprintf("[GET /objects/{id}][%d] objectsGetNotFound ", 404)
}

func (o *ObjectsGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewObjectsGetInternalServerError creates a ObjectsGetInternalServerError with default headers values
func NewObjectsGetInternalServerError() *ObjectsGetInternalServerError {
	return &ObjectsGetInternalServerError{}
}

/*
ObjectsGetInternalServerError describes a response with status code 500, with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type ObjectsGetInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this objects get internal server error response has a 2xx status code
func (o *ObjectsGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this objects get internal server error response has a 3xx status code
func (o *ObjectsGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this objects get internal server error response has a 4xx status code
func (o *ObjectsGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this objects get internal server error response has a 5xx status code
func (o *ObjectsGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this objects get internal server error response a status code equal to that given
func (o *ObjectsGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the objects get internal server error response
func (o *ObjectsGetInternalServerError) Code() int {
	return 500
}

func (o *ObjectsGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /objects/{id}][%d] objectsGetInternalServerError  %+v", 500, o.Payload)
}

func (o *ObjectsGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /objects/{id}][%d] objectsGetInternalServerError  %+v", 500, o.Payload)
}

func (o *ObjectsGetInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ObjectsGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
