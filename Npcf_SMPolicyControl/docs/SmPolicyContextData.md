# SmPolicyContextData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccNetChId** | Pointer to [**AccNetChId**](AccNetChId.md) |  | [optional] 
**ChargEntityAddr** | Pointer to [**AccNetChargingAddress**](AccNetChargingAddress.md) |  | [optional] 
**Gpsi** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**Supi** | **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | 
**InvalidSupi** | Pointer to **bool** | When this attribute is included and set to true, it indicates that the supi attribute contains an invalid value.This attribute shall be present if the SUPI is not available in the SMF or the SUPI is unauthenticated. When present it shall be set to true for an invalid SUPI and false (default) for a valid SUPI.  | [optional] 
**InterGrpIds** | Pointer to **[]string** |  | [optional] 
**PduSessionId** | **int32** | Unsigned integer identifying a PDU session, within the range 0 to 255, as specified in  clause 11.2.3.1b, bits 1 to 8, of 3GPP TS 24.007. If the PDU Session ID is allocated by the  Core Network for UEs not supporting N1 mode, reserved range 64 to 95 is used. PDU Session ID  within the reserved range is only visible in the Core Network.   | 
**PduSessionType** | [**PduSessionType**](PduSessionType.md) |  | 
**Chargingcharacteristics** | Pointer to **string** |  | [optional] 
**Dnn** | **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | 
**DnnSelMode** | Pointer to [**DnnSelectionMode**](DnnSelectionMode.md) |  | [optional] 
**NotificationUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**AccessType** | Pointer to [**AccessType**](AccessType.md) |  | [optional] 
**RatType** | Pointer to [**RatType**](RatType.md) |  | [optional] 
**AddAccessInfo** | Pointer to [**AdditionalAccessInfo**](AdditionalAccessInfo.md) |  | [optional] 
**ServingNetwork** | Pointer to [**PlmnIdNid**](PlmnIdNid.md) |  | [optional] 
**UserLocationInfo** | Pointer to [**UserLocation**](UserLocation.md) |  | [optional] 
**UeTimeZone** | Pointer to **string** | String with format \&quot;time-numoffset\&quot; optionally appended by \&quot;daylightSavingTime\&quot;, where  - \&quot;time-numoffset\&quot; shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \&quot;daylightSavingTime\&quot; shall represent the adjustment that has been made and shall be    encoded as \&quot;+1\&quot; or \&quot;+2\&quot; for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time.  | [optional] 
**Pei** | Pointer to **string** | String representing a Permanent Equipment Identifier that may contain - an IMEI or IMEISV, as  specified in clause 6.2 of 3GPP TS 23.003; a MAC address for a 5G-RG or FN-RG via  wireline  access, with an indication that this address cannot be trusted for regulatory purpose if this  address cannot be used as an Equipment Identifier of the FN-RG, as specified in clause 4.7.7  of 3GPP TS23.316. Examples are imei-012345678901234 or imeisv-0123456789012345.   | [optional] 
**Ipv4Address** | Pointer to **string** | String identifying a IPv4 address formatted in the &#39;dotted decimal&#39; notation as defined in RFC 1166.  | [optional] 
**Ipv6AddressPrefix** | Pointer to [**Ipv6Prefix**](Ipv6Prefix.md) |  | [optional] 
**IpDomain** | Pointer to **string** | Indicates the IPv4 address domain | [optional] 
**SubsSessAmbr** | Pointer to [**Ambr**](Ambr.md) |  | [optional] 
**AuthProfIndex** | Pointer to **string** | Indicates the DN-AAA authorization profile index | [optional] 
**SubsDefQos** | Pointer to [**SubscribedDefaultQos**](SubscribedDefaultQos.md) |  | [optional] 
**VplmnQos** | Pointer to [**VplmnQos**](VplmnQos.md) |  | [optional] 
**NumOfPackFilter** | Pointer to **int32** | Contains the number of supported packet filter for signalled QoS rules. | [optional] 
**Online** | Pointer to **bool** | If it is included and set to true, the online charging is applied to the PDU session. | [optional] 
**Offline** | Pointer to **bool** | If it is included and set to true, the offline charging is applied to the PDU session. | [optional] 
**Var3gppPsDataOffStatus** | Pointer to **bool** | If it is included and set to true, the 3GPP PS Data Off is activated by the UE. | [optional] 
**RefQosIndication** | Pointer to **bool** | If it is included and set to true, the reflective QoS is supported by the UE. | [optional] 
**TraceReq** | Pointer to [**NullableTraceData**](TraceData.md) |  | [optional] 
**SliceInfo** | [**Snssai**](Snssai.md) |  | 
**QosFlowUsage** | Pointer to [**QosFlowUsage**](QosFlowUsage.md) |  | [optional] 
**ServNfId** | Pointer to [**ServingNfIdentity**](ServingNfIdentity.md) |  | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**SmfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**RecoveryTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**MaPduInd** | Pointer to [**MaPduIndication**](MaPduIndication.md) |  | [optional] 
**AtsssCapab** | Pointer to [**AtsssCapability**](AtsssCapability.md) |  | [optional] 
**Ipv4FrameRouteList** | Pointer to **[]string** |  | [optional] 
**Ipv6FrameRouteList** | Pointer to [**[]Ipv6Prefix**](Ipv6Prefix.md) |  | [optional] 
**SatBackhaulCategory** | Pointer to [**SatelliteBackhaulCategory**](SatelliteBackhaulCategory.md) |  | [optional] 
**PcfUeInfo** | Pointer to [**NullablePcfUeCallbackInfo**](PcfUeCallbackInfo.md) |  | [optional] 
**PvsInfo** | Pointer to [**[]ServerAddressingInfo**](ServerAddressingInfo.md) |  | [optional] 
**OnboardInd** | Pointer to **bool** | If it is included and set to true, it indicates that the PDU session is used for UE Onboarding. | [optional] 
**NwdafDatas** | Pointer to [**[]NwdafData**](NwdafData.md) |  | [optional] 

## Methods

### NewSmPolicyContextData

`func NewSmPolicyContextData(supi string, pduSessionId int32, pduSessionType PduSessionType, dnn string, notificationUri string, sliceInfo Snssai, ) *SmPolicyContextData`

NewSmPolicyContextData instantiates a new SmPolicyContextData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmPolicyContextDataWithDefaults

`func NewSmPolicyContextDataWithDefaults() *SmPolicyContextData`

NewSmPolicyContextDataWithDefaults instantiates a new SmPolicyContextData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccNetChId

`func (o *SmPolicyContextData) GetAccNetChId() AccNetChId`

GetAccNetChId returns the AccNetChId field if non-nil, zero value otherwise.

### GetAccNetChIdOk

`func (o *SmPolicyContextData) GetAccNetChIdOk() (*AccNetChId, bool)`

GetAccNetChIdOk returns a tuple with the AccNetChId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccNetChId

`func (o *SmPolicyContextData) SetAccNetChId(v AccNetChId)`

SetAccNetChId sets AccNetChId field to given value.

### HasAccNetChId

`func (o *SmPolicyContextData) HasAccNetChId() bool`

HasAccNetChId returns a boolean if a field has been set.

### GetChargEntityAddr

`func (o *SmPolicyContextData) GetChargEntityAddr() AccNetChargingAddress`

GetChargEntityAddr returns the ChargEntityAddr field if non-nil, zero value otherwise.

### GetChargEntityAddrOk

`func (o *SmPolicyContextData) GetChargEntityAddrOk() (*AccNetChargingAddress, bool)`

GetChargEntityAddrOk returns a tuple with the ChargEntityAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargEntityAddr

`func (o *SmPolicyContextData) SetChargEntityAddr(v AccNetChargingAddress)`

SetChargEntityAddr sets ChargEntityAddr field to given value.

### HasChargEntityAddr

`func (o *SmPolicyContextData) HasChargEntityAddr() bool`

HasChargEntityAddr returns a boolean if a field has been set.

### GetGpsi

`func (o *SmPolicyContextData) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *SmPolicyContextData) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *SmPolicyContextData) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *SmPolicyContextData) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetSupi

`func (o *SmPolicyContextData) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *SmPolicyContextData) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *SmPolicyContextData) SetSupi(v string)`

SetSupi sets Supi field to given value.


### GetInvalidSupi

`func (o *SmPolicyContextData) GetInvalidSupi() bool`

GetInvalidSupi returns the InvalidSupi field if non-nil, zero value otherwise.

### GetInvalidSupiOk

`func (o *SmPolicyContextData) GetInvalidSupiOk() (*bool, bool)`

GetInvalidSupiOk returns a tuple with the InvalidSupi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidSupi

`func (o *SmPolicyContextData) SetInvalidSupi(v bool)`

SetInvalidSupi sets InvalidSupi field to given value.

### HasInvalidSupi

`func (o *SmPolicyContextData) HasInvalidSupi() bool`

HasInvalidSupi returns a boolean if a field has been set.

### GetInterGrpIds

`func (o *SmPolicyContextData) GetInterGrpIds() []string`

GetInterGrpIds returns the InterGrpIds field if non-nil, zero value otherwise.

### GetInterGrpIdsOk

`func (o *SmPolicyContextData) GetInterGrpIdsOk() (*[]string, bool)`

GetInterGrpIdsOk returns a tuple with the InterGrpIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterGrpIds

`func (o *SmPolicyContextData) SetInterGrpIds(v []string)`

SetInterGrpIds sets InterGrpIds field to given value.

### HasInterGrpIds

`func (o *SmPolicyContextData) HasInterGrpIds() bool`

HasInterGrpIds returns a boolean if a field has been set.

### GetPduSessionId

`func (o *SmPolicyContextData) GetPduSessionId() int32`

GetPduSessionId returns the PduSessionId field if non-nil, zero value otherwise.

### GetPduSessionIdOk

`func (o *SmPolicyContextData) GetPduSessionIdOk() (*int32, bool)`

GetPduSessionIdOk returns a tuple with the PduSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionId

`func (o *SmPolicyContextData) SetPduSessionId(v int32)`

SetPduSessionId sets PduSessionId field to given value.


### GetPduSessionType

`func (o *SmPolicyContextData) GetPduSessionType() PduSessionType`

GetPduSessionType returns the PduSessionType field if non-nil, zero value otherwise.

### GetPduSessionTypeOk

`func (o *SmPolicyContextData) GetPduSessionTypeOk() (*PduSessionType, bool)`

GetPduSessionTypeOk returns a tuple with the PduSessionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionType

`func (o *SmPolicyContextData) SetPduSessionType(v PduSessionType)`

SetPduSessionType sets PduSessionType field to given value.


### GetChargingcharacteristics

`func (o *SmPolicyContextData) GetChargingcharacteristics() string`

GetChargingcharacteristics returns the Chargingcharacteristics field if non-nil, zero value otherwise.

### GetChargingcharacteristicsOk

`func (o *SmPolicyContextData) GetChargingcharacteristicsOk() (*string, bool)`

GetChargingcharacteristicsOk returns a tuple with the Chargingcharacteristics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargingcharacteristics

`func (o *SmPolicyContextData) SetChargingcharacteristics(v string)`

SetChargingcharacteristics sets Chargingcharacteristics field to given value.

### HasChargingcharacteristics

`func (o *SmPolicyContextData) HasChargingcharacteristics() bool`

HasChargingcharacteristics returns a boolean if a field has been set.

### GetDnn

`func (o *SmPolicyContextData) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *SmPolicyContextData) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *SmPolicyContextData) SetDnn(v string)`

SetDnn sets Dnn field to given value.


### GetDnnSelMode

`func (o *SmPolicyContextData) GetDnnSelMode() DnnSelectionMode`

GetDnnSelMode returns the DnnSelMode field if non-nil, zero value otherwise.

### GetDnnSelModeOk

`func (o *SmPolicyContextData) GetDnnSelModeOk() (*DnnSelectionMode, bool)`

GetDnnSelModeOk returns a tuple with the DnnSelMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnnSelMode

`func (o *SmPolicyContextData) SetDnnSelMode(v DnnSelectionMode)`

SetDnnSelMode sets DnnSelMode field to given value.

### HasDnnSelMode

`func (o *SmPolicyContextData) HasDnnSelMode() bool`

HasDnnSelMode returns a boolean if a field has been set.

### GetNotificationUri

`func (o *SmPolicyContextData) GetNotificationUri() string`

GetNotificationUri returns the NotificationUri field if non-nil, zero value otherwise.

### GetNotificationUriOk

`func (o *SmPolicyContextData) GetNotificationUriOk() (*string, bool)`

GetNotificationUriOk returns a tuple with the NotificationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationUri

`func (o *SmPolicyContextData) SetNotificationUri(v string)`

SetNotificationUri sets NotificationUri field to given value.


### GetAccessType

`func (o *SmPolicyContextData) GetAccessType() AccessType`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *SmPolicyContextData) GetAccessTypeOk() (*AccessType, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *SmPolicyContextData) SetAccessType(v AccessType)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *SmPolicyContextData) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetRatType

`func (o *SmPolicyContextData) GetRatType() RatType`

GetRatType returns the RatType field if non-nil, zero value otherwise.

### GetRatTypeOk

`func (o *SmPolicyContextData) GetRatTypeOk() (*RatType, bool)`

GetRatTypeOk returns a tuple with the RatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatType

`func (o *SmPolicyContextData) SetRatType(v RatType)`

SetRatType sets RatType field to given value.

### HasRatType

`func (o *SmPolicyContextData) HasRatType() bool`

HasRatType returns a boolean if a field has been set.

### GetAddAccessInfo

`func (o *SmPolicyContextData) GetAddAccessInfo() AdditionalAccessInfo`

GetAddAccessInfo returns the AddAccessInfo field if non-nil, zero value otherwise.

### GetAddAccessInfoOk

`func (o *SmPolicyContextData) GetAddAccessInfoOk() (*AdditionalAccessInfo, bool)`

GetAddAccessInfoOk returns a tuple with the AddAccessInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddAccessInfo

`func (o *SmPolicyContextData) SetAddAccessInfo(v AdditionalAccessInfo)`

SetAddAccessInfo sets AddAccessInfo field to given value.

### HasAddAccessInfo

`func (o *SmPolicyContextData) HasAddAccessInfo() bool`

HasAddAccessInfo returns a boolean if a field has been set.

### GetServingNetwork

`func (o *SmPolicyContextData) GetServingNetwork() PlmnIdNid`

GetServingNetwork returns the ServingNetwork field if non-nil, zero value otherwise.

### GetServingNetworkOk

`func (o *SmPolicyContextData) GetServingNetworkOk() (*PlmnIdNid, bool)`

GetServingNetworkOk returns a tuple with the ServingNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingNetwork

`func (o *SmPolicyContextData) SetServingNetwork(v PlmnIdNid)`

SetServingNetwork sets ServingNetwork field to given value.

### HasServingNetwork

`func (o *SmPolicyContextData) HasServingNetwork() bool`

HasServingNetwork returns a boolean if a field has been set.

### GetUserLocationInfo

`func (o *SmPolicyContextData) GetUserLocationInfo() UserLocation`

GetUserLocationInfo returns the UserLocationInfo field if non-nil, zero value otherwise.

### GetUserLocationInfoOk

`func (o *SmPolicyContextData) GetUserLocationInfoOk() (*UserLocation, bool)`

GetUserLocationInfoOk returns a tuple with the UserLocationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLocationInfo

`func (o *SmPolicyContextData) SetUserLocationInfo(v UserLocation)`

SetUserLocationInfo sets UserLocationInfo field to given value.

### HasUserLocationInfo

`func (o *SmPolicyContextData) HasUserLocationInfo() bool`

HasUserLocationInfo returns a boolean if a field has been set.

### GetUeTimeZone

`func (o *SmPolicyContextData) GetUeTimeZone() string`

GetUeTimeZone returns the UeTimeZone field if non-nil, zero value otherwise.

### GetUeTimeZoneOk

`func (o *SmPolicyContextData) GetUeTimeZoneOk() (*string, bool)`

GetUeTimeZoneOk returns a tuple with the UeTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeTimeZone

`func (o *SmPolicyContextData) SetUeTimeZone(v string)`

SetUeTimeZone sets UeTimeZone field to given value.

### HasUeTimeZone

`func (o *SmPolicyContextData) HasUeTimeZone() bool`

HasUeTimeZone returns a boolean if a field has been set.

### GetPei

`func (o *SmPolicyContextData) GetPei() string`

GetPei returns the Pei field if non-nil, zero value otherwise.

### GetPeiOk

`func (o *SmPolicyContextData) GetPeiOk() (*string, bool)`

GetPeiOk returns a tuple with the Pei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPei

`func (o *SmPolicyContextData) SetPei(v string)`

SetPei sets Pei field to given value.

### HasPei

`func (o *SmPolicyContextData) HasPei() bool`

HasPei returns a boolean if a field has been set.

### GetIpv4Address

`func (o *SmPolicyContextData) GetIpv4Address() string`

GetIpv4Address returns the Ipv4Address field if non-nil, zero value otherwise.

### GetIpv4AddressOk

`func (o *SmPolicyContextData) GetIpv4AddressOk() (*string, bool)`

GetIpv4AddressOk returns a tuple with the Ipv4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Address

`func (o *SmPolicyContextData) SetIpv4Address(v string)`

SetIpv4Address sets Ipv4Address field to given value.

### HasIpv4Address

`func (o *SmPolicyContextData) HasIpv4Address() bool`

HasIpv4Address returns a boolean if a field has been set.

### GetIpv6AddressPrefix

`func (o *SmPolicyContextData) GetIpv6AddressPrefix() Ipv6Prefix`

GetIpv6AddressPrefix returns the Ipv6AddressPrefix field if non-nil, zero value otherwise.

### GetIpv6AddressPrefixOk

`func (o *SmPolicyContextData) GetIpv6AddressPrefixOk() (*Ipv6Prefix, bool)`

GetIpv6AddressPrefixOk returns a tuple with the Ipv6AddressPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6AddressPrefix

`func (o *SmPolicyContextData) SetIpv6AddressPrefix(v Ipv6Prefix)`

SetIpv6AddressPrefix sets Ipv6AddressPrefix field to given value.

### HasIpv6AddressPrefix

`func (o *SmPolicyContextData) HasIpv6AddressPrefix() bool`

HasIpv6AddressPrefix returns a boolean if a field has been set.

### GetIpDomain

`func (o *SmPolicyContextData) GetIpDomain() string`

GetIpDomain returns the IpDomain field if non-nil, zero value otherwise.

### GetIpDomainOk

`func (o *SmPolicyContextData) GetIpDomainOk() (*string, bool)`

GetIpDomainOk returns a tuple with the IpDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpDomain

`func (o *SmPolicyContextData) SetIpDomain(v string)`

SetIpDomain sets IpDomain field to given value.

### HasIpDomain

`func (o *SmPolicyContextData) HasIpDomain() bool`

HasIpDomain returns a boolean if a field has been set.

### GetSubsSessAmbr

`func (o *SmPolicyContextData) GetSubsSessAmbr() Ambr`

GetSubsSessAmbr returns the SubsSessAmbr field if non-nil, zero value otherwise.

### GetSubsSessAmbrOk

`func (o *SmPolicyContextData) GetSubsSessAmbrOk() (*Ambr, bool)`

GetSubsSessAmbrOk returns a tuple with the SubsSessAmbr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsSessAmbr

`func (o *SmPolicyContextData) SetSubsSessAmbr(v Ambr)`

SetSubsSessAmbr sets SubsSessAmbr field to given value.

### HasSubsSessAmbr

`func (o *SmPolicyContextData) HasSubsSessAmbr() bool`

HasSubsSessAmbr returns a boolean if a field has been set.

### GetAuthProfIndex

`func (o *SmPolicyContextData) GetAuthProfIndex() string`

GetAuthProfIndex returns the AuthProfIndex field if non-nil, zero value otherwise.

### GetAuthProfIndexOk

`func (o *SmPolicyContextData) GetAuthProfIndexOk() (*string, bool)`

GetAuthProfIndexOk returns a tuple with the AuthProfIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthProfIndex

`func (o *SmPolicyContextData) SetAuthProfIndex(v string)`

SetAuthProfIndex sets AuthProfIndex field to given value.

### HasAuthProfIndex

`func (o *SmPolicyContextData) HasAuthProfIndex() bool`

HasAuthProfIndex returns a boolean if a field has been set.

### GetSubsDefQos

`func (o *SmPolicyContextData) GetSubsDefQos() SubscribedDefaultQos`

GetSubsDefQos returns the SubsDefQos field if non-nil, zero value otherwise.

### GetSubsDefQosOk

`func (o *SmPolicyContextData) GetSubsDefQosOk() (*SubscribedDefaultQos, bool)`

GetSubsDefQosOk returns a tuple with the SubsDefQos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsDefQos

`func (o *SmPolicyContextData) SetSubsDefQos(v SubscribedDefaultQos)`

SetSubsDefQos sets SubsDefQos field to given value.

### HasSubsDefQos

`func (o *SmPolicyContextData) HasSubsDefQos() bool`

HasSubsDefQos returns a boolean if a field has been set.

### GetVplmnQos

`func (o *SmPolicyContextData) GetVplmnQos() VplmnQos`

GetVplmnQos returns the VplmnQos field if non-nil, zero value otherwise.

### GetVplmnQosOk

`func (o *SmPolicyContextData) GetVplmnQosOk() (*VplmnQos, bool)`

GetVplmnQosOk returns a tuple with the VplmnQos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVplmnQos

`func (o *SmPolicyContextData) SetVplmnQos(v VplmnQos)`

SetVplmnQos sets VplmnQos field to given value.

### HasVplmnQos

`func (o *SmPolicyContextData) HasVplmnQos() bool`

HasVplmnQos returns a boolean if a field has been set.

### GetNumOfPackFilter

`func (o *SmPolicyContextData) GetNumOfPackFilter() int32`

GetNumOfPackFilter returns the NumOfPackFilter field if non-nil, zero value otherwise.

### GetNumOfPackFilterOk

`func (o *SmPolicyContextData) GetNumOfPackFilterOk() (*int32, bool)`

GetNumOfPackFilterOk returns a tuple with the NumOfPackFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfPackFilter

`func (o *SmPolicyContextData) SetNumOfPackFilter(v int32)`

SetNumOfPackFilter sets NumOfPackFilter field to given value.

### HasNumOfPackFilter

`func (o *SmPolicyContextData) HasNumOfPackFilter() bool`

HasNumOfPackFilter returns a boolean if a field has been set.

### GetOnline

`func (o *SmPolicyContextData) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *SmPolicyContextData) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *SmPolicyContextData) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *SmPolicyContextData) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetOffline

`func (o *SmPolicyContextData) GetOffline() bool`

GetOffline returns the Offline field if non-nil, zero value otherwise.

### GetOfflineOk

`func (o *SmPolicyContextData) GetOfflineOk() (*bool, bool)`

GetOfflineOk returns a tuple with the Offline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffline

`func (o *SmPolicyContextData) SetOffline(v bool)`

SetOffline sets Offline field to given value.

### HasOffline

`func (o *SmPolicyContextData) HasOffline() bool`

HasOffline returns a boolean if a field has been set.

### GetVar3gppPsDataOffStatus

`func (o *SmPolicyContextData) GetVar3gppPsDataOffStatus() bool`

GetVar3gppPsDataOffStatus returns the Var3gppPsDataOffStatus field if non-nil, zero value otherwise.

### GetVar3gppPsDataOffStatusOk

`func (o *SmPolicyContextData) GetVar3gppPsDataOffStatusOk() (*bool, bool)`

GetVar3gppPsDataOffStatusOk returns a tuple with the Var3gppPsDataOffStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar3gppPsDataOffStatus

`func (o *SmPolicyContextData) SetVar3gppPsDataOffStatus(v bool)`

SetVar3gppPsDataOffStatus sets Var3gppPsDataOffStatus field to given value.

### HasVar3gppPsDataOffStatus

`func (o *SmPolicyContextData) HasVar3gppPsDataOffStatus() bool`

HasVar3gppPsDataOffStatus returns a boolean if a field has been set.

### GetRefQosIndication

`func (o *SmPolicyContextData) GetRefQosIndication() bool`

GetRefQosIndication returns the RefQosIndication field if non-nil, zero value otherwise.

### GetRefQosIndicationOk

`func (o *SmPolicyContextData) GetRefQosIndicationOk() (*bool, bool)`

GetRefQosIndicationOk returns a tuple with the RefQosIndication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefQosIndication

`func (o *SmPolicyContextData) SetRefQosIndication(v bool)`

SetRefQosIndication sets RefQosIndication field to given value.

### HasRefQosIndication

`func (o *SmPolicyContextData) HasRefQosIndication() bool`

HasRefQosIndication returns a boolean if a field has been set.

### GetTraceReq

`func (o *SmPolicyContextData) GetTraceReq() TraceData`

GetTraceReq returns the TraceReq field if non-nil, zero value otherwise.

### GetTraceReqOk

`func (o *SmPolicyContextData) GetTraceReqOk() (*TraceData, bool)`

GetTraceReqOk returns a tuple with the TraceReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceReq

`func (o *SmPolicyContextData) SetTraceReq(v TraceData)`

SetTraceReq sets TraceReq field to given value.

### HasTraceReq

`func (o *SmPolicyContextData) HasTraceReq() bool`

HasTraceReq returns a boolean if a field has been set.

### SetTraceReqNil

`func (o *SmPolicyContextData) SetTraceReqNil(b bool)`

 SetTraceReqNil sets the value for TraceReq to be an explicit nil

### UnsetTraceReq
`func (o *SmPolicyContextData) UnsetTraceReq()`

UnsetTraceReq ensures that no value is present for TraceReq, not even an explicit nil
### GetSliceInfo

`func (o *SmPolicyContextData) GetSliceInfo() Snssai`

GetSliceInfo returns the SliceInfo field if non-nil, zero value otherwise.

### GetSliceInfoOk

`func (o *SmPolicyContextData) GetSliceInfoOk() (*Snssai, bool)`

GetSliceInfoOk returns a tuple with the SliceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliceInfo

`func (o *SmPolicyContextData) SetSliceInfo(v Snssai)`

SetSliceInfo sets SliceInfo field to given value.


### GetQosFlowUsage

`func (o *SmPolicyContextData) GetQosFlowUsage() QosFlowUsage`

GetQosFlowUsage returns the QosFlowUsage field if non-nil, zero value otherwise.

### GetQosFlowUsageOk

`func (o *SmPolicyContextData) GetQosFlowUsageOk() (*QosFlowUsage, bool)`

GetQosFlowUsageOk returns a tuple with the QosFlowUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosFlowUsage

`func (o *SmPolicyContextData) SetQosFlowUsage(v QosFlowUsage)`

SetQosFlowUsage sets QosFlowUsage field to given value.

### HasQosFlowUsage

`func (o *SmPolicyContextData) HasQosFlowUsage() bool`

HasQosFlowUsage returns a boolean if a field has been set.

### GetServNfId

`func (o *SmPolicyContextData) GetServNfId() ServingNfIdentity`

GetServNfId returns the ServNfId field if non-nil, zero value otherwise.

### GetServNfIdOk

`func (o *SmPolicyContextData) GetServNfIdOk() (*ServingNfIdentity, bool)`

GetServNfIdOk returns a tuple with the ServNfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServNfId

`func (o *SmPolicyContextData) SetServNfId(v ServingNfIdentity)`

SetServNfId sets ServNfId field to given value.

### HasServNfId

`func (o *SmPolicyContextData) HasServNfId() bool`

HasServNfId returns a boolean if a field has been set.

### GetSuppFeat

`func (o *SmPolicyContextData) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *SmPolicyContextData) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *SmPolicyContextData) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *SmPolicyContextData) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.

### GetSmfId

`func (o *SmPolicyContextData) GetSmfId() string`

GetSmfId returns the SmfId field if non-nil, zero value otherwise.

### GetSmfIdOk

`func (o *SmPolicyContextData) GetSmfIdOk() (*string, bool)`

GetSmfIdOk returns a tuple with the SmfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfId

`func (o *SmPolicyContextData) SetSmfId(v string)`

SetSmfId sets SmfId field to given value.

### HasSmfId

`func (o *SmPolicyContextData) HasSmfId() bool`

HasSmfId returns a boolean if a field has been set.

### GetRecoveryTime

`func (o *SmPolicyContextData) GetRecoveryTime() time.Time`

GetRecoveryTime returns the RecoveryTime field if non-nil, zero value otherwise.

### GetRecoveryTimeOk

`func (o *SmPolicyContextData) GetRecoveryTimeOk() (*time.Time, bool)`

GetRecoveryTimeOk returns a tuple with the RecoveryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTime

`func (o *SmPolicyContextData) SetRecoveryTime(v time.Time)`

SetRecoveryTime sets RecoveryTime field to given value.

### HasRecoveryTime

`func (o *SmPolicyContextData) HasRecoveryTime() bool`

HasRecoveryTime returns a boolean if a field has been set.

### GetMaPduInd

`func (o *SmPolicyContextData) GetMaPduInd() MaPduIndication`

GetMaPduInd returns the MaPduInd field if non-nil, zero value otherwise.

### GetMaPduIndOk

`func (o *SmPolicyContextData) GetMaPduIndOk() (*MaPduIndication, bool)`

GetMaPduIndOk returns a tuple with the MaPduInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaPduInd

`func (o *SmPolicyContextData) SetMaPduInd(v MaPduIndication)`

SetMaPduInd sets MaPduInd field to given value.

### HasMaPduInd

`func (o *SmPolicyContextData) HasMaPduInd() bool`

HasMaPduInd returns a boolean if a field has been set.

### GetAtsssCapab

`func (o *SmPolicyContextData) GetAtsssCapab() AtsssCapability`

GetAtsssCapab returns the AtsssCapab field if non-nil, zero value otherwise.

### GetAtsssCapabOk

`func (o *SmPolicyContextData) GetAtsssCapabOk() (*AtsssCapability, bool)`

GetAtsssCapabOk returns a tuple with the AtsssCapab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtsssCapab

`func (o *SmPolicyContextData) SetAtsssCapab(v AtsssCapability)`

SetAtsssCapab sets AtsssCapab field to given value.

### HasAtsssCapab

`func (o *SmPolicyContextData) HasAtsssCapab() bool`

HasAtsssCapab returns a boolean if a field has been set.

### GetIpv4FrameRouteList

`func (o *SmPolicyContextData) GetIpv4FrameRouteList() []string`

GetIpv4FrameRouteList returns the Ipv4FrameRouteList field if non-nil, zero value otherwise.

### GetIpv4FrameRouteListOk

`func (o *SmPolicyContextData) GetIpv4FrameRouteListOk() (*[]string, bool)`

GetIpv4FrameRouteListOk returns a tuple with the Ipv4FrameRouteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4FrameRouteList

`func (o *SmPolicyContextData) SetIpv4FrameRouteList(v []string)`

SetIpv4FrameRouteList sets Ipv4FrameRouteList field to given value.

### HasIpv4FrameRouteList

`func (o *SmPolicyContextData) HasIpv4FrameRouteList() bool`

HasIpv4FrameRouteList returns a boolean if a field has been set.

### GetIpv6FrameRouteList

`func (o *SmPolicyContextData) GetIpv6FrameRouteList() []Ipv6Prefix`

GetIpv6FrameRouteList returns the Ipv6FrameRouteList field if non-nil, zero value otherwise.

### GetIpv6FrameRouteListOk

`func (o *SmPolicyContextData) GetIpv6FrameRouteListOk() (*[]Ipv6Prefix, bool)`

GetIpv6FrameRouteListOk returns a tuple with the Ipv6FrameRouteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6FrameRouteList

`func (o *SmPolicyContextData) SetIpv6FrameRouteList(v []Ipv6Prefix)`

SetIpv6FrameRouteList sets Ipv6FrameRouteList field to given value.

### HasIpv6FrameRouteList

`func (o *SmPolicyContextData) HasIpv6FrameRouteList() bool`

HasIpv6FrameRouteList returns a boolean if a field has been set.

### GetSatBackhaulCategory

`func (o *SmPolicyContextData) GetSatBackhaulCategory() SatelliteBackhaulCategory`

GetSatBackhaulCategory returns the SatBackhaulCategory field if non-nil, zero value otherwise.

### GetSatBackhaulCategoryOk

`func (o *SmPolicyContextData) GetSatBackhaulCategoryOk() (*SatelliteBackhaulCategory, bool)`

GetSatBackhaulCategoryOk returns a tuple with the SatBackhaulCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSatBackhaulCategory

`func (o *SmPolicyContextData) SetSatBackhaulCategory(v SatelliteBackhaulCategory)`

SetSatBackhaulCategory sets SatBackhaulCategory field to given value.

### HasSatBackhaulCategory

`func (o *SmPolicyContextData) HasSatBackhaulCategory() bool`

HasSatBackhaulCategory returns a boolean if a field has been set.

### GetPcfUeInfo

`func (o *SmPolicyContextData) GetPcfUeInfo() PcfUeCallbackInfo`

GetPcfUeInfo returns the PcfUeInfo field if non-nil, zero value otherwise.

### GetPcfUeInfoOk

`func (o *SmPolicyContextData) GetPcfUeInfoOk() (*PcfUeCallbackInfo, bool)`

GetPcfUeInfoOk returns a tuple with the PcfUeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfUeInfo

`func (o *SmPolicyContextData) SetPcfUeInfo(v PcfUeCallbackInfo)`

SetPcfUeInfo sets PcfUeInfo field to given value.

### HasPcfUeInfo

`func (o *SmPolicyContextData) HasPcfUeInfo() bool`

HasPcfUeInfo returns a boolean if a field has been set.

### SetPcfUeInfoNil

`func (o *SmPolicyContextData) SetPcfUeInfoNil(b bool)`

 SetPcfUeInfoNil sets the value for PcfUeInfo to be an explicit nil

### UnsetPcfUeInfo
`func (o *SmPolicyContextData) UnsetPcfUeInfo()`

UnsetPcfUeInfo ensures that no value is present for PcfUeInfo, not even an explicit nil
### GetPvsInfo

`func (o *SmPolicyContextData) GetPvsInfo() []ServerAddressingInfo`

GetPvsInfo returns the PvsInfo field if non-nil, zero value otherwise.

### GetPvsInfoOk

`func (o *SmPolicyContextData) GetPvsInfoOk() (*[]ServerAddressingInfo, bool)`

GetPvsInfoOk returns a tuple with the PvsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvsInfo

`func (o *SmPolicyContextData) SetPvsInfo(v []ServerAddressingInfo)`

SetPvsInfo sets PvsInfo field to given value.

### HasPvsInfo

`func (o *SmPolicyContextData) HasPvsInfo() bool`

HasPvsInfo returns a boolean if a field has been set.

### GetOnboardInd

`func (o *SmPolicyContextData) GetOnboardInd() bool`

GetOnboardInd returns the OnboardInd field if non-nil, zero value otherwise.

### GetOnboardIndOk

`func (o *SmPolicyContextData) GetOnboardIndOk() (*bool, bool)`

GetOnboardIndOk returns a tuple with the OnboardInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardInd

`func (o *SmPolicyContextData) SetOnboardInd(v bool)`

SetOnboardInd sets OnboardInd field to given value.

### HasOnboardInd

`func (o *SmPolicyContextData) HasOnboardInd() bool`

HasOnboardInd returns a boolean if a field has been set.

### GetNwdafDatas

`func (o *SmPolicyContextData) GetNwdafDatas() []NwdafData`

GetNwdafDatas returns the NwdafDatas field if non-nil, zero value otherwise.

### GetNwdafDatasOk

`func (o *SmPolicyContextData) GetNwdafDatasOk() (*[]NwdafData, bool)`

GetNwdafDatasOk returns a tuple with the NwdafDatas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwdafDatas

`func (o *SmPolicyContextData) SetNwdafDatas(v []NwdafData)`

SetNwdafDatas sets NwdafDatas field to given value.

### HasNwdafDatas

`func (o *SmPolicyContextData) HasNwdafDatas() bool`

HasNwdafDatas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


