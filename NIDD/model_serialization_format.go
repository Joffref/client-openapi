/*
3gpp-nidd

API for non IP data delivery.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package NIDD

import (
	"encoding/json"
	"fmt"
)

// SerializationFormat Possible values are - CBOR: The CBOR Serialzition format  - JSON: The JSON Serialzition format - XML: The XML Serialzition format 
type SerializationFormat struct {
	SerializationFormatAnyOf *SerializationFormatAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SerializationFormat) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into SerializationFormatAnyOf
	err = json.Unmarshal(data, &dst.SerializationFormatAnyOf);
	if err == nil {
		jsonSerializationFormatAnyOf, _ := json.Marshal(dst.SerializationFormatAnyOf)
		if string(jsonSerializationFormatAnyOf) == "{}" { // empty struct
			dst.SerializationFormatAnyOf = nil
		} else {
			return nil // data stored in dst.SerializationFormatAnyOf, return on the first match
		}
	} else {
		dst.SerializationFormatAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(SerializationFormat)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SerializationFormat) MarshalJSON() ([]byte, error) {
	if src.SerializationFormatAnyOf != nil {
		return json.Marshal(&src.SerializationFormatAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSerializationFormat struct {
	value *SerializationFormat
	isSet bool
}

func (v NullableSerializationFormat) Get() *SerializationFormat {
	return v.value
}

func (v *NullableSerializationFormat) Set(val *SerializationFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableSerializationFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableSerializationFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSerializationFormat(val *SerializationFormat) *NullableSerializationFormat {
	return &NullableSerializationFormat{value: val, isSet: true}
}

func (v NullableSerializationFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSerializationFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

