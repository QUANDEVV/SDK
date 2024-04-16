/*
proxima-core-engine

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// UserType3daEnum * `employee` - employee * `admin` - admin
type UserType3daEnum string

// List of UserType3daEnum
const (
	EMPLOYEE UserType3daEnum = "employee"
	ADMIN UserType3daEnum = "admin"
)

// All allowed values of UserType3daEnum enum
var AllowedUserType3daEnumEnumValues = []UserType3daEnum{
	"employee",
	"admin",
}

func (v *UserType3daEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UserType3daEnum(value)
	for _, existing := range AllowedUserType3daEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UserType3daEnum", value)
}

// NewUserType3daEnumFromValue returns a pointer to a valid UserType3daEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUserType3daEnumFromValue(v string) (*UserType3daEnum, error) {
	ev := UserType3daEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UserType3daEnum: valid values are %v", v, AllowedUserType3daEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UserType3daEnum) IsValid() bool {
	for _, existing := range AllowedUserType3daEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UserType3daEnum value
func (v UserType3daEnum) Ptr() *UserType3daEnum {
	return &v
}

type NullableUserType3daEnum struct {
	value *UserType3daEnum
	isSet bool
}

func (v NullableUserType3daEnum) Get() *UserType3daEnum {
	return v.value
}

func (v *NullableUserType3daEnum) Set(val *UserType3daEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableUserType3daEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableUserType3daEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserType3daEnum(val *UserType3daEnum) *NullableUserType3daEnum {
	return &NullableUserType3daEnum{value: val, isSet: true}
}

func (v NullableUserType3daEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserType3daEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

