# ServerFeedsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Feeds** | Pointer to [**[]ServerListFeed**](ServerListFeed.md) |  | [optional]

## Methods

### NewServerFeedsResponse

`func NewServerFeedsResponse() *ServerFeedsResponse`

NewServerFeedsResponse instantiates a new ServerFeedsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerFeedsResponseWithDefaults

`func NewServerFeedsResponseWithDefaults() *ServerFeedsResponse`

NewServerFeedsResponseWithDefaults instantiates a new ServerFeedsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeds

`func (o *ServerFeedsResponse) GetFeeds() []ServerListFeed`

GetFeeds returns the Feeds field if non-nil, zero value otherwise.

### GetFeedsOk

`func (o *ServerFeedsResponse) GetFeedsOk() (*[]ServerListFeed, bool)`

GetFeedsOk returns a tuple with the Feeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeds

`func (o *ServerFeedsResponse) SetFeeds(v []ServerListFeed)`

SetFeeds sets Feeds field to given value.

### HasFeeds

`func (o *ServerFeedsResponse) HasFeeds() bool`

HasFeeds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
