# ClientUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**FirstName** | **string** |  | 
**LastName** | **string** |  | 
**Username** | **string** |  | 
**Phonenumber** | Pointer to **NullableString** |  | [optional] 
**Email** | **string** |  | 
**Gender** | Pointer to [**OneOfGenderEnumBlankEnumNullEnum**](oneOf&lt;GenderEnum,BlankEnum,NullEnum&gt;.md) | Enter your gender.  * &#x60;Male&#x60; - Male * &#x60;Female&#x60; - Female * &#x60;Other&#x60; - Other | [optional] 
**DOB** | Pointer to **NullableString** |  | [optional] 
**Password** | **string** |  | 
**Token** | **string** |  | [readonly] 
**Profile** | [**ClientProfile**](ClientProfile.md) |  | 
**ProfilePhoto** | **string** |  | [readonly] 
**Country** | **string** |  | [readonly] 
**County** | **string** |  | [readonly] 
**City** | **string** |  | [readonly] 
**PostalCode** | **string** |  | [readonly] 
**Location** | **string** |  | [readonly] 

## Methods

### NewClientUser

`func NewClientUser(id int32, firstName string, lastName string, username string, email string, password string, token string, profile ClientProfile, profilePhoto string, country string, county string, city string, postalCode string, location string, ) *ClientUser`

NewClientUser instantiates a new ClientUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientUserWithDefaults

`func NewClientUserWithDefaults() *ClientUser`

NewClientUserWithDefaults instantiates a new ClientUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClientUser) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClientUser) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClientUser) SetId(v int32)`

SetId sets Id field to given value.


### GetFirstName

`func (o *ClientUser) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ClientUser) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ClientUser) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *ClientUser) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ClientUser) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ClientUser) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetUsername

`func (o *ClientUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ClientUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ClientUser) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPhonenumber

`func (o *ClientUser) GetPhonenumber() string`

GetPhonenumber returns the Phonenumber field if non-nil, zero value otherwise.

### GetPhonenumberOk

`func (o *ClientUser) GetPhonenumberOk() (*string, bool)`

GetPhonenumberOk returns a tuple with the Phonenumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhonenumber

`func (o *ClientUser) SetPhonenumber(v string)`

SetPhonenumber sets Phonenumber field to given value.

### HasPhonenumber

`func (o *ClientUser) HasPhonenumber() bool`

HasPhonenumber returns a boolean if a field has been set.

### SetPhonenumberNil

`func (o *ClientUser) SetPhonenumberNil(b bool)`

 SetPhonenumberNil sets the value for Phonenumber to be an explicit nil

### UnsetPhonenumber
`func (o *ClientUser) UnsetPhonenumber()`

UnsetPhonenumber ensures that no value is present for Phonenumber, not even an explicit nil
### GetEmail

`func (o *ClientUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ClientUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ClientUser) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetGender

`func (o *ClientUser) GetGender() OneOfGenderEnumBlankEnumNullEnum`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *ClientUser) GetGenderOk() (*OneOfGenderEnumBlankEnumNullEnum, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *ClientUser) SetGender(v OneOfGenderEnumBlankEnumNullEnum)`

SetGender sets Gender field to given value.

### HasGender

`func (o *ClientUser) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *ClientUser) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *ClientUser) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetDOB

`func (o *ClientUser) GetDOB() string`

GetDOB returns the DOB field if non-nil, zero value otherwise.

### GetDOBOk

`func (o *ClientUser) GetDOBOk() (*string, bool)`

GetDOBOk returns a tuple with the DOB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDOB

`func (o *ClientUser) SetDOB(v string)`

SetDOB sets DOB field to given value.

### HasDOB

`func (o *ClientUser) HasDOB() bool`

HasDOB returns a boolean if a field has been set.

### SetDOBNil

`func (o *ClientUser) SetDOBNil(b bool)`

 SetDOBNil sets the value for DOB to be an explicit nil

### UnsetDOB
`func (o *ClientUser) UnsetDOB()`

UnsetDOB ensures that no value is present for DOB, not even an explicit nil
### GetPassword

`func (o *ClientUser) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ClientUser) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ClientUser) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetToken

`func (o *ClientUser) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ClientUser) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ClientUser) SetToken(v string)`

SetToken sets Token field to given value.


### GetProfile

`func (o *ClientUser) GetProfile() ClientProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ClientUser) GetProfileOk() (*ClientProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ClientUser) SetProfile(v ClientProfile)`

SetProfile sets Profile field to given value.


### SetProfileNil

`func (o *ClientUser) SetProfileNil(b bool)`

 SetProfileNil sets the value for Profile to be an explicit nil

### UnsetProfile
`func (o *ClientUser) UnsetProfile()`

UnsetProfile ensures that no value is present for Profile, not even an explicit nil
### GetProfilePhoto

`func (o *ClientUser) GetProfilePhoto() string`

GetProfilePhoto returns the ProfilePhoto field if non-nil, zero value otherwise.

### GetProfilePhotoOk

`func (o *ClientUser) GetProfilePhotoOk() (*string, bool)`

GetProfilePhotoOk returns a tuple with the ProfilePhoto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePhoto

`func (o *ClientUser) SetProfilePhoto(v string)`

SetProfilePhoto sets ProfilePhoto field to given value.


### GetCountry

`func (o *ClientUser) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ClientUser) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ClientUser) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetCounty

`func (o *ClientUser) GetCounty() string`

GetCounty returns the County field if non-nil, zero value otherwise.

### GetCountyOk

`func (o *ClientUser) GetCountyOk() (*string, bool)`

GetCountyOk returns a tuple with the County field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounty

`func (o *ClientUser) SetCounty(v string)`

SetCounty sets County field to given value.


### GetCity

`func (o *ClientUser) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *ClientUser) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *ClientUser) SetCity(v string)`

SetCity sets City field to given value.


### GetPostalCode

`func (o *ClientUser) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *ClientUser) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *ClientUser) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.


### GetLocation

`func (o *ClientUser) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ClientUser) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ClientUser) SetLocation(v string)`

SetLocation sets Location field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


