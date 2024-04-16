# EmployeeProfile

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

### NewEmployeeProfile

`func NewEmployeeProfile(username string, profilePhoto string, ) *EmployeeProfile`

NewEmployeeProfile instantiates a new EmployeeProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployeeProfileWithDefaults

`func NewEmployeeProfileWithDefaults() *EmployeeProfile`

NewEmployeeProfileWithDefaults instantiates a new EmployeeProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *EmployeeProfile) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *EmployeeProfile) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *EmployeeProfile) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetProfilePhoto

`func (o *EmployeeProfile) GetProfilePhoto() string`

GetProfilePhoto returns the ProfilePhoto field if non-nil, zero value otherwise.

### GetProfilePhotoOk

`func (o *EmployeeProfile) GetProfilePhotoOk() (*string, bool)`

GetProfilePhotoOk returns a tuple with the ProfilePhoto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePhoto

`func (o *EmployeeProfile) SetProfilePhoto(v string)`

SetProfilePhoto sets ProfilePhoto field to given value.


### GetCountry

`func (o *EmployeeProfile) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *EmployeeProfile) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *EmployeeProfile) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *EmployeeProfile) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCounty

`func (o *EmployeeProfile) GetCounty() string`

GetCounty returns the County field if non-nil, zero value otherwise.

### GetCountyOk

`func (o *EmployeeProfile) GetCountyOk() (*string, bool)`

GetCountyOk returns a tuple with the County field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounty

`func (o *EmployeeProfile) SetCounty(v string)`

SetCounty sets County field to given value.

### HasCounty

`func (o *EmployeeProfile) HasCounty() bool`

HasCounty returns a boolean if a field has been set.

### GetCity

`func (o *EmployeeProfile) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *EmployeeProfile) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *EmployeeProfile) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *EmployeeProfile) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetPostalCode

`func (o *EmployeeProfile) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *EmployeeProfile) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *EmployeeProfile) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *EmployeeProfile) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetLocation

`func (o *EmployeeProfile) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *EmployeeProfile) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *EmployeeProfile) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *EmployeeProfile) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


