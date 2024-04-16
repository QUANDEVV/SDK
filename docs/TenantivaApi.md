# \TenantivaApi

All URIs are relative to *https://core.proximaai.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TenantivaCreateagentCreate**](TenantivaApi.md#TenantivaCreateagentCreate) | **Post** /api/tenantiva/createagent/ | 
[**TenantivaSampletoolsRetrieve**](TenantivaApi.md#TenantivaSampletoolsRetrieve) | **Get** /api/tenantiva/sampletools/ | 
[**TenantivaTenantknowlegebaseDestroy**](TenantivaApi.md#TenantivaTenantknowlegebaseDestroy) | **Delete** /api/tenantiva/tenantknowlegebase/ | 
[**TenantivaTenantknowlegebasePartialUpdate**](TenantivaApi.md#TenantivaTenantknowlegebasePartialUpdate) | **Patch** /api/tenantiva/tenantknowlegebase/ | 
[**TenantivaTenantknowlegebaseRetrieve**](TenantivaApi.md#TenantivaTenantknowlegebaseRetrieve) | **Get** /api/tenantiva/tenantknowlegebase/ | 
[**TenantivaTrainingprogressRetrieve**](TenantivaApi.md#TenantivaTrainingprogressRetrieve) | **Get** /api/tenantiva/trainingprogress/ | 



## TenantivaCreateagentCreate

> CreateAgent TenantivaCreateagentCreate(ctx).CreateAgent(createAgent).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    createAgent := *openapiclient.NewCreateAgent("TenantName_example", "ToolName_example") // CreateAgent | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantivaApi.TenantivaCreateagentCreate(context.Background()).CreateAgent(createAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantivaApi.TenantivaCreateagentCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantivaCreateagentCreate`: CreateAgent
    fmt.Fprintf(os.Stdout, "Response from `TenantivaApi.TenantivaCreateagentCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenantivaCreateagentCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAgent** | [**CreateAgent**](CreateAgent.md) |  | 

### Return type

[**CreateAgent**](CreateAgent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantivaSampletoolsRetrieve

> TenantivaSampletoolsRetrieve(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantivaApi.TenantivaSampletoolsRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantivaApi.TenantivaSampletoolsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTenantivaSampletoolsRetrieveRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantivaTenantknowlegebaseDestroy

> TenantivaTenantknowlegebaseDestroy(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantivaApi.TenantivaTenantknowlegebaseDestroy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantivaApi.TenantivaTenantknowlegebaseDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTenantivaTenantknowlegebaseDestroyRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantivaTenantknowlegebasePartialUpdate

> Knowledgebase TenantivaTenantknowlegebasePartialUpdate(ctx).PatchedKnowledgebase(patchedKnowledgebase).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    patchedKnowledgebase := *openapiclient.NewPatchedKnowledgebase() // PatchedKnowledgebase |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantivaApi.TenantivaTenantknowlegebasePartialUpdate(context.Background()).PatchedKnowledgebase(patchedKnowledgebase).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantivaApi.TenantivaTenantknowlegebasePartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantivaTenantknowlegebasePartialUpdate`: Knowledgebase
    fmt.Fprintf(os.Stdout, "Response from `TenantivaApi.TenantivaTenantknowlegebasePartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenantivaTenantknowlegebasePartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedKnowledgebase** | [**PatchedKnowledgebase**](PatchedKnowledgebase.md) |  | 

### Return type

[**Knowledgebase**](Knowledgebase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantivaTenantknowlegebaseRetrieve

> Knowledgebase TenantivaTenantknowlegebaseRetrieve(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantivaApi.TenantivaTenantknowlegebaseRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantivaApi.TenantivaTenantknowlegebaseRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantivaTenantknowlegebaseRetrieve`: Knowledgebase
    fmt.Fprintf(os.Stdout, "Response from `TenantivaApi.TenantivaTenantknowlegebaseRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTenantivaTenantknowlegebaseRetrieveRequest struct via the builder pattern


### Return type

[**Knowledgebase**](Knowledgebase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantivaTrainingprogressRetrieve

> AssistantTrainingProgressase TenantivaTrainingprogressRetrieve(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantivaApi.TenantivaTrainingprogressRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantivaApi.TenantivaTrainingprogressRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantivaTrainingprogressRetrieve`: AssistantTrainingProgressase
    fmt.Fprintf(os.Stdout, "Response from `TenantivaApi.TenantivaTrainingprogressRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTenantivaTrainingprogressRetrieveRequest struct via the builder pattern


### Return type

[**AssistantTrainingProgressase**](AssistantTrainingProgressase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

