# PatchedDelegation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DelegationId** | Pointer to **int32** |  | [optional] [readonly] 
**TenantId** | Pointer to [**TenantInfo**](TenantInfo.md) |  | [optional] 
**Employee** | Pointer to [**[]Employee**](Employee.md) |  | [optional] 
**Admin** | Pointer to [**[]Admin**](Admin.md) |  | [optional] 
**AssignedIssue** | Pointer to **NullableInt32** | The issue ID UUID . | [optional] 
**AssignedChat** | Pointer to **NullableInt32** | The chat ID UUID for an instance of a chat. | [optional] 
**AssignedService** | Pointer to **NullableInt32** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** | The timestamp of the chat. | [optional] [readonly] 
**DelegatedBy** | Pointer to **NullableInt32** |  | [optional] 
**DepartmentDelegatedTo** | Pointer to **NullableInt32** | The department ID. | [optional] 

## Methods

### NewPatchedDelegation

`func NewPatchedDelegation() *PatchedDelegation`

NewPatchedDelegation instantiates a new PatchedDelegation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedDelegationWithDefaults

`func NewPatchedDelegationWithDefaults() *PatchedDelegation`

NewPatchedDelegationWithDefaults instantiates a new PatchedDelegation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelegationId

`func (o *PatchedDelegation) GetDelegationId() int32`

GetDelegationId returns the DelegationId field if non-nil, zero value otherwise.

### GetDelegationIdOk

`func (o *PatchedDelegation) GetDelegationIdOk() (*int32, bool)`

GetDelegationIdOk returns a tuple with the DelegationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegationId

`func (o *PatchedDelegation) SetDelegationId(v int32)`

SetDelegationId sets DelegationId field to given value.

### HasDelegationId

`func (o *PatchedDelegation) HasDelegationId() bool`

HasDelegationId returns a boolean if a field has been set.

### GetTenantId

`func (o *PatchedDelegation) GetTenantId() TenantInfo`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *PatchedDelegation) GetTenantIdOk() (*TenantInfo, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *PatchedDelegation) SetTenantId(v TenantInfo)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *PatchedDelegation) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetEmployee

`func (o *PatchedDelegation) GetEmployee() []Employee`

GetEmployee returns the Employee field if non-nil, zero value otherwise.

### GetEmployeeOk

`func (o *PatchedDelegation) GetEmployeeOk() (*[]Employee, bool)`

GetEmployeeOk returns a tuple with the Employee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployee

`func (o *PatchedDelegation) SetEmployee(v []Employee)`

SetEmployee sets Employee field to given value.

### HasEmployee

`func (o *PatchedDelegation) HasEmployee() bool`

HasEmployee returns a boolean if a field has been set.

### GetAdmin

`func (o *PatchedDelegation) GetAdmin() []Admin`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *PatchedDelegation) GetAdminOk() (*[]Admin, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *PatchedDelegation) SetAdmin(v []Admin)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *PatchedDelegation) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetAssignedIssue

`func (o *PatchedDelegation) GetAssignedIssue() int32`

GetAssignedIssue returns the AssignedIssue field if non-nil, zero value otherwise.

### GetAssignedIssueOk

`func (o *PatchedDelegation) GetAssignedIssueOk() (*int32, bool)`

GetAssignedIssueOk returns a tuple with the AssignedIssue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedIssue

`func (o *PatchedDelegation) SetAssignedIssue(v int32)`

SetAssignedIssue sets AssignedIssue field to given value.

### HasAssignedIssue

`func (o *PatchedDelegation) HasAssignedIssue() bool`

HasAssignedIssue returns a boolean if a field has been set.

### SetAssignedIssueNil

`func (o *PatchedDelegation) SetAssignedIssueNil(b bool)`

 SetAssignedIssueNil sets the value for AssignedIssue to be an explicit nil

### UnsetAssignedIssue
`func (o *PatchedDelegation) UnsetAssignedIssue()`

UnsetAssignedIssue ensures that no value is present for AssignedIssue, not even an explicit nil
### GetAssignedChat

`func (o *PatchedDelegation) GetAssignedChat() int32`

GetAssignedChat returns the AssignedChat field if non-nil, zero value otherwise.

### GetAssignedChatOk

`func (o *PatchedDelegation) GetAssignedChatOk() (*int32, bool)`

GetAssignedChatOk returns a tuple with the AssignedChat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedChat

`func (o *PatchedDelegation) SetAssignedChat(v int32)`

SetAssignedChat sets AssignedChat field to given value.

### HasAssignedChat

`func (o *PatchedDelegation) HasAssignedChat() bool`

HasAssignedChat returns a boolean if a field has been set.

### SetAssignedChatNil

`func (o *PatchedDelegation) SetAssignedChatNil(b bool)`

 SetAssignedChatNil sets the value for AssignedChat to be an explicit nil

### UnsetAssignedChat
`func (o *PatchedDelegation) UnsetAssignedChat()`

UnsetAssignedChat ensures that no value is present for AssignedChat, not even an explicit nil
### GetAssignedService

`func (o *PatchedDelegation) GetAssignedService() int32`

GetAssignedService returns the AssignedService field if non-nil, zero value otherwise.

### GetAssignedServiceOk

`func (o *PatchedDelegation) GetAssignedServiceOk() (*int32, bool)`

GetAssignedServiceOk returns a tuple with the AssignedService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedService

`func (o *PatchedDelegation) SetAssignedService(v int32)`

SetAssignedService sets AssignedService field to given value.

### HasAssignedService

`func (o *PatchedDelegation) HasAssignedService() bool`

HasAssignedService returns a boolean if a field has been set.

### SetAssignedServiceNil

`func (o *PatchedDelegation) SetAssignedServiceNil(b bool)`

 SetAssignedServiceNil sets the value for AssignedService to be an explicit nil

### UnsetAssignedService
`func (o *PatchedDelegation) UnsetAssignedService()`

UnsetAssignedService ensures that no value is present for AssignedService, not even an explicit nil
### GetTimestamp

`func (o *PatchedDelegation) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PatchedDelegation) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PatchedDelegation) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *PatchedDelegation) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *PatchedDelegation) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *PatchedDelegation) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetDelegatedBy

`func (o *PatchedDelegation) GetDelegatedBy() int32`

GetDelegatedBy returns the DelegatedBy field if non-nil, zero value otherwise.

### GetDelegatedByOk

`func (o *PatchedDelegation) GetDelegatedByOk() (*int32, bool)`

GetDelegatedByOk returns a tuple with the DelegatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedBy

`func (o *PatchedDelegation) SetDelegatedBy(v int32)`

SetDelegatedBy sets DelegatedBy field to given value.

### HasDelegatedBy

`func (o *PatchedDelegation) HasDelegatedBy() bool`

HasDelegatedBy returns a boolean if a field has been set.

### SetDelegatedByNil

`func (o *PatchedDelegation) SetDelegatedByNil(b bool)`

 SetDelegatedByNil sets the value for DelegatedBy to be an explicit nil

### UnsetDelegatedBy
`func (o *PatchedDelegation) UnsetDelegatedBy()`

UnsetDelegatedBy ensures that no value is present for DelegatedBy, not even an explicit nil
### GetDepartmentDelegatedTo

`func (o *PatchedDelegation) GetDepartmentDelegatedTo() int32`

GetDepartmentDelegatedTo returns the DepartmentDelegatedTo field if non-nil, zero value otherwise.

### GetDepartmentDelegatedToOk

`func (o *PatchedDelegation) GetDepartmentDelegatedToOk() (*int32, bool)`

GetDepartmentDelegatedToOk returns a tuple with the DepartmentDelegatedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartmentDelegatedTo

`func (o *PatchedDelegation) SetDepartmentDelegatedTo(v int32)`

SetDepartmentDelegatedTo sets DepartmentDelegatedTo field to given value.

### HasDepartmentDelegatedTo

`func (o *PatchedDelegation) HasDepartmentDelegatedTo() bool`

HasDepartmentDelegatedTo returns a boolean if a field has been set.

### SetDepartmentDelegatedToNil

`func (o *PatchedDelegation) SetDepartmentDelegatedToNil(b bool)`

 SetDepartmentDelegatedToNil sets the value for DepartmentDelegatedTo to be an explicit nil

### UnsetDepartmentDelegatedTo
`func (o *PatchedDelegation) UnsetDepartmentDelegatedTo()`

UnsetDepartmentDelegatedTo ensures that no value is present for DepartmentDelegatedTo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


