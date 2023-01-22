/*
3gpp-cp-parameter-provisioning

API for provisioning communication pattern parameters.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package CpProvisioning

import (
	"encoding/json"
	"fmt"
)

// CommunicationIndicatorAnyOf the model 'CommunicationIndicatorAnyOf'
type CommunicationIndicatorAnyOf string

// List of CommunicationIndicator_anyOf
const (
	PERIODICALLY CommunicationIndicatorAnyOf = "PERIODICALLY"
	ON_DEMAND CommunicationIndicatorAnyOf = "ON_DEMAND"
)

// All allowed values of CommunicationIndicatorAnyOf enum
var AllowedCommunicationIndicatorAnyOfEnumValues = []CommunicationIndicatorAnyOf{
	"PERIODICALLY",
	"ON_DEMAND",
}

func (v *CommunicationIndicatorAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CommunicationIndicatorAnyOf(value)
	for _, existing := range AllowedCommunicationIndicatorAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CommunicationIndicatorAnyOf", value)
}

// NewCommunicationIndicatorAnyOfFromValue returns a pointer to a valid CommunicationIndicatorAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCommunicationIndicatorAnyOfFromValue(v string) (*CommunicationIndicatorAnyOf, error) {
	ev := CommunicationIndicatorAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CommunicationIndicatorAnyOf: valid values are %v", v, AllowedCommunicationIndicatorAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CommunicationIndicatorAnyOf) IsValid() bool {
	for _, existing := range AllowedCommunicationIndicatorAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CommunicationIndicator_anyOf value
func (v CommunicationIndicatorAnyOf) Ptr() *CommunicationIndicatorAnyOf {
	return &v
}

type NullableCommunicationIndicatorAnyOf struct {
	value *CommunicationIndicatorAnyOf
	isSet bool
}

func (v NullableCommunicationIndicatorAnyOf) Get() *CommunicationIndicatorAnyOf {
	return v.value
}

func (v *NullableCommunicationIndicatorAnyOf) Set(val *CommunicationIndicatorAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCommunicationIndicatorAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCommunicationIndicatorAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommunicationIndicatorAnyOf(val *CommunicationIndicatorAnyOf) *NullableCommunicationIndicatorAnyOf {
	return &NullableCommunicationIndicatorAnyOf{value: val, isSet: true}
}

func (v NullableCommunicationIndicatorAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommunicationIndicatorAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
