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

// ClientLikes struct for ClientLikes
type ClientLikes struct {
	Id int32 `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

// NewClientLikes instantiates a new ClientLikes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientLikes(id int32, username string, email string, firstName string, lastName string) *ClientLikes {
	this := ClientLikes{}
	this.Id = id
	this.Username = username
	this.Email = email
	this.FirstName = firstName
	this.LastName = lastName
	return &this
}

// NewClientLikesWithDefaults instantiates a new ClientLikes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientLikesWithDefaults() *ClientLikes {
	this := ClientLikes{}
	return &this
}

// GetId returns the Id field value
func (o *ClientLikes) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ClientLikes) GetIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ClientLikes) SetId(v int32) {
	o.Id = v
}

// GetUsername returns the Username field value
func (o *ClientLikes) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *ClientLikes) GetUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *ClientLikes) SetUsername(v string) {
	o.Username = v
}

// GetEmail returns the Email field value
func (o *ClientLikes) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *ClientLikes) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *ClientLikes) SetEmail(v string) {
	o.Email = v
}

// GetFirstName returns the FirstName field value
func (o *ClientLikes) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *ClientLikes) GetFirstNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *ClientLikes) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *ClientLikes) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *ClientLikes) GetLastNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *ClientLikes) SetLastName(v string) {
	o.LastName = v
}

func (o ClientLikes) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableClientLikes struct {
	value *ClientLikes
	isSet bool
}

func (v NullableClientLikes) Get() *ClientLikes {
	return v.value
}

func (v *NullableClientLikes) Set(val *ClientLikes) {
	v.value = val
	v.isSet = true
}

func (v NullableClientLikes) IsSet() bool {
	return v.isSet
}

func (v *NullableClientLikes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientLikes(val *ClientLikes) *NullableClientLikes {
	return &NullableClientLikes{value: val, isSet: true}
}

func (v NullableClientLikes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientLikes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


