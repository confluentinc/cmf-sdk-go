# SqlAssistantTurn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** |  | 
**Kind** | **string** |  | 
**Metadata** | [**SqlAssistantTurnMetadata**](SqlAssistantTurnMetadata.md) |  | 
**Spec** | [**SqlAssistantTurnSpec**](SqlAssistantTurnSpec.md) |  | 
**Status** | Pointer to [**SqlAssistantTurnStatus**](SqlAssistantTurnStatus.md) |  | [optional] 

## Methods

### NewSqlAssistantTurn

`func NewSqlAssistantTurn(apiVersion string, kind string, metadata SqlAssistantTurnMetadata, spec SqlAssistantTurnSpec, ) *SqlAssistantTurn`

NewSqlAssistantTurn instantiates a new SqlAssistantTurn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlAssistantTurnWithDefaults

`func NewSqlAssistantTurnWithDefaults() *SqlAssistantTurn`

NewSqlAssistantTurnWithDefaults instantiates a new SqlAssistantTurn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *SqlAssistantTurn) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *SqlAssistantTurn) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *SqlAssistantTurn) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *SqlAssistantTurn) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SqlAssistantTurn) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SqlAssistantTurn) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *SqlAssistantTurn) GetMetadata() SqlAssistantTurnMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SqlAssistantTurn) GetMetadataOk() (*SqlAssistantTurnMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SqlAssistantTurn) SetMetadata(v SqlAssistantTurnMetadata)`

SetMetadata sets Metadata field to given value.


### GetSpec

`func (o *SqlAssistantTurn) GetSpec() SqlAssistantTurnSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *SqlAssistantTurn) GetSpecOk() (*SqlAssistantTurnSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *SqlAssistantTurn) SetSpec(v SqlAssistantTurnSpec)`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *SqlAssistantTurn) GetStatus() SqlAssistantTurnStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SqlAssistantTurn) GetStatusOk() (*SqlAssistantTurnStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SqlAssistantTurn) SetStatus(v SqlAssistantTurnStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SqlAssistantTurn) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


