/*
VAE_MessageDelivery

API for VAE Message Delivery Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package VAE_MessageDelivery

import (
	"encoding/json"
	"fmt"
)

// ResultAnyOf the model 'ResultAnyOf'
type ResultAnyOf string

// List of Result_anyOf
const (
	SUCCESS ResultAnyOf = "SUCCESS"
	FAIL ResultAnyOf = "FAIL"
)

// All allowed values of ResultAnyOf enum
var AllowedResultAnyOfEnumValues = []ResultAnyOf{
	"SUCCESS",
	"FAIL",
}

func (v *ResultAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ResultAnyOf(value)
	for _, existing := range AllowedResultAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ResultAnyOf", value)
}

// NewResultAnyOfFromValue returns a pointer to a valid ResultAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewResultAnyOfFromValue(v string) (*ResultAnyOf, error) {
	ev := ResultAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ResultAnyOf: valid values are %v", v, AllowedResultAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ResultAnyOf) IsValid() bool {
	for _, existing := range AllowedResultAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Result_anyOf value
func (v ResultAnyOf) Ptr() *ResultAnyOf {
	return &v
}

type NullableResultAnyOf struct {
	value *ResultAnyOf
	isSet bool
}

func (v NullableResultAnyOf) Get() *ResultAnyOf {
	return v.value
}

func (v *NullableResultAnyOf) Set(val *ResultAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableResultAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableResultAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResultAnyOf(val *ResultAnyOf) *NullableResultAnyOf {
	return &NullableResultAnyOf{value: val, isSet: true}
}

func (v NullableResultAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResultAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
