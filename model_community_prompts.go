/*
proxima-core-engine

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// CommunityPrompts struct for CommunityPrompts
type CommunityPrompts struct {
	Id int32 `json:"id"`
	// Metadata
	Metadata interface{} `json:"metadata,omitempty"`
	CreatedAt NullableString `json:"created_at"`
	DateTimeCreatedAt NullableString `json:"date_time_created_at"`
	// The timestamp of the chat.
	Timestamp NullableTime `json:"timestamp"`
	UpdatedAt NullableTime `json:"updated_at"`
	CommunityDescription string `json:"community_description"`
	// Linking it to the knowledgebase
	Knowledgebase int32 `json:"knowledgebase"`
}

// NewCommunityPrompts instantiates a new CommunityPrompts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommunityPrompts(id int32, createdAt NullableString, dateTimeCreatedAt NullableString, timestamp NullableTime, updatedAt NullableTime, communityDescription string, knowledgebase int32) *CommunityPrompts {
	this := CommunityPrompts{}
	this.Id = id
	this.CreatedAt = createdAt
	this.DateTimeCreatedAt = dateTimeCreatedAt
	this.Timestamp = timestamp
	this.UpdatedAt = updatedAt
	this.CommunityDescription = communityDescription
	this.Knowledgebase = knowledgebase
	return &this
}

// NewCommunityPromptsWithDefaults instantiates a new CommunityPrompts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommunityPromptsWithDefaults() *CommunityPrompts {
	this := CommunityPrompts{}
	return &this
}

// GetId returns the Id field value
func (o *CommunityPrompts) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CommunityPrompts) GetIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CommunityPrompts) SetId(v int32) {
	o.Id = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommunityPrompts) GetMetadata() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommunityPrompts) GetMetadataOk() (*interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CommunityPrompts) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *CommunityPrompts) SetMetadata(v interface{}) {
	o.Metadata = v
}

// GetCreatedAt returns the CreatedAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CommunityPrompts) GetCreatedAt() string {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommunityPrompts) GetCreatedAtOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// SetCreatedAt sets field value
func (o *CommunityPrompts) SetCreatedAt(v string) {
	o.CreatedAt.Set(&v)
}

// GetDateTimeCreatedAt returns the DateTimeCreatedAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CommunityPrompts) GetDateTimeCreatedAt() string {
	if o == nil || o.DateTimeCreatedAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.DateTimeCreatedAt.Get()
}

// GetDateTimeCreatedAtOk returns a tuple with the DateTimeCreatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommunityPrompts) GetDateTimeCreatedAtOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DateTimeCreatedAt.Get(), o.DateTimeCreatedAt.IsSet()
}

// SetDateTimeCreatedAt sets field value
func (o *CommunityPrompts) SetDateTimeCreatedAt(v string) {
	o.DateTimeCreatedAt.Set(&v)
}

// GetTimestamp returns the Timestamp field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *CommunityPrompts) GetTimestamp() time.Time {
	if o == nil || o.Timestamp.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Timestamp.Get()
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommunityPrompts) GetTimestampOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Timestamp.Get(), o.Timestamp.IsSet()
}

// SetTimestamp sets field value
func (o *CommunityPrompts) SetTimestamp(v time.Time) {
	o.Timestamp.Set(&v)
}

// GetUpdatedAt returns the UpdatedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *CommunityPrompts) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommunityPrompts) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// SetUpdatedAt sets field value
func (o *CommunityPrompts) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}

// GetCommunityDescription returns the CommunityDescription field value
func (o *CommunityPrompts) GetCommunityDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CommunityDescription
}

// GetCommunityDescriptionOk returns a tuple with the CommunityDescription field value
// and a boolean to check if the value has been set.
func (o *CommunityPrompts) GetCommunityDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CommunityDescription, true
}

// SetCommunityDescription sets field value
func (o *CommunityPrompts) SetCommunityDescription(v string) {
	o.CommunityDescription = v
}

// GetKnowledgebase returns the Knowledgebase field value
func (o *CommunityPrompts) GetKnowledgebase() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Knowledgebase
}

// GetKnowledgebaseOk returns a tuple with the Knowledgebase field value
// and a boolean to check if the value has been set.
func (o *CommunityPrompts) GetKnowledgebaseOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Knowledgebase, true
}

// SetKnowledgebase sets field value
func (o *CommunityPrompts) SetKnowledgebase(v int32) {
	o.Knowledgebase = v
}

func (o CommunityPrompts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if true {
		toSerialize["date_time_created_at"] = o.DateTimeCreatedAt.Get()
	}
	if true {
		toSerialize["timestamp"] = o.Timestamp.Get()
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}
	if true {
		toSerialize["community_description"] = o.CommunityDescription
	}
	if true {
		toSerialize["knowledgebase"] = o.Knowledgebase
	}
	return json.Marshal(toSerialize)
}

type NullableCommunityPrompts struct {
	value *CommunityPrompts
	isSet bool
}

func (v NullableCommunityPrompts) Get() *CommunityPrompts {
	return v.value
}

func (v *NullableCommunityPrompts) Set(val *CommunityPrompts) {
	v.value = val
	v.isSet = true
}

func (v NullableCommunityPrompts) IsSet() bool {
	return v.isSet
}

func (v *NullableCommunityPrompts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommunityPrompts(val *CommunityPrompts) *NullableCommunityPrompts {
	return &NullableCommunityPrompts{value: val, isSet: true}
}

func (v NullableCommunityPrompts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommunityPrompts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


