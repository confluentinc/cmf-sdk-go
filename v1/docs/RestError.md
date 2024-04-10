# RestError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **int32** | The HTTP status code. | [readonly] 
**Error** | **string** | The short error message. | [readonly] 
**Path** | **string** | The path of the URL for this request. | [readonly] 
**Timestamp** | [**time.Time**](time.Time.md) | The time the error occured. | [readonly] 
**Message** | **string** | The long error message. | [readonly] 
**SchemaValidationErrors** | [**[]ValidationMessage**](ValidationMessage.md) | Validation errors against the OpenAPI schema. | 
**Trace** | **string** | The stacktrace for this error. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


