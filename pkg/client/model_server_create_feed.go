/*
Feed API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ServerCreateFeed type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerCreateFeed{}

// ServerCreateFeed struct for ServerCreateFeed
type ServerCreateFeed struct {
	Name *string `json:"name,omitempty"`
	Url string `json:"url"`
}

type _ServerCreateFeed ServerCreateFeed

// NewServerCreateFeed instantiates a new ServerCreateFeed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerCreateFeed(url string) *ServerCreateFeed {
	this := ServerCreateFeed{}
	this.Url = url
	return &this
}

// NewServerCreateFeedWithDefaults instantiates a new ServerCreateFeed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerCreateFeedWithDefaults() *ServerCreateFeed {
	this := ServerCreateFeed{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServerCreateFeed) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerCreateFeed) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServerCreateFeed) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServerCreateFeed) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value
func (o *ServerCreateFeed) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ServerCreateFeed) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ServerCreateFeed) SetUrl(v string) {
	o.Url = v
}

func (o ServerCreateFeed) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerCreateFeed) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

func (o *ServerCreateFeed) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"url",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varServerCreateFeed := _ServerCreateFeed{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServerCreateFeed)

	if err != nil {
		return err
	}

	*o = ServerCreateFeed(varServerCreateFeed)

	return err
}

type NullableServerCreateFeed struct {
	value *ServerCreateFeed
	isSet bool
}

func (v NullableServerCreateFeed) Get() *ServerCreateFeed {
	return v.value
}

func (v *NullableServerCreateFeed) Set(val *ServerCreateFeed) {
	v.value = val
	v.isSet = true
}

func (v NullableServerCreateFeed) IsSet() bool {
	return v.isSet
}

func (v *NullableServerCreateFeed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerCreateFeed(val *ServerCreateFeed) *NullableServerCreateFeed {
	return &NullableServerCreateFeed{value: val, isSet: true}
}

func (v NullableServerCreateFeed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerCreateFeed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


