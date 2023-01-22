/*
Slice NRM

OAS 3.0.1 specification of the Slice NRM @ 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package SliceNrm

import (
	"encoding/json"
)

// VsDataContainerSingle struct for VsDataContainerSingle
type VsDataContainerSingle struct {
	Id *string `json:"id,omitempty"`
	Attributes *VsDataContainerSingleAttributes `json:"attributes,omitempty"`
	VsDataContainer []VsDataContainerSingle `json:"VsDataContainer,omitempty"`
}

// NewVsDataContainerSingle instantiates a new VsDataContainerSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVsDataContainerSingle() *VsDataContainerSingle {
	this := VsDataContainerSingle{}
	return &this
}

// NewVsDataContainerSingleWithDefaults instantiates a new VsDataContainerSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVsDataContainerSingleWithDefaults() *VsDataContainerSingle {
	this := VsDataContainerSingle{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VsDataContainerSingle) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VsDataContainerSingle) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VsDataContainerSingle) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VsDataContainerSingle) SetId(v string) {
	o.Id = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *VsDataContainerSingle) GetAttributes() VsDataContainerSingleAttributes {
	if o == nil || isNil(o.Attributes) {
		var ret VsDataContainerSingleAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VsDataContainerSingle) GetAttributesOk() (*VsDataContainerSingleAttributes, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *VsDataContainerSingle) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given VsDataContainerSingleAttributes and assigns it to the Attributes field.
func (o *VsDataContainerSingle) SetAttributes(v VsDataContainerSingleAttributes) {
	o.Attributes = &v
}

// GetVsDataContainer returns the VsDataContainer field value if set, zero value otherwise.
func (o *VsDataContainerSingle) GetVsDataContainer() []VsDataContainerSingle {
	if o == nil || isNil(o.VsDataContainer) {
		var ret []VsDataContainerSingle
		return ret
	}
	return o.VsDataContainer
}

// GetVsDataContainerOk returns a tuple with the VsDataContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VsDataContainerSingle) GetVsDataContainerOk() ([]VsDataContainerSingle, bool) {
	if o == nil || isNil(o.VsDataContainer) {
    return nil, false
	}
	return o.VsDataContainer, true
}

// HasVsDataContainer returns a boolean if a field has been set.
func (o *VsDataContainerSingle) HasVsDataContainer() bool {
	if o != nil && !isNil(o.VsDataContainer) {
		return true
	}

	return false
}

// SetVsDataContainer gets a reference to the given []VsDataContainerSingle and assigns it to the VsDataContainer field.
func (o *VsDataContainerSingle) SetVsDataContainer(v []VsDataContainerSingle) {
	o.VsDataContainer = v
}

func (o VsDataContainerSingle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !isNil(o.VsDataContainer) {
		toSerialize["VsDataContainer"] = o.VsDataContainer
	}
	return json.Marshal(toSerialize)
}

type NullableVsDataContainerSingle struct {
	value *VsDataContainerSingle
	isSet bool
}

func (v NullableVsDataContainerSingle) Get() *VsDataContainerSingle {
	return v.value
}

func (v *NullableVsDataContainerSingle) Set(val *VsDataContainerSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableVsDataContainerSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableVsDataContainerSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVsDataContainerSingle(val *VsDataContainerSingle) *NullableVsDataContainerSingle {
	return &NullableVsDataContainerSingle{value: val, isSet: true}
}

func (v NullableVsDataContainerSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVsDataContainerSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

