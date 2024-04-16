# ServiceMessageChat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceRequestMessageId** | **int32** |  | [readonly] 
**ServiceRequestChat** | **int32** | The chat this message is a part of. | 
**MessageContent** | **string** | The content of the message. | 
**MessageSender** | [**MessageSenderEnum**](MessageSenderEnum.md) | Either the message is sent by the agent or by the tenant  * &#x60;client&#x60; - client * &#x60;tenant&#x60; - tenant * &#x60;tenant_iva&#x60; - tenant_iva * &#x60;anonymous_client&#x60; - anonymous_client | 

## Methods

### NewServiceMessageChat

`func NewServiceMessageChat(serviceRequestMessageId int32, serviceRequestChat int32, messageContent string, messageSender MessageSenderEnum, ) *ServiceMessageChat`

NewServiceMessageChat instantiates a new ServiceMessageChat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceMessageChatWithDefaults

`func NewServiceMessageChatWithDefaults() *ServiceMessageChat`

NewServiceMessageChatWithDefaults instantiates a new ServiceMessageChat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceRequestMessageId

`func (o *ServiceMessageChat) GetServiceRequestMessageId() int32`

GetServiceRequestMessageId returns the ServiceRequestMessageId field if non-nil, zero value otherwise.

### GetServiceRequestMessageIdOk

`func (o *ServiceMessageChat) GetServiceRequestMessageIdOk() (*int32, bool)`

GetServiceRequestMessageIdOk returns a tuple with the ServiceRequestMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRequestMessageId

`func (o *ServiceMessageChat) SetServiceRequestMessageId(v int32)`

SetServiceRequestMessageId sets ServiceRequestMessageId field to given value.


### GetServiceRequestChat

`func (o *ServiceMessageChat) GetServiceRequestChat() int32`

GetServiceRequestChat returns the ServiceRequestChat field if non-nil, zero value otherwise.

### GetServiceRequestChatOk

`func (o *ServiceMessageChat) GetServiceRequestChatOk() (*int32, bool)`

GetServiceRequestChatOk returns a tuple with the ServiceRequestChat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRequestChat

`func (o *ServiceMessageChat) SetServiceRequestChat(v int32)`

SetServiceRequestChat sets ServiceRequestChat field to given value.


### GetMessageContent

`func (o *ServiceMessageChat) GetMessageContent() string`

GetMessageContent returns the MessageContent field if non-nil, zero value otherwise.

### GetMessageContentOk

`func (o *ServiceMessageChat) GetMessageContentOk() (*string, bool)`

GetMessageContentOk returns a tuple with the MessageContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageContent

`func (o *ServiceMessageChat) SetMessageContent(v string)`

SetMessageContent sets MessageContent field to given value.


### GetMessageSender

`func (o *ServiceMessageChat) GetMessageSender() MessageSenderEnum`

GetMessageSender returns the MessageSender field if non-nil, zero value otherwise.

### GetMessageSenderOk

`func (o *ServiceMessageChat) GetMessageSenderOk() (*MessageSenderEnum, bool)`

GetMessageSenderOk returns a tuple with the MessageSender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageSender

`func (o *ServiceMessageChat) SetMessageSender(v MessageSenderEnum)`

SetMessageSender sets MessageSender field to given value.


### SetMessageSenderNil

`func (o *ServiceMessageChat) SetMessageSenderNil(b bool)`

 SetMessageSenderNil sets the value for MessageSender to be an explicit nil

### UnsetMessageSender
`func (o *ServiceMessageChat) UnsetMessageSender()`

UnsetMessageSender ensures that no value is present for MessageSender, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


