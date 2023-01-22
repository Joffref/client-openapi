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

// EPTransportSingleAllOfAttributes struct for EPTransportSingleAllOfAttributes
type EPTransportSingleAllOfAttributes struct {
	IpAddress *IpAddress `json:"ipAddress,omitempty"`
	LogicalInterfaceInfo *LogicalInterfaceInfo `json:"logicalInterfaceInfo,omitempty"`
	NextHopInfo *string `json:"nextHopInfo,omitempty"`
	QosProfile *string `json:"qosProfile,omitempty"`
	EpApplicationRefs []string `json:"epApplicationRefs,omitempty"`
}

// NewEPTransportSingleAllOfAttributes instantiates a new EPTransportSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEPTransportSingleAllOfAttributes() *EPTransportSingleAllOfAttributes {
	this := EPTransportSingleAllOfAttributes{}
	return &this
}

// NewEPTransportSingleAllOfAttributesWithDefaults instantiates a new EPTransportSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEPTransportSingleAllOfAttributesWithDefaults() *EPTransportSingleAllOfAttributes {
	this := EPTransportSingleAllOfAttributes{}
	return &this
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *EPTransportSingleAllOfAttributes) GetIpAddress() IpAddress {
	if o == nil || isNil(o.IpAddress) {
		var ret IpAddress
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EPTransportSingleAllOfAttributes) GetIpAddressOk() (*IpAddress, bool) {
	if o == nil || isNil(o.IpAddress) {
    return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *EPTransportSingleAllOfAttributes) HasIpAddress() bool {
	if o != nil && !isNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given IpAddress and assigns it to the IpAddress field.
func (o *EPTransportSingleAllOfAttributes) SetIpAddress(v IpAddress) {
	o.IpAddress = &v
}

// GetLogicalInterfaceInfo returns the LogicalInterfaceInfo field value if set, zero value otherwise.
func (o *EPTransportSingleAllOfAttributes) GetLogicalInterfaceInfo() LogicalInterfaceInfo {
	if o == nil || isNil(o.LogicalInterfaceInfo) {
		var ret LogicalInterfaceInfo
		return ret
	}
	return *o.LogicalInterfaceInfo
}

// GetLogicalInterfaceInfoOk returns a tuple with the LogicalInterfaceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EPTransportSingleAllOfAttributes) GetLogicalInterfaceInfoOk() (*LogicalInterfaceInfo, bool) {
	if o == nil || isNil(o.LogicalInterfaceInfo) {
    return nil, false
	}
	return o.LogicalInterfaceInfo, true
}

// HasLogicalInterfaceInfo returns a boolean if a field has been set.
func (o *EPTransportSingleAllOfAttributes) HasLogicalInterfaceInfo() bool {
	if o != nil && !isNil(o.LogicalInterfaceInfo) {
		return true
	}

	return false
}

// SetLogicalInterfaceInfo gets a reference to the given LogicalInterfaceInfo and assigns it to the LogicalInterfaceInfo field.
func (o *EPTransportSingleAllOfAttributes) SetLogicalInterfaceInfo(v LogicalInterfaceInfo) {
	o.LogicalInterfaceInfo = &v
}

// GetNextHopInfo returns the NextHopInfo field value if set, zero value otherwise.
func (o *EPTransportSingleAllOfAttributes) GetNextHopInfo() string {
	if o == nil || isNil(o.NextHopInfo) {
		var ret string
		return ret
	}
	return *o.NextHopInfo
}

// GetNextHopInfoOk returns a tuple with the NextHopInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EPTransportSingleAllOfAttributes) GetNextHopInfoOk() (*string, bool) {
	if o == nil || isNil(o.NextHopInfo) {
    return nil, false
	}
	return o.NextHopInfo, true
}

// HasNextHopInfo returns a boolean if a field has been set.
func (o *EPTransportSingleAllOfAttributes) HasNextHopInfo() bool {
	if o != nil && !isNil(o.NextHopInfo) {
		return true
	}

	return false
}

// SetNextHopInfo gets a reference to the given string and assigns it to the NextHopInfo field.
func (o *EPTransportSingleAllOfAttributes) SetNextHopInfo(v string) {
	o.NextHopInfo = &v
}

// GetQosProfile returns the QosProfile field value if set, zero value otherwise.
func (o *EPTransportSingleAllOfAttributes) GetQosProfile() string {
	if o == nil || isNil(o.QosProfile) {
		var ret string
		return ret
	}
	return *o.QosProfile
}

// GetQosProfileOk returns a tuple with the QosProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EPTransportSingleAllOfAttributes) GetQosProfileOk() (*string, bool) {
	if o == nil || isNil(o.QosProfile) {
    return nil, false
	}
	return o.QosProfile, true
}

// HasQosProfile returns a boolean if a field has been set.
func (o *EPTransportSingleAllOfAttributes) HasQosProfile() bool {
	if o != nil && !isNil(o.QosProfile) {
		return true
	}

	return false
}

// SetQosProfile gets a reference to the given string and assigns it to the QosProfile field.
func (o *EPTransportSingleAllOfAttributes) SetQosProfile(v string) {
	o.QosProfile = &v
}

// GetEpApplicationRefs returns the EpApplicationRefs field value if set, zero value otherwise.
func (o *EPTransportSingleAllOfAttributes) GetEpApplicationRefs() []string {
	if o == nil || isNil(o.EpApplicationRefs) {
		var ret []string
		return ret
	}
	return o.EpApplicationRefs
}

// GetEpApplicationRefsOk returns a tuple with the EpApplicationRefs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EPTransportSingleAllOfAttributes) GetEpApplicationRefsOk() ([]string, bool) {
	if o == nil || isNil(o.EpApplicationRefs) {
    return nil, false
	}
	return o.EpApplicationRefs, true
}

// HasEpApplicationRefs returns a boolean if a field has been set.
func (o *EPTransportSingleAllOfAttributes) HasEpApplicationRefs() bool {
	if o != nil && !isNil(o.EpApplicationRefs) {
		return true
	}

	return false
}

// SetEpApplicationRefs gets a reference to the given []string and assigns it to the EpApplicationRefs field.
func (o *EPTransportSingleAllOfAttributes) SetEpApplicationRefs(v []string) {
	o.EpApplicationRefs = v
}

func (o EPTransportSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IpAddress) {
		toSerialize["ipAddress"] = o.IpAddress
	}
	if !isNil(o.LogicalInterfaceInfo) {
		toSerialize["logicalInterfaceInfo"] = o.LogicalInterfaceInfo
	}
	if !isNil(o.NextHopInfo) {
		toSerialize["nextHopInfo"] = o.NextHopInfo
	}
	if !isNil(o.QosProfile) {
		toSerialize["qosProfile"] = o.QosProfile
	}
	if !isNil(o.EpApplicationRefs) {
		toSerialize["epApplicationRefs"] = o.EpApplicationRefs
	}
	return json.Marshal(toSerialize)
}

type NullableEPTransportSingleAllOfAttributes struct {
	value *EPTransportSingleAllOfAttributes
	isSet bool
}

func (v NullableEPTransportSingleAllOfAttributes) Get() *EPTransportSingleAllOfAttributes {
	return v.value
}

func (v *NullableEPTransportSingleAllOfAttributes) Set(val *EPTransportSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableEPTransportSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableEPTransportSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEPTransportSingleAllOfAttributes(val *EPTransportSingleAllOfAttributes) *NullableEPTransportSingleAllOfAttributes {
	return &NullableEPTransportSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableEPTransportSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEPTransportSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

