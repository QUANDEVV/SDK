# Knowledgebase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KnowledgebaseId** | **int32** | The tenant chats ID. | [readonly] 
**TenantId** | **int32** | Display name of the tenant | 
**Name** | **string** | The name of the knowledgebase | 
**Description** | Pointer to **NullableString** | description | [optional] 
**TypeOfDatasource** | Pointer to **NullableString** |  | [optional] 
**DatasourceTarget** | Pointer to **NullableString** | This will be responsible for ft&#x3D;etching the information source passed | [optional] 
**ChatPrompts** | [**[]ChatPrompts**](ChatPrompts.md) |  | [readonly] 
**CommunityPrompts** | [**[]CommunityPrompts**](CommunityPrompts.md) |  | [readonly] 
**ServicePrompts** | [**[]ServicePrompts**](ServicePrompts.md) |  | [readonly] 
**SurveyPrompts** | [**[]SurveyPrompts**](SurveyPrompts.md) |  | [readonly] 

## Methods

### NewKnowledgebase

`func NewKnowledgebase(knowledgebaseId int32, tenantId int32, name string, chatPrompts []ChatPrompts, communityPrompts []CommunityPrompts, servicePrompts []ServicePrompts, surveyPrompts []SurveyPrompts, ) *Knowledgebase`

NewKnowledgebase instantiates a new Knowledgebase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKnowledgebaseWithDefaults

`func NewKnowledgebaseWithDefaults() *Knowledgebase`

NewKnowledgebaseWithDefaults instantiates a new Knowledgebase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKnowledgebaseId

`func (o *Knowledgebase) GetKnowledgebaseId() int32`

GetKnowledgebaseId returns the KnowledgebaseId field if non-nil, zero value otherwise.

### GetKnowledgebaseIdOk

`func (o *Knowledgebase) GetKnowledgebaseIdOk() (*int32, bool)`

GetKnowledgebaseIdOk returns a tuple with the KnowledgebaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnowledgebaseId

`func (o *Knowledgebase) SetKnowledgebaseId(v int32)`

SetKnowledgebaseId sets KnowledgebaseId field to given value.


### GetTenantId

`func (o *Knowledgebase) GetTenantId() int32`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Knowledgebase) GetTenantIdOk() (*int32, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Knowledgebase) SetTenantId(v int32)`

SetTenantId sets TenantId field to given value.


### GetName

`func (o *Knowledgebase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Knowledgebase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Knowledgebase) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Knowledgebase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Knowledgebase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Knowledgebase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Knowledgebase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Knowledgebase) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Knowledgebase) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTypeOfDatasource

`func (o *Knowledgebase) GetTypeOfDatasource() string`

GetTypeOfDatasource returns the TypeOfDatasource field if non-nil, zero value otherwise.

### GetTypeOfDatasourceOk

`func (o *Knowledgebase) GetTypeOfDatasourceOk() (*string, bool)`

GetTypeOfDatasourceOk returns a tuple with the TypeOfDatasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeOfDatasource

`func (o *Knowledgebase) SetTypeOfDatasource(v string)`

SetTypeOfDatasource sets TypeOfDatasource field to given value.

### HasTypeOfDatasource

`func (o *Knowledgebase) HasTypeOfDatasource() bool`

HasTypeOfDatasource returns a boolean if a field has been set.

### SetTypeOfDatasourceNil

`func (o *Knowledgebase) SetTypeOfDatasourceNil(b bool)`

 SetTypeOfDatasourceNil sets the value for TypeOfDatasource to be an explicit nil

### UnsetTypeOfDatasource
`func (o *Knowledgebase) UnsetTypeOfDatasource()`

UnsetTypeOfDatasource ensures that no value is present for TypeOfDatasource, not even an explicit nil
### GetDatasourceTarget

`func (o *Knowledgebase) GetDatasourceTarget() string`

GetDatasourceTarget returns the DatasourceTarget field if non-nil, zero value otherwise.

### GetDatasourceTargetOk

`func (o *Knowledgebase) GetDatasourceTargetOk() (*string, bool)`

GetDatasourceTargetOk returns a tuple with the DatasourceTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasourceTarget

`func (o *Knowledgebase) SetDatasourceTarget(v string)`

SetDatasourceTarget sets DatasourceTarget field to given value.

### HasDatasourceTarget

`func (o *Knowledgebase) HasDatasourceTarget() bool`

HasDatasourceTarget returns a boolean if a field has been set.

### SetDatasourceTargetNil

`func (o *Knowledgebase) SetDatasourceTargetNil(b bool)`

 SetDatasourceTargetNil sets the value for DatasourceTarget to be an explicit nil

### UnsetDatasourceTarget
`func (o *Knowledgebase) UnsetDatasourceTarget()`

UnsetDatasourceTarget ensures that no value is present for DatasourceTarget, not even an explicit nil
### GetChatPrompts

`func (o *Knowledgebase) GetChatPrompts() []ChatPrompts`

GetChatPrompts returns the ChatPrompts field if non-nil, zero value otherwise.

### GetChatPromptsOk

`func (o *Knowledgebase) GetChatPromptsOk() (*[]ChatPrompts, bool)`

GetChatPromptsOk returns a tuple with the ChatPrompts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatPrompts

`func (o *Knowledgebase) SetChatPrompts(v []ChatPrompts)`

SetChatPrompts sets ChatPrompts field to given value.


### GetCommunityPrompts

`func (o *Knowledgebase) GetCommunityPrompts() []CommunityPrompts`

GetCommunityPrompts returns the CommunityPrompts field if non-nil, zero value otherwise.

### GetCommunityPromptsOk

`func (o *Knowledgebase) GetCommunityPromptsOk() (*[]CommunityPrompts, bool)`

GetCommunityPromptsOk returns a tuple with the CommunityPrompts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityPrompts

`func (o *Knowledgebase) SetCommunityPrompts(v []CommunityPrompts)`

SetCommunityPrompts sets CommunityPrompts field to given value.


### GetServicePrompts

`func (o *Knowledgebase) GetServicePrompts() []ServicePrompts`

GetServicePrompts returns the ServicePrompts field if non-nil, zero value otherwise.

### GetServicePromptsOk

`func (o *Knowledgebase) GetServicePromptsOk() (*[]ServicePrompts, bool)`

GetServicePromptsOk returns a tuple with the ServicePrompts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePrompts

`func (o *Knowledgebase) SetServicePrompts(v []ServicePrompts)`

SetServicePrompts sets ServicePrompts field to given value.


### GetSurveyPrompts

`func (o *Knowledgebase) GetSurveyPrompts() []SurveyPrompts`

GetSurveyPrompts returns the SurveyPrompts field if non-nil, zero value otherwise.

### GetSurveyPromptsOk

`func (o *Knowledgebase) GetSurveyPromptsOk() (*[]SurveyPrompts, bool)`

GetSurveyPromptsOk returns a tuple with the SurveyPrompts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveyPrompts

`func (o *Knowledgebase) SetSurveyPrompts(v []SurveyPrompts)`

SetSurveyPrompts sets SurveyPrompts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


