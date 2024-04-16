# PatchedServiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **NullableString** |  | [optional] [readonly] 
**ServiceRequestsId** | Pointer to **int32** |  | [optional] [readonly] 
**Client** | Pointer to [**ClientServicesInfo**](ClientServicesInfo.md) |  | [optional] 
**UnauthenticatedUser** | Pointer to [**AnonymousUser**](AnonymousUser.md) |  | [optional] 
**Service** | Pointer to [**Service**](Service.md) |  | [optional] 
**Status** | Pointer to [**ServiceRequestStatusEnum**](ServiceRequestStatusEnum.md) | Status of the request.  * &#x60;pending&#x60; - Pending * &#x60;completed&#x60; - Completed * &#x60;in_progress&#x60; - In Progress | [optional] 
**RequestDetails** | Pointer to **interface{}** | Specific request details in JSON format. | [optional] 
**AiInfo** | Pointer to **interface{}** | A json dump of the ai content generated | [optional] 
**ServiceActionPlanResponse** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewPatchedServiceRequest

`func NewPatchedServiceRequest() *PatchedServiceRequest`

NewPatchedServiceRequest instantiates a new PatchedServiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedServiceRequestWithDefaults

`func NewPatchedServiceRequestWithDefaults() *PatchedServiceRequest`

NewPatchedServiceRequestWithDefaults instantiates a new PatchedServiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *PatchedServiceRequest) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedServiceRequest) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedServiceRequest) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedServiceRequest) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *PatchedServiceRequest) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *PatchedServiceRequest) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetServiceRequestsId

`func (o *PatchedServiceRequest) GetServiceRequestsId() int32`

GetServiceRequestsId returns the ServiceRequestsId field if non-nil, zero value otherwise.

### GetServiceRequestsIdOk

`func (o *PatchedServiceRequest) GetServiceRequestsIdOk() (*int32, bool)`

GetServiceRequestsIdOk returns a tuple with the ServiceRequestsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRequestsId

`func (o *PatchedServiceRequest) SetServiceRequestsId(v int32)`

SetServiceRequestsId sets ServiceRequestsId field to given value.

### HasServiceRequestsId

`func (o *PatchedServiceRequest) HasServiceRequestsId() bool`

HasServiceRequestsId returns a boolean if a field has been set.

### GetClient

`func (o *PatchedServiceRequest) GetClient() ClientServicesInfo`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *PatchedServiceRequest) GetClientOk() (*ClientServicesInfo, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *PatchedServiceRequest) SetClient(v ClientServicesInfo)`

SetClient sets Client field to given value.

### HasClient

`func (o *PatchedServiceRequest) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetUnauthenticatedUser

`func (o *PatchedServiceRequest) GetUnauthenticatedUser() AnonymousUser`

GetUnauthenticatedUser returns the UnauthenticatedUser field if non-nil, zero value otherwise.

### GetUnauthenticatedUserOk

`func (o *PatchedServiceRequest) GetUnauthenticatedUserOk() (*AnonymousUser, bool)`

GetUnauthenticatedUserOk returns a tuple with the UnauthenticatedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnauthenticatedUser

`func (o *PatchedServiceRequest) SetUnauthenticatedUser(v AnonymousUser)`

SetUnauthenticatedUser sets UnauthenticatedUser field to given value.

### HasUnauthenticatedUser

`func (o *PatchedServiceRequest) HasUnauthenticatedUser() bool`

HasUnauthenticatedUser returns a boolean if a field has been set.

### GetService

`func (o *PatchedServiceRequest) GetService() Service`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *PatchedServiceRequest) GetServiceOk() (*Service, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *PatchedServiceRequest) SetService(v Service)`

SetService sets Service field to given value.

### HasService

`func (o *PatchedServiceRequest) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedServiceRequest) GetStatus() ServiceRequestStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedServiceRequest) GetStatusOk() (*ServiceRequestStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedServiceRequest) SetStatus(v ServiceRequestStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedServiceRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *PatchedServiceRequest) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *PatchedServiceRequest) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetRequestDetails

`func (o *PatchedServiceRequest) GetRequestDetails() interface{}`

GetRequestDetails returns the RequestDetails field if non-nil, zero value otherwise.

### GetRequestDetailsOk

`func (o *PatchedServiceRequest) GetRequestDetailsOk() (*interface{}, bool)`

GetRequestDetailsOk returns a tuple with the RequestDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestDetails

`func (o *PatchedServiceRequest) SetRequestDetails(v interface{})`

SetRequestDetails sets RequestDetails field to given value.

### HasRequestDetails

`func (o *PatchedServiceRequest) HasRequestDetails() bool`

HasRequestDetails returns a boolean if a field has been set.

### SetRequestDetailsNil

`func (o *PatchedServiceRequest) SetRequestDetailsNil(b bool)`

 SetRequestDetailsNil sets the value for RequestDetails to be an explicit nil

### UnsetRequestDetails
`func (o *PatchedServiceRequest) UnsetRequestDetails()`

UnsetRequestDetails ensures that no value is present for RequestDetails, not even an explicit nil
### GetAiInfo

`func (o *PatchedServiceRequest) GetAiInfo() interface{}`

GetAiInfo returns the AiInfo field if non-nil, zero value otherwise.

### GetAiInfoOk

`func (o *PatchedServiceRequest) GetAiInfoOk() (*interface{}, bool)`

GetAiInfoOk returns a tuple with the AiInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiInfo

`func (o *PatchedServiceRequest) SetAiInfo(v interface{})`

SetAiInfo sets AiInfo field to given value.

### HasAiInfo

`func (o *PatchedServiceRequest) HasAiInfo() bool`

HasAiInfo returns a boolean if a field has been set.

### SetAiInfoNil

`func (o *PatchedServiceRequest) SetAiInfoNil(b bool)`

 SetAiInfoNil sets the value for AiInfo to be an explicit nil

### UnsetAiInfo
`func (o *PatchedServiceRequest) UnsetAiInfo()`

UnsetAiInfo ensures that no value is present for AiInfo, not even an explicit nil
### GetServiceActionPlanResponse

`func (o *PatchedServiceRequest) GetServiceActionPlanResponse() interface{}`

GetServiceActionPlanResponse returns the ServiceActionPlanResponse field if non-nil, zero value otherwise.

### GetServiceActionPlanResponseOk

`func (o *PatchedServiceRequest) GetServiceActionPlanResponseOk() (*interface{}, bool)`

GetServiceActionPlanResponseOk returns a tuple with the ServiceActionPlanResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceActionPlanResponse

`func (o *PatchedServiceRequest) SetServiceActionPlanResponse(v interface{})`

SetServiceActionPlanResponse sets ServiceActionPlanResponse field to given value.

### HasServiceActionPlanResponse

`func (o *PatchedServiceRequest) HasServiceActionPlanResponse() bool`

HasServiceActionPlanResponse returns a boolean if a field has been set.

### SetServiceActionPlanResponseNil

`func (o *PatchedServiceRequest) SetServiceActionPlanResponseNil(b bool)`

 SetServiceActionPlanResponseNil sets the value for ServiceActionPlanResponse to be an explicit nil

### UnsetServiceActionPlanResponse
`func (o *PatchedServiceRequest) UnsetServiceActionPlanResponse()`

UnsetServiceActionPlanResponse ensures that no value is present for ServiceActionPlanResponse, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


