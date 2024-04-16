# SurveySubGroups

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SurveySubgroupsId** | **int32** | The survey report id | [readonly] 
**SurveyId** | Pointer to **string** |  | [optional] 
**SubgroupName** | **string** | The survey subgroup name | 
**SubgroupDescription** | **string** | The survey subgroup description | 
**SubgroupClients** | [**[]TargetAudience**](TargetAudience.md) |  | 

## Methods

### NewSurveySubGroups

`func NewSurveySubGroups(surveySubgroupsId int32, subgroupName string, subgroupDescription string, subgroupClients []TargetAudience, ) *SurveySubGroups`

NewSurveySubGroups instantiates a new SurveySubGroups object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSurveySubGroupsWithDefaults

`func NewSurveySubGroupsWithDefaults() *SurveySubGroups`

NewSurveySubGroupsWithDefaults instantiates a new SurveySubGroups object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSurveySubgroupsId

`func (o *SurveySubGroups) GetSurveySubgroupsId() int32`

GetSurveySubgroupsId returns the SurveySubgroupsId field if non-nil, zero value otherwise.

### GetSurveySubgroupsIdOk

`func (o *SurveySubGroups) GetSurveySubgroupsIdOk() (*int32, bool)`

GetSurveySubgroupsIdOk returns a tuple with the SurveySubgroupsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveySubgroupsId

`func (o *SurveySubGroups) SetSurveySubgroupsId(v int32)`

SetSurveySubgroupsId sets SurveySubgroupsId field to given value.


### GetSurveyId

`func (o *SurveySubGroups) GetSurveyId() string`

GetSurveyId returns the SurveyId field if non-nil, zero value otherwise.

### GetSurveyIdOk

`func (o *SurveySubGroups) GetSurveyIdOk() (*string, bool)`

GetSurveyIdOk returns a tuple with the SurveyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveyId

`func (o *SurveySubGroups) SetSurveyId(v string)`

SetSurveyId sets SurveyId field to given value.

### HasSurveyId

`func (o *SurveySubGroups) HasSurveyId() bool`

HasSurveyId returns a boolean if a field has been set.

### GetSubgroupName

`func (o *SurveySubGroups) GetSubgroupName() string`

GetSubgroupName returns the SubgroupName field if non-nil, zero value otherwise.

### GetSubgroupNameOk

`func (o *SurveySubGroups) GetSubgroupNameOk() (*string, bool)`

GetSubgroupNameOk returns a tuple with the SubgroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubgroupName

`func (o *SurveySubGroups) SetSubgroupName(v string)`

SetSubgroupName sets SubgroupName field to given value.


### GetSubgroupDescription

`func (o *SurveySubGroups) GetSubgroupDescription() string`

GetSubgroupDescription returns the SubgroupDescription field if non-nil, zero value otherwise.

### GetSubgroupDescriptionOk

`func (o *SurveySubGroups) GetSubgroupDescriptionOk() (*string, bool)`

GetSubgroupDescriptionOk returns a tuple with the SubgroupDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubgroupDescription

`func (o *SurveySubGroups) SetSubgroupDescription(v string)`

SetSubgroupDescription sets SubgroupDescription field to given value.


### GetSubgroupClients

`func (o *SurveySubGroups) GetSubgroupClients() []TargetAudience`

GetSubgroupClients returns the SubgroupClients field if non-nil, zero value otherwise.

### GetSubgroupClientsOk

`func (o *SurveySubGroups) GetSubgroupClientsOk() (*[]TargetAudience, bool)`

GetSubgroupClientsOk returns a tuple with the SubgroupClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubgroupClients

`func (o *SurveySubGroups) SetSubgroupClients(v []TargetAudience)`

SetSubgroupClients sets SubgroupClients field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


