# PatchedIssue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssueId** | Pointer to **int32** | The issue ID UUID . | [optional] [readonly] 
**Issue** | Pointer to **string** | Display name of the issue that&#39;s created | [optional] 
**Description** | Pointer to **string** | Description of the issue | [optional] 
**Solved** | Pointer to **bool** | Was the issue solved or not | [optional] 
**ClientId** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**CommunityId** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**CreatedAt** | Pointer to **NullableString** |  | [optional] [readonly] 
**Timestamp** | Pointer to **NullableTime** | The timestamp of the chat. | [optional] [readonly] 
**Tags** | Pointer to [**[]CommunityTag**](CommunityTag.md) |  | [optional] 
**NumComments** | Pointer to **int32** |  | [optional] [readonly] 
**NumUniqueUsers** | Pointer to **int32** |  | [optional] [readonly] 
**Visible** | Pointer to **NullableBool** | Was the issue solved or not | [optional] 

## Methods

### NewPatchedIssue

`func NewPatchedIssue() *PatchedIssue`

NewPatchedIssue instantiates a new PatchedIssue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedIssueWithDefaults

`func NewPatchedIssueWithDefaults() *PatchedIssue`

NewPatchedIssueWithDefaults instantiates a new PatchedIssue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssueId

`func (o *PatchedIssue) GetIssueId() int32`

GetIssueId returns the IssueId field if non-nil, zero value otherwise.

### GetIssueIdOk

`func (o *PatchedIssue) GetIssueIdOk() (*int32, bool)`

GetIssueIdOk returns a tuple with the IssueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueId

`func (o *PatchedIssue) SetIssueId(v int32)`

SetIssueId sets IssueId field to given value.

### HasIssueId

`func (o *PatchedIssue) HasIssueId() bool`

HasIssueId returns a boolean if a field has been set.

### GetIssue

`func (o *PatchedIssue) GetIssue() string`

GetIssue returns the Issue field if non-nil, zero value otherwise.

### GetIssueOk

`func (o *PatchedIssue) GetIssueOk() (*string, bool)`

GetIssueOk returns a tuple with the Issue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssue

`func (o *PatchedIssue) SetIssue(v string)`

SetIssue sets Issue field to given value.

### HasIssue

`func (o *PatchedIssue) HasIssue() bool`

HasIssue returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedIssue) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedIssue) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedIssue) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedIssue) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSolved

`func (o *PatchedIssue) GetSolved() bool`

GetSolved returns the Solved field if non-nil, zero value otherwise.

### GetSolvedOk

`func (o *PatchedIssue) GetSolvedOk() (*bool, bool)`

GetSolvedOk returns a tuple with the Solved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolved

`func (o *PatchedIssue) SetSolved(v bool)`

SetSolved sets Solved field to given value.

### HasSolved

`func (o *PatchedIssue) HasSolved() bool`

HasSolved returns a boolean if a field has been set.

### GetClientId

`func (o *PatchedIssue) GetClientId() map[string]interface{}`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *PatchedIssue) GetClientIdOk() (*map[string]interface{}, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *PatchedIssue) SetClientId(v map[string]interface{})`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *PatchedIssue) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetCommunityId

`func (o *PatchedIssue) GetCommunityId() map[string]interface{}`

GetCommunityId returns the CommunityId field if non-nil, zero value otherwise.

### GetCommunityIdOk

`func (o *PatchedIssue) GetCommunityIdOk() (*map[string]interface{}, bool)`

GetCommunityIdOk returns a tuple with the CommunityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityId

`func (o *PatchedIssue) SetCommunityId(v map[string]interface{})`

SetCommunityId sets CommunityId field to given value.

### HasCommunityId

`func (o *PatchedIssue) HasCommunityId() bool`

HasCommunityId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PatchedIssue) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedIssue) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedIssue) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedIssue) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *PatchedIssue) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *PatchedIssue) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetTimestamp

`func (o *PatchedIssue) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PatchedIssue) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PatchedIssue) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *PatchedIssue) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *PatchedIssue) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *PatchedIssue) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTags

`func (o *PatchedIssue) GetTags() []CommunityTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedIssue) GetTagsOk() (*[]CommunityTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedIssue) SetTags(v []CommunityTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedIssue) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetNumComments

`func (o *PatchedIssue) GetNumComments() int32`

GetNumComments returns the NumComments field if non-nil, zero value otherwise.

### GetNumCommentsOk

`func (o *PatchedIssue) GetNumCommentsOk() (*int32, bool)`

GetNumCommentsOk returns a tuple with the NumComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumComments

`func (o *PatchedIssue) SetNumComments(v int32)`

SetNumComments sets NumComments field to given value.

### HasNumComments

`func (o *PatchedIssue) HasNumComments() bool`

HasNumComments returns a boolean if a field has been set.

### GetNumUniqueUsers

`func (o *PatchedIssue) GetNumUniqueUsers() int32`

GetNumUniqueUsers returns the NumUniqueUsers field if non-nil, zero value otherwise.

### GetNumUniqueUsersOk

`func (o *PatchedIssue) GetNumUniqueUsersOk() (*int32, bool)`

GetNumUniqueUsersOk returns a tuple with the NumUniqueUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUniqueUsers

`func (o *PatchedIssue) SetNumUniqueUsers(v int32)`

SetNumUniqueUsers sets NumUniqueUsers field to given value.

### HasNumUniqueUsers

`func (o *PatchedIssue) HasNumUniqueUsers() bool`

HasNumUniqueUsers returns a boolean if a field has been set.

### GetVisible

`func (o *PatchedIssue) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *PatchedIssue) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *PatchedIssue) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *PatchedIssue) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### SetVisibleNil

`func (o *PatchedIssue) SetVisibleNil(b bool)`

 SetVisibleNil sets the value for Visible to be an explicit nil

### UnsetVisible
`func (o *PatchedIssue) UnsetVisible()`

UnsetVisible ensures that no value is present for Visible, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


