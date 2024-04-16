# PatchedKnowledgebase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KnowledgebaseId** | Pointer to **int32** | The tenant chats ID. | [optional] [readonly] 
**TenantId** | Pointer to **int32** | Display name of the tenant | [optional] 
**Name** | Pointer to **string** | The name of the knowledgebase | [optional] 
**Description** | Pointer to **NullableString** | description | [optional] 
**TypeOfDatasource** | Pointer to **NullableString** |  | [optional] 
**DatasourceTarget** | Pointer to **NullableString** | This will be responsible for ft&#x3D;etching the information source passed | [optional] 
**ChatPrompts** | Pointer to [**[]ChatPrompts**](ChatPrompts.md) |  | [optional] [readonly] 
**CommunityPrompts** | Pointer to [**[]CommunityPrompts**](CommunityPrompts.md) |  | [optional] [readonly] 
**ServicePrompts** | Pointer to [**[]ServicePrompts**](ServicePrompts.md) |  | [optional] [readonly] 
**SurveyPrompts** | Pointer to [**[]SurveyPrompts**](SurveyPrompts.md) |  | [optional] [readonly] 

## Methods

### NewPatchedKnowledgebase

`func NewPatchedKnowledgebase() *PatchedKnowledgebase`

NewPatchedKnowledgebase instantiates a new PatchedKnowledgebase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedKnowledgebaseWithDefaults

`func NewPatchedKnowledgebaseWithDefaults() *PatchedKnowledgebase`

NewPatchedKnowledgebaseWithDefaults instantiates a new PatchedKnowledgebase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKnowledgebaseId

`func (o *PatchedKnowledgebase) GetKnowledgebaseId() int32`

GetKnowledgebaseId returns the KnowledgebaseId field if non-nil, zero value otherwise.

### GetKnowledgebaseIdOk

`func (o *PatchedKnowledgebase) GetKnowledgebaseIdOk() (*int32, bool)`

GetKnowledgebaseIdOk returns a tuple with the KnowledgebaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnowledgebaseId

`func (o *PatchedKnowledgebase) SetKnowledgebaseId(v int32)`

SetKnowledgebaseId sets KnowledgebaseId field to given value.

### HasKnowledgebaseId

`func (o *PatchedKnowledgebase) HasKnowledgebaseId() bool`

HasKnowledgebaseId returns a boolean if a field has been set.

### GetTenantId

`func (o *PatchedKnowledgebase) GetTenantId() int32`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *PatchedKnowledgebase) GetTenantIdOk() (*int32, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *PatchedKnowledgebase) SetTenantId(v int32)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *PatchedKnowledgebase) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetName

`func (o *PatchedKnowledgebase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedKnowledgebase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedKnowledgebase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedKnowledgebase) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedKnowledgebase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedKnowledgebase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedKnowledgebase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedKnowledgebase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PatchedKnowledgebase) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PatchedKnowledgebase) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTypeOfDatasource

`func (o *PatchedKnowledgebase) GetTypeOfDatasource() string`

GetTypeOfDatasource returns the TypeOfDatasource field if non-nil, zero value otherwise.

### GetTypeOfDatasourceOk

`func (o *PatchedKnowledgebase) GetTypeOfDatasourceOk() (*string, bool)`

GetTypeOfDatasourceOk returns a tuple with the TypeOfDatasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeOfDatasource

`func (o *PatchedKnowledgebase) SetTypeOfDatasource(v string)`

SetTypeOfDatasource sets TypeOfDatasource field to given value.

### HasTypeOfDatasource

`func (o *PatchedKnowledgebase) HasTypeOfDatasource() bool`

HasTypeOfDatasource returns a boolean if a field has been set.

### SetTypeOfDatasourceNil

`func (o *PatchedKnowledgebase) SetTypeOfDatasourceNil(b bool)`

 SetTypeOfDatasourceNil sets the value for TypeOfDatasource to be an explicit nil

### UnsetTypeOfDatasource
`func (o *PatchedKnowledgebase) UnsetTypeOfDatasource()`

UnsetTypeOfDatasource ensures that no value is present for TypeOfDatasource, not even an explicit nil
### GetDatasourceTarget

`func (o *PatchedKnowledgebase) GetDatasourceTarget() string`

GetDatasourceTarget returns the DatasourceTarget field if non-nil, zero value otherwise.

### GetDatasourceTargetOk

`func (o *PatchedKnowledgebase) GetDatasourceTargetOk() (*string, bool)`

GetDatasourceTargetOk returns a tuple with the DatasourceTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasourceTarget

`func (o *PatchedKnowledgebase) SetDatasourceTarget(v string)`

SetDatasourceTarget sets DatasourceTarget field to given value.

### HasDatasourceTarget

`func (o *PatchedKnowledgebase) HasDatasourceTarget() bool`

HasDatasourceTarget returns a boolean if a field has been set.

### SetDatasourceTargetNil

`func (o *PatchedKnowledgebase) SetDatasourceTargetNil(b bool)`

 SetDatasourceTargetNil sets the value for DatasourceTarget to be an explicit nil

### UnsetDatasourceTarget
`func (o *PatchedKnowledgebase) UnsetDatasourceTarget()`

UnsetDatasourceTarget ensures that no value is present for DatasourceTarget, not even an explicit nil
### GetChatPrompts

`func (o *PatchedKnowledgebase) GetChatPrompts() []ChatPrompts`

GetChatPrompts returns the ChatPrompts field if non-nil, zero value otherwise.

### GetChatPromptsOk

`func (o *PatchedKnowledgebase) GetChatPromptsOk() (*[]ChatPrompts, bool)`

GetChatPromptsOk returns a tuple with the ChatPrompts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatPrompts

`func (o *PatchedKnowledgebase) SetChatPrompts(v []ChatPrompts)`

SetChatPrompts sets ChatPrompts field to given value.

### HasChatPrompts

`func (o *PatchedKnowledgebase) HasChatPrompts() bool`

HasChatPrompts returns a boolean if a field has been set.

### GetCommunityPrompts

`func (o *PatchedKnowledgebase) GetCommunityPrompts() []CommunityPrompts`

GetCommunityPrompts returns the CommunityPrompts field if non-nil, zero value otherwise.

### GetCommunityPromptsOk

`func (o *PatchedKnowledgebase) GetCommunityPromptsOk() (*[]CommunityPrompts, bool)`

GetCommunityPromptsOk returns a tuple with the CommunityPrompts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityPrompts

`func (o *PatchedKnowledgebase) SetCommunityPrompts(v []CommunityPrompts)`

SetCommunityPrompts sets CommunityPrompts field to given value.

### HasCommunityPrompts

`func (o *PatchedKnowledgebase) HasCommunityPrompts() bool`

HasCommunityPrompts returns a boolean if a field has been set.

### GetServicePrompts

`func (o *PatchedKnowledgebase) GetServicePrompts() []ServicePrompts`

GetServicePrompts returns the ServicePrompts field if non-nil, zero value otherwise.

### GetServicePromptsOk

`func (o *PatchedKnowledgebase) GetServicePromptsOk() (*[]ServicePrompts, bool)`

GetServicePromptsOk returns a tuple with the ServicePrompts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePrompts

`func (o *PatchedKnowledgebase) SetServicePrompts(v []ServicePrompts)`

SetServicePrompts sets ServicePrompts field to given value.

### HasServicePrompts

`func (o *PatchedKnowledgebase) HasServicePrompts() bool`

HasServicePrompts returns a boolean if a field has been set.

### GetSurveyPrompts

`func (o *PatchedKnowledgebase) GetSurveyPrompts() []SurveyPrompts`

GetSurveyPrompts returns the SurveyPrompts field if non-nil, zero value otherwise.

### GetSurveyPromptsOk

`func (o *PatchedKnowledgebase) GetSurveyPromptsOk() (*[]SurveyPrompts, bool)`

GetSurveyPromptsOk returns a tuple with the SurveyPrompts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveyPrompts

`func (o *PatchedKnowledgebase) SetSurveyPrompts(v []SurveyPrompts)`

SetSurveyPrompts sets SurveyPrompts field to given value.

### HasSurveyPrompts

`func (o *PatchedKnowledgebase) HasSurveyPrompts() bool`

HasSurveyPrompts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


