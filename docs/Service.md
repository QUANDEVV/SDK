# Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **int32** |  | [readonly] 
**TenantId** | **int32** | The tenant chats ID. | 
**NameOfService** | **string** |  | 
**ServiceDescription** | **string** |  | 
**ServiceType** | **int32** |  | 
**IsActive** | Pointer to **bool** |  | [optional] 
**ResolutionPeriod** | **string** |  | 
**ServiceAvailability** | [**ServiceAvailabilityEnum**](ServiceAvailabilityEnum.md) |  | 
**AiPowered** | Pointer to **bool** |  | [optional] 
**AiSolved** | Pointer to **NullableBool** |  | [optional] 
**Details** | **interface{}** |  | 
**ServiceActionPlan** | **interface{}** |  | 

## Methods

### NewService

`func NewService(serviceId int32, tenantId int32, nameOfService string, serviceDescription string, serviceType int32, resolutionPeriod string, serviceAvailability ServiceAvailabilityEnum, details interface{}, serviceActionPlan interface{}, ) *Service`

NewService instantiates a new Service object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceWithDefaults

`func NewServiceWithDefaults() *Service`

NewServiceWithDefaults instantiates a new Service object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *Service) GetServiceId() int32`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *Service) GetServiceIdOk() (*int32, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *Service) SetServiceId(v int32)`

SetServiceId sets ServiceId field to given value.


### GetTenantId

`func (o *Service) GetTenantId() int32`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Service) GetTenantIdOk() (*int32, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Service) SetTenantId(v int32)`

SetTenantId sets TenantId field to given value.


### GetNameOfService

`func (o *Service) GetNameOfService() string`

GetNameOfService returns the NameOfService field if non-nil, zero value otherwise.

### GetNameOfServiceOk

`func (o *Service) GetNameOfServiceOk() (*string, bool)`

GetNameOfServiceOk returns a tuple with the NameOfService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameOfService

`func (o *Service) SetNameOfService(v string)`

SetNameOfService sets NameOfService field to given value.


### GetServiceDescription

`func (o *Service) GetServiceDescription() string`

GetServiceDescription returns the ServiceDescription field if non-nil, zero value otherwise.

### GetServiceDescriptionOk

`func (o *Service) GetServiceDescriptionOk() (*string, bool)`

GetServiceDescriptionOk returns a tuple with the ServiceDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDescription

`func (o *Service) SetServiceDescription(v string)`

SetServiceDescription sets ServiceDescription field to given value.


### GetServiceType

`func (o *Service) GetServiceType() int32`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *Service) GetServiceTypeOk() (*int32, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *Service) SetServiceType(v int32)`

SetServiceType sets ServiceType field to given value.


### GetIsActive

`func (o *Service) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *Service) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *Service) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *Service) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetResolutionPeriod

`func (o *Service) GetResolutionPeriod() string`

GetResolutionPeriod returns the ResolutionPeriod field if non-nil, zero value otherwise.

### GetResolutionPeriodOk

`func (o *Service) GetResolutionPeriodOk() (*string, bool)`

GetResolutionPeriodOk returns a tuple with the ResolutionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionPeriod

`func (o *Service) SetResolutionPeriod(v string)`

SetResolutionPeriod sets ResolutionPeriod field to given value.


### GetServiceAvailability

`func (o *Service) GetServiceAvailability() ServiceAvailabilityEnum`

GetServiceAvailability returns the ServiceAvailability field if non-nil, zero value otherwise.

### GetServiceAvailabilityOk

`func (o *Service) GetServiceAvailabilityOk() (*ServiceAvailabilityEnum, bool)`

GetServiceAvailabilityOk returns a tuple with the ServiceAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAvailability

`func (o *Service) SetServiceAvailability(v ServiceAvailabilityEnum)`

SetServiceAvailability sets ServiceAvailability field to given value.


### GetAiPowered

`func (o *Service) GetAiPowered() bool`

GetAiPowered returns the AiPowered field if non-nil, zero value otherwise.

### GetAiPoweredOk

`func (o *Service) GetAiPoweredOk() (*bool, bool)`

GetAiPoweredOk returns a tuple with the AiPowered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiPowered

`func (o *Service) SetAiPowered(v bool)`

SetAiPowered sets AiPowered field to given value.

### HasAiPowered

`func (o *Service) HasAiPowered() bool`

HasAiPowered returns a boolean if a field has been set.

### GetAiSolved

`func (o *Service) GetAiSolved() bool`

GetAiSolved returns the AiSolved field if non-nil, zero value otherwise.

### GetAiSolvedOk

`func (o *Service) GetAiSolvedOk() (*bool, bool)`

GetAiSolvedOk returns a tuple with the AiSolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiSolved

`func (o *Service) SetAiSolved(v bool)`

SetAiSolved sets AiSolved field to given value.

### HasAiSolved

`func (o *Service) HasAiSolved() bool`

HasAiSolved returns a boolean if a field has been set.

### SetAiSolvedNil

`func (o *Service) SetAiSolvedNil(b bool)`

 SetAiSolvedNil sets the value for AiSolved to be an explicit nil

### UnsetAiSolved
`func (o *Service) UnsetAiSolved()`

UnsetAiSolved ensures that no value is present for AiSolved, not even an explicit nil
### GetDetails

`func (o *Service) GetDetails() interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *Service) GetDetailsOk() (*interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *Service) SetDetails(v interface{})`

SetDetails sets Details field to given value.


### SetDetailsNil

`func (o *Service) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *Service) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetServiceActionPlan

`func (o *Service) GetServiceActionPlan() interface{}`

GetServiceActionPlan returns the ServiceActionPlan field if non-nil, zero value otherwise.

### GetServiceActionPlanOk

`func (o *Service) GetServiceActionPlanOk() (*interface{}, bool)`

GetServiceActionPlanOk returns a tuple with the ServiceActionPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceActionPlan

`func (o *Service) SetServiceActionPlan(v interface{})`

SetServiceActionPlan sets ServiceActionPlan field to given value.


### SetServiceActionPlanNil

`func (o *Service) SetServiceActionPlanNil(b bool)`

 SetServiceActionPlanNil sets the value for ServiceActionPlan to be an explicit nil

### UnsetServiceActionPlan
`func (o *Service) UnsetServiceActionPlan()`

UnsetServiceActionPlan ensures that no value is present for ServiceActionPlan, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


