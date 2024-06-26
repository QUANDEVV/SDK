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

// ChannelEnum * `Mobile` - Mobile * `Website` - Website
type ChannelEnum string

// List of ChannelEnum
const (
	MOBILE ChannelEnum = "Mobile"
	WEBSITE ChannelEnum = "Website"
)

// All allowed values of ChannelEnum enum
var AllowedChannelEnumEnumValues = []ChannelEnum{
	"Mobile",
	"Website",
}

func (v *ChannelEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ChannelEnum(value)
	for _, existing := range AllowedChannelEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ChannelEnum", value)
}

// NewChannelEnumFromValue returns a pointer to a valid ChannelEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewChannelEnumFromValue(v string) (*ChannelEnum, error) {
	ev := ChannelEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ChannelEnum: valid values are %v", v, AllowedChannelEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ChannelEnum) IsValid() bool {
	for _, existing := range AllowedChannelEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ChannelEnum value
func (v ChannelEnum) Ptr() *ChannelEnum {
	return &v
}

type NullableChannelEnum struct {
	value *ChannelEnum
	isSet bool
}

func (v NullableChannelEnum) Get() *ChannelEnum {
	return v.value
}

func (v *NullableChannelEnum) Set(val *ChannelEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelEnum(val *ChannelEnum) *NullableChannelEnum {
	return &NullableChannelEnum{value: val, isSet: true}
}

func (v NullableChannelEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

