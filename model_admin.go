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

// Admin Serializers registration requests and creates a new user.
type Admin struct {
	Username string `json:"username"`
	Email string `json:"email"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Phonenumber NullableString `json:"phonenumber,omitempty"`
	// Enter your gender.  * `Male` - Male * `Female` - Female * `Other` - Other
	Gender OneOfGenderEnumBlankEnumNullEnum `json:"gender,omitempty"`
	DOB NullableString `json:"DOB,omitempty"`
	UserType OneOfUserType8ceEnumNullEnum `json:"user_type,omitempty"`
	// Display id of the tenant
	TenantId int32 `json:"tenant_id"`
	Password string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Token string `json:"token"`
}

// NewAdmin instantiates a new Admin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdmin(username string, email string, firstName string, lastName string, tenantId int32, password string, confirmPassword string, token string) *Admin {
	this := Admin{}
	this.Username = username
	this.Email = email
	this.FirstName = firstName
	this.LastName = lastName
	this.TenantId = tenantId
	this.Password = password
	this.ConfirmPassword = confirmPassword
	this.Token = token
	return &this
}

// NewAdminWithDefaults instantiates a new Admin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminWithDefaults() *Admin {
	this := Admin{}
	return &this
}

// GetUsername returns the Username field value
func (o *Admin) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *Admin) GetUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *Admin) SetUsername(v string) {
	o.Username = v
}

// GetEmail returns the Email field value
func (o *Admin) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *Admin) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *Admin) SetEmail(v string) {
	o.Email = v
}

// GetFirstName returns the FirstName field value
func (o *Admin) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *Admin) GetFirstNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *Admin) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *Admin) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *Admin) GetLastNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *Admin) SetLastName(v string) {
	o.LastName = v
}

// GetPhonenumber returns the Phonenumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Admin) GetPhonenumber() string {
	if o == nil || o.Phonenumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.Phonenumber.Get()
}

// GetPhonenumberOk returns a tuple with the Phonenumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Admin) GetPhonenumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Phonenumber.Get(), o.Phonenumber.IsSet()
}

// HasPhonenumber returns a boolean if a field has been set.
func (o *Admin) HasPhonenumber() bool {
	if o != nil && o.Phonenumber.IsSet() {
		return true
	}

	return false
}

// SetPhonenumber gets a reference to the given NullableString and assigns it to the Phonenumber field.
func (o *Admin) SetPhonenumber(v string) {
	o.Phonenumber.Set(&v)
}
// SetPhonenumberNil sets the value for Phonenumber to be an explicit nil
func (o *Admin) SetPhonenumberNil() {
	o.Phonenumber.Set(nil)
}

// UnsetPhonenumber ensures that no value is present for Phonenumber, not even an explicit nil
func (o *Admin) UnsetPhonenumber() {
	o.Phonenumber.Unset()
}

// GetGender returns the Gender field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Admin) GetGender() OneOfGenderEnumBlankEnumNullEnum {
	if o == nil  {
		var ret OneOfGenderEnumBlankEnumNullEnum
		return ret
	}
	return o.Gender
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Admin) GetGenderOk() (*OneOfGenderEnumBlankEnumNullEnum, bool) {
	if o == nil || o.Gender == nil {
		return nil, false
	}
	return &o.Gender, true
}

// HasGender returns a boolean if a field has been set.
func (o *Admin) HasGender() bool {
	if o != nil && o.Gender != nil {
		return true
	}

	return false
}

// SetGender gets a reference to the given OneOfGenderEnumBlankEnumNullEnum and assigns it to the Gender field.
func (o *Admin) SetGender(v OneOfGenderEnumBlankEnumNullEnum) {
	o.Gender = v
}

// GetDOB returns the DOB field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Admin) GetDOB() string {
	if o == nil || o.DOB.Get() == nil {
		var ret string
		return ret
	}
	return *o.DOB.Get()
}

// GetDOBOk returns a tuple with the DOB field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Admin) GetDOBOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DOB.Get(), o.DOB.IsSet()
}

// HasDOB returns a boolean if a field has been set.
func (o *Admin) HasDOB() bool {
	if o != nil && o.DOB.IsSet() {
		return true
	}

	return false
}

// SetDOB gets a reference to the given NullableString and assigns it to the DOB field.
func (o *Admin) SetDOB(v string) {
	o.DOB.Set(&v)
}
// SetDOBNil sets the value for DOB to be an explicit nil
func (o *Admin) SetDOBNil() {
	o.DOB.Set(nil)
}

// UnsetDOB ensures that no value is present for DOB, not even an explicit nil
func (o *Admin) UnsetDOB() {
	o.DOB.Unset()
}

// GetUserType returns the UserType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Admin) GetUserType() OneOfUserType8ceEnumNullEnum {
	if o == nil  {
		var ret OneOfUserType8ceEnumNullEnum
		return ret
	}
	return o.UserType
}

// GetUserTypeOk returns a tuple with the UserType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Admin) GetUserTypeOk() (*OneOfUserType8ceEnumNullEnum, bool) {
	if o == nil || o.UserType == nil {
		return nil, false
	}
	return &o.UserType, true
}

// HasUserType returns a boolean if a field has been set.
func (o *Admin) HasUserType() bool {
	if o != nil && o.UserType != nil {
		return true
	}

	return false
}

// SetUserType gets a reference to the given OneOfUserType8ceEnumNullEnum and assigns it to the UserType field.
func (o *Admin) SetUserType(v OneOfUserType8ceEnumNullEnum) {
	o.UserType = v
}

// GetTenantId returns the TenantId field value
func (o *Admin) GetTenantId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *Admin) GetTenantIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *Admin) SetTenantId(v int32) {
	o.TenantId = v
}

// GetPassword returns the Password field value
func (o *Admin) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *Admin) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *Admin) SetPassword(v string) {
	o.Password = v
}

// GetConfirmPassword returns the ConfirmPassword field value
func (o *Admin) GetConfirmPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfirmPassword
}

// GetConfirmPasswordOk returns a tuple with the ConfirmPassword field value
// and a boolean to check if the value has been set.
func (o *Admin) GetConfirmPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConfirmPassword, true
}

// SetConfirmPassword sets field value
func (o *Admin) SetConfirmPassword(v string) {
	o.ConfirmPassword = v
}

// GetToken returns the Token field value
func (o *Admin) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *Admin) GetTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *Admin) SetToken(v string) {
	o.Token = v
}

func (o Admin) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	if o.UserType != nil {
		toSerialize["user_type"] = o.UserType
	}
	if true {
		toSerialize["tenant_id"] = o.TenantId
	}
	if true {
		toSerialize["password"] = o.Password
	}
	if true {
		toSerialize["confirm_password"] = o.ConfirmPassword
	}
	if true {
		toSerialize["token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullableAdmin struct {
	value *Admin
	isSet bool
}

func (v NullableAdmin) Get() *Admin {
	return v.value
}

func (v *NullableAdmin) Set(val *Admin) {
	v.value = val
	v.isSet = true
}

func (v NullableAdmin) IsSet() bool {
	return v.isSet
}

func (v *NullableAdmin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdmin(val *Admin) *NullableAdmin {
	return &NullableAdmin{value: val, isSet: true}
}

func (v NullableAdmin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdmin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

