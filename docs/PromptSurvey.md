# PromptSurvey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PromptSurveyId** | **int32** | The lilling details id ID UUID | [readonly] 
**TenantId** | **int32** |  | [readonly] 
**SurveyTopic** | **string** | The survey topic | 
**SurveyDescription** | **string** | The survey description | 
**SurveyPrompt** | **string** | The survey context | 
**SurveyQuestions** | Pointer to **interface{}** | The survey questions | [optional] 
**Client** | [**Client**](Client.md) |  | 
**UnauthenticatedClient** | [**AnonymousUser**](AnonymousUser.md) |  | 
**SurveyType** | Pointer to [**OneOfSurveyTypeEnumBlankEnumNullEnum**](oneOf&lt;SurveyTypeEnum,BlankEnum,NullEnum&gt;.md) | The survey type  * &#x60;open_ended&#x60; - open_ended * &#x60;close_ended&#x60; - close_ended | [optional] 

## Methods

### NewPromptSurvey

`func NewPromptSurvey(promptSurveyId int32, tenantId int32, surveyTopic string, surveyDescription string, surveyPrompt string, client Client, unauthenticatedClient AnonymousUser, ) *PromptSurvey`

NewPromptSurvey instantiates a new PromptSurvey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromptSurveyWithDefaults

`func NewPromptSurveyWithDefaults() *PromptSurvey`

NewPromptSurveyWithDefaults instantiates a new PromptSurvey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPromptSurveyId

`func (o *PromptSurvey) GetPromptSurveyId() int32`

GetPromptSurveyId returns the PromptSurveyId field if non-nil, zero value otherwise.

### GetPromptSurveyIdOk

`func (o *PromptSurvey) GetPromptSurveyIdOk() (*int32, bool)`

GetPromptSurveyIdOk returns a tuple with the PromptSurveyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptSurveyId

`func (o *PromptSurvey) SetPromptSurveyId(v int32)`

SetPromptSurveyId sets PromptSurveyId field to given value.


### GetTenantId

`func (o *PromptSurvey) GetTenantId() int32`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *PromptSurvey) GetTenantIdOk() (*int32, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *PromptSurvey) SetTenantId(v int32)`

SetTenantId sets TenantId field to given value.


### GetSurveyTopic

`func (o *PromptSurvey) GetSurveyTopic() string`

GetSurveyTopic returns the SurveyTopic field if non-nil, zero value otherwise.

### GetSurveyTopicOk

`func (o *PromptSurvey) GetSurveyTopicOk() (*string, bool)`

GetSurveyTopicOk returns a tuple with the SurveyTopic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveyTopic

`func (o *PromptSurvey) SetSurveyTopic(v string)`

SetSurveyTopic sets SurveyTopic field to given value.


### GetSurveyDescription

`func (o *PromptSurvey) GetSurveyDescription() string`

GetSurveyDescription returns the SurveyDescription field if non-nil, zero value otherwise.

### GetSurveyDescriptionOk

`func (o *PromptSurvey) GetSurveyDescriptionOk() (*string, bool)`

GetSurveyDescriptionOk returns a tuple with the SurveyDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveyDescription

`func (o *PromptSurvey) SetSurveyDescription(v string)`

SetSurveyDescription sets SurveyDescription field to given value.


### GetSurveyPrompt

`func (o *PromptSurvey) GetSurveyPrompt() string`

GetSurveyPrompt returns the SurveyPrompt field if non-nil, zero value otherwise.

### GetSurveyPromptOk

`func (o *PromptSurvey) GetSurveyPromptOk() (*string, bool)`

GetSurveyPromptOk returns a tuple with the SurveyPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveyPrompt

`func (o *PromptSurvey) SetSurveyPrompt(v string)`

SetSurveyPrompt sets SurveyPrompt field to given value.


### GetSurveyQuestions

`func (o *PromptSurvey) GetSurveyQuestions() interface{}`

GetSurveyQuestions returns the SurveyQuestions field if non-nil, zero value otherwise.

### GetSurveyQuestionsOk

`func (o *PromptSurvey) GetSurveyQuestionsOk() (*interface{}, bool)`

GetSurveyQuestionsOk returns a tuple with the SurveyQuestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveyQuestions

`func (o *PromptSurvey) SetSurveyQuestions(v interface{})`

SetSurveyQuestions sets SurveyQuestions field to given value.

### HasSurveyQuestions

`func (o *PromptSurvey) HasSurveyQuestions() bool`

HasSurveyQuestions returns a boolean if a field has been set.

### SetSurveyQuestionsNil

`func (o *PromptSurvey) SetSurveyQuestionsNil(b bool)`

 SetSurveyQuestionsNil sets the value for SurveyQuestions to be an explicit nil

### UnsetSurveyQuestions
`func (o *PromptSurvey) UnsetSurveyQuestions()`

UnsetSurveyQuestions ensures that no value is present for SurveyQuestions, not even an explicit nil
### GetClient

`func (o *PromptSurvey) GetClient() Client`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *PromptSurvey) GetClientOk() (*Client, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *PromptSurvey) SetClient(v Client)`

SetClient sets Client field to given value.


### GetUnauthenticatedClient

`func (o *PromptSurvey) GetUnauthenticatedClient() AnonymousUser`

GetUnauthenticatedClient returns the UnauthenticatedClient field if non-nil, zero value otherwise.

### GetUnauthenticatedClientOk

`func (o *PromptSurvey) GetUnauthenticatedClientOk() (*AnonymousUser, bool)`

GetUnauthenticatedClientOk returns a tuple with the UnauthenticatedClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnauthenticatedClient

`func (o *PromptSurvey) SetUnauthenticatedClient(v AnonymousUser)`

SetUnauthenticatedClient sets UnauthenticatedClient field to given value.


### GetSurveyType

`func (o *PromptSurvey) GetSurveyType() OneOfSurveyTypeEnumBlankEnumNullEnum`

GetSurveyType returns the SurveyType field if non-nil, zero value otherwise.

### GetSurveyTypeOk

`func (o *PromptSurvey) GetSurveyTypeOk() (*OneOfSurveyTypeEnumBlankEnumNullEnum, bool)`

GetSurveyTypeOk returns a tuple with the SurveyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveyType

`func (o *PromptSurvey) SetSurveyType(v OneOfSurveyTypeEnumBlankEnumNullEnum)`

SetSurveyType sets SurveyType field to given value.

### HasSurveyType

`func (o *PromptSurvey) HasSurveyType() bool`

HasSurveyType returns a boolean if a field has been set.

### SetSurveyTypeNil

`func (o *PromptSurvey) SetSurveyTypeNil(b bool)`

 SetSurveyTypeNil sets the value for SurveyType to be an explicit nil

### UnsetSurveyType
`func (o *PromptSurvey) UnsetSurveyType()`

UnsetSurveyType ensures that no value is present for SurveyType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


