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

// TargetAudience struct for TargetAudience
type TargetAudience struct {
	Id int32 `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Phonenumber NullableString `json:"phonenumber,omitempty"`
	// Enter your gender.  * `Male` - Male * `Female` - Female * `Other` - Other
	Gender OneOfGenderEnumBlankEnumNullEnum `json:"gender,omitempty"`
	DOB NullableString `json:"DOB,omitempty"`
}

// NewTargetAudience instantiates a new TargetAudience object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetAudience(id int32, username string, email string, firstName string, lastName string) *TargetAudience {
	this := TargetAudience{}
	this.Id = id
	this.Username = username
	this.Email = email
	this.FirstName = firstName
	this.LastName = lastName
	return &this
}

// NewTargetAudienceWithDefaults instantiates a new TargetAudience object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetAudienceWithDefaults() *TargetAudience {
	this := TargetAudience{}
	return &this
}

// GetId returns the Id field value
func (o *TargetAudience) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TargetAudience) GetIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TargetAudience) SetId(v int32) {
	o.Id = v
}

// GetUsername returns the Username field value
func (o *TargetAudience) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *TargetAudience) GetUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *TargetAudience) SetUsername(v string) {
	o.Username = v
}

// GetEmail returns the Email field value
func (o *TargetAudience) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *TargetAudience) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *TargetAudience) SetEmail(v string) {
	o.Email = v
}

// GetFirstName returns the FirstName field value
func (o *TargetAudience) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *TargetAudience) GetFirstNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *TargetAudience) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *TargetAudience) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *TargetAudience) GetLastNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *TargetAudience) SetLastName(v string) {
	o.LastName = v
}

// GetPhonenumber returns the Phonenumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TargetAudience) GetPhonenumber() string {
	if o == nil || o.Phonenumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.Phonenumber.Get()
}

// GetPhonenumberOk returns a tuple with the Phonenumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TargetAudience) GetPhonenumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Phonenumber.Get(), o.Phonenumber.IsSet()
}

// HasPhonenumber returns a boolean if a field has been set.
func (o *TargetAudience) HasPhonenumber() bool {
	if o != nil && o.Phonenumber.IsSet() {
		return true
	}

	return false
}

// SetPhonenumber gets a reference to the given NullableString and assigns it to the Phonenumber field.
func (o *TargetAudience) SetPhonenumber(v string) {
	o.Phonenumber.Set(&v)
}
// SetPhonenumberNil sets the value for Phonenumber to be an explicit nil
func (o *TargetAudience) SetPhonenumberNil() {
	o.Phonenumber.Set(nil)
}

// UnsetPhonenumber ensures that no value is present for Phonenumber, not even an explicit nil
func (o *TargetAudience) UnsetPhonenumber() {
	o.Phonenumber.Unset()
}

// GetGender returns the Gender field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TargetAudience) GetGender() OneOfGenderEnumBlankEnumNullEnum {
	if o == nil  {
		var ret OneOfGenderEnumBlankEnumNullEnum
		return ret
	}
	return o.Gender
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TargetAudience) GetGenderOk() (*OneOfGenderEnumBlankEnumNullEnum, bool) {
	if o == nil || o.Gender == nil {
		return nil, false
	}
	return &o.Gender, true
}

// HasGender returns a boolean if a field has been set.
func (o *TargetAudience) HasGender() bool {
	if o != nil && o.Gender != nil {
		return true
	}

	return false
}

// SetGender gets a reference to the given OneOfGenderEnumBlankEnumNullEnum and assigns it to the Gender field.
func (o *TargetAudience) SetGender(v OneOfGenderEnumBlankEnumNullEnum) {
	o.Gender = v
}

// GetDOB returns the DOB field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TargetAudience) GetDOB() string {
	if o == nil || o.DOB.Get() == nil {
		var ret string
		return ret
	}
	return *o.DOB.Get()
}

// GetDOBOk returns a tuple with the DOB field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TargetAudience) GetDOBOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DOB.Get(), o.DOB.IsSet()
}

// HasDOB returns a boolean if a field has been set.
func (o *TargetAudience) HasDOB() bool {
	if o != nil && o.DOB.IsSet() {
		return true
	}

	return false
}

// SetDOB gets a reference to the given NullableString and assigns it to the DOB field.
func (o *TargetAudience) SetDOB(v string) {
	o.DOB.Set(&v)
}
// SetDOBNil sets the value for DOB to be an explicit nil
func (o *TargetAudience) SetDOBNil() {
	o.DOB.Set(nil)
}

// UnsetDOB ensures that no value is present for DOB, not even an explicit nil
func (o *TargetAudience) UnsetDOB() {
	o.DOB.Unset()
}

func (o TargetAudience) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["username"] = o.Username
	}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["first_name"] = o.FirstName
	}
	if true {
		toSerialize["last_name"] = o.LastName
	}
	if o.Phonenumber.IsSet() {
		toSerialize["phonenumber"] = o.Phonenumber.Get()
	}
	if o.Gender != nil {
		toSerialize["gender"] = o.Gender
	}
	if o.DOB.IsSet() {
		toSerialize["DOB"] = o.DOB.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableTargetAudience struct {
	value *TargetAudience
	isSet bool
}

func (v NullableTargetAudience) Get() *TargetAudience {
	return v.value
}

func (v *NullableTargetAudience) Set(val *TargetAudience) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetAudience) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetAudience) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetAudience(val *TargetAudience) *NullableTargetAudience {
	return &NullableTargetAudience{value: val, isSet: true}
}

func (v NullableTargetAudience) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTargetAudience) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


