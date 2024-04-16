# InviteEmployees

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**TenantId** | **int32** | The tenant chats ID. | 
**UserType** | Pointer to [**OneOfUserType3daEnumNullEnum**](oneOf&lt;UserType3daEnum,NullEnum&gt;.md) |  | [optional] 

## Methods

### NewInviteEmployees

`func NewInviteEmployees(email string, tenantId int32, ) *InviteEmployees`

NewInviteEmployees instantiates a new InviteEmployees object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInviteEmployeesWithDefaults

`func NewInviteEmployeesWithDefaults() *InviteEmployees`

NewInviteEmployeesWithDefaults instantiates a new InviteEmployees object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *InviteEmployees) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InviteEmployees) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InviteEmployees) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetTenantId

`func (o *InviteEmployees) GetTenantId() int32`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *InviteEmployees) GetTenantIdOk() (*int32, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *InviteEmployees) SetTenantId(v int32)`

SetTenantId sets TenantId field to given value.


### GetUserType

`func (o *InviteEmployees) GetUserType() OneOfUserType3daEnumNullEnum`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *InviteEmployees) GetUserTypeOk() (*OneOfUserType3daEnumNullEnum, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *InviteEmployees) SetUserType(v OneOfUserType3daEnumNullEnum)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *InviteEmployees) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### SetUserTypeNil

`func (o *InviteEmployees) SetUserTypeNil(b bool)`

 SetUserTypeNil sets the value for UserType to be an explicit nil

### UnsetUserType
`func (o *InviteEmployees) UnsetUserType()`

UnsetUserType ensures that no value is present for UserType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


