/*
Nsmf_PDUSession

SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nsmf_PDUSession

import (
	"encoding/json"
	"os"
)

// UpdatePduSessionRequest struct for UpdatePduSessionRequest
type UpdatePduSessionRequest struct {
	JsonData *HsmfUpdateData `json:"jsonData,omitempty"`
	BinaryDataN1SmInfoFromUe **os.File `json:"binaryDataN1SmInfoFromUe,omitempty"`
	BinaryDataUnknownN1SmInfo **os.File `json:"binaryDataUnknownN1SmInfo,omitempty"`
	BinaryDataN4Information **os.File `json:"binaryDataN4Information,omitempty"`
	BinaryDataN4InformationExt1 **os.File `json:"binaryDataN4InformationExt1,omitempty"`
	BinaryDataN4InformationExt2 **os.File `json:"binaryDataN4InformationExt2,omitempty"`
}

// NewUpdatePduSessionRequest instantiates a new UpdatePduSessionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePduSessionRequest() *UpdatePduSessionRequest {
	this := UpdatePduSessionRequest{}
	return &this
}

// NewUpdatePduSessionRequestWithDefaults instantiates a new UpdatePduSessionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePduSessionRequestWithDefaults() *UpdatePduSessionRequest {
	this := UpdatePduSessionRequest{}
	return &this
}

// GetJsonData returns the JsonData field value if set, zero value otherwise.
func (o *UpdatePduSessionRequest) GetJsonData() HsmfUpdateData {
	if o == nil || isNil(o.JsonData) {
		var ret HsmfUpdateData
		return ret
	}
	return *o.JsonData
}

// GetJsonDataOk returns a tuple with the JsonData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePduSessionRequest) GetJsonDataOk() (*HsmfUpdateData, bool) {
	if o == nil || isNil(o.JsonData) {
    return nil, false
	}
	return o.JsonData, true
}

// HasJsonData returns a boolean if a field has been set.
func (o *UpdatePduSessionRequest) HasJsonData() bool {
	if o != nil && !isNil(o.JsonData) {
		return true
	}

	return false
}

// SetJsonData gets a reference to the given HsmfUpdateData and assigns it to the JsonData field.
func (o *UpdatePduSessionRequest) SetJsonData(v HsmfUpdateData) {
	o.JsonData = &v
}

// GetBinaryDataN1SmInfoFromUe returns the BinaryDataN1SmInfoFromUe field value if set, zero value otherwise.
func (o *UpdatePduSessionRequest) GetBinaryDataN1SmInfoFromUe() *os.File {
	if o == nil || isNil(o.BinaryDataN1SmInfoFromUe) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN1SmInfoFromUe
}

// GetBinaryDataN1SmInfoFromUeOk returns a tuple with the BinaryDataN1SmInfoFromUe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePduSessionRequest) GetBinaryDataN1SmInfoFromUeOk() (**os.File, bool) {
	if o == nil || isNil(o.BinaryDataN1SmInfoFromUe) {
    return nil, false
	}
	return o.BinaryDataN1SmInfoFromUe, true
}

// HasBinaryDataN1SmInfoFromUe returns a boolean if a field has been set.
func (o *UpdatePduSessionRequest) HasBinaryDataN1SmInfoFromUe() bool {
	if o != nil && !isNil(o.BinaryDataN1SmInfoFromUe) {
		return true
	}

	return false
}

// SetBinaryDataN1SmInfoFromUe gets a reference to the given *os.File and assigns it to the BinaryDataN1SmInfoFromUe field.
func (o *UpdatePduSessionRequest) SetBinaryDataN1SmInfoFromUe(v *os.File) {
	o.BinaryDataN1SmInfoFromUe = &v
}

// GetBinaryDataUnknownN1SmInfo returns the BinaryDataUnknownN1SmInfo field value if set, zero value otherwise.
func (o *UpdatePduSessionRequest) GetBinaryDataUnknownN1SmInfo() *os.File {
	if o == nil || isNil(o.BinaryDataUnknownN1SmInfo) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataUnknownN1SmInfo
}

// GetBinaryDataUnknownN1SmInfoOk returns a tuple with the BinaryDataUnknownN1SmInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePduSessionRequest) GetBinaryDataUnknownN1SmInfoOk() (**os.File, bool) {
	if o == nil || isNil(o.BinaryDataUnknownN1SmInfo) {
    return nil, false
	}
	return o.BinaryDataUnknownN1SmInfo, true
}

// HasBinaryDataUnknownN1SmInfo returns a boolean if a field has been set.
func (o *UpdatePduSessionRequest) HasBinaryDataUnknownN1SmInfo() bool {
	if o != nil && !isNil(o.BinaryDataUnknownN1SmInfo) {
		return true
	}

	return false
}

// SetBinaryDataUnknownN1SmInfo gets a reference to the given *os.File and assigns it to the BinaryDataUnknownN1SmInfo field.
func (o *UpdatePduSessionRequest) SetBinaryDataUnknownN1SmInfo(v *os.File) {
	o.BinaryDataUnknownN1SmInfo = &v
}

// GetBinaryDataN4Information returns the BinaryDataN4Information field value if set, zero value otherwise.
func (o *UpdatePduSessionRequest) GetBinaryDataN4Information() *os.File {
	if o == nil || isNil(o.BinaryDataN4Information) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN4Information
}

// GetBinaryDataN4InformationOk returns a tuple with the BinaryDataN4Information field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePduSessionRequest) GetBinaryDataN4InformationOk() (**os.File, bool) {
	if o == nil || isNil(o.BinaryDataN4Information) {
    return nil, false
	}
	return o.BinaryDataN4Information, true
}

// HasBinaryDataN4Information returns a boolean if a field has been set.
func (o *UpdatePduSessionRequest) HasBinaryDataN4Information() bool {
	if o != nil && !isNil(o.BinaryDataN4Information) {
		return true
	}

	return false
}

// SetBinaryDataN4Information gets a reference to the given *os.File and assigns it to the BinaryDataN4Information field.
func (o *UpdatePduSessionRequest) SetBinaryDataN4Information(v *os.File) {
	o.BinaryDataN4Information = &v
}

// GetBinaryDataN4InformationExt1 returns the BinaryDataN4InformationExt1 field value if set, zero value otherwise.
func (o *UpdatePduSessionRequest) GetBinaryDataN4InformationExt1() *os.File {
	if o == nil || isNil(o.BinaryDataN4InformationExt1) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN4InformationExt1
}

// GetBinaryDataN4InformationExt1Ok returns a tuple with the BinaryDataN4InformationExt1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePduSessionRequest) GetBinaryDataN4InformationExt1Ok() (**os.File, bool) {
	if o == nil || isNil(o.BinaryDataN4InformationExt1) {
    return nil, false
	}
	return o.BinaryDataN4InformationExt1, true
}

// HasBinaryDataN4InformationExt1 returns a boolean if a field has been set.
func (o *UpdatePduSessionRequest) HasBinaryDataN4InformationExt1() bool {
	if o != nil && !isNil(o.BinaryDataN4InformationExt1) {
		return true
	}

	return false
}

// SetBinaryDataN4InformationExt1 gets a reference to the given *os.File and assigns it to the BinaryDataN4InformationExt1 field.
func (o *UpdatePduSessionRequest) SetBinaryDataN4InformationExt1(v *os.File) {
	o.BinaryDataN4InformationExt1 = &v
}

// GetBinaryDataN4InformationExt2 returns the BinaryDataN4InformationExt2 field value if set, zero value otherwise.
func (o *UpdatePduSessionRequest) GetBinaryDataN4InformationExt2() *os.File {
	if o == nil || isNil(o.BinaryDataN4InformationExt2) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN4InformationExt2
}

// GetBinaryDataN4InformationExt2Ok returns a tuple with the BinaryDataN4InformationExt2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePduSessionRequest) GetBinaryDataN4InformationExt2Ok() (**os.File, bool) {
	if o == nil || isNil(o.BinaryDataN4InformationExt2) {
    return nil, false
	}
	return o.BinaryDataN4InformationExt2, true
}

// HasBinaryDataN4InformationExt2 returns a boolean if a field has been set.
func (o *UpdatePduSessionRequest) HasBinaryDataN4InformationExt2() bool {
	if o != nil && !isNil(o.BinaryDataN4InformationExt2) {
		return true
	}

	return false
}

// SetBinaryDataN4InformationExt2 gets a reference to the given *os.File and assigns it to the BinaryDataN4InformationExt2 field.
func (o *UpdatePduSessionRequest) SetBinaryDataN4InformationExt2(v *os.File) {
	o.BinaryDataN4InformationExt2 = &v
}

func (o UpdatePduSessionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.JsonData) {
		toSerialize["jsonData"] = o.JsonData
	}
	if !isNil(o.BinaryDataN1SmInfoFromUe) {
		toSerialize["binaryDataN1SmInfoFromUe"] = o.BinaryDataN1SmInfoFromUe
	}
	if !isNil(o.BinaryDataUnknownN1SmInfo) {
		toSerialize["binaryDataUnknownN1SmInfo"] = o.BinaryDataUnknownN1SmInfo
	}
	if !isNil(o.BinaryDataN4Information) {
		toSerialize["binaryDataN4Information"] = o.BinaryDataN4Information
	}
	if !isNil(o.BinaryDataN4InformationExt1) {
		toSerialize["binaryDataN4InformationExt1"] = o.BinaryDataN4InformationExt1
	}
	if !isNil(o.BinaryDataN4InformationExt2) {
		toSerialize["binaryDataN4InformationExt2"] = o.BinaryDataN4InformationExt2
	}
	return json.Marshal(toSerialize)
}

type NullableUpdatePduSessionRequest struct {
	value *UpdatePduSessionRequest
	isSet bool
}

func (v NullableUpdatePduSessionRequest) Get() *UpdatePduSessionRequest {
	return v.value
}

func (v *NullableUpdatePduSessionRequest) Set(val *UpdatePduSessionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePduSessionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePduSessionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePduSessionRequest(val *UpdatePduSessionRequest) *NullableUpdatePduSessionRequest {
	return &NullableUpdatePduSessionRequest{value: val, isSet: true}
}

func (v NullableUpdatePduSessionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatePduSessionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

