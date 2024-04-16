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

// SurveyTypeEnum * `open_ended` - open_ended * `close_ended` - close_ended
type SurveyTypeEnum string

// List of SurveyTypeEnum
const (
	OPEN_ENDED SurveyTypeEnum = "open_ended"
	CLOSE_ENDED SurveyTypeEnum = "close_ended"
)

// All allowed values of SurveyTypeEnum enum
var AllowedSurveyTypeEnumEnumValues = []SurveyTypeEnum{
	"open_ended",
	"close_ended",
}

func (v *SurveyTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SurveyTypeEnum(value)
	for _, existing := range AllowedSurveyTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SurveyTypeEnum", value)
}

// NewSurveyTypeEnumFromValue returns a pointer to a valid SurveyTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSurveyTypeEnumFromValue(v string) (*SurveyTypeEnum, error) {
	ev := SurveyTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SurveyTypeEnum: valid values are %v", v, AllowedSurveyTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SurveyTypeEnum) IsValid() bool {
	for _, existing := range AllowedSurveyTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SurveyTypeEnum value
func (v SurveyTypeEnum) Ptr() *SurveyTypeEnum {
	return &v
}

type NullableSurveyTypeEnum struct {
	value *SurveyTypeEnum
	isSet bool
}

func (v NullableSurveyTypeEnum) Get() *SurveyTypeEnum {
	return v.value
}

func (v *NullableSurveyTypeEnum) Set(val *SurveyTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableSurveyTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableSurveyTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSurveyTypeEnum(val *SurveyTypeEnum) *NullableSurveyTypeEnum {
	return &NullableSurveyTypeEnum{value: val, isSet: true}
}

func (v NullableSurveyTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSurveyTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

