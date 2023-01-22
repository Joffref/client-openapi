/*
Nsmf_PDUSession

SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nsmf_PDUSession

import (
	"encoding/json"
	"time"
)

// VsmfUpdateError Error within Update Response from V-SMF, or from I-SMF to SMF
type VsmfUpdateError struct {
	Error ExtProblemDetails `json:"error"`
	// Procedure Transaction Identifier
	Pti *int32 `json:"pti,omitempty"`
	N1smCause *string `json:"n1smCause,omitempty"`
	N1SmInfoFromUe *RefToBinaryData `json:"n1SmInfoFromUe,omitempty"`
	UnknownN1SmInfo *RefToBinaryData `json:"unknownN1SmInfo,omitempty"`
	FailedToAssignEbiList []Arp `json:"failedToAssignEbiList,omitempty"`
	NgApCause *NgApCause `json:"ngApCause,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Var5gMmCauseValue *int32 `json:"5gMmCauseValue,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	RecoveryTime *time.Time `json:"recoveryTime,omitempty"`
	N4Info *N4Information `json:"n4Info,omitempty"`
	N4InfoExt1 *N4Information `json:"n4InfoExt1,omitempty"`
	N4InfoExt2 *N4Information `json:"n4InfoExt2,omitempty"`
	N4InfoExt3 *N4Information `json:"n4InfoExt3,omitempty"`
}

// NewVsmfUpdateError instantiates a new VsmfUpdateError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVsmfUpdateError(error_ ExtProblemDetails) *VsmfUpdateError {
	this := VsmfUpdateError{}
	this.Error = error_
	return &this
}

// NewVsmfUpdateErrorWithDefaults instantiates a new VsmfUpdateError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVsmfUpdateErrorWithDefaults() *VsmfUpdateError {
	this := VsmfUpdateError{}
	return &this
}

// GetError returns the Error field value
func (o *VsmfUpdateError) GetError() ExtProblemDetails {
	if o == nil {
		var ret ExtProblemDetails
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *VsmfUpdateError) GetErrorOk() (*ExtProblemDetails, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *VsmfUpdateError) SetError(v ExtProblemDetails) {
	o.Error = v
}

// GetPti returns the Pti field value if set, zero value otherwise.
func (o *VsmfUpdateError) GetPti() int32 {
	if o == nil || isNil(o.Pti) {
		var ret int32
		return ret
	}
	return *o.Pti
}

// GetPtiOk returns a tuple with the Pti field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VsmfUpdateError) GetPtiOk() (*int32, bool) {
	if o == nil || isNil(o.Pti) {
    return nil, false
	}
	return o.Pti, true
}

// HasPti returns a boolean if a field has been set.
func (o *VsmfUpdateError) HasPti() bool {
	if o != nil && !isNil(o.Pti) {
		return true
	}

	return false
}

// SetPti gets a reference to the given int32 and assigns it to the Pti field.
func (o *VsmfUpdateError) SetPti(v int32) {
	o.Pti = &v
}

// GetN1smCause returns the N1smCause field value if set, zero value otherwise.
func (o *VsmfUpdateError) GetN1smCause() string {
	if o == nil || isNil(o.N1smCause) {
		var ret string
		return ret
	}
	return *o.N1smCause
}

// GetN1smCauseOk returns a tuple with the N1smCause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VsmfUpdateError) GetN1smCauseOk() (*string, bool) {
	if o == nil || isNil(o.N1smCause) {
    return nil, false
	}
	return o.N1smCause, true
}

// HasN1smCause returns a boolean if a field has been set.
func (o *VsmfUpdateError) HasN1smCause() bool {
	if o != nil && !isNil(o.N1smCause) {
		return true
	}

	return false
}

// SetN1smCause gets a reference to the given string and assigns it to the N1smCause field.
func (o *VsmfUpdateError) SetN1smCause(v string) {
	o.N1smCause = &v
}

// GetN1SmInfoFromUe returns the N1SmInfoFromUe field value if set, zero value otherwise.
func (o *VsmfUpdateError) GetN1SmInfoFromUe() RefToBinaryData {
	if o == nil || isNil(o.N1SmInfoFromUe) {
		var ret RefToBinaryData
		return ret
	}
	return *o.N1SmInfoFromUe
}

// GetN1SmInfoFromUeOk returns a tuple with the N1SmInfoFromUe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VsmfUpdateError) GetN1SmInfoFromUeOk() (*RefToBinaryData, bool) {
	if o == nil || isNil(o.N1SmInfoFromUe) {
    return nil, false
	}
	return o.N1SmInfoFromUe, true
}

// HasN1SmInfoFromUe returns a boolean if a field has been set.
func (o *VsmfUpdateError) HasN1SmInfoFromUe() bool {
	if o != nil && !isNil(o.N1SmInfoFromUe) {
		return true
	}

	return false
}

// SetN1SmInfoFromUe gets a reference to the given RefToBinaryData and assigns it to the N1SmInfoFromUe field.
func (o *VsmfUpdateError) SetN1SmInfoFromUe(v RefToBinaryData) {
	o.N1SmInfoFromUe = &v
}

// GetUnknownN1SmInfo returns the UnknownN1SmInfo field value if set, zero value otherwise.
func (o *VsmfUpdateError) GetUnknownN1SmInfo() RefToBinaryData {
	if o == nil || isNil(o.UnknownN1SmInfo) {
		var ret RefToBinaryData
		return ret
	}
	return *o.UnknownN1SmInfo
}

// GetUnknownN1SmInfoOk returns a tuple with the UnknownN1SmInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VsmfUpdateError) GetUnknownN1SmInfoOk() (*RefToBinaryData, bool) {
	if o == nil || isNil(o.UnknownN1SmInfo) {
    return nil, false
	}
	return o.UnknownN1SmInfo, true
}

// HasUnknownN1SmInfo returns a boolean if a field has been set.
func (o *VsmfUpdateError) HasUnknownN1SmInfo() bool {
	if o != nil && !isNil(o.UnknownN1SmInfo) {
		return true
	}

	return false
}

// SetUnknownN1SmInfo gets a reference to the given RefToBinaryData and assigns it to the UnknownN1SmInfo field.
func (o *VsmfUpdateError) SetUnknownN1SmInfo(v RefToBinaryData) {
	o.UnknownN1SmInfo = &v
}

// GetFailedToAssignEbiList returns the FailedToAssignEbiList field value if set, zero value otherwise.
func (o *VsmfUpdateError) GetFailedToAssignEbiList() []Arp {
	if o == nil || isNil(o.FailedToAssignEbiList) {
		var ret []Arp
		return ret
	}
	return o.FailedToAssignEbiList
}

// GetFailedToAssignEbiListOk returns a tuple with the FailedToAssignEbiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VsmfUpdateError) GetFailedToAssignEbiListOk() ([]Arp, bool) {
	if o == nil || isNil(o.FailedToAssignEbiList) {
    return nil, false
	}
	return o.FailedToAssignEbiList, true
}

// HasFailedToAssignEbiList returns a boolean if a field has been set.
func (o *VsmfUpdateError) HasFailedToAssignEbiList() bool {
	if o != nil && !isNil(o.FailedToAssignEbiList) {
		return true
	}

	return false
}

// SetFailedToAssignEbiList gets a reference to the given []Arp and assigns it to the FailedToAssignEbiList field.
func (o *VsmfUpdateError) SetFailedToAssignEbiList(v []Arp) {
	o.FailedToAssignEbiList = v
}

// GetNgApCause returns the NgApCause field value if set, zero value otherwise.
func (o *VsmfUpdateError) GetNgApCause() NgApCause {
	if o == nil || isNil(o.NgApCause) {
		var ret NgApCause
		return ret
	}
	return *o.NgApCause
}

// GetNgApCauseOk returns a tuple with the NgApCause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VsmfUpdateError) GetNgApCauseOk() (*NgApCause, bool) {
	if o == nil || isNil(o.NgApCause) {
    return nil, false
	}
	return o.NgApCause, true
}

// HasNgApCause returns a boolean if a field has been set.
func (o *VsmfUpdateError) HasNgApCause() bool {
	if o != nil && !isNil(o.NgApCause) {
		return true
	}

	return false
}

// SetNgApCause gets a reference to the given NgApCause and assigns it to the NgApCause field.
func (o *VsmfUpdateError) SetNgApCause(v NgApCause) {
	o.NgApCause = &v
}

// GetVar5gMmCauseValue returns the Var5gMmCauseValue field value if set, zero value otherwise.
func (o *VsmfUpdateError) GetVar5gMmCauseValue() int32 {
	if o == nil || isNil(o.Var5gMmCauseValue) {
		var ret int32
		return ret
	}
	return *o.Var5gMmCauseValue
}

// GetVar5gMmCauseValueOk returns a tuple with the Var5gMmCauseValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VsmfUpdateError) GetVar5gMmCauseValueOk() (*int32, bool) {
	if o == nil || isNil(o.Var5gMmCauseValue) {
    return nil, false
	}
	return o.Var5gMmCauseValue, true
}

// HasVar5gMmCauseValue returns a boolean if a field has been set.
func (o *VsmfUpdateError) HasVar5gMmCauseValue() bool {
	if o != nil && !isNil(o.Var5gMmCauseValue) {
		return true
	}

	return false
}

// SetVar5gMmCauseValue gets a reference to the given int32 and assigns it to the Var5gMmCauseValue field.
func (o *VsmfUpdateError) SetVar5gMmCauseValue(v int32) {
	o.Var5gMmCauseValue = &v
}

// GetRecoveryTime returns the RecoveryTime field value if set, zero value otherwise.
func (o *VsmfUpdateError) GetRecoveryTime() time.Time {
	if o == nil || isNil(o.RecoveryTime) {
		var ret time.Time
		return ret
	}
	return *o.RecoveryTime
}

// GetRecoveryTimeOk returns a tuple with the RecoveryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VsmfUpdateError) GetRecoveryTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.RecoveryTime) {
    return nil, false
	}
	return o.RecoveryTime, true
}

// HasRecoveryTime returns a boolean if a field has been set.
func (o *VsmfUpdateError) HasRecoveryTime() bool {
	if o != nil && !isNil(o.RecoveryTime) {
		return true
	}

	return false
}

// SetRecoveryTime gets a reference to the given time.Time and assigns it to the RecoveryTime field.
func (o *VsmfUpdateError) SetRecoveryTime(v time.Time) {
	o.RecoveryTime = &v
}

// GetN4Info returns the N4Info field value if set, zero value otherwise.
func (o *VsmfUpdateError) GetN4Info() N4Information {
	if o == nil || isNil(o.N4Info) {
		var ret N4Information
		return ret
	}
	return *o.N4Info
}

// GetN4InfoOk returns a tuple with the N4Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VsmfUpdateError) GetN4InfoOk() (*N4Information, bool) {
	if o == nil || isNil(o.N4Info) {
    return nil, false
	}
	return o.N4Info, true
}

// HasN4Info returns a boolean if a field has been set.
func (o *VsmfUpdateError) HasN4Info() bool {
	if o != nil && !isNil(o.N4Info) {
		return true
	}

	return false
}

// SetN4Info gets a reference to the given N4Information and assigns it to the N4Info field.
func (o *VsmfUpdateError) SetN4Info(v N4Information) {
	o.N4Info = &v
}

// GetN4InfoExt1 returns the N4InfoExt1 field value if set, zero value otherwise.
func (o *VsmfUpdateError) GetN4InfoExt1() N4Information {
	if o == nil || isNil(o.N4InfoExt1) {
		var ret N4Information
		return ret
	}
	return *o.N4InfoExt1
}

// GetN4InfoExt1Ok returns a tuple with the N4InfoExt1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VsmfUpdateError) GetN4InfoExt1Ok() (*N4Information, bool) {
	if o == nil || isNil(o.N4InfoExt1) {
    return nil, false
	}
	return o.N4InfoExt1, true
}

// HasN4InfoExt1 returns a boolean if a field has been set.
func (o *VsmfUpdateError) HasN4InfoExt1() bool {
	if o != nil && !isNil(o.N4InfoExt1) {
		return true
	}

	return false
}

// SetN4InfoExt1 gets a reference to the given N4Information and assigns it to the N4InfoExt1 field.
func (o *VsmfUpdateError) SetN4InfoExt1(v N4Information) {
	o.N4InfoExt1 = &v
}

// GetN4InfoExt2 returns the N4InfoExt2 field value if set, zero value otherwise.
func (o *VsmfUpdateError) GetN4InfoExt2() N4Information {
	if o == nil || isNil(o.N4InfoExt2) {
		var ret N4Information
		return ret
	}
	return *o.N4InfoExt2
}

// GetN4InfoExt2Ok returns a tuple with the N4InfoExt2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VsmfUpdateError) GetN4InfoExt2Ok() (*N4Information, bool) {
	if o == nil || isNil(o.N4InfoExt2) {
    return nil, false
	}
	return o.N4InfoExt2, true
}

// HasN4InfoExt2 returns a boolean if a field has been set.
func (o *VsmfUpdateError) HasN4InfoExt2() bool {
	if o != nil && !isNil(o.N4InfoExt2) {
		return true
	}

	return false
}

// SetN4InfoExt2 gets a reference to the given N4Information and assigns it to the N4InfoExt2 field.
func (o *VsmfUpdateError) SetN4InfoExt2(v N4Information) {
	o.N4InfoExt2 = &v
}

// GetN4InfoExt3 returns the N4InfoExt3 field value if set, zero value otherwise.
func (o *VsmfUpdateError) GetN4InfoExt3() N4Information {
	if o == nil || isNil(o.N4InfoExt3) {
		var ret N4Information
		return ret
	}
	return *o.N4InfoExt3
}

// GetN4InfoExt3Ok returns a tuple with the N4InfoExt3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VsmfUpdateError) GetN4InfoExt3Ok() (*N4Information, bool) {
	if o == nil || isNil(o.N4InfoExt3) {
    return nil, false
	}
	return o.N4InfoExt3, true
}

// HasN4InfoExt3 returns a boolean if a field has been set.
func (o *VsmfUpdateError) HasN4InfoExt3() bool {
	if o != nil && !isNil(o.N4InfoExt3) {
		return true
	}

	return false
}

// SetN4InfoExt3 gets a reference to the given N4Information and assigns it to the N4InfoExt3 field.
func (o *VsmfUpdateError) SetN4InfoExt3(v N4Information) {
	o.N4InfoExt3 = &v
}

func (o VsmfUpdateError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["error"] = o.Error
	}
	if !isNil(o.Pti) {
		toSerialize["pti"] = o.Pti
	}
	if !isNil(o.N1smCause) {
		toSerialize["n1smCause"] = o.N1smCause
	}
	if !isNil(o.N1SmInfoFromUe) {
		toSerialize["n1SmInfoFromUe"] = o.N1SmInfoFromUe
	}
	if !isNil(o.UnknownN1SmInfo) {
		toSerialize["unknownN1SmInfo"] = o.UnknownN1SmInfo
	}
	if !isNil(o.FailedToAssignEbiList) {
		toSerialize["failedToAssignEbiList"] = o.FailedToAssignEbiList
	}
	if !isNil(o.NgApCause) {
		toSerialize["ngApCause"] = o.NgApCause
	}
	if !isNil(o.Var5gMmCauseValue) {
		toSerialize["5gMmCauseValue"] = o.Var5gMmCauseValue
	}
	if !isNil(o.RecoveryTime) {
		toSerialize["recoveryTime"] = o.RecoveryTime
	}
	if !isNil(o.N4Info) {
		toSerialize["n4Info"] = o.N4Info
	}
	if !isNil(o.N4InfoExt1) {
		toSerialize["n4InfoExt1"] = o.N4InfoExt1
	}
	if !isNil(o.N4InfoExt2) {
		toSerialize["n4InfoExt2"] = o.N4InfoExt2
	}
	if !isNil(o.N4InfoExt3) {
		toSerialize["n4InfoExt3"] = o.N4InfoExt3
	}
	return json.Marshal(toSerialize)
}

type NullableVsmfUpdateError struct {
	value *VsmfUpdateError
	isSet bool
}

func (v NullableVsmfUpdateError) Get() *VsmfUpdateError {
	return v.value
}

func (v *NullableVsmfUpdateError) Set(val *VsmfUpdateError) {
	v.value = val
	v.isSet = true
}

func (v NullableVsmfUpdateError) IsSet() bool {
	return v.isSet
}

func (v *NullableVsmfUpdateError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVsmfUpdateError(val *VsmfUpdateError) *NullableVsmfUpdateError {
	return &NullableVsmfUpdateError{value: val, isSet: true}
}

func (v NullableVsmfUpdateError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVsmfUpdateError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

