/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nchf_ConvergedCharging

import (
	"encoding/json"
	"time"
)

// SupplementaryService struct for SupplementaryService
type SupplementaryService struct {
	SupplementaryServiceType *SupplementaryServiceType `json:"supplementaryServiceType,omitempty"`
	SupplementaryServiceMode *SupplementaryServiceMode `json:"supplementaryServiceMode,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer. 
	NumberOfDiversions *int32 `json:"numberOfDiversions,omitempty"`
	AssociatedPartyAddress *string `json:"associatedPartyAddress,omitempty"`
	ConferenceId *string `json:"conferenceId,omitempty"`
	ParticipantActionType *ParticipantActionType `json:"participantActionType,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	ChangeTime *time.Time `json:"changeTime,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer. 
	NumberOfParticipants *int32 `json:"numberOfParticipants,omitempty"`
	CUGInformation *string `json:"cUGInformation,omitempty"`
}

// NewSupplementaryService instantiates a new SupplementaryService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupplementaryService() *SupplementaryService {
	this := SupplementaryService{}
	return &this
}

// NewSupplementaryServiceWithDefaults instantiates a new SupplementaryService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupplementaryServiceWithDefaults() *SupplementaryService {
	this := SupplementaryService{}
	return &this
}

// GetSupplementaryServiceType returns the SupplementaryServiceType field value if set, zero value otherwise.
func (o *SupplementaryService) GetSupplementaryServiceType() SupplementaryServiceType {
	if o == nil || isNil(o.SupplementaryServiceType) {
		var ret SupplementaryServiceType
		return ret
	}
	return *o.SupplementaryServiceType
}

// GetSupplementaryServiceTypeOk returns a tuple with the SupplementaryServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplementaryService) GetSupplementaryServiceTypeOk() (*SupplementaryServiceType, bool) {
	if o == nil || isNil(o.SupplementaryServiceType) {
    return nil, false
	}
	return o.SupplementaryServiceType, true
}

// HasSupplementaryServiceType returns a boolean if a field has been set.
func (o *SupplementaryService) HasSupplementaryServiceType() bool {
	if o != nil && !isNil(o.SupplementaryServiceType) {
		return true
	}

	return false
}

// SetSupplementaryServiceType gets a reference to the given SupplementaryServiceType and assigns it to the SupplementaryServiceType field.
func (o *SupplementaryService) SetSupplementaryServiceType(v SupplementaryServiceType) {
	o.SupplementaryServiceType = &v
}

// GetSupplementaryServiceMode returns the SupplementaryServiceMode field value if set, zero value otherwise.
func (o *SupplementaryService) GetSupplementaryServiceMode() SupplementaryServiceMode {
	if o == nil || isNil(o.SupplementaryServiceMode) {
		var ret SupplementaryServiceMode
		return ret
	}
	return *o.SupplementaryServiceMode
}

// GetSupplementaryServiceModeOk returns a tuple with the SupplementaryServiceMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplementaryService) GetSupplementaryServiceModeOk() (*SupplementaryServiceMode, bool) {
	if o == nil || isNil(o.SupplementaryServiceMode) {
    return nil, false
	}
	return o.SupplementaryServiceMode, true
}

// HasSupplementaryServiceMode returns a boolean if a field has been set.
func (o *SupplementaryService) HasSupplementaryServiceMode() bool {
	if o != nil && !isNil(o.SupplementaryServiceMode) {
		return true
	}

	return false
}

// SetSupplementaryServiceMode gets a reference to the given SupplementaryServiceMode and assigns it to the SupplementaryServiceMode field.
func (o *SupplementaryService) SetSupplementaryServiceMode(v SupplementaryServiceMode) {
	o.SupplementaryServiceMode = &v
}

// GetNumberOfDiversions returns the NumberOfDiversions field value if set, zero value otherwise.
func (o *SupplementaryService) GetNumberOfDiversions() int32 {
	if o == nil || isNil(o.NumberOfDiversions) {
		var ret int32
		return ret
	}
	return *o.NumberOfDiversions
}

// GetNumberOfDiversionsOk returns a tuple with the NumberOfDiversions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplementaryService) GetNumberOfDiversionsOk() (*int32, bool) {
	if o == nil || isNil(o.NumberOfDiversions) {
    return nil, false
	}
	return o.NumberOfDiversions, true
}

// HasNumberOfDiversions returns a boolean if a field has been set.
func (o *SupplementaryService) HasNumberOfDiversions() bool {
	if o != nil && !isNil(o.NumberOfDiversions) {
		return true
	}

	return false
}

// SetNumberOfDiversions gets a reference to the given int32 and assigns it to the NumberOfDiversions field.
func (o *SupplementaryService) SetNumberOfDiversions(v int32) {
	o.NumberOfDiversions = &v
}

// GetAssociatedPartyAddress returns the AssociatedPartyAddress field value if set, zero value otherwise.
func (o *SupplementaryService) GetAssociatedPartyAddress() string {
	if o == nil || isNil(o.AssociatedPartyAddress) {
		var ret string
		return ret
	}
	return *o.AssociatedPartyAddress
}

// GetAssociatedPartyAddressOk returns a tuple with the AssociatedPartyAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplementaryService) GetAssociatedPartyAddressOk() (*string, bool) {
	if o == nil || isNil(o.AssociatedPartyAddress) {
    return nil, false
	}
	return o.AssociatedPartyAddress, true
}

// HasAssociatedPartyAddress returns a boolean if a field has been set.
func (o *SupplementaryService) HasAssociatedPartyAddress() bool {
	if o != nil && !isNil(o.AssociatedPartyAddress) {
		return true
	}

	return false
}

// SetAssociatedPartyAddress gets a reference to the given string and assigns it to the AssociatedPartyAddress field.
func (o *SupplementaryService) SetAssociatedPartyAddress(v string) {
	o.AssociatedPartyAddress = &v
}

// GetConferenceId returns the ConferenceId field value if set, zero value otherwise.
func (o *SupplementaryService) GetConferenceId() string {
	if o == nil || isNil(o.ConferenceId) {
		var ret string
		return ret
	}
	return *o.ConferenceId
}

// GetConferenceIdOk returns a tuple with the ConferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplementaryService) GetConferenceIdOk() (*string, bool) {
	if o == nil || isNil(o.ConferenceId) {
    return nil, false
	}
	return o.ConferenceId, true
}

// HasConferenceId returns a boolean if a field has been set.
func (o *SupplementaryService) HasConferenceId() bool {
	if o != nil && !isNil(o.ConferenceId) {
		return true
	}

	return false
}

// SetConferenceId gets a reference to the given string and assigns it to the ConferenceId field.
func (o *SupplementaryService) SetConferenceId(v string) {
	o.ConferenceId = &v
}

// GetParticipantActionType returns the ParticipantActionType field value if set, zero value otherwise.
func (o *SupplementaryService) GetParticipantActionType() ParticipantActionType {
	if o == nil || isNil(o.ParticipantActionType) {
		var ret ParticipantActionType
		return ret
	}
	return *o.ParticipantActionType
}

// GetParticipantActionTypeOk returns a tuple with the ParticipantActionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplementaryService) GetParticipantActionTypeOk() (*ParticipantActionType, bool) {
	if o == nil || isNil(o.ParticipantActionType) {
    return nil, false
	}
	return o.ParticipantActionType, true
}

// HasParticipantActionType returns a boolean if a field has been set.
func (o *SupplementaryService) HasParticipantActionType() bool {
	if o != nil && !isNil(o.ParticipantActionType) {
		return true
	}

	return false
}

// SetParticipantActionType gets a reference to the given ParticipantActionType and assigns it to the ParticipantActionType field.
func (o *SupplementaryService) SetParticipantActionType(v ParticipantActionType) {
	o.ParticipantActionType = &v
}

// GetChangeTime returns the ChangeTime field value if set, zero value otherwise.
func (o *SupplementaryService) GetChangeTime() time.Time {
	if o == nil || isNil(o.ChangeTime) {
		var ret time.Time
		return ret
	}
	return *o.ChangeTime
}

// GetChangeTimeOk returns a tuple with the ChangeTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplementaryService) GetChangeTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.ChangeTime) {
    return nil, false
	}
	return o.ChangeTime, true
}

// HasChangeTime returns a boolean if a field has been set.
func (o *SupplementaryService) HasChangeTime() bool {
	if o != nil && !isNil(o.ChangeTime) {
		return true
	}

	return false
}

// SetChangeTime gets a reference to the given time.Time and assigns it to the ChangeTime field.
func (o *SupplementaryService) SetChangeTime(v time.Time) {
	o.ChangeTime = &v
}

// GetNumberOfParticipants returns the NumberOfParticipants field value if set, zero value otherwise.
func (o *SupplementaryService) GetNumberOfParticipants() int32 {
	if o == nil || isNil(o.NumberOfParticipants) {
		var ret int32
		return ret
	}
	return *o.NumberOfParticipants
}

// GetNumberOfParticipantsOk returns a tuple with the NumberOfParticipants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplementaryService) GetNumberOfParticipantsOk() (*int32, bool) {
	if o == nil || isNil(o.NumberOfParticipants) {
    return nil, false
	}
	return o.NumberOfParticipants, true
}

// HasNumberOfParticipants returns a boolean if a field has been set.
func (o *SupplementaryService) HasNumberOfParticipants() bool {
	if o != nil && !isNil(o.NumberOfParticipants) {
		return true
	}

	return false
}

// SetNumberOfParticipants gets a reference to the given int32 and assigns it to the NumberOfParticipants field.
func (o *SupplementaryService) SetNumberOfParticipants(v int32) {
	o.NumberOfParticipants = &v
}

// GetCUGInformation returns the CUGInformation field value if set, zero value otherwise.
func (o *SupplementaryService) GetCUGInformation() string {
	if o == nil || isNil(o.CUGInformation) {
		var ret string
		return ret
	}
	return *o.CUGInformation
}

// GetCUGInformationOk returns a tuple with the CUGInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplementaryService) GetCUGInformationOk() (*string, bool) {
	if o == nil || isNil(o.CUGInformation) {
    return nil, false
	}
	return o.CUGInformation, true
}

// HasCUGInformation returns a boolean if a field has been set.
func (o *SupplementaryService) HasCUGInformation() bool {
	if o != nil && !isNil(o.CUGInformation) {
		return true
	}

	return false
}

// SetCUGInformation gets a reference to the given string and assigns it to the CUGInformation field.
func (o *SupplementaryService) SetCUGInformation(v string) {
	o.CUGInformation = &v
}

func (o SupplementaryService) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SupplementaryServiceType) {
		toSerialize["supplementaryServiceType"] = o.SupplementaryServiceType
	}
	if !isNil(o.SupplementaryServiceMode) {
		toSerialize["supplementaryServiceMode"] = o.SupplementaryServiceMode
	}
	if !isNil(o.NumberOfDiversions) {
		toSerialize["numberOfDiversions"] = o.NumberOfDiversions
	}
	if !isNil(o.AssociatedPartyAddress) {
		toSerialize["associatedPartyAddress"] = o.AssociatedPartyAddress
	}
	if !isNil(o.ConferenceId) {
		toSerialize["conferenceId"] = o.ConferenceId
	}
	if !isNil(o.ParticipantActionType) {
		toSerialize["participantActionType"] = o.ParticipantActionType
	}
	if !isNil(o.ChangeTime) {
		toSerialize["changeTime"] = o.ChangeTime
	}
	if !isNil(o.NumberOfParticipants) {
		toSerialize["numberOfParticipants"] = o.NumberOfParticipants
	}
	if !isNil(o.CUGInformation) {
		toSerialize["cUGInformation"] = o.CUGInformation
	}
	return json.Marshal(toSerialize)
}

type NullableSupplementaryService struct {
	value *SupplementaryService
	isSet bool
}

func (v NullableSupplementaryService) Get() *SupplementaryService {
	return v.value
}

func (v *NullableSupplementaryService) Set(val *SupplementaryService) {
	v.value = val
	v.isSet = true
}

func (v NullableSupplementaryService) IsSet() bool {
	return v.isSet
}

func (v *NullableSupplementaryService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupplementaryService(val *SupplementaryService) *NullableSupplementaryService {
	return &NullableSupplementaryService{value: val, isSet: true}
}

func (v NullableSupplementaryService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupplementaryService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

