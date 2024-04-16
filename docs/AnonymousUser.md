# AnonymousUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnonymousUserId** | **int32** | The webhookevent chats ID UUID. | [readonly] 
**Token** | **string** |  | [readonly] 
**Email** | Pointer to **NullableString** | Email for anonymous users | [optional] 
**FullNames** | Pointer to **NullableString** |  | [optional] 
**Contact** | Pointer to **NullableString** | Contact related to an anonymous user | [optional] 

## Methods

### NewAnonymousUser

`func NewAnonymousUser(anonymousUserId int32, token string, ) *AnonymousUser`

NewAnonymousUser instantiates a new AnonymousUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnonymousUserWithDefaults

`func NewAnonymousUserWithDefaults() *AnonymousUser`

NewAnonymousUserWithDefaults instantiates a new AnonymousUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnonymousUserId

`func (o *AnonymousUser) GetAnonymousUserId() int32`

GetAnonymousUserId returns the AnonymousUserId field if non-nil, zero value otherwise.

### GetAnonymousUserIdOk

`func (o *AnonymousUser) GetAnonymousUserIdOk() (*int32, bool)`

GetAnonymousUserIdOk returns a tuple with the AnonymousUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymousUserId

`func (o *AnonymousUser) SetAnonymousUserId(v int32)`

SetAnonymousUserId sets AnonymousUserId field to given value.


### GetToken

`func (o *AnonymousUser) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AnonymousUser) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AnonymousUser) SetToken(v string)`

SetToken sets Token field to given value.


### GetEmail

`func (o *AnonymousUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AnonymousUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AnonymousUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AnonymousUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *AnonymousUser) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *AnonymousUser) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetFullNames

`func (o *AnonymousUser) GetFullNames() string`

GetFullNames returns the FullNames field if non-nil, zero value otherwise.

### GetFullNamesOk

`func (o *AnonymousUser) GetFullNamesOk() (*string, bool)`

GetFullNamesOk returns a tuple with the FullNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullNames

`func (o *AnonymousUser) SetFullNames(v string)`

SetFullNames sets FullNames field to given value.

### HasFullNames

`func (o *AnonymousUser) HasFullNames() bool`

HasFullNames returns a boolean if a field has been set.

### SetFullNamesNil

`func (o *AnonymousUser) SetFullNamesNil(b bool)`

 SetFullNamesNil sets the value for FullNames to be an explicit nil

### UnsetFullNames
`func (o *AnonymousUser) UnsetFullNames()`

UnsetFullNames ensures that no value is present for FullNames, not even an explicit nil
### GetContact

`func (o *AnonymousUser) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *AnonymousUser) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *AnonymousUser) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *AnonymousUser) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *AnonymousUser) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *AnonymousUser) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


