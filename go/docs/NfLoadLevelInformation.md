# NfLoadLevelInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NfType** | Pointer to [**NFType**](NFType.md) |  | [optional] 
**NfInstanceId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**NfSetId** | Pointer to **string** | NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \&quot;set&lt;Set ID&gt;.&lt;nftype&gt;set.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.&lt;NFType&gt;set.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)  &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NFType&gt; encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters &lt;Set ID&gt; encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.   | [optional] 
**NfCpuUsage** | Pointer to **int32** |  | [optional] 
**NfMemoryUsage** | Pointer to **int32** |  | [optional] 
**NfStorageUsage** | Pointer to **int32** |  | [optional] 

## Methods

### NewNfLoadLevelInformation

`func NewNfLoadLevelInformation() *NfLoadLevelInformation`

NewNfLoadLevelInformation instantiates a new NfLoadLevelInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNfLoadLevelInformationWithDefaults

`func NewNfLoadLevelInformationWithDefaults() *NfLoadLevelInformation`

NewNfLoadLevelInformationWithDefaults instantiates a new NfLoadLevelInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNfType

`func (o *NfLoadLevelInformation) GetNfType() NFType`

GetNfType returns the NfType field if non-nil, zero value otherwise.

### GetNfTypeOk

`func (o *NfLoadLevelInformation) GetNfTypeOk() (*NFType, bool)`

GetNfTypeOk returns a tuple with the NfType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfType

`func (o *NfLoadLevelInformation) SetNfType(v NFType)`

SetNfType sets NfType field to given value.

### HasNfType

`func (o *NfLoadLevelInformation) HasNfType() bool`

HasNfType returns a boolean if a field has been set.

### GetNfInstanceId

`func (o *NfLoadLevelInformation) GetNfInstanceId() string`

GetNfInstanceId returns the NfInstanceId field if non-nil, zero value otherwise.

### GetNfInstanceIdOk

`func (o *NfLoadLevelInformation) GetNfInstanceIdOk() (*string, bool)`

GetNfInstanceIdOk returns a tuple with the NfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfInstanceId

`func (o *NfLoadLevelInformation) SetNfInstanceId(v string)`

SetNfInstanceId sets NfInstanceId field to given value.

### HasNfInstanceId

`func (o *NfLoadLevelInformation) HasNfInstanceId() bool`

HasNfInstanceId returns a boolean if a field has been set.

### GetNfSetId

`func (o *NfLoadLevelInformation) GetNfSetId() string`

GetNfSetId returns the NfSetId field if non-nil, zero value otherwise.

### GetNfSetIdOk

`func (o *NfLoadLevelInformation) GetNfSetIdOk() (*string, bool)`

GetNfSetIdOk returns a tuple with the NfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfSetId

`func (o *NfLoadLevelInformation) SetNfSetId(v string)`

SetNfSetId sets NfSetId field to given value.

### HasNfSetId

`func (o *NfLoadLevelInformation) HasNfSetId() bool`

HasNfSetId returns a boolean if a field has been set.

### GetNfCpuUsage

`func (o *NfLoadLevelInformation) GetNfCpuUsage() int32`

GetNfCpuUsage returns the NfCpuUsage field if non-nil, zero value otherwise.

### GetNfCpuUsageOk

`func (o *NfLoadLevelInformation) GetNfCpuUsageOk() (*int32, bool)`

GetNfCpuUsageOk returns a tuple with the NfCpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfCpuUsage

`func (o *NfLoadLevelInformation) SetNfCpuUsage(v int32)`

SetNfCpuUsage sets NfCpuUsage field to given value.

### HasNfCpuUsage

`func (o *NfLoadLevelInformation) HasNfCpuUsage() bool`

HasNfCpuUsage returns a boolean if a field has been set.

### GetNfMemoryUsage

`func (o *NfLoadLevelInformation) GetNfMemoryUsage() int32`

GetNfMemoryUsage returns the NfMemoryUsage field if non-nil, zero value otherwise.

### GetNfMemoryUsageOk

`func (o *NfLoadLevelInformation) GetNfMemoryUsageOk() (*int32, bool)`

GetNfMemoryUsageOk returns a tuple with the NfMemoryUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfMemoryUsage

`func (o *NfLoadLevelInformation) SetNfMemoryUsage(v int32)`

SetNfMemoryUsage sets NfMemoryUsage field to given value.

### HasNfMemoryUsage

`func (o *NfLoadLevelInformation) HasNfMemoryUsage() bool`

HasNfMemoryUsage returns a boolean if a field has been set.

### GetNfStorageUsage

`func (o *NfLoadLevelInformation) GetNfStorageUsage() int32`

GetNfStorageUsage returns the NfStorageUsage field if non-nil, zero value otherwise.

### GetNfStorageUsageOk

`func (o *NfLoadLevelInformation) GetNfStorageUsageOk() (*int32, bool)`

GetNfStorageUsageOk returns a tuple with the NfStorageUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfStorageUsage

`func (o *NfLoadLevelInformation) SetNfStorageUsage(v int32)`

SetNfStorageUsage sets NfStorageUsage field to given value.

### HasNfStorageUsage

`func (o *NfLoadLevelInformation) HasNfStorageUsage() bool`

HasNfStorageUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


