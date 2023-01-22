/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudr_DR

import (
	"encoding/json"
)

// PolicyDataForIndividualUe Contains policy data for a given subscriber.
type PolicyDataForIndividualUe struct {
	UePolicyDataSet *UePolicySet `json:"uePolicyDataSet,omitempty"`
	SmPolicyDataSet *SmPolicyData `json:"smPolicyDataSet,omitempty"`
	AmPolicyDataSet *AmPolicyData `json:"amPolicyDataSet,omitempty"`
	// Contains UM policies. The value of the limit identifier is used as the key of the map. 
	UmData *map[string]UsageMonData `json:"umData,omitempty"`
	// Contains Operator Specific Data resource data. The key of the map is operator specific data element name and the value is the operator specific data of the UE. 
	OperatorSpecificDataSet *map[string]OperatorSpecificDataContainer `json:"operatorSpecificDataSet,omitempty"`
}

// NewPolicyDataForIndividualUe instantiates a new PolicyDataForIndividualUe object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyDataForIndividualUe() *PolicyDataForIndividualUe {
	this := PolicyDataForIndividualUe{}
	return &this
}

// NewPolicyDataForIndividualUeWithDefaults instantiates a new PolicyDataForIndividualUe object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyDataForIndividualUeWithDefaults() *PolicyDataForIndividualUe {
	this := PolicyDataForIndividualUe{}
	return &this
}

// GetUePolicyDataSet returns the UePolicyDataSet field value if set, zero value otherwise.
func (o *PolicyDataForIndividualUe) GetUePolicyDataSet() UePolicySet {
	if o == nil || isNil(o.UePolicyDataSet) {
		var ret UePolicySet
		return ret
	}
	return *o.UePolicyDataSet
}

// GetUePolicyDataSetOk returns a tuple with the UePolicyDataSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDataForIndividualUe) GetUePolicyDataSetOk() (*UePolicySet, bool) {
	if o == nil || isNil(o.UePolicyDataSet) {
    return nil, false
	}
	return o.UePolicyDataSet, true
}

// HasUePolicyDataSet returns a boolean if a field has been set.
func (o *PolicyDataForIndividualUe) HasUePolicyDataSet() bool {
	if o != nil && !isNil(o.UePolicyDataSet) {
		return true
	}

	return false
}

// SetUePolicyDataSet gets a reference to the given UePolicySet and assigns it to the UePolicyDataSet field.
func (o *PolicyDataForIndividualUe) SetUePolicyDataSet(v UePolicySet) {
	o.UePolicyDataSet = &v
}

// GetSmPolicyDataSet returns the SmPolicyDataSet field value if set, zero value otherwise.
func (o *PolicyDataForIndividualUe) GetSmPolicyDataSet() SmPolicyData {
	if o == nil || isNil(o.SmPolicyDataSet) {
		var ret SmPolicyData
		return ret
	}
	return *o.SmPolicyDataSet
}

// GetSmPolicyDataSetOk returns a tuple with the SmPolicyDataSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDataForIndividualUe) GetSmPolicyDataSetOk() (*SmPolicyData, bool) {
	if o == nil || isNil(o.SmPolicyDataSet) {
    return nil, false
	}
	return o.SmPolicyDataSet, true
}

// HasSmPolicyDataSet returns a boolean if a field has been set.
func (o *PolicyDataForIndividualUe) HasSmPolicyDataSet() bool {
	if o != nil && !isNil(o.SmPolicyDataSet) {
		return true
	}

	return false
}

// SetSmPolicyDataSet gets a reference to the given SmPolicyData and assigns it to the SmPolicyDataSet field.
func (o *PolicyDataForIndividualUe) SetSmPolicyDataSet(v SmPolicyData) {
	o.SmPolicyDataSet = &v
}

// GetAmPolicyDataSet returns the AmPolicyDataSet field value if set, zero value otherwise.
func (o *PolicyDataForIndividualUe) GetAmPolicyDataSet() AmPolicyData {
	if o == nil || isNil(o.AmPolicyDataSet) {
		var ret AmPolicyData
		return ret
	}
	return *o.AmPolicyDataSet
}

// GetAmPolicyDataSetOk returns a tuple with the AmPolicyDataSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDataForIndividualUe) GetAmPolicyDataSetOk() (*AmPolicyData, bool) {
	if o == nil || isNil(o.AmPolicyDataSet) {
    return nil, false
	}
	return o.AmPolicyDataSet, true
}

// HasAmPolicyDataSet returns a boolean if a field has been set.
func (o *PolicyDataForIndividualUe) HasAmPolicyDataSet() bool {
	if o != nil && !isNil(o.AmPolicyDataSet) {
		return true
	}

	return false
}

// SetAmPolicyDataSet gets a reference to the given AmPolicyData and assigns it to the AmPolicyDataSet field.
func (o *PolicyDataForIndividualUe) SetAmPolicyDataSet(v AmPolicyData) {
	o.AmPolicyDataSet = &v
}

// GetUmData returns the UmData field value if set, zero value otherwise.
func (o *PolicyDataForIndividualUe) GetUmData() map[string]UsageMonData {
	if o == nil || isNil(o.UmData) {
		var ret map[string]UsageMonData
		return ret
	}
	return *o.UmData
}

// GetUmDataOk returns a tuple with the UmData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDataForIndividualUe) GetUmDataOk() (*map[string]UsageMonData, bool) {
	if o == nil || isNil(o.UmData) {
    return nil, false
	}
	return o.UmData, true
}

// HasUmData returns a boolean if a field has been set.
func (o *PolicyDataForIndividualUe) HasUmData() bool {
	if o != nil && !isNil(o.UmData) {
		return true
	}

	return false
}

// SetUmData gets a reference to the given map[string]UsageMonData and assigns it to the UmData field.
func (o *PolicyDataForIndividualUe) SetUmData(v map[string]UsageMonData) {
	o.UmData = &v
}

// GetOperatorSpecificDataSet returns the OperatorSpecificDataSet field value if set, zero value otherwise.
func (o *PolicyDataForIndividualUe) GetOperatorSpecificDataSet() map[string]OperatorSpecificDataContainer {
	if o == nil || isNil(o.OperatorSpecificDataSet) {
		var ret map[string]OperatorSpecificDataContainer
		return ret
	}
	return *o.OperatorSpecificDataSet
}

// GetOperatorSpecificDataSetOk returns a tuple with the OperatorSpecificDataSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyDataForIndividualUe) GetOperatorSpecificDataSetOk() (*map[string]OperatorSpecificDataContainer, bool) {
	if o == nil || isNil(o.OperatorSpecificDataSet) {
    return nil, false
	}
	return o.OperatorSpecificDataSet, true
}

// HasOperatorSpecificDataSet returns a boolean if a field has been set.
func (o *PolicyDataForIndividualUe) HasOperatorSpecificDataSet() bool {
	if o != nil && !isNil(o.OperatorSpecificDataSet) {
		return true
	}

	return false
}

// SetOperatorSpecificDataSet gets a reference to the given map[string]OperatorSpecificDataContainer and assigns it to the OperatorSpecificDataSet field.
func (o *PolicyDataForIndividualUe) SetOperatorSpecificDataSet(v map[string]OperatorSpecificDataContainer) {
	o.OperatorSpecificDataSet = &v
}

func (o PolicyDataForIndividualUe) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.UePolicyDataSet) {
		toSerialize["uePolicyDataSet"] = o.UePolicyDataSet
	}
	if !isNil(o.SmPolicyDataSet) {
		toSerialize["smPolicyDataSet"] = o.SmPolicyDataSet
	}
	if !isNil(o.AmPolicyDataSet) {
		toSerialize["amPolicyDataSet"] = o.AmPolicyDataSet
	}
	if !isNil(o.UmData) {
		toSerialize["umData"] = o.UmData
	}
	if !isNil(o.OperatorSpecificDataSet) {
		toSerialize["operatorSpecificDataSet"] = o.OperatorSpecificDataSet
	}
	return json.Marshal(toSerialize)
}

type NullablePolicyDataForIndividualUe struct {
	value *PolicyDataForIndividualUe
	isSet bool
}

func (v NullablePolicyDataForIndividualUe) Get() *PolicyDataForIndividualUe {
	return v.value
}

func (v *NullablePolicyDataForIndividualUe) Set(val *PolicyDataForIndividualUe) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyDataForIndividualUe) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyDataForIndividualUe) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyDataForIndividualUe(val *PolicyDataForIndividualUe) *NullablePolicyDataForIndividualUe {
	return &NullablePolicyDataForIndividualUe{value: val, isSet: true}
}

func (v NullablePolicyDataForIndividualUe) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyDataForIndividualUe) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

