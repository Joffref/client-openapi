/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Subscription_Data

import (
	"encoding/json"
	"fmt"
)

// SensorMeasurementAnyOf the model 'SensorMeasurementAnyOf'
type SensorMeasurementAnyOf string

// List of SensorMeasurement_anyOf
const (
	BAROMETRIC_PRESSURE SensorMeasurementAnyOf = "BAROMETRIC_PRESSURE"
	UE_SPEED SensorMeasurementAnyOf = "UE_SPEED"
	UE_ORIENTATION SensorMeasurementAnyOf = "UE_ORIENTATION"
)

// All allowed values of SensorMeasurementAnyOf enum
var AllowedSensorMeasurementAnyOfEnumValues = []SensorMeasurementAnyOf{
	"BAROMETRIC_PRESSURE",
	"UE_SPEED",
	"UE_ORIENTATION",
}

func (v *SensorMeasurementAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SensorMeasurementAnyOf(value)
	for _, existing := range AllowedSensorMeasurementAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SensorMeasurementAnyOf", value)
}

// NewSensorMeasurementAnyOfFromValue returns a pointer to a valid SensorMeasurementAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSensorMeasurementAnyOfFromValue(v string) (*SensorMeasurementAnyOf, error) {
	ev := SensorMeasurementAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SensorMeasurementAnyOf: valid values are %v", v, AllowedSensorMeasurementAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SensorMeasurementAnyOf) IsValid() bool {
	for _, existing := range AllowedSensorMeasurementAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SensorMeasurement_anyOf value
func (v SensorMeasurementAnyOf) Ptr() *SensorMeasurementAnyOf {
	return &v
}

type NullableSensorMeasurementAnyOf struct {
	value *SensorMeasurementAnyOf
	isSet bool
}

func (v NullableSensorMeasurementAnyOf) Get() *SensorMeasurementAnyOf {
	return v.value
}

func (v *NullableSensorMeasurementAnyOf) Set(val *SensorMeasurementAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSensorMeasurementAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSensorMeasurementAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSensorMeasurementAnyOf(val *SensorMeasurementAnyOf) *NullableSensorMeasurementAnyOf {
	return &NullableSensorMeasurementAnyOf{value: val, isSet: true}
}

func (v NullableSensorMeasurementAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSensorMeasurementAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
