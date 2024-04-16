# SurveyReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SurveyReportId** | **int32** | The survey report id | [readonly] 
**SurveyId** | **int32** | Display name of the survey | 
**Conclusion** | **string** | The survey subgroup name | 
**SurveySuccess** | Pointer to **bool** | The survey success | [optional] 
**SurveyReporter** | Pointer to **NullableInt32** | Employee who made the survey review | [optional] 

## Methods

### NewSurveyReport

`func NewSurveyReport(surveyReportId int32, surveyId int32, conclusion string, ) *SurveyReport`

NewSurveyReport instantiates a new SurveyReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSurveyReportWithDefaults

`func NewSurveyReportWithDefaults() *SurveyReport`

NewSurveyReportWithDefaults instantiates a new SurveyReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSurveyReportId

`func (o *SurveyReport) GetSurveyReportId() int32`

GetSurveyReportId returns the SurveyReportId field if non-nil, zero value otherwise.

### GetSurveyReportIdOk

`func (o *SurveyReport) GetSurveyReportIdOk() (*int32, bool)`

GetSurveyReportIdOk returns a tuple with the SurveyReportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveyReportId

`func (o *SurveyReport) SetSurveyReportId(v int32)`

SetSurveyReportId sets SurveyReportId field to given value.


### GetSurveyId

`func (o *SurveyReport) GetSurveyId() int32`

GetSurveyId returns the SurveyId field if non-nil, zero value otherwise.

### GetSurveyIdOk

`func (o *SurveyReport) GetSurveyIdOk() (*int32, bool)`

GetSurveyIdOk returns a tuple with the SurveyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveyId

`func (o *SurveyReport) SetSurveyId(v int32)`

SetSurveyId sets SurveyId field to given value.


### GetConclusion

`func (o *SurveyReport) GetConclusion() string`

GetConclusion returns the Conclusion field if non-nil, zero value otherwise.

### GetConclusionOk

`func (o *SurveyReport) GetConclusionOk() (*string, bool)`

GetConclusionOk returns a tuple with the Conclusion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConclusion

`func (o *SurveyReport) SetConclusion(v string)`

SetConclusion sets Conclusion field to given value.


### GetSurveySuccess

`func (o *SurveyReport) GetSurveySuccess() bool`

GetSurveySuccess returns the SurveySuccess field if non-nil, zero value otherwise.

### GetSurveySuccessOk

`func (o *SurveyReport) GetSurveySuccessOk() (*bool, bool)`

GetSurveySuccessOk returns a tuple with the SurveySuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveySuccess

`func (o *SurveyReport) SetSurveySuccess(v bool)`

SetSurveySuccess sets SurveySuccess field to given value.

### HasSurveySuccess

`func (o *SurveyReport) HasSurveySuccess() bool`

HasSurveySuccess returns a boolean if a field has been set.

### GetSurveyReporter

`func (o *SurveyReport) GetSurveyReporter() int32`

GetSurveyReporter returns the SurveyReporter field if non-nil, zero value otherwise.

### GetSurveyReporterOk

`func (o *SurveyReport) GetSurveyReporterOk() (*int32, bool)`

GetSurveyReporterOk returns a tuple with the SurveyReporter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveyReporter

`func (o *SurveyReport) SetSurveyReporter(v int32)`

SetSurveyReporter sets SurveyReporter field to given value.

### HasSurveyReporter

`func (o *SurveyReport) HasSurveyReporter() bool`

HasSurveyReporter returns a boolean if a field has been set.

### SetSurveyReporterNil

`func (o *SurveyReport) SetSurveyReporterNil(b bool)`

 SetSurveyReporterNil sets the value for SurveyReporter to be an explicit nil

### UnsetSurveyReporter
`func (o *SurveyReport) UnsetSurveyReporter()`

UnsetSurveyReporter ensures that no value is present for SurveyReporter, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


