# UnverifiedUsers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UserType** | Pointer to [**OneOfUserType3daEnumNullEnum**](oneOf&lt;UserType3daEnum,NullEnum&gt;.md) |  | [optional] 

## Methods

### NewUnverifiedUsers

`func NewUnverifiedUsers(email string, ) *UnverifiedUsers`

NewUnverifiedUsers instantiates a new UnverifiedUsers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnverifiedUsersWithDefaults

`func NewUnverifiedUsersWithDefaults() *UnverifiedUsers`

NewUnverifiedUsersWithDefaults instantiates a new UnverifiedUsers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UnverifiedUsers) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnverifiedUsers) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnverifiedUsers) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetCreatedAt

`func (o *UnverifiedUsers) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UnverifiedUsers) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UnverifiedUsers) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UnverifiedUsers) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUserType

`func (o *UnverifiedUsers) GetUserType() OneOfUserType3daEnumNullEnum`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *UnverifiedUsers) GetUserTypeOk() (*OneOfUserType3daEnumNullEnum, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *UnverifiedUsers) SetUserType(v OneOfUserType3daEnumNullEnum)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *UnverifiedUsers) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### SetUserTypeNil

`func (o *UnverifiedUsers) SetUserTypeNil(b bool)`

 SetUserTypeNil sets the value for UserType to be an explicit nil

### UnsetUserType
`func (o *UnverifiedUsers) UnsetUserType()`

UnsetUserType ensures that no value is present for UserType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


