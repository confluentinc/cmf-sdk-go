# \SqlAssistantApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSqlAssistantConversation**](SqlAssistantApi.md#CreateSqlAssistantConversation) | **Post** /cmf/api/v1alpha1/environments/{envName}/sql-assistant/conversations | Create a SQL assistant conversation; the server assigns its name
[**CreateSqlAssistantTurn**](SqlAssistantApi.md#CreateSqlAssistantTurn) | **Post** /cmf/api/v1alpha1/environments/{envName}/sql-assistant/conversations/{conversationName}/turns | Append a turn; when a model call is due the server calls the model and returns its response
[**DeleteSqlAssistantConversation**](SqlAssistantApi.md#DeleteSqlAssistantConversation) | **Delete** /cmf/api/v1alpha1/environments/{envName}/sql-assistant/conversations/{conversationName} | Delete a conversation and all its turns
[**GetSqlAssistantConversation**](SqlAssistantApi.md#GetSqlAssistantConversation) | **Get** /cmf/api/v1alpha1/environments/{envName}/sql-assistant/conversations/{conversationName} | Get a SQL assistant conversation
[**ListSqlAssistantConversations**](SqlAssistantApi.md#ListSqlAssistantConversations) | **Get** /cmf/api/v1alpha1/environments/{envName}/sql-assistant/conversations | List the SQL assistant conversations in an environment
[**ListSqlAssistantTurns**](SqlAssistantApi.md#ListSqlAssistantTurns) | **Get** /cmf/api/v1alpha1/environments/{envName}/sql-assistant/conversations/{conversationName}/turns | List the turns of a conversation, in order
[**UpdateSqlAssistantConversation**](SqlAssistantApi.md#UpdateSqlAssistantConversation) | **Put** /cmf/api/v1alpha1/environments/{envName}/sql-assistant/conversations/{conversationName} | Update a conversation&#39;s compute pool, title, labels and annotations



## CreateSqlAssistantConversation

> SqlAssistantConversation CreateSqlAssistantConversation(ctx, envName).SqlAssistantConversation(sqlAssistantConversation).Execute()

Create a SQL assistant conversation; the server assigns its name

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
    sqlAssistantConversation := *openapiclient.NewSqlAssistantConversation("ApiVersion_example", "Kind_example", *openapiclient.NewSqlAssistantConversationMetadata(), *openapiclient.NewSqlAssistantConversationSpec("ComputePool_example")) // SqlAssistantConversation | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SqlAssistantApi.CreateSqlAssistantConversation(context.Background(), envName).SqlAssistantConversation(sqlAssistantConversation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SqlAssistantApi.CreateSqlAssistantConversation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSqlAssistantConversation`: SqlAssistantConversation
    fmt.Fprintf(os.Stdout, "Response from `SqlAssistantApi.CreateSqlAssistantConversation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSqlAssistantConversationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sqlAssistantConversation** | [**SqlAssistantConversation**](SqlAssistantConversation.md) |  | 

### Return type

[**SqlAssistantConversation**](SqlAssistantConversation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSqlAssistantTurn

> SqlAssistantTurn CreateSqlAssistantTurn(ctx, envName, conversationName).SqlAssistantTurn(sqlAssistantTurn).Execute()

Append a turn; when a model call is due the server calls the model and returns its response

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
    conversationName := "conversationName_example" // string | Server-assigned name of the conversation
    sqlAssistantTurn := *openapiclient.NewSqlAssistantTurn("ApiVersion_example", "Kind_example", *openapiclient.NewSqlAssistantTurnMetadata(), *openapiclient.NewSqlAssistantTurnSpec("Role_example")) // SqlAssistantTurn | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SqlAssistantApi.CreateSqlAssistantTurn(context.Background(), envName, conversationName).SqlAssistantTurn(sqlAssistantTurn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SqlAssistantApi.CreateSqlAssistantTurn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSqlAssistantTurn`: SqlAssistantTurn
    fmt.Fprintf(os.Stdout, "Response from `SqlAssistantApi.CreateSqlAssistantTurn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment | 
**conversationName** | **string** | Server-assigned name of the conversation | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSqlAssistantTurnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sqlAssistantTurn** | [**SqlAssistantTurn**](SqlAssistantTurn.md) |  | 

### Return type

[**SqlAssistantTurn**](SqlAssistantTurn.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSqlAssistantConversation

> DeleteSqlAssistantConversation(ctx, envName, conversationName).Execute()

Delete a conversation and all its turns

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
    conversationName := "conversationName_example" // string | Server-assigned name of the conversation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SqlAssistantApi.DeleteSqlAssistantConversation(context.Background(), envName, conversationName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SqlAssistantApi.DeleteSqlAssistantConversation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment | 
**conversationName** | **string** | Server-assigned name of the conversation | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSqlAssistantConversationRequest struct via the builder pattern


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


## GetSqlAssistantConversation

> SqlAssistantConversation GetSqlAssistantConversation(ctx, envName, conversationName).Execute()

Get a SQL assistant conversation

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
    conversationName := "conversationName_example" // string | Server-assigned name of the conversation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SqlAssistantApi.GetSqlAssistantConversation(context.Background(), envName, conversationName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SqlAssistantApi.GetSqlAssistantConversation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSqlAssistantConversation`: SqlAssistantConversation
    fmt.Fprintf(os.Stdout, "Response from `SqlAssistantApi.GetSqlAssistantConversation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment | 
**conversationName** | **string** | Server-assigned name of the conversation | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSqlAssistantConversationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SqlAssistantConversation**](SqlAssistantConversation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSqlAssistantConversations

> SqlAssistantConversationsPage ListSqlAssistantConversations(ctx, envName).Page(page).Size(size).Sort(sort).Filter(filter).Fields(fields).Execute()

List the SQL assistant conversations in an environment

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
    filter := "filter_example" // string | Filter query string with comma-separated expressions. Supports: - Name filtering: name=foo*bar (wildcards allowed) - Label equality: labels.key = value or labels.key != value - Label set-based: labels.key in (value1, value2) or labels.key notin (value1, value2) - Label existence: labels.key (exists) or !labels.key (does not exist) - State filtering (Applications only): state=RUNNING or state in (RUNNING, FAILED) or state notin (RUNNING, FAILED) - Status error filtering (Applications only): status.error (the Application's status reports an error, i.e. status.error is set and not blank) or !status.error (it reports none). An Application can report an error while its state is not FAILED, for example while its job keeps restarting. - Phase filtering (Statements and ComputePools): phase=PENDING or phase in (PENDING, RUNNING) or phase notin (PENDING, RUNNING) - Type filtering (Events only): type=CMF_STATUS or type in (CMF_STATUS, JOB_STATUS) or type notin (CMF_STATUS, JOB_STATUS) - Severity filtering (Events only): severity=WARNING or severity in (INFO, WARNING) or severity notin (WARNING). Values: INFO, WARNING. Severity is independent of type, so the two can be combined to narrow further (e.g. type=JOB_STATUS,severity=WARNING). An unrecognized severity value is rejected with HTTP 400. - Cluster status filtering (Environments only): clusterStatus=CONNECTED or clusterStatus in (CONNECTED, DISCONNECTED) or clusterStatus notin (CONNECTED, DISCONNECTED). Values match the effective state of the backing Kubernetes cluster (CONNECTED, DISCONNECTED, DECOMMISSIONED). - Source filtering (Savepoints only): source=MANUAL or source in (MANUAL, SCHEDULE) or source notin (UPGRADE). Identifies how the savepoint was created. Values: MANUAL (triggered via API, CLI, or UI), UPGRADE (created during an application upgrade), SCHEDULE (created by a periodic savepoint schedule). An unrecognized source value is rejected with HTTP 400 (unlike state, which simply returns no matches for an unknown value). Example: ?filter=name=foo*bar,labels.environment in (production, qa),!labels.development Example (with state): ?filter=name=prod*,state in (RUNNING, FAILED) Example (with status.error): ?filter=status.error,labels.team=payments Example (with phase): ?filter=name=my-stmt*,phase in (PENDING, RUNNING) Example (with type): ?filter=type=CMF_STATUS or ?filter=type in (CMF_STATUS, JOB_STATUS) Example (with severity): ?filter=severity=WARNING or ?filter=type=JOB_STATUS,severity=WARNING Example (with clusterStatus): ?filter=clusterStatus in (CONNECTED, DISCONNECTED) Example (with source): ?filter=state=COMPLETED,source=SCHEDULE (optional)
    fields := "fields_example" // string | Comma-separated list of field paths to include in the response. Supports nested fields using dot notation. Always includes apiVersion and kind fields even if not explicitly requested. Example: ?fields=metadata.name,metadata.createdTimestamp,status.phase (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SqlAssistantApi.ListSqlAssistantConversations(context.Background(), envName).Page(page).Size(size).Sort(sort).Filter(filter).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SqlAssistantApi.ListSqlAssistantConversations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSqlAssistantConversations`: SqlAssistantConversationsPage
    fmt.Fprintf(os.Stdout, "Response from `SqlAssistantApi.ListSqlAssistantConversations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSqlAssistantConversationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Zero-based page index (0..N) | 
 **size** | **int32** | The size of the page to be returned | 
 **sort** | **[]string** | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 
 **filter** | **string** | Filter query string with comma-separated expressions. Supports: - Name filtering: name&#x3D;foo*bar (wildcards allowed) - Label equality: labels.key &#x3D; value or labels.key !&#x3D; value - Label set-based: labels.key in (value1, value2) or labels.key notin (value1, value2) - Label existence: labels.key (exists) or !labels.key (does not exist) - State filtering (Applications only): state&#x3D;RUNNING or state in (RUNNING, FAILED) or state notin (RUNNING, FAILED) - Status error filtering (Applications only): status.error (the Application&#39;s status reports an error, i.e. status.error is set and not blank) or !status.error (it reports none). An Application can report an error while its state is not FAILED, for example while its job keeps restarting. - Phase filtering (Statements and ComputePools): phase&#x3D;PENDING or phase in (PENDING, RUNNING) or phase notin (PENDING, RUNNING) - Type filtering (Events only): type&#x3D;CMF_STATUS or type in (CMF_STATUS, JOB_STATUS) or type notin (CMF_STATUS, JOB_STATUS) - Severity filtering (Events only): severity&#x3D;WARNING or severity in (INFO, WARNING) or severity notin (WARNING). Values: INFO, WARNING. Severity is independent of type, so the two can be combined to narrow further (e.g. type&#x3D;JOB_STATUS,severity&#x3D;WARNING). An unrecognized severity value is rejected with HTTP 400. - Cluster status filtering (Environments only): clusterStatus&#x3D;CONNECTED or clusterStatus in (CONNECTED, DISCONNECTED) or clusterStatus notin (CONNECTED, DISCONNECTED). Values match the effective state of the backing Kubernetes cluster (CONNECTED, DISCONNECTED, DECOMMISSIONED). - Source filtering (Savepoints only): source&#x3D;MANUAL or source in (MANUAL, SCHEDULE) or source notin (UPGRADE). Identifies how the savepoint was created. Values: MANUAL (triggered via API, CLI, or UI), UPGRADE (created during an application upgrade), SCHEDULE (created by a periodic savepoint schedule). An unrecognized source value is rejected with HTTP 400 (unlike state, which simply returns no matches for an unknown value). Example: ?filter&#x3D;name&#x3D;foo*bar,labels.environment in (production, qa),!labels.development Example (with state): ?filter&#x3D;name&#x3D;prod*,state in (RUNNING, FAILED) Example (with status.error): ?filter&#x3D;status.error,labels.team&#x3D;payments Example (with phase): ?filter&#x3D;name&#x3D;my-stmt*,phase in (PENDING, RUNNING) Example (with type): ?filter&#x3D;type&#x3D;CMF_STATUS or ?filter&#x3D;type in (CMF_STATUS, JOB_STATUS) Example (with severity): ?filter&#x3D;severity&#x3D;WARNING or ?filter&#x3D;type&#x3D;JOB_STATUS,severity&#x3D;WARNING Example (with clusterStatus): ?filter&#x3D;clusterStatus in (CONNECTED, DISCONNECTED) Example (with source): ?filter&#x3D;state&#x3D;COMPLETED,source&#x3D;SCHEDULE | 
 **fields** | **string** | Comma-separated list of field paths to include in the response. Supports nested fields using dot notation. Always includes apiVersion and kind fields even if not explicitly requested. Example: ?fields&#x3D;metadata.name,metadata.createdTimestamp,status.phase | 

### Return type

[**SqlAssistantConversationsPage**](SqlAssistantConversationsPage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSqlAssistantTurns

> SqlAssistantTurnsPage ListSqlAssistantTurns(ctx, envName, conversationName).Page(page).Size(size).Execute()

List the turns of a conversation, in order

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
    conversationName := "conversationName_example" // string | Server-assigned name of the conversation
    page := int32(56) // int32 | Zero-based page index (0..N) (optional)
    size := int32(56) // int32 | The size of the page to be returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SqlAssistantApi.ListSqlAssistantTurns(context.Background(), envName, conversationName).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SqlAssistantApi.ListSqlAssistantTurns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSqlAssistantTurns`: SqlAssistantTurnsPage
    fmt.Fprintf(os.Stdout, "Response from `SqlAssistantApi.ListSqlAssistantTurns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment | 
**conversationName** | **string** | Server-assigned name of the conversation | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSqlAssistantTurnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Zero-based page index (0..N) | 
 **size** | **int32** | The size of the page to be returned | 

### Return type

[**SqlAssistantTurnsPage**](SqlAssistantTurnsPage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSqlAssistantConversation

> SqlAssistantConversation UpdateSqlAssistantConversation(ctx, envName, conversationName).SqlAssistantConversation(sqlAssistantConversation).Execute()

Update a conversation's compute pool, title, labels and annotations

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
    conversationName := "conversationName_example" // string | Server-assigned name of the conversation
    sqlAssistantConversation := *openapiclient.NewSqlAssistantConversation("ApiVersion_example", "Kind_example", *openapiclient.NewSqlAssistantConversationMetadata(), *openapiclient.NewSqlAssistantConversationSpec("ComputePool_example")) // SqlAssistantConversation | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SqlAssistantApi.UpdateSqlAssistantConversation(context.Background(), envName, conversationName).SqlAssistantConversation(sqlAssistantConversation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SqlAssistantApi.UpdateSqlAssistantConversation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSqlAssistantConversation`: SqlAssistantConversation
    fmt.Fprintf(os.Stdout, "Response from `SqlAssistantApi.UpdateSqlAssistantConversation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment | 
**conversationName** | **string** | Server-assigned name of the conversation | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSqlAssistantConversationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sqlAssistantConversation** | [**SqlAssistantConversation**](SqlAssistantConversation.md) |  | 

### Return type

[**SqlAssistantConversation**](SqlAssistantConversation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

