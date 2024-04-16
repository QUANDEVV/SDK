# InteractionTopics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InteractionTopicsId** | **int32** | The chat ID UUID for an instance of a chat. | [readonly] 
**Name** | **string** | Name of the topic | 
**Description** | **string** | Description of the topic | 
**Category** | **int32** | The topic category | 

## Methods

### NewInteractionTopics

`func NewInteractionTopics(interactionTopicsId int32, name string, description string, category int32, ) *InteractionTopics`

NewInteractionTopics instantiates a new InteractionTopics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInteractionTopicsWithDefaults

`func NewInteractionTopicsWithDefaults() *InteractionTopics`

NewInteractionTopicsWithDefaults instantiates a new InteractionTopics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInteractionTopicsId

`func (o *InteractionTopics) GetInteractionTopicsId() int32`

GetInteractionTopicsId returns the InteractionTopicsId field if non-nil, zero value otherwise.

### GetInteractionTopicsIdOk

`func (o *InteractionTopics) GetInteractionTopicsIdOk() (*int32, bool)`

GetInteractionTopicsIdOk returns a tuple with the InteractionTopicsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractionTopicsId

`func (o *InteractionTopics) SetInteractionTopicsId(v int32)`

SetInteractionTopicsId sets InteractionTopicsId field to given value.


### GetName

`func (o *InteractionTopics) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InteractionTopics) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InteractionTopics) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *InteractionTopics) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InteractionTopics) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InteractionTopics) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCategory

`func (o *InteractionTopics) GetCategory() int32`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *InteractionTopics) GetCategoryOk() (*int32, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *InteractionTopics) SetCategory(v int32)`

SetCategory sets Category field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


