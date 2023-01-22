/*
Npcf_MBSPolicyControl API

MBS Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Npcf_MBSPolicyControl

import (
	"encoding/json"
	"fmt"
)

// MbsPccRuleStatus Possible values are - ACTIVE: Indicates that the MBS PCC rule(s) are successfully installed. - INACTIVE: Indicates that the MBS PCC rule(s) are removed. 
type MbsPccRuleStatus struct {
	MbsPccRuleStatusAnyOf *MbsPccRuleStatusAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *MbsPccRuleStatus) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into MbsPccRuleStatusAnyOf
	err = json.Unmarshal(data, &dst.MbsPccRuleStatusAnyOf);
	if err == nil {
		jsonMbsPccRuleStatusAnyOf, _ := json.Marshal(dst.MbsPccRuleStatusAnyOf)
		if string(jsonMbsPccRuleStatusAnyOf) == "{}" { // empty struct
			dst.MbsPccRuleStatusAnyOf = nil
		} else {
			return nil // data stored in dst.MbsPccRuleStatusAnyOf, return on the first match
		}
	} else {
		dst.MbsPccRuleStatusAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(MbsPccRuleStatus)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *MbsPccRuleStatus) MarshalJSON() ([]byte, error) {
	if src.MbsPccRuleStatusAnyOf != nil {
		return json.Marshal(&src.MbsPccRuleStatusAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableMbsPccRuleStatus struct {
	value *MbsPccRuleStatus
	isSet bool
}

func (v NullableMbsPccRuleStatus) Get() *MbsPccRuleStatus {
	return v.value
}

func (v *NullableMbsPccRuleStatus) Set(val *MbsPccRuleStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableMbsPccRuleStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableMbsPccRuleStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbsPccRuleStatus(val *MbsPccRuleStatus) *NullableMbsPccRuleStatus {
	return &NullableMbsPccRuleStatus{value: val, isSet: true}
}

func (v NullableMbsPccRuleStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbsPccRuleStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

