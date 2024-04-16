# Employee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Username** | **string** |  | 
**Email** | **string** |  | 
**FirstName** | **string** |  | 
**LastName** | **string** |  | 
**Phonenumber** | Pointer to **NullableString** |  | [optional] 
**Gender** | Pointer to [**OneOfGenderEnumBlankEnumNullEnum**](oneOf&lt;GenderEnum,BlankEnum,NullEnum&gt;.md) | Enter your gender.  * &#x60;Male&#x60; - Male * &#x60;Female&#x60; - Female * &#x60;Other&#x60; - Other | [optional] 
**DOB** | Pointer to **NullableString** |  | [optional] 
**UserType** | Pointer to [**OneOfUserType8ceEnumNullEnum**](oneOf&lt;UserType8ceEnum,NullEnum&gt;.md) |  | [optional] 
**TenantId** | **int32** | Display name of the tenant | 
**Password** | **string** |  | 
**ConfirmPassword** | **string** |  | 
**Token** | **string** |  | [readonly] 

## Methods

### NewEmployee

`func NewEmployee(id int32, username string, email string, firstName string, lastName string, tenantId int32, password string, confirmPassword string, token string, ) *Employee`

NewEmployee instantiates a new Employee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployeeWithDefaults

`func NewEmployeeWithDefaults() *Employee`

NewEmployeeWithDefaults instantiates a new Employee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Employee) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Employee) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Employee) SetId(v int32)`

SetId sets Id field to given value.


### GetUsername

`func (o *Employee) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Employee) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Employee) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetEmail

`func (o *Employee) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Employee) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Employee) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFirstName

`func (o *Employee) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Employee) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Employee) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *Employee) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Employee) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Employee) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetPhonenumber

`func (o *Employee) GetPhonenumber() string`

GetPhonenumber returns the Phonenumber field if non-nil, zero value otherwise.

### GetPhonenumberOk

`func (o *Employee) GetPhonenumberOk() (*string, bool)`

GetPhonenumberOk returns a tuple with the Phonenumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhonenumber

`func (o *Employee) SetPhonenumber(v string)`

SetPhonenumber sets Phonenumber field to given value.

### HasPhonenumber

`func (o *Employee) HasPhonenumber() bool`

HasPhonenumber returns a boolean if a field has been set.

### SetPhonenumberNil

`func (o *Employee) SetPhonenumberNil(b bool)`

 SetPhonenumberNil sets the value for Phonenumber to be an explicit nil

### UnsetPhonenumber
`func (o *Employee) UnsetPhonenumber()`

UnsetPhonenumber ensures that no value is present for Phonenumber, not even an explicit nil
### GetGender

`func (o *Employee) GetGender() OneOfGenderEnumBlankEnumNullEnum`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *Employee) GetGenderOk() (*OneOfGenderEnumBlankEnumNullEnum, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *Employee) SetGender(v OneOfGenderEnumBlankEnumNullEnum)`

SetGender sets Gender field to given value.

### HasGender

`func (o *Employee) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *Employee) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *Employee) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetDOB

`func (o *Employee) GetDOB() string`

GetDOB returns the DOB field if non-nil, zero value otherwise.

### GetDOBOk

`func (o *Employee) GetDOBOk() (*string, bool)`

GetDOBOk returns a tuple with the DOB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDOB

`func (o *Employee) SetDOB(v string)`

SetDOB sets DOB field to given value.

### HasDOB

`func (o *Employee) HasDOB() bool`

HasDOB returns a boolean if a field has been set.

### SetDOBNil

`func (o *Employee) SetDOBNil(b bool)`

 SetDOBNil sets the value for DOB to be an explicit nil

### UnsetDOB
`func (o *Employee) UnsetDOB()`

UnsetDOB ensures that no value is present for DOB, not even an explicit nil
### GetUserType

`func (o *Employee) GetUserType() OneOfUserType8ceEnumNullEnum`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *Employee) GetUserTypeOk() (*OneOfUserType8ceEnumNullEnum, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *Employee) SetUserType(v OneOfUserType8ceEnumNullEnum)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *Employee) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### SetUserTypeNil

`func (o *Employee) SetUserTypeNil(b bool)`

 SetUserTypeNil sets the value for UserType to be an explicit nil

### UnsetUserType
`func (o *Employee) UnsetUserType()`

UnsetUserType ensures that no value is present for UserType, not even an explicit nil
### GetTenantId

`func (o *Employee) GetTenantId() int32`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Employee) GetTenantIdOk() (*int32, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Employee) SetTenantId(v int32)`

SetTenantId sets TenantId field to given value.


### GetPassword

`func (o *Employee) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Employee) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Employee) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetConfirmPassword

`func (o *Employee) GetConfirmPassword() string`

GetConfirmPassword returns the ConfirmPassword field if non-nil, zero value otherwise.

### GetConfirmPasswordOk

`func (o *Employee) GetConfirmPasswordOk() (*string, bool)`

GetConfirmPasswordOk returns a tuple with the ConfirmPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmPassword

`func (o *Employee) SetConfirmPassword(v string)`

SetConfirmPassword sets ConfirmPassword field to given value.


### GetToken

`func (o *Employee) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *Employee) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *Employee) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


