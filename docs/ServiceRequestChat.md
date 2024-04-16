# ServiceRequestChat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceRequestChatId** | **int32** |  | [readonly] 
**ServiceRequest** | [**ServiceRequest**](ServiceRequest.md) |  | [readonly] 
**CreatedAt** | **time.Time** | The date and time when the chat was created. | [readonly] 
**UpdatedAt** | **time.Time** | The date and time when the chat was last updated. | [readonly] 

## Methods

### NewServiceRequestChat

`func NewServiceRequestChat(serviceRequestChatId int32, serviceRequest ServiceRequest, createdAt time.Time, updatedAt time.Time, ) *ServiceRequestChat`

NewServiceRequestChat instantiates a new ServiceRequestChat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceRequestChatWithDefaults

`func NewServiceRequestChatWithDefaults() *ServiceRequestChat`

NewServiceRequestChatWithDefaults instantiates a new ServiceRequestChat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceRequestChatId

`func (o *ServiceRequestChat) GetServiceRequestChatId() int32`

GetServiceRequestChatId returns the ServiceRequestChatId field if non-nil, zero value otherwise.

### GetServiceRequestChatIdOk

`func (o *ServiceRequestChat) GetServiceRequestChatIdOk() (*int32, bool)`

GetServiceRequestChatIdOk returns a tuple with the ServiceRequestChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRequestChatId

`func (o *ServiceRequestChat) SetServiceRequestChatId(v int32)`

SetServiceRequestChatId sets ServiceRequestChatId field to given value.


### GetServiceRequest

`func (o *ServiceRequestChat) GetServiceRequest() ServiceRequest`

GetServiceRequest returns the ServiceRequest field if non-nil, zero value otherwise.

### GetServiceRequestOk

`func (o *ServiceRequestChat) GetServiceRequestOk() (*ServiceRequest, bool)`

GetServiceRequestOk returns a tuple with the ServiceRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRequest

`func (o *ServiceRequestChat) SetServiceRequest(v ServiceRequest)`

SetServiceRequest sets ServiceRequest field to given value.


### SetServiceRequestNil

`func (o *ServiceRequestChat) SetServiceRequestNil(b bool)`

 SetServiceRequestNil sets the value for ServiceRequest to be an explicit nil

### UnsetServiceRequest
`func (o *ServiceRequestChat) UnsetServiceRequest()`

UnsetServiceRequest ensures that no value is present for ServiceRequest, not even an explicit nil
### GetCreatedAt

`func (o *ServiceRequestChat) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServiceRequestChat) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServiceRequestChat) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ServiceRequestChat) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ServiceRequestChat) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ServiceRequestChat) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


