/*
Nsmf_PDUSession

SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nsmf_PDUSession

import (
	"encoding/json"
	"time"
)

// SmContextCreateError Error within Create SM Context Response
type SmContextCreateError struct {
	Error ExtProblemDetails `json:"error"`
	N1SmMsg *RefToBinaryData `json:"n1SmMsg,omitempty"`
	N2SmInfo *RefToBinaryData `json:"n2SmInfo,omitempty"`
	N2SmInfoType *N2SmInfoType `json:"n2SmInfoType,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	RecoveryTime *time.Time `json:"recoveryTime,omitempty"`
}

// NewSmContextCreateError instantiates a new SmContextCreateError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmContextCreateError(error_ ExtProblemDetails) *SmContextCreateError {
	this := SmContextCreateError{}
	this.Error = error_
	return &this
}

// NewSmContextCreateErrorWithDefaults instantiates a new SmContextCreateError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmContextCreateErrorWithDefaults() *SmContextCreateError {
	this := SmContextCreateError{}
	return &this
}

// GetError returns the Error field value
func (o *SmContextCreateError) GetError() ExtProblemDetails {
	if o == nil {
		var ret ExtProblemDetails
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *SmContextCreateError) GetErrorOk() (*ExtProblemDetails, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *SmContextCreateError) SetError(v ExtProblemDetails) {
	o.Error = v
}

// GetN1SmMsg returns the N1SmMsg field value if set, zero value otherwise.
func (o *SmContextCreateError) GetN1SmMsg() RefToBinaryData {
	if o == nil || isNil(o.N1SmMsg) {
		var ret RefToBinaryData
		return ret
	}
	return *o.N1SmMsg
}

// GetN1SmMsgOk returns a tuple with the N1SmMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmContextCreateError) GetN1SmMsgOk() (*RefToBinaryData, bool) {
	if o == nil || isNil(o.N1SmMsg) {
    return nil, false
	}
	return o.N1SmMsg, true
}

// HasN1SmMsg returns a boolean if a field has been set.
func (o *SmContextCreateError) HasN1SmMsg() bool {
	if o != nil && !isNil(o.N1SmMsg) {
		return true
	}

	return false
}

// SetN1SmMsg gets a reference to the given RefToBinaryData and assigns it to the N1SmMsg field.
func (o *SmContextCreateError) SetN1SmMsg(v RefToBinaryData) {
	o.N1SmMsg = &v
}

// GetN2SmInfo returns the N2SmInfo field value if set, zero value otherwise.
func (o *SmContextCreateError) GetN2SmInfo() RefToBinaryData {
	if o == nil || isNil(o.N2SmInfo) {
		var ret RefToBinaryData
		return ret
	}
	return *o.N2SmInfo
}

// GetN2SmInfoOk returns a tuple with the N2SmInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmContextCreateError) GetN2SmInfoOk() (*RefToBinaryData, bool) {
	if o == nil || isNil(o.N2SmInfo) {
    return nil, false
	}
	return o.N2SmInfo, true
}

// HasN2SmInfo returns a boolean if a field has been set.
func (o *SmContextCreateError) HasN2SmInfo() bool {
	if o != nil && !isNil(o.N2SmInfo) {
		return true
	}

	return false
}

// SetN2SmInfo gets a reference to the given RefToBinaryData and assigns it to the N2SmInfo field.
func (o *SmContextCreateError) SetN2SmInfo(v RefToBinaryData) {
	o.N2SmInfo = &v
}

// GetN2SmInfoType returns the N2SmInfoType field value if set, zero value otherwise.
func (o *SmContextCreateError) GetN2SmInfoType() N2SmInfoType {
	if o == nil || isNil(o.N2SmInfoType) {
		var ret N2SmInfoType
		return ret
	}
	return *o.N2SmInfoType
}

// GetN2SmInfoTypeOk returns a tuple with the N2SmInfoType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmContextCreateError) GetN2SmInfoTypeOk() (*N2SmInfoType, bool) {
	if o == nil || isNil(o.N2SmInfoType) {
    return nil, false
	}
	return o.N2SmInfoType, true
}

// HasN2SmInfoType returns a boolean if a field has been set.
func (o *SmContextCreateError) HasN2SmInfoType() bool {
	if o != nil && !isNil(o.N2SmInfoType) {
		return true
	}

	return false
}

// SetN2SmInfoType gets a reference to the given N2SmInfoType and assigns it to the N2SmInfoType field.
func (o *SmContextCreateError) SetN2SmInfoType(v N2SmInfoType) {
	o.N2SmInfoType = &v
}

// GetRecoveryTime returns the RecoveryTime field value if set, zero value otherwise.
func (o *SmContextCreateError) GetRecoveryTime() time.Time {
	if o == nil || isNil(o.RecoveryTime) {
		var ret time.Time
		return ret
	}
	return *o.RecoveryTime
}

// GetRecoveryTimeOk returns a tuple with the RecoveryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmContextCreateError) GetRecoveryTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.RecoveryTime) {
    return nil, false
	}
	return o.RecoveryTime, true
}

// HasRecoveryTime returns a boolean if a field has been set.
func (o *SmContextCreateError) HasRecoveryTime() bool {
	if o != nil && !isNil(o.RecoveryTime) {
		return true
	}

	return false
}

// SetRecoveryTime gets a reference to the given time.Time and assigns it to the RecoveryTime field.
func (o *SmContextCreateError) SetRecoveryTime(v time.Time) {
	o.RecoveryTime = &v
}

func (o SmContextCreateError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["error"] = o.Error
	}
	if !isNil(o.N1SmMsg) {
		toSerialize["n1SmMsg"] = o.N1SmMsg
	}
	if !isNil(o.N2SmInfo) {
		toSerialize["n2SmInfo"] = o.N2SmInfo
	}
	if !isNil(o.N2SmInfoType) {
		toSerialize["n2SmInfoType"] = o.N2SmInfoType
	}
	if !isNil(o.RecoveryTime) {
		toSerialize["recoveryTime"] = o.RecoveryTime
	}
	return json.Marshal(toSerialize)
}

type NullableSmContextCreateError struct {
	value *SmContextCreateError
	isSet bool
}

func (v NullableSmContextCreateError) Get() *SmContextCreateError {
	return v.value
}

func (v *NullableSmContextCreateError) Set(val *SmContextCreateError) {
	v.value = val
	v.isSet = true
}

func (v NullableSmContextCreateError) IsSet() bool {
	return v.isSet
}

func (v *NullableSmContextCreateError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmContextCreateError(val *SmContextCreateError) *NullableSmContextCreateError {
	return &NullableSmContextCreateError{value: val, isSet: true}
}

func (v NullableSmContextCreateError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmContextCreateError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

