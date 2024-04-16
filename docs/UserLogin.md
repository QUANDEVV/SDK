# UserLogin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**TenantId** | [**Tenant**](Tenant.md) |  | [readonly] 
**Username** | **string** |  | [readonly] 
**FirstName** | **string** |  | [readonly] 
**LastName** | **string** |  | [readonly] 
**Email** | **string** |  | 
**Password** | **string** |  | 
**Token** | **string** |  | [readonly] 
**UserType** | [**UserTypeFc6Enum**](UserTypeFc6Enum.md) |  | [readonly] 

## Methods

### NewUserLogin

`func NewUserLogin(id string, tenantId Tenant, username string, firstName string, lastName string, email string, password string, token string, userType UserTypeFc6Enum, ) *UserLogin`

NewUserLogin instantiates a new UserLogin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserLoginWithDefaults

`func NewUserLoginWithDefaults() *UserLogin`

NewUserLoginWithDefaults instantiates a new UserLogin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserLogin) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserLogin) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserLogin) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *UserLogin) GetTenantId() Tenant`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *UserLogin) GetTenantIdOk() (*Tenant, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *UserLogin) SetTenantId(v Tenant)`

SetTenantId sets TenantId field to given value.


### SetTenantIdNil

`func (o *UserLogin) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *UserLogin) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetUsername

`func (o *UserLogin) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserLogin) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserLogin) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetFirstName

`func (o *UserLogin) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserLogin) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserLogin) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *UserLogin) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserLogin) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserLogin) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetEmail

`func (o *UserLogin) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserLogin) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserLogin) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *UserLogin) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserLogin) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserLogin) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetToken

`func (o *UserLogin) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UserLogin) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UserLogin) SetToken(v string)`

SetToken sets Token field to given value.


### GetUserType

`func (o *UserLogin) GetUserType() UserTypeFc6Enum`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *UserLogin) GetUserTypeOk() (*UserTypeFc6Enum, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *UserLogin) SetUserType(v UserTypeFc6Enum)`

SetUserType sets UserType field to given value.


### SetUserTypeNil

`func (o *UserLogin) SetUserTypeNil(b bool)`

 SetUserTypeNil sets the value for UserType to be an explicit nil

### UnsetUserType
`func (o *UserLogin) UnsetUserType()`

UnsetUserType ensures that no value is present for UserType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


