# ServiceFormsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **NullableString** |  | [readonly] 
**ServiceFormResponseId** | **int32** |  | [readonly] 
**ServiceForm** | [**ServiceForms**](ServiceForms.md) |  | 
**Client** | [**ClientServicesInfo**](ClientServicesInfo.md) |  | 
**UnauthenticatedUser** | [**AnonymousUser**](AnonymousUser.md) |  | 
**FormResponse** | **interface{}** |  | 

## Methods

### NewServiceFormsResponse

`func NewServiceFormsResponse(createdAt NullableString, serviceFormResponseId int32, serviceForm ServiceForms, client ClientServicesInfo, unauthenticatedUser AnonymousUser, formResponse interface{}, ) *ServiceFormsResponse`

NewServiceFormsResponse instantiates a new ServiceFormsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceFormsResponseWithDefaults

`func NewServiceFormsResponseWithDefaults() *ServiceFormsResponse`

NewServiceFormsResponseWithDefaults instantiates a new ServiceFormsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ServiceFormsResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServiceFormsResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServiceFormsResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *ServiceFormsResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *ServiceFormsResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetServiceFormResponseId

`func (o *ServiceFormsResponse) GetServiceFormResponseId() int32`

GetServiceFormResponseId returns the ServiceFormResponseId field if non-nil, zero value otherwise.

### GetServiceFormResponseIdOk

`func (o *ServiceFormsResponse) GetServiceFormResponseIdOk() (*int32, bool)`

GetServiceFormResponseIdOk returns a tuple with the ServiceFormResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceFormResponseId

`func (o *ServiceFormsResponse) SetServiceFormResponseId(v int32)`

SetServiceFormResponseId sets ServiceFormResponseId field to given value.


### GetServiceForm

`func (o *ServiceFormsResponse) GetServiceForm() ServiceForms`

GetServiceForm returns the ServiceForm field if non-nil, zero value otherwise.

### GetServiceFormOk

`func (o *ServiceFormsResponse) GetServiceFormOk() (*ServiceForms, bool)`

GetServiceFormOk returns a tuple with the ServiceForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceForm

`func (o *ServiceFormsResponse) SetServiceForm(v ServiceForms)`

SetServiceForm sets ServiceForm field to given value.


### GetClient

`func (o *ServiceFormsResponse) GetClient() ClientServicesInfo`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *ServiceFormsResponse) GetClientOk() (*ClientServicesInfo, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *ServiceFormsResponse) SetClient(v ClientServicesInfo)`

SetClient sets Client field to given value.


### GetUnauthenticatedUser

`func (o *ServiceFormsResponse) GetUnauthenticatedUser() AnonymousUser`

GetUnauthenticatedUser returns the UnauthenticatedUser field if non-nil, zero value otherwise.

### GetUnauthenticatedUserOk

`func (o *ServiceFormsResponse) GetUnauthenticatedUserOk() (*AnonymousUser, bool)`

GetUnauthenticatedUserOk returns a tuple with the UnauthenticatedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnauthenticatedUser

`func (o *ServiceFormsResponse) SetUnauthenticatedUser(v AnonymousUser)`

SetUnauthenticatedUser sets UnauthenticatedUser field to given value.


### GetFormResponse

`func (o *ServiceFormsResponse) GetFormResponse() interface{}`

GetFormResponse returns the FormResponse field if non-nil, zero value otherwise.

### GetFormResponseOk

`func (o *ServiceFormsResponse) GetFormResponseOk() (*interface{}, bool)`

GetFormResponseOk returns a tuple with the FormResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormResponse

`func (o *ServiceFormsResponse) SetFormResponse(v interface{})`

SetFormResponse sets FormResponse field to given value.


### SetFormResponseNil

`func (o *ServiceFormsResponse) SetFormResponseNil(b bool)`

 SetFormResponseNil sets the value for FormResponse to be an explicit nil

### UnsetFormResponse
`func (o *ServiceFormsResponse) UnsetFormResponse()`

UnsetFormResponse ensures that no value is present for FormResponse, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


