# ResetPasswordEmailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**RedirectUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewResetPasswordEmailRequest

`func NewResetPasswordEmailRequest(email string, ) *ResetPasswordEmailRequest`

NewResetPasswordEmailRequest instantiates a new ResetPasswordEmailRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResetPasswordEmailRequestWithDefaults

`func NewResetPasswordEmailRequestWithDefaults() *ResetPasswordEmailRequest`

NewResetPasswordEmailRequestWithDefaults instantiates a new ResetPasswordEmailRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *ResetPasswordEmailRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ResetPasswordEmailRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ResetPasswordEmailRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetRedirectUrl

`func (o *ResetPasswordEmailRequest) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *ResetPasswordEmailRequest) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *ResetPasswordEmailRequest) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *ResetPasswordEmailRequest) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


