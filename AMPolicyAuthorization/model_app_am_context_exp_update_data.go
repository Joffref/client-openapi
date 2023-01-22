/*
3gpp-am-policyauthorization

API for AM policy authorization.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package AMPolicyAuthorization

import (
	"encoding/json"
)

// AppAmContextExpUpdateData Contains the modification(s) to be applied to the Individual application AM context exposure resource. 
type AppAmContextExpUpdateData struct {
	EvSubscs NullableAmEventsSubscDataRm `json:"evSubscs,omitempty"`
	HighThruInd *bool `json:"highThruInd,omitempty"`
	CovReqs []GeographicalArea `json:"covReqs,omitempty"`
	// Unsigned integer identifying a period of time in units of seconds.
	PolicyDuration *int32 `json:"policyDuration,omitempty"`
}

// NewAppAmContextExpUpdateData instantiates a new AppAmContextExpUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppAmContextExpUpdateData() *AppAmContextExpUpdateData {
	this := AppAmContextExpUpdateData{}
	return &this
}

// NewAppAmContextExpUpdateDataWithDefaults instantiates a new AppAmContextExpUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppAmContextExpUpdateDataWithDefaults() *AppAmContextExpUpdateData {
	this := AppAmContextExpUpdateData{}
	return &this
}

// GetEvSubscs returns the EvSubscs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppAmContextExpUpdateData) GetEvSubscs() AmEventsSubscDataRm {
	if o == nil || isNil(o.EvSubscs.Get()) {
		var ret AmEventsSubscDataRm
		return ret
	}
	return *o.EvSubscs.Get()
}

// GetEvSubscsOk returns a tuple with the EvSubscs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppAmContextExpUpdateData) GetEvSubscsOk() (*AmEventsSubscDataRm, bool) {
	if o == nil {
    return nil, false
	}
	return o.EvSubscs.Get(), o.EvSubscs.IsSet()
}

// HasEvSubscs returns a boolean if a field has been set.
func (o *AppAmContextExpUpdateData) HasEvSubscs() bool {
	if o != nil && o.EvSubscs.IsSet() {
		return true
	}

	return false
}

// SetEvSubscs gets a reference to the given NullableAmEventsSubscDataRm and assigns it to the EvSubscs field.
func (o *AppAmContextExpUpdateData) SetEvSubscs(v AmEventsSubscDataRm) {
	o.EvSubscs.Set(&v)
}
// SetEvSubscsNil sets the value for EvSubscs to be an explicit nil
func (o *AppAmContextExpUpdateData) SetEvSubscsNil() {
	o.EvSubscs.Set(nil)
}

// UnsetEvSubscs ensures that no value is present for EvSubscs, not even an explicit nil
func (o *AppAmContextExpUpdateData) UnsetEvSubscs() {
	o.EvSubscs.Unset()
}

// GetHighThruInd returns the HighThruInd field value if set, zero value otherwise.
func (o *AppAmContextExpUpdateData) GetHighThruInd() bool {
	if o == nil || isNil(o.HighThruInd) {
		var ret bool
		return ret
	}
	return *o.HighThruInd
}

// GetHighThruIndOk returns a tuple with the HighThruInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAmContextExpUpdateData) GetHighThruIndOk() (*bool, bool) {
	if o == nil || isNil(o.HighThruInd) {
    return nil, false
	}
	return o.HighThruInd, true
}

// HasHighThruInd returns a boolean if a field has been set.
func (o *AppAmContextExpUpdateData) HasHighThruInd() bool {
	if o != nil && !isNil(o.HighThruInd) {
		return true
	}

	return false
}

// SetHighThruInd gets a reference to the given bool and assigns it to the HighThruInd field.
func (o *AppAmContextExpUpdateData) SetHighThruInd(v bool) {
	o.HighThruInd = &v
}

// GetCovReqs returns the CovReqs field value if set, zero value otherwise.
func (o *AppAmContextExpUpdateData) GetCovReqs() []GeographicalArea {
	if o == nil || isNil(o.CovReqs) {
		var ret []GeographicalArea
		return ret
	}
	return o.CovReqs
}

// GetCovReqsOk returns a tuple with the CovReqs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAmContextExpUpdateData) GetCovReqsOk() ([]GeographicalArea, bool) {
	if o == nil || isNil(o.CovReqs) {
    return nil, false
	}
	return o.CovReqs, true
}

// HasCovReqs returns a boolean if a field has been set.
func (o *AppAmContextExpUpdateData) HasCovReqs() bool {
	if o != nil && !isNil(o.CovReqs) {
		return true
	}

	return false
}

// SetCovReqs gets a reference to the given []GeographicalArea and assigns it to the CovReqs field.
func (o *AppAmContextExpUpdateData) SetCovReqs(v []GeographicalArea) {
	o.CovReqs = v
}

// GetPolicyDuration returns the PolicyDuration field value if set, zero value otherwise.
func (o *AppAmContextExpUpdateData) GetPolicyDuration() int32 {
	if o == nil || isNil(o.PolicyDuration) {
		var ret int32
		return ret
	}
	return *o.PolicyDuration
}

// GetPolicyDurationOk returns a tuple with the PolicyDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAmContextExpUpdateData) GetPolicyDurationOk() (*int32, bool) {
	if o == nil || isNil(o.PolicyDuration) {
    return nil, false
	}
	return o.PolicyDuration, true
}

// HasPolicyDuration returns a boolean if a field has been set.
func (o *AppAmContextExpUpdateData) HasPolicyDuration() bool {
	if o != nil && !isNil(o.PolicyDuration) {
		return true
	}

	return false
}

// SetPolicyDuration gets a reference to the given int32 and assigns it to the PolicyDuration field.
func (o *AppAmContextExpUpdateData) SetPolicyDuration(v int32) {
	o.PolicyDuration = &v
}

func (o AppAmContextExpUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EvSubscs.IsSet() {
		toSerialize["evSubscs"] = o.EvSubscs.Get()
	}
	if !isNil(o.HighThruInd) {
		toSerialize["highThruInd"] = o.HighThruInd
	}
	if !isNil(o.CovReqs) {
		toSerialize["covReqs"] = o.CovReqs
	}
	if !isNil(o.PolicyDuration) {
		toSerialize["policyDuration"] = o.PolicyDuration
	}
	return json.Marshal(toSerialize)
}

type NullableAppAmContextExpUpdateData struct {
	value *AppAmContextExpUpdateData
	isSet bool
}

func (v NullableAppAmContextExpUpdateData) Get() *AppAmContextExpUpdateData {
	return v.value
}

func (v *NullableAppAmContextExpUpdateData) Set(val *AppAmContextExpUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppAmContextExpUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppAmContextExpUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppAmContextExpUpdateData(val *AppAmContextExpUpdateData) *NullableAppAmContextExpUpdateData {
	return &NullableAppAmContextExpUpdateData{value: val, isSet: true}
}

func (v NullableAppAmContextExpUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppAmContextExpUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

