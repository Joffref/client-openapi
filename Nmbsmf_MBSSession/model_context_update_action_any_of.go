/*
Nmbsmf-MBSSession

MB-SMF MBSSession Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nmbsmf_MBSSession

import (
	"encoding/json"
	"fmt"
)

// ContextUpdateActionAnyOf the model 'ContextUpdateActionAnyOf'
type ContextUpdateActionAnyOf string

// List of ContextUpdateAction_anyOf
const (
	START ContextUpdateActionAnyOf = "START"
	TERMINATE ContextUpdateActionAnyOf = "TERMINATE"
)

// All allowed values of ContextUpdateActionAnyOf enum
var AllowedContextUpdateActionAnyOfEnumValues = []ContextUpdateActionAnyOf{
	"START",
	"TERMINATE",
}

func (v *ContextUpdateActionAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ContextUpdateActionAnyOf(value)
	for _, existing := range AllowedContextUpdateActionAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ContextUpdateActionAnyOf", value)
}

// NewContextUpdateActionAnyOfFromValue returns a pointer to a valid ContextUpdateActionAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewContextUpdateActionAnyOfFromValue(v string) (*ContextUpdateActionAnyOf, error) {
	ev := ContextUpdateActionAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ContextUpdateActionAnyOf: valid values are %v", v, AllowedContextUpdateActionAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ContextUpdateActionAnyOf) IsValid() bool {
	for _, existing := range AllowedContextUpdateActionAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ContextUpdateAction_anyOf value
func (v ContextUpdateActionAnyOf) Ptr() *ContextUpdateActionAnyOf {
	return &v
}

type NullableContextUpdateActionAnyOf struct {
	value *ContextUpdateActionAnyOf
	isSet bool
}

func (v NullableContextUpdateActionAnyOf) Get() *ContextUpdateActionAnyOf {
	return v.value
}

func (v *NullableContextUpdateActionAnyOf) Set(val *ContextUpdateActionAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableContextUpdateActionAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableContextUpdateActionAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContextUpdateActionAnyOf(val *ContextUpdateActionAnyOf) *NullableContextUpdateActionAnyOf {
	return &NullableContextUpdateActionAnyOf{value: val, isSet: true}
}

func (v NullableContextUpdateActionAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContextUpdateActionAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
