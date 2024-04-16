# VirtualAssistantDocuments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Metadata** | Pointer to **interface{}** | Metadata | [optional] 
**CreatedAt** | **NullableString** |  | [readonly] 
**DateTimeCreatedAt** | **NullableString** |  | [readonly] 
**Timestamp** | **NullableTime** | The timestamp of the chat. | [readonly] 
**UpdatedAt** | **NullableTime** |  | [readonly] 
**Pdf** | **string** | PDF files associated with the virtual assistant. | 
**VirtualAssistantId** | **int32** | The virtual assistant ID. | 

## Methods

### NewVirtualAssistantDocuments

`func NewVirtualAssistantDocuments(id int32, createdAt NullableString, dateTimeCreatedAt NullableString, timestamp NullableTime, updatedAt NullableTime, pdf string, virtualAssistantId int32, ) *VirtualAssistantDocuments`

NewVirtualAssistantDocuments instantiates a new VirtualAssistantDocuments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualAssistantDocumentsWithDefaults

`func NewVirtualAssistantDocumentsWithDefaults() *VirtualAssistantDocuments`

NewVirtualAssistantDocumentsWithDefaults instantiates a new VirtualAssistantDocuments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VirtualAssistantDocuments) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VirtualAssistantDocuments) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VirtualAssistantDocuments) SetId(v int32)`

SetId sets Id field to given value.


### GetMetadata

`func (o *VirtualAssistantDocuments) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VirtualAssistantDocuments) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VirtualAssistantDocuments) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *VirtualAssistantDocuments) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *VirtualAssistantDocuments) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *VirtualAssistantDocuments) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetCreatedAt

`func (o *VirtualAssistantDocuments) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VirtualAssistantDocuments) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VirtualAssistantDocuments) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *VirtualAssistantDocuments) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *VirtualAssistantDocuments) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDateTimeCreatedAt

`func (o *VirtualAssistantDocuments) GetDateTimeCreatedAt() string`

GetDateTimeCreatedAt returns the DateTimeCreatedAt field if non-nil, zero value otherwise.

### GetDateTimeCreatedAtOk

`func (o *VirtualAssistantDocuments) GetDateTimeCreatedAtOk() (*string, bool)`

GetDateTimeCreatedAtOk returns a tuple with the DateTimeCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeCreatedAt

`func (o *VirtualAssistantDocuments) SetDateTimeCreatedAt(v string)`

SetDateTimeCreatedAt sets DateTimeCreatedAt field to given value.


### SetDateTimeCreatedAtNil

`func (o *VirtualAssistantDocuments) SetDateTimeCreatedAtNil(b bool)`

 SetDateTimeCreatedAtNil sets the value for DateTimeCreatedAt to be an explicit nil

### UnsetDateTimeCreatedAt
`func (o *VirtualAssistantDocuments) UnsetDateTimeCreatedAt()`

UnsetDateTimeCreatedAt ensures that no value is present for DateTimeCreatedAt, not even an explicit nil
### GetTimestamp

`func (o *VirtualAssistantDocuments) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *VirtualAssistantDocuments) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *VirtualAssistantDocuments) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### SetTimestampNil

`func (o *VirtualAssistantDocuments) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *VirtualAssistantDocuments) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetUpdatedAt

`func (o *VirtualAssistantDocuments) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VirtualAssistantDocuments) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VirtualAssistantDocuments) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *VirtualAssistantDocuments) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *VirtualAssistantDocuments) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetPdf

`func (o *VirtualAssistantDocuments) GetPdf() string`

GetPdf returns the Pdf field if non-nil, zero value otherwise.

### GetPdfOk

`func (o *VirtualAssistantDocuments) GetPdfOk() (*string, bool)`

GetPdfOk returns a tuple with the Pdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdf

`func (o *VirtualAssistantDocuments) SetPdf(v string)`

SetPdf sets Pdf field to given value.


### GetVirtualAssistantId

`func (o *VirtualAssistantDocuments) GetVirtualAssistantId() int32`

GetVirtualAssistantId returns the VirtualAssistantId field if non-nil, zero value otherwise.

### GetVirtualAssistantIdOk

`func (o *VirtualAssistantDocuments) GetVirtualAssistantIdOk() (*int32, bool)`

GetVirtualAssistantIdOk returns a tuple with the VirtualAssistantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualAssistantId

`func (o *VirtualAssistantDocuments) SetVirtualAssistantId(v int32)`

SetVirtualAssistantId sets VirtualAssistantId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


