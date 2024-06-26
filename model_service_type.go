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

// ServiceType struct for ServiceType
type ServiceType struct {
	Id int32 `json:"id"`
	// The tenant chats ID.
	TenantId int32 `json:"tenant_id"`
	Type ServiceTypeTypeEnum `json:"type"`
	Description string `json:"description"`
}

// NewServiceType instantiates a new ServiceType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceType(id int32, tenantId int32, type_ ServiceTypeTypeEnum, description string) *ServiceType {
	this := ServiceType{}
	this.Id = id
	this.TenantId = tenantId
	this.Type = type_
	this.Description = description
	return &this
}

// NewServiceTypeWithDefaults instantiates a new ServiceType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceTypeWithDefaults() *ServiceType {
	this := ServiceType{}
	return &this
}

// GetId returns the Id field value
func (o *ServiceType) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ServiceType) GetIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ServiceType) SetId(v int32) {
	o.Id = v
}

// GetTenantId returns the TenantId field value
func (o *ServiceType) GetTenantId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *ServiceType) GetTenantIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *ServiceType) SetTenantId(v int32) {
	o.TenantId = v
}

// GetType returns the Type field value
func (o *ServiceType) GetType() ServiceTypeTypeEnum {
	if o == nil {
		var ret ServiceTypeTypeEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ServiceType) GetTypeOk() (*ServiceTypeTypeEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ServiceType) SetType(v ServiceTypeTypeEnum) {
	o.Type = v
}

// GetDescription returns the Description field value
func (o *ServiceType) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ServiceType) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ServiceType) SetDescription(v string) {
	o.Description = v
}

func (o ServiceType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["tenant_id"] = o.TenantId
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableServiceType struct {
	value *ServiceType
	isSet bool
}

func (v NullableServiceType) Get() *ServiceType {
	return v.value
}

func (v *NullableServiceType) Set(val *ServiceType) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceType) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceType(val *ServiceType) *NullableServiceType {
	return &NullableServiceType{value: val, isSet: true}
}

func (v NullableServiceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


