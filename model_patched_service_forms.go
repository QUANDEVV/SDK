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

// PatchedServiceForms struct for PatchedServiceForms
type PatchedServiceForms struct {
	ServiceFormId *int32 `json:"service_form_id,omitempty"`
	Service *Service `json:"service,omitempty"`
	NameOfForm *string `json:"name_of_form,omitempty"`
	Form interface{} `json:"form,omitempty"`
}

// NewPatchedServiceForms instantiates a new PatchedServiceForms object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedServiceForms() *PatchedServiceForms {
	this := PatchedServiceForms{}
	return &this
}

// NewPatchedServiceFormsWithDefaults instantiates a new PatchedServiceForms object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedServiceFormsWithDefaults() *PatchedServiceForms {
	this := PatchedServiceForms{}
	return &this
}

// GetServiceFormId returns the ServiceFormId field value if set, zero value otherwise.
func (o *PatchedServiceForms) GetServiceFormId() int32 {
	if o == nil || o.ServiceFormId == nil {
		var ret int32
		return ret
	}
	return *o.ServiceFormId
}

// GetServiceFormIdOk returns a tuple with the ServiceFormId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedServiceForms) GetServiceFormIdOk() (*int32, bool) {
	if o == nil || o.ServiceFormId == nil {
		return nil, false
	}
	return o.ServiceFormId, true
}

// HasServiceFormId returns a boolean if a field has been set.
func (o *PatchedServiceForms) HasServiceFormId() bool {
	if o != nil && o.ServiceFormId != nil {
		return true
	}

	return false
}

// SetServiceFormId gets a reference to the given int32 and assigns it to the ServiceFormId field.
func (o *PatchedServiceForms) SetServiceFormId(v int32) {
	o.ServiceFormId = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *PatchedServiceForms) GetService() Service {
	if o == nil || o.Service == nil {
		var ret Service
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedServiceForms) GetServiceOk() (*Service, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *PatchedServiceForms) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given Service and assigns it to the Service field.
func (o *PatchedServiceForms) SetService(v Service) {
	o.Service = &v
}

// GetNameOfForm returns the NameOfForm field value if set, zero value otherwise.
func (o *PatchedServiceForms) GetNameOfForm() string {
	if o == nil || o.NameOfForm == nil {
		var ret string
		return ret
	}
	return *o.NameOfForm
}

// GetNameOfFormOk returns a tuple with the NameOfForm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedServiceForms) GetNameOfFormOk() (*string, bool) {
	if o == nil || o.NameOfForm == nil {
		return nil, false
	}
	return o.NameOfForm, true
}

// HasNameOfForm returns a boolean if a field has been set.
func (o *PatchedServiceForms) HasNameOfForm() bool {
	if o != nil && o.NameOfForm != nil {
		return true
	}

	return false
}

// SetNameOfForm gets a reference to the given string and assigns it to the NameOfForm field.
func (o *PatchedServiceForms) SetNameOfForm(v string) {
	o.NameOfForm = &v
}

// GetForm returns the Form field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedServiceForms) GetForm() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.Form
}

// GetFormOk returns a tuple with the Form field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedServiceForms) GetFormOk() (*interface{}, bool) {
	if o == nil || o.Form == nil {
		return nil, false
	}
	return &o.Form, true
}

// HasForm returns a boolean if a field has been set.
func (o *PatchedServiceForms) HasForm() bool {
	if o != nil && o.Form != nil {
		return true
	}

	return false
}

// SetForm gets a reference to the given interface{} and assigns it to the Form field.
func (o *PatchedServiceForms) SetForm(v interface{}) {
	o.Form = v
}

func (o PatchedServiceForms) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ServiceFormId != nil {
		toSerialize["service_form_id"] = o.ServiceFormId
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	if o.NameOfForm != nil {
		toSerialize["name_of_form"] = o.NameOfForm
	}
	if o.Form != nil {
		toSerialize["form"] = o.Form
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedServiceForms struct {
	value *PatchedServiceForms
	isSet bool
}

func (v NullablePatchedServiceForms) Get() *PatchedServiceForms {
	return v.value
}

func (v *NullablePatchedServiceForms) Set(val *PatchedServiceForms) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedServiceForms) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedServiceForms) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedServiceForms(val *PatchedServiceForms) *NullablePatchedServiceForms {
	return &NullablePatchedServiceForms{value: val, isSet: true}
}

func (v NullablePatchedServiceForms) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedServiceForms) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


