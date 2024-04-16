# Issue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssueId** | **int32** | The issue ID UUID . | [readonly] 
**Issue** | **string** | Display name of the issue that&#39;s created | 
**Description** | **string** | Description of the issue | 
**Solved** | Pointer to **bool** | Was the issue solved or not | [optional] 
**ClientId** | **map[string]interface{}** |  | [readonly] 
**CommunityId** | **map[string]interface{}** |  | [readonly] 
**CreatedAt** | **NullableString** |  | [readonly] 
**Timestamp** | **NullableTime** | The timestamp of the chat. | [readonly] 
**Tags** | [**[]CommunityTag**](CommunityTag.md) |  | 
**NumComments** | **int32** |  | [readonly] 
**NumUniqueUsers** | **int32** |  | [readonly] 
**Visible** | Pointer to **NullableBool** | Was the issue solved or not | [optional] 

## Methods

### NewIssue

`func NewIssue(issueId int32, issue string, description string, clientId map[string]interface{}, communityId map[string]interface{}, createdAt NullableString, timestamp NullableTime, tags []CommunityTag, numComments int32, numUniqueUsers int32, ) *Issue`

NewIssue instantiates a new Issue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueWithDefaults

`func NewIssueWithDefaults() *Issue`

NewIssueWithDefaults instantiates a new Issue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssueId

`func (o *Issue) GetIssueId() int32`

GetIssueId returns the IssueId field if non-nil, zero value otherwise.

### GetIssueIdOk

`func (o *Issue) GetIssueIdOk() (*int32, bool)`

GetIssueIdOk returns a tuple with the IssueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueId

`func (o *Issue) SetIssueId(v int32)`

SetIssueId sets IssueId field to given value.


### GetIssue

`func (o *Issue) GetIssue() string`

GetIssue returns the Issue field if non-nil, zero value otherwise.

### GetIssueOk

`func (o *Issue) GetIssueOk() (*string, bool)`

GetIssueOk returns a tuple with the Issue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssue

`func (o *Issue) SetIssue(v string)`

SetIssue sets Issue field to given value.


### GetDescription

`func (o *Issue) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Issue) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Issue) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSolved

`func (o *Issue) GetSolved() bool`

GetSolved returns the Solved field if non-nil, zero value otherwise.

### GetSolvedOk

`func (o *Issue) GetSolvedOk() (*bool, bool)`

GetSolvedOk returns a tuple with the Solved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolved

`func (o *Issue) SetSolved(v bool)`

SetSolved sets Solved field to given value.

### HasSolved

`func (o *Issue) HasSolved() bool`

HasSolved returns a boolean if a field has been set.

### GetClientId

`func (o *Issue) GetClientId() map[string]interface{}`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Issue) GetClientIdOk() (*map[string]interface{}, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Issue) SetClientId(v map[string]interface{})`

SetClientId sets ClientId field to given value.


### GetCommunityId

`func (o *Issue) GetCommunityId() map[string]interface{}`

GetCommunityId returns the CommunityId field if non-nil, zero value otherwise.

### GetCommunityIdOk

`func (o *Issue) GetCommunityIdOk() (*map[string]interface{}, bool)`

GetCommunityIdOk returns a tuple with the CommunityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityId

`func (o *Issue) SetCommunityId(v map[string]interface{})`

SetCommunityId sets CommunityId field to given value.


### GetCreatedAt

`func (o *Issue) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Issue) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Issue) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *Issue) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Issue) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetTimestamp

`func (o *Issue) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Issue) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Issue) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### SetTimestampNil

`func (o *Issue) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *Issue) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTags

`func (o *Issue) GetTags() []CommunityTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Issue) GetTagsOk() (*[]CommunityTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Issue) SetTags(v []CommunityTag)`

SetTags sets Tags field to given value.


### GetNumComments

`func (o *Issue) GetNumComments() int32`

GetNumComments returns the NumComments field if non-nil, zero value otherwise.

### GetNumCommentsOk

`func (o *Issue) GetNumCommentsOk() (*int32, bool)`

GetNumCommentsOk returns a tuple with the NumComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumComments

`func (o *Issue) SetNumComments(v int32)`

SetNumComments sets NumComments field to given value.


### GetNumUniqueUsers

`func (o *Issue) GetNumUniqueUsers() int32`

GetNumUniqueUsers returns the NumUniqueUsers field if non-nil, zero value otherwise.

### GetNumUniqueUsersOk

`func (o *Issue) GetNumUniqueUsersOk() (*int32, bool)`

GetNumUniqueUsersOk returns a tuple with the NumUniqueUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUniqueUsers

`func (o *Issue) SetNumUniqueUsers(v int32)`

SetNumUniqueUsers sets NumUniqueUsers field to given value.


### GetVisible

`func (o *Issue) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *Issue) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *Issue) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *Issue) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### SetVisibleNil

`func (o *Issue) SetVisibleNil(b bool)`

 SetVisibleNil sets the value for Visible to be an explicit nil

### UnsetVisible
`func (o *Issue) UnsetVisible()`

UnsetVisible ensures that no value is present for Visible, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


