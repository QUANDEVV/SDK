# \TenantmanagementApi

All URIs are relative to *https://core.proximaai.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TenantmanagementAddressCreate**](TenantmanagementApi.md#TenantmanagementAddressCreate) | **Post** /api/tenantmanagement/address/ | 
[**TenantmanagementAddressRetrieve**](TenantmanagementApi.md#TenantmanagementAddressRetrieve) | **Get** /api/tenantmanagement/address/ | 
[**TenantmanagementDepartmentCreate**](TenantmanagementApi.md#TenantmanagementDepartmentCreate) | **Post** /api/tenantmanagement/department | 
[**TenantmanagementDepartmentDestroy**](TenantmanagementApi.md#TenantmanagementDepartmentDestroy) | **Delete** /api/tenantmanagement/department | 
[**TenantmanagementDepartmentPartialUpdate**](TenantmanagementApi.md#TenantmanagementDepartmentPartialUpdate) | **Patch** /api/tenantmanagement/department | 
[**TenantmanagementDepartmentRetrieve**](TenantmanagementApi.md#TenantmanagementDepartmentRetrieve) | **Get** /api/tenantmanagement/department | 
[**TenantmanagementMetadataCreate**](TenantmanagementApi.md#TenantmanagementMetadataCreate) | **Post** /api/tenantmanagement/metadata/ | 
[**TenantmanagementMetadataDestroy**](TenantmanagementApi.md#TenantmanagementMetadataDestroy) | **Delete** /api/tenantmanagement/metadata/ | 
[**TenantmanagementMetadataRetrieve**](TenantmanagementApi.md#TenantmanagementMetadataRetrieve) | **Get** /api/tenantmanagement/metadata/ | 
[**TenantmanagementProductCreate**](TenantmanagementApi.md#TenantmanagementProductCreate) | **Post** /api/tenantmanagement/product/ | 
[**TenantmanagementProductDestroy**](TenantmanagementApi.md#TenantmanagementProductDestroy) | **Delete** /api/tenantmanagement/product/ | 
[**TenantmanagementProductRetrieve**](TenantmanagementApi.md#TenantmanagementProductRetrieve) | **Get** /api/tenantmanagement/product/ | 
[**TenantmanagementTenantCreate**](TenantmanagementApi.md#TenantmanagementTenantCreate) | **Post** /api/tenantmanagement/tenant/ | 
[**TenantmanagementTenantDestroy**](TenantmanagementApi.md#TenantmanagementTenantDestroy) | **Delete** /api/tenantmanagement/tenant/ | 
[**TenantmanagementTenantRetrieve**](TenantmanagementApi.md#TenantmanagementTenantRetrieve) | **Get** /api/tenantmanagement/tenant/ | 
[**TenantmanagementTenantdetailsRetrieve**](TenantmanagementApi.md#TenantmanagementTenantdetailsRetrieve) | **Get** /api/tenantmanagement/tenantdetails/ | 
[**TenantmanagementVirtualAssistantsCreate**](TenantmanagementApi.md#TenantmanagementVirtualAssistantsCreate) | **Post** /api/tenantmanagement/virtual-assistants/{id}/ | 
[**TenantmanagementVirtualAssistantsRetrieve**](TenantmanagementApi.md#TenantmanagementVirtualAssistantsRetrieve) | **Get** /api/tenantmanagement/virtual-assistants/{id}/ | 
[**TenantmanagementVirtualAssistantsUpdate**](TenantmanagementApi.md#TenantmanagementVirtualAssistantsUpdate) | **Put** /api/tenantmanagement/virtual-assistants/{id}/ | 
[**TenantmanagementVirtualassistantCreate**](TenantmanagementApi.md#TenantmanagementVirtualassistantCreate) | **Post** /api/tenantmanagement/virtualassistant/ | 
[**TenantmanagementVirtualassistantRetrieve**](TenantmanagementApi.md#TenantmanagementVirtualassistantRetrieve) | **Get** /api/tenantmanagement/virtualassistant/ | 
[**TenantmanagementVirtualassistantUpdate**](TenantmanagementApi.md#TenantmanagementVirtualassistantUpdate) | **Put** /api/tenantmanagement/virtualassistant/ | 
[**TenantmanagementVirtualassitantdocsCreate**](TenantmanagementApi.md#TenantmanagementVirtualassitantdocsCreate) | **Post** /api/tenantmanagement/virtualassitantdocs/ | 
[**TenantmanagementVirtualassitantdocsRetrieve**](TenantmanagementApi.md#TenantmanagementVirtualassitantdocsRetrieve) | **Get** /api/tenantmanagement/virtualassitantdocs/ | 



## TenantmanagementAddressCreate

> Address TenantmanagementAddressCreate(ctx).Address(address).Execute()





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
    address := *openapiclient.NewAddress(int32(123), int32(123), "City_example", "Country_example", "PostalCode_example", "State_example", "PaymentNumber_example") // Address | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantmanagementApi.TenantmanagementAddressCreate(context.Background()).Address(address).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantmanagementApi.TenantmanagementAddressCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantmanagementAddressCreate`: Address
    fmt.Fprintf(os.Stdout, "Response from `TenantmanagementApi.TenantmanagementAddressCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenantmanagementAddressCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **address** | [**Address**](Address.md) |  | 

### Return type

[**Address**](Address.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantmanagementAddressRetrieve

> Address TenantmanagementAddressRetrieve(ctx).Execute()





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
    resp, r, err := api_client.TenantmanagementApi.TenantmanagementAddressRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantmanagementApi.TenantmanagementAddressRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantmanagementAddressRetrieve`: Address
    fmt.Fprintf(os.Stdout, "Response from `TenantmanagementApi.TenantmanagementAddressRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTenantmanagementAddressRetrieveRequest struct via the builder pattern


### Return type

[**Address**](Address.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantmanagementDepartmentCreate

> Department TenantmanagementDepartmentCreate(ctx).Department(department).Execute()





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
    department := *openapiclient.NewDepartment(int32(123), int32(123), "Name_example", []int32{int32(123)}, []int32{int32(123)}) // Department | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantmanagementApi.TenantmanagementDepartmentCreate(context.Background()).Department(department).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantmanagementApi.TenantmanagementDepartmentCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantmanagementDepartmentCreate`: Department
    fmt.Fprintf(os.Stdout, "Response from `TenantmanagementApi.TenantmanagementDepartmentCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenantmanagementDepartmentCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **department** | [**Department**](Department.md) |  | 

### Return type

[**Department**](Department.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantmanagementDepartmentDestroy

> TenantmanagementDepartmentDestroy(ctx).Execute()





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
    resp, r, err := api_client.TenantmanagementApi.TenantmanagementDepartmentDestroy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantmanagementApi.TenantmanagementDepartmentDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTenantmanagementDepartmentDestroyRequest struct via the builder pattern


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


## TenantmanagementDepartmentPartialUpdate

> Department TenantmanagementDepartmentPartialUpdate(ctx).PatchedDepartment(patchedDepartment).Execute()





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
    patchedDepartment := *openapiclient.NewPatchedDepartment() // PatchedDepartment |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantmanagementApi.TenantmanagementDepartmentPartialUpdate(context.Background()).PatchedDepartment(patchedDepartment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantmanagementApi.TenantmanagementDepartmentPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantmanagementDepartmentPartialUpdate`: Department
    fmt.Fprintf(os.Stdout, "Response from `TenantmanagementApi.TenantmanagementDepartmentPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenantmanagementDepartmentPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedDepartment** | [**PatchedDepartment**](PatchedDepartment.md) |  | 

### Return type

[**Department**](Department.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantmanagementDepartmentRetrieve

> Department TenantmanagementDepartmentRetrieve(ctx).Execute()





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
    resp, r, err := api_client.TenantmanagementApi.TenantmanagementDepartmentRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantmanagementApi.TenantmanagementDepartmentRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantmanagementDepartmentRetrieve`: Department
    fmt.Fprintf(os.Stdout, "Response from `TenantmanagementApi.TenantmanagementDepartmentRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTenantmanagementDepartmentRetrieveRequest struct via the builder pattern


### Return type

[**Department**](Department.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantmanagementMetadataCreate

> Metadata TenantmanagementMetadataCreate(ctx).Metadata(metadata).Execute()





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
    metadata := *openapiclient.NewMetadata(int32(123), interface{}(123)) // Metadata | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantmanagementApi.TenantmanagementMetadataCreate(context.Background()).Metadata(metadata).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantmanagementApi.TenantmanagementMetadataCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantmanagementMetadataCreate`: Metadata
    fmt.Fprintf(os.Stdout, "Response from `TenantmanagementApi.TenantmanagementMetadataCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenantmanagementMetadataCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metadata** | [**Metadata**](Metadata.md) |  | 

### Return type

[**Metadata**](Metadata.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantmanagementMetadataDestroy

> TenantmanagementMetadataDestroy(ctx).Execute()





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
    resp, r, err := api_client.TenantmanagementApi.TenantmanagementMetadataDestroy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantmanagementApi.TenantmanagementMetadataDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTenantmanagementMetadataDestroyRequest struct via the builder pattern


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


## TenantmanagementMetadataRetrieve

> Metadata TenantmanagementMetadataRetrieve(ctx).Execute()





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
    resp, r, err := api_client.TenantmanagementApi.TenantmanagementMetadataRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantmanagementApi.TenantmanagementMetadataRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantmanagementMetadataRetrieve`: Metadata
    fmt.Fprintf(os.Stdout, "Response from `TenantmanagementApi.TenantmanagementMetadataRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTenantmanagementMetadataRetrieveRequest struct via the builder pattern


### Return type

[**Metadata**](Metadata.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantmanagementProductCreate

> Product TenantmanagementProductCreate(ctx).Product(product).Execute()





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
    product := *openapiclient.NewProduct(int32(123), int32(123), "Name_example", "Description_example", "Price_example") // Product | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantmanagementApi.TenantmanagementProductCreate(context.Background()).Product(product).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantmanagementApi.TenantmanagementProductCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantmanagementProductCreate`: Product
    fmt.Fprintf(os.Stdout, "Response from `TenantmanagementApi.TenantmanagementProductCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenantmanagementProductCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **product** | [**Product**](Product.md) |  | 

### Return type

[**Product**](Product.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantmanagementProductDestroy

> TenantmanagementProductDestroy(ctx).Execute()





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
    resp, r, err := api_client.TenantmanagementApi.TenantmanagementProductDestroy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantmanagementApi.TenantmanagementProductDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTenantmanagementProductDestroyRequest struct via the builder pattern


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


## TenantmanagementProductRetrieve

> Product TenantmanagementProductRetrieve(ctx).Execute()





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
    resp, r, err := api_client.TenantmanagementApi.TenantmanagementProductRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantmanagementApi.TenantmanagementProductRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantmanagementProductRetrieve`: Product
    fmt.Fprintf(os.Stdout, "Response from `TenantmanagementApi.TenantmanagementProductRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTenantmanagementProductRetrieveRequest struct via the builder pattern


### Return type

[**Product**](Product.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantmanagementTenantCreate

> Tenant TenantmanagementTenantCreate(ctx).Tenant(tenant).Execute()





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
    tenant := *openapiclient.NewTenant(int32(123), "TenantName_example", openapiclient.IndustryEnum("IT")) // Tenant | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantmanagementApi.TenantmanagementTenantCreate(context.Background()).Tenant(tenant).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantmanagementApi.TenantmanagementTenantCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantmanagementTenantCreate`: Tenant
    fmt.Fprintf(os.Stdout, "Response from `TenantmanagementApi.TenantmanagementTenantCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenantmanagementTenantCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenant** | [**Tenant**](Tenant.md) |  | 

### Return type

[**Tenant**](Tenant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantmanagementTenantDestroy

> TenantmanagementTenantDestroy(ctx).Execute()





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
    resp, r, err := api_client.TenantmanagementApi.TenantmanagementTenantDestroy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantmanagementApi.TenantmanagementTenantDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTenantmanagementTenantDestroyRequest struct via the builder pattern


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


## TenantmanagementTenantRetrieve

> Tenant TenantmanagementTenantRetrieve(ctx).Execute()





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
    resp, r, err := api_client.TenantmanagementApi.TenantmanagementTenantRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantmanagementApi.TenantmanagementTenantRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantmanagementTenantRetrieve`: Tenant
    fmt.Fprintf(os.Stdout, "Response from `TenantmanagementApi.TenantmanagementTenantRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTenantmanagementTenantRetrieveRequest struct via the builder pattern


### Return type

[**Tenant**](Tenant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantmanagementTenantdetailsRetrieve

> TenantmanagementTenantdetailsRetrieve(ctx).Execute()



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
    resp, r, err := api_client.TenantmanagementApi.TenantmanagementTenantdetailsRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantmanagementApi.TenantmanagementTenantdetailsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTenantmanagementTenantdetailsRetrieveRequest struct via the builder pattern


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


## TenantmanagementVirtualAssistantsCreate

> VirtualAssistant TenantmanagementVirtualAssistantsCreate(ctx, id).VirtualAssistant(virtualAssistant).Execute()



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
    id := int32(56) // int32 | 
    virtualAssistant := *openapiclient.NewVirtualAssistant(int32(123), "CreatedAt_example", time.Now(), time.Now(), time.Now(), "FirstName_example", "Industry_example", "Nickname_example", "Persona_example", "Temprature_example", "Accuracy_example", int32(123)) // VirtualAssistant | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantmanagementApi.TenantmanagementVirtualAssistantsCreate(context.Background(), id).VirtualAssistant(virtualAssistant).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantmanagementApi.TenantmanagementVirtualAssistantsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantmanagementVirtualAssistantsCreate`: VirtualAssistant
    fmt.Fprintf(os.Stdout, "Response from `TenantmanagementApi.TenantmanagementVirtualAssistantsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenantmanagementVirtualAssistantsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **virtualAssistant** | [**VirtualAssistant**](VirtualAssistant.md) |  | 

### Return type

[**VirtualAssistant**](VirtualAssistant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantmanagementVirtualAssistantsRetrieve

> VirtualAssistant TenantmanagementVirtualAssistantsRetrieve(ctx, id).Execute()



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
    id := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantmanagementApi.TenantmanagementVirtualAssistantsRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantmanagementApi.TenantmanagementVirtualAssistantsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantmanagementVirtualAssistantsRetrieve`: VirtualAssistant
    fmt.Fprintf(os.Stdout, "Response from `TenantmanagementApi.TenantmanagementVirtualAssistantsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenantmanagementVirtualAssistantsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VirtualAssistant**](VirtualAssistant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantmanagementVirtualAssistantsUpdate

> VirtualAssistant TenantmanagementVirtualAssistantsUpdate(ctx, id).VirtualAssistant(virtualAssistant).Execute()



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
    id := int32(56) // int32 | 
    virtualAssistant := *openapiclient.NewVirtualAssistant(int32(123), "CreatedAt_example", time.Now(), time.Now(), time.Now(), "FirstName_example", "Industry_example", "Nickname_example", "Persona_example", "Temprature_example", "Accuracy_example", int32(123)) // VirtualAssistant | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantmanagementApi.TenantmanagementVirtualAssistantsUpdate(context.Background(), id).VirtualAssistant(virtualAssistant).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantmanagementApi.TenantmanagementVirtualAssistantsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantmanagementVirtualAssistantsUpdate`: VirtualAssistant
    fmt.Fprintf(os.Stdout, "Response from `TenantmanagementApi.TenantmanagementVirtualAssistantsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenantmanagementVirtualAssistantsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **virtualAssistant** | [**VirtualAssistant**](VirtualAssistant.md) |  | 

### Return type

[**VirtualAssistant**](VirtualAssistant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantmanagementVirtualassistantCreate

> VirtualAssistant TenantmanagementVirtualassistantCreate(ctx).VirtualAssistant(virtualAssistant).Execute()



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
    virtualAssistant := *openapiclient.NewVirtualAssistant(int32(123), "CreatedAt_example", time.Now(), time.Now(), time.Now(), "FirstName_example", "Industry_example", "Nickname_example", "Persona_example", "Temprature_example", "Accuracy_example", int32(123)) // VirtualAssistant | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantmanagementApi.TenantmanagementVirtualassistantCreate(context.Background()).VirtualAssistant(virtualAssistant).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantmanagementApi.TenantmanagementVirtualassistantCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantmanagementVirtualassistantCreate`: VirtualAssistant
    fmt.Fprintf(os.Stdout, "Response from `TenantmanagementApi.TenantmanagementVirtualassistantCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenantmanagementVirtualassistantCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **virtualAssistant** | [**VirtualAssistant**](VirtualAssistant.md) |  | 

### Return type

[**VirtualAssistant**](VirtualAssistant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantmanagementVirtualassistantRetrieve

> VirtualAssistant TenantmanagementVirtualassistantRetrieve(ctx).Execute()



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
    resp, r, err := api_client.TenantmanagementApi.TenantmanagementVirtualassistantRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantmanagementApi.TenantmanagementVirtualassistantRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantmanagementVirtualassistantRetrieve`: VirtualAssistant
    fmt.Fprintf(os.Stdout, "Response from `TenantmanagementApi.TenantmanagementVirtualassistantRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTenantmanagementVirtualassistantRetrieveRequest struct via the builder pattern


### Return type

[**VirtualAssistant**](VirtualAssistant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantmanagementVirtualassistantUpdate

> VirtualAssistant TenantmanagementVirtualassistantUpdate(ctx).VirtualAssistant(virtualAssistant).Execute()



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
    virtualAssistant := *openapiclient.NewVirtualAssistant(int32(123), "CreatedAt_example", time.Now(), time.Now(), time.Now(), "FirstName_example", "Industry_example", "Nickname_example", "Persona_example", "Temprature_example", "Accuracy_example", int32(123)) // VirtualAssistant | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantmanagementApi.TenantmanagementVirtualassistantUpdate(context.Background()).VirtualAssistant(virtualAssistant).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantmanagementApi.TenantmanagementVirtualassistantUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantmanagementVirtualassistantUpdate`: VirtualAssistant
    fmt.Fprintf(os.Stdout, "Response from `TenantmanagementApi.TenantmanagementVirtualassistantUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenantmanagementVirtualassistantUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **virtualAssistant** | [**VirtualAssistant**](VirtualAssistant.md) |  | 

### Return type

[**VirtualAssistant**](VirtualAssistant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantmanagementVirtualassitantdocsCreate

> VirtualAssistantDocuments TenantmanagementVirtualassitantdocsCreate(ctx).VirtualAssistantDocuments(virtualAssistantDocuments).Execute()



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
    virtualAssistantDocuments := *openapiclient.NewVirtualAssistantDocuments(int32(123), "CreatedAt_example", time.Now(), time.Now(), time.Now(), "Pdf_example", int32(123)) // VirtualAssistantDocuments | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantmanagementApi.TenantmanagementVirtualassitantdocsCreate(context.Background()).VirtualAssistantDocuments(virtualAssistantDocuments).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantmanagementApi.TenantmanagementVirtualassitantdocsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantmanagementVirtualassitantdocsCreate`: VirtualAssistantDocuments
    fmt.Fprintf(os.Stdout, "Response from `TenantmanagementApi.TenantmanagementVirtualassitantdocsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenantmanagementVirtualassitantdocsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **virtualAssistantDocuments** | [**VirtualAssistantDocuments**](VirtualAssistantDocuments.md) |  | 

### Return type

[**VirtualAssistantDocuments**](VirtualAssistantDocuments.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantmanagementVirtualassitantdocsRetrieve

> VirtualAssistantDocuments TenantmanagementVirtualassitantdocsRetrieve(ctx).Execute()



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
    resp, r, err := api_client.TenantmanagementApi.TenantmanagementVirtualassitantdocsRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantmanagementApi.TenantmanagementVirtualassitantdocsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantmanagementVirtualassitantdocsRetrieve`: VirtualAssistantDocuments
    fmt.Fprintf(os.Stdout, "Response from `TenantmanagementApi.TenantmanagementVirtualassitantdocsRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTenantmanagementVirtualassitantdocsRetrieveRequest struct via the builder pattern


### Return type

[**VirtualAssistantDocuments**](VirtualAssistantDocuments.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

