/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Subscription_Data

import (
	"encoding/json"
)

// DefaultUnrelatedClass struct for DefaultUnrelatedClass
type DefaultUnrelatedClass struct {
	AllowedGeographicArea []GeographicArea `json:"allowedGeographicArea,omitempty"`
	PrivacyCheckRelatedAction *PrivacyCheckRelatedAction `json:"privacyCheckRelatedAction,omitempty"`
	CodeWordInd *CodeWordInd `json:"codeWordInd,omitempty"`
	ValidTimePeriod *ValidTimePeriod `json:"validTimePeriod,omitempty"`
	CodeWordList []string `json:"codeWordList,omitempty"`
}

// NewDefaultUnrelatedClass instantiates a new DefaultUnrelatedClass object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefaultUnrelatedClass() *DefaultUnrelatedClass {
	this := DefaultUnrelatedClass{}
	return &this
}

// NewDefaultUnrelatedClassWithDefaults instantiates a new DefaultUnrelatedClass object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefaultUnrelatedClassWithDefaults() *DefaultUnrelatedClass {
	this := DefaultUnrelatedClass{}
	return &this
}

// GetAllowedGeographicArea returns the AllowedGeographicArea field value if set, zero value otherwise.
func (o *DefaultUnrelatedClass) GetAllowedGeographicArea() []GeographicArea {
	if o == nil || isNil(o.AllowedGeographicArea) {
		var ret []GeographicArea
		return ret
	}
	return o.AllowedGeographicArea
}

// GetAllowedGeographicAreaOk returns a tuple with the AllowedGeographicArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultUnrelatedClass) GetAllowedGeographicAreaOk() ([]GeographicArea, bool) {
	if o == nil || isNil(o.AllowedGeographicArea) {
    return nil, false
	}
	return o.AllowedGeographicArea, true
}

// HasAllowedGeographicArea returns a boolean if a field has been set.
func (o *DefaultUnrelatedClass) HasAllowedGeographicArea() bool {
	if o != nil && !isNil(o.AllowedGeographicArea) {
		return true
	}

	return false
}

// SetAllowedGeographicArea gets a reference to the given []GeographicArea and assigns it to the AllowedGeographicArea field.
func (o *DefaultUnrelatedClass) SetAllowedGeographicArea(v []GeographicArea) {
	o.AllowedGeographicArea = v
}

// GetPrivacyCheckRelatedAction returns the PrivacyCheckRelatedAction field value if set, zero value otherwise.
func (o *DefaultUnrelatedClass) GetPrivacyCheckRelatedAction() PrivacyCheckRelatedAction {
	if o == nil || isNil(o.PrivacyCheckRelatedAction) {
		var ret PrivacyCheckRelatedAction
		return ret
	}
	return *o.PrivacyCheckRelatedAction
}

// GetPrivacyCheckRelatedActionOk returns a tuple with the PrivacyCheckRelatedAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultUnrelatedClass) GetPrivacyCheckRelatedActionOk() (*PrivacyCheckRelatedAction, bool) {
	if o == nil || isNil(o.PrivacyCheckRelatedAction) {
    return nil, false
	}
	return o.PrivacyCheckRelatedAction, true
}

// HasPrivacyCheckRelatedAction returns a boolean if a field has been set.
func (o *DefaultUnrelatedClass) HasPrivacyCheckRelatedAction() bool {
	if o != nil && !isNil(o.PrivacyCheckRelatedAction) {
		return true
	}

	return false
}

// SetPrivacyCheckRelatedAction gets a reference to the given PrivacyCheckRelatedAction and assigns it to the PrivacyCheckRelatedAction field.
func (o *DefaultUnrelatedClass) SetPrivacyCheckRelatedAction(v PrivacyCheckRelatedAction) {
	o.PrivacyCheckRelatedAction = &v
}

// GetCodeWordInd returns the CodeWordInd field value if set, zero value otherwise.
func (o *DefaultUnrelatedClass) GetCodeWordInd() CodeWordInd {
	if o == nil || isNil(o.CodeWordInd) {
		var ret CodeWordInd
		return ret
	}
	return *o.CodeWordInd
}

// GetCodeWordIndOk returns a tuple with the CodeWordInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultUnrelatedClass) GetCodeWordIndOk() (*CodeWordInd, bool) {
	if o == nil || isNil(o.CodeWordInd) {
    return nil, false
	}
	return o.CodeWordInd, true
}

// HasCodeWordInd returns a boolean if a field has been set.
func (o *DefaultUnrelatedClass) HasCodeWordInd() bool {
	if o != nil && !isNil(o.CodeWordInd) {
		return true
	}

	return false
}

// SetCodeWordInd gets a reference to the given CodeWordInd and assigns it to the CodeWordInd field.
func (o *DefaultUnrelatedClass) SetCodeWordInd(v CodeWordInd) {
	o.CodeWordInd = &v
}

// GetValidTimePeriod returns the ValidTimePeriod field value if set, zero value otherwise.
func (o *DefaultUnrelatedClass) GetValidTimePeriod() ValidTimePeriod {
	if o == nil || isNil(o.ValidTimePeriod) {
		var ret ValidTimePeriod
		return ret
	}
	return *o.ValidTimePeriod
}

// GetValidTimePeriodOk returns a tuple with the ValidTimePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultUnrelatedClass) GetValidTimePeriodOk() (*ValidTimePeriod, bool) {
	if o == nil || isNil(o.ValidTimePeriod) {
    return nil, false
	}
	return o.ValidTimePeriod, true
}

// HasValidTimePeriod returns a boolean if a field has been set.
func (o *DefaultUnrelatedClass) HasValidTimePeriod() bool {
	if o != nil && !isNil(o.ValidTimePeriod) {
		return true
	}

	return false
}

// SetValidTimePeriod gets a reference to the given ValidTimePeriod and assigns it to the ValidTimePeriod field.
func (o *DefaultUnrelatedClass) SetValidTimePeriod(v ValidTimePeriod) {
	o.ValidTimePeriod = &v
}

// GetCodeWordList returns the CodeWordList field value if set, zero value otherwise.
func (o *DefaultUnrelatedClass) GetCodeWordList() []string {
	if o == nil || isNil(o.CodeWordList) {
		var ret []string
		return ret
	}
	return o.CodeWordList
}

// GetCodeWordListOk returns a tuple with the CodeWordList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultUnrelatedClass) GetCodeWordListOk() ([]string, bool) {
	if o == nil || isNil(o.CodeWordList) {
    return nil, false
	}
	return o.CodeWordList, true
}

// HasCodeWordList returns a boolean if a field has been set.
func (o *DefaultUnrelatedClass) HasCodeWordList() bool {
	if o != nil && !isNil(o.CodeWordList) {
		return true
	}

	return false
}

// SetCodeWordList gets a reference to the given []string and assigns it to the CodeWordList field.
func (o *DefaultUnrelatedClass) SetCodeWordList(v []string) {
	o.CodeWordList = v
}

func (o DefaultUnrelatedClass) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AllowedGeographicArea) {
		toSerialize["allowedGeographicArea"] = o.AllowedGeographicArea
	}
	if !isNil(o.PrivacyCheckRelatedAction) {
		toSerialize["privacyCheckRelatedAction"] = o.PrivacyCheckRelatedAction
	}
	if !isNil(o.CodeWordInd) {
		toSerialize["codeWordInd"] = o.CodeWordInd
	}
	if !isNil(o.ValidTimePeriod) {
		toSerialize["validTimePeriod"] = o.ValidTimePeriod
	}
	if !isNil(o.CodeWordList) {
		toSerialize["codeWordList"] = o.CodeWordList
	}
	return json.Marshal(toSerialize)
}

type NullableDefaultUnrelatedClass struct {
	value *DefaultUnrelatedClass
	isSet bool
}

func (v NullableDefaultUnrelatedClass) Get() *DefaultUnrelatedClass {
	return v.value
}

func (v *NullableDefaultUnrelatedClass) Set(val *DefaultUnrelatedClass) {
	v.value = val
	v.isSet = true
}

func (v NullableDefaultUnrelatedClass) IsSet() bool {
	return v.isSet
}

func (v *NullableDefaultUnrelatedClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefaultUnrelatedClass(val *DefaultUnrelatedClass) *NullableDefaultUnrelatedClass {
	return &NullableDefaultUnrelatedClass{value: val, isSet: true}
}

func (v NullableDefaultUnrelatedClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefaultUnrelatedClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

