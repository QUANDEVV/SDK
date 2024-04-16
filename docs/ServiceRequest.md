# ServiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **NullableString** |  | [readonly] 
**ServiceRequestsId** | **int32** |  | [readonly] 
**Client** | [**ClientServicesInfo**](ClientServicesInfo.md) |  | 
**UnauthenticatedUser** | [**AnonymousUser**](AnonymousUser.md) |  | 
**Service** | [**Service**](Service.md) |  | 
**Status** | Pointer to [**ServiceRequestStatusEnum**](ServiceRequestStatusEnum.md) | Status of the request.  * &#x60;pending&#x60; - Pending * &#x60;completed&#x60; - Completed * &#x60;in_progress&#x60; - In Progress | [optional] 
**RequestDetails** | **interface{}** | Specific request details in JSON format. | 
**AiInfo** | **interface{}** | A json dump of the ai content generated | 
**ServiceActionPlanResponse** | **interface{}** |  | 

## Methods

### NewServiceRequest

`func NewServiceRequest(createdAt NullableString, serviceRequestsId int32, client ClientServicesInfo, unauthenticatedUser AnonymousUser, service Service, requestDetails interface{}, aiInfo interface{}, serviceActionPlanResponse interface{}, ) *ServiceRequest`

NewServiceRequest instantiates a new ServiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceRequestWithDefaults

`func NewServiceRequestWithDefaults() *ServiceRequest`

NewServiceRequestWithDefaults instantiates a new ServiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ServiceRequest) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServiceRequest) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServiceRequest) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *ServiceRequest) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *ServiceRequest) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetServiceRequestsId

`func (o *ServiceRequest) GetServiceRequestsId() int32`

GetServiceRequestsId returns the ServiceRequestsId field if non-nil, zero value otherwise.

### GetServiceRequestsIdOk

`func (o *ServiceRequest) GetServiceRequestsIdOk() (*int32, bool)`

GetServiceRequestsIdOk returns a tuple with the ServiceRequestsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRequestsId

`func (o *ServiceRequest) SetServiceRequestsId(v int32)`

SetServiceRequestsId sets ServiceRequestsId field to given value.


### GetClient

`func (o *ServiceRequest) GetClient() ClientServicesInfo`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *ServiceRequest) GetClientOk() (*ClientServicesInfo, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *ServiceRequest) SetClient(v ClientServicesInfo)`

SetClient sets Client field to given value.


### GetUnauthenticatedUser

`func (o *ServiceRequest) GetUnauthenticatedUser() AnonymousUser`

GetUnauthenticatedUser returns the UnauthenticatedUser field if non-nil, zero value otherwise.

### GetUnauthenticatedUserOk

`func (o *ServiceRequest) GetUnauthenticatedUserOk() (*AnonymousUser, bool)`

GetUnauthenticatedUserOk returns a tuple with the UnauthenticatedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnauthenticatedUser

`func (o *ServiceRequest) SetUnauthenticatedUser(v AnonymousUser)`

SetUnauthenticatedUser sets UnauthenticatedUser field to given value.


### GetService

`func (o *ServiceRequest) GetService() Service`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ServiceRequest) GetServiceOk() (*Service, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ServiceRequest) SetService(v Service)`

SetService sets Service field to given value.


### GetStatus

`func (o *ServiceRequest) GetStatus() ServiceRequestStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServiceRequest) GetStatusOk() (*ServiceRequestStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServiceRequest) SetStatus(v ServiceRequestStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ServiceRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ServiceRequest) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ServiceRequest) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetRequestDetails

`func (o *ServiceRequest) GetRequestDetails() interface{}`

GetRequestDetails returns the RequestDetails field if non-nil, zero value otherwise.

### GetRequestDetailsOk

`func (o *ServiceRequest) GetRequestDetailsOk() (*interface{}, bool)`

GetRequestDetailsOk returns a tuple with the RequestDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestDetails

`func (o *ServiceRequest) SetRequestDetails(v interface{})`

SetRequestDetails sets RequestDetails field to given value.


### SetRequestDetailsNil

`func (o *ServiceRequest) SetRequestDetailsNil(b bool)`

 SetRequestDetailsNil sets the value for RequestDetails to be an explicit nil

### UnsetRequestDetails
`func (o *ServiceRequest) UnsetRequestDetails()`

UnsetRequestDetails ensures that no value is present for RequestDetails, not even an explicit nil
### GetAiInfo

`func (o *ServiceRequest) GetAiInfo() interface{}`

GetAiInfo returns the AiInfo field if non-nil, zero value otherwise.

### GetAiInfoOk

`func (o *ServiceRequest) GetAiInfoOk() (*interface{}, bool)`

GetAiInfoOk returns a tuple with the AiInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiInfo

`func (o *ServiceRequest) SetAiInfo(v interface{})`

SetAiInfo sets AiInfo field to given value.


### SetAiInfoNil

`func (o *ServiceRequest) SetAiInfoNil(b bool)`

 SetAiInfoNil sets the value for AiInfo to be an explicit nil

### UnsetAiInfo
`func (o *ServiceRequest) UnsetAiInfo()`

UnsetAiInfo ensures that no value is present for AiInfo, not even an explicit nil
### GetServiceActionPlanResponse

`func (o *ServiceRequest) GetServiceActionPlanResponse() interface{}`

GetServiceActionPlanResponse returns the ServiceActionPlanResponse field if non-nil, zero value otherwise.

### GetServiceActionPlanResponseOk

`func (o *ServiceRequest) GetServiceActionPlanResponseOk() (*interface{}, bool)`

GetServiceActionPlanResponseOk returns a tuple with the ServiceActionPlanResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceActionPlanResponse

`func (o *ServiceRequest) SetServiceActionPlanResponse(v interface{})`

SetServiceActionPlanResponse sets ServiceActionPlanResponse field to given value.


### SetServiceActionPlanResponseNil

`func (o *ServiceRequest) SetServiceActionPlanResponseNil(b bool)`

 SetServiceActionPlanResponseNil sets the value for ServiceActionPlanResponse to be an explicit nil

### UnsetServiceActionPlanResponse
`func (o *ServiceRequest) UnsetServiceActionPlanResponse()`

UnsetServiceActionPlanResponse ensures that no value is present for ServiceActionPlanResponse, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


