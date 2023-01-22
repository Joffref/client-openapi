/*
Nnsacf_NSAC

Nnsacf_NSAC Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnsacf_NSAC

import (
	"encoding/json"
	"fmt"
)

// AcuFailureReasonAnyOf the model 'AcuFailureReasonAnyOf'
type AcuFailureReasonAnyOf string

// List of AcuFailureReason_anyOf
const (
	SLICE_NOT_FOUND AcuFailureReasonAnyOf = "SLICE_NOT_FOUND"
	EXCEED_MAX_UE_NUM AcuFailureReasonAnyOf = "EXCEED_MAX_UE_NUM"
	EXCEED_MAX_UE_NUM_3_GPP AcuFailureReasonAnyOf = "EXCEED_MAX_UE_NUM_3GPP"
	EXCEED_MAX_UE_NUM_N3_GPP AcuFailureReasonAnyOf = "EXCEED_MAX_UE_NUM_N3GPP"
	EXCEED_MAX_PDU_NUM AcuFailureReasonAnyOf = "EXCEED_MAX_PDU_NUM"
	EXCEED_MAX_PDU_NUM_3_GPP AcuFailureReasonAnyOf = "EXCEED_MAX_PDU_NUM_3GPP"
	EXCEED_MAX_PDU_NUM_N3_GPP AcuFailureReasonAnyOf = "EXCEED_MAX_PDU_NUM_N3GPP"
)

// All allowed values of AcuFailureReasonAnyOf enum
var AllowedAcuFailureReasonAnyOfEnumValues = []AcuFailureReasonAnyOf{
	"SLICE_NOT_FOUND",
	"EXCEED_MAX_UE_NUM",
	"EXCEED_MAX_UE_NUM_3GPP",
	"EXCEED_MAX_UE_NUM_N3GPP",
	"EXCEED_MAX_PDU_NUM",
	"EXCEED_MAX_PDU_NUM_3GPP",
	"EXCEED_MAX_PDU_NUM_N3GPP",
}

func (v *AcuFailureReasonAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AcuFailureReasonAnyOf(value)
	for _, existing := range AllowedAcuFailureReasonAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AcuFailureReasonAnyOf", value)
}

// NewAcuFailureReasonAnyOfFromValue returns a pointer to a valid AcuFailureReasonAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAcuFailureReasonAnyOfFromValue(v string) (*AcuFailureReasonAnyOf, error) {
	ev := AcuFailureReasonAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AcuFailureReasonAnyOf: valid values are %v", v, AllowedAcuFailureReasonAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AcuFailureReasonAnyOf) IsValid() bool {
	for _, existing := range AllowedAcuFailureReasonAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AcuFailureReason_anyOf value
func (v AcuFailureReasonAnyOf) Ptr() *AcuFailureReasonAnyOf {
	return &v
}

type NullableAcuFailureReasonAnyOf struct {
	value *AcuFailureReasonAnyOf
	isSet bool
}

func (v NullableAcuFailureReasonAnyOf) Get() *AcuFailureReasonAnyOf {
	return v.value
}

func (v *NullableAcuFailureReasonAnyOf) Set(val *AcuFailureReasonAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAcuFailureReasonAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAcuFailureReasonAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcuFailureReasonAnyOf(val *AcuFailureReasonAnyOf) *NullableAcuFailureReasonAnyOf {
	return &NullableAcuFailureReasonAnyOf{value: val, isSet: true}
}

func (v NullableAcuFailureReasonAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcuFailureReasonAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
