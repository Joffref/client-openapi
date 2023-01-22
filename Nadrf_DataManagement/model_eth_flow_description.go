/*
Nadrf_DataManagement

ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nadrf_DataManagement

import (
	"encoding/json"
)

// EthFlowDescription Identifies an Ethernet flow.
type EthFlowDescription struct {
	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042. 
	DestMacAddr *string `json:"destMacAddr,omitempty"`
	EthType string `json:"ethType"`
	// Defines a packet filter of an IP flow.
	FDesc *string `json:"fDesc,omitempty"`
	FDir *FlowDirection `json:"fDir,omitempty"`
	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042. 
	SourceMacAddr *string `json:"sourceMacAddr,omitempty"`
	VlanTags []string `json:"vlanTags,omitempty"`
	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042. 
	SrcMacAddrEnd *string `json:"srcMacAddrEnd,omitempty"`
	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042. 
	DestMacAddrEnd *string `json:"destMacAddrEnd,omitempty"`
}

// NewEthFlowDescription instantiates a new EthFlowDescription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEthFlowDescription(ethType string) *EthFlowDescription {
	this := EthFlowDescription{}
	this.EthType = ethType
	return &this
}

// NewEthFlowDescriptionWithDefaults instantiates a new EthFlowDescription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEthFlowDescriptionWithDefaults() *EthFlowDescription {
	this := EthFlowDescription{}
	return &this
}

// GetDestMacAddr returns the DestMacAddr field value if set, zero value otherwise.
func (o *EthFlowDescription) GetDestMacAddr() string {
	if o == nil || isNil(o.DestMacAddr) {
		var ret string
		return ret
	}
	return *o.DestMacAddr
}

// GetDestMacAddrOk returns a tuple with the DestMacAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthFlowDescription) GetDestMacAddrOk() (*string, bool) {
	if o == nil || isNil(o.DestMacAddr) {
    return nil, false
	}
	return o.DestMacAddr, true
}

// HasDestMacAddr returns a boolean if a field has been set.
func (o *EthFlowDescription) HasDestMacAddr() bool {
	if o != nil && !isNil(o.DestMacAddr) {
		return true
	}

	return false
}

// SetDestMacAddr gets a reference to the given string and assigns it to the DestMacAddr field.
func (o *EthFlowDescription) SetDestMacAddr(v string) {
	o.DestMacAddr = &v
}

// GetEthType returns the EthType field value
func (o *EthFlowDescription) GetEthType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EthType
}

// GetEthTypeOk returns a tuple with the EthType field value
// and a boolean to check if the value has been set.
func (o *EthFlowDescription) GetEthTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.EthType, true
}

// SetEthType sets field value
func (o *EthFlowDescription) SetEthType(v string) {
	o.EthType = v
}

// GetFDesc returns the FDesc field value if set, zero value otherwise.
func (o *EthFlowDescription) GetFDesc() string {
	if o == nil || isNil(o.FDesc) {
		var ret string
		return ret
	}
	return *o.FDesc
}

// GetFDescOk returns a tuple with the FDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthFlowDescription) GetFDescOk() (*string, bool) {
	if o == nil || isNil(o.FDesc) {
    return nil, false
	}
	return o.FDesc, true
}

// HasFDesc returns a boolean if a field has been set.
func (o *EthFlowDescription) HasFDesc() bool {
	if o != nil && !isNil(o.FDesc) {
		return true
	}

	return false
}

// SetFDesc gets a reference to the given string and assigns it to the FDesc field.
func (o *EthFlowDescription) SetFDesc(v string) {
	o.FDesc = &v
}

// GetFDir returns the FDir field value if set, zero value otherwise.
func (o *EthFlowDescription) GetFDir() FlowDirection {
	if o == nil || isNil(o.FDir) {
		var ret FlowDirection
		return ret
	}
	return *o.FDir
}

// GetFDirOk returns a tuple with the FDir field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthFlowDescription) GetFDirOk() (*FlowDirection, bool) {
	if o == nil || isNil(o.FDir) {
    return nil, false
	}
	return o.FDir, true
}

// HasFDir returns a boolean if a field has been set.
func (o *EthFlowDescription) HasFDir() bool {
	if o != nil && !isNil(o.FDir) {
		return true
	}

	return false
}

// SetFDir gets a reference to the given FlowDirection and assigns it to the FDir field.
func (o *EthFlowDescription) SetFDir(v FlowDirection) {
	o.FDir = &v
}

// GetSourceMacAddr returns the SourceMacAddr field value if set, zero value otherwise.
func (o *EthFlowDescription) GetSourceMacAddr() string {
	if o == nil || isNil(o.SourceMacAddr) {
		var ret string
		return ret
	}
	return *o.SourceMacAddr
}

// GetSourceMacAddrOk returns a tuple with the SourceMacAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthFlowDescription) GetSourceMacAddrOk() (*string, bool) {
	if o == nil || isNil(o.SourceMacAddr) {
    return nil, false
	}
	return o.SourceMacAddr, true
}

// HasSourceMacAddr returns a boolean if a field has been set.
func (o *EthFlowDescription) HasSourceMacAddr() bool {
	if o != nil && !isNil(o.SourceMacAddr) {
		return true
	}

	return false
}

// SetSourceMacAddr gets a reference to the given string and assigns it to the SourceMacAddr field.
func (o *EthFlowDescription) SetSourceMacAddr(v string) {
	o.SourceMacAddr = &v
}

// GetVlanTags returns the VlanTags field value if set, zero value otherwise.
func (o *EthFlowDescription) GetVlanTags() []string {
	if o == nil || isNil(o.VlanTags) {
		var ret []string
		return ret
	}
	return o.VlanTags
}

// GetVlanTagsOk returns a tuple with the VlanTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthFlowDescription) GetVlanTagsOk() ([]string, bool) {
	if o == nil || isNil(o.VlanTags) {
    return nil, false
	}
	return o.VlanTags, true
}

// HasVlanTags returns a boolean if a field has been set.
func (o *EthFlowDescription) HasVlanTags() bool {
	if o != nil && !isNil(o.VlanTags) {
		return true
	}

	return false
}

// SetVlanTags gets a reference to the given []string and assigns it to the VlanTags field.
func (o *EthFlowDescription) SetVlanTags(v []string) {
	o.VlanTags = v
}

// GetSrcMacAddrEnd returns the SrcMacAddrEnd field value if set, zero value otherwise.
func (o *EthFlowDescription) GetSrcMacAddrEnd() string {
	if o == nil || isNil(o.SrcMacAddrEnd) {
		var ret string
		return ret
	}
	return *o.SrcMacAddrEnd
}

// GetSrcMacAddrEndOk returns a tuple with the SrcMacAddrEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthFlowDescription) GetSrcMacAddrEndOk() (*string, bool) {
	if o == nil || isNil(o.SrcMacAddrEnd) {
    return nil, false
	}
	return o.SrcMacAddrEnd, true
}

// HasSrcMacAddrEnd returns a boolean if a field has been set.
func (o *EthFlowDescription) HasSrcMacAddrEnd() bool {
	if o != nil && !isNil(o.SrcMacAddrEnd) {
		return true
	}

	return false
}

// SetSrcMacAddrEnd gets a reference to the given string and assigns it to the SrcMacAddrEnd field.
func (o *EthFlowDescription) SetSrcMacAddrEnd(v string) {
	o.SrcMacAddrEnd = &v
}

// GetDestMacAddrEnd returns the DestMacAddrEnd field value if set, zero value otherwise.
func (o *EthFlowDescription) GetDestMacAddrEnd() string {
	if o == nil || isNil(o.DestMacAddrEnd) {
		var ret string
		return ret
	}
	return *o.DestMacAddrEnd
}

// GetDestMacAddrEndOk returns a tuple with the DestMacAddrEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthFlowDescription) GetDestMacAddrEndOk() (*string, bool) {
	if o == nil || isNil(o.DestMacAddrEnd) {
    return nil, false
	}
	return o.DestMacAddrEnd, true
}

// HasDestMacAddrEnd returns a boolean if a field has been set.
func (o *EthFlowDescription) HasDestMacAddrEnd() bool {
	if o != nil && !isNil(o.DestMacAddrEnd) {
		return true
	}

	return false
}

// SetDestMacAddrEnd gets a reference to the given string and assigns it to the DestMacAddrEnd field.
func (o *EthFlowDescription) SetDestMacAddrEnd(v string) {
	o.DestMacAddrEnd = &v
}

func (o EthFlowDescription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DestMacAddr) {
		toSerialize["destMacAddr"] = o.DestMacAddr
	}
	if true {
		toSerialize["ethType"] = o.EthType
	}
	if !isNil(o.FDesc) {
		toSerialize["fDesc"] = o.FDesc
	}
	if !isNil(o.FDir) {
		toSerialize["fDir"] = o.FDir
	}
	if !isNil(o.SourceMacAddr) {
		toSerialize["sourceMacAddr"] = o.SourceMacAddr
	}
	if !isNil(o.VlanTags) {
		toSerialize["vlanTags"] = o.VlanTags
	}
	if !isNil(o.SrcMacAddrEnd) {
		toSerialize["srcMacAddrEnd"] = o.SrcMacAddrEnd
	}
	if !isNil(o.DestMacAddrEnd) {
		toSerialize["destMacAddrEnd"] = o.DestMacAddrEnd
	}
	return json.Marshal(toSerialize)
}

type NullableEthFlowDescription struct {
	value *EthFlowDescription
	isSet bool
}

func (v NullableEthFlowDescription) Get() *EthFlowDescription {
	return v.value
}

func (v *NullableEthFlowDescription) Set(val *EthFlowDescription) {
	v.value = val
	v.isSet = true
}

func (v NullableEthFlowDescription) IsSet() bool {
	return v.isSet
}

func (v *NullableEthFlowDescription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEthFlowDescription(val *EthFlowDescription) *NullableEthFlowDescription {
	return &NullableEthFlowDescription{value: val, isSet: true}
}

func (v NullableEthFlowDescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEthFlowDescription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

