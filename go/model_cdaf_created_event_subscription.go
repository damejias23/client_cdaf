/*
Ncdaf_EventExposure

CDAF Event Exposure Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// CdafCreatedEventSubscription Data within a create CDAF event subscription response
type CdafCreatedEventSubscription struct {
	Subscription CdafEventSubscription `json:"subscription"`
	// String providing an URI formatted according to RFC 3986.
	SubscriptionId string `json:"subscriptionId"`
	ReportEvent *CdafEventReport `json:"reportEvent,omitempty"`
}

// NewCdafCreatedEventSubscription instantiates a new CdafCreatedEventSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCdafCreatedEventSubscription(subscription CdafEventSubscription, subscriptionId string) *CdafCreatedEventSubscription {
	this := CdafCreatedEventSubscription{}
	this.Subscription = subscription
	this.SubscriptionId = subscriptionId
	return &this
}

// NewCdafCreatedEventSubscriptionWithDefaults instantiates a new CdafCreatedEventSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCdafCreatedEventSubscriptionWithDefaults() *CdafCreatedEventSubscription {
	this := CdafCreatedEventSubscription{}
	return &this
}

// GetSubscription returns the Subscription field value
func (o *CdafCreatedEventSubscription) GetSubscription() CdafEventSubscription {
	if o == nil {
		var ret CdafEventSubscription
		return ret
	}

	return o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value
// and a boolean to check if the value has been set.
func (o *CdafCreatedEventSubscription) GetSubscriptionOk() (*CdafEventSubscription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subscription, true
}

// SetSubscription sets field value
func (o *CdafCreatedEventSubscription) SetSubscription(v CdafEventSubscription) {
	o.Subscription = v
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *CdafCreatedEventSubscription) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *CdafCreatedEventSubscription) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *CdafCreatedEventSubscription) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

// GetReportEvent returns the ReportEvent field value if set, zero value otherwise.
func (o *CdafCreatedEventSubscription) GetReportEvent() CdafEventReport {
	if o == nil || o.ReportEvent == nil {
		var ret CdafEventReport
		return ret
	}
	return *o.ReportEvent
}

// GetReportEventOk returns a tuple with the ReportEvent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdafCreatedEventSubscription) GetReportEventOk() (*CdafEventReport, bool) {
	if o == nil || o.ReportEvent == nil {
		return nil, false
	}
	return o.ReportEvent, true
}

// HasReportEvent returns a boolean if a field has been set.
func (o *CdafCreatedEventSubscription) HasReportEvent() bool {
	if o != nil && o.ReportEvent != nil {
		return true
	}

	return false
}

// SetReportEvent gets a reference to the given CdafEventReport and assigns it to the ReportEvent field.
func (o *CdafCreatedEventSubscription) SetReportEvent(v CdafEventReport) {
	o.ReportEvent = &v
}

func (o CdafCreatedEventSubscription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["subscription"] = o.Subscription
	}
	if true {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if o.ReportEvent != nil {
		toSerialize["reportEvent"] = o.ReportEvent
	}
	return json.Marshal(toSerialize)
}

type NullableCdafCreatedEventSubscription struct {
	value *CdafCreatedEventSubscription
	isSet bool
}

func (v NullableCdafCreatedEventSubscription) Get() *CdafCreatedEventSubscription {
	return v.value
}

func (v *NullableCdafCreatedEventSubscription) Set(val *CdafCreatedEventSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableCdafCreatedEventSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableCdafCreatedEventSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCdafCreatedEventSubscription(val *CdafCreatedEventSubscription) *NullableCdafCreatedEventSubscription {
	return &NullableCdafCreatedEventSubscription{value: val, isSet: true}
}

func (v NullableCdafCreatedEventSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCdafCreatedEventSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


