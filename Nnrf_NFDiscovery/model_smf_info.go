/*
NRF NFDiscovery Service

NRF NFDiscovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnrf_NFDiscovery

import (
	"encoding/json"
)

// SmfInfo Information of an SMF NF Instance
type SmfInfo struct {
	SNssaiSmfInfoList []SnssaiSmfInfoItem `json:"sNssaiSmfInfoList"`
	TaiList []Tai `json:"taiList,omitempty"`
	TaiRangeList []TaiRange `json:"taiRangeList,omitempty"`
	// Fully Qualified Domain Name
	PgwFqdn *string `json:"pgwFqdn,omitempty"`
	PgwIpAddrList []IpAddr `json:"pgwIpAddrList,omitempty"`
	AccessType []AccessType `json:"accessType,omitempty"`
	Priority *int32 `json:"priority,omitempty"`
	VsmfSupportInd *bool `json:"vsmfSupportInd,omitempty"`
	PgwFqdnList []string `json:"pgwFqdnList,omitempty"`
	// Deprecated
	SmfOnboardingCapability *bool `json:"smfOnboardingCapability,omitempty"`
	IsmfSupportInd *bool `json:"ismfSupportInd,omitempty"`
	SmfUPRPCapability *bool `json:"smfUPRPCapability,omitempty"`
}

// NewSmfInfo instantiates a new SmfInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmfInfo(sNssaiSmfInfoList []SnssaiSmfInfoItem) *SmfInfo {
	this := SmfInfo{}
	this.SNssaiSmfInfoList = sNssaiSmfInfoList
	var smfOnboardingCapability bool = false
	this.SmfOnboardingCapability = &smfOnboardingCapability
	var smfUPRPCapability bool = false
	this.SmfUPRPCapability = &smfUPRPCapability
	return &this
}

// NewSmfInfoWithDefaults instantiates a new SmfInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmfInfoWithDefaults() *SmfInfo {
	this := SmfInfo{}
	var smfOnboardingCapability bool = false
	this.SmfOnboardingCapability = &smfOnboardingCapability
	var smfUPRPCapability bool = false
	this.SmfUPRPCapability = &smfUPRPCapability
	return &this
}

// GetSNssaiSmfInfoList returns the SNssaiSmfInfoList field value
func (o *SmfInfo) GetSNssaiSmfInfoList() []SnssaiSmfInfoItem {
	if o == nil {
		var ret []SnssaiSmfInfoItem
		return ret
	}

	return o.SNssaiSmfInfoList
}

// GetSNssaiSmfInfoListOk returns a tuple with the SNssaiSmfInfoList field value
// and a boolean to check if the value has been set.
func (o *SmfInfo) GetSNssaiSmfInfoListOk() ([]SnssaiSmfInfoItem, bool) {
	if o == nil {
    return nil, false
	}
	return o.SNssaiSmfInfoList, true
}

// SetSNssaiSmfInfoList sets field value
func (o *SmfInfo) SetSNssaiSmfInfoList(v []SnssaiSmfInfoItem) {
	o.SNssaiSmfInfoList = v
}

// GetTaiList returns the TaiList field value if set, zero value otherwise.
func (o *SmfInfo) GetTaiList() []Tai {
	if o == nil || isNil(o.TaiList) {
		var ret []Tai
		return ret
	}
	return o.TaiList
}

// GetTaiListOk returns a tuple with the TaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfInfo) GetTaiListOk() ([]Tai, bool) {
	if o == nil || isNil(o.TaiList) {
    return nil, false
	}
	return o.TaiList, true
}

// HasTaiList returns a boolean if a field has been set.
func (o *SmfInfo) HasTaiList() bool {
	if o != nil && !isNil(o.TaiList) {
		return true
	}

	return false
}

// SetTaiList gets a reference to the given []Tai and assigns it to the TaiList field.
func (o *SmfInfo) SetTaiList(v []Tai) {
	o.TaiList = v
}

// GetTaiRangeList returns the TaiRangeList field value if set, zero value otherwise.
func (o *SmfInfo) GetTaiRangeList() []TaiRange {
	if o == nil || isNil(o.TaiRangeList) {
		var ret []TaiRange
		return ret
	}
	return o.TaiRangeList
}

// GetTaiRangeListOk returns a tuple with the TaiRangeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfInfo) GetTaiRangeListOk() ([]TaiRange, bool) {
	if o == nil || isNil(o.TaiRangeList) {
    return nil, false
	}
	return o.TaiRangeList, true
}

// HasTaiRangeList returns a boolean if a field has been set.
func (o *SmfInfo) HasTaiRangeList() bool {
	if o != nil && !isNil(o.TaiRangeList) {
		return true
	}

	return false
}

// SetTaiRangeList gets a reference to the given []TaiRange and assigns it to the TaiRangeList field.
func (o *SmfInfo) SetTaiRangeList(v []TaiRange) {
	o.TaiRangeList = v
}

// GetPgwFqdn returns the PgwFqdn field value if set, zero value otherwise.
func (o *SmfInfo) GetPgwFqdn() string {
	if o == nil || isNil(o.PgwFqdn) {
		var ret string
		return ret
	}
	return *o.PgwFqdn
}

// GetPgwFqdnOk returns a tuple with the PgwFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfInfo) GetPgwFqdnOk() (*string, bool) {
	if o == nil || isNil(o.PgwFqdn) {
    return nil, false
	}
	return o.PgwFqdn, true
}

// HasPgwFqdn returns a boolean if a field has been set.
func (o *SmfInfo) HasPgwFqdn() bool {
	if o != nil && !isNil(o.PgwFqdn) {
		return true
	}

	return false
}

// SetPgwFqdn gets a reference to the given string and assigns it to the PgwFqdn field.
func (o *SmfInfo) SetPgwFqdn(v string) {
	o.PgwFqdn = &v
}

// GetPgwIpAddrList returns the PgwIpAddrList field value if set, zero value otherwise.
func (o *SmfInfo) GetPgwIpAddrList() []IpAddr {
	if o == nil || isNil(o.PgwIpAddrList) {
		var ret []IpAddr
		return ret
	}
	return o.PgwIpAddrList
}

// GetPgwIpAddrListOk returns a tuple with the PgwIpAddrList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfInfo) GetPgwIpAddrListOk() ([]IpAddr, bool) {
	if o == nil || isNil(o.PgwIpAddrList) {
    return nil, false
	}
	return o.PgwIpAddrList, true
}

// HasPgwIpAddrList returns a boolean if a field has been set.
func (o *SmfInfo) HasPgwIpAddrList() bool {
	if o != nil && !isNil(o.PgwIpAddrList) {
		return true
	}

	return false
}

// SetPgwIpAddrList gets a reference to the given []IpAddr and assigns it to the PgwIpAddrList field.
func (o *SmfInfo) SetPgwIpAddrList(v []IpAddr) {
	o.PgwIpAddrList = v
}

// GetAccessType returns the AccessType field value if set, zero value otherwise.
func (o *SmfInfo) GetAccessType() []AccessType {
	if o == nil || isNil(o.AccessType) {
		var ret []AccessType
		return ret
	}
	return o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfInfo) GetAccessTypeOk() ([]AccessType, bool) {
	if o == nil || isNil(o.AccessType) {
    return nil, false
	}
	return o.AccessType, true
}

// HasAccessType returns a boolean if a field has been set.
func (o *SmfInfo) HasAccessType() bool {
	if o != nil && !isNil(o.AccessType) {
		return true
	}

	return false
}

// SetAccessType gets a reference to the given []AccessType and assigns it to the AccessType field.
func (o *SmfInfo) SetAccessType(v []AccessType) {
	o.AccessType = v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *SmfInfo) GetPriority() int32 {
	if o == nil || isNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfInfo) GetPriorityOk() (*int32, bool) {
	if o == nil || isNil(o.Priority) {
    return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *SmfInfo) HasPriority() bool {
	if o != nil && !isNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *SmfInfo) SetPriority(v int32) {
	o.Priority = &v
}

// GetVsmfSupportInd returns the VsmfSupportInd field value if set, zero value otherwise.
func (o *SmfInfo) GetVsmfSupportInd() bool {
	if o == nil || isNil(o.VsmfSupportInd) {
		var ret bool
		return ret
	}
	return *o.VsmfSupportInd
}

// GetVsmfSupportIndOk returns a tuple with the VsmfSupportInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfInfo) GetVsmfSupportIndOk() (*bool, bool) {
	if o == nil || isNil(o.VsmfSupportInd) {
    return nil, false
	}
	return o.VsmfSupportInd, true
}

// HasVsmfSupportInd returns a boolean if a field has been set.
func (o *SmfInfo) HasVsmfSupportInd() bool {
	if o != nil && !isNil(o.VsmfSupportInd) {
		return true
	}

	return false
}

// SetVsmfSupportInd gets a reference to the given bool and assigns it to the VsmfSupportInd field.
func (o *SmfInfo) SetVsmfSupportInd(v bool) {
	o.VsmfSupportInd = &v
}

// GetPgwFqdnList returns the PgwFqdnList field value if set, zero value otherwise.
func (o *SmfInfo) GetPgwFqdnList() []string {
	if o == nil || isNil(o.PgwFqdnList) {
		var ret []string
		return ret
	}
	return o.PgwFqdnList
}

// GetPgwFqdnListOk returns a tuple with the PgwFqdnList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfInfo) GetPgwFqdnListOk() ([]string, bool) {
	if o == nil || isNil(o.PgwFqdnList) {
    return nil, false
	}
	return o.PgwFqdnList, true
}

// HasPgwFqdnList returns a boolean if a field has been set.
func (o *SmfInfo) HasPgwFqdnList() bool {
	if o != nil && !isNil(o.PgwFqdnList) {
		return true
	}

	return false
}

// SetPgwFqdnList gets a reference to the given []string and assigns it to the PgwFqdnList field.
func (o *SmfInfo) SetPgwFqdnList(v []string) {
	o.PgwFqdnList = v
}

// GetSmfOnboardingCapability returns the SmfOnboardingCapability field value if set, zero value otherwise.
// Deprecated
func (o *SmfInfo) GetSmfOnboardingCapability() bool {
	if o == nil || isNil(o.SmfOnboardingCapability) {
		var ret bool
		return ret
	}
	return *o.SmfOnboardingCapability
}

// GetSmfOnboardingCapabilityOk returns a tuple with the SmfOnboardingCapability field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *SmfInfo) GetSmfOnboardingCapabilityOk() (*bool, bool) {
	if o == nil || isNil(o.SmfOnboardingCapability) {
    return nil, false
	}
	return o.SmfOnboardingCapability, true
}

// HasSmfOnboardingCapability returns a boolean if a field has been set.
func (o *SmfInfo) HasSmfOnboardingCapability() bool {
	if o != nil && !isNil(o.SmfOnboardingCapability) {
		return true
	}

	return false
}

// SetSmfOnboardingCapability gets a reference to the given bool and assigns it to the SmfOnboardingCapability field.
// Deprecated
func (o *SmfInfo) SetSmfOnboardingCapability(v bool) {
	o.SmfOnboardingCapability = &v
}

// GetIsmfSupportInd returns the IsmfSupportInd field value if set, zero value otherwise.
func (o *SmfInfo) GetIsmfSupportInd() bool {
	if o == nil || isNil(o.IsmfSupportInd) {
		var ret bool
		return ret
	}
	return *o.IsmfSupportInd
}

// GetIsmfSupportIndOk returns a tuple with the IsmfSupportInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfInfo) GetIsmfSupportIndOk() (*bool, bool) {
	if o == nil || isNil(o.IsmfSupportInd) {
    return nil, false
	}
	return o.IsmfSupportInd, true
}

// HasIsmfSupportInd returns a boolean if a field has been set.
func (o *SmfInfo) HasIsmfSupportInd() bool {
	if o != nil && !isNil(o.IsmfSupportInd) {
		return true
	}

	return false
}

// SetIsmfSupportInd gets a reference to the given bool and assigns it to the IsmfSupportInd field.
func (o *SmfInfo) SetIsmfSupportInd(v bool) {
	o.IsmfSupportInd = &v
}

// GetSmfUPRPCapability returns the SmfUPRPCapability field value if set, zero value otherwise.
func (o *SmfInfo) GetSmfUPRPCapability() bool {
	if o == nil || isNil(o.SmfUPRPCapability) {
		var ret bool
		return ret
	}
	return *o.SmfUPRPCapability
}

// GetSmfUPRPCapabilityOk returns a tuple with the SmfUPRPCapability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfInfo) GetSmfUPRPCapabilityOk() (*bool, bool) {
	if o == nil || isNil(o.SmfUPRPCapability) {
    return nil, false
	}
	return o.SmfUPRPCapability, true
}

// HasSmfUPRPCapability returns a boolean if a field has been set.
func (o *SmfInfo) HasSmfUPRPCapability() bool {
	if o != nil && !isNil(o.SmfUPRPCapability) {
		return true
	}

	return false
}

// SetSmfUPRPCapability gets a reference to the given bool and assigns it to the SmfUPRPCapability field.
func (o *SmfInfo) SetSmfUPRPCapability(v bool) {
	o.SmfUPRPCapability = &v
}

func (o SmfInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sNssaiSmfInfoList"] = o.SNssaiSmfInfoList
	}
	if !isNil(o.TaiList) {
		toSerialize["taiList"] = o.TaiList
	}
	if !isNil(o.TaiRangeList) {
		toSerialize["taiRangeList"] = o.TaiRangeList
	}
	if !isNil(o.PgwFqdn) {
		toSerialize["pgwFqdn"] = o.PgwFqdn
	}
	if !isNil(o.PgwIpAddrList) {
		toSerialize["pgwIpAddrList"] = o.PgwIpAddrList
	}
	if !isNil(o.AccessType) {
		toSerialize["accessType"] = o.AccessType
	}
	if !isNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !isNil(o.VsmfSupportInd) {
		toSerialize["vsmfSupportInd"] = o.VsmfSupportInd
	}
	if !isNil(o.PgwFqdnList) {
		toSerialize["pgwFqdnList"] = o.PgwFqdnList
	}
	if !isNil(o.SmfOnboardingCapability) {
		toSerialize["smfOnboardingCapability"] = o.SmfOnboardingCapability
	}
	if !isNil(o.IsmfSupportInd) {
		toSerialize["ismfSupportInd"] = o.IsmfSupportInd
	}
	if !isNil(o.SmfUPRPCapability) {
		toSerialize["smfUPRPCapability"] = o.SmfUPRPCapability
	}
	return json.Marshal(toSerialize)
}

type NullableSmfInfo struct {
	value *SmfInfo
	isSet bool
}

func (v NullableSmfInfo) Get() *SmfInfo {
	return v.value
}

func (v *NullableSmfInfo) Set(val *SmfInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSmfInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSmfInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmfInfo(val *SmfInfo) *NullableSmfInfo {
	return &NullableSmfInfo{value: val, isSet: true}
}

func (v NullableSmfInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmfInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

