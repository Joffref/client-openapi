/*
Nudsf_DataRepository

Nudsf Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudsf_DataRepository

import (
	"encoding/json"
	"fmt"
)

// KeyTypeAnyOf the model 'KeyTypeAnyOf'
type KeyTypeAnyOf string

// List of KeyType_anyOf
const (
	UNIQUE_KEY KeyTypeAnyOf = "UNIQUE_KEY"
	SEARCH_KEY KeyTypeAnyOf = "SEARCH_KEY"
	COUNT_KEY KeyTypeAnyOf = "COUNT_KEY"
	SEARCH_AND_COUNT_KEY KeyTypeAnyOf = "SEARCH_AND_COUNT_KEY"
	OTHER_TAG KeyTypeAnyOf = "OTHER_TAG"
)

// All allowed values of KeyTypeAnyOf enum
var AllowedKeyTypeAnyOfEnumValues = []KeyTypeAnyOf{
	"UNIQUE_KEY",
	"SEARCH_KEY",
	"COUNT_KEY",
	"SEARCH_AND_COUNT_KEY",
	"OTHER_TAG",
}

func (v *KeyTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := KeyTypeAnyOf(value)
	for _, existing := range AllowedKeyTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid KeyTypeAnyOf", value)
}

// NewKeyTypeAnyOfFromValue returns a pointer to a valid KeyTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewKeyTypeAnyOfFromValue(v string) (*KeyTypeAnyOf, error) {
	ev := KeyTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for KeyTypeAnyOf: valid values are %v", v, AllowedKeyTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KeyTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedKeyTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to KeyType_anyOf value
func (v KeyTypeAnyOf) Ptr() *KeyTypeAnyOf {
	return &v
}

type NullableKeyTypeAnyOf struct {
	value *KeyTypeAnyOf
	isSet bool
}

func (v NullableKeyTypeAnyOf) Get() *KeyTypeAnyOf {
	return v.value
}

func (v *NullableKeyTypeAnyOf) Set(val *KeyTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyTypeAnyOf(val *KeyTypeAnyOf) *NullableKeyTypeAnyOf {
	return &NullableKeyTypeAnyOf{value: val, isSet: true}
}

func (v NullableKeyTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
