# PatchedDepartment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepartmentId** | Pointer to **int32** | The department ID. | [optional] [readonly] 
**Tenant** | Pointer to **int32** | The tenant this department belongs to | [optional] 
**Name** | Pointer to **string** | Name of the department | [optional] 
**Description** | Pointer to **NullableString** | Description of the department | [optional] 
**Employees** | Pointer to **[]int32** |  | [optional] [readonly] 
**Admins** | Pointer to **[]int32** |  | [optional] [readonly] 

## Methods

### NewPatchedDepartment

`func NewPatchedDepartment() *PatchedDepartment`

NewPatchedDepartment instantiates a new PatchedDepartment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedDepartmentWithDefaults

`func NewPatchedDepartmentWithDefaults() *PatchedDepartment`

NewPatchedDepartmentWithDefaults instantiates a new PatchedDepartment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepartmentId

`func (o *PatchedDepartment) GetDepartmentId() int32`

GetDepartmentId returns the DepartmentId field if non-nil, zero value otherwise.

### GetDepartmentIdOk

`func (o *PatchedDepartment) GetDepartmentIdOk() (*int32, bool)`

GetDepartmentIdOk returns a tuple with the DepartmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartmentId

`func (o *PatchedDepartment) SetDepartmentId(v int32)`

SetDepartmentId sets DepartmentId field to given value.

### HasDepartmentId

`func (o *PatchedDepartment) HasDepartmentId() bool`

HasDepartmentId returns a boolean if a field has been set.

### GetTenant

`func (o *PatchedDepartment) GetTenant() int32`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedDepartment) GetTenantOk() (*int32, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedDepartment) SetTenant(v int32)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedDepartment) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetName

`func (o *PatchedDepartment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedDepartment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedDepartment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedDepartment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedDepartment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedDepartment) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedDepartment) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedDepartment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PatchedDepartment) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PatchedDepartment) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEmployees

`func (o *PatchedDepartment) GetEmployees() []int32`

GetEmployees returns the Employees field if non-nil, zero value otherwise.

### GetEmployeesOk

`func (o *PatchedDepartment) GetEmployeesOk() (*[]int32, bool)`

GetEmployeesOk returns a tuple with the Employees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployees

`func (o *PatchedDepartment) SetEmployees(v []int32)`

SetEmployees sets Employees field to given value.

### HasEmployees

`func (o *PatchedDepartment) HasEmployees() bool`

HasEmployees returns a boolean if a field has been set.

### GetAdmins

`func (o *PatchedDepartment) GetAdmins() []int32`

GetAdmins returns the Admins field if non-nil, zero value otherwise.

### GetAdminsOk

`func (o *PatchedDepartment) GetAdminsOk() (*[]int32, bool)`

GetAdminsOk returns a tuple with the Admins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmins

`func (o *PatchedDepartment) SetAdmins(v []int32)`

SetAdmins sets Admins field to given value.

### HasAdmins

`func (o *PatchedDepartment) HasAdmins() bool`

HasAdmins returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


