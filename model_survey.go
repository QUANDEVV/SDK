/*
proxima-core-engine

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// Survey struct for Survey
type Survey struct {
	// The lilling details id ID UUID
	SurveyId int32 `json:"survey_id"`
	// Display name of the tenant
	TenantId int32 `json:"tenant_id"`
	// The survey topic
	SurveyTopic string `json:"survey_topic"`
	// The survey description
	SurveyDescription string `json:"survey_description"`
	// The survey context
	SurveyContext string `json:"survey_context"`
	// The survey questions
	SurveyQuestions interface{} `json:"survey_questions,omitempty"`
	// The target audience/who to share with
	TargetAudience *[]int32 `json:"target_audience,omitempty"`
	// The survey type  * `open_ended` - open_ended * `close_ended` - close_ended
	SurveyType OneOfSurveyTypeEnumBlankEnumNullEnum `json:"survey_type,omitempty"`
	// The start day of the survey
	StartDay NullableString `json:"start_day"`
	// The end day of the survey
	EndDay NullableString `json:"end_day,omitempty"`
	SurveyStatus *bool `json:"survey_status,omitempty"`
	TargetAudienceCount int32 `json:"target_audience_count"`
	ResponseCount int32 `json:"response_count"`
	PercentageDifference float64 `json:"percentage_difference"`
	QuestionCount int32 `json:"question_count"`
}

// NewSurvey instantiates a new Survey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSurvey(surveyId int32, tenantId int32, surveyTopic string, surveyDescription string, surveyContext string, startDay NullableString, targetAudienceCount int32, responseCount int32, percentageDifference float64, questionCount int32) *Survey {
	this := Survey{}
	this.SurveyId = surveyId
	this.TenantId = tenantId
	this.SurveyTopic = surveyTopic
	this.SurveyDescription = surveyDescription
	this.SurveyContext = surveyContext
	this.StartDay = startDay
	this.TargetAudienceCount = targetAudienceCount
	this.ResponseCount = responseCount
	this.PercentageDifference = percentageDifference
	this.QuestionCount = questionCount
	return &this
}

// NewSurveyWithDefaults instantiates a new Survey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSurveyWithDefaults() *Survey {
	this := Survey{}
	return &this
}

// GetSurveyId returns the SurveyId field value
func (o *Survey) GetSurveyId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SurveyId
}

// GetSurveyIdOk returns a tuple with the SurveyId field value
// and a boolean to check if the value has been set.
func (o *Survey) GetSurveyIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SurveyId, true
}

// SetSurveyId sets field value
func (o *Survey) SetSurveyId(v int32) {
	o.SurveyId = v
}

// GetTenantId returns the TenantId field value
func (o *Survey) GetTenantId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *Survey) GetTenantIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *Survey) SetTenantId(v int32) {
	o.TenantId = v
}

// GetSurveyTopic returns the SurveyTopic field value
func (o *Survey) GetSurveyTopic() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SurveyTopic
}

// GetSurveyTopicOk returns a tuple with the SurveyTopic field value
// and a boolean to check if the value has been set.
func (o *Survey) GetSurveyTopicOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SurveyTopic, true
}

// SetSurveyTopic sets field value
func (o *Survey) SetSurveyTopic(v string) {
	o.SurveyTopic = v
}

// GetSurveyDescription returns the SurveyDescription field value
func (o *Survey) GetSurveyDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SurveyDescription
}

// GetSurveyDescriptionOk returns a tuple with the SurveyDescription field value
// and a boolean to check if the value has been set.
func (o *Survey) GetSurveyDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SurveyDescription, true
}

// SetSurveyDescription sets field value
func (o *Survey) SetSurveyDescription(v string) {
	o.SurveyDescription = v
}

// GetSurveyContext returns the SurveyContext field value
func (o *Survey) GetSurveyContext() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SurveyContext
}

// GetSurveyContextOk returns a tuple with the SurveyContext field value
// and a boolean to check if the value has been set.
func (o *Survey) GetSurveyContextOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SurveyContext, true
}

// SetSurveyContext sets field value
func (o *Survey) SetSurveyContext(v string) {
	o.SurveyContext = v
}

// GetSurveyQuestions returns the SurveyQuestions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Survey) GetSurveyQuestions() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.SurveyQuestions
}

// GetSurveyQuestionsOk returns a tuple with the SurveyQuestions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Survey) GetSurveyQuestionsOk() (*interface{}, bool) {
	if o == nil || o.SurveyQuestions == nil {
		return nil, false
	}
	return &o.SurveyQuestions, true
}

// HasSurveyQuestions returns a boolean if a field has been set.
func (o *Survey) HasSurveyQuestions() bool {
	if o != nil && o.SurveyQuestions != nil {
		return true
	}

	return false
}

// SetSurveyQuestions gets a reference to the given interface{} and assigns it to the SurveyQuestions field.
func (o *Survey) SetSurveyQuestions(v interface{}) {
	o.SurveyQuestions = v
}

// GetTargetAudience returns the TargetAudience field value if set, zero value otherwise.
func (o *Survey) GetTargetAudience() []int32 {
	if o == nil || o.TargetAudience == nil {
		var ret []int32
		return ret
	}
	return *o.TargetAudience
}

// GetTargetAudienceOk returns a tuple with the TargetAudience field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Survey) GetTargetAudienceOk() (*[]int32, bool) {
	if o == nil || o.TargetAudience == nil {
		return nil, false
	}
	return o.TargetAudience, true
}

// HasTargetAudience returns a boolean if a field has been set.
func (o *Survey) HasTargetAudience() bool {
	if o != nil && o.TargetAudience != nil {
		return true
	}

	return false
}

// SetTargetAudience gets a reference to the given []int32 and assigns it to the TargetAudience field.
func (o *Survey) SetTargetAudience(v []int32) {
	o.TargetAudience = &v
}

// GetSurveyType returns the SurveyType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Survey) GetSurveyType() OneOfSurveyTypeEnumBlankEnumNullEnum {
	if o == nil  {
		var ret OneOfSurveyTypeEnumBlankEnumNullEnum
		return ret
	}
	return o.SurveyType
}

// GetSurveyTypeOk returns a tuple with the SurveyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Survey) GetSurveyTypeOk() (*OneOfSurveyTypeEnumBlankEnumNullEnum, bool) {
	if o == nil || o.SurveyType == nil {
		return nil, false
	}
	return &o.SurveyType, true
}

// HasSurveyType returns a boolean if a field has been set.
func (o *Survey) HasSurveyType() bool {
	if o != nil && o.SurveyType != nil {
		return true
	}

	return false
}

// SetSurveyType gets a reference to the given OneOfSurveyTypeEnumBlankEnumNullEnum and assigns it to the SurveyType field.
func (o *Survey) SetSurveyType(v OneOfSurveyTypeEnumBlankEnumNullEnum) {
	o.SurveyType = v
}

// GetStartDay returns the StartDay field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Survey) GetStartDay() string {
	if o == nil || o.StartDay.Get() == nil {
		var ret string
		return ret
	}

	return *o.StartDay.Get()
}

// GetStartDayOk returns a tuple with the StartDay field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Survey) GetStartDayOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartDay.Get(), o.StartDay.IsSet()
}

// SetStartDay sets field value
func (o *Survey) SetStartDay(v string) {
	o.StartDay.Set(&v)
}

// GetEndDay returns the EndDay field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Survey) GetEndDay() string {
	if o == nil || o.EndDay.Get() == nil {
		var ret string
		return ret
	}
	return *o.EndDay.Get()
}

// GetEndDayOk returns a tuple with the EndDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Survey) GetEndDayOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndDay.Get(), o.EndDay.IsSet()
}

// HasEndDay returns a boolean if a field has been set.
func (o *Survey) HasEndDay() bool {
	if o != nil && o.EndDay.IsSet() {
		return true
	}

	return false
}

// SetEndDay gets a reference to the given NullableString and assigns it to the EndDay field.
func (o *Survey) SetEndDay(v string) {
	o.EndDay.Set(&v)
}
// SetEndDayNil sets the value for EndDay to be an explicit nil
func (o *Survey) SetEndDayNil() {
	o.EndDay.Set(nil)
}

// UnsetEndDay ensures that no value is present for EndDay, not even an explicit nil
func (o *Survey) UnsetEndDay() {
	o.EndDay.Unset()
}

// GetSurveyStatus returns the SurveyStatus field value if set, zero value otherwise.
func (o *Survey) GetSurveyStatus() bool {
	if o == nil || o.SurveyStatus == nil {
		var ret bool
		return ret
	}
	return *o.SurveyStatus
}

// GetSurveyStatusOk returns a tuple with the SurveyStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Survey) GetSurveyStatusOk() (*bool, bool) {
	if o == nil || o.SurveyStatus == nil {
		return nil, false
	}
	return o.SurveyStatus, true
}

// HasSurveyStatus returns a boolean if a field has been set.
func (o *Survey) HasSurveyStatus() bool {
	if o != nil && o.SurveyStatus != nil {
		return true
	}

	return false
}

// SetSurveyStatus gets a reference to the given bool and assigns it to the SurveyStatus field.
func (o *Survey) SetSurveyStatus(v bool) {
	o.SurveyStatus = &v
}

// GetTargetAudienceCount returns the TargetAudienceCount field value
func (o *Survey) GetTargetAudienceCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TargetAudienceCount
}

// GetTargetAudienceCountOk returns a tuple with the TargetAudienceCount field value
// and a boolean to check if the value has been set.
func (o *Survey) GetTargetAudienceCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TargetAudienceCount, true
}

// SetTargetAudienceCount sets field value
func (o *Survey) SetTargetAudienceCount(v int32) {
	o.TargetAudienceCount = v
}

// GetResponseCount returns the ResponseCount field value
func (o *Survey) GetResponseCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ResponseCount
}

// GetResponseCountOk returns a tuple with the ResponseCount field value
// and a boolean to check if the value has been set.
func (o *Survey) GetResponseCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ResponseCount, true
}

// SetResponseCount sets field value
func (o *Survey) SetResponseCount(v int32) {
	o.ResponseCount = v
}

// GetPercentageDifference returns the PercentageDifference field value
func (o *Survey) GetPercentageDifference() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.PercentageDifference
}

// GetPercentageDifferenceOk returns a tuple with the PercentageDifference field value
// and a boolean to check if the value has been set.
func (o *Survey) GetPercentageDifferenceOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PercentageDifference, true
}

// SetPercentageDifference sets field value
func (o *Survey) SetPercentageDifference(v float64) {
	o.PercentageDifference = v
}

// GetQuestionCount returns the QuestionCount field value
func (o *Survey) GetQuestionCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.QuestionCount
}

// GetQuestionCountOk returns a tuple with the QuestionCount field value
// and a boolean to check if the value has been set.
func (o *Survey) GetQuestionCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.QuestionCount, true
}

// SetQuestionCount sets field value
func (o *Survey) SetQuestionCount(v int32) {
	o.QuestionCount = v
}

func (o Survey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["survey_id"] = o.SurveyId
	}
	if true {
		toSerialize["tenant_id"] = o.TenantId
	}
	if true {
		toSerialize["survey_topic"] = o.SurveyTopic
	}
	if true {
		toSerialize["survey_description"] = o.SurveyDescription
	}
	if true {
		toSerialize["survey_context"] = o.SurveyContext
	}
	if o.SurveyQuestions != nil {
		toSerialize["survey_questions"] = o.SurveyQuestions
	}
	if o.TargetAudience != nil {
		toSerialize["target_audience"] = o.TargetAudience
	}
	if o.SurveyType != nil {
		toSerialize["survey_type"] = o.SurveyType
	}
	if true {
		toSerialize["start_day"] = o.StartDay.Get()
	}
	if o.EndDay.IsSet() {
		toSerialize["end_day"] = o.EndDay.Get()
	}
	if o.SurveyStatus != nil {
		toSerialize["survey_status"] = o.SurveyStatus
	}
	if true {
		toSerialize["target_audience_count"] = o.TargetAudienceCount
	}
	if true {
		toSerialize["response_count"] = o.ResponseCount
	}
	if true {
		toSerialize["percentage_difference"] = o.PercentageDifference
	}
	if true {
		toSerialize["question_count"] = o.QuestionCount
	}
	return json.Marshal(toSerialize)
}

type NullableSurvey struct {
	value *Survey
	isSet bool
}

func (v NullableSurvey) Get() *Survey {
	return v.value
}

func (v *NullableSurvey) Set(val *Survey) {
	v.value = val
	v.isSet = true
}

func (v NullableSurvey) IsSet() bool {
	return v.isSet
}

func (v *NullableSurvey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSurvey(val *Survey) *NullableSurvey {
	return &NullableSurvey{value: val, isSet: true}
}

func (v NullableSurvey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSurvey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


