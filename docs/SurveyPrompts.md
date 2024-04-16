# SurveyPrompts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Metadata** | Pointer to **interface{}** | Metadata | [optional] 
**CreatedAt** | **NullableString** |  | [readonly] 
**DateTimeCreatedAt** | **NullableString** |  | [readonly] 
**Timestamp** | **NullableTime** | The timestamp of the chat. | [readonly] 
**UpdatedAt** | **NullableTime** |  | [readonly] 
**SurveyDescription** | **string** |  | 
**Knowledgebase** | **int32** | Linking it to the knowledgebase | 

## Methods

### NewSurveyPrompts

`func NewSurveyPrompts(id int32, createdAt NullableString, dateTimeCreatedAt NullableString, timestamp NullableTime, updatedAt NullableTime, surveyDescription string, knowledgebase int32, ) *SurveyPrompts`

NewSurveyPrompts instantiates a new SurveyPrompts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSurveyPromptsWithDefaults

`func NewSurveyPromptsWithDefaults() *SurveyPrompts`

NewSurveyPromptsWithDefaults instantiates a new SurveyPrompts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SurveyPrompts) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SurveyPrompts) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SurveyPrompts) SetId(v int32)`

SetId sets Id field to given value.


### GetMetadata

`func (o *SurveyPrompts) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SurveyPrompts) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SurveyPrompts) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SurveyPrompts) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *SurveyPrompts) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *SurveyPrompts) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetCreatedAt

`func (o *SurveyPrompts) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SurveyPrompts) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SurveyPrompts) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *SurveyPrompts) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *SurveyPrompts) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDateTimeCreatedAt

`func (o *SurveyPrompts) GetDateTimeCreatedAt() string`

GetDateTimeCreatedAt returns the DateTimeCreatedAt field if non-nil, zero value otherwise.

### GetDateTimeCreatedAtOk

`func (o *SurveyPrompts) GetDateTimeCreatedAtOk() (*string, bool)`

GetDateTimeCreatedAtOk returns a tuple with the DateTimeCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeCreatedAt

`func (o *SurveyPrompts) SetDateTimeCreatedAt(v string)`

SetDateTimeCreatedAt sets DateTimeCreatedAt field to given value.


### SetDateTimeCreatedAtNil

`func (o *SurveyPrompts) SetDateTimeCreatedAtNil(b bool)`

 SetDateTimeCreatedAtNil sets the value for DateTimeCreatedAt to be an explicit nil

### UnsetDateTimeCreatedAt
`func (o *SurveyPrompts) UnsetDateTimeCreatedAt()`

UnsetDateTimeCreatedAt ensures that no value is present for DateTimeCreatedAt, not even an explicit nil
### GetTimestamp

`func (o *SurveyPrompts) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SurveyPrompts) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SurveyPrompts) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### SetTimestampNil

`func (o *SurveyPrompts) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *SurveyPrompts) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetUpdatedAt

`func (o *SurveyPrompts) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SurveyPrompts) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SurveyPrompts) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *SurveyPrompts) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *SurveyPrompts) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetSurveyDescription

`func (o *SurveyPrompts) GetSurveyDescription() string`

GetSurveyDescription returns the SurveyDescription field if non-nil, zero value otherwise.

### GetSurveyDescriptionOk

`func (o *SurveyPrompts) GetSurveyDescriptionOk() (*string, bool)`

GetSurveyDescriptionOk returns a tuple with the SurveyDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveyDescription

`func (o *SurveyPrompts) SetSurveyDescription(v string)`

SetSurveyDescription sets SurveyDescription field to given value.


### GetKnowledgebase

`func (o *SurveyPrompts) GetKnowledgebase() int32`

GetKnowledgebase returns the Knowledgebase field if non-nil, zero value otherwise.

### GetKnowledgebaseOk

`func (o *SurveyPrompts) GetKnowledgebaseOk() (*int32, bool)`

GetKnowledgebaseOk returns a tuple with the Knowledgebase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnowledgebase

`func (o *SurveyPrompts) SetKnowledgebase(v int32)`

SetKnowledgebase sets Knowledgebase field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


