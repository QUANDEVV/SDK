# Department

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepartmentId** | **int32** | The department ID. | [readonly] 
**Tenant** | **int32** | The tenant this department belongs to | 
**Name** | **string** | Name of the department | 
**Description** | Pointer to **NullableString** | Description of the department | [optional] 
**Employees** | **[]int32** |  | [readonly] 
**Admins** | **[]int32** |  | [readonly] 

## Methods

### NewDepartment

`func NewDepartment(departmentId int32, tenant int32, name string, employees []int32, admins []int32, ) *Department`

NewDepartment instantiates a new Department object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepartmentWithDefaults

`func NewDepartmentWithDefaults() *Department`

NewDepartmentWithDefaults instantiates a new Department object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepartmentId

`func (o *Department) GetDepartmentId() int32`

GetDepartmentId returns the DepartmentId field if non-nil, zero value otherwise.

### GetDepartmentIdOk

`func (o *Department) GetDepartmentIdOk() (*int32, bool)`

GetDepartmentIdOk returns a tuple with the DepartmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartmentId

`func (o *Department) SetDepartmentId(v int32)`

SetDepartmentId sets DepartmentId field to given value.


### GetTenant

`func (o *Department) GetTenant() int32`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *Department) GetTenantOk() (*int32, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *Department) SetTenant(v int32)`

SetTenant sets Tenant field to given value.


### GetName

`func (o *Department) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Department) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Department) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Department) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Department) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Department) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Department) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Department) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Department) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEmployees

`func (o *Department) GetEmployees() []int32`

GetEmployees returns the Employees field if non-nil, zero value otherwise.

### GetEmployeesOk

`func (o *Department) GetEmployeesOk() (*[]int32, bool)`

GetEmployeesOk returns a tuple with the Employees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployees

`func (o *Department) SetEmployees(v []int32)`

SetEmployees sets Employees field to given value.


### GetAdmins

`func (o *Department) GetAdmins() []int32`

GetAdmins returns the Admins field if non-nil, zero value otherwise.

### GetAdminsOk

`func (o *Department) GetAdminsOk() (*[]int32, bool)`

GetAdminsOk returns a tuple with the Admins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmins

`func (o *Department) SetAdmins(v []int32)`

SetAdmins sets Admins field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


