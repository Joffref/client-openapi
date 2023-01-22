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

// ExternalENBFunctionSingleAllOf struct for ExternalENBFunctionSingleAllOf
type ExternalENBFunctionSingleAllOf struct {
	Attributes *ManagedFunctionAttr `json:"attributes,omitempty"`
}

// NewExternalENBFunctionSingleAllOf instantiates a new ExternalENBFunctionSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalENBFunctionSingleAllOf() *ExternalENBFunctionSingleAllOf {
	this := ExternalENBFunctionSingleAllOf{}
	return &this
}

// NewExternalENBFunctionSingleAllOfWithDefaults instantiates a new ExternalENBFunctionSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalENBFunctionSingleAllOfWithDefaults() *ExternalENBFunctionSingleAllOf {
	this := ExternalENBFunctionSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ExternalENBFunctionSingleAllOf) GetAttributes() ManagedFunctionAttr {
	if o == nil || isNil(o.Attributes) {
		var ret ManagedFunctionAttr
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalENBFunctionSingleAllOf) GetAttributesOk() (*ManagedFunctionAttr, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ExternalENBFunctionSingleAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ManagedFunctionAttr and assigns it to the Attributes field.
func (o *ExternalENBFunctionSingleAllOf) SetAttributes(v ManagedFunctionAttr) {
	o.Attributes = &v
}

func (o ExternalENBFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableExternalENBFunctionSingleAllOf struct {
	value *ExternalENBFunctionSingleAllOf
	isSet bool
}

func (v NullableExternalENBFunctionSingleAllOf) Get() *ExternalENBFunctionSingleAllOf {
	return v.value
}

func (v *NullableExternalENBFunctionSingleAllOf) Set(val *ExternalENBFunctionSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalENBFunctionSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalENBFunctionSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalENBFunctionSingleAllOf(val *ExternalENBFunctionSingleAllOf) *NullableExternalENBFunctionSingleAllOf {
	return &NullableExternalENBFunctionSingleAllOf{value: val, isSet: true}
}

func (v NullableExternalENBFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalENBFunctionSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

