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

// checks if the ServerUpdateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerUpdateResponse{}

// ServerUpdateResponse struct for ServerUpdateResponse
type ServerUpdateResponse struct {
	Id *string `json:"id,omitempty"`
}

// NewServerUpdateResponse instantiates a new ServerUpdateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerUpdateResponse() *ServerUpdateResponse {
	this := ServerUpdateResponse{}
	return &this
}

// NewServerUpdateResponseWithDefaults instantiates a new ServerUpdateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerUpdateResponseWithDefaults() *ServerUpdateResponse {
	this := ServerUpdateResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServerUpdateResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerUpdateResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServerUpdateResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ServerUpdateResponse) SetId(v string) {
	o.Id = &v
}

func (o ServerUpdateResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerUpdateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableServerUpdateResponse struct {
	value *ServerUpdateResponse
	isSet bool
}

func (v NullableServerUpdateResponse) Get() *ServerUpdateResponse {
	return v.value
}

func (v *NullableServerUpdateResponse) Set(val *ServerUpdateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableServerUpdateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableServerUpdateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerUpdateResponse(val *ServerUpdateResponse) *NullableServerUpdateResponse {
	return &NullableServerUpdateResponse{value: val, isSet: true}
}

func (v NullableServerUpdateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerUpdateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
