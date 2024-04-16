# InteractionHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InteractionHistoryId** | **int32** | The call id. | [readonly] 
**Call** | [**Call**](Call.md) |  | [readonly] 
**InteractionType** | [**InteractionTypeEnum**](InteractionTypeEnum.md) | Type of interaction.  * &#x60;NOTE&#x60; - Note * &#x60;DECISION&#x60; - Decision * &#x60;ESCALATION&#x60; - Escalation | 
**Details** | **string** | Details of the interaction. | 

## Methods

### NewInteractionHistory

`func NewInteractionHistory(interactionHistoryId int32, call Call, interactionType InteractionTypeEnum, details string, ) *InteractionHistory`

NewInteractionHistory instantiates a new InteractionHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInteractionHistoryWithDefaults

`func NewInteractionHistoryWithDefaults() *InteractionHistory`

NewInteractionHistoryWithDefaults instantiates a new InteractionHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInteractionHistoryId

`func (o *InteractionHistory) GetInteractionHistoryId() int32`

GetInteractionHistoryId returns the InteractionHistoryId field if non-nil, zero value otherwise.

### GetInteractionHistoryIdOk

`func (o *InteractionHistory) GetInteractionHistoryIdOk() (*int32, bool)`

GetInteractionHistoryIdOk returns a tuple with the InteractionHistoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractionHistoryId

`func (o *InteractionHistory) SetInteractionHistoryId(v int32)`

SetInteractionHistoryId sets InteractionHistoryId field to given value.


### GetCall

`func (o *InteractionHistory) GetCall() Call`

GetCall returns the Call field if non-nil, zero value otherwise.

### GetCallOk

`func (o *InteractionHistory) GetCallOk() (*Call, bool)`

GetCallOk returns a tuple with the Call field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCall

`func (o *InteractionHistory) SetCall(v Call)`

SetCall sets Call field to given value.


### SetCallNil

`func (o *InteractionHistory) SetCallNil(b bool)`

 SetCallNil sets the value for Call to be an explicit nil

### UnsetCall
`func (o *InteractionHistory) UnsetCall()`

UnsetCall ensures that no value is present for Call, not even an explicit nil
### GetInteractionType

`func (o *InteractionHistory) GetInteractionType() InteractionTypeEnum`

GetInteractionType returns the InteractionType field if non-nil, zero value otherwise.

### GetInteractionTypeOk

`func (o *InteractionHistory) GetInteractionTypeOk() (*InteractionTypeEnum, bool)`

GetInteractionTypeOk returns a tuple with the InteractionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractionType

`func (o *InteractionHistory) SetInteractionType(v InteractionTypeEnum)`

SetInteractionType sets InteractionType field to given value.


### SetInteractionTypeNil

`func (o *InteractionHistory) SetInteractionTypeNil(b bool)`

 SetInteractionTypeNil sets the value for InteractionType to be an explicit nil

### UnsetInteractionType
`func (o *InteractionHistory) UnsetInteractionType()`

UnsetInteractionType ensures that no value is present for InteractionType, not even an explicit nil
### GetDetails

`func (o *InteractionHistory) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *InteractionHistory) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *InteractionHistory) SetDetails(v string)`

SetDetails sets Details field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


