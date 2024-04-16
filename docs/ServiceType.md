# ServiceType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**TenantId** | **int32** | The tenant chats ID. | 
**Type** | [**ServiceTypeTypeEnum**](ServiceTypeTypeEnum.md) |  | 
**Description** | **string** |  | 

## Methods

### NewServiceType

`func NewServiceType(id int32, tenantId int32, type_ ServiceTypeTypeEnum, description string, ) *ServiceType`

NewServiceType instantiates a new ServiceType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceTypeWithDefaults

`func NewServiceTypeWithDefaults() *ServiceType`

NewServiceTypeWithDefaults instantiates a new ServiceType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceType) SetId(v int32)`

SetId sets Id field to given value.


### GetTenantId

`func (o *ServiceType) GetTenantId() int32`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ServiceType) GetTenantIdOk() (*int32, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ServiceType) SetTenantId(v int32)`

SetTenantId sets TenantId field to given value.


### GetType

`func (o *ServiceType) GetType() ServiceTypeTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceType) GetTypeOk() (*ServiceTypeTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceType) SetType(v ServiceTypeTypeEnum)`

SetType sets Type field to given value.


### GetDescription

`func (o *ServiceType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceType) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


