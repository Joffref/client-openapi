/*
coslaNrm

OAS 3.0.1 specification of the Cosla NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package CoslaNrm

import (
	"encoding/json"
	"fmt"
)

// CollectionPeriodM6LTEType See details in 3GPP TS 32.422 clause 5.10.32.
type CollectionPeriodM6LTEType string

// List of collectionPeriodM6LTE-Type
const (
	_1024MS CollectionPeriodM6LTEType = "1024ms"
	_2048MS CollectionPeriodM6LTEType = "2048ms"
	_5120MS CollectionPeriodM6LTEType = "5120ms"
	_10240MS CollectionPeriodM6LTEType = "10240ms"
)

// All allowed values of CollectionPeriodM6LTEType enum
var AllowedCollectionPeriodM6LTETypeEnumValues = []CollectionPeriodM6LTEType{
	"1024ms",
	"2048ms",
	"5120ms",
	"10240ms",
}

func (v *CollectionPeriodM6LTEType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CollectionPeriodM6LTEType(value)
	for _, existing := range AllowedCollectionPeriodM6LTETypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CollectionPeriodM6LTEType", value)
}

// NewCollectionPeriodM6LTETypeFromValue returns a pointer to a valid CollectionPeriodM6LTEType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCollectionPeriodM6LTETypeFromValue(v string) (*CollectionPeriodM6LTEType, error) {
	ev := CollectionPeriodM6LTEType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CollectionPeriodM6LTEType: valid values are %v", v, AllowedCollectionPeriodM6LTETypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CollectionPeriodM6LTEType) IsValid() bool {
	for _, existing := range AllowedCollectionPeriodM6LTETypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to collectionPeriodM6LTE-Type value
func (v CollectionPeriodM6LTEType) Ptr() *CollectionPeriodM6LTEType {
	return &v
}

type NullableCollectionPeriodM6LTEType struct {
	value *CollectionPeriodM6LTEType
	isSet bool
}

func (v NullableCollectionPeriodM6LTEType) Get() *CollectionPeriodM6LTEType {
	return v.value
}

func (v *NullableCollectionPeriodM6LTEType) Set(val *CollectionPeriodM6LTEType) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionPeriodM6LTEType) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionPeriodM6LTEType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionPeriodM6LTEType(val *CollectionPeriodM6LTEType) *NullableCollectionPeriodM6LTEType {
	return &NullableCollectionPeriodM6LTEType{value: val, isSet: true}
}

func (v NullableCollectionPeriodM6LTEType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionPeriodM6LTEType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
