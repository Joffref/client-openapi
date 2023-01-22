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

// RadioNetworkExpectationObject This data type is the \"ExpectationObject\" data type with specialisations for RadioNetworkExpectation
type RadioNetworkExpectationObject struct {
	ObjectType *string `json:"objectType,omitempty"`
	ObjectInstance *string `json:"objectInstance,omitempty"`
	ObjectContexts []RadioNetworkExpectationObjectObjectContextsInner `json:"objectContexts,omitempty"`
}

// NewRadioNetworkExpectationObject instantiates a new RadioNetworkExpectationObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRadioNetworkExpectationObject() *RadioNetworkExpectationObject {
	this := RadioNetworkExpectationObject{}
	return &this
}

// NewRadioNetworkExpectationObjectWithDefaults instantiates a new RadioNetworkExpectationObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRadioNetworkExpectationObjectWithDefaults() *RadioNetworkExpectationObject {
	this := RadioNetworkExpectationObject{}
	return &this
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *RadioNetworkExpectationObject) GetObjectType() string {
	if o == nil || isNil(o.ObjectType) {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadioNetworkExpectationObject) GetObjectTypeOk() (*string, bool) {
	if o == nil || isNil(o.ObjectType) {
    return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *RadioNetworkExpectationObject) HasObjectType() bool {
	if o != nil && !isNil(o.ObjectType) {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *RadioNetworkExpectationObject) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetObjectInstance returns the ObjectInstance field value if set, zero value otherwise.
func (o *RadioNetworkExpectationObject) GetObjectInstance() string {
	if o == nil || isNil(o.ObjectInstance) {
		var ret string
		return ret
	}
	return *o.ObjectInstance
}

// GetObjectInstanceOk returns a tuple with the ObjectInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadioNetworkExpectationObject) GetObjectInstanceOk() (*string, bool) {
	if o == nil || isNil(o.ObjectInstance) {
    return nil, false
	}
	return o.ObjectInstance, true
}

// HasObjectInstance returns a boolean if a field has been set.
func (o *RadioNetworkExpectationObject) HasObjectInstance() bool {
	if o != nil && !isNil(o.ObjectInstance) {
		return true
	}

	return false
}

// SetObjectInstance gets a reference to the given string and assigns it to the ObjectInstance field.
func (o *RadioNetworkExpectationObject) SetObjectInstance(v string) {
	o.ObjectInstance = &v
}

// GetObjectContexts returns the ObjectContexts field value if set, zero value otherwise.
func (o *RadioNetworkExpectationObject) GetObjectContexts() []RadioNetworkExpectationObjectObjectContextsInner {
	if o == nil || isNil(o.ObjectContexts) {
		var ret []RadioNetworkExpectationObjectObjectContextsInner
		return ret
	}
	return o.ObjectContexts
}

// GetObjectContextsOk returns a tuple with the ObjectContexts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadioNetworkExpectationObject) GetObjectContextsOk() ([]RadioNetworkExpectationObjectObjectContextsInner, bool) {
	if o == nil || isNil(o.ObjectContexts) {
    return nil, false
	}
	return o.ObjectContexts, true
}

// HasObjectContexts returns a boolean if a field has been set.
func (o *RadioNetworkExpectationObject) HasObjectContexts() bool {
	if o != nil && !isNil(o.ObjectContexts) {
		return true
	}

	return false
}

// SetObjectContexts gets a reference to the given []RadioNetworkExpectationObjectObjectContextsInner and assigns it to the ObjectContexts field.
func (o *RadioNetworkExpectationObject) SetObjectContexts(v []RadioNetworkExpectationObjectObjectContextsInner) {
	o.ObjectContexts = v
}

func (o RadioNetworkExpectationObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ObjectType) {
		toSerialize["objectType"] = o.ObjectType
	}
	if !isNil(o.ObjectInstance) {
		toSerialize["objectInstance"] = o.ObjectInstance
	}
	if !isNil(o.ObjectContexts) {
		toSerialize["objectContexts"] = o.ObjectContexts
	}
	return json.Marshal(toSerialize)
}

type NullableRadioNetworkExpectationObject struct {
	value *RadioNetworkExpectationObject
	isSet bool
}

func (v NullableRadioNetworkExpectationObject) Get() *RadioNetworkExpectationObject {
	return v.value
}

func (v *NullableRadioNetworkExpectationObject) Set(val *RadioNetworkExpectationObject) {
	v.value = val
	v.isSet = true
}

func (v NullableRadioNetworkExpectationObject) IsSet() bool {
	return v.isSet
}

func (v *NullableRadioNetworkExpectationObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRadioNetworkExpectationObject(val *RadioNetworkExpectationObject) *NullableRadioNetworkExpectationObject {
	return &NullableRadioNetworkExpectationObject{value: val, isSet: true}
}

func (v NullableRadioNetworkExpectationObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRadioNetworkExpectationObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

