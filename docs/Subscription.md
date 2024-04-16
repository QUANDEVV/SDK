# Subscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Tenant** | [**PaymentsTenant**](PaymentsTenant.md) |  | [readonly] 
**Tier** | [**Tier**](Tier.md) |  | [readonly] 
**Quota** | [**Quota**](Quota.md) |  | [readonly] 
**Metadata** | Pointer to **interface{}** | Metadata | [optional] 
**CreatedAt** | **NullableString** |  | [readonly] 
**DateTimeCreatedAt** | **NullableString** |  | [readonly] 
**Timestamp** | **NullableTime** | The timestamp of the chat. | [readonly] 
**UpdatedAt** | **NullableTime** |  | [readonly] 
**IsValidFlag** | Pointer to **bool** |  | [optional] 
**StartDate** | Pointer to **NullableTime** |  | [optional] 
**ExpiryDate** | Pointer to **NullableTime** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**FreeTrialEnd** | Pointer to **NullableTime** |  | [optional] 
**FreeTrialUsed** | Pointer to **bool** |  | [optional] 
**Cancelled** | Pointer to **bool** |  | [optional] 

## Methods

### NewSubscription

`func NewSubscription(id int32, tenant PaymentsTenant, tier Tier, quota Quota, createdAt NullableString, dateTimeCreatedAt NullableString, timestamp NullableTime, updatedAt NullableTime, ) *Subscription`

NewSubscription instantiates a new Subscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionWithDefaults

`func NewSubscriptionWithDefaults() *Subscription`

NewSubscriptionWithDefaults instantiates a new Subscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Subscription) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Subscription) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Subscription) SetId(v int32)`

SetId sets Id field to given value.


### GetTenant

`func (o *Subscription) GetTenant() PaymentsTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *Subscription) GetTenantOk() (*PaymentsTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *Subscription) SetTenant(v PaymentsTenant)`

SetTenant sets Tenant field to given value.


### SetTenantNil

`func (o *Subscription) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *Subscription) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetTier

`func (o *Subscription) GetTier() Tier`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *Subscription) GetTierOk() (*Tier, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *Subscription) SetTier(v Tier)`

SetTier sets Tier field to given value.


### SetTierNil

`func (o *Subscription) SetTierNil(b bool)`

 SetTierNil sets the value for Tier to be an explicit nil

### UnsetTier
`func (o *Subscription) UnsetTier()`

UnsetTier ensures that no value is present for Tier, not even an explicit nil
### GetQuota

`func (o *Subscription) GetQuota() Quota`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *Subscription) GetQuotaOk() (*Quota, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *Subscription) SetQuota(v Quota)`

SetQuota sets Quota field to given value.


### SetQuotaNil

`func (o *Subscription) SetQuotaNil(b bool)`

 SetQuotaNil sets the value for Quota to be an explicit nil

### UnsetQuota
`func (o *Subscription) UnsetQuota()`

UnsetQuota ensures that no value is present for Quota, not even an explicit nil
### GetMetadata

`func (o *Subscription) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Subscription) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Subscription) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Subscription) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *Subscription) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *Subscription) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetCreatedAt

`func (o *Subscription) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Subscription) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Subscription) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *Subscription) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Subscription) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDateTimeCreatedAt

`func (o *Subscription) GetDateTimeCreatedAt() string`

GetDateTimeCreatedAt returns the DateTimeCreatedAt field if non-nil, zero value otherwise.

### GetDateTimeCreatedAtOk

`func (o *Subscription) GetDateTimeCreatedAtOk() (*string, bool)`

GetDateTimeCreatedAtOk returns a tuple with the DateTimeCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeCreatedAt

`func (o *Subscription) SetDateTimeCreatedAt(v string)`

SetDateTimeCreatedAt sets DateTimeCreatedAt field to given value.


### SetDateTimeCreatedAtNil

`func (o *Subscription) SetDateTimeCreatedAtNil(b bool)`

 SetDateTimeCreatedAtNil sets the value for DateTimeCreatedAt to be an explicit nil

### UnsetDateTimeCreatedAt
`func (o *Subscription) UnsetDateTimeCreatedAt()`

UnsetDateTimeCreatedAt ensures that no value is present for DateTimeCreatedAt, not even an explicit nil
### GetTimestamp

`func (o *Subscription) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Subscription) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Subscription) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### SetTimestampNil

`func (o *Subscription) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *Subscription) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetUpdatedAt

`func (o *Subscription) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Subscription) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Subscription) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *Subscription) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *Subscription) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetIsValidFlag

`func (o *Subscription) GetIsValidFlag() bool`

GetIsValidFlag returns the IsValidFlag field if non-nil, zero value otherwise.

### GetIsValidFlagOk

`func (o *Subscription) GetIsValidFlagOk() (*bool, bool)`

GetIsValidFlagOk returns a tuple with the IsValidFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValidFlag

`func (o *Subscription) SetIsValidFlag(v bool)`

SetIsValidFlag sets IsValidFlag field to given value.

### HasIsValidFlag

`func (o *Subscription) HasIsValidFlag() bool`

HasIsValidFlag returns a boolean if a field has been set.

### GetStartDate

`func (o *Subscription) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Subscription) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Subscription) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *Subscription) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *Subscription) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *Subscription) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetExpiryDate

`func (o *Subscription) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *Subscription) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *Subscription) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *Subscription) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDateNil

`func (o *Subscription) SetExpiryDateNil(b bool)`

 SetExpiryDateNil sets the value for ExpiryDate to be an explicit nil

### UnsetExpiryDate
`func (o *Subscription) UnsetExpiryDate()`

UnsetExpiryDate ensures that no value is present for ExpiryDate, not even an explicit nil
### GetEmail

`func (o *Subscription) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Subscription) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Subscription) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Subscription) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *Subscription) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *Subscription) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetFreeTrialEnd

`func (o *Subscription) GetFreeTrialEnd() time.Time`

GetFreeTrialEnd returns the FreeTrialEnd field if non-nil, zero value otherwise.

### GetFreeTrialEndOk

`func (o *Subscription) GetFreeTrialEndOk() (*time.Time, bool)`

GetFreeTrialEndOk returns a tuple with the FreeTrialEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeTrialEnd

`func (o *Subscription) SetFreeTrialEnd(v time.Time)`

SetFreeTrialEnd sets FreeTrialEnd field to given value.

### HasFreeTrialEnd

`func (o *Subscription) HasFreeTrialEnd() bool`

HasFreeTrialEnd returns a boolean if a field has been set.

### SetFreeTrialEndNil

`func (o *Subscription) SetFreeTrialEndNil(b bool)`

 SetFreeTrialEndNil sets the value for FreeTrialEnd to be an explicit nil

### UnsetFreeTrialEnd
`func (o *Subscription) UnsetFreeTrialEnd()`

UnsetFreeTrialEnd ensures that no value is present for FreeTrialEnd, not even an explicit nil
### GetFreeTrialUsed

`func (o *Subscription) GetFreeTrialUsed() bool`

GetFreeTrialUsed returns the FreeTrialUsed field if non-nil, zero value otherwise.

### GetFreeTrialUsedOk

`func (o *Subscription) GetFreeTrialUsedOk() (*bool, bool)`

GetFreeTrialUsedOk returns a tuple with the FreeTrialUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeTrialUsed

`func (o *Subscription) SetFreeTrialUsed(v bool)`

SetFreeTrialUsed sets FreeTrialUsed field to given value.

### HasFreeTrialUsed

`func (o *Subscription) HasFreeTrialUsed() bool`

HasFreeTrialUsed returns a boolean if a field has been set.

### GetCancelled

`func (o *Subscription) GetCancelled() bool`

GetCancelled returns the Cancelled field if non-nil, zero value otherwise.

### GetCancelledOk

`func (o *Subscription) GetCancelledOk() (*bool, bool)`

GetCancelledOk returns a tuple with the Cancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelled

`func (o *Subscription) SetCancelled(v bool)`

SetCancelled sets Cancelled field to given value.

### HasCancelled

`func (o *Subscription) HasCancelled() bool`

HasCancelled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


