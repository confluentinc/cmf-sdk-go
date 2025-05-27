# \DefaultApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateComputePool**](DefaultApi.md#CreateComputePool) | **Post** /cmf/api/v1/environments/{envName}/compute-pools | Creates a new Flink Compute Pool in the given Environment.
[**CreateEnvironmentSecretMapping**](DefaultApi.md#CreateEnvironmentSecretMapping) | **Post** /cmf/api/v1/environments/{envName}/secret-mappings | Creates the Environment Secret Mapping for the given Environment.
[**CreateKafkaCatalog**](DefaultApi.md#CreateKafkaCatalog) | **Post** /cmf/api/v1/catalogs/kafka | Creates a new Kafka Catalog that can be referenced by Flink Statements
[**CreateOrUpdateApplication**](DefaultApi.md#CreateOrUpdateApplication) | **Post** /cmf/api/v1/environments/{envName}/applications | Creates a new Flink Application or updates an existing one in the given Environment.
[**CreateOrUpdateEnvironment**](DefaultApi.md#CreateOrUpdateEnvironment) | **Post** /cmf/api/v1/environments | Create or update an Environment
[**CreateSecret**](DefaultApi.md#CreateSecret) | **Post** /cmf/api/v1/secrets | Create a Secret.
[**CreateStatement**](DefaultApi.md#CreateStatement) | **Post** /cmf/api/v1/environments/{envName}/statements | Creates a new Flink SQL Statement in the given Environment.
[**DeleteApplication**](DefaultApi.md#DeleteApplication) | **Delete** /cmf/api/v1/environments/{envName}/applications/{appName} | Deletes an Application of the given name in the given Environment.
[**DeleteComputePool**](DefaultApi.md#DeleteComputePool) | **Delete** /cmf/api/v1/environments/{envName}/compute-pools/{computePoolName} | Deletes the ComputePool of the given name in the given Environment.
[**DeleteEnvironment**](DefaultApi.md#DeleteEnvironment) | **Delete** /cmf/api/v1/environments/{envName} | 
[**DeleteEnvironmentSecretMapping**](DefaultApi.md#DeleteEnvironmentSecretMapping) | **Delete** /cmf/api/v1/environments/{envName}/secret-mappings/{name} | Deletes the Environment Secret Mapping for the given Environment and Secret.
[**DeleteKafkaCatalog**](DefaultApi.md#DeleteKafkaCatalog) | **Delete** /cmf/api/v1/catalogs/kafka/{catName} | Deletes the Kafka Catalog of the given name.
[**DeleteSecret**](DefaultApi.md#DeleteSecret) | **Delete** /cmf/api/v1/secrets/{secretName} | Delete the secret with the given name.
[**DeleteStatement**](DefaultApi.md#DeleteStatement) | **Delete** /cmf/api/v1/environments/{envName}/statements/{stmtName} | Deletes the Statement of the given name in the given Environment.
[**GetApplication**](DefaultApi.md#GetApplication) | **Get** /cmf/api/v1/environments/{envName}/applications/{appName} | Retrieve an Application of the given name in the given Environment.
[**GetApplicationEvents**](DefaultApi.md#GetApplicationEvents) | **Get** /cmf/api/v1alpha1/environments/{envName}/applications/{appName}/events | Get a paginated list of events of the given Application
[**GetApplicationInstance**](DefaultApi.md#GetApplicationInstance) | **Get** /cmf/api/v1/environments/{envName}/applications/{appName}/instances/{instName} | Retrieve an Instance of an Application
[**GetApplicationInstances**](DefaultApi.md#GetApplicationInstances) | **Get** /cmf/api/v1/environments/{envName}/applications/{appName}/instances | Get a paginated list of instances of the given Application
[**GetApplications**](DefaultApi.md#GetApplications) | **Get** /cmf/api/v1/environments/{envName}/applications | Retrieve a paginated list of all applications in the given Environment.
[**GetComputePool**](DefaultApi.md#GetComputePool) | **Get** /cmf/api/v1/environments/{envName}/compute-pools/{computePoolName} | Retrieve the Compute Pool of the given name in the given Environment.
[**GetComputePools**](DefaultApi.md#GetComputePools) | **Get** /cmf/api/v1/environments/{envName}/compute-pools | Retrieve a paginated list of Compute Pools in the given Environment.
[**GetEnvironment**](DefaultApi.md#GetEnvironment) | **Get** /cmf/api/v1/environments/{envName} | Get/Describe an environment with the given name.
[**GetEnvironmentSecretMapping**](DefaultApi.md#GetEnvironmentSecretMapping) | **Get** /cmf/api/v1/environments/{envName}/secret-mappings/{name} | Retrieve the Environment Secret Mapping for the given name in the given environment.
[**GetEnvironmentSecretMappings**](DefaultApi.md#GetEnvironmentSecretMappings) | **Get** /cmf/api/v1/environments/{envName}/secret-mappings | Retrieve a paginated list of all Environment Secret Mappings.
[**GetEnvironments**](DefaultApi.md#GetEnvironments) | **Get** /cmf/api/v1/environments | Retrieve a paginated list of all environments.
[**GetKafkaCatalog**](DefaultApi.md#GetKafkaCatalog) | **Get** /cmf/api/v1/catalogs/kafka/{catName} | Retrieve the Kafka Catalog of the given name.
[**GetKafkaCatalogs**](DefaultApi.md#GetKafkaCatalogs) | **Get** /cmf/api/v1/catalogs/kafka | Retrieve a paginated list of Kafka Catalogs
[**GetSecret**](DefaultApi.md#GetSecret) | **Get** /cmf/api/v1/secrets/{secretName} | Retrieve the Secret of the given name. Note that the secret data is not returned for security reasons.
[**GetSecrets**](DefaultApi.md#GetSecrets) | **Get** /cmf/api/v1/secrets | Retrieve a paginated list of all secrets. Note that the actual secret data is masked for security reasons.
[**GetStatement**](DefaultApi.md#GetStatement) | **Get** /cmf/api/v1/environments/{envName}/statements/{stmtName} | Retrieve the Statement of the given name in the given Environment.
[**GetStatementExceptions**](DefaultApi.md#GetStatementExceptions) | **Get** /cmf/api/v1/environments/{envName}/statements/{stmtName}/exceptions | Retrieves the last 10 exceptions of the Statement with the given name in the given Environment.
[**GetStatementResult**](DefaultApi.md#GetStatementResult) | **Get** /cmf/api/v1/environments/{envName}/statements/{stmtName}/results | Retrieve the result of the interactive Statement with the given name in the given Environment.
[**GetStatements**](DefaultApi.md#GetStatements) | **Get** /cmf/api/v1/environments/{envName}/statements | Retrieve a paginated list of Statements in the given Environment.
[**StartApplication**](DefaultApi.md#StartApplication) | **Post** /cmf/api/v1/environments/{envName}/applications/{appName}/start | Starts an earlier submitted Flink Application
[**SuspendApplication**](DefaultApi.md#SuspendApplication) | **Post** /cmf/api/v1/environments/{envName}/applications/{appName}/suspend | Suspends an earlier started Flink Application
[**UpdateEnvironmentSecretMapping**](DefaultApi.md#UpdateEnvironmentSecretMapping) | **Put** /cmf/api/v1/environments/{envName}/secret-mappings/{name} | Updates the Environment Secret Mapping for the given Environment.
[**UpdateSecret**](DefaultApi.md#UpdateSecret) | **Put** /cmf/api/v1/secrets/{secretName} | Update the secret.
[**UpdateStatement**](DefaultApi.md#UpdateStatement) | **Put** /cmf/api/v1/environments/{envName}/statements/{stmtName} | Updates a Statement of the given name in the given Environment.



## CreateComputePool

> ComputePool CreateComputePool(ctx, envName).ComputePool(computePool).Execute()

Creates a new Flink Compute Pool in the given Environment.

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
    computePool := *openapiclient.NewComputePool("ApiVersion_example", "Kind_example", *openapiclient.NewComputePoolMetadata("Name_example"), *openapiclient.NewComputePoolSpec("Type_example", map[string]interface{}(123))) // ComputePool | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateComputePool(context.Background(), envName).ComputePool(computePool).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateComputePool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateComputePool`: ComputePool
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateComputePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateComputePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **computePool** | [**ComputePool**](ComputePool.md) |  | 

### Return type

[**ComputePool**](ComputePool.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    resp, r, err := api_client.DefaultApi.CreateEnvironmentSecretMapping(context.Background(), envName).EnvironmentSecretMapping(environmentSecretMapping).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateEnvironmentSecretMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEnvironmentSecretMapping`: EnvironmentSecretMapping
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateEnvironmentSecretMapping`: %v\n", resp)
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


## CreateKafkaCatalog

> KafkaCatalog CreateKafkaCatalog(ctx).KafkaCatalog(kafkaCatalog).Execute()

Creates a new Kafka Catalog that can be referenced by Flink Statements

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
    kafkaCatalog := *openapiclient.NewKafkaCatalog("ApiVersion_example", "Kind_example", *openapiclient.NewCatalogMetadata("Name_example"), *openapiclient.NewKafkaCatalogSpec(*openapiclient.NewKafkaCatalogSpecSrInstance(map[string]string{"key": "Inner_example"}), []openapiclient.KafkaCatalogSpecKafkaClusters{*openapiclient.NewKafkaCatalogSpecKafkaClusters("DatabaseName_example", map[string]string{"key": "Inner_example"})})) // KafkaCatalog | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateKafkaCatalog(context.Background()).KafkaCatalog(kafkaCatalog).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateKafkaCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateKafkaCatalog`: KafkaCatalog
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateKafkaCatalog`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateKafkaCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kafkaCatalog** | [**KafkaCatalog**](KafkaCatalog.md) |  | 

### Return type

[**KafkaCatalog**](KafkaCatalog.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    resp, r, err := api_client.DefaultApi.CreateOrUpdateApplication(context.Background(), envName).FlinkApplication(flinkApplication).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateOrUpdateApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrUpdateApplication`: FlinkApplication
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateOrUpdateApplication`: %v\n", resp)
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
    resp, r, err := api_client.DefaultApi.CreateOrUpdateEnvironment(context.Background()).PostEnvironment(postEnvironment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateOrUpdateEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrUpdateEnvironment`: Environment
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateOrUpdateEnvironment`: %v\n", resp)
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
    resp, r, err := api_client.DefaultApi.CreateSecret(context.Background()).Secret(secret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSecret`: Secret
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSecret`: %v\n", resp)
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


## CreateStatement

> Statement CreateStatement(ctx, envName).Statement(statement).Execute()

Creates a new Flink SQL Statement in the given Environment.

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
    statement := *openapiclient.NewStatement("ApiVersion_example", "Kind_example", *openapiclient.NewStatementMetadata("Name_example"), *openapiclient.NewStatementSpec("Statement_example", "ComputePoolName_example")) // Statement | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateStatement(context.Background(), envName).Statement(statement).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateStatement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStatement`: Statement
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateStatement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStatementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **statement** | [**Statement**](Statement.md) |  | 

### Return type

[**Statement**](Statement.md)

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
    resp, r, err := api_client.DefaultApi.DeleteApplication(context.Background(), envName, appName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteApplication``: %v\n", err)
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


## DeleteComputePool

> DeleteComputePool(ctx, envName, computePoolName).Execute()

Deletes the ComputePool of the given name in the given Environment.

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
    computePoolName := "computePoolName_example" // string | Name of the ComputePool

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteComputePool(context.Background(), envName, computePoolName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteComputePool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment | 
**computePoolName** | **string** | Name of the ComputePool | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteComputePoolRequest struct via the builder pattern


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
    resp, r, err := api_client.DefaultApi.DeleteEnvironment(context.Background(), envName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteEnvironment``: %v\n", err)
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
    resp, r, err := api_client.DefaultApi.DeleteEnvironmentSecretMapping(context.Background(), envName, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteEnvironmentSecretMapping``: %v\n", err)
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


## DeleteKafkaCatalog

> DeleteKafkaCatalog(ctx, catName).Execute()

Deletes the Kafka Catalog of the given name.

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
    catName := "catName_example" // string | Name of the Kafka Catalog

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteKafkaCatalog(context.Background(), catName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteKafkaCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**catName** | **string** | Name of the Kafka Catalog | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKafkaCatalogRequest struct via the builder pattern


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
    resp, r, err := api_client.DefaultApi.DeleteSecret(context.Background(), secretName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSecret``: %v\n", err)
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


## DeleteStatement

> DeleteStatement(ctx, envName, stmtName).Execute()

Deletes the Statement of the given name in the given Environment.

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteStatement(context.Background(), envName, stmtName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteStatement``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStatementRequest struct via the builder pattern


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

> FlinkApplication GetApplication(ctx, envName, appName).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetApplication(context.Background(), envName, appName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplication`: FlinkApplication
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetApplication`: %v\n", resp)
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
    resp, r, err := api_client.DefaultApi.GetApplicationEvents(context.Background(), envName, appName).Page(page).Size(size).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetApplicationEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationEvents`: EventsPage
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetApplicationEvents`: %v\n", resp)
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
    resp, r, err := api_client.DefaultApi.GetApplicationInstance(context.Background(), envName, appName, instName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetApplicationInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationInstance`: FlinkApplicationInstance
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetApplicationInstance`: %v\n", resp)
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
    resp, r, err := api_client.DefaultApi.GetApplicationInstances(context.Background(), envName, appName).Page(page).Size(size).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetApplicationInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationInstances`: ApplicationInstancesPage
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetApplicationInstances`: %v\n", resp)
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

> ApplicationsPage GetApplications(ctx, envName).Page(page).Size(size).Sort(sort).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetApplications(context.Background(), envName).Page(page).Size(size).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplications`: ApplicationsPage
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetApplications`: %v\n", resp)
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


## GetComputePool

> ComputePool GetComputePool(ctx, envName, computePoolName).Execute()

Retrieve the Compute Pool of the given name in the given Environment.

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
    computePoolName := "computePoolName_example" // string | Name of the Compute Pool

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetComputePool(context.Background(), envName, computePoolName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetComputePool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetComputePool`: ComputePool
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetComputePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment | 
**computePoolName** | **string** | Name of the Compute Pool | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetComputePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ComputePool**](ComputePool.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComputePools

> ComputePoolsPage GetComputePools(ctx, envName).Page(page).Size(size).Sort(sort).Execute()

Retrieve a paginated list of Compute Pools in the given Environment.

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
    resp, r, err := api_client.DefaultApi.GetComputePools(context.Background(), envName).Page(page).Size(size).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetComputePools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetComputePools`: ComputePoolsPage
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetComputePools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetComputePoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Zero-based page index (0..N) | 
 **size** | **int32** | The size of the page to be returned | 
 **sort** | **[]string** | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**ComputePoolsPage**](ComputePoolsPage.md)

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
    resp, r, err := api_client.DefaultApi.GetEnvironment(context.Background(), envName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnvironment`: Environment
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetEnvironment`: %v\n", resp)
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
    resp, r, err := api_client.DefaultApi.GetEnvironmentSecretMapping(context.Background(), envName, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetEnvironmentSecretMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnvironmentSecretMapping`: EnvironmentSecretMapping
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetEnvironmentSecretMapping`: %v\n", resp)
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
    resp, r, err := api_client.DefaultApi.GetEnvironmentSecretMappings(context.Background(), envName).Page(page).Size(size).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetEnvironmentSecretMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnvironmentSecretMappings`: EnvironmentSecretMappingsPage
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetEnvironmentSecretMappings`: %v\n", resp)
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
    resp, r, err := api_client.DefaultApi.GetEnvironments(context.Background()).Page(page).Size(size).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetEnvironments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnvironments`: EnvironmentsPage
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetEnvironments`: %v\n", resp)
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


## GetKafkaCatalog

> KafkaCatalog GetKafkaCatalog(ctx, catName).Execute()

Retrieve the Kafka Catalog of the given name.

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
    catName := "catName_example" // string | Name of the Kafka Catalog

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetKafkaCatalog(context.Background(), catName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetKafkaCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKafkaCatalog`: KafkaCatalog
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetKafkaCatalog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**catName** | **string** | Name of the Kafka Catalog | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKafkaCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KafkaCatalog**](KafkaCatalog.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKafkaCatalogs

> KafkaCatalogsPage GetKafkaCatalogs(ctx).Page(page).Size(size).Sort(sort).Execute()

Retrieve a paginated list of Kafka Catalogs

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
    resp, r, err := api_client.DefaultApi.GetKafkaCatalogs(context.Background()).Page(page).Size(size).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetKafkaCatalogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKafkaCatalogs`: KafkaCatalogsPage
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetKafkaCatalogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKafkaCatalogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Zero-based page index (0..N) | 
 **size** | **int32** | The size of the page to be returned | 
 **sort** | **[]string** | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**KafkaCatalogsPage**](KafkaCatalogsPage.md)

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
    resp, r, err := api_client.DefaultApi.GetSecret(context.Background(), secretName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecret`: Secret
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSecret`: %v\n", resp)
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

> SecretsPage GetSecrets(ctx).Page(page).Size(size).Sort(sort).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetSecrets(context.Background()).Page(page).Size(size).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSecrets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecrets`: SecretsPage
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSecrets`: %v\n", resp)
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


## GetStatement

> Statement GetStatement(ctx, envName, stmtName).Execute()

Retrieve the Statement of the given name in the given Environment.

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetStatement(context.Background(), envName, stmtName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetStatement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatement`: Statement
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetStatement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment | 
**stmtName** | **string** | Name of the Statement | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Statement**](Statement.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatementExceptions

> StatementExceptionList GetStatementExceptions(ctx, envName, stmtName).Execute()

Retrieves the last 10 exceptions of the Statement with the given name in the given Environment.

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetStatementExceptions(context.Background(), envName, stmtName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetStatementExceptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatementExceptions`: StatementExceptionList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetStatementExceptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment | 
**stmtName** | **string** | Name of the Statement | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatementExceptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**StatementExceptionList**](StatementExceptionList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatementResult

> StatementResult GetStatementResult(ctx, envName, stmtName).PageToken(pageToken).Execute()

Retrieve the result of the interactive Statement with the given name in the given Environment.

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
    pageToken := "pageToken_example" // string | Token for the next page of results (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetStatementResult(context.Background(), envName, stmtName).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetStatementResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatementResult`: StatementResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetStatementResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment | 
**stmtName** | **string** | Name of the Statement | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatementResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageToken** | **string** | Token for the next page of results | 

### Return type

[**StatementResult**](StatementResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatements

> StatementsPage GetStatements(ctx, envName).Page(page).Size(size).Sort(sort).ComputePool(computePool).Phase(phase).Execute()

Retrieve a paginated list of Statements in the given Environment.

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
    computePool := "computePool_example" // string | Name of the ComputePool to filter on (optional)
    phase := "phase_example" // string | Phase to filter on (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetStatements(context.Background(), envName).Page(page).Size(size).Sort(sort).ComputePool(computePool).Phase(phase).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetStatements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatements`: StatementsPage
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetStatements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string** | Name of the Environment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Zero-based page index (0..N) | 
 **size** | **int32** | The size of the page to be returned | 
 **sort** | **[]string** | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 
 **computePool** | **string** | Name of the ComputePool to filter on | 
 **phase** | **string** | Phase to filter on | 

### Return type

[**StatementsPage**](StatementsPage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartApplication

> FlinkApplication StartApplication(ctx, envName, appName).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.StartApplication(context.Background(), envName, appName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.StartApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartApplication`: FlinkApplication
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.StartApplication`: %v\n", resp)
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
    resp, r, err := api_client.DefaultApi.SuspendApplication(context.Background(), envName, appName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SuspendApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SuspendApplication`: FlinkApplication
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SuspendApplication`: %v\n", resp)
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
    resp, r, err := api_client.DefaultApi.UpdateEnvironmentSecretMapping(context.Background(), envName, name).EnvironmentSecretMapping(environmentSecretMapping).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateEnvironmentSecretMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEnvironmentSecretMapping`: EnvironmentSecretMapping
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateEnvironmentSecretMapping`: %v\n", resp)
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
    resp, r, err := api_client.DefaultApi.UpdateSecret(context.Background(), secretName).Secret(secret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSecret`: Secret
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateSecret`: %v\n", resp)
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


## UpdateStatement

> UpdateStatement(ctx, envName, stmtName).Statement(statement).Execute()

Updates a Statement of the given name in the given Environment.

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
    statement := *openapiclient.NewStatement("ApiVersion_example", "Kind_example", *openapiclient.NewStatementMetadata("Name_example"), *openapiclient.NewStatementSpec("Statement_example", "ComputePoolName_example")) // Statement | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateStatement(context.Background(), envName, stmtName).Statement(statement).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateStatement``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStatementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **statement** | [**Statement**](Statement.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

