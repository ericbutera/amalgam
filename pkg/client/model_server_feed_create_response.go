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

// checks if the ServerFeedCreateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerFeedCreateResponse{}

// ServerFeedCreateResponse struct for ServerFeedCreateResponse
type ServerFeedCreateResponse struct {
	Id *string `json:"id,omitempty"`
}

// NewServerFeedCreateResponse instantiates a new ServerFeedCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerFeedCreateResponse() *ServerFeedCreateResponse {
	this := ServerFeedCreateResponse{}
	return &this
}

// NewServerFeedCreateResponseWithDefaults instantiates a new ServerFeedCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerFeedCreateResponseWithDefaults() *ServerFeedCreateResponse {
	this := ServerFeedCreateResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServerFeedCreateResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerFeedCreateResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServerFeedCreateResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ServerFeedCreateResponse) SetId(v string) {
	o.Id = &v
}

func (o ServerFeedCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerFeedCreateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableServerFeedCreateResponse struct {
	value *ServerFeedCreateResponse
	isSet bool
}

func (v NullableServerFeedCreateResponse) Get() *ServerFeedCreateResponse {
	return v.value
}

func (v *NullableServerFeedCreateResponse) Set(val *ServerFeedCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableServerFeedCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableServerFeedCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerFeedCreateResponse(val *ServerFeedCreateResponse) *NullableServerFeedCreateResponse {
	return &NullableServerFeedCreateResponse{value: val, isSet: true}
}

func (v NullableServerFeedCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerFeedCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


