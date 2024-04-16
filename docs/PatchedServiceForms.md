# PatchedServiceForms

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceFormId** | Pointer to **int32** |  | [optional] [readonly] 
**Service** | Pointer to [**Service**](Service.md) |  | [optional] 
**NameOfForm** | Pointer to **string** |  | [optional] 
**Form** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewPatchedServiceForms

`func NewPatchedServiceForms() *PatchedServiceForms`

NewPatchedServiceForms instantiates a new PatchedServiceForms object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedServiceFormsWithDefaults

`func NewPatchedServiceFormsWithDefaults() *PatchedServiceForms`

NewPatchedServiceFormsWithDefaults instantiates a new PatchedServiceForms object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceFormId

`func (o *PatchedServiceForms) GetServiceFormId() int32`

GetServiceFormId returns the ServiceFormId field if non-nil, zero value otherwise.

### GetServiceFormIdOk

`func (o *PatchedServiceForms) GetServiceFormIdOk() (*int32, bool)`

GetServiceFormIdOk returns a tuple with the ServiceFormId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceFormId

`func (o *PatchedServiceForms) SetServiceFormId(v int32)`

SetServiceFormId sets ServiceFormId field to given value.

### HasServiceFormId

`func (o *PatchedServiceForms) HasServiceFormId() bool`

HasServiceFormId returns a boolean if a field has been set.

### GetService

`func (o *PatchedServiceForms) GetService() Service`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *PatchedServiceForms) GetServiceOk() (*Service, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *PatchedServiceForms) SetService(v Service)`

SetService sets Service field to given value.

### HasService

`func (o *PatchedServiceForms) HasService() bool`

HasService returns a boolean if a field has been set.

### GetNameOfForm

`func (o *PatchedServiceForms) GetNameOfForm() string`

GetNameOfForm returns the NameOfForm field if non-nil, zero value otherwise.

### GetNameOfFormOk

`func (o *PatchedServiceForms) GetNameOfFormOk() (*string, bool)`

GetNameOfFormOk returns a tuple with the NameOfForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameOfForm

`func (o *PatchedServiceForms) SetNameOfForm(v string)`

SetNameOfForm sets NameOfForm field to given value.

### HasNameOfForm

`func (o *PatchedServiceForms) HasNameOfForm() bool`

HasNameOfForm returns a boolean if a field has been set.

### GetForm

`func (o *PatchedServiceForms) GetForm() interface{}`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *PatchedServiceForms) GetFormOk() (*interface{}, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *PatchedServiceForms) SetForm(v interface{})`

SetForm sets Form field to given value.

### HasForm

`func (o *PatchedServiceForms) HasForm() bool`

HasForm returns a boolean if a field has been set.

### SetFormNil

`func (o *PatchedServiceForms) SetFormNil(b bool)`

 SetFormNil sets the value for Form to be an explicit nil

### UnsetForm
`func (o *PatchedServiceForms) UnsetForm()`

UnsetForm ensures that no value is present for Form, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


