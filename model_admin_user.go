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

// AdminUser Handles serialization and deserialization of user objects
type AdminUser struct {
	Id int32 `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Username string `json:"username"`
	Phonenumber NullableString `json:"phonenumber,omitempty"`
	Email string `json:"email"`
	// Enter your gender.  * `Male` - Male * `Female` - Female * `Other` - Other
	Gender OneOfGenderEnumBlankEnumNullEnum `json:"gender,omitempty"`
	DOB NullableString `json:"DOB,omitempty"`
	Password string `json:"password"`
	Token string `json:"token"`
	Profile AdminProfile `json:"profile"`
	ProfilePhoto string `json:"profile_photo"`
	Country string `json:"country"`
	County string `json:"county"`
	City string `json:"city"`
	PostalCode string `json:"postal_code"`
	Location string `json:"location"`
}

// NewAdminUser instantiates a new AdminUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminUser(id int32, firstName string, lastName string, username string, email string, password string, token string, profile AdminProfile, profilePhoto string, country string, county string, city string, postalCode string, location string) *AdminUser {
	this := AdminUser{}
	this.Id = id
	this.FirstName = firstName
	this.LastName = lastName
	this.Username = username
	this.Email = email
	this.Password = password
	this.Token = token
	this.Profile = profile
	this.ProfilePhoto = profilePhoto
	this.Country = country
	this.County = county
	this.City = city
	this.PostalCode = postalCode
	this.Location = location
	return &this
}

// NewAdminUserWithDefaults instantiates a new AdminUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminUserWithDefaults() *AdminUser {
	this := AdminUser{}
	return &this
}

// GetId returns the Id field value
func (o *AdminUser) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AdminUser) GetIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AdminUser) SetId(v int32) {
	o.Id = v
}

// GetFirstName returns the FirstName field value
func (o *AdminUser) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *AdminUser) GetFirstNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *AdminUser) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *AdminUser) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *AdminUser) GetLastNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *AdminUser) SetLastName(v string) {
	o.LastName = v
}

// GetUsername returns the Username field value
func (o *AdminUser) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *AdminUser) GetUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *AdminUser) SetUsername(v string) {
	o.Username = v
}

// GetPhonenumber returns the Phonenumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdminUser) GetPhonenumber() string {
	if o == nil || o.Phonenumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.Phonenumber.Get()
}

// GetPhonenumberOk returns a tuple with the Phonenumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdminUser) GetPhonenumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Phonenumber.Get(), o.Phonenumber.IsSet()
}

// HasPhonenumber returns a boolean if a field has been set.
func (o *AdminUser) HasPhonenumber() bool {
	if o != nil && o.Phonenumber.IsSet() {
		return true
	}

	return false
}

// SetPhonenumber gets a reference to the given NullableString and assigns it to the Phonenumber field.
func (o *AdminUser) SetPhonenumber(v string) {
	o.Phonenumber.Set(&v)
}
// SetPhonenumberNil sets the value for Phonenumber to be an explicit nil
func (o *AdminUser) SetPhonenumberNil() {
	o.Phonenumber.Set(nil)
}

// UnsetPhonenumber ensures that no value is present for Phonenumber, not even an explicit nil
func (o *AdminUser) UnsetPhonenumber() {
	o.Phonenumber.Unset()
}

// GetEmail returns the Email field value
func (o *AdminUser) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *AdminUser) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *AdminUser) SetEmail(v string) {
	o.Email = v
}

// GetGender returns the Gender field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdminUser) GetGender() OneOfGenderEnumBlankEnumNullEnum {
	if o == nil  {
		var ret OneOfGenderEnumBlankEnumNullEnum
		return ret
	}
	return o.Gender
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdminUser) GetGenderOk() (*OneOfGenderEnumBlankEnumNullEnum, bool) {
	if o == nil || o.Gender == nil {
		return nil, false
	}
	return &o.Gender, true
}

// HasGender returns a boolean if a field has been set.
func (o *AdminUser) HasGender() bool {
	if o != nil && o.Gender != nil {
		return true
	}

	return false
}

// SetGender gets a reference to the given OneOfGenderEnumBlankEnumNullEnum and assigns it to the Gender field.
func (o *AdminUser) SetGender(v OneOfGenderEnumBlankEnumNullEnum) {
	o.Gender = v
}

// GetDOB returns the DOB field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdminUser) GetDOB() string {
	if o == nil || o.DOB.Get() == nil {
		var ret string
		return ret
	}
	return *o.DOB.Get()
}

// GetDOBOk returns a tuple with the DOB field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdminUser) GetDOBOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DOB.Get(), o.DOB.IsSet()
}

// HasDOB returns a boolean if a field has been set.
func (o *AdminUser) HasDOB() bool {
	if o != nil && o.DOB.IsSet() {
		return true
	}

	return false
}

// SetDOB gets a reference to the given NullableString and assigns it to the DOB field.
func (o *AdminUser) SetDOB(v string) {
	o.DOB.Set(&v)
}
// SetDOBNil sets the value for DOB to be an explicit nil
func (o *AdminUser) SetDOBNil() {
	o.DOB.Set(nil)
}

// UnsetDOB ensures that no value is present for DOB, not even an explicit nil
func (o *AdminUser) UnsetDOB() {
	o.DOB.Unset()
}

// GetPassword returns the Password field value
func (o *AdminUser) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *AdminUser) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *AdminUser) SetPassword(v string) {
	o.Password = v
}

// GetToken returns the Token field value
func (o *AdminUser) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *AdminUser) GetTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *AdminUser) SetToken(v string) {
	o.Token = v
}

// GetProfile returns the Profile field value
// If the value is explicit nil, the zero value for AdminProfile will be returned
func (o *AdminUser) GetProfile() AdminProfile {
	if o == nil {
		var ret AdminProfile
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdminUser) GetProfileOk() (*AdminProfile, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *AdminUser) SetProfile(v AdminProfile) {
	o.Profile = v
}

// GetProfilePhoto returns the ProfilePhoto field value
func (o *AdminUser) GetProfilePhoto() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProfilePhoto
}

// GetProfilePhotoOk returns a tuple with the ProfilePhoto field value
// and a boolean to check if the value has been set.
func (o *AdminUser) GetProfilePhotoOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProfilePhoto, true
}

// SetProfilePhoto sets field value
func (o *AdminUser) SetProfilePhoto(v string) {
	o.ProfilePhoto = v
}

// GetCountry returns the Country field value
func (o *AdminUser) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *AdminUser) GetCountryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *AdminUser) SetCountry(v string) {
	o.Country = v
}

// GetCounty returns the County field value
func (o *AdminUser) GetCounty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.County
}

// GetCountyOk returns a tuple with the County field value
// and a boolean to check if the value has been set.
func (o *AdminUser) GetCountyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.County, true
}

// SetCounty sets field value
func (o *AdminUser) SetCounty(v string) {
	o.County = v
}

// GetCity returns the City field value
func (o *AdminUser) GetCity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.City
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
func (o *AdminUser) GetCityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.City, true
}

// SetCity sets field value
func (o *AdminUser) SetCity(v string) {
	o.City = v
}

// GetPostalCode returns the PostalCode field value
func (o *AdminUser) GetPostalCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value
// and a boolean to check if the value has been set.
func (o *AdminUser) GetPostalCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PostalCode, true
}

// SetPostalCode sets field value
func (o *AdminUser) SetPostalCode(v string) {
	o.PostalCode = v
}

// GetLocation returns the Location field value
func (o *AdminUser) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *AdminUser) GetLocationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *AdminUser) SetLocation(v string) {
	o.Location = v
}

func (o AdminUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["first_name"] = o.FirstName
	}
	if true {
		toSerialize["last_name"] = o.LastName
	}
	if true {
		toSerialize["username"] = o.Username
	}
	if o.Phonenumber.IsSet() {
		toSerialize["phonenumber"] = o.Phonenumber.Get()
	}
	if true {
		toSerialize["email"] = o.Email
	}
	if o.Gender != nil {
		toSerialize["gender"] = o.Gender
	}
	if o.DOB.IsSet() {
		toSerialize["DOB"] = o.DOB.Get()
	}
	if true {
		toSerialize["password"] = o.Password
	}
	if true {
		toSerialize["token"] = o.Token
	}
	if o.Profile != nil {
		toSerialize["profile"] = o.Profile
	}
	if true {
		toSerialize["profile_photo"] = o.ProfilePhoto
	}
	if true {
		toSerialize["country"] = o.Country
	}
	if true {
		toSerialize["county"] = o.County
	}
	if true {
		toSerialize["city"] = o.City
	}
	if true {
		toSerialize["postal_code"] = o.PostalCode
	}
	if true {
		toSerialize["location"] = o.Location
	}
	return json.Marshal(toSerialize)
}

type NullableAdminUser struct {
	value *AdminUser
	isSet bool
}

func (v NullableAdminUser) Get() *AdminUser {
	return v.value
}

func (v *NullableAdminUser) Set(val *AdminUser) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminUser) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminUser(val *AdminUser) *NullableAdminUser {
	return &NullableAdminUser{value: val, isSet: true}
}

func (v NullableAdminUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

