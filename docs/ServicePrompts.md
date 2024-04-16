# ServicePrompts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Metadata** | Pointer to **interface{}** | Metadata | [optional] 
**CreatedAt** | **NullableString** |  | [readonly] 
**DateTimeCreatedAt** | **NullableString** |  | [readonly] 
**Timestamp** | **NullableTime** | The timestamp of the chat. | [readonly] 
**UpdatedAt** | **NullableTime** |  | [readonly] 
**ServiceDescription** | **string** |  | 
**Service** | **int32** | To tie services knowledge base to it | 
**Knowledgebase** | **int32** | Linking it to the knowledgebase | 

## Methods

### NewServicePrompts

`func NewServicePrompts(id int32, createdAt NullableString, dateTimeCreatedAt NullableString, timestamp NullableTime, updatedAt NullableTime, serviceDescription string, service int32, knowledgebase int32, ) *ServicePrompts`

NewServicePrompts instantiates a new ServicePrompts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicePromptsWithDefaults

`func NewServicePromptsWithDefaults() *ServicePrompts`

NewServicePromptsWithDefaults instantiates a new ServicePrompts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServicePrompts) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServicePrompts) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServicePrompts) SetId(v int32)`

SetId sets Id field to given value.


### GetMetadata

`func (o *ServicePrompts) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ServicePrompts) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ServicePrompts) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ServicePrompts) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *ServicePrompts) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *ServicePrompts) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetCreatedAt

`func (o *ServicePrompts) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServicePrompts) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServicePrompts) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *ServicePrompts) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *ServicePrompts) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDateTimeCreatedAt

`func (o *ServicePrompts) GetDateTimeCreatedAt() string`

GetDateTimeCreatedAt returns the DateTimeCreatedAt field if non-nil, zero value otherwise.

### GetDateTimeCreatedAtOk

`func (o *ServicePrompts) GetDateTimeCreatedAtOk() (*string, bool)`

GetDateTimeCreatedAtOk returns a tuple with the DateTimeCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeCreatedAt

`func (o *ServicePrompts) SetDateTimeCreatedAt(v string)`

SetDateTimeCreatedAt sets DateTimeCreatedAt field to given value.


### SetDateTimeCreatedAtNil

`func (o *ServicePrompts) SetDateTimeCreatedAtNil(b bool)`

 SetDateTimeCreatedAtNil sets the value for DateTimeCreatedAt to be an explicit nil

### UnsetDateTimeCreatedAt
`func (o *ServicePrompts) UnsetDateTimeCreatedAt()`

UnsetDateTimeCreatedAt ensures that no value is present for DateTimeCreatedAt, not even an explicit nil
### GetTimestamp

`func (o *ServicePrompts) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ServicePrompts) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ServicePrompts) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### SetTimestampNil

`func (o *ServicePrompts) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ServicePrompts) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetUpdatedAt

`func (o *ServicePrompts) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ServicePrompts) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ServicePrompts) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *ServicePrompts) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *ServicePrompts) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetServiceDescription

`func (o *ServicePrompts) GetServiceDescription() string`

GetServiceDescription returns the ServiceDescription field if non-nil, zero value otherwise.

### GetServiceDescriptionOk

`func (o *ServicePrompts) GetServiceDescriptionOk() (*string, bool)`

GetServiceDescriptionOk returns a tuple with the ServiceDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDescription

`func (o *ServicePrompts) SetServiceDescription(v string)`

SetServiceDescription sets ServiceDescription field to given value.


### GetService

`func (o *ServicePrompts) GetService() int32`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ServicePrompts) GetServiceOk() (*int32, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ServicePrompts) SetService(v int32)`

SetService sets Service field to given value.


### GetKnowledgebase

`func (o *ServicePrompts) GetKnowledgebase() int32`

GetKnowledgebase returns the Knowledgebase field if non-nil, zero value otherwise.

### GetKnowledgebaseOk

`func (o *ServicePrompts) GetKnowledgebaseOk() (*int32, bool)`

GetKnowledgebaseOk returns a tuple with the Knowledgebase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnowledgebase

`func (o *ServicePrompts) SetKnowledgebase(v int32)`

SetKnowledgebase sets Knowledgebase field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


