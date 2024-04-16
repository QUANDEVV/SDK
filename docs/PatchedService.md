# PatchedService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | Pointer to **int32** |  | [optional] [readonly] 
**TenantId** | Pointer to **int32** | The tenant chats ID. | [optional] 
**NameOfService** | Pointer to **string** |  | [optional] 
**ServiceDescription** | Pointer to **string** |  | [optional] 
**ServiceType** | Pointer to **int32** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**ResolutionPeriod** | Pointer to **string** |  | [optional] 
**ServiceAvailability** | Pointer to [**ServiceAvailabilityEnum**](ServiceAvailabilityEnum.md) |  | [optional] 
**AiPowered** | Pointer to **bool** |  | [optional] 
**AiSolved** | Pointer to **NullableBool** |  | [optional] 
**Details** | Pointer to **interface{}** |  | [optional] 
**ServiceActionPlan** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewPatchedService

`func NewPatchedService() *PatchedService`

NewPatchedService instantiates a new PatchedService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedServiceWithDefaults

`func NewPatchedServiceWithDefaults() *PatchedService`

NewPatchedServiceWithDefaults instantiates a new PatchedService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *PatchedService) GetServiceId() int32`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *PatchedService) GetServiceIdOk() (*int32, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *PatchedService) SetServiceId(v int32)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *PatchedService) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetTenantId

`func (o *PatchedService) GetTenantId() int32`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *PatchedService) GetTenantIdOk() (*int32, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *PatchedService) SetTenantId(v int32)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *PatchedService) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetNameOfService

`func (o *PatchedService) GetNameOfService() string`

GetNameOfService returns the NameOfService field if non-nil, zero value otherwise.

### GetNameOfServiceOk

`func (o *PatchedService) GetNameOfServiceOk() (*string, bool)`

GetNameOfServiceOk returns a tuple with the NameOfService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameOfService

`func (o *PatchedService) SetNameOfService(v string)`

SetNameOfService sets NameOfService field to given value.

### HasNameOfService

`func (o *PatchedService) HasNameOfService() bool`

HasNameOfService returns a boolean if a field has been set.

### GetServiceDescription

`func (o *PatchedService) GetServiceDescription() string`

GetServiceDescription returns the ServiceDescription field if non-nil, zero value otherwise.

### GetServiceDescriptionOk

`func (o *PatchedService) GetServiceDescriptionOk() (*string, bool)`

GetServiceDescriptionOk returns a tuple with the ServiceDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDescription

`func (o *PatchedService) SetServiceDescription(v string)`

SetServiceDescription sets ServiceDescription field to given value.

### HasServiceDescription

`func (o *PatchedService) HasServiceDescription() bool`

HasServiceDescription returns a boolean if a field has been set.

### GetServiceType

`func (o *PatchedService) GetServiceType() int32`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *PatchedService) GetServiceTypeOk() (*int32, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *PatchedService) SetServiceType(v int32)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *PatchedService) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetIsActive

`func (o *PatchedService) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *PatchedService) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *PatchedService) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *PatchedService) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetResolutionPeriod

`func (o *PatchedService) GetResolutionPeriod() string`

GetResolutionPeriod returns the ResolutionPeriod field if non-nil, zero value otherwise.

### GetResolutionPeriodOk

`func (o *PatchedService) GetResolutionPeriodOk() (*string, bool)`

GetResolutionPeriodOk returns a tuple with the ResolutionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionPeriod

`func (o *PatchedService) SetResolutionPeriod(v string)`

SetResolutionPeriod sets ResolutionPeriod field to given value.

### HasResolutionPeriod

`func (o *PatchedService) HasResolutionPeriod() bool`

HasResolutionPeriod returns a boolean if a field has been set.

### GetServiceAvailability

`func (o *PatchedService) GetServiceAvailability() ServiceAvailabilityEnum`

GetServiceAvailability returns the ServiceAvailability field if non-nil, zero value otherwise.

### GetServiceAvailabilityOk

`func (o *PatchedService) GetServiceAvailabilityOk() (*ServiceAvailabilityEnum, bool)`

GetServiceAvailabilityOk returns a tuple with the ServiceAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAvailability

`func (o *PatchedService) SetServiceAvailability(v ServiceAvailabilityEnum)`

SetServiceAvailability sets ServiceAvailability field to given value.

### HasServiceAvailability

`func (o *PatchedService) HasServiceAvailability() bool`

HasServiceAvailability returns a boolean if a field has been set.

### GetAiPowered

`func (o *PatchedService) GetAiPowered() bool`

GetAiPowered returns the AiPowered field if non-nil, zero value otherwise.

### GetAiPoweredOk

`func (o *PatchedService) GetAiPoweredOk() (*bool, bool)`

GetAiPoweredOk returns a tuple with the AiPowered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiPowered

`func (o *PatchedService) SetAiPowered(v bool)`

SetAiPowered sets AiPowered field to given value.

### HasAiPowered

`func (o *PatchedService) HasAiPowered() bool`

HasAiPowered returns a boolean if a field has been set.

### GetAiSolved

`func (o *PatchedService) GetAiSolved() bool`

GetAiSolved returns the AiSolved field if non-nil, zero value otherwise.

### GetAiSolvedOk

`func (o *PatchedService) GetAiSolvedOk() (*bool, bool)`

GetAiSolvedOk returns a tuple with the AiSolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiSolved

`func (o *PatchedService) SetAiSolved(v bool)`

SetAiSolved sets AiSolved field to given value.

### HasAiSolved

`func (o *PatchedService) HasAiSolved() bool`

HasAiSolved returns a boolean if a field has been set.

### SetAiSolvedNil

`func (o *PatchedService) SetAiSolvedNil(b bool)`

 SetAiSolvedNil sets the value for AiSolved to be an explicit nil

### UnsetAiSolved
`func (o *PatchedService) UnsetAiSolved()`

UnsetAiSolved ensures that no value is present for AiSolved, not even an explicit nil
### GetDetails

`func (o *PatchedService) GetDetails() interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *PatchedService) GetDetailsOk() (*interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *PatchedService) SetDetails(v interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *PatchedService) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *PatchedService) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *PatchedService) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetServiceActionPlan

`func (o *PatchedService) GetServiceActionPlan() interface{}`

GetServiceActionPlan returns the ServiceActionPlan field if non-nil, zero value otherwise.

### GetServiceActionPlanOk

`func (o *PatchedService) GetServiceActionPlanOk() (*interface{}, bool)`

GetServiceActionPlanOk returns a tuple with the ServiceActionPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceActionPlan

`func (o *PatchedService) SetServiceActionPlan(v interface{})`

SetServiceActionPlan sets ServiceActionPlan field to given value.

### HasServiceActionPlan

`func (o *PatchedService) HasServiceActionPlan() bool`

HasServiceActionPlan returns a boolean if a field has been set.

### SetServiceActionPlanNil

`func (o *PatchedService) SetServiceActionPlanNil(b bool)`

 SetServiceActionPlanNil sets the value for ServiceActionPlan to be an explicit nil

### UnsetServiceActionPlan
`func (o *PatchedService) UnsetServiceActionPlan()`

UnsetServiceActionPlan ensures that no value is present for ServiceActionPlan, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


