/*
3gpp-analyticsexposure

API for Analytics Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package AnalyticsExposure

import (
	"encoding/json"
	"fmt"
)

// NwdafFailureCode Possible values are: - UNAVAILABLE_DATA: Indicates the requested statistics information for the event is rejected since necessary data to perform the service is unavailable. - BOTH_STAT_PRED_NOT_ALLOWED: Indicates the requested analysis information for the event is rejected since the start time is in the past and the end time is in the future, which means the NF service consumer requested both statistics and prediction for the analytics. - UNSATISFIED_REQUESTED_ANALYTICS_TIME: Indicates that the requested event is rejected since the analytics information is not ready when the time indicated by the \"timeAnaNeeded\" attribute (as provided during the creation or modification of subscription) is reached. - OTHER: Indicates the requested analysis information for the event is rejected due to other reasons.  
type NwdafFailureCode struct {
	AnalyticsFailureCodeAnyOf *AnalyticsFailureCodeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NwdafFailureCode) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AnalyticsFailureCodeAnyOf
	err = json.Unmarshal(data, &dst.AnalyticsFailureCodeAnyOf);
	if err == nil {
		jsonAnalyticsFailureCodeAnyOf, _ := json.Marshal(dst.AnalyticsFailureCodeAnyOf)
		if string(jsonAnalyticsFailureCodeAnyOf) == "{}" { // empty struct
			dst.AnalyticsFailureCodeAnyOf = nil
		} else {
			return nil // data stored in dst.AnalyticsFailureCodeAnyOf, return on the first match
		}
	} else {
		dst.AnalyticsFailureCodeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(NwdafFailureCode)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NwdafFailureCode) MarshalJSON() ([]byte, error) {
	if src.AnalyticsFailureCodeAnyOf != nil {
		return json.Marshal(&src.AnalyticsFailureCodeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNwdafFailureCode struct {
	value *NwdafFailureCode
	isSet bool
}

func (v NullableNwdafFailureCode) Get() *NwdafFailureCode {
	return v.value
}

func (v *NullableNwdafFailureCode) Set(val *NwdafFailureCode) {
	v.value = val
	v.isSet = true
}

func (v NullableNwdafFailureCode) IsSet() bool {
	return v.isSet
}

func (v *NullableNwdafFailureCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNwdafFailureCode(val *NwdafFailureCode) *NullableNwdafFailureCode {
	return &NullableNwdafFailureCode{value: val, isSet: true}
}

func (v NullableNwdafFailureCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNwdafFailureCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

