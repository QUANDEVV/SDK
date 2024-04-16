# QuestionInducedSurvey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuestionInducedSurveyId** | **int32** | The lilling details id ID UUID | [readonly] 
**TenantId** | **int32** |  | [readonly] 
**SurveyTopic** | **string** | The survey topic | 
**SurveyDescription** | **string** | The survey description | 
**QuestionsInduced** | Pointer to **interface{}** | The survey questions | [optional] 
**SurveyQuestions** | Pointer to **interface{}** | The survey questions | [optional] 
**Client** | [**Client**](Client.md) |  | 
**UnauthenticatedClient** | [**AnonymousUser**](AnonymousUser.md) |  | 
**SurveyType** | Pointer to [**OneOfSurveyTypeEnumBlankEnumNullEnum**](oneOf&lt;SurveyTypeEnum,BlankEnum,NullEnum&gt;.md) | The survey type  * &#x60;open_ended&#x60; - open_ended * &#x60;close_ended&#x60; - close_ended | [optional] 

## Methods

### NewQuestionInducedSurvey

`func NewQuestionInducedSurvey(questionInducedSurveyId int32, tenantId int32, surveyTopic string, surveyDescription string, client Client, unauthenticatedClient AnonymousUser, ) *QuestionInducedSurvey`

NewQuestionInducedSurvey instantiates a new QuestionInducedSurvey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuestionInducedSurveyWithDefaults

`func NewQuestionInducedSurveyWithDefaults() *QuestionInducedSurvey`

NewQuestionInducedSurveyWithDefaults instantiates a new QuestionInducedSurvey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuestionInducedSurveyId

`func (o *QuestionInducedSurvey) GetQuestionInducedSurveyId() int32`

GetQuestionInducedSurveyId returns the QuestionInducedSurveyId field if non-nil, zero value otherwise.

### GetQuestionInducedSurveyIdOk

`func (o *QuestionInducedSurvey) GetQuestionInducedSurveyIdOk() (*int32, bool)`

GetQuestionInducedSurveyIdOk returns a tuple with the QuestionInducedSurveyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionInducedSurveyId

`func (o *QuestionInducedSurvey) SetQuestionInducedSurveyId(v int32)`

SetQuestionInducedSurveyId sets QuestionInducedSurveyId field to given value.


### GetTenantId

`func (o *QuestionInducedSurvey) GetTenantId() int32`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *QuestionInducedSurvey) GetTenantIdOk() (*int32, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *QuestionInducedSurvey) SetTenantId(v int32)`

SetTenantId sets TenantId field to given value.


### GetSurveyTopic

`func (o *QuestionInducedSurvey) GetSurveyTopic() string`

GetSurveyTopic returns the SurveyTopic field if non-nil, zero value otherwise.

### GetSurveyTopicOk

`func (o *QuestionInducedSurvey) GetSurveyTopicOk() (*string, bool)`

GetSurveyTopicOk returns a tuple with the SurveyTopic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveyTopic

`func (o *QuestionInducedSurvey) SetSurveyTopic(v string)`

SetSurveyTopic sets SurveyTopic field to given value.


### GetSurveyDescription

`func (o *QuestionInducedSurvey) GetSurveyDescription() string`

GetSurveyDescription returns the SurveyDescription field if non-nil, zero value otherwise.

### GetSurveyDescriptionOk

`func (o *QuestionInducedSurvey) GetSurveyDescriptionOk() (*string, bool)`

GetSurveyDescriptionOk returns a tuple with the SurveyDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveyDescription

`func (o *QuestionInducedSurvey) SetSurveyDescription(v string)`

SetSurveyDescription sets SurveyDescription field to given value.


### GetQuestionsInduced

`func (o *QuestionInducedSurvey) GetQuestionsInduced() interface{}`

GetQuestionsInduced returns the QuestionsInduced field if non-nil, zero value otherwise.

### GetQuestionsInducedOk

`func (o *QuestionInducedSurvey) GetQuestionsInducedOk() (*interface{}, bool)`

GetQuestionsInducedOk returns a tuple with the QuestionsInduced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionsInduced

`func (o *QuestionInducedSurvey) SetQuestionsInduced(v interface{})`

SetQuestionsInduced sets QuestionsInduced field to given value.

### HasQuestionsInduced

`func (o *QuestionInducedSurvey) HasQuestionsInduced() bool`

HasQuestionsInduced returns a boolean if a field has been set.

### SetQuestionsInducedNil

`func (o *QuestionInducedSurvey) SetQuestionsInducedNil(b bool)`

 SetQuestionsInducedNil sets the value for QuestionsInduced to be an explicit nil

### UnsetQuestionsInduced
`func (o *QuestionInducedSurvey) UnsetQuestionsInduced()`

UnsetQuestionsInduced ensures that no value is present for QuestionsInduced, not even an explicit nil
### GetSurveyQuestions

`func (o *QuestionInducedSurvey) GetSurveyQuestions() interface{}`

GetSurveyQuestions returns the SurveyQuestions field if non-nil, zero value otherwise.

### GetSurveyQuestionsOk

`func (o *QuestionInducedSurvey) GetSurveyQuestionsOk() (*interface{}, bool)`

GetSurveyQuestionsOk returns a tuple with the SurveyQuestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveyQuestions

`func (o *QuestionInducedSurvey) SetSurveyQuestions(v interface{})`

SetSurveyQuestions sets SurveyQuestions field to given value.

### HasSurveyQuestions

`func (o *QuestionInducedSurvey) HasSurveyQuestions() bool`

HasSurveyQuestions returns a boolean if a field has been set.

### SetSurveyQuestionsNil

`func (o *QuestionInducedSurvey) SetSurveyQuestionsNil(b bool)`

 SetSurveyQuestionsNil sets the value for SurveyQuestions to be an explicit nil

### UnsetSurveyQuestions
`func (o *QuestionInducedSurvey) UnsetSurveyQuestions()`

UnsetSurveyQuestions ensures that no value is present for SurveyQuestions, not even an explicit nil
### GetClient

`func (o *QuestionInducedSurvey) GetClient() Client`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *QuestionInducedSurvey) GetClientOk() (*Client, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *QuestionInducedSurvey) SetClient(v Client)`

SetClient sets Client field to given value.


### GetUnauthenticatedClient

`func (o *QuestionInducedSurvey) GetUnauthenticatedClient() AnonymousUser`

GetUnauthenticatedClient returns the UnauthenticatedClient field if non-nil, zero value otherwise.

### GetUnauthenticatedClientOk

`func (o *QuestionInducedSurvey) GetUnauthenticatedClientOk() (*AnonymousUser, bool)`

GetUnauthenticatedClientOk returns a tuple with the UnauthenticatedClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnauthenticatedClient

`func (o *QuestionInducedSurvey) SetUnauthenticatedClient(v AnonymousUser)`

SetUnauthenticatedClient sets UnauthenticatedClient field to given value.


### GetSurveyType

`func (o *QuestionInducedSurvey) GetSurveyType() OneOfSurveyTypeEnumBlankEnumNullEnum`

GetSurveyType returns the SurveyType field if non-nil, zero value otherwise.

### GetSurveyTypeOk

`func (o *QuestionInducedSurvey) GetSurveyTypeOk() (*OneOfSurveyTypeEnumBlankEnumNullEnum, bool)`

GetSurveyTypeOk returns a tuple with the SurveyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveyType

`func (o *QuestionInducedSurvey) SetSurveyType(v OneOfSurveyTypeEnumBlankEnumNullEnum)`

SetSurveyType sets SurveyType field to given value.

### HasSurveyType

`func (o *QuestionInducedSurvey) HasSurveyType() bool`

HasSurveyType returns a boolean if a field has been set.

### SetSurveyTypeNil

`func (o *QuestionInducedSurvey) SetSurveyTypeNil(b bool)`

 SetSurveyTypeNil sets the value for SurveyType to be an explicit nil

### UnsetSurveyType
`func (o *QuestionInducedSurvey) UnsetSurveyType()`

UnsetSurveyType ensures that no value is present for SurveyType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


