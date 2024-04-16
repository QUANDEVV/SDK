# Message

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | **int32** | The message ID UUID for an instance of a chat. | [readonly] 
**ChatId** | **int32** | The chat ID UUID for an instance of a chat. | 
**TextContent** | Pointer to **NullableString** | Message text content | [optional] 
**VoiceContent** | Pointer to **NullableString** | The voice note sent | [optional] 
**MessageSender** | [**MessageSenderEnum**](MessageSenderEnum.md) | Either the message is sent by the agent or by the tenant  * &#x60;client&#x60; - client * &#x60;tenant&#x60; - tenant * &#x60;tenant_iva&#x60; - tenant_iva * &#x60;anonymous_client&#x60; - anonymous_client | 
**Escalated** | Pointer to **NullableBool** | Say whether a client escalated a chat to a human agent. | [optional] 
**Channel** | Pointer to [**OneOfChannelEnumNullEnum**](oneOf&lt;ChannelEnum,NullEnum&gt;.md) | Channel from which message was sent  * &#x60;Mobile&#x60; - Mobile * &#x60;Website&#x60; - Website | [optional] 
**Topic** | Pointer to [**OneOfTopicEnumNullEnum**](oneOf&lt;TopicEnum,NullEnum&gt;.md) | Message topic  * &#x60;Mobile&#x60; - Mobile * &#x60;Website&#x60; - Website | [optional] 

## Methods

### NewMessage

`func NewMessage(messageId int32, chatId int32, messageSender MessageSenderEnum, ) *Message`

NewMessage instantiates a new Message object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageWithDefaults

`func NewMessageWithDefaults() *Message`

NewMessageWithDefaults instantiates a new Message object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageId

`func (o *Message) GetMessageId() int32`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *Message) GetMessageIdOk() (*int32, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *Message) SetMessageId(v int32)`

SetMessageId sets MessageId field to given value.


### GetChatId

`func (o *Message) GetChatId() int32`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *Message) GetChatIdOk() (*int32, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *Message) SetChatId(v int32)`

SetChatId sets ChatId field to given value.


### GetTextContent

`func (o *Message) GetTextContent() string`

GetTextContent returns the TextContent field if non-nil, zero value otherwise.

### GetTextContentOk

`func (o *Message) GetTextContentOk() (*string, bool)`

GetTextContentOk returns a tuple with the TextContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextContent

`func (o *Message) SetTextContent(v string)`

SetTextContent sets TextContent field to given value.

### HasTextContent

`func (o *Message) HasTextContent() bool`

HasTextContent returns a boolean if a field has been set.

### SetTextContentNil

`func (o *Message) SetTextContentNil(b bool)`

 SetTextContentNil sets the value for TextContent to be an explicit nil

### UnsetTextContent
`func (o *Message) UnsetTextContent()`

UnsetTextContent ensures that no value is present for TextContent, not even an explicit nil
### GetVoiceContent

`func (o *Message) GetVoiceContent() string`

GetVoiceContent returns the VoiceContent field if non-nil, zero value otherwise.

### GetVoiceContentOk

`func (o *Message) GetVoiceContentOk() (*string, bool)`

GetVoiceContentOk returns a tuple with the VoiceContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceContent

`func (o *Message) SetVoiceContent(v string)`

SetVoiceContent sets VoiceContent field to given value.

### HasVoiceContent

`func (o *Message) HasVoiceContent() bool`

HasVoiceContent returns a boolean if a field has been set.

### SetVoiceContentNil

`func (o *Message) SetVoiceContentNil(b bool)`

 SetVoiceContentNil sets the value for VoiceContent to be an explicit nil

### UnsetVoiceContent
`func (o *Message) UnsetVoiceContent()`

UnsetVoiceContent ensures that no value is present for VoiceContent, not even an explicit nil
### GetMessageSender

`func (o *Message) GetMessageSender() MessageSenderEnum`

GetMessageSender returns the MessageSender field if non-nil, zero value otherwise.

### GetMessageSenderOk

`func (o *Message) GetMessageSenderOk() (*MessageSenderEnum, bool)`

GetMessageSenderOk returns a tuple with the MessageSender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageSender

`func (o *Message) SetMessageSender(v MessageSenderEnum)`

SetMessageSender sets MessageSender field to given value.


### SetMessageSenderNil

`func (o *Message) SetMessageSenderNil(b bool)`

 SetMessageSenderNil sets the value for MessageSender to be an explicit nil

### UnsetMessageSender
`func (o *Message) UnsetMessageSender()`

UnsetMessageSender ensures that no value is present for MessageSender, not even an explicit nil
### GetEscalated

`func (o *Message) GetEscalated() bool`

GetEscalated returns the Escalated field if non-nil, zero value otherwise.

### GetEscalatedOk

`func (o *Message) GetEscalatedOk() (*bool, bool)`

GetEscalatedOk returns a tuple with the Escalated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEscalated

`func (o *Message) SetEscalated(v bool)`

SetEscalated sets Escalated field to given value.

### HasEscalated

`func (o *Message) HasEscalated() bool`

HasEscalated returns a boolean if a field has been set.

### SetEscalatedNil

`func (o *Message) SetEscalatedNil(b bool)`

 SetEscalatedNil sets the value for Escalated to be an explicit nil

### UnsetEscalated
`func (o *Message) UnsetEscalated()`

UnsetEscalated ensures that no value is present for Escalated, not even an explicit nil
### GetChannel

`func (o *Message) GetChannel() OneOfChannelEnumNullEnum`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *Message) GetChannelOk() (*OneOfChannelEnumNullEnum, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *Message) SetChannel(v OneOfChannelEnumNullEnum)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *Message) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### SetChannelNil

`func (o *Message) SetChannelNil(b bool)`

 SetChannelNil sets the value for Channel to be an explicit nil

### UnsetChannel
`func (o *Message) UnsetChannel()`

UnsetChannel ensures that no value is present for Channel, not even an explicit nil
### GetTopic

`func (o *Message) GetTopic() OneOfTopicEnumNullEnum`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *Message) GetTopicOk() (*OneOfTopicEnumNullEnum, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *Message) SetTopic(v OneOfTopicEnumNullEnum)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *Message) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### SetTopicNil

`func (o *Message) SetTopicNil(b bool)`

 SetTopicNil sets the value for Topic to be an explicit nil

### UnsetTopic
`func (o *Message) UnsetTopic()`

UnsetTopic ensures that no value is present for Topic, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


