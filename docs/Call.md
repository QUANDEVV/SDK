# Call

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallId** | **int32** | The call id. | [readonly] 
**TenantId** | [**TenantInfo**](TenantInfo.md) |  | [readonly] 
**Duration** | **string** |  | 
**Success** | **bool** |  | 
**ClientCaller** | [**ClientInfo**](ClientInfo.md) |  | 
**GuestCalled** | Pointer to **NullableInt32** | Display name of the client | [optional] 

## Methods

### NewCall

`func NewCall(callId int32, tenantId TenantInfo, duration string, success bool, clientCaller ClientInfo, ) *Call`

NewCall instantiates a new Call object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallWithDefaults

`func NewCallWithDefaults() *Call`

NewCallWithDefaults instantiates a new Call object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallId

`func (o *Call) GetCallId() int32`

GetCallId returns the CallId field if non-nil, zero value otherwise.

### GetCallIdOk

`func (o *Call) GetCallIdOk() (*int32, bool)`

GetCallIdOk returns a tuple with the CallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallId

`func (o *Call) SetCallId(v int32)`

SetCallId sets CallId field to given value.


### GetTenantId

`func (o *Call) GetTenantId() TenantInfo`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Call) GetTenantIdOk() (*TenantInfo, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Call) SetTenantId(v TenantInfo)`

SetTenantId sets TenantId field to given value.


### SetTenantIdNil

`func (o *Call) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *Call) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetDuration

`func (o *Call) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *Call) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *Call) SetDuration(v string)`

SetDuration sets Duration field to given value.


### GetSuccess

`func (o *Call) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *Call) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *Call) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetClientCaller

`func (o *Call) GetClientCaller() ClientInfo`

GetClientCaller returns the ClientCaller field if non-nil, zero value otherwise.

### GetClientCallerOk

`func (o *Call) GetClientCallerOk() (*ClientInfo, bool)`

GetClientCallerOk returns a tuple with the ClientCaller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCaller

`func (o *Call) SetClientCaller(v ClientInfo)`

SetClientCaller sets ClientCaller field to given value.


### GetGuestCalled

`func (o *Call) GetGuestCalled() int32`

GetGuestCalled returns the GuestCalled field if non-nil, zero value otherwise.

### GetGuestCalledOk

`func (o *Call) GetGuestCalledOk() (*int32, bool)`

GetGuestCalledOk returns a tuple with the GuestCalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestCalled

`func (o *Call) SetGuestCalled(v int32)`

SetGuestCalled sets GuestCalled field to given value.

### HasGuestCalled

`func (o *Call) HasGuestCalled() bool`

HasGuestCalled returns a boolean if a field has been set.

### SetGuestCalledNil

`func (o *Call) SetGuestCalledNil(b bool)`

 SetGuestCalledNil sets the value for GuestCalled to be an explicit nil

### UnsetGuestCalled
`func (o *Call) UnsetGuestCalled()`

UnsetGuestCalled ensures that no value is present for GuestCalled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


