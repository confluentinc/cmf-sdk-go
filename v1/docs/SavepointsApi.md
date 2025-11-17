# \SavepointsApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSavepointForFlinkApplication**](SavepointsApi.md#CreateSavepointForFlinkApplication) | **Post** /cmf/api/v1/environments/{envName}/applications/{appName}/savepoints | Creates a new Savepoint for the given Application in the given Environment.
[**CreateSavepointForFlinkStatement**](SavepointsApi.md#CreateSavepointForFlinkStatement) | **Post** /cmf/api/v1/environments/{envName}/statements/{stmtName}/savepoints | Creates a new Savepoint for the given Statement in the given Environment.
[**DeleteSavepointForFlinkApplication**](SavepointsApi.md#DeleteSavepointForFlinkApplication) | **Delete** /cmf/api/v1/environments/{envName}/applications/{appName}/savepoints/{savepointName} | Deletes the Savepoint of the given name for the given Application in the given Environment.
[**DeleteSavepointForFlinkStatement**](SavepointsApi.md#DeleteSavepointForFlinkStatement) | **Delete** /cmf/api/v1/environments/{envName}/statements/{stmtName}/savepoints/{savepointName} | Deletes the Savepoint of the given name for the given Statement in the given Environment.
[**DetachSavepointFromFlinkApplication**](SavepointsApi.md#DetachSavepointFromFlinkApplication) | **Post** /cmf/api/v1/environments/{envName}/applications/{appName}/savepoints/{savepointName}/detach | Detaches the Savepoint of the given name for the given Application in the given Environment.
[**GetSavepointForFlinkApplication**](SavepointsApi.md#GetSavepointForFlinkApplication) | **Get** /cmf/api/v1/environments/{envName}/applications/{appName}/savepoints/{savepointName} | Retrieve the Savepoint of the given name for the given Application in the given Environment.
[**GetSavepointForFlinkStatement**](SavepointsApi.md#GetSavepointForFlinkStatement) | **Get** /cmf/api/v1/environments/{envName}/statements/{stmtName}/savepoints/{savepointName} | Retrieve the Savepoint of the given name for the given Statement in the given Environment.
[**GetSavepointsForFlinkApplication**](SavepointsApi.md#GetSavepointsForFlinkApplication) | **Get** /cmf/api/v1/environments/{envName}/applications/{appName}/savepoints | Retrieve a paginated list of all Savepoints for the given Application in the given Environment.
[**GetSavepointsForFlinkStatement**](SavepointsApi.md#GetSavepointsForFlinkStatement) | **Get** /cmf/api/v1/environments/{envName}/statements/{stmtName}/savepoints | Retrieve a paginated list of all Savepoints for the given Statement in the given Environment.



## CreateSavepointForFlinkApplication

> Savepoint CreateSavepointForFlinkApplication(ctx, envName, appName).Savepoint(savepoint).Execute()

Creates a new Savepoint for the given Application in the given Environment.

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
    envName := "envName_example" // string | Name of the Environment
    appName := "appName_example" // string | Name of the Application
    savepoint := *openapiclient.NewSavepoint("ApiVersion_example", "Kind_example", *openapiclient.NewSavepointMetadata(), *openapiclient.NewSavepointSpec()) // Savepoint |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SavepointsApi.CreateSavepointForFlinkApplication(context.Background(), envName, appName).Savepoint(savepoint).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SavepointsApi.CreateSavepointForFlinkApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSavepointForFlinkApplication`: Savepoint
    fmt.Fprintf(os.Stdout, "Response from `SavepointsApi.CreateSavepointForFlinkApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment | 
**appName** | **string** | Name of the Application | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSavepointForFlinkApplicationRequest struct via the builder pattern


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


## CreateSavepointForFlinkStatement

> Savepoint CreateSavepointForFlinkStatement(ctx, envName, stmtName).Savepoint(savepoint).Execute()

Creates a new Savepoint for the given Statement in the given Environment.

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
    envName := "envName_example" // string | Name of the Environment
    stmtName := "stmtName_example" // string | Name of the Statement
    savepoint := *openapiclient.NewSavepoint("ApiVersion_example", "Kind_example", *openapiclient.NewSavepointMetadata(), *openapiclient.NewSavepointSpec()) // Savepoint |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SavepointsApi.CreateSavepointForFlinkStatement(context.Background(), envName, stmtName).Savepoint(savepoint).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SavepointsApi.CreateSavepointForFlinkStatement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSavepointForFlinkStatement`: Savepoint
    fmt.Fprintf(os.Stdout, "Response from `SavepointsApi.CreateSavepointForFlinkStatement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment | 
**stmtName** | **string** | Name of the Statement | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSavepointForFlinkStatementRequest struct via the builder pattern


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


## DeleteSavepointForFlinkApplication

> DeleteSavepointForFlinkApplication(ctx, envName, appName, savepointName).Force(force).Execute()

Deletes the Savepoint of the given name for the given Application in the given Environment.

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
    envName := "envName_example" // string | Name of the Environment
    appName := "appName_example" // string | Name of the Application
    savepointName := "savepointName_example" // string | Name of the Savepoint to be deleted.
    force := true // bool | If a Savepoint is marked for deletion, it can be force deleted. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SavepointsApi.DeleteSavepointForFlinkApplication(context.Background(), envName, appName, savepointName).Force(force).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SavepointsApi.DeleteSavepointForFlinkApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment | 
**appName** | **string** | Name of the Application | 
**savepointName** | **string** | Name of the Savepoint to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSavepointForFlinkApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **force** | **bool** | If a Savepoint is marked for deletion, it can be force deleted. | [default to false]

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


## DeleteSavepointForFlinkStatement

> DeleteSavepointForFlinkStatement(ctx, envName, stmtName, savepointName).Force(force).Execute()

Deletes the Savepoint of the given name for the given Statement in the given Environment.

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
    envName := "envName_example" // string | Name of the Environment
    stmtName := "stmtName_example" // string | Name of the Statement
    savepointName := "savepointName_example" // string | Name of the Savepoint to be deleted.
    force := true // bool | If a Savepoint is marked for deletion, it can be force deleted. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SavepointsApi.DeleteSavepointForFlinkStatement(context.Background(), envName, stmtName, savepointName).Force(force).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SavepointsApi.DeleteSavepointForFlinkStatement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment | 
**stmtName** | **string** | Name of the Statement | 
**savepointName** | **string** | Name of the Savepoint to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSavepointForFlinkStatementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **force** | **bool** | If a Savepoint is marked for deletion, it can be force deleted. | [default to false]

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


## DetachSavepointFromFlinkApplication

> Savepoint DetachSavepointFromFlinkApplication(ctx, envName, appName, savepointName).Execute()

Detaches the Savepoint of the given name for the given Application in the given Environment.

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
    envName := "envName_example" // string | Name of the Environment
    appName := "appName_example" // string | Name of the Application
    savepointName := "savepointName_example" // string | Name of the Savepoint to be detached.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SavepointsApi.DetachSavepointFromFlinkApplication(context.Background(), envName, appName, savepointName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SavepointsApi.DetachSavepointFromFlinkApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DetachSavepointFromFlinkApplication`: Savepoint
    fmt.Fprintf(os.Stdout, "Response from `SavepointsApi.DetachSavepointFromFlinkApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment | 
**appName** | **string** | Name of the Application | 
**savepointName** | **string** | Name of the Savepoint to be detached. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachSavepointFromFlinkApplicationRequest struct via the builder pattern


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


## GetSavepointForFlinkApplication

> Savepoint GetSavepointForFlinkApplication(ctx, envName, appName, savepointName).Execute()

Retrieve the Savepoint of the given name for the given Application in the given Environment.

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
    envName := "envName_example" // string | Name of the Environment
    appName := "appName_example" // string | Name of the Application
    savepointName := "savepointName_example" // string | Name of the Savepoint

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SavepointsApi.GetSavepointForFlinkApplication(context.Background(), envName, appName, savepointName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SavepointsApi.GetSavepointForFlinkApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSavepointForFlinkApplication`: Savepoint
    fmt.Fprintf(os.Stdout, "Response from `SavepointsApi.GetSavepointForFlinkApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment | 
**appName** | **string** | Name of the Application | 
**savepointName** | **string** | Name of the Savepoint | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSavepointForFlinkApplicationRequest struct via the builder pattern


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


## GetSavepointForFlinkStatement

> Savepoint GetSavepointForFlinkStatement(ctx, envName, stmtName, savepointName).Execute()

Retrieve the Savepoint of the given name for the given Statement in the given Environment.

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
    envName := "envName_example" // string | Name of the Environment
    stmtName := "stmtName_example" // string | Name of the Statement
    savepointName := "savepointName_example" // string | Name of the Savepoint

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SavepointsApi.GetSavepointForFlinkStatement(context.Background(), envName, stmtName, savepointName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SavepointsApi.GetSavepointForFlinkStatement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSavepointForFlinkStatement`: Savepoint
    fmt.Fprintf(os.Stdout, "Response from `SavepointsApi.GetSavepointForFlinkStatement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment | 
**stmtName** | **string** | Name of the Statement | 
**savepointName** | **string** | Name of the Savepoint | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSavepointForFlinkStatementRequest struct via the builder pattern


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


## GetSavepointsForFlinkApplication

> SavepointsPage GetSavepointsForFlinkApplication(ctx, envName, appName).Page(page).Size(size).Sort(sort).Execute()

Retrieve a paginated list of all Savepoints for the given Application in the given Environment.

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
    envName := "envName_example" // string | Name of the Environment
    appName := "appName_example" // string | Name of the Application
    page := int32(56) // int32 | Zero-based page index (0..N) (optional)
    size := int32(56) // int32 | The size of the page to be returned (optional)
    sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SavepointsApi.GetSavepointsForFlinkApplication(context.Background(), envName, appName).Page(page).Size(size).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SavepointsApi.GetSavepointsForFlinkApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSavepointsForFlinkApplication`: SavepointsPage
    fmt.Fprintf(os.Stdout, "Response from `SavepointsApi.GetSavepointsForFlinkApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment | 
**appName** | **string** | Name of the Application | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSavepointsForFlinkApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Zero-based page index (0..N) | 
 **size** | **int32** | The size of the page to be returned | 
 **sort** | **[]string** | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

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


## GetSavepointsForFlinkStatement

> SavepointsPage GetSavepointsForFlinkStatement(ctx, envName, stmtName).Page(page).Size(size).Sort(sort).Execute()

Retrieve a paginated list of all Savepoints for the given Statement in the given Environment.

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
    envName := "envName_example" // string | Name of the Environment
    stmtName := "stmtName_example" // string | Name of the Statement
    page := int32(56) // int32 | Zero-based page index (0..N) (optional)
    size := int32(56) // int32 | The size of the page to be returned (optional)
    sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SavepointsApi.GetSavepointsForFlinkStatement(context.Background(), envName, stmtName).Page(page).Size(size).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SavepointsApi.GetSavepointsForFlinkStatement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSavepointsForFlinkStatement`: SavepointsPage
    fmt.Fprintf(os.Stdout, "Response from `SavepointsApi.GetSavepointsForFlinkStatement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment | 
**stmtName** | **string** | Name of the Statement | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSavepointsForFlinkStatementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Zero-based page index (0..N) | 
 **size** | **int32** | The size of the page to be returned | 
 **sort** | **[]string** | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

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

