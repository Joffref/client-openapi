/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nchf_ConvergedCharging

import (
	"encoding/json"
	"time"
)

// AccessTransferInformation struct for AccessTransferInformation
type AccessTransferInformation struct {
	AccessTransferType *AccessTransferType `json:"accessTransferType,omitempty"`
	AccessNetworkInformation []string `json:"accessNetworkInformation,omitempty"`
	CellularNetworkInformation *string `json:"cellularNetworkInformation,omitempty"`
	InterUETransfer *UETransferType `json:"interUETransfer,omitempty"`
	// String representing a Permanent Equipment Identifier that may contain - an IMEI or IMEISV, as  specified in clause 6.2 of 3GPP TS 23.003; a MAC address for a 5G-RG or FN-RG via  wireline  access, with an indication that this address cannot be trusted for regulatory purpose if this  address cannot be used as an Equipment Identifier of the FN-RG, as specified in clause 4.7.7  of 3GPP TS23.316. Examples are imei-012345678901234 or imeisv-0123456789012345.  
	UserEquipmentInfo *string `json:"userEquipmentInfo,omitempty"`
	InstanceId *string `json:"instanceId,omitempty"`
	RelatedIMSChargingIdentifier *string `json:"relatedIMSChargingIdentifier,omitempty"`
	RelatedIMSChargingIdentifierNode *IMSAddress `json:"relatedIMSChargingIdentifierNode,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	ChangeTime *time.Time `json:"changeTime,omitempty"`
}

// NewAccessTransferInformation instantiates a new AccessTransferInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessTransferInformation() *AccessTransferInformation {
	this := AccessTransferInformation{}
	return &this
}

// NewAccessTransferInformationWithDefaults instantiates a new AccessTransferInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessTransferInformationWithDefaults() *AccessTransferInformation {
	this := AccessTransferInformation{}
	return &this
}

// GetAccessTransferType returns the AccessTransferType field value if set, zero value otherwise.
func (o *AccessTransferInformation) GetAccessTransferType() AccessTransferType {
	if o == nil || isNil(o.AccessTransferType) {
		var ret AccessTransferType
		return ret
	}
	return *o.AccessTransferType
}

// GetAccessTransferTypeOk returns a tuple with the AccessTransferType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTransferInformation) GetAccessTransferTypeOk() (*AccessTransferType, bool) {
	if o == nil || isNil(o.AccessTransferType) {
    return nil, false
	}
	return o.AccessTransferType, true
}

// HasAccessTransferType returns a boolean if a field has been set.
func (o *AccessTransferInformation) HasAccessTransferType() bool {
	if o != nil && !isNil(o.AccessTransferType) {
		return true
	}

	return false
}

// SetAccessTransferType gets a reference to the given AccessTransferType and assigns it to the AccessTransferType field.
func (o *AccessTransferInformation) SetAccessTransferType(v AccessTransferType) {
	o.AccessTransferType = &v
}

// GetAccessNetworkInformation returns the AccessNetworkInformation field value if set, zero value otherwise.
func (o *AccessTransferInformation) GetAccessNetworkInformation() []string {
	if o == nil || isNil(o.AccessNetworkInformation) {
		var ret []string
		return ret
	}
	return o.AccessNetworkInformation
}

// GetAccessNetworkInformationOk returns a tuple with the AccessNetworkInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTransferInformation) GetAccessNetworkInformationOk() ([]string, bool) {
	if o == nil || isNil(o.AccessNetworkInformation) {
    return nil, false
	}
	return o.AccessNetworkInformation, true
}

// HasAccessNetworkInformation returns a boolean if a field has been set.
func (o *AccessTransferInformation) HasAccessNetworkInformation() bool {
	if o != nil && !isNil(o.AccessNetworkInformation) {
		return true
	}

	return false
}

// SetAccessNetworkInformation gets a reference to the given []string and assigns it to the AccessNetworkInformation field.
func (o *AccessTransferInformation) SetAccessNetworkInformation(v []string) {
	o.AccessNetworkInformation = v
}

// GetCellularNetworkInformation returns the CellularNetworkInformation field value if set, zero value otherwise.
func (o *AccessTransferInformation) GetCellularNetworkInformation() string {
	if o == nil || isNil(o.CellularNetworkInformation) {
		var ret string
		return ret
	}
	return *o.CellularNetworkInformation
}

// GetCellularNetworkInformationOk returns a tuple with the CellularNetworkInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTransferInformation) GetCellularNetworkInformationOk() (*string, bool) {
	if o == nil || isNil(o.CellularNetworkInformation) {
    return nil, false
	}
	return o.CellularNetworkInformation, true
}

// HasCellularNetworkInformation returns a boolean if a field has been set.
func (o *AccessTransferInformation) HasCellularNetworkInformation() bool {
	if o != nil && !isNil(o.CellularNetworkInformation) {
		return true
	}

	return false
}

// SetCellularNetworkInformation gets a reference to the given string and assigns it to the CellularNetworkInformation field.
func (o *AccessTransferInformation) SetCellularNetworkInformation(v string) {
	o.CellularNetworkInformation = &v
}

// GetInterUETransfer returns the InterUETransfer field value if set, zero value otherwise.
func (o *AccessTransferInformation) GetInterUETransfer() UETransferType {
	if o == nil || isNil(o.InterUETransfer) {
		var ret UETransferType
		return ret
	}
	return *o.InterUETransfer
}

// GetInterUETransferOk returns a tuple with the InterUETransfer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTransferInformation) GetInterUETransferOk() (*UETransferType, bool) {
	if o == nil || isNil(o.InterUETransfer) {
    return nil, false
	}
	return o.InterUETransfer, true
}

// HasInterUETransfer returns a boolean if a field has been set.
func (o *AccessTransferInformation) HasInterUETransfer() bool {
	if o != nil && !isNil(o.InterUETransfer) {
		return true
	}

	return false
}

// SetInterUETransfer gets a reference to the given UETransferType and assigns it to the InterUETransfer field.
func (o *AccessTransferInformation) SetInterUETransfer(v UETransferType) {
	o.InterUETransfer = &v
}

// GetUserEquipmentInfo returns the UserEquipmentInfo field value if set, zero value otherwise.
func (o *AccessTransferInformation) GetUserEquipmentInfo() string {
	if o == nil || isNil(o.UserEquipmentInfo) {
		var ret string
		return ret
	}
	return *o.UserEquipmentInfo
}

// GetUserEquipmentInfoOk returns a tuple with the UserEquipmentInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTransferInformation) GetUserEquipmentInfoOk() (*string, bool) {
	if o == nil || isNil(o.UserEquipmentInfo) {
    return nil, false
	}
	return o.UserEquipmentInfo, true
}

// HasUserEquipmentInfo returns a boolean if a field has been set.
func (o *AccessTransferInformation) HasUserEquipmentInfo() bool {
	if o != nil && !isNil(o.UserEquipmentInfo) {
		return true
	}

	return false
}

// SetUserEquipmentInfo gets a reference to the given string and assigns it to the UserEquipmentInfo field.
func (o *AccessTransferInformation) SetUserEquipmentInfo(v string) {
	o.UserEquipmentInfo = &v
}

// GetInstanceId returns the InstanceId field value if set, zero value otherwise.
func (o *AccessTransferInformation) GetInstanceId() string {
	if o == nil || isNil(o.InstanceId) {
		var ret string
		return ret
	}
	return *o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTransferInformation) GetInstanceIdOk() (*string, bool) {
	if o == nil || isNil(o.InstanceId) {
    return nil, false
	}
	return o.InstanceId, true
}

// HasInstanceId returns a boolean if a field has been set.
func (o *AccessTransferInformation) HasInstanceId() bool {
	if o != nil && !isNil(o.InstanceId) {
		return true
	}

	return false
}

// SetInstanceId gets a reference to the given string and assigns it to the InstanceId field.
func (o *AccessTransferInformation) SetInstanceId(v string) {
	o.InstanceId = &v
}

// GetRelatedIMSChargingIdentifier returns the RelatedIMSChargingIdentifier field value if set, zero value otherwise.
func (o *AccessTransferInformation) GetRelatedIMSChargingIdentifier() string {
	if o == nil || isNil(o.RelatedIMSChargingIdentifier) {
		var ret string
		return ret
	}
	return *o.RelatedIMSChargingIdentifier
}

// GetRelatedIMSChargingIdentifierOk returns a tuple with the RelatedIMSChargingIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTransferInformation) GetRelatedIMSChargingIdentifierOk() (*string, bool) {
	if o == nil || isNil(o.RelatedIMSChargingIdentifier) {
    return nil, false
	}
	return o.RelatedIMSChargingIdentifier, true
}

// HasRelatedIMSChargingIdentifier returns a boolean if a field has been set.
func (o *AccessTransferInformation) HasRelatedIMSChargingIdentifier() bool {
	if o != nil && !isNil(o.RelatedIMSChargingIdentifier) {
		return true
	}

	return false
}

// SetRelatedIMSChargingIdentifier gets a reference to the given string and assigns it to the RelatedIMSChargingIdentifier field.
func (o *AccessTransferInformation) SetRelatedIMSChargingIdentifier(v string) {
	o.RelatedIMSChargingIdentifier = &v
}

// GetRelatedIMSChargingIdentifierNode returns the RelatedIMSChargingIdentifierNode field value if set, zero value otherwise.
func (o *AccessTransferInformation) GetRelatedIMSChargingIdentifierNode() IMSAddress {
	if o == nil || isNil(o.RelatedIMSChargingIdentifierNode) {
		var ret IMSAddress
		return ret
	}
	return *o.RelatedIMSChargingIdentifierNode
}

// GetRelatedIMSChargingIdentifierNodeOk returns a tuple with the RelatedIMSChargingIdentifierNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTransferInformation) GetRelatedIMSChargingIdentifierNodeOk() (*IMSAddress, bool) {
	if o == nil || isNil(o.RelatedIMSChargingIdentifierNode) {
    return nil, false
	}
	return o.RelatedIMSChargingIdentifierNode, true
}

// HasRelatedIMSChargingIdentifierNode returns a boolean if a field has been set.
func (o *AccessTransferInformation) HasRelatedIMSChargingIdentifierNode() bool {
	if o != nil && !isNil(o.RelatedIMSChargingIdentifierNode) {
		return true
	}

	return false
}

// SetRelatedIMSChargingIdentifierNode gets a reference to the given IMSAddress and assigns it to the RelatedIMSChargingIdentifierNode field.
func (o *AccessTransferInformation) SetRelatedIMSChargingIdentifierNode(v IMSAddress) {
	o.RelatedIMSChargingIdentifierNode = &v
}

// GetChangeTime returns the ChangeTime field value if set, zero value otherwise.
func (o *AccessTransferInformation) GetChangeTime() time.Time {
	if o == nil || isNil(o.ChangeTime) {
		var ret time.Time
		return ret
	}
	return *o.ChangeTime
}

// GetChangeTimeOk returns a tuple with the ChangeTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTransferInformation) GetChangeTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.ChangeTime) {
    return nil, false
	}
	return o.ChangeTime, true
}

// HasChangeTime returns a boolean if a field has been set.
func (o *AccessTransferInformation) HasChangeTime() bool {
	if o != nil && !isNil(o.ChangeTime) {
		return true
	}

	return false
}

// SetChangeTime gets a reference to the given time.Time and assigns it to the ChangeTime field.
func (o *AccessTransferInformation) SetChangeTime(v time.Time) {
	o.ChangeTime = &v
}

func (o AccessTransferInformation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AccessTransferType) {
		toSerialize["accessTransferType"] = o.AccessTransferType
	}
	if !isNil(o.AccessNetworkInformation) {
		toSerialize["accessNetworkInformation"] = o.AccessNetworkInformation
	}
	if !isNil(o.CellularNetworkInformation) {
		toSerialize["cellularNetworkInformation"] = o.CellularNetworkInformation
	}
	if !isNil(o.InterUETransfer) {
		toSerialize["interUETransfer"] = o.InterUETransfer
	}
	if !isNil(o.UserEquipmentInfo) {
		toSerialize["userEquipmentInfo"] = o.UserEquipmentInfo
	}
	if !isNil(o.InstanceId) {
		toSerialize["instanceId"] = o.InstanceId
	}
	if !isNil(o.RelatedIMSChargingIdentifier) {
		toSerialize["relatedIMSChargingIdentifier"] = o.RelatedIMSChargingIdentifier
	}
	if !isNil(o.RelatedIMSChargingIdentifierNode) {
		toSerialize["relatedIMSChargingIdentifierNode"] = o.RelatedIMSChargingIdentifierNode
	}
	if !isNil(o.ChangeTime) {
		toSerialize["changeTime"] = o.ChangeTime
	}
	return json.Marshal(toSerialize)
}

type NullableAccessTransferInformation struct {
	value *AccessTransferInformation
	isSet bool
}

func (v NullableAccessTransferInformation) Get() *AccessTransferInformation {
	return v.value
}

func (v *NullableAccessTransferInformation) Set(val *AccessTransferInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessTransferInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessTransferInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessTransferInformation(val *AccessTransferInformation) *NullableAccessTransferInformation {
	return &NullableAccessTransferInformation{value: val, isSet: true}
}

func (v NullableAccessTransferInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessTransferInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

