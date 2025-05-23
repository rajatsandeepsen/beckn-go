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

// checks if the StatusMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatusMessage{}

// StatusMessage struct for StatusMessage
type StatusMessage struct {
	// Human-readable ID of the order. This is generated at the BPP layer. The BPP can either generate order id within its system or forward the order ID created at the provider level.
	OrderId string `json:"order_id"`
}

type _StatusMessage StatusMessage

// NewStatusMessage instantiates a new StatusMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusMessage(orderId string) *StatusMessage {
	this := StatusMessage{}
	this.OrderId = orderId
	return &this
}

// NewStatusMessageWithDefaults instantiates a new StatusMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusMessageWithDefaults() *StatusMessage {
	this := StatusMessage{}
	return &this
}

// GetOrderId returns the OrderId field value
func (o *StatusMessage) GetOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value
// and a boolean to check if the value has been set.
func (o *StatusMessage) GetOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderId, true
}

// SetOrderId sets field value
func (o *StatusMessage) SetOrderId(v string) {
	o.OrderId = v
}

func (o StatusMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatusMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["order_id"] = o.OrderId
	return toSerialize, nil
}

func (o *StatusMessage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"order_id",
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

	varStatusMessage := _StatusMessage{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStatusMessage)

	if err != nil {
		return err
	}

	*o = StatusMessage(varStatusMessage)

	return err
}

type NullableStatusMessage struct {
	value *StatusMessage
	isSet bool
}

func (v NullableStatusMessage) Get() *StatusMessage {
	return v.value
}

func (v *NullableStatusMessage) Set(val *StatusMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusMessage(val *StatusMessage) *NullableStatusMessage {
	return &NullableStatusMessage{value: val, isSet: true}
}

func (v NullableStatusMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


