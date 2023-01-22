/*
Nudm_UEAU

UDM UE Authentication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudm_UEAU

import (
	"encoding/json"
)

// ResynchronizationInfo struct for ResynchronizationInfo
type ResynchronizationInfo struct {
	Rand string `json:"rand"`
	Auts string `json:"auts"`
}

// NewResynchronizationInfo instantiates a new ResynchronizationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResynchronizationInfo(rand string, auts string) *ResynchronizationInfo {
	this := ResynchronizationInfo{}
	this.Rand = rand
	this.Auts = auts
	return &this
}

// NewResynchronizationInfoWithDefaults instantiates a new ResynchronizationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResynchronizationInfoWithDefaults() *ResynchronizationInfo {
	this := ResynchronizationInfo{}
	return &this
}

// GetRand returns the Rand field value
func (o *ResynchronizationInfo) GetRand() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Rand
}

// GetRandOk returns a tuple with the Rand field value
// and a boolean to check if the value has been set.
func (o *ResynchronizationInfo) GetRandOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Rand, true
}

// SetRand sets field value
func (o *ResynchronizationInfo) SetRand(v string) {
	o.Rand = v
}

// GetAuts returns the Auts field value
func (o *ResynchronizationInfo) GetAuts() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Auts
}

// GetAutsOk returns a tuple with the Auts field value
// and a boolean to check if the value has been set.
func (o *ResynchronizationInfo) GetAutsOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Auts, true
}

// SetAuts sets field value
func (o *ResynchronizationInfo) SetAuts(v string) {
	o.Auts = v
}

func (o ResynchronizationInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["rand"] = o.Rand
	}
	if true {
		toSerialize["auts"] = o.Auts
	}
	return json.Marshal(toSerialize)
}

type NullableResynchronizationInfo struct {
	value *ResynchronizationInfo
	isSet bool
}

func (v NullableResynchronizationInfo) Get() *ResynchronizationInfo {
	return v.value
}

func (v *NullableResynchronizationInfo) Set(val *ResynchronizationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableResynchronizationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableResynchronizationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResynchronizationInfo(val *ResynchronizationInfo) *NullableResynchronizationInfo {
	return &NullableResynchronizationInfo{value: val, isSet: true}
}

func (v NullableResynchronizationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResynchronizationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

