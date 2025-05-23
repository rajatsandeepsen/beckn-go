/*
Beckn Protocol Core

Beckn Core Transaction API specification

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transaction

import (
	"encoding/json"
)

// checks if the XInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &XInput{}

// XInput Contains any additional or extended inputs required to confirm an order. This is typically a Form Input. Sometimes, selection of catalog elements is not enough for the BPP to confirm an order. For example, to confirm a flight ticket, the airline requires details of the passengers along with information on baggage, identity, in addition to the class of ticket. Similarly, a logistics company may require details on the nature of shipment in order to confirm the shipping. A recruiting firm may require additional details on the applicant in order to confirm a job application. For all such purposes, the BPP can choose to send this object attached to any object in the catalog that is required to be sent while placing the order. This object can typically be sent at an item level or at the order level. The item level XInput will override the Order level XInput as it indicates a special requirement of information for that particular item. Hence the BAP must render a separate form for the Item and another form at the Order level before confirmation.
type XInput struct {
	Form *Form `json:"form,omitempty"`
	// Indicates whether the form data is mandatorily required by the BPP to confirm the order.
	Required *bool `json:"required,omitempty"`
}

// NewXInput instantiates a new XInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXInput() *XInput {
	this := XInput{}
	return &this
}

// NewXInputWithDefaults instantiates a new XInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXInputWithDefaults() *XInput {
	this := XInput{}
	return &this
}

// GetForm returns the Form field value if set, zero value otherwise.
func (o *XInput) GetForm() Form {
	if o == nil || IsNil(o.Form) {
		var ret Form
		return ret
	}
	return *o.Form
}

// GetFormOk returns a tuple with the Form field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XInput) GetFormOk() (*Form, bool) {
	if o == nil || IsNil(o.Form) {
		return nil, false
	}
	return o.Form, true
}

// HasForm returns a boolean if a field has been set.
func (o *XInput) HasForm() bool {
	if o != nil && !IsNil(o.Form) {
		return true
	}

	return false
}

// SetForm gets a reference to the given Form and assigns it to the Form field.
func (o *XInput) SetForm(v Form) {
	o.Form = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *XInput) GetRequired() bool {
	if o == nil || IsNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XInput) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *XInput) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *XInput) SetRequired(v bool) {
	o.Required = &v
}

func (o XInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o XInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Form) {
		toSerialize["form"] = o.Form
	}
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	return toSerialize, nil
}

type NullableXInput struct {
	value *XInput
	isSet bool
}

func (v NullableXInput) Get() *XInput {
	return v.value
}

func (v *NullableXInput) Set(val *XInput) {
	v.value = val
	v.isSet = true
}

func (v NullableXInput) IsSet() bool {
	return v.isSet
}

func (v *NullableXInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXInput(val *XInput) *NullableXInput {
	return &NullableXInput{value: val, isSet: true}
}

func (v NullableXInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


