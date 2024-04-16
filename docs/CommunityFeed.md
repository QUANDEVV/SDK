# CommunityFeed

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
**NumLikes** | **string** |  | [readonly] 

## Methods

### NewCommunityFeed

`func NewCommunityFeed(issueId int32, issue string, description string, clientId map[string]interface{}, communityId map[string]interface{}, createdAt NullableString, timestamp NullableTime, tags []CommunityTag, numComments int32, numUniqueUsers int32, numLikes string, ) *CommunityFeed`

NewCommunityFeed instantiates a new CommunityFeed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommunityFeedWithDefaults

`func NewCommunityFeedWithDefaults() *CommunityFeed`

NewCommunityFeedWithDefaults instantiates a new CommunityFeed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssueId

`func (o *CommunityFeed) GetIssueId() int32`

GetIssueId returns the IssueId field if non-nil, zero value otherwise.

### GetIssueIdOk

`func (o *CommunityFeed) GetIssueIdOk() (*int32, bool)`

GetIssueIdOk returns a tuple with the IssueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueId

`func (o *CommunityFeed) SetIssueId(v int32)`

SetIssueId sets IssueId field to given value.


### GetIssue

`func (o *CommunityFeed) GetIssue() string`

GetIssue returns the Issue field if non-nil, zero value otherwise.

### GetIssueOk

`func (o *CommunityFeed) GetIssueOk() (*string, bool)`

GetIssueOk returns a tuple with the Issue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssue

`func (o *CommunityFeed) SetIssue(v string)`

SetIssue sets Issue field to given value.


### GetDescription

`func (o *CommunityFeed) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CommunityFeed) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CommunityFeed) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSolved

`func (o *CommunityFeed) GetSolved() bool`

GetSolved returns the Solved field if non-nil, zero value otherwise.

### GetSolvedOk

`func (o *CommunityFeed) GetSolvedOk() (*bool, bool)`

GetSolvedOk returns a tuple with the Solved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolved

`func (o *CommunityFeed) SetSolved(v bool)`

SetSolved sets Solved field to given value.

### HasSolved

`func (o *CommunityFeed) HasSolved() bool`

HasSolved returns a boolean if a field has been set.

### GetClientId

`func (o *CommunityFeed) GetClientId() map[string]interface{}`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *CommunityFeed) GetClientIdOk() (*map[string]interface{}, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *CommunityFeed) SetClientId(v map[string]interface{})`

SetClientId sets ClientId field to given value.


### GetCommunityId

`func (o *CommunityFeed) GetCommunityId() map[string]interface{}`

GetCommunityId returns the CommunityId field if non-nil, zero value otherwise.

### GetCommunityIdOk

`func (o *CommunityFeed) GetCommunityIdOk() (*map[string]interface{}, bool)`

GetCommunityIdOk returns a tuple with the CommunityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityId

`func (o *CommunityFeed) SetCommunityId(v map[string]interface{})`

SetCommunityId sets CommunityId field to given value.


### GetCreatedAt

`func (o *CommunityFeed) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CommunityFeed) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CommunityFeed) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *CommunityFeed) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *CommunityFeed) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetTimestamp

`func (o *CommunityFeed) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CommunityFeed) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CommunityFeed) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### SetTimestampNil

`func (o *CommunityFeed) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *CommunityFeed) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTags

`func (o *CommunityFeed) GetTags() []CommunityTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CommunityFeed) GetTagsOk() (*[]CommunityTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CommunityFeed) SetTags(v []CommunityTag)`

SetTags sets Tags field to given value.


### GetNumComments

`func (o *CommunityFeed) GetNumComments() int32`

GetNumComments returns the NumComments field if non-nil, zero value otherwise.

### GetNumCommentsOk

`func (o *CommunityFeed) GetNumCommentsOk() (*int32, bool)`

GetNumCommentsOk returns a tuple with the NumComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumComments

`func (o *CommunityFeed) SetNumComments(v int32)`

SetNumComments sets NumComments field to given value.


### GetNumUniqueUsers

`func (o *CommunityFeed) GetNumUniqueUsers() int32`

GetNumUniqueUsers returns the NumUniqueUsers field if non-nil, zero value otherwise.

### GetNumUniqueUsersOk

`func (o *CommunityFeed) GetNumUniqueUsersOk() (*int32, bool)`

GetNumUniqueUsersOk returns a tuple with the NumUniqueUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUniqueUsers

`func (o *CommunityFeed) SetNumUniqueUsers(v int32)`

SetNumUniqueUsers sets NumUniqueUsers field to given value.


### GetNumLikes

`func (o *CommunityFeed) GetNumLikes() string`

GetNumLikes returns the NumLikes field if non-nil, zero value otherwise.

### GetNumLikesOk

`func (o *CommunityFeed) GetNumLikesOk() (*string, bool)`

GetNumLikesOk returns a tuple with the NumLikes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumLikes

`func (o *CommunityFeed) SetNumLikes(v string)`

SetNumLikes sets NumLikes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


