# ClientInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Username** | **string** |  | 
**Email** | **string** |  | 
**FirstName** | **string** |  | 
**LastName** | **string** |  | 

## Methods

### NewClientInfo

`func NewClientInfo(id int32, username string, email string, firstName string, lastName string, ) *ClientInfo`

NewClientInfo instantiates a new ClientInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientInfoWithDefaults

`func NewClientInfoWithDefaults() *ClientInfo`

NewClientInfoWithDefaults instantiates a new ClientInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClientInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClientInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClientInfo) SetId(v int32)`

SetId sets Id field to given value.


### GetUsername

`func (o *ClientInfo) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ClientInfo) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ClientInfo) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetEmail

`func (o *ClientInfo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ClientInfo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ClientInfo) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFirstName

`func (o *ClientInfo) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ClientInfo) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ClientInfo) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *ClientInfo) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ClientInfo) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ClientInfo) SetLastName(v string)`

SetLastName sets LastName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


