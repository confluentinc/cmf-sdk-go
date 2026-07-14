# \ArtifactsApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateArtifact**](ArtifactsApi.md#CreateArtifact) | **Post** /cmf/api/v1/environments/{envName}/artifacts | Upload a new Artifact to an Environment.
[**DeleteArtifact**](ArtifactsApi.md#DeleteArtifact) | **Delete** /cmf/api/v1/environments/{envName}/artifacts/{artifactName} | Deletes the Artifact of the given name in the given Environment.
[**DownloadArtifactContent**](ArtifactsApi.md#DownloadArtifactContent) | **Get** /cmf/api/v1/environments/{envName}/artifacts/{artifactName}/content | Download the binary content of an Artifact.
[**GetArtifact**](ArtifactsApi.md#GetArtifact) | **Get** /cmf/api/v1/environments/{envName}/artifacts/{artifactName} | Retrieve an Artifact of the given name in the given Environment.
[**ListArtifactVersions**](ArtifactsApi.md#ListArtifactVersions) | **Get** /cmf/api/v1/environments/{envName}/artifacts/{artifactName}/versions | Retrieve a paginated list of all versions of an Artifact, ordered newest-first.
[**ListArtifacts**](ArtifactsApi.md#ListArtifacts) | **Get** /cmf/api/v1/environments/{envName}/artifacts | Retrieve a paginated list of all Artifacts in the given Environment.
[**UpdateArtifact**](ArtifactsApi.md#UpdateArtifact) | **Put** /cmf/api/v1/environments/{envName}/artifacts/{artifactName} | Update an Artifact&#39;s labels and annotations, and/or upload a new version.



## CreateArtifact

> Artifact CreateArtifact(ctx, envName).Artifact(artifact).File(file).Execute()

Upload a new Artifact to an Environment.



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
    artifact := *openapiclient.NewArtifact("ApiVersion_example", "Kind_example", *openapiclient.NewArtifactMetadata("Name_example"), map[string]interface{}(123)) // Artifact | 
    file := os.NewFile(1234, "some_file") // *os.File | Binary content of the Artifact.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArtifactsApi.CreateArtifact(context.Background(), envName).Artifact(artifact).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsApi.CreateArtifact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateArtifact`: Artifact
    fmt.Fprintf(os.Stdout, "Response from `ArtifactsApi.CreateArtifact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateArtifactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **artifact** | [**Artifact**](Artifact.md) |  | 
 **file** | ***os.File** | Binary content of the Artifact. | 

### Return type

[**Artifact**](Artifact.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteArtifact

> DeleteArtifact(ctx, envName, artifactName).Version(version).Execute()

Deletes the Artifact of the given name in the given Environment.



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
    artifactName := "artifactName_example" // string | Name of the Artifact
    version := "version_example" // string | Specific version number to delete, or `all`. Defaults to `all`. (optional) (default to "all")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArtifactsApi.DeleteArtifact(context.Background(), envName, artifactName).Version(version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsApi.DeleteArtifact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment | 
**artifactName** | **string** | Name of the Artifact | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteArtifactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **version** | **string** | Specific version number to delete, or &#x60;all&#x60;. Defaults to &#x60;all&#x60;. | [default to &quot;all&quot;]

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


## DownloadArtifactContent

> *os.File DownloadArtifactContent(ctx, envName, artifactName).Version(version).Execute()

Download the binary content of an Artifact.



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
    artifactName := "artifactName_example" // string | Name of the Artifact
    version := "version_example" // string | Specific version number to download, or `latest`. Defaults to `latest`. (optional) (default to "latest")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArtifactsApi.DownloadArtifactContent(context.Background(), envName, artifactName).Version(version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsApi.DownloadArtifactContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadArtifactContent`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ArtifactsApi.DownloadArtifactContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment | 
**artifactName** | **string** | Name of the Artifact | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadArtifactContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **version** | **string** | Specific version number to download, or &#x60;latest&#x60;. Defaults to &#x60;latest&#x60;. | [default to &quot;latest&quot;]

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArtifact

> Artifact GetArtifact(ctx, envName, artifactName).Version(version).Execute()

Retrieve an Artifact of the given name in the given Environment.

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
    artifactName := "artifactName_example" // string | Name of the Artifact
    version := "version_example" // string | Specific version number to retrieve, or `latest`. Defaults to `latest`. (optional) (default to "latest")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArtifactsApi.GetArtifact(context.Background(), envName, artifactName).Version(version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsApi.GetArtifact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetArtifact`: Artifact
    fmt.Fprintf(os.Stdout, "Response from `ArtifactsApi.GetArtifact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment | 
**artifactName** | **string** | Name of the Artifact | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetArtifactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **version** | **string** | Specific version number to retrieve, or &#x60;latest&#x60;. Defaults to &#x60;latest&#x60;. | [default to &quot;latest&quot;]

### Return type

[**Artifact**](Artifact.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListArtifactVersions

> ArtifactsPage ListArtifactVersions(ctx, envName, artifactName).Page(page).Size(size).Sort(sort).Filter(filter).Fields(fields).Execute()

Retrieve a paginated list of all versions of an Artifact, ordered newest-first.



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
    artifactName := "artifactName_example" // string | Name of the Artifact
    page := int32(56) // int32 | Zero-based page index (0..N) (optional)
    size := int32(56) // int32 | The size of the page to be returned (optional)
    sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. (optional)
    filter := "filter_example" // string | Filter query string with comma-separated expressions. Supports: - Name filtering: name=foo*bar (wildcards allowed) - Label equality: labels.key = value or labels.key != value - Label set-based: labels.key in (value1, value2) or labels.key notin (value1, value2) - Label existence: labels.key (exists) or !labels.key (does not exist) - State filtering (Applications only): state=RUNNING or state in (RUNNING, FAILED) or state notin (RUNNING, FAILED) - Phase filtering (Statements and ComputePools): phase=PENDING or phase in (PENDING, RUNNING) or phase notin (PENDING, RUNNING) - Type filtering (Events only): type=CMF_STATUS or type in (CMF_STATUS, JOB_STATUS) or type notin (CMF_STATUS, JOB_STATUS) - Cluster status filtering (Environments only): clusterStatus=CONNECTED or clusterStatus in (CONNECTED, DISCONNECTED) or clusterStatus notin (CONNECTED, DISCONNECTED). Values match the effective state of the backing Kubernetes cluster (CONNECTED, DISCONNECTED, DECOMMISSIONED). - Source filtering (Savepoints only): source=MANUAL or source in (MANUAL, SCHEDULE) or source notin (UPGRADE). Identifies how the savepoint was created. Values: MANUAL (triggered via API, CLI, or UI), UPGRADE (created during an application upgrade), SCHEDULE (created by a periodic savepoint schedule). An unrecognized source value is rejected with HTTP 400 (unlike state, which simply returns no matches for an unknown value). Example: ?filter=name=foo*bar,labels.environment in (production, qa),!labels.development Example (with state): ?filter=name=prod*,state in (RUNNING, FAILED) Example (with phase): ?filter=name=my-stmt*,phase in (PENDING, RUNNING) Example (with type): ?filter=type=CMF_STATUS or ?filter=type in (CMF_STATUS, JOB_STATUS) Example (with clusterStatus): ?filter=clusterStatus in (CONNECTED, DISCONNECTED) Example (with source): ?filter=state=COMPLETED,source=SCHEDULE (optional)
    fields := "fields_example" // string | Comma-separated list of field paths to include in the response. Supports nested fields using dot notation. Always includes apiVersion and kind fields even if not explicitly requested. Example: ?fields=metadata.name,metadata.createdTimestamp,status.phase (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArtifactsApi.ListArtifactVersions(context.Background(), envName, artifactName).Page(page).Size(size).Sort(sort).Filter(filter).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsApi.ListArtifactVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListArtifactVersions`: ArtifactsPage
    fmt.Fprintf(os.Stdout, "Response from `ArtifactsApi.ListArtifactVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment | 
**artifactName** | **string** | Name of the Artifact | 

### Other Parameters

Other parameters are passed through a pointer to a apiListArtifactVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Zero-based page index (0..N) | 
 **size** | **int32** | The size of the page to be returned | 
 **sort** | **[]string** | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 
 **filter** | **string** | Filter query string with comma-separated expressions. Supports: - Name filtering: name&#x3D;foo*bar (wildcards allowed) - Label equality: labels.key &#x3D; value or labels.key !&#x3D; value - Label set-based: labels.key in (value1, value2) or labels.key notin (value1, value2) - Label existence: labels.key (exists) or !labels.key (does not exist) - State filtering (Applications only): state&#x3D;RUNNING or state in (RUNNING, FAILED) or state notin (RUNNING, FAILED) - Phase filtering (Statements and ComputePools): phase&#x3D;PENDING or phase in (PENDING, RUNNING) or phase notin (PENDING, RUNNING) - Type filtering (Events only): type&#x3D;CMF_STATUS or type in (CMF_STATUS, JOB_STATUS) or type notin (CMF_STATUS, JOB_STATUS) - Cluster status filtering (Environments only): clusterStatus&#x3D;CONNECTED or clusterStatus in (CONNECTED, DISCONNECTED) or clusterStatus notin (CONNECTED, DISCONNECTED). Values match the effective state of the backing Kubernetes cluster (CONNECTED, DISCONNECTED, DECOMMISSIONED). - Source filtering (Savepoints only): source&#x3D;MANUAL or source in (MANUAL, SCHEDULE) or source notin (UPGRADE). Identifies how the savepoint was created. Values: MANUAL (triggered via API, CLI, or UI), UPGRADE (created during an application upgrade), SCHEDULE (created by a periodic savepoint schedule). An unrecognized source value is rejected with HTTP 400 (unlike state, which simply returns no matches for an unknown value). Example: ?filter&#x3D;name&#x3D;foo*bar,labels.environment in (production, qa),!labels.development Example (with state): ?filter&#x3D;name&#x3D;prod*,state in (RUNNING, FAILED) Example (with phase): ?filter&#x3D;name&#x3D;my-stmt*,phase in (PENDING, RUNNING) Example (with type): ?filter&#x3D;type&#x3D;CMF_STATUS or ?filter&#x3D;type in (CMF_STATUS, JOB_STATUS) Example (with clusterStatus): ?filter&#x3D;clusterStatus in (CONNECTED, DISCONNECTED) Example (with source): ?filter&#x3D;state&#x3D;COMPLETED,source&#x3D;SCHEDULE | 
 **fields** | **string** | Comma-separated list of field paths to include in the response. Supports nested fields using dot notation. Always includes apiVersion and kind fields even if not explicitly requested. Example: ?fields&#x3D;metadata.name,metadata.createdTimestamp,status.phase | 

### Return type

[**ArtifactsPage**](ArtifactsPage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListArtifacts

> ArtifactsPage ListArtifacts(ctx, envName).Page(page).Size(size).Sort(sort).Filter(filter).Fields(fields).Execute()

Retrieve a paginated list of all Artifacts in the given Environment.



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
    filter := "filter_example" // string | Filter query string with comma-separated expressions. Supports: - Name filtering: name=foo*bar (wildcards allowed) - Label equality: labels.key = value or labels.key != value - Label set-based: labels.key in (value1, value2) or labels.key notin (value1, value2) - Label existence: labels.key (exists) or !labels.key (does not exist) - State filtering (Applications only): state=RUNNING or state in (RUNNING, FAILED) or state notin (RUNNING, FAILED) - Phase filtering (Statements and ComputePools): phase=PENDING or phase in (PENDING, RUNNING) or phase notin (PENDING, RUNNING) - Type filtering (Events only): type=CMF_STATUS or type in (CMF_STATUS, JOB_STATUS) or type notin (CMF_STATUS, JOB_STATUS) - Cluster status filtering (Environments only): clusterStatus=CONNECTED or clusterStatus in (CONNECTED, DISCONNECTED) or clusterStatus notin (CONNECTED, DISCONNECTED). Values match the effective state of the backing Kubernetes cluster (CONNECTED, DISCONNECTED, DECOMMISSIONED). - Source filtering (Savepoints only): source=MANUAL or source in (MANUAL, SCHEDULE) or source notin (UPGRADE). Identifies how the savepoint was created. Values: MANUAL (triggered via API, CLI, or UI), UPGRADE (created during an application upgrade), SCHEDULE (created by a periodic savepoint schedule). An unrecognized source value is rejected with HTTP 400 (unlike state, which simply returns no matches for an unknown value). Example: ?filter=name=foo*bar,labels.environment in (production, qa),!labels.development Example (with state): ?filter=name=prod*,state in (RUNNING, FAILED) Example (with phase): ?filter=name=my-stmt*,phase in (PENDING, RUNNING) Example (with type): ?filter=type=CMF_STATUS or ?filter=type in (CMF_STATUS, JOB_STATUS) Example (with clusterStatus): ?filter=clusterStatus in (CONNECTED, DISCONNECTED) Example (with source): ?filter=state=COMPLETED,source=SCHEDULE (optional)
    fields := "fields_example" // string | Comma-separated list of field paths to include in the response. Supports nested fields using dot notation. Always includes apiVersion and kind fields even if not explicitly requested. Example: ?fields=metadata.name,metadata.createdTimestamp,status.phase (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArtifactsApi.ListArtifacts(context.Background(), envName).Page(page).Size(size).Sort(sort).Filter(filter).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsApi.ListArtifacts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListArtifacts`: ArtifactsPage
    fmt.Fprintf(os.Stdout, "Response from `ArtifactsApi.ListArtifacts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment | 

### Other Parameters

Other parameters are passed through a pointer to a apiListArtifactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Zero-based page index (0..N) | 
 **size** | **int32** | The size of the page to be returned | 
 **sort** | **[]string** | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 
 **filter** | **string** | Filter query string with comma-separated expressions. Supports: - Name filtering: name&#x3D;foo*bar (wildcards allowed) - Label equality: labels.key &#x3D; value or labels.key !&#x3D; value - Label set-based: labels.key in (value1, value2) or labels.key notin (value1, value2) - Label existence: labels.key (exists) or !labels.key (does not exist) - State filtering (Applications only): state&#x3D;RUNNING or state in (RUNNING, FAILED) or state notin (RUNNING, FAILED) - Phase filtering (Statements and ComputePools): phase&#x3D;PENDING or phase in (PENDING, RUNNING) or phase notin (PENDING, RUNNING) - Type filtering (Events only): type&#x3D;CMF_STATUS or type in (CMF_STATUS, JOB_STATUS) or type notin (CMF_STATUS, JOB_STATUS) - Cluster status filtering (Environments only): clusterStatus&#x3D;CONNECTED or clusterStatus in (CONNECTED, DISCONNECTED) or clusterStatus notin (CONNECTED, DISCONNECTED). Values match the effective state of the backing Kubernetes cluster (CONNECTED, DISCONNECTED, DECOMMISSIONED). - Source filtering (Savepoints only): source&#x3D;MANUAL or source in (MANUAL, SCHEDULE) or source notin (UPGRADE). Identifies how the savepoint was created. Values: MANUAL (triggered via API, CLI, or UI), UPGRADE (created during an application upgrade), SCHEDULE (created by a periodic savepoint schedule). An unrecognized source value is rejected with HTTP 400 (unlike state, which simply returns no matches for an unknown value). Example: ?filter&#x3D;name&#x3D;foo*bar,labels.environment in (production, qa),!labels.development Example (with state): ?filter&#x3D;name&#x3D;prod*,state in (RUNNING, FAILED) Example (with phase): ?filter&#x3D;name&#x3D;my-stmt*,phase in (PENDING, RUNNING) Example (with type): ?filter&#x3D;type&#x3D;CMF_STATUS or ?filter&#x3D;type in (CMF_STATUS, JOB_STATUS) Example (with clusterStatus): ?filter&#x3D;clusterStatus in (CONNECTED, DISCONNECTED) Example (with source): ?filter&#x3D;state&#x3D;COMPLETED,source&#x3D;SCHEDULE | 
 **fields** | **string** | Comma-separated list of field paths to include in the response. Supports nested fields using dot notation. Always includes apiVersion and kind fields even if not explicitly requested. Example: ?fields&#x3D;metadata.name,metadata.createdTimestamp,status.phase | 

### Return type

[**ArtifactsPage**](ArtifactsPage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateArtifact

> Artifact UpdateArtifact(ctx, envName, artifactName).Artifact(artifact).File(file).Execute()

Update an Artifact's labels and annotations, and/or upload a new version.



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
    artifactName := "artifactName_example" // string | Name of the Artifact
    artifact := *openapiclient.NewArtifact("ApiVersion_example", "Kind_example", *openapiclient.NewArtifactMetadata("Name_example"), map[string]interface{}(123)) // Artifact | 
    file := os.NewFile(1234, "some_file") // *os.File | Optional new binary content. When present, a new version is created. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArtifactsApi.UpdateArtifact(context.Background(), envName, artifactName).Artifact(artifact).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsApi.UpdateArtifact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateArtifact`: Artifact
    fmt.Fprintf(os.Stdout, "Response from `ArtifactsApi.UpdateArtifact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment | 
**artifactName** | **string** | Name of the Artifact | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateArtifactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **artifact** | [**Artifact**](Artifact.md) |  | 
 **file** | ***os.File** | Optional new binary content. When present, a new version is created. | 

### Return type

[**Artifact**](Artifact.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

