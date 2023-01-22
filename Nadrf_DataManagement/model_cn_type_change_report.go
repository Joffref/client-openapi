/*
Nadrf_DataManagement

ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nadrf_DataManagement

import (
	"encoding/json"
)

// CnTypeChangeReport struct for CnTypeChangeReport
type CnTypeChangeReport struct {
	NewCnType CnType `json:"newCnType"`
	OldCnType *CnType `json:"oldCnType,omitempty"`
}

// NewCnTypeChangeReport instantiates a new CnTypeChangeReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCnTypeChangeReport(newCnType CnType) *CnTypeChangeReport {
	this := CnTypeChangeReport{}
	this.NewCnType = newCnType
	return &this
}

// NewCnTypeChangeReportWithDefaults instantiates a new CnTypeChangeReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCnTypeChangeReportWithDefaults() *CnTypeChangeReport {
	this := CnTypeChangeReport{}
	return &this
}

// GetNewCnType returns the NewCnType field value
func (o *CnTypeChangeReport) GetNewCnType() CnType {
	if o == nil {
		var ret CnType
		return ret
	}

	return o.NewCnType
}

// GetNewCnTypeOk returns a tuple with the NewCnType field value
// and a boolean to check if the value has been set.
func (o *CnTypeChangeReport) GetNewCnTypeOk() (*CnType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NewCnType, true
}

// SetNewCnType sets field value
func (o *CnTypeChangeReport) SetNewCnType(v CnType) {
	o.NewCnType = v
}

// GetOldCnType returns the OldCnType field value if set, zero value otherwise.
func (o *CnTypeChangeReport) GetOldCnType() CnType {
	if o == nil || isNil(o.OldCnType) {
		var ret CnType
		return ret
	}
	return *o.OldCnType
}

// GetOldCnTypeOk returns a tuple with the OldCnType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CnTypeChangeReport) GetOldCnTypeOk() (*CnType, bool) {
	if o == nil || isNil(o.OldCnType) {
    return nil, false
	}
	return o.OldCnType, true
}

// HasOldCnType returns a boolean if a field has been set.
func (o *CnTypeChangeReport) HasOldCnType() bool {
	if o != nil && !isNil(o.OldCnType) {
		return true
	}

	return false
}

// SetOldCnType gets a reference to the given CnType and assigns it to the OldCnType field.
func (o *CnTypeChangeReport) SetOldCnType(v CnType) {
	o.OldCnType = &v
}

func (o CnTypeChangeReport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["newCnType"] = o.NewCnType
	}
	if !isNil(o.OldCnType) {
		toSerialize["oldCnType"] = o.OldCnType
	}
	return json.Marshal(toSerialize)
}

type NullableCnTypeChangeReport struct {
	value *CnTypeChangeReport
	isSet bool
}

func (v NullableCnTypeChangeReport) Get() *CnTypeChangeReport {
	return v.value
}

func (v *NullableCnTypeChangeReport) Set(val *CnTypeChangeReport) {
	v.value = val
	v.isSet = true
}

func (v NullableCnTypeChangeReport) IsSet() bool {
	return v.isSet
}

func (v *NullableCnTypeChangeReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCnTypeChangeReport(val *CnTypeChangeReport) *NullableCnTypeChangeReport {
	return &NullableCnTypeChangeReport{value: val, isSet: true}
}

func (v NullableCnTypeChangeReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCnTypeChangeReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

