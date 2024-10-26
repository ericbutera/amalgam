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

// checks if the ServerArticleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerArticleResponse{}

// ServerArticleResponse struct for ServerArticleResponse
type ServerArticleResponse struct {
	Article *ServiceArticle `json:"article,omitempty"`
}

// NewServerArticleResponse instantiates a new ServerArticleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerArticleResponse() *ServerArticleResponse {
	this := ServerArticleResponse{}
	return &this
}

// NewServerArticleResponseWithDefaults instantiates a new ServerArticleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerArticleResponseWithDefaults() *ServerArticleResponse {
	this := ServerArticleResponse{}
	return &this
}

// GetArticle returns the Article field value if set, zero value otherwise.
func (o *ServerArticleResponse) GetArticle() ServiceArticle {
	if o == nil || IsNil(o.Article) {
		var ret ServiceArticle
		return ret
	}
	return *o.Article
}

// GetArticleOk returns a tuple with the Article field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerArticleResponse) GetArticleOk() (*ServiceArticle, bool) {
	if o == nil || IsNil(o.Article) {
		return nil, false
	}
	return o.Article, true
}

// HasArticle returns a boolean if a field has been set.
func (o *ServerArticleResponse) HasArticle() bool {
	if o != nil && !IsNil(o.Article) {
		return true
	}

	return false
}

// SetArticle gets a reference to the given ServiceArticle and assigns it to the Article field.
func (o *ServerArticleResponse) SetArticle(v ServiceArticle) {
	o.Article = &v
}

func (o ServerArticleResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerArticleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Article) {
		toSerialize["article"] = o.Article
	}
	return toSerialize, nil
}

type NullableServerArticleResponse struct {
	value *ServerArticleResponse
	isSet bool
}

func (v NullableServerArticleResponse) Get() *ServerArticleResponse {
	return v.value
}

func (v *NullableServerArticleResponse) Set(val *ServerArticleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableServerArticleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableServerArticleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerArticleResponse(val *ServerArticleResponse) *NullableServerArticleResponse {
	return &NullableServerArticleResponse{value: val, isSet: true}
}

func (v NullableServerArticleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerArticleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
