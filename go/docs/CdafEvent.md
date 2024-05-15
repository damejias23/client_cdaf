# CdafEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CdafEventType**](CdafEventType.md) |  | 
**MaxReports** | Pointer to **int32** |  | [optional] 
**MaxResponseTime** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**MinInterval** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**NextReport** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**NextPeriodicReportTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 

## Methods

### NewCdafEvent

`func NewCdafEvent(type_ CdafEventType, ) *CdafEvent`

NewCdafEvent instantiates a new CdafEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdafEventWithDefaults

`func NewCdafEventWithDefaults() *CdafEvent`

NewCdafEventWithDefaults instantiates a new CdafEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CdafEvent) GetType() CdafEventType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CdafEvent) GetTypeOk() (*CdafEventType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CdafEvent) SetType(v CdafEventType)`

SetType sets Type field to given value.


### GetMaxReports

`func (o *CdafEvent) GetMaxReports() int32`

GetMaxReports returns the MaxReports field if non-nil, zero value otherwise.

### GetMaxReportsOk

`func (o *CdafEvent) GetMaxReportsOk() (*int32, bool)`

GetMaxReportsOk returns a tuple with the MaxReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReports

`func (o *CdafEvent) SetMaxReports(v int32)`

SetMaxReports sets MaxReports field to given value.

### HasMaxReports

`func (o *CdafEvent) HasMaxReports() bool`

HasMaxReports returns a boolean if a field has been set.

### GetMaxResponseTime

`func (o *CdafEvent) GetMaxResponseTime() int32`

GetMaxResponseTime returns the MaxResponseTime field if non-nil, zero value otherwise.

### GetMaxResponseTimeOk

`func (o *CdafEvent) GetMaxResponseTimeOk() (*int32, bool)`

GetMaxResponseTimeOk returns a tuple with the MaxResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResponseTime

`func (o *CdafEvent) SetMaxResponseTime(v int32)`

SetMaxResponseTime sets MaxResponseTime field to given value.

### HasMaxResponseTime

`func (o *CdafEvent) HasMaxResponseTime() bool`

HasMaxResponseTime returns a boolean if a field has been set.

### GetMinInterval

`func (o *CdafEvent) GetMinInterval() int32`

GetMinInterval returns the MinInterval field if non-nil, zero value otherwise.

### GetMinIntervalOk

`func (o *CdafEvent) GetMinIntervalOk() (*int32, bool)`

GetMinIntervalOk returns a tuple with the MinInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinInterval

`func (o *CdafEvent) SetMinInterval(v int32)`

SetMinInterval sets MinInterval field to given value.

### HasMinInterval

`func (o *CdafEvent) HasMinInterval() bool`

HasMinInterval returns a boolean if a field has been set.

### GetNextReport

`func (o *CdafEvent) GetNextReport() time.Time`

GetNextReport returns the NextReport field if non-nil, zero value otherwise.

### GetNextReportOk

`func (o *CdafEvent) GetNextReportOk() (*time.Time, bool)`

GetNextReportOk returns a tuple with the NextReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextReport

`func (o *CdafEvent) SetNextReport(v time.Time)`

SetNextReport sets NextReport field to given value.

### HasNextReport

`func (o *CdafEvent) HasNextReport() bool`

HasNextReport returns a boolean if a field has been set.

### GetNextPeriodicReportTime

`func (o *CdafEvent) GetNextPeriodicReportTime() time.Time`

GetNextPeriodicReportTime returns the NextPeriodicReportTime field if non-nil, zero value otherwise.

### GetNextPeriodicReportTimeOk

`func (o *CdafEvent) GetNextPeriodicReportTimeOk() (*time.Time, bool)`

GetNextPeriodicReportTimeOk returns a tuple with the NextPeriodicReportTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPeriodicReportTime

`func (o *CdafEvent) SetNextPeriodicReportTime(v time.Time)`

SetNextPeriodicReportTime sets NextPeriodicReportTime field to given value.

### HasNextPeriodicReportTime

`func (o *CdafEvent) HasNextPeriodicReportTime() bool`

HasNextPeriodicReportTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


