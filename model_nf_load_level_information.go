/*
Ncdaf_EventExposure

CDAF Event Exposure Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cdaf

import (
	"encoding/json"
)

// NfLoadLevelInformation Represents load level information of a given NF instance.
type NfLoadLevelInformation struct {
	NfType *NFType `json:"nfType,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	NfInstanceId *string `json:"nfInstanceId,omitempty"`
	// NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \"set<Set ID>.<nftype>set.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.<NFType>set.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)  <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NFType> encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters <Set ID> encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.  
	NfSetId *string `json:"nfSetId,omitempty"`
	NfCpuUsage *int32 `json:"nfCpuUsage,omitempty"`
	NfMemoryUsage *int32 `json:"nfMemoryUsage,omitempty"`
	NfStorageUsage *int32 `json:"nfStorageUsage,omitempty"`
	NfStatus *NfStatus `json:"nfStatus,omitempty"`
}

// NewNfLoadLevelInformation instantiates a new NfLoadLevelInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNfLoadLevelInformation() *NfLoadLevelInformation {
	this := NfLoadLevelInformation{}
	return &this
}

// NewNfLoadLevelInformationWithDefaults instantiates a new NfLoadLevelInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNfLoadLevelInformationWithDefaults() *NfLoadLevelInformation {
	this := NfLoadLevelInformation{}
	return &this
}

// GetNfType returns the NfType field value if set, zero value otherwise.
func (o *NfLoadLevelInformation) GetNfType() NFType {
	if o == nil || o.NfType == nil {
		var ret NFType
		return ret
	}
	return *o.NfType
}

// GetNfTypeOk returns a tuple with the NfType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NfLoadLevelInformation) GetNfTypeOk() (*NFType, bool) {
	if o == nil || o.NfType == nil {
		return nil, false
	}
	return o.NfType, true
}

// HasNfType returns a boolean if a field has been set.
func (o *NfLoadLevelInformation) HasNfType() bool {
	if o != nil && o.NfType != nil {
		return true
	}

	return false
}

// SetNfType gets a reference to the given NFType and assigns it to the NfType field.
func (o *NfLoadLevelInformation) SetNfType(v NFType) {
	o.NfType = &v
}

// GetNfInstanceId returns the NfInstanceId field value if set, zero value otherwise.
func (o *NfLoadLevelInformation) GetNfInstanceId() string {
	if o == nil || o.NfInstanceId == nil {
		var ret string
		return ret
	}
	return *o.NfInstanceId
}

// GetNfInstanceIdOk returns a tuple with the NfInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NfLoadLevelInformation) GetNfInstanceIdOk() (*string, bool) {
	if o == nil || o.NfInstanceId == nil {
		return nil, false
	}
	return o.NfInstanceId, true
}

// HasNfInstanceId returns a boolean if a field has been set.
func (o *NfLoadLevelInformation) HasNfInstanceId() bool {
	if o != nil && o.NfInstanceId != nil {
		return true
	}

	return false
}

// SetNfInstanceId gets a reference to the given string and assigns it to the NfInstanceId field.
func (o *NfLoadLevelInformation) SetNfInstanceId(v string) {
	o.NfInstanceId = &v
}

// GetNfSetId returns the NfSetId field value if set, zero value otherwise.
func (o *NfLoadLevelInformation) GetNfSetId() string {
	if o == nil || o.NfSetId == nil {
		var ret string
		return ret
	}
	return *o.NfSetId
}

// GetNfSetIdOk returns a tuple with the NfSetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NfLoadLevelInformation) GetNfSetIdOk() (*string, bool) {
	if o == nil || o.NfSetId == nil {
		return nil, false
	}
	return o.NfSetId, true
}

// HasNfSetId returns a boolean if a field has been set.
func (o *NfLoadLevelInformation) HasNfSetId() bool {
	if o != nil && o.NfSetId != nil {
		return true
	}

	return false
}

// SetNfSetId gets a reference to the given string and assigns it to the NfSetId field.
func (o *NfLoadLevelInformation) SetNfSetId(v string) {
	o.NfSetId = &v
}

// GetNfCpuUsage returns the NfCpuUsage field value if set, zero value otherwise.
func (o *NfLoadLevelInformation) GetNfCpuUsage() int32 {
	if o == nil || o.NfCpuUsage == nil {
		var ret int32
		return ret
	}
	return *o.NfCpuUsage
}

// GetNfCpuUsageOk returns a tuple with the NfCpuUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NfLoadLevelInformation) GetNfCpuUsageOk() (*int32, bool) {
	if o == nil || o.NfCpuUsage == nil {
		return nil, false
	}
	return o.NfCpuUsage, true
}

// HasNfCpuUsage returns a boolean if a field has been set.
func (o *NfLoadLevelInformation) HasNfCpuUsage() bool {
	if o != nil && o.NfCpuUsage != nil {
		return true
	}

	return false
}

// SetNfCpuUsage gets a reference to the given int32 and assigns it to the NfCpuUsage field.
func (o *NfLoadLevelInformation) SetNfCpuUsage(v int32) {
	o.NfCpuUsage = &v
}

// GetNfMemoryUsage returns the NfMemoryUsage field value if set, zero value otherwise.
func (o *NfLoadLevelInformation) GetNfMemoryUsage() int32 {
	if o == nil || o.NfMemoryUsage == nil {
		var ret int32
		return ret
	}
	return *o.NfMemoryUsage
}

// GetNfMemoryUsageOk returns a tuple with the NfMemoryUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NfLoadLevelInformation) GetNfMemoryUsageOk() (*int32, bool) {
	if o == nil || o.NfMemoryUsage == nil {
		return nil, false
	}
	return o.NfMemoryUsage, true
}

// HasNfMemoryUsage returns a boolean if a field has been set.
func (o *NfLoadLevelInformation) HasNfMemoryUsage() bool {
	if o != nil && o.NfMemoryUsage != nil {
		return true
	}

	return false
}

// SetNfMemoryUsage gets a reference to the given int32 and assigns it to the NfMemoryUsage field.
func (o *NfLoadLevelInformation) SetNfMemoryUsage(v int32) {
	o.NfMemoryUsage = &v
}

// GetNfStorageUsage returns the NfStorageUsage field value if set, zero value otherwise.
func (o *NfLoadLevelInformation) GetNfStorageUsage() int32 {
	if o == nil || o.NfStorageUsage == nil {
		var ret int32
		return ret
	}
	return *o.NfStorageUsage
}

// GetNfStorageUsageOk returns a tuple with the NfStorageUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NfLoadLevelInformation) GetNfStorageUsageOk() (*int32, bool) {
	if o == nil || o.NfStorageUsage == nil {
		return nil, false
	}
	return o.NfStorageUsage, true
}

// HasNfStorageUsage returns a boolean if a field has been set.
func (o *NfLoadLevelInformation) HasNfStorageUsage() bool {
	if o != nil && o.NfStorageUsage != nil {
		return true
	}

	return false
}

// SetNfStorageUsage gets a reference to the given int32 and assigns it to the NfStorageUsage field.
func (o *NfLoadLevelInformation) SetNfStorageUsage(v int32) {
	o.NfStorageUsage = &v
}

// GetNfStatus returns the NfStatus field value if set, zero value otherwise.
func (o *NfLoadLevelInformation) GetNfStatus() NfStatus {
	if o == nil || o.NfStatus == nil {
		var ret NfStatus
		return ret
	}
	return *o.NfStatus
}

// GetNfStatusOk returns a tuple with the NfStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NfLoadLevelInformation) GetNfStatusOk() (*NfStatus, bool) {
	if o == nil || o.NfStatus == nil {
		return nil, false
	}
	return o.NfStatus, true
}

// HasNfStatus returns a boolean if a field has been set.
func (o *NfLoadLevelInformation) HasNfStatus() bool {
	if o != nil && o.NfStatus != nil {
		return true
	}

	return false
}

// SetNfStatus gets a reference to the given NfStatus and assigns it to the NfStatus field.
func (o *NfLoadLevelInformation) SetNfStatus(v NfStatus) {
	o.NfStatus = &v
}

func (o NfLoadLevelInformation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NfType != nil {
		toSerialize["nfType"] = o.NfType
	}
	if o.NfInstanceId != nil {
		toSerialize["nfInstanceId"] = o.NfInstanceId
	}
	if o.NfSetId != nil {
		toSerialize["nfSetId"] = o.NfSetId
	}
	if o.NfCpuUsage != nil {
		toSerialize["nfCpuUsage"] = o.NfCpuUsage
	}
	if o.NfMemoryUsage != nil {
		toSerialize["nfMemoryUsage"] = o.NfMemoryUsage
	}
	if o.NfStorageUsage != nil {
		toSerialize["nfStorageUsage"] = o.NfStorageUsage
	}
	if o.NfStatus != nil {
		toSerialize["nfStatus"] = o.NfStatus
	}
	return json.Marshal(toSerialize)
}

type NullableNfLoadLevelInformation struct {
	value *NfLoadLevelInformation
	isSet bool
}

func (v NullableNfLoadLevelInformation) Get() *NfLoadLevelInformation {
	return v.value
}

func (v *NullableNfLoadLevelInformation) Set(val *NfLoadLevelInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableNfLoadLevelInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableNfLoadLevelInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNfLoadLevelInformation(val *NfLoadLevelInformation) *NullableNfLoadLevelInformation {
	return &NullableNfLoadLevelInformation{value: val, isSet: true}
}

func (v NullableNfLoadLevelInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNfLoadLevelInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


