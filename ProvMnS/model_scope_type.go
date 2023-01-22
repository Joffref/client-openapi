/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ProvMnS

import (
	"encoding/json"
	"fmt"
)

// ScopeType the model 'ScopeType'
type ScopeType string

// List of ScopeType
const (
	ONLY ScopeType = "BASE_ONLY"
	NTH_LEVEL ScopeType = "BASE_NTH_LEVEL"
	SUBTREE ScopeType = "BASE_SUBTREE"
	ALL ScopeType = "BASE_ALL"
)

// All allowed values of ScopeType enum
var AllowedScopeTypeEnumValues = []ScopeType{
	"BASE_ONLY",
	"BASE_NTH_LEVEL",
	"BASE_SUBTREE",
	"BASE_ALL",
}

func (v *ScopeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ScopeType(value)
	for _, existing := range AllowedScopeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ScopeType", value)
}

// NewScopeTypeFromValue returns a pointer to a valid ScopeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewScopeTypeFromValue(v string) (*ScopeType, error) {
	ev := ScopeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ScopeType: valid values are %v", v, AllowedScopeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ScopeType) IsValid() bool {
	for _, existing := range AllowedScopeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScopeType value
func (v ScopeType) Ptr() *ScopeType {
	return &v
}

type NullableScopeType struct {
	value *ScopeType
	isSet bool
}

func (v NullableScopeType) Get() *ScopeType {
	return v.value
}

func (v *NullableScopeType) Set(val *ScopeType) {
	v.value = val
	v.isSet = true
}

func (v NullableScopeType) IsSet() bool {
	return v.isSet
}

func (v *NullableScopeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScopeType(val *ScopeType) *NullableScopeType {
	return &NullableScopeType{value: val, isSet: true}
}

func (v NullableScopeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScopeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
