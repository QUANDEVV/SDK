# PatchedNotifications

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationsId** | Pointer to **int32** | The notifications for an instance of a chat. | [optional] [readonly] 
**Employee** | Pointer to **NullableInt32** |  | [optional] 
**Admin** | Pointer to **NullableInt32** |  | [optional] 
**Client** | Pointer to **NullableInt32** |  | [optional] 
**GuestUser** | Pointer to **NullableInt32** | The webhookevent chats ID UUID. | [optional] 
**Delegation** | Pointer to [**Delegation**](Delegation.md) |  | [optional] 
**NotificationMessage** | Pointer to **NullableString** | Message content of the notification | [optional] 
**Read** | Pointer to **bool** | Whether the notification has been read or not | [optional] 
**Timestamp** | Pointer to **NullableTime** | The timestamp of the chat. | [optional] [readonly] 

## Methods

### NewPatchedNotifications

`func NewPatchedNotifications() *PatchedNotifications`

NewPatchedNotifications instantiates a new PatchedNotifications object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedNotificationsWithDefaults

`func NewPatchedNotificationsWithDefaults() *PatchedNotifications`

NewPatchedNotificationsWithDefaults instantiates a new PatchedNotifications object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationsId

`func (o *PatchedNotifications) GetNotificationsId() int32`

GetNotificationsId returns the NotificationsId field if non-nil, zero value otherwise.

### GetNotificationsIdOk

`func (o *PatchedNotifications) GetNotificationsIdOk() (*int32, bool)`

GetNotificationsIdOk returns a tuple with the NotificationsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationsId

`func (o *PatchedNotifications) SetNotificationsId(v int32)`

SetNotificationsId sets NotificationsId field to given value.

### HasNotificationsId

`func (o *PatchedNotifications) HasNotificationsId() bool`

HasNotificationsId returns a boolean if a field has been set.

### GetEmployee

`func (o *PatchedNotifications) GetEmployee() int32`

GetEmployee returns the Employee field if non-nil, zero value otherwise.

### GetEmployeeOk

`func (o *PatchedNotifications) GetEmployeeOk() (*int32, bool)`

GetEmployeeOk returns a tuple with the Employee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployee

`func (o *PatchedNotifications) SetEmployee(v int32)`

SetEmployee sets Employee field to given value.

### HasEmployee

`func (o *PatchedNotifications) HasEmployee() bool`

HasEmployee returns a boolean if a field has been set.

### SetEmployeeNil

`func (o *PatchedNotifications) SetEmployeeNil(b bool)`

 SetEmployeeNil sets the value for Employee to be an explicit nil

### UnsetEmployee
`func (o *PatchedNotifications) UnsetEmployee()`

UnsetEmployee ensures that no value is present for Employee, not even an explicit nil
### GetAdmin

`func (o *PatchedNotifications) GetAdmin() int32`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *PatchedNotifications) GetAdminOk() (*int32, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *PatchedNotifications) SetAdmin(v int32)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *PatchedNotifications) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### SetAdminNil

`func (o *PatchedNotifications) SetAdminNil(b bool)`

 SetAdminNil sets the value for Admin to be an explicit nil

### UnsetAdmin
`func (o *PatchedNotifications) UnsetAdmin()`

UnsetAdmin ensures that no value is present for Admin, not even an explicit nil
### GetClient

`func (o *PatchedNotifications) GetClient() int32`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *PatchedNotifications) GetClientOk() (*int32, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *PatchedNotifications) SetClient(v int32)`

SetClient sets Client field to given value.

### HasClient

`func (o *PatchedNotifications) HasClient() bool`

HasClient returns a boolean if a field has been set.

### SetClientNil

`func (o *PatchedNotifications) SetClientNil(b bool)`

 SetClientNil sets the value for Client to be an explicit nil

### UnsetClient
`func (o *PatchedNotifications) UnsetClient()`

UnsetClient ensures that no value is present for Client, not even an explicit nil
### GetGuestUser

`func (o *PatchedNotifications) GetGuestUser() int32`

GetGuestUser returns the GuestUser field if non-nil, zero value otherwise.

### GetGuestUserOk

`func (o *PatchedNotifications) GetGuestUserOk() (*int32, bool)`

GetGuestUserOk returns a tuple with the GuestUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestUser

`func (o *PatchedNotifications) SetGuestUser(v int32)`

SetGuestUser sets GuestUser field to given value.

### HasGuestUser

`func (o *PatchedNotifications) HasGuestUser() bool`

HasGuestUser returns a boolean if a field has been set.

### SetGuestUserNil

`func (o *PatchedNotifications) SetGuestUserNil(b bool)`

 SetGuestUserNil sets the value for GuestUser to be an explicit nil

### UnsetGuestUser
`func (o *PatchedNotifications) UnsetGuestUser()`

UnsetGuestUser ensures that no value is present for GuestUser, not even an explicit nil
### GetDelegation

`func (o *PatchedNotifications) GetDelegation() Delegation`

GetDelegation returns the Delegation field if non-nil, zero value otherwise.

### GetDelegationOk

`func (o *PatchedNotifications) GetDelegationOk() (*Delegation, bool)`

GetDelegationOk returns a tuple with the Delegation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegation

`func (o *PatchedNotifications) SetDelegation(v Delegation)`

SetDelegation sets Delegation field to given value.

### HasDelegation

`func (o *PatchedNotifications) HasDelegation() bool`

HasDelegation returns a boolean if a field has been set.

### GetNotificationMessage

`func (o *PatchedNotifications) GetNotificationMessage() string`

GetNotificationMessage returns the NotificationMessage field if non-nil, zero value otherwise.

### GetNotificationMessageOk

`func (o *PatchedNotifications) GetNotificationMessageOk() (*string, bool)`

GetNotificationMessageOk returns a tuple with the NotificationMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationMessage

`func (o *PatchedNotifications) SetNotificationMessage(v string)`

SetNotificationMessage sets NotificationMessage field to given value.

### HasNotificationMessage

`func (o *PatchedNotifications) HasNotificationMessage() bool`

HasNotificationMessage returns a boolean if a field has been set.

### SetNotificationMessageNil

`func (o *PatchedNotifications) SetNotificationMessageNil(b bool)`

 SetNotificationMessageNil sets the value for NotificationMessage to be an explicit nil

### UnsetNotificationMessage
`func (o *PatchedNotifications) UnsetNotificationMessage()`

UnsetNotificationMessage ensures that no value is present for NotificationMessage, not even an explicit nil
### GetRead

`func (o *PatchedNotifications) GetRead() bool`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *PatchedNotifications) GetReadOk() (*bool, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *PatchedNotifications) SetRead(v bool)`

SetRead sets Read field to given value.

### HasRead

`func (o *PatchedNotifications) HasRead() bool`

HasRead returns a boolean if a field has been set.

### GetTimestamp

`func (o *PatchedNotifications) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PatchedNotifications) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PatchedNotifications) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *PatchedNotifications) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *PatchedNotifications) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *PatchedNotifications) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


