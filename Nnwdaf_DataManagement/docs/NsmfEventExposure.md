# NsmfEventExposure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Supi** | Pointer to **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | [optional] 
**Gpsi** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**AnyUeInd** | Pointer to **bool** | Any UE indication. This IE shall be present if the event subscription is applicable to any UE. Default value \&quot;false\&quot; is used, if not present.  | [optional] 
**GroupId** | Pointer to **string** | String identifying a group of devices network internal globally unique ID which identifies a set of IMSIs, as specified in clause 19.9 of 3GPP TS 23.003.   | [optional] 
**PduSeId** | Pointer to **int32** | Unsigned integer identifying a PDU session, within the range 0 to 255, as specified in  clause 11.2.3.1b, bits 1 to 8, of 3GPP TS 24.007. If the PDU Session ID is allocated by the  Core Network for UEs not supporting N1 mode, reserved range 64 to 95 is used. PDU Session ID  within the reserved range is only visible in the Core Network.   | [optional] 
**Dnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**Snssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**SubId** | Pointer to **string** | Identifies an Individual SMF Notification Subscription. To enable that the value is used as part of a URI, the string shall only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention defined in 3GPP TS 29.501. In an OpenAPI schema, the format shall be designated as \&quot;SubId\&quot;.  | [optional] 
**NotifId** | **string** | Notification Correlation ID assigned by the NF service consumer. | 
**NotifUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**AltNotifIpv4Addrs** | Pointer to **[]string** | Alternate or backup IPv4 address(es) where to send Notifications. | [optional] 
**AltNotifIpv6Addrs** | Pointer to [**[]Ipv6Addr**](Ipv6Addr.md) | Alternate or backup IPv6 address(es) where to send Notifications. | [optional] 
**AltNotifFqdns** | Pointer to **[]string** | Alternate or backup FQDN(s) where to send Notifications. | [optional] 
**EventSubs** | [**[]EventSubscription1**](EventSubscription1.md) | Subscribed events | 
**EventNotifs** | Pointer to [**[]EventNotification1**](EventNotification1.md) |  | [optional] 
**ImmeRep** | Pointer to **bool** |  | [optional] 
**NotifMethod** | Pointer to [**NotificationMethod1**](NotificationMethod1.md) |  | [optional] 
**MaxReportNbr** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**Expiry** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**RepPeriod** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**Guami** | Pointer to [**Guami**](Guami.md) |  | [optional] 
**ServiveName** | Pointer to [**ServiceName**](ServiceName.md) |  | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**SampRatio** | Pointer to **int32** | Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.   | [optional] 
**PartitionCriteria** | Pointer to [**[]PartitioningCriteria**](PartitioningCriteria.md) | Criteria for partitioning the UEs before applying the sampling ratio. | [optional] 
**GrpRepTime** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**NotifFlag** | Pointer to [**NotificationFlag**](NotificationFlag.md) |  | [optional] 

## Methods

### NewNsmfEventExposure

`func NewNsmfEventExposure(notifId string, notifUri string, eventSubs []EventSubscription1, ) *NsmfEventExposure`

NewNsmfEventExposure instantiates a new NsmfEventExposure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNsmfEventExposureWithDefaults

`func NewNsmfEventExposureWithDefaults() *NsmfEventExposure`

NewNsmfEventExposureWithDefaults instantiates a new NsmfEventExposure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupi

`func (o *NsmfEventExposure) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *NsmfEventExposure) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *NsmfEventExposure) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *NsmfEventExposure) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetGpsi

`func (o *NsmfEventExposure) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *NsmfEventExposure) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *NsmfEventExposure) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *NsmfEventExposure) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetAnyUeInd

`func (o *NsmfEventExposure) GetAnyUeInd() bool`

GetAnyUeInd returns the AnyUeInd field if non-nil, zero value otherwise.

### GetAnyUeIndOk

`func (o *NsmfEventExposure) GetAnyUeIndOk() (*bool, bool)`

GetAnyUeIndOk returns a tuple with the AnyUeInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyUeInd

`func (o *NsmfEventExposure) SetAnyUeInd(v bool)`

SetAnyUeInd sets AnyUeInd field to given value.

### HasAnyUeInd

`func (o *NsmfEventExposure) HasAnyUeInd() bool`

HasAnyUeInd returns a boolean if a field has been set.

### GetGroupId

`func (o *NsmfEventExposure) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *NsmfEventExposure) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *NsmfEventExposure) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *NsmfEventExposure) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetPduSeId

`func (o *NsmfEventExposure) GetPduSeId() int32`

GetPduSeId returns the PduSeId field if non-nil, zero value otherwise.

### GetPduSeIdOk

`func (o *NsmfEventExposure) GetPduSeIdOk() (*int32, bool)`

GetPduSeIdOk returns a tuple with the PduSeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSeId

`func (o *NsmfEventExposure) SetPduSeId(v int32)`

SetPduSeId sets PduSeId field to given value.

### HasPduSeId

`func (o *NsmfEventExposure) HasPduSeId() bool`

HasPduSeId returns a boolean if a field has been set.

### GetDnn

`func (o *NsmfEventExposure) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *NsmfEventExposure) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *NsmfEventExposure) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *NsmfEventExposure) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetSnssai

`func (o *NsmfEventExposure) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *NsmfEventExposure) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *NsmfEventExposure) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.

### HasSnssai

`func (o *NsmfEventExposure) HasSnssai() bool`

HasSnssai returns a boolean if a field has been set.

### GetSubId

`func (o *NsmfEventExposure) GetSubId() string`

GetSubId returns the SubId field if non-nil, zero value otherwise.

### GetSubIdOk

`func (o *NsmfEventExposure) GetSubIdOk() (*string, bool)`

GetSubIdOk returns a tuple with the SubId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubId

`func (o *NsmfEventExposure) SetSubId(v string)`

SetSubId sets SubId field to given value.

### HasSubId

`func (o *NsmfEventExposure) HasSubId() bool`

HasSubId returns a boolean if a field has been set.

### GetNotifId

`func (o *NsmfEventExposure) GetNotifId() string`

GetNotifId returns the NotifId field if non-nil, zero value otherwise.

### GetNotifIdOk

`func (o *NsmfEventExposure) GetNotifIdOk() (*string, bool)`

GetNotifIdOk returns a tuple with the NotifId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifId

`func (o *NsmfEventExposure) SetNotifId(v string)`

SetNotifId sets NotifId field to given value.


### GetNotifUri

`func (o *NsmfEventExposure) GetNotifUri() string`

GetNotifUri returns the NotifUri field if non-nil, zero value otherwise.

### GetNotifUriOk

`func (o *NsmfEventExposure) GetNotifUriOk() (*string, bool)`

GetNotifUriOk returns a tuple with the NotifUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifUri

`func (o *NsmfEventExposure) SetNotifUri(v string)`

SetNotifUri sets NotifUri field to given value.


### GetAltNotifIpv4Addrs

`func (o *NsmfEventExposure) GetAltNotifIpv4Addrs() []string`

GetAltNotifIpv4Addrs returns the AltNotifIpv4Addrs field if non-nil, zero value otherwise.

### GetAltNotifIpv4AddrsOk

`func (o *NsmfEventExposure) GetAltNotifIpv4AddrsOk() (*[]string, bool)`

GetAltNotifIpv4AddrsOk returns a tuple with the AltNotifIpv4Addrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltNotifIpv4Addrs

`func (o *NsmfEventExposure) SetAltNotifIpv4Addrs(v []string)`

SetAltNotifIpv4Addrs sets AltNotifIpv4Addrs field to given value.

### HasAltNotifIpv4Addrs

`func (o *NsmfEventExposure) HasAltNotifIpv4Addrs() bool`

HasAltNotifIpv4Addrs returns a boolean if a field has been set.

### GetAltNotifIpv6Addrs

`func (o *NsmfEventExposure) GetAltNotifIpv6Addrs() []Ipv6Addr`

GetAltNotifIpv6Addrs returns the AltNotifIpv6Addrs field if non-nil, zero value otherwise.

### GetAltNotifIpv6AddrsOk

`func (o *NsmfEventExposure) GetAltNotifIpv6AddrsOk() (*[]Ipv6Addr, bool)`

GetAltNotifIpv6AddrsOk returns a tuple with the AltNotifIpv6Addrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltNotifIpv6Addrs

`func (o *NsmfEventExposure) SetAltNotifIpv6Addrs(v []Ipv6Addr)`

SetAltNotifIpv6Addrs sets AltNotifIpv6Addrs field to given value.

### HasAltNotifIpv6Addrs

`func (o *NsmfEventExposure) HasAltNotifIpv6Addrs() bool`

HasAltNotifIpv6Addrs returns a boolean if a field has been set.

### GetAltNotifFqdns

`func (o *NsmfEventExposure) GetAltNotifFqdns() []string`

GetAltNotifFqdns returns the AltNotifFqdns field if non-nil, zero value otherwise.

### GetAltNotifFqdnsOk

`func (o *NsmfEventExposure) GetAltNotifFqdnsOk() (*[]string, bool)`

GetAltNotifFqdnsOk returns a tuple with the AltNotifFqdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltNotifFqdns

`func (o *NsmfEventExposure) SetAltNotifFqdns(v []string)`

SetAltNotifFqdns sets AltNotifFqdns field to given value.

### HasAltNotifFqdns

`func (o *NsmfEventExposure) HasAltNotifFqdns() bool`

HasAltNotifFqdns returns a boolean if a field has been set.

### GetEventSubs

`func (o *NsmfEventExposure) GetEventSubs() []EventSubscription1`

GetEventSubs returns the EventSubs field if non-nil, zero value otherwise.

### GetEventSubsOk

`func (o *NsmfEventExposure) GetEventSubsOk() (*[]EventSubscription1, bool)`

GetEventSubsOk returns a tuple with the EventSubs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSubs

`func (o *NsmfEventExposure) SetEventSubs(v []EventSubscription1)`

SetEventSubs sets EventSubs field to given value.


### GetEventNotifs

`func (o *NsmfEventExposure) GetEventNotifs() []EventNotification1`

GetEventNotifs returns the EventNotifs field if non-nil, zero value otherwise.

### GetEventNotifsOk

`func (o *NsmfEventExposure) GetEventNotifsOk() (*[]EventNotification1, bool)`

GetEventNotifsOk returns a tuple with the EventNotifs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNotifs

`func (o *NsmfEventExposure) SetEventNotifs(v []EventNotification1)`

SetEventNotifs sets EventNotifs field to given value.

### HasEventNotifs

`func (o *NsmfEventExposure) HasEventNotifs() bool`

HasEventNotifs returns a boolean if a field has been set.

### GetImmeRep

`func (o *NsmfEventExposure) GetImmeRep() bool`

GetImmeRep returns the ImmeRep field if non-nil, zero value otherwise.

### GetImmeRepOk

`func (o *NsmfEventExposure) GetImmeRepOk() (*bool, bool)`

GetImmeRepOk returns a tuple with the ImmeRep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmeRep

`func (o *NsmfEventExposure) SetImmeRep(v bool)`

SetImmeRep sets ImmeRep field to given value.

### HasImmeRep

`func (o *NsmfEventExposure) HasImmeRep() bool`

HasImmeRep returns a boolean if a field has been set.

### GetNotifMethod

`func (o *NsmfEventExposure) GetNotifMethod() NotificationMethod1`

GetNotifMethod returns the NotifMethod field if non-nil, zero value otherwise.

### GetNotifMethodOk

`func (o *NsmfEventExposure) GetNotifMethodOk() (*NotificationMethod1, bool)`

GetNotifMethodOk returns a tuple with the NotifMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifMethod

`func (o *NsmfEventExposure) SetNotifMethod(v NotificationMethod1)`

SetNotifMethod sets NotifMethod field to given value.

### HasNotifMethod

`func (o *NsmfEventExposure) HasNotifMethod() bool`

HasNotifMethod returns a boolean if a field has been set.

### GetMaxReportNbr

`func (o *NsmfEventExposure) GetMaxReportNbr() int32`

GetMaxReportNbr returns the MaxReportNbr field if non-nil, zero value otherwise.

### GetMaxReportNbrOk

`func (o *NsmfEventExposure) GetMaxReportNbrOk() (*int32, bool)`

GetMaxReportNbrOk returns a tuple with the MaxReportNbr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReportNbr

`func (o *NsmfEventExposure) SetMaxReportNbr(v int32)`

SetMaxReportNbr sets MaxReportNbr field to given value.

### HasMaxReportNbr

`func (o *NsmfEventExposure) HasMaxReportNbr() bool`

HasMaxReportNbr returns a boolean if a field has been set.

### GetExpiry

`func (o *NsmfEventExposure) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *NsmfEventExposure) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *NsmfEventExposure) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *NsmfEventExposure) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetRepPeriod

`func (o *NsmfEventExposure) GetRepPeriod() int32`

GetRepPeriod returns the RepPeriod field if non-nil, zero value otherwise.

### GetRepPeriodOk

`func (o *NsmfEventExposure) GetRepPeriodOk() (*int32, bool)`

GetRepPeriodOk returns a tuple with the RepPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepPeriod

`func (o *NsmfEventExposure) SetRepPeriod(v int32)`

SetRepPeriod sets RepPeriod field to given value.

### HasRepPeriod

`func (o *NsmfEventExposure) HasRepPeriod() bool`

HasRepPeriod returns a boolean if a field has been set.

### GetGuami

`func (o *NsmfEventExposure) GetGuami() Guami`

GetGuami returns the Guami field if non-nil, zero value otherwise.

### GetGuamiOk

`func (o *NsmfEventExposure) GetGuamiOk() (*Guami, bool)`

GetGuamiOk returns a tuple with the Guami field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuami

`func (o *NsmfEventExposure) SetGuami(v Guami)`

SetGuami sets Guami field to given value.

### HasGuami

`func (o *NsmfEventExposure) HasGuami() bool`

HasGuami returns a boolean if a field has been set.

### GetServiveName

`func (o *NsmfEventExposure) GetServiveName() ServiceName`

GetServiveName returns the ServiveName field if non-nil, zero value otherwise.

### GetServiveNameOk

`func (o *NsmfEventExposure) GetServiveNameOk() (*ServiceName, bool)`

GetServiveNameOk returns a tuple with the ServiveName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiveName

`func (o *NsmfEventExposure) SetServiveName(v ServiceName)`

SetServiveName sets ServiveName field to given value.

### HasServiveName

`func (o *NsmfEventExposure) HasServiveName() bool`

HasServiveName returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *NsmfEventExposure) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *NsmfEventExposure) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *NsmfEventExposure) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *NsmfEventExposure) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetSampRatio

`func (o *NsmfEventExposure) GetSampRatio() int32`

GetSampRatio returns the SampRatio field if non-nil, zero value otherwise.

### GetSampRatioOk

`func (o *NsmfEventExposure) GetSampRatioOk() (*int32, bool)`

GetSampRatioOk returns a tuple with the SampRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampRatio

`func (o *NsmfEventExposure) SetSampRatio(v int32)`

SetSampRatio sets SampRatio field to given value.

### HasSampRatio

`func (o *NsmfEventExposure) HasSampRatio() bool`

HasSampRatio returns a boolean if a field has been set.

### GetPartitionCriteria

`func (o *NsmfEventExposure) GetPartitionCriteria() []PartitioningCriteria`

GetPartitionCriteria returns the PartitionCriteria field if non-nil, zero value otherwise.

### GetPartitionCriteriaOk

`func (o *NsmfEventExposure) GetPartitionCriteriaOk() (*[]PartitioningCriteria, bool)`

GetPartitionCriteriaOk returns a tuple with the PartitionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionCriteria

`func (o *NsmfEventExposure) SetPartitionCriteria(v []PartitioningCriteria)`

SetPartitionCriteria sets PartitionCriteria field to given value.

### HasPartitionCriteria

`func (o *NsmfEventExposure) HasPartitionCriteria() bool`

HasPartitionCriteria returns a boolean if a field has been set.

### GetGrpRepTime

`func (o *NsmfEventExposure) GetGrpRepTime() int32`

GetGrpRepTime returns the GrpRepTime field if non-nil, zero value otherwise.

### GetGrpRepTimeOk

`func (o *NsmfEventExposure) GetGrpRepTimeOk() (*int32, bool)`

GetGrpRepTimeOk returns a tuple with the GrpRepTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrpRepTime

`func (o *NsmfEventExposure) SetGrpRepTime(v int32)`

SetGrpRepTime sets GrpRepTime field to given value.

### HasGrpRepTime

`func (o *NsmfEventExposure) HasGrpRepTime() bool`

HasGrpRepTime returns a boolean if a field has been set.

### GetNotifFlag

`func (o *NsmfEventExposure) GetNotifFlag() NotificationFlag`

GetNotifFlag returns the NotifFlag field if non-nil, zero value otherwise.

### GetNotifFlagOk

`func (o *NsmfEventExposure) GetNotifFlagOk() (*NotificationFlag, bool)`

GetNotifFlagOk returns a tuple with the NotifFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifFlag

`func (o *NsmfEventExposure) SetNotifFlag(v NotificationFlag)`

SetNotifFlag sets NotifFlag field to given value.

### HasNotifFlag

`func (o *NsmfEventExposure) HasNotifFlag() bool`

HasNotifFlag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


