# \AuthApi

All URIs are relative to *https://core.proximaai.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthAdminCreate**](AuthApi.md#AuthAdminCreate) | **Post** /api/auth/admin/ | 
[**AuthAdminverificationconfirmationRetrieve**](AuthApi.md#AuthAdminverificationconfirmationRetrieve) | **Get** /api/auth/adminverificationconfirmation/{uidb64}/ | 
[**AuthAdminverificationrequestCreate**](AuthApi.md#AuthAdminverificationrequestCreate) | **Post** /api/auth/adminverificationrequest/ | 
[**AuthAnonymoususerCreate**](AuthApi.md#AuthAnonymoususerCreate) | **Post** /api/auth/anonymoususer/ | 
[**AuthAnonymoususerDestroy**](AuthApi.md#AuthAnonymoususerDestroy) | **Delete** /api/auth/anonymoususer/ | 
[**AuthAnonymoususerRetrieve**](AuthApi.md#AuthAnonymoususerRetrieve) | **Get** /api/auth/anonymoususer/ | 
[**AuthAuthGoogleCreate**](AuthApi.md#AuthAuthGoogleCreate) | **Post** /api/auth/auth/google/ | 
[**AuthClientCreate**](AuthApi.md#AuthClientCreate) | **Post** /api/auth/client/ | 
[**AuthClientUsersSignupSigninCreate**](AuthApi.md#AuthClientUsersSignupSigninCreate) | **Post** /api/auth/client_users_signup_signin/ | 
[**AuthClientverificationconfirmationRetrieve**](AuthApi.md#AuthClientverificationconfirmationRetrieve) | **Get** /api/auth/clientverificationconfirmation/{uidb64}/ | 
[**AuthClientverificationrequestCreate**](AuthApi.md#AuthClientverificationrequestCreate) | **Post** /api/auth/clientverificationrequest/ | 
[**AuthEmployeeCreate**](AuthApi.md#AuthEmployeeCreate) | **Post** /api/auth/employee/ | 
[**AuthEmployeeverificationconfirmationRetrieve**](AuthApi.md#AuthEmployeeverificationconfirmationRetrieve) | **Get** /api/auth/employeeverificationconfirmation/{uidb64}/ | 
[**AuthEmployeeverificationrequestCreate**](AuthApi.md#AuthEmployeeverificationrequestCreate) | **Post** /api/auth/employeeverificationrequest/ | 
[**AuthInviteusersCreate**](AuthApi.md#AuthInviteusersCreate) | **Post** /api/auth/inviteusers/ | 
[**AuthPasswordResetCompletePartialUpdate**](AuthApi.md#AuthPasswordResetCompletePartialUpdate) | **Patch** /api/auth/password-reset-complete/ | 
[**AuthPasswordResetRetrieve**](AuthApi.md#AuthPasswordResetRetrieve) | **Get** /api/auth/password-reset/{uidb64}/{token}/ | 
[**AuthRequestlogintokenCreate**](AuthApi.md#AuthRequestlogintokenCreate) | **Post** /api/auth/requestlogintoken/ | 
[**AuthRequestpasswordtokenCreate**](AuthApi.md#AuthRequestpasswordtokenCreate) | **Post** /api/auth/requestpasswordtoken/ | 
[**AuthRequestresetpassowrdCreate**](AuthApi.md#AuthRequestresetpassowrdCreate) | **Post** /api/auth/requestresetpassowrd/ | 
[**AuthSigninCreate**](AuthApi.md#AuthSigninCreate) | **Post** /api/auth/signin/ | 
[**AuthTokenloginCreate**](AuthApi.md#AuthTokenloginCreate) | **Post** /api/auth/tokenlogin/ | 
[**AuthTokenpasswordchangeCreate**](AuthApi.md#AuthTokenpasswordchangeCreate) | **Post** /api/auth/tokenpasswordchange/ | 
[**AuthTokensigninPartialUpdate**](AuthApi.md#AuthTokensigninPartialUpdate) | **Patch** /api/auth/tokensignin/ | 
[**AuthTokensigninconfirmationCreate**](AuthApi.md#AuthTokensigninconfirmationCreate) | **Post** /api/auth/tokensigninconfirmation/{uidb64}/ | 
[**AuthTokensigninrequestRetrieve**](AuthApi.md#AuthTokensigninrequestRetrieve) | **Get** /api/auth/tokensigninrequest/ | 
[**AuthTokenverificationCreate**](AuthApi.md#AuthTokenverificationCreate) | **Post** /api/auth/tokenverification/ | 
[**AuthUnverifiedusersList**](AuthApi.md#AuthUnverifiedusersList) | **Get** /api/auth/unverifiedusers/ | 
[**AuthVerifyinvitedusersCreate**](AuthApi.md#AuthVerifyinvitedusersCreate) | **Post** /api/auth/verifyinvitedusers/ | 



## AuthAdminCreate

> Admin AuthAdminCreate(ctx).Admin(admin).Execute()





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
    admin := *openapiclient.NewAdmin("Username_example", "Email_example", "FirstName_example", "LastName_example", int32(123), "Password_example", "ConfirmPassword_example", "Token_example") // Admin | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthApi.AuthAdminCreate(context.Background()).Admin(admin).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.AuthAdminCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthAdminCreate`: Admin
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.AuthAdminCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthAdminCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **admin** | [**Admin**](Admin.md) |  | 

### Return type

[**Admin**](Admin.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthAdminverificationconfirmationRetrieve

> SetNewPassword AuthAdminverificationconfirmationRetrieve(ctx, uidb64).Execute()



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
    uidb64 := "uidb64_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthApi.AuthAdminverificationconfirmationRetrieve(context.Background(), uidb64).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.AuthAdminverificationconfirmationRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthAdminverificationconfirmationRetrieve`: SetNewPassword
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.AuthAdminverificationconfirmationRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uidb64** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthAdminverificationconfirmationRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SetNewPassword**](SetNewPassword.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthAdminverificationrequestCreate

> ResetPasswordEmailRequest AuthAdminverificationrequestCreate(ctx).ResetPasswordEmailRequest(resetPasswordEmailRequest).Execute()



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
    resetPasswordEmailRequest := *openapiclient.NewResetPasswordEmailRequest("Email_example") // ResetPasswordEmailRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthApi.AuthAdminverificationrequestCreate(context.Background()).ResetPasswordEmailRequest(resetPasswordEmailRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.AuthAdminverificationrequestCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthAdminverificationrequestCreate`: ResetPasswordEmailRequest
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.AuthAdminverificationrequestCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthAdminverificationrequestCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resetPasswordEmailRequest** | [**ResetPasswordEmailRequest**](ResetPasswordEmailRequest.md) |  | 

### Return type

[**ResetPasswordEmailRequest**](ResetPasswordEmailRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthAnonymoususerCreate

> AnonymousUser AuthAnonymoususerCreate(ctx).AnonymousUser(anonymousUser).Execute()





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
    anonymousUser := *openapiclient.NewAnonymousUser(int32(123), "Token_example") // AnonymousUser |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthApi.AuthAnonymoususerCreate(context.Background()).AnonymousUser(anonymousUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.AuthAnonymoususerCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthAnonymoususerCreate`: AnonymousUser
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.AuthAnonymoususerCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthAnonymoususerCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **anonymousUser** | [**AnonymousUser**](AnonymousUser.md) |  | 

### Return type

[**AnonymousUser**](AnonymousUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthAnonymoususerDestroy

> AuthAnonymoususerDestroy(ctx).Execute()





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
    resp, r, err := api_client.AuthApi.AuthAnonymoususerDestroy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.AuthAnonymoususerDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAuthAnonymoususerDestroyRequest struct via the builder pattern


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


## AuthAnonymoususerRetrieve

> AnonymousUser AuthAnonymoususerRetrieve(ctx).Execute()





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
    resp, r, err := api_client.AuthApi.AuthAnonymoususerRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.AuthAnonymoususerRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthAnonymoususerRetrieve`: AnonymousUser
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.AuthAnonymoususerRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAuthAnonymoususerRetrieveRequest struct via the builder pattern


### Return type

[**AnonymousUser**](AnonymousUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthAuthGoogleCreate

> UserLogin AuthAuthGoogleCreate(ctx).UserLogin(userLogin).Execute()



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
    userLogin := *openapiclient.NewUserLogin("Id_example", *openapiclient.NewTenant(int32(123), "TenantName_example", openapiclient.IndustryEnum("IT")), "Username_example", "FirstName_example", "LastName_example", "Email_example", "Password_example", "Token_example", openapiclient.UserTypeFc6Enum("Admin")) // UserLogin | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthApi.AuthAuthGoogleCreate(context.Background()).UserLogin(userLogin).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.AuthAuthGoogleCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthAuthGoogleCreate`: UserLogin
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.AuthAuthGoogleCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthAuthGoogleCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userLogin** | [**UserLogin**](UserLogin.md) |  | 

### Return type

[**UserLogin**](UserLogin.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthClientCreate

> Client AuthClientCreate(ctx).Client(client).Execute()





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
    client := *openapiclient.NewClient(int32(123), "Username_example", "Email_example", "FirstName_example", "LastName_example", "Password_example", "ConfirmPassword_example", "Token_example") // Client | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthApi.AuthClientCreate(context.Background()).Client(client).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.AuthClientCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthClientCreate`: Client
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.AuthClientCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthClientCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **client** | [**Client**](Client.md) |  | 

### Return type

[**Client**](Client.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthClientUsersSignupSigninCreate

> UserLogin AuthClientUsersSignupSigninCreate(ctx).UserLogin(userLogin).Execute()



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
    userLogin := *openapiclient.NewUserLogin("Id_example", *openapiclient.NewTenant(int32(123), "TenantName_example", openapiclient.IndustryEnum("IT")), "Username_example", "FirstName_example", "LastName_example", "Email_example", "Password_example", "Token_example", openapiclient.UserTypeFc6Enum("Admin")) // UserLogin | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthApi.AuthClientUsersSignupSigninCreate(context.Background()).UserLogin(userLogin).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.AuthClientUsersSignupSigninCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthClientUsersSignupSigninCreate`: UserLogin
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.AuthClientUsersSignupSigninCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthClientUsersSignupSigninCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userLogin** | [**UserLogin**](UserLogin.md) |  | 

### Return type

[**UserLogin**](UserLogin.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthClientverificationconfirmationRetrieve

> SetNewPassword AuthClientverificationconfirmationRetrieve(ctx, uidb64).Execute()



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
    uidb64 := "uidb64_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthApi.AuthClientverificationconfirmationRetrieve(context.Background(), uidb64).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.AuthClientverificationconfirmationRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthClientverificationconfirmationRetrieve`: SetNewPassword
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.AuthClientverificationconfirmationRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uidb64** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthClientverificationconfirmationRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SetNewPassword**](SetNewPassword.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthClientverificationrequestCreate

> ResetPasswordEmailRequest AuthClientverificationrequestCreate(ctx).ResetPasswordEmailRequest(resetPasswordEmailRequest).Execute()



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
    resetPasswordEmailRequest := *openapiclient.NewResetPasswordEmailRequest("Email_example") // ResetPasswordEmailRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthApi.AuthClientverificationrequestCreate(context.Background()).ResetPasswordEmailRequest(resetPasswordEmailRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.AuthClientverificationrequestCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthClientverificationrequestCreate`: ResetPasswordEmailRequest
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.AuthClientverificationrequestCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthClientverificationrequestCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resetPasswordEmailRequest** | [**ResetPasswordEmailRequest**](ResetPasswordEmailRequest.md) |  | 

### Return type

[**ResetPasswordEmailRequest**](ResetPasswordEmailRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthEmployeeCreate

> Employee AuthEmployeeCreate(ctx).Employee(employee).Execute()





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
    employee := *openapiclient.NewEmployee(int32(123), "Username_example", "Email_example", "FirstName_example", "LastName_example", int32(123), "Password_example", "ConfirmPassword_example", "Token_example") // Employee | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthApi.AuthEmployeeCreate(context.Background()).Employee(employee).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.AuthEmployeeCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthEmployeeCreate`: Employee
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.AuthEmployeeCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthEmployeeCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **employee** | [**Employee**](Employee.md) |  | 

### Return type

[**Employee**](Employee.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthEmployeeverificationconfirmationRetrieve

> SetNewPassword AuthEmployeeverificationconfirmationRetrieve(ctx, uidb64).Execute()



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
    uidb64 := "uidb64_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthApi.AuthEmployeeverificationconfirmationRetrieve(context.Background(), uidb64).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.AuthEmployeeverificationconfirmationRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthEmployeeverificationconfirmationRetrieve`: SetNewPassword
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.AuthEmployeeverificationconfirmationRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uidb64** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthEmployeeverificationconfirmationRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SetNewPassword**](SetNewPassword.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthEmployeeverificationrequestCreate

> ResetPasswordEmailRequest AuthEmployeeverificationrequestCreate(ctx).ResetPasswordEmailRequest(resetPasswordEmailRequest).Execute()



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
    resetPasswordEmailRequest := *openapiclient.NewResetPasswordEmailRequest("Email_example") // ResetPasswordEmailRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthApi.AuthEmployeeverificationrequestCreate(context.Background()).ResetPasswordEmailRequest(resetPasswordEmailRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.AuthEmployeeverificationrequestCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthEmployeeverificationrequestCreate`: ResetPasswordEmailRequest
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.AuthEmployeeverificationrequestCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthEmployeeverificationrequestCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resetPasswordEmailRequest** | [**ResetPasswordEmailRequest**](ResetPasswordEmailRequest.md) |  | 

### Return type

[**ResetPasswordEmailRequest**](ResetPasswordEmailRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthInviteusersCreate

> InviteEmployees AuthInviteusersCreate(ctx).InviteEmployees(inviteEmployees).Execute()



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
    inviteEmployees := *openapiclient.NewInviteEmployees("Email_example", int32(123)) // InviteEmployees | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthApi.AuthInviteusersCreate(context.Background()).InviteEmployees(inviteEmployees).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.AuthInviteusersCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthInviteusersCreate`: InviteEmployees
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.AuthInviteusersCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthInviteusersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inviteEmployees** | [**InviteEmployees**](InviteEmployees.md) |  | 

### Return type

[**InviteEmployees**](InviteEmployees.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthPasswordResetCompletePartialUpdate

> SetNewPassword AuthPasswordResetCompletePartialUpdate(ctx).PatchedSetNewPassword(patchedSetNewPassword).Execute()



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
    patchedSetNewPassword := *openapiclient.NewPatchedSetNewPassword() // PatchedSetNewPassword |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthApi.AuthPasswordResetCompletePartialUpdate(context.Background()).PatchedSetNewPassword(patchedSetNewPassword).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.AuthPasswordResetCompletePartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthPasswordResetCompletePartialUpdate`: SetNewPassword
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.AuthPasswordResetCompletePartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthPasswordResetCompletePartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedSetNewPassword** | [**PatchedSetNewPassword**](PatchedSetNewPassword.md) |  | 

### Return type

[**SetNewPassword**](SetNewPassword.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthPasswordResetRetrieve

> SetNewPassword AuthPasswordResetRetrieve(ctx, token, uidb64).Execute()



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
    token := "token_example" // string | 
    uidb64 := "uidb64_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthApi.AuthPasswordResetRetrieve(context.Background(), token, uidb64).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.AuthPasswordResetRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthPasswordResetRetrieve`: SetNewPassword
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.AuthPasswordResetRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 
**uidb64** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthPasswordResetRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SetNewPassword**](SetNewPassword.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthRequestlogintokenCreate

> ResetPasswordEmailRequest AuthRequestlogintokenCreate(ctx).ResetPasswordEmailRequest(resetPasswordEmailRequest).Execute()





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
    resetPasswordEmailRequest := *openapiclient.NewResetPasswordEmailRequest("Email_example") // ResetPasswordEmailRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthApi.AuthRequestlogintokenCreate(context.Background()).ResetPasswordEmailRequest(resetPasswordEmailRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.AuthRequestlogintokenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthRequestlogintokenCreate`: ResetPasswordEmailRequest
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.AuthRequestlogintokenCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthRequestlogintokenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resetPasswordEmailRequest** | [**ResetPasswordEmailRequest**](ResetPasswordEmailRequest.md) |  | 

### Return type

[**ResetPasswordEmailRequest**](ResetPasswordEmailRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthRequestpasswordtokenCreate

> ResetPasswordEmailRequest AuthRequestpasswordtokenCreate(ctx).ResetPasswordEmailRequest(resetPasswordEmailRequest).Execute()





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
    resetPasswordEmailRequest := *openapiclient.NewResetPasswordEmailRequest("Email_example") // ResetPasswordEmailRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthApi.AuthRequestpasswordtokenCreate(context.Background()).ResetPasswordEmailRequest(resetPasswordEmailRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.AuthRequestpasswordtokenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthRequestpasswordtokenCreate`: ResetPasswordEmailRequest
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.AuthRequestpasswordtokenCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthRequestpasswordtokenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resetPasswordEmailRequest** | [**ResetPasswordEmailRequest**](ResetPasswordEmailRequest.md) |  | 

### Return type

[**ResetPasswordEmailRequest**](ResetPasswordEmailRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthRequestresetpassowrdCreate

> ResetPasswordEmailRequest AuthRequestresetpassowrdCreate(ctx).ResetPasswordEmailRequest(resetPasswordEmailRequest).Execute()





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
    resetPasswordEmailRequest := *openapiclient.NewResetPasswordEmailRequest("Email_example") // ResetPasswordEmailRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthApi.AuthRequestresetpassowrdCreate(context.Background()).ResetPasswordEmailRequest(resetPasswordEmailRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.AuthRequestresetpassowrdCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthRequestresetpassowrdCreate`: ResetPasswordEmailRequest
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.AuthRequestresetpassowrdCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthRequestresetpassowrdCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resetPasswordEmailRequest** | [**ResetPasswordEmailRequest**](ResetPasswordEmailRequest.md) |  | 

### Return type

[**ResetPasswordEmailRequest**](ResetPasswordEmailRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthSigninCreate

> UserLogin AuthSigninCreate(ctx).UserLogin(userLogin).Execute()



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
    userLogin := *openapiclient.NewUserLogin("Id_example", *openapiclient.NewTenant(int32(123), "TenantName_example", openapiclient.IndustryEnum("IT")), "Username_example", "FirstName_example", "LastName_example", "Email_example", "Password_example", "Token_example", openapiclient.UserTypeFc6Enum("Admin")) // UserLogin | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthApi.AuthSigninCreate(context.Background()).UserLogin(userLogin).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.AuthSigninCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthSigninCreate`: UserLogin
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.AuthSigninCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthSigninCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userLogin** | [**UserLogin**](UserLogin.md) |  | 

### Return type

[**UserLogin**](UserLogin.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthTokenloginCreate

> TokenLogin AuthTokenloginCreate(ctx).TokenLogin(tokenLogin).Execute()





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
    tokenLogin := *openapiclient.NewTokenLogin(int32(123), "Id_example", *openapiclient.NewTenant(int32(123), "TenantName_example", openapiclient.IndustryEnum("IT")), "Username_example", "FirstName_example", "LastName_example", "Email_example", openapiclient.UserTypeFc6Enum("Admin")) // TokenLogin | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthApi.AuthTokenloginCreate(context.Background()).TokenLogin(tokenLogin).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.AuthTokenloginCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthTokenloginCreate`: TokenLogin
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.AuthTokenloginCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthTokenloginCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenLogin** | [**TokenLogin**](TokenLogin.md) |  | 

### Return type

[**TokenLogin**](TokenLogin.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthTokenpasswordchangeCreate

> TokenSetNewPassword AuthTokenpasswordchangeCreate(ctx).TokenSetNewPassword(tokenSetNewPassword).Execute()





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
    tokenSetNewPassword := *openapiclient.NewTokenSetNewPassword("Password_example", int32(123)) // TokenSetNewPassword | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthApi.AuthTokenpasswordchangeCreate(context.Background()).TokenSetNewPassword(tokenSetNewPassword).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.AuthTokenpasswordchangeCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthTokenpasswordchangeCreate`: TokenSetNewPassword
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.AuthTokenpasswordchangeCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthTokenpasswordchangeCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenSetNewPassword** | [**TokenSetNewPassword**](TokenSetNewPassword.md) |  | 

### Return type

[**TokenSetNewPassword**](TokenSetNewPassword.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthTokensigninPartialUpdate

> SetNewPassword AuthTokensigninPartialUpdate(ctx).PatchedSetNewPassword(patchedSetNewPassword).Execute()



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
    patchedSetNewPassword := *openapiclient.NewPatchedSetNewPassword() // PatchedSetNewPassword |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthApi.AuthTokensigninPartialUpdate(context.Background()).PatchedSetNewPassword(patchedSetNewPassword).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.AuthTokensigninPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthTokensigninPartialUpdate`: SetNewPassword
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.AuthTokensigninPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthTokensigninPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedSetNewPassword** | [**PatchedSetNewPassword**](PatchedSetNewPassword.md) |  | 

### Return type

[**SetNewPassword**](SetNewPassword.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthTokensigninconfirmationCreate

> ResetPasswordEmailRequest AuthTokensigninconfirmationCreate(ctx, uidb64).ResetPasswordEmailRequest(resetPasswordEmailRequest).Execute()



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
    uidb64 := "uidb64_example" // string | 
    resetPasswordEmailRequest := *openapiclient.NewResetPasswordEmailRequest("Email_example") // ResetPasswordEmailRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthApi.AuthTokensigninconfirmationCreate(context.Background(), uidb64).ResetPasswordEmailRequest(resetPasswordEmailRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.AuthTokensigninconfirmationCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthTokensigninconfirmationCreate`: ResetPasswordEmailRequest
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.AuthTokensigninconfirmationCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uidb64** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthTokensigninconfirmationCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resetPasswordEmailRequest** | [**ResetPasswordEmailRequest**](ResetPasswordEmailRequest.md) |  | 

### Return type

[**ResetPasswordEmailRequest**](ResetPasswordEmailRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthTokensigninrequestRetrieve

> SetNewPassword AuthTokensigninrequestRetrieve(ctx).Execute()



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
    resp, r, err := api_client.AuthApi.AuthTokensigninrequestRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.AuthTokensigninrequestRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthTokensigninrequestRetrieve`: SetNewPassword
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.AuthTokensigninrequestRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAuthTokensigninrequestRetrieveRequest struct via the builder pattern


### Return type

[**SetNewPassword**](SetNewPassword.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthTokenverificationCreate

> EmailVerification AuthTokenverificationCreate(ctx).EmailVerification(emailVerification).Execute()





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
    emailVerification := *openapiclient.NewEmailVerification(int32(123)) // EmailVerification | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthApi.AuthTokenverificationCreate(context.Background()).EmailVerification(emailVerification).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.AuthTokenverificationCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthTokenverificationCreate`: EmailVerification
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.AuthTokenverificationCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthTokenverificationCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailVerification** | [**EmailVerification**](EmailVerification.md) |  | 

### Return type

[**EmailVerification**](EmailVerification.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthUnverifiedusersList

> PaginatedUnverifiedUsersList AuthUnverifiedusersList(ctx).Limit(limit).Offset(offset).Execute()



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
    resp, r, err := api_client.AuthApi.AuthUnverifiedusersList(context.Background()).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.AuthUnverifiedusersList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthUnverifiedusersList`: PaginatedUnverifiedUsersList
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.AuthUnverifiedusersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthUnverifiedusersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedUnverifiedUsersList**](PaginatedUnverifiedUsersList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthVerifyinvitedusersCreate

> VerifyUsersToken AuthVerifyinvitedusersCreate(ctx).VerifyUsersToken(verifyUsersToken).Execute()





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
    verifyUsersToken := *openapiclient.NewVerifyUsersToken(int32(123)) // VerifyUsersToken | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthApi.AuthVerifyinvitedusersCreate(context.Background()).VerifyUsersToken(verifyUsersToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.AuthVerifyinvitedusersCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthVerifyinvitedusersCreate`: VerifyUsersToken
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.AuthVerifyinvitedusersCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthVerifyinvitedusersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **verifyUsersToken** | [**VerifyUsersToken**](VerifyUsersToken.md) |  | 

### Return type

[**VerifyUsersToken**](VerifyUsersToken.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

