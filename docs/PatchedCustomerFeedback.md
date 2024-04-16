# PatchedCustomerFeedback

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerFeedbackId** | Pointer to **int32** |  | [optional] [readonly] 
**Call** | Pointer to [**Call**](Call.md) |  | [optional] [readonly] 
**Rating** | Pointer to **int32** | Rating on a scale of 1 to 5. | [optional] 
**FeedbackText** | Pointer to **NullableString** | Detailed text feedback from the customer. | [optional] 
**FeedbackCategory** | Pointer to [**FeedbackCategoryEnum**](FeedbackCategoryEnum.md) | Category of the feedback.  * &#x60;SERVICE&#x60; - Service * &#x60;TECHNICAL&#x60; - Technical Issue * &#x60;GENERAL&#x60; - General Feedback | [optional] 

## Methods

### NewPatchedCustomerFeedback

`func NewPatchedCustomerFeedback() *PatchedCustomerFeedback`

NewPatchedCustomerFeedback instantiates a new PatchedCustomerFeedback object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedCustomerFeedbackWithDefaults

`func NewPatchedCustomerFeedbackWithDefaults() *PatchedCustomerFeedback`

NewPatchedCustomerFeedbackWithDefaults instantiates a new PatchedCustomerFeedback object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerFeedbackId

`func (o *PatchedCustomerFeedback) GetCustomerFeedbackId() int32`

GetCustomerFeedbackId returns the CustomerFeedbackId field if non-nil, zero value otherwise.

### GetCustomerFeedbackIdOk

`func (o *PatchedCustomerFeedback) GetCustomerFeedbackIdOk() (*int32, bool)`

GetCustomerFeedbackIdOk returns a tuple with the CustomerFeedbackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerFeedbackId

`func (o *PatchedCustomerFeedback) SetCustomerFeedbackId(v int32)`

SetCustomerFeedbackId sets CustomerFeedbackId field to given value.

### HasCustomerFeedbackId

`func (o *PatchedCustomerFeedback) HasCustomerFeedbackId() bool`

HasCustomerFeedbackId returns a boolean if a field has been set.

### GetCall

`func (o *PatchedCustomerFeedback) GetCall() Call`

GetCall returns the Call field if non-nil, zero value otherwise.

### GetCallOk

`func (o *PatchedCustomerFeedback) GetCallOk() (*Call, bool)`

GetCallOk returns a tuple with the Call field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCall

`func (o *PatchedCustomerFeedback) SetCall(v Call)`

SetCall sets Call field to given value.

### HasCall

`func (o *PatchedCustomerFeedback) HasCall() bool`

HasCall returns a boolean if a field has been set.

### SetCallNil

`func (o *PatchedCustomerFeedback) SetCallNil(b bool)`

 SetCallNil sets the value for Call to be an explicit nil

### UnsetCall
`func (o *PatchedCustomerFeedback) UnsetCall()`

UnsetCall ensures that no value is present for Call, not even an explicit nil
### GetRating

`func (o *PatchedCustomerFeedback) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *PatchedCustomerFeedback) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *PatchedCustomerFeedback) SetRating(v int32)`

SetRating sets Rating field to given value.

### HasRating

`func (o *PatchedCustomerFeedback) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetFeedbackText

`func (o *PatchedCustomerFeedback) GetFeedbackText() string`

GetFeedbackText returns the FeedbackText field if non-nil, zero value otherwise.

### GetFeedbackTextOk

`func (o *PatchedCustomerFeedback) GetFeedbackTextOk() (*string, bool)`

GetFeedbackTextOk returns a tuple with the FeedbackText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackText

`func (o *PatchedCustomerFeedback) SetFeedbackText(v string)`

SetFeedbackText sets FeedbackText field to given value.

### HasFeedbackText

`func (o *PatchedCustomerFeedback) HasFeedbackText() bool`

HasFeedbackText returns a boolean if a field has been set.

### SetFeedbackTextNil

`func (o *PatchedCustomerFeedback) SetFeedbackTextNil(b bool)`

 SetFeedbackTextNil sets the value for FeedbackText to be an explicit nil

### UnsetFeedbackText
`func (o *PatchedCustomerFeedback) UnsetFeedbackText()`

UnsetFeedbackText ensures that no value is present for FeedbackText, not even an explicit nil
### GetFeedbackCategory

`func (o *PatchedCustomerFeedback) GetFeedbackCategory() FeedbackCategoryEnum`

GetFeedbackCategory returns the FeedbackCategory field if non-nil, zero value otherwise.

### GetFeedbackCategoryOk

`func (o *PatchedCustomerFeedback) GetFeedbackCategoryOk() (*FeedbackCategoryEnum, bool)`

GetFeedbackCategoryOk returns a tuple with the FeedbackCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackCategory

`func (o *PatchedCustomerFeedback) SetFeedbackCategory(v FeedbackCategoryEnum)`

SetFeedbackCategory sets FeedbackCategory field to given value.

### HasFeedbackCategory

`func (o *PatchedCustomerFeedback) HasFeedbackCategory() bool`

HasFeedbackCategory returns a boolean if a field has been set.

### SetFeedbackCategoryNil

`func (o *PatchedCustomerFeedback) SetFeedbackCategoryNil(b bool)`

 SetFeedbackCategoryNil sets the value for FeedbackCategory to be an explicit nil

### UnsetFeedbackCategory
`func (o *PatchedCustomerFeedback) UnsetFeedbackCategory()`

UnsetFeedbackCategory ensures that no value is present for FeedbackCategory, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


