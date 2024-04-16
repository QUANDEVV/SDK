/*
proxima-core-engine

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// Tier struct for Tier
type Tier struct {
	Name string `json:"name"`
	Features []Feature `json:"features"`
	Prices []Price `json:"prices"`
}

// NewTier instantiates a new Tier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTier(name string, features []Feature, prices []Price) *Tier {
	this := Tier{}
	this.Name = name
	this.Features = features
	this.Prices = prices
	return &this
}

// NewTierWithDefaults instantiates a new Tier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTierWithDefaults() *Tier {
	this := Tier{}
	return &this
}

// GetName returns the Name field value
func (o *Tier) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Tier) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Tier) SetName(v string) {
	o.Name = v
}

// GetFeatures returns the Features field value
func (o *Tier) GetFeatures() []Feature {
	if o == nil {
		var ret []Feature
		return ret
	}

	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value
// and a boolean to check if the value has been set.
func (o *Tier) GetFeaturesOk() (*[]Feature, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Features, true
}

// SetFeatures sets field value
func (o *Tier) SetFeatures(v []Feature) {
	o.Features = v
}

// GetPrices returns the Prices field value
func (o *Tier) GetPrices() []Price {
	if o == nil {
		var ret []Price
		return ret
	}

	return o.Prices
}

// GetPricesOk returns a tuple with the Prices field value
// and a boolean to check if the value has been set.
func (o *Tier) GetPricesOk() (*[]Price, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Prices, true
}

// SetPrices sets field value
func (o *Tier) SetPrices(v []Price) {
	o.Prices = v
}

func (o Tier) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["features"] = o.Features
	}
	if true {
		toSerialize["prices"] = o.Prices
	}
	return json.Marshal(toSerialize)
}

type NullableTier struct {
	value *Tier
	isSet bool
}

func (v NullableTier) Get() *Tier {
	return v.value
}

func (v *NullableTier) Set(val *Tier) {
	v.value = val
	v.isSet = true
}

func (v NullableTier) IsSet() bool {
	return v.isSet
}

func (v *NullableTier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTier(val *Tier) *NullableTier {
	return &NullableTier{value: val, isSet: true}
}

func (v NullableTier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


