# CallResolution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallResolutionId** | **int32** | The call resolution id. | [readonly] 
**Call** | [**Call**](Call.md) |  | [readonly] 
**Status** | Pointer to [**CallResolutionStatusEnum**](CallResolutionStatusEnum.md) |  | [optional] 
**PendingActions** | Pointer to **interface{}** |  | [optional] 
**FollowUpRequired** | Pointer to **bool** |  | [optional] 
**FollowUpDate** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewCallResolution

`func NewCallResolution(callResolutionId int32, call Call, ) *CallResolution`

NewCallResolution instantiates a new CallResolution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallResolutionWithDefaults

`func NewCallResolutionWithDefaults() *CallResolution`

NewCallResolutionWithDefaults instantiates a new CallResolution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallResolutionId

`func (o *CallResolution) GetCallResolutionId() int32`

GetCallResolutionId returns the CallResolutionId field if non-nil, zero value otherwise.

### GetCallResolutionIdOk

`func (o *CallResolution) GetCallResolutionIdOk() (*int32, bool)`

GetCallResolutionIdOk returns a tuple with the CallResolutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallResolutionId

`func (o *CallResolution) SetCallResolutionId(v int32)`

SetCallResolutionId sets CallResolutionId field to given value.


### GetCall

`func (o *CallResolution) GetCall() Call`

GetCall returns the Call field if non-nil, zero value otherwise.

### GetCallOk

`func (o *CallResolution) GetCallOk() (*Call, bool)`

GetCallOk returns a tuple with the Call field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCall

`func (o *CallResolution) SetCall(v Call)`

SetCall sets Call field to given value.


### SetCallNil

`func (o *CallResolution) SetCallNil(b bool)`

 SetCallNil sets the value for Call to be an explicit nil

### UnsetCall
`func (o *CallResolution) UnsetCall()`

UnsetCall ensures that no value is present for Call, not even an explicit nil
### GetStatus

`func (o *CallResolution) GetStatus() CallResolutionStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CallResolution) GetStatusOk() (*CallResolutionStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CallResolution) SetStatus(v CallResolutionStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CallResolution) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPendingActions

`func (o *CallResolution) GetPendingActions() interface{}`

GetPendingActions returns the PendingActions field if non-nil, zero value otherwise.

### GetPendingActionsOk

`func (o *CallResolution) GetPendingActionsOk() (*interface{}, bool)`

GetPendingActionsOk returns a tuple with the PendingActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingActions

`func (o *CallResolution) SetPendingActions(v interface{})`

SetPendingActions sets PendingActions field to given value.

### HasPendingActions

`func (o *CallResolution) HasPendingActions() bool`

HasPendingActions returns a boolean if a field has been set.

### SetPendingActionsNil

`func (o *CallResolution) SetPendingActionsNil(b bool)`

 SetPendingActionsNil sets the value for PendingActions to be an explicit nil

### UnsetPendingActions
`func (o *CallResolution) UnsetPendingActions()`

UnsetPendingActions ensures that no value is present for PendingActions, not even an explicit nil
### GetFollowUpRequired

`func (o *CallResolution) GetFollowUpRequired() bool`

GetFollowUpRequired returns the FollowUpRequired field if non-nil, zero value otherwise.

### GetFollowUpRequiredOk

`func (o *CallResolution) GetFollowUpRequiredOk() (*bool, bool)`

GetFollowUpRequiredOk returns a tuple with the FollowUpRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowUpRequired

`func (o *CallResolution) SetFollowUpRequired(v bool)`

SetFollowUpRequired sets FollowUpRequired field to given value.

### HasFollowUpRequired

`func (o *CallResolution) HasFollowUpRequired() bool`

HasFollowUpRequired returns a boolean if a field has been set.

### GetFollowUpDate

`func (o *CallResolution) GetFollowUpDate() time.Time`

GetFollowUpDate returns the FollowUpDate field if non-nil, zero value otherwise.

### GetFollowUpDateOk

`func (o *CallResolution) GetFollowUpDateOk() (*time.Time, bool)`

GetFollowUpDateOk returns a tuple with the FollowUpDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowUpDate

`func (o *CallResolution) SetFollowUpDate(v time.Time)`

SetFollowUpDate sets FollowUpDate field to given value.

### HasFollowUpDate

`func (o *CallResolution) HasFollowUpDate() bool`

HasFollowUpDate returns a boolean if a field has been set.

### SetFollowUpDateNil

`func (o *CallResolution) SetFollowUpDateNil(b bool)`

 SetFollowUpDateNil sets the value for FollowUpDate to be an explicit nil

### UnsetFollowUpDate
`func (o *CallResolution) UnsetFollowUpDate()`

UnsetFollowUpDate ensures that no value is present for FollowUpDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


