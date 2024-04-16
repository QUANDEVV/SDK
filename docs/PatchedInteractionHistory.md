# PatchedInteractionHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InteractionHistoryId** | Pointer to **int32** | The call id. | [optional] [readonly] 
**Call** | Pointer to [**Call**](Call.md) |  | [optional] [readonly] 
**InteractionType** | Pointer to [**InteractionTypeEnum**](InteractionTypeEnum.md) | Type of interaction.  * &#x60;NOTE&#x60; - Note * &#x60;DECISION&#x60; - Decision * &#x60;ESCALATION&#x60; - Escalation | [optional] 
**Details** | Pointer to **string** | Details of the interaction. | [optional] 

## Methods

### NewPatchedInteractionHistory

`func NewPatchedInteractionHistory() *PatchedInteractionHistory`

NewPatchedInteractionHistory instantiates a new PatchedInteractionHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedInteractionHistoryWithDefaults

`func NewPatchedInteractionHistoryWithDefaults() *PatchedInteractionHistory`

NewPatchedInteractionHistoryWithDefaults instantiates a new PatchedInteractionHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInteractionHistoryId

`func (o *PatchedInteractionHistory) GetInteractionHistoryId() int32`

GetInteractionHistoryId returns the InteractionHistoryId field if non-nil, zero value otherwise.

### GetInteractionHistoryIdOk

`func (o *PatchedInteractionHistory) GetInteractionHistoryIdOk() (*int32, bool)`

GetInteractionHistoryIdOk returns a tuple with the InteractionHistoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractionHistoryId

`func (o *PatchedInteractionHistory) SetInteractionHistoryId(v int32)`

SetInteractionHistoryId sets InteractionHistoryId field to given value.

### HasInteractionHistoryId

`func (o *PatchedInteractionHistory) HasInteractionHistoryId() bool`

HasInteractionHistoryId returns a boolean if a field has been set.

### GetCall

`func (o *PatchedInteractionHistory) GetCall() Call`

GetCall returns the Call field if non-nil, zero value otherwise.

### GetCallOk

`func (o *PatchedInteractionHistory) GetCallOk() (*Call, bool)`

GetCallOk returns a tuple with the Call field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCall

`func (o *PatchedInteractionHistory) SetCall(v Call)`

SetCall sets Call field to given value.

### HasCall

`func (o *PatchedInteractionHistory) HasCall() bool`

HasCall returns a boolean if a field has been set.

### SetCallNil

`func (o *PatchedInteractionHistory) SetCallNil(b bool)`

 SetCallNil sets the value for Call to be an explicit nil

### UnsetCall
`func (o *PatchedInteractionHistory) UnsetCall()`

UnsetCall ensures that no value is present for Call, not even an explicit nil
### GetInteractionType

`func (o *PatchedInteractionHistory) GetInteractionType() InteractionTypeEnum`

GetInteractionType returns the InteractionType field if non-nil, zero value otherwise.

### GetInteractionTypeOk

`func (o *PatchedInteractionHistory) GetInteractionTypeOk() (*InteractionTypeEnum, bool)`

GetInteractionTypeOk returns a tuple with the InteractionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractionType

`func (o *PatchedInteractionHistory) SetInteractionType(v InteractionTypeEnum)`

SetInteractionType sets InteractionType field to given value.

### HasInteractionType

`func (o *PatchedInteractionHistory) HasInteractionType() bool`

HasInteractionType returns a boolean if a field has been set.

### SetInteractionTypeNil

`func (o *PatchedInteractionHistory) SetInteractionTypeNil(b bool)`

 SetInteractionTypeNil sets the value for InteractionType to be an explicit nil

### UnsetInteractionType
`func (o *PatchedInteractionHistory) UnsetInteractionType()`

UnsetInteractionType ensures that no value is present for InteractionType, not even an explicit nil
### GetDetails

`func (o *PatchedInteractionHistory) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *PatchedInteractionHistory) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *PatchedInteractionHistory) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *PatchedInteractionHistory) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


