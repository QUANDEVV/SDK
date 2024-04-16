# PatchedChat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | Pointer to **int32** | The chat ID UUID for an instance of a chat. | [optional] [readonly] 
**TenantId** | Pointer to [**TenantInfo**](TenantInfo.md) |  | [optional] 
**GuestClient** | Pointer to [**ClientInfo**](ClientInfo.md) |  | [optional] 
**ChatOwner** | Pointer to [**ClientInfo**](ClientInfo.md) |  | [optional] 
**AnonymousUser** | Pointer to [**AnonymousUser**](AnonymousUser.md) |  | [optional] 
**ClientSatisfaction** | Pointer to **NullableBool** | Whether client is satisfied or not | [optional] 
**IvaEnabled** | Pointer to **bool** | Enable or disable whether the virtual assistant should talk to a client | [optional] 
**Timestamp** | Pointer to **NullableTime** | The timestamp of the chat. | [optional] [readonly] 
**CreatedAt** | Pointer to **NullableString** |  | [optional] [readonly] 

## Methods

### NewPatchedChat

`func NewPatchedChat() *PatchedChat`

NewPatchedChat instantiates a new PatchedChat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedChatWithDefaults

`func NewPatchedChatWithDefaults() *PatchedChat`

NewPatchedChatWithDefaults instantiates a new PatchedChat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *PatchedChat) GetChatId() int32`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *PatchedChat) GetChatIdOk() (*int32, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *PatchedChat) SetChatId(v int32)`

SetChatId sets ChatId field to given value.

### HasChatId

`func (o *PatchedChat) HasChatId() bool`

HasChatId returns a boolean if a field has been set.

### GetTenantId

`func (o *PatchedChat) GetTenantId() TenantInfo`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *PatchedChat) GetTenantIdOk() (*TenantInfo, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *PatchedChat) SetTenantId(v TenantInfo)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *PatchedChat) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetGuestClient

`func (o *PatchedChat) GetGuestClient() ClientInfo`

GetGuestClient returns the GuestClient field if non-nil, zero value otherwise.

### GetGuestClientOk

`func (o *PatchedChat) GetGuestClientOk() (*ClientInfo, bool)`

GetGuestClientOk returns a tuple with the GuestClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestClient

`func (o *PatchedChat) SetGuestClient(v ClientInfo)`

SetGuestClient sets GuestClient field to given value.

### HasGuestClient

`func (o *PatchedChat) HasGuestClient() bool`

HasGuestClient returns a boolean if a field has been set.

### GetChatOwner

`func (o *PatchedChat) GetChatOwner() ClientInfo`

GetChatOwner returns the ChatOwner field if non-nil, zero value otherwise.

### GetChatOwnerOk

`func (o *PatchedChat) GetChatOwnerOk() (*ClientInfo, bool)`

GetChatOwnerOk returns a tuple with the ChatOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatOwner

`func (o *PatchedChat) SetChatOwner(v ClientInfo)`

SetChatOwner sets ChatOwner field to given value.

### HasChatOwner

`func (o *PatchedChat) HasChatOwner() bool`

HasChatOwner returns a boolean if a field has been set.

### GetAnonymousUser

`func (o *PatchedChat) GetAnonymousUser() AnonymousUser`

GetAnonymousUser returns the AnonymousUser field if non-nil, zero value otherwise.

### GetAnonymousUserOk

`func (o *PatchedChat) GetAnonymousUserOk() (*AnonymousUser, bool)`

GetAnonymousUserOk returns a tuple with the AnonymousUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymousUser

`func (o *PatchedChat) SetAnonymousUser(v AnonymousUser)`

SetAnonymousUser sets AnonymousUser field to given value.

### HasAnonymousUser

`func (o *PatchedChat) HasAnonymousUser() bool`

HasAnonymousUser returns a boolean if a field has been set.

### GetClientSatisfaction

`func (o *PatchedChat) GetClientSatisfaction() bool`

GetClientSatisfaction returns the ClientSatisfaction field if non-nil, zero value otherwise.

### GetClientSatisfactionOk

`func (o *PatchedChat) GetClientSatisfactionOk() (*bool, bool)`

GetClientSatisfactionOk returns a tuple with the ClientSatisfaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSatisfaction

`func (o *PatchedChat) SetClientSatisfaction(v bool)`

SetClientSatisfaction sets ClientSatisfaction field to given value.

### HasClientSatisfaction

`func (o *PatchedChat) HasClientSatisfaction() bool`

HasClientSatisfaction returns a boolean if a field has been set.

### SetClientSatisfactionNil

`func (o *PatchedChat) SetClientSatisfactionNil(b bool)`

 SetClientSatisfactionNil sets the value for ClientSatisfaction to be an explicit nil

### UnsetClientSatisfaction
`func (o *PatchedChat) UnsetClientSatisfaction()`

UnsetClientSatisfaction ensures that no value is present for ClientSatisfaction, not even an explicit nil
### GetIvaEnabled

`func (o *PatchedChat) GetIvaEnabled() bool`

GetIvaEnabled returns the IvaEnabled field if non-nil, zero value otherwise.

### GetIvaEnabledOk

`func (o *PatchedChat) GetIvaEnabledOk() (*bool, bool)`

GetIvaEnabledOk returns a tuple with the IvaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIvaEnabled

`func (o *PatchedChat) SetIvaEnabled(v bool)`

SetIvaEnabled sets IvaEnabled field to given value.

### HasIvaEnabled

`func (o *PatchedChat) HasIvaEnabled() bool`

HasIvaEnabled returns a boolean if a field has been set.

### GetTimestamp

`func (o *PatchedChat) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PatchedChat) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PatchedChat) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *PatchedChat) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *PatchedChat) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *PatchedChat) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetCreatedAt

`func (o *PatchedChat) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedChat) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedChat) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedChat) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *PatchedChat) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *PatchedChat) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


