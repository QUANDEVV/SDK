# \CallApi

All URIs are relative to *https://core.proximaai.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CallAnalyticsCallResolutionAnalyticsRetrieve**](CallApi.md#CallAnalyticsCallResolutionAnalyticsRetrieve) | **Get** /api/call/analytics/call-resolution-analytics/ | 
[**CallAnalyticsCallTagAnalyticsRetrieve**](CallApi.md#CallAnalyticsCallTagAnalyticsRetrieve) | **Get** /api/call/analytics/call-tag-analytics/ | 
[**CallAnalyticsCallVolumeAnalyticsRetrieve**](CallApi.md#CallAnalyticsCallVolumeAnalyticsRetrieve) | **Get** /api/call/analytics/call-volume-analytics/ | 
[**CallAnalyticsCustomerSatisfactionAnalyticsRetrieve**](CallApi.md#CallAnalyticsCustomerSatisfactionAnalyticsRetrieve) | **Get** /api/call/analytics/customer-satisfaction-analytics/ | 
[**CallAnalyticsFeedbackCategoryDistributionRetrieve**](CallApi.md#CallAnalyticsFeedbackCategoryDistributionRetrieve) | **Get** /api/call/analytics/feedback-category-distribution/ | 
[**CallAnalyticsInteractionTypeAnalyticsRetrieve**](CallApi.md#CallAnalyticsInteractionTypeAnalyticsRetrieve) | **Get** /api/call/analytics/interaction-type-analytics/ | 
[**CallCallAgentPerformanceCreate**](CallApi.md#CallCallAgentPerformanceCreate) | **Post** /api/call/call/agent-performance/ | 
[**CallCallAgentPerformanceDestroy**](CallApi.md#CallCallAgentPerformanceDestroy) | **Delete** /api/call/call/agent-performance/ | 
[**CallCallAgentPerformancePartialUpdate**](CallApi.md#CallCallAgentPerformancePartialUpdate) | **Patch** /api/call/call/agent-performance/ | 
[**CallCallAgentPerformanceRetrieve**](CallApi.md#CallCallAgentPerformanceRetrieve) | **Get** /api/call/call/agent-performance/ | 
[**CallCallCallMetricsCreate**](CallApi.md#CallCallCallMetricsCreate) | **Post** /api/call/call/call-metrics/ | 
[**CallCallCallMetricsDestroy**](CallApi.md#CallCallCallMetricsDestroy) | **Delete** /api/call/call/call-metrics/ | 
[**CallCallCallMetricsPartialUpdate**](CallApi.md#CallCallCallMetricsPartialUpdate) | **Patch** /api/call/call/call-metrics/ | 
[**CallCallCallMetricsRetrieve**](CallApi.md#CallCallCallMetricsRetrieve) | **Get** /api/call/call/call-metrics/ | 
[**CallCallCallResolutionCreate**](CallApi.md#CallCallCallResolutionCreate) | **Post** /api/call/call/call-resolution/ | 
[**CallCallCallResolutionDestroy**](CallApi.md#CallCallCallResolutionDestroy) | **Delete** /api/call/call/call-resolution/ | 
[**CallCallCallResolutionPartialUpdate**](CallApi.md#CallCallCallResolutionPartialUpdate) | **Patch** /api/call/call/call-resolution/ | 
[**CallCallCallResolutionRetrieve**](CallApi.md#CallCallCallResolutionRetrieve) | **Get** /api/call/call/call-resolution/ | 
[**CallCallCallTagCreate**](CallApi.md#CallCallCallTagCreate) | **Post** /api/call/call/call-tag/ | 
[**CallCallCallTagDestroy**](CallApi.md#CallCallCallTagDestroy) | **Delete** /api/call/call/call-tag/ | 
[**CallCallCallTagMappingCreate**](CallApi.md#CallCallCallTagMappingCreate) | **Post** /api/call/call/call-tag-mapping/ | 
[**CallCallCallTagMappingDestroy**](CallApi.md#CallCallCallTagMappingDestroy) | **Delete** /api/call/call/call-tag-mapping/ | 
[**CallCallCallTagMappingRetrieve**](CallApi.md#CallCallCallTagMappingRetrieve) | **Get** /api/call/call/call-tag-mapping/ | 
[**CallCallCallTagRetrieve**](CallApi.md#CallCallCallTagRetrieve) | **Get** /api/call/call/call-tag/ | 
[**CallCallCallsCreate**](CallApi.md#CallCallCallsCreate) | **Post** /api/call/call/calls/ | 
[**CallCallCallsDestroy**](CallApi.md#CallCallCallsDestroy) | **Delete** /api/call/call/calls/ | 
[**CallCallCallsPartialUpdate**](CallApi.md#CallCallCallsPartialUpdate) | **Patch** /api/call/call/calls/ | 
[**CallCallCallsRetrieve**](CallApi.md#CallCallCallsRetrieve) | **Get** /api/call/call/calls/ | 
[**CallCallCustomerFeedbackCreate**](CallApi.md#CallCallCustomerFeedbackCreate) | **Post** /api/call/call/customer-feedback/ | 
[**CallCallCustomerFeedbackDestroy**](CallApi.md#CallCallCustomerFeedbackDestroy) | **Delete** /api/call/call/customer-feedback/ | 
[**CallCallCustomerFeedbackPartialUpdate**](CallApi.md#CallCallCustomerFeedbackPartialUpdate) | **Patch** /api/call/call/customer-feedback/ | 
[**CallCallCustomerFeedbackRetrieve**](CallApi.md#CallCallCustomerFeedbackRetrieve) | **Get** /api/call/call/customer-feedback/ | 
[**CallCallInteractionHistoryCreate**](CallApi.md#CallCallInteractionHistoryCreate) | **Post** /api/call/call/interaction-history/ | 
[**CallCallInteractionHistoryDestroy**](CallApi.md#CallCallInteractionHistoryDestroy) | **Delete** /api/call/call/interaction-history/ | 
[**CallCallInteractionHistoryPartialUpdate**](CallApi.md#CallCallInteractionHistoryPartialUpdate) | **Patch** /api/call/call/interaction-history/ | 
[**CallCallInteractionHistoryRetrieve**](CallApi.md#CallCallInteractionHistoryRetrieve) | **Get** /api/call/call/interaction-history/ | 



## CallAnalyticsCallResolutionAnalyticsRetrieve

> CallAnalyticsCallResolutionAnalyticsRetrieve(ctx).Execute()





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
    resp, r, err := api_client.CallApi.CallAnalyticsCallResolutionAnalyticsRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallApi.CallAnalyticsCallResolutionAnalyticsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCallAnalyticsCallResolutionAnalyticsRetrieveRequest struct via the builder pattern


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


## CallAnalyticsCallTagAnalyticsRetrieve

> CallAnalyticsCallTagAnalyticsRetrieve(ctx).Execute()





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
    resp, r, err := api_client.CallApi.CallAnalyticsCallTagAnalyticsRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallApi.CallAnalyticsCallTagAnalyticsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCallAnalyticsCallTagAnalyticsRetrieveRequest struct via the builder pattern


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


## CallAnalyticsCallVolumeAnalyticsRetrieve

> CallAnalyticsCallVolumeAnalyticsRetrieve(ctx).Execute()





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
    resp, r, err := api_client.CallApi.CallAnalyticsCallVolumeAnalyticsRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallApi.CallAnalyticsCallVolumeAnalyticsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCallAnalyticsCallVolumeAnalyticsRetrieveRequest struct via the builder pattern


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


## CallAnalyticsCustomerSatisfactionAnalyticsRetrieve

> CallAnalyticsCustomerSatisfactionAnalyticsRetrieve(ctx).Execute()





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
    resp, r, err := api_client.CallApi.CallAnalyticsCustomerSatisfactionAnalyticsRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallApi.CallAnalyticsCustomerSatisfactionAnalyticsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCallAnalyticsCustomerSatisfactionAnalyticsRetrieveRequest struct via the builder pattern


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


## CallAnalyticsFeedbackCategoryDistributionRetrieve

> CallAnalyticsFeedbackCategoryDistributionRetrieve(ctx).Execute()





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
    resp, r, err := api_client.CallApi.CallAnalyticsFeedbackCategoryDistributionRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallApi.CallAnalyticsFeedbackCategoryDistributionRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCallAnalyticsFeedbackCategoryDistributionRetrieveRequest struct via the builder pattern


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


## CallAnalyticsInteractionTypeAnalyticsRetrieve

> CallAnalyticsInteractionTypeAnalyticsRetrieve(ctx).Execute()





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
    resp, r, err := api_client.CallApi.CallAnalyticsInteractionTypeAnalyticsRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallApi.CallAnalyticsInteractionTypeAnalyticsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCallAnalyticsInteractionTypeAnalyticsRetrieveRequest struct via the builder pattern


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


## CallCallAgentPerformanceCreate

> AgentPerformance CallCallAgentPerformanceCreate(ctx).AgentPerformance(agentPerformance).Execute()





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
    agentPerformance := *openapiclient.NewAgentPerformance(int32(123), *openapiclient.NewTenantInfo(int32(123), "TenantName_example"), "AverageCallDuration_example", "AverageSatisfactionScore_example", "ResolutionRate_example", int32(123)) // AgentPerformance | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CallApi.CallCallAgentPerformanceCreate(context.Background()).AgentPerformance(agentPerformance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallApi.CallCallAgentPerformanceCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CallCallAgentPerformanceCreate`: AgentPerformance
    fmt.Fprintf(os.Stdout, "Response from `CallApi.CallCallAgentPerformanceCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCallCallAgentPerformanceCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agentPerformance** | [**AgentPerformance**](AgentPerformance.md) |  | 

### Return type

[**AgentPerformance**](AgentPerformance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallCallAgentPerformanceDestroy

> CallCallAgentPerformanceDestroy(ctx).Execute()





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
    resp, r, err := api_client.CallApi.CallCallAgentPerformanceDestroy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallApi.CallCallAgentPerformanceDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCallCallAgentPerformanceDestroyRequest struct via the builder pattern


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


## CallCallAgentPerformancePartialUpdate

> AgentPerformance CallCallAgentPerformancePartialUpdate(ctx).PatchedAgentPerformance(patchedAgentPerformance).Execute()





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
    patchedAgentPerformance := *openapiclient.NewPatchedAgentPerformance() // PatchedAgentPerformance |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CallApi.CallCallAgentPerformancePartialUpdate(context.Background()).PatchedAgentPerformance(patchedAgentPerformance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallApi.CallCallAgentPerformancePartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CallCallAgentPerformancePartialUpdate`: AgentPerformance
    fmt.Fprintf(os.Stdout, "Response from `CallApi.CallCallAgentPerformancePartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCallCallAgentPerformancePartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedAgentPerformance** | [**PatchedAgentPerformance**](PatchedAgentPerformance.md) |  | 

### Return type

[**AgentPerformance**](AgentPerformance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallCallAgentPerformanceRetrieve

> AgentPerformance CallCallAgentPerformanceRetrieve(ctx).Execute()





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
    resp, r, err := api_client.CallApi.CallCallAgentPerformanceRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallApi.CallCallAgentPerformanceRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CallCallAgentPerformanceRetrieve`: AgentPerformance
    fmt.Fprintf(os.Stdout, "Response from `CallApi.CallCallAgentPerformanceRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCallCallAgentPerformanceRetrieveRequest struct via the builder pattern


### Return type

[**AgentPerformance**](AgentPerformance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallCallCallMetricsCreate

> CallMetrics CallCallCallMetricsCreate(ctx).CallMetrics(callMetrics).Execute()





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
    callMetrics := *openapiclient.NewCallMetrics(int32(123), *openapiclient.NewCall(int32(123), *openapiclient.NewTenantInfo(int32(123), "TenantName_example"), "Duration_example", false, *openapiclient.NewClientInfo(int32(123), "Username_example", "Email_example", "FirstName_example", "LastName_example")), interface{}(123)) // CallMetrics | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CallApi.CallCallCallMetricsCreate(context.Background()).CallMetrics(callMetrics).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallApi.CallCallCallMetricsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CallCallCallMetricsCreate`: CallMetrics
    fmt.Fprintf(os.Stdout, "Response from `CallApi.CallCallCallMetricsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCallCallCallMetricsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **callMetrics** | [**CallMetrics**](CallMetrics.md) |  | 

### Return type

[**CallMetrics**](CallMetrics.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallCallCallMetricsDestroy

> CallCallCallMetricsDestroy(ctx).Execute()





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
    resp, r, err := api_client.CallApi.CallCallCallMetricsDestroy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallApi.CallCallCallMetricsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCallCallCallMetricsDestroyRequest struct via the builder pattern


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


## CallCallCallMetricsPartialUpdate

> CallMetrics CallCallCallMetricsPartialUpdate(ctx).PatchedCallMetrics(patchedCallMetrics).Execute()





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
    patchedCallMetrics := *openapiclient.NewPatchedCallMetrics() // PatchedCallMetrics |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CallApi.CallCallCallMetricsPartialUpdate(context.Background()).PatchedCallMetrics(patchedCallMetrics).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallApi.CallCallCallMetricsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CallCallCallMetricsPartialUpdate`: CallMetrics
    fmt.Fprintf(os.Stdout, "Response from `CallApi.CallCallCallMetricsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCallCallCallMetricsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedCallMetrics** | [**PatchedCallMetrics**](PatchedCallMetrics.md) |  | 

### Return type

[**CallMetrics**](CallMetrics.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallCallCallMetricsRetrieve

> CallMetrics CallCallCallMetricsRetrieve(ctx).Execute()





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
    resp, r, err := api_client.CallApi.CallCallCallMetricsRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallApi.CallCallCallMetricsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CallCallCallMetricsRetrieve`: CallMetrics
    fmt.Fprintf(os.Stdout, "Response from `CallApi.CallCallCallMetricsRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCallCallCallMetricsRetrieveRequest struct via the builder pattern


### Return type

[**CallMetrics**](CallMetrics.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallCallCallResolutionCreate

> CallResolution CallCallCallResolutionCreate(ctx).CallResolution(callResolution).Execute()





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
    callResolution := *openapiclient.NewCallResolution(int32(123), *openapiclient.NewCall(int32(123), *openapiclient.NewTenantInfo(int32(123), "TenantName_example"), "Duration_example", false, *openapiclient.NewClientInfo(int32(123), "Username_example", "Email_example", "FirstName_example", "LastName_example"))) // CallResolution |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CallApi.CallCallCallResolutionCreate(context.Background()).CallResolution(callResolution).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallApi.CallCallCallResolutionCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CallCallCallResolutionCreate`: CallResolution
    fmt.Fprintf(os.Stdout, "Response from `CallApi.CallCallCallResolutionCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCallCallCallResolutionCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **callResolution** | [**CallResolution**](CallResolution.md) |  | 

### Return type

[**CallResolution**](CallResolution.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallCallCallResolutionDestroy

> CallCallCallResolutionDestroy(ctx).Execute()





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
    resp, r, err := api_client.CallApi.CallCallCallResolutionDestroy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallApi.CallCallCallResolutionDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCallCallCallResolutionDestroyRequest struct via the builder pattern


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


## CallCallCallResolutionPartialUpdate

> CallResolution CallCallCallResolutionPartialUpdate(ctx).PatchedCallResolution(patchedCallResolution).Execute()





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
    patchedCallResolution := *openapiclient.NewPatchedCallResolution() // PatchedCallResolution |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CallApi.CallCallCallResolutionPartialUpdate(context.Background()).PatchedCallResolution(patchedCallResolution).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallApi.CallCallCallResolutionPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CallCallCallResolutionPartialUpdate`: CallResolution
    fmt.Fprintf(os.Stdout, "Response from `CallApi.CallCallCallResolutionPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCallCallCallResolutionPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedCallResolution** | [**PatchedCallResolution**](PatchedCallResolution.md) |  | 

### Return type

[**CallResolution**](CallResolution.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallCallCallResolutionRetrieve

> CallResolution CallCallCallResolutionRetrieve(ctx).Execute()





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
    resp, r, err := api_client.CallApi.CallCallCallResolutionRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallApi.CallCallCallResolutionRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CallCallCallResolutionRetrieve`: CallResolution
    fmt.Fprintf(os.Stdout, "Response from `CallApi.CallCallCallResolutionRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCallCallCallResolutionRetrieveRequest struct via the builder pattern


### Return type

[**CallResolution**](CallResolution.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallCallCallTagCreate

> CallTag CallCallCallTagCreate(ctx).CallTag(callTag).Execute()





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
    callTag := *openapiclient.NewCallTag(int32(123), "Name_example") // CallTag | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CallApi.CallCallCallTagCreate(context.Background()).CallTag(callTag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallApi.CallCallCallTagCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CallCallCallTagCreate`: CallTag
    fmt.Fprintf(os.Stdout, "Response from `CallApi.CallCallCallTagCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCallCallCallTagCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **callTag** | [**CallTag**](CallTag.md) |  | 

### Return type

[**CallTag**](CallTag.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallCallCallTagDestroy

> CallCallCallTagDestroy(ctx).Execute()





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
    resp, r, err := api_client.CallApi.CallCallCallTagDestroy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallApi.CallCallCallTagDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCallCallCallTagDestroyRequest struct via the builder pattern


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


## CallCallCallTagMappingCreate

> CallTagMapping CallCallCallTagMappingCreate(ctx).CallTagMapping(callTagMapping).Execute()





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
    callTagMapping := *openapiclient.NewCallTagMapping(int32(123), *openapiclient.NewCall(int32(123), *openapiclient.NewTenantInfo(int32(123), "TenantName_example"), "Duration_example", false, *openapiclient.NewClientInfo(int32(123), "Username_example", "Email_example", "FirstName_example", "LastName_example")), *openapiclient.NewCallTag(int32(123), "Name_example")) // CallTagMapping |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CallApi.CallCallCallTagMappingCreate(context.Background()).CallTagMapping(callTagMapping).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallApi.CallCallCallTagMappingCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CallCallCallTagMappingCreate`: CallTagMapping
    fmt.Fprintf(os.Stdout, "Response from `CallApi.CallCallCallTagMappingCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCallCallCallTagMappingCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **callTagMapping** | [**CallTagMapping**](CallTagMapping.md) |  | 

### Return type

[**CallTagMapping**](CallTagMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallCallCallTagMappingDestroy

> CallCallCallTagMappingDestroy(ctx).Execute()





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
    resp, r, err := api_client.CallApi.CallCallCallTagMappingDestroy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallApi.CallCallCallTagMappingDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCallCallCallTagMappingDestroyRequest struct via the builder pattern


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


## CallCallCallTagMappingRetrieve

> CallTagMapping CallCallCallTagMappingRetrieve(ctx).Execute()





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
    resp, r, err := api_client.CallApi.CallCallCallTagMappingRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallApi.CallCallCallTagMappingRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CallCallCallTagMappingRetrieve`: CallTagMapping
    fmt.Fprintf(os.Stdout, "Response from `CallApi.CallCallCallTagMappingRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCallCallCallTagMappingRetrieveRequest struct via the builder pattern


### Return type

[**CallTagMapping**](CallTagMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallCallCallTagRetrieve

> CallTag CallCallCallTagRetrieve(ctx).Execute()





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
    resp, r, err := api_client.CallApi.CallCallCallTagRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallApi.CallCallCallTagRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CallCallCallTagRetrieve`: CallTag
    fmt.Fprintf(os.Stdout, "Response from `CallApi.CallCallCallTagRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCallCallCallTagRetrieveRequest struct via the builder pattern


### Return type

[**CallTag**](CallTag.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallCallCallsCreate

> Call CallCallCallsCreate(ctx).Call(call).Execute()





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
    call := *openapiclient.NewCall(int32(123), *openapiclient.NewTenantInfo(int32(123), "TenantName_example"), "Duration_example", false, *openapiclient.NewClientInfo(int32(123), "Username_example", "Email_example", "FirstName_example", "LastName_example")) // Call | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CallApi.CallCallCallsCreate(context.Background()).Call(call).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallApi.CallCallCallsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CallCallCallsCreate`: Call
    fmt.Fprintf(os.Stdout, "Response from `CallApi.CallCallCallsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCallCallCallsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **call** | [**Call**](Call.md) |  | 

### Return type

[**Call**](Call.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallCallCallsDestroy

> CallCallCallsDestroy(ctx).Execute()





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
    resp, r, err := api_client.CallApi.CallCallCallsDestroy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallApi.CallCallCallsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCallCallCallsDestroyRequest struct via the builder pattern


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


## CallCallCallsPartialUpdate

> Call CallCallCallsPartialUpdate(ctx).PatchedCall(patchedCall).Execute()





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
    patchedCall := *openapiclient.NewPatchedCall() // PatchedCall |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CallApi.CallCallCallsPartialUpdate(context.Background()).PatchedCall(patchedCall).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallApi.CallCallCallsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CallCallCallsPartialUpdate`: Call
    fmt.Fprintf(os.Stdout, "Response from `CallApi.CallCallCallsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCallCallCallsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedCall** | [**PatchedCall**](PatchedCall.md) |  | 

### Return type

[**Call**](Call.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallCallCallsRetrieve

> Call CallCallCallsRetrieve(ctx).Execute()





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
    resp, r, err := api_client.CallApi.CallCallCallsRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallApi.CallCallCallsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CallCallCallsRetrieve`: Call
    fmt.Fprintf(os.Stdout, "Response from `CallApi.CallCallCallsRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCallCallCallsRetrieveRequest struct via the builder pattern


### Return type

[**Call**](Call.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallCallCustomerFeedbackCreate

> CustomerFeedback CallCallCustomerFeedbackCreate(ctx).CustomerFeedback(customerFeedback).Execute()





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
    customerFeedback := *openapiclient.NewCustomerFeedback(int32(123), *openapiclient.NewCall(int32(123), *openapiclient.NewTenantInfo(int32(123), "TenantName_example"), "Duration_example", false, *openapiclient.NewClientInfo(int32(123), "Username_example", "Email_example", "FirstName_example", "LastName_example")), int32(123), openapiclient.FeedbackCategoryEnum("SERVICE")) // CustomerFeedback | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CallApi.CallCallCustomerFeedbackCreate(context.Background()).CustomerFeedback(customerFeedback).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallApi.CallCallCustomerFeedbackCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CallCallCustomerFeedbackCreate`: CustomerFeedback
    fmt.Fprintf(os.Stdout, "Response from `CallApi.CallCallCustomerFeedbackCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCallCallCustomerFeedbackCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerFeedback** | [**CustomerFeedback**](CustomerFeedback.md) |  | 

### Return type

[**CustomerFeedback**](CustomerFeedback.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallCallCustomerFeedbackDestroy

> CallCallCustomerFeedbackDestroy(ctx).Execute()





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
    resp, r, err := api_client.CallApi.CallCallCustomerFeedbackDestroy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallApi.CallCallCustomerFeedbackDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCallCallCustomerFeedbackDestroyRequest struct via the builder pattern


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


## CallCallCustomerFeedbackPartialUpdate

> CustomerFeedback CallCallCustomerFeedbackPartialUpdate(ctx).PatchedCustomerFeedback(patchedCustomerFeedback).Execute()





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
    patchedCustomerFeedback := *openapiclient.NewPatchedCustomerFeedback() // PatchedCustomerFeedback |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CallApi.CallCallCustomerFeedbackPartialUpdate(context.Background()).PatchedCustomerFeedback(patchedCustomerFeedback).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallApi.CallCallCustomerFeedbackPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CallCallCustomerFeedbackPartialUpdate`: CustomerFeedback
    fmt.Fprintf(os.Stdout, "Response from `CallApi.CallCallCustomerFeedbackPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCallCallCustomerFeedbackPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedCustomerFeedback** | [**PatchedCustomerFeedback**](PatchedCustomerFeedback.md) |  | 

### Return type

[**CustomerFeedback**](CustomerFeedback.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallCallCustomerFeedbackRetrieve

> CustomerFeedback CallCallCustomerFeedbackRetrieve(ctx).Execute()





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
    resp, r, err := api_client.CallApi.CallCallCustomerFeedbackRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallApi.CallCallCustomerFeedbackRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CallCallCustomerFeedbackRetrieve`: CustomerFeedback
    fmt.Fprintf(os.Stdout, "Response from `CallApi.CallCallCustomerFeedbackRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCallCallCustomerFeedbackRetrieveRequest struct via the builder pattern


### Return type

[**CustomerFeedback**](CustomerFeedback.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallCallInteractionHistoryCreate

> InteractionHistory CallCallInteractionHistoryCreate(ctx).InteractionHistory(interactionHistory).Execute()





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
    interactionHistory := *openapiclient.NewInteractionHistory(int32(123), *openapiclient.NewCall(int32(123), *openapiclient.NewTenantInfo(int32(123), "TenantName_example"), "Duration_example", false, *openapiclient.NewClientInfo(int32(123), "Username_example", "Email_example", "FirstName_example", "LastName_example")), openapiclient.InteractionTypeEnum("NOTE"), "Details_example") // InteractionHistory | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CallApi.CallCallInteractionHistoryCreate(context.Background()).InteractionHistory(interactionHistory).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallApi.CallCallInteractionHistoryCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CallCallInteractionHistoryCreate`: InteractionHistory
    fmt.Fprintf(os.Stdout, "Response from `CallApi.CallCallInteractionHistoryCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCallCallInteractionHistoryCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **interactionHistory** | [**InteractionHistory**](InteractionHistory.md) |  | 

### Return type

[**InteractionHistory**](InteractionHistory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallCallInteractionHistoryDestroy

> CallCallInteractionHistoryDestroy(ctx).Execute()





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
    resp, r, err := api_client.CallApi.CallCallInteractionHistoryDestroy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallApi.CallCallInteractionHistoryDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCallCallInteractionHistoryDestroyRequest struct via the builder pattern


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


## CallCallInteractionHistoryPartialUpdate

> InteractionHistory CallCallInteractionHistoryPartialUpdate(ctx).PatchedInteractionHistory(patchedInteractionHistory).Execute()





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
    patchedInteractionHistory := *openapiclient.NewPatchedInteractionHistory() // PatchedInteractionHistory |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CallApi.CallCallInteractionHistoryPartialUpdate(context.Background()).PatchedInteractionHistory(patchedInteractionHistory).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallApi.CallCallInteractionHistoryPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CallCallInteractionHistoryPartialUpdate`: InteractionHistory
    fmt.Fprintf(os.Stdout, "Response from `CallApi.CallCallInteractionHistoryPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCallCallInteractionHistoryPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedInteractionHistory** | [**PatchedInteractionHistory**](PatchedInteractionHistory.md) |  | 

### Return type

[**InteractionHistory**](InteractionHistory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallCallInteractionHistoryRetrieve

> InteractionHistory CallCallInteractionHistoryRetrieve(ctx).Execute()





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
    resp, r, err := api_client.CallApi.CallCallInteractionHistoryRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CallApi.CallCallInteractionHistoryRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CallCallInteractionHistoryRetrieve`: InteractionHistory
    fmt.Fprintf(os.Stdout, "Response from `CallApi.CallCallInteractionHistoryRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCallCallInteractionHistoryRetrieveRequest struct via the builder pattern


### Return type

[**InteractionHistory**](InteractionHistory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

