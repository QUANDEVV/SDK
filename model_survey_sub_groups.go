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

// SurveySubGroups struct for SurveySubGroups
type SurveySubGroups struct {
	// The survey report id
	SurveySubgroupsId int32 `json:"survey_subgroups_id"`
	SurveyId *string `json:"survey_id,omitempty"`
	// The survey subgroup name
	SubgroupName string `json:"subgroup_name"`
	// The survey subgroup description
	SubgroupDescription string `json:"subgroup_description"`
	SubgroupClients []TargetAudience `json:"subgroup_clients"`
}

// NewSurveySubGroups instantiates a new SurveySubGroups object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSurveySubGroups(surveySubgroupsId int32, subgroupName string, subgroupDescription string, subgroupClients []TargetAudience) *SurveySubGroups {
	this := SurveySubGroups{}
	this.SurveySubgroupsId = surveySubgroupsId
	this.SubgroupName = subgroupName
	this.SubgroupDescription = subgroupDescription
	this.SubgroupClients = subgroupClients
	return &this
}

// NewSurveySubGroupsWithDefaults instantiates a new SurveySubGroups object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSurveySubGroupsWithDefaults() *SurveySubGroups {
	this := SurveySubGroups{}
	return &this
}

// GetSurveySubgroupsId returns the SurveySubgroupsId field value
func (o *SurveySubGroups) GetSurveySubgroupsId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SurveySubgroupsId
}

// GetSurveySubgroupsIdOk returns a tuple with the SurveySubgroupsId field value
// and a boolean to check if the value has been set.
func (o *SurveySubGroups) GetSurveySubgroupsIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SurveySubgroupsId, true
}

// SetSurveySubgroupsId sets field value
func (o *SurveySubGroups) SetSurveySubgroupsId(v int32) {
	o.SurveySubgroupsId = v
}

// GetSurveyId returns the SurveyId field value if set, zero value otherwise.
func (o *SurveySubGroups) GetSurveyId() string {
	if o == nil || o.SurveyId == nil {
		var ret string
		return ret
	}
	return *o.SurveyId
}

// GetSurveyIdOk returns a tuple with the SurveyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SurveySubGroups) GetSurveyIdOk() (*string, bool) {
	if o == nil || o.SurveyId == nil {
		return nil, false
	}
	return o.SurveyId, true
}

// HasSurveyId returns a boolean if a field has been set.
func (o *SurveySubGroups) HasSurveyId() bool {
	if o != nil && o.SurveyId != nil {
		return true
	}

	return false
}

// SetSurveyId gets a reference to the given string and assigns it to the SurveyId field.
func (o *SurveySubGroups) SetSurveyId(v string) {
	o.SurveyId = &v
}

// GetSubgroupName returns the SubgroupName field value
func (o *SurveySubGroups) GetSubgroupName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubgroupName
}

// GetSubgroupNameOk returns a tuple with the SubgroupName field value
// and a boolean to check if the value has been set.
func (o *SurveySubGroups) GetSubgroupNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SubgroupName, true
}

// SetSubgroupName sets field value
func (o *SurveySubGroups) SetSubgroupName(v string) {
	o.SubgroupName = v
}

// GetSubgroupDescription returns the SubgroupDescription field value
func (o *SurveySubGroups) GetSubgroupDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubgroupDescription
}

// GetSubgroupDescriptionOk returns a tuple with the SubgroupDescription field value
// and a boolean to check if the value has been set.
func (o *SurveySubGroups) GetSubgroupDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SubgroupDescription, true
}

// SetSubgroupDescription sets field value
func (o *SurveySubGroups) SetSubgroupDescription(v string) {
	o.SubgroupDescription = v
}

// GetSubgroupClients returns the SubgroupClients field value
func (o *SurveySubGroups) GetSubgroupClients() []TargetAudience {
	if o == nil {
		var ret []TargetAudience
		return ret
	}

	return o.SubgroupClients
}

// GetSubgroupClientsOk returns a tuple with the SubgroupClients field value
// and a boolean to check if the value has been set.
func (o *SurveySubGroups) GetSubgroupClientsOk() (*[]TargetAudience, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SubgroupClients, true
}

// SetSubgroupClients sets field value
func (o *SurveySubGroups) SetSubgroupClients(v []TargetAudience) {
	o.SubgroupClients = v
}

func (o SurveySubGroups) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["survey_subgroups_id"] = o.SurveySubgroupsId
	}
	if o.SurveyId != nil {
		toSerialize["survey_id"] = o.SurveyId
	}
	if true {
		toSerialize["subgroup_name"] = o.SubgroupName
	}
	if true {
		toSerialize["subgroup_description"] = o.SubgroupDescription
	}
	if true {
		toSerialize["subgroup_clients"] = o.SubgroupClients
	}
	return json.Marshal(toSerialize)
}

type NullableSurveySubGroups struct {
	value *SurveySubGroups
	isSet bool
}

func (v NullableSurveySubGroups) Get() *SurveySubGroups {
	return v.value
}

func (v *NullableSurveySubGroups) Set(val *SurveySubGroups) {
	v.value = val
	v.isSet = true
}

func (v NullableSurveySubGroups) IsSet() bool {
	return v.isSet
}

func (v *NullableSurveySubGroups) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSurveySubGroups(val *SurveySubGroups) *NullableSurveySubGroups {
	return &NullableSurveySubGroups{value: val, isSet: true}
}

func (v NullableSurveySubGroups) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSurveySubGroups) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

