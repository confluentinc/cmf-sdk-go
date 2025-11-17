# \DetachedSavepointsApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDetachedSavepoint**](DetachedSavepointsApi.md#CreateDetachedSavepoint) | **Post** /cmf/api/v1/detached-savepoints | Creates a new detached savepoint.
[**DeleteDetachedSavepoint**](DetachedSavepointsApi.md#DeleteDetachedSavepoint) | **Delete** /cmf/api/v1/detached-savepoints/{detachedSavepointName} | Deletes the Detached Savepoint of the given name.
[**GetDetachedSavepoint**](DetachedSavepointsApi.md#GetDetachedSavepoint) | **Get** /cmf/api/v1/detached-savepoints/{detachedSavepointName} | Retrieve the Detached Savepoint of the given name.
[**ListDetachedSavepoints**](DetachedSavepointsApi.md#ListDetachedSavepoints) | **Get** /cmf/api/v1/detached-savepoints | Retrieve a paginated list of all Detached Savepoints.



## CreateDetachedSavepoint

> Savepoint CreateDetachedSavepoint(ctx).Savepoint(savepoint).Execute()

Creates a new detached savepoint.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    savepoint := *openapiclient.NewSavepoint("ApiVersion_example", "Kind_example", *openapiclient.NewSavepointMetadata(), *openapiclient.NewSavepointSpec()) // Savepoint |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DetachedSavepointsApi.CreateDetachedSavepoint(context.Background()).Savepoint(savepoint).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DetachedSavepointsApi.CreateDetachedSavepoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDetachedSavepoint`: Savepoint
    fmt.Fprintf(os.Stdout, "Response from `DetachedSavepointsApi.CreateDetachedSavepoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDetachedSavepointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **savepoint** | [**Savepoint**](Savepoint.md) |  | 

### Return type

[**Savepoint**](Savepoint.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDetachedSavepoint

> DeleteDetachedSavepoint(ctx, detachedSavepointName).Execute()

Deletes the Detached Savepoint of the given name.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    detachedSavepointName := "detachedSavepointName_example" // string | Name of the Detached Savepoint

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DetachedSavepointsApi.DeleteDetachedSavepoint(context.Background(), detachedSavepointName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DetachedSavepointsApi.DeleteDetachedSavepoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**detachedSavepointName** | **string** | Name of the Detached Savepoint | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDetachedSavepointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDetachedSavepoint

> Savepoint GetDetachedSavepoint(ctx, detachedSavepointName).Execute()

Retrieve the Detached Savepoint of the given name.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    detachedSavepointName := "detachedSavepointName_example" // string | Name of the Detached Savepoint

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DetachedSavepointsApi.GetDetachedSavepoint(context.Background(), detachedSavepointName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DetachedSavepointsApi.GetDetachedSavepoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDetachedSavepoint`: Savepoint
    fmt.Fprintf(os.Stdout, "Response from `DetachedSavepointsApi.GetDetachedSavepoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**detachedSavepointName** | **string** | Name of the Detached Savepoint | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDetachedSavepointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Savepoint**](Savepoint.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDetachedSavepoints

> SavepointsPage ListDetachedSavepoints(ctx).Page(page).Size(size).Sort(sort).Name(name).Execute()

Retrieve a paginated list of all Detached Savepoints.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    page := int32(56) // int32 | Zero-based page index (0..N) (optional)
    size := int32(56) // int32 | The size of the page to be returned (optional)
    sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. (optional)
    name := "name_example" // string | Filter by detached savepoint name prefix (e.g. ?name=abc) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DetachedSavepointsApi.ListDetachedSavepoints(context.Background()).Page(page).Size(size).Sort(sort).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DetachedSavepointsApi.ListDetachedSavepoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDetachedSavepoints`: SavepointsPage
    fmt.Fprintf(os.Stdout, "Response from `DetachedSavepointsApi.ListDetachedSavepoints`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDetachedSavepointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Zero-based page index (0..N) | 
 **size** | **int32** | The size of the page to be returned | 
 **sort** | **[]string** | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 
 **name** | **string** | Filter by detached savepoint name prefix (e.g. ?name&#x3D;abc) | 

### Return type

[**SavepointsPage**](SavepointsPage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

