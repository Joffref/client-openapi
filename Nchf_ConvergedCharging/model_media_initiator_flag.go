/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nchf_ConvergedCharging

import (
	"encoding/json"
	"fmt"
)

// MediaInitiatorFlag struct for MediaInitiatorFlag
type MediaInitiatorFlag struct {
	MediaInitiatorFlagAnyOf *MediaInitiatorFlagAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *MediaInitiatorFlag) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into MediaInitiatorFlagAnyOf
	err = json.Unmarshal(data, &dst.MediaInitiatorFlagAnyOf);
	if err == nil {
		jsonMediaInitiatorFlagAnyOf, _ := json.Marshal(dst.MediaInitiatorFlagAnyOf)
		if string(jsonMediaInitiatorFlagAnyOf) == "{}" { // empty struct
			dst.MediaInitiatorFlagAnyOf = nil
		} else {
			return nil // data stored in dst.MediaInitiatorFlagAnyOf, return on the first match
		}
	} else {
		dst.MediaInitiatorFlagAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(MediaInitiatorFlag)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *MediaInitiatorFlag) MarshalJSON() ([]byte, error) {
	if src.MediaInitiatorFlagAnyOf != nil {
		return json.Marshal(&src.MediaInitiatorFlagAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableMediaInitiatorFlag struct {
	value *MediaInitiatorFlag
	isSet bool
}

func (v NullableMediaInitiatorFlag) Get() *MediaInitiatorFlag {
	return v.value
}

func (v *NullableMediaInitiatorFlag) Set(val *MediaInitiatorFlag) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaInitiatorFlag) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaInitiatorFlag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaInitiatorFlag(val *MediaInitiatorFlag) *NullableMediaInitiatorFlag {
	return &NullableMediaInitiatorFlag{value: val, isSet: true}
}

func (v NullableMediaInitiatorFlag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaInitiatorFlag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

