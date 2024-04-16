# Tier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Features** | [**[]Feature**](Feature.md) |  | [readonly] 
**Prices** | [**[]Price**](Price.md) |  | [readonly] 

## Methods

### NewTier

`func NewTier(name string, features []Feature, prices []Price, ) *Tier`

NewTier instantiates a new Tier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTierWithDefaults

`func NewTierWithDefaults() *Tier`

NewTierWithDefaults instantiates a new Tier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Tier) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Tier) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Tier) SetName(v string)`

SetName sets Name field to given value.


### GetFeatures

`func (o *Tier) GetFeatures() []Feature`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *Tier) GetFeaturesOk() (*[]Feature, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *Tier) SetFeatures(v []Feature)`

SetFeatures sets Features field to given value.


### GetPrices

`func (o *Tier) GetPrices() []Price`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *Tier) GetPricesOk() (*[]Price, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *Tier) SetPrices(v []Price)`

SetPrices sets Prices field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


