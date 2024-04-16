/*
proxima-core-engine

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// CommunityFeed struct for CommunityFeed
type CommunityFeed struct {
	// The issue ID UUID .
	IssueId int32 `json:"issue_id"`
	// Display name of the issue that's created
	Issue string `json:"issue"`
	// Description of the issue
	Description string `json:"description"`
	// Was the issue solved or not
	Solved *bool `json:"solved,omitempty"`
	ClientId map[string]interface{} `json:"client_id"`
	CommunityId map[string]interface{} `json:"community_id"`
	CreatedAt NullableString `json:"created_at"`
	// The timestamp of the chat.
	Timestamp NullableTime `json:"timestamp"`
	Tags []CommunityTag `json:"tags"`
	NumComments int32 `json:"num_comments"`
	NumUniqueUsers int32 `json:"num_unique_users"`
	NumLikes string `json:"num_likes"`
}

// NewCommunityFeed instantiates a new CommunityFeed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommunityFeed(issueId int32, issue string, description string, clientId map[string]interface{}, communityId map[string]interface{}, createdAt NullableString, timestamp NullableTime, tags []CommunityTag, numComments int32, numUniqueUsers int32, numLikes string) *CommunityFeed {
	this := CommunityFeed{}
	this.IssueId = issueId
	this.Issue = issue
	this.Description = description
	this.ClientId = clientId
	this.CommunityId = communityId
	this.CreatedAt = createdAt
	this.Timestamp = timestamp
	this.Tags = tags
	this.NumComments = numComments
	this.NumUniqueUsers = numUniqueUsers
	this.NumLikes = numLikes
	return &this
}

// NewCommunityFeedWithDefaults instantiates a new CommunityFeed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommunityFeedWithDefaults() *CommunityFeed {
	this := CommunityFeed{}
	return &this
}

// GetIssueId returns the IssueId field value
func (o *CommunityFeed) GetIssueId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IssueId
}

// GetIssueIdOk returns a tuple with the IssueId field value
// and a boolean to check if the value has been set.
func (o *CommunityFeed) GetIssueIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IssueId, true
}

// SetIssueId sets field value
func (o *CommunityFeed) SetIssueId(v int32) {
	o.IssueId = v
}

// GetIssue returns the Issue field value
func (o *CommunityFeed) GetIssue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Issue
}

// GetIssueOk returns a tuple with the Issue field value
// and a boolean to check if the value has been set.
func (o *CommunityFeed) GetIssueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Issue, true
}

// SetIssue sets field value
func (o *CommunityFeed) SetIssue(v string) {
	o.Issue = v
}

// GetDescription returns the Description field value
func (o *CommunityFeed) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CommunityFeed) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CommunityFeed) SetDescription(v string) {
	o.Description = v
}

// GetSolved returns the Solved field value if set, zero value otherwise.
func (o *CommunityFeed) GetSolved() bool {
	if o == nil || o.Solved == nil {
		var ret bool
		return ret
	}
	return *o.Solved
}

// GetSolvedOk returns a tuple with the Solved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunityFeed) GetSolvedOk() (*bool, bool) {
	if o == nil || o.Solved == nil {
		return nil, false
	}
	return o.Solved, true
}

// HasSolved returns a boolean if a field has been set.
func (o *CommunityFeed) HasSolved() bool {
	if o != nil && o.Solved != nil {
		return true
	}

	return false
}

// SetSolved gets a reference to the given bool and assigns it to the Solved field.
func (o *CommunityFeed) SetSolved(v bool) {
	o.Solved = &v
}

// GetClientId returns the ClientId field value
func (o *CommunityFeed) GetClientId() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *CommunityFeed) GetClientIdOk() (*map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *CommunityFeed) SetClientId(v map[string]interface{}) {
	o.ClientId = v
}

// GetCommunityId returns the CommunityId field value
func (o *CommunityFeed) GetCommunityId() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.CommunityId
}

// GetCommunityIdOk returns a tuple with the CommunityId field value
// and a boolean to check if the value has been set.
func (o *CommunityFeed) GetCommunityIdOk() (*map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CommunityId, true
}

// SetCommunityId sets field value
func (o *CommunityFeed) SetCommunityId(v map[string]interface{}) {
	o.CommunityId = v
}

// GetCreatedAt returns the CreatedAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CommunityFeed) GetCreatedAt() string {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommunityFeed) GetCreatedAtOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// SetCreatedAt sets field value
func (o *CommunityFeed) SetCreatedAt(v string) {
	o.CreatedAt.Set(&v)
}

// GetTimestamp returns the Timestamp field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *CommunityFeed) GetTimestamp() time.Time {
	if o == nil || o.Timestamp.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Timestamp.Get()
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommunityFeed) GetTimestampOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Timestamp.Get(), o.Timestamp.IsSet()
}

// SetTimestamp sets field value
func (o *CommunityFeed) SetTimestamp(v time.Time) {
	o.Timestamp.Set(&v)
}

// GetTags returns the Tags field value
func (o *CommunityFeed) GetTags() []CommunityTag {
	if o == nil {
		var ret []CommunityTag
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *CommunityFeed) GetTagsOk() (*[]CommunityTag, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value
func (o *CommunityFeed) SetTags(v []CommunityTag) {
	o.Tags = v
}

// GetNumComments returns the NumComments field value
func (o *CommunityFeed) GetNumComments() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumComments
}

// GetNumCommentsOk returns a tuple with the NumComments field value
// and a boolean to check if the value has been set.
func (o *CommunityFeed) GetNumCommentsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NumComments, true
}

// SetNumComments sets field value
func (o *CommunityFeed) SetNumComments(v int32) {
	o.NumComments = v
}

// GetNumUniqueUsers returns the NumUniqueUsers field value
func (o *CommunityFeed) GetNumUniqueUsers() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumUniqueUsers
}

// GetNumUniqueUsersOk returns a tuple with the NumUniqueUsers field value
// and a boolean to check if the value has been set.
func (o *CommunityFeed) GetNumUniqueUsersOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NumUniqueUsers, true
}

// SetNumUniqueUsers sets field value
func (o *CommunityFeed) SetNumUniqueUsers(v int32) {
	o.NumUniqueUsers = v
}

// GetNumLikes returns the NumLikes field value
func (o *CommunityFeed) GetNumLikes() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NumLikes
}

// GetNumLikesOk returns a tuple with the NumLikes field value
// and a boolean to check if the value has been set.
func (o *CommunityFeed) GetNumLikesOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NumLikes, true
}

// SetNumLikes sets field value
func (o *CommunityFeed) SetNumLikes(v string) {
	o.NumLikes = v
}

func (o CommunityFeed) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["issue_id"] = o.IssueId
	}
	if true {
		toSerialize["issue"] = o.Issue
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if o.Solved != nil {
		toSerialize["solved"] = o.Solved
	}
	if true {
		toSerialize["client_id"] = o.ClientId
	}
	if true {
		toSerialize["community_id"] = o.CommunityId
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if true {
		toSerialize["timestamp"] = o.Timestamp.Get()
	}
	if true {
		toSerialize["tags"] = o.Tags
	}
	if true {
		toSerialize["num_comments"] = o.NumComments
	}
	if true {
		toSerialize["num_unique_users"] = o.NumUniqueUsers
	}
	if true {
		toSerialize["num_likes"] = o.NumLikes
	}
	return json.Marshal(toSerialize)
}

type NullableCommunityFeed struct {
	value *CommunityFeed
	isSet bool
}

func (v NullableCommunityFeed) Get() *CommunityFeed {
	return v.value
}

func (v *NullableCommunityFeed) Set(val *CommunityFeed) {
	v.value = val
	v.isSet = true
}

func (v NullableCommunityFeed) IsSet() bool {
	return v.isSet
}

func (v *NullableCommunityFeed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommunityFeed(val *CommunityFeed) *NullableCommunityFeed {
	return &NullableCommunityFeed{value: val, isSet: true}
}

func (v NullableCommunityFeed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommunityFeed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


