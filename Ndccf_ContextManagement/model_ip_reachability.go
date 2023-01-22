/*
Ndccf_ContextManagement

DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Ndccf_ContextManagement

import (
	"encoding/json"
	"fmt"
)

// IpReachability Indicates the type(s) of IP addresses reachable via an SCP
type IpReachability struct {
	IpReachabilityAnyOf *IpReachabilityAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *IpReachability) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into IpReachabilityAnyOf
	err = json.Unmarshal(data, &dst.IpReachabilityAnyOf);
	if err == nil {
		jsonIpReachabilityAnyOf, _ := json.Marshal(dst.IpReachabilityAnyOf)
		if string(jsonIpReachabilityAnyOf) == "{}" { // empty struct
			dst.IpReachabilityAnyOf = nil
		} else {
			return nil // data stored in dst.IpReachabilityAnyOf, return on the first match
		}
	} else {
		dst.IpReachabilityAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(IpReachability)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *IpReachability) MarshalJSON() ([]byte, error) {
	if src.IpReachabilityAnyOf != nil {
		return json.Marshal(&src.IpReachabilityAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableIpReachability struct {
	value *IpReachability
	isSet bool
}

func (v NullableIpReachability) Get() *IpReachability {
	return v.value
}

func (v *NullableIpReachability) Set(val *IpReachability) {
	v.value = val
	v.isSet = true
}

func (v NullableIpReachability) IsSet() bool {
	return v.isSet
}

func (v *NullableIpReachability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpReachability(val *IpReachability) *NullableIpReachability {
	return &NullableIpReachability{value: val, isSet: true}
}

func (v NullableIpReachability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpReachability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

