# CdafEventSubscriptionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubId** | **string** | String providing an URI formatted according to RFC 3986. | 
**OldSubId** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 

## Methods

### NewCdafEventSubscriptionInfo

`func NewCdafEventSubscriptionInfo(subId string, ) *CdafEventSubscriptionInfo`

NewCdafEventSubscriptionInfo instantiates a new CdafEventSubscriptionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdafEventSubscriptionInfoWithDefaults

`func NewCdafEventSubscriptionInfoWithDefaults() *CdafEventSubscriptionInfo`

NewCdafEventSubscriptionInfoWithDefaults instantiates a new CdafEventSubscriptionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubId

`func (o *CdafEventSubscriptionInfo) GetSubId() string`

GetSubId returns the SubId field if non-nil, zero value otherwise.

### GetSubIdOk

`func (o *CdafEventSubscriptionInfo) GetSubIdOk() (*string, bool)`

GetSubIdOk returns a tuple with the SubId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubId

`func (o *CdafEventSubscriptionInfo) SetSubId(v string)`

SetSubId sets SubId field to given value.


### GetOldSubId

`func (o *CdafEventSubscriptionInfo) GetOldSubId() string`

GetOldSubId returns the OldSubId field if non-nil, zero value otherwise.

### GetOldSubIdOk

`func (o *CdafEventSubscriptionInfo) GetOldSubIdOk() (*string, bool)`

GetOldSubIdOk returns a tuple with the OldSubId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldSubId

`func (o *CdafEventSubscriptionInfo) SetOldSubId(v string)`

SetOldSubId sets OldSubId field to given value.

### HasOldSubId

`func (o *CdafEventSubscriptionInfo) HasOldSubId() bool`

HasOldSubId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


