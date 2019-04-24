/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@creativesoftwarefdn.org
 */ // Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/creativesoftwarefdn/weaviate/entities/models"
)

// WeaviateActionsDeleteNoContentCode is the HTTP code returned for type WeaviateActionsDeleteNoContent
const WeaviateActionsDeleteNoContentCode int = 204

/*WeaviateActionsDeleteNoContent Successfully deleted.

swagger:response weaviateActionsDeleteNoContent
*/
type WeaviateActionsDeleteNoContent struct {
}

// NewWeaviateActionsDeleteNoContent creates WeaviateActionsDeleteNoContent with default headers values
func NewWeaviateActionsDeleteNoContent() *WeaviateActionsDeleteNoContent {

	return &WeaviateActionsDeleteNoContent{}
}

// WriteResponse to the client
func (o *WeaviateActionsDeleteNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// WeaviateActionsDeleteUnauthorizedCode is the HTTP code returned for type WeaviateActionsDeleteUnauthorized
const WeaviateActionsDeleteUnauthorizedCode int = 401

/*WeaviateActionsDeleteUnauthorized Unauthorized or invalid credentials.

swagger:response weaviateActionsDeleteUnauthorized
*/
type WeaviateActionsDeleteUnauthorized struct {
}

// NewWeaviateActionsDeleteUnauthorized creates WeaviateActionsDeleteUnauthorized with default headers values
func NewWeaviateActionsDeleteUnauthorized() *WeaviateActionsDeleteUnauthorized {

	return &WeaviateActionsDeleteUnauthorized{}
}

// WriteResponse to the client
func (o *WeaviateActionsDeleteUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// WeaviateActionsDeleteForbiddenCode is the HTTP code returned for type WeaviateActionsDeleteForbidden
const WeaviateActionsDeleteForbiddenCode int = 403

/*WeaviateActionsDeleteForbidden Insufficient permissions.

swagger:response weaviateActionsDeleteForbidden
*/
type WeaviateActionsDeleteForbidden struct {
}

// NewWeaviateActionsDeleteForbidden creates WeaviateActionsDeleteForbidden with default headers values
func NewWeaviateActionsDeleteForbidden() *WeaviateActionsDeleteForbidden {

	return &WeaviateActionsDeleteForbidden{}
}

// WriteResponse to the client
func (o *WeaviateActionsDeleteForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// WeaviateActionsDeleteNotFoundCode is the HTTP code returned for type WeaviateActionsDeleteNotFound
const WeaviateActionsDeleteNotFoundCode int = 404

/*WeaviateActionsDeleteNotFound Successful query result but no resource was found.

swagger:response weaviateActionsDeleteNotFound
*/
type WeaviateActionsDeleteNotFound struct {
}

// NewWeaviateActionsDeleteNotFound creates WeaviateActionsDeleteNotFound with default headers values
func NewWeaviateActionsDeleteNotFound() *WeaviateActionsDeleteNotFound {

	return &WeaviateActionsDeleteNotFound{}
}

// WriteResponse to the client
func (o *WeaviateActionsDeleteNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// WeaviateActionsDeleteInternalServerErrorCode is the HTTP code returned for type WeaviateActionsDeleteInternalServerError
const WeaviateActionsDeleteInternalServerErrorCode int = 500

/*WeaviateActionsDeleteInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response weaviateActionsDeleteInternalServerError
*/
type WeaviateActionsDeleteInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateActionsDeleteInternalServerError creates WeaviateActionsDeleteInternalServerError with default headers values
func NewWeaviateActionsDeleteInternalServerError() *WeaviateActionsDeleteInternalServerError {

	return &WeaviateActionsDeleteInternalServerError{}
}

// WithPayload adds the payload to the weaviate actions delete internal server error response
func (o *WeaviateActionsDeleteInternalServerError) WithPayload(payload *models.ErrorResponse) *WeaviateActionsDeleteInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate actions delete internal server error response
func (o *WeaviateActionsDeleteInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateActionsDeleteInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}