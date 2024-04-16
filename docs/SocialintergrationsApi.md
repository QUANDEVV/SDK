# \SocialintergrationsApi

All URIs are relative to *https://core.proximaai.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SocialintergrationsFacebookintergrationCreate**](SocialintergrationsApi.md#SocialintergrationsFacebookintergrationCreate) | **Post** /api/socialintergrations/facebookintergration/ | 
[**SocialintergrationsFacebookintergrationRetrieve**](SocialintergrationsApi.md#SocialintergrationsFacebookintergrationRetrieve) | **Get** /api/socialintergrations/facebookintergration/ | 
[**SocialintergrationsInstagramintergrationCreate**](SocialintergrationsApi.md#SocialintergrationsInstagramintergrationCreate) | **Post** /api/socialintergrations/instagramintergration/ | 
[**SocialintergrationsInstagramintergrationRetrieve**](SocialintergrationsApi.md#SocialintergrationsInstagramintergrationRetrieve) | **Get** /api/socialintergrations/instagramintergration/ | 
[**SocialintergrationsProximawhatsappintergrationCreate**](SocialintergrationsApi.md#SocialintergrationsProximawhatsappintergrationCreate) | **Post** /api/socialintergrations/proximawhatsappintergration/ | 
[**SocialintergrationsProximawhatsappintergrationRetrieve**](SocialintergrationsApi.md#SocialintergrationsProximawhatsappintergrationRetrieve) | **Get** /api/socialintergrations/proximawhatsappintergration/ | 
[**SocialintergrationsSavefacebookintergrationCreate**](SocialintergrationsApi.md#SocialintergrationsSavefacebookintergrationCreate) | **Post** /api/socialintergrations/savefacebookintergration/ | 
[**SocialintergrationsSavefacebookintergrationRetrieve**](SocialintergrationsApi.md#SocialintergrationsSavefacebookintergrationRetrieve) | **Get** /api/socialintergrations/savefacebookintergration/ | 
[**SocialintergrationsSavefacebookintergrationUpdate**](SocialintergrationsApi.md#SocialintergrationsSavefacebookintergrationUpdate) | **Put** /api/socialintergrations/savefacebookintergration/ | 
[**SocialintergrationsSaveinstagramintergrationCreate**](SocialintergrationsApi.md#SocialintergrationsSaveinstagramintergrationCreate) | **Post** /api/socialintergrations/saveinstagramintergration/ | 
[**SocialintergrationsSaveinstagramintergrationRetrieve**](SocialintergrationsApi.md#SocialintergrationsSaveinstagramintergrationRetrieve) | **Get** /api/socialintergrations/saveinstagramintergration/ | 
[**SocialintergrationsSaveinstagramintergrationUpdate**](SocialintergrationsApi.md#SocialintergrationsSaveinstagramintergrationUpdate) | **Put** /api/socialintergrations/saveinstagramintergration/ | 
[**SocialintergrationsSavewhatsappintergrationCreate**](SocialintergrationsApi.md#SocialintergrationsSavewhatsappintergrationCreate) | **Post** /api/socialintergrations/savewhatsappintergration/ | 
[**SocialintergrationsSavewhatsappintergrationRetrieve**](SocialintergrationsApi.md#SocialintergrationsSavewhatsappintergrationRetrieve) | **Get** /api/socialintergrations/savewhatsappintergration/ | 
[**SocialintergrationsSavewhatsappintergrationUpdate**](SocialintergrationsApi.md#SocialintergrationsSavewhatsappintergrationUpdate) | **Put** /api/socialintergrations/savewhatsappintergration/ | 
[**SocialintergrationsWhatsappintergrationCreate**](SocialintergrationsApi.md#SocialintergrationsWhatsappintergrationCreate) | **Post** /api/socialintergrations/whatsappintergration/ | 
[**SocialintergrationsWhatsappintergrationRetrieve**](SocialintergrationsApi.md#SocialintergrationsWhatsappintergrationRetrieve) | **Get** /api/socialintergrations/whatsappintergration/ | 



## SocialintergrationsFacebookintergrationCreate

> SocialintergrationsFacebookintergrationCreate(ctx).Execute()



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
    resp, r, err := api_client.SocialintergrationsApi.SocialintergrationsFacebookintergrationCreate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SocialintergrationsApi.SocialintergrationsFacebookintergrationCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSocialintergrationsFacebookintergrationCreateRequest struct via the builder pattern


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


## SocialintergrationsFacebookintergrationRetrieve

> SocialintergrationsFacebookintergrationRetrieve(ctx).Execute()



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
    resp, r, err := api_client.SocialintergrationsApi.SocialintergrationsFacebookintergrationRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SocialintergrationsApi.SocialintergrationsFacebookintergrationRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSocialintergrationsFacebookintergrationRetrieveRequest struct via the builder pattern


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


## SocialintergrationsInstagramintergrationCreate

> SocialintergrationsInstagramintergrationCreate(ctx).Execute()



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
    resp, r, err := api_client.SocialintergrationsApi.SocialintergrationsInstagramintergrationCreate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SocialintergrationsApi.SocialintergrationsInstagramintergrationCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSocialintergrationsInstagramintergrationCreateRequest struct via the builder pattern


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


## SocialintergrationsInstagramintergrationRetrieve

> SocialintergrationsInstagramintergrationRetrieve(ctx).Execute()



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
    resp, r, err := api_client.SocialintergrationsApi.SocialintergrationsInstagramintergrationRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SocialintergrationsApi.SocialintergrationsInstagramintergrationRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSocialintergrationsInstagramintergrationRetrieveRequest struct via the builder pattern


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


## SocialintergrationsProximawhatsappintergrationCreate

> SocialintergrationsProximawhatsappintergrationCreate(ctx).Execute()



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
    resp, r, err := api_client.SocialintergrationsApi.SocialintergrationsProximawhatsappintergrationCreate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SocialintergrationsApi.SocialintergrationsProximawhatsappintergrationCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSocialintergrationsProximawhatsappintergrationCreateRequest struct via the builder pattern


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


## SocialintergrationsProximawhatsappintergrationRetrieve

> SocialintergrationsProximawhatsappintergrationRetrieve(ctx).Execute()



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
    resp, r, err := api_client.SocialintergrationsApi.SocialintergrationsProximawhatsappintergrationRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SocialintergrationsApi.SocialintergrationsProximawhatsappintergrationRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSocialintergrationsProximawhatsappintergrationRetrieveRequest struct via the builder pattern


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


## SocialintergrationsSavefacebookintergrationCreate

> SocialintergrationsSavefacebookintergrationCreate(ctx).Execute()





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
    resp, r, err := api_client.SocialintergrationsApi.SocialintergrationsSavefacebookintergrationCreate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SocialintergrationsApi.SocialintergrationsSavefacebookintergrationCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSocialintergrationsSavefacebookintergrationCreateRequest struct via the builder pattern


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


## SocialintergrationsSavefacebookintergrationRetrieve

> SocialintergrationsSavefacebookintergrationRetrieve(ctx).Execute()





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
    resp, r, err := api_client.SocialintergrationsApi.SocialintergrationsSavefacebookintergrationRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SocialintergrationsApi.SocialintergrationsSavefacebookintergrationRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSocialintergrationsSavefacebookintergrationRetrieveRequest struct via the builder pattern


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


## SocialintergrationsSavefacebookintergrationUpdate

> SocialintergrationsSavefacebookintergrationUpdate(ctx).Execute()





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
    resp, r, err := api_client.SocialintergrationsApi.SocialintergrationsSavefacebookintergrationUpdate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SocialintergrationsApi.SocialintergrationsSavefacebookintergrationUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSocialintergrationsSavefacebookintergrationUpdateRequest struct via the builder pattern


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


## SocialintergrationsSaveinstagramintergrationCreate

> SocialintergrationsSaveinstagramintergrationCreate(ctx).Execute()





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
    resp, r, err := api_client.SocialintergrationsApi.SocialintergrationsSaveinstagramintergrationCreate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SocialintergrationsApi.SocialintergrationsSaveinstagramintergrationCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSocialintergrationsSaveinstagramintergrationCreateRequest struct via the builder pattern


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


## SocialintergrationsSaveinstagramintergrationRetrieve

> SocialintergrationsSaveinstagramintergrationRetrieve(ctx).Execute()





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
    resp, r, err := api_client.SocialintergrationsApi.SocialintergrationsSaveinstagramintergrationRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SocialintergrationsApi.SocialintergrationsSaveinstagramintergrationRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSocialintergrationsSaveinstagramintergrationRetrieveRequest struct via the builder pattern


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


## SocialintergrationsSaveinstagramintergrationUpdate

> SocialintergrationsSaveinstagramintergrationUpdate(ctx).Execute()





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
    resp, r, err := api_client.SocialintergrationsApi.SocialintergrationsSaveinstagramintergrationUpdate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SocialintergrationsApi.SocialintergrationsSaveinstagramintergrationUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSocialintergrationsSaveinstagramintergrationUpdateRequest struct via the builder pattern


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


## SocialintergrationsSavewhatsappintergrationCreate

> SocialintergrationsSavewhatsappintergrationCreate(ctx).Execute()





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
    resp, r, err := api_client.SocialintergrationsApi.SocialintergrationsSavewhatsappintergrationCreate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SocialintergrationsApi.SocialintergrationsSavewhatsappintergrationCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSocialintergrationsSavewhatsappintergrationCreateRequest struct via the builder pattern


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


## SocialintergrationsSavewhatsappintergrationRetrieve

> SocialintergrationsSavewhatsappintergrationRetrieve(ctx).Execute()





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
    resp, r, err := api_client.SocialintergrationsApi.SocialintergrationsSavewhatsappintergrationRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SocialintergrationsApi.SocialintergrationsSavewhatsappintergrationRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSocialintergrationsSavewhatsappintergrationRetrieveRequest struct via the builder pattern


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


## SocialintergrationsSavewhatsappintergrationUpdate

> SocialintergrationsSavewhatsappintergrationUpdate(ctx).Execute()





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
    resp, r, err := api_client.SocialintergrationsApi.SocialintergrationsSavewhatsappintergrationUpdate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SocialintergrationsApi.SocialintergrationsSavewhatsappintergrationUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSocialintergrationsSavewhatsappintergrationUpdateRequest struct via the builder pattern


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


## SocialintergrationsWhatsappintergrationCreate

> SocialintergrationsWhatsappintergrationCreate(ctx).Execute()



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
    resp, r, err := api_client.SocialintergrationsApi.SocialintergrationsWhatsappintergrationCreate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SocialintergrationsApi.SocialintergrationsWhatsappintergrationCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSocialintergrationsWhatsappintergrationCreateRequest struct via the builder pattern


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


## SocialintergrationsWhatsappintergrationRetrieve

> SocialintergrationsWhatsappintergrationRetrieve(ctx).Execute()



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
    resp, r, err := api_client.SocialintergrationsApi.SocialintergrationsWhatsappintergrationRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SocialintergrationsApi.SocialintergrationsWhatsappintergrationRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSocialintergrationsWhatsappintergrationRetrieveRequest struct via the builder pattern


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

