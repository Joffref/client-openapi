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

// Condition the model 'Condition'
type Condition string

// List of Condition
const (
	EQUAL_TO Condition = "Is_equal_to"
	LESS_THAN Condition = "Is_less_than"
	GREATER_THAN Condition = "Is_greater_than"
	WITHIN_THE_RANGE Condition = "Is_within_the_range"
)

// All allowed values of Condition enum
var AllowedConditionEnumValues = []Condition{
	"Is_equal_to",
	"Is_less_than",
	"Is_greater_than",
	"Is_within_the_range",
}

func (v *Condition) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Condition(value)
	for _, existing := range AllowedConditionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Condition", value)
}

// NewConditionFromValue returns a pointer to a valid Condition
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConditionFromValue(v string) (*Condition, error) {
	ev := Condition(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Condition: valid values are %v", v, AllowedConditionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Condition) IsValid() bool {
	for _, existing := range AllowedConditionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Condition value
func (v Condition) Ptr() *Condition {
	return &v
}

type NullableCondition struct {
	value *Condition
	isSet bool
}

func (v NullableCondition) Get() *Condition {
	return v.value
}

func (v *NullableCondition) Set(val *Condition) {
	v.value = val
	v.isSet = true
}

func (v NullableCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCondition(val *Condition) *NullableCondition {
	return &NullableCondition{value: val, isSet: true}
}

func (v NullableCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
