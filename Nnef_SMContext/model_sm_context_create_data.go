/*
Nnef_SMContext

Nnef SMContext Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnef_SMContext

import (
	"encoding/json"
)

// SmContextCreateData Representation of the Individual SM context to be created.
type SmContextCreateData struct {
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501. 
	Supi string `json:"supi"`
	// Unsigned integer identifying a PDU session, within the range 0 to 255, as specified in  clause 11.2.3.1b, bits 1 to 8, of 3GPP TS 24.007. If the PDU Session ID is allocated by the  Core Network for UEs not supporting N1 mode, reserved range 64 to 95 is used. PDU Session ID  within the reserved range is only visible in the Core Network.  
	PduSessionId int32 `json:"pduSessionId"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn string `json:"dnn"`
	Snssai Snssai `json:"snssai"`
	// This IE shall contain the NEF ID of the target NEF.
	NefId string `json:"nefId"`
	// String providing an URI formatted according to RFC 3986.
	DlNiddEndPoint string `json:"dlNiddEndPoint"`
	// String providing an URI formatted according to RFC 3986.
	NotificationUri string `json:"notificationUri"`
	NiddInfo *NiddInformation `json:"niddInfo,omitempty"`
	// When present, this IE shall indicate the UE capability to support RDS. The value of this IE shall be set as following  - true  UE supports RDS  - false (default)  UE does not support RDS 
	RdsSupport *bool `json:"rdsSupport,omitempty"`
	SmContextConfig *SmContextConfiguration `json:"smContextConfig,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
}

// NewSmContextCreateData instantiates a new SmContextCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmContextCreateData(supi string, pduSessionId int32, dnn string, snssai Snssai, nefId string, dlNiddEndPoint string, notificationUri string) *SmContextCreateData {
	this := SmContextCreateData{}
	this.Supi = supi
	this.PduSessionId = pduSessionId
	this.Dnn = dnn
	this.Snssai = snssai
	this.NefId = nefId
	this.DlNiddEndPoint = dlNiddEndPoint
	this.NotificationUri = notificationUri
	return &this
}

// NewSmContextCreateDataWithDefaults instantiates a new SmContextCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmContextCreateDataWithDefaults() *SmContextCreateData {
	this := SmContextCreateData{}
	return &this
}

// GetSupi returns the Supi field value
func (o *SmContextCreateData) GetSupi() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Supi
}

// GetSupiOk returns a tuple with the Supi field value
// and a boolean to check if the value has been set.
func (o *SmContextCreateData) GetSupiOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Supi, true
}

// SetSupi sets field value
func (o *SmContextCreateData) SetSupi(v string) {
	o.Supi = v
}

// GetPduSessionId returns the PduSessionId field value
func (o *SmContextCreateData) GetPduSessionId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PduSessionId
}

// GetPduSessionIdOk returns a tuple with the PduSessionId field value
// and a boolean to check if the value has been set.
func (o *SmContextCreateData) GetPduSessionIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PduSessionId, true
}

// SetPduSessionId sets field value
func (o *SmContextCreateData) SetPduSessionId(v int32) {
	o.PduSessionId = v
}

// GetDnn returns the Dnn field value
func (o *SmContextCreateData) GetDnn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value
// and a boolean to check if the value has been set.
func (o *SmContextCreateData) GetDnnOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Dnn, true
}

// SetDnn sets field value
func (o *SmContextCreateData) SetDnn(v string) {
	o.Dnn = v
}

// GetSnssai returns the Snssai field value
func (o *SmContextCreateData) GetSnssai() Snssai {
	if o == nil {
		var ret Snssai
		return ret
	}

	return o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value
// and a boolean to check if the value has been set.
func (o *SmContextCreateData) GetSnssaiOk() (*Snssai, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Snssai, true
}

// SetSnssai sets field value
func (o *SmContextCreateData) SetSnssai(v Snssai) {
	o.Snssai = v
}

// GetNefId returns the NefId field value
func (o *SmContextCreateData) GetNefId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NefId
}

// GetNefIdOk returns a tuple with the NefId field value
// and a boolean to check if the value has been set.
func (o *SmContextCreateData) GetNefIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NefId, true
}

// SetNefId sets field value
func (o *SmContextCreateData) SetNefId(v string) {
	o.NefId = v
}

// GetDlNiddEndPoint returns the DlNiddEndPoint field value
func (o *SmContextCreateData) GetDlNiddEndPoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DlNiddEndPoint
}

// GetDlNiddEndPointOk returns a tuple with the DlNiddEndPoint field value
// and a boolean to check if the value has been set.
func (o *SmContextCreateData) GetDlNiddEndPointOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DlNiddEndPoint, true
}

// SetDlNiddEndPoint sets field value
func (o *SmContextCreateData) SetDlNiddEndPoint(v string) {
	o.DlNiddEndPoint = v
}

// GetNotificationUri returns the NotificationUri field value
func (o *SmContextCreateData) GetNotificationUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotificationUri
}

// GetNotificationUriOk returns a tuple with the NotificationUri field value
// and a boolean to check if the value has been set.
func (o *SmContextCreateData) GetNotificationUriOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NotificationUri, true
}

// SetNotificationUri sets field value
func (o *SmContextCreateData) SetNotificationUri(v string) {
	o.NotificationUri = v
}

// GetNiddInfo returns the NiddInfo field value if set, zero value otherwise.
func (o *SmContextCreateData) GetNiddInfo() NiddInformation {
	if o == nil || isNil(o.NiddInfo) {
		var ret NiddInformation
		return ret
	}
	return *o.NiddInfo
}

// GetNiddInfoOk returns a tuple with the NiddInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmContextCreateData) GetNiddInfoOk() (*NiddInformation, bool) {
	if o == nil || isNil(o.NiddInfo) {
    return nil, false
	}
	return o.NiddInfo, true
}

// HasNiddInfo returns a boolean if a field has been set.
func (o *SmContextCreateData) HasNiddInfo() bool {
	if o != nil && !isNil(o.NiddInfo) {
		return true
	}

	return false
}

// SetNiddInfo gets a reference to the given NiddInformation and assigns it to the NiddInfo field.
func (o *SmContextCreateData) SetNiddInfo(v NiddInformation) {
	o.NiddInfo = &v
}

// GetRdsSupport returns the RdsSupport field value if set, zero value otherwise.
func (o *SmContextCreateData) GetRdsSupport() bool {
	if o == nil || isNil(o.RdsSupport) {
		var ret bool
		return ret
	}
	return *o.RdsSupport
}

// GetRdsSupportOk returns a tuple with the RdsSupport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmContextCreateData) GetRdsSupportOk() (*bool, bool) {
	if o == nil || isNil(o.RdsSupport) {
    return nil, false
	}
	return o.RdsSupport, true
}

// HasRdsSupport returns a boolean if a field has been set.
func (o *SmContextCreateData) HasRdsSupport() bool {
	if o != nil && !isNil(o.RdsSupport) {
		return true
	}

	return false
}

// SetRdsSupport gets a reference to the given bool and assigns it to the RdsSupport field.
func (o *SmContextCreateData) SetRdsSupport(v bool) {
	o.RdsSupport = &v
}

// GetSmContextConfig returns the SmContextConfig field value if set, zero value otherwise.
func (o *SmContextCreateData) GetSmContextConfig() SmContextConfiguration {
	if o == nil || isNil(o.SmContextConfig) {
		var ret SmContextConfiguration
		return ret
	}
	return *o.SmContextConfig
}

// GetSmContextConfigOk returns a tuple with the SmContextConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmContextCreateData) GetSmContextConfigOk() (*SmContextConfiguration, bool) {
	if o == nil || isNil(o.SmContextConfig) {
    return nil, false
	}
	return o.SmContextConfig, true
}

// HasSmContextConfig returns a boolean if a field has been set.
func (o *SmContextCreateData) HasSmContextConfig() bool {
	if o != nil && !isNil(o.SmContextConfig) {
		return true
	}

	return false
}

// SetSmContextConfig gets a reference to the given SmContextConfiguration and assigns it to the SmContextConfig field.
func (o *SmContextCreateData) SetSmContextConfig(v SmContextConfiguration) {
	o.SmContextConfig = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *SmContextCreateData) GetSupportedFeatures() string {
	if o == nil || isNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmContextCreateData) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || isNil(o.SupportedFeatures) {
    return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *SmContextCreateData) HasSupportedFeatures() bool {
	if o != nil && !isNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *SmContextCreateData) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

func (o SmContextCreateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["supi"] = o.Supi
	}
	if true {
		toSerialize["pduSessionId"] = o.PduSessionId
	}
	if true {
		toSerialize["dnn"] = o.Dnn
	}
	if true {
		toSerialize["snssai"] = o.Snssai
	}
	if true {
		toSerialize["nefId"] = o.NefId
	}
	if true {
		toSerialize["dlNiddEndPoint"] = o.DlNiddEndPoint
	}
	if true {
		toSerialize["notificationUri"] = o.NotificationUri
	}
	if !isNil(o.NiddInfo) {
		toSerialize["niddInfo"] = o.NiddInfo
	}
	if !isNil(o.RdsSupport) {
		toSerialize["rdsSupport"] = o.RdsSupport
	}
	if !isNil(o.SmContextConfig) {
		toSerialize["smContextConfig"] = o.SmContextConfig
	}
	if !isNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	return json.Marshal(toSerialize)
}

type NullableSmContextCreateData struct {
	value *SmContextCreateData
	isSet bool
}

func (v NullableSmContextCreateData) Get() *SmContextCreateData {
	return v.value
}

func (v *NullableSmContextCreateData) Set(val *SmContextCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableSmContextCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableSmContextCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmContextCreateData(val *SmContextCreateData) *NullableSmContextCreateData {
	return &NullableSmContextCreateData{value: val, isSet: true}
}

func (v NullableSmContextCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmContextCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

