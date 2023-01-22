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

// ImsAssociatedIdentities A list of identities belonging to the same Implicit Registration Set (IRS), along with the registration state of the IRS 
type ImsAssociatedIdentities struct {
	IrsState ImsRegistrationState `json:"irsState"`
	PublicIdentities PublicIdentities `json:"publicIdentities"`
}

// NewImsAssociatedIdentities instantiates a new ImsAssociatedIdentities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImsAssociatedIdentities(irsState ImsRegistrationState, publicIdentities PublicIdentities) *ImsAssociatedIdentities {
	this := ImsAssociatedIdentities{}
	this.IrsState = irsState
	this.PublicIdentities = publicIdentities
	return &this
}

// NewImsAssociatedIdentitiesWithDefaults instantiates a new ImsAssociatedIdentities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImsAssociatedIdentitiesWithDefaults() *ImsAssociatedIdentities {
	this := ImsAssociatedIdentities{}
	return &this
}

// GetIrsState returns the IrsState field value
func (o *ImsAssociatedIdentities) GetIrsState() ImsRegistrationState {
	if o == nil {
		var ret ImsRegistrationState
		return ret
	}

	return o.IrsState
}

// GetIrsStateOk returns a tuple with the IrsState field value
// and a boolean to check if the value has been set.
func (o *ImsAssociatedIdentities) GetIrsStateOk() (*ImsRegistrationState, bool) {
	if o == nil {
    return nil, false
	}
	return &o.IrsState, true
}

// SetIrsState sets field value
func (o *ImsAssociatedIdentities) SetIrsState(v ImsRegistrationState) {
	o.IrsState = v
}

// GetPublicIdentities returns the PublicIdentities field value
func (o *ImsAssociatedIdentities) GetPublicIdentities() PublicIdentities {
	if o == nil {
		var ret PublicIdentities
		return ret
	}

	return o.PublicIdentities
}

// GetPublicIdentitiesOk returns a tuple with the PublicIdentities field value
// and a boolean to check if the value has been set.
func (o *ImsAssociatedIdentities) GetPublicIdentitiesOk() (*PublicIdentities, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PublicIdentities, true
}

// SetPublicIdentities sets field value
func (o *ImsAssociatedIdentities) SetPublicIdentities(v PublicIdentities) {
	o.PublicIdentities = v
}

func (o ImsAssociatedIdentities) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["irsState"] = o.IrsState
	}
	if true {
		toSerialize["publicIdentities"] = o.PublicIdentities
	}
	return json.Marshal(toSerialize)
}

type NullableImsAssociatedIdentities struct {
	value *ImsAssociatedIdentities
	isSet bool
}

func (v NullableImsAssociatedIdentities) Get() *ImsAssociatedIdentities {
	return v.value
}

func (v *NullableImsAssociatedIdentities) Set(val *ImsAssociatedIdentities) {
	v.value = val
	v.isSet = true
}

func (v NullableImsAssociatedIdentities) IsSet() bool {
	return v.isSet
}

func (v *NullableImsAssociatedIdentities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImsAssociatedIdentities(val *ImsAssociatedIdentities) *NullableImsAssociatedIdentities {
	return &NullableImsAssociatedIdentities{value: val, isSet: true}
}

func (v NullableImsAssociatedIdentities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImsAssociatedIdentities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

