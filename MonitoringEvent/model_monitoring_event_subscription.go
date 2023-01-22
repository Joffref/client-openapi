/*
3gpp-monitoring-event

API for Monitoring Event.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package MonitoringEvent

import (
	"encoding/json"
	"time"
	"fmt"
)

// MonitoringEventSubscription Represents a subscription to event(s) monitoring.
type MonitoringEventSubscription struct {
	interface{} *interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *MonitoringEventSubscription) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into interface{}
	err = json.Unmarshal(data, &dst.interface{});
	if err == nil {
		jsoninterface{}, _ := json.Marshal(dst.interface{})
		if string(jsoninterface{}) == "{}" { // empty struct
			dst.interface{} = nil
		} else {
			return nil // data stored in dst.interface{}, return on the first match
		}
	} else {
		dst.interface{} = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(MonitoringEventSubscription)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *MonitoringEventSubscription) MarshalJSON() ([]byte, error) {
	if src.interface{} != nil {
		return json.Marshal(&src.interface{})
	}

	return nil, nil // no data in anyOf schemas
}

type NullableMonitoringEventSubscription struct {
	value *MonitoringEventSubscription
	isSet bool
}

func (v NullableMonitoringEventSubscription) Get() *MonitoringEventSubscription {
	return v.value
}

func (v *NullableMonitoringEventSubscription) Set(val *MonitoringEventSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringEventSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringEventSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringEventSubscription(val *MonitoringEventSubscription) *NullableMonitoringEventSubscription {
	return &NullableMonitoringEventSubscription{value: val, isSet: true}
}

func (v NullableMonitoringEventSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringEventSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

