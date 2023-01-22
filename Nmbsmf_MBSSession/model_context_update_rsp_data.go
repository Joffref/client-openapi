/*
Nmbsmf-MBSSession

MB-SMF MBSSession Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nmbsmf_MBSSession

import (
	"encoding/json"
)

// ContextUpdateRspData Data within ContextUpdate Response
type ContextUpdateRspData struct {
	LlSsm *Ssm `json:"llSsm,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer. 
	CTeid *int32 `json:"cTeid,omitempty"`
	N2MbsSmInfo *N2MbsSmInfo `json:"n2MbsSmInfo,omitempty"`
}

// NewContextUpdateRspData instantiates a new ContextUpdateRspData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContextUpdateRspData() *ContextUpdateRspData {
	this := ContextUpdateRspData{}
	return &this
}

// NewContextUpdateRspDataWithDefaults instantiates a new ContextUpdateRspData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContextUpdateRspDataWithDefaults() *ContextUpdateRspData {
	this := ContextUpdateRspData{}
	return &this
}

// GetLlSsm returns the LlSsm field value if set, zero value otherwise.
func (o *ContextUpdateRspData) GetLlSsm() Ssm {
	if o == nil || isNil(o.LlSsm) {
		var ret Ssm
		return ret
	}
	return *o.LlSsm
}

// GetLlSsmOk returns a tuple with the LlSsm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextUpdateRspData) GetLlSsmOk() (*Ssm, bool) {
	if o == nil || isNil(o.LlSsm) {
    return nil, false
	}
	return o.LlSsm, true
}

// HasLlSsm returns a boolean if a field has been set.
func (o *ContextUpdateRspData) HasLlSsm() bool {
	if o != nil && !isNil(o.LlSsm) {
		return true
	}

	return false
}

// SetLlSsm gets a reference to the given Ssm and assigns it to the LlSsm field.
func (o *ContextUpdateRspData) SetLlSsm(v Ssm) {
	o.LlSsm = &v
}

// GetCTeid returns the CTeid field value if set, zero value otherwise.
func (o *ContextUpdateRspData) GetCTeid() int32 {
	if o == nil || isNil(o.CTeid) {
		var ret int32
		return ret
	}
	return *o.CTeid
}

// GetCTeidOk returns a tuple with the CTeid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextUpdateRspData) GetCTeidOk() (*int32, bool) {
	if o == nil || isNil(o.CTeid) {
    return nil, false
	}
	return o.CTeid, true
}

// HasCTeid returns a boolean if a field has been set.
func (o *ContextUpdateRspData) HasCTeid() bool {
	if o != nil && !isNil(o.CTeid) {
		return true
	}

	return false
}

// SetCTeid gets a reference to the given int32 and assigns it to the CTeid field.
func (o *ContextUpdateRspData) SetCTeid(v int32) {
	o.CTeid = &v
}

// GetN2MbsSmInfo returns the N2MbsSmInfo field value if set, zero value otherwise.
func (o *ContextUpdateRspData) GetN2MbsSmInfo() N2MbsSmInfo {
	if o == nil || isNil(o.N2MbsSmInfo) {
		var ret N2MbsSmInfo
		return ret
	}
	return *o.N2MbsSmInfo
}

// GetN2MbsSmInfoOk returns a tuple with the N2MbsSmInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextUpdateRspData) GetN2MbsSmInfoOk() (*N2MbsSmInfo, bool) {
	if o == nil || isNil(o.N2MbsSmInfo) {
    return nil, false
	}
	return o.N2MbsSmInfo, true
}

// HasN2MbsSmInfo returns a boolean if a field has been set.
func (o *ContextUpdateRspData) HasN2MbsSmInfo() bool {
	if o != nil && !isNil(o.N2MbsSmInfo) {
		return true
	}

	return false
}

// SetN2MbsSmInfo gets a reference to the given N2MbsSmInfo and assigns it to the N2MbsSmInfo field.
func (o *ContextUpdateRspData) SetN2MbsSmInfo(v N2MbsSmInfo) {
	o.N2MbsSmInfo = &v
}

func (o ContextUpdateRspData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.LlSsm) {
		toSerialize["llSsm"] = o.LlSsm
	}
	if !isNil(o.CTeid) {
		toSerialize["cTeid"] = o.CTeid
	}
	if !isNil(o.N2MbsSmInfo) {
		toSerialize["n2MbsSmInfo"] = o.N2MbsSmInfo
	}
	return json.Marshal(toSerialize)
}

type NullableContextUpdateRspData struct {
	value *ContextUpdateRspData
	isSet bool
}

func (v NullableContextUpdateRspData) Get() *ContextUpdateRspData {
	return v.value
}

func (v *NullableContextUpdateRspData) Set(val *ContextUpdateRspData) {
	v.value = val
	v.isSet = true
}

func (v NullableContextUpdateRspData) IsSet() bool {
	return v.isSet
}

func (v *NullableContextUpdateRspData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContextUpdateRspData(val *ContextUpdateRspData) *NullableContextUpdateRspData {
	return &NullableContextUpdateRspData{value: val, isSet: true}
}

func (v NullableContextUpdateRspData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContextUpdateRspData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

