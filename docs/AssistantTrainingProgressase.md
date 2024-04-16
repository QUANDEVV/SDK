# AssistantTrainingProgressase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssistantTrainingId** | **int32** | The tenant chats ID. | [readonly] 
**Status** | Pointer to [**OneOfAssistantTrainingProgressaseStatusEnumBlankEnumNullEnum**](oneOf&lt;AssistantTrainingProgressaseStatusEnum,BlankEnum,NullEnum&gt;.md) |  | [optional] 
**TenantId** | **int32** | Display name of the tenant | 
**Succesful** | Pointer to **bool** | Know the progress of updating an assistant | [optional] 
**Message** | Pointer to **NullableString** | Updating | [optional] 

## Methods

### NewAssistantTrainingProgressase

`func NewAssistantTrainingProgressase(assistantTrainingId int32, tenantId int32, ) *AssistantTrainingProgressase`

NewAssistantTrainingProgressase instantiates a new AssistantTrainingProgressase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssistantTrainingProgressaseWithDefaults

`func NewAssistantTrainingProgressaseWithDefaults() *AssistantTrainingProgressase`

NewAssistantTrainingProgressaseWithDefaults instantiates a new AssistantTrainingProgressase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssistantTrainingId

`func (o *AssistantTrainingProgressase) GetAssistantTrainingId() int32`

GetAssistantTrainingId returns the AssistantTrainingId field if non-nil, zero value otherwise.

### GetAssistantTrainingIdOk

`func (o *AssistantTrainingProgressase) GetAssistantTrainingIdOk() (*int32, bool)`

GetAssistantTrainingIdOk returns a tuple with the AssistantTrainingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssistantTrainingId

`func (o *AssistantTrainingProgressase) SetAssistantTrainingId(v int32)`

SetAssistantTrainingId sets AssistantTrainingId field to given value.


### GetStatus

`func (o *AssistantTrainingProgressase) GetStatus() OneOfAssistantTrainingProgressaseStatusEnumBlankEnumNullEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AssistantTrainingProgressase) GetStatusOk() (*OneOfAssistantTrainingProgressaseStatusEnumBlankEnumNullEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AssistantTrainingProgressase) SetStatus(v OneOfAssistantTrainingProgressaseStatusEnumBlankEnumNullEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AssistantTrainingProgressase) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *AssistantTrainingProgressase) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *AssistantTrainingProgressase) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTenantId

`func (o *AssistantTrainingProgressase) GetTenantId() int32`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AssistantTrainingProgressase) GetTenantIdOk() (*int32, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AssistantTrainingProgressase) SetTenantId(v int32)`

SetTenantId sets TenantId field to given value.


### GetSuccesful

`func (o *AssistantTrainingProgressase) GetSuccesful() bool`

GetSuccesful returns the Succesful field if non-nil, zero value otherwise.

### GetSuccesfulOk

`func (o *AssistantTrainingProgressase) GetSuccesfulOk() (*bool, bool)`

GetSuccesfulOk returns a tuple with the Succesful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccesful

`func (o *AssistantTrainingProgressase) SetSuccesful(v bool)`

SetSuccesful sets Succesful field to given value.

### HasSuccesful

`func (o *AssistantTrainingProgressase) HasSuccesful() bool`

HasSuccesful returns a boolean if a field has been set.

### GetMessage

`func (o *AssistantTrainingProgressase) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AssistantTrainingProgressase) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AssistantTrainingProgressase) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AssistantTrainingProgressase) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *AssistantTrainingProgressase) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *AssistantTrainingProgressase) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


