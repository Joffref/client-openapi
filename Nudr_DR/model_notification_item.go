/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudr_DR

import (
	"encoding/json"
)

// NotificationItem Identifies a data change notification when the change occurs in a fragment (subset of resource data) of a given resource. 
type NotificationItem struct {
	// String providing an URI formatted according to RFC 3986.
	ResourceId string `json:"resourceId"`
	NotifItems []UpdatedItem `json:"notifItems"`
}

// NewNotificationItem instantiates a new NotificationItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationItem(resourceId string, notifItems []UpdatedItem) *NotificationItem {
	this := NotificationItem{}
	this.ResourceId = resourceId
	this.NotifItems = notifItems
	return &this
}

// NewNotificationItemWithDefaults instantiates a new NotificationItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationItemWithDefaults() *NotificationItem {
	this := NotificationItem{}
	return &this
}

// GetResourceId returns the ResourceId field value
func (o *NotificationItem) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *NotificationItem) GetResourceIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *NotificationItem) SetResourceId(v string) {
	o.ResourceId = v
}

// GetNotifItems returns the NotifItems field value
func (o *NotificationItem) GetNotifItems() []UpdatedItem {
	if o == nil {
		var ret []UpdatedItem
		return ret
	}

	return o.NotifItems
}

// GetNotifItemsOk returns a tuple with the NotifItems field value
// and a boolean to check if the value has been set.
func (o *NotificationItem) GetNotifItemsOk() ([]UpdatedItem, bool) {
	if o == nil {
    return nil, false
	}
	return o.NotifItems, true
}

// SetNotifItems sets field value
func (o *NotificationItem) SetNotifItems(v []UpdatedItem) {
	o.NotifItems = v
}

func (o NotificationItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["resourceId"] = o.ResourceId
	}
	if true {
		toSerialize["notifItems"] = o.NotifItems
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationItem struct {
	value *NotificationItem
	isSet bool
}

func (v NullableNotificationItem) Get() *NotificationItem {
	return v.value
}

func (v *NullableNotificationItem) Set(val *NotificationItem) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationItem) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationItem(val *NotificationItem) *NullableNotificationItem {
	return &NullableNotificationItem{value: val, isSet: true}
}

func (v NullableNotificationItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

