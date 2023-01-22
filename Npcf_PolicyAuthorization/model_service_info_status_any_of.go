/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Npcf_PolicyAuthorization

import (
	"encoding/json"
	"fmt"
)

// ServiceInfoStatusAnyOf the model 'ServiceInfoStatusAnyOf'
type ServiceInfoStatusAnyOf string

// List of ServiceInfoStatus_anyOf
const (
	FINAL ServiceInfoStatusAnyOf = "FINAL"
	PRELIMINARY ServiceInfoStatusAnyOf = "PRELIMINARY"
)

// All allowed values of ServiceInfoStatusAnyOf enum
var AllowedServiceInfoStatusAnyOfEnumValues = []ServiceInfoStatusAnyOf{
	"FINAL",
	"PRELIMINARY",
}

func (v *ServiceInfoStatusAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ServiceInfoStatusAnyOf(value)
	for _, existing := range AllowedServiceInfoStatusAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ServiceInfoStatusAnyOf", value)
}

// NewServiceInfoStatusAnyOfFromValue returns a pointer to a valid ServiceInfoStatusAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewServiceInfoStatusAnyOfFromValue(v string) (*ServiceInfoStatusAnyOf, error) {
	ev := ServiceInfoStatusAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ServiceInfoStatusAnyOf: valid values are %v", v, AllowedServiceInfoStatusAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ServiceInfoStatusAnyOf) IsValid() bool {
	for _, existing := range AllowedServiceInfoStatusAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceInfoStatus_anyOf value
func (v ServiceInfoStatusAnyOf) Ptr() *ServiceInfoStatusAnyOf {
	return &v
}

type NullableServiceInfoStatusAnyOf struct {
	value *ServiceInfoStatusAnyOf
	isSet bool
}

func (v NullableServiceInfoStatusAnyOf) Get() *ServiceInfoStatusAnyOf {
	return v.value
}

func (v *NullableServiceInfoStatusAnyOf) Set(val *ServiceInfoStatusAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceInfoStatusAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceInfoStatusAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceInfoStatusAnyOf(val *ServiceInfoStatusAnyOf) *NullableServiceInfoStatusAnyOf {
	return &NullableServiceInfoStatusAnyOf{value: val, isSet: true}
}

func (v NullableServiceInfoStatusAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceInfoStatusAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
