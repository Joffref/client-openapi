/*
Nnef_SMContext

Nnef SMContext Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnef_SMContext

import (
	"encoding/json"
)

// SmContextConfiguration NIDD Configuration for the SM context.
type SmContextConfiguration struct {
	SmalDataRateControl *SmallDataRateControl `json:"smalDataRateControl,omitempty"`
	SmallDataRateStatus *SmallDataRateStatus `json:"smallDataRateStatus,omitempty"`
	// When present, this IE shall contain the maximum allowed number of Downlink NAS Data PDUs per deci hour of the serving PLMN, as specified in subclause 5.31.14.2 of 3GPP TS 23.501 [2].   Minimum  10 
	ServPlmnDataRateCtl NullableInt32 `json:"servPlmnDataRateCtl,omitempty"`
}

// NewSmContextConfiguration instantiates a new SmContextConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmContextConfiguration() *SmContextConfiguration {
	this := SmContextConfiguration{}
	return &this
}

// NewSmContextConfigurationWithDefaults instantiates a new SmContextConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmContextConfigurationWithDefaults() *SmContextConfiguration {
	this := SmContextConfiguration{}
	return &this
}

// GetSmalDataRateControl returns the SmalDataRateControl field value if set, zero value otherwise.
func (o *SmContextConfiguration) GetSmalDataRateControl() SmallDataRateControl {
	if o == nil || isNil(o.SmalDataRateControl) {
		var ret SmallDataRateControl
		return ret
	}
	return *o.SmalDataRateControl
}

// GetSmalDataRateControlOk returns a tuple with the SmalDataRateControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmContextConfiguration) GetSmalDataRateControlOk() (*SmallDataRateControl, bool) {
	if o == nil || isNil(o.SmalDataRateControl) {
    return nil, false
	}
	return o.SmalDataRateControl, true
}

// HasSmalDataRateControl returns a boolean if a field has been set.
func (o *SmContextConfiguration) HasSmalDataRateControl() bool {
	if o != nil && !isNil(o.SmalDataRateControl) {
		return true
	}

	return false
}

// SetSmalDataRateControl gets a reference to the given SmallDataRateControl and assigns it to the SmalDataRateControl field.
func (o *SmContextConfiguration) SetSmalDataRateControl(v SmallDataRateControl) {
	o.SmalDataRateControl = &v
}

// GetSmallDataRateStatus returns the SmallDataRateStatus field value if set, zero value otherwise.
func (o *SmContextConfiguration) GetSmallDataRateStatus() SmallDataRateStatus {
	if o == nil || isNil(o.SmallDataRateStatus) {
		var ret SmallDataRateStatus
		return ret
	}
	return *o.SmallDataRateStatus
}

// GetSmallDataRateStatusOk returns a tuple with the SmallDataRateStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmContextConfiguration) GetSmallDataRateStatusOk() (*SmallDataRateStatus, bool) {
	if o == nil || isNil(o.SmallDataRateStatus) {
    return nil, false
	}
	return o.SmallDataRateStatus, true
}

// HasSmallDataRateStatus returns a boolean if a field has been set.
func (o *SmContextConfiguration) HasSmallDataRateStatus() bool {
	if o != nil && !isNil(o.SmallDataRateStatus) {
		return true
	}

	return false
}

// SetSmallDataRateStatus gets a reference to the given SmallDataRateStatus and assigns it to the SmallDataRateStatus field.
func (o *SmContextConfiguration) SetSmallDataRateStatus(v SmallDataRateStatus) {
	o.SmallDataRateStatus = &v
}

// GetServPlmnDataRateCtl returns the ServPlmnDataRateCtl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmContextConfiguration) GetServPlmnDataRateCtl() int32 {
	if o == nil || isNil(o.ServPlmnDataRateCtl.Get()) {
		var ret int32
		return ret
	}
	return *o.ServPlmnDataRateCtl.Get()
}

// GetServPlmnDataRateCtlOk returns a tuple with the ServPlmnDataRateCtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmContextConfiguration) GetServPlmnDataRateCtlOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.ServPlmnDataRateCtl.Get(), o.ServPlmnDataRateCtl.IsSet()
}

// HasServPlmnDataRateCtl returns a boolean if a field has been set.
func (o *SmContextConfiguration) HasServPlmnDataRateCtl() bool {
	if o != nil && o.ServPlmnDataRateCtl.IsSet() {
		return true
	}

	return false
}

// SetServPlmnDataRateCtl gets a reference to the given NullableInt32 and assigns it to the ServPlmnDataRateCtl field.
func (o *SmContextConfiguration) SetServPlmnDataRateCtl(v int32) {
	o.ServPlmnDataRateCtl.Set(&v)
}
// SetServPlmnDataRateCtlNil sets the value for ServPlmnDataRateCtl to be an explicit nil
func (o *SmContextConfiguration) SetServPlmnDataRateCtlNil() {
	o.ServPlmnDataRateCtl.Set(nil)
}

// UnsetServPlmnDataRateCtl ensures that no value is present for ServPlmnDataRateCtl, not even an explicit nil
func (o *SmContextConfiguration) UnsetServPlmnDataRateCtl() {
	o.ServPlmnDataRateCtl.Unset()
}

func (o SmContextConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SmalDataRateControl) {
		toSerialize["smalDataRateControl"] = o.SmalDataRateControl
	}
	if !isNil(o.SmallDataRateStatus) {
		toSerialize["smallDataRateStatus"] = o.SmallDataRateStatus
	}
	if o.ServPlmnDataRateCtl.IsSet() {
		toSerialize["servPlmnDataRateCtl"] = o.ServPlmnDataRateCtl.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSmContextConfiguration struct {
	value *SmContextConfiguration
	isSet bool
}

func (v NullableSmContextConfiguration) Get() *SmContextConfiguration {
	return v.value
}

func (v *NullableSmContextConfiguration) Set(val *SmContextConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableSmContextConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableSmContextConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmContextConfiguration(val *SmContextConfiguration) *NullableSmContextConfiguration {
	return &NullableSmContextConfiguration{value: val, isSet: true}
}

func (v NullableSmContextConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmContextConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

