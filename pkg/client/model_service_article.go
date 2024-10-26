/*
Feed API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ServiceArticle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceArticle{}

// ServiceArticle struct for ServiceArticle
type ServiceArticle struct {
	AuthorEmail *string `json:"authorEmail,omitempty"`
	AuthorName *string `json:"authorName,omitempty"`
	Content *string `json:"content,omitempty"`
	FeedId *string `json:"feedId,omitempty"`
	Guid *string `json:"guid,omitempty"`
	Id *string `json:"id,omitempty"`
	ImageUrl *string `json:"imageUrl,omitempty"`
	Preview *string `json:"preview,omitempty"`
	Title *string `json:"title,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewServiceArticle instantiates a new ServiceArticle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceArticle() *ServiceArticle {
	this := ServiceArticle{}
	return &this
}

// NewServiceArticleWithDefaults instantiates a new ServiceArticle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceArticleWithDefaults() *ServiceArticle {
	this := ServiceArticle{}
	return &this
}

// GetAuthorEmail returns the AuthorEmail field value if set, zero value otherwise.
func (o *ServiceArticle) GetAuthorEmail() string {
	if o == nil || IsNil(o.AuthorEmail) {
		var ret string
		return ret
	}
	return *o.AuthorEmail
}

// GetAuthorEmailOk returns a tuple with the AuthorEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceArticle) GetAuthorEmailOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorEmail) {
		return nil, false
	}
	return o.AuthorEmail, true
}

// HasAuthorEmail returns a boolean if a field has been set.
func (o *ServiceArticle) HasAuthorEmail() bool {
	if o != nil && !IsNil(o.AuthorEmail) {
		return true
	}

	return false
}

// SetAuthorEmail gets a reference to the given string and assigns it to the AuthorEmail field.
func (o *ServiceArticle) SetAuthorEmail(v string) {
	o.AuthorEmail = &v
}

// GetAuthorName returns the AuthorName field value if set, zero value otherwise.
func (o *ServiceArticle) GetAuthorName() string {
	if o == nil || IsNil(o.AuthorName) {
		var ret string
		return ret
	}
	return *o.AuthorName
}

// GetAuthorNameOk returns a tuple with the AuthorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceArticle) GetAuthorNameOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorName) {
		return nil, false
	}
	return o.AuthorName, true
}

// HasAuthorName returns a boolean if a field has been set.
func (o *ServiceArticle) HasAuthorName() bool {
	if o != nil && !IsNil(o.AuthorName) {
		return true
	}

	return false
}

// SetAuthorName gets a reference to the given string and assigns it to the AuthorName field.
func (o *ServiceArticle) SetAuthorName(v string) {
	o.AuthorName = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *ServiceArticle) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceArticle) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *ServiceArticle) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *ServiceArticle) SetContent(v string) {
	o.Content = &v
}

// GetFeedId returns the FeedId field value if set, zero value otherwise.
func (o *ServiceArticle) GetFeedId() string {
	if o == nil || IsNil(o.FeedId) {
		var ret string
		return ret
	}
	return *o.FeedId
}

// GetFeedIdOk returns a tuple with the FeedId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceArticle) GetFeedIdOk() (*string, bool) {
	if o == nil || IsNil(o.FeedId) {
		return nil, false
	}
	return o.FeedId, true
}

// HasFeedId returns a boolean if a field has been set.
func (o *ServiceArticle) HasFeedId() bool {
	if o != nil && !IsNil(o.FeedId) {
		return true
	}

	return false
}

// SetFeedId gets a reference to the given string and assigns it to the FeedId field.
func (o *ServiceArticle) SetFeedId(v string) {
	o.FeedId = &v
}

// GetGuid returns the Guid field value if set, zero value otherwise.
func (o *ServiceArticle) GetGuid() string {
	if o == nil || IsNil(o.Guid) {
		var ret string
		return ret
	}
	return *o.Guid
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceArticle) GetGuidOk() (*string, bool) {
	if o == nil || IsNil(o.Guid) {
		return nil, false
	}
	return o.Guid, true
}

// HasGuid returns a boolean if a field has been set.
func (o *ServiceArticle) HasGuid() bool {
	if o != nil && !IsNil(o.Guid) {
		return true
	}

	return false
}

// SetGuid gets a reference to the given string and assigns it to the Guid field.
func (o *ServiceArticle) SetGuid(v string) {
	o.Guid = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServiceArticle) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceArticle) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServiceArticle) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ServiceArticle) SetId(v string) {
	o.Id = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *ServiceArticle) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceArticle) GetImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ImageUrl) {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *ServiceArticle) HasImageUrl() bool {
	if o != nil && !IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *ServiceArticle) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetPreview returns the Preview field value if set, zero value otherwise.
func (o *ServiceArticle) GetPreview() string {
	if o == nil || IsNil(o.Preview) {
		var ret string
		return ret
	}
	return *o.Preview
}

// GetPreviewOk returns a tuple with the Preview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceArticle) GetPreviewOk() (*string, bool) {
	if o == nil || IsNil(o.Preview) {
		return nil, false
	}
	return o.Preview, true
}

// HasPreview returns a boolean if a field has been set.
func (o *ServiceArticle) HasPreview() bool {
	if o != nil && !IsNil(o.Preview) {
		return true
	}

	return false
}

// SetPreview gets a reference to the given string and assigns it to the Preview field.
func (o *ServiceArticle) SetPreview(v string) {
	o.Preview = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ServiceArticle) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceArticle) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ServiceArticle) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ServiceArticle) SetTitle(v string) {
	o.Title = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ServiceArticle) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceArticle) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ServiceArticle) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ServiceArticle) SetUrl(v string) {
	o.Url = &v
}

func (o ServiceArticle) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceArticle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthorEmail) {
		toSerialize["authorEmail"] = o.AuthorEmail
	}
	if !IsNil(o.AuthorName) {
		toSerialize["authorName"] = o.AuthorName
	}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.FeedId) {
		toSerialize["feedId"] = o.FeedId
	}
	if !IsNil(o.Guid) {
		toSerialize["guid"] = o.Guid
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ImageUrl) {
		toSerialize["imageUrl"] = o.ImageUrl
	}
	if !IsNil(o.Preview) {
		toSerialize["preview"] = o.Preview
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableServiceArticle struct {
	value *ServiceArticle
	isSet bool
}

func (v NullableServiceArticle) Get() *ServiceArticle {
	return v.value
}

func (v *NullableServiceArticle) Set(val *ServiceArticle) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceArticle) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceArticle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceArticle(val *ServiceArticle) *NullableServiceArticle {
	return &NullableServiceArticle{value: val, isSet: true}
}

func (v NullableServiceArticle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceArticle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}