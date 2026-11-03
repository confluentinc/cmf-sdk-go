# \FlinkApplicationsApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrUpdateApplication**](FlinkApplicationsApi.md#CreateOrUpdateApplication) | **Post** /cmf/api/v1/environments/{envName}/applications | Creates a new Flink Application or updates an existing one in the given Environment.
[**DeleteApplication**](FlinkApplicationsApi.md#DeleteApplication) | **Delete** /cmf/api/v1/environments/{envName}/applications/{appName} | Deletes an Application of the given name in the given Environment.
[**GetApplication**](FlinkApplicationsApi.md#GetApplication) | **Get** /cmf/api/v1/environments/{envName}/applications/{appName} | Retrieve an Application of the given name in the given Environment.
[**GetApplicationAutoscaler**](FlinkApplicationsApi.md#GetApplicationAutoscaler) | **Get** /cmf/api/v1alpha1/environments/{envName}/applications/{appName}/autoscaler | Get the current autoscaler state of the given Application
[**GetApplicationEvents**](FlinkApplicationsApi.md#GetApplicationEvents) | **Get** /cmf/api/v1alpha1/environments/{envName}/applications/{appName}/events | Get a paginated list of events of the given Application
[**GetApplicationInstance**](FlinkApplicationsApi.md#GetApplicationInstance) | **Get** /cmf/api/v1/environments/{envName}/applications/{appName}/instances/{instName} | Retrieve an Instance of an Application
[**GetApplicationInstances**](FlinkApplicationsApi.md#GetApplicationInstances) | **Get** /cmf/api/v1/environments/{envName}/applications/{appName}/instances | Get a paginated list of instances of the given Application
[**GetApplicationResourceEvents**](FlinkApplicationsApi.md#GetApplicationResourceEvents) | **Get** /cmf/api/v1/environments/{envName}/applications/{appName}/events | Get a paginated list of events of the given Application
[**GetApplications**](FlinkApplicationsApi.md#GetApplications) | **Get** /cmf/api/v1/environments/{envName}/applications | Retrieve a paginated list of all applications in the given Environment.
[**GetPodLogs**](FlinkApplicationsApi.md#GetPodLogs) | **Get** /cmf/api/v1/environments/{envName}/applications/{appName}/pod-logs | Get a snapshot of the container logs of one of the Application&#39;s pods
[**ListApplicationPods**](FlinkApplicationsApi.md#ListApplicationPods) | **Get** /cmf/api/v1/environments/{envName}/applications/{appName}/pods | List the Kubernetes pods the Application currently runs on
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
    envName := "production" // string | Name of the Environment
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

> DeleteApplication(ctx, envName, appName).Force(force).SkipCleanup(skipCleanup).Execute()

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
    force := true // bool | If set to `true`, the platform deletes the application from CMF metadata only, without requiring Kubernetes cluster connectivity. (optional) (default to false)
    skipCleanup := true // bool | If set to `true`, forcibly clears any finalizers blocking the deletion of the underlying Kubernetes resources. Only use this if the application has remained in a pending-deletion state for longer than expected. This skips any cleanup those finalizers would otherwise have performed (for example, taking a final savepoint), so it should only be used once you've confirmed the resource is genuinely stuck rather than still in the middle of a normal, in-progress cleanup. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlinkApplicationsApi.DeleteApplication(context.Background(), envName, appName).Force(force).SkipCleanup(skipCleanup).Execute()
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


 **force** | **bool** | If set to &#x60;true&#x60;, the platform deletes the application from CMF metadata only, without requiring Kubernetes cluster connectivity. | [default to false]
 **skipCleanup** | **bool** | If set to &#x60;true&#x60;, forcibly clears any finalizers blocking the deletion of the underlying Kubernetes resources. Only use this if the application has remained in a pending-deletion state for longer than expected. This skips any cleanup those finalizers would otherwise have performed (for example, taking a final savepoint), so it should only be used once you&#39;ve confirmed the resource is genuinely stuck rather than still in the middle of a normal, in-progress cleanup. | [default to false]

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


## GetApplicationAutoscaler

> FlinkApplicationAutoscaler GetApplicationAutoscaler(ctx, envName, appName).Execute()

Get the current autoscaler state of the given Application



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
    resp, r, err := api_client.FlinkApplicationsApi.GetApplicationAutoscaler(context.Background(), envName, appName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlinkApplicationsApi.GetApplicationAutoscaler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationAutoscaler`: FlinkApplicationAutoscaler
    fmt.Fprintf(os.Stdout, "Response from `FlinkApplicationsApi.GetApplicationAutoscaler`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment | 
**appName** | **string** | Name of the Application | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationAutoscalerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FlinkApplicationAutoscaler**](FlinkApplicationAutoscaler.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationEvents

> EventsPage GetApplicationEvents(ctx, envName, appName).Page(page).Size(size).Sort(sort).Filter(filter).Search(search).SearchScope(searchScope).Fields(fields).Execute()

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
    filter := "filter_example" // string | Filter query string with comma-separated expressions. Supports: - Name filtering: name=foo*bar (wildcards allowed) - Label equality: labels.key = value or labels.key != value - Label set-based: labels.key in (value1, value2) or labels.key notin (value1, value2) - Label existence: labels.key (exists) or !labels.key (does not exist) - State filtering (Applications only): state=RUNNING or state in (RUNNING, FAILED) or state notin (RUNNING, FAILED) - Status error filtering (Applications only): status.error (the Application's status reports an error, i.e. status.error is set and not blank) or !status.error (it reports none). An Application can report an error while its state is not FAILED, for example while its job keeps restarting. - Phase filtering (Statements and ComputePools): phase=PENDING or phase in (PENDING, RUNNING) or phase notin (PENDING, RUNNING) - Type filtering (Events only): type=CMF_STATUS or type in (CMF_STATUS, JOB_STATUS) or type notin (CMF_STATUS, JOB_STATUS) - Severity filtering (Events only): severity=WARNING or severity in (INFO, WARNING) or severity notin (WARNING). Values: INFO, WARNING. Severity is independent of type, so the two can be combined to narrow further (e.g. type=JOB_STATUS,severity=WARNING). An unrecognized severity value is rejected with HTTP 400. - Cluster status filtering (Environments only): clusterStatus=CONNECTED or clusterStatus in (CONNECTED, DISCONNECTED) or clusterStatus notin (CONNECTED, DISCONNECTED). Values match the effective state of the backing Kubernetes cluster (CONNECTED, DISCONNECTED, DECOMMISSIONED). - Source filtering (Savepoints only): source=MANUAL or source in (MANUAL, SCHEDULE) or source notin (UPGRADE). Identifies how the savepoint was created. Values: MANUAL (triggered via API, CLI, or UI), UPGRADE (created during an application upgrade), SCHEDULE (created by a periodic savepoint schedule). An unrecognized source value is rejected with HTTP 400 (unlike state, which simply returns no matches for an unknown value). Example: ?filter=name=foo*bar,labels.environment in (production, qa),!labels.development Example (with state): ?filter=name=prod*,state in (RUNNING, FAILED) Example (with status.error): ?filter=status.error,labels.team=payments Example (with phase): ?filter=name=my-stmt*,phase in (PENDING, RUNNING) Example (with type): ?filter=type=CMF_STATUS or ?filter=type in (CMF_STATUS, JOB_STATUS) Example (with severity): ?filter=severity=WARNING or ?filter=type=JOB_STATUS,severity=WARNING Example (with clusterStatus): ?filter=clusterStatus in (CONNECTED, DISCONNECTED) Example (with source): ?filter=state=COMPLETED,source=SCHEDULE (optional)
    search := "search_example" // string | Search term to match against fields specified in searchScope. Note: Both search and searchScope must be provided together. If only one is provided, the request will be rejected. Example: ?search=foo&searchScope=name,kubernetesNamespace (optional)
    searchScope := "searchScope_example" // string | Comma-separated list of fields to search in. Must be provided together with the search parameter. Unsupported field names will result in a 400 Bad Request. For Environments: supported fields are name, kubernetesNamespace. For Statements: supported fields are name, statement. For Events: supported fields are message; application events also support flinkApplicationInstance. For Secrets: supported fields are name, environments. For Kafka Catalogs: supported fields are name. When multiple fields are specified, the search uses OR logic. Example (Environments): ?search=foo&searchScope=name,kubernetesNamespace means (name contains foo OR kubernetesNamespace contains foo) Example (Statements): ?search=SELECT&searchScope=name,statement means (name contains SELECT OR statement contains SELECT) Example (Application events): ?search=RUNNING&searchScope=message,flinkApplicationInstance means (message contains RUNNING OR flinkApplicationInstance equals RUNNING) (optional)
    fields := "fields_example" // string | Comma-separated list of field paths to include in the response. Supports nested fields using dot notation. Always includes apiVersion and kind fields even if not explicitly requested. Example: ?fields=metadata.name,metadata.createdTimestamp,status.phase (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlinkApplicationsApi.GetApplicationEvents(context.Background(), envName, appName).Page(page).Size(size).Sort(sort).Filter(filter).Search(search).SearchScope(searchScope).Fields(fields).Execute()
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
 **filter** | **string** | Filter query string with comma-separated expressions. Supports: - Name filtering: name&#x3D;foo*bar (wildcards allowed) - Label equality: labels.key &#x3D; value or labels.key !&#x3D; value - Label set-based: labels.key in (value1, value2) or labels.key notin (value1, value2) - Label existence: labels.key (exists) or !labels.key (does not exist) - State filtering (Applications only): state&#x3D;RUNNING or state in (RUNNING, FAILED) or state notin (RUNNING, FAILED) - Status error filtering (Applications only): status.error (the Application&#39;s status reports an error, i.e. status.error is set and not blank) or !status.error (it reports none). An Application can report an error while its state is not FAILED, for example while its job keeps restarting. - Phase filtering (Statements and ComputePools): phase&#x3D;PENDING or phase in (PENDING, RUNNING) or phase notin (PENDING, RUNNING) - Type filtering (Events only): type&#x3D;CMF_STATUS or type in (CMF_STATUS, JOB_STATUS) or type notin (CMF_STATUS, JOB_STATUS) - Severity filtering (Events only): severity&#x3D;WARNING or severity in (INFO, WARNING) or severity notin (WARNING). Values: INFO, WARNING. Severity is independent of type, so the two can be combined to narrow further (e.g. type&#x3D;JOB_STATUS,severity&#x3D;WARNING). An unrecognized severity value is rejected with HTTP 400. - Cluster status filtering (Environments only): clusterStatus&#x3D;CONNECTED or clusterStatus in (CONNECTED, DISCONNECTED) or clusterStatus notin (CONNECTED, DISCONNECTED). Values match the effective state of the backing Kubernetes cluster (CONNECTED, DISCONNECTED, DECOMMISSIONED). - Source filtering (Savepoints only): source&#x3D;MANUAL or source in (MANUAL, SCHEDULE) or source notin (UPGRADE). Identifies how the savepoint was created. Values: MANUAL (triggered via API, CLI, or UI), UPGRADE (created during an application upgrade), SCHEDULE (created by a periodic savepoint schedule). An unrecognized source value is rejected with HTTP 400 (unlike state, which simply returns no matches for an unknown value). Example: ?filter&#x3D;name&#x3D;foo*bar,labels.environment in (production, qa),!labels.development Example (with state): ?filter&#x3D;name&#x3D;prod*,state in (RUNNING, FAILED) Example (with status.error): ?filter&#x3D;status.error,labels.team&#x3D;payments Example (with phase): ?filter&#x3D;name&#x3D;my-stmt*,phase in (PENDING, RUNNING) Example (with type): ?filter&#x3D;type&#x3D;CMF_STATUS or ?filter&#x3D;type in (CMF_STATUS, JOB_STATUS) Example (with severity): ?filter&#x3D;severity&#x3D;WARNING or ?filter&#x3D;type&#x3D;JOB_STATUS,severity&#x3D;WARNING Example (with clusterStatus): ?filter&#x3D;clusterStatus in (CONNECTED, DISCONNECTED) Example (with source): ?filter&#x3D;state&#x3D;COMPLETED,source&#x3D;SCHEDULE | 
 **search** | **string** | Search term to match against fields specified in searchScope. Note: Both search and searchScope must be provided together. If only one is provided, the request will be rejected. Example: ?search&#x3D;foo&amp;searchScope&#x3D;name,kubernetesNamespace | 
 **searchScope** | **string** | Comma-separated list of fields to search in. Must be provided together with the search parameter. Unsupported field names will result in a 400 Bad Request. For Environments: supported fields are name, kubernetesNamespace. For Statements: supported fields are name, statement. For Events: supported fields are message; application events also support flinkApplicationInstance. For Secrets: supported fields are name, environments. For Kafka Catalogs: supported fields are name. When multiple fields are specified, the search uses OR logic. Example (Environments): ?search&#x3D;foo&amp;searchScope&#x3D;name,kubernetesNamespace means (name contains foo OR kubernetesNamespace contains foo) Example (Statements): ?search&#x3D;SELECT&amp;searchScope&#x3D;name,statement means (name contains SELECT OR statement contains SELECT) Example (Application events): ?search&#x3D;RUNNING&amp;searchScope&#x3D;message,flinkApplicationInstance means (message contains RUNNING OR flinkApplicationInstance equals RUNNING) | 
 **fields** | **string** | Comma-separated list of field paths to include in the response. Supports nested fields using dot notation. Always includes apiVersion and kind fields even if not explicitly requested. Example: ?fields&#x3D;metadata.name,metadata.createdTimestamp,status.phase | 

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


## GetApplicationResourceEvents

> ResourceEventsPage GetApplicationResourceEvents(ctx, envName, appName).Page(page).Size(size).Sort(sort).Filter(filter).Search(search).SearchScope(searchScope).Fields(fields).Execute()

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
    filter := "filter_example" // string | Filter query string with comma-separated expressions. Supports: - Name filtering: name=foo*bar (wildcards allowed) - Label equality: labels.key = value or labels.key != value - Label set-based: labels.key in (value1, value2) or labels.key notin (value1, value2) - Label existence: labels.key (exists) or !labels.key (does not exist) - State filtering (Applications only): state=RUNNING or state in (RUNNING, FAILED) or state notin (RUNNING, FAILED) - Status error filtering (Applications only): status.error (the Application's status reports an error, i.e. status.error is set and not blank) or !status.error (it reports none). An Application can report an error while its state is not FAILED, for example while its job keeps restarting. - Phase filtering (Statements and ComputePools): phase=PENDING or phase in (PENDING, RUNNING) or phase notin (PENDING, RUNNING) - Type filtering (Events only): type=CMF_STATUS or type in (CMF_STATUS, JOB_STATUS) or type notin (CMF_STATUS, JOB_STATUS) - Severity filtering (Events only): severity=WARNING or severity in (INFO, WARNING) or severity notin (WARNING). Values: INFO, WARNING. Severity is independent of type, so the two can be combined to narrow further (e.g. type=JOB_STATUS,severity=WARNING). An unrecognized severity value is rejected with HTTP 400. - Cluster status filtering (Environments only): clusterStatus=CONNECTED or clusterStatus in (CONNECTED, DISCONNECTED) or clusterStatus notin (CONNECTED, DISCONNECTED). Values match the effective state of the backing Kubernetes cluster (CONNECTED, DISCONNECTED, DECOMMISSIONED). - Source filtering (Savepoints only): source=MANUAL or source in (MANUAL, SCHEDULE) or source notin (UPGRADE). Identifies how the savepoint was created. Values: MANUAL (triggered via API, CLI, or UI), UPGRADE (created during an application upgrade), SCHEDULE (created by a periodic savepoint schedule). An unrecognized source value is rejected with HTTP 400 (unlike state, which simply returns no matches for an unknown value). Example: ?filter=name=foo*bar,labels.environment in (production, qa),!labels.development Example (with state): ?filter=name=prod*,state in (RUNNING, FAILED) Example (with status.error): ?filter=status.error,labels.team=payments Example (with phase): ?filter=name=my-stmt*,phase in (PENDING, RUNNING) Example (with type): ?filter=type=CMF_STATUS or ?filter=type in (CMF_STATUS, JOB_STATUS) Example (with severity): ?filter=severity=WARNING or ?filter=type=JOB_STATUS,severity=WARNING Example (with clusterStatus): ?filter=clusterStatus in (CONNECTED, DISCONNECTED) Example (with source): ?filter=state=COMPLETED,source=SCHEDULE (optional)
    search := "search_example" // string | Search term to match against fields specified in searchScope. Note: Both search and searchScope must be provided together. If only one is provided, the request will be rejected. Example: ?search=foo&searchScope=name,kubernetesNamespace (optional)
    searchScope := "searchScope_example" // string | Comma-separated list of fields to search in. Must be provided together with the search parameter. Unsupported field names will result in a 400 Bad Request. For Environments: supported fields are name, kubernetesNamespace. For Statements: supported fields are name, statement. For Events: supported fields are message; application events also support flinkApplicationInstance. For Secrets: supported fields are name, environments. For Kafka Catalogs: supported fields are name. When multiple fields are specified, the search uses OR logic. Example (Environments): ?search=foo&searchScope=name,kubernetesNamespace means (name contains foo OR kubernetesNamespace contains foo) Example (Statements): ?search=SELECT&searchScope=name,statement means (name contains SELECT OR statement contains SELECT) Example (Application events): ?search=RUNNING&searchScope=message,flinkApplicationInstance means (message contains RUNNING OR flinkApplicationInstance equals RUNNING) (optional)
    fields := "fields_example" // string | Comma-separated list of field paths to include in the response. Supports nested fields using dot notation. Always includes apiVersion and kind fields even if not explicitly requested. Example: ?fields=metadata.name,metadata.createdTimestamp,status.phase (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlinkApplicationsApi.GetApplicationResourceEvents(context.Background(), envName, appName).Page(page).Size(size).Sort(sort).Filter(filter).Search(search).SearchScope(searchScope).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlinkApplicationsApi.GetApplicationResourceEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationResourceEvents`: ResourceEventsPage
    fmt.Fprintf(os.Stdout, "Response from `FlinkApplicationsApi.GetApplicationResourceEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment | 
**appName** | **string** | Name of the Application | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationResourceEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Zero-based page index (0..N) | 
 **size** | **int32** | The size of the page to be returned | 
 **sort** | **[]string** | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 
 **filter** | **string** | Filter query string with comma-separated expressions. Supports: - Name filtering: name&#x3D;foo*bar (wildcards allowed) - Label equality: labels.key &#x3D; value or labels.key !&#x3D; value - Label set-based: labels.key in (value1, value2) or labels.key notin (value1, value2) - Label existence: labels.key (exists) or !labels.key (does not exist) - State filtering (Applications only): state&#x3D;RUNNING or state in (RUNNING, FAILED) or state notin (RUNNING, FAILED) - Status error filtering (Applications only): status.error (the Application&#39;s status reports an error, i.e. status.error is set and not blank) or !status.error (it reports none). An Application can report an error while its state is not FAILED, for example while its job keeps restarting. - Phase filtering (Statements and ComputePools): phase&#x3D;PENDING or phase in (PENDING, RUNNING) or phase notin (PENDING, RUNNING) - Type filtering (Events only): type&#x3D;CMF_STATUS or type in (CMF_STATUS, JOB_STATUS) or type notin (CMF_STATUS, JOB_STATUS) - Severity filtering (Events only): severity&#x3D;WARNING or severity in (INFO, WARNING) or severity notin (WARNING). Values: INFO, WARNING. Severity is independent of type, so the two can be combined to narrow further (e.g. type&#x3D;JOB_STATUS,severity&#x3D;WARNING). An unrecognized severity value is rejected with HTTP 400. - Cluster status filtering (Environments only): clusterStatus&#x3D;CONNECTED or clusterStatus in (CONNECTED, DISCONNECTED) or clusterStatus notin (CONNECTED, DISCONNECTED). Values match the effective state of the backing Kubernetes cluster (CONNECTED, DISCONNECTED, DECOMMISSIONED). - Source filtering (Savepoints only): source&#x3D;MANUAL or source in (MANUAL, SCHEDULE) or source notin (UPGRADE). Identifies how the savepoint was created. Values: MANUAL (triggered via API, CLI, or UI), UPGRADE (created during an application upgrade), SCHEDULE (created by a periodic savepoint schedule). An unrecognized source value is rejected with HTTP 400 (unlike state, which simply returns no matches for an unknown value). Example: ?filter&#x3D;name&#x3D;foo*bar,labels.environment in (production, qa),!labels.development Example (with state): ?filter&#x3D;name&#x3D;prod*,state in (RUNNING, FAILED) Example (with status.error): ?filter&#x3D;status.error,labels.team&#x3D;payments Example (with phase): ?filter&#x3D;name&#x3D;my-stmt*,phase in (PENDING, RUNNING) Example (with type): ?filter&#x3D;type&#x3D;CMF_STATUS or ?filter&#x3D;type in (CMF_STATUS, JOB_STATUS) Example (with severity): ?filter&#x3D;severity&#x3D;WARNING or ?filter&#x3D;type&#x3D;JOB_STATUS,severity&#x3D;WARNING Example (with clusterStatus): ?filter&#x3D;clusterStatus in (CONNECTED, DISCONNECTED) Example (with source): ?filter&#x3D;state&#x3D;COMPLETED,source&#x3D;SCHEDULE | 
 **search** | **string** | Search term to match against fields specified in searchScope. Note: Both search and searchScope must be provided together. If only one is provided, the request will be rejected. Example: ?search&#x3D;foo&amp;searchScope&#x3D;name,kubernetesNamespace | 
 **searchScope** | **string** | Comma-separated list of fields to search in. Must be provided together with the search parameter. Unsupported field names will result in a 400 Bad Request. For Environments: supported fields are name, kubernetesNamespace. For Statements: supported fields are name, statement. For Events: supported fields are message; application events also support flinkApplicationInstance. For Secrets: supported fields are name, environments. For Kafka Catalogs: supported fields are name. When multiple fields are specified, the search uses OR logic. Example (Environments): ?search&#x3D;foo&amp;searchScope&#x3D;name,kubernetesNamespace means (name contains foo OR kubernetesNamespace contains foo) Example (Statements): ?search&#x3D;SELECT&amp;searchScope&#x3D;name,statement means (name contains SELECT OR statement contains SELECT) Example (Application events): ?search&#x3D;RUNNING&amp;searchScope&#x3D;message,flinkApplicationInstance means (message contains RUNNING OR flinkApplicationInstance equals RUNNING) | 
 **fields** | **string** | Comma-separated list of field paths to include in the response. Supports nested fields using dot notation. Always includes apiVersion and kind fields even if not explicitly requested. Example: ?fields&#x3D;metadata.name,metadata.createdTimestamp,status.phase | 

### Return type

[**ResourceEventsPage**](ResourceEventsPage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplications

> ApplicationsPage GetApplications(ctx, envName).Page(page).Size(size).Sort(sort).IncludeResourceInformation(includeResourceInformation).Filter(filter).Fields(fields).Execute()

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
    sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Supported properties: name, metadata.name, creationTimestamp, metadata.creationTimestamp, updateTimestamp, metadata.updateTimestamp. Note: updateTimestamp sorts by the most recent spec update time. The displayed metadata.updateTimestamp may differ as it reflects the latest of spec or status changes. (optional)
    includeResourceInformation := true // bool | Whether to include resource summary in the response. (optional) (default to false)
    filter := "filter_example" // string | Filter query string with comma-separated expressions. Supports: - Name filtering: name=foo*bar (wildcards allowed) - Label equality: labels.key = value or labels.key != value - Label set-based: labels.key in (value1, value2) or labels.key notin (value1, value2) - Label existence: labels.key (exists) or !labels.key (does not exist) - State filtering (Applications only): state=RUNNING or state in (RUNNING, FAILED) or state notin (RUNNING, FAILED) - Status error filtering (Applications only): status.error (the Application's status reports an error, i.e. status.error is set and not blank) or !status.error (it reports none). An Application can report an error while its state is not FAILED, for example while its job keeps restarting. - Phase filtering (Statements and ComputePools): phase=PENDING or phase in (PENDING, RUNNING) or phase notin (PENDING, RUNNING) - Type filtering (Events only): type=CMF_STATUS or type in (CMF_STATUS, JOB_STATUS) or type notin (CMF_STATUS, JOB_STATUS) - Severity filtering (Events only): severity=WARNING or severity in (INFO, WARNING) or severity notin (WARNING). Values: INFO, WARNING. Severity is independent of type, so the two can be combined to narrow further (e.g. type=JOB_STATUS,severity=WARNING). An unrecognized severity value is rejected with HTTP 400. - Cluster status filtering (Environments only): clusterStatus=CONNECTED or clusterStatus in (CONNECTED, DISCONNECTED) or clusterStatus notin (CONNECTED, DISCONNECTED). Values match the effective state of the backing Kubernetes cluster (CONNECTED, DISCONNECTED, DECOMMISSIONED). - Source filtering (Savepoints only): source=MANUAL or source in (MANUAL, SCHEDULE) or source notin (UPGRADE). Identifies how the savepoint was created. Values: MANUAL (triggered via API, CLI, or UI), UPGRADE (created during an application upgrade), SCHEDULE (created by a periodic savepoint schedule). An unrecognized source value is rejected with HTTP 400 (unlike state, which simply returns no matches for an unknown value). Example: ?filter=name=foo*bar,labels.environment in (production, qa),!labels.development Example (with state): ?filter=name=prod*,state in (RUNNING, FAILED) Example (with status.error): ?filter=status.error,labels.team=payments Example (with phase): ?filter=name=my-stmt*,phase in (PENDING, RUNNING) Example (with type): ?filter=type=CMF_STATUS or ?filter=type in (CMF_STATUS, JOB_STATUS) Example (with severity): ?filter=severity=WARNING or ?filter=type=JOB_STATUS,severity=WARNING Example (with clusterStatus): ?filter=clusterStatus in (CONNECTED, DISCONNECTED) Example (with source): ?filter=state=COMPLETED,source=SCHEDULE (optional)
    fields := "fields_example" // string | Comma-separated list of field paths to include in the response. Supports nested fields using dot notation. Always includes apiVersion and kind fields even if not explicitly requested. Example: ?fields=metadata.name,metadata.createdTimestamp,status.phase (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlinkApplicationsApi.GetApplications(context.Background(), envName).Page(page).Size(size).Sort(sort).IncludeResourceInformation(includeResourceInformation).Filter(filter).Fields(fields).Execute()
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
 **sort** | **[]string** | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Supported properties: name, metadata.name, creationTimestamp, metadata.creationTimestamp, updateTimestamp, metadata.updateTimestamp. Note: updateTimestamp sorts by the most recent spec update time. The displayed metadata.updateTimestamp may differ as it reflects the latest of spec or status changes. | 
 **includeResourceInformation** | **bool** | Whether to include resource summary in the response. | [default to false]
 **filter** | **string** | Filter query string with comma-separated expressions. Supports: - Name filtering: name&#x3D;foo*bar (wildcards allowed) - Label equality: labels.key &#x3D; value or labels.key !&#x3D; value - Label set-based: labels.key in (value1, value2) or labels.key notin (value1, value2) - Label existence: labels.key (exists) or !labels.key (does not exist) - State filtering (Applications only): state&#x3D;RUNNING or state in (RUNNING, FAILED) or state notin (RUNNING, FAILED) - Status error filtering (Applications only): status.error (the Application&#39;s status reports an error, i.e. status.error is set and not blank) or !status.error (it reports none). An Application can report an error while its state is not FAILED, for example while its job keeps restarting. - Phase filtering (Statements and ComputePools): phase&#x3D;PENDING or phase in (PENDING, RUNNING) or phase notin (PENDING, RUNNING) - Type filtering (Events only): type&#x3D;CMF_STATUS or type in (CMF_STATUS, JOB_STATUS) or type notin (CMF_STATUS, JOB_STATUS) - Severity filtering (Events only): severity&#x3D;WARNING or severity in (INFO, WARNING) or severity notin (WARNING). Values: INFO, WARNING. Severity is independent of type, so the two can be combined to narrow further (e.g. type&#x3D;JOB_STATUS,severity&#x3D;WARNING). An unrecognized severity value is rejected with HTTP 400. - Cluster status filtering (Environments only): clusterStatus&#x3D;CONNECTED or clusterStatus in (CONNECTED, DISCONNECTED) or clusterStatus notin (CONNECTED, DISCONNECTED). Values match the effective state of the backing Kubernetes cluster (CONNECTED, DISCONNECTED, DECOMMISSIONED). - Source filtering (Savepoints only): source&#x3D;MANUAL or source in (MANUAL, SCHEDULE) or source notin (UPGRADE). Identifies how the savepoint was created. Values: MANUAL (triggered via API, CLI, or UI), UPGRADE (created during an application upgrade), SCHEDULE (created by a periodic savepoint schedule). An unrecognized source value is rejected with HTTP 400 (unlike state, which simply returns no matches for an unknown value). Example: ?filter&#x3D;name&#x3D;foo*bar,labels.environment in (production, qa),!labels.development Example (with state): ?filter&#x3D;name&#x3D;prod*,state in (RUNNING, FAILED) Example (with status.error): ?filter&#x3D;status.error,labels.team&#x3D;payments Example (with phase): ?filter&#x3D;name&#x3D;my-stmt*,phase in (PENDING, RUNNING) Example (with type): ?filter&#x3D;type&#x3D;CMF_STATUS or ?filter&#x3D;type in (CMF_STATUS, JOB_STATUS) Example (with severity): ?filter&#x3D;severity&#x3D;WARNING or ?filter&#x3D;type&#x3D;JOB_STATUS,severity&#x3D;WARNING Example (with clusterStatus): ?filter&#x3D;clusterStatus in (CONNECTED, DISCONNECTED) Example (with source): ?filter&#x3D;state&#x3D;COMPLETED,source&#x3D;SCHEDULE | 
 **fields** | **string** | Comma-separated list of field paths to include in the response. Supports nested fields using dot notation. Always includes apiVersion and kind fields even if not explicitly requested. Example: ?fields&#x3D;metadata.name,metadata.createdTimestamp,status.phase | 

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


## GetPodLogs

> KubernetesPodLogs GetPodLogs(ctx, envName, appName).Component(component).PodName(podName).Container(container).TailLines(tailLines).SinceSeconds(sinceSeconds).SinceTime(sinceTime).Timestamps(timestamps).FromStart(fromStart).Previous(previous).Execute()

Get a snapshot of the container logs of one of the Application's pods



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    envName := "envName_example" // string | Name of the Environment
    appName := "appName_example" // string | Name of the Application
    component := "component_example" // string | Which component of the Flink cluster to read the logs of.
    podName := "podName_example" // string | Name of the pod to read the logs of. Must be one of the pods reported in `availablePods`. When omitted, the logs of the first available pod of the requested component are returned. (optional)
    container := "container_example" // string | Name of the container within the pod to read the logs of. Must be one of the containers of that pod. When omitted, the pod's main Flink container is used. (optional)
    tailLines := int32(56) // int32 | Number of most recent log lines to return. Ignored when `fromStart` is set. (optional) (default to 200)
    sinceSeconds := int32(56) // int32 | Only return log lines written within the last N seconds. The `tailLines` cap still applies, so a wide window can still be truncated. (optional)
    sinceTime := time.Now() // time.Time | Only return log lines written at or after this timestamp (RFC 3339). Use it together with `timestamps` to continue reading where a previous response ended. Mutually exclusive with `sinceSeconds`. (optional)
    timestamps := true // bool | Prefix every returned line with the RFC 3339 timestamp at which the container wrote it. (optional) (default to false)
    fromStart := true // bool | Return the log from the oldest line the Kubernetes cluster still retains instead of the most recent ones. The response is capped in size; `truncated` reports whether it hit that cap. (optional) (default to false)
    previous := true // bool | Read the logs of the previous, terminated container instead of the running one. Use this to inspect why a container crashed. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlinkApplicationsApi.GetPodLogs(context.Background(), envName, appName).Component(component).PodName(podName).Container(container).TailLines(tailLines).SinceSeconds(sinceSeconds).SinceTime(sinceTime).Timestamps(timestamps).FromStart(fromStart).Previous(previous).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlinkApplicationsApi.GetPodLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPodLogs`: KubernetesPodLogs
    fmt.Fprintf(os.Stdout, "Response from `FlinkApplicationsApi.GetPodLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment | 
**appName** | **string** | Name of the Application | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPodLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **component** | **string** | Which component of the Flink cluster to read the logs of. | 
 **podName** | **string** | Name of the pod to read the logs of. Must be one of the pods reported in &#x60;availablePods&#x60;. When omitted, the logs of the first available pod of the requested component are returned. | 
 **container** | **string** | Name of the container within the pod to read the logs of. Must be one of the containers of that pod. When omitted, the pod&#39;s main Flink container is used. | 
 **tailLines** | **int32** | Number of most recent log lines to return. Ignored when &#x60;fromStart&#x60; is set. | [default to 200]
 **sinceSeconds** | **int32** | Only return log lines written within the last N seconds. The &#x60;tailLines&#x60; cap still applies, so a wide window can still be truncated. | 
 **sinceTime** | **time.Time** | Only return log lines written at or after this timestamp (RFC 3339). Use it together with &#x60;timestamps&#x60; to continue reading where a previous response ended. Mutually exclusive with &#x60;sinceSeconds&#x60;. | 
 **timestamps** | **bool** | Prefix every returned line with the RFC 3339 timestamp at which the container wrote it. | [default to false]
 **fromStart** | **bool** | Return the log from the oldest line the Kubernetes cluster still retains instead of the most recent ones. The response is capped in size; &#x60;truncated&#x60; reports whether it hit that cap. | [default to false]
 **previous** | **bool** | Read the logs of the previous, terminated container instead of the running one. Use this to inspect why a container crashed. | [default to false]

### Return type

[**KubernetesPodLogs**](KubernetesPodLogs.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicationPods

> FlinkClusterPodsPage ListApplicationPods(ctx, envName, appName).Page(page).Size(size).Sort(sort).Filter(filter).Fields(fields).Execute()

List the Kubernetes pods the Application currently runs on



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
    filter := "filter_example" // string | Filter query string with comma-separated expressions. Supports: - Name filtering: name=foo*bar (wildcards allowed) - Label equality: labels.key = value or labels.key != value - Label set-based: labels.key in (value1, value2) or labels.key notin (value1, value2) - Label existence: labels.key (exists) or !labels.key (does not exist) - State filtering (Applications only): state=RUNNING or state in (RUNNING, FAILED) or state notin (RUNNING, FAILED) - Status error filtering (Applications only): status.error (the Application's status reports an error, i.e. status.error is set and not blank) or !status.error (it reports none). An Application can report an error while its state is not FAILED, for example while its job keeps restarting. - Phase filtering (Statements and ComputePools): phase=PENDING or phase in (PENDING, RUNNING) or phase notin (PENDING, RUNNING) - Type filtering (Events only): type=CMF_STATUS or type in (CMF_STATUS, JOB_STATUS) or type notin (CMF_STATUS, JOB_STATUS) - Severity filtering (Events only): severity=WARNING or severity in (INFO, WARNING) or severity notin (WARNING). Values: INFO, WARNING. Severity is independent of type, so the two can be combined to narrow further (e.g. type=JOB_STATUS,severity=WARNING). An unrecognized severity value is rejected with HTTP 400. - Cluster status filtering (Environments only): clusterStatus=CONNECTED or clusterStatus in (CONNECTED, DISCONNECTED) or clusterStatus notin (CONNECTED, DISCONNECTED). Values match the effective state of the backing Kubernetes cluster (CONNECTED, DISCONNECTED, DECOMMISSIONED). - Source filtering (Savepoints only): source=MANUAL or source in (MANUAL, SCHEDULE) or source notin (UPGRADE). Identifies how the savepoint was created. Values: MANUAL (triggered via API, CLI, or UI), UPGRADE (created during an application upgrade), SCHEDULE (created by a periodic savepoint schedule). An unrecognized source value is rejected with HTTP 400 (unlike state, which simply returns no matches for an unknown value). Example: ?filter=name=foo*bar,labels.environment in (production, qa),!labels.development Example (with state): ?filter=name=prod*,state in (RUNNING, FAILED) Example (with status.error): ?filter=status.error,labels.team=payments Example (with phase): ?filter=name=my-stmt*,phase in (PENDING, RUNNING) Example (with type): ?filter=type=CMF_STATUS or ?filter=type in (CMF_STATUS, JOB_STATUS) Example (with severity): ?filter=severity=WARNING or ?filter=type=JOB_STATUS,severity=WARNING Example (with clusterStatus): ?filter=clusterStatus in (CONNECTED, DISCONNECTED) Example (with source): ?filter=state=COMPLETED,source=SCHEDULE (optional)
    fields := "fields_example" // string | Comma-separated list of field paths to include in the response. Supports nested fields using dot notation. Always includes apiVersion and kind fields even if not explicitly requested. Example: ?fields=metadata.name,metadata.createdTimestamp,status.phase (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlinkApplicationsApi.ListApplicationPods(context.Background(), envName, appName).Page(page).Size(size).Sort(sort).Filter(filter).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlinkApplicationsApi.ListApplicationPods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApplicationPods`: FlinkClusterPodsPage
    fmt.Fprintf(os.Stdout, "Response from `FlinkApplicationsApi.ListApplicationPods`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment | 
**appName** | **string** | Name of the Application | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationPodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Zero-based page index (0..N) | 
 **size** | **int32** | The size of the page to be returned | 
 **sort** | **[]string** | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 
 **filter** | **string** | Filter query string with comma-separated expressions. Supports: - Name filtering: name&#x3D;foo*bar (wildcards allowed) - Label equality: labels.key &#x3D; value or labels.key !&#x3D; value - Label set-based: labels.key in (value1, value2) or labels.key notin (value1, value2) - Label existence: labels.key (exists) or !labels.key (does not exist) - State filtering (Applications only): state&#x3D;RUNNING or state in (RUNNING, FAILED) or state notin (RUNNING, FAILED) - Status error filtering (Applications only): status.error (the Application&#39;s status reports an error, i.e. status.error is set and not blank) or !status.error (it reports none). An Application can report an error while its state is not FAILED, for example while its job keeps restarting. - Phase filtering (Statements and ComputePools): phase&#x3D;PENDING or phase in (PENDING, RUNNING) or phase notin (PENDING, RUNNING) - Type filtering (Events only): type&#x3D;CMF_STATUS or type in (CMF_STATUS, JOB_STATUS) or type notin (CMF_STATUS, JOB_STATUS) - Severity filtering (Events only): severity&#x3D;WARNING or severity in (INFO, WARNING) or severity notin (WARNING). Values: INFO, WARNING. Severity is independent of type, so the two can be combined to narrow further (e.g. type&#x3D;JOB_STATUS,severity&#x3D;WARNING). An unrecognized severity value is rejected with HTTP 400. - Cluster status filtering (Environments only): clusterStatus&#x3D;CONNECTED or clusterStatus in (CONNECTED, DISCONNECTED) or clusterStatus notin (CONNECTED, DISCONNECTED). Values match the effective state of the backing Kubernetes cluster (CONNECTED, DISCONNECTED, DECOMMISSIONED). - Source filtering (Savepoints only): source&#x3D;MANUAL or source in (MANUAL, SCHEDULE) or source notin (UPGRADE). Identifies how the savepoint was created. Values: MANUAL (triggered via API, CLI, or UI), UPGRADE (created during an application upgrade), SCHEDULE (created by a periodic savepoint schedule). An unrecognized source value is rejected with HTTP 400 (unlike state, which simply returns no matches for an unknown value). Example: ?filter&#x3D;name&#x3D;foo*bar,labels.environment in (production, qa),!labels.development Example (with state): ?filter&#x3D;name&#x3D;prod*,state in (RUNNING, FAILED) Example (with status.error): ?filter&#x3D;status.error,labels.team&#x3D;payments Example (with phase): ?filter&#x3D;name&#x3D;my-stmt*,phase in (PENDING, RUNNING) Example (with type): ?filter&#x3D;type&#x3D;CMF_STATUS or ?filter&#x3D;type in (CMF_STATUS, JOB_STATUS) Example (with severity): ?filter&#x3D;severity&#x3D;WARNING or ?filter&#x3D;type&#x3D;JOB_STATUS,severity&#x3D;WARNING Example (with clusterStatus): ?filter&#x3D;clusterStatus in (CONNECTED, DISCONNECTED) Example (with source): ?filter&#x3D;state&#x3D;COMPLETED,source&#x3D;SCHEDULE | 
 **fields** | **string** | Comma-separated list of field paths to include in the response. Supports nested fields using dot notation. Always includes apiVersion and kind fields even if not explicitly requested. Example: ?fields&#x3D;metadata.name,metadata.createdTimestamp,status.phase | 

### Return type

[**FlinkClusterPodsPage**](FlinkClusterPodsPage.md)

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
    startFromSavepointUid := "startFromSavepointUid_example" // string | UID of the Savepoint from which the application should be started. This savepoint could belong to the application or can be a detached savepoint. (optional)

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


 **startFromSavepointUid** | **string** | UID of the Savepoint from which the application should be started. This savepoint could belong to the application or can be a detached savepoint. | 

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

