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
	"context"
	"net/http"
	"errors"
)

// SubscriberApiService is a service that implements the logic for the SubscriberApiServicer
// This service should implement the business logic for every endpoint for the SubscriberApi API.
// Include any external packages or services that will be required by this service.
type SubscriberApiService struct {
}

// NewSubscriberApiService creates a default api service
func NewSubscriberApiService() SubscriberApiServicer {
	return &SubscriberApiService{}
}

// GetSubIdLocation - Get last reported location
func (s *SubscriberApiService) GetSubIdLocation(ctx context.Context, subId string) (ImplResponse, error) {
	// TODO - update GetSubIdLocation with the required logic for this service method.
	// Add api_subscriber_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, LocationResponse{}) or use other options such as http.Ok ...
	//return Response(200, LocationResponse{}), nil

	//TODO: Uncomment the next line to return response Response(0, {}) or use other options such as http.Ok ...
	//return Response(0, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetSubIdLocation method not implemented")
}