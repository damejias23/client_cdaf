# CdafEventNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportEvent** | Pointer to [**CdafEventReport**](CdafEventReport.md) |  | [optional] 
**EventSubsSyncInfo** | Pointer to [**CdafEventSubsSyncInfo**](CdafEventSubsSyncInfo.md) |  | [optional] 

## Methods

### NewCdafEventNotification

`func NewCdafEventNotification() *CdafEventNotification`

NewCdafEventNotification instantiates a new CdafEventNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdafEventNotificationWithDefaults

`func NewCdafEventNotificationWithDefaults() *CdafEventNotification`

NewCdafEventNotificationWithDefaults instantiates a new CdafEventNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportEvent

`func (o *CdafEventNotification) GetReportEvent() CdafEventReport`

GetReportEvent returns the ReportEvent field if non-nil, zero value otherwise.

### GetReportEventOk

`func (o *CdafEventNotification) GetReportEventOk() (*CdafEventReport, bool)`

GetReportEventOk returns a tuple with the ReportEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportEvent

`func (o *CdafEventNotification) SetReportEvent(v CdafEventReport)`

SetReportEvent sets ReportEvent field to given value.

### HasReportEvent

`func (o *CdafEventNotification) HasReportEvent() bool`

HasReportEvent returns a boolean if a field has been set.

### GetEventSubsSyncInfo

`func (o *CdafEventNotification) GetEventSubsSyncInfo() CdafEventSubsSyncInfo`

GetEventSubsSyncInfo returns the EventSubsSyncInfo field if non-nil, zero value otherwise.

### GetEventSubsSyncInfoOk

`func (o *CdafEventNotification) GetEventSubsSyncInfoOk() (*CdafEventSubsSyncInfo, bool)`

GetEventSubsSyncInfoOk returns a tuple with the EventSubsSyncInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSubsSyncInfo

`func (o *CdafEventNotification) SetEventSubsSyncInfo(v CdafEventSubsSyncInfo)`

SetEventSubsSyncInfo sets EventSubsSyncInfo field to given value.

### HasEventSubsSyncInfo

`func (o *CdafEventNotification) HasEventSubsSyncInfo() bool`

HasEventSubsSyncInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


