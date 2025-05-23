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

// checks if the OnSearchMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OnSearchMessage{}

// OnSearchMessage struct for OnSearchMessage
type OnSearchMessage struct {
	Catalog Catalog `json:"catalog"`
}

type _OnSearchMessage OnSearchMessage

// NewOnSearchMessage instantiates a new OnSearchMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnSearchMessage(catalog Catalog) *OnSearchMessage {
	this := OnSearchMessage{}
	this.Catalog = catalog
	return &this
}

// NewOnSearchMessageWithDefaults instantiates a new OnSearchMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnSearchMessageWithDefaults() *OnSearchMessage {
	this := OnSearchMessage{}
	return &this
}

// GetCatalog returns the Catalog field value
func (o *OnSearchMessage) GetCatalog() Catalog {
	if o == nil {
		var ret Catalog
		return ret
	}

	return o.Catalog
}

// GetCatalogOk returns a tuple with the Catalog field value
// and a boolean to check if the value has been set.
func (o *OnSearchMessage) GetCatalogOk() (*Catalog, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Catalog, true
}

// SetCatalog sets field value
func (o *OnSearchMessage) SetCatalog(v Catalog) {
	o.Catalog = v
}

func (o OnSearchMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OnSearchMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["catalog"] = o.Catalog
	return toSerialize, nil
}

func (o *OnSearchMessage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"catalog",
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

	varOnSearchMessage := _OnSearchMessage{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOnSearchMessage)

	if err != nil {
		return err
	}

	*o = OnSearchMessage(varOnSearchMessage)

	return err
}

type NullableOnSearchMessage struct {
	value *OnSearchMessage
	isSet bool
}

func (v NullableOnSearchMessage) Get() *OnSearchMessage {
	return v.value
}

func (v *NullableOnSearchMessage) Set(val *OnSearchMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableOnSearchMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableOnSearchMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnSearchMessage(val *OnSearchMessage) *NullableOnSearchMessage {
	return &NullableOnSearchMessage{value: val, isSet: true}
}

func (v NullableOnSearchMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnSearchMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


