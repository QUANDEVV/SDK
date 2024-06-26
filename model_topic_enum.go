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

// TopicEnum * `Mobile` - Mobile * `Website` - Website
type TopicEnum string

// List of TopicEnum
const (
	MOBILE TopicEnum = "Mobile"
	WEBSITE TopicEnum = "Website"
)

// All allowed values of TopicEnum enum
var AllowedTopicEnumEnumValues = []TopicEnum{
	"Mobile",
	"Website",
}

func (v *TopicEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TopicEnum(value)
	for _, existing := range AllowedTopicEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TopicEnum", value)
}

// NewTopicEnumFromValue returns a pointer to a valid TopicEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTopicEnumFromValue(v string) (*TopicEnum, error) {
	ev := TopicEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TopicEnum: valid values are %v", v, AllowedTopicEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TopicEnum) IsValid() bool {
	for _, existing := range AllowedTopicEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TopicEnum value
func (v TopicEnum) Ptr() *TopicEnum {
	return &v
}

type NullableTopicEnum struct {
	value *TopicEnum
	isSet bool
}

func (v NullableTopicEnum) Get() *TopicEnum {
	return v.value
}

func (v *NullableTopicEnum) Set(val *TopicEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableTopicEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableTopicEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTopicEnum(val *TopicEnum) *NullableTopicEnum {
	return &NullableTopicEnum{value: val, isSet: true}
}

func (v NullableTopicEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTopicEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

