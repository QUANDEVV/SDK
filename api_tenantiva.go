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

// TenantivaApiService TenantivaApi service
type TenantivaApiService service

type ApiTenantivaCreateagentCreateRequest struct {
	ctx _context.Context
	ApiService *TenantivaApiService
	createAgent *CreateAgent
}

func (r ApiTenantivaCreateagentCreateRequest) CreateAgent(createAgent CreateAgent) ApiTenantivaCreateagentCreateRequest {
	r.createAgent = &createAgent
	return r
}

func (r ApiTenantivaCreateagentCreateRequest) Execute() (CreateAgent, *_nethttp.Response, error) {
	return r.ApiService.TenantivaCreateagentCreateExecute(r)
}

/*
TenantivaCreateagentCreate Method for TenantivaCreateagentCreate

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiTenantivaCreateagentCreateRequest
*/
func (a *TenantivaApiService) TenantivaCreateagentCreate(ctx _context.Context) ApiTenantivaCreateagentCreateRequest {
	return ApiTenantivaCreateagentCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreateAgent
func (a *TenantivaApiService) TenantivaCreateagentCreateExecute(r ApiTenantivaCreateagentCreateRequest) (CreateAgent, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CreateAgent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TenantivaApiService.TenantivaCreateagentCreate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/tenantiva/createagent/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.createAgent == nil {
		return localVarReturnValue, nil, reportError("createAgent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/x-www-form-urlencoded", "multipart/form-data"}

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
	// body params
	localVarPostBody = r.createAgent
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
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

type ApiTenantivaSampletoolsRetrieveRequest struct {
	ctx _context.Context
	ApiService *TenantivaApiService
}


func (r ApiTenantivaSampletoolsRetrieveRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.TenantivaSampletoolsRetrieveExecute(r)
}

/*
TenantivaSampletoolsRetrieve Method for TenantivaSampletoolsRetrieve

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiTenantivaSampletoolsRetrieveRequest
*/
func (a *TenantivaApiService) TenantivaSampletoolsRetrieve(ctx _context.Context) ApiTenantivaSampletoolsRetrieveRequest {
	return ApiTenantivaSampletoolsRetrieveRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *TenantivaApiService) TenantivaSampletoolsRetrieveExecute(r ApiTenantivaSampletoolsRetrieveRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TenantivaApiService.TenantivaSampletoolsRetrieve")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/tenantiva/sampletools/"

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

type ApiTenantivaTenantknowlegebaseDestroyRequest struct {
	ctx _context.Context
	ApiService *TenantivaApiService
}


func (r ApiTenantivaTenantknowlegebaseDestroyRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.TenantivaTenantknowlegebaseDestroyExecute(r)
}

/*
TenantivaTenantknowlegebaseDestroy Method for TenantivaTenantknowlegebaseDestroy

Mixin for applying DRF's pagination.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiTenantivaTenantknowlegebaseDestroyRequest
*/
func (a *TenantivaApiService) TenantivaTenantknowlegebaseDestroy(ctx _context.Context) ApiTenantivaTenantknowlegebaseDestroyRequest {
	return ApiTenantivaTenantknowlegebaseDestroyRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *TenantivaApiService) TenantivaTenantknowlegebaseDestroyExecute(r ApiTenantivaTenantknowlegebaseDestroyRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TenantivaApiService.TenantivaTenantknowlegebaseDestroy")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/tenantiva/tenantknowlegebase/"

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

type ApiTenantivaTenantknowlegebasePartialUpdateRequest struct {
	ctx _context.Context
	ApiService *TenantivaApiService
	patchedKnowledgebase *PatchedKnowledgebase
}

func (r ApiTenantivaTenantknowlegebasePartialUpdateRequest) PatchedKnowledgebase(patchedKnowledgebase PatchedKnowledgebase) ApiTenantivaTenantknowlegebasePartialUpdateRequest {
	r.patchedKnowledgebase = &patchedKnowledgebase
	return r
}

func (r ApiTenantivaTenantknowlegebasePartialUpdateRequest) Execute() (Knowledgebase, *_nethttp.Response, error) {
	return r.ApiService.TenantivaTenantknowlegebasePartialUpdateExecute(r)
}

/*
TenantivaTenantknowlegebasePartialUpdate Method for TenantivaTenantknowlegebasePartialUpdate

Mixin for applying DRF's pagination.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiTenantivaTenantknowlegebasePartialUpdateRequest
*/
func (a *TenantivaApiService) TenantivaTenantknowlegebasePartialUpdate(ctx _context.Context) ApiTenantivaTenantknowlegebasePartialUpdateRequest {
	return ApiTenantivaTenantknowlegebasePartialUpdateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Knowledgebase
func (a *TenantivaApiService) TenantivaTenantknowlegebasePartialUpdateExecute(r ApiTenantivaTenantknowlegebasePartialUpdateRequest) (Knowledgebase, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Knowledgebase
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TenantivaApiService.TenantivaTenantknowlegebasePartialUpdate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/tenantiva/tenantknowlegebase/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/x-www-form-urlencoded", "multipart/form-data"}

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
	// body params
	localVarPostBody = r.patchedKnowledgebase
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
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

type ApiTenantivaTenantknowlegebaseRetrieveRequest struct {
	ctx _context.Context
	ApiService *TenantivaApiService
}


func (r ApiTenantivaTenantknowlegebaseRetrieveRequest) Execute() (Knowledgebase, *_nethttp.Response, error) {
	return r.ApiService.TenantivaTenantknowlegebaseRetrieveExecute(r)
}

/*
TenantivaTenantknowlegebaseRetrieve Method for TenantivaTenantknowlegebaseRetrieve

Mixin for applying DRF's pagination.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiTenantivaTenantknowlegebaseRetrieveRequest
*/
func (a *TenantivaApiService) TenantivaTenantknowlegebaseRetrieve(ctx _context.Context) ApiTenantivaTenantknowlegebaseRetrieveRequest {
	return ApiTenantivaTenantknowlegebaseRetrieveRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Knowledgebase
func (a *TenantivaApiService) TenantivaTenantknowlegebaseRetrieveExecute(r ApiTenantivaTenantknowlegebaseRetrieveRequest) (Knowledgebase, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Knowledgebase
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TenantivaApiService.TenantivaTenantknowlegebaseRetrieve")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/tenantiva/tenantknowlegebase/"

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
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

type ApiTenantivaTrainingprogressRetrieveRequest struct {
	ctx _context.Context
	ApiService *TenantivaApiService
}


func (r ApiTenantivaTrainingprogressRetrieveRequest) Execute() (AssistantTrainingProgressase, *_nethttp.Response, error) {
	return r.ApiService.TenantivaTrainingprogressRetrieveExecute(r)
}

/*
TenantivaTrainingprogressRetrieve Method for TenantivaTrainingprogressRetrieve

GET
Params:
assistant_training_id
tenant_id

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiTenantivaTrainingprogressRetrieveRequest
*/
func (a *TenantivaApiService) TenantivaTrainingprogressRetrieve(ctx _context.Context) ApiTenantivaTrainingprogressRetrieveRequest {
	return ApiTenantivaTrainingprogressRetrieveRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AssistantTrainingProgressase
func (a *TenantivaApiService) TenantivaTrainingprogressRetrieveExecute(r ApiTenantivaTrainingprogressRetrieveRequest) (AssistantTrainingProgressase, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AssistantTrainingProgressase
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TenantivaApiService.TenantivaTrainingprogressRetrieve")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/tenantiva/trainingprogress/"

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
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
