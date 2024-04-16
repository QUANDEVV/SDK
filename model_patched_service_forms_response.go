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

// PatchedServiceFormsResponse struct for PatchedServiceFormsResponse
type PatchedServiceFormsResponse struct {
	CreatedAt NullableString `json:"created_at,omitempty"`
	ServiceFormResponseId *int32 `json:"service_form_response_id,omitempty"`
	ServiceForm *ServiceForms `json:"service_form,omitempty"`
	Client *ClientServicesInfo `json:"client,omitempty"`
	UnauthenticatedUser *AnonymousUser `json:"unauthenticated_user,omitempty"`
	FormResponse interface{} `json:"form_response,omitempty"`
}

// NewPatchedServiceFormsResponse instantiates a new PatchedServiceFormsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedServiceFormsResponse() *PatchedServiceFormsResponse {
	this := PatchedServiceFormsResponse{}
	return &this
}

// NewPatchedServiceFormsResponseWithDefaults instantiates a new PatchedServiceFormsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedServiceFormsResponseWithDefaults() *PatchedServiceFormsResponse {
	this := PatchedServiceFormsResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedServiceFormsResponse) GetCreatedAt() string {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedServiceFormsResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PatchedServiceFormsResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableString and assigns it to the CreatedAt field.
func (o *PatchedServiceFormsResponse) SetCreatedAt(v string) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *PatchedServiceFormsResponse) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *PatchedServiceFormsResponse) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetServiceFormResponseId returns the ServiceFormResponseId field value if set, zero value otherwise.
func (o *PatchedServiceFormsResponse) GetServiceFormResponseId() int32 {
	if o == nil || o.ServiceFormResponseId == nil {
		var ret int32
		return ret
	}
	return *o.ServiceFormResponseId
}

// GetServiceFormResponseIdOk returns a tuple with the ServiceFormResponseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedServiceFormsResponse) GetServiceFormResponseIdOk() (*int32, bool) {
	if o == nil || o.ServiceFormResponseId == nil {
		return nil, false
	}
	return o.ServiceFormResponseId, true
}

// HasServiceFormResponseId returns a boolean if a field has been set.
func (o *PatchedServiceFormsResponse) HasServiceFormResponseId() bool {
	if o != nil && o.ServiceFormResponseId != nil {
		return true
	}

	return false
}

// SetServiceFormResponseId gets a reference to the given int32 and assigns it to the ServiceFormResponseId field.
func (o *PatchedServiceFormsResponse) SetServiceFormResponseId(v int32) {
	o.ServiceFormResponseId = &v
}

// GetServiceForm returns the ServiceForm field value if set, zero value otherwise.
func (o *PatchedServiceFormsResponse) GetServiceForm() ServiceForms {
	if o == nil || o.ServiceForm == nil {
		var ret ServiceForms
		return ret
	}
	return *o.ServiceForm
}

// GetServiceFormOk returns a tuple with the ServiceForm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedServiceFormsResponse) GetServiceFormOk() (*ServiceForms, bool) {
	if o == nil || o.ServiceForm == nil {
		return nil, false
	}
	return o.ServiceForm, true
}

// HasServiceForm returns a boolean if a field has been set.
func (o *PatchedServiceFormsResponse) HasServiceForm() bool {
	if o != nil && o.ServiceForm != nil {
		return true
	}

	return false
}

// SetServiceForm gets a reference to the given ServiceForms and assigns it to the ServiceForm field.
func (o *PatchedServiceFormsResponse) SetServiceForm(v ServiceForms) {
	o.ServiceForm = &v
}

// GetClient returns the Client field value if set, zero value otherwise.
func (o *PatchedServiceFormsResponse) GetClient() ClientServicesInfo {
	if o == nil || o.Client == nil {
		var ret ClientServicesInfo
		return ret
	}
	return *o.Client
}

// GetClientOk returns a tuple with the Client field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedServiceFormsResponse) GetClientOk() (*ClientServicesInfo, bool) {
	if o == nil || o.Client == nil {
		return nil, false
	}
	return o.Client, true
}

// HasClient returns a boolean if a field has been set.
func (o *PatchedServiceFormsResponse) HasClient() bool {
	if o != nil && o.Client != nil {
		return true
	}

	return false
}

// SetClient gets a reference to the given ClientServicesInfo and assigns it to the Client field.
func (o *PatchedServiceFormsResponse) SetClient(v ClientServicesInfo) {
	o.Client = &v
}

// GetUnauthenticatedUser returns the UnauthenticatedUser field value if set, zero value otherwise.
func (o *PatchedServiceFormsResponse) GetUnauthenticatedUser() AnonymousUser {
	if o == nil || o.UnauthenticatedUser == nil {
		var ret AnonymousUser
		return ret
	}
	return *o.UnauthenticatedUser
}

// GetUnauthenticatedUserOk returns a tuple with the UnauthenticatedUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedServiceFormsResponse) GetUnauthenticatedUserOk() (*AnonymousUser, bool) {
	if o == nil || o.UnauthenticatedUser == nil {
		return nil, false
	}
	return o.UnauthenticatedUser, true
}

// HasUnauthenticatedUser returns a boolean if a field has been set.
func (o *PatchedServiceFormsResponse) HasUnauthenticatedUser() bool {
	if o != nil && o.UnauthenticatedUser != nil {
		return true
	}

	return false
}

// SetUnauthenticatedUser gets a reference to the given AnonymousUser and assigns it to the UnauthenticatedUser field.
func (o *PatchedServiceFormsResponse) SetUnauthenticatedUser(v AnonymousUser) {
	o.UnauthenticatedUser = &v
}

// GetFormResponse returns the FormResponse field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedServiceFormsResponse) GetFormResponse() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.FormResponse
}

// GetFormResponseOk returns a tuple with the FormResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedServiceFormsResponse) GetFormResponseOk() (*interface{}, bool) {
	if o == nil || o.FormResponse == nil {
		return nil, false
	}
	return &o.FormResponse, true
}

// HasFormResponse returns a boolean if a field has been set.
func (o *PatchedServiceFormsResponse) HasFormResponse() bool {
	if o != nil && o.FormResponse != nil {
		return true
	}

	return false
}

// SetFormResponse gets a reference to the given interface{} and assigns it to the FormResponse field.
func (o *PatchedServiceFormsResponse) SetFormResponse(v interface{}) {
	o.FormResponse = v
}

func (o PatchedServiceFormsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.ServiceFormResponseId != nil {
		toSerialize["service_form_response_id"] = o.ServiceFormResponseId
	}
	if o.ServiceForm != nil {
		toSerialize["service_form"] = o.ServiceForm
	}
	if o.Client != nil {
		toSerialize["client"] = o.Client
	}
	if o.UnauthenticatedUser != nil {
		toSerialize["unauthenticated_user"] = o.UnauthenticatedUser
	}
	if o.FormResponse != nil {
		toSerialize["form_response"] = o.FormResponse
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedServiceFormsResponse struct {
	value *PatchedServiceFormsResponse
	isSet bool
}

func (v NullablePatchedServiceFormsResponse) Get() *PatchedServiceFormsResponse {
	return v.value
}

func (v *NullablePatchedServiceFormsResponse) Set(val *PatchedServiceFormsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedServiceFormsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedServiceFormsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedServiceFormsResponse(val *PatchedServiceFormsResponse) *NullablePatchedServiceFormsResponse {
	return &NullablePatchedServiceFormsResponse{value: val, isSet: true}
}

func (v NullablePatchedServiceFormsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedServiceFormsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

