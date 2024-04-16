# Thread

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThreadId** | **int32** | The thread ID UUID for an instance of a thread. | [readonly] 
**Issue** | [**Nested**](Nested.md) |  | [readonly] 

## Methods

### NewThread

`func NewThread(threadId int32, issue Nested, ) *Thread`

NewThread instantiates a new Thread object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreadWithDefaults

`func NewThreadWithDefaults() *Thread`

NewThreadWithDefaults instantiates a new Thread object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreadId

`func (o *Thread) GetThreadId() int32`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *Thread) GetThreadIdOk() (*int32, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *Thread) SetThreadId(v int32)`

SetThreadId sets ThreadId field to given value.


### GetIssue

`func (o *Thread) GetIssue() Nested`

GetIssue returns the Issue field if non-nil, zero value otherwise.

### GetIssueOk

`func (o *Thread) GetIssueOk() (*Nested, bool)`

GetIssueOk returns a tuple with the Issue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssue

`func (o *Thread) SetIssue(v Nested)`

SetIssue sets Issue field to given value.


### SetIssueNil

`func (o *Thread) SetIssueNil(b bool)`

 SetIssueNil sets the value for Issue to be an explicit nil

### UnsetIssue
`func (o *Thread) UnsetIssue()`

UnsetIssue ensures that no value is present for Issue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


