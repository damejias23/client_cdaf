/*
Ncdaf_EventExposure

CDAF Event Exposure Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// CdafEventReport Represents a report triggered by a subscribed event type
type CdafEventReport struct {
	Type CdafEventTypeAnyOf `json:"type"`
	State CdafEventState `json:"state"`
	// string with format 'date-time' as defined in OpenAPI.
	TimeStamp time.Time `json:"timeStamp"`
	// String providing an URI formatted according to RFC 3986.
	SubscriptionId *string `json:"subscriptionId,omitempty"`
	NfLoadLevelInfos []NfLoadLevelInformation `json:"nfLoadLevelInfos,omitempty"`
}

// NewCdafEventReport instantiates a new CdafEventReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCdafEventReport(type_ CdafEventTypeAnyOf, state CdafEventState, timeStamp time.Time) *CdafEventReport {
	this := CdafEventReport{}
	this.Type = type_
	this.State = state
	this.TimeStamp = timeStamp
	return &this
}

// NewCdafEventReportWithDefaults instantiates a new CdafEventReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCdafEventReportWithDefaults() *CdafEventReport {
	this := CdafEventReport{}
	return &this
}

// GetType returns the Type field value
func (o *CdafEventReport) GetType() CdafEventTypeAnyOf {
	if o == nil {
		var ret CdafEventTypeAnyOf
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CdafEventReport) GetTypeOk() (*CdafEventTypeAnyOf, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CdafEventReport) SetType(v CdafEventTypeAnyOf) {
	o.Type = v
}

// GetState returns the State field value
func (o *CdafEventReport) GetState() CdafEventState {
	if o == nil {
		var ret CdafEventState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *CdafEventReport) GetStateOk() (*CdafEventState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *CdafEventReport) SetState(v CdafEventState) {
	o.State = v
}

// GetTimeStamp returns the TimeStamp field value
func (o *CdafEventReport) GetTimeStamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.TimeStamp
}

// GetTimeStampOk returns a tuple with the TimeStamp field value
// and a boolean to check if the value has been set.
func (o *CdafEventReport) GetTimeStampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeStamp, true
}

// SetTimeStamp sets field value
func (o *CdafEventReport) SetTimeStamp(v time.Time) {
	o.TimeStamp = v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *CdafEventReport) GetSubscriptionId() string {
	if o == nil || o.SubscriptionId == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdafEventReport) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || o.SubscriptionId == nil {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *CdafEventReport) HasSubscriptionId() bool {
	if o != nil && o.SubscriptionId != nil {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *CdafEventReport) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetNfLoadLevelInfos returns the NfLoadLevelInfos field value if set, zero value otherwise.
func (o *CdafEventReport) GetNfLoadLevelInfos() []NfLoadLevelInformation {
	if o == nil || o.NfLoadLevelInfos == nil {
		var ret []NfLoadLevelInformation
		return ret
	}
	return o.NfLoadLevelInfos
}

// GetNfLoadLevelInfosOk returns a tuple with the NfLoadLevelInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdafEventReport) GetNfLoadLevelInfosOk() ([]NfLoadLevelInformation, bool) {
	if o == nil || o.NfLoadLevelInfos == nil {
		return nil, false
	}
	return o.NfLoadLevelInfos, true
}

// HasNfLoadLevelInfos returns a boolean if a field has been set.
func (o *CdafEventReport) HasNfLoadLevelInfos() bool {
	if o != nil && o.NfLoadLevelInfos != nil {
		return true
	}

	return false
}

// SetNfLoadLevelInfos gets a reference to the given []NfLoadLevelInformation and assigns it to the NfLoadLevelInfos field.
func (o *CdafEventReport) SetNfLoadLevelInfos(v []NfLoadLevelInformation) {
	o.NfLoadLevelInfos = v
}

func (o CdafEventReport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["timeStamp"] = o.TimeStamp
	}
	if o.SubscriptionId != nil {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if o.NfLoadLevelInfos != nil {
		toSerialize["nfLoadLevelInfos"] = o.NfLoadLevelInfos
	}
	return json.Marshal(toSerialize)
}

type NullableCdafEventReport struct {
	value *CdafEventReport
	isSet bool
}

func (v NullableCdafEventReport) Get() *CdafEventReport {
	return v.value
}

func (v *NullableCdafEventReport) Set(val *CdafEventReport) {
	v.value = val
	v.isSet = true
}

func (v NullableCdafEventReport) IsSet() bool {
	return v.isSet
}

func (v *NullableCdafEventReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCdafEventReport(val *CdafEventReport) *NullableCdafEventReport {
	return &NullableCdafEventReport{value: val, isSet: true}
}

func (v NullableCdafEventReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCdafEventReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


