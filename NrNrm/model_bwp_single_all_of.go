/*
NR NRM

OAS 3.0.1 specification of the NR NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package NrNrm

import (
	"encoding/json"
)

// BwpSingleAllOf struct for BwpSingleAllOf
type BwpSingleAllOf struct {
	Attributes *ManagedFunctionAttr `json:"attributes,omitempty"`
}

// NewBwpSingleAllOf instantiates a new BwpSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBwpSingleAllOf() *BwpSingleAllOf {
	this := BwpSingleAllOf{}
	return &this
}

// NewBwpSingleAllOfWithDefaults instantiates a new BwpSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBwpSingleAllOfWithDefaults() *BwpSingleAllOf {
	this := BwpSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *BwpSingleAllOf) GetAttributes() ManagedFunctionAttr {
	if o == nil || isNil(o.Attributes) {
		var ret ManagedFunctionAttr
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BwpSingleAllOf) GetAttributesOk() (*ManagedFunctionAttr, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *BwpSingleAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ManagedFunctionAttr and assigns it to the Attributes field.
func (o *BwpSingleAllOf) SetAttributes(v ManagedFunctionAttr) {
	o.Attributes = &v
}

func (o BwpSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableBwpSingleAllOf struct {
	value *BwpSingleAllOf
	isSet bool
}

func (v NullableBwpSingleAllOf) Get() *BwpSingleAllOf {
	return v.value
}

func (v *NullableBwpSingleAllOf) Set(val *BwpSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBwpSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBwpSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBwpSingleAllOf(val *BwpSingleAllOf) *NullableBwpSingleAllOf {
	return &NullableBwpSingleAllOf{value: val, isSet: true}
}

func (v NullableBwpSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBwpSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

