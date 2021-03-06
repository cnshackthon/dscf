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
	"encoding/json"
	"errors"
	"net/http"

	//"net/http"
	//"errors"
	"fmt"
	"os"

	clientapi "github.com/cnshackthon/dscf/api"
	clientmodel "github.com/cnshackthon/dscf/model"
)

// LocationApiService is a service that implements the logic for the LocationApiServicer
// This service should implement the business logic for every endpoint for the LocationApi API.
// Include any external packages or services that will be required by this service.
type LocationApiService struct {
}

// NewLocationApiService creates a default api service
func NewLocationApiService() LocationApiServicer {
	return &LocationApiService{}
}

// GetSubIdLocation - Get last reported location
func (s *LocationApiService) GetSubIdLocation(ctx context.Context, subId string) (ImplResponse, error) {
	// TODO - update GetSubIdLocation with the required logic for this service method.
	// Add api_location_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, LocationResponse{}) or use other options such as http.Ok ...
	//return Response(200, LocationResponse{}), nil

	//TODO: Uncomment the next line to return response Response(0, {}) or use other options such as http.Ok ...
	//return Response(0, nil),nil
	// get location from api gee sim

	imsi, ok := globCorrMap[subId]
	if ok {
		fmt.Fprintf(os.Stdout, "sub id is : %s, convert to %s\n", subId, imsi)
		subscriberId := *clientmodel.NewSubscriberId(globCorrMap[subId]) // SubscriberId | External ID of the subscriber
		xTestmode := "true"                                              // string | Enables test mode when set to \"true\".  Assumes \"false\" if not present. (optional)

		configuration := clientapi.NewConfiguration()
		api_client := clientapi.NewAPIClient(configuration)
		resp, r, err := api_client.LocationApi.GetSubscriberLocation(context.Background()).SubscriberId(subscriberId).XTestmode(xTestmode).Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `LocationApi.GetSubscriberLocation``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}

		resp.Id = &subId
		// response from `GetSubscriberLocation`: LocationResponse
		a, _ := json.Marshal(resp)
		fmt.Fprintf(os.Stdout, "Response from `LocationApi.GetSubscriberLocation`: %s\n", a)

		//return Response(http.StatusNotImplemented, nil), errors.New("GetSubIdLocation method not implemented")
		return Response(200, resp), nil
	} else {

		return Response(http.StatusBadRequest, nil), errors.New("no such subId in the data base! sud id is : " +subId)
	}
}
