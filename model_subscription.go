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

// Subscription struct for Subscription
type Subscription struct {
	Id int32 `json:"id"`
	Tenant PaymentsTenant `json:"tenant"`
	Tier Tier `json:"tier"`
	Quota Quota `json:"quota"`
	// Metadata
	Metadata interface{} `json:"metadata,omitempty"`
	CreatedAt NullableString `json:"created_at"`
	DateTimeCreatedAt NullableString `json:"date_time_created_at"`
	// The timestamp of the chat.
	Timestamp NullableTime `json:"timestamp"`
	UpdatedAt NullableTime `json:"updated_at"`
	IsValidFlag *bool `json:"is_valid_flag,omitempty"`
	StartDate NullableTime `json:"start_date,omitempty"`
	ExpiryDate NullableTime `json:"expiry_date,omitempty"`
	Email NullableString `json:"email,omitempty"`
	FreeTrialEnd NullableTime `json:"free_trial_end,omitempty"`
	FreeTrialUsed *bool `json:"free_trial_used,omitempty"`
	Cancelled *bool `json:"cancelled,omitempty"`
}

// NewSubscription instantiates a new Subscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscription(id int32, tenant PaymentsTenant, tier Tier, quota Quota, createdAt NullableString, dateTimeCreatedAt NullableString, timestamp NullableTime, updatedAt NullableTime) *Subscription {
	this := Subscription{}
	this.Id = id
	this.Tenant = tenant
	this.Tier = tier
	this.Quota = quota
	this.CreatedAt = createdAt
	this.DateTimeCreatedAt = dateTimeCreatedAt
	this.Timestamp = timestamp
	this.UpdatedAt = updatedAt
	return &this
}

// NewSubscriptionWithDefaults instantiates a new Subscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionWithDefaults() *Subscription {
	this := Subscription{}
	return &this
}

// GetId returns the Id field value
func (o *Subscription) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Subscription) GetIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Subscription) SetId(v int32) {
	o.Id = v
}

// GetTenant returns the Tenant field value
// If the value is explicit nil, the zero value for PaymentsTenant will be returned
func (o *Subscription) GetTenant() PaymentsTenant {
	if o == nil {
		var ret PaymentsTenant
		return ret
	}

	return o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Subscription) GetTenantOk() (*PaymentsTenant, bool) {
	if o == nil || o.Tenant == nil {
		return nil, false
	}
	return &o.Tenant, true
}

// SetTenant sets field value
func (o *Subscription) SetTenant(v PaymentsTenant) {
	o.Tenant = v
}

// GetTier returns the Tier field value
// If the value is explicit nil, the zero value for Tier will be returned
func (o *Subscription) GetTier() Tier {
	if o == nil {
		var ret Tier
		return ret
	}

	return o.Tier
}

// GetTierOk returns a tuple with the Tier field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Subscription) GetTierOk() (*Tier, bool) {
	if o == nil || o.Tier == nil {
		return nil, false
	}
	return &o.Tier, true
}

// SetTier sets field value
func (o *Subscription) SetTier(v Tier) {
	o.Tier = v
}

// GetQuota returns the Quota field value
// If the value is explicit nil, the zero value for Quota will be returned
func (o *Subscription) GetQuota() Quota {
	if o == nil {
		var ret Quota
		return ret
	}

	return o.Quota
}

// GetQuotaOk returns a tuple with the Quota field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Subscription) GetQuotaOk() (*Quota, bool) {
	if o == nil || o.Quota == nil {
		return nil, false
	}
	return &o.Quota, true
}

// SetQuota sets field value
func (o *Subscription) SetQuota(v Quota) {
	o.Quota = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Subscription) GetMetadata() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Subscription) GetMetadataOk() (*interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Subscription) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *Subscription) SetMetadata(v interface{}) {
	o.Metadata = v
}

// GetCreatedAt returns the CreatedAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Subscription) GetCreatedAt() string {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Subscription) GetCreatedAtOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// SetCreatedAt sets field value
func (o *Subscription) SetCreatedAt(v string) {
	o.CreatedAt.Set(&v)
}

// GetDateTimeCreatedAt returns the DateTimeCreatedAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Subscription) GetDateTimeCreatedAt() string {
	if o == nil || o.DateTimeCreatedAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.DateTimeCreatedAt.Get()
}

// GetDateTimeCreatedAtOk returns a tuple with the DateTimeCreatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Subscription) GetDateTimeCreatedAtOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DateTimeCreatedAt.Get(), o.DateTimeCreatedAt.IsSet()
}

// SetDateTimeCreatedAt sets field value
func (o *Subscription) SetDateTimeCreatedAt(v string) {
	o.DateTimeCreatedAt.Set(&v)
}

// GetTimestamp returns the Timestamp field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Subscription) GetTimestamp() time.Time {
	if o == nil || o.Timestamp.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Timestamp.Get()
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Subscription) GetTimestampOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Timestamp.Get(), o.Timestamp.IsSet()
}

// SetTimestamp sets field value
func (o *Subscription) SetTimestamp(v time.Time) {
	o.Timestamp.Set(&v)
}

// GetUpdatedAt returns the UpdatedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Subscription) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Subscription) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// SetUpdatedAt sets field value
func (o *Subscription) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}

// GetIsValidFlag returns the IsValidFlag field value if set, zero value otherwise.
func (o *Subscription) GetIsValidFlag() bool {
	if o == nil || o.IsValidFlag == nil {
		var ret bool
		return ret
	}
	return *o.IsValidFlag
}

// GetIsValidFlagOk returns a tuple with the IsValidFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetIsValidFlagOk() (*bool, bool) {
	if o == nil || o.IsValidFlag == nil {
		return nil, false
	}
	return o.IsValidFlag, true
}

// HasIsValidFlag returns a boolean if a field has been set.
func (o *Subscription) HasIsValidFlag() bool {
	if o != nil && o.IsValidFlag != nil {
		return true
	}

	return false
}

// SetIsValidFlag gets a reference to the given bool and assigns it to the IsValidFlag field.
func (o *Subscription) SetIsValidFlag(v bool) {
	o.IsValidFlag = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Subscription) GetStartDate() time.Time {
	if o == nil || o.StartDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate.Get()
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Subscription) GetStartDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartDate.Get(), o.StartDate.IsSet()
}

// HasStartDate returns a boolean if a field has been set.
func (o *Subscription) HasStartDate() bool {
	if o != nil && o.StartDate.IsSet() {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given NullableTime and assigns it to the StartDate field.
func (o *Subscription) SetStartDate(v time.Time) {
	o.StartDate.Set(&v)
}
// SetStartDateNil sets the value for StartDate to be an explicit nil
func (o *Subscription) SetStartDateNil() {
	o.StartDate.Set(nil)
}

// UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
func (o *Subscription) UnsetStartDate() {
	o.StartDate.Unset()
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Subscription) GetExpiryDate() time.Time {
	if o == nil || o.ExpiryDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiryDate.Get()
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Subscription) GetExpiryDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExpiryDate.Get(), o.ExpiryDate.IsSet()
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *Subscription) HasExpiryDate() bool {
	if o != nil && o.ExpiryDate.IsSet() {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given NullableTime and assigns it to the ExpiryDate field.
func (o *Subscription) SetExpiryDate(v time.Time) {
	o.ExpiryDate.Set(&v)
}
// SetExpiryDateNil sets the value for ExpiryDate to be an explicit nil
func (o *Subscription) SetExpiryDateNil() {
	o.ExpiryDate.Set(nil)
}

// UnsetExpiryDate ensures that no value is present for ExpiryDate, not even an explicit nil
func (o *Subscription) UnsetExpiryDate() {
	o.ExpiryDate.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Subscription) GetEmail() string {
	if o == nil || o.Email.Get() == nil {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Subscription) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *Subscription) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *Subscription) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *Subscription) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *Subscription) UnsetEmail() {
	o.Email.Unset()
}

// GetFreeTrialEnd returns the FreeTrialEnd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Subscription) GetFreeTrialEnd() time.Time {
	if o == nil || o.FreeTrialEnd.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.FreeTrialEnd.Get()
}

// GetFreeTrialEndOk returns a tuple with the FreeTrialEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Subscription) GetFreeTrialEndOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FreeTrialEnd.Get(), o.FreeTrialEnd.IsSet()
}

// HasFreeTrialEnd returns a boolean if a field has been set.
func (o *Subscription) HasFreeTrialEnd() bool {
	if o != nil && o.FreeTrialEnd.IsSet() {
		return true
	}

	return false
}

// SetFreeTrialEnd gets a reference to the given NullableTime and assigns it to the FreeTrialEnd field.
func (o *Subscription) SetFreeTrialEnd(v time.Time) {
	o.FreeTrialEnd.Set(&v)
}
// SetFreeTrialEndNil sets the value for FreeTrialEnd to be an explicit nil
func (o *Subscription) SetFreeTrialEndNil() {
	o.FreeTrialEnd.Set(nil)
}

// UnsetFreeTrialEnd ensures that no value is present for FreeTrialEnd, not even an explicit nil
func (o *Subscription) UnsetFreeTrialEnd() {
	o.FreeTrialEnd.Unset()
}

// GetFreeTrialUsed returns the FreeTrialUsed field value if set, zero value otherwise.
func (o *Subscription) GetFreeTrialUsed() bool {
	if o == nil || o.FreeTrialUsed == nil {
		var ret bool
		return ret
	}
	return *o.FreeTrialUsed
}

// GetFreeTrialUsedOk returns a tuple with the FreeTrialUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetFreeTrialUsedOk() (*bool, bool) {
	if o == nil || o.FreeTrialUsed == nil {
		return nil, false
	}
	return o.FreeTrialUsed, true
}

// HasFreeTrialUsed returns a boolean if a field has been set.
func (o *Subscription) HasFreeTrialUsed() bool {
	if o != nil && o.FreeTrialUsed != nil {
		return true
	}

	return false
}

// SetFreeTrialUsed gets a reference to the given bool and assigns it to the FreeTrialUsed field.
func (o *Subscription) SetFreeTrialUsed(v bool) {
	o.FreeTrialUsed = &v
}

// GetCancelled returns the Cancelled field value if set, zero value otherwise.
func (o *Subscription) GetCancelled() bool {
	if o == nil || o.Cancelled == nil {
		var ret bool
		return ret
	}
	return *o.Cancelled
}

// GetCancelledOk returns a tuple with the Cancelled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetCancelledOk() (*bool, bool) {
	if o == nil || o.Cancelled == nil {
		return nil, false
	}
	return o.Cancelled, true
}

// HasCancelled returns a boolean if a field has been set.
func (o *Subscription) HasCancelled() bool {
	if o != nil && o.Cancelled != nil {
		return true
	}

	return false
}

// SetCancelled gets a reference to the given bool and assigns it to the Cancelled field.
func (o *Subscription) SetCancelled(v bool) {
	o.Cancelled = &v
}

func (o Subscription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Tenant != nil {
		toSerialize["tenant"] = o.Tenant
	}
	if o.Tier != nil {
		toSerialize["tier"] = o.Tier
	}
	if o.Quota != nil {
		toSerialize["quota"] = o.Quota
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
	if o.IsValidFlag != nil {
		toSerialize["is_valid_flag"] = o.IsValidFlag
	}
	if o.StartDate.IsSet() {
		toSerialize["start_date"] = o.StartDate.Get()
	}
	if o.ExpiryDate.IsSet() {
		toSerialize["expiry_date"] = o.ExpiryDate.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.FreeTrialEnd.IsSet() {
		toSerialize["free_trial_end"] = o.FreeTrialEnd.Get()
	}
	if o.FreeTrialUsed != nil {
		toSerialize["free_trial_used"] = o.FreeTrialUsed
	}
	if o.Cancelled != nil {
		toSerialize["cancelled"] = o.Cancelled
	}
	return json.Marshal(toSerialize)
}

type NullableSubscription struct {
	value *Subscription
	isSet bool
}

func (v NullableSubscription) Get() *Subscription {
	return v.value
}

func (v *NullableSubscription) Set(val *Subscription) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscription(val *Subscription) *NullableSubscription {
	return &NullableSubscription{value: val, isSet: true}
}

func (v NullableSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


