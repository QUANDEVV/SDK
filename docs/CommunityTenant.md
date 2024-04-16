# CommunityTenant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **int32** | The tenant chats ID. | [readonly] 
**TenantName** | **string** | name of the company | 
**Industry** | [**IndustryEnum**](IndustryEnum.md) | Industry of the company  * &#x60;IT&#x60; - Information Technology * &#x60;Finance&#x60; - Finance * &#x60;Healthcare&#x60; - Healthcare * &#x60;Education&#x60; - Education * &#x60;Retail&#x60; - Retail * &#x60;Manufacturing&#x60; - Manufacturing * &#x60;Automotive&#x60; - Automotive * &#x60;Hospitality&#x60; - Hospitality * &#x60;RealEstate&#x60; - Real Estate * &#x60;Media&#x60; - Media * &#x60;Telecommunications&#x60; - Telecommunications * &#x60;Energy&#x60; - Energy * &#x60;Transportation&#x60; - Transportation * &#x60;Agriculture&#x60; - Agriculture | 
**Token** | Pointer to **NullableString** | Token for the tenant | [optional] 
**SubCategory** | Pointer to **NullableString** |  | [optional] 
**RegistrationNumber** | Pointer to **NullableString** | The business registration number | [optional] 

## Methods

### NewCommunityTenant

`func NewCommunityTenant(tenantId int32, tenantName string, industry IndustryEnum, ) *CommunityTenant`

NewCommunityTenant instantiates a new CommunityTenant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommunityTenantWithDefaults

`func NewCommunityTenantWithDefaults() *CommunityTenant`

NewCommunityTenantWithDefaults instantiates a new CommunityTenant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *CommunityTenant) GetTenantId() int32`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CommunityTenant) GetTenantIdOk() (*int32, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CommunityTenant) SetTenantId(v int32)`

SetTenantId sets TenantId field to given value.


### GetTenantName

`func (o *CommunityTenant) GetTenantName() string`

GetTenantName returns the TenantName field if non-nil, zero value otherwise.

### GetTenantNameOk

`func (o *CommunityTenant) GetTenantNameOk() (*string, bool)`

GetTenantNameOk returns a tuple with the TenantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantName

`func (o *CommunityTenant) SetTenantName(v string)`

SetTenantName sets TenantName field to given value.


### GetIndustry

`func (o *CommunityTenant) GetIndustry() IndustryEnum`

GetIndustry returns the Industry field if non-nil, zero value otherwise.

### GetIndustryOk

`func (o *CommunityTenant) GetIndustryOk() (*IndustryEnum, bool)`

GetIndustryOk returns a tuple with the Industry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustry

`func (o *CommunityTenant) SetIndustry(v IndustryEnum)`

SetIndustry sets Industry field to given value.


### SetIndustryNil

`func (o *CommunityTenant) SetIndustryNil(b bool)`

 SetIndustryNil sets the value for Industry to be an explicit nil

### UnsetIndustry
`func (o *CommunityTenant) UnsetIndustry()`

UnsetIndustry ensures that no value is present for Industry, not even an explicit nil
### GetToken

`func (o *CommunityTenant) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CommunityTenant) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CommunityTenant) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CommunityTenant) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *CommunityTenant) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *CommunityTenant) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetSubCategory

`func (o *CommunityTenant) GetSubCategory() string`

GetSubCategory returns the SubCategory field if non-nil, zero value otherwise.

### GetSubCategoryOk

`func (o *CommunityTenant) GetSubCategoryOk() (*string, bool)`

GetSubCategoryOk returns a tuple with the SubCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubCategory

`func (o *CommunityTenant) SetSubCategory(v string)`

SetSubCategory sets SubCategory field to given value.

### HasSubCategory

`func (o *CommunityTenant) HasSubCategory() bool`

HasSubCategory returns a boolean if a field has been set.

### SetSubCategoryNil

`func (o *CommunityTenant) SetSubCategoryNil(b bool)`

 SetSubCategoryNil sets the value for SubCategory to be an explicit nil

### UnsetSubCategory
`func (o *CommunityTenant) UnsetSubCategory()`

UnsetSubCategory ensures that no value is present for SubCategory, not even an explicit nil
### GetRegistrationNumber

`func (o *CommunityTenant) GetRegistrationNumber() string`

GetRegistrationNumber returns the RegistrationNumber field if non-nil, zero value otherwise.

### GetRegistrationNumberOk

`func (o *CommunityTenant) GetRegistrationNumberOk() (*string, bool)`

GetRegistrationNumberOk returns a tuple with the RegistrationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationNumber

`func (o *CommunityTenant) SetRegistrationNumber(v string)`

SetRegistrationNumber sets RegistrationNumber field to given value.

### HasRegistrationNumber

`func (o *CommunityTenant) HasRegistrationNumber() bool`

HasRegistrationNumber returns a boolean if a field has been set.

### SetRegistrationNumberNil

`func (o *CommunityTenant) SetRegistrationNumberNil(b bool)`

 SetRegistrationNumberNil sets the value for RegistrationNumber to be an explicit nil

### UnsetRegistrationNumber
`func (o *CommunityTenant) UnsetRegistrationNumber()`

UnsetRegistrationNumber ensures that no value is present for RegistrationNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


