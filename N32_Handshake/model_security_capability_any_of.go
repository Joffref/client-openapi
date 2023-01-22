/*
N32 Handshake API

N32-c Handshake Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package N32_Handshake

import (
	"encoding/json"
	"fmt"
)

// SecurityCapabilityAnyOf the model 'SecurityCapabilityAnyOf'
type SecurityCapabilityAnyOf string

// List of SecurityCapability_anyOf
const (
	TLS SecurityCapabilityAnyOf = "TLS"
	PRINS SecurityCapabilityAnyOf = "PRINS"
	NONE SecurityCapabilityAnyOf = "NONE"
)

// All allowed values of SecurityCapabilityAnyOf enum
var AllowedSecurityCapabilityAnyOfEnumValues = []SecurityCapabilityAnyOf{
	"TLS",
	"PRINS",
	"NONE",
}

func (v *SecurityCapabilityAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SecurityCapabilityAnyOf(value)
	for _, existing := range AllowedSecurityCapabilityAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SecurityCapabilityAnyOf", value)
}

// NewSecurityCapabilityAnyOfFromValue returns a pointer to a valid SecurityCapabilityAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSecurityCapabilityAnyOfFromValue(v string) (*SecurityCapabilityAnyOf, error) {
	ev := SecurityCapabilityAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SecurityCapabilityAnyOf: valid values are %v", v, AllowedSecurityCapabilityAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SecurityCapabilityAnyOf) IsValid() bool {
	for _, existing := range AllowedSecurityCapabilityAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityCapability_anyOf value
func (v SecurityCapabilityAnyOf) Ptr() *SecurityCapabilityAnyOf {
	return &v
}

type NullableSecurityCapabilityAnyOf struct {
	value *SecurityCapabilityAnyOf
	isSet bool
}

func (v NullableSecurityCapabilityAnyOf) Get() *SecurityCapabilityAnyOf {
	return v.value
}

func (v *NullableSecurityCapabilityAnyOf) Set(val *SecurityCapabilityAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityCapabilityAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityCapabilityAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityCapabilityAnyOf(val *SecurityCapabilityAnyOf) *NullableSecurityCapabilityAnyOf {
	return &NullableSecurityCapabilityAnyOf{value: val, isSet: true}
}

func (v NullableSecurityCapabilityAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityCapabilityAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
