# AdminProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** |  | 
**ProfilePhoto** | **string** |  | [readonly] 
**Country** | Pointer to **string** |  | [optional] 
**County** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**PostalCode** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 

## Methods

### NewAdminProfile

`func NewAdminProfile(username string, profilePhoto string, ) *AdminProfile`

NewAdminProfile instantiates a new AdminProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminProfileWithDefaults

`func NewAdminProfileWithDefaults() *AdminProfile`

NewAdminProfileWithDefaults instantiates a new AdminProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *AdminProfile) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AdminProfile) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AdminProfile) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetProfilePhoto

`func (o *AdminProfile) GetProfilePhoto() string`

GetProfilePhoto returns the ProfilePhoto field if non-nil, zero value otherwise.

### GetProfilePhotoOk

`func (o *AdminProfile) GetProfilePhotoOk() (*string, bool)`

GetProfilePhotoOk returns a tuple with the ProfilePhoto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePhoto

`func (o *AdminProfile) SetProfilePhoto(v string)`

SetProfilePhoto sets ProfilePhoto field to given value.


### GetCountry

`func (o *AdminProfile) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *AdminProfile) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *AdminProfile) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *AdminProfile) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCounty

`func (o *AdminProfile) GetCounty() string`

GetCounty returns the County field if non-nil, zero value otherwise.

### GetCountyOk

`func (o *AdminProfile) GetCountyOk() (*string, bool)`

GetCountyOk returns a tuple with the County field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounty

`func (o *AdminProfile) SetCounty(v string)`

SetCounty sets County field to given value.

### HasCounty

`func (o *AdminProfile) HasCounty() bool`

HasCounty returns a boolean if a field has been set.

### GetCity

`func (o *AdminProfile) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *AdminProfile) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *AdminProfile) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *AdminProfile) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetPostalCode

`func (o *AdminProfile) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *AdminProfile) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *AdminProfile) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *AdminProfile) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetLocation

`func (o *AdminProfile) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AdminProfile) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AdminProfile) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *AdminProfile) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


