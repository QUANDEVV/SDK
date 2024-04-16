# Comment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommentId** | **int32** | The chat ID UUID for an instance of a issue. | [readonly] 
**Thread** | **int32** | Display name of the Thread | 
**Client** | [**ClientLikes**](ClientLikes.md) |  | 
**CommentDescription** | Pointer to **NullableString** | Description of the comment | [optional] 
**Likes** | [**[]ClientLikes**](ClientLikes.md) |  | 
**Dislikes** | Pointer to **[]int32** | Users who liked the comment | [optional] 

## Methods

### NewComment

`func NewComment(commentId int32, thread int32, client ClientLikes, likes []ClientLikes, ) *Comment`

NewComment instantiates a new Comment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentWithDefaults

`func NewCommentWithDefaults() *Comment`

NewCommentWithDefaults instantiates a new Comment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommentId

`func (o *Comment) GetCommentId() int32`

GetCommentId returns the CommentId field if non-nil, zero value otherwise.

### GetCommentIdOk

`func (o *Comment) GetCommentIdOk() (*int32, bool)`

GetCommentIdOk returns a tuple with the CommentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentId

`func (o *Comment) SetCommentId(v int32)`

SetCommentId sets CommentId field to given value.


### GetThread

`func (o *Comment) GetThread() int32`

GetThread returns the Thread field if non-nil, zero value otherwise.

### GetThreadOk

`func (o *Comment) GetThreadOk() (*int32, bool)`

GetThreadOk returns a tuple with the Thread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThread

`func (o *Comment) SetThread(v int32)`

SetThread sets Thread field to given value.


### GetClient

`func (o *Comment) GetClient() ClientLikes`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *Comment) GetClientOk() (*ClientLikes, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *Comment) SetClient(v ClientLikes)`

SetClient sets Client field to given value.


### GetCommentDescription

`func (o *Comment) GetCommentDescription() string`

GetCommentDescription returns the CommentDescription field if non-nil, zero value otherwise.

### GetCommentDescriptionOk

`func (o *Comment) GetCommentDescriptionOk() (*string, bool)`

GetCommentDescriptionOk returns a tuple with the CommentDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentDescription

`func (o *Comment) SetCommentDescription(v string)`

SetCommentDescription sets CommentDescription field to given value.

### HasCommentDescription

`func (o *Comment) HasCommentDescription() bool`

HasCommentDescription returns a boolean if a field has been set.

### SetCommentDescriptionNil

`func (o *Comment) SetCommentDescriptionNil(b bool)`

 SetCommentDescriptionNil sets the value for CommentDescription to be an explicit nil

### UnsetCommentDescription
`func (o *Comment) UnsetCommentDescription()`

UnsetCommentDescription ensures that no value is present for CommentDescription, not even an explicit nil
### GetLikes

`func (o *Comment) GetLikes() []ClientLikes`

GetLikes returns the Likes field if non-nil, zero value otherwise.

### GetLikesOk

`func (o *Comment) GetLikesOk() (*[]ClientLikes, bool)`

GetLikesOk returns a tuple with the Likes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikes

`func (o *Comment) SetLikes(v []ClientLikes)`

SetLikes sets Likes field to given value.


### GetDislikes

`func (o *Comment) GetDislikes() []int32`

GetDislikes returns the Dislikes field if non-nil, zero value otherwise.

### GetDislikesOk

`func (o *Comment) GetDislikesOk() (*[]int32, bool)`

GetDislikesOk returns a tuple with the Dislikes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDislikes

`func (o *Comment) SetDislikes(v []int32)`

SetDislikes sets Dislikes field to given value.

### HasDislikes

`func (o *Comment) HasDislikes() bool`

HasDislikes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


