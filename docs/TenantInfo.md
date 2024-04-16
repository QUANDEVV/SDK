# TenantInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **int32** | The tenant chats ID. | [readonly] 
**TenantName** | **string** | name of the company | 

## Methods

### NewTenantInfo

`func NewTenantInfo(tenantId int32, tenantName string, ) *TenantInfo`

NewTenantInfo instantiates a new TenantInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantInfoWithDefaults

`func NewTenantInfoWithDefaults() *TenantInfo`

NewTenantInfoWithDefaults instantiates a new TenantInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *TenantInfo) GetTenantId() int32`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TenantInfo) GetTenantIdOk() (*int32, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TenantInfo) SetTenantId(v int32)`

SetTenantId sets TenantId field to given value.


### GetTenantName

`func (o *TenantInfo) GetTenantName() string`

GetTenantName returns the TenantName field if non-nil, zero value otherwise.

### GetTenantNameOk

`func (o *TenantInfo) GetTenantNameOk() (*string, bool)`

GetTenantNameOk returns a tuple with the TenantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantName

`func (o *TenantInfo) SetTenantName(v string)`

SetTenantName sets TenantName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


