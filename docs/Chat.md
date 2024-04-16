# Chat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | **int32** | The chat ID UUID for an instance of a chat. | [readonly] 
**TenantId** | [**TenantInfo**](TenantInfo.md) |  | 
**GuestClient** | [**ClientInfo**](ClientInfo.md) |  | 
**ChatOwner** | [**ClientInfo**](ClientInfo.md) |  | 
**AnonymousUser** | [**AnonymousUser**](AnonymousUser.md) |  | 
**ClientSatisfaction** | Pointer to **NullableBool** | Whether client is satisfied or not | [optional] 
**IvaEnabled** | Pointer to **bool** | Enable or disable whether the virtual assistant should talk to a client | [optional] 
**Timestamp** | **NullableTime** | The timestamp of the chat. | [readonly] 
**CreatedAt** | **NullableString** |  | [readonly] 

## Methods

### NewChat

`func NewChat(chatId int32, tenantId TenantInfo, guestClient ClientInfo, chatOwner ClientInfo, anonymousUser AnonymousUser, timestamp NullableTime, createdAt NullableString, ) *Chat`

NewChat instantiates a new Chat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatWithDefaults

`func NewChatWithDefaults() *Chat`

NewChatWithDefaults instantiates a new Chat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *Chat) GetChatId() int32`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *Chat) GetChatIdOk() (*int32, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *Chat) SetChatId(v int32)`

SetChatId sets ChatId field to given value.


### GetTenantId

`func (o *Chat) GetTenantId() TenantInfo`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Chat) GetTenantIdOk() (*TenantInfo, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Chat) SetTenantId(v TenantInfo)`

SetTenantId sets TenantId field to given value.


### GetGuestClient

`func (o *Chat) GetGuestClient() ClientInfo`

GetGuestClient returns the GuestClient field if non-nil, zero value otherwise.

### GetGuestClientOk

`func (o *Chat) GetGuestClientOk() (*ClientInfo, bool)`

GetGuestClientOk returns a tuple with the GuestClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestClient

`func (o *Chat) SetGuestClient(v ClientInfo)`

SetGuestClient sets GuestClient field to given value.


### GetChatOwner

`func (o *Chat) GetChatOwner() ClientInfo`

GetChatOwner returns the ChatOwner field if non-nil, zero value otherwise.

### GetChatOwnerOk

`func (o *Chat) GetChatOwnerOk() (*ClientInfo, bool)`

GetChatOwnerOk returns a tuple with the ChatOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatOwner

`func (o *Chat) SetChatOwner(v ClientInfo)`

SetChatOwner sets ChatOwner field to given value.


### GetAnonymousUser

`func (o *Chat) GetAnonymousUser() AnonymousUser`

GetAnonymousUser returns the AnonymousUser field if non-nil, zero value otherwise.

### GetAnonymousUserOk

`func (o *Chat) GetAnonymousUserOk() (*AnonymousUser, bool)`

GetAnonymousUserOk returns a tuple with the AnonymousUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymousUser

`func (o *Chat) SetAnonymousUser(v AnonymousUser)`

SetAnonymousUser sets AnonymousUser field to given value.


### GetClientSatisfaction

`func (o *Chat) GetClientSatisfaction() bool`

GetClientSatisfaction returns the ClientSatisfaction field if non-nil, zero value otherwise.

### GetClientSatisfactionOk

`func (o *Chat) GetClientSatisfactionOk() (*bool, bool)`

GetClientSatisfactionOk returns a tuple with the ClientSatisfaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSatisfaction

`func (o *Chat) SetClientSatisfaction(v bool)`

SetClientSatisfaction sets ClientSatisfaction field to given value.

### HasClientSatisfaction

`func (o *Chat) HasClientSatisfaction() bool`

HasClientSatisfaction returns a boolean if a field has been set.

### SetClientSatisfactionNil

`func (o *Chat) SetClientSatisfactionNil(b bool)`

 SetClientSatisfactionNil sets the value for ClientSatisfaction to be an explicit nil

### UnsetClientSatisfaction
`func (o *Chat) UnsetClientSatisfaction()`

UnsetClientSatisfaction ensures that no value is present for ClientSatisfaction, not even an explicit nil
### GetIvaEnabled

`func (o *Chat) GetIvaEnabled() bool`

GetIvaEnabled returns the IvaEnabled field if non-nil, zero value otherwise.

### GetIvaEnabledOk

`func (o *Chat) GetIvaEnabledOk() (*bool, bool)`

GetIvaEnabledOk returns a tuple with the IvaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIvaEnabled

`func (o *Chat) SetIvaEnabled(v bool)`

SetIvaEnabled sets IvaEnabled field to given value.

### HasIvaEnabled

`func (o *Chat) HasIvaEnabled() bool`

HasIvaEnabled returns a boolean if a field has been set.

### GetTimestamp

`func (o *Chat) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Chat) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Chat) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### SetTimestampNil

`func (o *Chat) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *Chat) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetCreatedAt

`func (o *Chat) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Chat) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Chat) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *Chat) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Chat) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


