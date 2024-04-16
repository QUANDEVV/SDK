# CallMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallMetricsId** | **int32** | The call metrics id. | [readonly] 
**Call** | [**Call**](Call.md) |  | [readonly] 
**Transcripts** | **interface{}** |  | 
**CallSatisfaction** | Pointer to **bool** |  | [optional] 
**IssueResolution** | Pointer to **bool** |  | [optional] 

## Methods

### NewCallMetrics

`func NewCallMetrics(callMetricsId int32, call Call, transcripts interface{}, ) *CallMetrics`

NewCallMetrics instantiates a new CallMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallMetricsWithDefaults

`func NewCallMetricsWithDefaults() *CallMetrics`

NewCallMetricsWithDefaults instantiates a new CallMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallMetricsId

`func (o *CallMetrics) GetCallMetricsId() int32`

GetCallMetricsId returns the CallMetricsId field if non-nil, zero value otherwise.

### GetCallMetricsIdOk

`func (o *CallMetrics) GetCallMetricsIdOk() (*int32, bool)`

GetCallMetricsIdOk returns a tuple with the CallMetricsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallMetricsId

`func (o *CallMetrics) SetCallMetricsId(v int32)`

SetCallMetricsId sets CallMetricsId field to given value.


### GetCall

`func (o *CallMetrics) GetCall() Call`

GetCall returns the Call field if non-nil, zero value otherwise.

### GetCallOk

`func (o *CallMetrics) GetCallOk() (*Call, bool)`

GetCallOk returns a tuple with the Call field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCall

`func (o *CallMetrics) SetCall(v Call)`

SetCall sets Call field to given value.


### SetCallNil

`func (o *CallMetrics) SetCallNil(b bool)`

 SetCallNil sets the value for Call to be an explicit nil

### UnsetCall
`func (o *CallMetrics) UnsetCall()`

UnsetCall ensures that no value is present for Call, not even an explicit nil
### GetTranscripts

`func (o *CallMetrics) GetTranscripts() interface{}`

GetTranscripts returns the Transcripts field if non-nil, zero value otherwise.

### GetTranscriptsOk

`func (o *CallMetrics) GetTranscriptsOk() (*interface{}, bool)`

GetTranscriptsOk returns a tuple with the Transcripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscripts

`func (o *CallMetrics) SetTranscripts(v interface{})`

SetTranscripts sets Transcripts field to given value.


### SetTranscriptsNil

`func (o *CallMetrics) SetTranscriptsNil(b bool)`

 SetTranscriptsNil sets the value for Transcripts to be an explicit nil

### UnsetTranscripts
`func (o *CallMetrics) UnsetTranscripts()`

UnsetTranscripts ensures that no value is present for Transcripts, not even an explicit nil
### GetCallSatisfaction

`func (o *CallMetrics) GetCallSatisfaction() bool`

GetCallSatisfaction returns the CallSatisfaction field if non-nil, zero value otherwise.

### GetCallSatisfactionOk

`func (o *CallMetrics) GetCallSatisfactionOk() (*bool, bool)`

GetCallSatisfactionOk returns a tuple with the CallSatisfaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSatisfaction

`func (o *CallMetrics) SetCallSatisfaction(v bool)`

SetCallSatisfaction sets CallSatisfaction field to given value.

### HasCallSatisfaction

`func (o *CallMetrics) HasCallSatisfaction() bool`

HasCallSatisfaction returns a boolean if a field has been set.

### GetIssueResolution

`func (o *CallMetrics) GetIssueResolution() bool`

GetIssueResolution returns the IssueResolution field if non-nil, zero value otherwise.

### GetIssueResolutionOk

`func (o *CallMetrics) GetIssueResolutionOk() (*bool, bool)`

GetIssueResolutionOk returns a tuple with the IssueResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueResolution

`func (o *CallMetrics) SetIssueResolution(v bool)`

SetIssueResolution sets IssueResolution field to given value.

### HasIssueResolution

`func (o *CallMetrics) HasIssueResolution() bool`

HasIssueResolution returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


