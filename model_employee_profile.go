/*
proxima-core-engine

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// EmployeeProfile struct for EmployeeProfile
type EmployeeProfile struct {
	Username string `json:"username"`
	ProfilePhoto string `json:"profile_photo"`
	Country *string `json:"country,omitempty"`
	County *string `json:"county,omitempty"`
	City *string `json:"city,omitempty"`
	PostalCode *string `json:"postal_code,omitempty"`
	Location *string `json:"location,omitempty"`
}

// NewEmployeeProfile instantiates a new EmployeeProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmployeeProfile(username string, profilePhoto string) *EmployeeProfile {
	this := EmployeeProfile{}
	this.Username = username
	this.ProfilePhoto = profilePhoto
	return &this
}

// NewEmployeeProfileWithDefaults instantiates a new EmployeeProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmployeeProfileWithDefaults() *EmployeeProfile {
	this := EmployeeProfile{}
	return &this
}

// GetUsername returns the Username field value
func (o *EmployeeProfile) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *EmployeeProfile) GetUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *EmployeeProfile) SetUsername(v string) {
	o.Username = v
}

// GetProfilePhoto returns the ProfilePhoto field value
func (o *EmployeeProfile) GetProfilePhoto() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProfilePhoto
}

// GetProfilePhotoOk returns a tuple with the ProfilePhoto field value
// and a boolean to check if the value has been set.
func (o *EmployeeProfile) GetProfilePhotoOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProfilePhoto, true
}

// SetProfilePhoto sets field value
func (o *EmployeeProfile) SetProfilePhoto(v string) {
	o.ProfilePhoto = v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *EmployeeProfile) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployeeProfile) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *EmployeeProfile) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *EmployeeProfile) SetCountry(v string) {
	o.Country = &v
}

// GetCounty returns the County field value if set, zero value otherwise.
func (o *EmployeeProfile) GetCounty() string {
	if o == nil || o.County == nil {
		var ret string
		return ret
	}
	return *o.County
}

// GetCountyOk returns a tuple with the County field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployeeProfile) GetCountyOk() (*string, bool) {
	if o == nil || o.County == nil {
		return nil, false
	}
	return o.County, true
}

// HasCounty returns a boolean if a field has been set.
func (o *EmployeeProfile) HasCounty() bool {
	if o != nil && o.County != nil {
		return true
	}

	return false
}

// SetCounty gets a reference to the given string and assigns it to the County field.
func (o *EmployeeProfile) SetCounty(v string) {
	o.County = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *EmployeeProfile) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployeeProfile) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *EmployeeProfile) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *EmployeeProfile) SetCity(v string) {
	o.City = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *EmployeeProfile) GetPostalCode() string {
	if o == nil || o.PostalCode == nil {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployeeProfile) GetPostalCodeOk() (*string, bool) {
	if o == nil || o.PostalCode == nil {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *EmployeeProfile) HasPostalCode() bool {
	if o != nil && o.PostalCode != nil {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *EmployeeProfile) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *EmployeeProfile) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployeeProfile) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *EmployeeProfile) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *EmployeeProfile) SetLocation(v string) {
	o.Location = &v
}

func (o EmployeeProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["username"] = o.Username
	}
	if true {
		toSerialize["profile_photo"] = o.ProfilePhoto
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	if o.County != nil {
		toSerialize["county"] = o.County
	}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	if o.PostalCode != nil {
		toSerialize["postal_code"] = o.PostalCode
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	return json.Marshal(toSerialize)
}

type NullableEmployeeProfile struct {
	value *EmployeeProfile
	isSet bool
}

func (v NullableEmployeeProfile) Get() *EmployeeProfile {
	return v.value
}

func (v *NullableEmployeeProfile) Set(val *EmployeeProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableEmployeeProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableEmployeeProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmployeeProfile(val *EmployeeProfile) *NullableEmployeeProfile {
	return &NullableEmployeeProfile{value: val, isSet: true}
}

func (v NullableEmployeeProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmployeeProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


