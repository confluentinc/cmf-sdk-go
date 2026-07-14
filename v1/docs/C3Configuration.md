# C3Configuration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**C3ConfigurationStatus**](C3ConfigurationStatus.md) |  | [optional] 

## Methods

### NewC3Configuration

`func NewC3Configuration() *C3Configuration`

NewC3Configuration instantiates a new C3Configuration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC3ConfigurationWithDefaults

`func NewC3ConfigurationWithDefaults() *C3Configuration`

NewC3ConfigurationWithDefaults instantiates a new C3Configuration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *C3Configuration) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *C3Configuration) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *C3Configuration) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *C3Configuration) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetStatus

`func (o *C3Configuration) GetStatus() C3ConfigurationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *C3Configuration) GetStatusOk() (*C3ConfigurationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *C3Configuration) SetStatus(v C3ConfigurationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *C3Configuration) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


