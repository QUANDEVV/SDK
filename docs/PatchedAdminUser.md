# PatchedAdminUser

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
**Profile** | Pointer to [**AdminProfile**](AdminProfile.md) |  | [optional] 
**ProfilePhoto** | Pointer to **string** |  | [optional] [readonly] 
**Country** | Pointer to **string** |  | [optional] [readonly] 
**County** | Pointer to **string** |  | [optional] [readonly] 
**City** | Pointer to **string** |  | [optional] [readonly] 
**PostalCode** | Pointer to **string** |  | [optional] [readonly] 
**Location** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewPatchedAdminUser

`func NewPatchedAdminUser() *PatchedAdminUser`

NewPatchedAdminUser instantiates a new PatchedAdminUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedAdminUserWithDefaults

`func NewPatchedAdminUserWithDefaults() *PatchedAdminUser`

NewPatchedAdminUserWithDefaults instantiates a new PatchedAdminUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedAdminUser) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedAdminUser) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedAdminUser) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedAdminUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFirstName

`func (o *PatchedAdminUser) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *PatchedAdminUser) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *PatchedAdminUser) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *PatchedAdminUser) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *PatchedAdminUser) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *PatchedAdminUser) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *PatchedAdminUser) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *PatchedAdminUser) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetUsername

`func (o *PatchedAdminUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PatchedAdminUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PatchedAdminUser) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *PatchedAdminUser) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPhonenumber

`func (o *PatchedAdminUser) GetPhonenumber() string`

GetPhonenumber returns the Phonenumber field if non-nil, zero value otherwise.

### GetPhonenumberOk

`func (o *PatchedAdminUser) GetPhonenumberOk() (*string, bool)`

GetPhonenumberOk returns a tuple with the Phonenumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhonenumber

`func (o *PatchedAdminUser) SetPhonenumber(v string)`

SetPhonenumber sets Phonenumber field to given value.

### HasPhonenumber

`func (o *PatchedAdminUser) HasPhonenumber() bool`

HasPhonenumber returns a boolean if a field has been set.

### SetPhonenumberNil

`func (o *PatchedAdminUser) SetPhonenumberNil(b bool)`

 SetPhonenumberNil sets the value for Phonenumber to be an explicit nil

### UnsetPhonenumber
`func (o *PatchedAdminUser) UnsetPhonenumber()`

UnsetPhonenumber ensures that no value is present for Phonenumber, not even an explicit nil
### GetEmail

`func (o *PatchedAdminUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PatchedAdminUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PatchedAdminUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PatchedAdminUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetGender

`func (o *PatchedAdminUser) GetGender() OneOfGenderEnumBlankEnumNullEnum`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *PatchedAdminUser) GetGenderOk() (*OneOfGenderEnumBlankEnumNullEnum, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *PatchedAdminUser) SetGender(v OneOfGenderEnumBlankEnumNullEnum)`

SetGender sets Gender field to given value.

### HasGender

`func (o *PatchedAdminUser) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *PatchedAdminUser) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *PatchedAdminUser) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetDOB

`func (o *PatchedAdminUser) GetDOB() string`

GetDOB returns the DOB field if non-nil, zero value otherwise.

### GetDOBOk

`func (o *PatchedAdminUser) GetDOBOk() (*string, bool)`

GetDOBOk returns a tuple with the DOB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDOB

`func (o *PatchedAdminUser) SetDOB(v string)`

SetDOB sets DOB field to given value.

### HasDOB

`func (o *PatchedAdminUser) HasDOB() bool`

HasDOB returns a boolean if a field has been set.

### SetDOBNil

`func (o *PatchedAdminUser) SetDOBNil(b bool)`

 SetDOBNil sets the value for DOB to be an explicit nil

### UnsetDOB
`func (o *PatchedAdminUser) UnsetDOB()`

UnsetDOB ensures that no value is present for DOB, not even an explicit nil
### GetPassword

`func (o *PatchedAdminUser) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PatchedAdminUser) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PatchedAdminUser) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PatchedAdminUser) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetToken

`func (o *PatchedAdminUser) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PatchedAdminUser) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PatchedAdminUser) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PatchedAdminUser) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetProfile

`func (o *PatchedAdminUser) GetProfile() AdminProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *PatchedAdminUser) GetProfileOk() (*AdminProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *PatchedAdminUser) SetProfile(v AdminProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *PatchedAdminUser) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### SetProfileNil

`func (o *PatchedAdminUser) SetProfileNil(b bool)`

 SetProfileNil sets the value for Profile to be an explicit nil

### UnsetProfile
`func (o *PatchedAdminUser) UnsetProfile()`

UnsetProfile ensures that no value is present for Profile, not even an explicit nil
### GetProfilePhoto

`func (o *PatchedAdminUser) GetProfilePhoto() string`

GetProfilePhoto returns the ProfilePhoto field if non-nil, zero value otherwise.

### GetProfilePhotoOk

`func (o *PatchedAdminUser) GetProfilePhotoOk() (*string, bool)`

GetProfilePhotoOk returns a tuple with the ProfilePhoto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePhoto

`func (o *PatchedAdminUser) SetProfilePhoto(v string)`

SetProfilePhoto sets ProfilePhoto field to given value.

### HasProfilePhoto

`func (o *PatchedAdminUser) HasProfilePhoto() bool`

HasProfilePhoto returns a boolean if a field has been set.

### GetCountry

`func (o *PatchedAdminUser) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PatchedAdminUser) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PatchedAdminUser) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PatchedAdminUser) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCounty

`func (o *PatchedAdminUser) GetCounty() string`

GetCounty returns the County field if non-nil, zero value otherwise.

### GetCountyOk

`func (o *PatchedAdminUser) GetCountyOk() (*string, bool)`

GetCountyOk returns a tuple with the County field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounty

`func (o *PatchedAdminUser) SetCounty(v string)`

SetCounty sets County field to given value.

### HasCounty

`func (o *PatchedAdminUser) HasCounty() bool`

HasCounty returns a boolean if a field has been set.

### GetCity

`func (o *PatchedAdminUser) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *PatchedAdminUser) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *PatchedAdminUser) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *PatchedAdminUser) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetPostalCode

`func (o *PatchedAdminUser) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *PatchedAdminUser) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *PatchedAdminUser) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *PatchedAdminUser) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetLocation

`func (o *PatchedAdminUser) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PatchedAdminUser) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PatchedAdminUser) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PatchedAdminUser) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


