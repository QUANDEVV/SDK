# PaymentsTenant

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

### NewPaymentsTenant

`func NewPaymentsTenant(tenantId int32, tenantName string, industry IndustryEnum, ) *PaymentsTenant`

NewPaymentsTenant instantiates a new PaymentsTenant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentsTenantWithDefaults

`func NewPaymentsTenantWithDefaults() *PaymentsTenant`

NewPaymentsTenantWithDefaults instantiates a new PaymentsTenant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *PaymentsTenant) GetTenantId() int32`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *PaymentsTenant) GetTenantIdOk() (*int32, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *PaymentsTenant) SetTenantId(v int32)`

SetTenantId sets TenantId field to given value.


### GetTenantName

`func (o *PaymentsTenant) GetTenantName() string`

GetTenantName returns the TenantName field if non-nil, zero value otherwise.

### GetTenantNameOk

`func (o *PaymentsTenant) GetTenantNameOk() (*string, bool)`

GetTenantNameOk returns a tuple with the TenantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantName

`func (o *PaymentsTenant) SetTenantName(v string)`

SetTenantName sets TenantName field to given value.


### GetIndustry

`func (o *PaymentsTenant) GetIndustry() IndustryEnum`

GetIndustry returns the Industry field if non-nil, zero value otherwise.

### GetIndustryOk

`func (o *PaymentsTenant) GetIndustryOk() (*IndustryEnum, bool)`

GetIndustryOk returns a tuple with the Industry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustry

`func (o *PaymentsTenant) SetIndustry(v IndustryEnum)`

SetIndustry sets Industry field to given value.


### SetIndustryNil

`func (o *PaymentsTenant) SetIndustryNil(b bool)`

 SetIndustryNil sets the value for Industry to be an explicit nil

### UnsetIndustry
`func (o *PaymentsTenant) UnsetIndustry()`

UnsetIndustry ensures that no value is present for Industry, not even an explicit nil
### GetToken

`func (o *PaymentsTenant) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PaymentsTenant) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PaymentsTenant) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PaymentsTenant) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *PaymentsTenant) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *PaymentsTenant) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetSubCategory

`func (o *PaymentsTenant) GetSubCategory() string`

GetSubCategory returns the SubCategory field if non-nil, zero value otherwise.

### GetSubCategoryOk

`func (o *PaymentsTenant) GetSubCategoryOk() (*string, bool)`

GetSubCategoryOk returns a tuple with the SubCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubCategory

`func (o *PaymentsTenant) SetSubCategory(v string)`

SetSubCategory sets SubCategory field to given value.

### HasSubCategory

`func (o *PaymentsTenant) HasSubCategory() bool`

HasSubCategory returns a boolean if a field has been set.

### SetSubCategoryNil

`func (o *PaymentsTenant) SetSubCategoryNil(b bool)`

 SetSubCategoryNil sets the value for SubCategory to be an explicit nil

### UnsetSubCategory
`func (o *PaymentsTenant) UnsetSubCategory()`

UnsetSubCategory ensures that no value is present for SubCategory, not even an explicit nil
### GetRegistrationNumber

`func (o *PaymentsTenant) GetRegistrationNumber() string`

GetRegistrationNumber returns the RegistrationNumber field if non-nil, zero value otherwise.

### GetRegistrationNumberOk

`func (o *PaymentsTenant) GetRegistrationNumberOk() (*string, bool)`

GetRegistrationNumberOk returns a tuple with the RegistrationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationNumber

`func (o *PaymentsTenant) SetRegistrationNumber(v string)`

SetRegistrationNumber sets RegistrationNumber field to given value.

### HasRegistrationNumber

`func (o *PaymentsTenant) HasRegistrationNumber() bool`

HasRegistrationNumber returns a boolean if a field has been set.

### SetRegistrationNumberNil

`func (o *PaymentsTenant) SetRegistrationNumberNil(b bool)`

 SetRegistrationNumberNil sets the value for RegistrationNumber to be an explicit nil

### UnsetRegistrationNumber
`func (o *PaymentsTenant) UnsetRegistrationNumber()`

UnsetRegistrationNumber ensures that no value is present for RegistrationNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


