/*
Npcf_SMPolicyControl API

Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Npcf_SMPolicyControl

import (
	"encoding/json"
	"fmt"
)

// PolicyDecisionFailureCode Indicates the type of the failed policy decision and/or condition data.
type PolicyDecisionFailureCode struct {
	PolicyDecisionFailureCodeAnyOf *PolicyDecisionFailureCodeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PolicyDecisionFailureCode) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into PolicyDecisionFailureCodeAnyOf
	err = json.Unmarshal(data, &dst.PolicyDecisionFailureCodeAnyOf);
	if err == nil {
		jsonPolicyDecisionFailureCodeAnyOf, _ := json.Marshal(dst.PolicyDecisionFailureCodeAnyOf)
		if string(jsonPolicyDecisionFailureCodeAnyOf) == "{}" { // empty struct
			dst.PolicyDecisionFailureCodeAnyOf = nil
		} else {
			return nil // data stored in dst.PolicyDecisionFailureCodeAnyOf, return on the first match
		}
	} else {
		dst.PolicyDecisionFailureCodeAnyOf = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(PolicyDecisionFailureCode)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PolicyDecisionFailureCode) MarshalJSON() ([]byte, error) {
	if src.PolicyDecisionFailureCodeAnyOf != nil {
		return json.Marshal(&src.PolicyDecisionFailureCodeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePolicyDecisionFailureCode struct {
	value *PolicyDecisionFailureCode
	isSet bool
}

func (v NullablePolicyDecisionFailureCode) Get() *PolicyDecisionFailureCode {
	return v.value
}

func (v *NullablePolicyDecisionFailureCode) Set(val *PolicyDecisionFailureCode) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyDecisionFailureCode) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyDecisionFailureCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyDecisionFailureCode(val *PolicyDecisionFailureCode) *NullablePolicyDecisionFailureCode {
	return &NullablePolicyDecisionFailureCode{value: val, isSet: true}
}

func (v NullablePolicyDecisionFailureCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyDecisionFailureCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

