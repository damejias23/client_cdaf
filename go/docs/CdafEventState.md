# CdafEventState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** |  | 
**RemainReports** | Pointer to **int32** |  | [optional] 
**RemainDuration** | Pointer to **int32** | indicating a time in seconds. | [optional] 

## Methods

### NewCdafEventState

`func NewCdafEventState(active bool, ) *CdafEventState`

NewCdafEventState instantiates a new CdafEventState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdafEventStateWithDefaults

`func NewCdafEventStateWithDefaults() *CdafEventState`

NewCdafEventStateWithDefaults instantiates a new CdafEventState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *CdafEventState) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CdafEventState) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CdafEventState) SetActive(v bool)`

SetActive sets Active field to given value.


### GetRemainReports

`func (o *CdafEventState) GetRemainReports() int32`

GetRemainReports returns the RemainReports field if non-nil, zero value otherwise.

### GetRemainReportsOk

`func (o *CdafEventState) GetRemainReportsOk() (*int32, bool)`

GetRemainReportsOk returns a tuple with the RemainReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainReports

`func (o *CdafEventState) SetRemainReports(v int32)`

SetRemainReports sets RemainReports field to given value.

### HasRemainReports

`func (o *CdafEventState) HasRemainReports() bool`

HasRemainReports returns a boolean if a field has been set.

### GetRemainDuration

`func (o *CdafEventState) GetRemainDuration() int32`

GetRemainDuration returns the RemainDuration field if non-nil, zero value otherwise.

### GetRemainDurationOk

`func (o *CdafEventState) GetRemainDurationOk() (*int32, bool)`

GetRemainDurationOk returns a tuple with the RemainDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainDuration

`func (o *CdafEventState) SetRemainDuration(v int32)`

SetRemainDuration sets RemainDuration field to given value.

### HasRemainDuration

`func (o *CdafEventState) HasRemainDuration() bool`

HasRemainDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


