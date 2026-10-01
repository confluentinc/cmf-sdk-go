# \ServiceAccountAllowlistApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListServiceAccountAllowlistEntries**](ServiceAccountAllowlistApi.md#ListServiceAccountAllowlistEntries) | **Get** /cmf/api/v1/service-account-allowlist-entries | List the service accounts on the grandfathered allowlist.



## ListServiceAccountAllowlistEntries

> ServiceAccountAllowlistEntriesPage ListServiceAccountAllowlistEntries(ctx).Page(page).Size(size).Sort(sort).Filter(filter).Fields(fields).Execute()

List the service accounts on the grandfathered allowlist.



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
    filter := "filter_example" // string | Filter query string with comma-separated expressions. Supports: - Name filtering: name=foo*bar (wildcards allowed) - Label equality: labels.key = value or labels.key != value - Label set-based: labels.key in (value1, value2) or labels.key notin (value1, value2) - Label existence: labels.key (exists) or !labels.key (does not exist) - State filtering (Applications only): state=RUNNING or state in (RUNNING, FAILED) or state notin (RUNNING, FAILED) - Phase filtering (Statements and ComputePools): phase=PENDING or phase in (PENDING, RUNNING) or phase notin (PENDING, RUNNING) - Type filtering (Events only): type=CMF_STATUS or type in (CMF_STATUS, JOB_STATUS) or type notin (CMF_STATUS, JOB_STATUS) - Cluster status filtering (Environments only): clusterStatus=CONNECTED or clusterStatus in (CONNECTED, DISCONNECTED) or clusterStatus notin (CONNECTED, DISCONNECTED). Values match the effective state of the backing Kubernetes cluster (CONNECTED, DISCONNECTED, DECOMMISSIONED). - Source filtering (Savepoints only): source=MANUAL or source in (MANUAL, SCHEDULE) or source notin (UPGRADE). Identifies how the savepoint was created. Values: MANUAL (triggered via API, CLI, or UI), UPGRADE (created during an application upgrade), SCHEDULE (created by a periodic savepoint schedule). An unrecognized source value is rejected with HTTP 400 (unlike state, which simply returns no matches for an unknown value). Example: ?filter=name=foo*bar,labels.environment in (production, qa),!labels.development Example (with state): ?filter=name=prod*,state in (RUNNING, FAILED) Example (with phase): ?filter=name=my-stmt*,phase in (PENDING, RUNNING) Example (with type): ?filter=type=CMF_STATUS or ?filter=type in (CMF_STATUS, JOB_STATUS) Example (with clusterStatus): ?filter=clusterStatus in (CONNECTED, DISCONNECTED) Example (with source): ?filter=state=COMPLETED,source=SCHEDULE (optional)
    fields := "fields_example" // string | Comma-separated list of field paths to include in the response. Supports nested fields using dot notation. Always includes apiVersion and kind fields even if not explicitly requested. Example: ?fields=metadata.name,metadata.createdTimestamp,status.phase (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceAccountAllowlistApi.ListServiceAccountAllowlistEntries(context.Background()).Page(page).Size(size).Sort(sort).Filter(filter).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountAllowlistApi.ListServiceAccountAllowlistEntries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServiceAccountAllowlistEntries`: ServiceAccountAllowlistEntriesPage
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountAllowlistApi.ListServiceAccountAllowlistEntries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListServiceAccountAllowlistEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Zero-based page index (0..N) | 
 **size** | **int32** | The size of the page to be returned | 
 **sort** | **[]string** | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 
 **filter** | **string** | Filter query string with comma-separated expressions. Supports: - Name filtering: name&#x3D;foo*bar (wildcards allowed) - Label equality: labels.key &#x3D; value or labels.key !&#x3D; value - Label set-based: labels.key in (value1, value2) or labels.key notin (value1, value2) - Label existence: labels.key (exists) or !labels.key (does not exist) - State filtering (Applications only): state&#x3D;RUNNING or state in (RUNNING, FAILED) or state notin (RUNNING, FAILED) - Phase filtering (Statements and ComputePools): phase&#x3D;PENDING or phase in (PENDING, RUNNING) or phase notin (PENDING, RUNNING) - Type filtering (Events only): type&#x3D;CMF_STATUS or type in (CMF_STATUS, JOB_STATUS) or type notin (CMF_STATUS, JOB_STATUS) - Cluster status filtering (Environments only): clusterStatus&#x3D;CONNECTED or clusterStatus in (CONNECTED, DISCONNECTED) or clusterStatus notin (CONNECTED, DISCONNECTED). Values match the effective state of the backing Kubernetes cluster (CONNECTED, DISCONNECTED, DECOMMISSIONED). - Source filtering (Savepoints only): source&#x3D;MANUAL or source in (MANUAL, SCHEDULE) or source notin (UPGRADE). Identifies how the savepoint was created. Values: MANUAL (triggered via API, CLI, or UI), UPGRADE (created during an application upgrade), SCHEDULE (created by a periodic savepoint schedule). An unrecognized source value is rejected with HTTP 400 (unlike state, which simply returns no matches for an unknown value). Example: ?filter&#x3D;name&#x3D;foo*bar,labels.environment in (production, qa),!labels.development Example (with state): ?filter&#x3D;name&#x3D;prod*,state in (RUNNING, FAILED) Example (with phase): ?filter&#x3D;name&#x3D;my-stmt*,phase in (PENDING, RUNNING) Example (with type): ?filter&#x3D;type&#x3D;CMF_STATUS or ?filter&#x3D;type in (CMF_STATUS, JOB_STATUS) Example (with clusterStatus): ?filter&#x3D;clusterStatus in (CONNECTED, DISCONNECTED) Example (with source): ?filter&#x3D;state&#x3D;COMPLETED,source&#x3D;SCHEDULE | 
 **fields** | **string** | Comma-separated list of field paths to include in the response. Supports nested fields using dot notation. Always includes apiVersion and kind fields even if not explicitly requested. Example: ?fields&#x3D;metadata.name,metadata.createdTimestamp,status.phase | 

### Return type

[**ServiceAccountAllowlistEntriesPage**](ServiceAccountAllowlistEntriesPage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

