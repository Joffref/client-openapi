/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ProvMnS

import (
	"encoding/json"
)

// NotifyMoiAttributeValueChanges struct for NotifyMoiAttributeValueChanges
type NotifyMoiAttributeValueChanges struct {
	Href string `json:"href"`
	NotificationId int32 `json:"notificationId"`
	NotificationType NotificationType `json:"notificationType"`
	EventTime time.Time `json:"eventTime"`
	SystemDN string `json:"systemDN"`
	CorrelatedNotifications []CorrelatedNotification `json:"correlatedNotifications,omitempty"`
	AdditionalText *string `json:"additionalText,omitempty"`
	SourceIndicator *SourceIndicator `json:"sourceIndicator,omitempty"`
	// The first array item contains the attribute name value pairs with the new values, and the second array item the attribute name value pairs with the optional old values.
	AttributeListValueChanges []map[string]interface{} `json:"attributeListValueChanges"`
}

// NewNotifyMoiAttributeValueChanges instantiates a new NotifyMoiAttributeValueChanges object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotifyMoiAttributeValueChanges(href string, notificationId int32, notificationType NotificationType, eventTime time.Time, systemDN string, attributeListValueChanges []map[string]interface{}) *NotifyMoiAttributeValueChanges {
	this := NotifyMoiAttributeValueChanges{}
	this.Href = href
	this.NotificationId = notificationId
	this.NotificationType = notificationType
	this.EventTime = eventTime
	this.SystemDN = systemDN
	this.AttributeListValueChanges = attributeListValueChanges
	return &this
}

// NewNotifyMoiAttributeValueChangesWithDefaults instantiates a new NotifyMoiAttributeValueChanges object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotifyMoiAttributeValueChangesWithDefaults() *NotifyMoiAttributeValueChanges {
	this := NotifyMoiAttributeValueChanges{}
	return &this
}

// GetHref returns the Href field value
func (o *NotifyMoiAttributeValueChanges) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *NotifyMoiAttributeValueChanges) GetHrefOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *NotifyMoiAttributeValueChanges) SetHref(v string) {
	o.Href = v
}

// GetNotificationId returns the NotificationId field value
func (o *NotifyMoiAttributeValueChanges) GetNotificationId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NotificationId
}

// GetNotificationIdOk returns a tuple with the NotificationId field value
// and a boolean to check if the value has been set.
func (o *NotifyMoiAttributeValueChanges) GetNotificationIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NotificationId, true
}

// SetNotificationId sets field value
func (o *NotifyMoiAttributeValueChanges) SetNotificationId(v int32) {
	o.NotificationId = v
}

// GetNotificationType returns the NotificationType field value
func (o *NotifyMoiAttributeValueChanges) GetNotificationType() NotificationType {
	if o == nil {
		var ret NotificationType
		return ret
	}

	return o.NotificationType
}

// GetNotificationTypeOk returns a tuple with the NotificationType field value
// and a boolean to check if the value has been set.
func (o *NotifyMoiAttributeValueChanges) GetNotificationTypeOk() (*NotificationType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NotificationType, true
}

// SetNotificationType sets field value
func (o *NotifyMoiAttributeValueChanges) SetNotificationType(v NotificationType) {
	o.NotificationType = v
}

// GetEventTime returns the EventTime field value
func (o *NotifyMoiAttributeValueChanges) GetEventTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EventTime
}

// GetEventTimeOk returns a tuple with the EventTime field value
// and a boolean to check if the value has been set.
func (o *NotifyMoiAttributeValueChanges) GetEventTimeOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.EventTime, true
}

// SetEventTime sets field value
func (o *NotifyMoiAttributeValueChanges) SetEventTime(v time.Time) {
	o.EventTime = v
}

// GetSystemDN returns the SystemDN field value
func (o *NotifyMoiAttributeValueChanges) GetSystemDN() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SystemDN
}

// GetSystemDNOk returns a tuple with the SystemDN field value
// and a boolean to check if the value has been set.
func (o *NotifyMoiAttributeValueChanges) GetSystemDNOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SystemDN, true
}

// SetSystemDN sets field value
func (o *NotifyMoiAttributeValueChanges) SetSystemDN(v string) {
	o.SystemDN = v
}

// GetCorrelatedNotifications returns the CorrelatedNotifications field value if set, zero value otherwise.
func (o *NotifyMoiAttributeValueChanges) GetCorrelatedNotifications() []CorrelatedNotification {
	if o == nil || isNil(o.CorrelatedNotifications) {
		var ret []CorrelatedNotification
		return ret
	}
	return o.CorrelatedNotifications
}

// GetCorrelatedNotificationsOk returns a tuple with the CorrelatedNotifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyMoiAttributeValueChanges) GetCorrelatedNotificationsOk() ([]CorrelatedNotification, bool) {
	if o == nil || isNil(o.CorrelatedNotifications) {
    return nil, false
	}
	return o.CorrelatedNotifications, true
}

// HasCorrelatedNotifications returns a boolean if a field has been set.
func (o *NotifyMoiAttributeValueChanges) HasCorrelatedNotifications() bool {
	if o != nil && !isNil(o.CorrelatedNotifications) {
		return true
	}

	return false
}

// SetCorrelatedNotifications gets a reference to the given []CorrelatedNotification and assigns it to the CorrelatedNotifications field.
func (o *NotifyMoiAttributeValueChanges) SetCorrelatedNotifications(v []CorrelatedNotification) {
	o.CorrelatedNotifications = v
}

// GetAdditionalText returns the AdditionalText field value if set, zero value otherwise.
func (o *NotifyMoiAttributeValueChanges) GetAdditionalText() string {
	if o == nil || isNil(o.AdditionalText) {
		var ret string
		return ret
	}
	return *o.AdditionalText
}

// GetAdditionalTextOk returns a tuple with the AdditionalText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyMoiAttributeValueChanges) GetAdditionalTextOk() (*string, bool) {
	if o == nil || isNil(o.AdditionalText) {
    return nil, false
	}
	return o.AdditionalText, true
}

// HasAdditionalText returns a boolean if a field has been set.
func (o *NotifyMoiAttributeValueChanges) HasAdditionalText() bool {
	if o != nil && !isNil(o.AdditionalText) {
		return true
	}

	return false
}

// SetAdditionalText gets a reference to the given string and assigns it to the AdditionalText field.
func (o *NotifyMoiAttributeValueChanges) SetAdditionalText(v string) {
	o.AdditionalText = &v
}

// GetSourceIndicator returns the SourceIndicator field value if set, zero value otherwise.
func (o *NotifyMoiAttributeValueChanges) GetSourceIndicator() SourceIndicator {
	if o == nil || isNil(o.SourceIndicator) {
		var ret SourceIndicator
		return ret
	}
	return *o.SourceIndicator
}

// GetSourceIndicatorOk returns a tuple with the SourceIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyMoiAttributeValueChanges) GetSourceIndicatorOk() (*SourceIndicator, bool) {
	if o == nil || isNil(o.SourceIndicator) {
    return nil, false
	}
	return o.SourceIndicator, true
}

// HasSourceIndicator returns a boolean if a field has been set.
func (o *NotifyMoiAttributeValueChanges) HasSourceIndicator() bool {
	if o != nil && !isNil(o.SourceIndicator) {
		return true
	}

	return false
}

// SetSourceIndicator gets a reference to the given SourceIndicator and assigns it to the SourceIndicator field.
func (o *NotifyMoiAttributeValueChanges) SetSourceIndicator(v SourceIndicator) {
	o.SourceIndicator = &v
}

// GetAttributeListValueChanges returns the AttributeListValueChanges field value
func (o *NotifyMoiAttributeValueChanges) GetAttributeListValueChanges() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.AttributeListValueChanges
}

// GetAttributeListValueChangesOk returns a tuple with the AttributeListValueChanges field value
// and a boolean to check if the value has been set.
func (o *NotifyMoiAttributeValueChanges) GetAttributeListValueChangesOk() ([]map[string]interface{}, bool) {
	if o == nil {
    return nil, false
	}
	return o.AttributeListValueChanges, true
}

// SetAttributeListValueChanges sets field value
func (o *NotifyMoiAttributeValueChanges) SetAttributeListValueChanges(v []map[string]interface{}) {
	o.AttributeListValueChanges = v
}

func (o NotifyMoiAttributeValueChanges) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["href"] = o.Href
	}
	if true {
		toSerialize["notificationId"] = o.NotificationId
	}
	if true {
		toSerialize["notificationType"] = o.NotificationType
	}
	if true {
		toSerialize["eventTime"] = o.EventTime
	}
	if true {
		toSerialize["systemDN"] = o.SystemDN
	}
	if !isNil(o.CorrelatedNotifications) {
		toSerialize["correlatedNotifications"] = o.CorrelatedNotifications
	}
	if !isNil(o.AdditionalText) {
		toSerialize["additionalText"] = o.AdditionalText
	}
	if !isNil(o.SourceIndicator) {
		toSerialize["sourceIndicator"] = o.SourceIndicator
	}
	if true {
		toSerialize["attributeListValueChanges"] = o.AttributeListValueChanges
	}
	return json.Marshal(toSerialize)
}

type NullableNotifyMoiAttributeValueChanges struct {
	value *NotifyMoiAttributeValueChanges
	isSet bool
}

func (v NullableNotifyMoiAttributeValueChanges) Get() *NotifyMoiAttributeValueChanges {
	return v.value
}

func (v *NullableNotifyMoiAttributeValueChanges) Set(val *NotifyMoiAttributeValueChanges) {
	v.value = val
	v.isSet = true
}

func (v NullableNotifyMoiAttributeValueChanges) IsSet() bool {
	return v.isSet
}

func (v *NullableNotifyMoiAttributeValueChanges) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotifyMoiAttributeValueChanges(val *NotifyMoiAttributeValueChanges) *NullableNotifyMoiAttributeValueChanges {
	return &NullableNotifyMoiAttributeValueChanges{value: val, isSet: true}
}

func (v NullableNotifyMoiAttributeValueChanges) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotifyMoiAttributeValueChanges) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

