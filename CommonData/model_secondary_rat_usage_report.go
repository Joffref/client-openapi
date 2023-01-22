/*
Common Data Types

Common Data Types for Service Based Interfaces.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.   

API version: 1.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package CommonData

import (
	"encoding/json"
)

// SecondaryRatUsageReport Secondary RAT Usage Report to report usage data for a secondary RAT for QoS flows.
type SecondaryRatUsageReport struct {
	SecondaryRatType RatType `json:"secondaryRatType"`
	QosFlowsUsageData []QosFlowUsageReport `json:"qosFlowsUsageData"`
}

// NewSecondaryRatUsageReport instantiates a new SecondaryRatUsageReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecondaryRatUsageReport(secondaryRatType RatType, qosFlowsUsageData []QosFlowUsageReport) *SecondaryRatUsageReport {
	this := SecondaryRatUsageReport{}
	this.SecondaryRatType = secondaryRatType
	this.QosFlowsUsageData = qosFlowsUsageData
	return &this
}

// NewSecondaryRatUsageReportWithDefaults instantiates a new SecondaryRatUsageReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecondaryRatUsageReportWithDefaults() *SecondaryRatUsageReport {
	this := SecondaryRatUsageReport{}
	return &this
}

// GetSecondaryRatType returns the SecondaryRatType field value
func (o *SecondaryRatUsageReport) GetSecondaryRatType() RatType {
	if o == nil {
		var ret RatType
		return ret
	}

	return o.SecondaryRatType
}

// GetSecondaryRatTypeOk returns a tuple with the SecondaryRatType field value
// and a boolean to check if the value has been set.
func (o *SecondaryRatUsageReport) GetSecondaryRatTypeOk() (*RatType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SecondaryRatType, true
}

// SetSecondaryRatType sets field value
func (o *SecondaryRatUsageReport) SetSecondaryRatType(v RatType) {
	o.SecondaryRatType = v
}

// GetQosFlowsUsageData returns the QosFlowsUsageData field value
func (o *SecondaryRatUsageReport) GetQosFlowsUsageData() []QosFlowUsageReport {
	if o == nil {
		var ret []QosFlowUsageReport
		return ret
	}

	return o.QosFlowsUsageData
}

// GetQosFlowsUsageDataOk returns a tuple with the QosFlowsUsageData field value
// and a boolean to check if the value has been set.
func (o *SecondaryRatUsageReport) GetQosFlowsUsageDataOk() ([]QosFlowUsageReport, bool) {
	if o == nil {
    return nil, false
	}
	return o.QosFlowsUsageData, true
}

// SetQosFlowsUsageData sets field value
func (o *SecondaryRatUsageReport) SetQosFlowsUsageData(v []QosFlowUsageReport) {
	o.QosFlowsUsageData = v
}

func (o SecondaryRatUsageReport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["secondaryRatType"] = o.SecondaryRatType
	}
	if true {
		toSerialize["qosFlowsUsageData"] = o.QosFlowsUsageData
	}
	return json.Marshal(toSerialize)
}

type NullableSecondaryRatUsageReport struct {
	value *SecondaryRatUsageReport
	isSet bool
}

func (v NullableSecondaryRatUsageReport) Get() *SecondaryRatUsageReport {
	return v.value
}

func (v *NullableSecondaryRatUsageReport) Set(val *SecondaryRatUsageReport) {
	v.value = val
	v.isSet = true
}

func (v NullableSecondaryRatUsageReport) IsSet() bool {
	return v.isSet
}

func (v *NullableSecondaryRatUsageReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecondaryRatUsageReport(val *SecondaryRatUsageReport) *NullableSecondaryRatUsageReport {
	return &NullableSecondaryRatUsageReport{value: val, isSet: true}
}

func (v NullableSecondaryRatUsageReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecondaryRatUsageReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

