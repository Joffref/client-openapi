/*
Nchf_SpendingLimitControl

Nchf Spending Limit Control Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nchf_SpendingLimitControl

import (
	"encoding/json"
	"time"
)

// SpendingLimitStatus Represents the data structure presenting the statuses of policy counters. 
type SpendingLimitStatus struct {
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501. 
	Supi *string `json:"supi,omitempty"`
	NotifId *string `json:"notifId,omitempty"`
	// Status of the requested policy counters. The key of the map is the attribute \"policyCounterId\". 
	StatusInfos *map[string]PolicyCounterInfo `json:"statusInfos,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	Expiry *time.Time `json:"expiry,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
}

// NewSpendingLimitStatus instantiates a new SpendingLimitStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpendingLimitStatus() *SpendingLimitStatus {
	this := SpendingLimitStatus{}
	return &this
}

// NewSpendingLimitStatusWithDefaults instantiates a new SpendingLimitStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpendingLimitStatusWithDefaults() *SpendingLimitStatus {
	this := SpendingLimitStatus{}
	return &this
}

// GetSupi returns the Supi field value if set, zero value otherwise.
func (o *SpendingLimitStatus) GetSupi() string {
	if o == nil || isNil(o.Supi) {
		var ret string
		return ret
	}
	return *o.Supi
}

// GetSupiOk returns a tuple with the Supi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpendingLimitStatus) GetSupiOk() (*string, bool) {
	if o == nil || isNil(o.Supi) {
    return nil, false
	}
	return o.Supi, true
}

// HasSupi returns a boolean if a field has been set.
func (o *SpendingLimitStatus) HasSupi() bool {
	if o != nil && !isNil(o.Supi) {
		return true
	}

	return false
}

// SetSupi gets a reference to the given string and assigns it to the Supi field.
func (o *SpendingLimitStatus) SetSupi(v string) {
	o.Supi = &v
}

// GetNotifId returns the NotifId field value if set, zero value otherwise.
func (o *SpendingLimitStatus) GetNotifId() string {
	if o == nil || isNil(o.NotifId) {
		var ret string
		return ret
	}
	return *o.NotifId
}

// GetNotifIdOk returns a tuple with the NotifId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpendingLimitStatus) GetNotifIdOk() (*string, bool) {
	if o == nil || isNil(o.NotifId) {
    return nil, false
	}
	return o.NotifId, true
}

// HasNotifId returns a boolean if a field has been set.
func (o *SpendingLimitStatus) HasNotifId() bool {
	if o != nil && !isNil(o.NotifId) {
		return true
	}

	return false
}

// SetNotifId gets a reference to the given string and assigns it to the NotifId field.
func (o *SpendingLimitStatus) SetNotifId(v string) {
	o.NotifId = &v
}

// GetStatusInfos returns the StatusInfos field value if set, zero value otherwise.
func (o *SpendingLimitStatus) GetStatusInfos() map[string]PolicyCounterInfo {
	if o == nil || isNil(o.StatusInfos) {
		var ret map[string]PolicyCounterInfo
		return ret
	}
	return *o.StatusInfos
}

// GetStatusInfosOk returns a tuple with the StatusInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpendingLimitStatus) GetStatusInfosOk() (*map[string]PolicyCounterInfo, bool) {
	if o == nil || isNil(o.StatusInfos) {
    return nil, false
	}
	return o.StatusInfos, true
}

// HasStatusInfos returns a boolean if a field has been set.
func (o *SpendingLimitStatus) HasStatusInfos() bool {
	if o != nil && !isNil(o.StatusInfos) {
		return true
	}

	return false
}

// SetStatusInfos gets a reference to the given map[string]PolicyCounterInfo and assigns it to the StatusInfos field.
func (o *SpendingLimitStatus) SetStatusInfos(v map[string]PolicyCounterInfo) {
	o.StatusInfos = &v
}

// GetExpiry returns the Expiry field value if set, zero value otherwise.
func (o *SpendingLimitStatus) GetExpiry() time.Time {
	if o == nil || isNil(o.Expiry) {
		var ret time.Time
		return ret
	}
	return *o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpendingLimitStatus) GetExpiryOk() (*time.Time, bool) {
	if o == nil || isNil(o.Expiry) {
    return nil, false
	}
	return o.Expiry, true
}

// HasExpiry returns a boolean if a field has been set.
func (o *SpendingLimitStatus) HasExpiry() bool {
	if o != nil && !isNil(o.Expiry) {
		return true
	}

	return false
}

// SetExpiry gets a reference to the given time.Time and assigns it to the Expiry field.
func (o *SpendingLimitStatus) SetExpiry(v time.Time) {
	o.Expiry = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *SpendingLimitStatus) GetSupportedFeatures() string {
	if o == nil || isNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpendingLimitStatus) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || isNil(o.SupportedFeatures) {
    return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *SpendingLimitStatus) HasSupportedFeatures() bool {
	if o != nil && !isNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *SpendingLimitStatus) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

func (o SpendingLimitStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Supi) {
		toSerialize["supi"] = o.Supi
	}
	if !isNil(o.NotifId) {
		toSerialize["notifId"] = o.NotifId
	}
	if !isNil(o.StatusInfos) {
		toSerialize["statusInfos"] = o.StatusInfos
	}
	if !isNil(o.Expiry) {
		toSerialize["expiry"] = o.Expiry
	}
	if !isNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	return json.Marshal(toSerialize)
}

type NullableSpendingLimitStatus struct {
	value *SpendingLimitStatus
	isSet bool
}

func (v NullableSpendingLimitStatus) Get() *SpendingLimitStatus {
	return v.value
}

func (v *NullableSpendingLimitStatus) Set(val *SpendingLimitStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSpendingLimitStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSpendingLimitStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpendingLimitStatus(val *SpendingLimitStatus) *NullableSpendingLimitStatus {
	return &NullableSpendingLimitStatus{value: val, isSet: true}
}

func (v NullableSpendingLimitStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpendingLimitStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

