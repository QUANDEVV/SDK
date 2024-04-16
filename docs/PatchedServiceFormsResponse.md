# PatchedServiceFormsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **NullableString** |  | [optional] [readonly] 
**ServiceFormResponseId** | Pointer to **int32** |  | [optional] [readonly] 
**ServiceForm** | Pointer to [**ServiceForms**](ServiceForms.md) |  | [optional] 
**Client** | Pointer to [**ClientServicesInfo**](ClientServicesInfo.md) |  | [optional] 
**UnauthenticatedUser** | Pointer to [**AnonymousUser**](AnonymousUser.md) |  | [optional] 
**FormResponse** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewPatchedServiceFormsResponse

`func NewPatchedServiceFormsResponse() *PatchedServiceFormsResponse`

NewPatchedServiceFormsResponse instantiates a new PatchedServiceFormsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedServiceFormsResponseWithDefaults

`func NewPatchedServiceFormsResponseWithDefaults() *PatchedServiceFormsResponse`

NewPatchedServiceFormsResponseWithDefaults instantiates a new PatchedServiceFormsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *PatchedServiceFormsResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedServiceFormsResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedServiceFormsResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedServiceFormsResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *PatchedServiceFormsResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *PatchedServiceFormsResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetServiceFormResponseId

`func (o *PatchedServiceFormsResponse) GetServiceFormResponseId() int32`

GetServiceFormResponseId returns the ServiceFormResponseId field if non-nil, zero value otherwise.

### GetServiceFormResponseIdOk

`func (o *PatchedServiceFormsResponse) GetServiceFormResponseIdOk() (*int32, bool)`

GetServiceFormResponseIdOk returns a tuple with the ServiceFormResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceFormResponseId

`func (o *PatchedServiceFormsResponse) SetServiceFormResponseId(v int32)`

SetServiceFormResponseId sets ServiceFormResponseId field to given value.

### HasServiceFormResponseId

`func (o *PatchedServiceFormsResponse) HasServiceFormResponseId() bool`

HasServiceFormResponseId returns a boolean if a field has been set.

### GetServiceForm

`func (o *PatchedServiceFormsResponse) GetServiceForm() ServiceForms`

GetServiceForm returns the ServiceForm field if non-nil, zero value otherwise.

### GetServiceFormOk

`func (o *PatchedServiceFormsResponse) GetServiceFormOk() (*ServiceForms, bool)`

GetServiceFormOk returns a tuple with the ServiceForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceForm

`func (o *PatchedServiceFormsResponse) SetServiceForm(v ServiceForms)`

SetServiceForm sets ServiceForm field to given value.

### HasServiceForm

`func (o *PatchedServiceFormsResponse) HasServiceForm() bool`

HasServiceForm returns a boolean if a field has been set.

### GetClient

`func (o *PatchedServiceFormsResponse) GetClient() ClientServicesInfo`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *PatchedServiceFormsResponse) GetClientOk() (*ClientServicesInfo, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *PatchedServiceFormsResponse) SetClient(v ClientServicesInfo)`

SetClient sets Client field to given value.

### HasClient

`func (o *PatchedServiceFormsResponse) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetUnauthenticatedUser

`func (o *PatchedServiceFormsResponse) GetUnauthenticatedUser() AnonymousUser`

GetUnauthenticatedUser returns the UnauthenticatedUser field if non-nil, zero value otherwise.

### GetUnauthenticatedUserOk

`func (o *PatchedServiceFormsResponse) GetUnauthenticatedUserOk() (*AnonymousUser, bool)`

GetUnauthenticatedUserOk returns a tuple with the UnauthenticatedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnauthenticatedUser

`func (o *PatchedServiceFormsResponse) SetUnauthenticatedUser(v AnonymousUser)`

SetUnauthenticatedUser sets UnauthenticatedUser field to given value.

### HasUnauthenticatedUser

`func (o *PatchedServiceFormsResponse) HasUnauthenticatedUser() bool`

HasUnauthenticatedUser returns a boolean if a field has been set.

### GetFormResponse

`func (o *PatchedServiceFormsResponse) GetFormResponse() interface{}`

GetFormResponse returns the FormResponse field if non-nil, zero value otherwise.

### GetFormResponseOk

`func (o *PatchedServiceFormsResponse) GetFormResponseOk() (*interface{}, bool)`

GetFormResponseOk returns a tuple with the FormResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormResponse

`func (o *PatchedServiceFormsResponse) SetFormResponse(v interface{})`

SetFormResponse sets FormResponse field to given value.

### HasFormResponse

`func (o *PatchedServiceFormsResponse) HasFormResponse() bool`

HasFormResponse returns a boolean if a field has been set.

### SetFormResponseNil

`func (o *PatchedServiceFormsResponse) SetFormResponseNil(b bool)`

 SetFormResponseNil sets the value for FormResponse to be an explicit nil

### UnsetFormResponse
`func (o *PatchedServiceFormsResponse) UnsetFormResponse()`

UnsetFormResponse ensures that no value is present for FormResponse, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


