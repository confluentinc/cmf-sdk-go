# \EnvironmentsApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEnvironmentSecretMapping**](EnvironmentsApi.md#CreateEnvironmentSecretMapping) | **Post** /cmf/api/v1/environments/{envName}/secret-mappings | Creates the Environment Secret Mapping for the given Environment.
[**CreateOrUpdateEnvironment**](EnvironmentsApi.md#CreateOrUpdateEnvironment) | **Post** /cmf/api/v1/environments | Create or update an Environment
[**DeleteEnvironment**](EnvironmentsApi.md#DeleteEnvironment) | **Delete** /cmf/api/v1/environments/{envName} | 
[**DeleteEnvironmentSecretMapping**](EnvironmentsApi.md#DeleteEnvironmentSecretMapping) | **Delete** /cmf/api/v1/environments/{envName}/secret-mappings/{name} | Deletes the Environment Secret Mapping for the given Environment and Secret.
[**GetEnvironment**](EnvironmentsApi.md#GetEnvironment) | **Get** /cmf/api/v1/environments/{envName} | Get/Describe an environment with the given name.
[**GetEnvironmentSecretMapping**](EnvironmentsApi.md#GetEnvironmentSecretMapping) | **Get** /cmf/api/v1/environments/{envName}/secret-mappings/{name} | Retrieve the Environment Secret Mapping for the given name in the given environment.
[**GetEnvironmentSecretMappings**](EnvironmentsApi.md#GetEnvironmentSecretMappings) | **Get** /cmf/api/v1/environments/{envName}/secret-mappings | Retrieve a paginated list of all Environment Secret Mappings.
[**GetEnvironments**](EnvironmentsApi.md#GetEnvironments) | **Get** /cmf/api/v1/environments | Retrieve a paginated list of all environments.
[**UpdateEnvironmentSecretMapping**](EnvironmentsApi.md#UpdateEnvironmentSecretMapping) | **Put** /cmf/api/v1/environments/{envName}/secret-mappings/{name} | Updates the Environment Secret Mapping for the given Environment.



## CreateEnvironmentSecretMapping

> EnvironmentSecretMapping CreateEnvironmentSecretMapping(ctx, envName).EnvironmentSecretMapping(environmentSecretMapping).Execute()

Creates the Environment Secret Mapping for the given Environment.

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
    environmentSecretMapping := *openapiclient.NewEnvironmentSecretMapping("ApiVersion_example", "Kind_example") // EnvironmentSecretMapping | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnvironmentsApi.CreateEnvironmentSecretMapping(context.Background(), envName).EnvironmentSecretMapping(environmentSecretMapping).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.CreateEnvironmentSecretMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEnvironmentSecretMapping`: EnvironmentSecretMapping
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.CreateEnvironmentSecretMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEnvironmentSecretMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentSecretMapping** | [**EnvironmentSecretMapping**](EnvironmentSecretMapping.md) |  | 

### Return type

[**EnvironmentSecretMapping**](EnvironmentSecretMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrUpdateEnvironment

> Environment CreateOrUpdateEnvironment(ctx).PostEnvironment(postEnvironment).Execute()

Create or update an Environment

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
    postEnvironment := *openapiclient.NewPostEnvironment("Name_example") // PostEnvironment | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnvironmentsApi.CreateOrUpdateEnvironment(context.Background()).PostEnvironment(postEnvironment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.CreateOrUpdateEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrUpdateEnvironment`: Environment
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.CreateOrUpdateEnvironment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postEnvironment** | [**PostEnvironment**](PostEnvironment.md) |  | 

### Return type

[**Environment**](Environment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEnvironment

> DeleteEnvironment(ctx, envName).Execute()



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
    envName := "envName_example" // string | Name of the Environment to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnvironmentsApi.DeleteEnvironment(context.Background(), envName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.DeleteEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEnvironmentRequest struct via the builder pattern


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


## DeleteEnvironmentSecretMapping

> DeleteEnvironmentSecretMapping(ctx, envName, name).Execute()

Deletes the Environment Secret Mapping for the given Environment and Secret.

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
    envName := "envName_example" // string | Name of the Environment in which the mapping has to be deleted.
    name := "name_example" // string | Name of the environment secret mapping to be deleted in the given environment.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnvironmentsApi.DeleteEnvironmentSecretMapping(context.Background(), envName, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.DeleteEnvironmentSecretMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment in which the mapping has to be deleted. | 
**name** | **string** | Name of the environment secret mapping to be deleted in the given environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEnvironmentSecretMappingRequest struct via the builder pattern


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


## GetEnvironment

> Environment GetEnvironment(ctx, envName).Execute()

Get/Describe an environment with the given name.

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
    envName := "envName_example" // string | Name of the Environment to be retrieved.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnvironmentsApi.GetEnvironment(context.Background(), envName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.GetEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnvironment`: Environment
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.GetEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Environment**](Environment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironmentSecretMapping

> EnvironmentSecretMapping GetEnvironmentSecretMapping(ctx, envName, name).Execute()

Retrieve the Environment Secret Mapping for the given name in the given environment.

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
    name := "name_example" // string | Name of the environment secret mapping to be retrieved.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnvironmentsApi.GetEnvironmentSecretMapping(context.Background(), envName, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.GetEnvironmentSecretMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnvironmentSecretMapping`: EnvironmentSecretMapping
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.GetEnvironmentSecretMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment | 
**name** | **string** | Name of the environment secret mapping to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentSecretMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EnvironmentSecretMapping**](EnvironmentSecretMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironmentSecretMappings

> EnvironmentSecretMappingsPage GetEnvironmentSecretMappings(ctx, envName).Page(page).Size(size).Sort(sort).Execute()

Retrieve a paginated list of all Environment Secret Mappings.

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnvironmentsApi.GetEnvironmentSecretMappings(context.Background(), envName).Page(page).Size(size).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.GetEnvironmentSecretMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnvironmentSecretMappings`: EnvironmentSecretMappingsPage
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.GetEnvironmentSecretMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentSecretMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Zero-based page index (0..N) | 
 **size** | **int32** | The size of the page to be returned | 
 **sort** | **[]string** | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**EnvironmentSecretMappingsPage**](EnvironmentSecretMappingsPage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironments

> EnvironmentsPage GetEnvironments(ctx).Page(page).Size(size).Sort(sort).Execute()

Retrieve a paginated list of all environments.

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnvironmentsApi.GetEnvironments(context.Background()).Page(page).Size(size).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.GetEnvironments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnvironments`: EnvironmentsPage
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.GetEnvironments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Zero-based page index (0..N) | 
 **size** | **int32** | The size of the page to be returned | 
 **sort** | **[]string** | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**EnvironmentsPage**](EnvironmentsPage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEnvironmentSecretMapping

> EnvironmentSecretMapping UpdateEnvironmentSecretMapping(ctx, envName, name).EnvironmentSecretMapping(environmentSecretMapping).Execute()

Updates the Environment Secret Mapping for the given Environment.

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
    name := "name_example" // string | Name of the environment secret mapping to be updated
    environmentSecretMapping := *openapiclient.NewEnvironmentSecretMapping("ApiVersion_example", "Kind_example") // EnvironmentSecretMapping | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnvironmentsApi.UpdateEnvironmentSecretMapping(context.Background(), envName, name).EnvironmentSecretMapping(environmentSecretMapping).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.UpdateEnvironmentSecretMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEnvironmentSecretMapping`: EnvironmentSecretMapping
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.UpdateEnvironmentSecretMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment | 
**name** | **string** | Name of the environment secret mapping to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEnvironmentSecretMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **environmentSecretMapping** | [**EnvironmentSecretMapping**](EnvironmentSecretMapping.md) |  | 

### Return type

[**EnvironmentSecretMapping**](EnvironmentSecretMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

