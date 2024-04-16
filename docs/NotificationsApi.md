# \NotificationsApi

All URIs are relative to *https://core.proximaai.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NotificationsDelegationCreate**](NotificationsApi.md#NotificationsDelegationCreate) | **Post** /api/notifications/delegation/ | 
[**NotificationsDelegationDestroy**](NotificationsApi.md#NotificationsDelegationDestroy) | **Delete** /api/notifications/delegation/ | 
[**NotificationsDelegationPartialUpdate**](NotificationsApi.md#NotificationsDelegationPartialUpdate) | **Patch** /api/notifications/delegation/ | 
[**NotificationsDelegationRetrieve**](NotificationsApi.md#NotificationsDelegationRetrieve) | **Get** /api/notifications/delegation/ | 
[**NotificationsNotificationsDestroy**](NotificationsApi.md#NotificationsNotificationsDestroy) | **Delete** /api/notifications/notifications/ | 
[**NotificationsNotificationsPartialUpdate**](NotificationsApi.md#NotificationsNotificationsPartialUpdate) | **Patch** /api/notifications/notifications/ | 
[**NotificationsNotificationsRetrieve**](NotificationsApi.md#NotificationsNotificationsRetrieve) | **Get** /api/notifications/notifications/ | 



## NotificationsDelegationCreate

> Delegation NotificationsDelegationCreate(ctx).Delegation(delegation).Execute()





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
    delegation := *openapiclient.NewDelegation(int32(123), *openapiclient.NewTenantInfo(int32(123), "TenantName_example"), []openapiclient.Employee{*openapiclient.NewEmployee(int32(123), "Username_example", "Email_example", "FirstName_example", "LastName_example", int32(123), "Password_example", "ConfirmPassword_example", "Token_example")}, []openapiclient.Admin{*openapiclient.NewAdmin("Username_example", "Email_example", "FirstName_example", "LastName_example", int32(123), "Password_example", "ConfirmPassword_example", "Token_example")}, time.Now()) // Delegation | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotificationsApi.NotificationsDelegationCreate(context.Background()).Delegation(delegation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.NotificationsDelegationCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NotificationsDelegationCreate`: Delegation
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.NotificationsDelegationCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNotificationsDelegationCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **delegation** | [**Delegation**](Delegation.md) |  | 

### Return type

[**Delegation**](Delegation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotificationsDelegationDestroy

> NotificationsDelegationDestroy(ctx).Execute()





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
    resp, r, err := api_client.NotificationsApi.NotificationsDelegationDestroy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.NotificationsDelegationDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNotificationsDelegationDestroyRequest struct via the builder pattern


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


## NotificationsDelegationPartialUpdate

> Delegation NotificationsDelegationPartialUpdate(ctx).PatchedDelegation(patchedDelegation).Execute()





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
    patchedDelegation := *openapiclient.NewPatchedDelegation() // PatchedDelegation |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotificationsApi.NotificationsDelegationPartialUpdate(context.Background()).PatchedDelegation(patchedDelegation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.NotificationsDelegationPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NotificationsDelegationPartialUpdate`: Delegation
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.NotificationsDelegationPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNotificationsDelegationPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedDelegation** | [**PatchedDelegation**](PatchedDelegation.md) |  | 

### Return type

[**Delegation**](Delegation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotificationsDelegationRetrieve

> Delegation NotificationsDelegationRetrieve(ctx).Execute()





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
    resp, r, err := api_client.NotificationsApi.NotificationsDelegationRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.NotificationsDelegationRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NotificationsDelegationRetrieve`: Delegation
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.NotificationsDelegationRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNotificationsDelegationRetrieveRequest struct via the builder pattern


### Return type

[**Delegation**](Delegation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotificationsNotificationsDestroy

> NotificationsNotificationsDestroy(ctx).Execute()





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
    resp, r, err := api_client.NotificationsApi.NotificationsNotificationsDestroy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.NotificationsNotificationsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNotificationsNotificationsDestroyRequest struct via the builder pattern


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


## NotificationsNotificationsPartialUpdate

> Notifications NotificationsNotificationsPartialUpdate(ctx).PatchedNotifications(patchedNotifications).Execute()





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
    patchedNotifications := *openapiclient.NewPatchedNotifications() // PatchedNotifications |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotificationsApi.NotificationsNotificationsPartialUpdate(context.Background()).PatchedNotifications(patchedNotifications).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.NotificationsNotificationsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NotificationsNotificationsPartialUpdate`: Notifications
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.NotificationsNotificationsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNotificationsNotificationsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedNotifications** | [**PatchedNotifications**](PatchedNotifications.md) |  | 

### Return type

[**Notifications**](Notifications.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotificationsNotificationsRetrieve

> Notifications NotificationsNotificationsRetrieve(ctx).Execute()





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
    resp, r, err := api_client.NotificationsApi.NotificationsNotificationsRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.NotificationsNotificationsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NotificationsNotificationsRetrieve`: Notifications
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.NotificationsNotificationsRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNotificationsNotificationsRetrieveRequest struct via the builder pattern


### Return type

[**Notifications**](Notifications.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

