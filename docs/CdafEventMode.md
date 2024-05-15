# CdafEventMode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Trigger** | [**CdafEventTrigger**](CdafEventTrigger.md) |  | 
**MaxReports** | Pointer to **int32** |  | [optional] 
**Expiry** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**RepPeriod** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**SampRatio** | Pointer to **int32** | Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.   | [optional] 

## Methods

### NewCdafEventMode

`func NewCdafEventMode(trigger CdafEventTrigger, ) *CdafEventMode`

NewCdafEventMode instantiates a new CdafEventMode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdafEventModeWithDefaults

`func NewCdafEventModeWithDefaults() *CdafEventMode`

NewCdafEventModeWithDefaults instantiates a new CdafEventMode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrigger

`func (o *CdafEventMode) GetTrigger() CdafEventTrigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *CdafEventMode) GetTriggerOk() (*CdafEventTrigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *CdafEventMode) SetTrigger(v CdafEventTrigger)`

SetTrigger sets Trigger field to given value.


### GetMaxReports

`func (o *CdafEventMode) GetMaxReports() int32`

GetMaxReports returns the MaxReports field if non-nil, zero value otherwise.

### GetMaxReportsOk

`func (o *CdafEventMode) GetMaxReportsOk() (*int32, bool)`

GetMaxReportsOk returns a tuple with the MaxReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReports

`func (o *CdafEventMode) SetMaxReports(v int32)`

SetMaxReports sets MaxReports field to given value.

### HasMaxReports

`func (o *CdafEventMode) HasMaxReports() bool`

HasMaxReports returns a boolean if a field has been set.

### GetExpiry

`func (o *CdafEventMode) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *CdafEventMode) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *CdafEventMode) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *CdafEventMode) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetRepPeriod

`func (o *CdafEventMode) GetRepPeriod() int32`

GetRepPeriod returns the RepPeriod field if non-nil, zero value otherwise.

### GetRepPeriodOk

`func (o *CdafEventMode) GetRepPeriodOk() (*int32, bool)`

GetRepPeriodOk returns a tuple with the RepPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepPeriod

`func (o *CdafEventMode) SetRepPeriod(v int32)`

SetRepPeriod sets RepPeriod field to given value.

### HasRepPeriod

`func (o *CdafEventMode) HasRepPeriod() bool`

HasRepPeriod returns a boolean if a field has been set.

### GetSampRatio

`func (o *CdafEventMode) GetSampRatio() int32`

GetSampRatio returns the SampRatio field if non-nil, zero value otherwise.

### GetSampRatioOk

`func (o *CdafEventMode) GetSampRatioOk() (*int32, bool)`

GetSampRatioOk returns a tuple with the SampRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampRatio

`func (o *CdafEventMode) SetSampRatio(v int32)`

SetSampRatio sets SampRatio field to given value.

### HasSampRatio

`func (o *CdafEventMode) HasSampRatio() bool`

HasSampRatio returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


