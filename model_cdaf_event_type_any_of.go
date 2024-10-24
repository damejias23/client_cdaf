/*
Ncdaf_EventExposure

CDAF Event Exposure Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cdaf

import (
	"encoding/json"
	"fmt"
)

// CdafEventTypeAnyOf the model 'CdafEventTypeAnyOf'
type CdafEventTypeAnyOf string

// List of CdafEventType_anyOf
const (
	REPORT_RESOURCE_USAGE CdafEventTypeAnyOf = "REPORT_RESOURCE_USAGE"
)

// All allowed values of CdafEventTypeAnyOf enum
var AllowedCdafEventTypeAnyOfEnumValues = []CdafEventTypeAnyOf{
	"REPORT_RESOURCE_USAGE",
}

func (v *CdafEventTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CdafEventTypeAnyOf(value)
	for _, existing := range AllowedCdafEventTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CdafEventTypeAnyOf", value)
}

// NewCdafEventTypeAnyOfFromValue returns a pointer to a valid CdafEventTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCdafEventTypeAnyOfFromValue(v string) (*CdafEventTypeAnyOf, error) {
	ev := CdafEventTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CdafEventTypeAnyOf: valid values are %v", v, AllowedCdafEventTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CdafEventTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedCdafEventTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CdafEventType_anyOf value
func (v CdafEventTypeAnyOf) Ptr() *CdafEventTypeAnyOf {
	return &v
}

type NullableCdafEventTypeAnyOf struct {
	value *CdafEventTypeAnyOf
	isSet bool
}

func (v NullableCdafEventTypeAnyOf) Get() *CdafEventTypeAnyOf {
	return v.value
}

func (v *NullableCdafEventTypeAnyOf) Set(val *CdafEventTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCdafEventTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCdafEventTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCdafEventTypeAnyOf(val *CdafEventTypeAnyOf) *NullableCdafEventTypeAnyOf {
	return &NullableCdafEventTypeAnyOf{value: val, isSet: true}
}

func (v NullableCdafEventTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCdafEventTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

