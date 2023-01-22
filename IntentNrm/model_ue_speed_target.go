/*
Intent NRM

OAS 3.0.1 definition of the Intent NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package IntentNrm

import (
	"encoding/json"
)

// UESpeedTarget This data type is the \"ExpectationTarget\" data type with specialisations for UESpeedTarget
type UESpeedTarget struct {
	TargetAttribute *string `json:"targetAttribute,omitempty"`
	TargetCondition *string `json:"targetCondition,omitempty"`
	TargetValueRange *int32 `json:"targetValueRange,omitempty"`
}

// NewUESpeedTarget instantiates a new UESpeedTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUESpeedTarget() *UESpeedTarget {
	this := UESpeedTarget{}
	return &this
}

// NewUESpeedTargetWithDefaults instantiates a new UESpeedTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUESpeedTargetWithDefaults() *UESpeedTarget {
	this := UESpeedTarget{}
	return &this
}

// GetTargetAttribute returns the TargetAttribute field value if set, zero value otherwise.
func (o *UESpeedTarget) GetTargetAttribute() string {
	if o == nil || isNil(o.TargetAttribute) {
		var ret string
		return ret
	}
	return *o.TargetAttribute
}

// GetTargetAttributeOk returns a tuple with the TargetAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UESpeedTarget) GetTargetAttributeOk() (*string, bool) {
	if o == nil || isNil(o.TargetAttribute) {
    return nil, false
	}
	return o.TargetAttribute, true
}

// HasTargetAttribute returns a boolean if a field has been set.
func (o *UESpeedTarget) HasTargetAttribute() bool {
	if o != nil && !isNil(o.TargetAttribute) {
		return true
	}

	return false
}

// SetTargetAttribute gets a reference to the given string and assigns it to the TargetAttribute field.
func (o *UESpeedTarget) SetTargetAttribute(v string) {
	o.TargetAttribute = &v
}

// GetTargetCondition returns the TargetCondition field value if set, zero value otherwise.
func (o *UESpeedTarget) GetTargetCondition() string {
	if o == nil || isNil(o.TargetCondition) {
		var ret string
		return ret
	}
	return *o.TargetCondition
}

// GetTargetConditionOk returns a tuple with the TargetCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UESpeedTarget) GetTargetConditionOk() (*string, bool) {
	if o == nil || isNil(o.TargetCondition) {
    return nil, false
	}
	return o.TargetCondition, true
}

// HasTargetCondition returns a boolean if a field has been set.
func (o *UESpeedTarget) HasTargetCondition() bool {
	if o != nil && !isNil(o.TargetCondition) {
		return true
	}

	return false
}

// SetTargetCondition gets a reference to the given string and assigns it to the TargetCondition field.
func (o *UESpeedTarget) SetTargetCondition(v string) {
	o.TargetCondition = &v
}

// GetTargetValueRange returns the TargetValueRange field value if set, zero value otherwise.
func (o *UESpeedTarget) GetTargetValueRange() int32 {
	if o == nil || isNil(o.TargetValueRange) {
		var ret int32
		return ret
	}
	return *o.TargetValueRange
}

// GetTargetValueRangeOk returns a tuple with the TargetValueRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UESpeedTarget) GetTargetValueRangeOk() (*int32, bool) {
	if o == nil || isNil(o.TargetValueRange) {
    return nil, false
	}
	return o.TargetValueRange, true
}

// HasTargetValueRange returns a boolean if a field has been set.
func (o *UESpeedTarget) HasTargetValueRange() bool {
	if o != nil && !isNil(o.TargetValueRange) {
		return true
	}

	return false
}

// SetTargetValueRange gets a reference to the given int32 and assigns it to the TargetValueRange field.
func (o *UESpeedTarget) SetTargetValueRange(v int32) {
	o.TargetValueRange = &v
}

func (o UESpeedTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TargetAttribute) {
		toSerialize["targetAttribute"] = o.TargetAttribute
	}
	if !isNil(o.TargetCondition) {
		toSerialize["targetCondition"] = o.TargetCondition
	}
	if !isNil(o.TargetValueRange) {
		toSerialize["targetValueRange"] = o.TargetValueRange
	}
	return json.Marshal(toSerialize)
}

type NullableUESpeedTarget struct {
	value *UESpeedTarget
	isSet bool
}

func (v NullableUESpeedTarget) Get() *UESpeedTarget {
	return v.value
}

func (v *NullableUESpeedTarget) Set(val *UESpeedTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableUESpeedTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableUESpeedTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUESpeedTarget(val *UESpeedTarget) *NullableUESpeedTarget {
	return &NullableUESpeedTarget{value: val, isSet: true}
}

func (v NullableUESpeedTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUESpeedTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

