/*
Naf_ProSe API

Naf_ProSe Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Naf_ProSe

import (
	"encoding/json"
)

// RestrictedCodeSuffixRange Contains a range of the Restricted Code Suffixes which are consecutive.
type RestrictedCodeSuffixRange struct {
	// Contains the ProSe Restricted Code Suffix.
	BeginningSuffix string `json:"beginningSuffix"`
	// Contains the ProSe Restricted Code Suffix.
	EndingSuffix string `json:"endingSuffix"`
}

// NewRestrictedCodeSuffixRange instantiates a new RestrictedCodeSuffixRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestrictedCodeSuffixRange(beginningSuffix string, endingSuffix string) *RestrictedCodeSuffixRange {
	this := RestrictedCodeSuffixRange{}
	this.BeginningSuffix = beginningSuffix
	this.EndingSuffix = endingSuffix
	return &this
}

// NewRestrictedCodeSuffixRangeWithDefaults instantiates a new RestrictedCodeSuffixRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestrictedCodeSuffixRangeWithDefaults() *RestrictedCodeSuffixRange {
	this := RestrictedCodeSuffixRange{}
	return &this
}

// GetBeginningSuffix returns the BeginningSuffix field value
func (o *RestrictedCodeSuffixRange) GetBeginningSuffix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BeginningSuffix
}

// GetBeginningSuffixOk returns a tuple with the BeginningSuffix field value
// and a boolean to check if the value has been set.
func (o *RestrictedCodeSuffixRange) GetBeginningSuffixOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.BeginningSuffix, true
}

// SetBeginningSuffix sets field value
func (o *RestrictedCodeSuffixRange) SetBeginningSuffix(v string) {
	o.BeginningSuffix = v
}

// GetEndingSuffix returns the EndingSuffix field value
func (o *RestrictedCodeSuffixRange) GetEndingSuffix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndingSuffix
}

// GetEndingSuffixOk returns a tuple with the EndingSuffix field value
// and a boolean to check if the value has been set.
func (o *RestrictedCodeSuffixRange) GetEndingSuffixOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.EndingSuffix, true
}

// SetEndingSuffix sets field value
func (o *RestrictedCodeSuffixRange) SetEndingSuffix(v string) {
	o.EndingSuffix = v
}

func (o RestrictedCodeSuffixRange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["beginningSuffix"] = o.BeginningSuffix
	}
	if true {
		toSerialize["endingSuffix"] = o.EndingSuffix
	}
	return json.Marshal(toSerialize)
}

type NullableRestrictedCodeSuffixRange struct {
	value *RestrictedCodeSuffixRange
	isSet bool
}

func (v NullableRestrictedCodeSuffixRange) Get() *RestrictedCodeSuffixRange {
	return v.value
}

func (v *NullableRestrictedCodeSuffixRange) Set(val *RestrictedCodeSuffixRange) {
	v.value = val
	v.isSet = true
}

func (v NullableRestrictedCodeSuffixRange) IsSet() bool {
	return v.isSet
}

func (v *NullableRestrictedCodeSuffixRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestrictedCodeSuffixRange(val *RestrictedCodeSuffixRange) *NullableRestrictedCodeSuffixRange {
	return &NullableRestrictedCodeSuffixRange{value: val, isSet: true}
}

func (v NullableRestrictedCodeSuffixRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestrictedCodeSuffixRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

