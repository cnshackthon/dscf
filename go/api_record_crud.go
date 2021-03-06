/*
 * Ndscf_DataRepository
 *
 * Ndscf Data Repository Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package datarepository

import (
	//"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// RecordCRUDApiController binds http requests to an api service and writes the service results to the http response
type RecordCRUDApiController struct {
	service RecordCRUDApiServicer
	errorHandler ErrorHandler
}

// RecordCRUDApiOption for how the controller is set up.
type RecordCRUDApiOption func(*RecordCRUDApiController)

// WithRecordCRUDApiErrorHandler inject ErrorHandler into controller
func WithRecordCRUDApiErrorHandler(h ErrorHandler) RecordCRUDApiOption {
	return func(c *RecordCRUDApiController) {
		c.errorHandler = h
	}
}

// NewRecordCRUDApiController creates a default api controller
func NewRecordCRUDApiController(s RecordCRUDApiServicer, opts ...RecordCRUDApiOption) Router {
	controller := &RecordCRUDApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all of the api route for the RecordCRUDApiController
func (c *RecordCRUDApiController) Routes() Routes {
	return Routes{ 
		{
			"DeleteCorrelation",
			strings.ToUpper("Delete"),
			"/ndscf/v1/correlation/{ueId}",
			c.DeleteCorrelation,
		},
	}
}

// DeleteCorrelation - Delete a Record with an user provided RecordId
func (c *RecordCRUDApiController) DeleteCorrelation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	ueIdParam := params["ueId"]
	
	subIdsParam := strings.Split(query.Get("subIds"), ",")
	result, err := c.service.DeleteCorrelation(r.Context(), ueIdParam, subIdsParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
