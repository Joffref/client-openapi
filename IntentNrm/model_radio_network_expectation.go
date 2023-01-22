/*
Intent NRM

OAS 3.0.1 definition of the Intent NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package IntentNrm

import (
	"encoding/json"
)

// RadioNetworkExpectation This data type is the \"IntentExpectation\" data type with specialisations to represent MnS consumer's expectations for radio network  delivering and performance assurance    
type RadioNetworkExpectation struct {
	ExpectationId *string `json:"expectationId,omitempty"`
	ExpectationVerb *ExpectationVerb `json:"expectationVerb,omitempty"`
	ExpectationObjects []RadioNetworkExpectationObject `json:"expectationObjects,omitempty"`
	ExpectationTargets []RadioNetworkExpectationExpectationTargetsInner `json:"expectationTargets,omitempty"`
	ExpectationContexts []ExpectationContext `json:"expectationContexts,omitempty"`
	ExpectationfulfilmentInfo *FulfilmentInfo `json:"expectationfulfilmentInfo,omitempty"`
}

// NewRadioNetworkExpectation instantiates a new RadioNetworkExpectation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRadioNetworkExpectation() *RadioNetworkExpectation {
	this := RadioNetworkExpectation{}
	return &this
}

// NewRadioNetworkExpectationWithDefaults instantiates a new RadioNetworkExpectation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRadioNetworkExpectationWithDefaults() *RadioNetworkExpectation {
	this := RadioNetworkExpectation{}
	return &this
}

// GetExpectationId returns the ExpectationId field value if set, zero value otherwise.
func (o *RadioNetworkExpectation) GetExpectationId() string {
	if o == nil || isNil(o.ExpectationId) {
		var ret string
		return ret
	}
	return *o.ExpectationId
}

// GetExpectationIdOk returns a tuple with the ExpectationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadioNetworkExpectation) GetExpectationIdOk() (*string, bool) {
	if o == nil || isNil(o.ExpectationId) {
    return nil, false
	}
	return o.ExpectationId, true
}

// HasExpectationId returns a boolean if a field has been set.
func (o *RadioNetworkExpectation) HasExpectationId() bool {
	if o != nil && !isNil(o.ExpectationId) {
		return true
	}

	return false
}

// SetExpectationId gets a reference to the given string and assigns it to the ExpectationId field.
func (o *RadioNetworkExpectation) SetExpectationId(v string) {
	o.ExpectationId = &v
}

// GetExpectationVerb returns the ExpectationVerb field value if set, zero value otherwise.
func (o *RadioNetworkExpectation) GetExpectationVerb() ExpectationVerb {
	if o == nil || isNil(o.ExpectationVerb) {
		var ret ExpectationVerb
		return ret
	}
	return *o.ExpectationVerb
}

// GetExpectationVerbOk returns a tuple with the ExpectationVerb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadioNetworkExpectation) GetExpectationVerbOk() (*ExpectationVerb, bool) {
	if o == nil || isNil(o.ExpectationVerb) {
    return nil, false
	}
	return o.ExpectationVerb, true
}

// HasExpectationVerb returns a boolean if a field has been set.
func (o *RadioNetworkExpectation) HasExpectationVerb() bool {
	if o != nil && !isNil(o.ExpectationVerb) {
		return true
	}

	return false
}

// SetExpectationVerb gets a reference to the given ExpectationVerb and assigns it to the ExpectationVerb field.
func (o *RadioNetworkExpectation) SetExpectationVerb(v ExpectationVerb) {
	o.ExpectationVerb = &v
}

// GetExpectationObjects returns the ExpectationObjects field value if set, zero value otherwise.
func (o *RadioNetworkExpectation) GetExpectationObjects() []RadioNetworkExpectationObject {
	if o == nil || isNil(o.ExpectationObjects) {
		var ret []RadioNetworkExpectationObject
		return ret
	}
	return o.ExpectationObjects
}

// GetExpectationObjectsOk returns a tuple with the ExpectationObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadioNetworkExpectation) GetExpectationObjectsOk() ([]RadioNetworkExpectationObject, bool) {
	if o == nil || isNil(o.ExpectationObjects) {
    return nil, false
	}
	return o.ExpectationObjects, true
}

// HasExpectationObjects returns a boolean if a field has been set.
func (o *RadioNetworkExpectation) HasExpectationObjects() bool {
	if o != nil && !isNil(o.ExpectationObjects) {
		return true
	}

	return false
}

// SetExpectationObjects gets a reference to the given []RadioNetworkExpectationObject and assigns it to the ExpectationObjects field.
func (o *RadioNetworkExpectation) SetExpectationObjects(v []RadioNetworkExpectationObject) {
	o.ExpectationObjects = v
}

// GetExpectationTargets returns the ExpectationTargets field value if set, zero value otherwise.
func (o *RadioNetworkExpectation) GetExpectationTargets() []RadioNetworkExpectationExpectationTargetsInner {
	if o == nil || isNil(o.ExpectationTargets) {
		var ret []RadioNetworkExpectationExpectationTargetsInner
		return ret
	}
	return o.ExpectationTargets
}

// GetExpectationTargetsOk returns a tuple with the ExpectationTargets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadioNetworkExpectation) GetExpectationTargetsOk() ([]RadioNetworkExpectationExpectationTargetsInner, bool) {
	if o == nil || isNil(o.ExpectationTargets) {
    return nil, false
	}
	return o.ExpectationTargets, true
}

// HasExpectationTargets returns a boolean if a field has been set.
func (o *RadioNetworkExpectation) HasExpectationTargets() bool {
	if o != nil && !isNil(o.ExpectationTargets) {
		return true
	}

	return false
}

// SetExpectationTargets gets a reference to the given []RadioNetworkExpectationExpectationTargetsInner and assigns it to the ExpectationTargets field.
func (o *RadioNetworkExpectation) SetExpectationTargets(v []RadioNetworkExpectationExpectationTargetsInner) {
	o.ExpectationTargets = v
}

// GetExpectationContexts returns the ExpectationContexts field value if set, zero value otherwise.
func (o *RadioNetworkExpectation) GetExpectationContexts() []ExpectationContext {
	if o == nil || isNil(o.ExpectationContexts) {
		var ret []ExpectationContext
		return ret
	}
	return o.ExpectationContexts
}

// GetExpectationContextsOk returns a tuple with the ExpectationContexts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadioNetworkExpectation) GetExpectationContextsOk() ([]ExpectationContext, bool) {
	if o == nil || isNil(o.ExpectationContexts) {
    return nil, false
	}
	return o.ExpectationContexts, true
}

// HasExpectationContexts returns a boolean if a field has been set.
func (o *RadioNetworkExpectation) HasExpectationContexts() bool {
	if o != nil && !isNil(o.ExpectationContexts) {
		return true
	}

	return false
}

// SetExpectationContexts gets a reference to the given []ExpectationContext and assigns it to the ExpectationContexts field.
func (o *RadioNetworkExpectation) SetExpectationContexts(v []ExpectationContext) {
	o.ExpectationContexts = v
}

// GetExpectationfulfilmentInfo returns the ExpectationfulfilmentInfo field value if set, zero value otherwise.
func (o *RadioNetworkExpectation) GetExpectationfulfilmentInfo() FulfilmentInfo {
	if o == nil || isNil(o.ExpectationfulfilmentInfo) {
		var ret FulfilmentInfo
		return ret
	}
	return *o.ExpectationfulfilmentInfo
}

// GetExpectationfulfilmentInfoOk returns a tuple with the ExpectationfulfilmentInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadioNetworkExpectation) GetExpectationfulfilmentInfoOk() (*FulfilmentInfo, bool) {
	if o == nil || isNil(o.ExpectationfulfilmentInfo) {
    return nil, false
	}
	return o.ExpectationfulfilmentInfo, true
}

// HasExpectationfulfilmentInfo returns a boolean if a field has been set.
func (o *RadioNetworkExpectation) HasExpectationfulfilmentInfo() bool {
	if o != nil && !isNil(o.ExpectationfulfilmentInfo) {
		return true
	}

	return false
}

// SetExpectationfulfilmentInfo gets a reference to the given FulfilmentInfo and assigns it to the ExpectationfulfilmentInfo field.
func (o *RadioNetworkExpectation) SetExpectationfulfilmentInfo(v FulfilmentInfo) {
	o.ExpectationfulfilmentInfo = &v
}

func (o RadioNetworkExpectation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ExpectationId) {
		toSerialize["expectationId"] = o.ExpectationId
	}
	if !isNil(o.ExpectationVerb) {
		toSerialize["expectationVerb"] = o.ExpectationVerb
	}
	if !isNil(o.ExpectationObjects) {
		toSerialize["expectationObjects"] = o.ExpectationObjects
	}
	if !isNil(o.ExpectationTargets) {
		toSerialize["expectationTargets"] = o.ExpectationTargets
	}
	if !isNil(o.ExpectationContexts) {
		toSerialize["expectationContexts"] = o.ExpectationContexts
	}
	if !isNil(o.ExpectationfulfilmentInfo) {
		toSerialize["expectationfulfilmentInfo"] = o.ExpectationfulfilmentInfo
	}
	return json.Marshal(toSerialize)
}

type NullableRadioNetworkExpectation struct {
	value *RadioNetworkExpectation
	isSet bool
}

func (v NullableRadioNetworkExpectation) Get() *RadioNetworkExpectation {
	return v.value
}

func (v *NullableRadioNetworkExpectation) Set(val *RadioNetworkExpectation) {
	v.value = val
	v.isSet = true
}

func (v NullableRadioNetworkExpectation) IsSet() bool {
	return v.isSet
}

func (v *NullableRadioNetworkExpectation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRadioNetworkExpectation(val *RadioNetworkExpectation) *NullableRadioNetworkExpectation {
	return &NullableRadioNetworkExpectation{value: val, isSet: true}
}

func (v NullableRadioNetworkExpectation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRadioNetworkExpectation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

