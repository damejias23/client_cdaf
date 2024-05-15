/*
Ncdaf_EventExposure

CDAF Event Exposure Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// CdafEventTriggerAnyOf the model 'CdafEventTriggerAnyOf'
type CdafEventTriggerAnyOf string

// List of CdafEventTrigger_anyOf
const (
	ONE_TIME CdafEventTriggerAnyOf = "ONE_TIME"
	CONTINUOUS CdafEventTriggerAnyOf = "CONTINUOUS"
	PERIODIC CdafEventTriggerAnyOf = "PERIODIC"
)

// All allowed values of CdafEventTriggerAnyOf enum
var AllowedCdafEventTriggerAnyOfEnumValues = []CdafEventTriggerAnyOf{
	"ONE_TIME",
	"CONTINUOUS",
	"PERIODIC",
}

func (v *CdafEventTriggerAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CdafEventTriggerAnyOf(value)
	for _, existing := range AllowedCdafEventTriggerAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CdafEventTriggerAnyOf", value)
}

// NewCdafEventTriggerAnyOfFromValue returns a pointer to a valid CdafEventTriggerAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCdafEventTriggerAnyOfFromValue(v string) (*CdafEventTriggerAnyOf, error) {
	ev := CdafEventTriggerAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CdafEventTriggerAnyOf: valid values are %v", v, AllowedCdafEventTriggerAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CdafEventTriggerAnyOf) IsValid() bool {
	for _, existing := range AllowedCdafEventTriggerAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CdafEventTrigger_anyOf value
func (v CdafEventTriggerAnyOf) Ptr() *CdafEventTriggerAnyOf {
	return &v
}

type NullableCdafEventTriggerAnyOf struct {
	value *CdafEventTriggerAnyOf
	isSet bool
}

func (v NullableCdafEventTriggerAnyOf) Get() *CdafEventTriggerAnyOf {
	return v.value
}

func (v *NullableCdafEventTriggerAnyOf) Set(val *CdafEventTriggerAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCdafEventTriggerAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCdafEventTriggerAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCdafEventTriggerAnyOf(val *CdafEventTriggerAnyOf) *NullableCdafEventTriggerAnyOf {
	return &NullableCdafEventTriggerAnyOf{value: val, isSet: true}
}

func (v NullableCdafEventTriggerAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCdafEventTriggerAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

