/*
Unified Data Repository Service API file for Application Data

The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Application_Data

import (
	"encoding/json"
)

// AmInfluDataPatch Represents the AM Influence Data that can be updated.
type AmInfluDataPatch struct {
	// Identifies one or more applications.
	AppIds []string `json:"appIds,omitempty"`
	// Identifies one or more DNN, S-NSSAI combinations.
	DnnSnssaiInfos []DnnSnssaiInformation `json:"dnnSnssaiInfos,omitempty"`
	// List of AM related events for which a subscription is required.
	EvSubs []AmInfluEvent `json:"evSubs,omitempty"`
	// Contains the headers provisioned by the NEF.
	Headers []string `json:"headers,omitempty"`
	// Indicates whether high throughput is desired for the indicated UE traffic.
	ThruReq NullableBool `json:"thruReq,omitempty"`
	// String providing an URI formatted according to RFC 3986 with the OpenAPI 'nullable: true' property. 
	NotifUri NullableString `json:"notifUri,omitempty"`
	// Notification correlation identifier.
	NotifCorrId NullableString `json:"notifCorrId,omitempty"`
	// Indicates the service area coverage requirement.
	CovReq []ServiceAreaCoverageInfo `json:"covReq,omitempty"`
}

// NewAmInfluDataPatch instantiates a new AmInfluDataPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmInfluDataPatch() *AmInfluDataPatch {
	this := AmInfluDataPatch{}
	return &this
}

// NewAmInfluDataPatchWithDefaults instantiates a new AmInfluDataPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmInfluDataPatchWithDefaults() *AmInfluDataPatch {
	this := AmInfluDataPatch{}
	return &this
}

// GetAppIds returns the AppIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AmInfluDataPatch) GetAppIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AppIds
}

// GetAppIdsOk returns a tuple with the AppIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AmInfluDataPatch) GetAppIdsOk() ([]string, bool) {
	if o == nil || isNil(o.AppIds) {
    return nil, false
	}
	return o.AppIds, true
}

// HasAppIds returns a boolean if a field has been set.
func (o *AmInfluDataPatch) HasAppIds() bool {
	if o != nil && isNil(o.AppIds) {
		return true
	}

	return false
}

// SetAppIds gets a reference to the given []string and assigns it to the AppIds field.
func (o *AmInfluDataPatch) SetAppIds(v []string) {
	o.AppIds = v
}

// GetDnnSnssaiInfos returns the DnnSnssaiInfos field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AmInfluDataPatch) GetDnnSnssaiInfos() []DnnSnssaiInformation {
	if o == nil {
		var ret []DnnSnssaiInformation
		return ret
	}
	return o.DnnSnssaiInfos
}

// GetDnnSnssaiInfosOk returns a tuple with the DnnSnssaiInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AmInfluDataPatch) GetDnnSnssaiInfosOk() ([]DnnSnssaiInformation, bool) {
	if o == nil || isNil(o.DnnSnssaiInfos) {
    return nil, false
	}
	return o.DnnSnssaiInfos, true
}

// HasDnnSnssaiInfos returns a boolean if a field has been set.
func (o *AmInfluDataPatch) HasDnnSnssaiInfos() bool {
	if o != nil && isNil(o.DnnSnssaiInfos) {
		return true
	}

	return false
}

// SetDnnSnssaiInfos gets a reference to the given []DnnSnssaiInformation and assigns it to the DnnSnssaiInfos field.
func (o *AmInfluDataPatch) SetDnnSnssaiInfos(v []DnnSnssaiInformation) {
	o.DnnSnssaiInfos = v
}

// GetEvSubs returns the EvSubs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AmInfluDataPatch) GetEvSubs() []AmInfluEvent {
	if o == nil {
		var ret []AmInfluEvent
		return ret
	}
	return o.EvSubs
}

// GetEvSubsOk returns a tuple with the EvSubs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AmInfluDataPatch) GetEvSubsOk() ([]AmInfluEvent, bool) {
	if o == nil || isNil(o.EvSubs) {
    return nil, false
	}
	return o.EvSubs, true
}

// HasEvSubs returns a boolean if a field has been set.
func (o *AmInfluDataPatch) HasEvSubs() bool {
	if o != nil && isNil(o.EvSubs) {
		return true
	}

	return false
}

// SetEvSubs gets a reference to the given []AmInfluEvent and assigns it to the EvSubs field.
func (o *AmInfluDataPatch) SetEvSubs(v []AmInfluEvent) {
	o.EvSubs = v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *AmInfluDataPatch) GetHeaders() []string {
	if o == nil || isNil(o.Headers) {
		var ret []string
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmInfluDataPatch) GetHeadersOk() ([]string, bool) {
	if o == nil || isNil(o.Headers) {
    return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *AmInfluDataPatch) HasHeaders() bool {
	if o != nil && !isNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []string and assigns it to the Headers field.
func (o *AmInfluDataPatch) SetHeaders(v []string) {
	o.Headers = v
}

// GetThruReq returns the ThruReq field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AmInfluDataPatch) GetThruReq() bool {
	if o == nil || isNil(o.ThruReq.Get()) {
		var ret bool
		return ret
	}
	return *o.ThruReq.Get()
}

// GetThruReqOk returns a tuple with the ThruReq field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AmInfluDataPatch) GetThruReqOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.ThruReq.Get(), o.ThruReq.IsSet()
}

// HasThruReq returns a boolean if a field has been set.
func (o *AmInfluDataPatch) HasThruReq() bool {
	if o != nil && o.ThruReq.IsSet() {
		return true
	}

	return false
}

// SetThruReq gets a reference to the given NullableBool and assigns it to the ThruReq field.
func (o *AmInfluDataPatch) SetThruReq(v bool) {
	o.ThruReq.Set(&v)
}
// SetThruReqNil sets the value for ThruReq to be an explicit nil
func (o *AmInfluDataPatch) SetThruReqNil() {
	o.ThruReq.Set(nil)
}

// UnsetThruReq ensures that no value is present for ThruReq, not even an explicit nil
func (o *AmInfluDataPatch) UnsetThruReq() {
	o.ThruReq.Unset()
}

// GetNotifUri returns the NotifUri field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AmInfluDataPatch) GetNotifUri() string {
	if o == nil || isNil(o.NotifUri.Get()) {
		var ret string
		return ret
	}
	return *o.NotifUri.Get()
}

// GetNotifUriOk returns a tuple with the NotifUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AmInfluDataPatch) GetNotifUriOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.NotifUri.Get(), o.NotifUri.IsSet()
}

// HasNotifUri returns a boolean if a field has been set.
func (o *AmInfluDataPatch) HasNotifUri() bool {
	if o != nil && o.NotifUri.IsSet() {
		return true
	}

	return false
}

// SetNotifUri gets a reference to the given NullableString and assigns it to the NotifUri field.
func (o *AmInfluDataPatch) SetNotifUri(v string) {
	o.NotifUri.Set(&v)
}
// SetNotifUriNil sets the value for NotifUri to be an explicit nil
func (o *AmInfluDataPatch) SetNotifUriNil() {
	o.NotifUri.Set(nil)
}

// UnsetNotifUri ensures that no value is present for NotifUri, not even an explicit nil
func (o *AmInfluDataPatch) UnsetNotifUri() {
	o.NotifUri.Unset()
}

// GetNotifCorrId returns the NotifCorrId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AmInfluDataPatch) GetNotifCorrId() string {
	if o == nil || isNil(o.NotifCorrId.Get()) {
		var ret string
		return ret
	}
	return *o.NotifCorrId.Get()
}

// GetNotifCorrIdOk returns a tuple with the NotifCorrId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AmInfluDataPatch) GetNotifCorrIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.NotifCorrId.Get(), o.NotifCorrId.IsSet()
}

// HasNotifCorrId returns a boolean if a field has been set.
func (o *AmInfluDataPatch) HasNotifCorrId() bool {
	if o != nil && o.NotifCorrId.IsSet() {
		return true
	}

	return false
}

// SetNotifCorrId gets a reference to the given NullableString and assigns it to the NotifCorrId field.
func (o *AmInfluDataPatch) SetNotifCorrId(v string) {
	o.NotifCorrId.Set(&v)
}
// SetNotifCorrIdNil sets the value for NotifCorrId to be an explicit nil
func (o *AmInfluDataPatch) SetNotifCorrIdNil() {
	o.NotifCorrId.Set(nil)
}

// UnsetNotifCorrId ensures that no value is present for NotifCorrId, not even an explicit nil
func (o *AmInfluDataPatch) UnsetNotifCorrId() {
	o.NotifCorrId.Unset()
}

// GetCovReq returns the CovReq field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AmInfluDataPatch) GetCovReq() []ServiceAreaCoverageInfo {
	if o == nil {
		var ret []ServiceAreaCoverageInfo
		return ret
	}
	return o.CovReq
}

// GetCovReqOk returns a tuple with the CovReq field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AmInfluDataPatch) GetCovReqOk() ([]ServiceAreaCoverageInfo, bool) {
	if o == nil || isNil(o.CovReq) {
    return nil, false
	}
	return o.CovReq, true
}

// HasCovReq returns a boolean if a field has been set.
func (o *AmInfluDataPatch) HasCovReq() bool {
	if o != nil && isNil(o.CovReq) {
		return true
	}

	return false
}

// SetCovReq gets a reference to the given []ServiceAreaCoverageInfo and assigns it to the CovReq field.
func (o *AmInfluDataPatch) SetCovReq(v []ServiceAreaCoverageInfo) {
	o.CovReq = v
}

func (o AmInfluDataPatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AppIds != nil {
		toSerialize["appIds"] = o.AppIds
	}
	if o.DnnSnssaiInfos != nil {
		toSerialize["dnnSnssaiInfos"] = o.DnnSnssaiInfos
	}
	if o.EvSubs != nil {
		toSerialize["evSubs"] = o.EvSubs
	}
	if !isNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	if o.ThruReq.IsSet() {
		toSerialize["thruReq"] = o.ThruReq.Get()
	}
	if o.NotifUri.IsSet() {
		toSerialize["notifUri"] = o.NotifUri.Get()
	}
	if o.NotifCorrId.IsSet() {
		toSerialize["notifCorrId"] = o.NotifCorrId.Get()
	}
	if o.CovReq != nil {
		toSerialize["covReq"] = o.CovReq
	}
	return json.Marshal(toSerialize)
}

type NullableAmInfluDataPatch struct {
	value *AmInfluDataPatch
	isSet bool
}

func (v NullableAmInfluDataPatch) Get() *AmInfluDataPatch {
	return v.value
}

func (v *NullableAmInfluDataPatch) Set(val *AmInfluDataPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableAmInfluDataPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableAmInfluDataPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmInfluDataPatch(val *AmInfluDataPatch) *NullableAmInfluDataPatch {
	return &NullableAmInfluDataPatch{value: val, isSet: true}
}

func (v NullableAmInfluDataPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmInfluDataPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

