/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudr_DR

import (
	"encoding/json"
)

// SmsManagementSubscriptionData struct for SmsManagementSubscriptionData
type SmsManagementSubscriptionData struct {
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
	MtSmsSubscribed *bool `json:"mtSmsSubscribed,omitempty"`
	MtSmsBarringAll *bool `json:"mtSmsBarringAll,omitempty"`
	MtSmsBarringRoaming *bool `json:"mtSmsBarringRoaming,omitempty"`
	MoSmsSubscribed *bool `json:"moSmsSubscribed,omitempty"`
	MoSmsBarringAll *bool `json:"moSmsBarringAll,omitempty"`
	MoSmsBarringRoaming *bool `json:"moSmsBarringRoaming,omitempty"`
	SharedSmsMngDataIds []string `json:"sharedSmsMngDataIds,omitempty"`
	TraceData NullableTraceData1 `json:"traceData,omitempty"`
}

// NewSmsManagementSubscriptionData instantiates a new SmsManagementSubscriptionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmsManagementSubscriptionData() *SmsManagementSubscriptionData {
	this := SmsManagementSubscriptionData{}
	return &this
}

// NewSmsManagementSubscriptionDataWithDefaults instantiates a new SmsManagementSubscriptionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmsManagementSubscriptionDataWithDefaults() *SmsManagementSubscriptionData {
	this := SmsManagementSubscriptionData{}
	return &this
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *SmsManagementSubscriptionData) GetSupportedFeatures() string {
	if o == nil || isNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsManagementSubscriptionData) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || isNil(o.SupportedFeatures) {
    return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *SmsManagementSubscriptionData) HasSupportedFeatures() bool {
	if o != nil && !isNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *SmsManagementSubscriptionData) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetMtSmsSubscribed returns the MtSmsSubscribed field value if set, zero value otherwise.
func (o *SmsManagementSubscriptionData) GetMtSmsSubscribed() bool {
	if o == nil || isNil(o.MtSmsSubscribed) {
		var ret bool
		return ret
	}
	return *o.MtSmsSubscribed
}

// GetMtSmsSubscribedOk returns a tuple with the MtSmsSubscribed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsManagementSubscriptionData) GetMtSmsSubscribedOk() (*bool, bool) {
	if o == nil || isNil(o.MtSmsSubscribed) {
    return nil, false
	}
	return o.MtSmsSubscribed, true
}

// HasMtSmsSubscribed returns a boolean if a field has been set.
func (o *SmsManagementSubscriptionData) HasMtSmsSubscribed() bool {
	if o != nil && !isNil(o.MtSmsSubscribed) {
		return true
	}

	return false
}

// SetMtSmsSubscribed gets a reference to the given bool and assigns it to the MtSmsSubscribed field.
func (o *SmsManagementSubscriptionData) SetMtSmsSubscribed(v bool) {
	o.MtSmsSubscribed = &v
}

// GetMtSmsBarringAll returns the MtSmsBarringAll field value if set, zero value otherwise.
func (o *SmsManagementSubscriptionData) GetMtSmsBarringAll() bool {
	if o == nil || isNil(o.MtSmsBarringAll) {
		var ret bool
		return ret
	}
	return *o.MtSmsBarringAll
}

// GetMtSmsBarringAllOk returns a tuple with the MtSmsBarringAll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsManagementSubscriptionData) GetMtSmsBarringAllOk() (*bool, bool) {
	if o == nil || isNil(o.MtSmsBarringAll) {
    return nil, false
	}
	return o.MtSmsBarringAll, true
}

// HasMtSmsBarringAll returns a boolean if a field has been set.
func (o *SmsManagementSubscriptionData) HasMtSmsBarringAll() bool {
	if o != nil && !isNil(o.MtSmsBarringAll) {
		return true
	}

	return false
}

// SetMtSmsBarringAll gets a reference to the given bool and assigns it to the MtSmsBarringAll field.
func (o *SmsManagementSubscriptionData) SetMtSmsBarringAll(v bool) {
	o.MtSmsBarringAll = &v
}

// GetMtSmsBarringRoaming returns the MtSmsBarringRoaming field value if set, zero value otherwise.
func (o *SmsManagementSubscriptionData) GetMtSmsBarringRoaming() bool {
	if o == nil || isNil(o.MtSmsBarringRoaming) {
		var ret bool
		return ret
	}
	return *o.MtSmsBarringRoaming
}

// GetMtSmsBarringRoamingOk returns a tuple with the MtSmsBarringRoaming field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsManagementSubscriptionData) GetMtSmsBarringRoamingOk() (*bool, bool) {
	if o == nil || isNil(o.MtSmsBarringRoaming) {
    return nil, false
	}
	return o.MtSmsBarringRoaming, true
}

// HasMtSmsBarringRoaming returns a boolean if a field has been set.
func (o *SmsManagementSubscriptionData) HasMtSmsBarringRoaming() bool {
	if o != nil && !isNil(o.MtSmsBarringRoaming) {
		return true
	}

	return false
}

// SetMtSmsBarringRoaming gets a reference to the given bool and assigns it to the MtSmsBarringRoaming field.
func (o *SmsManagementSubscriptionData) SetMtSmsBarringRoaming(v bool) {
	o.MtSmsBarringRoaming = &v
}

// GetMoSmsSubscribed returns the MoSmsSubscribed field value if set, zero value otherwise.
func (o *SmsManagementSubscriptionData) GetMoSmsSubscribed() bool {
	if o == nil || isNil(o.MoSmsSubscribed) {
		var ret bool
		return ret
	}
	return *o.MoSmsSubscribed
}

// GetMoSmsSubscribedOk returns a tuple with the MoSmsSubscribed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsManagementSubscriptionData) GetMoSmsSubscribedOk() (*bool, bool) {
	if o == nil || isNil(o.MoSmsSubscribed) {
    return nil, false
	}
	return o.MoSmsSubscribed, true
}

// HasMoSmsSubscribed returns a boolean if a field has been set.
func (o *SmsManagementSubscriptionData) HasMoSmsSubscribed() bool {
	if o != nil && !isNil(o.MoSmsSubscribed) {
		return true
	}

	return false
}

// SetMoSmsSubscribed gets a reference to the given bool and assigns it to the MoSmsSubscribed field.
func (o *SmsManagementSubscriptionData) SetMoSmsSubscribed(v bool) {
	o.MoSmsSubscribed = &v
}

// GetMoSmsBarringAll returns the MoSmsBarringAll field value if set, zero value otherwise.
func (o *SmsManagementSubscriptionData) GetMoSmsBarringAll() bool {
	if o == nil || isNil(o.MoSmsBarringAll) {
		var ret bool
		return ret
	}
	return *o.MoSmsBarringAll
}

// GetMoSmsBarringAllOk returns a tuple with the MoSmsBarringAll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsManagementSubscriptionData) GetMoSmsBarringAllOk() (*bool, bool) {
	if o == nil || isNil(o.MoSmsBarringAll) {
    return nil, false
	}
	return o.MoSmsBarringAll, true
}

// HasMoSmsBarringAll returns a boolean if a field has been set.
func (o *SmsManagementSubscriptionData) HasMoSmsBarringAll() bool {
	if o != nil && !isNil(o.MoSmsBarringAll) {
		return true
	}

	return false
}

// SetMoSmsBarringAll gets a reference to the given bool and assigns it to the MoSmsBarringAll field.
func (o *SmsManagementSubscriptionData) SetMoSmsBarringAll(v bool) {
	o.MoSmsBarringAll = &v
}

// GetMoSmsBarringRoaming returns the MoSmsBarringRoaming field value if set, zero value otherwise.
func (o *SmsManagementSubscriptionData) GetMoSmsBarringRoaming() bool {
	if o == nil || isNil(o.MoSmsBarringRoaming) {
		var ret bool
		return ret
	}
	return *o.MoSmsBarringRoaming
}

// GetMoSmsBarringRoamingOk returns a tuple with the MoSmsBarringRoaming field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsManagementSubscriptionData) GetMoSmsBarringRoamingOk() (*bool, bool) {
	if o == nil || isNil(o.MoSmsBarringRoaming) {
    return nil, false
	}
	return o.MoSmsBarringRoaming, true
}

// HasMoSmsBarringRoaming returns a boolean if a field has been set.
func (o *SmsManagementSubscriptionData) HasMoSmsBarringRoaming() bool {
	if o != nil && !isNil(o.MoSmsBarringRoaming) {
		return true
	}

	return false
}

// SetMoSmsBarringRoaming gets a reference to the given bool and assigns it to the MoSmsBarringRoaming field.
func (o *SmsManagementSubscriptionData) SetMoSmsBarringRoaming(v bool) {
	o.MoSmsBarringRoaming = &v
}

// GetSharedSmsMngDataIds returns the SharedSmsMngDataIds field value if set, zero value otherwise.
func (o *SmsManagementSubscriptionData) GetSharedSmsMngDataIds() []string {
	if o == nil || isNil(o.SharedSmsMngDataIds) {
		var ret []string
		return ret
	}
	return o.SharedSmsMngDataIds
}

// GetSharedSmsMngDataIdsOk returns a tuple with the SharedSmsMngDataIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsManagementSubscriptionData) GetSharedSmsMngDataIdsOk() ([]string, bool) {
	if o == nil || isNil(o.SharedSmsMngDataIds) {
    return nil, false
	}
	return o.SharedSmsMngDataIds, true
}

// HasSharedSmsMngDataIds returns a boolean if a field has been set.
func (o *SmsManagementSubscriptionData) HasSharedSmsMngDataIds() bool {
	if o != nil && !isNil(o.SharedSmsMngDataIds) {
		return true
	}

	return false
}

// SetSharedSmsMngDataIds gets a reference to the given []string and assigns it to the SharedSmsMngDataIds field.
func (o *SmsManagementSubscriptionData) SetSharedSmsMngDataIds(v []string) {
	o.SharedSmsMngDataIds = v
}

// GetTraceData returns the TraceData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmsManagementSubscriptionData) GetTraceData() TraceData1 {
	if o == nil || isNil(o.TraceData.Get()) {
		var ret TraceData1
		return ret
	}
	return *o.TraceData.Get()
}

// GetTraceDataOk returns a tuple with the TraceData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmsManagementSubscriptionData) GetTraceDataOk() (*TraceData1, bool) {
	if o == nil {
    return nil, false
	}
	return o.TraceData.Get(), o.TraceData.IsSet()
}

// HasTraceData returns a boolean if a field has been set.
func (o *SmsManagementSubscriptionData) HasTraceData() bool {
	if o != nil && o.TraceData.IsSet() {
		return true
	}

	return false
}

// SetTraceData gets a reference to the given NullableTraceData1 and assigns it to the TraceData field.
func (o *SmsManagementSubscriptionData) SetTraceData(v TraceData1) {
	o.TraceData.Set(&v)
}
// SetTraceDataNil sets the value for TraceData to be an explicit nil
func (o *SmsManagementSubscriptionData) SetTraceDataNil() {
	o.TraceData.Set(nil)
}

// UnsetTraceData ensures that no value is present for TraceData, not even an explicit nil
func (o *SmsManagementSubscriptionData) UnsetTraceData() {
	o.TraceData.Unset()
}

func (o SmsManagementSubscriptionData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	if !isNil(o.MtSmsSubscribed) {
		toSerialize["mtSmsSubscribed"] = o.MtSmsSubscribed
	}
	if !isNil(o.MtSmsBarringAll) {
		toSerialize["mtSmsBarringAll"] = o.MtSmsBarringAll
	}
	if !isNil(o.MtSmsBarringRoaming) {
		toSerialize["mtSmsBarringRoaming"] = o.MtSmsBarringRoaming
	}
	if !isNil(o.MoSmsSubscribed) {
		toSerialize["moSmsSubscribed"] = o.MoSmsSubscribed
	}
	if !isNil(o.MoSmsBarringAll) {
		toSerialize["moSmsBarringAll"] = o.MoSmsBarringAll
	}
	if !isNil(o.MoSmsBarringRoaming) {
		toSerialize["moSmsBarringRoaming"] = o.MoSmsBarringRoaming
	}
	if !isNil(o.SharedSmsMngDataIds) {
		toSerialize["sharedSmsMngDataIds"] = o.SharedSmsMngDataIds
	}
	if o.TraceData.IsSet() {
		toSerialize["traceData"] = o.TraceData.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSmsManagementSubscriptionData struct {
	value *SmsManagementSubscriptionData
	isSet bool
}

func (v NullableSmsManagementSubscriptionData) Get() *SmsManagementSubscriptionData {
	return v.value
}

func (v *NullableSmsManagementSubscriptionData) Set(val *SmsManagementSubscriptionData) {
	v.value = val
	v.isSet = true
}

func (v NullableSmsManagementSubscriptionData) IsSet() bool {
	return v.isSet
}

func (v *NullableSmsManagementSubscriptionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmsManagementSubscriptionData(val *SmsManagementSubscriptionData) *NullableSmsManagementSubscriptionData {
	return &NullableSmsManagementSubscriptionData{value: val, isSet: true}
}

func (v NullableSmsManagementSubscriptionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmsManagementSubscriptionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

