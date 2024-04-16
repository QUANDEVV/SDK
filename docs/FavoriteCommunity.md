# FavoriteCommunity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommunityId** | **int32** | The tenant community ID UUID. | [readonly] 
**TenantId** | [**CommunityTenant**](CommunityTenant.md) |  | 
**Description** | Pointer to **NullableString** | Description of the community. | [optional] 

## Methods

### NewFavoriteCommunity

`func NewFavoriteCommunity(communityId int32, tenantId CommunityTenant, ) *FavoriteCommunity`

NewFavoriteCommunity instantiates a new FavoriteCommunity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFavoriteCommunityWithDefaults

`func NewFavoriteCommunityWithDefaults() *FavoriteCommunity`

NewFavoriteCommunityWithDefaults instantiates a new FavoriteCommunity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommunityId

`func (o *FavoriteCommunity) GetCommunityId() int32`

GetCommunityId returns the CommunityId field if non-nil, zero value otherwise.

### GetCommunityIdOk

`func (o *FavoriteCommunity) GetCommunityIdOk() (*int32, bool)`

GetCommunityIdOk returns a tuple with the CommunityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityId

`func (o *FavoriteCommunity) SetCommunityId(v int32)`

SetCommunityId sets CommunityId field to given value.


### GetTenantId

`func (o *FavoriteCommunity) GetTenantId() CommunityTenant`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *FavoriteCommunity) GetTenantIdOk() (*CommunityTenant, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *FavoriteCommunity) SetTenantId(v CommunityTenant)`

SetTenantId sets TenantId field to given value.


### GetDescription

`func (o *FavoriteCommunity) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FavoriteCommunity) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FavoriteCommunity) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FavoriteCommunity) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *FavoriteCommunity) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *FavoriteCommunity) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


