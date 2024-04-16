# PatchedCall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallId** | Pointer to **int32** | The call id. | [optional] [readonly] 
**TenantId** | Pointer to [**TenantInfo**](TenantInfo.md) |  | [optional] [readonly] 
**Duration** | Pointer to **string** |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**ClientCaller** | Pointer to [**ClientInfo**](ClientInfo.md) |  | [optional] 
**GuestCalled** | Pointer to **NullableInt32** | Display name of the client | [optional] 

## Methods

### NewPatchedCall

`func NewPatchedCall() *PatchedCall`

NewPatchedCall instantiates a new PatchedCall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedCallWithDefaults

`func NewPatchedCallWithDefaults() *PatchedCall`

NewPatchedCallWithDefaults instantiates a new PatchedCall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallId

`func (o *PatchedCall) GetCallId() int32`

GetCallId returns the CallId field if non-nil, zero value otherwise.

### GetCallIdOk

`func (o *PatchedCall) GetCallIdOk() (*int32, bool)`

GetCallIdOk returns a tuple with the CallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallId

`func (o *PatchedCall) SetCallId(v int32)`

SetCallId sets CallId field to given value.

### HasCallId

`func (o *PatchedCall) HasCallId() bool`

HasCallId returns a boolean if a field has been set.

### GetTenantId

`func (o *PatchedCall) GetTenantId() TenantInfo`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *PatchedCall) GetTenantIdOk() (*TenantInfo, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *PatchedCall) SetTenantId(v TenantInfo)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *PatchedCall) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *PatchedCall) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *PatchedCall) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetDuration

`func (o *PatchedCall) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *PatchedCall) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *PatchedCall) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *PatchedCall) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetSuccess

`func (o *PatchedCall) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *PatchedCall) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *PatchedCall) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *PatchedCall) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetClientCaller

`func (o *PatchedCall) GetClientCaller() ClientInfo`

GetClientCaller returns the ClientCaller field if non-nil, zero value otherwise.

### GetClientCallerOk

`func (o *PatchedCall) GetClientCallerOk() (*ClientInfo, bool)`

GetClientCallerOk returns a tuple with the ClientCaller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCaller

`func (o *PatchedCall) SetClientCaller(v ClientInfo)`

SetClientCaller sets ClientCaller field to given value.

### HasClientCaller

`func (o *PatchedCall) HasClientCaller() bool`

HasClientCaller returns a boolean if a field has been set.

### GetGuestCalled

`func (o *PatchedCall) GetGuestCalled() int32`

GetGuestCalled returns the GuestCalled field if non-nil, zero value otherwise.

### GetGuestCalledOk

`func (o *PatchedCall) GetGuestCalledOk() (*int32, bool)`

GetGuestCalledOk returns a tuple with the GuestCalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestCalled

`func (o *PatchedCall) SetGuestCalled(v int32)`

SetGuestCalled sets GuestCalled field to given value.

### HasGuestCalled

`func (o *PatchedCall) HasGuestCalled() bool`

HasGuestCalled returns a boolean if a field has been set.

### SetGuestCalledNil

`func (o *PatchedCall) SetGuestCalledNil(b bool)`

 SetGuestCalledNil sets the value for GuestCalled to be an explicit nil

### UnsetGuestCalled
`func (o *PatchedCall) UnsetGuestCalled()`

UnsetGuestCalled ensures that no value is present for GuestCalled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


