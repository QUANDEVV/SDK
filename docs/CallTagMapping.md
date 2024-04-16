# CallTagMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallTagMappingId** | **int32** | The call tag mapping id. | [readonly] 
**Call** | [**Call**](Call.md) |  | [readonly] 
**Tag** | [**CallTag**](CallTag.md) |  | [readonly] 

## Methods

### NewCallTagMapping

`func NewCallTagMapping(callTagMappingId int32, call Call, tag CallTag, ) *CallTagMapping`

NewCallTagMapping instantiates a new CallTagMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallTagMappingWithDefaults

`func NewCallTagMappingWithDefaults() *CallTagMapping`

NewCallTagMappingWithDefaults instantiates a new CallTagMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallTagMappingId

`func (o *CallTagMapping) GetCallTagMappingId() int32`

GetCallTagMappingId returns the CallTagMappingId field if non-nil, zero value otherwise.

### GetCallTagMappingIdOk

`func (o *CallTagMapping) GetCallTagMappingIdOk() (*int32, bool)`

GetCallTagMappingIdOk returns a tuple with the CallTagMappingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallTagMappingId

`func (o *CallTagMapping) SetCallTagMappingId(v int32)`

SetCallTagMappingId sets CallTagMappingId field to given value.


### GetCall

`func (o *CallTagMapping) GetCall() Call`

GetCall returns the Call field if non-nil, zero value otherwise.

### GetCallOk

`func (o *CallTagMapping) GetCallOk() (*Call, bool)`

GetCallOk returns a tuple with the Call field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCall

`func (o *CallTagMapping) SetCall(v Call)`

SetCall sets Call field to given value.


### SetCallNil

`func (o *CallTagMapping) SetCallNil(b bool)`

 SetCallNil sets the value for Call to be an explicit nil

### UnsetCall
`func (o *CallTagMapping) UnsetCall()`

UnsetCall ensures that no value is present for Call, not even an explicit nil
### GetTag

`func (o *CallTagMapping) GetTag() CallTag`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CallTagMapping) GetTagOk() (*CallTag, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CallTagMapping) SetTag(v CallTag)`

SetTag sets Tag field to given value.


### SetTagNil

`func (o *CallTagMapping) SetTagNil(b bool)`

 SetTagNil sets the value for Tag to be an explicit nil

### UnsetTag
`func (o *CallTagMapping) UnsetTag()`

UnsetTag ensures that no value is present for Tag, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


