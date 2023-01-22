/*
NR NRM

OAS 3.0.1 specification of the NR NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package NrNrm

import (
	"encoding/json"
	"fmt"
)

// IsInitialBwp the model 'IsInitialBwp'
type IsInitialBwp string

// List of IsInitialBwp
const (
	INITIAL IsInitialBwp = "INITIAL"
	OTHER IsInitialBwp = "OTHER"
	SUL IsInitialBwp = "SUL"
)

// All allowed values of IsInitialBwp enum
var AllowedIsInitialBwpEnumValues = []IsInitialBwp{
	"INITIAL",
	"OTHER",
	"SUL",
}

func (v *IsInitialBwp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IsInitialBwp(value)
	for _, existing := range AllowedIsInitialBwpEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IsInitialBwp", value)
}

// NewIsInitialBwpFromValue returns a pointer to a valid IsInitialBwp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIsInitialBwpFromValue(v string) (*IsInitialBwp, error) {
	ev := IsInitialBwp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IsInitialBwp: valid values are %v", v, AllowedIsInitialBwpEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IsInitialBwp) IsValid() bool {
	for _, existing := range AllowedIsInitialBwpEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IsInitialBwp value
func (v IsInitialBwp) Ptr() *IsInitialBwp {
	return &v
}

type NullableIsInitialBwp struct {
	value *IsInitialBwp
	isSet bool
}

func (v NullableIsInitialBwp) Get() *IsInitialBwp {
	return v.value
}

func (v *NullableIsInitialBwp) Set(val *IsInitialBwp) {
	v.value = val
	v.isSet = true
}

func (v NullableIsInitialBwp) IsSet() bool {
	return v.isSet
}

func (v *NullableIsInitialBwp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIsInitialBwp(val *IsInitialBwp) *NullableIsInitialBwp {
	return &NullableIsInitialBwp{value: val, isSet: true}
}

func (v NullableIsInitialBwp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIsInitialBwp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
