/*
EES EEL Managed ACR Service

EES EEL Managed ACR Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Eees_EELManagedACR

import (
	"encoding/json"
)

// EELACRReq Represents the parameters to request the EES (e.g. S-EES) to handle all the operations of an ACR. 
type EELACRReq struct {
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.  
	UeId string `json:"ueId"`
	EasCharacs []EasCharacteristics `json:"easCharacs"`
	// string providing an URI formatted according to IETF RFC 3986.
	AppCtxtStoreAddr *string `json:"appCtxtStoreAddr,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SuppFeat *string `json:"suppFeat,omitempty"`
}

// NewEELACRReq instantiates a new EELACRReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEELACRReq(ueId string, easCharacs []EasCharacteristics) *EELACRReq {
	this := EELACRReq{}
	this.UeId = ueId
	this.EasCharacs = easCharacs
	return &this
}

// NewEELACRReqWithDefaults instantiates a new EELACRReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEELACRReqWithDefaults() *EELACRReq {
	this := EELACRReq{}
	return &this
}

// GetUeId returns the UeId field value
func (o *EELACRReq) GetUeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UeId
}

// GetUeIdOk returns a tuple with the UeId field value
// and a boolean to check if the value has been set.
func (o *EELACRReq) GetUeIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UeId, true
}

// SetUeId sets field value
func (o *EELACRReq) SetUeId(v string) {
	o.UeId = v
}

// GetEasCharacs returns the EasCharacs field value
func (o *EELACRReq) GetEasCharacs() []EasCharacteristics {
	if o == nil {
		var ret []EasCharacteristics
		return ret
	}

	return o.EasCharacs
}

// GetEasCharacsOk returns a tuple with the EasCharacs field value
// and a boolean to check if the value has been set.
func (o *EELACRReq) GetEasCharacsOk() ([]EasCharacteristics, bool) {
	if o == nil {
    return nil, false
	}
	return o.EasCharacs, true
}

// SetEasCharacs sets field value
func (o *EELACRReq) SetEasCharacs(v []EasCharacteristics) {
	o.EasCharacs = v
}

// GetAppCtxtStoreAddr returns the AppCtxtStoreAddr field value if set, zero value otherwise.
func (o *EELACRReq) GetAppCtxtStoreAddr() string {
	if o == nil || isNil(o.AppCtxtStoreAddr) {
		var ret string
		return ret
	}
	return *o.AppCtxtStoreAddr
}

// GetAppCtxtStoreAddrOk returns a tuple with the AppCtxtStoreAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EELACRReq) GetAppCtxtStoreAddrOk() (*string, bool) {
	if o == nil || isNil(o.AppCtxtStoreAddr) {
    return nil, false
	}
	return o.AppCtxtStoreAddr, true
}

// HasAppCtxtStoreAddr returns a boolean if a field has been set.
func (o *EELACRReq) HasAppCtxtStoreAddr() bool {
	if o != nil && !isNil(o.AppCtxtStoreAddr) {
		return true
	}

	return false
}

// SetAppCtxtStoreAddr gets a reference to the given string and assigns it to the AppCtxtStoreAddr field.
func (o *EELACRReq) SetAppCtxtStoreAddr(v string) {
	o.AppCtxtStoreAddr = &v
}

// GetSuppFeat returns the SuppFeat field value if set, zero value otherwise.
func (o *EELACRReq) GetSuppFeat() string {
	if o == nil || isNil(o.SuppFeat) {
		var ret string
		return ret
	}
	return *o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EELACRReq) GetSuppFeatOk() (*string, bool) {
	if o == nil || isNil(o.SuppFeat) {
    return nil, false
	}
	return o.SuppFeat, true
}

// HasSuppFeat returns a boolean if a field has been set.
func (o *EELACRReq) HasSuppFeat() bool {
	if o != nil && !isNil(o.SuppFeat) {
		return true
	}

	return false
}

// SetSuppFeat gets a reference to the given string and assigns it to the SuppFeat field.
func (o *EELACRReq) SetSuppFeat(v string) {
	o.SuppFeat = &v
}

func (o EELACRReq) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ueId"] = o.UeId
	}
	if true {
		toSerialize["easCharacs"] = o.EasCharacs
	}
	if !isNil(o.AppCtxtStoreAddr) {
		toSerialize["appCtxtStoreAddr"] = o.AppCtxtStoreAddr
	}
	if !isNil(o.SuppFeat) {
		toSerialize["suppFeat"] = o.SuppFeat
	}
	return json.Marshal(toSerialize)
}

type NullableEELACRReq struct {
	value *EELACRReq
	isSet bool
}

func (v NullableEELACRReq) Get() *EELACRReq {
	return v.value
}

func (v *NullableEELACRReq) Set(val *EELACRReq) {
	v.value = val
	v.isSet = true
}

func (v NullableEELACRReq) IsSet() bool {
	return v.isSet
}

func (v *NullableEELACRReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEELACRReq(val *EELACRReq) *NullableEELACRReq {
	return &NullableEELACRReq{value: val, isSet: true}
}

func (v NullableEELACRReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEELACRReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

