# ProvidePosInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationEstimate** | Pointer to [**GeographicArea**](GeographicArea.md) |  | [optional] 
**LocalLocationEstimate** | Pointer to [**LocalArea**](LocalArea.md) |  | [optional] 
**AccuracyFulfilmentIndicator** | Pointer to [**AccuracyFulfilmentIndicator**](AccuracyFulfilmentIndicator.md) |  | [optional] 
**AgeOfLocationEstimate** | Pointer to **int32** | Indicates value of the age of the location estimate. | [optional] 
**TimestampOfLocationEstimate** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**VelocityEstimate** | Pointer to [**VelocityEstimate**](VelocityEstimate.md) |  | [optional] 
**PositioningDataList** | Pointer to [**[]PositioningMethodAndUsage**](PositioningMethodAndUsage.md) |  | [optional] 
**GnssPositioningDataList** | Pointer to [**[]GnssPositioningMethodAndUsage**](GnssPositioningMethodAndUsage.md) |  | [optional] 
**Ecgi** | Pointer to [**Ecgi**](Ecgi.md) |  | [optional] 
**Ncgi** | Pointer to [**Ncgi**](Ncgi.md) |  | [optional] 
**TargetServingNode** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**TargetMmeName** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**TargetMmeRealm** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**UtranSrvccInd** | Pointer to **bool** |  | [optional] 
**CivicAddress** | Pointer to [**CivicAddress**](CivicAddress.md) |  | [optional] 
**BarometricPressure** | Pointer to **int32** | Specifies the measured uncompensated atmospheric pressure. | [optional] 
**Altitude** | Pointer to **float64** | Indicates value of altitude. | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**ServingLMFIdentification** | Pointer to **string** | LMF identification. | [optional] 
**LocationPrivacyVerResult** | Pointer to [**LocationPrivacyVerResult**](LocationPrivacyVerResult.md) |  | [optional] 
**AchievedQos** | Pointer to [**MinorLocationQoS**](MinorLocationQoS.md) |  | [optional] 

## Methods

### NewProvidePosInfo

`func NewProvidePosInfo() *ProvidePosInfo`

NewProvidePosInfo instantiates a new ProvidePosInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvidePosInfoWithDefaults

`func NewProvidePosInfoWithDefaults() *ProvidePosInfo`

NewProvidePosInfoWithDefaults instantiates a new ProvidePosInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocationEstimate

`func (o *ProvidePosInfo) GetLocationEstimate() GeographicArea`

GetLocationEstimate returns the LocationEstimate field if non-nil, zero value otherwise.

### GetLocationEstimateOk

`func (o *ProvidePosInfo) GetLocationEstimateOk() (*GeographicArea, bool)`

GetLocationEstimateOk returns a tuple with the LocationEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationEstimate

`func (o *ProvidePosInfo) SetLocationEstimate(v GeographicArea)`

SetLocationEstimate sets LocationEstimate field to given value.

### HasLocationEstimate

`func (o *ProvidePosInfo) HasLocationEstimate() bool`

HasLocationEstimate returns a boolean if a field has been set.

### GetLocalLocationEstimate

`func (o *ProvidePosInfo) GetLocalLocationEstimate() LocalArea`

GetLocalLocationEstimate returns the LocalLocationEstimate field if non-nil, zero value otherwise.

### GetLocalLocationEstimateOk

`func (o *ProvidePosInfo) GetLocalLocationEstimateOk() (*LocalArea, bool)`

GetLocalLocationEstimateOk returns a tuple with the LocalLocationEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalLocationEstimate

`func (o *ProvidePosInfo) SetLocalLocationEstimate(v LocalArea)`

SetLocalLocationEstimate sets LocalLocationEstimate field to given value.

### HasLocalLocationEstimate

`func (o *ProvidePosInfo) HasLocalLocationEstimate() bool`

HasLocalLocationEstimate returns a boolean if a field has been set.

### GetAccuracyFulfilmentIndicator

`func (o *ProvidePosInfo) GetAccuracyFulfilmentIndicator() AccuracyFulfilmentIndicator`

GetAccuracyFulfilmentIndicator returns the AccuracyFulfilmentIndicator field if non-nil, zero value otherwise.

### GetAccuracyFulfilmentIndicatorOk

`func (o *ProvidePosInfo) GetAccuracyFulfilmentIndicatorOk() (*AccuracyFulfilmentIndicator, bool)`

GetAccuracyFulfilmentIndicatorOk returns a tuple with the AccuracyFulfilmentIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccuracyFulfilmentIndicator

`func (o *ProvidePosInfo) SetAccuracyFulfilmentIndicator(v AccuracyFulfilmentIndicator)`

SetAccuracyFulfilmentIndicator sets AccuracyFulfilmentIndicator field to given value.

### HasAccuracyFulfilmentIndicator

`func (o *ProvidePosInfo) HasAccuracyFulfilmentIndicator() bool`

HasAccuracyFulfilmentIndicator returns a boolean if a field has been set.

### GetAgeOfLocationEstimate

`func (o *ProvidePosInfo) GetAgeOfLocationEstimate() int32`

GetAgeOfLocationEstimate returns the AgeOfLocationEstimate field if non-nil, zero value otherwise.

### GetAgeOfLocationEstimateOk

`func (o *ProvidePosInfo) GetAgeOfLocationEstimateOk() (*int32, bool)`

GetAgeOfLocationEstimateOk returns a tuple with the AgeOfLocationEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeOfLocationEstimate

`func (o *ProvidePosInfo) SetAgeOfLocationEstimate(v int32)`

SetAgeOfLocationEstimate sets AgeOfLocationEstimate field to given value.

### HasAgeOfLocationEstimate

`func (o *ProvidePosInfo) HasAgeOfLocationEstimate() bool`

HasAgeOfLocationEstimate returns a boolean if a field has been set.

### GetTimestampOfLocationEstimate

`func (o *ProvidePosInfo) GetTimestampOfLocationEstimate() time.Time`

GetTimestampOfLocationEstimate returns the TimestampOfLocationEstimate field if non-nil, zero value otherwise.

### GetTimestampOfLocationEstimateOk

`func (o *ProvidePosInfo) GetTimestampOfLocationEstimateOk() (*time.Time, bool)`

GetTimestampOfLocationEstimateOk returns a tuple with the TimestampOfLocationEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampOfLocationEstimate

`func (o *ProvidePosInfo) SetTimestampOfLocationEstimate(v time.Time)`

SetTimestampOfLocationEstimate sets TimestampOfLocationEstimate field to given value.

### HasTimestampOfLocationEstimate

`func (o *ProvidePosInfo) HasTimestampOfLocationEstimate() bool`

HasTimestampOfLocationEstimate returns a boolean if a field has been set.

### GetVelocityEstimate

`func (o *ProvidePosInfo) GetVelocityEstimate() VelocityEstimate`

GetVelocityEstimate returns the VelocityEstimate field if non-nil, zero value otherwise.

### GetVelocityEstimateOk

`func (o *ProvidePosInfo) GetVelocityEstimateOk() (*VelocityEstimate, bool)`

GetVelocityEstimateOk returns a tuple with the VelocityEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVelocityEstimate

`func (o *ProvidePosInfo) SetVelocityEstimate(v VelocityEstimate)`

SetVelocityEstimate sets VelocityEstimate field to given value.

### HasVelocityEstimate

`func (o *ProvidePosInfo) HasVelocityEstimate() bool`

HasVelocityEstimate returns a boolean if a field has been set.

### GetPositioningDataList

`func (o *ProvidePosInfo) GetPositioningDataList() []PositioningMethodAndUsage`

GetPositioningDataList returns the PositioningDataList field if non-nil, zero value otherwise.

### GetPositioningDataListOk

`func (o *ProvidePosInfo) GetPositioningDataListOk() (*[]PositioningMethodAndUsage, bool)`

GetPositioningDataListOk returns a tuple with the PositioningDataList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositioningDataList

`func (o *ProvidePosInfo) SetPositioningDataList(v []PositioningMethodAndUsage)`

SetPositioningDataList sets PositioningDataList field to given value.

### HasPositioningDataList

`func (o *ProvidePosInfo) HasPositioningDataList() bool`

HasPositioningDataList returns a boolean if a field has been set.

### GetGnssPositioningDataList

`func (o *ProvidePosInfo) GetGnssPositioningDataList() []GnssPositioningMethodAndUsage`

GetGnssPositioningDataList returns the GnssPositioningDataList field if non-nil, zero value otherwise.

### GetGnssPositioningDataListOk

`func (o *ProvidePosInfo) GetGnssPositioningDataListOk() (*[]GnssPositioningMethodAndUsage, bool)`

GetGnssPositioningDataListOk returns a tuple with the GnssPositioningDataList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGnssPositioningDataList

`func (o *ProvidePosInfo) SetGnssPositioningDataList(v []GnssPositioningMethodAndUsage)`

SetGnssPositioningDataList sets GnssPositioningDataList field to given value.

### HasGnssPositioningDataList

`func (o *ProvidePosInfo) HasGnssPositioningDataList() bool`

HasGnssPositioningDataList returns a boolean if a field has been set.

### GetEcgi

`func (o *ProvidePosInfo) GetEcgi() Ecgi`

GetEcgi returns the Ecgi field if non-nil, zero value otherwise.

### GetEcgiOk

`func (o *ProvidePosInfo) GetEcgiOk() (*Ecgi, bool)`

GetEcgiOk returns a tuple with the Ecgi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcgi

`func (o *ProvidePosInfo) SetEcgi(v Ecgi)`

SetEcgi sets Ecgi field to given value.

### HasEcgi

`func (o *ProvidePosInfo) HasEcgi() bool`

HasEcgi returns a boolean if a field has been set.

### GetNcgi

`func (o *ProvidePosInfo) GetNcgi() Ncgi`

GetNcgi returns the Ncgi field if non-nil, zero value otherwise.

### GetNcgiOk

`func (o *ProvidePosInfo) GetNcgiOk() (*Ncgi, bool)`

GetNcgiOk returns a tuple with the Ncgi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNcgi

`func (o *ProvidePosInfo) SetNcgi(v Ncgi)`

SetNcgi sets Ncgi field to given value.

### HasNcgi

`func (o *ProvidePosInfo) HasNcgi() bool`

HasNcgi returns a boolean if a field has been set.

### GetTargetServingNode

`func (o *ProvidePosInfo) GetTargetServingNode() string`

GetTargetServingNode returns the TargetServingNode field if non-nil, zero value otherwise.

### GetTargetServingNodeOk

`func (o *ProvidePosInfo) GetTargetServingNodeOk() (*string, bool)`

GetTargetServingNodeOk returns a tuple with the TargetServingNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetServingNode

`func (o *ProvidePosInfo) SetTargetServingNode(v string)`

SetTargetServingNode sets TargetServingNode field to given value.

### HasTargetServingNode

`func (o *ProvidePosInfo) HasTargetServingNode() bool`

HasTargetServingNode returns a boolean if a field has been set.

### GetTargetMmeName

`func (o *ProvidePosInfo) GetTargetMmeName() string`

GetTargetMmeName returns the TargetMmeName field if non-nil, zero value otherwise.

### GetTargetMmeNameOk

`func (o *ProvidePosInfo) GetTargetMmeNameOk() (*string, bool)`

GetTargetMmeNameOk returns a tuple with the TargetMmeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMmeName

`func (o *ProvidePosInfo) SetTargetMmeName(v string)`

SetTargetMmeName sets TargetMmeName field to given value.

### HasTargetMmeName

`func (o *ProvidePosInfo) HasTargetMmeName() bool`

HasTargetMmeName returns a boolean if a field has been set.

### GetTargetMmeRealm

`func (o *ProvidePosInfo) GetTargetMmeRealm() string`

GetTargetMmeRealm returns the TargetMmeRealm field if non-nil, zero value otherwise.

### GetTargetMmeRealmOk

`func (o *ProvidePosInfo) GetTargetMmeRealmOk() (*string, bool)`

GetTargetMmeRealmOk returns a tuple with the TargetMmeRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMmeRealm

`func (o *ProvidePosInfo) SetTargetMmeRealm(v string)`

SetTargetMmeRealm sets TargetMmeRealm field to given value.

### HasTargetMmeRealm

`func (o *ProvidePosInfo) HasTargetMmeRealm() bool`

HasTargetMmeRealm returns a boolean if a field has been set.

### GetUtranSrvccInd

`func (o *ProvidePosInfo) GetUtranSrvccInd() bool`

GetUtranSrvccInd returns the UtranSrvccInd field if non-nil, zero value otherwise.

### GetUtranSrvccIndOk

`func (o *ProvidePosInfo) GetUtranSrvccIndOk() (*bool, bool)`

GetUtranSrvccIndOk returns a tuple with the UtranSrvccInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtranSrvccInd

`func (o *ProvidePosInfo) SetUtranSrvccInd(v bool)`

SetUtranSrvccInd sets UtranSrvccInd field to given value.

### HasUtranSrvccInd

`func (o *ProvidePosInfo) HasUtranSrvccInd() bool`

HasUtranSrvccInd returns a boolean if a field has been set.

### GetCivicAddress

`func (o *ProvidePosInfo) GetCivicAddress() CivicAddress`

GetCivicAddress returns the CivicAddress field if non-nil, zero value otherwise.

### GetCivicAddressOk

`func (o *ProvidePosInfo) GetCivicAddressOk() (*CivicAddress, bool)`

GetCivicAddressOk returns a tuple with the CivicAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCivicAddress

`func (o *ProvidePosInfo) SetCivicAddress(v CivicAddress)`

SetCivicAddress sets CivicAddress field to given value.

### HasCivicAddress

`func (o *ProvidePosInfo) HasCivicAddress() bool`

HasCivicAddress returns a boolean if a field has been set.

### GetBarometricPressure

`func (o *ProvidePosInfo) GetBarometricPressure() int32`

GetBarometricPressure returns the BarometricPressure field if non-nil, zero value otherwise.

### GetBarometricPressureOk

`func (o *ProvidePosInfo) GetBarometricPressureOk() (*int32, bool)`

GetBarometricPressureOk returns a tuple with the BarometricPressure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarometricPressure

`func (o *ProvidePosInfo) SetBarometricPressure(v int32)`

SetBarometricPressure sets BarometricPressure field to given value.

### HasBarometricPressure

`func (o *ProvidePosInfo) HasBarometricPressure() bool`

HasBarometricPressure returns a boolean if a field has been set.

### GetAltitude

`func (o *ProvidePosInfo) GetAltitude() float64`

GetAltitude returns the Altitude field if non-nil, zero value otherwise.

### GetAltitudeOk

`func (o *ProvidePosInfo) GetAltitudeOk() (*float64, bool)`

GetAltitudeOk returns a tuple with the Altitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltitude

`func (o *ProvidePosInfo) SetAltitude(v float64)`

SetAltitude sets Altitude field to given value.

### HasAltitude

`func (o *ProvidePosInfo) HasAltitude() bool`

HasAltitude returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *ProvidePosInfo) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *ProvidePosInfo) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *ProvidePosInfo) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *ProvidePosInfo) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetServingLMFIdentification

`func (o *ProvidePosInfo) GetServingLMFIdentification() string`

GetServingLMFIdentification returns the ServingLMFIdentification field if non-nil, zero value otherwise.

### GetServingLMFIdentificationOk

`func (o *ProvidePosInfo) GetServingLMFIdentificationOk() (*string, bool)`

GetServingLMFIdentificationOk returns a tuple with the ServingLMFIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingLMFIdentification

`func (o *ProvidePosInfo) SetServingLMFIdentification(v string)`

SetServingLMFIdentification sets ServingLMFIdentification field to given value.

### HasServingLMFIdentification

`func (o *ProvidePosInfo) HasServingLMFIdentification() bool`

HasServingLMFIdentification returns a boolean if a field has been set.

### GetLocationPrivacyVerResult

`func (o *ProvidePosInfo) GetLocationPrivacyVerResult() LocationPrivacyVerResult`

GetLocationPrivacyVerResult returns the LocationPrivacyVerResult field if non-nil, zero value otherwise.

### GetLocationPrivacyVerResultOk

`func (o *ProvidePosInfo) GetLocationPrivacyVerResultOk() (*LocationPrivacyVerResult, bool)`

GetLocationPrivacyVerResultOk returns a tuple with the LocationPrivacyVerResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationPrivacyVerResult

`func (o *ProvidePosInfo) SetLocationPrivacyVerResult(v LocationPrivacyVerResult)`

SetLocationPrivacyVerResult sets LocationPrivacyVerResult field to given value.

### HasLocationPrivacyVerResult

`func (o *ProvidePosInfo) HasLocationPrivacyVerResult() bool`

HasLocationPrivacyVerResult returns a boolean if a field has been set.

### GetAchievedQos

`func (o *ProvidePosInfo) GetAchievedQos() MinorLocationQoS`

GetAchievedQos returns the AchievedQos field if non-nil, zero value otherwise.

### GetAchievedQosOk

`func (o *ProvidePosInfo) GetAchievedQosOk() (*MinorLocationQoS, bool)`

GetAchievedQosOk returns a tuple with the AchievedQos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievedQos

`func (o *ProvidePosInfo) SetAchievedQos(v MinorLocationQoS)`

SetAchievedQos sets AchievedQos field to given value.

### HasAchievedQos

`func (o *ProvidePosInfo) HasAchievedQos() bool`

HasAchievedQos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


