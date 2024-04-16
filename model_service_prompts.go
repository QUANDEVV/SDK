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

// ServicePrompts struct for ServicePrompts
type ServicePrompts struct {
	Id int32 `json:"id"`
	// Metadata
	Metadata interface{} `json:"metadata,omitempty"`
	CreatedAt NullableString `json:"created_at"`
	DateTimeCreatedAt NullableString `json:"date_time_created_at"`
	// The timestamp of the chat.
	Timestamp NullableTime `json:"timestamp"`
	UpdatedAt NullableTime `json:"updated_at"`
	ServiceDescription string `json:"service_description"`
	// To tie services knowledge base to it
	Service int32 `json:"service"`
	// Linking it to the knowledgebase
	Knowledgebase int32 `json:"knowledgebase"`
}

// NewServicePrompts instantiates a new ServicePrompts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServicePrompts(id int32, createdAt NullableString, dateTimeCreatedAt NullableString, timestamp NullableTime, updatedAt NullableTime, serviceDescription string, service int32, knowledgebase int32) *ServicePrompts {
	this := ServicePrompts{}
	this.Id = id
	this.CreatedAt = createdAt
	this.DateTimeCreatedAt = dateTimeCreatedAt
	this.Timestamp = timestamp
	this.UpdatedAt = updatedAt
	this.ServiceDescription = serviceDescription
	this.Service = service
	this.Knowledgebase = knowledgebase
	return &this
}

// NewServicePromptsWithDefaults instantiates a new ServicePrompts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicePromptsWithDefaults() *ServicePrompts {
	this := ServicePrompts{}
	return &this
}

// GetId returns the Id field value
func (o *ServicePrompts) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ServicePrompts) GetIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ServicePrompts) SetId(v int32) {
	o.Id = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServicePrompts) GetMetadata() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServicePrompts) GetMetadataOk() (*interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ServicePrompts) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *ServicePrompts) SetMetadata(v interface{}) {
	o.Metadata = v
}

// GetCreatedAt returns the CreatedAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ServicePrompts) GetCreatedAt() string {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServicePrompts) GetCreatedAtOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// SetCreatedAt sets field value
func (o *ServicePrompts) SetCreatedAt(v string) {
	o.CreatedAt.Set(&v)
}

// GetDateTimeCreatedAt returns the DateTimeCreatedAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ServicePrompts) GetDateTimeCreatedAt() string {
	if o == nil || o.DateTimeCreatedAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.DateTimeCreatedAt.Get()
}

// GetDateTimeCreatedAtOk returns a tuple with the DateTimeCreatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServicePrompts) GetDateTimeCreatedAtOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DateTimeCreatedAt.Get(), o.DateTimeCreatedAt.IsSet()
}

// SetDateTimeCreatedAt sets field value
func (o *ServicePrompts) SetDateTimeCreatedAt(v string) {
	o.DateTimeCreatedAt.Set(&v)
}

// GetTimestamp returns the Timestamp field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *ServicePrompts) GetTimestamp() time.Time {
	if o == nil || o.Timestamp.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Timestamp.Get()
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServicePrompts) GetTimestampOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Timestamp.Get(), o.Timestamp.IsSet()
}

// SetTimestamp sets field value
func (o *ServicePrompts) SetTimestamp(v time.Time) {
	o.Timestamp.Set(&v)
}

// GetUpdatedAt returns the UpdatedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *ServicePrompts) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServicePrompts) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// SetUpdatedAt sets field value
func (o *ServicePrompts) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}

// GetServiceDescription returns the ServiceDescription field value
func (o *ServicePrompts) GetServiceDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceDescription
}

// GetServiceDescriptionOk returns a tuple with the ServiceDescription field value
// and a boolean to check if the value has been set.
func (o *ServicePrompts) GetServiceDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ServiceDescription, true
}

// SetServiceDescription sets field value
func (o *ServicePrompts) SetServiceDescription(v string) {
	o.ServiceDescription = v
}

// GetService returns the Service field value
func (o *ServicePrompts) GetService() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *ServicePrompts) GetServiceOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value
func (o *ServicePrompts) SetService(v int32) {
	o.Service = v
}

// GetKnowledgebase returns the Knowledgebase field value
func (o *ServicePrompts) GetKnowledgebase() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Knowledgebase
}

// GetKnowledgebaseOk returns a tuple with the Knowledgebase field value
// and a boolean to check if the value has been set.
func (o *ServicePrompts) GetKnowledgebaseOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Knowledgebase, true
}

// SetKnowledgebase sets field value
func (o *ServicePrompts) SetKnowledgebase(v int32) {
	o.Knowledgebase = v
}

func (o ServicePrompts) MarshalJSON() ([]byte, error) {
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
		toSerialize["service_description"] = o.ServiceDescription
	}
	if true {
		toSerialize["service"] = o.Service
	}
	if true {
		toSerialize["knowledgebase"] = o.Knowledgebase
	}
	return json.Marshal(toSerialize)
}

type NullableServicePrompts struct {
	value *ServicePrompts
	isSet bool
}

func (v NullableServicePrompts) Get() *ServicePrompts {
	return v.value
}

func (v *NullableServicePrompts) Set(val *ServicePrompts) {
	v.value = val
	v.isSet = true
}

func (v NullableServicePrompts) IsSet() bool {
	return v.isSet
}

func (v *NullableServicePrompts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicePrompts(val *ServicePrompts) *NullableServicePrompts {
	return &NullableServicePrompts{value: val, isSet: true}
}

func (v NullableServicePrompts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicePrompts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

