# CdafEventSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventRequest** | [**CdafEvent**](CdafEvent.md) |  | 
**EventNotifyUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**Options** | Pointer to [**CdafEventMode**](CdafEventMode.md) |  | [optional] 

## Methods

### NewCdafEventSubscription

`func NewCdafEventSubscription(eventRequest CdafEvent, eventNotifyUri string, ) *CdafEventSubscription`

NewCdafEventSubscription instantiates a new CdafEventSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdafEventSubscriptionWithDefaults

`func NewCdafEventSubscriptionWithDefaults() *CdafEventSubscription`

NewCdafEventSubscriptionWithDefaults instantiates a new CdafEventSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventRequest

`func (o *CdafEventSubscription) GetEventRequest() CdafEvent`

GetEventRequest returns the EventRequest field if non-nil, zero value otherwise.

### GetEventRequestOk

`func (o *CdafEventSubscription) GetEventRequestOk() (*CdafEvent, bool)`

GetEventRequestOk returns a tuple with the EventRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventRequest

`func (o *CdafEventSubscription) SetEventRequest(v CdafEvent)`

SetEventRequest sets EventRequest field to given value.


### GetEventNotifyUri

`func (o *CdafEventSubscription) GetEventNotifyUri() string`

GetEventNotifyUri returns the EventNotifyUri field if non-nil, zero value otherwise.

### GetEventNotifyUriOk

`func (o *CdafEventSubscription) GetEventNotifyUriOk() (*string, bool)`

GetEventNotifyUriOk returns a tuple with the EventNotifyUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNotifyUri

`func (o *CdafEventSubscription) SetEventNotifyUri(v string)`

SetEventNotifyUri sets EventNotifyUri field to given value.


### GetOptions

`func (o *CdafEventSubscription) GetOptions() CdafEventMode`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CdafEventSubscription) GetOptionsOk() (*CdafEventMode, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CdafEventSubscription) SetOptions(v CdafEventMode)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *CdafEventSubscription) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


