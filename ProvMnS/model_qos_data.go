/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ProvMnS

import (
	"encoding/json"
)

// QosData struct for QosData
type QosData struct {
	QosId *string `json:"qosId,omitempty"`
	FiveQIValue *int32 `json:"fiveQIValue,omitempty"`
	// This data type is defined in the same way as the 'BitRate' data type, but with the OpenAPI 'nullable: true' property. 
	MaxbrUl NullableString `json:"maxbrUl,omitempty"`
	// This data type is defined in the same way as the 'BitRate' data type, but with the OpenAPI 'nullable: true' property. 
	MaxbrDl NullableString `json:"maxbrDl,omitempty"`
	// This data type is defined in the same way as the 'BitRate' data type, but with the OpenAPI 'nullable: true' property. 
	GbrUl NullableString `json:"gbrUl,omitempty"`
	// This data type is defined in the same way as the 'BitRate' data type, but with the OpenAPI 'nullable: true' property. 
	GbrDl NullableString `json:"gbrDl,omitempty"`
	Arp *Arp `json:"arp,omitempty"`
	QosNotificationControl *bool `json:"qosNotificationControl,omitempty"`
	ReflectiveQos *bool `json:"reflectiveQos,omitempty"`
	SharingKeyDl *string `json:"sharingKeyDl,omitempty"`
	SharingKeyUl *string `json:"sharingKeyUl,omitempty"`
	// This data type is defined in the same way as the 'PacketLossRate' data type, but with the OpenAPI 'nullable: true' property. 
	MaxPacketLossRateDl NullableInt32 `json:"maxPacketLossRateDl,omitempty"`
	// This data type is defined in the same way as the 'PacketLossRate' data type, but with the OpenAPI 'nullable: true' property. 
	MaxPacketLossRateUl NullableInt32 `json:"maxPacketLossRateUl,omitempty"`
	// This data type is defined in the same way as the 'ExtMaxDataBurstVol' data type, but with the OpenAPI 'nullable: true' property. 
	ExtMaxDataBurstVol NullableInt32 `json:"extMaxDataBurstVol,omitempty"`
}

// NewQosData instantiates a new QosData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQosData() *QosData {
	this := QosData{}
	return &this
}

// NewQosDataWithDefaults instantiates a new QosData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQosDataWithDefaults() *QosData {
	this := QosData{}
	return &this
}

// GetQosId returns the QosId field value if set, zero value otherwise.
func (o *QosData) GetQosId() string {
	if o == nil || isNil(o.QosId) {
		var ret string
		return ret
	}
	return *o.QosId
}

// GetQosIdOk returns a tuple with the QosId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosData) GetQosIdOk() (*string, bool) {
	if o == nil || isNil(o.QosId) {
    return nil, false
	}
	return o.QosId, true
}

// HasQosId returns a boolean if a field has been set.
func (o *QosData) HasQosId() bool {
	if o != nil && !isNil(o.QosId) {
		return true
	}

	return false
}

// SetQosId gets a reference to the given string and assigns it to the QosId field.
func (o *QosData) SetQosId(v string) {
	o.QosId = &v
}

// GetFiveQIValue returns the FiveQIValue field value if set, zero value otherwise.
func (o *QosData) GetFiveQIValue() int32 {
	if o == nil || isNil(o.FiveQIValue) {
		var ret int32
		return ret
	}
	return *o.FiveQIValue
}

// GetFiveQIValueOk returns a tuple with the FiveQIValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosData) GetFiveQIValueOk() (*int32, bool) {
	if o == nil || isNil(o.FiveQIValue) {
    return nil, false
	}
	return o.FiveQIValue, true
}

// HasFiveQIValue returns a boolean if a field has been set.
func (o *QosData) HasFiveQIValue() bool {
	if o != nil && !isNil(o.FiveQIValue) {
		return true
	}

	return false
}

// SetFiveQIValue gets a reference to the given int32 and assigns it to the FiveQIValue field.
func (o *QosData) SetFiveQIValue(v int32) {
	o.FiveQIValue = &v
}

// GetMaxbrUl returns the MaxbrUl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QosData) GetMaxbrUl() string {
	if o == nil || isNil(o.MaxbrUl.Get()) {
		var ret string
		return ret
	}
	return *o.MaxbrUl.Get()
}

// GetMaxbrUlOk returns a tuple with the MaxbrUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QosData) GetMaxbrUlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.MaxbrUl.Get(), o.MaxbrUl.IsSet()
}

// HasMaxbrUl returns a boolean if a field has been set.
func (o *QosData) HasMaxbrUl() bool {
	if o != nil && o.MaxbrUl.IsSet() {
		return true
	}

	return false
}

// SetMaxbrUl gets a reference to the given NullableString and assigns it to the MaxbrUl field.
func (o *QosData) SetMaxbrUl(v string) {
	o.MaxbrUl.Set(&v)
}
// SetMaxbrUlNil sets the value for MaxbrUl to be an explicit nil
func (o *QosData) SetMaxbrUlNil() {
	o.MaxbrUl.Set(nil)
}

// UnsetMaxbrUl ensures that no value is present for MaxbrUl, not even an explicit nil
func (o *QosData) UnsetMaxbrUl() {
	o.MaxbrUl.Unset()
}

// GetMaxbrDl returns the MaxbrDl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QosData) GetMaxbrDl() string {
	if o == nil || isNil(o.MaxbrDl.Get()) {
		var ret string
		return ret
	}
	return *o.MaxbrDl.Get()
}

// GetMaxbrDlOk returns a tuple with the MaxbrDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QosData) GetMaxbrDlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.MaxbrDl.Get(), o.MaxbrDl.IsSet()
}

// HasMaxbrDl returns a boolean if a field has been set.
func (o *QosData) HasMaxbrDl() bool {
	if o != nil && o.MaxbrDl.IsSet() {
		return true
	}

	return false
}

// SetMaxbrDl gets a reference to the given NullableString and assigns it to the MaxbrDl field.
func (o *QosData) SetMaxbrDl(v string) {
	o.MaxbrDl.Set(&v)
}
// SetMaxbrDlNil sets the value for MaxbrDl to be an explicit nil
func (o *QosData) SetMaxbrDlNil() {
	o.MaxbrDl.Set(nil)
}

// UnsetMaxbrDl ensures that no value is present for MaxbrDl, not even an explicit nil
func (o *QosData) UnsetMaxbrDl() {
	o.MaxbrDl.Unset()
}

// GetGbrUl returns the GbrUl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QosData) GetGbrUl() string {
	if o == nil || isNil(o.GbrUl.Get()) {
		var ret string
		return ret
	}
	return *o.GbrUl.Get()
}

// GetGbrUlOk returns a tuple with the GbrUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QosData) GetGbrUlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.GbrUl.Get(), o.GbrUl.IsSet()
}

// HasGbrUl returns a boolean if a field has been set.
func (o *QosData) HasGbrUl() bool {
	if o != nil && o.GbrUl.IsSet() {
		return true
	}

	return false
}

// SetGbrUl gets a reference to the given NullableString and assigns it to the GbrUl field.
func (o *QosData) SetGbrUl(v string) {
	o.GbrUl.Set(&v)
}
// SetGbrUlNil sets the value for GbrUl to be an explicit nil
func (o *QosData) SetGbrUlNil() {
	o.GbrUl.Set(nil)
}

// UnsetGbrUl ensures that no value is present for GbrUl, not even an explicit nil
func (o *QosData) UnsetGbrUl() {
	o.GbrUl.Unset()
}

// GetGbrDl returns the GbrDl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QosData) GetGbrDl() string {
	if o == nil || isNil(o.GbrDl.Get()) {
		var ret string
		return ret
	}
	return *o.GbrDl.Get()
}

// GetGbrDlOk returns a tuple with the GbrDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QosData) GetGbrDlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.GbrDl.Get(), o.GbrDl.IsSet()
}

// HasGbrDl returns a boolean if a field has been set.
func (o *QosData) HasGbrDl() bool {
	if o != nil && o.GbrDl.IsSet() {
		return true
	}

	return false
}

// SetGbrDl gets a reference to the given NullableString and assigns it to the GbrDl field.
func (o *QosData) SetGbrDl(v string) {
	o.GbrDl.Set(&v)
}
// SetGbrDlNil sets the value for GbrDl to be an explicit nil
func (o *QosData) SetGbrDlNil() {
	o.GbrDl.Set(nil)
}

// UnsetGbrDl ensures that no value is present for GbrDl, not even an explicit nil
func (o *QosData) UnsetGbrDl() {
	o.GbrDl.Unset()
}

// GetArp returns the Arp field value if set, zero value otherwise.
func (o *QosData) GetArp() Arp {
	if o == nil || isNil(o.Arp) {
		var ret Arp
		return ret
	}
	return *o.Arp
}

// GetArpOk returns a tuple with the Arp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosData) GetArpOk() (*Arp, bool) {
	if o == nil || isNil(o.Arp) {
    return nil, false
	}
	return o.Arp, true
}

// HasArp returns a boolean if a field has been set.
func (o *QosData) HasArp() bool {
	if o != nil && !isNil(o.Arp) {
		return true
	}

	return false
}

// SetArp gets a reference to the given Arp and assigns it to the Arp field.
func (o *QosData) SetArp(v Arp) {
	o.Arp = &v
}

// GetQosNotificationControl returns the QosNotificationControl field value if set, zero value otherwise.
func (o *QosData) GetQosNotificationControl() bool {
	if o == nil || isNil(o.QosNotificationControl) {
		var ret bool
		return ret
	}
	return *o.QosNotificationControl
}

// GetQosNotificationControlOk returns a tuple with the QosNotificationControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosData) GetQosNotificationControlOk() (*bool, bool) {
	if o == nil || isNil(o.QosNotificationControl) {
    return nil, false
	}
	return o.QosNotificationControl, true
}

// HasQosNotificationControl returns a boolean if a field has been set.
func (o *QosData) HasQosNotificationControl() bool {
	if o != nil && !isNil(o.QosNotificationControl) {
		return true
	}

	return false
}

// SetQosNotificationControl gets a reference to the given bool and assigns it to the QosNotificationControl field.
func (o *QosData) SetQosNotificationControl(v bool) {
	o.QosNotificationControl = &v
}

// GetReflectiveQos returns the ReflectiveQos field value if set, zero value otherwise.
func (o *QosData) GetReflectiveQos() bool {
	if o == nil || isNil(o.ReflectiveQos) {
		var ret bool
		return ret
	}
	return *o.ReflectiveQos
}

// GetReflectiveQosOk returns a tuple with the ReflectiveQos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosData) GetReflectiveQosOk() (*bool, bool) {
	if o == nil || isNil(o.ReflectiveQos) {
    return nil, false
	}
	return o.ReflectiveQos, true
}

// HasReflectiveQos returns a boolean if a field has been set.
func (o *QosData) HasReflectiveQos() bool {
	if o != nil && !isNil(o.ReflectiveQos) {
		return true
	}

	return false
}

// SetReflectiveQos gets a reference to the given bool and assigns it to the ReflectiveQos field.
func (o *QosData) SetReflectiveQos(v bool) {
	o.ReflectiveQos = &v
}

// GetSharingKeyDl returns the SharingKeyDl field value if set, zero value otherwise.
func (o *QosData) GetSharingKeyDl() string {
	if o == nil || isNil(o.SharingKeyDl) {
		var ret string
		return ret
	}
	return *o.SharingKeyDl
}

// GetSharingKeyDlOk returns a tuple with the SharingKeyDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosData) GetSharingKeyDlOk() (*string, bool) {
	if o == nil || isNil(o.SharingKeyDl) {
    return nil, false
	}
	return o.SharingKeyDl, true
}

// HasSharingKeyDl returns a boolean if a field has been set.
func (o *QosData) HasSharingKeyDl() bool {
	if o != nil && !isNil(o.SharingKeyDl) {
		return true
	}

	return false
}

// SetSharingKeyDl gets a reference to the given string and assigns it to the SharingKeyDl field.
func (o *QosData) SetSharingKeyDl(v string) {
	o.SharingKeyDl = &v
}

// GetSharingKeyUl returns the SharingKeyUl field value if set, zero value otherwise.
func (o *QosData) GetSharingKeyUl() string {
	if o == nil || isNil(o.SharingKeyUl) {
		var ret string
		return ret
	}
	return *o.SharingKeyUl
}

// GetSharingKeyUlOk returns a tuple with the SharingKeyUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosData) GetSharingKeyUlOk() (*string, bool) {
	if o == nil || isNil(o.SharingKeyUl) {
    return nil, false
	}
	return o.SharingKeyUl, true
}

// HasSharingKeyUl returns a boolean if a field has been set.
func (o *QosData) HasSharingKeyUl() bool {
	if o != nil && !isNil(o.SharingKeyUl) {
		return true
	}

	return false
}

// SetSharingKeyUl gets a reference to the given string and assigns it to the SharingKeyUl field.
func (o *QosData) SetSharingKeyUl(v string) {
	o.SharingKeyUl = &v
}

// GetMaxPacketLossRateDl returns the MaxPacketLossRateDl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QosData) GetMaxPacketLossRateDl() int32 {
	if o == nil || isNil(o.MaxPacketLossRateDl.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxPacketLossRateDl.Get()
}

// GetMaxPacketLossRateDlOk returns a tuple with the MaxPacketLossRateDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QosData) GetMaxPacketLossRateDlOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.MaxPacketLossRateDl.Get(), o.MaxPacketLossRateDl.IsSet()
}

// HasMaxPacketLossRateDl returns a boolean if a field has been set.
func (o *QosData) HasMaxPacketLossRateDl() bool {
	if o != nil && o.MaxPacketLossRateDl.IsSet() {
		return true
	}

	return false
}

// SetMaxPacketLossRateDl gets a reference to the given NullableInt32 and assigns it to the MaxPacketLossRateDl field.
func (o *QosData) SetMaxPacketLossRateDl(v int32) {
	o.MaxPacketLossRateDl.Set(&v)
}
// SetMaxPacketLossRateDlNil sets the value for MaxPacketLossRateDl to be an explicit nil
func (o *QosData) SetMaxPacketLossRateDlNil() {
	o.MaxPacketLossRateDl.Set(nil)
}

// UnsetMaxPacketLossRateDl ensures that no value is present for MaxPacketLossRateDl, not even an explicit nil
func (o *QosData) UnsetMaxPacketLossRateDl() {
	o.MaxPacketLossRateDl.Unset()
}

// GetMaxPacketLossRateUl returns the MaxPacketLossRateUl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QosData) GetMaxPacketLossRateUl() int32 {
	if o == nil || isNil(o.MaxPacketLossRateUl.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxPacketLossRateUl.Get()
}

// GetMaxPacketLossRateUlOk returns a tuple with the MaxPacketLossRateUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QosData) GetMaxPacketLossRateUlOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.MaxPacketLossRateUl.Get(), o.MaxPacketLossRateUl.IsSet()
}

// HasMaxPacketLossRateUl returns a boolean if a field has been set.
func (o *QosData) HasMaxPacketLossRateUl() bool {
	if o != nil && o.MaxPacketLossRateUl.IsSet() {
		return true
	}

	return false
}

// SetMaxPacketLossRateUl gets a reference to the given NullableInt32 and assigns it to the MaxPacketLossRateUl field.
func (o *QosData) SetMaxPacketLossRateUl(v int32) {
	o.MaxPacketLossRateUl.Set(&v)
}
// SetMaxPacketLossRateUlNil sets the value for MaxPacketLossRateUl to be an explicit nil
func (o *QosData) SetMaxPacketLossRateUlNil() {
	o.MaxPacketLossRateUl.Set(nil)
}

// UnsetMaxPacketLossRateUl ensures that no value is present for MaxPacketLossRateUl, not even an explicit nil
func (o *QosData) UnsetMaxPacketLossRateUl() {
	o.MaxPacketLossRateUl.Unset()
}

// GetExtMaxDataBurstVol returns the ExtMaxDataBurstVol field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QosData) GetExtMaxDataBurstVol() int32 {
	if o == nil || isNil(o.ExtMaxDataBurstVol.Get()) {
		var ret int32
		return ret
	}
	return *o.ExtMaxDataBurstVol.Get()
}

// GetExtMaxDataBurstVolOk returns a tuple with the ExtMaxDataBurstVol field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QosData) GetExtMaxDataBurstVolOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.ExtMaxDataBurstVol.Get(), o.ExtMaxDataBurstVol.IsSet()
}

// HasExtMaxDataBurstVol returns a boolean if a field has been set.
func (o *QosData) HasExtMaxDataBurstVol() bool {
	if o != nil && o.ExtMaxDataBurstVol.IsSet() {
		return true
	}

	return false
}

// SetExtMaxDataBurstVol gets a reference to the given NullableInt32 and assigns it to the ExtMaxDataBurstVol field.
func (o *QosData) SetExtMaxDataBurstVol(v int32) {
	o.ExtMaxDataBurstVol.Set(&v)
}
// SetExtMaxDataBurstVolNil sets the value for ExtMaxDataBurstVol to be an explicit nil
func (o *QosData) SetExtMaxDataBurstVolNil() {
	o.ExtMaxDataBurstVol.Set(nil)
}

// UnsetExtMaxDataBurstVol ensures that no value is present for ExtMaxDataBurstVol, not even an explicit nil
func (o *QosData) UnsetExtMaxDataBurstVol() {
	o.ExtMaxDataBurstVol.Unset()
}

func (o QosData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.QosId) {
		toSerialize["qosId"] = o.QosId
	}
	if !isNil(o.FiveQIValue) {
		toSerialize["fiveQIValue"] = o.FiveQIValue
	}
	if o.MaxbrUl.IsSet() {
		toSerialize["maxbrUl"] = o.MaxbrUl.Get()
	}
	if o.MaxbrDl.IsSet() {
		toSerialize["maxbrDl"] = o.MaxbrDl.Get()
	}
	if o.GbrUl.IsSet() {
		toSerialize["gbrUl"] = o.GbrUl.Get()
	}
	if o.GbrDl.IsSet() {
		toSerialize["gbrDl"] = o.GbrDl.Get()
	}
	if !isNil(o.Arp) {
		toSerialize["arp"] = o.Arp
	}
	if !isNil(o.QosNotificationControl) {
		toSerialize["qosNotificationControl"] = o.QosNotificationControl
	}
	if !isNil(o.ReflectiveQos) {
		toSerialize["reflectiveQos"] = o.ReflectiveQos
	}
	if !isNil(o.SharingKeyDl) {
		toSerialize["sharingKeyDl"] = o.SharingKeyDl
	}
	if !isNil(o.SharingKeyUl) {
		toSerialize["sharingKeyUl"] = o.SharingKeyUl
	}
	if o.MaxPacketLossRateDl.IsSet() {
		toSerialize["maxPacketLossRateDl"] = o.MaxPacketLossRateDl.Get()
	}
	if o.MaxPacketLossRateUl.IsSet() {
		toSerialize["maxPacketLossRateUl"] = o.MaxPacketLossRateUl.Get()
	}
	if o.ExtMaxDataBurstVol.IsSet() {
		toSerialize["extMaxDataBurstVol"] = o.ExtMaxDataBurstVol.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableQosData struct {
	value *QosData
	isSet bool
}

func (v NullableQosData) Get() *QosData {
	return v.value
}

func (v *NullableQosData) Set(val *QosData) {
	v.value = val
	v.isSet = true
}

func (v NullableQosData) IsSet() bool {
	return v.isSet
}

func (v *NullableQosData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQosData(val *QosData) *NullableQosData {
	return &NullableQosData{value: val, isSet: true}
}

func (v NullableQosData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQosData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

