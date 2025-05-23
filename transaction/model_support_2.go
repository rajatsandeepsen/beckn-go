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

// checks if the Support2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Support2{}

// Support2 struct for Support2
type Support2 struct {
	Context Support2Context `json:"context"`
	Message Support2Message `json:"message"`
}

type _Support2 Support2

// NewSupport2 instantiates a new Support2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupport2(context Support2Context, message Support2Message) *Support2 {
	this := Support2{}
	this.Context = context
	this.Message = message
	return &this
}

// NewSupport2WithDefaults instantiates a new Support2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupport2WithDefaults() *Support2 {
	this := Support2{}
	return &this
}

// GetContext returns the Context field value
func (o *Support2) GetContext() Support2Context {
	if o == nil {
		var ret Support2Context
		return ret
	}

	return o.Context
}

// GetContextOk returns a tuple with the Context field value
// and a boolean to check if the value has been set.
func (o *Support2) GetContextOk() (*Support2Context, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Context, true
}

// SetContext sets field value
func (o *Support2) SetContext(v Support2Context) {
	o.Context = v
}

// GetMessage returns the Message field value
func (o *Support2) GetMessage() Support2Message {
	if o == nil {
		var ret Support2Message
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *Support2) GetMessageOk() (*Support2Message, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *Support2) SetMessage(v Support2Message) {
	o.Message = v
}

func (o Support2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Support2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["context"] = o.Context
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *Support2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"context",
		"message",
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

	varSupport2 := _Support2{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSupport2)

	if err != nil {
		return err
	}

	*o = Support2(varSupport2)

	return err
}

type NullableSupport2 struct {
	value *Support2
	isSet bool
}

func (v NullableSupport2) Get() *Support2 {
	return v.value
}

func (v *NullableSupport2) Set(val *Support2) {
	v.value = val
	v.isSet = true
}

func (v NullableSupport2) IsSet() bool {
	return v.isSet
}

func (v *NullableSupport2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupport2(val *Support2) *NullableSupport2 {
	return &NullableSupport2{value: val, isSet: true}
}

func (v NullableSupport2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupport2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


