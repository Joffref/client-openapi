# UeContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Supi** | Pointer to **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | [optional] 
**SupiUnauthInd** | Pointer to **bool** |  | [optional] 
**GpsiList** | Pointer to **[]string** |  | [optional] 
**Pei** | Pointer to **string** | String representing a Permanent Equipment Identifier that may contain - an IMEI or IMEISV, as  specified in clause 6.2 of 3GPP TS 23.003; a MAC address for a 5G-RG or FN-RG via  wireline  access, with an indication that this address cannot be trusted for regulatory purpose if this  address cannot be used as an Equipment Identifier of the FN-RG, as specified in clause 4.7.7  of 3GPP TS23.316. Examples are imei-012345678901234 or imeisv-0123456789012345.   | [optional] 
**UdmGroupId** | Pointer to **string** | Identifier of a group of NFs. | [optional] 
**AusfGroupId** | Pointer to **string** | Identifier of a group of NFs. | [optional] 
**PcfGroupId** | Pointer to **string** | Identifier of a group of NFs. | [optional] 
**RoutingIndicator** | Pointer to **string** |  | [optional] 
**HNwPubKeyId** | Pointer to **int32** |  | [optional] 
**GroupList** | Pointer to **[]string** |  | [optional] 
**DrxParameter** | Pointer to **string** | string with format &#39;bytes&#39; as defined in OpenAPI | [optional] 
**SubRfsp** | Pointer to **int32** | Unsigned integer representing the \&quot;Subscriber Profile ID for RAT/Frequency Priority\&quot;  as specified in 3GPP TS 36.413.  | [optional] 
**UsedRfsp** | Pointer to **int32** | Unsigned integer representing the \&quot;Subscriber Profile ID for RAT/Frequency Priority\&quot;  as specified in 3GPP TS 36.413.  | [optional] 
**SubUeAmbr** | Pointer to [**Ambr**](Ambr.md) |  | [optional] 
**SubUeSliceMbrList** | Pointer to [**map[string]SliceMbr**](SliceMbr.md) | A map(list of key-value pairs) where Snssai serves as key. | [optional] 
**SmsfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**SeafData** | Pointer to [**SeafData**](SeafData.md) |  | [optional] 
**Var5gMmCapability** | Pointer to **string** | string with format &#39;bytes&#39; as defined in OpenAPI | [optional] 
**PcfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**PcfSetId** | Pointer to **string** | NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \&quot;set&lt;Set ID&gt;.&lt;nftype&gt;set.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.&lt;NFType&gt;set.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)  &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NFType&gt; encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters &lt;Set ID&gt; encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.   | [optional] 
**PcfAmpServiceSetId** | Pointer to **string** | NF Service Set Identifier (see clause 28.12 of 3GPP TS 23.003) formatted as the following  string \&quot;set&lt;Set ID&gt;.sn&lt;Service Name&gt;.nfi&lt;NF Instance ID&gt;.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.sn&lt;ServiceName&gt;.nfi&lt;NFInstanceID&gt;.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)   &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NID&gt; encoded as defined in clause 5.4.2 (\&quot;Nid\&quot; data type definition)  &lt;NFInstanceId&gt; encoded as defined in clause 5.3.2  &lt;ServiceName&gt; encoded as defined in 3GPP TS 29.510  &lt;Set ID&gt; encoded as a string of characters consisting of alphabetic    characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that shall end    with either an alphabetic character or a digit.  | [optional] 
**PcfUepServiceSetId** | Pointer to **string** | NF Service Set Identifier (see clause 28.12 of 3GPP TS 23.003) formatted as the following  string \&quot;set&lt;Set ID&gt;.sn&lt;Service Name&gt;.nfi&lt;NF Instance ID&gt;.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.sn&lt;ServiceName&gt;.nfi&lt;NFInstanceID&gt;.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)   &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NID&gt; encoded as defined in clause 5.4.2 (\&quot;Nid\&quot; data type definition)  &lt;NFInstanceId&gt; encoded as defined in clause 5.3.2  &lt;ServiceName&gt; encoded as defined in 3GPP TS 29.510  &lt;Set ID&gt; encoded as a string of characters consisting of alphabetic    characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that shall end    with either an alphabetic character or a digit.  | [optional] 
**PcfBinding** | Pointer to [**SbiBindingLevel**](SbiBindingLevel.md) |  | [optional] 
**PcfAmPolicyUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**AmPolicyReqTriggerList** | Pointer to [**[]PolicyReqTrigger**](PolicyReqTrigger.md) |  | [optional] 
**PcfUePolicyUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**UePolicyReqTriggerList** | Pointer to [**[]PolicyReqTrigger**](PolicyReqTrigger.md) |  | [optional] 
**HpcfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**HpcfSetId** | Pointer to **string** | NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \&quot;set&lt;Set ID&gt;.&lt;nftype&gt;set.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.&lt;NFType&gt;set.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)  &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NFType&gt; encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters &lt;Set ID&gt; encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.   | [optional] 
**RestrictedRatList** | Pointer to [**[]RatType**](RatType.md) |  | [optional] 
**ForbiddenAreaList** | Pointer to [**[]Area**](Area.md) |  | [optional] 
**ServiceAreaRestriction** | Pointer to [**ServiceAreaRestriction**](ServiceAreaRestriction.md) |  | [optional] 
**RestrictedCoreNwTypeList** | Pointer to [**[]CoreNetworkType**](CoreNetworkType.md) |  | [optional] 
**EventSubscriptionList** | Pointer to [**[]ExtAmfEventSubscription**](ExtAmfEventSubscription.md) |  | [optional] 
**MmContextList** | Pointer to [**[]MmContext**](MmContext.md) |  | [optional] 
**SessionContextList** | Pointer to [**[]PduSessionContext**](PduSessionContext.md) |  | [optional] 
**EpsInterworkingInfo** | Pointer to [**EpsInterworkingInfo**](EpsInterworkingInfo.md) |  | [optional] 
**TraceData** | Pointer to [**NullableTraceData**](TraceData.md) |  | [optional] 
**ServiceGapExpiryTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**StnSr** | Pointer to **string** | String representing the STN-SR as defined in clause 18.6 of 3GPP TS 23.003. | [optional] 
**CMsisdn** | Pointer to **string** | String representing the C-MSISDN as defined in clause 18.7 of 3GPP TS 23.003. | [optional] 
**MsClassmark2** | Pointer to **string** | string with format &#39;bytes&#39; as defined in OpenAPI | [optional] 
**SupportedCodecList** | Pointer to **[]string** |  | [optional] 
**SmallDataRateStatusInfos** | Pointer to [**[]SmallDataRateStatusInfo**](SmallDataRateStatusInfo.md) |  | [optional] 
**RestrictedPrimaryRatList** | Pointer to [**[]RatType**](RatType.md) |  | [optional] 
**RestrictedSecondaryRatList** | Pointer to [**[]RatType**](RatType.md) |  | [optional] 
**V2xContext** | Pointer to [**V2xContext**](V2xContext.md) |  | [optional] 
**LteCatMInd** | Pointer to **bool** |  | [optional] [default to false]
**RedCapInd** | Pointer to **bool** |  | [optional] [default to false]
**MoExpDataCounter** | Pointer to [**MoExpDataCounter**](MoExpDataCounter.md) |  | [optional] 
**CagData** | Pointer to [**CagData**](CagData.md) |  | [optional] 
**ManagementMdtInd** | Pointer to **bool** |  | [optional] [default to false]
**ImmediateMdtConf** | Pointer to [**ImmediateMdtConf**](ImmediateMdtConf.md) |  | [optional] 
**EcRestrictionDataWb** | Pointer to [**EcRestrictionDataWb**](EcRestrictionDataWb.md) |  | [optional] 
**EcRestrictionDataNb** | Pointer to **bool** |  | [optional] [default to false]
**IabOperationAllowed** | Pointer to **bool** |  | [optional] 
**ProseContext** | Pointer to [**ProseContext**](ProseContext.md) |  | [optional] 
**AnalyticsSubscriptionList** | Pointer to [**[]AnalyticsSubscription**](AnalyticsSubscription.md) |  | [optional] 
**PcfAmpBindingInfo** | Pointer to **string** |  | [optional] 
**PcfUepBindingInfo** | Pointer to **string** |  | [optional] 
**UsedServiceAreaRestriction** | Pointer to [**ServiceAreaRestriction**](ServiceAreaRestriction.md) |  | [optional] 
**PraInAmPolicy** | Pointer to [**map[string]PresenceInfo**](PresenceInfo.md) | A map(list of key-value pairs) where praId serves as key. | [optional] 
**PraInUePolicy** | Pointer to [**map[string]PresenceInfo**](PresenceInfo.md) | A map(list of key-value pairs) where praId serves as key. | [optional] 
**UpdpSubscriptionData** | Pointer to [**UpdpSubscriptionData**](UpdpSubscriptionData.md) |  | [optional] 
**SmPolicyNotifyPduList** | Pointer to [**[]PduSessionInfo**](PduSessionInfo.md) |  | [optional] 
**PcfUeCallbackInfo** | Pointer to [**NullablePcfUeCallbackInfo**](PcfUeCallbackInfo.md) |  | [optional] 
**UePositioningCap** | Pointer to **string** | Positioning capabilities supported by the UE. A string encoding the \&quot;ProvideCapabilities-r9-IEs\&quot; IE as specified in clause 6.3 of 3GPP TS 37.355 (start from octet 1). | [optional] 
**AstiDistributionIndication** | Pointer to **bool** |  | [optional] [default to false]
**TsErrorBudget** | Pointer to **int32** |  | [optional] 
**SnpnOnboardInd** | Pointer to **bool** |  | [optional] [default to false]
**SmfSelInfo** | Pointer to [**NullableSmfSelectionData**](SmfSelectionData.md) |  | [optional] 
**PcfUeSliceMbrList** | Pointer to [**map[string]SliceMbr**](SliceMbr.md) | A map(list of key-value pairs) where Snssai serves as key. | [optional] 
**SmsfSetId** | Pointer to **string** | NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \&quot;set&lt;Set ID&gt;.&lt;nftype&gt;set.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.&lt;NFType&gt;set.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)  &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NFType&gt; encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters &lt;Set ID&gt; encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.   | [optional] 
**SmsfServiceSetId** | Pointer to **string** | NF Service Set Identifier (see clause 28.12 of 3GPP TS 23.003) formatted as the following  string \&quot;set&lt;Set ID&gt;.sn&lt;Service Name&gt;.nfi&lt;NF Instance ID&gt;.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.sn&lt;ServiceName&gt;.nfi&lt;NFInstanceID&gt;.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)   &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NID&gt; encoded as defined in clause 5.4.2 (\&quot;Nid\&quot; data type definition)  &lt;NFInstanceId&gt; encoded as defined in clause 5.3.2  &lt;ServiceName&gt; encoded as defined in 3GPP TS 29.510  &lt;Set ID&gt; encoded as a string of characters consisting of alphabetic    characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that shall end    with either an alphabetic character or a digit.  | [optional] 
**SmsfBindingInfo** | Pointer to **string** |  | [optional] 

## Methods

### NewUeContext

`func NewUeContext() *UeContext`

NewUeContext instantiates a new UeContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUeContextWithDefaults

`func NewUeContextWithDefaults() *UeContext`

NewUeContextWithDefaults instantiates a new UeContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupi

`func (o *UeContext) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *UeContext) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *UeContext) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *UeContext) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetSupiUnauthInd

`func (o *UeContext) GetSupiUnauthInd() bool`

GetSupiUnauthInd returns the SupiUnauthInd field if non-nil, zero value otherwise.

### GetSupiUnauthIndOk

`func (o *UeContext) GetSupiUnauthIndOk() (*bool, bool)`

GetSupiUnauthIndOk returns a tuple with the SupiUnauthInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupiUnauthInd

`func (o *UeContext) SetSupiUnauthInd(v bool)`

SetSupiUnauthInd sets SupiUnauthInd field to given value.

### HasSupiUnauthInd

`func (o *UeContext) HasSupiUnauthInd() bool`

HasSupiUnauthInd returns a boolean if a field has been set.

### GetGpsiList

`func (o *UeContext) GetGpsiList() []string`

GetGpsiList returns the GpsiList field if non-nil, zero value otherwise.

### GetGpsiListOk

`func (o *UeContext) GetGpsiListOk() (*[]string, bool)`

GetGpsiListOk returns a tuple with the GpsiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsiList

`func (o *UeContext) SetGpsiList(v []string)`

SetGpsiList sets GpsiList field to given value.

### HasGpsiList

`func (o *UeContext) HasGpsiList() bool`

HasGpsiList returns a boolean if a field has been set.

### GetPei

`func (o *UeContext) GetPei() string`

GetPei returns the Pei field if non-nil, zero value otherwise.

### GetPeiOk

`func (o *UeContext) GetPeiOk() (*string, bool)`

GetPeiOk returns a tuple with the Pei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPei

`func (o *UeContext) SetPei(v string)`

SetPei sets Pei field to given value.

### HasPei

`func (o *UeContext) HasPei() bool`

HasPei returns a boolean if a field has been set.

### GetUdmGroupId

`func (o *UeContext) GetUdmGroupId() string`

GetUdmGroupId returns the UdmGroupId field if non-nil, zero value otherwise.

### GetUdmGroupIdOk

`func (o *UeContext) GetUdmGroupIdOk() (*string, bool)`

GetUdmGroupIdOk returns a tuple with the UdmGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdmGroupId

`func (o *UeContext) SetUdmGroupId(v string)`

SetUdmGroupId sets UdmGroupId field to given value.

### HasUdmGroupId

`func (o *UeContext) HasUdmGroupId() bool`

HasUdmGroupId returns a boolean if a field has been set.

### GetAusfGroupId

`func (o *UeContext) GetAusfGroupId() string`

GetAusfGroupId returns the AusfGroupId field if non-nil, zero value otherwise.

### GetAusfGroupIdOk

`func (o *UeContext) GetAusfGroupIdOk() (*string, bool)`

GetAusfGroupIdOk returns a tuple with the AusfGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAusfGroupId

`func (o *UeContext) SetAusfGroupId(v string)`

SetAusfGroupId sets AusfGroupId field to given value.

### HasAusfGroupId

`func (o *UeContext) HasAusfGroupId() bool`

HasAusfGroupId returns a boolean if a field has been set.

### GetPcfGroupId

`func (o *UeContext) GetPcfGroupId() string`

GetPcfGroupId returns the PcfGroupId field if non-nil, zero value otherwise.

### GetPcfGroupIdOk

`func (o *UeContext) GetPcfGroupIdOk() (*string, bool)`

GetPcfGroupIdOk returns a tuple with the PcfGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfGroupId

`func (o *UeContext) SetPcfGroupId(v string)`

SetPcfGroupId sets PcfGroupId field to given value.

### HasPcfGroupId

`func (o *UeContext) HasPcfGroupId() bool`

HasPcfGroupId returns a boolean if a field has been set.

### GetRoutingIndicator

`func (o *UeContext) GetRoutingIndicator() string`

GetRoutingIndicator returns the RoutingIndicator field if non-nil, zero value otherwise.

### GetRoutingIndicatorOk

`func (o *UeContext) GetRoutingIndicatorOk() (*string, bool)`

GetRoutingIndicatorOk returns a tuple with the RoutingIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingIndicator

`func (o *UeContext) SetRoutingIndicator(v string)`

SetRoutingIndicator sets RoutingIndicator field to given value.

### HasRoutingIndicator

`func (o *UeContext) HasRoutingIndicator() bool`

HasRoutingIndicator returns a boolean if a field has been set.

### GetHNwPubKeyId

`func (o *UeContext) GetHNwPubKeyId() int32`

GetHNwPubKeyId returns the HNwPubKeyId field if non-nil, zero value otherwise.

### GetHNwPubKeyIdOk

`func (o *UeContext) GetHNwPubKeyIdOk() (*int32, bool)`

GetHNwPubKeyIdOk returns a tuple with the HNwPubKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHNwPubKeyId

`func (o *UeContext) SetHNwPubKeyId(v int32)`

SetHNwPubKeyId sets HNwPubKeyId field to given value.

### HasHNwPubKeyId

`func (o *UeContext) HasHNwPubKeyId() bool`

HasHNwPubKeyId returns a boolean if a field has been set.

### GetGroupList

`func (o *UeContext) GetGroupList() []string`

GetGroupList returns the GroupList field if non-nil, zero value otherwise.

### GetGroupListOk

`func (o *UeContext) GetGroupListOk() (*[]string, bool)`

GetGroupListOk returns a tuple with the GroupList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupList

`func (o *UeContext) SetGroupList(v []string)`

SetGroupList sets GroupList field to given value.

### HasGroupList

`func (o *UeContext) HasGroupList() bool`

HasGroupList returns a boolean if a field has been set.

### GetDrxParameter

`func (o *UeContext) GetDrxParameter() string`

GetDrxParameter returns the DrxParameter field if non-nil, zero value otherwise.

### GetDrxParameterOk

`func (o *UeContext) GetDrxParameterOk() (*string, bool)`

GetDrxParameterOk returns a tuple with the DrxParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrxParameter

`func (o *UeContext) SetDrxParameter(v string)`

SetDrxParameter sets DrxParameter field to given value.

### HasDrxParameter

`func (o *UeContext) HasDrxParameter() bool`

HasDrxParameter returns a boolean if a field has been set.

### GetSubRfsp

`func (o *UeContext) GetSubRfsp() int32`

GetSubRfsp returns the SubRfsp field if non-nil, zero value otherwise.

### GetSubRfspOk

`func (o *UeContext) GetSubRfspOk() (*int32, bool)`

GetSubRfspOk returns a tuple with the SubRfsp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubRfsp

`func (o *UeContext) SetSubRfsp(v int32)`

SetSubRfsp sets SubRfsp field to given value.

### HasSubRfsp

`func (o *UeContext) HasSubRfsp() bool`

HasSubRfsp returns a boolean if a field has been set.

### GetUsedRfsp

`func (o *UeContext) GetUsedRfsp() int32`

GetUsedRfsp returns the UsedRfsp field if non-nil, zero value otherwise.

### GetUsedRfspOk

`func (o *UeContext) GetUsedRfspOk() (*int32, bool)`

GetUsedRfspOk returns a tuple with the UsedRfsp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedRfsp

`func (o *UeContext) SetUsedRfsp(v int32)`

SetUsedRfsp sets UsedRfsp field to given value.

### HasUsedRfsp

`func (o *UeContext) HasUsedRfsp() bool`

HasUsedRfsp returns a boolean if a field has been set.

### GetSubUeAmbr

`func (o *UeContext) GetSubUeAmbr() Ambr`

GetSubUeAmbr returns the SubUeAmbr field if non-nil, zero value otherwise.

### GetSubUeAmbrOk

`func (o *UeContext) GetSubUeAmbrOk() (*Ambr, bool)`

GetSubUeAmbrOk returns a tuple with the SubUeAmbr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubUeAmbr

`func (o *UeContext) SetSubUeAmbr(v Ambr)`

SetSubUeAmbr sets SubUeAmbr field to given value.

### HasSubUeAmbr

`func (o *UeContext) HasSubUeAmbr() bool`

HasSubUeAmbr returns a boolean if a field has been set.

### GetSubUeSliceMbrList

`func (o *UeContext) GetSubUeSliceMbrList() map[string]SliceMbr`

GetSubUeSliceMbrList returns the SubUeSliceMbrList field if non-nil, zero value otherwise.

### GetSubUeSliceMbrListOk

`func (o *UeContext) GetSubUeSliceMbrListOk() (*map[string]SliceMbr, bool)`

GetSubUeSliceMbrListOk returns a tuple with the SubUeSliceMbrList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubUeSliceMbrList

`func (o *UeContext) SetSubUeSliceMbrList(v map[string]SliceMbr)`

SetSubUeSliceMbrList sets SubUeSliceMbrList field to given value.

### HasSubUeSliceMbrList

`func (o *UeContext) HasSubUeSliceMbrList() bool`

HasSubUeSliceMbrList returns a boolean if a field has been set.

### GetSmsfId

`func (o *UeContext) GetSmsfId() string`

GetSmsfId returns the SmsfId field if non-nil, zero value otherwise.

### GetSmsfIdOk

`func (o *UeContext) GetSmsfIdOk() (*string, bool)`

GetSmsfIdOk returns a tuple with the SmsfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsfId

`func (o *UeContext) SetSmsfId(v string)`

SetSmsfId sets SmsfId field to given value.

### HasSmsfId

`func (o *UeContext) HasSmsfId() bool`

HasSmsfId returns a boolean if a field has been set.

### GetSeafData

`func (o *UeContext) GetSeafData() SeafData`

GetSeafData returns the SeafData field if non-nil, zero value otherwise.

### GetSeafDataOk

`func (o *UeContext) GetSeafDataOk() (*SeafData, bool)`

GetSeafDataOk returns a tuple with the SeafData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeafData

`func (o *UeContext) SetSeafData(v SeafData)`

SetSeafData sets SeafData field to given value.

### HasSeafData

`func (o *UeContext) HasSeafData() bool`

HasSeafData returns a boolean if a field has been set.

### GetVar5gMmCapability

`func (o *UeContext) GetVar5gMmCapability() string`

GetVar5gMmCapability returns the Var5gMmCapability field if non-nil, zero value otherwise.

### GetVar5gMmCapabilityOk

`func (o *UeContext) GetVar5gMmCapabilityOk() (*string, bool)`

GetVar5gMmCapabilityOk returns a tuple with the Var5gMmCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar5gMmCapability

`func (o *UeContext) SetVar5gMmCapability(v string)`

SetVar5gMmCapability sets Var5gMmCapability field to given value.

### HasVar5gMmCapability

`func (o *UeContext) HasVar5gMmCapability() bool`

HasVar5gMmCapability returns a boolean if a field has been set.

### GetPcfId

`func (o *UeContext) GetPcfId() string`

GetPcfId returns the PcfId field if non-nil, zero value otherwise.

### GetPcfIdOk

`func (o *UeContext) GetPcfIdOk() (*string, bool)`

GetPcfIdOk returns a tuple with the PcfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfId

`func (o *UeContext) SetPcfId(v string)`

SetPcfId sets PcfId field to given value.

### HasPcfId

`func (o *UeContext) HasPcfId() bool`

HasPcfId returns a boolean if a field has been set.

### GetPcfSetId

`func (o *UeContext) GetPcfSetId() string`

GetPcfSetId returns the PcfSetId field if non-nil, zero value otherwise.

### GetPcfSetIdOk

`func (o *UeContext) GetPcfSetIdOk() (*string, bool)`

GetPcfSetIdOk returns a tuple with the PcfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfSetId

`func (o *UeContext) SetPcfSetId(v string)`

SetPcfSetId sets PcfSetId field to given value.

### HasPcfSetId

`func (o *UeContext) HasPcfSetId() bool`

HasPcfSetId returns a boolean if a field has been set.

### GetPcfAmpServiceSetId

`func (o *UeContext) GetPcfAmpServiceSetId() string`

GetPcfAmpServiceSetId returns the PcfAmpServiceSetId field if non-nil, zero value otherwise.

### GetPcfAmpServiceSetIdOk

`func (o *UeContext) GetPcfAmpServiceSetIdOk() (*string, bool)`

GetPcfAmpServiceSetIdOk returns a tuple with the PcfAmpServiceSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfAmpServiceSetId

`func (o *UeContext) SetPcfAmpServiceSetId(v string)`

SetPcfAmpServiceSetId sets PcfAmpServiceSetId field to given value.

### HasPcfAmpServiceSetId

`func (o *UeContext) HasPcfAmpServiceSetId() bool`

HasPcfAmpServiceSetId returns a boolean if a field has been set.

### GetPcfUepServiceSetId

`func (o *UeContext) GetPcfUepServiceSetId() string`

GetPcfUepServiceSetId returns the PcfUepServiceSetId field if non-nil, zero value otherwise.

### GetPcfUepServiceSetIdOk

`func (o *UeContext) GetPcfUepServiceSetIdOk() (*string, bool)`

GetPcfUepServiceSetIdOk returns a tuple with the PcfUepServiceSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfUepServiceSetId

`func (o *UeContext) SetPcfUepServiceSetId(v string)`

SetPcfUepServiceSetId sets PcfUepServiceSetId field to given value.

### HasPcfUepServiceSetId

`func (o *UeContext) HasPcfUepServiceSetId() bool`

HasPcfUepServiceSetId returns a boolean if a field has been set.

### GetPcfBinding

`func (o *UeContext) GetPcfBinding() SbiBindingLevel`

GetPcfBinding returns the PcfBinding field if non-nil, zero value otherwise.

### GetPcfBindingOk

`func (o *UeContext) GetPcfBindingOk() (*SbiBindingLevel, bool)`

GetPcfBindingOk returns a tuple with the PcfBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfBinding

`func (o *UeContext) SetPcfBinding(v SbiBindingLevel)`

SetPcfBinding sets PcfBinding field to given value.

### HasPcfBinding

`func (o *UeContext) HasPcfBinding() bool`

HasPcfBinding returns a boolean if a field has been set.

### GetPcfAmPolicyUri

`func (o *UeContext) GetPcfAmPolicyUri() string`

GetPcfAmPolicyUri returns the PcfAmPolicyUri field if non-nil, zero value otherwise.

### GetPcfAmPolicyUriOk

`func (o *UeContext) GetPcfAmPolicyUriOk() (*string, bool)`

GetPcfAmPolicyUriOk returns a tuple with the PcfAmPolicyUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfAmPolicyUri

`func (o *UeContext) SetPcfAmPolicyUri(v string)`

SetPcfAmPolicyUri sets PcfAmPolicyUri field to given value.

### HasPcfAmPolicyUri

`func (o *UeContext) HasPcfAmPolicyUri() bool`

HasPcfAmPolicyUri returns a boolean if a field has been set.

### GetAmPolicyReqTriggerList

`func (o *UeContext) GetAmPolicyReqTriggerList() []PolicyReqTrigger`

GetAmPolicyReqTriggerList returns the AmPolicyReqTriggerList field if non-nil, zero value otherwise.

### GetAmPolicyReqTriggerListOk

`func (o *UeContext) GetAmPolicyReqTriggerListOk() (*[]PolicyReqTrigger, bool)`

GetAmPolicyReqTriggerListOk returns a tuple with the AmPolicyReqTriggerList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmPolicyReqTriggerList

`func (o *UeContext) SetAmPolicyReqTriggerList(v []PolicyReqTrigger)`

SetAmPolicyReqTriggerList sets AmPolicyReqTriggerList field to given value.

### HasAmPolicyReqTriggerList

`func (o *UeContext) HasAmPolicyReqTriggerList() bool`

HasAmPolicyReqTriggerList returns a boolean if a field has been set.

### GetPcfUePolicyUri

`func (o *UeContext) GetPcfUePolicyUri() string`

GetPcfUePolicyUri returns the PcfUePolicyUri field if non-nil, zero value otherwise.

### GetPcfUePolicyUriOk

`func (o *UeContext) GetPcfUePolicyUriOk() (*string, bool)`

GetPcfUePolicyUriOk returns a tuple with the PcfUePolicyUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfUePolicyUri

`func (o *UeContext) SetPcfUePolicyUri(v string)`

SetPcfUePolicyUri sets PcfUePolicyUri field to given value.

### HasPcfUePolicyUri

`func (o *UeContext) HasPcfUePolicyUri() bool`

HasPcfUePolicyUri returns a boolean if a field has been set.

### GetUePolicyReqTriggerList

`func (o *UeContext) GetUePolicyReqTriggerList() []PolicyReqTrigger`

GetUePolicyReqTriggerList returns the UePolicyReqTriggerList field if non-nil, zero value otherwise.

### GetUePolicyReqTriggerListOk

`func (o *UeContext) GetUePolicyReqTriggerListOk() (*[]PolicyReqTrigger, bool)`

GetUePolicyReqTriggerListOk returns a tuple with the UePolicyReqTriggerList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUePolicyReqTriggerList

`func (o *UeContext) SetUePolicyReqTriggerList(v []PolicyReqTrigger)`

SetUePolicyReqTriggerList sets UePolicyReqTriggerList field to given value.

### HasUePolicyReqTriggerList

`func (o *UeContext) HasUePolicyReqTriggerList() bool`

HasUePolicyReqTriggerList returns a boolean if a field has been set.

### GetHpcfId

`func (o *UeContext) GetHpcfId() string`

GetHpcfId returns the HpcfId field if non-nil, zero value otherwise.

### GetHpcfIdOk

`func (o *UeContext) GetHpcfIdOk() (*string, bool)`

GetHpcfIdOk returns a tuple with the HpcfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHpcfId

`func (o *UeContext) SetHpcfId(v string)`

SetHpcfId sets HpcfId field to given value.

### HasHpcfId

`func (o *UeContext) HasHpcfId() bool`

HasHpcfId returns a boolean if a field has been set.

### GetHpcfSetId

`func (o *UeContext) GetHpcfSetId() string`

GetHpcfSetId returns the HpcfSetId field if non-nil, zero value otherwise.

### GetHpcfSetIdOk

`func (o *UeContext) GetHpcfSetIdOk() (*string, bool)`

GetHpcfSetIdOk returns a tuple with the HpcfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHpcfSetId

`func (o *UeContext) SetHpcfSetId(v string)`

SetHpcfSetId sets HpcfSetId field to given value.

### HasHpcfSetId

`func (o *UeContext) HasHpcfSetId() bool`

HasHpcfSetId returns a boolean if a field has been set.

### GetRestrictedRatList

`func (o *UeContext) GetRestrictedRatList() []RatType`

GetRestrictedRatList returns the RestrictedRatList field if non-nil, zero value otherwise.

### GetRestrictedRatListOk

`func (o *UeContext) GetRestrictedRatListOk() (*[]RatType, bool)`

GetRestrictedRatListOk returns a tuple with the RestrictedRatList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedRatList

`func (o *UeContext) SetRestrictedRatList(v []RatType)`

SetRestrictedRatList sets RestrictedRatList field to given value.

### HasRestrictedRatList

`func (o *UeContext) HasRestrictedRatList() bool`

HasRestrictedRatList returns a boolean if a field has been set.

### GetForbiddenAreaList

`func (o *UeContext) GetForbiddenAreaList() []Area`

GetForbiddenAreaList returns the ForbiddenAreaList field if non-nil, zero value otherwise.

### GetForbiddenAreaListOk

`func (o *UeContext) GetForbiddenAreaListOk() (*[]Area, bool)`

GetForbiddenAreaListOk returns a tuple with the ForbiddenAreaList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForbiddenAreaList

`func (o *UeContext) SetForbiddenAreaList(v []Area)`

SetForbiddenAreaList sets ForbiddenAreaList field to given value.

### HasForbiddenAreaList

`func (o *UeContext) HasForbiddenAreaList() bool`

HasForbiddenAreaList returns a boolean if a field has been set.

### GetServiceAreaRestriction

`func (o *UeContext) GetServiceAreaRestriction() ServiceAreaRestriction`

GetServiceAreaRestriction returns the ServiceAreaRestriction field if non-nil, zero value otherwise.

### GetServiceAreaRestrictionOk

`func (o *UeContext) GetServiceAreaRestrictionOk() (*ServiceAreaRestriction, bool)`

GetServiceAreaRestrictionOk returns a tuple with the ServiceAreaRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAreaRestriction

`func (o *UeContext) SetServiceAreaRestriction(v ServiceAreaRestriction)`

SetServiceAreaRestriction sets ServiceAreaRestriction field to given value.

### HasServiceAreaRestriction

`func (o *UeContext) HasServiceAreaRestriction() bool`

HasServiceAreaRestriction returns a boolean if a field has been set.

### GetRestrictedCoreNwTypeList

`func (o *UeContext) GetRestrictedCoreNwTypeList() []CoreNetworkType`

GetRestrictedCoreNwTypeList returns the RestrictedCoreNwTypeList field if non-nil, zero value otherwise.

### GetRestrictedCoreNwTypeListOk

`func (o *UeContext) GetRestrictedCoreNwTypeListOk() (*[]CoreNetworkType, bool)`

GetRestrictedCoreNwTypeListOk returns a tuple with the RestrictedCoreNwTypeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedCoreNwTypeList

`func (o *UeContext) SetRestrictedCoreNwTypeList(v []CoreNetworkType)`

SetRestrictedCoreNwTypeList sets RestrictedCoreNwTypeList field to given value.

### HasRestrictedCoreNwTypeList

`func (o *UeContext) HasRestrictedCoreNwTypeList() bool`

HasRestrictedCoreNwTypeList returns a boolean if a field has been set.

### GetEventSubscriptionList

`func (o *UeContext) GetEventSubscriptionList() []ExtAmfEventSubscription`

GetEventSubscriptionList returns the EventSubscriptionList field if non-nil, zero value otherwise.

### GetEventSubscriptionListOk

`func (o *UeContext) GetEventSubscriptionListOk() (*[]ExtAmfEventSubscription, bool)`

GetEventSubscriptionListOk returns a tuple with the EventSubscriptionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSubscriptionList

`func (o *UeContext) SetEventSubscriptionList(v []ExtAmfEventSubscription)`

SetEventSubscriptionList sets EventSubscriptionList field to given value.

### HasEventSubscriptionList

`func (o *UeContext) HasEventSubscriptionList() bool`

HasEventSubscriptionList returns a boolean if a field has been set.

### GetMmContextList

`func (o *UeContext) GetMmContextList() []MmContext`

GetMmContextList returns the MmContextList field if non-nil, zero value otherwise.

### GetMmContextListOk

`func (o *UeContext) GetMmContextListOk() (*[]MmContext, bool)`

GetMmContextListOk returns a tuple with the MmContextList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMmContextList

`func (o *UeContext) SetMmContextList(v []MmContext)`

SetMmContextList sets MmContextList field to given value.

### HasMmContextList

`func (o *UeContext) HasMmContextList() bool`

HasMmContextList returns a boolean if a field has been set.

### GetSessionContextList

`func (o *UeContext) GetSessionContextList() []PduSessionContext`

GetSessionContextList returns the SessionContextList field if non-nil, zero value otherwise.

### GetSessionContextListOk

`func (o *UeContext) GetSessionContextListOk() (*[]PduSessionContext, bool)`

GetSessionContextListOk returns a tuple with the SessionContextList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionContextList

`func (o *UeContext) SetSessionContextList(v []PduSessionContext)`

SetSessionContextList sets SessionContextList field to given value.

### HasSessionContextList

`func (o *UeContext) HasSessionContextList() bool`

HasSessionContextList returns a boolean if a field has been set.

### GetEpsInterworkingInfo

`func (o *UeContext) GetEpsInterworkingInfo() EpsInterworkingInfo`

GetEpsInterworkingInfo returns the EpsInterworkingInfo field if non-nil, zero value otherwise.

### GetEpsInterworkingInfoOk

`func (o *UeContext) GetEpsInterworkingInfoOk() (*EpsInterworkingInfo, bool)`

GetEpsInterworkingInfoOk returns a tuple with the EpsInterworkingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsInterworkingInfo

`func (o *UeContext) SetEpsInterworkingInfo(v EpsInterworkingInfo)`

SetEpsInterworkingInfo sets EpsInterworkingInfo field to given value.

### HasEpsInterworkingInfo

`func (o *UeContext) HasEpsInterworkingInfo() bool`

HasEpsInterworkingInfo returns a boolean if a field has been set.

### GetTraceData

`func (o *UeContext) GetTraceData() TraceData`

GetTraceData returns the TraceData field if non-nil, zero value otherwise.

### GetTraceDataOk

`func (o *UeContext) GetTraceDataOk() (*TraceData, bool)`

GetTraceDataOk returns a tuple with the TraceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceData

`func (o *UeContext) SetTraceData(v TraceData)`

SetTraceData sets TraceData field to given value.

### HasTraceData

`func (o *UeContext) HasTraceData() bool`

HasTraceData returns a boolean if a field has been set.

### SetTraceDataNil

`func (o *UeContext) SetTraceDataNil(b bool)`

 SetTraceDataNil sets the value for TraceData to be an explicit nil

### UnsetTraceData
`func (o *UeContext) UnsetTraceData()`

UnsetTraceData ensures that no value is present for TraceData, not even an explicit nil
### GetServiceGapExpiryTime

`func (o *UeContext) GetServiceGapExpiryTime() time.Time`

GetServiceGapExpiryTime returns the ServiceGapExpiryTime field if non-nil, zero value otherwise.

### GetServiceGapExpiryTimeOk

`func (o *UeContext) GetServiceGapExpiryTimeOk() (*time.Time, bool)`

GetServiceGapExpiryTimeOk returns a tuple with the ServiceGapExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceGapExpiryTime

`func (o *UeContext) SetServiceGapExpiryTime(v time.Time)`

SetServiceGapExpiryTime sets ServiceGapExpiryTime field to given value.

### HasServiceGapExpiryTime

`func (o *UeContext) HasServiceGapExpiryTime() bool`

HasServiceGapExpiryTime returns a boolean if a field has been set.

### GetStnSr

`func (o *UeContext) GetStnSr() string`

GetStnSr returns the StnSr field if non-nil, zero value otherwise.

### GetStnSrOk

`func (o *UeContext) GetStnSrOk() (*string, bool)`

GetStnSrOk returns a tuple with the StnSr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStnSr

`func (o *UeContext) SetStnSr(v string)`

SetStnSr sets StnSr field to given value.

### HasStnSr

`func (o *UeContext) HasStnSr() bool`

HasStnSr returns a boolean if a field has been set.

### GetCMsisdn

`func (o *UeContext) GetCMsisdn() string`

GetCMsisdn returns the CMsisdn field if non-nil, zero value otherwise.

### GetCMsisdnOk

`func (o *UeContext) GetCMsisdnOk() (*string, bool)`

GetCMsisdnOk returns a tuple with the CMsisdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCMsisdn

`func (o *UeContext) SetCMsisdn(v string)`

SetCMsisdn sets CMsisdn field to given value.

### HasCMsisdn

`func (o *UeContext) HasCMsisdn() bool`

HasCMsisdn returns a boolean if a field has been set.

### GetMsClassmark2

`func (o *UeContext) GetMsClassmark2() string`

GetMsClassmark2 returns the MsClassmark2 field if non-nil, zero value otherwise.

### GetMsClassmark2Ok

`func (o *UeContext) GetMsClassmark2Ok() (*string, bool)`

GetMsClassmark2Ok returns a tuple with the MsClassmark2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsClassmark2

`func (o *UeContext) SetMsClassmark2(v string)`

SetMsClassmark2 sets MsClassmark2 field to given value.

### HasMsClassmark2

`func (o *UeContext) HasMsClassmark2() bool`

HasMsClassmark2 returns a boolean if a field has been set.

### GetSupportedCodecList

`func (o *UeContext) GetSupportedCodecList() []string`

GetSupportedCodecList returns the SupportedCodecList field if non-nil, zero value otherwise.

### GetSupportedCodecListOk

`func (o *UeContext) GetSupportedCodecListOk() (*[]string, bool)`

GetSupportedCodecListOk returns a tuple with the SupportedCodecList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedCodecList

`func (o *UeContext) SetSupportedCodecList(v []string)`

SetSupportedCodecList sets SupportedCodecList field to given value.

### HasSupportedCodecList

`func (o *UeContext) HasSupportedCodecList() bool`

HasSupportedCodecList returns a boolean if a field has been set.

### GetSmallDataRateStatusInfos

`func (o *UeContext) GetSmallDataRateStatusInfos() []SmallDataRateStatusInfo`

GetSmallDataRateStatusInfos returns the SmallDataRateStatusInfos field if non-nil, zero value otherwise.

### GetSmallDataRateStatusInfosOk

`func (o *UeContext) GetSmallDataRateStatusInfosOk() (*[]SmallDataRateStatusInfo, bool)`

GetSmallDataRateStatusInfosOk returns a tuple with the SmallDataRateStatusInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmallDataRateStatusInfos

`func (o *UeContext) SetSmallDataRateStatusInfos(v []SmallDataRateStatusInfo)`

SetSmallDataRateStatusInfos sets SmallDataRateStatusInfos field to given value.

### HasSmallDataRateStatusInfos

`func (o *UeContext) HasSmallDataRateStatusInfos() bool`

HasSmallDataRateStatusInfos returns a boolean if a field has been set.

### GetRestrictedPrimaryRatList

`func (o *UeContext) GetRestrictedPrimaryRatList() []RatType`

GetRestrictedPrimaryRatList returns the RestrictedPrimaryRatList field if non-nil, zero value otherwise.

### GetRestrictedPrimaryRatListOk

`func (o *UeContext) GetRestrictedPrimaryRatListOk() (*[]RatType, bool)`

GetRestrictedPrimaryRatListOk returns a tuple with the RestrictedPrimaryRatList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedPrimaryRatList

`func (o *UeContext) SetRestrictedPrimaryRatList(v []RatType)`

SetRestrictedPrimaryRatList sets RestrictedPrimaryRatList field to given value.

### HasRestrictedPrimaryRatList

`func (o *UeContext) HasRestrictedPrimaryRatList() bool`

HasRestrictedPrimaryRatList returns a boolean if a field has been set.

### GetRestrictedSecondaryRatList

`func (o *UeContext) GetRestrictedSecondaryRatList() []RatType`

GetRestrictedSecondaryRatList returns the RestrictedSecondaryRatList field if non-nil, zero value otherwise.

### GetRestrictedSecondaryRatListOk

`func (o *UeContext) GetRestrictedSecondaryRatListOk() (*[]RatType, bool)`

GetRestrictedSecondaryRatListOk returns a tuple with the RestrictedSecondaryRatList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedSecondaryRatList

`func (o *UeContext) SetRestrictedSecondaryRatList(v []RatType)`

SetRestrictedSecondaryRatList sets RestrictedSecondaryRatList field to given value.

### HasRestrictedSecondaryRatList

`func (o *UeContext) HasRestrictedSecondaryRatList() bool`

HasRestrictedSecondaryRatList returns a boolean if a field has been set.

### GetV2xContext

`func (o *UeContext) GetV2xContext() V2xContext`

GetV2xContext returns the V2xContext field if non-nil, zero value otherwise.

### GetV2xContextOk

`func (o *UeContext) GetV2xContextOk() (*V2xContext, bool)`

GetV2xContextOk returns a tuple with the V2xContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV2xContext

`func (o *UeContext) SetV2xContext(v V2xContext)`

SetV2xContext sets V2xContext field to given value.

### HasV2xContext

`func (o *UeContext) HasV2xContext() bool`

HasV2xContext returns a boolean if a field has been set.

### GetLteCatMInd

`func (o *UeContext) GetLteCatMInd() bool`

GetLteCatMInd returns the LteCatMInd field if non-nil, zero value otherwise.

### GetLteCatMIndOk

`func (o *UeContext) GetLteCatMIndOk() (*bool, bool)`

GetLteCatMIndOk returns a tuple with the LteCatMInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLteCatMInd

`func (o *UeContext) SetLteCatMInd(v bool)`

SetLteCatMInd sets LteCatMInd field to given value.

### HasLteCatMInd

`func (o *UeContext) HasLteCatMInd() bool`

HasLteCatMInd returns a boolean if a field has been set.

### GetRedCapInd

`func (o *UeContext) GetRedCapInd() bool`

GetRedCapInd returns the RedCapInd field if non-nil, zero value otherwise.

### GetRedCapIndOk

`func (o *UeContext) GetRedCapIndOk() (*bool, bool)`

GetRedCapIndOk returns a tuple with the RedCapInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedCapInd

`func (o *UeContext) SetRedCapInd(v bool)`

SetRedCapInd sets RedCapInd field to given value.

### HasRedCapInd

`func (o *UeContext) HasRedCapInd() bool`

HasRedCapInd returns a boolean if a field has been set.

### GetMoExpDataCounter

`func (o *UeContext) GetMoExpDataCounter() MoExpDataCounter`

GetMoExpDataCounter returns the MoExpDataCounter field if non-nil, zero value otherwise.

### GetMoExpDataCounterOk

`func (o *UeContext) GetMoExpDataCounterOk() (*MoExpDataCounter, bool)`

GetMoExpDataCounterOk returns a tuple with the MoExpDataCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoExpDataCounter

`func (o *UeContext) SetMoExpDataCounter(v MoExpDataCounter)`

SetMoExpDataCounter sets MoExpDataCounter field to given value.

### HasMoExpDataCounter

`func (o *UeContext) HasMoExpDataCounter() bool`

HasMoExpDataCounter returns a boolean if a field has been set.

### GetCagData

`func (o *UeContext) GetCagData() CagData`

GetCagData returns the CagData field if non-nil, zero value otherwise.

### GetCagDataOk

`func (o *UeContext) GetCagDataOk() (*CagData, bool)`

GetCagDataOk returns a tuple with the CagData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCagData

`func (o *UeContext) SetCagData(v CagData)`

SetCagData sets CagData field to given value.

### HasCagData

`func (o *UeContext) HasCagData() bool`

HasCagData returns a boolean if a field has been set.

### GetManagementMdtInd

`func (o *UeContext) GetManagementMdtInd() bool`

GetManagementMdtInd returns the ManagementMdtInd field if non-nil, zero value otherwise.

### GetManagementMdtIndOk

`func (o *UeContext) GetManagementMdtIndOk() (*bool, bool)`

GetManagementMdtIndOk returns a tuple with the ManagementMdtInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementMdtInd

`func (o *UeContext) SetManagementMdtInd(v bool)`

SetManagementMdtInd sets ManagementMdtInd field to given value.

### HasManagementMdtInd

`func (o *UeContext) HasManagementMdtInd() bool`

HasManagementMdtInd returns a boolean if a field has been set.

### GetImmediateMdtConf

`func (o *UeContext) GetImmediateMdtConf() ImmediateMdtConf`

GetImmediateMdtConf returns the ImmediateMdtConf field if non-nil, zero value otherwise.

### GetImmediateMdtConfOk

`func (o *UeContext) GetImmediateMdtConfOk() (*ImmediateMdtConf, bool)`

GetImmediateMdtConfOk returns a tuple with the ImmediateMdtConf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmediateMdtConf

`func (o *UeContext) SetImmediateMdtConf(v ImmediateMdtConf)`

SetImmediateMdtConf sets ImmediateMdtConf field to given value.

### HasImmediateMdtConf

`func (o *UeContext) HasImmediateMdtConf() bool`

HasImmediateMdtConf returns a boolean if a field has been set.

### GetEcRestrictionDataWb

`func (o *UeContext) GetEcRestrictionDataWb() EcRestrictionDataWb`

GetEcRestrictionDataWb returns the EcRestrictionDataWb field if non-nil, zero value otherwise.

### GetEcRestrictionDataWbOk

`func (o *UeContext) GetEcRestrictionDataWbOk() (*EcRestrictionDataWb, bool)`

GetEcRestrictionDataWbOk returns a tuple with the EcRestrictionDataWb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcRestrictionDataWb

`func (o *UeContext) SetEcRestrictionDataWb(v EcRestrictionDataWb)`

SetEcRestrictionDataWb sets EcRestrictionDataWb field to given value.

### HasEcRestrictionDataWb

`func (o *UeContext) HasEcRestrictionDataWb() bool`

HasEcRestrictionDataWb returns a boolean if a field has been set.

### GetEcRestrictionDataNb

`func (o *UeContext) GetEcRestrictionDataNb() bool`

GetEcRestrictionDataNb returns the EcRestrictionDataNb field if non-nil, zero value otherwise.

### GetEcRestrictionDataNbOk

`func (o *UeContext) GetEcRestrictionDataNbOk() (*bool, bool)`

GetEcRestrictionDataNbOk returns a tuple with the EcRestrictionDataNb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcRestrictionDataNb

`func (o *UeContext) SetEcRestrictionDataNb(v bool)`

SetEcRestrictionDataNb sets EcRestrictionDataNb field to given value.

### HasEcRestrictionDataNb

`func (o *UeContext) HasEcRestrictionDataNb() bool`

HasEcRestrictionDataNb returns a boolean if a field has been set.

### GetIabOperationAllowed

`func (o *UeContext) GetIabOperationAllowed() bool`

GetIabOperationAllowed returns the IabOperationAllowed field if non-nil, zero value otherwise.

### GetIabOperationAllowedOk

`func (o *UeContext) GetIabOperationAllowedOk() (*bool, bool)`

GetIabOperationAllowedOk returns a tuple with the IabOperationAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIabOperationAllowed

`func (o *UeContext) SetIabOperationAllowed(v bool)`

SetIabOperationAllowed sets IabOperationAllowed field to given value.

### HasIabOperationAllowed

`func (o *UeContext) HasIabOperationAllowed() bool`

HasIabOperationAllowed returns a boolean if a field has been set.

### GetProseContext

`func (o *UeContext) GetProseContext() ProseContext`

GetProseContext returns the ProseContext field if non-nil, zero value otherwise.

### GetProseContextOk

`func (o *UeContext) GetProseContextOk() (*ProseContext, bool)`

GetProseContextOk returns a tuple with the ProseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProseContext

`func (o *UeContext) SetProseContext(v ProseContext)`

SetProseContext sets ProseContext field to given value.

### HasProseContext

`func (o *UeContext) HasProseContext() bool`

HasProseContext returns a boolean if a field has been set.

### GetAnalyticsSubscriptionList

`func (o *UeContext) GetAnalyticsSubscriptionList() []AnalyticsSubscription`

GetAnalyticsSubscriptionList returns the AnalyticsSubscriptionList field if non-nil, zero value otherwise.

### GetAnalyticsSubscriptionListOk

`func (o *UeContext) GetAnalyticsSubscriptionListOk() (*[]AnalyticsSubscription, bool)`

GetAnalyticsSubscriptionListOk returns a tuple with the AnalyticsSubscriptionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticsSubscriptionList

`func (o *UeContext) SetAnalyticsSubscriptionList(v []AnalyticsSubscription)`

SetAnalyticsSubscriptionList sets AnalyticsSubscriptionList field to given value.

### HasAnalyticsSubscriptionList

`func (o *UeContext) HasAnalyticsSubscriptionList() bool`

HasAnalyticsSubscriptionList returns a boolean if a field has been set.

### GetPcfAmpBindingInfo

`func (o *UeContext) GetPcfAmpBindingInfo() string`

GetPcfAmpBindingInfo returns the PcfAmpBindingInfo field if non-nil, zero value otherwise.

### GetPcfAmpBindingInfoOk

`func (o *UeContext) GetPcfAmpBindingInfoOk() (*string, bool)`

GetPcfAmpBindingInfoOk returns a tuple with the PcfAmpBindingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfAmpBindingInfo

`func (o *UeContext) SetPcfAmpBindingInfo(v string)`

SetPcfAmpBindingInfo sets PcfAmpBindingInfo field to given value.

### HasPcfAmpBindingInfo

`func (o *UeContext) HasPcfAmpBindingInfo() bool`

HasPcfAmpBindingInfo returns a boolean if a field has been set.

### GetPcfUepBindingInfo

`func (o *UeContext) GetPcfUepBindingInfo() string`

GetPcfUepBindingInfo returns the PcfUepBindingInfo field if non-nil, zero value otherwise.

### GetPcfUepBindingInfoOk

`func (o *UeContext) GetPcfUepBindingInfoOk() (*string, bool)`

GetPcfUepBindingInfoOk returns a tuple with the PcfUepBindingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfUepBindingInfo

`func (o *UeContext) SetPcfUepBindingInfo(v string)`

SetPcfUepBindingInfo sets PcfUepBindingInfo field to given value.

### HasPcfUepBindingInfo

`func (o *UeContext) HasPcfUepBindingInfo() bool`

HasPcfUepBindingInfo returns a boolean if a field has been set.

### GetUsedServiceAreaRestriction

`func (o *UeContext) GetUsedServiceAreaRestriction() ServiceAreaRestriction`

GetUsedServiceAreaRestriction returns the UsedServiceAreaRestriction field if non-nil, zero value otherwise.

### GetUsedServiceAreaRestrictionOk

`func (o *UeContext) GetUsedServiceAreaRestrictionOk() (*ServiceAreaRestriction, bool)`

GetUsedServiceAreaRestrictionOk returns a tuple with the UsedServiceAreaRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedServiceAreaRestriction

`func (o *UeContext) SetUsedServiceAreaRestriction(v ServiceAreaRestriction)`

SetUsedServiceAreaRestriction sets UsedServiceAreaRestriction field to given value.

### HasUsedServiceAreaRestriction

`func (o *UeContext) HasUsedServiceAreaRestriction() bool`

HasUsedServiceAreaRestriction returns a boolean if a field has been set.

### GetPraInAmPolicy

`func (o *UeContext) GetPraInAmPolicy() map[string]PresenceInfo`

GetPraInAmPolicy returns the PraInAmPolicy field if non-nil, zero value otherwise.

### GetPraInAmPolicyOk

`func (o *UeContext) GetPraInAmPolicyOk() (*map[string]PresenceInfo, bool)`

GetPraInAmPolicyOk returns a tuple with the PraInAmPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPraInAmPolicy

`func (o *UeContext) SetPraInAmPolicy(v map[string]PresenceInfo)`

SetPraInAmPolicy sets PraInAmPolicy field to given value.

### HasPraInAmPolicy

`func (o *UeContext) HasPraInAmPolicy() bool`

HasPraInAmPolicy returns a boolean if a field has been set.

### GetPraInUePolicy

`func (o *UeContext) GetPraInUePolicy() map[string]PresenceInfo`

GetPraInUePolicy returns the PraInUePolicy field if non-nil, zero value otherwise.

### GetPraInUePolicyOk

`func (o *UeContext) GetPraInUePolicyOk() (*map[string]PresenceInfo, bool)`

GetPraInUePolicyOk returns a tuple with the PraInUePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPraInUePolicy

`func (o *UeContext) SetPraInUePolicy(v map[string]PresenceInfo)`

SetPraInUePolicy sets PraInUePolicy field to given value.

### HasPraInUePolicy

`func (o *UeContext) HasPraInUePolicy() bool`

HasPraInUePolicy returns a boolean if a field has been set.

### GetUpdpSubscriptionData

`func (o *UeContext) GetUpdpSubscriptionData() UpdpSubscriptionData`

GetUpdpSubscriptionData returns the UpdpSubscriptionData field if non-nil, zero value otherwise.

### GetUpdpSubscriptionDataOk

`func (o *UeContext) GetUpdpSubscriptionDataOk() (*UpdpSubscriptionData, bool)`

GetUpdpSubscriptionDataOk returns a tuple with the UpdpSubscriptionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdpSubscriptionData

`func (o *UeContext) SetUpdpSubscriptionData(v UpdpSubscriptionData)`

SetUpdpSubscriptionData sets UpdpSubscriptionData field to given value.

### HasUpdpSubscriptionData

`func (o *UeContext) HasUpdpSubscriptionData() bool`

HasUpdpSubscriptionData returns a boolean if a field has been set.

### GetSmPolicyNotifyPduList

`func (o *UeContext) GetSmPolicyNotifyPduList() []PduSessionInfo`

GetSmPolicyNotifyPduList returns the SmPolicyNotifyPduList field if non-nil, zero value otherwise.

### GetSmPolicyNotifyPduListOk

`func (o *UeContext) GetSmPolicyNotifyPduListOk() (*[]PduSessionInfo, bool)`

GetSmPolicyNotifyPduListOk returns a tuple with the SmPolicyNotifyPduList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmPolicyNotifyPduList

`func (o *UeContext) SetSmPolicyNotifyPduList(v []PduSessionInfo)`

SetSmPolicyNotifyPduList sets SmPolicyNotifyPduList field to given value.

### HasSmPolicyNotifyPduList

`func (o *UeContext) HasSmPolicyNotifyPduList() bool`

HasSmPolicyNotifyPduList returns a boolean if a field has been set.

### GetPcfUeCallbackInfo

`func (o *UeContext) GetPcfUeCallbackInfo() PcfUeCallbackInfo`

GetPcfUeCallbackInfo returns the PcfUeCallbackInfo field if non-nil, zero value otherwise.

### GetPcfUeCallbackInfoOk

`func (o *UeContext) GetPcfUeCallbackInfoOk() (*PcfUeCallbackInfo, bool)`

GetPcfUeCallbackInfoOk returns a tuple with the PcfUeCallbackInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfUeCallbackInfo

`func (o *UeContext) SetPcfUeCallbackInfo(v PcfUeCallbackInfo)`

SetPcfUeCallbackInfo sets PcfUeCallbackInfo field to given value.

### HasPcfUeCallbackInfo

`func (o *UeContext) HasPcfUeCallbackInfo() bool`

HasPcfUeCallbackInfo returns a boolean if a field has been set.

### SetPcfUeCallbackInfoNil

`func (o *UeContext) SetPcfUeCallbackInfoNil(b bool)`

 SetPcfUeCallbackInfoNil sets the value for PcfUeCallbackInfo to be an explicit nil

### UnsetPcfUeCallbackInfo
`func (o *UeContext) UnsetPcfUeCallbackInfo()`

UnsetPcfUeCallbackInfo ensures that no value is present for PcfUeCallbackInfo, not even an explicit nil
### GetUePositioningCap

`func (o *UeContext) GetUePositioningCap() string`

GetUePositioningCap returns the UePositioningCap field if non-nil, zero value otherwise.

### GetUePositioningCapOk

`func (o *UeContext) GetUePositioningCapOk() (*string, bool)`

GetUePositioningCapOk returns a tuple with the UePositioningCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUePositioningCap

`func (o *UeContext) SetUePositioningCap(v string)`

SetUePositioningCap sets UePositioningCap field to given value.

### HasUePositioningCap

`func (o *UeContext) HasUePositioningCap() bool`

HasUePositioningCap returns a boolean if a field has been set.

### GetAstiDistributionIndication

`func (o *UeContext) GetAstiDistributionIndication() bool`

GetAstiDistributionIndication returns the AstiDistributionIndication field if non-nil, zero value otherwise.

### GetAstiDistributionIndicationOk

`func (o *UeContext) GetAstiDistributionIndicationOk() (*bool, bool)`

GetAstiDistributionIndicationOk returns a tuple with the AstiDistributionIndication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAstiDistributionIndication

`func (o *UeContext) SetAstiDistributionIndication(v bool)`

SetAstiDistributionIndication sets AstiDistributionIndication field to given value.

### HasAstiDistributionIndication

`func (o *UeContext) HasAstiDistributionIndication() bool`

HasAstiDistributionIndication returns a boolean if a field has been set.

### GetTsErrorBudget

`func (o *UeContext) GetTsErrorBudget() int32`

GetTsErrorBudget returns the TsErrorBudget field if non-nil, zero value otherwise.

### GetTsErrorBudgetOk

`func (o *UeContext) GetTsErrorBudgetOk() (*int32, bool)`

GetTsErrorBudgetOk returns a tuple with the TsErrorBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsErrorBudget

`func (o *UeContext) SetTsErrorBudget(v int32)`

SetTsErrorBudget sets TsErrorBudget field to given value.

### HasTsErrorBudget

`func (o *UeContext) HasTsErrorBudget() bool`

HasTsErrorBudget returns a boolean if a field has been set.

### GetSnpnOnboardInd

`func (o *UeContext) GetSnpnOnboardInd() bool`

GetSnpnOnboardInd returns the SnpnOnboardInd field if non-nil, zero value otherwise.

### GetSnpnOnboardIndOk

`func (o *UeContext) GetSnpnOnboardIndOk() (*bool, bool)`

GetSnpnOnboardIndOk returns a tuple with the SnpnOnboardInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnpnOnboardInd

`func (o *UeContext) SetSnpnOnboardInd(v bool)`

SetSnpnOnboardInd sets SnpnOnboardInd field to given value.

### HasSnpnOnboardInd

`func (o *UeContext) HasSnpnOnboardInd() bool`

HasSnpnOnboardInd returns a boolean if a field has been set.

### GetSmfSelInfo

`func (o *UeContext) GetSmfSelInfo() SmfSelectionData`

GetSmfSelInfo returns the SmfSelInfo field if non-nil, zero value otherwise.

### GetSmfSelInfoOk

`func (o *UeContext) GetSmfSelInfoOk() (*SmfSelectionData, bool)`

GetSmfSelInfoOk returns a tuple with the SmfSelInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfSelInfo

`func (o *UeContext) SetSmfSelInfo(v SmfSelectionData)`

SetSmfSelInfo sets SmfSelInfo field to given value.

### HasSmfSelInfo

`func (o *UeContext) HasSmfSelInfo() bool`

HasSmfSelInfo returns a boolean if a field has been set.

### SetSmfSelInfoNil

`func (o *UeContext) SetSmfSelInfoNil(b bool)`

 SetSmfSelInfoNil sets the value for SmfSelInfo to be an explicit nil

### UnsetSmfSelInfo
`func (o *UeContext) UnsetSmfSelInfo()`

UnsetSmfSelInfo ensures that no value is present for SmfSelInfo, not even an explicit nil
### GetPcfUeSliceMbrList

`func (o *UeContext) GetPcfUeSliceMbrList() map[string]SliceMbr`

GetPcfUeSliceMbrList returns the PcfUeSliceMbrList field if non-nil, zero value otherwise.

### GetPcfUeSliceMbrListOk

`func (o *UeContext) GetPcfUeSliceMbrListOk() (*map[string]SliceMbr, bool)`

GetPcfUeSliceMbrListOk returns a tuple with the PcfUeSliceMbrList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfUeSliceMbrList

`func (o *UeContext) SetPcfUeSliceMbrList(v map[string]SliceMbr)`

SetPcfUeSliceMbrList sets PcfUeSliceMbrList field to given value.

### HasPcfUeSliceMbrList

`func (o *UeContext) HasPcfUeSliceMbrList() bool`

HasPcfUeSliceMbrList returns a boolean if a field has been set.

### GetSmsfSetId

`func (o *UeContext) GetSmsfSetId() string`

GetSmsfSetId returns the SmsfSetId field if non-nil, zero value otherwise.

### GetSmsfSetIdOk

`func (o *UeContext) GetSmsfSetIdOk() (*string, bool)`

GetSmsfSetIdOk returns a tuple with the SmsfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsfSetId

`func (o *UeContext) SetSmsfSetId(v string)`

SetSmsfSetId sets SmsfSetId field to given value.

### HasSmsfSetId

`func (o *UeContext) HasSmsfSetId() bool`

HasSmsfSetId returns a boolean if a field has been set.

### GetSmsfServiceSetId

`func (o *UeContext) GetSmsfServiceSetId() string`

GetSmsfServiceSetId returns the SmsfServiceSetId field if non-nil, zero value otherwise.

### GetSmsfServiceSetIdOk

`func (o *UeContext) GetSmsfServiceSetIdOk() (*string, bool)`

GetSmsfServiceSetIdOk returns a tuple with the SmsfServiceSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsfServiceSetId

`func (o *UeContext) SetSmsfServiceSetId(v string)`

SetSmsfServiceSetId sets SmsfServiceSetId field to given value.

### HasSmsfServiceSetId

`func (o *UeContext) HasSmsfServiceSetId() bool`

HasSmsfServiceSetId returns a boolean if a field has been set.

### GetSmsfBindingInfo

`func (o *UeContext) GetSmsfBindingInfo() string`

GetSmsfBindingInfo returns the SmsfBindingInfo field if non-nil, zero value otherwise.

### GetSmsfBindingInfoOk

`func (o *UeContext) GetSmsfBindingInfoOk() (*string, bool)`

GetSmsfBindingInfoOk returns a tuple with the SmsfBindingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsfBindingInfo

`func (o *UeContext) SetSmsfBindingInfo(v string)`

SetSmsfBindingInfo sets SmsfBindingInfo field to given value.

### HasSmsfBindingInfo

`func (o *UeContext) HasSmsfBindingInfo() bool`

HasSmsfBindingInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


