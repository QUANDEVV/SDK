# Nested

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssueId** | **int32** | The issue ID UUID . | [readonly] 
**Metadata** | Pointer to **interface{}** | Metadata | [optional] 
**CreatedAt** | **NullableString** |  | [readonly] 
**DateTimeCreatedAt** | **NullableString** |  | [readonly] 
**Timestamp** | **NullableTime** | The timestamp of the chat. | [readonly] 
**UpdatedAt** | **NullableTime** |  | [readonly] 
**CommunityAiAssistant** | Pointer to **NullableString** |  | [optional] 
**Issue** | **string** | Display name of the issue that&#39;s created | 
**Description** | **string** | Description of the issue | 
**Solved** | Pointer to **bool** | Was the issue solved or not | [optional] 
**Visible** | Pointer to **NullableBool** | Was the issue solved or not | [optional] 
**ClientId** | Pointer to **NullableInt32** | Display name of the client | [optional] 
**UnauthenticatedClient** | Pointer to **NullableInt32** | The webhookevent chats ID UUID. | [optional] 
**CommunityId** | **int32** | Display name of the community | 
**Tags** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewNested

`func NewNested(issueId int32, createdAt NullableString, dateTimeCreatedAt NullableString, timestamp NullableTime, updatedAt NullableTime, issue string, description string, communityId int32, ) *Nested`

NewNested instantiates a new Nested object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNestedWithDefaults

`func NewNestedWithDefaults() *Nested`

NewNestedWithDefaults instantiates a new Nested object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssueId

`func (o *Nested) GetIssueId() int32`

GetIssueId returns the IssueId field if non-nil, zero value otherwise.

### GetIssueIdOk

`func (o *Nested) GetIssueIdOk() (*int32, bool)`

GetIssueIdOk returns a tuple with the IssueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueId

`func (o *Nested) SetIssueId(v int32)`

SetIssueId sets IssueId field to given value.


### GetMetadata

`func (o *Nested) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Nested) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Nested) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Nested) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *Nested) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *Nested) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetCreatedAt

`func (o *Nested) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Nested) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Nested) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *Nested) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Nested) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDateTimeCreatedAt

`func (o *Nested) GetDateTimeCreatedAt() string`

GetDateTimeCreatedAt returns the DateTimeCreatedAt field if non-nil, zero value otherwise.

### GetDateTimeCreatedAtOk

`func (o *Nested) GetDateTimeCreatedAtOk() (*string, bool)`

GetDateTimeCreatedAtOk returns a tuple with the DateTimeCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeCreatedAt

`func (o *Nested) SetDateTimeCreatedAt(v string)`

SetDateTimeCreatedAt sets DateTimeCreatedAt field to given value.


### SetDateTimeCreatedAtNil

`func (o *Nested) SetDateTimeCreatedAtNil(b bool)`

 SetDateTimeCreatedAtNil sets the value for DateTimeCreatedAt to be an explicit nil

### UnsetDateTimeCreatedAt
`func (o *Nested) UnsetDateTimeCreatedAt()`

UnsetDateTimeCreatedAt ensures that no value is present for DateTimeCreatedAt, not even an explicit nil
### GetTimestamp

`func (o *Nested) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Nested) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Nested) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### SetTimestampNil

`func (o *Nested) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *Nested) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetUpdatedAt

`func (o *Nested) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Nested) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Nested) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *Nested) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *Nested) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetCommunityAiAssistant

`func (o *Nested) GetCommunityAiAssistant() string`

GetCommunityAiAssistant returns the CommunityAiAssistant field if non-nil, zero value otherwise.

### GetCommunityAiAssistantOk

`func (o *Nested) GetCommunityAiAssistantOk() (*string, bool)`

GetCommunityAiAssistantOk returns a tuple with the CommunityAiAssistant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityAiAssistant

`func (o *Nested) SetCommunityAiAssistant(v string)`

SetCommunityAiAssistant sets CommunityAiAssistant field to given value.

### HasCommunityAiAssistant

`func (o *Nested) HasCommunityAiAssistant() bool`

HasCommunityAiAssistant returns a boolean if a field has been set.

### SetCommunityAiAssistantNil

`func (o *Nested) SetCommunityAiAssistantNil(b bool)`

 SetCommunityAiAssistantNil sets the value for CommunityAiAssistant to be an explicit nil

### UnsetCommunityAiAssistant
`func (o *Nested) UnsetCommunityAiAssistant()`

UnsetCommunityAiAssistant ensures that no value is present for CommunityAiAssistant, not even an explicit nil
### GetIssue

`func (o *Nested) GetIssue() string`

GetIssue returns the Issue field if non-nil, zero value otherwise.

### GetIssueOk

`func (o *Nested) GetIssueOk() (*string, bool)`

GetIssueOk returns a tuple with the Issue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssue

`func (o *Nested) SetIssue(v string)`

SetIssue sets Issue field to given value.


### GetDescription

`func (o *Nested) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Nested) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Nested) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSolved

`func (o *Nested) GetSolved() bool`

GetSolved returns the Solved field if non-nil, zero value otherwise.

### GetSolvedOk

`func (o *Nested) GetSolvedOk() (*bool, bool)`

GetSolvedOk returns a tuple with the Solved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolved

`func (o *Nested) SetSolved(v bool)`

SetSolved sets Solved field to given value.

### HasSolved

`func (o *Nested) HasSolved() bool`

HasSolved returns a boolean if a field has been set.

### GetVisible

`func (o *Nested) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *Nested) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *Nested) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *Nested) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### SetVisibleNil

`func (o *Nested) SetVisibleNil(b bool)`

 SetVisibleNil sets the value for Visible to be an explicit nil

### UnsetVisible
`func (o *Nested) UnsetVisible()`

UnsetVisible ensures that no value is present for Visible, not even an explicit nil
### GetClientId

`func (o *Nested) GetClientId() int32`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Nested) GetClientIdOk() (*int32, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Nested) SetClientId(v int32)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *Nested) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### SetClientIdNil

`func (o *Nested) SetClientIdNil(b bool)`

 SetClientIdNil sets the value for ClientId to be an explicit nil

### UnsetClientId
`func (o *Nested) UnsetClientId()`

UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
### GetUnauthenticatedClient

`func (o *Nested) GetUnauthenticatedClient() int32`

GetUnauthenticatedClient returns the UnauthenticatedClient field if non-nil, zero value otherwise.

### GetUnauthenticatedClientOk

`func (o *Nested) GetUnauthenticatedClientOk() (*int32, bool)`

GetUnauthenticatedClientOk returns a tuple with the UnauthenticatedClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnauthenticatedClient

`func (o *Nested) SetUnauthenticatedClient(v int32)`

SetUnauthenticatedClient sets UnauthenticatedClient field to given value.

### HasUnauthenticatedClient

`func (o *Nested) HasUnauthenticatedClient() bool`

HasUnauthenticatedClient returns a boolean if a field has been set.

### SetUnauthenticatedClientNil

`func (o *Nested) SetUnauthenticatedClientNil(b bool)`

 SetUnauthenticatedClientNil sets the value for UnauthenticatedClient to be an explicit nil

### UnsetUnauthenticatedClient
`func (o *Nested) UnsetUnauthenticatedClient()`

UnsetUnauthenticatedClient ensures that no value is present for UnauthenticatedClient, not even an explicit nil
### GetCommunityId

`func (o *Nested) GetCommunityId() int32`

GetCommunityId returns the CommunityId field if non-nil, zero value otherwise.

### GetCommunityIdOk

`func (o *Nested) GetCommunityIdOk() (*int32, bool)`

GetCommunityIdOk returns a tuple with the CommunityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityId

`func (o *Nested) SetCommunityId(v int32)`

SetCommunityId sets CommunityId field to given value.


### GetTags

`func (o *Nested) GetTags() []int32`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Nested) GetTagsOk() (*[]int32, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Nested) SetTags(v []int32)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Nested) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


