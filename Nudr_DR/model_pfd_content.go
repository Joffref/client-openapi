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

// PfdContent Represents the content of a PFD for an application identifier.
type PfdContent struct {
	// Identifies a PDF of an application identifier.
	PfdId *string `json:"pfdId,omitempty"`
	// Represents a 3-tuple with protocol, server ip and server port for UL/DL application traffic. 
	FlowDescriptions []string `json:"flowDescriptions,omitempty"`
	// Indicates a URL or a regular expression which is used to match the significant parts of the URL. 
	Urls []string `json:"urls,omitempty"`
	// Indicates an FQDN or a regular expression as a domain name matching criteria.
	DomainNames []string `json:"domainNames,omitempty"`
	DnProtocol *DomainNameProtocol `json:"dnProtocol,omitempty"`
}

// NewPfdContent instantiates a new PfdContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPfdContent() *PfdContent {
	this := PfdContent{}
	return &this
}

// NewPfdContentWithDefaults instantiates a new PfdContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPfdContentWithDefaults() *PfdContent {
	this := PfdContent{}
	return &this
}

// GetPfdId returns the PfdId field value if set, zero value otherwise.
func (o *PfdContent) GetPfdId() string {
	if o == nil || isNil(o.PfdId) {
		var ret string
		return ret
	}
	return *o.PfdId
}

// GetPfdIdOk returns a tuple with the PfdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PfdContent) GetPfdIdOk() (*string, bool) {
	if o == nil || isNil(o.PfdId) {
    return nil, false
	}
	return o.PfdId, true
}

// HasPfdId returns a boolean if a field has been set.
func (o *PfdContent) HasPfdId() bool {
	if o != nil && !isNil(o.PfdId) {
		return true
	}

	return false
}

// SetPfdId gets a reference to the given string and assigns it to the PfdId field.
func (o *PfdContent) SetPfdId(v string) {
	o.PfdId = &v
}

// GetFlowDescriptions returns the FlowDescriptions field value if set, zero value otherwise.
func (o *PfdContent) GetFlowDescriptions() []string {
	if o == nil || isNil(o.FlowDescriptions) {
		var ret []string
		return ret
	}
	return o.FlowDescriptions
}

// GetFlowDescriptionsOk returns a tuple with the FlowDescriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PfdContent) GetFlowDescriptionsOk() ([]string, bool) {
	if o == nil || isNil(o.FlowDescriptions) {
    return nil, false
	}
	return o.FlowDescriptions, true
}

// HasFlowDescriptions returns a boolean if a field has been set.
func (o *PfdContent) HasFlowDescriptions() bool {
	if o != nil && !isNil(o.FlowDescriptions) {
		return true
	}

	return false
}

// SetFlowDescriptions gets a reference to the given []string and assigns it to the FlowDescriptions field.
func (o *PfdContent) SetFlowDescriptions(v []string) {
	o.FlowDescriptions = v
}

// GetUrls returns the Urls field value if set, zero value otherwise.
func (o *PfdContent) GetUrls() []string {
	if o == nil || isNil(o.Urls) {
		var ret []string
		return ret
	}
	return o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PfdContent) GetUrlsOk() ([]string, bool) {
	if o == nil || isNil(o.Urls) {
    return nil, false
	}
	return o.Urls, true
}

// HasUrls returns a boolean if a field has been set.
func (o *PfdContent) HasUrls() bool {
	if o != nil && !isNil(o.Urls) {
		return true
	}

	return false
}

// SetUrls gets a reference to the given []string and assigns it to the Urls field.
func (o *PfdContent) SetUrls(v []string) {
	o.Urls = v
}

// GetDomainNames returns the DomainNames field value if set, zero value otherwise.
func (o *PfdContent) GetDomainNames() []string {
	if o == nil || isNil(o.DomainNames) {
		var ret []string
		return ret
	}
	return o.DomainNames
}

// GetDomainNamesOk returns a tuple with the DomainNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PfdContent) GetDomainNamesOk() ([]string, bool) {
	if o == nil || isNil(o.DomainNames) {
    return nil, false
	}
	return o.DomainNames, true
}

// HasDomainNames returns a boolean if a field has been set.
func (o *PfdContent) HasDomainNames() bool {
	if o != nil && !isNil(o.DomainNames) {
		return true
	}

	return false
}

// SetDomainNames gets a reference to the given []string and assigns it to the DomainNames field.
func (o *PfdContent) SetDomainNames(v []string) {
	o.DomainNames = v
}

// GetDnProtocol returns the DnProtocol field value if set, zero value otherwise.
func (o *PfdContent) GetDnProtocol() DomainNameProtocol {
	if o == nil || isNil(o.DnProtocol) {
		var ret DomainNameProtocol
		return ret
	}
	return *o.DnProtocol
}

// GetDnProtocolOk returns a tuple with the DnProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PfdContent) GetDnProtocolOk() (*DomainNameProtocol, bool) {
	if o == nil || isNil(o.DnProtocol) {
    return nil, false
	}
	return o.DnProtocol, true
}

// HasDnProtocol returns a boolean if a field has been set.
func (o *PfdContent) HasDnProtocol() bool {
	if o != nil && !isNil(o.DnProtocol) {
		return true
	}

	return false
}

// SetDnProtocol gets a reference to the given DomainNameProtocol and assigns it to the DnProtocol field.
func (o *PfdContent) SetDnProtocol(v DomainNameProtocol) {
	o.DnProtocol = &v
}

func (o PfdContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PfdId) {
		toSerialize["pfdId"] = o.PfdId
	}
	if !isNil(o.FlowDescriptions) {
		toSerialize["flowDescriptions"] = o.FlowDescriptions
	}
	if !isNil(o.Urls) {
		toSerialize["urls"] = o.Urls
	}
	if !isNil(o.DomainNames) {
		toSerialize["domainNames"] = o.DomainNames
	}
	if !isNil(o.DnProtocol) {
		toSerialize["dnProtocol"] = o.DnProtocol
	}
	return json.Marshal(toSerialize)
}

type NullablePfdContent struct {
	value *PfdContent
	isSet bool
}

func (v NullablePfdContent) Get() *PfdContent {
	return v.value
}

func (v *NullablePfdContent) Set(val *PfdContent) {
	v.value = val
	v.isSet = true
}

func (v NullablePfdContent) IsSet() bool {
	return v.isSet
}

func (v *NullablePfdContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePfdContent(val *PfdContent) *NullablePfdContent {
	return &NullablePfdContent{value: val, isSet: true}
}

func (v NullablePfdContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePfdContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

