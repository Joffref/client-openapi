/*
Namf_EventExposure

AMF Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Namf_EventExposure

import (
	"encoding/json"
)

// UserLocation At least one of eutraLocation, nrLocation and n3gaLocation shall be present. Several of them may be present. 
type UserLocation struct {
	EutraLocation *EutraLocation `json:"eutraLocation,omitempty"`
	NrLocation *NrLocation `json:"nrLocation,omitempty"`
	N3gaLocation *N3gaLocation `json:"n3gaLocation,omitempty"`
	UtraLocation *UtraLocation `json:"utraLocation,omitempty"`
	GeraLocation *GeraLocation `json:"geraLocation,omitempty"`
}

// NewUserLocation instantiates a new UserLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserLocation() *UserLocation {
	this := UserLocation{}
	return &this
}

// NewUserLocationWithDefaults instantiates a new UserLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserLocationWithDefaults() *UserLocation {
	this := UserLocation{}
	return &this
}

// GetEutraLocation returns the EutraLocation field value if set, zero value otherwise.
func (o *UserLocation) GetEutraLocation() EutraLocation {
	if o == nil || isNil(o.EutraLocation) {
		var ret EutraLocation
		return ret
	}
	return *o.EutraLocation
}

// GetEutraLocationOk returns a tuple with the EutraLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLocation) GetEutraLocationOk() (*EutraLocation, bool) {
	if o == nil || isNil(o.EutraLocation) {
    return nil, false
	}
	return o.EutraLocation, true
}

// HasEutraLocation returns a boolean if a field has been set.
func (o *UserLocation) HasEutraLocation() bool {
	if o != nil && !isNil(o.EutraLocation) {
		return true
	}

	return false
}

// SetEutraLocation gets a reference to the given EutraLocation and assigns it to the EutraLocation field.
func (o *UserLocation) SetEutraLocation(v EutraLocation) {
	o.EutraLocation = &v
}

// GetNrLocation returns the NrLocation field value if set, zero value otherwise.
func (o *UserLocation) GetNrLocation() NrLocation {
	if o == nil || isNil(o.NrLocation) {
		var ret NrLocation
		return ret
	}
	return *o.NrLocation
}

// GetNrLocationOk returns a tuple with the NrLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLocation) GetNrLocationOk() (*NrLocation, bool) {
	if o == nil || isNil(o.NrLocation) {
    return nil, false
	}
	return o.NrLocation, true
}

// HasNrLocation returns a boolean if a field has been set.
func (o *UserLocation) HasNrLocation() bool {
	if o != nil && !isNil(o.NrLocation) {
		return true
	}

	return false
}

// SetNrLocation gets a reference to the given NrLocation and assigns it to the NrLocation field.
func (o *UserLocation) SetNrLocation(v NrLocation) {
	o.NrLocation = &v
}

// GetN3gaLocation returns the N3gaLocation field value if set, zero value otherwise.
func (o *UserLocation) GetN3gaLocation() N3gaLocation {
	if o == nil || isNil(o.N3gaLocation) {
		var ret N3gaLocation
		return ret
	}
	return *o.N3gaLocation
}

// GetN3gaLocationOk returns a tuple with the N3gaLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLocation) GetN3gaLocationOk() (*N3gaLocation, bool) {
	if o == nil || isNil(o.N3gaLocation) {
    return nil, false
	}
	return o.N3gaLocation, true
}

// HasN3gaLocation returns a boolean if a field has been set.
func (o *UserLocation) HasN3gaLocation() bool {
	if o != nil && !isNil(o.N3gaLocation) {
		return true
	}

	return false
}

// SetN3gaLocation gets a reference to the given N3gaLocation and assigns it to the N3gaLocation field.
func (o *UserLocation) SetN3gaLocation(v N3gaLocation) {
	o.N3gaLocation = &v
}

// GetUtraLocation returns the UtraLocation field value if set, zero value otherwise.
func (o *UserLocation) GetUtraLocation() UtraLocation {
	if o == nil || isNil(o.UtraLocation) {
		var ret UtraLocation
		return ret
	}
	return *o.UtraLocation
}

// GetUtraLocationOk returns a tuple with the UtraLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLocation) GetUtraLocationOk() (*UtraLocation, bool) {
	if o == nil || isNil(o.UtraLocation) {
    return nil, false
	}
	return o.UtraLocation, true
}

// HasUtraLocation returns a boolean if a field has been set.
func (o *UserLocation) HasUtraLocation() bool {
	if o != nil && !isNil(o.UtraLocation) {
		return true
	}

	return false
}

// SetUtraLocation gets a reference to the given UtraLocation and assigns it to the UtraLocation field.
func (o *UserLocation) SetUtraLocation(v UtraLocation) {
	o.UtraLocation = &v
}

// GetGeraLocation returns the GeraLocation field value if set, zero value otherwise.
func (o *UserLocation) GetGeraLocation() GeraLocation {
	if o == nil || isNil(o.GeraLocation) {
		var ret GeraLocation
		return ret
	}
	return *o.GeraLocation
}

// GetGeraLocationOk returns a tuple with the GeraLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLocation) GetGeraLocationOk() (*GeraLocation, bool) {
	if o == nil || isNil(o.GeraLocation) {
    return nil, false
	}
	return o.GeraLocation, true
}

// HasGeraLocation returns a boolean if a field has been set.
func (o *UserLocation) HasGeraLocation() bool {
	if o != nil && !isNil(o.GeraLocation) {
		return true
	}

	return false
}

// SetGeraLocation gets a reference to the given GeraLocation and assigns it to the GeraLocation field.
func (o *UserLocation) SetGeraLocation(v GeraLocation) {
	o.GeraLocation = &v
}

func (o UserLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EutraLocation) {
		toSerialize["eutraLocation"] = o.EutraLocation
	}
	if !isNil(o.NrLocation) {
		toSerialize["nrLocation"] = o.NrLocation
	}
	if !isNil(o.N3gaLocation) {
		toSerialize["n3gaLocation"] = o.N3gaLocation
	}
	if !isNil(o.UtraLocation) {
		toSerialize["utraLocation"] = o.UtraLocation
	}
	if !isNil(o.GeraLocation) {
		toSerialize["geraLocation"] = o.GeraLocation
	}
	return json.Marshal(toSerialize)
}

type NullableUserLocation struct {
	value *UserLocation
	isSet bool
}

func (v NullableUserLocation) Get() *UserLocation {
	return v.value
}

func (v *NullableUserLocation) Set(val *UserLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableUserLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableUserLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserLocation(val *UserLocation) *NullableUserLocation {
	return &NullableUserLocation{value: val, isSet: true}
}

func (v NullableUserLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

