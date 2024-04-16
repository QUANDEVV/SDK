# PatchedAgentPerformance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentPerformanceId** | Pointer to **int32** |  | [optional] [readonly] 
**TenantId** | Pointer to [**TenantInfo**](TenantInfo.md) |  | [optional] [readonly] 
**AverageCallDuration** | Pointer to **string** | Average duration of calls handled by the agent. | [optional] 
**AverageSatisfactionScore** | Pointer to **float64** | Average satisfaction score from customer feedback. | [optional] 
**ResolutionRate** | Pointer to **float64** | Percentage of calls resolved by the agent. | [optional] 
**TotalCallsHandled** | Pointer to **int32** | Total number of calls handled by the agent. | [optional] 

## Methods

### NewPatchedAgentPerformance

`func NewPatchedAgentPerformance() *PatchedAgentPerformance`

NewPatchedAgentPerformance instantiates a new PatchedAgentPerformance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedAgentPerformanceWithDefaults

`func NewPatchedAgentPerformanceWithDefaults() *PatchedAgentPerformance`

NewPatchedAgentPerformanceWithDefaults instantiates a new PatchedAgentPerformance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentPerformanceId

`func (o *PatchedAgentPerformance) GetAgentPerformanceId() int32`

GetAgentPerformanceId returns the AgentPerformanceId field if non-nil, zero value otherwise.

### GetAgentPerformanceIdOk

`func (o *PatchedAgentPerformance) GetAgentPerformanceIdOk() (*int32, bool)`

GetAgentPerformanceIdOk returns a tuple with the AgentPerformanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentPerformanceId

`func (o *PatchedAgentPerformance) SetAgentPerformanceId(v int32)`

SetAgentPerformanceId sets AgentPerformanceId field to given value.

### HasAgentPerformanceId

`func (o *PatchedAgentPerformance) HasAgentPerformanceId() bool`

HasAgentPerformanceId returns a boolean if a field has been set.

### GetTenantId

`func (o *PatchedAgentPerformance) GetTenantId() TenantInfo`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *PatchedAgentPerformance) GetTenantIdOk() (*TenantInfo, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *PatchedAgentPerformance) SetTenantId(v TenantInfo)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *PatchedAgentPerformance) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *PatchedAgentPerformance) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *PatchedAgentPerformance) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetAverageCallDuration

`func (o *PatchedAgentPerformance) GetAverageCallDuration() string`

GetAverageCallDuration returns the AverageCallDuration field if non-nil, zero value otherwise.

### GetAverageCallDurationOk

`func (o *PatchedAgentPerformance) GetAverageCallDurationOk() (*string, bool)`

GetAverageCallDurationOk returns a tuple with the AverageCallDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageCallDuration

`func (o *PatchedAgentPerformance) SetAverageCallDuration(v string)`

SetAverageCallDuration sets AverageCallDuration field to given value.

### HasAverageCallDuration

`func (o *PatchedAgentPerformance) HasAverageCallDuration() bool`

HasAverageCallDuration returns a boolean if a field has been set.

### GetAverageSatisfactionScore

`func (o *PatchedAgentPerformance) GetAverageSatisfactionScore() float64`

GetAverageSatisfactionScore returns the AverageSatisfactionScore field if non-nil, zero value otherwise.

### GetAverageSatisfactionScoreOk

`func (o *PatchedAgentPerformance) GetAverageSatisfactionScoreOk() (*float64, bool)`

GetAverageSatisfactionScoreOk returns a tuple with the AverageSatisfactionScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageSatisfactionScore

`func (o *PatchedAgentPerformance) SetAverageSatisfactionScore(v float64)`

SetAverageSatisfactionScore sets AverageSatisfactionScore field to given value.

### HasAverageSatisfactionScore

`func (o *PatchedAgentPerformance) HasAverageSatisfactionScore() bool`

HasAverageSatisfactionScore returns a boolean if a field has been set.

### GetResolutionRate

`func (o *PatchedAgentPerformance) GetResolutionRate() float64`

GetResolutionRate returns the ResolutionRate field if non-nil, zero value otherwise.

### GetResolutionRateOk

`func (o *PatchedAgentPerformance) GetResolutionRateOk() (*float64, bool)`

GetResolutionRateOk returns a tuple with the ResolutionRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionRate

`func (o *PatchedAgentPerformance) SetResolutionRate(v float64)`

SetResolutionRate sets ResolutionRate field to given value.

### HasResolutionRate

`func (o *PatchedAgentPerformance) HasResolutionRate() bool`

HasResolutionRate returns a boolean if a field has been set.

### GetTotalCallsHandled

`func (o *PatchedAgentPerformance) GetTotalCallsHandled() int32`

GetTotalCallsHandled returns the TotalCallsHandled field if non-nil, zero value otherwise.

### GetTotalCallsHandledOk

`func (o *PatchedAgentPerformance) GetTotalCallsHandledOk() (*int32, bool)`

GetTotalCallsHandledOk returns a tuple with the TotalCallsHandled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCallsHandled

`func (o *PatchedAgentPerformance) SetTotalCallsHandled(v int32)`

SetTotalCallsHandled sets TotalCallsHandled field to given value.

### HasTotalCallsHandled

`func (o *PatchedAgentPerformance) HasTotalCallsHandled() bool`

HasTotalCallsHandled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


