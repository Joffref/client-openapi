/*
Unified Data Repository Service API file for policy data

The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Policy_Data

import (
	"encoding/json"
)

// ChargingInformation Contains the addresses of the charging functions.
type ChargingInformation struct {
	// String providing an URI formatted according to RFC 3986.
	PrimaryChfAddress string `json:"primaryChfAddress"`
	// String providing an URI formatted according to RFC 3986.
	SecondaryChfAddress *string `json:"secondaryChfAddress,omitempty"`
	// NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \"set<Set ID>.<nftype>set.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.<NFType>set.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)  <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NFType> encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters <Set ID> encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.  
	PrimaryChfSetId *string `json:"primaryChfSetId,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	PrimaryChfInstanceId *string `json:"primaryChfInstanceId,omitempty"`
	// NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \"set<Set ID>.<nftype>set.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.<NFType>set.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)  <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NFType> encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters <Set ID> encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.  
	SecondaryChfSetId *string `json:"secondaryChfSetId,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	SecondaryChfInstanceId *string `json:"secondaryChfInstanceId,omitempty"`
}

// NewChargingInformation instantiates a new ChargingInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChargingInformation(primaryChfAddress string) *ChargingInformation {
	this := ChargingInformation{}
	this.PrimaryChfAddress = primaryChfAddress
	return &this
}

// NewChargingInformationWithDefaults instantiates a new ChargingInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChargingInformationWithDefaults() *ChargingInformation {
	this := ChargingInformation{}
	return &this
}

// GetPrimaryChfAddress returns the PrimaryChfAddress field value
func (o *ChargingInformation) GetPrimaryChfAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrimaryChfAddress
}

// GetPrimaryChfAddressOk returns a tuple with the PrimaryChfAddress field value
// and a boolean to check if the value has been set.
func (o *ChargingInformation) GetPrimaryChfAddressOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PrimaryChfAddress, true
}

// SetPrimaryChfAddress sets field value
func (o *ChargingInformation) SetPrimaryChfAddress(v string) {
	o.PrimaryChfAddress = v
}

// GetSecondaryChfAddress returns the SecondaryChfAddress field value if set, zero value otherwise.
func (o *ChargingInformation) GetSecondaryChfAddress() string {
	if o == nil || isNil(o.SecondaryChfAddress) {
		var ret string
		return ret
	}
	return *o.SecondaryChfAddress
}

// GetSecondaryChfAddressOk returns a tuple with the SecondaryChfAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargingInformation) GetSecondaryChfAddressOk() (*string, bool) {
	if o == nil || isNil(o.SecondaryChfAddress) {
    return nil, false
	}
	return o.SecondaryChfAddress, true
}

// HasSecondaryChfAddress returns a boolean if a field has been set.
func (o *ChargingInformation) HasSecondaryChfAddress() bool {
	if o != nil && !isNil(o.SecondaryChfAddress) {
		return true
	}

	return false
}

// SetSecondaryChfAddress gets a reference to the given string and assigns it to the SecondaryChfAddress field.
func (o *ChargingInformation) SetSecondaryChfAddress(v string) {
	o.SecondaryChfAddress = &v
}

// GetPrimaryChfSetId returns the PrimaryChfSetId field value if set, zero value otherwise.
func (o *ChargingInformation) GetPrimaryChfSetId() string {
	if o == nil || isNil(o.PrimaryChfSetId) {
		var ret string
		return ret
	}
	return *o.PrimaryChfSetId
}

// GetPrimaryChfSetIdOk returns a tuple with the PrimaryChfSetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargingInformation) GetPrimaryChfSetIdOk() (*string, bool) {
	if o == nil || isNil(o.PrimaryChfSetId) {
    return nil, false
	}
	return o.PrimaryChfSetId, true
}

// HasPrimaryChfSetId returns a boolean if a field has been set.
func (o *ChargingInformation) HasPrimaryChfSetId() bool {
	if o != nil && !isNil(o.PrimaryChfSetId) {
		return true
	}

	return false
}

// SetPrimaryChfSetId gets a reference to the given string and assigns it to the PrimaryChfSetId field.
func (o *ChargingInformation) SetPrimaryChfSetId(v string) {
	o.PrimaryChfSetId = &v
}

// GetPrimaryChfInstanceId returns the PrimaryChfInstanceId field value if set, zero value otherwise.
func (o *ChargingInformation) GetPrimaryChfInstanceId() string {
	if o == nil || isNil(o.PrimaryChfInstanceId) {
		var ret string
		return ret
	}
	return *o.PrimaryChfInstanceId
}

// GetPrimaryChfInstanceIdOk returns a tuple with the PrimaryChfInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargingInformation) GetPrimaryChfInstanceIdOk() (*string, bool) {
	if o == nil || isNil(o.PrimaryChfInstanceId) {
    return nil, false
	}
	return o.PrimaryChfInstanceId, true
}

// HasPrimaryChfInstanceId returns a boolean if a field has been set.
func (o *ChargingInformation) HasPrimaryChfInstanceId() bool {
	if o != nil && !isNil(o.PrimaryChfInstanceId) {
		return true
	}

	return false
}

// SetPrimaryChfInstanceId gets a reference to the given string and assigns it to the PrimaryChfInstanceId field.
func (o *ChargingInformation) SetPrimaryChfInstanceId(v string) {
	o.PrimaryChfInstanceId = &v
}

// GetSecondaryChfSetId returns the SecondaryChfSetId field value if set, zero value otherwise.
func (o *ChargingInformation) GetSecondaryChfSetId() string {
	if o == nil || isNil(o.SecondaryChfSetId) {
		var ret string
		return ret
	}
	return *o.SecondaryChfSetId
}

// GetSecondaryChfSetIdOk returns a tuple with the SecondaryChfSetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargingInformation) GetSecondaryChfSetIdOk() (*string, bool) {
	if o == nil || isNil(o.SecondaryChfSetId) {
    return nil, false
	}
	return o.SecondaryChfSetId, true
}

// HasSecondaryChfSetId returns a boolean if a field has been set.
func (o *ChargingInformation) HasSecondaryChfSetId() bool {
	if o != nil && !isNil(o.SecondaryChfSetId) {
		return true
	}

	return false
}

// SetSecondaryChfSetId gets a reference to the given string and assigns it to the SecondaryChfSetId field.
func (o *ChargingInformation) SetSecondaryChfSetId(v string) {
	o.SecondaryChfSetId = &v
}

// GetSecondaryChfInstanceId returns the SecondaryChfInstanceId field value if set, zero value otherwise.
func (o *ChargingInformation) GetSecondaryChfInstanceId() string {
	if o == nil || isNil(o.SecondaryChfInstanceId) {
		var ret string
		return ret
	}
	return *o.SecondaryChfInstanceId
}

// GetSecondaryChfInstanceIdOk returns a tuple with the SecondaryChfInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargingInformation) GetSecondaryChfInstanceIdOk() (*string, bool) {
	if o == nil || isNil(o.SecondaryChfInstanceId) {
    return nil, false
	}
	return o.SecondaryChfInstanceId, true
}

// HasSecondaryChfInstanceId returns a boolean if a field has been set.
func (o *ChargingInformation) HasSecondaryChfInstanceId() bool {
	if o != nil && !isNil(o.SecondaryChfInstanceId) {
		return true
	}

	return false
}

// SetSecondaryChfInstanceId gets a reference to the given string and assigns it to the SecondaryChfInstanceId field.
func (o *ChargingInformation) SetSecondaryChfInstanceId(v string) {
	o.SecondaryChfInstanceId = &v
}

func (o ChargingInformation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["primaryChfAddress"] = o.PrimaryChfAddress
	}
	if !isNil(o.SecondaryChfAddress) {
		toSerialize["secondaryChfAddress"] = o.SecondaryChfAddress
	}
	if !isNil(o.PrimaryChfSetId) {
		toSerialize["primaryChfSetId"] = o.PrimaryChfSetId
	}
	if !isNil(o.PrimaryChfInstanceId) {
		toSerialize["primaryChfInstanceId"] = o.PrimaryChfInstanceId
	}
	if !isNil(o.SecondaryChfSetId) {
		toSerialize["secondaryChfSetId"] = o.SecondaryChfSetId
	}
	if !isNil(o.SecondaryChfInstanceId) {
		toSerialize["secondaryChfInstanceId"] = o.SecondaryChfInstanceId
	}
	return json.Marshal(toSerialize)
}

type NullableChargingInformation struct {
	value *ChargingInformation
	isSet bool
}

func (v NullableChargingInformation) Get() *ChargingInformation {
	return v.value
}

func (v *NullableChargingInformation) Set(val *ChargingInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableChargingInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableChargingInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChargingInformation(val *ChargingInformation) *NullableChargingInformation {
	return &NullableChargingInformation{value: val, isSet: true}
}

func (v NullableChargingInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChargingInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

