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

// checks if the Intent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Intent{}

// Intent The intent to buy or avail a product or a service. The BAP can declare the intent of the consumer containing <ul><li>What they want (A product, service, offer)</li><li>Who they want (A seller, service provider, agent etc)</li><li>Where they want it and where they want it from</li><li>When they want it (start and end time of fulfillment</li><li>How they want to pay for it</li></ul><br>This has properties like descriptor,provider,fulfillment,payment,category,offer,item,tags<br>This is typically used by the BAP to send the purpose of the user's search to the BPP. This will be used by the BPP to find products or services it offers that may match the user's intent.<br>For example, in Mobility, the mobility consumer declares a mobility intent. In this case, the mobility consumer declares information that describes various aspects of their journey like,<ul><li>Where would they like to begin their journey (intent.fulfillment.start.location)</li><li>Where would they like to end their journey (intent.fulfillment.end.location)</li><li>When would they like to begin their journey (intent.fulfillment.start.time)</li><li>When would they like to end their journey (intent.fulfillment.end.time)</li><li>Who is the transport service provider they would like to avail services from (intent.provider)</li><li>Who is traveling (This is not recommended in public networks) (intent.fulfillment.customer)</li><li>What kind of fare product would they like to purchase (intent.item)</li><li>What add-on services would they like to avail</li><li>What offers would they like to apply on their booking (intent.offer)</li><li>What category of services would they like to avail (intent.category)</li><li>What additional luggage are they carrying</li><li>How would they like to pay for their journey (intent.payment)</li></ul><br>For example, in health domain, a consumer declares the intent for a lab booking the describes various aspects of their booking like,<ul><li>Where would they like to get their scan/test done (intent.fulfillment.start.location)</li><li>When would they like to get their scan/test done (intent.fulfillment.start.time)</li><li>When would they like to get the results of their test/scan (intent.fulfillment.end.time)</li><li>Who is the service provider they would like to avail services from (intent.provider)</li><li>Who is getting the test/scan (intent.fulfillment.customer)</li><li>What kind of test/scan would they like to purchase (intent.item)</li><li>What category of services would they like to avail (intent.category)</li><li>How would they like to pay for their journey (intent.payment)</li></ul>
type Intent struct {
	// A raw description of the search intent. Free text search strings, raw audio, etc can be sent in this object.
	Descriptor *Descriptor `json:"descriptor,omitempty"`
	// The provider from which the customer wants to place to the order from
	Provider *Provider `json:"provider,omitempty"`
	// Details on how the customer wants their order fulfilled
	Fulfillment *Fulfillment `json:"fulfillment,omitempty"`
	// Details on how the customer wants to pay for the order
	Payment *Payment `json:"payment,omitempty"`
	// Details on the item category
	Category *Category `json:"category,omitempty"`
	// details on the offer the customer wants to avail
	Offer *Offer `json:"offer,omitempty"`
	// Details of the item that the consumer wants to order
	Item *Item `json:"item,omitempty"`
	Tags []TagGroup `json:"tags,omitempty"`
}

// NewIntent instantiates a new Intent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntent() *Intent {
	this := Intent{}
	return &this
}

// NewIntentWithDefaults instantiates a new Intent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntentWithDefaults() *Intent {
	this := Intent{}
	return &this
}

// GetDescriptor returns the Descriptor field value if set, zero value otherwise.
func (o *Intent) GetDescriptor() Descriptor {
	if o == nil || IsNil(o.Descriptor) {
		var ret Descriptor
		return ret
	}
	return *o.Descriptor
}

// GetDescriptorOk returns a tuple with the Descriptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Intent) GetDescriptorOk() (*Descriptor, bool) {
	if o == nil || IsNil(o.Descriptor) {
		return nil, false
	}
	return o.Descriptor, true
}

// HasDescriptor returns a boolean if a field has been set.
func (o *Intent) HasDescriptor() bool {
	if o != nil && !IsNil(o.Descriptor) {
		return true
	}

	return false
}

// SetDescriptor gets a reference to the given Descriptor and assigns it to the Descriptor field.
func (o *Intent) SetDescriptor(v Descriptor) {
	o.Descriptor = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *Intent) GetProvider() Provider {
	if o == nil || IsNil(o.Provider) {
		var ret Provider
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Intent) GetProviderOk() (*Provider, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *Intent) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given Provider and assigns it to the Provider field.
func (o *Intent) SetProvider(v Provider) {
	o.Provider = &v
}

// GetFulfillment returns the Fulfillment field value if set, zero value otherwise.
func (o *Intent) GetFulfillment() Fulfillment {
	if o == nil || IsNil(o.Fulfillment) {
		var ret Fulfillment
		return ret
	}
	return *o.Fulfillment
}

// GetFulfillmentOk returns a tuple with the Fulfillment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Intent) GetFulfillmentOk() (*Fulfillment, bool) {
	if o == nil || IsNil(o.Fulfillment) {
		return nil, false
	}
	return o.Fulfillment, true
}

// HasFulfillment returns a boolean if a field has been set.
func (o *Intent) HasFulfillment() bool {
	if o != nil && !IsNil(o.Fulfillment) {
		return true
	}

	return false
}

// SetFulfillment gets a reference to the given Fulfillment and assigns it to the Fulfillment field.
func (o *Intent) SetFulfillment(v Fulfillment) {
	o.Fulfillment = &v
}

// GetPayment returns the Payment field value if set, zero value otherwise.
func (o *Intent) GetPayment() Payment {
	if o == nil || IsNil(o.Payment) {
		var ret Payment
		return ret
	}
	return *o.Payment
}

// GetPaymentOk returns a tuple with the Payment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Intent) GetPaymentOk() (*Payment, bool) {
	if o == nil || IsNil(o.Payment) {
		return nil, false
	}
	return o.Payment, true
}

// HasPayment returns a boolean if a field has been set.
func (o *Intent) HasPayment() bool {
	if o != nil && !IsNil(o.Payment) {
		return true
	}

	return false
}

// SetPayment gets a reference to the given Payment and assigns it to the Payment field.
func (o *Intent) SetPayment(v Payment) {
	o.Payment = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *Intent) GetCategory() Category {
	if o == nil || IsNil(o.Category) {
		var ret Category
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Intent) GetCategoryOk() (*Category, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *Intent) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given Category and assigns it to the Category field.
func (o *Intent) SetCategory(v Category) {
	o.Category = &v
}

// GetOffer returns the Offer field value if set, zero value otherwise.
func (o *Intent) GetOffer() Offer {
	if o == nil || IsNil(o.Offer) {
		var ret Offer
		return ret
	}
	return *o.Offer
}

// GetOfferOk returns a tuple with the Offer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Intent) GetOfferOk() (*Offer, bool) {
	if o == nil || IsNil(o.Offer) {
		return nil, false
	}
	return o.Offer, true
}

// HasOffer returns a boolean if a field has been set.
func (o *Intent) HasOffer() bool {
	if o != nil && !IsNil(o.Offer) {
		return true
	}

	return false
}

// SetOffer gets a reference to the given Offer and assigns it to the Offer field.
func (o *Intent) SetOffer(v Offer) {
	o.Offer = &v
}

// GetItem returns the Item field value if set, zero value otherwise.
func (o *Intent) GetItem() Item {
	if o == nil || IsNil(o.Item) {
		var ret Item
		return ret
	}
	return *o.Item
}

// GetItemOk returns a tuple with the Item field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Intent) GetItemOk() (*Item, bool) {
	if o == nil || IsNil(o.Item) {
		return nil, false
	}
	return o.Item, true
}

// HasItem returns a boolean if a field has been set.
func (o *Intent) HasItem() bool {
	if o != nil && !IsNil(o.Item) {
		return true
	}

	return false
}

// SetItem gets a reference to the given Item and assigns it to the Item field.
func (o *Intent) SetItem(v Item) {
	o.Item = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Intent) GetTags() []TagGroup {
	if o == nil || IsNil(o.Tags) {
		var ret []TagGroup
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Intent) GetTagsOk() ([]TagGroup, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Intent) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagGroup and assigns it to the Tags field.
func (o *Intent) SetTags(v []TagGroup) {
	o.Tags = v
}

func (o Intent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Intent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Descriptor) {
		toSerialize["descriptor"] = o.Descriptor
	}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	if !IsNil(o.Fulfillment) {
		toSerialize["fulfillment"] = o.Fulfillment
	}
	if !IsNil(o.Payment) {
		toSerialize["payment"] = o.Payment
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.Offer) {
		toSerialize["offer"] = o.Offer
	}
	if !IsNil(o.Item) {
		toSerialize["item"] = o.Item
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableIntent struct {
	value *Intent
	isSet bool
}

func (v NullableIntent) Get() *Intent {
	return v.value
}

func (v *NullableIntent) Set(val *Intent) {
	v.value = val
	v.isSet = true
}

func (v NullableIntent) IsSet() bool {
	return v.isSet
}

func (v *NullableIntent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntent(val *Intent) *NullableIntent {
	return &NullableIntent{value: val, isSet: true}
}

func (v NullableIntent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


