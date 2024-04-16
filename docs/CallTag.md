# CallTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallTagId** | **int32** | The call tag id. | [readonly] 
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCallTag

`func NewCallTag(callTagId int32, name string, ) *CallTag`

NewCallTag instantiates a new CallTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallTagWithDefaults

`func NewCallTagWithDefaults() *CallTag`

NewCallTagWithDefaults instantiates a new CallTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallTagId

`func (o *CallTag) GetCallTagId() int32`

GetCallTagId returns the CallTagId field if non-nil, zero value otherwise.

### GetCallTagIdOk

`func (o *CallTag) GetCallTagIdOk() (*int32, bool)`

GetCallTagIdOk returns a tuple with the CallTagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallTagId

`func (o *CallTag) SetCallTagId(v int32)`

SetCallTagId sets CallTagId field to given value.


### GetName

`func (o *CallTag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CallTag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CallTag) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CallTag) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CallTag) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CallTag) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CallTag) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CallTag) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CallTag) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


