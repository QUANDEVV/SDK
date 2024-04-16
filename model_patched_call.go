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

// PatchedCall struct for PatchedCall
type PatchedCall struct {
	// The call id.
	CallId *int32 `json:"call_id,omitempty"`
	TenantId TenantInfo `json:"tenant_id,omitempty"`
	Duration *string `json:"duration,omitempty"`
	Success *bool `json:"success,omitempty"`
	ClientCaller *ClientInfo `json:"client_caller,omitempty"`
	// Display name of the client
	GuestCalled NullableInt32 `json:"guest_called,omitempty"`
}

// NewPatchedCall instantiates a new PatchedCall object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedCall() *PatchedCall {
	this := PatchedCall{}
	return &this
}

// NewPatchedCallWithDefaults instantiates a new PatchedCall object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedCallWithDefaults() *PatchedCall {
	this := PatchedCall{}
	return &this
}

// GetCallId returns the CallId field value if set, zero value otherwise.
func (o *PatchedCall) GetCallId() int32 {
	if o == nil || o.CallId == nil {
		var ret int32
		return ret
	}
	return *o.CallId
}

// GetCallIdOk returns a tuple with the CallId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCall) GetCallIdOk() (*int32, bool) {
	if o == nil || o.CallId == nil {
		return nil, false
	}
	return o.CallId, true
}

// HasCallId returns a boolean if a field has been set.
func (o *PatchedCall) HasCallId() bool {
	if o != nil && o.CallId != nil {
		return true
	}

	return false
}

// SetCallId gets a reference to the given int32 and assigns it to the CallId field.
func (o *PatchedCall) SetCallId(v int32) {
	o.CallId = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedCall) GetTenantId() TenantInfo {
	if o == nil  {
		var ret TenantInfo
		return ret
	}
	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedCall) GetTenantIdOk() (*TenantInfo, bool) {
	if o == nil || o.TenantId == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *PatchedCall) HasTenantId() bool {
	if o != nil && o.TenantId != nil {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given TenantInfo and assigns it to the TenantId field.
func (o *PatchedCall) SetTenantId(v TenantInfo) {
	o.TenantId = v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *PatchedCall) GetDuration() string {
	if o == nil || o.Duration == nil {
		var ret string
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCall) GetDurationOk() (*string, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *PatchedCall) HasDuration() bool {
	if o != nil && o.Duration != nil {
		return true
	}

	return false
}

// SetDuration gets a reference to the given string and assigns it to the Duration field.
func (o *PatchedCall) SetDuration(v string) {
	o.Duration = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *PatchedCall) GetSuccess() bool {
	if o == nil || o.Success == nil {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCall) GetSuccessOk() (*bool, bool) {
	if o == nil || o.Success == nil {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *PatchedCall) HasSuccess() bool {
	if o != nil && o.Success != nil {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *PatchedCall) SetSuccess(v bool) {
	o.Success = &v
}

// GetClientCaller returns the ClientCaller field value if set, zero value otherwise.
func (o *PatchedCall) GetClientCaller() ClientInfo {
	if o == nil || o.ClientCaller == nil {
		var ret ClientInfo
		return ret
	}
	return *o.ClientCaller
}

// GetClientCallerOk returns a tuple with the ClientCaller field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCall) GetClientCallerOk() (*ClientInfo, bool) {
	if o == nil || o.ClientCaller == nil {
		return nil, false
	}
	return o.ClientCaller, true
}

// HasClientCaller returns a boolean if a field has been set.
func (o *PatchedCall) HasClientCaller() bool {
	if o != nil && o.ClientCaller != nil {
		return true
	}

	return false
}

// SetClientCaller gets a reference to the given ClientInfo and assigns it to the ClientCaller field.
func (o *PatchedCall) SetClientCaller(v ClientInfo) {
	o.ClientCaller = &v
}

// GetGuestCalled returns the GuestCalled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedCall) GetGuestCalled() int32 {
	if o == nil || o.GuestCalled.Get() == nil {
		var ret int32
		return ret
	}
	return *o.GuestCalled.Get()
}

// GetGuestCalledOk returns a tuple with the GuestCalled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedCall) GetGuestCalledOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.GuestCalled.Get(), o.GuestCalled.IsSet()
}

// HasGuestCalled returns a boolean if a field has been set.
func (o *PatchedCall) HasGuestCalled() bool {
	if o != nil && o.GuestCalled.IsSet() {
		return true
	}

	return false
}

// SetGuestCalled gets a reference to the given NullableInt32 and assigns it to the GuestCalled field.
func (o *PatchedCall) SetGuestCalled(v int32) {
	o.GuestCalled.Set(&v)
}
// SetGuestCalledNil sets the value for GuestCalled to be an explicit nil
func (o *PatchedCall) SetGuestCalledNil() {
	o.GuestCalled.Set(nil)
}

// UnsetGuestCalled ensures that no value is present for GuestCalled, not even an explicit nil
func (o *PatchedCall) UnsetGuestCalled() {
	o.GuestCalled.Unset()
}

func (o PatchedCall) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CallId != nil {
		toSerialize["call_id"] = o.CallId
	}
	if o.TenantId != nil {
		toSerialize["tenant_id"] = o.TenantId
	}
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}
	if o.Success != nil {
		toSerialize["success"] = o.Success
	}
	if o.ClientCaller != nil {
		toSerialize["client_caller"] = o.ClientCaller
	}
	if o.GuestCalled.IsSet() {
		toSerialize["guest_called"] = o.GuestCalled.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedCall struct {
	value *PatchedCall
	isSet bool
}

func (v NullablePatchedCall) Get() *PatchedCall {
	return v.value
}

func (v *NullablePatchedCall) Set(val *PatchedCall) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedCall) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedCall) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedCall(val *PatchedCall) *NullablePatchedCall {
	return &NullablePatchedCall{value: val, isSet: true}
}

func (v NullablePatchedCall) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedCall) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

