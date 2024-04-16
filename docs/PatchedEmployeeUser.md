# PatchedEmployeeUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Phonenumber** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Gender** | Pointer to [**OneOfGenderEnumBlankEnumNullEnum**](oneOf&lt;GenderEnum,BlankEnum,NullEnum&gt;.md) | Enter your gender.  * &#x60;Male&#x60; - Male * &#x60;Female&#x60; - Female * &#x60;Other&#x60; - Other | [optional] 
**DOB** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] [readonly] 
**Profile** | Pointer to [**EmployeeProfile**](EmployeeProfile.md) |  | [optional] 
**ProfilePhoto** | Pointer to **string** |  | [optional] [readonly] 
**Country** | Pointer to **string** |  | [optional] [readonly] 
**County** | Pointer to **string** |  | [optional] [readonly] 
**City** | Pointer to **string** |  | [optional] [readonly] 
**PostalCode** | Pointer to **string** |  | [optional] [readonly] 
**Location** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewPatchedEmployeeUser

`func NewPatchedEmployeeUser() *PatchedEmployeeUser`

NewPatchedEmployeeUser instantiates a new PatchedEmployeeUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedEmployeeUserWithDefaults

`func NewPatchedEmployeeUserWithDefaults() *PatchedEmployeeUser`

NewPatchedEmployeeUserWithDefaults instantiates a new PatchedEmployeeUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedEmployeeUser) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedEmployeeUser) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedEmployeeUser) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedEmployeeUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFirstName

`func (o *PatchedEmployeeUser) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *PatchedEmployeeUser) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *PatchedEmployeeUser) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *PatchedEmployeeUser) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *PatchedEmployeeUser) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *PatchedEmployeeUser) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *PatchedEmployeeUser) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *PatchedEmployeeUser) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetUsername

`func (o *PatchedEmployeeUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PatchedEmployeeUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PatchedEmployeeUser) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *PatchedEmployeeUser) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPhonenumber

`func (o *PatchedEmployeeUser) GetPhonenumber() string`

GetPhonenumber returns the Phonenumber field if non-nil, zero value otherwise.

### GetPhonenumberOk

`func (o *PatchedEmployeeUser) GetPhonenumberOk() (*string, bool)`

GetPhonenumberOk returns a tuple with the Phonenumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhonenumber

`func (o *PatchedEmployeeUser) SetPhonenumber(v string)`

SetPhonenumber sets Phonenumber field to given value.

### HasPhonenumber

`func (o *PatchedEmployeeUser) HasPhonenumber() bool`

HasPhonenumber returns a boolean if a field has been set.

### SetPhonenumberNil

`func (o *PatchedEmployeeUser) SetPhonenumberNil(b bool)`

 SetPhonenumberNil sets the value for Phonenumber to be an explicit nil

### UnsetPhonenumber
`func (o *PatchedEmployeeUser) UnsetPhonenumber()`

UnsetPhonenumber ensures that no value is present for Phonenumber, not even an explicit nil
### GetEmail

`func (o *PatchedEmployeeUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PatchedEmployeeUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PatchedEmployeeUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PatchedEmployeeUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetGender

`func (o *PatchedEmployeeUser) GetGender() OneOfGenderEnumBlankEnumNullEnum`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *PatchedEmployeeUser) GetGenderOk() (*OneOfGenderEnumBlankEnumNullEnum, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *PatchedEmployeeUser) SetGender(v OneOfGenderEnumBlankEnumNullEnum)`

SetGender sets Gender field to given value.

### HasGender

`func (o *PatchedEmployeeUser) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *PatchedEmployeeUser) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *PatchedEmployeeUser) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetDOB

`func (o *PatchedEmployeeUser) GetDOB() string`

GetDOB returns the DOB field if non-nil, zero value otherwise.

### GetDOBOk

`func (o *PatchedEmployeeUser) GetDOBOk() (*string, bool)`

GetDOBOk returns a tuple with the DOB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDOB

`func (o *PatchedEmployeeUser) SetDOB(v string)`

SetDOB sets DOB field to given value.

### HasDOB

`func (o *PatchedEmployeeUser) HasDOB() bool`

HasDOB returns a boolean if a field has been set.

### SetDOBNil

`func (o *PatchedEmployeeUser) SetDOBNil(b bool)`

 SetDOBNil sets the value for DOB to be an explicit nil

### UnsetDOB
`func (o *PatchedEmployeeUser) UnsetDOB()`

UnsetDOB ensures that no value is present for DOB, not even an explicit nil
### GetPassword

`func (o *PatchedEmployeeUser) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PatchedEmployeeUser) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PatchedEmployeeUser) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PatchedEmployeeUser) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetToken

`func (o *PatchedEmployeeUser) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PatchedEmployeeUser) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PatchedEmployeeUser) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PatchedEmployeeUser) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetProfile

`func (o *PatchedEmployeeUser) GetProfile() EmployeeProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *PatchedEmployeeUser) GetProfileOk() (*EmployeeProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *PatchedEmployeeUser) SetProfile(v EmployeeProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *PatchedEmployeeUser) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### SetProfileNil

`func (o *PatchedEmployeeUser) SetProfileNil(b bool)`

 SetProfileNil sets the value for Profile to be an explicit nil

### UnsetProfile
`func (o *PatchedEmployeeUser) UnsetProfile()`

UnsetProfile ensures that no value is present for Profile, not even an explicit nil
### GetProfilePhoto

`func (o *PatchedEmployeeUser) GetProfilePhoto() string`

GetProfilePhoto returns the ProfilePhoto field if non-nil, zero value otherwise.

### GetProfilePhotoOk

`func (o *PatchedEmployeeUser) GetProfilePhotoOk() (*string, bool)`

GetProfilePhotoOk returns a tuple with the ProfilePhoto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePhoto

`func (o *PatchedEmployeeUser) SetProfilePhoto(v string)`

SetProfilePhoto sets ProfilePhoto field to given value.

### HasProfilePhoto

`func (o *PatchedEmployeeUser) HasProfilePhoto() bool`

HasProfilePhoto returns a boolean if a field has been set.

### GetCountry

`func (o *PatchedEmployeeUser) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PatchedEmployeeUser) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PatchedEmployeeUser) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PatchedEmployeeUser) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCounty

`func (o *PatchedEmployeeUser) GetCounty() string`

GetCounty returns the County field if non-nil, zero value otherwise.

### GetCountyOk

`func (o *PatchedEmployeeUser) GetCountyOk() (*string, bool)`

GetCountyOk returns a tuple with the County field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounty

`func (o *PatchedEmployeeUser) SetCounty(v string)`

SetCounty sets County field to given value.

### HasCounty

`func (o *PatchedEmployeeUser) HasCounty() bool`

HasCounty returns a boolean if a field has been set.

### GetCity

`func (o *PatchedEmployeeUser) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *PatchedEmployeeUser) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *PatchedEmployeeUser) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *PatchedEmployeeUser) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetPostalCode

`func (o *PatchedEmployeeUser) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *PatchedEmployeeUser) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *PatchedEmployeeUser) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *PatchedEmployeeUser) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetLocation

`func (o *PatchedEmployeeUser) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PatchedEmployeeUser) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PatchedEmployeeUser) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PatchedEmployeeUser) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


