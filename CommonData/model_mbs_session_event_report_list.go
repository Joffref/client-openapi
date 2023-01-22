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

// MbsSessionEventReportList MBS session event report list
type MbsSessionEventReportList struct {
	EventReportList []MbsSessionEventReport `json:"eventReportList"`
	NotifyCorrelationId *string `json:"notifyCorrelationId,omitempty"`
}

// NewMbsSessionEventReportList instantiates a new MbsSessionEventReportList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMbsSessionEventReportList(eventReportList []MbsSessionEventReport) *MbsSessionEventReportList {
	this := MbsSessionEventReportList{}
	this.EventReportList = eventReportList
	return &this
}

// NewMbsSessionEventReportListWithDefaults instantiates a new MbsSessionEventReportList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMbsSessionEventReportListWithDefaults() *MbsSessionEventReportList {
	this := MbsSessionEventReportList{}
	return &this
}

// GetEventReportList returns the EventReportList field value
func (o *MbsSessionEventReportList) GetEventReportList() []MbsSessionEventReport {
	if o == nil {
		var ret []MbsSessionEventReport
		return ret
	}

	return o.EventReportList
}

// GetEventReportListOk returns a tuple with the EventReportList field value
// and a boolean to check if the value has been set.
func (o *MbsSessionEventReportList) GetEventReportListOk() ([]MbsSessionEventReport, bool) {
	if o == nil {
    return nil, false
	}
	return o.EventReportList, true
}

// SetEventReportList sets field value
func (o *MbsSessionEventReportList) SetEventReportList(v []MbsSessionEventReport) {
	o.EventReportList = v
}

// GetNotifyCorrelationId returns the NotifyCorrelationId field value if set, zero value otherwise.
func (o *MbsSessionEventReportList) GetNotifyCorrelationId() string {
	if o == nil || isNil(o.NotifyCorrelationId) {
		var ret string
		return ret
	}
	return *o.NotifyCorrelationId
}

// GetNotifyCorrelationIdOk returns a tuple with the NotifyCorrelationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsSessionEventReportList) GetNotifyCorrelationIdOk() (*string, bool) {
	if o == nil || isNil(o.NotifyCorrelationId) {
    return nil, false
	}
	return o.NotifyCorrelationId, true
}

// HasNotifyCorrelationId returns a boolean if a field has been set.
func (o *MbsSessionEventReportList) HasNotifyCorrelationId() bool {
	if o != nil && !isNil(o.NotifyCorrelationId) {
		return true
	}

	return false
}

// SetNotifyCorrelationId gets a reference to the given string and assigns it to the NotifyCorrelationId field.
func (o *MbsSessionEventReportList) SetNotifyCorrelationId(v string) {
	o.NotifyCorrelationId = &v
}

func (o MbsSessionEventReportList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["eventReportList"] = o.EventReportList
	}
	if !isNil(o.NotifyCorrelationId) {
		toSerialize["notifyCorrelationId"] = o.NotifyCorrelationId
	}
	return json.Marshal(toSerialize)
}

type NullableMbsSessionEventReportList struct {
	value *MbsSessionEventReportList
	isSet bool
}

func (v NullableMbsSessionEventReportList) Get() *MbsSessionEventReportList {
	return v.value
}

func (v *NullableMbsSessionEventReportList) Set(val *MbsSessionEventReportList) {
	v.value = val
	v.isSet = true
}

func (v NullableMbsSessionEventReportList) IsSet() bool {
	return v.isSet
}

func (v *NullableMbsSessionEventReportList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbsSessionEventReportList(val *MbsSessionEventReportList) *NullableMbsSessionEventReportList {
	return &NullableMbsSessionEventReportList{value: val, isSet: true}
}

func (v NullableMbsSessionEventReportList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbsSessionEventReportList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

