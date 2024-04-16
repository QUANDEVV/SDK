# PatchedMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | Pointer to **int32** | The message ID UUID for an instance of a chat. | [optional] [readonly] 
**ChatId** | Pointer to **int32** | The chat ID UUID for an instance of a chat. | [optional] 
**TextContent** | Pointer to **NullableString** | Message text content | [optional] 
**VoiceContent** | Pointer to **NullableString** | The voice note sent | [optional] 
**MessageSender** | Pointer to [**MessageSenderEnum**](MessageSenderEnum.md) | Either the message is sent by the agent or by the tenant  * &#x60;client&#x60; - client * &#x60;tenant&#x60; - tenant * &#x60;tenant_iva&#x60; - tenant_iva * &#x60;anonymous_client&#x60; - anonymous_client | [optional] 
**Escalated** | Pointer to **NullableBool** | Say whether a client escalated a chat to a human agent. | [optional] 
**Channel** | Pointer to [**OneOfChannelEnumNullEnum**](oneOf&lt;ChannelEnum,NullEnum&gt;.md) | Channel from which message was sent  * &#x60;Mobile&#x60; - Mobile * &#x60;Website&#x60; - Website | [optional] 
**Topic** | Pointer to [**OneOfTopicEnumNullEnum**](oneOf&lt;TopicEnum,NullEnum&gt;.md) | Message topic  * &#x60;Mobile&#x60; - Mobile * &#x60;Website&#x60; - Website | [optional] 

## Methods

### NewPatchedMessage

`func NewPatchedMessage() *PatchedMessage`

NewPatchedMessage instantiates a new PatchedMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedMessageWithDefaults

`func NewPatchedMessageWithDefaults() *PatchedMessage`

NewPatchedMessageWithDefaults instantiates a new PatchedMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageId

`func (o *PatchedMessage) GetMessageId() int32`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *PatchedMessage) GetMessageIdOk() (*int32, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *PatchedMessage) SetMessageId(v int32)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *PatchedMessage) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetChatId

`func (o *PatchedMessage) GetChatId() int32`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *PatchedMessage) GetChatIdOk() (*int32, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *PatchedMessage) SetChatId(v int32)`

SetChatId sets ChatId field to given value.

### HasChatId

`func (o *PatchedMessage) HasChatId() bool`

HasChatId returns a boolean if a field has been set.

### GetTextContent

`func (o *PatchedMessage) GetTextContent() string`

GetTextContent returns the TextContent field if non-nil, zero value otherwise.

### GetTextContentOk

`func (o *PatchedMessage) GetTextContentOk() (*string, bool)`

GetTextContentOk returns a tuple with the TextContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextContent

`func (o *PatchedMessage) SetTextContent(v string)`

SetTextContent sets TextContent field to given value.

### HasTextContent

`func (o *PatchedMessage) HasTextContent() bool`

HasTextContent returns a boolean if a field has been set.

### SetTextContentNil

`func (o *PatchedMessage) SetTextContentNil(b bool)`

 SetTextContentNil sets the value for TextContent to be an explicit nil

### UnsetTextContent
`func (o *PatchedMessage) UnsetTextContent()`

UnsetTextContent ensures that no value is present for TextContent, not even an explicit nil
### GetVoiceContent

`func (o *PatchedMessage) GetVoiceContent() string`

GetVoiceContent returns the VoiceContent field if non-nil, zero value otherwise.

### GetVoiceContentOk

`func (o *PatchedMessage) GetVoiceContentOk() (*string, bool)`

GetVoiceContentOk returns a tuple with the VoiceContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceContent

`func (o *PatchedMessage) SetVoiceContent(v string)`

SetVoiceContent sets VoiceContent field to given value.

### HasVoiceContent

`func (o *PatchedMessage) HasVoiceContent() bool`

HasVoiceContent returns a boolean if a field has been set.

### SetVoiceContentNil

`func (o *PatchedMessage) SetVoiceContentNil(b bool)`

 SetVoiceContentNil sets the value for VoiceContent to be an explicit nil

### UnsetVoiceContent
`func (o *PatchedMessage) UnsetVoiceContent()`

UnsetVoiceContent ensures that no value is present for VoiceContent, not even an explicit nil
### GetMessageSender

`func (o *PatchedMessage) GetMessageSender() MessageSenderEnum`

GetMessageSender returns the MessageSender field if non-nil, zero value otherwise.

### GetMessageSenderOk

`func (o *PatchedMessage) GetMessageSenderOk() (*MessageSenderEnum, bool)`

GetMessageSenderOk returns a tuple with the MessageSender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageSender

`func (o *PatchedMessage) SetMessageSender(v MessageSenderEnum)`

SetMessageSender sets MessageSender field to given value.

### HasMessageSender

`func (o *PatchedMessage) HasMessageSender() bool`

HasMessageSender returns a boolean if a field has been set.

### SetMessageSenderNil

`func (o *PatchedMessage) SetMessageSenderNil(b bool)`

 SetMessageSenderNil sets the value for MessageSender to be an explicit nil

### UnsetMessageSender
`func (o *PatchedMessage) UnsetMessageSender()`

UnsetMessageSender ensures that no value is present for MessageSender, not even an explicit nil
### GetEscalated

`func (o *PatchedMessage) GetEscalated() bool`

GetEscalated returns the Escalated field if non-nil, zero value otherwise.

### GetEscalatedOk

`func (o *PatchedMessage) GetEscalatedOk() (*bool, bool)`

GetEscalatedOk returns a tuple with the Escalated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEscalated

`func (o *PatchedMessage) SetEscalated(v bool)`

SetEscalated sets Escalated field to given value.

### HasEscalated

`func (o *PatchedMessage) HasEscalated() bool`

HasEscalated returns a boolean if a field has been set.

### SetEscalatedNil

`func (o *PatchedMessage) SetEscalatedNil(b bool)`

 SetEscalatedNil sets the value for Escalated to be an explicit nil

### UnsetEscalated
`func (o *PatchedMessage) UnsetEscalated()`

UnsetEscalated ensures that no value is present for Escalated, not even an explicit nil
### GetChannel

`func (o *PatchedMessage) GetChannel() OneOfChannelEnumNullEnum`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *PatchedMessage) GetChannelOk() (*OneOfChannelEnumNullEnum, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *PatchedMessage) SetChannel(v OneOfChannelEnumNullEnum)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *PatchedMessage) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### SetChannelNil

`func (o *PatchedMessage) SetChannelNil(b bool)`

 SetChannelNil sets the value for Channel to be an explicit nil

### UnsetChannel
`func (o *PatchedMessage) UnsetChannel()`

UnsetChannel ensures that no value is present for Channel, not even an explicit nil
### GetTopic

`func (o *PatchedMessage) GetTopic() OneOfTopicEnumNullEnum`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *PatchedMessage) GetTopicOk() (*OneOfTopicEnumNullEnum, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *PatchedMessage) SetTopic(v OneOfTopicEnumNullEnum)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *PatchedMessage) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### SetTopicNil

`func (o *PatchedMessage) SetTopicNil(b bool)`

 SetTopicNil sets the value for Topic to be an explicit nil

### UnsetTopic
`func (o *PatchedMessage) UnsetTopic()`

UnsetTopic ensures that no value is present for Topic, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


