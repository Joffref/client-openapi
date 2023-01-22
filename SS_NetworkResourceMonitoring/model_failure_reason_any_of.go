/*
SS_NetworkResourceMonitoring

API for SEAL Network Resource Monitoring.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package SS_NetworkResourceMonitoring

import (
	"encoding/json"
	"fmt"
)

// FailureReasonAnyOf the model 'FailureReasonAnyOf'
type FailureReasonAnyOf string

// List of FailureReason_anyOf
const (
	USER_NOT_FOUND FailureReasonAnyOf = "USER_NOT_FOUND"
	STREAM_NOT_FOUND FailureReasonAnyOf = "STREAM_NOT_FOUND"
	DATA_NOT_AVAILABLE FailureReasonAnyOf = "DATA_NOT_AVAILABLE"
	OTHER_REASON FailureReasonAnyOf = "OTHER_REASON"
)

// All allowed values of FailureReasonAnyOf enum
var AllowedFailureReasonAnyOfEnumValues = []FailureReasonAnyOf{
	"USER_NOT_FOUND",
	"STREAM_NOT_FOUND",
	"DATA_NOT_AVAILABLE",
	"OTHER_REASON",
}

func (v *FailureReasonAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FailureReasonAnyOf(value)
	for _, existing := range AllowedFailureReasonAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FailureReasonAnyOf", value)
}

// NewFailureReasonAnyOfFromValue returns a pointer to a valid FailureReasonAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFailureReasonAnyOfFromValue(v string) (*FailureReasonAnyOf, error) {
	ev := FailureReasonAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FailureReasonAnyOf: valid values are %v", v, AllowedFailureReasonAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FailureReasonAnyOf) IsValid() bool {
	for _, existing := range AllowedFailureReasonAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FailureReason_anyOf value
func (v FailureReasonAnyOf) Ptr() *FailureReasonAnyOf {
	return &v
}

type NullableFailureReasonAnyOf struct {
	value *FailureReasonAnyOf
	isSet bool
}

func (v NullableFailureReasonAnyOf) Get() *FailureReasonAnyOf {
	return v.value
}

func (v *NullableFailureReasonAnyOf) Set(val *FailureReasonAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFailureReasonAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFailureReasonAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFailureReasonAnyOf(val *FailureReasonAnyOf) *NullableFailureReasonAnyOf {
	return &NullableFailureReasonAnyOf{value: val, isSet: true}
}

func (v NullableFailureReasonAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFailureReasonAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
