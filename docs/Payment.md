# Payment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | **time.Time** |  | [readonly] 
**MerchantRef** | **string** |  | 
**TransactionId** | Pointer to **NullableString** |  | [optional] 
**Amount** | **float64** |  | 
**Complete** | Pointer to **bool** |  | [optional] 
**Payload** | Pointer to **interface{}** |  | [optional] 
**PaymentPayload** | Pointer to **interface{}** |  | [optional] 
**Subscription** | **int32** |  | 

## Methods

### NewPayment

`func NewPayment(id int32, createdAt time.Time, updatedAt time.Time, merchantRef string, amount float64, subscription int32, ) *Payment`

NewPayment instantiates a new Payment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentWithDefaults

`func NewPaymentWithDefaults() *Payment`

NewPaymentWithDefaults instantiates a new Payment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Payment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Payment) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Payment) SetId(v int32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *Payment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Payment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Payment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Payment) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Payment) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Payment) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetMerchantRef

`func (o *Payment) GetMerchantRef() string`

GetMerchantRef returns the MerchantRef field if non-nil, zero value otherwise.

### GetMerchantRefOk

`func (o *Payment) GetMerchantRefOk() (*string, bool)`

GetMerchantRefOk returns a tuple with the MerchantRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantRef

`func (o *Payment) SetMerchantRef(v string)`

SetMerchantRef sets MerchantRef field to given value.


### GetTransactionId

`func (o *Payment) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *Payment) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *Payment) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *Payment) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### SetTransactionIdNil

`func (o *Payment) SetTransactionIdNil(b bool)`

 SetTransactionIdNil sets the value for TransactionId to be an explicit nil

### UnsetTransactionId
`func (o *Payment) UnsetTransactionId()`

UnsetTransactionId ensures that no value is present for TransactionId, not even an explicit nil
### GetAmount

`func (o *Payment) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Payment) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Payment) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetComplete

`func (o *Payment) GetComplete() bool`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *Payment) GetCompleteOk() (*bool, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *Payment) SetComplete(v bool)`

SetComplete sets Complete field to given value.

### HasComplete

`func (o *Payment) HasComplete() bool`

HasComplete returns a boolean if a field has been set.

### GetPayload

`func (o *Payment) GetPayload() interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *Payment) GetPayloadOk() (*interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *Payment) SetPayload(v interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *Payment) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### SetPayloadNil

`func (o *Payment) SetPayloadNil(b bool)`

 SetPayloadNil sets the value for Payload to be an explicit nil

### UnsetPayload
`func (o *Payment) UnsetPayload()`

UnsetPayload ensures that no value is present for Payload, not even an explicit nil
### GetPaymentPayload

`func (o *Payment) GetPaymentPayload() interface{}`

GetPaymentPayload returns the PaymentPayload field if non-nil, zero value otherwise.

### GetPaymentPayloadOk

`func (o *Payment) GetPaymentPayloadOk() (*interface{}, bool)`

GetPaymentPayloadOk returns a tuple with the PaymentPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentPayload

`func (o *Payment) SetPaymentPayload(v interface{})`

SetPaymentPayload sets PaymentPayload field to given value.

### HasPaymentPayload

`func (o *Payment) HasPaymentPayload() bool`

HasPaymentPayload returns a boolean if a field has been set.

### SetPaymentPayloadNil

`func (o *Payment) SetPaymentPayloadNil(b bool)`

 SetPaymentPayloadNil sets the value for PaymentPayload to be an explicit nil

### UnsetPaymentPayload
`func (o *Payment) UnsetPaymentPayload()`

UnsetPaymentPayload ensures that no value is present for PaymentPayload, not even an explicit nil
### GetSubscription

`func (o *Payment) GetSubscription() int32`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *Payment) GetSubscriptionOk() (*int32, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *Payment) SetSubscription(v int32)`

SetSubscription sets Subscription field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


