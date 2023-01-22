/*
NR NRM

OAS 3.0.1 specification of the NR NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package NrNrm

import (
	"encoding/json"
)

// GnbCuCpFunctionSingleAllOf1 struct for GnbCuCpFunctionSingleAllOf1
type GnbCuCpFunctionSingleAllOf1 struct {
	RRMPolicyRatio []RRMPolicyRatioSingle `json:"RRMPolicyRatio,omitempty"`
	NrCellCu []NrCellCuSingle `json:"NrCellCu,omitempty"`
	EPXnC []EPXnCSingle `json:"EP_XnC,omitempty"`
	EPE1 []EPE1Single `json:"EP_E1,omitempty"`
	EPF1C []EPF1CSingle `json:"EP_F1C,omitempty"`
	EPNgC []EPNgCSingle `json:"EP_NgC,omitempty"`
	EPX2C []EPX2CSingle `json:"EP_X2C,omitempty"`
	DANRManagementFunction *DANRManagementFunctionSingle `json:"DANRManagementFunction,omitempty"`
	DESManagementFunction *DESManagementFunctionSingle `json:"DESManagementFunction,omitempty"`
	DMROFunction *DMROFunctionSingle `json:"DMROFunction,omitempty"`
	DLBOFunction *DLBOFunctionSingle `json:"DLBOFunction,omitempty"`
}

// NewGnbCuCpFunctionSingleAllOf1 instantiates a new GnbCuCpFunctionSingleAllOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGnbCuCpFunctionSingleAllOf1() *GnbCuCpFunctionSingleAllOf1 {
	this := GnbCuCpFunctionSingleAllOf1{}
	return &this
}

// NewGnbCuCpFunctionSingleAllOf1WithDefaults instantiates a new GnbCuCpFunctionSingleAllOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGnbCuCpFunctionSingleAllOf1WithDefaults() *GnbCuCpFunctionSingleAllOf1 {
	this := GnbCuCpFunctionSingleAllOf1{}
	return &this
}

// GetRRMPolicyRatio returns the RRMPolicyRatio field value if set, zero value otherwise.
func (o *GnbCuCpFunctionSingleAllOf1) GetRRMPolicyRatio() []RRMPolicyRatioSingle {
	if o == nil || isNil(o.RRMPolicyRatio) {
		var ret []RRMPolicyRatioSingle
		return ret
	}
	return o.RRMPolicyRatio
}

// GetRRMPolicyRatioOk returns a tuple with the RRMPolicyRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GnbCuCpFunctionSingleAllOf1) GetRRMPolicyRatioOk() ([]RRMPolicyRatioSingle, bool) {
	if o == nil || isNil(o.RRMPolicyRatio) {
    return nil, false
	}
	return o.RRMPolicyRatio, true
}

// HasRRMPolicyRatio returns a boolean if a field has been set.
func (o *GnbCuCpFunctionSingleAllOf1) HasRRMPolicyRatio() bool {
	if o != nil && !isNil(o.RRMPolicyRatio) {
		return true
	}

	return false
}

// SetRRMPolicyRatio gets a reference to the given []RRMPolicyRatioSingle and assigns it to the RRMPolicyRatio field.
func (o *GnbCuCpFunctionSingleAllOf1) SetRRMPolicyRatio(v []RRMPolicyRatioSingle) {
	o.RRMPolicyRatio = v
}

// GetNrCellCu returns the NrCellCu field value if set, zero value otherwise.
func (o *GnbCuCpFunctionSingleAllOf1) GetNrCellCu() []NrCellCuSingle {
	if o == nil || isNil(o.NrCellCu) {
		var ret []NrCellCuSingle
		return ret
	}
	return o.NrCellCu
}

// GetNrCellCuOk returns a tuple with the NrCellCu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GnbCuCpFunctionSingleAllOf1) GetNrCellCuOk() ([]NrCellCuSingle, bool) {
	if o == nil || isNil(o.NrCellCu) {
    return nil, false
	}
	return o.NrCellCu, true
}

// HasNrCellCu returns a boolean if a field has been set.
func (o *GnbCuCpFunctionSingleAllOf1) HasNrCellCu() bool {
	if o != nil && !isNil(o.NrCellCu) {
		return true
	}

	return false
}

// SetNrCellCu gets a reference to the given []NrCellCuSingle and assigns it to the NrCellCu field.
func (o *GnbCuCpFunctionSingleAllOf1) SetNrCellCu(v []NrCellCuSingle) {
	o.NrCellCu = v
}

// GetEPXnC returns the EPXnC field value if set, zero value otherwise.
func (o *GnbCuCpFunctionSingleAllOf1) GetEPXnC() []EPXnCSingle {
	if o == nil || isNil(o.EPXnC) {
		var ret []EPXnCSingle
		return ret
	}
	return o.EPXnC
}

// GetEPXnCOk returns a tuple with the EPXnC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GnbCuCpFunctionSingleAllOf1) GetEPXnCOk() ([]EPXnCSingle, bool) {
	if o == nil || isNil(o.EPXnC) {
    return nil, false
	}
	return o.EPXnC, true
}

// HasEPXnC returns a boolean if a field has been set.
func (o *GnbCuCpFunctionSingleAllOf1) HasEPXnC() bool {
	if o != nil && !isNil(o.EPXnC) {
		return true
	}

	return false
}

// SetEPXnC gets a reference to the given []EPXnCSingle and assigns it to the EPXnC field.
func (o *GnbCuCpFunctionSingleAllOf1) SetEPXnC(v []EPXnCSingle) {
	o.EPXnC = v
}

// GetEPE1 returns the EPE1 field value if set, zero value otherwise.
func (o *GnbCuCpFunctionSingleAllOf1) GetEPE1() []EPE1Single {
	if o == nil || isNil(o.EPE1) {
		var ret []EPE1Single
		return ret
	}
	return o.EPE1
}

// GetEPE1Ok returns a tuple with the EPE1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GnbCuCpFunctionSingleAllOf1) GetEPE1Ok() ([]EPE1Single, bool) {
	if o == nil || isNil(o.EPE1) {
    return nil, false
	}
	return o.EPE1, true
}

// HasEPE1 returns a boolean if a field has been set.
func (o *GnbCuCpFunctionSingleAllOf1) HasEPE1() bool {
	if o != nil && !isNil(o.EPE1) {
		return true
	}

	return false
}

// SetEPE1 gets a reference to the given []EPE1Single and assigns it to the EPE1 field.
func (o *GnbCuCpFunctionSingleAllOf1) SetEPE1(v []EPE1Single) {
	o.EPE1 = v
}

// GetEPF1C returns the EPF1C field value if set, zero value otherwise.
func (o *GnbCuCpFunctionSingleAllOf1) GetEPF1C() []EPF1CSingle {
	if o == nil || isNil(o.EPF1C) {
		var ret []EPF1CSingle
		return ret
	}
	return o.EPF1C
}

// GetEPF1COk returns a tuple with the EPF1C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GnbCuCpFunctionSingleAllOf1) GetEPF1COk() ([]EPF1CSingle, bool) {
	if o == nil || isNil(o.EPF1C) {
    return nil, false
	}
	return o.EPF1C, true
}

// HasEPF1C returns a boolean if a field has been set.
func (o *GnbCuCpFunctionSingleAllOf1) HasEPF1C() bool {
	if o != nil && !isNil(o.EPF1C) {
		return true
	}

	return false
}

// SetEPF1C gets a reference to the given []EPF1CSingle and assigns it to the EPF1C field.
func (o *GnbCuCpFunctionSingleAllOf1) SetEPF1C(v []EPF1CSingle) {
	o.EPF1C = v
}

// GetEPNgC returns the EPNgC field value if set, zero value otherwise.
func (o *GnbCuCpFunctionSingleAllOf1) GetEPNgC() []EPNgCSingle {
	if o == nil || isNil(o.EPNgC) {
		var ret []EPNgCSingle
		return ret
	}
	return o.EPNgC
}

// GetEPNgCOk returns a tuple with the EPNgC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GnbCuCpFunctionSingleAllOf1) GetEPNgCOk() ([]EPNgCSingle, bool) {
	if o == nil || isNil(o.EPNgC) {
    return nil, false
	}
	return o.EPNgC, true
}

// HasEPNgC returns a boolean if a field has been set.
func (o *GnbCuCpFunctionSingleAllOf1) HasEPNgC() bool {
	if o != nil && !isNil(o.EPNgC) {
		return true
	}

	return false
}

// SetEPNgC gets a reference to the given []EPNgCSingle and assigns it to the EPNgC field.
func (o *GnbCuCpFunctionSingleAllOf1) SetEPNgC(v []EPNgCSingle) {
	o.EPNgC = v
}

// GetEPX2C returns the EPX2C field value if set, zero value otherwise.
func (o *GnbCuCpFunctionSingleAllOf1) GetEPX2C() []EPX2CSingle {
	if o == nil || isNil(o.EPX2C) {
		var ret []EPX2CSingle
		return ret
	}
	return o.EPX2C
}

// GetEPX2COk returns a tuple with the EPX2C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GnbCuCpFunctionSingleAllOf1) GetEPX2COk() ([]EPX2CSingle, bool) {
	if o == nil || isNil(o.EPX2C) {
    return nil, false
	}
	return o.EPX2C, true
}

// HasEPX2C returns a boolean if a field has been set.
func (o *GnbCuCpFunctionSingleAllOf1) HasEPX2C() bool {
	if o != nil && !isNil(o.EPX2C) {
		return true
	}

	return false
}

// SetEPX2C gets a reference to the given []EPX2CSingle and assigns it to the EPX2C field.
func (o *GnbCuCpFunctionSingleAllOf1) SetEPX2C(v []EPX2CSingle) {
	o.EPX2C = v
}

// GetDANRManagementFunction returns the DANRManagementFunction field value if set, zero value otherwise.
func (o *GnbCuCpFunctionSingleAllOf1) GetDANRManagementFunction() DANRManagementFunctionSingle {
	if o == nil || isNil(o.DANRManagementFunction) {
		var ret DANRManagementFunctionSingle
		return ret
	}
	return *o.DANRManagementFunction
}

// GetDANRManagementFunctionOk returns a tuple with the DANRManagementFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GnbCuCpFunctionSingleAllOf1) GetDANRManagementFunctionOk() (*DANRManagementFunctionSingle, bool) {
	if o == nil || isNil(o.DANRManagementFunction) {
    return nil, false
	}
	return o.DANRManagementFunction, true
}

// HasDANRManagementFunction returns a boolean if a field has been set.
func (o *GnbCuCpFunctionSingleAllOf1) HasDANRManagementFunction() bool {
	if o != nil && !isNil(o.DANRManagementFunction) {
		return true
	}

	return false
}

// SetDANRManagementFunction gets a reference to the given DANRManagementFunctionSingle and assigns it to the DANRManagementFunction field.
func (o *GnbCuCpFunctionSingleAllOf1) SetDANRManagementFunction(v DANRManagementFunctionSingle) {
	o.DANRManagementFunction = &v
}

// GetDESManagementFunction returns the DESManagementFunction field value if set, zero value otherwise.
func (o *GnbCuCpFunctionSingleAllOf1) GetDESManagementFunction() DESManagementFunctionSingle {
	if o == nil || isNil(o.DESManagementFunction) {
		var ret DESManagementFunctionSingle
		return ret
	}
	return *o.DESManagementFunction
}

// GetDESManagementFunctionOk returns a tuple with the DESManagementFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GnbCuCpFunctionSingleAllOf1) GetDESManagementFunctionOk() (*DESManagementFunctionSingle, bool) {
	if o == nil || isNil(o.DESManagementFunction) {
    return nil, false
	}
	return o.DESManagementFunction, true
}

// HasDESManagementFunction returns a boolean if a field has been set.
func (o *GnbCuCpFunctionSingleAllOf1) HasDESManagementFunction() bool {
	if o != nil && !isNil(o.DESManagementFunction) {
		return true
	}

	return false
}

// SetDESManagementFunction gets a reference to the given DESManagementFunctionSingle and assigns it to the DESManagementFunction field.
func (o *GnbCuCpFunctionSingleAllOf1) SetDESManagementFunction(v DESManagementFunctionSingle) {
	o.DESManagementFunction = &v
}

// GetDMROFunction returns the DMROFunction field value if set, zero value otherwise.
func (o *GnbCuCpFunctionSingleAllOf1) GetDMROFunction() DMROFunctionSingle {
	if o == nil || isNil(o.DMROFunction) {
		var ret DMROFunctionSingle
		return ret
	}
	return *o.DMROFunction
}

// GetDMROFunctionOk returns a tuple with the DMROFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GnbCuCpFunctionSingleAllOf1) GetDMROFunctionOk() (*DMROFunctionSingle, bool) {
	if o == nil || isNil(o.DMROFunction) {
    return nil, false
	}
	return o.DMROFunction, true
}

// HasDMROFunction returns a boolean if a field has been set.
func (o *GnbCuCpFunctionSingleAllOf1) HasDMROFunction() bool {
	if o != nil && !isNil(o.DMROFunction) {
		return true
	}

	return false
}

// SetDMROFunction gets a reference to the given DMROFunctionSingle and assigns it to the DMROFunction field.
func (o *GnbCuCpFunctionSingleAllOf1) SetDMROFunction(v DMROFunctionSingle) {
	o.DMROFunction = &v
}

// GetDLBOFunction returns the DLBOFunction field value if set, zero value otherwise.
func (o *GnbCuCpFunctionSingleAllOf1) GetDLBOFunction() DLBOFunctionSingle {
	if o == nil || isNil(o.DLBOFunction) {
		var ret DLBOFunctionSingle
		return ret
	}
	return *o.DLBOFunction
}

// GetDLBOFunctionOk returns a tuple with the DLBOFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GnbCuCpFunctionSingleAllOf1) GetDLBOFunctionOk() (*DLBOFunctionSingle, bool) {
	if o == nil || isNil(o.DLBOFunction) {
    return nil, false
	}
	return o.DLBOFunction, true
}

// HasDLBOFunction returns a boolean if a field has been set.
func (o *GnbCuCpFunctionSingleAllOf1) HasDLBOFunction() bool {
	if o != nil && !isNil(o.DLBOFunction) {
		return true
	}

	return false
}

// SetDLBOFunction gets a reference to the given DLBOFunctionSingle and assigns it to the DLBOFunction field.
func (o *GnbCuCpFunctionSingleAllOf1) SetDLBOFunction(v DLBOFunctionSingle) {
	o.DLBOFunction = &v
}

func (o GnbCuCpFunctionSingleAllOf1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RRMPolicyRatio) {
		toSerialize["RRMPolicyRatio"] = o.RRMPolicyRatio
	}
	if !isNil(o.NrCellCu) {
		toSerialize["NrCellCu"] = o.NrCellCu
	}
	if !isNil(o.EPXnC) {
		toSerialize["EP_XnC"] = o.EPXnC
	}
	if !isNil(o.EPE1) {
		toSerialize["EP_E1"] = o.EPE1
	}
	if !isNil(o.EPF1C) {
		toSerialize["EP_F1C"] = o.EPF1C
	}
	if !isNil(o.EPNgC) {
		toSerialize["EP_NgC"] = o.EPNgC
	}
	if !isNil(o.EPX2C) {
		toSerialize["EP_X2C"] = o.EPX2C
	}
	if !isNil(o.DANRManagementFunction) {
		toSerialize["DANRManagementFunction"] = o.DANRManagementFunction
	}
	if !isNil(o.DESManagementFunction) {
		toSerialize["DESManagementFunction"] = o.DESManagementFunction
	}
	if !isNil(o.DMROFunction) {
		toSerialize["DMROFunction"] = o.DMROFunction
	}
	if !isNil(o.DLBOFunction) {
		toSerialize["DLBOFunction"] = o.DLBOFunction
	}
	return json.Marshal(toSerialize)
}

type NullableGnbCuCpFunctionSingleAllOf1 struct {
	value *GnbCuCpFunctionSingleAllOf1
	isSet bool
}

func (v NullableGnbCuCpFunctionSingleAllOf1) Get() *GnbCuCpFunctionSingleAllOf1 {
	return v.value
}

func (v *NullableGnbCuCpFunctionSingleAllOf1) Set(val *GnbCuCpFunctionSingleAllOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableGnbCuCpFunctionSingleAllOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableGnbCuCpFunctionSingleAllOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGnbCuCpFunctionSingleAllOf1(val *GnbCuCpFunctionSingleAllOf1) *NullableGnbCuCpFunctionSingleAllOf1 {
	return &NullableGnbCuCpFunctionSingleAllOf1{value: val, isSet: true}
}

func (v NullableGnbCuCpFunctionSingleAllOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGnbCuCpFunctionSingleAllOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

