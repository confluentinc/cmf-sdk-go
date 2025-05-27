# Pageable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **int32** |  | [optional] 
**Size** | Pointer to **int32** | The number of items in a page. | [optional] 
**Sort** | Pointer to [**Sort**](Sort.md) |  | [optional] 

## Methods

### NewPageable

`func NewPageable() *Pageable`

NewPageable instantiates a new Pageable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageableWithDefaults

`func NewPageableWithDefaults() *Pageable`

NewPageableWithDefaults instantiates a new Pageable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *Pageable) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *Pageable) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *Pageable) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *Pageable) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetSize

`func (o *Pageable) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Pageable) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Pageable) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *Pageable) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSort

`func (o *Pageable) GetSort() Sort`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *Pageable) GetSortOk() (*Sort, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *Pageable) SetSort(v Sort)`

SetSort sets Sort field to given value.

### HasSort

`func (o *Pageable) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


