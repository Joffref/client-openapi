/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nchf_ConvergedCharging

import (
	"encoding/json"
	"fmt"
)

// RedundantTransmissionTypeAnyOf the model 'RedundantTransmissionTypeAnyOf'
type RedundantTransmissionTypeAnyOf string

// List of RedundantTransmissionType_anyOf
const (
	NON_TRANSMISSION RedundantTransmissionTypeAnyOf = "NON_TRANSMISSION"
	END_TO_END_USER_PLANE_PATHS RedundantTransmissionTypeAnyOf = "END_TO_END_USER_PLANE_PATHS"
	N3_N9 RedundantTransmissionTypeAnyOf = "N3/N9"
	TRANSPORT_LAYER RedundantTransmissionTypeAnyOf = "TRANSPORT_LAYER"
)

// All allowed values of RedundantTransmissionTypeAnyOf enum
var AllowedRedundantTransmissionTypeAnyOfEnumValues = []RedundantTransmissionTypeAnyOf{
	"NON_TRANSMISSION",
	"END_TO_END_USER_PLANE_PATHS",
	"N3/N9",
	"TRANSPORT_LAYER",
}

func (v *RedundantTransmissionTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RedundantTransmissionTypeAnyOf(value)
	for _, existing := range AllowedRedundantTransmissionTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RedundantTransmissionTypeAnyOf", value)
}

// NewRedundantTransmissionTypeAnyOfFromValue returns a pointer to a valid RedundantTransmissionTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRedundantTransmissionTypeAnyOfFromValue(v string) (*RedundantTransmissionTypeAnyOf, error) {
	ev := RedundantTransmissionTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RedundantTransmissionTypeAnyOf: valid values are %v", v, AllowedRedundantTransmissionTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RedundantTransmissionTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedRedundantTransmissionTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RedundantTransmissionType_anyOf value
func (v RedundantTransmissionTypeAnyOf) Ptr() *RedundantTransmissionTypeAnyOf {
	return &v
}

type NullableRedundantTransmissionTypeAnyOf struct {
	value *RedundantTransmissionTypeAnyOf
	isSet bool
}

func (v NullableRedundantTransmissionTypeAnyOf) Get() *RedundantTransmissionTypeAnyOf {
	return v.value
}

func (v *NullableRedundantTransmissionTypeAnyOf) Set(val *RedundantTransmissionTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRedundantTransmissionTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRedundantTransmissionTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedundantTransmissionTypeAnyOf(val *RedundantTransmissionTypeAnyOf) *NullableRedundantTransmissionTypeAnyOf {
	return &NullableRedundantTransmissionTypeAnyOf{value: val, isSet: true}
}

func (v NullableRedundantTransmissionTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedundantTransmissionTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
