/*
proxima-core-engine

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

// PredictiveApiService PredictiveApi service
type PredictiveApiService service

type ApiPredictiveEscalatedchatsRetrieveRequest struct {
	ctx _context.Context
	ApiService *PredictiveApiService
}


func (r ApiPredictiveEscalatedchatsRetrieveRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.PredictiveEscalatedchatsRetrieveExecute(r)
}

/*
PredictiveEscalatedchatsRetrieve Method for PredictiveEscalatedchatsRetrieve

GET
Retrieve EscalatedChatsss matching query.

Params:
escalated_chats_id
tenant
date

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPredictiveEscalatedchatsRetrieveRequest
*/
func (a *PredictiveApiService) PredictiveEscalatedchatsRetrieve(ctx _context.Context) ApiPredictiveEscalatedchatsRetrieveRequest {
	return ApiPredictiveEscalatedchatsRetrieveRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *PredictiveApiService) PredictiveEscalatedchatsRetrieveExecute(r ApiPredictiveEscalatedchatsRetrieveRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PredictiveApiService.PredictiveEscalatedchatsRetrieve")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/predictive/escalatedchats/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiPredictiveGenderdistributionRetrieveRequest struct {
	ctx _context.Context
	ApiService *PredictiveApiService
}


func (r ApiPredictiveGenderdistributionRetrieveRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.PredictiveGenderdistributionRetrieveExecute(r)
}

/*
PredictiveGenderdistributionRetrieve Method for PredictiveGenderdistributionRetrieve

GET
Retrieve GenderDistributionss matching query.

Params:
gender_distribution_id
tenant
date

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPredictiveGenderdistributionRetrieveRequest
*/
func (a *PredictiveApiService) PredictiveGenderdistributionRetrieve(ctx _context.Context) ApiPredictiveGenderdistributionRetrieveRequest {
	return ApiPredictiveGenderdistributionRetrieveRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *PredictiveApiService) PredictiveGenderdistributionRetrieveExecute(r ApiPredictiveGenderdistributionRetrieveRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PredictiveApiService.PredictiveGenderdistributionRetrieve")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/predictive/genderdistribution/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiPredictiveGrowthrateRetrieveRequest struct {
	ctx _context.Context
	ApiService *PredictiveApiService
}


func (r ApiPredictiveGrowthrateRetrieveRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.PredictiveGrowthrateRetrieveExecute(r)
}

/*
PredictiveGrowthrateRetrieve Method for PredictiveGrowthrateRetrieve

GET
Retrieve GrowthRatess matching query.

Params:
growth_rate_id
tenant
date

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPredictiveGrowthrateRetrieveRequest
*/
func (a *PredictiveApiService) PredictiveGrowthrateRetrieve(ctx _context.Context) ApiPredictiveGrowthrateRetrieveRequest {
	return ApiPredictiveGrowthrateRetrieveRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *PredictiveApiService) PredictiveGrowthrateRetrieveExecute(r ApiPredictiveGrowthrateRetrieveRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PredictiveApiService.PredictiveGrowthrateRetrieve")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/predictive/growthrate/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiPredictiveInteractionsRetrieveRequest struct {
	ctx _context.Context
	ApiService *PredictiveApiService
}


func (r ApiPredictiveInteractionsRetrieveRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.PredictiveInteractionsRetrieveExecute(r)
}

/*
PredictiveInteractionsRetrieve Method for PredictiveInteractionsRetrieve

GET
Retrieve interactionss matching query.

Params:
interactions_id
tenant
date

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPredictiveInteractionsRetrieveRequest
*/
func (a *PredictiveApiService) PredictiveInteractionsRetrieve(ctx _context.Context) ApiPredictiveInteractionsRetrieveRequest {
	return ApiPredictiveInteractionsRetrieveRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *PredictiveApiService) PredictiveInteractionsRetrieveExecute(r ApiPredictiveInteractionsRetrieveRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PredictiveApiService.PredictiveInteractionsRetrieve")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/predictive/interactions/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiPredictiveInvolvementsRetrieveRequest struct {
	ctx _context.Context
	ApiService *PredictiveApiService
}


func (r ApiPredictiveInvolvementsRetrieveRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.PredictiveInvolvementsRetrieveExecute(r)
}

/*
PredictiveInvolvementsRetrieve Method for PredictiveInvolvementsRetrieve

GET
Retrieve Involvementss matching query.

Params:
involvements_id
tenant
date

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPredictiveInvolvementsRetrieveRequest
*/
func (a *PredictiveApiService) PredictiveInvolvementsRetrieve(ctx _context.Context) ApiPredictiveInvolvementsRetrieveRequest {
	return ApiPredictiveInvolvementsRetrieveRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *PredictiveApiService) PredictiveInvolvementsRetrieveExecute(r ApiPredictiveInvolvementsRetrieveRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PredictiveApiService.PredictiveInvolvementsRetrieve")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/predictive/involvements/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiPredictiveTopicsRetrieveRequest struct {
	ctx _context.Context
	ApiService *PredictiveApiService
}


func (r ApiPredictiveTopicsRetrieveRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.PredictiveTopicsRetrieveExecute(r)
}

/*
PredictiveTopicsRetrieve Method for PredictiveTopicsRetrieve

GET
Retrieve topics matching query.

Params:
topics_id
tenant
date

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPredictiveTopicsRetrieveRequest
*/
func (a *PredictiveApiService) PredictiveTopicsRetrieve(ctx _context.Context) ApiPredictiveTopicsRetrieveRequest {
	return ApiPredictiveTopicsRetrieveRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *PredictiveApiService) PredictiveTopicsRetrieveExecute(r ApiPredictiveTopicsRetrieveRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PredictiveApiService.PredictiveTopicsRetrieve")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/predictive/topics/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiPredictiveUniqueclientsRetrieveRequest struct {
	ctx _context.Context
	ApiService *PredictiveApiService
}


func (r ApiPredictiveUniqueclientsRetrieveRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.PredictiveUniqueclientsRetrieveExecute(r)
}

/*
PredictiveUniqueclientsRetrieve Method for PredictiveUniqueclientsRetrieve

GET
Retrieve UniqueClientsss matching query.

Params:
unique_clients_id
tenant
date

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPredictiveUniqueclientsRetrieveRequest
*/
func (a *PredictiveApiService) PredictiveUniqueclientsRetrieve(ctx _context.Context) ApiPredictiveUniqueclientsRetrieveRequest {
	return ApiPredictiveUniqueclientsRetrieveRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *PredictiveApiService) PredictiveUniqueclientsRetrieveExecute(r ApiPredictiveUniqueclientsRetrieveRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PredictiveApiService.PredictiveUniqueclientsRetrieve")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/predictive/uniqueclients/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
