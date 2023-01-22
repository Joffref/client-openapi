/*
TS 28.532 Streaming data reporting service

OAS 3.0.1 specification for the Streaming data reporting service (Streaming MnS) © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package StreamingDataMnS

import (
	"encoding/json"
	"fmt"
)

// CollectionPeriodRRMUMTSType See details in 3GPP TS 32.422 clause 5.10.21.
type CollectionPeriodRRMUMTSType string

// List of collectionPeriodRRMUMTS-Type
const (
	_100MS CollectionPeriodRRMUMTSType = "100ms"
	_250MS CollectionPeriodRRMUMTSType = "250ms"
	_500MS CollectionPeriodRRMUMTSType = "500ms"
	_1000MS CollectionPeriodRRMUMTSType = "1000ms"
	_2000MS CollectionPeriodRRMUMTSType = "2000ms"
	_3000MS CollectionPeriodRRMUMTSType = "3000ms"
	_4000MS CollectionPeriodRRMUMTSType = "4000ms"
	_6000MS CollectionPeriodRRMUMTSType = "6000ms"
)

// All allowed values of CollectionPeriodRRMUMTSType enum
var AllowedCollectionPeriodRRMUMTSTypeEnumValues = []CollectionPeriodRRMUMTSType{
	"100ms",
	"250ms",
	"500ms",
	"1000ms",
	"2000ms",
	"3000ms",
	"4000ms",
	"6000ms",
}

func (v *CollectionPeriodRRMUMTSType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CollectionPeriodRRMUMTSType(value)
	for _, existing := range AllowedCollectionPeriodRRMUMTSTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CollectionPeriodRRMUMTSType", value)
}

// NewCollectionPeriodRRMUMTSTypeFromValue returns a pointer to a valid CollectionPeriodRRMUMTSType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCollectionPeriodRRMUMTSTypeFromValue(v string) (*CollectionPeriodRRMUMTSType, error) {
	ev := CollectionPeriodRRMUMTSType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CollectionPeriodRRMUMTSType: valid values are %v", v, AllowedCollectionPeriodRRMUMTSTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CollectionPeriodRRMUMTSType) IsValid() bool {
	for _, existing := range AllowedCollectionPeriodRRMUMTSTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to collectionPeriodRRMUMTS-Type value
func (v CollectionPeriodRRMUMTSType) Ptr() *CollectionPeriodRRMUMTSType {
	return &v
}

type NullableCollectionPeriodRRMUMTSType struct {
	value *CollectionPeriodRRMUMTSType
	isSet bool
}

func (v NullableCollectionPeriodRRMUMTSType) Get() *CollectionPeriodRRMUMTSType {
	return v.value
}

func (v *NullableCollectionPeriodRRMUMTSType) Set(val *CollectionPeriodRRMUMTSType) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionPeriodRRMUMTSType) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionPeriodRRMUMTSType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionPeriodRRMUMTSType(val *CollectionPeriodRRMUMTSType) *NullableCollectionPeriodRRMUMTSType {
	return &NullableCollectionPeriodRRMUMTSType{value: val, isSet: true}
}

func (v NullableCollectionPeriodRRMUMTSType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionPeriodRRMUMTSType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
