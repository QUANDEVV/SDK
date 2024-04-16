# \PaymentsApi

All URIs are relative to *https://core.proximaai.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PaymentsFreetrialCreate**](PaymentsApi.md#PaymentsFreetrialCreate) | **Post** /api/payments/freetrial/ | 
[**PaymentsGetpaymentsRetrieve**](PaymentsApi.md#PaymentsGetpaymentsRetrieve) | **Get** /api/payments/getpayments/ | 
[**PaymentsSubscriptionRetrieve**](PaymentsApi.md#PaymentsSubscriptionRetrieve) | **Get** /api/payments/subscription/ | 
[**PaymentsTabularpaymentsRetrieve**](PaymentsApi.md#PaymentsTabularpaymentsRetrieve) | **Get** /api/payments/tabularpayments/ | 
[**PaymentsTiersList**](PaymentsApi.md#PaymentsTiersList) | **Get** /api/payments/tiers/ | 
[**PaymentsUpdateSubscriptionCreate**](PaymentsApi.md#PaymentsUpdateSubscriptionCreate) | **Post** /api/payments/update_subscription/{id}/ | 



## PaymentsFreetrialCreate

> FreeTrialEnroll PaymentsFreetrialCreate(ctx).FreeTrialEnroll(freeTrialEnroll).Execute()



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
    freeTrialEnroll := *openapiclient.NewFreeTrialEnroll("Email_example", int32(123)) // FreeTrialEnroll | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PaymentsApi.PaymentsFreetrialCreate(context.Background()).FreeTrialEnroll(freeTrialEnroll).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsApi.PaymentsFreetrialCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaymentsFreetrialCreate`: FreeTrialEnroll
    fmt.Fprintf(os.Stdout, "Response from `PaymentsApi.PaymentsFreetrialCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentsFreetrialCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **freeTrialEnroll** | [**FreeTrialEnroll**](FreeTrialEnroll.md) |  | 

### Return type

[**FreeTrialEnroll**](FreeTrialEnroll.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentsGetpaymentsRetrieve

> Payment PaymentsGetpaymentsRetrieve(ctx).Execute()





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
    resp, r, err := api_client.PaymentsApi.PaymentsGetpaymentsRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsApi.PaymentsGetpaymentsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaymentsGetpaymentsRetrieve`: Payment
    fmt.Fprintf(os.Stdout, "Response from `PaymentsApi.PaymentsGetpaymentsRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentsGetpaymentsRetrieveRequest struct via the builder pattern


### Return type

[**Payment**](Payment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentsSubscriptionRetrieve

> Subscription PaymentsSubscriptionRetrieve(ctx).Execute()





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
    resp, r, err := api_client.PaymentsApi.PaymentsSubscriptionRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsApi.PaymentsSubscriptionRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaymentsSubscriptionRetrieve`: Subscription
    fmt.Fprintf(os.Stdout, "Response from `PaymentsApi.PaymentsSubscriptionRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentsSubscriptionRetrieveRequest struct via the builder pattern


### Return type

[**Subscription**](Subscription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentsTabularpaymentsRetrieve

> PaymentsTabularpaymentsRetrieve(ctx).Execute()



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
    resp, r, err := api_client.PaymentsApi.PaymentsTabularpaymentsRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsApi.PaymentsTabularpaymentsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentsTabularpaymentsRetrieveRequest struct via the builder pattern


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


## PaymentsTiersList

> PaginatedTierList PaymentsTiersList(ctx).Limit(limit).Offset(offset).Execute()



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
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PaymentsApi.PaymentsTiersList(context.Background()).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsApi.PaymentsTiersList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaymentsTiersList`: PaginatedTierList
    fmt.Fprintf(os.Stdout, "Response from `PaymentsApi.PaymentsTiersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentsTiersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedTierList**](PaginatedTierList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentsUpdateSubscriptionCreate

> Subscription PaymentsUpdateSubscriptionCreate(ctx, id).Subscription(subscription).Execute()



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
    id := "id_example" // string | 
    subscription := *openapiclient.NewSubscription(int32(123), *openapiclient.NewPaymentsTenant(int32(123), "TenantName_example", openapiclient.IndustryEnum("IT")), *openapiclient.NewTier("Name_example", []openapiclient.Feature{*openapiclient.NewFeature("Name_example", "Description_example")}, []openapiclient.Price{*openapiclient.NewPrice(*openapiclient.NewQuota("Name_example", int32(123)), "Amount_example", *openapiclient.NewCompanyType(openapiclient.CompanyTypeTypeEnum("small")))}), *openapiclient.NewQuota("Name_example", int32(123)), "CreatedAt_example", time.Now(), time.Now(), time.Now()) // Subscription |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PaymentsApi.PaymentsUpdateSubscriptionCreate(context.Background(), id).Subscription(subscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsApi.PaymentsUpdateSubscriptionCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaymentsUpdateSubscriptionCreate`: Subscription
    fmt.Fprintf(os.Stdout, "Response from `PaymentsApi.PaymentsUpdateSubscriptionCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentsUpdateSubscriptionCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscription** | [**Subscription**](Subscription.md) |  | 

### Return type

[**Subscription**](Subscription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

