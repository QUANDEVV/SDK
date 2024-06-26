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

// VirtualAssistantDocuments struct for VirtualAssistantDocuments
type VirtualAssistantDocuments struct {
	Id int32 `json:"id"`
	// Metadata
	Metadata interface{} `json:"metadata,omitempty"`
	CreatedAt NullableString `json:"created_at"`
	DateTimeCreatedAt NullableString `json:"date_time_created_at"`
	// The timestamp of the chat.
	Timestamp NullableTime `json:"timestamp"`
	UpdatedAt NullableTime `json:"updated_at"`
	// PDF files associated with the virtual assistant.
	Pdf string `json:"pdf"`
	// The virtual assistant ID.
	VirtualAssistantId int32 `json:"virtual_assistant_id"`
}

// NewVirtualAssistantDocuments instantiates a new VirtualAssistantDocuments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualAssistantDocuments(id int32, createdAt NullableString, dateTimeCreatedAt NullableString, timestamp NullableTime, updatedAt NullableTime, pdf string, virtualAssistantId int32) *VirtualAssistantDocuments {
	this := VirtualAssistantDocuments{}
	this.Id = id
	this.CreatedAt = createdAt
	this.DateTimeCreatedAt = dateTimeCreatedAt
	this.Timestamp = timestamp
	this.UpdatedAt = updatedAt
	this.Pdf = pdf
	this.VirtualAssistantId = virtualAssistantId
	return &this
}

// NewVirtualAssistantDocumentsWithDefaults instantiates a new VirtualAssistantDocuments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualAssistantDocumentsWithDefaults() *VirtualAssistantDocuments {
	this := VirtualAssistantDocuments{}
	return &this
}

// GetId returns the Id field value
func (o *VirtualAssistantDocuments) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VirtualAssistantDocuments) GetIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VirtualAssistantDocuments) SetId(v int32) {
	o.Id = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualAssistantDocuments) GetMetadata() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualAssistantDocuments) GetMetadataOk() (*interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *VirtualAssistantDocuments) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *VirtualAssistantDocuments) SetMetadata(v interface{}) {
	o.Metadata = v
}

// GetCreatedAt returns the CreatedAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *VirtualAssistantDocuments) GetCreatedAt() string {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualAssistantDocuments) GetCreatedAtOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// SetCreatedAt sets field value
func (o *VirtualAssistantDocuments) SetCreatedAt(v string) {
	o.CreatedAt.Set(&v)
}

// GetDateTimeCreatedAt returns the DateTimeCreatedAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *VirtualAssistantDocuments) GetDateTimeCreatedAt() string {
	if o == nil || o.DateTimeCreatedAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.DateTimeCreatedAt.Get()
}

// GetDateTimeCreatedAtOk returns a tuple with the DateTimeCreatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualAssistantDocuments) GetDateTimeCreatedAtOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DateTimeCreatedAt.Get(), o.DateTimeCreatedAt.IsSet()
}

// SetDateTimeCreatedAt sets field value
func (o *VirtualAssistantDocuments) SetDateTimeCreatedAt(v string) {
	o.DateTimeCreatedAt.Set(&v)
}

// GetTimestamp returns the Timestamp field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *VirtualAssistantDocuments) GetTimestamp() time.Time {
	if o == nil || o.Timestamp.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Timestamp.Get()
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualAssistantDocuments) GetTimestampOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Timestamp.Get(), o.Timestamp.IsSet()
}

// SetTimestamp sets field value
func (o *VirtualAssistantDocuments) SetTimestamp(v time.Time) {
	o.Timestamp.Set(&v)
}

// GetUpdatedAt returns the UpdatedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *VirtualAssistantDocuments) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualAssistantDocuments) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// SetUpdatedAt sets field value
func (o *VirtualAssistantDocuments) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}

// GetPdf returns the Pdf field value
func (o *VirtualAssistantDocuments) GetPdf() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pdf
}

// GetPdfOk returns a tuple with the Pdf field value
// and a boolean to check if the value has been set.
func (o *VirtualAssistantDocuments) GetPdfOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Pdf, true
}

// SetPdf sets field value
func (o *VirtualAssistantDocuments) SetPdf(v string) {
	o.Pdf = v
}

// GetVirtualAssistantId returns the VirtualAssistantId field value
func (o *VirtualAssistantDocuments) GetVirtualAssistantId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VirtualAssistantId
}

// GetVirtualAssistantIdOk returns a tuple with the VirtualAssistantId field value
// and a boolean to check if the value has been set.
func (o *VirtualAssistantDocuments) GetVirtualAssistantIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VirtualAssistantId, true
}

// SetVirtualAssistantId sets field value
func (o *VirtualAssistantDocuments) SetVirtualAssistantId(v int32) {
	o.VirtualAssistantId = v
}

func (o VirtualAssistantDocuments) MarshalJSON() ([]byte, error) {
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
		toSerialize["pdf"] = o.Pdf
	}
	if true {
		toSerialize["virtual_assistant_id"] = o.VirtualAssistantId
	}
	return json.Marshal(toSerialize)
}

type NullableVirtualAssistantDocuments struct {
	value *VirtualAssistantDocuments
	isSet bool
}

func (v NullableVirtualAssistantDocuments) Get() *VirtualAssistantDocuments {
	return v.value
}

func (v *NullableVirtualAssistantDocuments) Set(val *VirtualAssistantDocuments) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualAssistantDocuments) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualAssistantDocuments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualAssistantDocuments(val *VirtualAssistantDocuments) *NullableVirtualAssistantDocuments {
	return &NullableVirtualAssistantDocuments{value: val, isSet: true}
}

func (v NullableVirtualAssistantDocuments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualAssistantDocuments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


