/*
Npcf_EventExposure

PCF Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Npcf_EventExposure

import (
	"encoding/json"
	"fmt"
)

// SatelliteBackhaulCategoryAnyOf the model 'SatelliteBackhaulCategoryAnyOf'
type SatelliteBackhaulCategoryAnyOf string

// List of SatelliteBackhaulCategory_anyOf
const (
	GEO SatelliteBackhaulCategoryAnyOf = "GEO"
	MEO SatelliteBackhaulCategoryAnyOf = "MEO"
	LEO SatelliteBackhaulCategoryAnyOf = "LEO"
	OTHER_SAT SatelliteBackhaulCategoryAnyOf = "OTHER_SAT"
	NON_SATELLITE SatelliteBackhaulCategoryAnyOf = "NON_SATELLITE"
)

// All allowed values of SatelliteBackhaulCategoryAnyOf enum
var AllowedSatelliteBackhaulCategoryAnyOfEnumValues = []SatelliteBackhaulCategoryAnyOf{
	"GEO",
	"MEO",
	"LEO",
	"OTHER_SAT",
	"NON_SATELLITE",
}

func (v *SatelliteBackhaulCategoryAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SatelliteBackhaulCategoryAnyOf(value)
	for _, existing := range AllowedSatelliteBackhaulCategoryAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SatelliteBackhaulCategoryAnyOf", value)
}

// NewSatelliteBackhaulCategoryAnyOfFromValue returns a pointer to a valid SatelliteBackhaulCategoryAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSatelliteBackhaulCategoryAnyOfFromValue(v string) (*SatelliteBackhaulCategoryAnyOf, error) {
	ev := SatelliteBackhaulCategoryAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SatelliteBackhaulCategoryAnyOf: valid values are %v", v, AllowedSatelliteBackhaulCategoryAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SatelliteBackhaulCategoryAnyOf) IsValid() bool {
	for _, existing := range AllowedSatelliteBackhaulCategoryAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SatelliteBackhaulCategory_anyOf value
func (v SatelliteBackhaulCategoryAnyOf) Ptr() *SatelliteBackhaulCategoryAnyOf {
	return &v
}

type NullableSatelliteBackhaulCategoryAnyOf struct {
	value *SatelliteBackhaulCategoryAnyOf
	isSet bool
}

func (v NullableSatelliteBackhaulCategoryAnyOf) Get() *SatelliteBackhaulCategoryAnyOf {
	return v.value
}

func (v *NullableSatelliteBackhaulCategoryAnyOf) Set(val *SatelliteBackhaulCategoryAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSatelliteBackhaulCategoryAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSatelliteBackhaulCategoryAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSatelliteBackhaulCategoryAnyOf(val *SatelliteBackhaulCategoryAnyOf) *NullableSatelliteBackhaulCategoryAnyOf {
	return &NullableSatelliteBackhaulCategoryAnyOf{value: val, isSet: true}
}

func (v NullableSatelliteBackhaulCategoryAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSatelliteBackhaulCategoryAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
