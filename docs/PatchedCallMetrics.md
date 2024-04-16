# PatchedCallMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallMetricsId** | Pointer to **int32** | The call metrics id. | [optional] [readonly] 
**Call** | Pointer to [**Call**](Call.md) |  | [optional] [readonly] 
**Transcripts** | Pointer to **interface{}** |  | [optional] 
**CallSatisfaction** | Pointer to **bool** |  | [optional] 
**IssueResolution** | Pointer to **bool** |  | [optional] 

## Methods

### NewPatchedCallMetrics

`func NewPatchedCallMetrics() *PatchedCallMetrics`

NewPatchedCallMetrics instantiates a new PatchedCallMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedCallMetricsWithDefaults

`func NewPatchedCallMetricsWithDefaults() *PatchedCallMetrics`

NewPatchedCallMetricsWithDefaults instantiates a new PatchedCallMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallMetricsId

`func (o *PatchedCallMetrics) GetCallMetricsId() int32`

GetCallMetricsId returns the CallMetricsId field if non-nil, zero value otherwise.

### GetCallMetricsIdOk

`func (o *PatchedCallMetrics) GetCallMetricsIdOk() (*int32, bool)`

GetCallMetricsIdOk returns a tuple with the CallMetricsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallMetricsId

`func (o *PatchedCallMetrics) SetCallMetricsId(v int32)`

SetCallMetricsId sets CallMetricsId field to given value.

### HasCallMetricsId

`func (o *PatchedCallMetrics) HasCallMetricsId() bool`

HasCallMetricsId returns a boolean if a field has been set.

### GetCall

`func (o *PatchedCallMetrics) GetCall() Call`

GetCall returns the Call field if non-nil, zero value otherwise.

### GetCallOk

`func (o *PatchedCallMetrics) GetCallOk() (*Call, bool)`

GetCallOk returns a tuple with the Call field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCall

`func (o *PatchedCallMetrics) SetCall(v Call)`

SetCall sets Call field to given value.

### HasCall

`func (o *PatchedCallMetrics) HasCall() bool`

HasCall returns a boolean if a field has been set.

### SetCallNil

`func (o *PatchedCallMetrics) SetCallNil(b bool)`

 SetCallNil sets the value for Call to be an explicit nil

### UnsetCall
`func (o *PatchedCallMetrics) UnsetCall()`

UnsetCall ensures that no value is present for Call, not even an explicit nil
### GetTranscripts

`func (o *PatchedCallMetrics) GetTranscripts() interface{}`

GetTranscripts returns the Transcripts field if non-nil, zero value otherwise.

### GetTranscriptsOk

`func (o *PatchedCallMetrics) GetTranscriptsOk() (*interface{}, bool)`

GetTranscriptsOk returns a tuple with the Transcripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscripts

`func (o *PatchedCallMetrics) SetTranscripts(v interface{})`

SetTranscripts sets Transcripts field to given value.

### HasTranscripts

`func (o *PatchedCallMetrics) HasTranscripts() bool`

HasTranscripts returns a boolean if a field has been set.

### SetTranscriptsNil

`func (o *PatchedCallMetrics) SetTranscriptsNil(b bool)`

 SetTranscriptsNil sets the value for Transcripts to be an explicit nil

### UnsetTranscripts
`func (o *PatchedCallMetrics) UnsetTranscripts()`

UnsetTranscripts ensures that no value is present for Transcripts, not even an explicit nil
### GetCallSatisfaction

`func (o *PatchedCallMetrics) GetCallSatisfaction() bool`

GetCallSatisfaction returns the CallSatisfaction field if non-nil, zero value otherwise.

### GetCallSatisfactionOk

`func (o *PatchedCallMetrics) GetCallSatisfactionOk() (*bool, bool)`

GetCallSatisfactionOk returns a tuple with the CallSatisfaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSatisfaction

`func (o *PatchedCallMetrics) SetCallSatisfaction(v bool)`

SetCallSatisfaction sets CallSatisfaction field to given value.

### HasCallSatisfaction

`func (o *PatchedCallMetrics) HasCallSatisfaction() bool`

HasCallSatisfaction returns a boolean if a field has been set.

### GetIssueResolution

`func (o *PatchedCallMetrics) GetIssueResolution() bool`

GetIssueResolution returns the IssueResolution field if non-nil, zero value otherwise.

### GetIssueResolutionOk

`func (o *PatchedCallMetrics) GetIssueResolutionOk() (*bool, bool)`

GetIssueResolutionOk returns a tuple with the IssueResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueResolution

`func (o *PatchedCallMetrics) SetIssueResolution(v bool)`

SetIssueResolution sets IssueResolution field to given value.

### HasIssueResolution

`func (o *PatchedCallMetrics) HasIssueResolution() bool`

HasIssueResolution returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


