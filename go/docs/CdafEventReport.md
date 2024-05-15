# CdafEventReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CdafEventType**](CdafEventType.md) |  | 
**State** | [**CdafEventState**](CdafEventState.md) |  | 
**TimeStamp** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 
**SubscriptionId** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**NfLoadLevelInfos** | Pointer to [**[]NfLoadLevelInformation**](NfLoadLevelInformation.md) |  | [optional] 

## Methods

### NewCdafEventReport

`func NewCdafEventReport(type_ CdafEventType, state CdafEventState, timeStamp time.Time, ) *CdafEventReport`

NewCdafEventReport instantiates a new CdafEventReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdafEventReportWithDefaults

`func NewCdafEventReportWithDefaults() *CdafEventReport`

NewCdafEventReportWithDefaults instantiates a new CdafEventReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CdafEventReport) GetType() CdafEventType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CdafEventReport) GetTypeOk() (*CdafEventType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CdafEventReport) SetType(v CdafEventType)`

SetType sets Type field to given value.


### GetState

`func (o *CdafEventReport) GetState() CdafEventState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CdafEventReport) GetStateOk() (*CdafEventState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CdafEventReport) SetState(v CdafEventState)`

SetState sets State field to given value.


### GetTimeStamp

`func (o *CdafEventReport) GetTimeStamp() time.Time`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *CdafEventReport) GetTimeStampOk() (*time.Time, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *CdafEventReport) SetTimeStamp(v time.Time)`

SetTimeStamp sets TimeStamp field to given value.


### GetSubscriptionId

`func (o *CdafEventReport) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *CdafEventReport) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *CdafEventReport) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *CdafEventReport) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetNfLoadLevelInfos

`func (o *CdafEventReport) GetNfLoadLevelInfos() []NfLoadLevelInformation`

GetNfLoadLevelInfos returns the NfLoadLevelInfos field if non-nil, zero value otherwise.

### GetNfLoadLevelInfosOk

`func (o *CdafEventReport) GetNfLoadLevelInfosOk() (*[]NfLoadLevelInformation, bool)`

GetNfLoadLevelInfosOk returns a tuple with the NfLoadLevelInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfLoadLevelInfos

`func (o *CdafEventReport) SetNfLoadLevelInfos(v []NfLoadLevelInformation)`

SetNfLoadLevelInfos sets NfLoadLevelInfos field to given value.

### HasNfLoadLevelInfos

`func (o *CdafEventReport) HasNfLoadLevelInfos() bool`

HasNfLoadLevelInfos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


