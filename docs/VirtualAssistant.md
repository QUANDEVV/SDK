# VirtualAssistant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VirtualAssistantId** | **int32** | The tenant chats ID. | [readonly] 
**Metadata** | Pointer to **interface{}** | Metadata | [optional] 
**CreatedAt** | **NullableString** |  | [readonly] 
**DateTimeCreatedAt** | **NullableString** |  | [readonly] 
**Timestamp** | **NullableTime** | The timestamp of the chat. | [readonly] 
**UpdatedAt** | **NullableTime** |  | [readonly] 
**FirstName** | **string** | name of the company | 
**Industry** | **string** | industry of the company | 
**Nickname** | **string** | industry of the company | 
**Persona** | **string** | industry of the company | 
**Temprature** | **string** | industry of the company | 
**Accuracy** | **string** | industry of the company | 
**TenantId** | **int32** | Display name of the tenant | 

## Methods

### NewVirtualAssistant

`func NewVirtualAssistant(virtualAssistantId int32, createdAt NullableString, dateTimeCreatedAt NullableString, timestamp NullableTime, updatedAt NullableTime, firstName string, industry string, nickname string, persona string, temprature string, accuracy string, tenantId int32, ) *VirtualAssistant`

NewVirtualAssistant instantiates a new VirtualAssistant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualAssistantWithDefaults

`func NewVirtualAssistantWithDefaults() *VirtualAssistant`

NewVirtualAssistantWithDefaults instantiates a new VirtualAssistant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVirtualAssistantId

`func (o *VirtualAssistant) GetVirtualAssistantId() int32`

GetVirtualAssistantId returns the VirtualAssistantId field if non-nil, zero value otherwise.

### GetVirtualAssistantIdOk

`func (o *VirtualAssistant) GetVirtualAssistantIdOk() (*int32, bool)`

GetVirtualAssistantIdOk returns a tuple with the VirtualAssistantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualAssistantId

`func (o *VirtualAssistant) SetVirtualAssistantId(v int32)`

SetVirtualAssistantId sets VirtualAssistantId field to given value.


### GetMetadata

`func (o *VirtualAssistant) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VirtualAssistant) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VirtualAssistant) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *VirtualAssistant) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *VirtualAssistant) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *VirtualAssistant) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetCreatedAt

`func (o *VirtualAssistant) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VirtualAssistant) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VirtualAssistant) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *VirtualAssistant) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *VirtualAssistant) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDateTimeCreatedAt

`func (o *VirtualAssistant) GetDateTimeCreatedAt() string`

GetDateTimeCreatedAt returns the DateTimeCreatedAt field if non-nil, zero value otherwise.

### GetDateTimeCreatedAtOk

`func (o *VirtualAssistant) GetDateTimeCreatedAtOk() (*string, bool)`

GetDateTimeCreatedAtOk returns a tuple with the DateTimeCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeCreatedAt

`func (o *VirtualAssistant) SetDateTimeCreatedAt(v string)`

SetDateTimeCreatedAt sets DateTimeCreatedAt field to given value.


### SetDateTimeCreatedAtNil

`func (o *VirtualAssistant) SetDateTimeCreatedAtNil(b bool)`

 SetDateTimeCreatedAtNil sets the value for DateTimeCreatedAt to be an explicit nil

### UnsetDateTimeCreatedAt
`func (o *VirtualAssistant) UnsetDateTimeCreatedAt()`

UnsetDateTimeCreatedAt ensures that no value is present for DateTimeCreatedAt, not even an explicit nil
### GetTimestamp

`func (o *VirtualAssistant) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *VirtualAssistant) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *VirtualAssistant) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### SetTimestampNil

`func (o *VirtualAssistant) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *VirtualAssistant) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetUpdatedAt

`func (o *VirtualAssistant) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VirtualAssistant) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VirtualAssistant) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *VirtualAssistant) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *VirtualAssistant) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetFirstName

`func (o *VirtualAssistant) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *VirtualAssistant) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *VirtualAssistant) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetIndustry

`func (o *VirtualAssistant) GetIndustry() string`

GetIndustry returns the Industry field if non-nil, zero value otherwise.

### GetIndustryOk

`func (o *VirtualAssistant) GetIndustryOk() (*string, bool)`

GetIndustryOk returns a tuple with the Industry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustry

`func (o *VirtualAssistant) SetIndustry(v string)`

SetIndustry sets Industry field to given value.


### GetNickname

`func (o *VirtualAssistant) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *VirtualAssistant) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *VirtualAssistant) SetNickname(v string)`

SetNickname sets Nickname field to given value.


### GetPersona

`func (o *VirtualAssistant) GetPersona() string`

GetPersona returns the Persona field if non-nil, zero value otherwise.

### GetPersonaOk

`func (o *VirtualAssistant) GetPersonaOk() (*string, bool)`

GetPersonaOk returns a tuple with the Persona field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersona

`func (o *VirtualAssistant) SetPersona(v string)`

SetPersona sets Persona field to given value.


### GetTemprature

`func (o *VirtualAssistant) GetTemprature() string`

GetTemprature returns the Temprature field if non-nil, zero value otherwise.

### GetTempratureOk

`func (o *VirtualAssistant) GetTempratureOk() (*string, bool)`

GetTempratureOk returns a tuple with the Temprature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemprature

`func (o *VirtualAssistant) SetTemprature(v string)`

SetTemprature sets Temprature field to given value.


### GetAccuracy

`func (o *VirtualAssistant) GetAccuracy() string`

GetAccuracy returns the Accuracy field if non-nil, zero value otherwise.

### GetAccuracyOk

`func (o *VirtualAssistant) GetAccuracyOk() (*string, bool)`

GetAccuracyOk returns a tuple with the Accuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccuracy

`func (o *VirtualAssistant) SetAccuracy(v string)`

SetAccuracy sets Accuracy field to given value.


### GetTenantId

`func (o *VirtualAssistant) GetTenantId() int32`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *VirtualAssistant) GetTenantIdOk() (*int32, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *VirtualAssistant) SetTenantId(v int32)`

SetTenantId sets TenantId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


