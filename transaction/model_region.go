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

// checks if the Region type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Region{}

// Region Describes an arbitrary region of space. The network policy should contain a published list of supported regions by the network.
type Region struct {
	// The number of dimensions that are used to describe any point inside that region. The most common dimensionality of a region is 2, that represents an area on a map. There are regions on the map that can be approximated to one-dimensional regions like roads, railway lines, or shipping lines. 3 dimensional regions are rarer, but are gaining popularity as flying drones are being adopted for various fulfillment services.
	Dimensions *string `json:"dimensions,omitempty"`
	// The type of region. This is used to specify the granularity of the region represented by this object. Various examples of two-dimensional region types are city, country, state, district, and so on. The network policy should contain a list of all possible region types supported by the network.
	Type *string `json:"type,omitempty"`
	// Name of the region as specified on the map where that region exists.
	Name *string `json:"name,omitempty"`
	// A standard code representing the region. This should be interpreted in the same way by all network participants.
	Code *string `json:"code,omitempty"`
	// A string representing the boundary of the region. One-dimensional regions are represented by polylines. Two-dimensional regions are represented by polygons, and three-dimensional regions can represented by polyhedra.
	Boundary *string `json:"boundary,omitempty"`
	// The url to the map of the region. This can be a globally recognized map or the one specified by the network policy.
	MapUrl *string `json:"map_url,omitempty"`
}

// NewRegion instantiates a new Region object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegion() *Region {
	this := Region{}
	return &this
}

// NewRegionWithDefaults instantiates a new Region object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegionWithDefaults() *Region {
	this := Region{}
	return &this
}

// GetDimensions returns the Dimensions field value if set, zero value otherwise.
func (o *Region) GetDimensions() string {
	if o == nil || IsNil(o.Dimensions) {
		var ret string
		return ret
	}
	return *o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Region) GetDimensionsOk() (*string, bool) {
	if o == nil || IsNil(o.Dimensions) {
		return nil, false
	}
	return o.Dimensions, true
}

// HasDimensions returns a boolean if a field has been set.
func (o *Region) HasDimensions() bool {
	if o != nil && !IsNil(o.Dimensions) {
		return true
	}

	return false
}

// SetDimensions gets a reference to the given string and assigns it to the Dimensions field.
func (o *Region) SetDimensions(v string) {
	o.Dimensions = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Region) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Region) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Region) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Region) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Region) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Region) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Region) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Region) SetName(v string) {
	o.Name = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *Region) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Region) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *Region) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *Region) SetCode(v string) {
	o.Code = &v
}

// GetBoundary returns the Boundary field value if set, zero value otherwise.
func (o *Region) GetBoundary() string {
	if o == nil || IsNil(o.Boundary) {
		var ret string
		return ret
	}
	return *o.Boundary
}

// GetBoundaryOk returns a tuple with the Boundary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Region) GetBoundaryOk() (*string, bool) {
	if o == nil || IsNil(o.Boundary) {
		return nil, false
	}
	return o.Boundary, true
}

// HasBoundary returns a boolean if a field has been set.
func (o *Region) HasBoundary() bool {
	if o != nil && !IsNil(o.Boundary) {
		return true
	}

	return false
}

// SetBoundary gets a reference to the given string and assigns it to the Boundary field.
func (o *Region) SetBoundary(v string) {
	o.Boundary = &v
}

// GetMapUrl returns the MapUrl field value if set, zero value otherwise.
func (o *Region) GetMapUrl() string {
	if o == nil || IsNil(o.MapUrl) {
		var ret string
		return ret
	}
	return *o.MapUrl
}

// GetMapUrlOk returns a tuple with the MapUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Region) GetMapUrlOk() (*string, bool) {
	if o == nil || IsNil(o.MapUrl) {
		return nil, false
	}
	return o.MapUrl, true
}

// HasMapUrl returns a boolean if a field has been set.
func (o *Region) HasMapUrl() bool {
	if o != nil && !IsNil(o.MapUrl) {
		return true
	}

	return false
}

// SetMapUrl gets a reference to the given string and assigns it to the MapUrl field.
func (o *Region) SetMapUrl(v string) {
	o.MapUrl = &v
}

func (o Region) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Region) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Dimensions) {
		toSerialize["dimensions"] = o.Dimensions
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Boundary) {
		toSerialize["boundary"] = o.Boundary
	}
	if !IsNil(o.MapUrl) {
		toSerialize["map_url"] = o.MapUrl
	}
	return toSerialize, nil
}

type NullableRegion struct {
	value *Region
	isSet bool
}

func (v NullableRegion) Get() *Region {
	return v.value
}

func (v *NullableRegion) Set(val *Region) {
	v.value = val
	v.isSet = true
}

func (v NullableRegion) IsSet() bool {
	return v.isSet
}

func (v *NullableRegion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegion(val *Region) *NullableRegion {
	return &NullableRegion{value: val, isSet: true}
}

func (v NullableRegion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


