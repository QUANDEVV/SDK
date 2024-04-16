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

// PatchedDelegation struct for PatchedDelegation
type PatchedDelegation struct {
	DelegationId *int32 `json:"delegation_id,omitempty"`
	TenantId *TenantInfo `json:"tenant_id,omitempty"`
	Employee *[]Employee `json:"employee,omitempty"`
	Admin *[]Admin `json:"admin,omitempty"`
	// The issue ID UUID .
	AssignedIssue NullableInt32 `json:"assigned_issue,omitempty"`
	// The chat ID UUID for an instance of a chat.
	AssignedChat NullableInt32 `json:"assigned_chat,omitempty"`
	AssignedService NullableInt32 `json:"assigned_service,omitempty"`
	// The timestamp of the chat.
	Timestamp NullableTime `json:"timestamp,omitempty"`
	DelegatedBy NullableInt32 `json:"delegated_by,omitempty"`
	// The department ID.
	DepartmentDelegatedTo NullableInt32 `json:"department_delegated_to,omitempty"`
}

// NewPatchedDelegation instantiates a new PatchedDelegation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedDelegation() *PatchedDelegation {
	this := PatchedDelegation{}
	return &this
}

// NewPatchedDelegationWithDefaults instantiates a new PatchedDelegation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedDelegationWithDefaults() *PatchedDelegation {
	this := PatchedDelegation{}
	return &this
}

// GetDelegationId returns the DelegationId field value if set, zero value otherwise.
func (o *PatchedDelegation) GetDelegationId() int32 {
	if o == nil || o.DelegationId == nil {
		var ret int32
		return ret
	}
	return *o.DelegationId
}

// GetDelegationIdOk returns a tuple with the DelegationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDelegation) GetDelegationIdOk() (*int32, bool) {
	if o == nil || o.DelegationId == nil {
		return nil, false
	}
	return o.DelegationId, true
}

// HasDelegationId returns a boolean if a field has been set.
func (o *PatchedDelegation) HasDelegationId() bool {
	if o != nil && o.DelegationId != nil {
		return true
	}

	return false
}

// SetDelegationId gets a reference to the given int32 and assigns it to the DelegationId field.
func (o *PatchedDelegation) SetDelegationId(v int32) {
	o.DelegationId = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *PatchedDelegation) GetTenantId() TenantInfo {
	if o == nil || o.TenantId == nil {
		var ret TenantInfo
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDelegation) GetTenantIdOk() (*TenantInfo, bool) {
	if o == nil || o.TenantId == nil {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *PatchedDelegation) HasTenantId() bool {
	if o != nil && o.TenantId != nil {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given TenantInfo and assigns it to the TenantId field.
func (o *PatchedDelegation) SetTenantId(v TenantInfo) {
	o.TenantId = &v
}

// GetEmployee returns the Employee field value if set, zero value otherwise.
func (o *PatchedDelegation) GetEmployee() []Employee {
	if o == nil || o.Employee == nil {
		var ret []Employee
		return ret
	}
	return *o.Employee
}

// GetEmployeeOk returns a tuple with the Employee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDelegation) GetEmployeeOk() (*[]Employee, bool) {
	if o == nil || o.Employee == nil {
		return nil, false
	}
	return o.Employee, true
}

// HasEmployee returns a boolean if a field has been set.
func (o *PatchedDelegation) HasEmployee() bool {
	if o != nil && o.Employee != nil {
		return true
	}

	return false
}

// SetEmployee gets a reference to the given []Employee and assigns it to the Employee field.
func (o *PatchedDelegation) SetEmployee(v []Employee) {
	o.Employee = &v
}

// GetAdmin returns the Admin field value if set, zero value otherwise.
func (o *PatchedDelegation) GetAdmin() []Admin {
	if o == nil || o.Admin == nil {
		var ret []Admin
		return ret
	}
	return *o.Admin
}

// GetAdminOk returns a tuple with the Admin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDelegation) GetAdminOk() (*[]Admin, bool) {
	if o == nil || o.Admin == nil {
		return nil, false
	}
	return o.Admin, true
}

// HasAdmin returns a boolean if a field has been set.
func (o *PatchedDelegation) HasAdmin() bool {
	if o != nil && o.Admin != nil {
		return true
	}

	return false
}

// SetAdmin gets a reference to the given []Admin and assigns it to the Admin field.
func (o *PatchedDelegation) SetAdmin(v []Admin) {
	o.Admin = &v
}

// GetAssignedIssue returns the AssignedIssue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedDelegation) GetAssignedIssue() int32 {
	if o == nil || o.AssignedIssue.Get() == nil {
		var ret int32
		return ret
	}
	return *o.AssignedIssue.Get()
}

// GetAssignedIssueOk returns a tuple with the AssignedIssue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedDelegation) GetAssignedIssueOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AssignedIssue.Get(), o.AssignedIssue.IsSet()
}

// HasAssignedIssue returns a boolean if a field has been set.
func (o *PatchedDelegation) HasAssignedIssue() bool {
	if o != nil && o.AssignedIssue.IsSet() {
		return true
	}

	return false
}

// SetAssignedIssue gets a reference to the given NullableInt32 and assigns it to the AssignedIssue field.
func (o *PatchedDelegation) SetAssignedIssue(v int32) {
	o.AssignedIssue.Set(&v)
}
// SetAssignedIssueNil sets the value for AssignedIssue to be an explicit nil
func (o *PatchedDelegation) SetAssignedIssueNil() {
	o.AssignedIssue.Set(nil)
}

// UnsetAssignedIssue ensures that no value is present for AssignedIssue, not even an explicit nil
func (o *PatchedDelegation) UnsetAssignedIssue() {
	o.AssignedIssue.Unset()
}

// GetAssignedChat returns the AssignedChat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedDelegation) GetAssignedChat() int32 {
	if o == nil || o.AssignedChat.Get() == nil {
		var ret int32
		return ret
	}
	return *o.AssignedChat.Get()
}

// GetAssignedChatOk returns a tuple with the AssignedChat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedDelegation) GetAssignedChatOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AssignedChat.Get(), o.AssignedChat.IsSet()
}

// HasAssignedChat returns a boolean if a field has been set.
func (o *PatchedDelegation) HasAssignedChat() bool {
	if o != nil && o.AssignedChat.IsSet() {
		return true
	}

	return false
}

// SetAssignedChat gets a reference to the given NullableInt32 and assigns it to the AssignedChat field.
func (o *PatchedDelegation) SetAssignedChat(v int32) {
	o.AssignedChat.Set(&v)
}
// SetAssignedChatNil sets the value for AssignedChat to be an explicit nil
func (o *PatchedDelegation) SetAssignedChatNil() {
	o.AssignedChat.Set(nil)
}

// UnsetAssignedChat ensures that no value is present for AssignedChat, not even an explicit nil
func (o *PatchedDelegation) UnsetAssignedChat() {
	o.AssignedChat.Unset()
}

// GetAssignedService returns the AssignedService field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedDelegation) GetAssignedService() int32 {
	if o == nil || o.AssignedService.Get() == nil {
		var ret int32
		return ret
	}
	return *o.AssignedService.Get()
}

// GetAssignedServiceOk returns a tuple with the AssignedService field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedDelegation) GetAssignedServiceOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AssignedService.Get(), o.AssignedService.IsSet()
}

// HasAssignedService returns a boolean if a field has been set.
func (o *PatchedDelegation) HasAssignedService() bool {
	if o != nil && o.AssignedService.IsSet() {
		return true
	}

	return false
}

// SetAssignedService gets a reference to the given NullableInt32 and assigns it to the AssignedService field.
func (o *PatchedDelegation) SetAssignedService(v int32) {
	o.AssignedService.Set(&v)
}
// SetAssignedServiceNil sets the value for AssignedService to be an explicit nil
func (o *PatchedDelegation) SetAssignedServiceNil() {
	o.AssignedService.Set(nil)
}

// UnsetAssignedService ensures that no value is present for AssignedService, not even an explicit nil
func (o *PatchedDelegation) UnsetAssignedService() {
	o.AssignedService.Unset()
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedDelegation) GetTimestamp() time.Time {
	if o == nil || o.Timestamp.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Timestamp.Get()
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedDelegation) GetTimestampOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Timestamp.Get(), o.Timestamp.IsSet()
}

// HasTimestamp returns a boolean if a field has been set.
func (o *PatchedDelegation) HasTimestamp() bool {
	if o != nil && o.Timestamp.IsSet() {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given NullableTime and assigns it to the Timestamp field.
func (o *PatchedDelegation) SetTimestamp(v time.Time) {
	o.Timestamp.Set(&v)
}
// SetTimestampNil sets the value for Timestamp to be an explicit nil
func (o *PatchedDelegation) SetTimestampNil() {
	o.Timestamp.Set(nil)
}

// UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
func (o *PatchedDelegation) UnsetTimestamp() {
	o.Timestamp.Unset()
}

// GetDelegatedBy returns the DelegatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedDelegation) GetDelegatedBy() int32 {
	if o == nil || o.DelegatedBy.Get() == nil {
		var ret int32
		return ret
	}
	return *o.DelegatedBy.Get()
}

// GetDelegatedByOk returns a tuple with the DelegatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedDelegation) GetDelegatedByOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DelegatedBy.Get(), o.DelegatedBy.IsSet()
}

// HasDelegatedBy returns a boolean if a field has been set.
func (o *PatchedDelegation) HasDelegatedBy() bool {
	if o != nil && o.DelegatedBy.IsSet() {
		return true
	}

	return false
}

// SetDelegatedBy gets a reference to the given NullableInt32 and assigns it to the DelegatedBy field.
func (o *PatchedDelegation) SetDelegatedBy(v int32) {
	o.DelegatedBy.Set(&v)
}
// SetDelegatedByNil sets the value for DelegatedBy to be an explicit nil
func (o *PatchedDelegation) SetDelegatedByNil() {
	o.DelegatedBy.Set(nil)
}

// UnsetDelegatedBy ensures that no value is present for DelegatedBy, not even an explicit nil
func (o *PatchedDelegation) UnsetDelegatedBy() {
	o.DelegatedBy.Unset()
}

// GetDepartmentDelegatedTo returns the DepartmentDelegatedTo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedDelegation) GetDepartmentDelegatedTo() int32 {
	if o == nil || o.DepartmentDelegatedTo.Get() == nil {
		var ret int32
		return ret
	}
	return *o.DepartmentDelegatedTo.Get()
}

// GetDepartmentDelegatedToOk returns a tuple with the DepartmentDelegatedTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedDelegation) GetDepartmentDelegatedToOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DepartmentDelegatedTo.Get(), o.DepartmentDelegatedTo.IsSet()
}

// HasDepartmentDelegatedTo returns a boolean if a field has been set.
func (o *PatchedDelegation) HasDepartmentDelegatedTo() bool {
	if o != nil && o.DepartmentDelegatedTo.IsSet() {
		return true
	}

	return false
}

// SetDepartmentDelegatedTo gets a reference to the given NullableInt32 and assigns it to the DepartmentDelegatedTo field.
func (o *PatchedDelegation) SetDepartmentDelegatedTo(v int32) {
	o.DepartmentDelegatedTo.Set(&v)
}
// SetDepartmentDelegatedToNil sets the value for DepartmentDelegatedTo to be an explicit nil
func (o *PatchedDelegation) SetDepartmentDelegatedToNil() {
	o.DepartmentDelegatedTo.Set(nil)
}

// UnsetDepartmentDelegatedTo ensures that no value is present for DepartmentDelegatedTo, not even an explicit nil
func (o *PatchedDelegation) UnsetDepartmentDelegatedTo() {
	o.DepartmentDelegatedTo.Unset()
}

func (o PatchedDelegation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DelegationId != nil {
		toSerialize["delegation_id"] = o.DelegationId
	}
	if o.TenantId != nil {
		toSerialize["tenant_id"] = o.TenantId
	}
	if o.Employee != nil {
		toSerialize["employee"] = o.Employee
	}
	if o.Admin != nil {
		toSerialize["admin"] = o.Admin
	}
	if o.AssignedIssue.IsSet() {
		toSerialize["assigned_issue"] = o.AssignedIssue.Get()
	}
	if o.AssignedChat.IsSet() {
		toSerialize["assigned_chat"] = o.AssignedChat.Get()
	}
	if o.AssignedService.IsSet() {
		toSerialize["assigned_service"] = o.AssignedService.Get()
	}
	if o.Timestamp.IsSet() {
		toSerialize["timestamp"] = o.Timestamp.Get()
	}
	if o.DelegatedBy.IsSet() {
		toSerialize["delegated_by"] = o.DelegatedBy.Get()
	}
	if o.DepartmentDelegatedTo.IsSet() {
		toSerialize["department_delegated_to"] = o.DepartmentDelegatedTo.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedDelegation struct {
	value *PatchedDelegation
	isSet bool
}

func (v NullablePatchedDelegation) Get() *PatchedDelegation {
	return v.value
}

func (v *NullablePatchedDelegation) Set(val *PatchedDelegation) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedDelegation) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedDelegation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedDelegation(val *PatchedDelegation) *NullablePatchedDelegation {
	return &NullablePatchedDelegation{value: val, isSet: true}
}

func (v NullablePatchedDelegation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedDelegation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

