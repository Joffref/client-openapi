# ServiceParameterData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **string** | Identifies an application. | [optional] 
**Dnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**Snssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**InterGroupId** | Pointer to **string** | String identifying a group of devices network internal globally unique ID which identifies a set of IMSIs, as specified in clause 19.9 of 3GPP TS 23.003.   | [optional] 
**Supi** | Pointer to **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | [optional] 
**UeIpv4** | Pointer to **string** | string identifying a Ipv4 address formatted in the \&quot;dotted decimal\&quot; notation as defined in IETF RFC 1166. | [optional] 
**UeIpv6** | Pointer to **string** | string identifying a Ipv6 address formatted according to clause 4 in IETF RFC 5952. The mixed Ipv4 Ipv6 notation according to clause 5 of IETF RFC 5952 shall not be used. | [optional] 
**UeMac** | Pointer to **string** | String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042.  | [optional] 
**AnyUeInd** | Pointer to **bool** |  | [optional] 
**ParamOverPc5** | Pointer to **string** | Represents configuration parameters for V2X communications over PC5 reference point.  | [optional] 
**ParamOverUu** | Pointer to **string** | Represents configuration parameters for V2X communications over Uu reference point.  | [optional] 
**ParamForProSeDd** | Pointer to **string** | Represents the service parameters for 5G ProSe direct discovery. | [optional] 
**ParamForProSeDc** | Pointer to **string** | Represents the service parameters for 5G ProSe direct communications. | [optional] 
**ParamForProSeU2NRelUe** | Pointer to **string** | Represents the service parameters for 5G ProSe UE-to-network relay UE. | [optional] 
**ParamForProSeRemUe** | Pointer to **string** | Represents the service parameters for 5G ProSe Remate UE. | [optional] 
**UrspGuidance** | Pointer to [**[]UrspRuleRequest**](UrspRuleRequest.md) | Contains the service parameter used to guide the URSP. | [optional] 
**DeliveryEvents** | Pointer to [**[]Event**](Event.md) | Contains the outcome of the UE Policy Delivery. | [optional] 
**PolicDelivNotifCorreId** | Pointer to **string** | Contains the Notification Correlation Id allocated by the NEF for the notification of UE Policy delivery outcome.  | [optional] 
**PolicDelivNotifUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**ResUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**Headers** | Pointer to **[]string** |  | [optional] 
**ResetIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewServiceParameterData

`func NewServiceParameterData() *ServiceParameterData`

NewServiceParameterData instantiates a new ServiceParameterData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceParameterDataWithDefaults

`func NewServiceParameterDataWithDefaults() *ServiceParameterData`

NewServiceParameterDataWithDefaults instantiates a new ServiceParameterData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *ServiceParameterData) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *ServiceParameterData) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *ServiceParameterData) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *ServiceParameterData) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetDnn

`func (o *ServiceParameterData) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *ServiceParameterData) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *ServiceParameterData) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *ServiceParameterData) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetSnssai

`func (o *ServiceParameterData) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *ServiceParameterData) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *ServiceParameterData) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.

### HasSnssai

`func (o *ServiceParameterData) HasSnssai() bool`

HasSnssai returns a boolean if a field has been set.

### GetInterGroupId

`func (o *ServiceParameterData) GetInterGroupId() string`

GetInterGroupId returns the InterGroupId field if non-nil, zero value otherwise.

### GetInterGroupIdOk

`func (o *ServiceParameterData) GetInterGroupIdOk() (*string, bool)`

GetInterGroupIdOk returns a tuple with the InterGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterGroupId

`func (o *ServiceParameterData) SetInterGroupId(v string)`

SetInterGroupId sets InterGroupId field to given value.

### HasInterGroupId

`func (o *ServiceParameterData) HasInterGroupId() bool`

HasInterGroupId returns a boolean if a field has been set.

### GetSupi

`func (o *ServiceParameterData) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *ServiceParameterData) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *ServiceParameterData) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *ServiceParameterData) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetUeIpv4

`func (o *ServiceParameterData) GetUeIpv4() string`

GetUeIpv4 returns the UeIpv4 field if non-nil, zero value otherwise.

### GetUeIpv4Ok

`func (o *ServiceParameterData) GetUeIpv4Ok() (*string, bool)`

GetUeIpv4Ok returns a tuple with the UeIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpv4

`func (o *ServiceParameterData) SetUeIpv4(v string)`

SetUeIpv4 sets UeIpv4 field to given value.

### HasUeIpv4

`func (o *ServiceParameterData) HasUeIpv4() bool`

HasUeIpv4 returns a boolean if a field has been set.

### GetUeIpv6

`func (o *ServiceParameterData) GetUeIpv6() string`

GetUeIpv6 returns the UeIpv6 field if non-nil, zero value otherwise.

### GetUeIpv6Ok

`func (o *ServiceParameterData) GetUeIpv6Ok() (*string, bool)`

GetUeIpv6Ok returns a tuple with the UeIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpv6

`func (o *ServiceParameterData) SetUeIpv6(v string)`

SetUeIpv6 sets UeIpv6 field to given value.

### HasUeIpv6

`func (o *ServiceParameterData) HasUeIpv6() bool`

HasUeIpv6 returns a boolean if a field has been set.

### GetUeMac

`func (o *ServiceParameterData) GetUeMac() string`

GetUeMac returns the UeMac field if non-nil, zero value otherwise.

### GetUeMacOk

`func (o *ServiceParameterData) GetUeMacOk() (*string, bool)`

GetUeMacOk returns a tuple with the UeMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeMac

`func (o *ServiceParameterData) SetUeMac(v string)`

SetUeMac sets UeMac field to given value.

### HasUeMac

`func (o *ServiceParameterData) HasUeMac() bool`

HasUeMac returns a boolean if a field has been set.

### GetAnyUeInd

`func (o *ServiceParameterData) GetAnyUeInd() bool`

GetAnyUeInd returns the AnyUeInd field if non-nil, zero value otherwise.

### GetAnyUeIndOk

`func (o *ServiceParameterData) GetAnyUeIndOk() (*bool, bool)`

GetAnyUeIndOk returns a tuple with the AnyUeInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyUeInd

`func (o *ServiceParameterData) SetAnyUeInd(v bool)`

SetAnyUeInd sets AnyUeInd field to given value.

### HasAnyUeInd

`func (o *ServiceParameterData) HasAnyUeInd() bool`

HasAnyUeInd returns a boolean if a field has been set.

### GetParamOverPc5

`func (o *ServiceParameterData) GetParamOverPc5() string`

GetParamOverPc5 returns the ParamOverPc5 field if non-nil, zero value otherwise.

### GetParamOverPc5Ok

`func (o *ServiceParameterData) GetParamOverPc5Ok() (*string, bool)`

GetParamOverPc5Ok returns a tuple with the ParamOverPc5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamOverPc5

`func (o *ServiceParameterData) SetParamOverPc5(v string)`

SetParamOverPc5 sets ParamOverPc5 field to given value.

### HasParamOverPc5

`func (o *ServiceParameterData) HasParamOverPc5() bool`

HasParamOverPc5 returns a boolean if a field has been set.

### GetParamOverUu

`func (o *ServiceParameterData) GetParamOverUu() string`

GetParamOverUu returns the ParamOverUu field if non-nil, zero value otherwise.

### GetParamOverUuOk

`func (o *ServiceParameterData) GetParamOverUuOk() (*string, bool)`

GetParamOverUuOk returns a tuple with the ParamOverUu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamOverUu

`func (o *ServiceParameterData) SetParamOverUu(v string)`

SetParamOverUu sets ParamOverUu field to given value.

### HasParamOverUu

`func (o *ServiceParameterData) HasParamOverUu() bool`

HasParamOverUu returns a boolean if a field has been set.

### GetParamForProSeDd

`func (o *ServiceParameterData) GetParamForProSeDd() string`

GetParamForProSeDd returns the ParamForProSeDd field if non-nil, zero value otherwise.

### GetParamForProSeDdOk

`func (o *ServiceParameterData) GetParamForProSeDdOk() (*string, bool)`

GetParamForProSeDdOk returns a tuple with the ParamForProSeDd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamForProSeDd

`func (o *ServiceParameterData) SetParamForProSeDd(v string)`

SetParamForProSeDd sets ParamForProSeDd field to given value.

### HasParamForProSeDd

`func (o *ServiceParameterData) HasParamForProSeDd() bool`

HasParamForProSeDd returns a boolean if a field has been set.

### GetParamForProSeDc

`func (o *ServiceParameterData) GetParamForProSeDc() string`

GetParamForProSeDc returns the ParamForProSeDc field if non-nil, zero value otherwise.

### GetParamForProSeDcOk

`func (o *ServiceParameterData) GetParamForProSeDcOk() (*string, bool)`

GetParamForProSeDcOk returns a tuple with the ParamForProSeDc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamForProSeDc

`func (o *ServiceParameterData) SetParamForProSeDc(v string)`

SetParamForProSeDc sets ParamForProSeDc field to given value.

### HasParamForProSeDc

`func (o *ServiceParameterData) HasParamForProSeDc() bool`

HasParamForProSeDc returns a boolean if a field has been set.

### GetParamForProSeU2NRelUe

`func (o *ServiceParameterData) GetParamForProSeU2NRelUe() string`

GetParamForProSeU2NRelUe returns the ParamForProSeU2NRelUe field if non-nil, zero value otherwise.

### GetParamForProSeU2NRelUeOk

`func (o *ServiceParameterData) GetParamForProSeU2NRelUeOk() (*string, bool)`

GetParamForProSeU2NRelUeOk returns a tuple with the ParamForProSeU2NRelUe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamForProSeU2NRelUe

`func (o *ServiceParameterData) SetParamForProSeU2NRelUe(v string)`

SetParamForProSeU2NRelUe sets ParamForProSeU2NRelUe field to given value.

### HasParamForProSeU2NRelUe

`func (o *ServiceParameterData) HasParamForProSeU2NRelUe() bool`

HasParamForProSeU2NRelUe returns a boolean if a field has been set.

### GetParamForProSeRemUe

`func (o *ServiceParameterData) GetParamForProSeRemUe() string`

GetParamForProSeRemUe returns the ParamForProSeRemUe field if non-nil, zero value otherwise.

### GetParamForProSeRemUeOk

`func (o *ServiceParameterData) GetParamForProSeRemUeOk() (*string, bool)`

GetParamForProSeRemUeOk returns a tuple with the ParamForProSeRemUe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamForProSeRemUe

`func (o *ServiceParameterData) SetParamForProSeRemUe(v string)`

SetParamForProSeRemUe sets ParamForProSeRemUe field to given value.

### HasParamForProSeRemUe

`func (o *ServiceParameterData) HasParamForProSeRemUe() bool`

HasParamForProSeRemUe returns a boolean if a field has been set.

### GetUrspGuidance

`func (o *ServiceParameterData) GetUrspGuidance() []UrspRuleRequest`

GetUrspGuidance returns the UrspGuidance field if non-nil, zero value otherwise.

### GetUrspGuidanceOk

`func (o *ServiceParameterData) GetUrspGuidanceOk() (*[]UrspRuleRequest, bool)`

GetUrspGuidanceOk returns a tuple with the UrspGuidance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrspGuidance

`func (o *ServiceParameterData) SetUrspGuidance(v []UrspRuleRequest)`

SetUrspGuidance sets UrspGuidance field to given value.

### HasUrspGuidance

`func (o *ServiceParameterData) HasUrspGuidance() bool`

HasUrspGuidance returns a boolean if a field has been set.

### GetDeliveryEvents

`func (o *ServiceParameterData) GetDeliveryEvents() []Event`

GetDeliveryEvents returns the DeliveryEvents field if non-nil, zero value otherwise.

### GetDeliveryEventsOk

`func (o *ServiceParameterData) GetDeliveryEventsOk() (*[]Event, bool)`

GetDeliveryEventsOk returns a tuple with the DeliveryEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryEvents

`func (o *ServiceParameterData) SetDeliveryEvents(v []Event)`

SetDeliveryEvents sets DeliveryEvents field to given value.

### HasDeliveryEvents

`func (o *ServiceParameterData) HasDeliveryEvents() bool`

HasDeliveryEvents returns a boolean if a field has been set.

### GetPolicDelivNotifCorreId

`func (o *ServiceParameterData) GetPolicDelivNotifCorreId() string`

GetPolicDelivNotifCorreId returns the PolicDelivNotifCorreId field if non-nil, zero value otherwise.

### GetPolicDelivNotifCorreIdOk

`func (o *ServiceParameterData) GetPolicDelivNotifCorreIdOk() (*string, bool)`

GetPolicDelivNotifCorreIdOk returns a tuple with the PolicDelivNotifCorreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicDelivNotifCorreId

`func (o *ServiceParameterData) SetPolicDelivNotifCorreId(v string)`

SetPolicDelivNotifCorreId sets PolicDelivNotifCorreId field to given value.

### HasPolicDelivNotifCorreId

`func (o *ServiceParameterData) HasPolicDelivNotifCorreId() bool`

HasPolicDelivNotifCorreId returns a boolean if a field has been set.

### GetPolicDelivNotifUri

`func (o *ServiceParameterData) GetPolicDelivNotifUri() string`

GetPolicDelivNotifUri returns the PolicDelivNotifUri field if non-nil, zero value otherwise.

### GetPolicDelivNotifUriOk

`func (o *ServiceParameterData) GetPolicDelivNotifUriOk() (*string, bool)`

GetPolicDelivNotifUriOk returns a tuple with the PolicDelivNotifUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicDelivNotifUri

`func (o *ServiceParameterData) SetPolicDelivNotifUri(v string)`

SetPolicDelivNotifUri sets PolicDelivNotifUri field to given value.

### HasPolicDelivNotifUri

`func (o *ServiceParameterData) HasPolicDelivNotifUri() bool`

HasPolicDelivNotifUri returns a boolean if a field has been set.

### GetSuppFeat

`func (o *ServiceParameterData) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *ServiceParameterData) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *ServiceParameterData) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *ServiceParameterData) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.

### GetResUri

`func (o *ServiceParameterData) GetResUri() string`

GetResUri returns the ResUri field if non-nil, zero value otherwise.

### GetResUriOk

`func (o *ServiceParameterData) GetResUriOk() (*string, bool)`

GetResUriOk returns a tuple with the ResUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResUri

`func (o *ServiceParameterData) SetResUri(v string)`

SetResUri sets ResUri field to given value.

### HasResUri

`func (o *ServiceParameterData) HasResUri() bool`

HasResUri returns a boolean if a field has been set.

### GetHeaders

`func (o *ServiceParameterData) GetHeaders() []string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *ServiceParameterData) GetHeadersOk() (*[]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *ServiceParameterData) SetHeaders(v []string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *ServiceParameterData) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetResetIds

`func (o *ServiceParameterData) GetResetIds() []string`

GetResetIds returns the ResetIds field if non-nil, zero value otherwise.

### GetResetIdsOk

`func (o *ServiceParameterData) GetResetIdsOk() (*[]string, bool)`

GetResetIdsOk returns a tuple with the ResetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetIds

`func (o *ServiceParameterData) SetResetIds(v []string)`

SetResetIds sets ResetIds field to given value.

### HasResetIds

`func (o *ServiceParameterData) HasResetIds() bool`

HasResetIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


