# \ServicesApi

All URIs are relative to *https://core.proximaai.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServicesServiceformsCreate**](ServicesApi.md#ServicesServiceformsCreate) | **Post** /api/services/serviceforms/ | 
[**ServicesServiceformsDestroy**](ServicesApi.md#ServicesServiceformsDestroy) | **Delete** /api/services/serviceforms/ | 
[**ServicesServiceformsPartialUpdate**](ServicesApi.md#ServicesServiceformsPartialUpdate) | **Patch** /api/services/serviceforms/ | 
[**ServicesServiceformsRetrieve**](ServicesApi.md#ServicesServiceformsRetrieve) | **Get** /api/services/serviceforms/ | 
[**ServicesServiceformsresponseCreate**](ServicesApi.md#ServicesServiceformsresponseCreate) | **Post** /api/services/serviceformsresponse/ | 
[**ServicesServiceformsresponseDestroy**](ServicesApi.md#ServicesServiceformsresponseDestroy) | **Delete** /api/services/serviceformsresponse/ | 
[**ServicesServiceformsresponsePartialUpdate**](ServicesApi.md#ServicesServiceformsresponsePartialUpdate) | **Patch** /api/services/serviceformsresponse/ | 
[**ServicesServiceformsresponseRetrieve**](ServicesApi.md#ServicesServiceformsresponseRetrieve) | **Get** /api/services/serviceformsresponse/ | 
[**ServicesServicerequestchatCreate**](ServicesApi.md#ServicesServicerequestchatCreate) | **Post** /api/services/servicerequestchat/ | 
[**ServicesServicerequestchatRetrieve**](ServicesApi.md#ServicesServicerequestchatRetrieve) | **Get** /api/services/servicerequestchat/ | 
[**ServicesServicerequestchatmessageCreate**](ServicesApi.md#ServicesServicerequestchatmessageCreate) | **Post** /api/services/servicerequestchatmessage/ | 
[**ServicesServicerequestchatmessageRetrieve**](ServicesApi.md#ServicesServicerequestchatmessageRetrieve) | **Get** /api/services/servicerequestchatmessage/ | 
[**ServicesServiceresquestCreate**](ServicesApi.md#ServicesServiceresquestCreate) | **Post** /api/services/serviceresquest/ | 
[**ServicesServiceresquestDestroy**](ServicesApi.md#ServicesServiceresquestDestroy) | **Delete** /api/services/serviceresquest/ | 
[**ServicesServiceresquestPartialUpdate**](ServicesApi.md#ServicesServiceresquestPartialUpdate) | **Patch** /api/services/serviceresquest/ | 
[**ServicesServiceresquestRetrieve**](ServicesApi.md#ServicesServiceresquestRetrieve) | **Get** /api/services/serviceresquest/ | 
[**ServicesServicesCreate**](ServicesApi.md#ServicesServicesCreate) | **Post** /api/services/services/ | 
[**ServicesServicesDestroy**](ServicesApi.md#ServicesServicesDestroy) | **Delete** /api/services/services/ | 
[**ServicesServicesPartialUpdate**](ServicesApi.md#ServicesServicesPartialUpdate) | **Patch** /api/services/services/ | 
[**ServicesServicesRetrieve**](ServicesApi.md#ServicesServicesRetrieve) | **Get** /api/services/services/ | 
[**ServicesServicetypeCreate**](ServicesApi.md#ServicesServicetypeCreate) | **Post** /api/services/servicetype/ | 
[**ServicesServicetypeDestroy**](ServicesApi.md#ServicesServicetypeDestroy) | **Delete** /api/services/servicetype/ | 
[**ServicesServicetypeRetrieve**](ServicesApi.md#ServicesServicetypeRetrieve) | **Get** /api/services/servicetype/ | 



## ServicesServiceformsCreate

> ServiceForms ServicesServiceformsCreate(ctx).ServiceForms(serviceForms).Execute()





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
    serviceForms := *openapiclient.NewServiceForms(int32(123), *openapiclient.NewService(int32(123), int32(123), "NameOfService_example", "ServiceDescription_example", int32(123), "ResolutionPeriod_example", openapiclient.ServiceAvailabilityEnum("All_Clients"), interface{}(123), interface{}(123)), "NameOfForm_example", interface{}(123)) // ServiceForms | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicesApi.ServicesServiceformsCreate(context.Background()).ServiceForms(serviceForms).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.ServicesServiceformsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServicesServiceformsCreate`: ServiceForms
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.ServicesServiceformsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServicesServiceformsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceForms** | [**ServiceForms**](ServiceForms.md) |  | 

### Return type

[**ServiceForms**](ServiceForms.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicesServiceformsDestroy

> ServicesServiceformsDestroy(ctx).Execute()





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
    resp, r, err := api_client.ServicesApi.ServicesServiceformsDestroy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.ServicesServiceformsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiServicesServiceformsDestroyRequest struct via the builder pattern


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


## ServicesServiceformsPartialUpdate

> ServiceForms ServicesServiceformsPartialUpdate(ctx).PatchedServiceForms(patchedServiceForms).Execute()





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
    patchedServiceForms := *openapiclient.NewPatchedServiceForms() // PatchedServiceForms |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicesApi.ServicesServiceformsPartialUpdate(context.Background()).PatchedServiceForms(patchedServiceForms).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.ServicesServiceformsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServicesServiceformsPartialUpdate`: ServiceForms
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.ServicesServiceformsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServicesServiceformsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedServiceForms** | [**PatchedServiceForms**](PatchedServiceForms.md) |  | 

### Return type

[**ServiceForms**](ServiceForms.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicesServiceformsRetrieve

> ServiceForms ServicesServiceformsRetrieve(ctx).Execute()





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
    resp, r, err := api_client.ServicesApi.ServicesServiceformsRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.ServicesServiceformsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServicesServiceformsRetrieve`: ServiceForms
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.ServicesServiceformsRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiServicesServiceformsRetrieveRequest struct via the builder pattern


### Return type

[**ServiceForms**](ServiceForms.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicesServiceformsresponseCreate

> ServiceFormsResponse ServicesServiceformsresponseCreate(ctx).ServiceFormsResponse(serviceFormsResponse).Execute()





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
    serviceFormsResponse := *openapiclient.NewServiceFormsResponse("CreatedAt_example", int32(123), *openapiclient.NewServiceForms(int32(123), *openapiclient.NewService(int32(123), int32(123), "NameOfService_example", "ServiceDescription_example", int32(123), "ResolutionPeriod_example", openapiclient.ServiceAvailabilityEnum("All_Clients"), interface{}(123), interface{}(123)), "NameOfForm_example", interface{}(123)), *openapiclient.NewClientServicesInfo(int32(123), "Username_example", "Email_example", "FirstName_example", "LastName_example"), *openapiclient.NewAnonymousUser(int32(123), "Token_example"), interface{}(123)) // ServiceFormsResponse | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicesApi.ServicesServiceformsresponseCreate(context.Background()).ServiceFormsResponse(serviceFormsResponse).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.ServicesServiceformsresponseCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServicesServiceformsresponseCreate`: ServiceFormsResponse
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.ServicesServiceformsresponseCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServicesServiceformsresponseCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceFormsResponse** | [**ServiceFormsResponse**](ServiceFormsResponse.md) |  | 

### Return type

[**ServiceFormsResponse**](ServiceFormsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicesServiceformsresponseDestroy

> ServicesServiceformsresponseDestroy(ctx).Execute()





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
    resp, r, err := api_client.ServicesApi.ServicesServiceformsresponseDestroy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.ServicesServiceformsresponseDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiServicesServiceformsresponseDestroyRequest struct via the builder pattern


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


## ServicesServiceformsresponsePartialUpdate

> ServiceFormsResponse ServicesServiceformsresponsePartialUpdate(ctx).PatchedServiceFormsResponse(patchedServiceFormsResponse).Execute()





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
    patchedServiceFormsResponse := *openapiclient.NewPatchedServiceFormsResponse() // PatchedServiceFormsResponse |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicesApi.ServicesServiceformsresponsePartialUpdate(context.Background()).PatchedServiceFormsResponse(patchedServiceFormsResponse).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.ServicesServiceformsresponsePartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServicesServiceformsresponsePartialUpdate`: ServiceFormsResponse
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.ServicesServiceformsresponsePartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServicesServiceformsresponsePartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedServiceFormsResponse** | [**PatchedServiceFormsResponse**](PatchedServiceFormsResponse.md) |  | 

### Return type

[**ServiceFormsResponse**](ServiceFormsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicesServiceformsresponseRetrieve

> ServiceFormsResponse ServicesServiceformsresponseRetrieve(ctx).Execute()





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
    resp, r, err := api_client.ServicesApi.ServicesServiceformsresponseRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.ServicesServiceformsresponseRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServicesServiceformsresponseRetrieve`: ServiceFormsResponse
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.ServicesServiceformsresponseRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiServicesServiceformsresponseRetrieveRequest struct via the builder pattern


### Return type

[**ServiceFormsResponse**](ServiceFormsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicesServicerequestchatCreate

> ServiceMessageChat ServicesServicerequestchatCreate(ctx).ServiceMessageChat(serviceMessageChat).Execute()





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
    serviceMessageChat := *openapiclient.NewServiceMessageChat(int32(123), int32(123), "MessageContent_example", openapiclient.MessageSenderEnum("client")) // ServiceMessageChat | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicesApi.ServicesServicerequestchatCreate(context.Background()).ServiceMessageChat(serviceMessageChat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.ServicesServicerequestchatCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServicesServicerequestchatCreate`: ServiceMessageChat
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.ServicesServicerequestchatCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServicesServicerequestchatCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceMessageChat** | [**ServiceMessageChat**](ServiceMessageChat.md) |  | 

### Return type

[**ServiceMessageChat**](ServiceMessageChat.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicesServicerequestchatRetrieve

> ServiceMessageChat ServicesServicerequestchatRetrieve(ctx).Execute()





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
    resp, r, err := api_client.ServicesApi.ServicesServicerequestchatRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.ServicesServicerequestchatRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServicesServicerequestchatRetrieve`: ServiceMessageChat
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.ServicesServicerequestchatRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiServicesServicerequestchatRetrieveRequest struct via the builder pattern


### Return type

[**ServiceMessageChat**](ServiceMessageChat.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicesServicerequestchatmessageCreate

> ServiceRequestChat ServicesServicerequestchatmessageCreate(ctx).ServiceRequestChat(serviceRequestChat).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    serviceRequestChat := *openapiclient.NewServiceRequestChat(int32(123), *openapiclient.NewServiceRequest("CreatedAt_example", int32(123), *openapiclient.NewClientServicesInfo(int32(123), "Username_example", "Email_example", "FirstName_example", "LastName_example"), *openapiclient.NewAnonymousUser(int32(123), "Token_example"), *openapiclient.NewService(int32(123), int32(123), "NameOfService_example", "ServiceDescription_example", int32(123), "ResolutionPeriod_example", openapiclient.ServiceAvailabilityEnum("All_Clients"), interface{}(123), interface{}(123)), interface{}(123), interface{}(123), interface{}(123)), time.Now(), time.Now()) // ServiceRequestChat |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicesApi.ServicesServicerequestchatmessageCreate(context.Background()).ServiceRequestChat(serviceRequestChat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.ServicesServicerequestchatmessageCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServicesServicerequestchatmessageCreate`: ServiceRequestChat
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.ServicesServicerequestchatmessageCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServicesServicerequestchatmessageCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceRequestChat** | [**ServiceRequestChat**](ServiceRequestChat.md) |  | 

### Return type

[**ServiceRequestChat**](ServiceRequestChat.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicesServicerequestchatmessageRetrieve

> ServiceRequestChat ServicesServicerequestchatmessageRetrieve(ctx).Execute()





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
    resp, r, err := api_client.ServicesApi.ServicesServicerequestchatmessageRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.ServicesServicerequestchatmessageRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServicesServicerequestchatmessageRetrieve`: ServiceRequestChat
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.ServicesServicerequestchatmessageRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiServicesServicerequestchatmessageRetrieveRequest struct via the builder pattern


### Return type

[**ServiceRequestChat**](ServiceRequestChat.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicesServiceresquestCreate

> ServiceRequest ServicesServiceresquestCreate(ctx).ServiceRequest(serviceRequest).Execute()





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
    serviceRequest := *openapiclient.NewServiceRequest("CreatedAt_example", int32(123), *openapiclient.NewClientServicesInfo(int32(123), "Username_example", "Email_example", "FirstName_example", "LastName_example"), *openapiclient.NewAnonymousUser(int32(123), "Token_example"), *openapiclient.NewService(int32(123), int32(123), "NameOfService_example", "ServiceDescription_example", int32(123), "ResolutionPeriod_example", openapiclient.ServiceAvailabilityEnum("All_Clients"), interface{}(123), interface{}(123)), interface{}(123), interface{}(123), interface{}(123)) // ServiceRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicesApi.ServicesServiceresquestCreate(context.Background()).ServiceRequest(serviceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.ServicesServiceresquestCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServicesServiceresquestCreate`: ServiceRequest
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.ServicesServiceresquestCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServicesServiceresquestCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceRequest** | [**ServiceRequest**](ServiceRequest.md) |  | 

### Return type

[**ServiceRequest**](ServiceRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicesServiceresquestDestroy

> ServicesServiceresquestDestroy(ctx).Execute()





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
    resp, r, err := api_client.ServicesApi.ServicesServiceresquestDestroy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.ServicesServiceresquestDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiServicesServiceresquestDestroyRequest struct via the builder pattern


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


## ServicesServiceresquestPartialUpdate

> ServiceRequest ServicesServiceresquestPartialUpdate(ctx).PatchedServiceRequest(patchedServiceRequest).Execute()





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
    patchedServiceRequest := *openapiclient.NewPatchedServiceRequest() // PatchedServiceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicesApi.ServicesServiceresquestPartialUpdate(context.Background()).PatchedServiceRequest(patchedServiceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.ServicesServiceresquestPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServicesServiceresquestPartialUpdate`: ServiceRequest
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.ServicesServiceresquestPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServicesServiceresquestPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedServiceRequest** | [**PatchedServiceRequest**](PatchedServiceRequest.md) |  | 

### Return type

[**ServiceRequest**](ServiceRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicesServiceresquestRetrieve

> ServiceRequest ServicesServiceresquestRetrieve(ctx).Execute()





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
    resp, r, err := api_client.ServicesApi.ServicesServiceresquestRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.ServicesServiceresquestRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServicesServiceresquestRetrieve`: ServiceRequest
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.ServicesServiceresquestRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiServicesServiceresquestRetrieveRequest struct via the builder pattern


### Return type

[**ServiceRequest**](ServiceRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicesServicesCreate

> Service ServicesServicesCreate(ctx).Service(service).Execute()





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
    service := *openapiclient.NewService(int32(123), int32(123), "NameOfService_example", "ServiceDescription_example", int32(123), "ResolutionPeriod_example", openapiclient.ServiceAvailabilityEnum("All_Clients"), interface{}(123), interface{}(123)) // Service | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicesApi.ServicesServicesCreate(context.Background()).Service(service).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.ServicesServicesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServicesServicesCreate`: Service
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.ServicesServicesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServicesServicesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **service** | [**Service**](Service.md) |  | 

### Return type

[**Service**](Service.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicesServicesDestroy

> ServicesServicesDestroy(ctx).Execute()





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
    resp, r, err := api_client.ServicesApi.ServicesServicesDestroy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.ServicesServicesDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiServicesServicesDestroyRequest struct via the builder pattern


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


## ServicesServicesPartialUpdate

> Service ServicesServicesPartialUpdate(ctx).PatchedService(patchedService).Execute()





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
    patchedService := *openapiclient.NewPatchedService() // PatchedService |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicesApi.ServicesServicesPartialUpdate(context.Background()).PatchedService(patchedService).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.ServicesServicesPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServicesServicesPartialUpdate`: Service
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.ServicesServicesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServicesServicesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedService** | [**PatchedService**](PatchedService.md) |  | 

### Return type

[**Service**](Service.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicesServicesRetrieve

> Service ServicesServicesRetrieve(ctx).Execute()





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
    resp, r, err := api_client.ServicesApi.ServicesServicesRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.ServicesServicesRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServicesServicesRetrieve`: Service
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.ServicesServicesRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiServicesServicesRetrieveRequest struct via the builder pattern


### Return type

[**Service**](Service.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicesServicetypeCreate

> ServiceType ServicesServicetypeCreate(ctx).ServiceType(serviceType).Execute()





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
    serviceType := *openapiclient.NewServiceType(int32(123), int32(123), openapiclient.ServiceTypeTypeEnum("problem_resolution"), "Description_example") // ServiceType | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicesApi.ServicesServicetypeCreate(context.Background()).ServiceType(serviceType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.ServicesServicetypeCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServicesServicetypeCreate`: ServiceType
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.ServicesServicetypeCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServicesServicetypeCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceType** | [**ServiceType**](ServiceType.md) |  | 

### Return type

[**ServiceType**](ServiceType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicesServicetypeDestroy

> ServicesServicetypeDestroy(ctx).Execute()





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
    resp, r, err := api_client.ServicesApi.ServicesServicetypeDestroy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.ServicesServicetypeDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiServicesServicetypeDestroyRequest struct via the builder pattern


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


## ServicesServicetypeRetrieve

> ServiceType ServicesServicetypeRetrieve(ctx).Execute()





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
    resp, r, err := api_client.ServicesApi.ServicesServicetypeRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.ServicesServicetypeRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServicesServicetypeRetrieve`: ServiceType
    fmt.Fprintf(os.Stdout, "Response from `ServicesApi.ServicesServicetypeRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiServicesServicetypeRetrieveRequest struct via the builder pattern


### Return type

[**ServiceType**](ServiceType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

