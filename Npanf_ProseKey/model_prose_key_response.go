/*
Npanf_ProseKey

PAnF ProseKey Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Npanf_ProseKey

import (
	"encoding/json"
)

// ProseKeyResponse Prose Key Response.
type ProseKeyResponse struct {
	// 5G Prose Remote User Key
	Var5gPruk string `json:"5gPruk"`
}

// NewProseKeyResponse instantiates a new ProseKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProseKeyResponse(var5gPruk string) *ProseKeyResponse {
	this := ProseKeyResponse{}
	this.Var5gPruk = var5gPruk
	return &this
}

// NewProseKeyResponseWithDefaults instantiates a new ProseKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProseKeyResponseWithDefaults() *ProseKeyResponse {
	this := ProseKeyResponse{}
	return &this
}

// GetVar5gPruk returns the Var5gPruk field value
func (o *ProseKeyResponse) GetVar5gPruk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Var5gPruk
}

// GetVar5gPrukOk returns a tuple with the Var5gPruk field value
// and a boolean to check if the value has been set.
func (o *ProseKeyResponse) GetVar5gPrukOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Var5gPruk, true
}

// SetVar5gPruk sets field value
func (o *ProseKeyResponse) SetVar5gPruk(v string) {
	o.Var5gPruk = v
}

func (o ProseKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["5gPruk"] = o.Var5gPruk
	}
	return json.Marshal(toSerialize)
}

type NullableProseKeyResponse struct {
	value *ProseKeyResponse
	isSet bool
}

func (v NullableProseKeyResponse) Get() *ProseKeyResponse {
	return v.value
}

func (v *NullableProseKeyResponse) Set(val *ProseKeyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProseKeyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProseKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProseKeyResponse(val *ProseKeyResponse) *NullableProseKeyResponse {
	return &NullableProseKeyResponse{value: val, isSet: true}
}

func (v NullableProseKeyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProseKeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

