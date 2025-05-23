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

// checks if the Contact type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Contact{}

// Contact Describes the contact information of an entity
type Contact struct {
	Phone *string `json:"phone,omitempty"`
	Email *string `json:"email,omitempty"`
	// A Jcard object as per draft-ietf-jcardcal-jcard-03 specification
	Jcard map[string]interface{} `json:"jcard,omitempty"`
}

// NewContact instantiates a new Contact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContact() *Contact {
	this := Contact{}
	return &this
}

// NewContactWithDefaults instantiates a new Contact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactWithDefaults() *Contact {
	this := Contact{}
	return &this
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *Contact) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contact) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *Contact) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *Contact) SetPhone(v string) {
	o.Phone = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *Contact) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contact) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *Contact) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *Contact) SetEmail(v string) {
	o.Email = &v
}

// GetJcard returns the Jcard field value if set, zero value otherwise.
func (o *Contact) GetJcard() map[string]interface{} {
	if o == nil || IsNil(o.Jcard) {
		var ret map[string]interface{}
		return ret
	}
	return o.Jcard
}

// GetJcardOk returns a tuple with the Jcard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contact) GetJcardOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Jcard) {
		return map[string]interface{}{}, false
	}
	return o.Jcard, true
}

// HasJcard returns a boolean if a field has been set.
func (o *Contact) HasJcard() bool {
	if o != nil && !IsNil(o.Jcard) {
		return true
	}

	return false
}

// SetJcard gets a reference to the given map[string]interface{} and assigns it to the Jcard field.
func (o *Contact) SetJcard(v map[string]interface{}) {
	o.Jcard = v
}

func (o Contact) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Contact) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Jcard) {
		toSerialize["jcard"] = o.Jcard
	}
	return toSerialize, nil
}

type NullableContact struct {
	value *Contact
	isSet bool
}

func (v NullableContact) Get() *Contact {
	return v.value
}

func (v *NullableContact) Set(val *Contact) {
	v.value = val
	v.isSet = true
}

func (v NullableContact) IsSet() bool {
	return v.isSet
}

func (v *NullableContact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContact(val *Contact) *NullableContact {
	return &NullableContact{value: val, isSet: true}
}

func (v NullableContact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


