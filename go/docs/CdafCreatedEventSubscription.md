# CdafCreatedEventSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subscription** | [**CdafEventSubscription**](CdafEventSubscription.md) |  | 
**SubscriptionId** | **string** | String providing an URI formatted according to RFC 3986. | 
**ReportEvent** | Pointer to [**CdafEventReport**](CdafEventReport.md) |  | [optional] 

## Methods

### NewCdafCreatedEventSubscription

`func NewCdafCreatedEventSubscription(subscription CdafEventSubscription, subscriptionId string, ) *CdafCreatedEventSubscription`

NewCdafCreatedEventSubscription instantiates a new CdafCreatedEventSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdafCreatedEventSubscriptionWithDefaults

`func NewCdafCreatedEventSubscriptionWithDefaults() *CdafCreatedEventSubscription`

NewCdafCreatedEventSubscriptionWithDefaults instantiates a new CdafCreatedEventSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscription

`func (o *CdafCreatedEventSubscription) GetSubscription() CdafEventSubscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *CdafCreatedEventSubscription) GetSubscriptionOk() (*CdafEventSubscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *CdafCreatedEventSubscription) SetSubscription(v CdafEventSubscription)`

SetSubscription sets Subscription field to given value.


### GetSubscriptionId

`func (o *CdafCreatedEventSubscription) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *CdafCreatedEventSubscription) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *CdafCreatedEventSubscription) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetReportEvent

`func (o *CdafCreatedEventSubscription) GetReportEvent() CdafEventReport`

GetReportEvent returns the ReportEvent field if non-nil, zero value otherwise.

### GetReportEventOk

`func (o *CdafCreatedEventSubscription) GetReportEventOk() (*CdafEventReport, bool)`

GetReportEventOk returns a tuple with the ReportEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportEvent

`func (o *CdafCreatedEventSubscription) SetReportEvent(v CdafEventReport)`

SetReportEvent sets ReportEvent field to given value.

### HasReportEvent

`func (o *CdafCreatedEventSubscription) HasReportEvent() bool`

HasReportEvent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


