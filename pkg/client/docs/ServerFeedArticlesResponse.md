# ServerFeedArticlesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Articles** | Pointer to [**[]ModelsArticle**](ModelsArticle.md) |  | [optional] 

## Methods

### NewServerFeedArticlesResponse

`func NewServerFeedArticlesResponse() *ServerFeedArticlesResponse`

NewServerFeedArticlesResponse instantiates a new ServerFeedArticlesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerFeedArticlesResponseWithDefaults

`func NewServerFeedArticlesResponseWithDefaults() *ServerFeedArticlesResponse`

NewServerFeedArticlesResponseWithDefaults instantiates a new ServerFeedArticlesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArticles

`func (o *ServerFeedArticlesResponse) GetArticles() []ModelsArticle`

GetArticles returns the Articles field if non-nil, zero value otherwise.

### GetArticlesOk

`func (o *ServerFeedArticlesResponse) GetArticlesOk() (*[]ModelsArticle, bool)`

GetArticlesOk returns a tuple with the Articles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticles

`func (o *ServerFeedArticlesResponse) SetArticles(v []ModelsArticle)`

SetArticles sets Articles field to given value.

### HasArticles

`func (o *ServerFeedArticlesResponse) HasArticles() bool`

HasArticles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


