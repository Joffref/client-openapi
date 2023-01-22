/*
Eees_EASDiscovery

API for EAS Discovery. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Eees_EASDiscovery

import (
	"encoding/json"
)

// EasDiscoveryResp ECS discovery response.
type EasDiscoveryResp struct {
	// List of EAS discovery information.
	DiscoveredEas []DiscoveredEas `json:"discoveredEas"`
}

// NewEasDiscoveryResp instantiates a new EasDiscoveryResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEasDiscoveryResp(discoveredEas []DiscoveredEas) *EasDiscoveryResp {
	this := EasDiscoveryResp{}
	this.DiscoveredEas = discoveredEas
	return &this
}

// NewEasDiscoveryRespWithDefaults instantiates a new EasDiscoveryResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEasDiscoveryRespWithDefaults() *EasDiscoveryResp {
	this := EasDiscoveryResp{}
	return &this
}

// GetDiscoveredEas returns the DiscoveredEas field value
func (o *EasDiscoveryResp) GetDiscoveredEas() []DiscoveredEas {
	if o == nil {
		var ret []DiscoveredEas
		return ret
	}

	return o.DiscoveredEas
}

// GetDiscoveredEasOk returns a tuple with the DiscoveredEas field value
// and a boolean to check if the value has been set.
func (o *EasDiscoveryResp) GetDiscoveredEasOk() ([]DiscoveredEas, bool) {
	if o == nil {
    return nil, false
	}
	return o.DiscoveredEas, true
}

// SetDiscoveredEas sets field value
func (o *EasDiscoveryResp) SetDiscoveredEas(v []DiscoveredEas) {
	o.DiscoveredEas = v
}

func (o EasDiscoveryResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["discoveredEas"] = o.DiscoveredEas
	}
	return json.Marshal(toSerialize)
}

type NullableEasDiscoveryResp struct {
	value *EasDiscoveryResp
	isSet bool
}

func (v NullableEasDiscoveryResp) Get() *EasDiscoveryResp {
	return v.value
}

func (v *NullableEasDiscoveryResp) Set(val *EasDiscoveryResp) {
	v.value = val
	v.isSet = true
}

func (v NullableEasDiscoveryResp) IsSet() bool {
	return v.isSet
}

func (v *NullableEasDiscoveryResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEasDiscoveryResp(val *EasDiscoveryResp) *NullableEasDiscoveryResp {
	return &NullableEasDiscoveryResp{value: val, isSet: true}
}

func (v NullableEasDiscoveryResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEasDiscoveryResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

