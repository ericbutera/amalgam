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

// checks if the ServerFeedArticlesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerFeedArticlesResponse{}

// ServerFeedArticlesResponse struct for ServerFeedArticlesResponse
type ServerFeedArticlesResponse struct {
	Articles []ServiceArticle `json:"articles,omitempty"`
}

// NewServerFeedArticlesResponse instantiates a new ServerFeedArticlesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerFeedArticlesResponse() *ServerFeedArticlesResponse {
	this := ServerFeedArticlesResponse{}
	return &this
}

// NewServerFeedArticlesResponseWithDefaults instantiates a new ServerFeedArticlesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerFeedArticlesResponseWithDefaults() *ServerFeedArticlesResponse {
	this := ServerFeedArticlesResponse{}
	return &this
}

// GetArticles returns the Articles field value if set, zero value otherwise.
func (o *ServerFeedArticlesResponse) GetArticles() []ServiceArticle {
	if o == nil || IsNil(o.Articles) {
		var ret []ServiceArticle
		return ret
	}
	return o.Articles
}

// GetArticlesOk returns a tuple with the Articles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerFeedArticlesResponse) GetArticlesOk() ([]ServiceArticle, bool) {
	if o == nil || IsNil(o.Articles) {
		return nil, false
	}
	return o.Articles, true
}

// HasArticles returns a boolean if a field has been set.
func (o *ServerFeedArticlesResponse) HasArticles() bool {
	if o != nil && !IsNil(o.Articles) {
		return true
	}

	return false
}

// SetArticles gets a reference to the given []ServiceArticle and assigns it to the Articles field.
func (o *ServerFeedArticlesResponse) SetArticles(v []ServiceArticle) {
	o.Articles = v
}

func (o ServerFeedArticlesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerFeedArticlesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Articles) {
		toSerialize["articles"] = o.Articles
	}
	return toSerialize, nil
}

type NullableServerFeedArticlesResponse struct {
	value *ServerFeedArticlesResponse
	isSet bool
}

func (v NullableServerFeedArticlesResponse) Get() *ServerFeedArticlesResponse {
	return v.value
}

func (v *NullableServerFeedArticlesResponse) Set(val *ServerFeedArticlesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableServerFeedArticlesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableServerFeedArticlesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerFeedArticlesResponse(val *ServerFeedArticlesResponse) *NullableServerFeedArticlesResponse {
	return &NullableServerFeedArticlesResponse{value: val, isSet: true}
}

func (v NullableServerFeedArticlesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerFeedArticlesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
