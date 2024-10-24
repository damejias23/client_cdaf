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

// CdafEventTrigger Describes how CDAF should generate the report for the event
type CdafEventTrigger struct {
	CdafEventTriggerAnyOf *CdafEventTriggerAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *CdafEventTrigger) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into CdafEventTriggerAnyOf
	err = json.Unmarshal(data, &dst.CdafEventTriggerAnyOf);
	if err == nil {
		jsonCdafEventTriggerAnyOf, _ := json.Marshal(dst.CdafEventTriggerAnyOf)
		if string(jsonCdafEventTriggerAnyOf) == "{}" { // empty struct
			dst.CdafEventTriggerAnyOf = nil
		} else {
			return nil // data stored in dst.CdafEventTriggerAnyOf, return on the first match
		}
	} else {
		dst.CdafEventTriggerAnyOf = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("Data failed to match schemas in anyOf(CdafEventTrigger)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *CdafEventTrigger) MarshalJSON() ([]byte, error) {
	if src.CdafEventTriggerAnyOf != nil {
		return json.Marshal(&src.CdafEventTriggerAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableCdafEventTrigger struct {
	value *CdafEventTrigger
	isSet bool
}

func (v NullableCdafEventTrigger) Get() *CdafEventTrigger {
	return v.value
}

func (v *NullableCdafEventTrigger) Set(val *CdafEventTrigger) {
	v.value = val
	v.isSet = true
}

func (v NullableCdafEventTrigger) IsSet() bool {
	return v.isSet
}

func (v *NullableCdafEventTrigger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCdafEventTrigger(val *CdafEventTrigger) *NullableCdafEventTrigger {
	return &NullableCdafEventTrigger{value: val, isSet: true}
}

func (v NullableCdafEventTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCdafEventTrigger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


