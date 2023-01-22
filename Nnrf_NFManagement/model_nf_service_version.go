/*
NRF NFManagement Service

NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnrf_NFManagement

import (
	"encoding/json"
	"time"
)

// NFServiceVersion Contains the version details of an NF service
type NFServiceVersion struct {
	ApiVersionInUri string `json:"apiVersionInUri"`
	ApiFullVersion string `json:"apiFullVersion"`
	// string with format 'date-time' as defined in OpenAPI.
	Expiry *time.Time `json:"expiry,omitempty"`
}

// NewNFServiceVersion instantiates a new NFServiceVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNFServiceVersion(apiVersionInUri string, apiFullVersion string) *NFServiceVersion {
	this := NFServiceVersion{}
	this.ApiVersionInUri = apiVersionInUri
	this.ApiFullVersion = apiFullVersion
	return &this
}

// NewNFServiceVersionWithDefaults instantiates a new NFServiceVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNFServiceVersionWithDefaults() *NFServiceVersion {
	this := NFServiceVersion{}
	return &this
}

// GetApiVersionInUri returns the ApiVersionInUri field value
func (o *NFServiceVersion) GetApiVersionInUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersionInUri
}

// GetApiVersionInUriOk returns a tuple with the ApiVersionInUri field value
// and a boolean to check if the value has been set.
func (o *NFServiceVersion) GetApiVersionInUriOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ApiVersionInUri, true
}

// SetApiVersionInUri sets field value
func (o *NFServiceVersion) SetApiVersionInUri(v string) {
	o.ApiVersionInUri = v
}

// GetApiFullVersion returns the ApiFullVersion field value
func (o *NFServiceVersion) GetApiFullVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiFullVersion
}

// GetApiFullVersionOk returns a tuple with the ApiFullVersion field value
// and a boolean to check if the value has been set.
func (o *NFServiceVersion) GetApiFullVersionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ApiFullVersion, true
}

// SetApiFullVersion sets field value
func (o *NFServiceVersion) SetApiFullVersion(v string) {
	o.ApiFullVersion = v
}

// GetExpiry returns the Expiry field value if set, zero value otherwise.
func (o *NFServiceVersion) GetExpiry() time.Time {
	if o == nil || isNil(o.Expiry) {
		var ret time.Time
		return ret
	}
	return *o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFServiceVersion) GetExpiryOk() (*time.Time, bool) {
	if o == nil || isNil(o.Expiry) {
    return nil, false
	}
	return o.Expiry, true
}

// HasExpiry returns a boolean if a field has been set.
func (o *NFServiceVersion) HasExpiry() bool {
	if o != nil && !isNil(o.Expiry) {
		return true
	}

	return false
}

// SetExpiry gets a reference to the given time.Time and assigns it to the Expiry field.
func (o *NFServiceVersion) SetExpiry(v time.Time) {
	o.Expiry = &v
}

func (o NFServiceVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["apiVersionInUri"] = o.ApiVersionInUri
	}
	if true {
		toSerialize["apiFullVersion"] = o.ApiFullVersion
	}
	if !isNil(o.Expiry) {
		toSerialize["expiry"] = o.Expiry
	}
	return json.Marshal(toSerialize)
}

type NullableNFServiceVersion struct {
	value *NFServiceVersion
	isSet bool
}

func (v NullableNFServiceVersion) Get() *NFServiceVersion {
	return v.value
}

func (v *NullableNFServiceVersion) Set(val *NFServiceVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableNFServiceVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableNFServiceVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNFServiceVersion(val *NFServiceVersion) *NullableNFServiceVersion {
	return &NullableNFServiceVersion{value: val, isSet: true}
}

func (v NullableNFServiceVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNFServiceVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

