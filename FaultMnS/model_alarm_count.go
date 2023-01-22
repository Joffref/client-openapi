/*
Fault Supervision MnS

OAS 3.0.1 definition of the Fault Supervision MnS © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package FaultMnS

import (
	"encoding/json"
)

// AlarmCount struct for AlarmCount
type AlarmCount struct {
	CriticalCount int32 `json:"criticalCount"`
	MajorCount int32 `json:"majorCount"`
	MinorCount int32 `json:"minorCount"`
	WarningCount int32 `json:"warningCount"`
	IndeterminateCount int32 `json:"indeterminateCount"`
	ClearedCount int32 `json:"clearedCount"`
}

// NewAlarmCount instantiates a new AlarmCount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlarmCount(criticalCount int32, majorCount int32, minorCount int32, warningCount int32, indeterminateCount int32, clearedCount int32) *AlarmCount {
	this := AlarmCount{}
	this.CriticalCount = criticalCount
	this.MajorCount = majorCount
	this.MinorCount = minorCount
	this.WarningCount = warningCount
	this.IndeterminateCount = indeterminateCount
	this.ClearedCount = clearedCount
	return &this
}

// NewAlarmCountWithDefaults instantiates a new AlarmCount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlarmCountWithDefaults() *AlarmCount {
	this := AlarmCount{}
	return &this
}

// GetCriticalCount returns the CriticalCount field value
func (o *AlarmCount) GetCriticalCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CriticalCount
}

// GetCriticalCountOk returns a tuple with the CriticalCount field value
// and a boolean to check if the value has been set.
func (o *AlarmCount) GetCriticalCountOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CriticalCount, true
}

// SetCriticalCount sets field value
func (o *AlarmCount) SetCriticalCount(v int32) {
	o.CriticalCount = v
}

// GetMajorCount returns the MajorCount field value
func (o *AlarmCount) GetMajorCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MajorCount
}

// GetMajorCountOk returns a tuple with the MajorCount field value
// and a boolean to check if the value has been set.
func (o *AlarmCount) GetMajorCountOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MajorCount, true
}

// SetMajorCount sets field value
func (o *AlarmCount) SetMajorCount(v int32) {
	o.MajorCount = v
}

// GetMinorCount returns the MinorCount field value
func (o *AlarmCount) GetMinorCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MinorCount
}

// GetMinorCountOk returns a tuple with the MinorCount field value
// and a boolean to check if the value has been set.
func (o *AlarmCount) GetMinorCountOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MinorCount, true
}

// SetMinorCount sets field value
func (o *AlarmCount) SetMinorCount(v int32) {
	o.MinorCount = v
}

// GetWarningCount returns the WarningCount field value
func (o *AlarmCount) GetWarningCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.WarningCount
}

// GetWarningCountOk returns a tuple with the WarningCount field value
// and a boolean to check if the value has been set.
func (o *AlarmCount) GetWarningCountOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.WarningCount, true
}

// SetWarningCount sets field value
func (o *AlarmCount) SetWarningCount(v int32) {
	o.WarningCount = v
}

// GetIndeterminateCount returns the IndeterminateCount field value
func (o *AlarmCount) GetIndeterminateCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IndeterminateCount
}

// GetIndeterminateCountOk returns a tuple with the IndeterminateCount field value
// and a boolean to check if the value has been set.
func (o *AlarmCount) GetIndeterminateCountOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.IndeterminateCount, true
}

// SetIndeterminateCount sets field value
func (o *AlarmCount) SetIndeterminateCount(v int32) {
	o.IndeterminateCount = v
}

// GetClearedCount returns the ClearedCount field value
func (o *AlarmCount) GetClearedCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ClearedCount
}

// GetClearedCountOk returns a tuple with the ClearedCount field value
// and a boolean to check if the value has been set.
func (o *AlarmCount) GetClearedCountOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ClearedCount, true
}

// SetClearedCount sets field value
func (o *AlarmCount) SetClearedCount(v int32) {
	o.ClearedCount = v
}

func (o AlarmCount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["criticalCount"] = o.CriticalCount
	}
	if true {
		toSerialize["majorCount"] = o.MajorCount
	}
	if true {
		toSerialize["minorCount"] = o.MinorCount
	}
	if true {
		toSerialize["warningCount"] = o.WarningCount
	}
	if true {
		toSerialize["indeterminateCount"] = o.IndeterminateCount
	}
	if true {
		toSerialize["clearedCount"] = o.ClearedCount
	}
	return json.Marshal(toSerialize)
}

type NullableAlarmCount struct {
	value *AlarmCount
	isSet bool
}

func (v NullableAlarmCount) Get() *AlarmCount {
	return v.value
}

func (v *NullableAlarmCount) Set(val *AlarmCount) {
	v.value = val
	v.isSet = true
}

func (v NullableAlarmCount) IsSet() bool {
	return v.isSet
}

func (v *NullableAlarmCount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlarmCount(val *AlarmCount) *NullableAlarmCount {
	return &NullableAlarmCount{value: val, isSet: true}
}

func (v NullableAlarmCount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlarmCount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

