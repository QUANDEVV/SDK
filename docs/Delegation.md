# Delegation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DelegationId** | **int32** |  | [readonly] 
**TenantId** | [**TenantInfo**](TenantInfo.md) |  | 
**Employee** | [**[]Employee**](Employee.md) |  | 
**Admin** | [**[]Admin**](Admin.md) |  | 
**AssignedIssue** | Pointer to **NullableInt32** | The issue ID UUID . | [optional] 
**AssignedChat** | Pointer to **NullableInt32** | The chat ID UUID for an instance of a chat. | [optional] 
**AssignedService** | Pointer to **NullableInt32** |  | [optional] 
**Timestamp** | **NullableTime** | The timestamp of the chat. | [readonly] 
**DelegatedBy** | Pointer to **NullableInt32** |  | [optional] 
**DepartmentDelegatedTo** | Pointer to **NullableInt32** | The department ID. | [optional] 

## Methods

### NewDelegation

`func NewDelegation(delegationId int32, tenantId TenantInfo, employee []Employee, admin []Admin, timestamp NullableTime, ) *Delegation`

NewDelegation instantiates a new Delegation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDelegationWithDefaults

`func NewDelegationWithDefaults() *Delegation`

NewDelegationWithDefaults instantiates a new Delegation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelegationId

`func (o *Delegation) GetDelegationId() int32`

GetDelegationId returns the DelegationId field if non-nil, zero value otherwise.

### GetDelegationIdOk

`func (o *Delegation) GetDelegationIdOk() (*int32, bool)`

GetDelegationIdOk returns a tuple with the DelegationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegationId

`func (o *Delegation) SetDelegationId(v int32)`

SetDelegationId sets DelegationId field to given value.


### GetTenantId

`func (o *Delegation) GetTenantId() TenantInfo`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Delegation) GetTenantIdOk() (*TenantInfo, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Delegation) SetTenantId(v TenantInfo)`

SetTenantId sets TenantId field to given value.


### GetEmployee

`func (o *Delegation) GetEmployee() []Employee`

GetEmployee returns the Employee field if non-nil, zero value otherwise.

### GetEmployeeOk

`func (o *Delegation) GetEmployeeOk() (*[]Employee, bool)`

GetEmployeeOk returns a tuple with the Employee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployee

`func (o *Delegation) SetEmployee(v []Employee)`

SetEmployee sets Employee field to given value.


### GetAdmin

`func (o *Delegation) GetAdmin() []Admin`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *Delegation) GetAdminOk() (*[]Admin, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *Delegation) SetAdmin(v []Admin)`

SetAdmin sets Admin field to given value.


### GetAssignedIssue

`func (o *Delegation) GetAssignedIssue() int32`

GetAssignedIssue returns the AssignedIssue field if non-nil, zero value otherwise.

### GetAssignedIssueOk

`func (o *Delegation) GetAssignedIssueOk() (*int32, bool)`

GetAssignedIssueOk returns a tuple with the AssignedIssue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedIssue

`func (o *Delegation) SetAssignedIssue(v int32)`

SetAssignedIssue sets AssignedIssue field to given value.

### HasAssignedIssue

`func (o *Delegation) HasAssignedIssue() bool`

HasAssignedIssue returns a boolean if a field has been set.

### SetAssignedIssueNil

`func (o *Delegation) SetAssignedIssueNil(b bool)`

 SetAssignedIssueNil sets the value for AssignedIssue to be an explicit nil

### UnsetAssignedIssue
`func (o *Delegation) UnsetAssignedIssue()`

UnsetAssignedIssue ensures that no value is present for AssignedIssue, not even an explicit nil
### GetAssignedChat

`func (o *Delegation) GetAssignedChat() int32`

GetAssignedChat returns the AssignedChat field if non-nil, zero value otherwise.

### GetAssignedChatOk

`func (o *Delegation) GetAssignedChatOk() (*int32, bool)`

GetAssignedChatOk returns a tuple with the AssignedChat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedChat

`func (o *Delegation) SetAssignedChat(v int32)`

SetAssignedChat sets AssignedChat field to given value.

### HasAssignedChat

`func (o *Delegation) HasAssignedChat() bool`

HasAssignedChat returns a boolean if a field has been set.

### SetAssignedChatNil

`func (o *Delegation) SetAssignedChatNil(b bool)`

 SetAssignedChatNil sets the value for AssignedChat to be an explicit nil

### UnsetAssignedChat
`func (o *Delegation) UnsetAssignedChat()`

UnsetAssignedChat ensures that no value is present for AssignedChat, not even an explicit nil
### GetAssignedService

`func (o *Delegation) GetAssignedService() int32`

GetAssignedService returns the AssignedService field if non-nil, zero value otherwise.

### GetAssignedServiceOk

`func (o *Delegation) GetAssignedServiceOk() (*int32, bool)`

GetAssignedServiceOk returns a tuple with the AssignedService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedService

`func (o *Delegation) SetAssignedService(v int32)`

SetAssignedService sets AssignedService field to given value.

### HasAssignedService

`func (o *Delegation) HasAssignedService() bool`

HasAssignedService returns a boolean if a field has been set.

### SetAssignedServiceNil

`func (o *Delegation) SetAssignedServiceNil(b bool)`

 SetAssignedServiceNil sets the value for AssignedService to be an explicit nil

### UnsetAssignedService
`func (o *Delegation) UnsetAssignedService()`

UnsetAssignedService ensures that no value is present for AssignedService, not even an explicit nil
### GetTimestamp

`func (o *Delegation) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Delegation) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Delegation) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### SetTimestampNil

`func (o *Delegation) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *Delegation) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetDelegatedBy

`func (o *Delegation) GetDelegatedBy() int32`

GetDelegatedBy returns the DelegatedBy field if non-nil, zero value otherwise.

### GetDelegatedByOk

`func (o *Delegation) GetDelegatedByOk() (*int32, bool)`

GetDelegatedByOk returns a tuple with the DelegatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedBy

`func (o *Delegation) SetDelegatedBy(v int32)`

SetDelegatedBy sets DelegatedBy field to given value.

### HasDelegatedBy

`func (o *Delegation) HasDelegatedBy() bool`

HasDelegatedBy returns a boolean if a field has been set.

### SetDelegatedByNil

`func (o *Delegation) SetDelegatedByNil(b bool)`

 SetDelegatedByNil sets the value for DelegatedBy to be an explicit nil

### UnsetDelegatedBy
`func (o *Delegation) UnsetDelegatedBy()`

UnsetDelegatedBy ensures that no value is present for DelegatedBy, not even an explicit nil
### GetDepartmentDelegatedTo

`func (o *Delegation) GetDepartmentDelegatedTo() int32`

GetDepartmentDelegatedTo returns the DepartmentDelegatedTo field if non-nil, zero value otherwise.

### GetDepartmentDelegatedToOk

`func (o *Delegation) GetDepartmentDelegatedToOk() (*int32, bool)`

GetDepartmentDelegatedToOk returns a tuple with the DepartmentDelegatedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartmentDelegatedTo

`func (o *Delegation) SetDepartmentDelegatedTo(v int32)`

SetDepartmentDelegatedTo sets DepartmentDelegatedTo field to given value.

### HasDepartmentDelegatedTo

`func (o *Delegation) HasDepartmentDelegatedTo() bool`

HasDepartmentDelegatedTo returns a boolean if a field has been set.

### SetDepartmentDelegatedToNil

`func (o *Delegation) SetDepartmentDelegatedToNil(b bool)`

 SetDepartmentDelegatedToNil sets the value for DepartmentDelegatedTo to be an explicit nil

### UnsetDepartmentDelegatedTo
`func (o *Delegation) UnsetDepartmentDelegatedTo()`

UnsetDepartmentDelegatedTo ensures that no value is present for DepartmentDelegatedTo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


