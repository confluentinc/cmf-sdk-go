# \SecretsApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSecret**](SecretsApi.md#CreateSecret) | **Post** /cmf/api/v1/secrets | Create a Secret.
[**DeleteSecret**](SecretsApi.md#DeleteSecret) | **Delete** /cmf/api/v1/secrets/{secretName} | Delete the secret with the given name.
[**GetSecret**](SecretsApi.md#GetSecret) | **Get** /cmf/api/v1/secrets/{secretName} | Retrieve the Secret of the given name. Note that the secret data is not returned for security reasons.
[**GetSecrets**](SecretsApi.md#GetSecrets) | **Get** /cmf/api/v1/secrets | Retrieve a paginated list of all secrets. Note that the actual secret data is masked for security reasons.
[**UpdateSecret**](SecretsApi.md#UpdateSecret) | **Put** /cmf/api/v1/secrets/{secretName} | Update the secret.



## CreateSecret

> Secret CreateSecret(ctx).Secret(secret).Execute()

Create a Secret.



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
    secret := *openapiclient.NewSecret("ApiVersion_example", "Kind_example", *openapiclient.NewSecretMetadata("Name_example"), *openapiclient.NewSecretSpec()) // Secret | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SecretsApi.CreateSecret(context.Background()).Secret(secret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.CreateSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSecret`: Secret
    fmt.Fprintf(os.Stdout, "Response from `SecretsApi.CreateSecret`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **secret** | [**Secret**](Secret.md) |  | 

### Return type

[**Secret**](Secret.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSecret

> DeleteSecret(ctx, secretName).Execute()

Delete the secret with the given name.

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
    secretName := "secretName_example" // string | Name of the Secret

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SecretsApi.DeleteSecret(context.Background(), secretName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.DeleteSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secretName** | **string** | Name of the Secret | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSecretRequest struct via the builder pattern


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


## GetSecret

> Secret GetSecret(ctx, secretName).Execute()

Retrieve the Secret of the given name. Note that the secret data is not returned for security reasons.

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
    secretName := "secretName_example" // string | Name of the Secret

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SecretsApi.GetSecret(context.Background(), secretName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecret`: Secret
    fmt.Fprintf(os.Stdout, "Response from `SecretsApi.GetSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secretName** | **string** | Name of the Secret | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Secret**](Secret.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecrets

> SecretsPage GetSecrets(ctx).Page(page).Size(size).Sort(sort).Filter(filter).Search(search).SearchScope(searchScope).Fields(fields).Execute()

Retrieve a paginated list of all secrets. Note that the actual secret data is masked for security reasons.

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
    filter := "filter_example" // string | Filter query string with comma-separated expressions. Supports: - Name filtering: name=foo*bar (wildcards allowed) - Label equality: labels.key = value or labels.key != value - Label set-based: labels.key in (value1, value2) or labels.key notin (value1, value2) - Label existence: labels.key (exists) or !labels.key (does not exist) - State filtering (Applications only): state=RUNNING or state in (RUNNING, FAILED) or state notin (RUNNING, FAILED) - Phase filtering (Statements and ComputePools): phase=PENDING or phase in (PENDING, RUNNING) or phase notin (PENDING, RUNNING) - Type filtering (Events only): type=CMF_STATUS or type in (CMF_STATUS, JOB_STATUS) or type notin (CMF_STATUS, JOB_STATUS) Example: ?filter=name=foo*bar,labels.environment in (production, qa),!labels.development Example (with state): ?filter=name=prod*,state in (RUNNING, FAILED) Example (with phase): ?filter=name=my-stmt*,phase in (PENDING, RUNNING) Example (with type): ?filter=type=CMF_STATUS or ?filter=type in (CMF_STATUS, JOB_STATUS) (optional)
    search := "search_example" // string | Search term to match against fields specified in searchScope. Note: Both search and searchScope must be provided together. If only one is provided, the request will be rejected. Example: ?search=foo&searchScope=name,kubernetesNamespace (optional)
    searchScope := "searchScope_example" // string | Comma-separated list of fields to search in. Must be provided together with the search parameter. Unsupported field names will result in a 400 Bad Request. For Environments: supported fields are name, kubernetesNamespace. For Statements: supported fields are name, statement. For Events: supported fields are message, flinkApplicationInstance. For Secrets: supported fields are name, environments. When multiple fields are specified, the search uses OR logic. Example (Environments): ?search=foo&searchScope=name,kubernetesNamespace means (name contains foo OR kubernetesNamespace contains foo) Example (Statements): ?search=SELECT&searchScope=name,statement means (name contains SELECT OR statement contains SELECT) Example (Events): ?search=RUNNING&searchScope=message,flinkApplicationInstance means (message contains RUNNING OR flinkApplicationInstance equals RUNNING) (optional)
    fields := "fields_example" // string | Comma-separated list of field paths to include in the response. Supports nested fields using dot notation. Always includes apiVersion and kind fields even if not explicitly requested. Example: ?fields=metadata.name,metadata.createdTimestamp,status.phase (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SecretsApi.GetSecrets(context.Background()).Page(page).Size(size).Sort(sort).Filter(filter).Search(search).SearchScope(searchScope).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.GetSecrets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecrets`: SecretsPage
    fmt.Fprintf(os.Stdout, "Response from `SecretsApi.GetSecrets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Zero-based page index (0..N) | 
 **size** | **int32** | The size of the page to be returned | 
 **sort** | **[]string** | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 
 **filter** | **string** | Filter query string with comma-separated expressions. Supports: - Name filtering: name&#x3D;foo*bar (wildcards allowed) - Label equality: labels.key &#x3D; value or labels.key !&#x3D; value - Label set-based: labels.key in (value1, value2) or labels.key notin (value1, value2) - Label existence: labels.key (exists) or !labels.key (does not exist) - State filtering (Applications only): state&#x3D;RUNNING or state in (RUNNING, FAILED) or state notin (RUNNING, FAILED) - Phase filtering (Statements and ComputePools): phase&#x3D;PENDING or phase in (PENDING, RUNNING) or phase notin (PENDING, RUNNING) - Type filtering (Events only): type&#x3D;CMF_STATUS or type in (CMF_STATUS, JOB_STATUS) or type notin (CMF_STATUS, JOB_STATUS) Example: ?filter&#x3D;name&#x3D;foo*bar,labels.environment in (production, qa),!labels.development Example (with state): ?filter&#x3D;name&#x3D;prod*,state in (RUNNING, FAILED) Example (with phase): ?filter&#x3D;name&#x3D;my-stmt*,phase in (PENDING, RUNNING) Example (with type): ?filter&#x3D;type&#x3D;CMF_STATUS or ?filter&#x3D;type in (CMF_STATUS, JOB_STATUS) | 
 **search** | **string** | Search term to match against fields specified in searchScope. Note: Both search and searchScope must be provided together. If only one is provided, the request will be rejected. Example: ?search&#x3D;foo&amp;searchScope&#x3D;name,kubernetesNamespace | 
 **searchScope** | **string** | Comma-separated list of fields to search in. Must be provided together with the search parameter. Unsupported field names will result in a 400 Bad Request. For Environments: supported fields are name, kubernetesNamespace. For Statements: supported fields are name, statement. For Events: supported fields are message, flinkApplicationInstance. For Secrets: supported fields are name, environments. When multiple fields are specified, the search uses OR logic. Example (Environments): ?search&#x3D;foo&amp;searchScope&#x3D;name,kubernetesNamespace means (name contains foo OR kubernetesNamespace contains foo) Example (Statements): ?search&#x3D;SELECT&amp;searchScope&#x3D;name,statement means (name contains SELECT OR statement contains SELECT) Example (Events): ?search&#x3D;RUNNING&amp;searchScope&#x3D;message,flinkApplicationInstance means (message contains RUNNING OR flinkApplicationInstance equals RUNNING) | 
 **fields** | **string** | Comma-separated list of field paths to include in the response. Supports nested fields using dot notation. Always includes apiVersion and kind fields even if not explicitly requested. Example: ?fields&#x3D;metadata.name,metadata.createdTimestamp,status.phase | 

### Return type

[**SecretsPage**](SecretsPage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSecret

> Secret UpdateSecret(ctx, secretName).Secret(secret).Execute()

Update the secret.

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
    secretName := "secretName_example" // string | Name of the Secret
    secret := *openapiclient.NewSecret("ApiVersion_example", "Kind_example", *openapiclient.NewSecretMetadata("Name_example"), *openapiclient.NewSecretSpec()) // Secret | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SecretsApi.UpdateSecret(context.Background(), secretName).Secret(secret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.UpdateSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSecret`: Secret
    fmt.Fprintf(os.Stdout, "Response from `SecretsApi.UpdateSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secretName** | **string** | Name of the Secret | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **secret** | [**Secret**](Secret.md) |  | 

### Return type

[**Secret**](Secret.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

