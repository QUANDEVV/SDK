# TokenLogin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **int32** |  | 
**Id** | **string** |  | [readonly] 
**TenantId** | [**Tenant**](Tenant.md) |  | [readonly] 
**Username** | **string** |  | [readonly] 
**FirstName** | **string** |  | [readonly] 
**LastName** | **string** |  | [readonly] 
**Email** | **string** |  | [readonly] 
**UserType** | [**UserTypeFc6Enum**](UserTypeFc6Enum.md) |  | [readonly] 

## Methods

### NewTokenLogin

`func NewTokenLogin(token int32, id string, tenantId Tenant, username string, firstName string, lastName string, email string, userType UserTypeFc6Enum, ) *TokenLogin`

NewTokenLogin instantiates a new TokenLogin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenLoginWithDefaults

`func NewTokenLoginWithDefaults() *TokenLogin`

NewTokenLoginWithDefaults instantiates a new TokenLogin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *TokenLogin) GetToken() int32`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TokenLogin) GetTokenOk() (*int32, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TokenLogin) SetToken(v int32)`

SetToken sets Token field to given value.


### GetId

`func (o *TokenLogin) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TokenLogin) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TokenLogin) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *TokenLogin) GetTenantId() Tenant`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TokenLogin) GetTenantIdOk() (*Tenant, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TokenLogin) SetTenantId(v Tenant)`

SetTenantId sets TenantId field to given value.


### SetTenantIdNil

`func (o *TokenLogin) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *TokenLogin) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetUsername

`func (o *TokenLogin) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *TokenLogin) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *TokenLogin) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetFirstName

`func (o *TokenLogin) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *TokenLogin) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *TokenLogin) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *TokenLogin) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *TokenLogin) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *TokenLogin) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetEmail

`func (o *TokenLogin) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *TokenLogin) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *TokenLogin) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetUserType

`func (o *TokenLogin) GetUserType() UserTypeFc6Enum`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *TokenLogin) GetUserTypeOk() (*UserTypeFc6Enum, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *TokenLogin) SetUserType(v UserTypeFc6Enum)`

SetUserType sets UserType field to given value.


### SetUserTypeNil

`func (o *TokenLogin) SetUserTypeNil(b bool)`

 SetUserTypeNil sets the value for UserType to be an explicit nil

### UnsetUserType
`func (o *TokenLogin) UnsetUserType()`

UnsetUserType ensures that no value is present for UserType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


