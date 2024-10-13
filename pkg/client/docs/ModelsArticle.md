# ModelsArticle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorEmail** | Pointer to **string** |  | [optional] 
**AuthorName** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**DeletedAt** | Pointer to [**GormDeletedAt**](GormDeletedAt.md) |  | [optional] 
**Feed** | Pointer to [**ModelsFeed**](ModelsFeed.md) |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**ImageUrl** | Pointer to **string** |  | [optional] 
**Preview** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewModelsArticle

`func NewModelsArticle() *ModelsArticle`

NewModelsArticle instantiates a new ModelsArticle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsArticleWithDefaults

`func NewModelsArticleWithDefaults() *ModelsArticle`

NewModelsArticleWithDefaults instantiates a new ModelsArticle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorEmail

`func (o *ModelsArticle) GetAuthorEmail() string`

GetAuthorEmail returns the AuthorEmail field if non-nil, zero value otherwise.

### GetAuthorEmailOk

`func (o *ModelsArticle) GetAuthorEmailOk() (*string, bool)`

GetAuthorEmailOk returns a tuple with the AuthorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorEmail

`func (o *ModelsArticle) SetAuthorEmail(v string)`

SetAuthorEmail sets AuthorEmail field to given value.

### HasAuthorEmail

`func (o *ModelsArticle) HasAuthorEmail() bool`

HasAuthorEmail returns a boolean if a field has been set.

### GetAuthorName

`func (o *ModelsArticle) GetAuthorName() string`

GetAuthorName returns the AuthorName field if non-nil, zero value otherwise.

### GetAuthorNameOk

`func (o *ModelsArticle) GetAuthorNameOk() (*string, bool)`

GetAuthorNameOk returns a tuple with the AuthorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorName

`func (o *ModelsArticle) SetAuthorName(v string)`

SetAuthorName sets AuthorName field to given value.

### HasAuthorName

`func (o *ModelsArticle) HasAuthorName() bool`

HasAuthorName returns a boolean if a field has been set.

### GetContent

`func (o *ModelsArticle) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ModelsArticle) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ModelsArticle) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *ModelsArticle) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ModelsArticle) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelsArticle) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelsArticle) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelsArticle) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ModelsArticle) GetDeletedAt() GormDeletedAt`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ModelsArticle) GetDeletedAtOk() (*GormDeletedAt, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ModelsArticle) SetDeletedAt(v GormDeletedAt)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ModelsArticle) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetFeed

`func (o *ModelsArticle) GetFeed() ModelsFeed`

GetFeed returns the Feed field if non-nil, zero value otherwise.

### GetFeedOk

`func (o *ModelsArticle) GetFeedOk() (*ModelsFeed, bool)`

GetFeedOk returns a tuple with the Feed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeed

`func (o *ModelsArticle) SetFeed(v ModelsFeed)`

SetFeed sets Feed field to given value.

### HasFeed

`func (o *ModelsArticle) HasFeed() bool`

HasFeed returns a boolean if a field has been set.

### GetGuid

`func (o *ModelsArticle) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *ModelsArticle) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *ModelsArticle) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *ModelsArticle) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetId

`func (o *ModelsArticle) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsArticle) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsArticle) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsArticle) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImageUrl

`func (o *ModelsArticle) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *ModelsArticle) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *ModelsArticle) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *ModelsArticle) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetPreview

`func (o *ModelsArticle) GetPreview() string`

GetPreview returns the Preview field if non-nil, zero value otherwise.

### GetPreviewOk

`func (o *ModelsArticle) GetPreviewOk() (*string, bool)`

GetPreviewOk returns a tuple with the Preview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreview

`func (o *ModelsArticle) SetPreview(v string)`

SetPreview sets Preview field to given value.

### HasPreview

`func (o *ModelsArticle) HasPreview() bool`

HasPreview returns a boolean if a field has been set.

### GetTitle

`func (o *ModelsArticle) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ModelsArticle) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ModelsArticle) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ModelsArticle) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ModelsArticle) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelsArticle) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelsArticle) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ModelsArticle) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *ModelsArticle) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ModelsArticle) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ModelsArticle) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ModelsArticle) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


