# Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EventTypeEnum**](EventTypeEnum.md) |  | 
**Client** | [**ClientCommunity**](ClientCommunity.md) |  | [readonly] 
**Timestamp** | **time.Time** |  | [readonly] 

## Methods

### NewEvent

`func NewEvent(type_ EventTypeEnum, client ClientCommunity, timestamp time.Time, ) *Event`

NewEvent instantiates a new Event object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventWithDefaults

`func NewEventWithDefaults() *Event`

NewEventWithDefaults instantiates a new Event object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Event) GetType() EventTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Event) GetTypeOk() (*EventTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Event) SetType(v EventTypeEnum)`

SetType sets Type field to given value.


### GetClient

`func (o *Event) GetClient() ClientCommunity`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *Event) GetClientOk() (*ClientCommunity, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *Event) SetClient(v ClientCommunity)`

SetClient sets Client field to given value.


### SetClientNil

`func (o *Event) SetClientNil(b bool)`

 SetClientNil sets the value for Client to be an explicit nil

### UnsetClient
`func (o *Event) UnsetClient()`

UnsetClient ensures that no value is present for Client, not even an explicit nil
### GetTimestamp

`func (o *Event) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Event) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Event) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


