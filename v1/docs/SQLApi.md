# \SQLApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateComputePool**](SQLApi.md#CreateComputePool) | **Post** /cmf/api/v1/environments/{envName}/compute-pools | Creates a new Flink Compute Pool in the given Environment.
[**CreateKafkaCatalog**](SQLApi.md#CreateKafkaCatalog) | **Post** /cmf/api/v1/catalogs/kafka | Creates a new Kafka Catalog that can be referenced by Flink Statements
[**CreateKafkaDatabase**](SQLApi.md#CreateKafkaDatabase) | **Post** /cmf/api/v1/catalogs/kafka/{catName}/databases | Creates a new Kafka Database
[**CreateStatement**](SQLApi.md#CreateStatement) | **Post** /cmf/api/v1/environments/{envName}/statements | Creates a new Flink SQL Statement in the given Environment.
[**DeleteComputePool**](SQLApi.md#DeleteComputePool) | **Delete** /cmf/api/v1/environments/{envName}/compute-pools/{computePoolName} | Deletes the ComputePool of the given name in the given Environment.
[**DeleteKafkaCatalog**](SQLApi.md#DeleteKafkaCatalog) | **Delete** /cmf/api/v1/catalogs/kafka/{catName} | Deletes the Kafka Catalog of the given name.
[**DeleteKafkaDatabase**](SQLApi.md#DeleteKafkaDatabase) | **Delete** /cmf/api/v1/catalogs/kafka/{catName}/databases/{dbName} | Deletes the Kafka Database of the given name in the given KafkaCatalog.
[**DeleteStatement**](SQLApi.md#DeleteStatement) | **Delete** /cmf/api/v1/environments/{envName}/statements/{stmtName} | Deletes the Statement of the given name in the given Environment.
[**GetComputePool**](SQLApi.md#GetComputePool) | **Get** /cmf/api/v1/environments/{envName}/compute-pools/{computePoolName} | Retrieve the Compute Pool of the given name in the given Environment.
[**GetComputePools**](SQLApi.md#GetComputePools) | **Get** /cmf/api/v1/environments/{envName}/compute-pools | Retrieve a paginated list of Compute Pools in the given Environment.
[**GetKafkaCatalog**](SQLApi.md#GetKafkaCatalog) | **Get** /cmf/api/v1/catalogs/kafka/{catName} | Retrieve the Kafka Catalog of the given name.
[**GetKafkaCatalogs**](SQLApi.md#GetKafkaCatalogs) | **Get** /cmf/api/v1/catalogs/kafka | Retrieve a paginated list of Kafka Catalogs
[**GetKafkaDatabase**](SQLApi.md#GetKafkaDatabase) | **Get** /cmf/api/v1/catalogs/kafka/{catName}/databases/{dbName} | Retrieve the Kafka Database of the given name in the given KafkaCatalog.
[**GetKafkaDatabases**](SQLApi.md#GetKafkaDatabases) | **Get** /cmf/api/v1/catalogs/kafka/{catName}/databases | Retrieve a paginated list of Kafka Databases
[**GetStatement**](SQLApi.md#GetStatement) | **Get** /cmf/api/v1/environments/{envName}/statements/{stmtName} | Retrieve the Statement of the given name in the given Environment.
[**GetStatementExceptions**](SQLApi.md#GetStatementExceptions) | **Get** /cmf/api/v1/environments/{envName}/statements/{stmtName}/exceptions | Retrieves the last 10 exceptions of the Statement with the given name in the given Environment.
[**GetStatementResult**](SQLApi.md#GetStatementResult) | **Get** /cmf/api/v1/environments/{envName}/statements/{stmtName}/results | Retrieve the result of the interactive Statement with the given name in the given Environment.
[**GetStatements**](SQLApi.md#GetStatements) | **Get** /cmf/api/v1/environments/{envName}/statements | Retrieve a paginated list of Statements in the given Environment.
[**UpdateKafkaCatalog**](SQLApi.md#UpdateKafkaCatalog) | **Put** /cmf/api/v1/catalogs/kafka/{catName} | Updates a KafkaCatalog of the given name.
[**UpdateKafkaDatabase**](SQLApi.md#UpdateKafkaDatabase) | **Put** /cmf/api/v1/catalogs/kafka/{catName}/databases/{dbName} | Updates a KafkaDatabase of the given name in the given KafkaCatalog.
[**UpdateStatement**](SQLApi.md#UpdateStatement) | **Put** /cmf/api/v1/environments/{envName}/statements/{stmtName} | Updates a Statement of the given name in the given Environment.



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
    resp, r, err := api_client.SQLApi.CreateComputePool(context.Background(), envName).ComputePool(computePool).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SQLApi.CreateComputePool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateComputePool`: ComputePool
    fmt.Fprintf(os.Stdout, "Response from `SQLApi.CreateComputePool`: %v\n", resp)
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
    kafkaCatalog := *openapiclient.NewKafkaCatalog("ApiVersion_example", "Kind_example", *openapiclient.NewCatalogMetadata("Name_example"), *openapiclient.NewKafkaCatalogSpec(*openapiclient.NewKafkaCatalogSpecSrInstance(map[string]string{"key": "Inner_example"}))) // KafkaCatalog | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SQLApi.CreateKafkaCatalog(context.Background()).KafkaCatalog(kafkaCatalog).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SQLApi.CreateKafkaCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateKafkaCatalog`: KafkaCatalog
    fmt.Fprintf(os.Stdout, "Response from `SQLApi.CreateKafkaCatalog`: %v\n", resp)
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


## CreateKafkaDatabase

> KafkaDatabase CreateKafkaDatabase(ctx, catName).KafkaDatabase(kafkaDatabase).Execute()

Creates a new Kafka Database

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
    catName := "catName_example" // string | Name of the Catalog
    kafkaDatabase := *openapiclient.NewKafkaDatabase("ApiVersion_example", "Kind_example", *openapiclient.NewDatabaseMetadata("Name_example"), *openapiclient.NewKafkaDatabaseSpec(*openapiclient.NewKafkaDatabaseSpecKafkaCluster(map[string]string{"key": "Inner_example"}))) // KafkaDatabase | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SQLApi.CreateKafkaDatabase(context.Background(), catName).KafkaDatabase(kafkaDatabase).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SQLApi.CreateKafkaDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateKafkaDatabase`: KafkaDatabase
    fmt.Fprintf(os.Stdout, "Response from `SQLApi.CreateKafkaDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**catName** | **string** | Name of the Catalog | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateKafkaDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kafkaDatabase** | [**KafkaDatabase**](KafkaDatabase.md) |  | 

### Return type

[**KafkaDatabase**](KafkaDatabase.md)

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
    resp, r, err := api_client.SQLApi.CreateStatement(context.Background(), envName).Statement(statement).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SQLApi.CreateStatement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStatement`: Statement
    fmt.Fprintf(os.Stdout, "Response from `SQLApi.CreateStatement`: %v\n", resp)
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
    resp, r, err := api_client.SQLApi.DeleteComputePool(context.Background(), envName, computePoolName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SQLApi.DeleteComputePool``: %v\n", err)
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
    resp, r, err := api_client.SQLApi.DeleteKafkaCatalog(context.Background(), catName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SQLApi.DeleteKafkaCatalog``: %v\n", err)
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


## DeleteKafkaDatabase

> DeleteKafkaDatabase(ctx, catName, dbName).Execute()

Deletes the Kafka Database of the given name in the given KafkaCatalog.

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
    dbName := "dbName_example" // string | Name of the Kafka Database

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SQLApi.DeleteKafkaDatabase(context.Background(), catName, dbName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SQLApi.DeleteKafkaDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**catName** | **string** | Name of the Kafka Catalog | 
**dbName** | **string** | Name of the Kafka Database | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKafkaDatabaseRequest struct via the builder pattern


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
    resp, r, err := api_client.SQLApi.DeleteStatement(context.Background(), envName, stmtName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SQLApi.DeleteStatement``: %v\n", err)
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
    resp, r, err := api_client.SQLApi.GetComputePool(context.Background(), envName, computePoolName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SQLApi.GetComputePool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetComputePool`: ComputePool
    fmt.Fprintf(os.Stdout, "Response from `SQLApi.GetComputePool`: %v\n", resp)
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
    resp, r, err := api_client.SQLApi.GetComputePools(context.Background(), envName).Page(page).Size(size).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SQLApi.GetComputePools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetComputePools`: ComputePoolsPage
    fmt.Fprintf(os.Stdout, "Response from `SQLApi.GetComputePools`: %v\n", resp)
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
    resp, r, err := api_client.SQLApi.GetKafkaCatalog(context.Background(), catName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SQLApi.GetKafkaCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKafkaCatalog`: KafkaCatalog
    fmt.Fprintf(os.Stdout, "Response from `SQLApi.GetKafkaCatalog`: %v\n", resp)
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
    resp, r, err := api_client.SQLApi.GetKafkaCatalogs(context.Background()).Page(page).Size(size).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SQLApi.GetKafkaCatalogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKafkaCatalogs`: KafkaCatalogsPage
    fmt.Fprintf(os.Stdout, "Response from `SQLApi.GetKafkaCatalogs`: %v\n", resp)
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


## GetKafkaDatabase

> KafkaDatabase GetKafkaDatabase(ctx, catName, dbName).Execute()

Retrieve the Kafka Database of the given name in the given KafkaCatalog.

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
    dbName := "dbName_example" // string | Name of the Kafka Database

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SQLApi.GetKafkaDatabase(context.Background(), catName, dbName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SQLApi.GetKafkaDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKafkaDatabase`: KafkaDatabase
    fmt.Fprintf(os.Stdout, "Response from `SQLApi.GetKafkaDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**catName** | **string** | Name of the Kafka Catalog | 
**dbName** | **string** | Name of the Kafka Database | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKafkaDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**KafkaDatabase**](KafkaDatabase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKafkaDatabases

> KafkaDatabasesPage GetKafkaDatabases(ctx, catName).Page(page).Size(size).Sort(sort).Execute()

Retrieve a paginated list of Kafka Databases

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
    catName := "catName_example" // string | Name of the Catalog
    page := int32(56) // int32 | Zero-based page index (0..N) (optional)
    size := int32(56) // int32 | The size of the page to be returned (optional)
    sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SQLApi.GetKafkaDatabases(context.Background(), catName).Page(page).Size(size).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SQLApi.GetKafkaDatabases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKafkaDatabases`: KafkaDatabasesPage
    fmt.Fprintf(os.Stdout, "Response from `SQLApi.GetKafkaDatabases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**catName** | **string** | Name of the Catalog | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKafkaDatabasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Zero-based page index (0..N) | 
 **size** | **int32** | The size of the page to be returned | 
 **sort** | **[]string** | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**KafkaDatabasesPage**](KafkaDatabasesPage.md)

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
    resp, r, err := api_client.SQLApi.GetStatement(context.Background(), envName, stmtName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SQLApi.GetStatement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatement`: Statement
    fmt.Fprintf(os.Stdout, "Response from `SQLApi.GetStatement`: %v\n", resp)
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
    resp, r, err := api_client.SQLApi.GetStatementExceptions(context.Background(), envName, stmtName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SQLApi.GetStatementExceptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatementExceptions`: StatementExceptionList
    fmt.Fprintf(os.Stdout, "Response from `SQLApi.GetStatementExceptions`: %v\n", resp)
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
    resp, r, err := api_client.SQLApi.GetStatementResult(context.Background(), envName, stmtName).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SQLApi.GetStatementResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatementResult`: StatementResult
    fmt.Fprintf(os.Stdout, "Response from `SQLApi.GetStatementResult`: %v\n", resp)
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
    resp, r, err := api_client.SQLApi.GetStatements(context.Background(), envName).Page(page).Size(size).Sort(sort).ComputePool(computePool).Phase(phase).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SQLApi.GetStatements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatements`: StatementsPage
    fmt.Fprintf(os.Stdout, "Response from `SQLApi.GetStatements`: %v\n", resp)
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


## UpdateKafkaCatalog

> UpdateKafkaCatalog(ctx, catName).KafkaCatalog(kafkaCatalog).Execute()

Updates a KafkaCatalog of the given name.

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
    catName := "catName_example" // string | Name of the KafkaCatalog
    kafkaCatalog := *openapiclient.NewKafkaCatalog("ApiVersion_example", "Kind_example", *openapiclient.NewCatalogMetadata("Name_example"), *openapiclient.NewKafkaCatalogSpec(*openapiclient.NewKafkaCatalogSpecSrInstance(map[string]string{"key": "Inner_example"}))) // KafkaCatalog | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SQLApi.UpdateKafkaCatalog(context.Background(), catName).KafkaCatalog(kafkaCatalog).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SQLApi.UpdateKafkaCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**catName** | **string** | Name of the KafkaCatalog | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKafkaCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kafkaCatalog** | [**KafkaCatalog**](KafkaCatalog.md) |  | 

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


## UpdateKafkaDatabase

> UpdateKafkaDatabase(ctx, catName, dbName).KafkaDatabase(kafkaDatabase).Execute()

Updates a KafkaDatabase of the given name in the given KafkaCatalog.

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
    catName := "catName_example" // string | Name of the KafkaCatalog
    dbName := "dbName_example" // string | Name of the KafkaDatabase
    kafkaDatabase := *openapiclient.NewKafkaDatabase("ApiVersion_example", "Kind_example", *openapiclient.NewDatabaseMetadata("Name_example"), *openapiclient.NewKafkaDatabaseSpec(*openapiclient.NewKafkaDatabaseSpecKafkaCluster(map[string]string{"key": "Inner_example"}))) // KafkaDatabase | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SQLApi.UpdateKafkaDatabase(context.Background(), catName, dbName).KafkaDatabase(kafkaDatabase).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SQLApi.UpdateKafkaDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**catName** | **string** | Name of the KafkaCatalog | 
**dbName** | **string** | Name of the KafkaDatabase | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKafkaDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **kafkaDatabase** | [**KafkaDatabase**](KafkaDatabase.md) |  | 

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
    resp, r, err := api_client.SQLApi.UpdateStatement(context.Background(), envName, stmtName).Statement(statement).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SQLApi.UpdateStatement``: %v\n", err)
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

