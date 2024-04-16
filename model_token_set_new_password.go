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

// TokenSetNewPassword struct for TokenSetNewPassword
type TokenSetNewPassword struct {
	Password string `json:"password"`
	Token int32 `json:"token"`
}

// NewTokenSetNewPassword instantiates a new TokenSetNewPassword object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenSetNewPassword(password string, token int32) *TokenSetNewPassword {
	this := TokenSetNewPassword{}
	this.Password = password
	this.Token = token
	return &this
}

// NewTokenSetNewPasswordWithDefaults instantiates a new TokenSetNewPassword object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenSetNewPasswordWithDefaults() *TokenSetNewPassword {
	this := TokenSetNewPassword{}
	return &this
}

// GetPassword returns the Password field value
func (o *TokenSetNewPassword) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *TokenSetNewPassword) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *TokenSetNewPassword) SetPassword(v string) {
	o.Password = v
}

// GetToken returns the Token field value
func (o *TokenSetNewPassword) GetToken() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *TokenSetNewPassword) GetTokenOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *TokenSetNewPassword) SetToken(v int32) {
	o.Token = v
}

func (o TokenSetNewPassword) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["password"] = o.Password
	}
	if true {
		toSerialize["token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullableTokenSetNewPassword struct {
	value *TokenSetNewPassword
	isSet bool
}

func (v NullableTokenSetNewPassword) Get() *TokenSetNewPassword {
	return v.value
}

func (v *NullableTokenSetNewPassword) Set(val *TokenSetNewPassword) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenSetNewPassword) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenSetNewPassword) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenSetNewPassword(val *TokenSetNewPassword) *NullableTokenSetNewPassword {
	return &NullableTokenSetNewPassword{value: val, isSet: true}
}

func (v NullableTokenSetNewPassword) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenSetNewPassword) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

