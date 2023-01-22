/*
Nhss_imsSDM

Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nhss_imsSDM

import (
	"encoding/json"
)

// CsLocation Location data in CS domain
type CsLocation struct {
	MscNumber string `json:"mscNumber"`
	VlrNumber string `json:"vlrNumber"`
	PlmnId PlmnId `json:"plmnId"`
	VlrLocation *GeraLocation `json:"vlrLocation,omitempty"`
	CsgInformation *CsgInformation `json:"csgInformation,omitempty"`
	// String with format \"time-numoffset\" optionally appended by \"daylightSavingTime\", where  - \"time-numoffset\" shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \"daylightSavingTime\" shall represent the adjustment that has been made and shall be    encoded as \"+1\" or \"+2\" for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time. 
	TimeZone *string `json:"timeZone,omitempty"`
	EUtranCgi *Ecgi `json:"eUtranCgi,omitempty"`
	Tai *Tai `json:"tai,omitempty"`
}

// NewCsLocation instantiates a new CsLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCsLocation(mscNumber string, vlrNumber string, plmnId PlmnId) *CsLocation {
	this := CsLocation{}
	this.MscNumber = mscNumber
	this.VlrNumber = vlrNumber
	this.PlmnId = plmnId
	return &this
}

// NewCsLocationWithDefaults instantiates a new CsLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCsLocationWithDefaults() *CsLocation {
	this := CsLocation{}
	return &this
}

// GetMscNumber returns the MscNumber field value
func (o *CsLocation) GetMscNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MscNumber
}

// GetMscNumberOk returns a tuple with the MscNumber field value
// and a boolean to check if the value has been set.
func (o *CsLocation) GetMscNumberOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MscNumber, true
}

// SetMscNumber sets field value
func (o *CsLocation) SetMscNumber(v string) {
	o.MscNumber = v
}

// GetVlrNumber returns the VlrNumber field value
func (o *CsLocation) GetVlrNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VlrNumber
}

// GetVlrNumberOk returns a tuple with the VlrNumber field value
// and a boolean to check if the value has been set.
func (o *CsLocation) GetVlrNumberOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.VlrNumber, true
}

// SetVlrNumber sets field value
func (o *CsLocation) SetVlrNumber(v string) {
	o.VlrNumber = v
}

// GetPlmnId returns the PlmnId field value
func (o *CsLocation) GetPlmnId() PlmnId {
	if o == nil {
		var ret PlmnId
		return ret
	}

	return o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value
// and a boolean to check if the value has been set.
func (o *CsLocation) GetPlmnIdOk() (*PlmnId, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PlmnId, true
}

// SetPlmnId sets field value
func (o *CsLocation) SetPlmnId(v PlmnId) {
	o.PlmnId = v
}

// GetVlrLocation returns the VlrLocation field value if set, zero value otherwise.
func (o *CsLocation) GetVlrLocation() GeraLocation {
	if o == nil || isNil(o.VlrLocation) {
		var ret GeraLocation
		return ret
	}
	return *o.VlrLocation
}

// GetVlrLocationOk returns a tuple with the VlrLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsLocation) GetVlrLocationOk() (*GeraLocation, bool) {
	if o == nil || isNil(o.VlrLocation) {
    return nil, false
	}
	return o.VlrLocation, true
}

// HasVlrLocation returns a boolean if a field has been set.
func (o *CsLocation) HasVlrLocation() bool {
	if o != nil && !isNil(o.VlrLocation) {
		return true
	}

	return false
}

// SetVlrLocation gets a reference to the given GeraLocation and assigns it to the VlrLocation field.
func (o *CsLocation) SetVlrLocation(v GeraLocation) {
	o.VlrLocation = &v
}

// GetCsgInformation returns the CsgInformation field value if set, zero value otherwise.
func (o *CsLocation) GetCsgInformation() CsgInformation {
	if o == nil || isNil(o.CsgInformation) {
		var ret CsgInformation
		return ret
	}
	return *o.CsgInformation
}

// GetCsgInformationOk returns a tuple with the CsgInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsLocation) GetCsgInformationOk() (*CsgInformation, bool) {
	if o == nil || isNil(o.CsgInformation) {
    return nil, false
	}
	return o.CsgInformation, true
}

// HasCsgInformation returns a boolean if a field has been set.
func (o *CsLocation) HasCsgInformation() bool {
	if o != nil && !isNil(o.CsgInformation) {
		return true
	}

	return false
}

// SetCsgInformation gets a reference to the given CsgInformation and assigns it to the CsgInformation field.
func (o *CsLocation) SetCsgInformation(v CsgInformation) {
	o.CsgInformation = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *CsLocation) GetTimeZone() string {
	if o == nil || isNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsLocation) GetTimeZoneOk() (*string, bool) {
	if o == nil || isNil(o.TimeZone) {
    return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *CsLocation) HasTimeZone() bool {
	if o != nil && !isNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *CsLocation) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetEUtranCgi returns the EUtranCgi field value if set, zero value otherwise.
func (o *CsLocation) GetEUtranCgi() Ecgi {
	if o == nil || isNil(o.EUtranCgi) {
		var ret Ecgi
		return ret
	}
	return *o.EUtranCgi
}

// GetEUtranCgiOk returns a tuple with the EUtranCgi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsLocation) GetEUtranCgiOk() (*Ecgi, bool) {
	if o == nil || isNil(o.EUtranCgi) {
    return nil, false
	}
	return o.EUtranCgi, true
}

// HasEUtranCgi returns a boolean if a field has been set.
func (o *CsLocation) HasEUtranCgi() bool {
	if o != nil && !isNil(o.EUtranCgi) {
		return true
	}

	return false
}

// SetEUtranCgi gets a reference to the given Ecgi and assigns it to the EUtranCgi field.
func (o *CsLocation) SetEUtranCgi(v Ecgi) {
	o.EUtranCgi = &v
}

// GetTai returns the Tai field value if set, zero value otherwise.
func (o *CsLocation) GetTai() Tai {
	if o == nil || isNil(o.Tai) {
		var ret Tai
		return ret
	}
	return *o.Tai
}

// GetTaiOk returns a tuple with the Tai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsLocation) GetTaiOk() (*Tai, bool) {
	if o == nil || isNil(o.Tai) {
    return nil, false
	}
	return o.Tai, true
}

// HasTai returns a boolean if a field has been set.
func (o *CsLocation) HasTai() bool {
	if o != nil && !isNil(o.Tai) {
		return true
	}

	return false
}

// SetTai gets a reference to the given Tai and assigns it to the Tai field.
func (o *CsLocation) SetTai(v Tai) {
	o.Tai = &v
}

func (o CsLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mscNumber"] = o.MscNumber
	}
	if true {
		toSerialize["vlrNumber"] = o.VlrNumber
	}
	if true {
		toSerialize["plmnId"] = o.PlmnId
	}
	if !isNil(o.VlrLocation) {
		toSerialize["vlrLocation"] = o.VlrLocation
	}
	if !isNil(o.CsgInformation) {
		toSerialize["csgInformation"] = o.CsgInformation
	}
	if !isNil(o.TimeZone) {
		toSerialize["timeZone"] = o.TimeZone
	}
	if !isNil(o.EUtranCgi) {
		toSerialize["eUtranCgi"] = o.EUtranCgi
	}
	if !isNil(o.Tai) {
		toSerialize["tai"] = o.Tai
	}
	return json.Marshal(toSerialize)
}

type NullableCsLocation struct {
	value *CsLocation
	isSet bool
}

func (v NullableCsLocation) Get() *CsLocation {
	return v.value
}

func (v *NullableCsLocation) Set(val *CsLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableCsLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableCsLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCsLocation(val *CsLocation) *NullableCsLocation {
	return &NullableCsLocation{value: val, isSet: true}
}

func (v NullableCsLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCsLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

