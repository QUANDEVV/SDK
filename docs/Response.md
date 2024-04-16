# Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseId** | **int32** | The lilling details id ID UUID | [readonly] 
**SurveyId** | Pointer to **NullableInt32** | Display name of the tenant | [optional] 
**PromptSurveyId** | Pointer to **NullableInt32** | Display name of the tenant | [optional] 
**QuestionsInducedSurveyId** | Pointer to **NullableInt32** | Display name of the tenant | [optional] 
**Client** | [**TargetAudience**](TargetAudience.md) |  | 
**SurveyResponse** | Pointer to **interface{}** | The survey response | [optional] 
**Timestamp** | **NullableTime** | The timestamp of the chat. | [readonly] 
**CreatedAt** | **NullableString** |  | [readonly] 

## Methods

### NewResponse

`func NewResponse(responseId int32, client TargetAudience, timestamp NullableTime, createdAt NullableString, ) *Response`

NewResponse instantiates a new Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseWithDefaults

`func NewResponseWithDefaults() *Response`

NewResponseWithDefaults instantiates a new Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseId

`func (o *Response) GetResponseId() int32`

GetResponseId returns the ResponseId field if non-nil, zero value otherwise.

### GetResponseIdOk

`func (o *Response) GetResponseIdOk() (*int32, bool)`

GetResponseIdOk returns a tuple with the ResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseId

`func (o *Response) SetResponseId(v int32)`

SetResponseId sets ResponseId field to given value.


### GetSurveyId

`func (o *Response) GetSurveyId() int32`

GetSurveyId returns the SurveyId field if non-nil, zero value otherwise.

### GetSurveyIdOk

`func (o *Response) GetSurveyIdOk() (*int32, bool)`

GetSurveyIdOk returns a tuple with the SurveyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveyId

`func (o *Response) SetSurveyId(v int32)`

SetSurveyId sets SurveyId field to given value.

### HasSurveyId

`func (o *Response) HasSurveyId() bool`

HasSurveyId returns a boolean if a field has been set.

### SetSurveyIdNil

`func (o *Response) SetSurveyIdNil(b bool)`

 SetSurveyIdNil sets the value for SurveyId to be an explicit nil

### UnsetSurveyId
`func (o *Response) UnsetSurveyId()`

UnsetSurveyId ensures that no value is present for SurveyId, not even an explicit nil
### GetPromptSurveyId

`func (o *Response) GetPromptSurveyId() int32`

GetPromptSurveyId returns the PromptSurveyId field if non-nil, zero value otherwise.

### GetPromptSurveyIdOk

`func (o *Response) GetPromptSurveyIdOk() (*int32, bool)`

GetPromptSurveyIdOk returns a tuple with the PromptSurveyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptSurveyId

`func (o *Response) SetPromptSurveyId(v int32)`

SetPromptSurveyId sets PromptSurveyId field to given value.

### HasPromptSurveyId

`func (o *Response) HasPromptSurveyId() bool`

HasPromptSurveyId returns a boolean if a field has been set.

### SetPromptSurveyIdNil

`func (o *Response) SetPromptSurveyIdNil(b bool)`

 SetPromptSurveyIdNil sets the value for PromptSurveyId to be an explicit nil

### UnsetPromptSurveyId
`func (o *Response) UnsetPromptSurveyId()`

UnsetPromptSurveyId ensures that no value is present for PromptSurveyId, not even an explicit nil
### GetQuestionsInducedSurveyId

`func (o *Response) GetQuestionsInducedSurveyId() int32`

GetQuestionsInducedSurveyId returns the QuestionsInducedSurveyId field if non-nil, zero value otherwise.

### GetQuestionsInducedSurveyIdOk

`func (o *Response) GetQuestionsInducedSurveyIdOk() (*int32, bool)`

GetQuestionsInducedSurveyIdOk returns a tuple with the QuestionsInducedSurveyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionsInducedSurveyId

`func (o *Response) SetQuestionsInducedSurveyId(v int32)`

SetQuestionsInducedSurveyId sets QuestionsInducedSurveyId field to given value.

### HasQuestionsInducedSurveyId

`func (o *Response) HasQuestionsInducedSurveyId() bool`

HasQuestionsInducedSurveyId returns a boolean if a field has been set.

### SetQuestionsInducedSurveyIdNil

`func (o *Response) SetQuestionsInducedSurveyIdNil(b bool)`

 SetQuestionsInducedSurveyIdNil sets the value for QuestionsInducedSurveyId to be an explicit nil

### UnsetQuestionsInducedSurveyId
`func (o *Response) UnsetQuestionsInducedSurveyId()`

UnsetQuestionsInducedSurveyId ensures that no value is present for QuestionsInducedSurveyId, not even an explicit nil
### GetClient

`func (o *Response) GetClient() TargetAudience`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *Response) GetClientOk() (*TargetAudience, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *Response) SetClient(v TargetAudience)`

SetClient sets Client field to given value.


### GetSurveyResponse

`func (o *Response) GetSurveyResponse() interface{}`

GetSurveyResponse returns the SurveyResponse field if non-nil, zero value otherwise.

### GetSurveyResponseOk

`func (o *Response) GetSurveyResponseOk() (*interface{}, bool)`

GetSurveyResponseOk returns a tuple with the SurveyResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveyResponse

`func (o *Response) SetSurveyResponse(v interface{})`

SetSurveyResponse sets SurveyResponse field to given value.

### HasSurveyResponse

`func (o *Response) HasSurveyResponse() bool`

HasSurveyResponse returns a boolean if a field has been set.

### SetSurveyResponseNil

`func (o *Response) SetSurveyResponseNil(b bool)`

 SetSurveyResponseNil sets the value for SurveyResponse to be an explicit nil

### UnsetSurveyResponse
`func (o *Response) UnsetSurveyResponse()`

UnsetSurveyResponse ensures that no value is present for SurveyResponse, not even an explicit nil
### GetTimestamp

`func (o *Response) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Response) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Response) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### SetTimestampNil

`func (o *Response) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *Response) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetCreatedAt

`func (o *Response) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Response) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Response) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *Response) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Response) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


