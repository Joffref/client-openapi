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

// AreaOfInterest - struct for AreaOfInterest
type AreaOfInterest struct {
	GeoAreaToCellMapping *GeoAreaToCellMapping
	ArrayOfEutraCellId *[]EutraCellId
	ArrayOfNrCellId *[]NrCellId
	ArrayOfTai1 *[]Tai1
	ArrayOfUtraCellId *[]UtraCellId
}

// GeoAreaToCellMappingAsAreaOfInterest is a convenience function that returns GeoAreaToCellMapping wrapped in AreaOfInterest
func GeoAreaToCellMappingAsAreaOfInterest(v *GeoAreaToCellMapping) AreaOfInterest {
	return AreaOfInterest{
		GeoAreaToCellMapping: v,
	}
}

// []EutraCellIdAsAreaOfInterest is a convenience function that returns []EutraCellId wrapped in AreaOfInterest
func ArrayOfEutraCellIdAsAreaOfInterest(v *[]EutraCellId) AreaOfInterest {
	return AreaOfInterest{
		ArrayOfEutraCellId: v,
	}
}

// []NrCellIdAsAreaOfInterest is a convenience function that returns []NrCellId wrapped in AreaOfInterest
func ArrayOfNrCellIdAsAreaOfInterest(v *[]NrCellId) AreaOfInterest {
	return AreaOfInterest{
		ArrayOfNrCellId: v,
	}
}

// []Tai1AsAreaOfInterest is a convenience function that returns []Tai1 wrapped in AreaOfInterest
func ArrayOfTai1AsAreaOfInterest(v *[]Tai1) AreaOfInterest {
	return AreaOfInterest{
		ArrayOfTai1: v,
	}
}

// []UtraCellIdAsAreaOfInterest is a convenience function that returns []UtraCellId wrapped in AreaOfInterest
func ArrayOfUtraCellIdAsAreaOfInterest(v *[]UtraCellId) AreaOfInterest {
	return AreaOfInterest{
		ArrayOfUtraCellId: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AreaOfInterest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GeoAreaToCellMapping
	err = newStrictDecoder(data).Decode(&dst.GeoAreaToCellMapping)
	if err == nil {
		jsonGeoAreaToCellMapping, _ := json.Marshal(dst.GeoAreaToCellMapping)
		if string(jsonGeoAreaToCellMapping) == "{}" { // empty struct
			dst.GeoAreaToCellMapping = nil
		} else {
			match++
		}
	} else {
		dst.GeoAreaToCellMapping = nil
	}

	// try to unmarshal data into ArrayOfEutraCellId
	err = newStrictDecoder(data).Decode(&dst.ArrayOfEutraCellId)
	if err == nil {
		jsonArrayOfEutraCellId, _ := json.Marshal(dst.ArrayOfEutraCellId)
		if string(jsonArrayOfEutraCellId) == "{}" { // empty struct
			dst.ArrayOfEutraCellId = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfEutraCellId = nil
	}

	// try to unmarshal data into ArrayOfNrCellId
	err = newStrictDecoder(data).Decode(&dst.ArrayOfNrCellId)
	if err == nil {
		jsonArrayOfNrCellId, _ := json.Marshal(dst.ArrayOfNrCellId)
		if string(jsonArrayOfNrCellId) == "{}" { // empty struct
			dst.ArrayOfNrCellId = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfNrCellId = nil
	}

	// try to unmarshal data into ArrayOfTai1
	err = newStrictDecoder(data).Decode(&dst.ArrayOfTai1)
	if err == nil {
		jsonArrayOfTai1, _ := json.Marshal(dst.ArrayOfTai1)
		if string(jsonArrayOfTai1) == "{}" { // empty struct
			dst.ArrayOfTai1 = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfTai1 = nil
	}

	// try to unmarshal data into ArrayOfUtraCellId
	err = newStrictDecoder(data).Decode(&dst.ArrayOfUtraCellId)
	if err == nil {
		jsonArrayOfUtraCellId, _ := json.Marshal(dst.ArrayOfUtraCellId)
		if string(jsonArrayOfUtraCellId) == "{}" { // empty struct
			dst.ArrayOfUtraCellId = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfUtraCellId = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GeoAreaToCellMapping = nil
		dst.ArrayOfEutraCellId = nil
		dst.ArrayOfNrCellId = nil
		dst.ArrayOfTai1 = nil
		dst.ArrayOfUtraCellId = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AreaOfInterest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AreaOfInterest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AreaOfInterest) MarshalJSON() ([]byte, error) {
	if src.GeoAreaToCellMapping != nil {
		return json.Marshal(&src.GeoAreaToCellMapping)
	}

	if src.ArrayOfEutraCellId != nil {
		return json.Marshal(&src.ArrayOfEutraCellId)
	}

	if src.ArrayOfNrCellId != nil {
		return json.Marshal(&src.ArrayOfNrCellId)
	}

	if src.ArrayOfTai1 != nil {
		return json.Marshal(&src.ArrayOfTai1)
	}

	if src.ArrayOfUtraCellId != nil {
		return json.Marshal(&src.ArrayOfUtraCellId)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AreaOfInterest) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.GeoAreaToCellMapping != nil {
		return obj.GeoAreaToCellMapping
	}

	if obj.ArrayOfEutraCellId != nil {
		return obj.ArrayOfEutraCellId
	}

	if obj.ArrayOfNrCellId != nil {
		return obj.ArrayOfNrCellId
	}

	if obj.ArrayOfTai1 != nil {
		return obj.ArrayOfTai1
	}

	if obj.ArrayOfUtraCellId != nil {
		return obj.ArrayOfUtraCellId
	}

	// all schemas are nil
	return nil
}

type NullableAreaOfInterest struct {
	value *AreaOfInterest
	isSet bool
}

func (v NullableAreaOfInterest) Get() *AreaOfInterest {
	return v.value
}

func (v *NullableAreaOfInterest) Set(val *AreaOfInterest) {
	v.value = val
	v.isSet = true
}

func (v NullableAreaOfInterest) IsSet() bool {
	return v.isSet
}

func (v *NullableAreaOfInterest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAreaOfInterest(val *AreaOfInterest) *NullableAreaOfInterest {
	return &NullableAreaOfInterest{value: val, isSet: true}
}

func (v NullableAreaOfInterest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAreaOfInterest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

