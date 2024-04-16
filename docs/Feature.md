# Feature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | **string** |  | 
**ExtraFeatures** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewFeature

`func NewFeature(name string, description string, ) *Feature`

NewFeature instantiates a new Feature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureWithDefaults

`func NewFeatureWithDefaults() *Feature`

NewFeatureWithDefaults instantiates a new Feature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Feature) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Feature) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Feature) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Feature) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Feature) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Feature) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetExtraFeatures

`func (o *Feature) GetExtraFeatures() interface{}`

GetExtraFeatures returns the ExtraFeatures field if non-nil, zero value otherwise.

### GetExtraFeaturesOk

`func (o *Feature) GetExtraFeaturesOk() (*interface{}, bool)`

GetExtraFeaturesOk returns a tuple with the ExtraFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraFeatures

`func (o *Feature) SetExtraFeatures(v interface{})`

SetExtraFeatures sets ExtraFeatures field to given value.

### HasExtraFeatures

`func (o *Feature) HasExtraFeatures() bool`

HasExtraFeatures returns a boolean if a field has been set.

### SetExtraFeaturesNil

`func (o *Feature) SetExtraFeaturesNil(b bool)`

 SetExtraFeaturesNil sets the value for ExtraFeatures to be an explicit nil

### UnsetExtraFeatures
`func (o *Feature) UnsetExtraFeatures()`

UnsetExtraFeatures ensures that no value is present for ExtraFeatures, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


