# PatchedClientUser

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
**Profile** | Pointer to [**ClientProfile**](ClientProfile.md) |  | [optional] 
**ProfilePhoto** | Pointer to **string** |  | [optional] [readonly] 
**Country** | Pointer to **string** |  | [optional] [readonly] 
**County** | Pointer to **string** |  | [optional] [readonly] 
**City** | Pointer to **string** |  | [optional] [readonly] 
**PostalCode** | Pointer to **string** |  | [optional] [readonly] 
**Location** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewPatchedClientUser

`func NewPatchedClientUser() *PatchedClientUser`

NewPatchedClientUser instantiates a new PatchedClientUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedClientUserWithDefaults

`func NewPatchedClientUserWithDefaults() *PatchedClientUser`

NewPatchedClientUserWithDefaults instantiates a new PatchedClientUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedClientUser) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedClientUser) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedClientUser) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedClientUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFirstName

`func (o *PatchedClientUser) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *PatchedClientUser) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *PatchedClientUser) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *PatchedClientUser) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *PatchedClientUser) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *PatchedClientUser) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *PatchedClientUser) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *PatchedClientUser) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetUsername

`func (o *PatchedClientUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PatchedClientUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PatchedClientUser) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *PatchedClientUser) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPhonenumber

`func (o *PatchedClientUser) GetPhonenumber() string`

GetPhonenumber returns the Phonenumber field if non-nil, zero value otherwise.

### GetPhonenumberOk

`func (o *PatchedClientUser) GetPhonenumberOk() (*string, bool)`

GetPhonenumberOk returns a tuple with the Phonenumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhonenumber

`func (o *PatchedClientUser) SetPhonenumber(v string)`

SetPhonenumber sets Phonenumber field to given value.

### HasPhonenumber

`func (o *PatchedClientUser) HasPhonenumber() bool`

HasPhonenumber returns a boolean if a field has been set.

### SetPhonenumberNil

`func (o *PatchedClientUser) SetPhonenumberNil(b bool)`

 SetPhonenumberNil sets the value for Phonenumber to be an explicit nil

### UnsetPhonenumber
`func (o *PatchedClientUser) UnsetPhonenumber()`

UnsetPhonenumber ensures that no value is present for Phonenumber, not even an explicit nil
### GetEmail

`func (o *PatchedClientUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PatchedClientUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PatchedClientUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PatchedClientUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetGender

`func (o *PatchedClientUser) GetGender() OneOfGenderEnumBlankEnumNullEnum`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *PatchedClientUser) GetGenderOk() (*OneOfGenderEnumBlankEnumNullEnum, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *PatchedClientUser) SetGender(v OneOfGenderEnumBlankEnumNullEnum)`

SetGender sets Gender field to given value.

### HasGender

`func (o *PatchedClientUser) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *PatchedClientUser) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *PatchedClientUser) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetDOB

`func (o *PatchedClientUser) GetDOB() string`

GetDOB returns the DOB field if non-nil, zero value otherwise.

### GetDOBOk

`func (o *PatchedClientUser) GetDOBOk() (*string, bool)`

GetDOBOk returns a tuple with the DOB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDOB

`func (o *PatchedClientUser) SetDOB(v string)`

SetDOB sets DOB field to given value.

### HasDOB

`func (o *PatchedClientUser) HasDOB() bool`

HasDOB returns a boolean if a field has been set.

### SetDOBNil

`func (o *PatchedClientUser) SetDOBNil(b bool)`

 SetDOBNil sets the value for DOB to be an explicit nil

### UnsetDOB
`func (o *PatchedClientUser) UnsetDOB()`

UnsetDOB ensures that no value is present for DOB, not even an explicit nil
### GetPassword

`func (o *PatchedClientUser) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PatchedClientUser) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PatchedClientUser) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PatchedClientUser) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetToken

`func (o *PatchedClientUser) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PatchedClientUser) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PatchedClientUser) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PatchedClientUser) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetProfile

`func (o *PatchedClientUser) GetProfile() ClientProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *PatchedClientUser) GetProfileOk() (*ClientProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *PatchedClientUser) SetProfile(v ClientProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *PatchedClientUser) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### SetProfileNil

`func (o *PatchedClientUser) SetProfileNil(b bool)`

 SetProfileNil sets the value for Profile to be an explicit nil

### UnsetProfile
`func (o *PatchedClientUser) UnsetProfile()`

UnsetProfile ensures that no value is present for Profile, not even an explicit nil
### GetProfilePhoto

`func (o *PatchedClientUser) GetProfilePhoto() string`

GetProfilePhoto returns the ProfilePhoto field if non-nil, zero value otherwise.

### GetProfilePhotoOk

`func (o *PatchedClientUser) GetProfilePhotoOk() (*string, bool)`

GetProfilePhotoOk returns a tuple with the ProfilePhoto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePhoto

`func (o *PatchedClientUser) SetProfilePhoto(v string)`

SetProfilePhoto sets ProfilePhoto field to given value.

### HasProfilePhoto

`func (o *PatchedClientUser) HasProfilePhoto() bool`

HasProfilePhoto returns a boolean if a field has been set.

### GetCountry

`func (o *PatchedClientUser) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PatchedClientUser) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PatchedClientUser) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PatchedClientUser) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCounty

`func (o *PatchedClientUser) GetCounty() string`

GetCounty returns the County field if non-nil, zero value otherwise.

### GetCountyOk

`func (o *PatchedClientUser) GetCountyOk() (*string, bool)`

GetCountyOk returns a tuple with the County field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounty

`func (o *PatchedClientUser) SetCounty(v string)`

SetCounty sets County field to given value.

### HasCounty

`func (o *PatchedClientUser) HasCounty() bool`

HasCounty returns a boolean if a field has been set.

### GetCity

`func (o *PatchedClientUser) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *PatchedClientUser) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *PatchedClientUser) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *PatchedClientUser) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetPostalCode

`func (o *PatchedClientUser) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *PatchedClientUser) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *PatchedClientUser) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *PatchedClientUser) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetLocation

`func (o *PatchedClientUser) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PatchedClientUser) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PatchedClientUser) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PatchedClientUser) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


