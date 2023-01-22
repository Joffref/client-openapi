/*
Nnwdaf_EventsSubscription

Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnwdaf_EventsSubscription

import (
	"encoding/json"
	"fmt"
)

// MLModelAddr - Addresses of ML model files.
type MLModelAddr struct {
	Interface{} *interface{}
}

// interface{}AsMLModelAddr is a convenience function that returns interface{} wrapped in MLModelAddr
func Interface{}AsMLModelAddr(v *interface{}) MLModelAddr {
	return MLModelAddr{
		Interface{}: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *MLModelAddr) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Interface{}
	err = newStrictDecoder(data).Decode(&dst.Interface{})
	if err == nil {
		jsonInterface{}, _ := json.Marshal(dst.Interface{})
		if string(jsonInterface{}) == "{}" { // empty struct
			dst.Interface{} = nil
		} else {
			match++
		}
	} else {
		dst.Interface{} = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Interface{} = nil

		return fmt.Errorf("data matches more than one schema in oneOf(MLModelAddr)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(MLModelAddr)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MLModelAddr) MarshalJSON() ([]byte, error) {
	if src.Interface{} != nil {
		return json.Marshal(&src.Interface{})
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *MLModelAddr) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Interface{} != nil {
		return obj.Interface{}
	}

	// all schemas are nil
	return nil
}

type NullableMLModelAddr struct {
	value *MLModelAddr
	isSet bool
}

func (v NullableMLModelAddr) Get() *MLModelAddr {
	return v.value
}

func (v *NullableMLModelAddr) Set(val *MLModelAddr) {
	v.value = val
	v.isSet = true
}

func (v NullableMLModelAddr) IsSet() bool {
	return v.isSet
}

func (v *NullableMLModelAddr) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMLModelAddr(val *MLModelAddr) *NullableMLModelAddr {
	return &NullableMLModelAddr{value: val, isSet: true}
}

func (v NullableMLModelAddr) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMLModelAddr) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

