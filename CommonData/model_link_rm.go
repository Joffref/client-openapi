/*
Common Data Types

Common Data Types for Service Based Interfaces.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.   

API version: 1.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package CommonData

import (
	"encoding/json"
)

// LinkRm It contains the URI of the linked resource with the OpenAPI 'nullable: true' property. 
type LinkRm struct {
	// String providing an URI formatted according to RFC 3986.
	Href *string `json:"href,omitempty"`
}

// NewLinkRm instantiates a new LinkRm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkRm() *LinkRm {
	this := LinkRm{}
	return &this
}

// NewLinkRmWithDefaults instantiates a new LinkRm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkRmWithDefaults() *LinkRm {
	this := LinkRm{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *LinkRm) GetHref() string {
	if o == nil || isNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkRm) GetHrefOk() (*string, bool) {
	if o == nil || isNil(o.Href) {
    return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *LinkRm) HasHref() bool {
	if o != nil && !isNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *LinkRm) SetHref(v string) {
	o.Href = &v
}

func (o LinkRm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	return json.Marshal(toSerialize)
}

type NullableLinkRm struct {
	value *LinkRm
	isSet bool
}

func (v NullableLinkRm) Get() *LinkRm {
	return v.value
}

func (v *NullableLinkRm) Set(val *LinkRm) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkRm) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkRm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkRm(val *LinkRm) *NullableLinkRm {
	return &NullableLinkRm{value: val, isSet: true}
}

func (v NullableLinkRm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkRm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

