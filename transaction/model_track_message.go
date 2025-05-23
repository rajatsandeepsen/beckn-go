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

// checks if the TrackMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrackMessage{}

// TrackMessage struct for TrackMessage
type TrackMessage struct {
	// Human-readable ID of the order. This is generated at the BPP layer. The BPP can either generate order id within its system or forward the order ID created at the provider level.
	OrderId string `json:"order_id"`
	CallbackUrl *string `json:"callback_url,omitempty"`
}

type _TrackMessage TrackMessage

// NewTrackMessage instantiates a new TrackMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrackMessage(orderId string) *TrackMessage {
	this := TrackMessage{}
	this.OrderId = orderId
	return &this
}

// NewTrackMessageWithDefaults instantiates a new TrackMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrackMessageWithDefaults() *TrackMessage {
	this := TrackMessage{}
	return &this
}

// GetOrderId returns the OrderId field value
func (o *TrackMessage) GetOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value
// and a boolean to check if the value has been set.
func (o *TrackMessage) GetOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderId, true
}

// SetOrderId sets field value
func (o *TrackMessage) SetOrderId(v string) {
	o.OrderId = v
}

// GetCallbackUrl returns the CallbackUrl field value if set, zero value otherwise.
func (o *TrackMessage) GetCallbackUrl() string {
	if o == nil || IsNil(o.CallbackUrl) {
		var ret string
		return ret
	}
	return *o.CallbackUrl
}

// GetCallbackUrlOk returns a tuple with the CallbackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackMessage) GetCallbackUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CallbackUrl) {
		return nil, false
	}
	return o.CallbackUrl, true
}

// HasCallbackUrl returns a boolean if a field has been set.
func (o *TrackMessage) HasCallbackUrl() bool {
	if o != nil && !IsNil(o.CallbackUrl) {
		return true
	}

	return false
}

// SetCallbackUrl gets a reference to the given string and assigns it to the CallbackUrl field.
func (o *TrackMessage) SetCallbackUrl(v string) {
	o.CallbackUrl = &v
}

func (o TrackMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrackMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["order_id"] = o.OrderId
	if !IsNil(o.CallbackUrl) {
		toSerialize["callback_url"] = o.CallbackUrl
	}
	return toSerialize, nil
}

func (o *TrackMessage) UnmarshalJSON(data []byte) (err error) {
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

	varTrackMessage := _TrackMessage{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTrackMessage)

	if err != nil {
		return err
	}

	*o = TrackMessage(varTrackMessage)

	return err
}

type NullableTrackMessage struct {
	value *TrackMessage
	isSet bool
}

func (v NullableTrackMessage) Get() *TrackMessage {
	return v.value
}

func (v *NullableTrackMessage) Set(val *TrackMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableTrackMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableTrackMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrackMessage(val *TrackMessage) *NullableTrackMessage {
	return &NullableTrackMessage{value: val, isSet: true}
}

func (v NullableTrackMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrackMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


