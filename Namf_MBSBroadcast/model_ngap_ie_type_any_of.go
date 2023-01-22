/*
Namf_MBSBroadcast

AMF MBSBroadcast Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Namf_MBSBroadcast

import (
	"encoding/json"
	"fmt"
)

// NgapIeTypeAnyOf the model 'NgapIeTypeAnyOf'
type NgapIeTypeAnyOf string

// List of NgapIeType_anyOf
const (
	REQ NgapIeTypeAnyOf = "MBS_SES_REQ"
	RSP NgapIeTypeAnyOf = "MBS_SES_RSP"
	FAIL NgapIeTypeAnyOf = "MBS_SES_FAIL"
	REL_RSP NgapIeTypeAnyOf = "MBS_SES_REL_RSP"
)

// All allowed values of NgapIeTypeAnyOf enum
var AllowedNgapIeTypeAnyOfEnumValues = []NgapIeTypeAnyOf{
	"MBS_SES_REQ",
	"MBS_SES_RSP",
	"MBS_SES_FAIL",
	"MBS_SES_REL_RSP",
}

func (v *NgapIeTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NgapIeTypeAnyOf(value)
	for _, existing := range AllowedNgapIeTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NgapIeTypeAnyOf", value)
}

// NewNgapIeTypeAnyOfFromValue returns a pointer to a valid NgapIeTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNgapIeTypeAnyOfFromValue(v string) (*NgapIeTypeAnyOf, error) {
	ev := NgapIeTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NgapIeTypeAnyOf: valid values are %v", v, AllowedNgapIeTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NgapIeTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedNgapIeTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NgapIeType_anyOf value
func (v NgapIeTypeAnyOf) Ptr() *NgapIeTypeAnyOf {
	return &v
}

type NullableNgapIeTypeAnyOf struct {
	value *NgapIeTypeAnyOf
	isSet bool
}

func (v NullableNgapIeTypeAnyOf) Get() *NgapIeTypeAnyOf {
	return v.value
}

func (v *NullableNgapIeTypeAnyOf) Set(val *NgapIeTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNgapIeTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNgapIeTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNgapIeTypeAnyOf(val *NgapIeTypeAnyOf) *NullableNgapIeTypeAnyOf {
	return &NullableNgapIeTypeAnyOf{value: val, isSet: true}
}

func (v NullableNgapIeTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNgapIeTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
