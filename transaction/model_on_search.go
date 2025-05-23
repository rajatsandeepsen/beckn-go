/*
Beckn Protocol Core

Beckn Core Transaction API specification

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transaction

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the OnSearch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OnSearch{}

// OnSearch struct for OnSearch
type OnSearch struct {
	Context OnSearchContext `json:"context"`
	Message *OnSearchMessage `json:"message,omitempty"`
	Error *Error `json:"error,omitempty"`
}

type _OnSearch OnSearch

// NewOnSearch instantiates a new OnSearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnSearch(context OnSearchContext) *OnSearch {
	this := OnSearch{}
	this.Context = context
	return &this
}

// NewOnSearchWithDefaults instantiates a new OnSearch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnSearchWithDefaults() *OnSearch {
	this := OnSearch{}
	return &this
}

// GetContext returns the Context field value
func (o *OnSearch) GetContext() OnSearchContext {
	if o == nil {
		var ret OnSearchContext
		return ret
	}

	return o.Context
}

// GetContextOk returns a tuple with the Context field value
// and a boolean to check if the value has been set.
func (o *OnSearch) GetContextOk() (*OnSearchContext, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Context, true
}

// SetContext sets field value
func (o *OnSearch) SetContext(v OnSearchContext) {
	o.Context = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *OnSearch) GetMessage() OnSearchMessage {
	if o == nil || IsNil(o.Message) {
		var ret OnSearchMessage
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnSearch) GetMessageOk() (*OnSearchMessage, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *OnSearch) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given OnSearchMessage and assigns it to the Message field.
func (o *OnSearch) SetMessage(v OnSearchMessage) {
	o.Message = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *OnSearch) GetError() Error {
	if o == nil || IsNil(o.Error) {
		var ret Error
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnSearch) GetErrorOk() (*Error, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *OnSearch) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given Error and assigns it to the Error field.
func (o *OnSearch) SetError(v Error) {
	o.Error = &v
}

func (o OnSearch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OnSearch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["context"] = o.Context
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

func (o *OnSearch) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"context",
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

	varOnSearch := _OnSearch{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOnSearch)

	if err != nil {
		return err
	}

	*o = OnSearch(varOnSearch)

	return err
}

type NullableOnSearch struct {
	value *OnSearch
	isSet bool
}

func (v NullableOnSearch) Get() *OnSearch {
	return v.value
}

func (v *NullableOnSearch) Set(val *OnSearch) {
	v.value = val
	v.isSet = true
}

func (v NullableOnSearch) IsSet() bool {
	return v.isSet
}

func (v *NullableOnSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnSearch(val *OnSearch) *NullableOnSearch {
	return &NullableOnSearch{value: val, isSet: true}
}

func (v NullableOnSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


