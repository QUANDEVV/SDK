# \SchemaApi

All URIs are relative to *https://core.proximaai.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SchemaRetrieve**](SchemaApi.md#SchemaRetrieve) | **Get** /api/schema/ | 



## SchemaRetrieve

> map[string]interface{} SchemaRetrieve(ctx).Format(format).Lang(lang).Execute()





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
    format := "format_example" // string |  (optional)
    lang := "lang_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SchemaApi.SchemaRetrieve(context.Background()).Format(format).Lang(lang).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemaApi.SchemaRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SchemaRetrieve`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SchemaApi.SchemaRetrieve`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSchemaRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string** |  | 
 **lang** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.oai.openapi, application/yaml, application/vnd.oai.openapi+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

