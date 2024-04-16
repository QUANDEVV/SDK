# Survey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SurveyId** | **int32** | The lilling details id ID UUID | [readonly] 
**TenantId** | **int32** | Display name of the tenant | 
**SurveyTopic** | **string** | The survey topic | 
**SurveyDescription** | **string** | The survey description | 
**SurveyContext** | **string** | The survey context | 
**SurveyQuestions** | Pointer to **interface{}** | The survey questions | [optional] 
**TargetAudience** | Pointer to **[]int32** | The target audience/who to share with | [optional] 
**SurveyType** | Pointer to [**OneOfSurveyTypeEnumBlankEnumNullEnum**](oneOf&lt;SurveyTypeEnum,BlankEnum,NullEnum&gt;.md) | The survey type  * &#x60;open_ended&#x60; - open_ended * &#x60;close_ended&#x60; - close_ended | [optional] 
**StartDay** | **NullableString** | The start day of the survey | [readonly] 
**EndDay** | Pointer to **NullableString** | The end day of the survey | [optional] 
**SurveyStatus** | Pointer to **bool** |  | [optional] 
**TargetAudienceCount** | **int32** |  | [readonly] 
**ResponseCount** | **int32** |  | [readonly] 
**PercentageDifference** | **float64** |  | [readonly] 
**QuestionCount** | **int32** |  | [readonly] 

## Methods

### NewSurvey

`func NewSurvey(surveyId int32, tenantId int32, surveyTopic string, surveyDescription string, surveyContext string, startDay NullableString, targetAudienceCount int32, responseCount int32, percentageDifference float64, questionCount int32, ) *Survey`

NewSurvey instantiates a new Survey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSurveyWithDefaults

`func NewSurveyWithDefaults() *Survey`

NewSurveyWithDefaults instantiates a new Survey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSurveyId

`func (o *Survey) GetSurveyId() int32`

GetSurveyId returns the SurveyId field if non-nil, zero value otherwise.

### GetSurveyIdOk

`func (o *Survey) GetSurveyIdOk() (*int32, bool)`

GetSurveyIdOk returns a tuple with the SurveyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveyId

`func (o *Survey) SetSurveyId(v int32)`

SetSurveyId sets SurveyId field to given value.


### GetTenantId

`func (o *Survey) GetTenantId() int32`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Survey) GetTenantIdOk() (*int32, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Survey) SetTenantId(v int32)`

SetTenantId sets TenantId field to given value.


### GetSurveyTopic

`func (o *Survey) GetSurveyTopic() string`

GetSurveyTopic returns the SurveyTopic field if non-nil, zero value otherwise.

### GetSurveyTopicOk

`func (o *Survey) GetSurveyTopicOk() (*string, bool)`

GetSurveyTopicOk returns a tuple with the SurveyTopic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveyTopic

`func (o *Survey) SetSurveyTopic(v string)`

SetSurveyTopic sets SurveyTopic field to given value.


### GetSurveyDescription

`func (o *Survey) GetSurveyDescription() string`

GetSurveyDescription returns the SurveyDescription field if non-nil, zero value otherwise.

### GetSurveyDescriptionOk

`func (o *Survey) GetSurveyDescriptionOk() (*string, bool)`

GetSurveyDescriptionOk returns a tuple with the SurveyDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveyDescription

`func (o *Survey) SetSurveyDescription(v string)`

SetSurveyDescription sets SurveyDescription field to given value.


### GetSurveyContext

`func (o *Survey) GetSurveyContext() string`

GetSurveyContext returns the SurveyContext field if non-nil, zero value otherwise.

### GetSurveyContextOk

`func (o *Survey) GetSurveyContextOk() (*string, bool)`

GetSurveyContextOk returns a tuple with the SurveyContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveyContext

`func (o *Survey) SetSurveyContext(v string)`

SetSurveyContext sets SurveyContext field to given value.


### GetSurveyQuestions

`func (o *Survey) GetSurveyQuestions() interface{}`

GetSurveyQuestions returns the SurveyQuestions field if non-nil, zero value otherwise.

### GetSurveyQuestionsOk

`func (o *Survey) GetSurveyQuestionsOk() (*interface{}, bool)`

GetSurveyQuestionsOk returns a tuple with the SurveyQuestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveyQuestions

`func (o *Survey) SetSurveyQuestions(v interface{})`

SetSurveyQuestions sets SurveyQuestions field to given value.

### HasSurveyQuestions

`func (o *Survey) HasSurveyQuestions() bool`

HasSurveyQuestions returns a boolean if a field has been set.

### SetSurveyQuestionsNil

`func (o *Survey) SetSurveyQuestionsNil(b bool)`

 SetSurveyQuestionsNil sets the value for SurveyQuestions to be an explicit nil

### UnsetSurveyQuestions
`func (o *Survey) UnsetSurveyQuestions()`

UnsetSurveyQuestions ensures that no value is present for SurveyQuestions, not even an explicit nil
### GetTargetAudience

`func (o *Survey) GetTargetAudience() []int32`

GetTargetAudience returns the TargetAudience field if non-nil, zero value otherwise.

### GetTargetAudienceOk

`func (o *Survey) GetTargetAudienceOk() (*[]int32, bool)`

GetTargetAudienceOk returns a tuple with the TargetAudience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAudience

`func (o *Survey) SetTargetAudience(v []int32)`

SetTargetAudience sets TargetAudience field to given value.

### HasTargetAudience

`func (o *Survey) HasTargetAudience() bool`

HasTargetAudience returns a boolean if a field has been set.

### GetSurveyType

`func (o *Survey) GetSurveyType() OneOfSurveyTypeEnumBlankEnumNullEnum`

GetSurveyType returns the SurveyType field if non-nil, zero value otherwise.

### GetSurveyTypeOk

`func (o *Survey) GetSurveyTypeOk() (*OneOfSurveyTypeEnumBlankEnumNullEnum, bool)`

GetSurveyTypeOk returns a tuple with the SurveyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveyType

`func (o *Survey) SetSurveyType(v OneOfSurveyTypeEnumBlankEnumNullEnum)`

SetSurveyType sets SurveyType field to given value.

### HasSurveyType

`func (o *Survey) HasSurveyType() bool`

HasSurveyType returns a boolean if a field has been set.

### SetSurveyTypeNil

`func (o *Survey) SetSurveyTypeNil(b bool)`

 SetSurveyTypeNil sets the value for SurveyType to be an explicit nil

### UnsetSurveyType
`func (o *Survey) UnsetSurveyType()`

UnsetSurveyType ensures that no value is present for SurveyType, not even an explicit nil
### GetStartDay

`func (o *Survey) GetStartDay() string`

GetStartDay returns the StartDay field if non-nil, zero value otherwise.

### GetStartDayOk

`func (o *Survey) GetStartDayOk() (*string, bool)`

GetStartDayOk returns a tuple with the StartDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDay

`func (o *Survey) SetStartDay(v string)`

SetStartDay sets StartDay field to given value.


### SetStartDayNil

`func (o *Survey) SetStartDayNil(b bool)`

 SetStartDayNil sets the value for StartDay to be an explicit nil

### UnsetStartDay
`func (o *Survey) UnsetStartDay()`

UnsetStartDay ensures that no value is present for StartDay, not even an explicit nil
### GetEndDay

`func (o *Survey) GetEndDay() string`

GetEndDay returns the EndDay field if non-nil, zero value otherwise.

### GetEndDayOk

`func (o *Survey) GetEndDayOk() (*string, bool)`

GetEndDayOk returns a tuple with the EndDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDay

`func (o *Survey) SetEndDay(v string)`

SetEndDay sets EndDay field to given value.

### HasEndDay

`func (o *Survey) HasEndDay() bool`

HasEndDay returns a boolean if a field has been set.

### SetEndDayNil

`func (o *Survey) SetEndDayNil(b bool)`

 SetEndDayNil sets the value for EndDay to be an explicit nil

### UnsetEndDay
`func (o *Survey) UnsetEndDay()`

UnsetEndDay ensures that no value is present for EndDay, not even an explicit nil
### GetSurveyStatus

`func (o *Survey) GetSurveyStatus() bool`

GetSurveyStatus returns the SurveyStatus field if non-nil, zero value otherwise.

### GetSurveyStatusOk

`func (o *Survey) GetSurveyStatusOk() (*bool, bool)`

GetSurveyStatusOk returns a tuple with the SurveyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveyStatus

`func (o *Survey) SetSurveyStatus(v bool)`

SetSurveyStatus sets SurveyStatus field to given value.

### HasSurveyStatus

`func (o *Survey) HasSurveyStatus() bool`

HasSurveyStatus returns a boolean if a field has been set.

### GetTargetAudienceCount

`func (o *Survey) GetTargetAudienceCount() int32`

GetTargetAudienceCount returns the TargetAudienceCount field if non-nil, zero value otherwise.

### GetTargetAudienceCountOk

`func (o *Survey) GetTargetAudienceCountOk() (*int32, bool)`

GetTargetAudienceCountOk returns a tuple with the TargetAudienceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAudienceCount

`func (o *Survey) SetTargetAudienceCount(v int32)`

SetTargetAudienceCount sets TargetAudienceCount field to given value.


### GetResponseCount

`func (o *Survey) GetResponseCount() int32`

GetResponseCount returns the ResponseCount field if non-nil, zero value otherwise.

### GetResponseCountOk

`func (o *Survey) GetResponseCountOk() (*int32, bool)`

GetResponseCountOk returns a tuple with the ResponseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCount

`func (o *Survey) SetResponseCount(v int32)`

SetResponseCount sets ResponseCount field to given value.


### GetPercentageDifference

`func (o *Survey) GetPercentageDifference() float64`

GetPercentageDifference returns the PercentageDifference field if non-nil, zero value otherwise.

### GetPercentageDifferenceOk

`func (o *Survey) GetPercentageDifferenceOk() (*float64, bool)`

GetPercentageDifferenceOk returns a tuple with the PercentageDifference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageDifference

`func (o *Survey) SetPercentageDifference(v float64)`

SetPercentageDifference sets PercentageDifference field to given value.


### GetQuestionCount

`func (o *Survey) GetQuestionCount() int32`

GetQuestionCount returns the QuestionCount field if non-nil, zero value otherwise.

### GetQuestionCountOk

`func (o *Survey) GetQuestionCountOk() (*int32, bool)`

GetQuestionCountOk returns a tuple with the QuestionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionCount

`func (o *Survey) SetQuestionCount(v int32)`

SetQuestionCount sets QuestionCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


