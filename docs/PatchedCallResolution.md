# PatchedCallResolution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallResolutionId** | Pointer to **int32** | The call resolution id. | [optional] [readonly] 
**Call** | Pointer to [**Call**](Call.md) |  | [optional] [readonly] 
**Status** | Pointer to [**CallResolutionStatusEnum**](CallResolutionStatusEnum.md) |  | [optional] 
**PendingActions** | Pointer to **interface{}** |  | [optional] 
**FollowUpRequired** | Pointer to **bool** |  | [optional] 
**FollowUpDate** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewPatchedCallResolution

`func NewPatchedCallResolution() *PatchedCallResolution`

NewPatchedCallResolution instantiates a new PatchedCallResolution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedCallResolutionWithDefaults

`func NewPatchedCallResolutionWithDefaults() *PatchedCallResolution`

NewPatchedCallResolutionWithDefaults instantiates a new PatchedCallResolution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallResolutionId

`func (o *PatchedCallResolution) GetCallResolutionId() int32`

GetCallResolutionId returns the CallResolutionId field if non-nil, zero value otherwise.

### GetCallResolutionIdOk

`func (o *PatchedCallResolution) GetCallResolutionIdOk() (*int32, bool)`

GetCallResolutionIdOk returns a tuple with the CallResolutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallResolutionId

`func (o *PatchedCallResolution) SetCallResolutionId(v int32)`

SetCallResolutionId sets CallResolutionId field to given value.

### HasCallResolutionId

`func (o *PatchedCallResolution) HasCallResolutionId() bool`

HasCallResolutionId returns a boolean if a field has been set.

### GetCall

`func (o *PatchedCallResolution) GetCall() Call`

GetCall returns the Call field if non-nil, zero value otherwise.

### GetCallOk

`func (o *PatchedCallResolution) GetCallOk() (*Call, bool)`

GetCallOk returns a tuple with the Call field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCall

`func (o *PatchedCallResolution) SetCall(v Call)`

SetCall sets Call field to given value.

### HasCall

`func (o *PatchedCallResolution) HasCall() bool`

HasCall returns a boolean if a field has been set.

### SetCallNil

`func (o *PatchedCallResolution) SetCallNil(b bool)`

 SetCallNil sets the value for Call to be an explicit nil

### UnsetCall
`func (o *PatchedCallResolution) UnsetCall()`

UnsetCall ensures that no value is present for Call, not even an explicit nil
### GetStatus

`func (o *PatchedCallResolution) GetStatus() CallResolutionStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedCallResolution) GetStatusOk() (*CallResolutionStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedCallResolution) SetStatus(v CallResolutionStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedCallResolution) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPendingActions

`func (o *PatchedCallResolution) GetPendingActions() interface{}`

GetPendingActions returns the PendingActions field if non-nil, zero value otherwise.

### GetPendingActionsOk

`func (o *PatchedCallResolution) GetPendingActionsOk() (*interface{}, bool)`

GetPendingActionsOk returns a tuple with the PendingActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingActions

`func (o *PatchedCallResolution) SetPendingActions(v interface{})`

SetPendingActions sets PendingActions field to given value.

### HasPendingActions

`func (o *PatchedCallResolution) HasPendingActions() bool`

HasPendingActions returns a boolean if a field has been set.

### SetPendingActionsNil

`func (o *PatchedCallResolution) SetPendingActionsNil(b bool)`

 SetPendingActionsNil sets the value for PendingActions to be an explicit nil

### UnsetPendingActions
`func (o *PatchedCallResolution) UnsetPendingActions()`

UnsetPendingActions ensures that no value is present for PendingActions, not even an explicit nil
### GetFollowUpRequired

`func (o *PatchedCallResolution) GetFollowUpRequired() bool`

GetFollowUpRequired returns the FollowUpRequired field if non-nil, zero value otherwise.

### GetFollowUpRequiredOk

`func (o *PatchedCallResolution) GetFollowUpRequiredOk() (*bool, bool)`

GetFollowUpRequiredOk returns a tuple with the FollowUpRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowUpRequired

`func (o *PatchedCallResolution) SetFollowUpRequired(v bool)`

SetFollowUpRequired sets FollowUpRequired field to given value.

### HasFollowUpRequired

`func (o *PatchedCallResolution) HasFollowUpRequired() bool`

HasFollowUpRequired returns a boolean if a field has been set.

### GetFollowUpDate

`func (o *PatchedCallResolution) GetFollowUpDate() time.Time`

GetFollowUpDate returns the FollowUpDate field if non-nil, zero value otherwise.

### GetFollowUpDateOk

`func (o *PatchedCallResolution) GetFollowUpDateOk() (*time.Time, bool)`

GetFollowUpDateOk returns a tuple with the FollowUpDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowUpDate

`func (o *PatchedCallResolution) SetFollowUpDate(v time.Time)`

SetFollowUpDate sets FollowUpDate field to given value.

### HasFollowUpDate

`func (o *PatchedCallResolution) HasFollowUpDate() bool`

HasFollowUpDate returns a boolean if a field has been set.

### SetFollowUpDateNil

`func (o *PatchedCallResolution) SetFollowUpDateNil(b bool)`

 SetFollowUpDateNil sets the value for FollowUpDate to be an explicit nil

### UnsetFollowUpDate
`func (o *PatchedCallResolution) UnsetFollowUpDate()`

UnsetFollowUpDate ensures that no value is present for FollowUpDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


