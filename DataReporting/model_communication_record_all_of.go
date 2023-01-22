/*
3gpp-data-reporting

API for 3GPP Data Reporting.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package DataReporting

import (
	"encoding/json"
)

// CommunicationRecordAllOf struct for CommunicationRecordAllOf
type CommunicationRecordAllOf struct {
	TimeInterval TimeWindow `json:"timeInterval"`
	// Unsigned integer identifying a volume in units of bytes.
	UplinkVolume *int64 `json:"uplinkVolume,omitempty"`
	// Unsigned integer identifying a volume in units of bytes.
	DownlinkVolume *int64 `json:"downlinkVolume,omitempty"`
}

// NewCommunicationRecordAllOf instantiates a new CommunicationRecordAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommunicationRecordAllOf(timeInterval TimeWindow) *CommunicationRecordAllOf {
	this := CommunicationRecordAllOf{}
	this.TimeInterval = timeInterval
	return &this
}

// NewCommunicationRecordAllOfWithDefaults instantiates a new CommunicationRecordAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommunicationRecordAllOfWithDefaults() *CommunicationRecordAllOf {
	this := CommunicationRecordAllOf{}
	return &this
}

// GetTimeInterval returns the TimeInterval field value
func (o *CommunicationRecordAllOf) GetTimeInterval() TimeWindow {
	if o == nil {
		var ret TimeWindow
		return ret
	}

	return o.TimeInterval
}

// GetTimeIntervalOk returns a tuple with the TimeInterval field value
// and a boolean to check if the value has been set.
func (o *CommunicationRecordAllOf) GetTimeIntervalOk() (*TimeWindow, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TimeInterval, true
}

// SetTimeInterval sets field value
func (o *CommunicationRecordAllOf) SetTimeInterval(v TimeWindow) {
	o.TimeInterval = v
}

// GetUplinkVolume returns the UplinkVolume field value if set, zero value otherwise.
func (o *CommunicationRecordAllOf) GetUplinkVolume() int64 {
	if o == nil || isNil(o.UplinkVolume) {
		var ret int64
		return ret
	}
	return *o.UplinkVolume
}

// GetUplinkVolumeOk returns a tuple with the UplinkVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationRecordAllOf) GetUplinkVolumeOk() (*int64, bool) {
	if o == nil || isNil(o.UplinkVolume) {
    return nil, false
	}
	return o.UplinkVolume, true
}

// HasUplinkVolume returns a boolean if a field has been set.
func (o *CommunicationRecordAllOf) HasUplinkVolume() bool {
	if o != nil && !isNil(o.UplinkVolume) {
		return true
	}

	return false
}

// SetUplinkVolume gets a reference to the given int64 and assigns it to the UplinkVolume field.
func (o *CommunicationRecordAllOf) SetUplinkVolume(v int64) {
	o.UplinkVolume = &v
}

// GetDownlinkVolume returns the DownlinkVolume field value if set, zero value otherwise.
func (o *CommunicationRecordAllOf) GetDownlinkVolume() int64 {
	if o == nil || isNil(o.DownlinkVolume) {
		var ret int64
		return ret
	}
	return *o.DownlinkVolume
}

// GetDownlinkVolumeOk returns a tuple with the DownlinkVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationRecordAllOf) GetDownlinkVolumeOk() (*int64, bool) {
	if o == nil || isNil(o.DownlinkVolume) {
    return nil, false
	}
	return o.DownlinkVolume, true
}

// HasDownlinkVolume returns a boolean if a field has been set.
func (o *CommunicationRecordAllOf) HasDownlinkVolume() bool {
	if o != nil && !isNil(o.DownlinkVolume) {
		return true
	}

	return false
}

// SetDownlinkVolume gets a reference to the given int64 and assigns it to the DownlinkVolume field.
func (o *CommunicationRecordAllOf) SetDownlinkVolume(v int64) {
	o.DownlinkVolume = &v
}

func (o CommunicationRecordAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["timeInterval"] = o.TimeInterval
	}
	if !isNil(o.UplinkVolume) {
		toSerialize["uplinkVolume"] = o.UplinkVolume
	}
	if !isNil(o.DownlinkVolume) {
		toSerialize["downlinkVolume"] = o.DownlinkVolume
	}
	return json.Marshal(toSerialize)
}

type NullableCommunicationRecordAllOf struct {
	value *CommunicationRecordAllOf
	isSet bool
}

func (v NullableCommunicationRecordAllOf) Get() *CommunicationRecordAllOf {
	return v.value
}

func (v *NullableCommunicationRecordAllOf) Set(val *CommunicationRecordAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCommunicationRecordAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCommunicationRecordAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommunicationRecordAllOf(val *CommunicationRecordAllOf) *NullableCommunicationRecordAllOf {
	return &NullableCommunicationRecordAllOf{value: val, isSet: true}
}

func (v NullableCommunicationRecordAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommunicationRecordAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

