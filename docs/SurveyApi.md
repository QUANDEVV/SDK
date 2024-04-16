# \SurveyApi

All URIs are relative to *https://core.proximaai.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SurveyEndsurveyCreate**](SurveyApi.md#SurveyEndsurveyCreate) | **Post** /api/survey/endsurvey/ | 
[**SurveyPromptsurveyCreate**](SurveyApi.md#SurveyPromptsurveyCreate) | **Post** /api/survey/promptsurvey/ | 
[**SurveyPromptsurveyRetrieve**](SurveyApi.md#SurveyPromptsurveyRetrieve) | **Get** /api/survey/promptsurvey/ | 
[**SurveyQuestioninducedCreate**](SurveyApi.md#SurveyQuestioninducedCreate) | **Post** /api/survey/questioninduced/ | 
[**SurveyQuestioninducedRetrieve**](SurveyApi.md#SurveyQuestioninducedRetrieve) | **Get** /api/survey/questioninduced/ | 
[**SurveyResponseCreate**](SurveyApi.md#SurveyResponseCreate) | **Post** /api/survey/response/ | 
[**SurveyResponseRetrieve**](SurveyApi.md#SurveyResponseRetrieve) | **Get** /api/survey/response/ | 
[**SurveySurveyCreate**](SurveyApi.md#SurveySurveyCreate) | **Post** /api/survey/survey/ | 
[**SurveySurveyRetrieve**](SurveyApi.md#SurveySurveyRetrieve) | **Get** /api/survey/survey/ | 
[**SurveySurveyreportviewCreate**](SurveyApi.md#SurveySurveyreportviewCreate) | **Post** /api/survey/surveyreportview/ | 
[**SurveySurveyreportviewRetrieve**](SurveyApi.md#SurveySurveyreportviewRetrieve) | **Get** /api/survey/surveyreportview/ | 
[**SurveySurveysubgroupCreate**](SurveyApi.md#SurveySurveysubgroupCreate) | **Post** /api/survey/surveysubgroup/ | 
[**SurveySurveysubgroupRetrieve**](SurveyApi.md#SurveySurveysubgroupRetrieve) | **Get** /api/survey/surveysubgroup/ | 



## SurveyEndsurveyCreate

> SurveyEndsurveyCreate(ctx).Execute()





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
    resp, r, err := api_client.SurveyApi.SurveyEndsurveyCreate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SurveyApi.SurveyEndsurveyCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSurveyEndsurveyCreateRequest struct via the builder pattern


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


## SurveyPromptsurveyCreate

> PromptSurvey SurveyPromptsurveyCreate(ctx).PromptSurvey(promptSurvey).Execute()



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
    promptSurvey := *openapiclient.NewPromptSurvey(int32(123), int32(123), "SurveyTopic_example", "SurveyDescription_example", "SurveyPrompt_example", *openapiclient.NewClient(int32(123), "Username_example", "Email_example", "FirstName_example", "LastName_example", "Password_example", "ConfirmPassword_example", "Token_example"), *openapiclient.NewAnonymousUser(int32(123), "Token_example")) // PromptSurvey | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SurveyApi.SurveyPromptsurveyCreate(context.Background()).PromptSurvey(promptSurvey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SurveyApi.SurveyPromptsurveyCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SurveyPromptsurveyCreate`: PromptSurvey
    fmt.Fprintf(os.Stdout, "Response from `SurveyApi.SurveyPromptsurveyCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSurveyPromptsurveyCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **promptSurvey** | [**PromptSurvey**](PromptSurvey.md) |  | 

### Return type

[**PromptSurvey**](PromptSurvey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SurveyPromptsurveyRetrieve

> PromptSurvey SurveyPromptsurveyRetrieve(ctx).Execute()





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
    resp, r, err := api_client.SurveyApi.SurveyPromptsurveyRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SurveyApi.SurveyPromptsurveyRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SurveyPromptsurveyRetrieve`: PromptSurvey
    fmt.Fprintf(os.Stdout, "Response from `SurveyApi.SurveyPromptsurveyRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSurveyPromptsurveyRetrieveRequest struct via the builder pattern


### Return type

[**PromptSurvey**](PromptSurvey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SurveyQuestioninducedCreate

> QuestionInducedSurvey SurveyQuestioninducedCreate(ctx).QuestionInducedSurvey(questionInducedSurvey).Execute()



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
    questionInducedSurvey := *openapiclient.NewQuestionInducedSurvey(int32(123), int32(123), "SurveyTopic_example", "SurveyDescription_example", *openapiclient.NewClient(int32(123), "Username_example", "Email_example", "FirstName_example", "LastName_example", "Password_example", "ConfirmPassword_example", "Token_example"), *openapiclient.NewAnonymousUser(int32(123), "Token_example")) // QuestionInducedSurvey | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SurveyApi.SurveyQuestioninducedCreate(context.Background()).QuestionInducedSurvey(questionInducedSurvey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SurveyApi.SurveyQuestioninducedCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SurveyQuestioninducedCreate`: QuestionInducedSurvey
    fmt.Fprintf(os.Stdout, "Response from `SurveyApi.SurveyQuestioninducedCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSurveyQuestioninducedCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **questionInducedSurvey** | [**QuestionInducedSurvey**](QuestionInducedSurvey.md) |  | 

### Return type

[**QuestionInducedSurvey**](QuestionInducedSurvey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SurveyQuestioninducedRetrieve

> QuestionInducedSurvey SurveyQuestioninducedRetrieve(ctx).Execute()





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
    resp, r, err := api_client.SurveyApi.SurveyQuestioninducedRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SurveyApi.SurveyQuestioninducedRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SurveyQuestioninducedRetrieve`: QuestionInducedSurvey
    fmt.Fprintf(os.Stdout, "Response from `SurveyApi.SurveyQuestioninducedRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSurveyQuestioninducedRetrieveRequest struct via the builder pattern


### Return type

[**QuestionInducedSurvey**](QuestionInducedSurvey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SurveyResponseCreate

> Response SurveyResponseCreate(ctx).Response(response).Execute()





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
    response := *openapiclient.NewResponse(int32(123), *openapiclient.NewTargetAudience(int32(123), "Username_example", "Email_example", "FirstName_example", "LastName_example"), time.Now(), "CreatedAt_example") // Response | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SurveyApi.SurveyResponseCreate(context.Background()).Response(response).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SurveyApi.SurveyResponseCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SurveyResponseCreate`: Response
    fmt.Fprintf(os.Stdout, "Response from `SurveyApi.SurveyResponseCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSurveyResponseCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **response** | [**Response**](Response.md) |  | 

### Return type

[**Response**](Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SurveyResponseRetrieve

> Response SurveyResponseRetrieve(ctx).Execute()





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
    resp, r, err := api_client.SurveyApi.SurveyResponseRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SurveyApi.SurveyResponseRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SurveyResponseRetrieve`: Response
    fmt.Fprintf(os.Stdout, "Response from `SurveyApi.SurveyResponseRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSurveyResponseRetrieveRequest struct via the builder pattern


### Return type

[**Response**](Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SurveySurveyCreate

> Survey SurveySurveyCreate(ctx).Survey(survey).Execute()





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
    survey := *openapiclient.NewSurvey(int32(123), int32(123), "SurveyTopic_example", "SurveyDescription_example", "SurveyContext_example", time.Now(), int32(123), int32(123), float64(123), int32(123)) // Survey | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SurveyApi.SurveySurveyCreate(context.Background()).Survey(survey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SurveyApi.SurveySurveyCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SurveySurveyCreate`: Survey
    fmt.Fprintf(os.Stdout, "Response from `SurveyApi.SurveySurveyCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSurveySurveyCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **survey** | [**Survey**](Survey.md) |  | 

### Return type

[**Survey**](Survey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SurveySurveyRetrieve

> Survey SurveySurveyRetrieve(ctx).Execute()





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
    resp, r, err := api_client.SurveyApi.SurveySurveyRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SurveyApi.SurveySurveyRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SurveySurveyRetrieve`: Survey
    fmt.Fprintf(os.Stdout, "Response from `SurveyApi.SurveySurveyRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSurveySurveyRetrieveRequest struct via the builder pattern


### Return type

[**Survey**](Survey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SurveySurveyreportviewCreate

> SurveyReport SurveySurveyreportviewCreate(ctx).SurveyReport(surveyReport).Execute()





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
    surveyReport := *openapiclient.NewSurveyReport(int32(123), int32(123), "Conclusion_example") // SurveyReport | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SurveyApi.SurveySurveyreportviewCreate(context.Background()).SurveyReport(surveyReport).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SurveyApi.SurveySurveyreportviewCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SurveySurveyreportviewCreate`: SurveyReport
    fmt.Fprintf(os.Stdout, "Response from `SurveyApi.SurveySurveyreportviewCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSurveySurveyreportviewCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **surveyReport** | [**SurveyReport**](SurveyReport.md) |  | 

### Return type

[**SurveyReport**](SurveyReport.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SurveySurveyreportviewRetrieve

> SurveyReport SurveySurveyreportviewRetrieve(ctx).Execute()





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
    resp, r, err := api_client.SurveyApi.SurveySurveyreportviewRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SurveyApi.SurveySurveyreportviewRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SurveySurveyreportviewRetrieve`: SurveyReport
    fmt.Fprintf(os.Stdout, "Response from `SurveyApi.SurveySurveyreportviewRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSurveySurveyreportviewRetrieveRequest struct via the builder pattern


### Return type

[**SurveyReport**](SurveyReport.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SurveySurveysubgroupCreate

> SurveySubGroups SurveySurveysubgroupCreate(ctx).SurveySubGroups(surveySubGroups).Execute()





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
    surveySubGroups := *openapiclient.NewSurveySubGroups(int32(123), "SubgroupName_example", "SubgroupDescription_example", []openapiclient.TargetAudience{*openapiclient.NewTargetAudience(int32(123), "Username_example", "Email_example", "FirstName_example", "LastName_example")}) // SurveySubGroups | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SurveyApi.SurveySurveysubgroupCreate(context.Background()).SurveySubGroups(surveySubGroups).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SurveyApi.SurveySurveysubgroupCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SurveySurveysubgroupCreate`: SurveySubGroups
    fmt.Fprintf(os.Stdout, "Response from `SurveyApi.SurveySurveysubgroupCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSurveySurveysubgroupCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **surveySubGroups** | [**SurveySubGroups**](SurveySubGroups.md) |  | 

### Return type

[**SurveySubGroups**](SurveySubGroups.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SurveySurveysubgroupRetrieve

> SurveySubGroups SurveySurveysubgroupRetrieve(ctx).Execute()





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
    resp, r, err := api_client.SurveyApi.SurveySurveysubgroupRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SurveyApi.SurveySurveysubgroupRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SurveySurveysubgroupRetrieve`: SurveySubGroups
    fmt.Fprintf(os.Stdout, "Response from `SurveyApi.SurveySurveysubgroupRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSurveySurveysubgroupRetrieveRequest struct via the builder pattern


### Return type

[**SurveySubGroups**](SurveySubGroups.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

