/*
Beckn Protocol Core

Beckn Core Transaction API specification

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transaction

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the OnTrackContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OnTrackContext{}

// OnTrackContext struct for OnTrackContext
type OnTrackContext struct {
	// Domain code that is relevant to this transaction context
	Domain interface{} `json:"domain,omitempty"`
	// The location where the transaction is intended to be fulfilled.
	Location *Location `json:"location,omitempty"`
	Action string `json:"action"`
	// Version of transaction protocol being used by the sender.
	Version *string `json:"version,omitempty"`
	// Subscriber ID of the BAP
	BapId *string `json:"bap_id,omitempty"`
	// Subscriber URL of the BAP for accepting callbacks from BPPs.
	BapUri *string `json:"bap_uri,omitempty"`
	// Subscriber ID of the BPP
	BppId *string `json:"bpp_id,omitempty"`
	// Subscriber URL of the BPP for accepting calls from BAPs.
	BppUri *string `json:"bpp_uri,omitempty"`
	// This is a unique value which persists across all API calls from `search` through `confirm`. This is done to indicate an active user session across multiple requests. The BPPs can use this value to push personalized recommendations, and dynamic offerings related to an ongoing transaction despite being unaware of the user active on the BAP.
	TransactionId *string `json:"transaction_id,omitempty"`
	// This is a unique value which persists during a request / callback cycle. Since beckn protocol APIs are asynchronous, BAPs need a common value to match an incoming callback from a BPP to an earlier call. This value can also be used to ignore duplicate messages coming from the BPP. It is recommended to generate a fresh message_id for every new interaction. When sending unsolicited callbacks, BPPs must generate a new message_id.
	MessageId *string `json:"message_id,omitempty"`
	// Time of request generation in RFC3339 format
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// The encryption public key of the sender
	Key *string `json:"key,omitempty"`
	// The duration in ISO8601 format after timestamp for which this message holds valid
	Ttl *string `json:"ttl,omitempty"`
}

type _OnTrackContext OnTrackContext

// NewOnTrackContext instantiates a new OnTrackContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnTrackContext(action string) *OnTrackContext {
	this := OnTrackContext{}
	this.Action = action
	return &this
}

// NewOnTrackContextWithDefaults instantiates a new OnTrackContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnTrackContextWithDefaults() *OnTrackContext {
	this := OnTrackContext{}
	return &this
}

// GetDomain returns the Domain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnTrackContext) GetDomain() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnTrackContext) GetDomainOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return &o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *OnTrackContext) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given interface{} and assigns it to the Domain field.
func (o *OnTrackContext) SetDomain(v interface{}) {
	o.Domain = v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *OnTrackContext) GetLocation() Location {
	if o == nil || IsNil(o.Location) {
		var ret Location
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnTrackContext) GetLocationOk() (*Location, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *OnTrackContext) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given Location and assigns it to the Location field.
func (o *OnTrackContext) SetLocation(v Location) {
	o.Location = &v
}

// GetAction returns the Action field value
func (o *OnTrackContext) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *OnTrackContext) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *OnTrackContext) SetAction(v string) {
	o.Action = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *OnTrackContext) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnTrackContext) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *OnTrackContext) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *OnTrackContext) SetVersion(v string) {
	o.Version = &v
}

// GetBapId returns the BapId field value if set, zero value otherwise.
func (o *OnTrackContext) GetBapId() string {
	if o == nil || IsNil(o.BapId) {
		var ret string
		return ret
	}
	return *o.BapId
}

// GetBapIdOk returns a tuple with the BapId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnTrackContext) GetBapIdOk() (*string, bool) {
	if o == nil || IsNil(o.BapId) {
		return nil, false
	}
	return o.BapId, true
}

// HasBapId returns a boolean if a field has been set.
func (o *OnTrackContext) HasBapId() bool {
	if o != nil && !IsNil(o.BapId) {
		return true
	}

	return false
}

// SetBapId gets a reference to the given string and assigns it to the BapId field.
func (o *OnTrackContext) SetBapId(v string) {
	o.BapId = &v
}

// GetBapUri returns the BapUri field value if set, zero value otherwise.
func (o *OnTrackContext) GetBapUri() string {
	if o == nil || IsNil(o.BapUri) {
		var ret string
		return ret
	}
	return *o.BapUri
}

// GetBapUriOk returns a tuple with the BapUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnTrackContext) GetBapUriOk() (*string, bool) {
	if o == nil || IsNil(o.BapUri) {
		return nil, false
	}
	return o.BapUri, true
}

// HasBapUri returns a boolean if a field has been set.
func (o *OnTrackContext) HasBapUri() bool {
	if o != nil && !IsNil(o.BapUri) {
		return true
	}

	return false
}

// SetBapUri gets a reference to the given string and assigns it to the BapUri field.
func (o *OnTrackContext) SetBapUri(v string) {
	o.BapUri = &v
}

// GetBppId returns the BppId field value if set, zero value otherwise.
func (o *OnTrackContext) GetBppId() string {
	if o == nil || IsNil(o.BppId) {
		var ret string
		return ret
	}
	return *o.BppId
}

// GetBppIdOk returns a tuple with the BppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnTrackContext) GetBppIdOk() (*string, bool) {
	if o == nil || IsNil(o.BppId) {
		return nil, false
	}
	return o.BppId, true
}

// HasBppId returns a boolean if a field has been set.
func (o *OnTrackContext) HasBppId() bool {
	if o != nil && !IsNil(o.BppId) {
		return true
	}

	return false
}

// SetBppId gets a reference to the given string and assigns it to the BppId field.
func (o *OnTrackContext) SetBppId(v string) {
	o.BppId = &v
}

// GetBppUri returns the BppUri field value if set, zero value otherwise.
func (o *OnTrackContext) GetBppUri() string {
	if o == nil || IsNil(o.BppUri) {
		var ret string
		return ret
	}
	return *o.BppUri
}

// GetBppUriOk returns a tuple with the BppUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnTrackContext) GetBppUriOk() (*string, bool) {
	if o == nil || IsNil(o.BppUri) {
		return nil, false
	}
	return o.BppUri, true
}

// HasBppUri returns a boolean if a field has been set.
func (o *OnTrackContext) HasBppUri() bool {
	if o != nil && !IsNil(o.BppUri) {
		return true
	}

	return false
}

// SetBppUri gets a reference to the given string and assigns it to the BppUri field.
func (o *OnTrackContext) SetBppUri(v string) {
	o.BppUri = &v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *OnTrackContext) GetTransactionId() string {
	if o == nil || IsNil(o.TransactionId) {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnTrackContext) GetTransactionIdOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionId) {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *OnTrackContext) HasTransactionId() bool {
	if o != nil && !IsNil(o.TransactionId) {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *OnTrackContext) SetTransactionId(v string) {
	o.TransactionId = &v
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *OnTrackContext) GetMessageId() string {
	if o == nil || IsNil(o.MessageId) {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnTrackContext) GetMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessageId) {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *OnTrackContext) HasMessageId() bool {
	if o != nil && !IsNil(o.MessageId) {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *OnTrackContext) SetMessageId(v string) {
	o.MessageId = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *OnTrackContext) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnTrackContext) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *OnTrackContext) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *OnTrackContext) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *OnTrackContext) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnTrackContext) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *OnTrackContext) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *OnTrackContext) SetKey(v string) {
	o.Key = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *OnTrackContext) GetTtl() string {
	if o == nil || IsNil(o.Ttl) {
		var ret string
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnTrackContext) GetTtlOk() (*string, bool) {
	if o == nil || IsNil(o.Ttl) {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *OnTrackContext) HasTtl() bool {
	if o != nil && !IsNil(o.Ttl) {
		return true
	}

	return false
}

// SetTtl gets a reference to the given string and assigns it to the Ttl field.
func (o *OnTrackContext) SetTtl(v string) {
	o.Ttl = &v
}

func (o OnTrackContext) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OnTrackContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	toSerialize["action"] = o.Action
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.BapId) {
		toSerialize["bap_id"] = o.BapId
	}
	if !IsNil(o.BapUri) {
		toSerialize["bap_uri"] = o.BapUri
	}
	if !IsNil(o.BppId) {
		toSerialize["bpp_id"] = o.BppId
	}
	if !IsNil(o.BppUri) {
		toSerialize["bpp_uri"] = o.BppUri
	}
	if !IsNil(o.TransactionId) {
		toSerialize["transaction_id"] = o.TransactionId
	}
	if !IsNil(o.MessageId) {
		toSerialize["message_id"] = o.MessageId
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Ttl) {
		toSerialize["ttl"] = o.Ttl
	}
	return toSerialize, nil
}

func (o *OnTrackContext) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action",
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

	varOnTrackContext := _OnTrackContext{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOnTrackContext)

	if err != nil {
		return err
	}

	*o = OnTrackContext(varOnTrackContext)

	return err
}

type NullableOnTrackContext struct {
	value *OnTrackContext
	isSet bool
}

func (v NullableOnTrackContext) Get() *OnTrackContext {
	return v.value
}

func (v *NullableOnTrackContext) Set(val *OnTrackContext) {
	v.value = val
	v.isSet = true
}

func (v NullableOnTrackContext) IsSet() bool {
	return v.isSet
}

func (v *NullableOnTrackContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnTrackContext(val *OnTrackContext) *NullableOnTrackContext {
	return &NullableOnTrackContext{value: val, isSet: true}
}

func (v NullableOnTrackContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnTrackContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


