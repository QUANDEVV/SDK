# \ChatApi

All URIs are relative to *https://core.proximaai.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChatCategoryCreate**](ChatApi.md#ChatCategoryCreate) | **Post** /api/chat/category/ | 
[**ChatCategoryRetrieve**](ChatApi.md#ChatCategoryRetrieve) | **Get** /api/chat/category/ | 
[**ChatCategoryUpdate**](ChatApi.md#ChatCategoryUpdate) | **Put** /api/chat/category/ | 
[**ChatChatCreate**](ChatApi.md#ChatChatCreate) | **Post** /api/chat/chat/ | 
[**ChatChatDestroy**](ChatApi.md#ChatChatDestroy) | **Delete** /api/chat/chat/ | 
[**ChatChatPartialUpdate**](ChatApi.md#ChatChatPartialUpdate) | **Patch** /api/chat/chat/ | 
[**ChatChatRetrieve**](ChatApi.md#ChatChatRetrieve) | **Get** /api/chat/chat/ | 
[**ChatEscalatedChatsRetrieve**](ChatApi.md#ChatEscalatedChatsRetrieve) | **Get** /api/chat/escalated_chats/ | 
[**ChatMessageCreate**](ChatApi.md#ChatMessageCreate) | **Post** /api/chat/message/ | 
[**ChatMessageDestroy**](ChatApi.md#ChatMessageDestroy) | **Delete** /api/chat/message/ | 
[**ChatMessagePartialUpdate**](ChatApi.md#ChatMessagePartialUpdate) | **Patch** /api/chat/message/ | 
[**ChatMessageRetrieve**](ChatApi.md#ChatMessageRetrieve) | **Get** /api/chat/message/ | 
[**ChatTopicsCreate**](ChatApi.md#ChatTopicsCreate) | **Post** /api/chat/topics/ | 
[**ChatTopicsRetrieve**](ChatApi.md#ChatTopicsRetrieve) | **Get** /api/chat/topics/ | 
[**ChatTopicsUpdate**](ChatApi.md#ChatTopicsUpdate) | **Put** /api/chat/topics/ | 



## ChatCategoryCreate

> Category ChatCategoryCreate(ctx).Category(category).Execute()





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
    category := *openapiclient.NewCategory(int32(123), int32(123), "Name_example", "Description_example") // Category | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatApi.ChatCategoryCreate(context.Background()).Category(category).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatApi.ChatCategoryCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChatCategoryCreate`: Category
    fmt.Fprintf(os.Stdout, "Response from `ChatApi.ChatCategoryCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatCategoryCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **category** | [**Category**](Category.md) |  | 

### Return type

[**Category**](Category.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatCategoryRetrieve

> Category ChatCategoryRetrieve(ctx).Execute()





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
    resp, r, err := api_client.ChatApi.ChatCategoryRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatApi.ChatCategoryRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChatCategoryRetrieve`: Category
    fmt.Fprintf(os.Stdout, "Response from `ChatApi.ChatCategoryRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChatCategoryRetrieveRequest struct via the builder pattern


### Return type

[**Category**](Category.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatCategoryUpdate

> Category ChatCategoryUpdate(ctx).Category(category).Execute()





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
    category := *openapiclient.NewCategory(int32(123), int32(123), "Name_example", "Description_example") // Category | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatApi.ChatCategoryUpdate(context.Background()).Category(category).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatApi.ChatCategoryUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChatCategoryUpdate`: Category
    fmt.Fprintf(os.Stdout, "Response from `ChatApi.ChatCategoryUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatCategoryUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **category** | [**Category**](Category.md) |  | 

### Return type

[**Category**](Category.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatChatCreate

> Chat ChatChatCreate(ctx).Chat(chat).Execute()





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
    chat := *openapiclient.NewChat(int32(123), *openapiclient.NewTenantInfo(int32(123), "TenantName_example"), *openapiclient.NewClientInfo(int32(123), "Username_example", "Email_example", "FirstName_example", "LastName_example"), *openapiclient.NewClientInfo(int32(123), "Username_example", "Email_example", "FirstName_example", "LastName_example"), *openapiclient.NewAnonymousUser(int32(123), "Token_example"), time.Now(), "CreatedAt_example") // Chat | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatApi.ChatChatCreate(context.Background()).Chat(chat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatApi.ChatChatCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChatChatCreate`: Chat
    fmt.Fprintf(os.Stdout, "Response from `ChatApi.ChatChatCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatChatCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chat** | [**Chat**](Chat.md) |  | 

### Return type

[**Chat**](Chat.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatChatDestroy

> ChatChatDestroy(ctx).Execute()





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
    resp, r, err := api_client.ChatApi.ChatChatDestroy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatApi.ChatChatDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChatChatDestroyRequest struct via the builder pattern


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


## ChatChatPartialUpdate

> Chat ChatChatPartialUpdate(ctx).PatchedChat(patchedChat).Execute()





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
    patchedChat := *openapiclient.NewPatchedChat() // PatchedChat |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatApi.ChatChatPartialUpdate(context.Background()).PatchedChat(patchedChat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatApi.ChatChatPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChatChatPartialUpdate`: Chat
    fmt.Fprintf(os.Stdout, "Response from `ChatApi.ChatChatPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatChatPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedChat** | [**PatchedChat**](PatchedChat.md) |  | 

### Return type

[**Chat**](Chat.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatChatRetrieve

> Chat ChatChatRetrieve(ctx).Execute()





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
    resp, r, err := api_client.ChatApi.ChatChatRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatApi.ChatChatRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChatChatRetrieve`: Chat
    fmt.Fprintf(os.Stdout, "Response from `ChatApi.ChatChatRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChatChatRetrieveRequest struct via the builder pattern


### Return type

[**Chat**](Chat.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatEscalatedChatsRetrieve

> Chat ChatEscalatedChatsRetrieve(ctx).Execute()





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
    resp, r, err := api_client.ChatApi.ChatEscalatedChatsRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatApi.ChatEscalatedChatsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChatEscalatedChatsRetrieve`: Chat
    fmt.Fprintf(os.Stdout, "Response from `ChatApi.ChatEscalatedChatsRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChatEscalatedChatsRetrieveRequest struct via the builder pattern


### Return type

[**Chat**](Chat.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatMessageCreate

> Message ChatMessageCreate(ctx).Message(message).Execute()





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
    message := *openapiclient.NewMessage(int32(123), int32(123), openapiclient.MessageSenderEnum("client")) // Message | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatApi.ChatMessageCreate(context.Background()).Message(message).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatApi.ChatMessageCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChatMessageCreate`: Message
    fmt.Fprintf(os.Stdout, "Response from `ChatApi.ChatMessageCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatMessageCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **message** | [**Message**](Message.md) |  | 

### Return type

[**Message**](Message.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatMessageDestroy

> ChatMessageDestroy(ctx).Execute()





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
    resp, r, err := api_client.ChatApi.ChatMessageDestroy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatApi.ChatMessageDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChatMessageDestroyRequest struct via the builder pattern


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


## ChatMessagePartialUpdate

> Message ChatMessagePartialUpdate(ctx).PatchedMessage(patchedMessage).Execute()





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
    patchedMessage := *openapiclient.NewPatchedMessage() // PatchedMessage |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatApi.ChatMessagePartialUpdate(context.Background()).PatchedMessage(patchedMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatApi.ChatMessagePartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChatMessagePartialUpdate`: Message
    fmt.Fprintf(os.Stdout, "Response from `ChatApi.ChatMessagePartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatMessagePartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedMessage** | [**PatchedMessage**](PatchedMessage.md) |  | 

### Return type

[**Message**](Message.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatMessageRetrieve

> Message ChatMessageRetrieve(ctx).Execute()





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
    resp, r, err := api_client.ChatApi.ChatMessageRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatApi.ChatMessageRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChatMessageRetrieve`: Message
    fmt.Fprintf(os.Stdout, "Response from `ChatApi.ChatMessageRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChatMessageRetrieveRequest struct via the builder pattern


### Return type

[**Message**](Message.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatTopicsCreate

> InteractionTopics ChatTopicsCreate(ctx).InteractionTopics(interactionTopics).Execute()





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
    interactionTopics := *openapiclient.NewInteractionTopics(int32(123), "Name_example", "Description_example", int32(123)) // InteractionTopics | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatApi.ChatTopicsCreate(context.Background()).InteractionTopics(interactionTopics).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatApi.ChatTopicsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChatTopicsCreate`: InteractionTopics
    fmt.Fprintf(os.Stdout, "Response from `ChatApi.ChatTopicsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatTopicsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **interactionTopics** | [**InteractionTopics**](InteractionTopics.md) |  | 

### Return type

[**InteractionTopics**](InteractionTopics.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatTopicsRetrieve

> InteractionTopics ChatTopicsRetrieve(ctx).Execute()





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
    resp, r, err := api_client.ChatApi.ChatTopicsRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatApi.ChatTopicsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChatTopicsRetrieve`: InteractionTopics
    fmt.Fprintf(os.Stdout, "Response from `ChatApi.ChatTopicsRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChatTopicsRetrieveRequest struct via the builder pattern


### Return type

[**InteractionTopics**](InteractionTopics.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatTopicsUpdate

> InteractionTopics ChatTopicsUpdate(ctx).InteractionTopics(interactionTopics).Execute()





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
    interactionTopics := *openapiclient.NewInteractionTopics(int32(123), "Name_example", "Description_example", int32(123)) // InteractionTopics | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatApi.ChatTopicsUpdate(context.Background()).InteractionTopics(interactionTopics).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatApi.ChatTopicsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChatTopicsUpdate`: InteractionTopics
    fmt.Fprintf(os.Stdout, "Response from `ChatApi.ChatTopicsUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatTopicsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **interactionTopics** | [**InteractionTopics**](InteractionTopics.md) |  | 

### Return type

[**InteractionTopics**](InteractionTopics.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

