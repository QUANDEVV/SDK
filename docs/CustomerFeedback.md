# CustomerFeedback

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerFeedbackId** | **int32** |  | [readonly] 
**Call** | [**Call**](Call.md) |  | [readonly] 
**Rating** | **int32** | Rating on a scale of 1 to 5. | 
**FeedbackText** | Pointer to **NullableString** | Detailed text feedback from the customer. | [optional] 
**FeedbackCategory** | [**FeedbackCategoryEnum**](FeedbackCategoryEnum.md) | Category of the feedback.  * &#x60;SERVICE&#x60; - Service * &#x60;TECHNICAL&#x60; - Technical Issue * &#x60;GENERAL&#x60; - General Feedback | 

## Methods

### NewCustomerFeedback

`func NewCustomerFeedback(customerFeedbackId int32, call Call, rating int32, feedbackCategory FeedbackCategoryEnum, ) *CustomerFeedback`

NewCustomerFeedback instantiates a new CustomerFeedback object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerFeedbackWithDefaults

`func NewCustomerFeedbackWithDefaults() *CustomerFeedback`

NewCustomerFeedbackWithDefaults instantiates a new CustomerFeedback object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerFeedbackId

`func (o *CustomerFeedback) GetCustomerFeedbackId() int32`

GetCustomerFeedbackId returns the CustomerFeedbackId field if non-nil, zero value otherwise.

### GetCustomerFeedbackIdOk

`func (o *CustomerFeedback) GetCustomerFeedbackIdOk() (*int32, bool)`

GetCustomerFeedbackIdOk returns a tuple with the CustomerFeedbackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerFeedbackId

`func (o *CustomerFeedback) SetCustomerFeedbackId(v int32)`

SetCustomerFeedbackId sets CustomerFeedbackId field to given value.


### GetCall

`func (o *CustomerFeedback) GetCall() Call`

GetCall returns the Call field if non-nil, zero value otherwise.

### GetCallOk

`func (o *CustomerFeedback) GetCallOk() (*Call, bool)`

GetCallOk returns a tuple with the Call field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCall

`func (o *CustomerFeedback) SetCall(v Call)`

SetCall sets Call field to given value.


### SetCallNil

`func (o *CustomerFeedback) SetCallNil(b bool)`

 SetCallNil sets the value for Call to be an explicit nil

### UnsetCall
`func (o *CustomerFeedback) UnsetCall()`

UnsetCall ensures that no value is present for Call, not even an explicit nil
### GetRating

`func (o *CustomerFeedback) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *CustomerFeedback) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *CustomerFeedback) SetRating(v int32)`

SetRating sets Rating field to given value.


### GetFeedbackText

`func (o *CustomerFeedback) GetFeedbackText() string`

GetFeedbackText returns the FeedbackText field if non-nil, zero value otherwise.

### GetFeedbackTextOk

`func (o *CustomerFeedback) GetFeedbackTextOk() (*string, bool)`

GetFeedbackTextOk returns a tuple with the FeedbackText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackText

`func (o *CustomerFeedback) SetFeedbackText(v string)`

SetFeedbackText sets FeedbackText field to given value.

### HasFeedbackText

`func (o *CustomerFeedback) HasFeedbackText() bool`

HasFeedbackText returns a boolean if a field has been set.

### SetFeedbackTextNil

`func (o *CustomerFeedback) SetFeedbackTextNil(b bool)`

 SetFeedbackTextNil sets the value for FeedbackText to be an explicit nil

### UnsetFeedbackText
`func (o *CustomerFeedback) UnsetFeedbackText()`

UnsetFeedbackText ensures that no value is present for FeedbackText, not even an explicit nil
### GetFeedbackCategory

`func (o *CustomerFeedback) GetFeedbackCategory() FeedbackCategoryEnum`

GetFeedbackCategory returns the FeedbackCategory field if non-nil, zero value otherwise.

### GetFeedbackCategoryOk

`func (o *CustomerFeedback) GetFeedbackCategoryOk() (*FeedbackCategoryEnum, bool)`

GetFeedbackCategoryOk returns a tuple with the FeedbackCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackCategory

`func (o *CustomerFeedback) SetFeedbackCategory(v FeedbackCategoryEnum)`

SetFeedbackCategory sets FeedbackCategory field to given value.


### SetFeedbackCategoryNil

`func (o *CustomerFeedback) SetFeedbackCategoryNil(b bool)`

 SetFeedbackCategoryNil sets the value for FeedbackCategory to be an explicit nil

### UnsetFeedbackCategory
`func (o *CustomerFeedback) UnsetFeedbackCategory()`

UnsetFeedbackCategory ensures that no value is present for FeedbackCategory, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


