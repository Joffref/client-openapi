/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Namf_Communication

import (
	"encoding/json"
)

// N1N2MessageTransferReqData Data within a N1/N2 message transfer request
type N1N2MessageTransferReqData struct {
	N1MessageContainer *N1MessageContainer `json:"n1MessageContainer,omitempty"`
	N2InfoContainer *N2InfoContainer `json:"n2InfoContainer,omitempty"`
	MtData *RefToBinaryData `json:"mtData,omitempty"`
	SkipInd *bool `json:"skipInd,omitempty"`
	LastMsgIndication *bool `json:"lastMsgIndication,omitempty"`
	// Unsigned integer identifying a PDU session, within the range 0 to 255, as specified in  clause 11.2.3.1b, bits 1 to 8, of 3GPP TS 24.007. If the PDU Session ID is allocated by the  Core Network for UEs not supporting N1 mode, reserved range 64 to 95 is used. PDU Session ID  within the reserved range is only visible in the Core Network.  
	PduSessionId *int32 `json:"pduSessionId,omitempty"`
	// LCS Correlation ID.
	LcsCorrelationId *string `json:"lcsCorrelationId,omitempty"`
	// Paging Policy Indicator
	Ppi *int32 `json:"ppi,omitempty"`
	Arp *Arp `json:"arp,omitempty"`
	// Unsigned integer representing a 5G QoS Identifier (see clause 5.7.2.1 of 3GPP TS 23.501, within the range 0 to 255. 
	Var5qi *int32 `json:"5qi,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	N1n2FailureTxfNotifURI *string `json:"n1n2FailureTxfNotifURI,omitempty"`
	SmfReallocationInd *bool `json:"smfReallocationInd,omitempty"`
	AreaOfValidity *AreaOfValidity `json:"areaOfValidity,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
	OldGuami *Guami `json:"oldGuami,omitempty"`
	MaAcceptedInd *bool `json:"maAcceptedInd,omitempty"`
	ExtBufSupport *bool `json:"extBufSupport,omitempty"`
	TargetAccess *AccessType `json:"targetAccess,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	NfId *string `json:"nfId,omitempty"`
}

// NewN1N2MessageTransferReqData instantiates a new N1N2MessageTransferReqData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewN1N2MessageTransferReqData() *N1N2MessageTransferReqData {
	this := N1N2MessageTransferReqData{}
	var skipInd bool = false
	this.SkipInd = &skipInd
	var smfReallocationInd bool = false
	this.SmfReallocationInd = &smfReallocationInd
	var maAcceptedInd bool = false
	this.MaAcceptedInd = &maAcceptedInd
	var extBufSupport bool = false
	this.ExtBufSupport = &extBufSupport
	return &this
}

// NewN1N2MessageTransferReqDataWithDefaults instantiates a new N1N2MessageTransferReqData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewN1N2MessageTransferReqDataWithDefaults() *N1N2MessageTransferReqData {
	this := N1N2MessageTransferReqData{}
	var skipInd bool = false
	this.SkipInd = &skipInd
	var smfReallocationInd bool = false
	this.SmfReallocationInd = &smfReallocationInd
	var maAcceptedInd bool = false
	this.MaAcceptedInd = &maAcceptedInd
	var extBufSupport bool = false
	this.ExtBufSupport = &extBufSupport
	return &this
}

// GetN1MessageContainer returns the N1MessageContainer field value if set, zero value otherwise.
func (o *N1N2MessageTransferReqData) GetN1MessageContainer() N1MessageContainer {
	if o == nil || isNil(o.N1MessageContainer) {
		var ret N1MessageContainer
		return ret
	}
	return *o.N1MessageContainer
}

// GetN1MessageContainerOk returns a tuple with the N1MessageContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N1N2MessageTransferReqData) GetN1MessageContainerOk() (*N1MessageContainer, bool) {
	if o == nil || isNil(o.N1MessageContainer) {
    return nil, false
	}
	return o.N1MessageContainer, true
}

// HasN1MessageContainer returns a boolean if a field has been set.
func (o *N1N2MessageTransferReqData) HasN1MessageContainer() bool {
	if o != nil && !isNil(o.N1MessageContainer) {
		return true
	}

	return false
}

// SetN1MessageContainer gets a reference to the given N1MessageContainer and assigns it to the N1MessageContainer field.
func (o *N1N2MessageTransferReqData) SetN1MessageContainer(v N1MessageContainer) {
	o.N1MessageContainer = &v
}

// GetN2InfoContainer returns the N2InfoContainer field value if set, zero value otherwise.
func (o *N1N2MessageTransferReqData) GetN2InfoContainer() N2InfoContainer {
	if o == nil || isNil(o.N2InfoContainer) {
		var ret N2InfoContainer
		return ret
	}
	return *o.N2InfoContainer
}

// GetN2InfoContainerOk returns a tuple with the N2InfoContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N1N2MessageTransferReqData) GetN2InfoContainerOk() (*N2InfoContainer, bool) {
	if o == nil || isNil(o.N2InfoContainer) {
    return nil, false
	}
	return o.N2InfoContainer, true
}

// HasN2InfoContainer returns a boolean if a field has been set.
func (o *N1N2MessageTransferReqData) HasN2InfoContainer() bool {
	if o != nil && !isNil(o.N2InfoContainer) {
		return true
	}

	return false
}

// SetN2InfoContainer gets a reference to the given N2InfoContainer and assigns it to the N2InfoContainer field.
func (o *N1N2MessageTransferReqData) SetN2InfoContainer(v N2InfoContainer) {
	o.N2InfoContainer = &v
}

// GetMtData returns the MtData field value if set, zero value otherwise.
func (o *N1N2MessageTransferReqData) GetMtData() RefToBinaryData {
	if o == nil || isNil(o.MtData) {
		var ret RefToBinaryData
		return ret
	}
	return *o.MtData
}

// GetMtDataOk returns a tuple with the MtData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N1N2MessageTransferReqData) GetMtDataOk() (*RefToBinaryData, bool) {
	if o == nil || isNil(o.MtData) {
    return nil, false
	}
	return o.MtData, true
}

// HasMtData returns a boolean if a field has been set.
func (o *N1N2MessageTransferReqData) HasMtData() bool {
	if o != nil && !isNil(o.MtData) {
		return true
	}

	return false
}

// SetMtData gets a reference to the given RefToBinaryData and assigns it to the MtData field.
func (o *N1N2MessageTransferReqData) SetMtData(v RefToBinaryData) {
	o.MtData = &v
}

// GetSkipInd returns the SkipInd field value if set, zero value otherwise.
func (o *N1N2MessageTransferReqData) GetSkipInd() bool {
	if o == nil || isNil(o.SkipInd) {
		var ret bool
		return ret
	}
	return *o.SkipInd
}

// GetSkipIndOk returns a tuple with the SkipInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N1N2MessageTransferReqData) GetSkipIndOk() (*bool, bool) {
	if o == nil || isNil(o.SkipInd) {
    return nil, false
	}
	return o.SkipInd, true
}

// HasSkipInd returns a boolean if a field has been set.
func (o *N1N2MessageTransferReqData) HasSkipInd() bool {
	if o != nil && !isNil(o.SkipInd) {
		return true
	}

	return false
}

// SetSkipInd gets a reference to the given bool and assigns it to the SkipInd field.
func (o *N1N2MessageTransferReqData) SetSkipInd(v bool) {
	o.SkipInd = &v
}

// GetLastMsgIndication returns the LastMsgIndication field value if set, zero value otherwise.
func (o *N1N2MessageTransferReqData) GetLastMsgIndication() bool {
	if o == nil || isNil(o.LastMsgIndication) {
		var ret bool
		return ret
	}
	return *o.LastMsgIndication
}

// GetLastMsgIndicationOk returns a tuple with the LastMsgIndication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N1N2MessageTransferReqData) GetLastMsgIndicationOk() (*bool, bool) {
	if o == nil || isNil(o.LastMsgIndication) {
    return nil, false
	}
	return o.LastMsgIndication, true
}

// HasLastMsgIndication returns a boolean if a field has been set.
func (o *N1N2MessageTransferReqData) HasLastMsgIndication() bool {
	if o != nil && !isNil(o.LastMsgIndication) {
		return true
	}

	return false
}

// SetLastMsgIndication gets a reference to the given bool and assigns it to the LastMsgIndication field.
func (o *N1N2MessageTransferReqData) SetLastMsgIndication(v bool) {
	o.LastMsgIndication = &v
}

// GetPduSessionId returns the PduSessionId field value if set, zero value otherwise.
func (o *N1N2MessageTransferReqData) GetPduSessionId() int32 {
	if o == nil || isNil(o.PduSessionId) {
		var ret int32
		return ret
	}
	return *o.PduSessionId
}

// GetPduSessionIdOk returns a tuple with the PduSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N1N2MessageTransferReqData) GetPduSessionIdOk() (*int32, bool) {
	if o == nil || isNil(o.PduSessionId) {
    return nil, false
	}
	return o.PduSessionId, true
}

// HasPduSessionId returns a boolean if a field has been set.
func (o *N1N2MessageTransferReqData) HasPduSessionId() bool {
	if o != nil && !isNil(o.PduSessionId) {
		return true
	}

	return false
}

// SetPduSessionId gets a reference to the given int32 and assigns it to the PduSessionId field.
func (o *N1N2MessageTransferReqData) SetPduSessionId(v int32) {
	o.PduSessionId = &v
}

// GetLcsCorrelationId returns the LcsCorrelationId field value if set, zero value otherwise.
func (o *N1N2MessageTransferReqData) GetLcsCorrelationId() string {
	if o == nil || isNil(o.LcsCorrelationId) {
		var ret string
		return ret
	}
	return *o.LcsCorrelationId
}

// GetLcsCorrelationIdOk returns a tuple with the LcsCorrelationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N1N2MessageTransferReqData) GetLcsCorrelationIdOk() (*string, bool) {
	if o == nil || isNil(o.LcsCorrelationId) {
    return nil, false
	}
	return o.LcsCorrelationId, true
}

// HasLcsCorrelationId returns a boolean if a field has been set.
func (o *N1N2MessageTransferReqData) HasLcsCorrelationId() bool {
	if o != nil && !isNil(o.LcsCorrelationId) {
		return true
	}

	return false
}

// SetLcsCorrelationId gets a reference to the given string and assigns it to the LcsCorrelationId field.
func (o *N1N2MessageTransferReqData) SetLcsCorrelationId(v string) {
	o.LcsCorrelationId = &v
}

// GetPpi returns the Ppi field value if set, zero value otherwise.
func (o *N1N2MessageTransferReqData) GetPpi() int32 {
	if o == nil || isNil(o.Ppi) {
		var ret int32
		return ret
	}
	return *o.Ppi
}

// GetPpiOk returns a tuple with the Ppi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N1N2MessageTransferReqData) GetPpiOk() (*int32, bool) {
	if o == nil || isNil(o.Ppi) {
    return nil, false
	}
	return o.Ppi, true
}

// HasPpi returns a boolean if a field has been set.
func (o *N1N2MessageTransferReqData) HasPpi() bool {
	if o != nil && !isNil(o.Ppi) {
		return true
	}

	return false
}

// SetPpi gets a reference to the given int32 and assigns it to the Ppi field.
func (o *N1N2MessageTransferReqData) SetPpi(v int32) {
	o.Ppi = &v
}

// GetArp returns the Arp field value if set, zero value otherwise.
func (o *N1N2MessageTransferReqData) GetArp() Arp {
	if o == nil || isNil(o.Arp) {
		var ret Arp
		return ret
	}
	return *o.Arp
}

// GetArpOk returns a tuple with the Arp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N1N2MessageTransferReqData) GetArpOk() (*Arp, bool) {
	if o == nil || isNil(o.Arp) {
    return nil, false
	}
	return o.Arp, true
}

// HasArp returns a boolean if a field has been set.
func (o *N1N2MessageTransferReqData) HasArp() bool {
	if o != nil && !isNil(o.Arp) {
		return true
	}

	return false
}

// SetArp gets a reference to the given Arp and assigns it to the Arp field.
func (o *N1N2MessageTransferReqData) SetArp(v Arp) {
	o.Arp = &v
}

// GetVar5qi returns the Var5qi field value if set, zero value otherwise.
func (o *N1N2MessageTransferReqData) GetVar5qi() int32 {
	if o == nil || isNil(o.Var5qi) {
		var ret int32
		return ret
	}
	return *o.Var5qi
}

// GetVar5qiOk returns a tuple with the Var5qi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N1N2MessageTransferReqData) GetVar5qiOk() (*int32, bool) {
	if o == nil || isNil(o.Var5qi) {
    return nil, false
	}
	return o.Var5qi, true
}

// HasVar5qi returns a boolean if a field has been set.
func (o *N1N2MessageTransferReqData) HasVar5qi() bool {
	if o != nil && !isNil(o.Var5qi) {
		return true
	}

	return false
}

// SetVar5qi gets a reference to the given int32 and assigns it to the Var5qi field.
func (o *N1N2MessageTransferReqData) SetVar5qi(v int32) {
	o.Var5qi = &v
}

// GetN1n2FailureTxfNotifURI returns the N1n2FailureTxfNotifURI field value if set, zero value otherwise.
func (o *N1N2MessageTransferReqData) GetN1n2FailureTxfNotifURI() string {
	if o == nil || isNil(o.N1n2FailureTxfNotifURI) {
		var ret string
		return ret
	}
	return *o.N1n2FailureTxfNotifURI
}

// GetN1n2FailureTxfNotifURIOk returns a tuple with the N1n2FailureTxfNotifURI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N1N2MessageTransferReqData) GetN1n2FailureTxfNotifURIOk() (*string, bool) {
	if o == nil || isNil(o.N1n2FailureTxfNotifURI) {
    return nil, false
	}
	return o.N1n2FailureTxfNotifURI, true
}

// HasN1n2FailureTxfNotifURI returns a boolean if a field has been set.
func (o *N1N2MessageTransferReqData) HasN1n2FailureTxfNotifURI() bool {
	if o != nil && !isNil(o.N1n2FailureTxfNotifURI) {
		return true
	}

	return false
}

// SetN1n2FailureTxfNotifURI gets a reference to the given string and assigns it to the N1n2FailureTxfNotifURI field.
func (o *N1N2MessageTransferReqData) SetN1n2FailureTxfNotifURI(v string) {
	o.N1n2FailureTxfNotifURI = &v
}

// GetSmfReallocationInd returns the SmfReallocationInd field value if set, zero value otherwise.
func (o *N1N2MessageTransferReqData) GetSmfReallocationInd() bool {
	if o == nil || isNil(o.SmfReallocationInd) {
		var ret bool
		return ret
	}
	return *o.SmfReallocationInd
}

// GetSmfReallocationIndOk returns a tuple with the SmfReallocationInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N1N2MessageTransferReqData) GetSmfReallocationIndOk() (*bool, bool) {
	if o == nil || isNil(o.SmfReallocationInd) {
    return nil, false
	}
	return o.SmfReallocationInd, true
}

// HasSmfReallocationInd returns a boolean if a field has been set.
func (o *N1N2MessageTransferReqData) HasSmfReallocationInd() bool {
	if o != nil && !isNil(o.SmfReallocationInd) {
		return true
	}

	return false
}

// SetSmfReallocationInd gets a reference to the given bool and assigns it to the SmfReallocationInd field.
func (o *N1N2MessageTransferReqData) SetSmfReallocationInd(v bool) {
	o.SmfReallocationInd = &v
}

// GetAreaOfValidity returns the AreaOfValidity field value if set, zero value otherwise.
func (o *N1N2MessageTransferReqData) GetAreaOfValidity() AreaOfValidity {
	if o == nil || isNil(o.AreaOfValidity) {
		var ret AreaOfValidity
		return ret
	}
	return *o.AreaOfValidity
}

// GetAreaOfValidityOk returns a tuple with the AreaOfValidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N1N2MessageTransferReqData) GetAreaOfValidityOk() (*AreaOfValidity, bool) {
	if o == nil || isNil(o.AreaOfValidity) {
    return nil, false
	}
	return o.AreaOfValidity, true
}

// HasAreaOfValidity returns a boolean if a field has been set.
func (o *N1N2MessageTransferReqData) HasAreaOfValidity() bool {
	if o != nil && !isNil(o.AreaOfValidity) {
		return true
	}

	return false
}

// SetAreaOfValidity gets a reference to the given AreaOfValidity and assigns it to the AreaOfValidity field.
func (o *N1N2MessageTransferReqData) SetAreaOfValidity(v AreaOfValidity) {
	o.AreaOfValidity = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *N1N2MessageTransferReqData) GetSupportedFeatures() string {
	if o == nil || isNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N1N2MessageTransferReqData) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || isNil(o.SupportedFeatures) {
    return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *N1N2MessageTransferReqData) HasSupportedFeatures() bool {
	if o != nil && !isNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *N1N2MessageTransferReqData) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetOldGuami returns the OldGuami field value if set, zero value otherwise.
func (o *N1N2MessageTransferReqData) GetOldGuami() Guami {
	if o == nil || isNil(o.OldGuami) {
		var ret Guami
		return ret
	}
	return *o.OldGuami
}

// GetOldGuamiOk returns a tuple with the OldGuami field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N1N2MessageTransferReqData) GetOldGuamiOk() (*Guami, bool) {
	if o == nil || isNil(o.OldGuami) {
    return nil, false
	}
	return o.OldGuami, true
}

// HasOldGuami returns a boolean if a field has been set.
func (o *N1N2MessageTransferReqData) HasOldGuami() bool {
	if o != nil && !isNil(o.OldGuami) {
		return true
	}

	return false
}

// SetOldGuami gets a reference to the given Guami and assigns it to the OldGuami field.
func (o *N1N2MessageTransferReqData) SetOldGuami(v Guami) {
	o.OldGuami = &v
}

// GetMaAcceptedInd returns the MaAcceptedInd field value if set, zero value otherwise.
func (o *N1N2MessageTransferReqData) GetMaAcceptedInd() bool {
	if o == nil || isNil(o.MaAcceptedInd) {
		var ret bool
		return ret
	}
	return *o.MaAcceptedInd
}

// GetMaAcceptedIndOk returns a tuple with the MaAcceptedInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N1N2MessageTransferReqData) GetMaAcceptedIndOk() (*bool, bool) {
	if o == nil || isNil(o.MaAcceptedInd) {
    return nil, false
	}
	return o.MaAcceptedInd, true
}

// HasMaAcceptedInd returns a boolean if a field has been set.
func (o *N1N2MessageTransferReqData) HasMaAcceptedInd() bool {
	if o != nil && !isNil(o.MaAcceptedInd) {
		return true
	}

	return false
}

// SetMaAcceptedInd gets a reference to the given bool and assigns it to the MaAcceptedInd field.
func (o *N1N2MessageTransferReqData) SetMaAcceptedInd(v bool) {
	o.MaAcceptedInd = &v
}

// GetExtBufSupport returns the ExtBufSupport field value if set, zero value otherwise.
func (o *N1N2MessageTransferReqData) GetExtBufSupport() bool {
	if o == nil || isNil(o.ExtBufSupport) {
		var ret bool
		return ret
	}
	return *o.ExtBufSupport
}

// GetExtBufSupportOk returns a tuple with the ExtBufSupport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N1N2MessageTransferReqData) GetExtBufSupportOk() (*bool, bool) {
	if o == nil || isNil(o.ExtBufSupport) {
    return nil, false
	}
	return o.ExtBufSupport, true
}

// HasExtBufSupport returns a boolean if a field has been set.
func (o *N1N2MessageTransferReqData) HasExtBufSupport() bool {
	if o != nil && !isNil(o.ExtBufSupport) {
		return true
	}

	return false
}

// SetExtBufSupport gets a reference to the given bool and assigns it to the ExtBufSupport field.
func (o *N1N2MessageTransferReqData) SetExtBufSupport(v bool) {
	o.ExtBufSupport = &v
}

// GetTargetAccess returns the TargetAccess field value if set, zero value otherwise.
func (o *N1N2MessageTransferReqData) GetTargetAccess() AccessType {
	if o == nil || isNil(o.TargetAccess) {
		var ret AccessType
		return ret
	}
	return *o.TargetAccess
}

// GetTargetAccessOk returns a tuple with the TargetAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N1N2MessageTransferReqData) GetTargetAccessOk() (*AccessType, bool) {
	if o == nil || isNil(o.TargetAccess) {
    return nil, false
	}
	return o.TargetAccess, true
}

// HasTargetAccess returns a boolean if a field has been set.
func (o *N1N2MessageTransferReqData) HasTargetAccess() bool {
	if o != nil && !isNil(o.TargetAccess) {
		return true
	}

	return false
}

// SetTargetAccess gets a reference to the given AccessType and assigns it to the TargetAccess field.
func (o *N1N2MessageTransferReqData) SetTargetAccess(v AccessType) {
	o.TargetAccess = &v
}

// GetNfId returns the NfId field value if set, zero value otherwise.
func (o *N1N2MessageTransferReqData) GetNfId() string {
	if o == nil || isNil(o.NfId) {
		var ret string
		return ret
	}
	return *o.NfId
}

// GetNfIdOk returns a tuple with the NfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N1N2MessageTransferReqData) GetNfIdOk() (*string, bool) {
	if o == nil || isNil(o.NfId) {
    return nil, false
	}
	return o.NfId, true
}

// HasNfId returns a boolean if a field has been set.
func (o *N1N2MessageTransferReqData) HasNfId() bool {
	if o != nil && !isNil(o.NfId) {
		return true
	}

	return false
}

// SetNfId gets a reference to the given string and assigns it to the NfId field.
func (o *N1N2MessageTransferReqData) SetNfId(v string) {
	o.NfId = &v
}

func (o N1N2MessageTransferReqData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.N1MessageContainer) {
		toSerialize["n1MessageContainer"] = o.N1MessageContainer
	}
	if !isNil(o.N2InfoContainer) {
		toSerialize["n2InfoContainer"] = o.N2InfoContainer
	}
	if !isNil(o.MtData) {
		toSerialize["mtData"] = o.MtData
	}
	if !isNil(o.SkipInd) {
		toSerialize["skipInd"] = o.SkipInd
	}
	if !isNil(o.LastMsgIndication) {
		toSerialize["lastMsgIndication"] = o.LastMsgIndication
	}
	if !isNil(o.PduSessionId) {
		toSerialize["pduSessionId"] = o.PduSessionId
	}
	if !isNil(o.LcsCorrelationId) {
		toSerialize["lcsCorrelationId"] = o.LcsCorrelationId
	}
	if !isNil(o.Ppi) {
		toSerialize["ppi"] = o.Ppi
	}
	if !isNil(o.Arp) {
		toSerialize["arp"] = o.Arp
	}
	if !isNil(o.Var5qi) {
		toSerialize["5qi"] = o.Var5qi
	}
	if !isNil(o.N1n2FailureTxfNotifURI) {
		toSerialize["n1n2FailureTxfNotifURI"] = o.N1n2FailureTxfNotifURI
	}
	if !isNil(o.SmfReallocationInd) {
		toSerialize["smfReallocationInd"] = o.SmfReallocationInd
	}
	if !isNil(o.AreaOfValidity) {
		toSerialize["areaOfValidity"] = o.AreaOfValidity
	}
	if !isNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	if !isNil(o.OldGuami) {
		toSerialize["oldGuami"] = o.OldGuami
	}
	if !isNil(o.MaAcceptedInd) {
		toSerialize["maAcceptedInd"] = o.MaAcceptedInd
	}
	if !isNil(o.ExtBufSupport) {
		toSerialize["extBufSupport"] = o.ExtBufSupport
	}
	if !isNil(o.TargetAccess) {
		toSerialize["targetAccess"] = o.TargetAccess
	}
	if !isNil(o.NfId) {
		toSerialize["nfId"] = o.NfId
	}
	return json.Marshal(toSerialize)
}

type NullableN1N2MessageTransferReqData struct {
	value *N1N2MessageTransferReqData
	isSet bool
}

func (v NullableN1N2MessageTransferReqData) Get() *N1N2MessageTransferReqData {
	return v.value
}

func (v *NullableN1N2MessageTransferReqData) Set(val *N1N2MessageTransferReqData) {
	v.value = val
	v.isSet = true
}

func (v NullableN1N2MessageTransferReqData) IsSet() bool {
	return v.isSet
}

func (v *NullableN1N2MessageTransferReqData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableN1N2MessageTransferReqData(val *N1N2MessageTransferReqData) *NullableN1N2MessageTransferReqData {
	return &NullableN1N2MessageTransferReqData{value: val, isSet: true}
}

func (v NullableN1N2MessageTransferReqData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableN1N2MessageTransferReqData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

