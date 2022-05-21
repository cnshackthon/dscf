/*
Network as Code

Manipulate network conditions via simplified REST calls

API version: 2
Contact: todd.levi@nokia.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datarepository

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	clientmodel "github.com/cnshackthon/dscf/model"
)

// Linger please
var (
	_ _context.Context
)

// BandwidthApiService BandwidthApi service
type BandwidthApiService service

type ApiGetCustomBandwidthRequest struct {
	ctx _context.Context
	ApiService *BandwidthApiService
	subscriberId *clientmodel.SubscriberId
	xTestmode *string
}

// External ID of the subscriber
func (r ApiGetCustomBandwidthRequest) SubscriberId(subscriberId clientmodel.SubscriberId) ApiGetCustomBandwidthRequest {
	r.subscriberId = &subscriberId
	return r
}
// Enables test mode when set to \&quot;true\&quot;.  Assumes \&quot;false\&quot; if not present.
func (r ApiGetCustomBandwidthRequest) XTestmode(xTestmode string) ApiGetCustomBandwidthRequest {
	r.xTestmode = &xTestmode
	return r
}

func (r ApiGetCustomBandwidthRequest) Execute() (clientmodel.CustomLimits, *_nethttp.Response, error) {
	return r.ApiService.GetCustomBandwidthExecute(r)
}

/*
GetCustomBandwidth Get upload/download limit

Get the current upload/download bandwidth limits for the subscriber identified by the given ID

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetCustomBandwidthRequest
*/
func (a *BandwidthApiService) GetCustomBandwidth(ctx _context.Context) ApiGetCustomBandwidthRequest {
	return ApiGetCustomBandwidthRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CustomLimits
func (a *BandwidthApiService) GetCustomBandwidthExecute(r ApiGetCustomBandwidthRequest) (clientmodel.CustomLimits, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  clientmodel.CustomLimits
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BandwidthApiService.GetCustomBandwidth")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscriber/bandwidth/custom"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.subscriberId == nil {
		return localVarReturnValue, nil, reportError("subscriberId is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xTestmode != nil {
		localVarHeaderParams["x-testmode"] = parameterToString(*r.xTestmode, "")
	}
	// body params
	localVarPostBody = r.subscriberId
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apigee_apikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-apikey"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v clientmodel.ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetSubscriberBandwidthRequest struct {
	ctx _context.Context
	ApiService *BandwidthApiService
	subscriberId *clientmodel.SubscriberId
	xTestmode *string
}

// External ID of the subscriber
func (r ApiGetSubscriberBandwidthRequest) SubscriberId(subscriberId clientmodel.SubscriberId) ApiGetSubscriberBandwidthRequest {
	r.subscriberId = &subscriberId
	return r
}
// Enables test mode when set to \&quot;true\&quot;.  Assumes \&quot;false\&quot; if not present.
func (r ApiGetSubscriberBandwidthRequest) XTestmode(xTestmode string) ApiGetSubscriberBandwidthRequest {
	r.xTestmode = &xTestmode
	return r
}

func (r ApiGetSubscriberBandwidthRequest) Execute() (clientmodel.BandwidthResponse, *_nethttp.Response, error) {
	return r.ApiService.GetSubscriberBandwidthExecute(r)
}

/*
GetSubscriberBandwidth Get the current subscriber bandwidth

Get the bandwidth identifier for the subscriber identified by the given IMSI

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSubscriberBandwidthRequest
*/
func (a *BandwidthApiService) GetSubscriberBandwidth(ctx _context.Context) ApiGetSubscriberBandwidthRequest {
	return ApiGetSubscriberBandwidthRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BandwidthResponse
func (a *BandwidthApiService) GetSubscriberBandwidthExecute(r ApiGetSubscriberBandwidthRequest) (clientmodel.BandwidthResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  clientmodel.BandwidthResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BandwidthApiService.GetSubscriberBandwidth")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscriber/bandwidth"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.subscriberId == nil {
		return localVarReturnValue, nil, reportError("subscriberId is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xTestmode != nil {
		localVarHeaderParams["x-testmode"] = parameterToString(*r.xTestmode, "")
	}
	// body params
	localVarPostBody = r.subscriberId
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apigee_apikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-apikey"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v clientmodel.ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateCustomBandwidthRequest struct {
	ctx _context.Context
	ApiService *BandwidthApiService
	customLimits *clientmodel.CustomLimits
	xTestmode *string
}

// New upload/download limits for the subscriber
func (r ApiUpdateCustomBandwidthRequest) CustomLimits(customLimits clientmodel.CustomLimits) ApiUpdateCustomBandwidthRequest {
	r.customLimits = &customLimits
	return r
}
// Enables test mode when set to \&quot;true\&quot;.  Assumes \&quot;false\&quot; if not present.
func (r ApiUpdateCustomBandwidthRequest) XTestmode(xTestmode string) ApiUpdateCustomBandwidthRequest {
	r.xTestmode = &xTestmode
	return r
}

func (r ApiUpdateCustomBandwidthRequest) Execute() (clientmodel.CustomLimits, *_nethttp.Response, error) {
	return r.ApiService.UpdateCustomBandwidthExecute(r)
}

/*
UpdateCustomBandwidth Set upload limit

Set the upload/download bandwidth limits for the subscriber identified by the given ID

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUpdateCustomBandwidthRequest
*/
func (a *BandwidthApiService) UpdateCustomBandwidth(ctx _context.Context) ApiUpdateCustomBandwidthRequest {
	return ApiUpdateCustomBandwidthRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CustomLimits
func (a *BandwidthApiService) UpdateCustomBandwidthExecute(r ApiUpdateCustomBandwidthRequest) (clientmodel.CustomLimits, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  clientmodel.CustomLimits
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BandwidthApiService.UpdateCustomBandwidth")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscriber/bandwidth/custom"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.customLimits == nil {
		return localVarReturnValue, nil, reportError("customLimits is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xTestmode != nil {
		localVarHeaderParams["x-testmode"] = parameterToString(*r.xTestmode, "")
	}
	// body params
	localVarPostBody = r.customLimits
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apigee_apikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-apikey"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v clientmodel.ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateSubscriberBandwidthRequest struct {
	ctx _context.Context
	ApiService *BandwidthApiService
	bandwidthUpdate *clientmodel.BandwidthUpdate
	xTestmode *string
}

// New bandwidth for subscriber
func (r ApiUpdateSubscriberBandwidthRequest) BandwidthUpdate(bandwidthUpdate clientmodel.BandwidthUpdate) ApiUpdateSubscriberBandwidthRequest {
	r.bandwidthUpdate = &bandwidthUpdate
	return r
}
// Enables test mode when set to \&quot;true\&quot;.  Assumes \&quot;false\&quot; if not present.
func (r ApiUpdateSubscriberBandwidthRequest) XTestmode(xTestmode string) ApiUpdateSubscriberBandwidthRequest {
	r.xTestmode = &xTestmode
	return r
}

func (r ApiUpdateSubscriberBandwidthRequest) Execute() (clientmodel.BandwidthResponse, *_nethttp.Response, error) {
	return r.ApiService.UpdateSubscriberBandwidthExecute(r)
}

/*
UpdateSubscriberBandwidth Update the bandwidth of the subscriber

Update the bandwidth identifier for the subscriber identified by the given IMSI

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUpdateSubscriberBandwidthRequest
*/
func (a *BandwidthApiService) UpdateSubscriberBandwidth(ctx _context.Context) ApiUpdateSubscriberBandwidthRequest {
	return ApiUpdateSubscriberBandwidthRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BandwidthResponse
func (a *BandwidthApiService) UpdateSubscriberBandwidthExecute(r ApiUpdateSubscriberBandwidthRequest) (clientmodel.BandwidthResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  clientmodel.BandwidthResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BandwidthApiService.UpdateSubscriberBandwidth")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscriber/bandwidth"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.bandwidthUpdate == nil {
		return localVarReturnValue, nil, reportError("bandwidthUpdate is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xTestmode != nil {
		localVarHeaderParams["x-testmode"] = parameterToString(*r.xTestmode, "")
	}
	// body params
	localVarPostBody = r.bandwidthUpdate
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apigee_apikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-apikey"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v clientmodel.ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v clientmodel.ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
