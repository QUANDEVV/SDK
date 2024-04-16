# ChatPrompts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Metadata** | Pointer to **interface{}** | Metadata | [optional] 
**CreatedAt** | **NullableString** |  | [readonly] 
**DateTimeCreatedAt** | **NullableString** |  | [readonly] 
**Timestamp** | **NullableTime** | The timestamp of the chat. | [readonly] 
**UpdatedAt** | **NullableTime** |  | [readonly] 
**ChatDescription** | **string** |  | 
**Knowledgebase** | **int32** | Linking it to the knowledgebase | 

## Methods

### NewChatPrompts

`func NewChatPrompts(id int32, createdAt NullableString, dateTimeCreatedAt NullableString, timestamp NullableTime, updatedAt NullableTime, chatDescription string, knowledgebase int32, ) *ChatPrompts`

NewChatPrompts instantiates a new ChatPrompts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatPromptsWithDefaults

`func NewChatPromptsWithDefaults() *ChatPrompts`

NewChatPromptsWithDefaults instantiates a new ChatPrompts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChatPrompts) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChatPrompts) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChatPrompts) SetId(v int32)`

SetId sets Id field to given value.


### GetMetadata

`func (o *ChatPrompts) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ChatPrompts) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ChatPrompts) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ChatPrompts) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *ChatPrompts) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *ChatPrompts) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetCreatedAt

`func (o *ChatPrompts) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ChatPrompts) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ChatPrompts) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *ChatPrompts) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *ChatPrompts) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDateTimeCreatedAt

`func (o *ChatPrompts) GetDateTimeCreatedAt() string`

GetDateTimeCreatedAt returns the DateTimeCreatedAt field if non-nil, zero value otherwise.

### GetDateTimeCreatedAtOk

`func (o *ChatPrompts) GetDateTimeCreatedAtOk() (*string, bool)`

GetDateTimeCreatedAtOk returns a tuple with the DateTimeCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeCreatedAt

`func (o *ChatPrompts) SetDateTimeCreatedAt(v string)`

SetDateTimeCreatedAt sets DateTimeCreatedAt field to given value.


### SetDateTimeCreatedAtNil

`func (o *ChatPrompts) SetDateTimeCreatedAtNil(b bool)`

 SetDateTimeCreatedAtNil sets the value for DateTimeCreatedAt to be an explicit nil

### UnsetDateTimeCreatedAt
`func (o *ChatPrompts) UnsetDateTimeCreatedAt()`

UnsetDateTimeCreatedAt ensures that no value is present for DateTimeCreatedAt, not even an explicit nil
### GetTimestamp

`func (o *ChatPrompts) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ChatPrompts) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ChatPrompts) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### SetTimestampNil

`func (o *ChatPrompts) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ChatPrompts) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetUpdatedAt

`func (o *ChatPrompts) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ChatPrompts) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ChatPrompts) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *ChatPrompts) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *ChatPrompts) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetChatDescription

`func (o *ChatPrompts) GetChatDescription() string`

GetChatDescription returns the ChatDescription field if non-nil, zero value otherwise.

### GetChatDescriptionOk

`func (o *ChatPrompts) GetChatDescriptionOk() (*string, bool)`

GetChatDescriptionOk returns a tuple with the ChatDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatDescription

`func (o *ChatPrompts) SetChatDescription(v string)`

SetChatDescription sets ChatDescription field to given value.


### GetKnowledgebase

`func (o *ChatPrompts) GetKnowledgebase() int32`

GetKnowledgebase returns the Knowledgebase field if non-nil, zero value otherwise.

### GetKnowledgebaseOk

`func (o *ChatPrompts) GetKnowledgebaseOk() (*int32, bool)`

GetKnowledgebaseOk returns a tuple with the Knowledgebase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnowledgebase

`func (o *ChatPrompts) SetKnowledgebase(v int32)`

SetKnowledgebase sets Knowledgebase field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


