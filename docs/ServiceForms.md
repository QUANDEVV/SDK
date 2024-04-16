# ServiceForms

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceFormId** | **int32** |  | [readonly] 
**Service** | [**Service**](Service.md) |  | 
**NameOfForm** | **string** |  | 
**Form** | **interface{}** |  | 

## Methods

### NewServiceForms

`func NewServiceForms(serviceFormId int32, service Service, nameOfForm string, form interface{}, ) *ServiceForms`

NewServiceForms instantiates a new ServiceForms object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceFormsWithDefaults

`func NewServiceFormsWithDefaults() *ServiceForms`

NewServiceFormsWithDefaults instantiates a new ServiceForms object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceFormId

`func (o *ServiceForms) GetServiceFormId() int32`

GetServiceFormId returns the ServiceFormId field if non-nil, zero value otherwise.

### GetServiceFormIdOk

`func (o *ServiceForms) GetServiceFormIdOk() (*int32, bool)`

GetServiceFormIdOk returns a tuple with the ServiceFormId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceFormId

`func (o *ServiceForms) SetServiceFormId(v int32)`

SetServiceFormId sets ServiceFormId field to given value.


### GetService

`func (o *ServiceForms) GetService() Service`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ServiceForms) GetServiceOk() (*Service, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ServiceForms) SetService(v Service)`

SetService sets Service field to given value.


### GetNameOfForm

`func (o *ServiceForms) GetNameOfForm() string`

GetNameOfForm returns the NameOfForm field if non-nil, zero value otherwise.

### GetNameOfFormOk

`func (o *ServiceForms) GetNameOfFormOk() (*string, bool)`

GetNameOfFormOk returns a tuple with the NameOfForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameOfForm

`func (o *ServiceForms) SetNameOfForm(v string)`

SetNameOfForm sets NameOfForm field to given value.


### GetForm

`func (o *ServiceForms) GetForm() interface{}`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *ServiceForms) GetFormOk() (*interface{}, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *ServiceForms) SetForm(v interface{})`

SetForm sets Form field to given value.


### SetFormNil

`func (o *ServiceForms) SetFormNil(b bool)`

 SetFormNil sets the value for Form to be an explicit nil

### UnsetForm
`func (o *ServiceForms) UnsetForm()`

UnsetForm ensures that no value is present for Form, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


