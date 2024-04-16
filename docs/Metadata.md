# Metadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetadataId** | **int32** | The tenant chats ID UUID for all chats. | [readonly] 
**TenantId** | **interface{}** | Display name of the tenant | 

## Methods

### NewMetadata

`func NewMetadata(metadataId int32, tenantId interface{}, ) *Metadata`

NewMetadata instantiates a new Metadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataWithDefaults

`func NewMetadataWithDefaults() *Metadata`

NewMetadataWithDefaults instantiates a new Metadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadataId

`func (o *Metadata) GetMetadataId() int32`

GetMetadataId returns the MetadataId field if non-nil, zero value otherwise.

### GetMetadataIdOk

`func (o *Metadata) GetMetadataIdOk() (*int32, bool)`

GetMetadataIdOk returns a tuple with the MetadataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataId

`func (o *Metadata) SetMetadataId(v int32)`

SetMetadataId sets MetadataId field to given value.


### GetTenantId

`func (o *Metadata) GetTenantId() interface{}`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Metadata) GetTenantIdOk() (*interface{}, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Metadata) SetTenantId(v interface{})`

SetTenantId sets TenantId field to given value.


### SetTenantIdNil

`func (o *Metadata) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *Metadata) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


