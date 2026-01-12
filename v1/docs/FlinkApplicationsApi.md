# \FlinkApplicationsApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrUpdateApplication**](FlinkApplicationsApi.md#CreateOrUpdateApplication) | **Post** /cmf/api/v1/environments/{envName}/applications | Creates a new Flink Application or updates an existing one in the given Environment.
[**DeleteApplication**](FlinkApplicationsApi.md#DeleteApplication) | **Delete** /cmf/api/v1/environments/{envName}/applications/{appName} | Deletes an Application of the given name in the given Environment.
[**GetApplication**](FlinkApplicationsApi.md#GetApplication) | **Get** /cmf/api/v1/environments/{envName}/applications/{appName} | Retrieve an Application of the given name in the given Environment.
[**GetApplicationEvents**](FlinkApplicationsApi.md#GetApplicationEvents) | **Get** /cmf/api/v1alpha1/environments/{envName}/applications/{appName}/events | Get a paginated list of events of the given Application
[**GetApplicationInstance**](FlinkApplicationsApi.md#GetApplicationInstance) | **Get** /cmf/api/v1/environments/{envName}/applications/{appName}/instances/{instName} | Retrieve an Instance of an Application
[**GetApplicationInstances**](FlinkApplicationsApi.md#GetApplicationInstances) | **Get** /cmf/api/v1/environments/{envName}/applications/{appName}/instances | Get a paginated list of instances of the given Application
[**GetApplications**](FlinkApplicationsApi.md#GetApplications) | **Get** /cmf/api/v1/environments/{envName}/applications | Retrieve a paginated list of all applications in the given Environment.
[**StartApplication**](FlinkApplicationsApi.md#StartApplication) | **Post** /cmf/api/v1/environments/{envName}/applications/{appName}/start | Starts an earlier submitted Flink Application
[**SuspendApplication**](FlinkApplicationsApi.md#SuspendApplication) | **Post** /cmf/api/v1/environments/{envName}/applications/{appName}/suspend | Suspends an earlier started Flink Application



## CreateOrUpdateApplication

> FlinkApplication CreateOrUpdateApplication(ctx, envName).FlinkApplication(flinkApplication).Execute()

Creates a new Flink Application or updates an existing one in the given Environment.

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
    flinkApplication := *openapiclient.NewFlinkApplication("ApiVersion_example", "Kind_example", map[string]interface{}(123), map[string]interface{}(123)) // FlinkApplication | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlinkApplicationsApi.CreateOrUpdateApplication(context.Background(), envName).FlinkApplication(flinkApplication).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlinkApplicationsApi.CreateOrUpdateApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrUpdateApplication`: FlinkApplication
    fmt.Fprintf(os.Stdout, "Response from `FlinkApplicationsApi.CreateOrUpdateApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **flinkApplication** | [**FlinkApplication**](FlinkApplication.md) |  | 

### Return type

[**FlinkApplication**](FlinkApplication.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplication

> DeleteApplication(ctx, envName, appName).Execute()

Deletes an Application of the given name in the given Environment.

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlinkApplicationsApi.DeleteApplication(context.Background(), envName, appName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlinkApplicationsApi.DeleteApplication``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationRequest struct via the builder pattern


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


## GetApplication

> FlinkApplication GetApplication(ctx, envName, appName).IncludeResourceInformation(includeResourceInformation).Execute()

Retrieve an Application of the given name in the given Environment.

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
    includeResourceInformation := true // bool | Whether to include resource summary in the response. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlinkApplicationsApi.GetApplication(context.Background(), envName, appName).IncludeResourceInformation(includeResourceInformation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlinkApplicationsApi.GetApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplication`: FlinkApplication
    fmt.Fprintf(os.Stdout, "Response from `FlinkApplicationsApi.GetApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment | 
**appName** | **string** | Name of the Application | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeResourceInformation** | **bool** | Whether to include resource summary in the response. | [default to false]

### Return type

[**FlinkApplication**](FlinkApplication.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationEvents

> EventsPage GetApplicationEvents(ctx, envName, appName).Page(page).Size(size).Sort(sort).Execute()

Get a paginated list of events of the given Application

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
    resp, r, err := api_client.FlinkApplicationsApi.GetApplicationEvents(context.Background(), envName, appName).Page(page).Size(size).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlinkApplicationsApi.GetApplicationEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationEvents`: EventsPage
    fmt.Fprintf(os.Stdout, "Response from `FlinkApplicationsApi.GetApplicationEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment | 
**appName** | **string** | Name of the Application | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Zero-based page index (0..N) | 
 **size** | **int32** | The size of the page to be returned | 
 **sort** | **[]string** | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**EventsPage**](EventsPage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationInstance

> FlinkApplicationInstance GetApplicationInstance(ctx, envName, appName, instName).Execute()

Retrieve an Instance of an Application

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
    instName := "instName_example" // string | Name of the ApplicationInstance

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlinkApplicationsApi.GetApplicationInstance(context.Background(), envName, appName, instName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlinkApplicationsApi.GetApplicationInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationInstance`: FlinkApplicationInstance
    fmt.Fprintf(os.Stdout, "Response from `FlinkApplicationsApi.GetApplicationInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment | 
**appName** | **string** | Name of the Application | 
**instName** | **string** | Name of the ApplicationInstance | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**FlinkApplicationInstance**](FlinkApplicationInstance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationInstances

> ApplicationInstancesPage GetApplicationInstances(ctx, envName, appName).Page(page).Size(size).Sort(sort).Execute()

Get a paginated list of instances of the given Application

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
    resp, r, err := api_client.FlinkApplicationsApi.GetApplicationInstances(context.Background(), envName, appName).Page(page).Size(size).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlinkApplicationsApi.GetApplicationInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationInstances`: ApplicationInstancesPage
    fmt.Fprintf(os.Stdout, "Response from `FlinkApplicationsApi.GetApplicationInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment | 
**appName** | **string** | Name of the Application | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Zero-based page index (0..N) | 
 **size** | **int32** | The size of the page to be returned | 
 **sort** | **[]string** | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**ApplicationInstancesPage**](ApplicationInstancesPage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplications

> ApplicationsPage GetApplications(ctx, envName).Page(page).Size(size).Sort(sort).IncludeResourceInformation(includeResourceInformation).Execute()

Retrieve a paginated list of all applications in the given Environment.

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
    page := int32(56) // int32 | Zero-based page index (0..N) (optional)
    size := int32(56) // int32 | The size of the page to be returned (optional)
    sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. (optional)
    includeResourceInformation := true // bool | Whether to include resource summary in the response. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlinkApplicationsApi.GetApplications(context.Background(), envName).Page(page).Size(size).Sort(sort).IncludeResourceInformation(includeResourceInformation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlinkApplicationsApi.GetApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplications`: ApplicationsPage
    fmt.Fprintf(os.Stdout, "Response from `FlinkApplicationsApi.GetApplications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Zero-based page index (0..N) | 
 **size** | **int32** | The size of the page to be returned | 
 **sort** | **[]string** | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 
 **includeResourceInformation** | **bool** | Whether to include resource summary in the response. | [default to false]

### Return type

[**ApplicationsPage**](ApplicationsPage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartApplication

> FlinkApplication StartApplication(ctx, envName, appName).StartFromSavepointUid(startFromSavepointUid).Execute()

Starts an earlier submitted Flink Application

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
    startFromSavepointUid := "startFromSavepointUid_example" // string | UID of the Savepoint from which the application should be started. This savepoint could belong to the application or can be a deatched savepoint. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlinkApplicationsApi.StartApplication(context.Background(), envName, appName).StartFromSavepointUid(startFromSavepointUid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlinkApplicationsApi.StartApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartApplication`: FlinkApplication
    fmt.Fprintf(os.Stdout, "Response from `FlinkApplicationsApi.StartApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment | 
**appName** | **string** | Name of the Application | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startFromSavepointUid** | **string** | UID of the Savepoint from which the application should be started. This savepoint could belong to the application or can be a deatched savepoint. | 

### Return type

[**FlinkApplication**](FlinkApplication.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuspendApplication

> FlinkApplication SuspendApplication(ctx, envName, appName).Execute()

Suspends an earlier started Flink Application

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlinkApplicationsApi.SuspendApplication(context.Background(), envName, appName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlinkApplicationsApi.SuspendApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SuspendApplication`: FlinkApplication
    fmt.Fprintf(os.Stdout, "Response from `FlinkApplicationsApi.SuspendApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment | 
**appName** | **string** | Name of the Application | 

### Other Parameters

Other parameters are passed through a pointer to a apiSuspendApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FlinkApplication**](FlinkApplication.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

