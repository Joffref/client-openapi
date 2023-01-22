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

// MbsMediaComp Represents an MBS Media Component.
type MbsMediaComp struct {
	MbsMedCompNum int32 `json:"mbsMedCompNum"`
	MbsFlowDescs []string `json:"mbsFlowDescs,omitempty"`
	MbsSdfResPrio *ReservPriority `json:"mbsSdfResPrio,omitempty"`
	MbsMediaInfo *MbsMediaInfo `json:"mbsMediaInfo,omitempty"`
	QosRef *string `json:"qosRef,omitempty"`
	MbsQoSReq *MbsQoSReq `json:"mbsQoSReq,omitempty"`
}

// NewMbsMediaComp instantiates a new MbsMediaComp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMbsMediaComp(mbsMedCompNum int32) *MbsMediaComp {
	this := MbsMediaComp{}
	this.MbsMedCompNum = mbsMedCompNum
	return &this
}

// NewMbsMediaCompWithDefaults instantiates a new MbsMediaComp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMbsMediaCompWithDefaults() *MbsMediaComp {
	this := MbsMediaComp{}
	return &this
}

// GetMbsMedCompNum returns the MbsMedCompNum field value
func (o *MbsMediaComp) GetMbsMedCompNum() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MbsMedCompNum
}

// GetMbsMedCompNumOk returns a tuple with the MbsMedCompNum field value
// and a boolean to check if the value has been set.
func (o *MbsMediaComp) GetMbsMedCompNumOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MbsMedCompNum, true
}

// SetMbsMedCompNum sets field value
func (o *MbsMediaComp) SetMbsMedCompNum(v int32) {
	o.MbsMedCompNum = v
}

// GetMbsFlowDescs returns the MbsFlowDescs field value if set, zero value otherwise.
func (o *MbsMediaComp) GetMbsFlowDescs() []string {
	if o == nil || isNil(o.MbsFlowDescs) {
		var ret []string
		return ret
	}
	return o.MbsFlowDescs
}

// GetMbsFlowDescsOk returns a tuple with the MbsFlowDescs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsMediaComp) GetMbsFlowDescsOk() ([]string, bool) {
	if o == nil || isNil(o.MbsFlowDescs) {
    return nil, false
	}
	return o.MbsFlowDescs, true
}

// HasMbsFlowDescs returns a boolean if a field has been set.
func (o *MbsMediaComp) HasMbsFlowDescs() bool {
	if o != nil && !isNil(o.MbsFlowDescs) {
		return true
	}

	return false
}

// SetMbsFlowDescs gets a reference to the given []string and assigns it to the MbsFlowDescs field.
func (o *MbsMediaComp) SetMbsFlowDescs(v []string) {
	o.MbsFlowDescs = v
}

// GetMbsSdfResPrio returns the MbsSdfResPrio field value if set, zero value otherwise.
func (o *MbsMediaComp) GetMbsSdfResPrio() ReservPriority {
	if o == nil || isNil(o.MbsSdfResPrio) {
		var ret ReservPriority
		return ret
	}
	return *o.MbsSdfResPrio
}

// GetMbsSdfResPrioOk returns a tuple with the MbsSdfResPrio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsMediaComp) GetMbsSdfResPrioOk() (*ReservPriority, bool) {
	if o == nil || isNil(o.MbsSdfResPrio) {
    return nil, false
	}
	return o.MbsSdfResPrio, true
}

// HasMbsSdfResPrio returns a boolean if a field has been set.
func (o *MbsMediaComp) HasMbsSdfResPrio() bool {
	if o != nil && !isNil(o.MbsSdfResPrio) {
		return true
	}

	return false
}

// SetMbsSdfResPrio gets a reference to the given ReservPriority and assigns it to the MbsSdfResPrio field.
func (o *MbsMediaComp) SetMbsSdfResPrio(v ReservPriority) {
	o.MbsSdfResPrio = &v
}

// GetMbsMediaInfo returns the MbsMediaInfo field value if set, zero value otherwise.
func (o *MbsMediaComp) GetMbsMediaInfo() MbsMediaInfo {
	if o == nil || isNil(o.MbsMediaInfo) {
		var ret MbsMediaInfo
		return ret
	}
	return *o.MbsMediaInfo
}

// GetMbsMediaInfoOk returns a tuple with the MbsMediaInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsMediaComp) GetMbsMediaInfoOk() (*MbsMediaInfo, bool) {
	if o == nil || isNil(o.MbsMediaInfo) {
    return nil, false
	}
	return o.MbsMediaInfo, true
}

// HasMbsMediaInfo returns a boolean if a field has been set.
func (o *MbsMediaComp) HasMbsMediaInfo() bool {
	if o != nil && !isNil(o.MbsMediaInfo) {
		return true
	}

	return false
}

// SetMbsMediaInfo gets a reference to the given MbsMediaInfo and assigns it to the MbsMediaInfo field.
func (o *MbsMediaComp) SetMbsMediaInfo(v MbsMediaInfo) {
	o.MbsMediaInfo = &v
}

// GetQosRef returns the QosRef field value if set, zero value otherwise.
func (o *MbsMediaComp) GetQosRef() string {
	if o == nil || isNil(o.QosRef) {
		var ret string
		return ret
	}
	return *o.QosRef
}

// GetQosRefOk returns a tuple with the QosRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsMediaComp) GetQosRefOk() (*string, bool) {
	if o == nil || isNil(o.QosRef) {
    return nil, false
	}
	return o.QosRef, true
}

// HasQosRef returns a boolean if a field has been set.
func (o *MbsMediaComp) HasQosRef() bool {
	if o != nil && !isNil(o.QosRef) {
		return true
	}

	return false
}

// SetQosRef gets a reference to the given string and assigns it to the QosRef field.
func (o *MbsMediaComp) SetQosRef(v string) {
	o.QosRef = &v
}

// GetMbsQoSReq returns the MbsQoSReq field value if set, zero value otherwise.
func (o *MbsMediaComp) GetMbsQoSReq() MbsQoSReq {
	if o == nil || isNil(o.MbsQoSReq) {
		var ret MbsQoSReq
		return ret
	}
	return *o.MbsQoSReq
}

// GetMbsQoSReqOk returns a tuple with the MbsQoSReq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsMediaComp) GetMbsQoSReqOk() (*MbsQoSReq, bool) {
	if o == nil || isNil(o.MbsQoSReq) {
    return nil, false
	}
	return o.MbsQoSReq, true
}

// HasMbsQoSReq returns a boolean if a field has been set.
func (o *MbsMediaComp) HasMbsQoSReq() bool {
	if o != nil && !isNil(o.MbsQoSReq) {
		return true
	}

	return false
}

// SetMbsQoSReq gets a reference to the given MbsQoSReq and assigns it to the MbsQoSReq field.
func (o *MbsMediaComp) SetMbsQoSReq(v MbsQoSReq) {
	o.MbsQoSReq = &v
}

func (o MbsMediaComp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mbsMedCompNum"] = o.MbsMedCompNum
	}
	if !isNil(o.MbsFlowDescs) {
		toSerialize["mbsFlowDescs"] = o.MbsFlowDescs
	}
	if !isNil(o.MbsSdfResPrio) {
		toSerialize["mbsSdfResPrio"] = o.MbsSdfResPrio
	}
	if !isNil(o.MbsMediaInfo) {
		toSerialize["mbsMediaInfo"] = o.MbsMediaInfo
	}
	if !isNil(o.QosRef) {
		toSerialize["qosRef"] = o.QosRef
	}
	if !isNil(o.MbsQoSReq) {
		toSerialize["mbsQoSReq"] = o.MbsQoSReq
	}
	return json.Marshal(toSerialize)
}

type NullableMbsMediaComp struct {
	value *MbsMediaComp
	isSet bool
}

func (v NullableMbsMediaComp) Get() *MbsMediaComp {
	return v.value
}

func (v *NullableMbsMediaComp) Set(val *MbsMediaComp) {
	v.value = val
	v.isSet = true
}

func (v NullableMbsMediaComp) IsSet() bool {
	return v.isSet
}

func (v *NullableMbsMediaComp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbsMediaComp(val *MbsMediaComp) *NullableMbsMediaComp {
	return &NullableMbsMediaComp{value: val, isSet: true}
}

func (v NullableMbsMediaComp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbsMediaComp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

