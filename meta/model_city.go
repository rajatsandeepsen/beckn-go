/*
Beckn Protocol Meta API

This document contains all the meta API endpoints that are implemented by the network participants. The information returned from these endpoints typically contain cacheable information.

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package meta

import (
	"encoding/json"
)

// checks if the City type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &City{}

// City Describes a city
type City struct {
	// Name of the city
	Name *string `json:"name,omitempty"`
	// City code
	Code *string `json:"code,omitempty"`
}

// NewCity instantiates a new City object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCity() *City {
	this := City{}
	return &this
}

// NewCityWithDefaults instantiates a new City object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCityWithDefaults() *City {
	this := City{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *City) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *City) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *City) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *City) SetName(v string) {
	o.Name = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *City) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *City) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *City) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *City) SetCode(v string) {
	o.Code = &v
}

func (o City) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o City) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	return toSerialize, nil
}

type NullableCity struct {
	value *City
	isSet bool
}

func (v NullableCity) Get() *City {
	return v.value
}

func (v *NullableCity) Set(val *City) {
	v.value = val
	v.isSet = true
}

func (v NullableCity) IsSet() bool {
	return v.isSet
}

func (v *NullableCity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCity(val *City) *NullableCity {
	return &NullableCity{value: val, isSet: true}
}

func (v NullableCity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


