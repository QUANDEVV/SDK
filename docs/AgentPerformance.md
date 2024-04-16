# AgentPerformance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentPerformanceId** | **int32** |  | [readonly] 
**TenantId** | [**TenantInfo**](TenantInfo.md) |  | [readonly] 
**AverageCallDuration** | **string** | Average duration of calls handled by the agent. | 
**AverageSatisfactionScore** | **float64** | Average satisfaction score from customer feedback. | 
**ResolutionRate** | **float64** | Percentage of calls resolved by the agent. | 
**TotalCallsHandled** | **int32** | Total number of calls handled by the agent. | 

## Methods

### NewAgentPerformance

`func NewAgentPerformance(agentPerformanceId int32, tenantId TenantInfo, averageCallDuration string, averageSatisfactionScore float64, resolutionRate float64, totalCallsHandled int32, ) *AgentPerformance`

NewAgentPerformance instantiates a new AgentPerformance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentPerformanceWithDefaults

`func NewAgentPerformanceWithDefaults() *AgentPerformance`

NewAgentPerformanceWithDefaults instantiates a new AgentPerformance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentPerformanceId

`func (o *AgentPerformance) GetAgentPerformanceId() int32`

GetAgentPerformanceId returns the AgentPerformanceId field if non-nil, zero value otherwise.

### GetAgentPerformanceIdOk

`func (o *AgentPerformance) GetAgentPerformanceIdOk() (*int32, bool)`

GetAgentPerformanceIdOk returns a tuple with the AgentPerformanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentPerformanceId

`func (o *AgentPerformance) SetAgentPerformanceId(v int32)`

SetAgentPerformanceId sets AgentPerformanceId field to given value.


### GetTenantId

`func (o *AgentPerformance) GetTenantId() TenantInfo`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AgentPerformance) GetTenantIdOk() (*TenantInfo, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AgentPerformance) SetTenantId(v TenantInfo)`

SetTenantId sets TenantId field to given value.


### SetTenantIdNil

`func (o *AgentPerformance) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *AgentPerformance) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetAverageCallDuration

`func (o *AgentPerformance) GetAverageCallDuration() string`

GetAverageCallDuration returns the AverageCallDuration field if non-nil, zero value otherwise.

### GetAverageCallDurationOk

`func (o *AgentPerformance) GetAverageCallDurationOk() (*string, bool)`

GetAverageCallDurationOk returns a tuple with the AverageCallDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageCallDuration

`func (o *AgentPerformance) SetAverageCallDuration(v string)`

SetAverageCallDuration sets AverageCallDuration field to given value.


### GetAverageSatisfactionScore

`func (o *AgentPerformance) GetAverageSatisfactionScore() float64`

GetAverageSatisfactionScore returns the AverageSatisfactionScore field if non-nil, zero value otherwise.

### GetAverageSatisfactionScoreOk

`func (o *AgentPerformance) GetAverageSatisfactionScoreOk() (*float64, bool)`

GetAverageSatisfactionScoreOk returns a tuple with the AverageSatisfactionScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageSatisfactionScore

`func (o *AgentPerformance) SetAverageSatisfactionScore(v float64)`

SetAverageSatisfactionScore sets AverageSatisfactionScore field to given value.


### GetResolutionRate

`func (o *AgentPerformance) GetResolutionRate() float64`

GetResolutionRate returns the ResolutionRate field if non-nil, zero value otherwise.

### GetResolutionRateOk

`func (o *AgentPerformance) GetResolutionRateOk() (*float64, bool)`

GetResolutionRateOk returns a tuple with the ResolutionRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionRate

`func (o *AgentPerformance) SetResolutionRate(v float64)`

SetResolutionRate sets ResolutionRate field to given value.


### GetTotalCallsHandled

`func (o *AgentPerformance) GetTotalCallsHandled() int32`

GetTotalCallsHandled returns the TotalCallsHandled field if non-nil, zero value otherwise.

### GetTotalCallsHandledOk

`func (o *AgentPerformance) GetTotalCallsHandledOk() (*int32, bool)`

GetTotalCallsHandledOk returns a tuple with the TotalCallsHandled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCallsHandled

`func (o *AgentPerformance) SetTotalCallsHandled(v int32)`

SetTotalCallsHandled sets TotalCallsHandled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


