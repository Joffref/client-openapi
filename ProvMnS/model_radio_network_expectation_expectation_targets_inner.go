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

// RadioNetworkExpectationExpectationTargetsInner - struct for RadioNetworkExpectationExpectationTargetsInner
type RadioNetworkExpectationExpectationTargetsInner struct {
	AveDLRANUEThptTarget *AveDLRANUEThptTarget
	AveULRANUEThptTarget *AveULRANUEThptTarget
	ExpectationTarget *ExpectationTarget
	LowDLRANUEThptRatioTarget *LowDLRANUEThptRatioTarget
	LowSINRRatioTarget *LowSINRRatioTarget
	LowULRANUEThptRatioTarget *LowULRANUEThptRatioTarget
	WeakRSRPRatioTarget *WeakRSRPRatioTarget
}

// AveDLRANUEThptTargetAsRadioNetworkExpectationExpectationTargetsInner is a convenience function that returns AveDLRANUEThptTarget wrapped in RadioNetworkExpectationExpectationTargetsInner
func AveDLRANUEThptTargetAsRadioNetworkExpectationExpectationTargetsInner(v *AveDLRANUEThptTarget) RadioNetworkExpectationExpectationTargetsInner {
	return RadioNetworkExpectationExpectationTargetsInner{
		AveDLRANUEThptTarget: v,
	}
}

// AveULRANUEThptTargetAsRadioNetworkExpectationExpectationTargetsInner is a convenience function that returns AveULRANUEThptTarget wrapped in RadioNetworkExpectationExpectationTargetsInner
func AveULRANUEThptTargetAsRadioNetworkExpectationExpectationTargetsInner(v *AveULRANUEThptTarget) RadioNetworkExpectationExpectationTargetsInner {
	return RadioNetworkExpectationExpectationTargetsInner{
		AveULRANUEThptTarget: v,
	}
}

// ExpectationTargetAsRadioNetworkExpectationExpectationTargetsInner is a convenience function that returns ExpectationTarget wrapped in RadioNetworkExpectationExpectationTargetsInner
func ExpectationTargetAsRadioNetworkExpectationExpectationTargetsInner(v *ExpectationTarget) RadioNetworkExpectationExpectationTargetsInner {
	return RadioNetworkExpectationExpectationTargetsInner{
		ExpectationTarget: v,
	}
}

// LowDLRANUEThptRatioTargetAsRadioNetworkExpectationExpectationTargetsInner is a convenience function that returns LowDLRANUEThptRatioTarget wrapped in RadioNetworkExpectationExpectationTargetsInner
func LowDLRANUEThptRatioTargetAsRadioNetworkExpectationExpectationTargetsInner(v *LowDLRANUEThptRatioTarget) RadioNetworkExpectationExpectationTargetsInner {
	return RadioNetworkExpectationExpectationTargetsInner{
		LowDLRANUEThptRatioTarget: v,
	}
}

// LowSINRRatioTargetAsRadioNetworkExpectationExpectationTargetsInner is a convenience function that returns LowSINRRatioTarget wrapped in RadioNetworkExpectationExpectationTargetsInner
func LowSINRRatioTargetAsRadioNetworkExpectationExpectationTargetsInner(v *LowSINRRatioTarget) RadioNetworkExpectationExpectationTargetsInner {
	return RadioNetworkExpectationExpectationTargetsInner{
		LowSINRRatioTarget: v,
	}
}

// LowULRANUEThptRatioTargetAsRadioNetworkExpectationExpectationTargetsInner is a convenience function that returns LowULRANUEThptRatioTarget wrapped in RadioNetworkExpectationExpectationTargetsInner
func LowULRANUEThptRatioTargetAsRadioNetworkExpectationExpectationTargetsInner(v *LowULRANUEThptRatioTarget) RadioNetworkExpectationExpectationTargetsInner {
	return RadioNetworkExpectationExpectationTargetsInner{
		LowULRANUEThptRatioTarget: v,
	}
}

// WeakRSRPRatioTargetAsRadioNetworkExpectationExpectationTargetsInner is a convenience function that returns WeakRSRPRatioTarget wrapped in RadioNetworkExpectationExpectationTargetsInner
func WeakRSRPRatioTargetAsRadioNetworkExpectationExpectationTargetsInner(v *WeakRSRPRatioTarget) RadioNetworkExpectationExpectationTargetsInner {
	return RadioNetworkExpectationExpectationTargetsInner{
		WeakRSRPRatioTarget: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *RadioNetworkExpectationExpectationTargetsInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AveDLRANUEThptTarget
	err = newStrictDecoder(data).Decode(&dst.AveDLRANUEThptTarget)
	if err == nil {
		jsonAveDLRANUEThptTarget, _ := json.Marshal(dst.AveDLRANUEThptTarget)
		if string(jsonAveDLRANUEThptTarget) == "{}" { // empty struct
			dst.AveDLRANUEThptTarget = nil
		} else {
			match++
		}
	} else {
		dst.AveDLRANUEThptTarget = nil
	}

	// try to unmarshal data into AveULRANUEThptTarget
	err = newStrictDecoder(data).Decode(&dst.AveULRANUEThptTarget)
	if err == nil {
		jsonAveULRANUEThptTarget, _ := json.Marshal(dst.AveULRANUEThptTarget)
		if string(jsonAveULRANUEThptTarget) == "{}" { // empty struct
			dst.AveULRANUEThptTarget = nil
		} else {
			match++
		}
	} else {
		dst.AveULRANUEThptTarget = nil
	}

	// try to unmarshal data into ExpectationTarget
	err = newStrictDecoder(data).Decode(&dst.ExpectationTarget)
	if err == nil {
		jsonExpectationTarget, _ := json.Marshal(dst.ExpectationTarget)
		if string(jsonExpectationTarget) == "{}" { // empty struct
			dst.ExpectationTarget = nil
		} else {
			match++
		}
	} else {
		dst.ExpectationTarget = nil
	}

	// try to unmarshal data into LowDLRANUEThptRatioTarget
	err = newStrictDecoder(data).Decode(&dst.LowDLRANUEThptRatioTarget)
	if err == nil {
		jsonLowDLRANUEThptRatioTarget, _ := json.Marshal(dst.LowDLRANUEThptRatioTarget)
		if string(jsonLowDLRANUEThptRatioTarget) == "{}" { // empty struct
			dst.LowDLRANUEThptRatioTarget = nil
		} else {
			match++
		}
	} else {
		dst.LowDLRANUEThptRatioTarget = nil
	}

	// try to unmarshal data into LowSINRRatioTarget
	err = newStrictDecoder(data).Decode(&dst.LowSINRRatioTarget)
	if err == nil {
		jsonLowSINRRatioTarget, _ := json.Marshal(dst.LowSINRRatioTarget)
		if string(jsonLowSINRRatioTarget) == "{}" { // empty struct
			dst.LowSINRRatioTarget = nil
		} else {
			match++
		}
	} else {
		dst.LowSINRRatioTarget = nil
	}

	// try to unmarshal data into LowULRANUEThptRatioTarget
	err = newStrictDecoder(data).Decode(&dst.LowULRANUEThptRatioTarget)
	if err == nil {
		jsonLowULRANUEThptRatioTarget, _ := json.Marshal(dst.LowULRANUEThptRatioTarget)
		if string(jsonLowULRANUEThptRatioTarget) == "{}" { // empty struct
			dst.LowULRANUEThptRatioTarget = nil
		} else {
			match++
		}
	} else {
		dst.LowULRANUEThptRatioTarget = nil
	}

	// try to unmarshal data into WeakRSRPRatioTarget
	err = newStrictDecoder(data).Decode(&dst.WeakRSRPRatioTarget)
	if err == nil {
		jsonWeakRSRPRatioTarget, _ := json.Marshal(dst.WeakRSRPRatioTarget)
		if string(jsonWeakRSRPRatioTarget) == "{}" { // empty struct
			dst.WeakRSRPRatioTarget = nil
		} else {
			match++
		}
	} else {
		dst.WeakRSRPRatioTarget = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AveDLRANUEThptTarget = nil
		dst.AveULRANUEThptTarget = nil
		dst.ExpectationTarget = nil
		dst.LowDLRANUEThptRatioTarget = nil
		dst.LowSINRRatioTarget = nil
		dst.LowULRANUEThptRatioTarget = nil
		dst.WeakRSRPRatioTarget = nil

		return fmt.Errorf("data matches more than one schema in oneOf(RadioNetworkExpectationExpectationTargetsInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(RadioNetworkExpectationExpectationTargetsInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RadioNetworkExpectationExpectationTargetsInner) MarshalJSON() ([]byte, error) {
	if src.AveDLRANUEThptTarget != nil {
		return json.Marshal(&src.AveDLRANUEThptTarget)
	}

	if src.AveULRANUEThptTarget != nil {
		return json.Marshal(&src.AveULRANUEThptTarget)
	}

	if src.ExpectationTarget != nil {
		return json.Marshal(&src.ExpectationTarget)
	}

	if src.LowDLRANUEThptRatioTarget != nil {
		return json.Marshal(&src.LowDLRANUEThptRatioTarget)
	}

	if src.LowSINRRatioTarget != nil {
		return json.Marshal(&src.LowSINRRatioTarget)
	}

	if src.LowULRANUEThptRatioTarget != nil {
		return json.Marshal(&src.LowULRANUEThptRatioTarget)
	}

	if src.WeakRSRPRatioTarget != nil {
		return json.Marshal(&src.WeakRSRPRatioTarget)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RadioNetworkExpectationExpectationTargetsInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AveDLRANUEThptTarget != nil {
		return obj.AveDLRANUEThptTarget
	}

	if obj.AveULRANUEThptTarget != nil {
		return obj.AveULRANUEThptTarget
	}

	if obj.ExpectationTarget != nil {
		return obj.ExpectationTarget
	}

	if obj.LowDLRANUEThptRatioTarget != nil {
		return obj.LowDLRANUEThptRatioTarget
	}

	if obj.LowSINRRatioTarget != nil {
		return obj.LowSINRRatioTarget
	}

	if obj.LowULRANUEThptRatioTarget != nil {
		return obj.LowULRANUEThptRatioTarget
	}

	if obj.WeakRSRPRatioTarget != nil {
		return obj.WeakRSRPRatioTarget
	}

	// all schemas are nil
	return nil
}

type NullableRadioNetworkExpectationExpectationTargetsInner struct {
	value *RadioNetworkExpectationExpectationTargetsInner
	isSet bool
}

func (v NullableRadioNetworkExpectationExpectationTargetsInner) Get() *RadioNetworkExpectationExpectationTargetsInner {
	return v.value
}

func (v *NullableRadioNetworkExpectationExpectationTargetsInner) Set(val *RadioNetworkExpectationExpectationTargetsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableRadioNetworkExpectationExpectationTargetsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableRadioNetworkExpectationExpectationTargetsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRadioNetworkExpectationExpectationTargetsInner(val *RadioNetworkExpectationExpectationTargetsInner) *NullableRadioNetworkExpectationExpectationTargetsInner {
	return &NullableRadioNetworkExpectationExpectationTargetsInner{value: val, isSet: true}
}

func (v NullableRadioNetworkExpectationExpectationTargetsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRadioNetworkExpectationExpectationTargetsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

