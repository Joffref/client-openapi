/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Subscription_Data

import (
	"encoding/json"
	"fmt"
)

// PpDataType struct for PpDataType
type PpDataType struct {
	PpDataTypeAnyOf *PpDataTypeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PpDataType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into PpDataTypeAnyOf
	err = json.Unmarshal(data, &dst.PpDataTypeAnyOf);
	if err == nil {
		jsonPpDataTypeAnyOf, _ := json.Marshal(dst.PpDataTypeAnyOf)
		if string(jsonPpDataTypeAnyOf) == "{}" { // empty struct
			dst.PpDataTypeAnyOf = nil
		} else {
			return nil // data stored in dst.PpDataTypeAnyOf, return on the first match
		}
	} else {
		dst.PpDataTypeAnyOf = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(PpDataType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PpDataType) MarshalJSON() ([]byte, error) {
	if src.PpDataTypeAnyOf != nil {
		return json.Marshal(&src.PpDataTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePpDataType struct {
	value *PpDataType
	isSet bool
}

func (v NullablePpDataType) Get() *PpDataType {
	return v.value
}

func (v *NullablePpDataType) Set(val *PpDataType) {
	v.value = val
	v.isSet = true
}

func (v NullablePpDataType) IsSet() bool {
	return v.isSet
}

func (v *NullablePpDataType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePpDataType(val *PpDataType) *NullablePpDataType {
	return &NullablePpDataType{value: val, isSet: true}
}

func (v NullablePpDataType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePpDataType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

