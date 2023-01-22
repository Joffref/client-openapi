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

// AlarmsGet200ResponseValueAllOf struct for AlarmsGet200ResponseValueAllOf
type AlarmsGet200ResponseValueAllOf struct {
	LastNotificationHeader *NotificationHeader `json:"lastNotificationHeader,omitempty"`
}

// NewAlarmsGet200ResponseValueAllOf instantiates a new AlarmsGet200ResponseValueAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlarmsGet200ResponseValueAllOf() *AlarmsGet200ResponseValueAllOf {
	this := AlarmsGet200ResponseValueAllOf{}
	return &this
}

// NewAlarmsGet200ResponseValueAllOfWithDefaults instantiates a new AlarmsGet200ResponseValueAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlarmsGet200ResponseValueAllOfWithDefaults() *AlarmsGet200ResponseValueAllOf {
	this := AlarmsGet200ResponseValueAllOf{}
	return &this
}

// GetLastNotificationHeader returns the LastNotificationHeader field value if set, zero value otherwise.
func (o *AlarmsGet200ResponseValueAllOf) GetLastNotificationHeader() NotificationHeader {
	if o == nil || isNil(o.LastNotificationHeader) {
		var ret NotificationHeader
		return ret
	}
	return *o.LastNotificationHeader
}

// GetLastNotificationHeaderOk returns a tuple with the LastNotificationHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmsGet200ResponseValueAllOf) GetLastNotificationHeaderOk() (*NotificationHeader, bool) {
	if o == nil || isNil(o.LastNotificationHeader) {
    return nil, false
	}
	return o.LastNotificationHeader, true
}

// HasLastNotificationHeader returns a boolean if a field has been set.
func (o *AlarmsGet200ResponseValueAllOf) HasLastNotificationHeader() bool {
	if o != nil && !isNil(o.LastNotificationHeader) {
		return true
	}

	return false
}

// SetLastNotificationHeader gets a reference to the given NotificationHeader and assigns it to the LastNotificationHeader field.
func (o *AlarmsGet200ResponseValueAllOf) SetLastNotificationHeader(v NotificationHeader) {
	o.LastNotificationHeader = &v
}

func (o AlarmsGet200ResponseValueAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.LastNotificationHeader) {
		toSerialize["lastNotificationHeader"] = o.LastNotificationHeader
	}
	return json.Marshal(toSerialize)
}

type NullableAlarmsGet200ResponseValueAllOf struct {
	value *AlarmsGet200ResponseValueAllOf
	isSet bool
}

func (v NullableAlarmsGet200ResponseValueAllOf) Get() *AlarmsGet200ResponseValueAllOf {
	return v.value
}

func (v *NullableAlarmsGet200ResponseValueAllOf) Set(val *AlarmsGet200ResponseValueAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAlarmsGet200ResponseValueAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAlarmsGet200ResponseValueAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlarmsGet200ResponseValueAllOf(val *AlarmsGet200ResponseValueAllOf) *NullableAlarmsGet200ResponseValueAllOf {
	return &NullableAlarmsGet200ResponseValueAllOf{value: val, isSet: true}
}

func (v NullableAlarmsGet200ResponseValueAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlarmsGet200ResponseValueAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

