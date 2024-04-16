# FavoriteCommunitiesClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** |  | 
**Email** | **string** |  | 
**Favorites** | [**[]FavoriteCommunity**](FavoriteCommunity.md) |  | 

## Methods

### NewFavoriteCommunitiesClient

`func NewFavoriteCommunitiesClient(username string, email string, favorites []FavoriteCommunity, ) *FavoriteCommunitiesClient`

NewFavoriteCommunitiesClient instantiates a new FavoriteCommunitiesClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFavoriteCommunitiesClientWithDefaults

`func NewFavoriteCommunitiesClientWithDefaults() *FavoriteCommunitiesClient`

NewFavoriteCommunitiesClientWithDefaults instantiates a new FavoriteCommunitiesClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *FavoriteCommunitiesClient) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *FavoriteCommunitiesClient) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *FavoriteCommunitiesClient) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetEmail

`func (o *FavoriteCommunitiesClient) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *FavoriteCommunitiesClient) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *FavoriteCommunitiesClient) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFavorites

`func (o *FavoriteCommunitiesClient) GetFavorites() []FavoriteCommunity`

GetFavorites returns the Favorites field if non-nil, zero value otherwise.

### GetFavoritesOk

`func (o *FavoriteCommunitiesClient) GetFavoritesOk() (*[]FavoriteCommunity, bool)`

GetFavoritesOk returns a tuple with the Favorites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavorites

`func (o *FavoriteCommunitiesClient) SetFavorites(v []FavoriteCommunity)`

SetFavorites sets Favorites field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


