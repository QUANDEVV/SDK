# Member

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Username** | **string** |  | 
**Email** | **string** |  | 
**FirstName** | **string** |  | 
**LastName** | **string** |  | 
**Phonenumber** | Pointer to **NullableString** |  | [optional] 
**Gender** | Pointer to [**OneOfGenderEnumBlankEnumNullEnum**](oneOf&lt;GenderEnum,BlankEnum,NullEnum&gt;.md) | Enter your gender.  * &#x60;Male&#x60; - Male * &#x60;Female&#x60; - Female * &#x60;Other&#x60; - Other | [optional] 
**DOB** | Pointer to **NullableString** |  | [optional] 
**JoinTimestamp** | **time.Time** |  | [readonly] 
**IssuesCount** | **int32** |  | [readonly] 
**Rating** | **float64** |  | [readonly] 

## Methods

### NewMember

`func NewMember(id int32, username string, email string, firstName string, lastName string, joinTimestamp time.Time, issuesCount int32, rating float64, ) *Member`

NewMember instantiates a new Member object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberWithDefaults

`func NewMemberWithDefaults() *Member`

NewMemberWithDefaults instantiates a new Member object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Member) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Member) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Member) SetId(v int32)`

SetId sets Id field to given value.


### GetUsername

`func (o *Member) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Member) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Member) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetEmail

`func (o *Member) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Member) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Member) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFirstName

`func (o *Member) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Member) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Member) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *Member) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Member) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Member) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetPhonenumber

`func (o *Member) GetPhonenumber() string`

GetPhonenumber returns the Phonenumber field if non-nil, zero value otherwise.

### GetPhonenumberOk

`func (o *Member) GetPhonenumberOk() (*string, bool)`

GetPhonenumberOk returns a tuple with the Phonenumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhonenumber

`func (o *Member) SetPhonenumber(v string)`

SetPhonenumber sets Phonenumber field to given value.

### HasPhonenumber

`func (o *Member) HasPhonenumber() bool`

HasPhonenumber returns a boolean if a field has been set.

### SetPhonenumberNil

`func (o *Member) SetPhonenumberNil(b bool)`

 SetPhonenumberNil sets the value for Phonenumber to be an explicit nil

### UnsetPhonenumber
`func (o *Member) UnsetPhonenumber()`

UnsetPhonenumber ensures that no value is present for Phonenumber, not even an explicit nil
### GetGender

`func (o *Member) GetGender() OneOfGenderEnumBlankEnumNullEnum`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *Member) GetGenderOk() (*OneOfGenderEnumBlankEnumNullEnum, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *Member) SetGender(v OneOfGenderEnumBlankEnumNullEnum)`

SetGender sets Gender field to given value.

### HasGender

`func (o *Member) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *Member) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *Member) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetDOB

`func (o *Member) GetDOB() string`

GetDOB returns the DOB field if non-nil, zero value otherwise.

### GetDOBOk

`func (o *Member) GetDOBOk() (*string, bool)`

GetDOBOk returns a tuple with the DOB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDOB

`func (o *Member) SetDOB(v string)`

SetDOB sets DOB field to given value.

### HasDOB

`func (o *Member) HasDOB() bool`

HasDOB returns a boolean if a field has been set.

### SetDOBNil

`func (o *Member) SetDOBNil(b bool)`

 SetDOBNil sets the value for DOB to be an explicit nil

### UnsetDOB
`func (o *Member) UnsetDOB()`

UnsetDOB ensures that no value is present for DOB, not even an explicit nil
### GetJoinTimestamp

`func (o *Member) GetJoinTimestamp() time.Time`

GetJoinTimestamp returns the JoinTimestamp field if non-nil, zero value otherwise.

### GetJoinTimestampOk

`func (o *Member) GetJoinTimestampOk() (*time.Time, bool)`

GetJoinTimestampOk returns a tuple with the JoinTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinTimestamp

`func (o *Member) SetJoinTimestamp(v time.Time)`

SetJoinTimestamp sets JoinTimestamp field to given value.


### GetIssuesCount

`func (o *Member) GetIssuesCount() int32`

GetIssuesCount returns the IssuesCount field if non-nil, zero value otherwise.

### GetIssuesCountOk

`func (o *Member) GetIssuesCountOk() (*int32, bool)`

GetIssuesCountOk returns a tuple with the IssuesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuesCount

`func (o *Member) SetIssuesCount(v int32)`

SetIssuesCount sets IssuesCount field to given value.


### GetRating

`func (o *Member) GetRating() float64`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *Member) GetRatingOk() (*float64, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *Member) SetRating(v float64)`

SetRating sets Rating field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


