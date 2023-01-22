/*
3gpp-chargeable-party

API for Chargeable Party management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ChargeableParty

import (
	"encoding/json"
)

// ChargeableParty Represents the configuration of a chargeable party.
type ChargeableParty struct {
	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	Self *string `json:"self,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn *string `json:"dnn,omitempty"`
	Snssai *Snssai `json:"snssai,omitempty"`
	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	NotificationDestination string `json:"notificationDestination"`
	// Set to true by the SCS/AS to request the SCEF to send a test notification as defined in clause 5.2.5.3. Set to false or omitted otherwise.
	RequestTestNotification *bool `json:"requestTestNotification,omitempty"`
	WebsockNotifConfig *WebsockNotifConfig `json:"websockNotifConfig,omitempty"`
	// Identifies the external Application Identifier.
	ExterAppId *string `json:"exterAppId,omitempty"`
	// string identifying a Ipv4 address formatted in the \"dotted decimal\" notation as defined in IETF RFC 1166.
	Ipv4Addr *string `json:"ipv4Addr,omitempty"`
	IpDomain *string `json:"ipDomain,omitempty"`
	// string identifying a Ipv6 address formatted according to clause 4 in IETF RFC 5952. The mixed Ipv4 Ipv6 notation according to clause 5 of IETF RFC 5952 shall not be used.
	Ipv6Addr *string `json:"ipv6Addr,omitempty"`
	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042. 
	MacAddr *string `json:"macAddr,omitempty"`
	// Describes the application flows.
	FlowInfo []FlowInfo `json:"flowInfo,omitempty"`
	// Identifies Ethernet packet flows.
	EthFlowInfo []EthFlowDescription `json:"ethFlowInfo,omitempty"`
	SponsorInformation SponsorInformation `json:"sponsorInformation"`
	// Indicates whether the sponsoring data connectivity is enabled (true) or not (false). 
	SponsoringEnabled bool `json:"sponsoringEnabled"`
	// string identifying a BDT Reference ID as defined in clause 5.3.3 of 3GPP TS 29.154.
	ReferenceId *string `json:"referenceId,omitempty"`
	ServAuthInfo *ServAuthInfo `json:"servAuthInfo,omitempty"`
	UsageThreshold *UsageThreshold `json:"usageThreshold,omitempty"`
	// Represents the list of event(s) to which the SCS/AS requests to subscribe to.
	Events []Event `json:"events,omitempty"`
}

// NewChargeableParty instantiates a new ChargeableParty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChargeableParty(notificationDestination string, sponsorInformation SponsorInformation, sponsoringEnabled bool) *ChargeableParty {
	this := ChargeableParty{}
	this.NotificationDestination = notificationDestination
	this.SponsorInformation = sponsorInformation
	this.SponsoringEnabled = sponsoringEnabled
	return &this
}

// NewChargeablePartyWithDefaults instantiates a new ChargeableParty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChargeablePartyWithDefaults() *ChargeableParty {
	this := ChargeableParty{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *ChargeableParty) GetSelf() string {
	if o == nil || isNil(o.Self) {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeableParty) GetSelfOk() (*string, bool) {
	if o == nil || isNil(o.Self) {
    return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *ChargeableParty) HasSelf() bool {
	if o != nil && !isNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *ChargeableParty) SetSelf(v string) {
	o.Self = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *ChargeableParty) GetSupportedFeatures() string {
	if o == nil || isNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeableParty) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || isNil(o.SupportedFeatures) {
    return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *ChargeableParty) HasSupportedFeatures() bool {
	if o != nil && !isNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *ChargeableParty) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetDnn returns the Dnn field value if set, zero value otherwise.
func (o *ChargeableParty) GetDnn() string {
	if o == nil || isNil(o.Dnn) {
		var ret string
		return ret
	}
	return *o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeableParty) GetDnnOk() (*string, bool) {
	if o == nil || isNil(o.Dnn) {
    return nil, false
	}
	return o.Dnn, true
}

// HasDnn returns a boolean if a field has been set.
func (o *ChargeableParty) HasDnn() bool {
	if o != nil && !isNil(o.Dnn) {
		return true
	}

	return false
}

// SetDnn gets a reference to the given string and assigns it to the Dnn field.
func (o *ChargeableParty) SetDnn(v string) {
	o.Dnn = &v
}

// GetSnssai returns the Snssai field value if set, zero value otherwise.
func (o *ChargeableParty) GetSnssai() Snssai {
	if o == nil || isNil(o.Snssai) {
		var ret Snssai
		return ret
	}
	return *o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeableParty) GetSnssaiOk() (*Snssai, bool) {
	if o == nil || isNil(o.Snssai) {
    return nil, false
	}
	return o.Snssai, true
}

// HasSnssai returns a boolean if a field has been set.
func (o *ChargeableParty) HasSnssai() bool {
	if o != nil && !isNil(o.Snssai) {
		return true
	}

	return false
}

// SetSnssai gets a reference to the given Snssai and assigns it to the Snssai field.
func (o *ChargeableParty) SetSnssai(v Snssai) {
	o.Snssai = &v
}

// GetNotificationDestination returns the NotificationDestination field value
func (o *ChargeableParty) GetNotificationDestination() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotificationDestination
}

// GetNotificationDestinationOk returns a tuple with the NotificationDestination field value
// and a boolean to check if the value has been set.
func (o *ChargeableParty) GetNotificationDestinationOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NotificationDestination, true
}

// SetNotificationDestination sets field value
func (o *ChargeableParty) SetNotificationDestination(v string) {
	o.NotificationDestination = v
}

// GetRequestTestNotification returns the RequestTestNotification field value if set, zero value otherwise.
func (o *ChargeableParty) GetRequestTestNotification() bool {
	if o == nil || isNil(o.RequestTestNotification) {
		var ret bool
		return ret
	}
	return *o.RequestTestNotification
}

// GetRequestTestNotificationOk returns a tuple with the RequestTestNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeableParty) GetRequestTestNotificationOk() (*bool, bool) {
	if o == nil || isNil(o.RequestTestNotification) {
    return nil, false
	}
	return o.RequestTestNotification, true
}

// HasRequestTestNotification returns a boolean if a field has been set.
func (o *ChargeableParty) HasRequestTestNotification() bool {
	if o != nil && !isNil(o.RequestTestNotification) {
		return true
	}

	return false
}

// SetRequestTestNotification gets a reference to the given bool and assigns it to the RequestTestNotification field.
func (o *ChargeableParty) SetRequestTestNotification(v bool) {
	o.RequestTestNotification = &v
}

// GetWebsockNotifConfig returns the WebsockNotifConfig field value if set, zero value otherwise.
func (o *ChargeableParty) GetWebsockNotifConfig() WebsockNotifConfig {
	if o == nil || isNil(o.WebsockNotifConfig) {
		var ret WebsockNotifConfig
		return ret
	}
	return *o.WebsockNotifConfig
}

// GetWebsockNotifConfigOk returns a tuple with the WebsockNotifConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeableParty) GetWebsockNotifConfigOk() (*WebsockNotifConfig, bool) {
	if o == nil || isNil(o.WebsockNotifConfig) {
    return nil, false
	}
	return o.WebsockNotifConfig, true
}

// HasWebsockNotifConfig returns a boolean if a field has been set.
func (o *ChargeableParty) HasWebsockNotifConfig() bool {
	if o != nil && !isNil(o.WebsockNotifConfig) {
		return true
	}

	return false
}

// SetWebsockNotifConfig gets a reference to the given WebsockNotifConfig and assigns it to the WebsockNotifConfig field.
func (o *ChargeableParty) SetWebsockNotifConfig(v WebsockNotifConfig) {
	o.WebsockNotifConfig = &v
}

// GetExterAppId returns the ExterAppId field value if set, zero value otherwise.
func (o *ChargeableParty) GetExterAppId() string {
	if o == nil || isNil(o.ExterAppId) {
		var ret string
		return ret
	}
	return *o.ExterAppId
}

// GetExterAppIdOk returns a tuple with the ExterAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeableParty) GetExterAppIdOk() (*string, bool) {
	if o == nil || isNil(o.ExterAppId) {
    return nil, false
	}
	return o.ExterAppId, true
}

// HasExterAppId returns a boolean if a field has been set.
func (o *ChargeableParty) HasExterAppId() bool {
	if o != nil && !isNil(o.ExterAppId) {
		return true
	}

	return false
}

// SetExterAppId gets a reference to the given string and assigns it to the ExterAppId field.
func (o *ChargeableParty) SetExterAppId(v string) {
	o.ExterAppId = &v
}

// GetIpv4Addr returns the Ipv4Addr field value if set, zero value otherwise.
func (o *ChargeableParty) GetIpv4Addr() string {
	if o == nil || isNil(o.Ipv4Addr) {
		var ret string
		return ret
	}
	return *o.Ipv4Addr
}

// GetIpv4AddrOk returns a tuple with the Ipv4Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeableParty) GetIpv4AddrOk() (*string, bool) {
	if o == nil || isNil(o.Ipv4Addr) {
    return nil, false
	}
	return o.Ipv4Addr, true
}

// HasIpv4Addr returns a boolean if a field has been set.
func (o *ChargeableParty) HasIpv4Addr() bool {
	if o != nil && !isNil(o.Ipv4Addr) {
		return true
	}

	return false
}

// SetIpv4Addr gets a reference to the given string and assigns it to the Ipv4Addr field.
func (o *ChargeableParty) SetIpv4Addr(v string) {
	o.Ipv4Addr = &v
}

// GetIpDomain returns the IpDomain field value if set, zero value otherwise.
func (o *ChargeableParty) GetIpDomain() string {
	if o == nil || isNil(o.IpDomain) {
		var ret string
		return ret
	}
	return *o.IpDomain
}

// GetIpDomainOk returns a tuple with the IpDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeableParty) GetIpDomainOk() (*string, bool) {
	if o == nil || isNil(o.IpDomain) {
    return nil, false
	}
	return o.IpDomain, true
}

// HasIpDomain returns a boolean if a field has been set.
func (o *ChargeableParty) HasIpDomain() bool {
	if o != nil && !isNil(o.IpDomain) {
		return true
	}

	return false
}

// SetIpDomain gets a reference to the given string and assigns it to the IpDomain field.
func (o *ChargeableParty) SetIpDomain(v string) {
	o.IpDomain = &v
}

// GetIpv6Addr returns the Ipv6Addr field value if set, zero value otherwise.
func (o *ChargeableParty) GetIpv6Addr() string {
	if o == nil || isNil(o.Ipv6Addr) {
		var ret string
		return ret
	}
	return *o.Ipv6Addr
}

// GetIpv6AddrOk returns a tuple with the Ipv6Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeableParty) GetIpv6AddrOk() (*string, bool) {
	if o == nil || isNil(o.Ipv6Addr) {
    return nil, false
	}
	return o.Ipv6Addr, true
}

// HasIpv6Addr returns a boolean if a field has been set.
func (o *ChargeableParty) HasIpv6Addr() bool {
	if o != nil && !isNil(o.Ipv6Addr) {
		return true
	}

	return false
}

// SetIpv6Addr gets a reference to the given string and assigns it to the Ipv6Addr field.
func (o *ChargeableParty) SetIpv6Addr(v string) {
	o.Ipv6Addr = &v
}

// GetMacAddr returns the MacAddr field value if set, zero value otherwise.
func (o *ChargeableParty) GetMacAddr() string {
	if o == nil || isNil(o.MacAddr) {
		var ret string
		return ret
	}
	return *o.MacAddr
}

// GetMacAddrOk returns a tuple with the MacAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeableParty) GetMacAddrOk() (*string, bool) {
	if o == nil || isNil(o.MacAddr) {
    return nil, false
	}
	return o.MacAddr, true
}

// HasMacAddr returns a boolean if a field has been set.
func (o *ChargeableParty) HasMacAddr() bool {
	if o != nil && !isNil(o.MacAddr) {
		return true
	}

	return false
}

// SetMacAddr gets a reference to the given string and assigns it to the MacAddr field.
func (o *ChargeableParty) SetMacAddr(v string) {
	o.MacAddr = &v
}

// GetFlowInfo returns the FlowInfo field value if set, zero value otherwise.
func (o *ChargeableParty) GetFlowInfo() []FlowInfo {
	if o == nil || isNil(o.FlowInfo) {
		var ret []FlowInfo
		return ret
	}
	return o.FlowInfo
}

// GetFlowInfoOk returns a tuple with the FlowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeableParty) GetFlowInfoOk() ([]FlowInfo, bool) {
	if o == nil || isNil(o.FlowInfo) {
    return nil, false
	}
	return o.FlowInfo, true
}

// HasFlowInfo returns a boolean if a field has been set.
func (o *ChargeableParty) HasFlowInfo() bool {
	if o != nil && !isNil(o.FlowInfo) {
		return true
	}

	return false
}

// SetFlowInfo gets a reference to the given []FlowInfo and assigns it to the FlowInfo field.
func (o *ChargeableParty) SetFlowInfo(v []FlowInfo) {
	o.FlowInfo = v
}

// GetEthFlowInfo returns the EthFlowInfo field value if set, zero value otherwise.
func (o *ChargeableParty) GetEthFlowInfo() []EthFlowDescription {
	if o == nil || isNil(o.EthFlowInfo) {
		var ret []EthFlowDescription
		return ret
	}
	return o.EthFlowInfo
}

// GetEthFlowInfoOk returns a tuple with the EthFlowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeableParty) GetEthFlowInfoOk() ([]EthFlowDescription, bool) {
	if o == nil || isNil(o.EthFlowInfo) {
    return nil, false
	}
	return o.EthFlowInfo, true
}

// HasEthFlowInfo returns a boolean if a field has been set.
func (o *ChargeableParty) HasEthFlowInfo() bool {
	if o != nil && !isNil(o.EthFlowInfo) {
		return true
	}

	return false
}

// SetEthFlowInfo gets a reference to the given []EthFlowDescription and assigns it to the EthFlowInfo field.
func (o *ChargeableParty) SetEthFlowInfo(v []EthFlowDescription) {
	o.EthFlowInfo = v
}

// GetSponsorInformation returns the SponsorInformation field value
func (o *ChargeableParty) GetSponsorInformation() SponsorInformation {
	if o == nil {
		var ret SponsorInformation
		return ret
	}

	return o.SponsorInformation
}

// GetSponsorInformationOk returns a tuple with the SponsorInformation field value
// and a boolean to check if the value has been set.
func (o *ChargeableParty) GetSponsorInformationOk() (*SponsorInformation, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SponsorInformation, true
}

// SetSponsorInformation sets field value
func (o *ChargeableParty) SetSponsorInformation(v SponsorInformation) {
	o.SponsorInformation = v
}

// GetSponsoringEnabled returns the SponsoringEnabled field value
func (o *ChargeableParty) GetSponsoringEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SponsoringEnabled
}

// GetSponsoringEnabledOk returns a tuple with the SponsoringEnabled field value
// and a boolean to check if the value has been set.
func (o *ChargeableParty) GetSponsoringEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SponsoringEnabled, true
}

// SetSponsoringEnabled sets field value
func (o *ChargeableParty) SetSponsoringEnabled(v bool) {
	o.SponsoringEnabled = v
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise.
func (o *ChargeableParty) GetReferenceId() string {
	if o == nil || isNil(o.ReferenceId) {
		var ret string
		return ret
	}
	return *o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeableParty) GetReferenceIdOk() (*string, bool) {
	if o == nil || isNil(o.ReferenceId) {
    return nil, false
	}
	return o.ReferenceId, true
}

// HasReferenceId returns a boolean if a field has been set.
func (o *ChargeableParty) HasReferenceId() bool {
	if o != nil && !isNil(o.ReferenceId) {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given string and assigns it to the ReferenceId field.
func (o *ChargeableParty) SetReferenceId(v string) {
	o.ReferenceId = &v
}

// GetServAuthInfo returns the ServAuthInfo field value if set, zero value otherwise.
func (o *ChargeableParty) GetServAuthInfo() ServAuthInfo {
	if o == nil || isNil(o.ServAuthInfo) {
		var ret ServAuthInfo
		return ret
	}
	return *o.ServAuthInfo
}

// GetServAuthInfoOk returns a tuple with the ServAuthInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeableParty) GetServAuthInfoOk() (*ServAuthInfo, bool) {
	if o == nil || isNil(o.ServAuthInfo) {
    return nil, false
	}
	return o.ServAuthInfo, true
}

// HasServAuthInfo returns a boolean if a field has been set.
func (o *ChargeableParty) HasServAuthInfo() bool {
	if o != nil && !isNil(o.ServAuthInfo) {
		return true
	}

	return false
}

// SetServAuthInfo gets a reference to the given ServAuthInfo and assigns it to the ServAuthInfo field.
func (o *ChargeableParty) SetServAuthInfo(v ServAuthInfo) {
	o.ServAuthInfo = &v
}

// GetUsageThreshold returns the UsageThreshold field value if set, zero value otherwise.
func (o *ChargeableParty) GetUsageThreshold() UsageThreshold {
	if o == nil || isNil(o.UsageThreshold) {
		var ret UsageThreshold
		return ret
	}
	return *o.UsageThreshold
}

// GetUsageThresholdOk returns a tuple with the UsageThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeableParty) GetUsageThresholdOk() (*UsageThreshold, bool) {
	if o == nil || isNil(o.UsageThreshold) {
    return nil, false
	}
	return o.UsageThreshold, true
}

// HasUsageThreshold returns a boolean if a field has been set.
func (o *ChargeableParty) HasUsageThreshold() bool {
	if o != nil && !isNil(o.UsageThreshold) {
		return true
	}

	return false
}

// SetUsageThreshold gets a reference to the given UsageThreshold and assigns it to the UsageThreshold field.
func (o *ChargeableParty) SetUsageThreshold(v UsageThreshold) {
	o.UsageThreshold = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *ChargeableParty) GetEvents() []Event {
	if o == nil || isNil(o.Events) {
		var ret []Event
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeableParty) GetEventsOk() ([]Event, bool) {
	if o == nil || isNil(o.Events) {
    return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *ChargeableParty) HasEvents() bool {
	if o != nil && !isNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []Event and assigns it to the Events field.
func (o *ChargeableParty) SetEvents(v []Event) {
	o.Events = v
}

func (o ChargeableParty) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !isNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	if !isNil(o.Dnn) {
		toSerialize["dnn"] = o.Dnn
	}
	if !isNil(o.Snssai) {
		toSerialize["snssai"] = o.Snssai
	}
	if true {
		toSerialize["notificationDestination"] = o.NotificationDestination
	}
	if !isNil(o.RequestTestNotification) {
		toSerialize["requestTestNotification"] = o.RequestTestNotification
	}
	if !isNil(o.WebsockNotifConfig) {
		toSerialize["websockNotifConfig"] = o.WebsockNotifConfig
	}
	if !isNil(o.ExterAppId) {
		toSerialize["exterAppId"] = o.ExterAppId
	}
	if !isNil(o.Ipv4Addr) {
		toSerialize["ipv4Addr"] = o.Ipv4Addr
	}
	if !isNil(o.IpDomain) {
		toSerialize["ipDomain"] = o.IpDomain
	}
	if !isNil(o.Ipv6Addr) {
		toSerialize["ipv6Addr"] = o.Ipv6Addr
	}
	if !isNil(o.MacAddr) {
		toSerialize["macAddr"] = o.MacAddr
	}
	if !isNil(o.FlowInfo) {
		toSerialize["flowInfo"] = o.FlowInfo
	}
	if !isNil(o.EthFlowInfo) {
		toSerialize["ethFlowInfo"] = o.EthFlowInfo
	}
	if true {
		toSerialize["sponsorInformation"] = o.SponsorInformation
	}
	if true {
		toSerialize["sponsoringEnabled"] = o.SponsoringEnabled
	}
	if !isNil(o.ReferenceId) {
		toSerialize["referenceId"] = o.ReferenceId
	}
	if !isNil(o.ServAuthInfo) {
		toSerialize["servAuthInfo"] = o.ServAuthInfo
	}
	if !isNil(o.UsageThreshold) {
		toSerialize["usageThreshold"] = o.UsageThreshold
	}
	if !isNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableChargeableParty struct {
	value *ChargeableParty
	isSet bool
}

func (v NullableChargeableParty) Get() *ChargeableParty {
	return v.value
}

func (v *NullableChargeableParty) Set(val *ChargeableParty) {
	v.value = val
	v.isSet = true
}

func (v NullableChargeableParty) IsSet() bool {
	return v.isSet
}

func (v *NullableChargeableParty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChargeableParty(val *ChargeableParty) *NullableChargeableParty {
	return &NullableChargeableParty{value: val, isSet: true}
}

func (v NullableChargeableParty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChargeableParty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

