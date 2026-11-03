# \C3Api

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAuthConfig**](C3Api.md#GetAuthConfig) | **Get** /cmf/api/v1/c3/auth-config | Discover the CMF authentication/authorization configuration.
[**GetC3Configuration**](C3Api.md#GetC3Configuration) | **Get** /cmf/api/v1/c3/configuration | Retrieve the effective CMF runtime configuration.
[**GetC3LicenseInformation**](C3Api.md#GetC3LicenseInformation) | **Get** /cmf/api/v1/c3/license | Retrieve license information for Confluent Control Center integration.
[**GetFeatures**](C3Api.md#GetFeatures) | **Get** /cmf/api/v1/c3/features | Discover which optional CMF UI features the deployment exposes.
[**GetFlinkConfigReference**](C3Api.md#GetFlinkConfigReference) | **Get** /cmf/api/v1/c3/flink-config-references/{flinkVersionName} | Retrieve the reference of the Flink configuration options of a Flink version.
[**GetWhoami**](C3Api.md#GetWhoami) | **Get** /cmf/api/v1/c3/whoami | Return the authenticated principal&#39;s identity.



## GetAuthConfig

> AuthConfig GetAuthConfig(ctx).Execute()

Discover the CMF authentication/authorization configuration.



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.C3Api.GetAuthConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `C3Api.GetAuthConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthConfig`: AuthConfig
    fmt.Fprintf(os.Stdout, "Response from `C3Api.GetAuthConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthConfigRequest struct via the builder pattern


### Return type

[**AuthConfig**](AuthConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetC3Configuration

> C3Configuration GetC3Configuration(ctx).Execute()

Retrieve the effective CMF runtime configuration.



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.C3Api.GetC3Configuration(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `C3Api.GetC3Configuration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetC3Configuration`: C3Configuration
    fmt.Fprintf(os.Stdout, "Response from `C3Api.GetC3Configuration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetC3ConfigurationRequest struct via the builder pattern


### Return type

[**C3Configuration**](C3Configuration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetC3LicenseInformation

> C3LicenseInformation GetC3LicenseInformation(ctx).Execute()

Retrieve license information for Confluent Control Center integration.

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.C3Api.GetC3LicenseInformation(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `C3Api.GetC3LicenseInformation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetC3LicenseInformation`: C3LicenseInformation
    fmt.Fprintf(os.Stdout, "Response from `C3Api.GetC3LicenseInformation`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetC3LicenseInformationRequest struct via the builder pattern


### Return type

[**C3LicenseInformation**](C3LicenseInformation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeatures

> C3Features GetFeatures(ctx).Execute()

Discover which optional CMF UI features the deployment exposes.



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.C3Api.GetFeatures(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `C3Api.GetFeatures``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFeatures`: C3Features
    fmt.Fprintf(os.Stdout, "Response from `C3Api.GetFeatures`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeaturesRequest struct via the builder pattern


### Return type

[**C3Features**](C3Features.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlinkConfigReference

> FlinkConfigReference GetFlinkConfigReference(ctx, flinkVersionName).Execute()

Retrieve the reference of the Flink configuration options of a Flink version.



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
    flinkVersionName := "v1_19" // string | The Flink version, as in the flinkVersion of a compute pool's or application's spec, for example v1_19.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.C3Api.GetFlinkConfigReference(context.Background(), flinkVersionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `C3Api.GetFlinkConfigReference``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFlinkConfigReference`: FlinkConfigReference
    fmt.Fprintf(os.Stdout, "Response from `C3Api.GetFlinkConfigReference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flinkVersionName** | **string** | The Flink version, as in the flinkVersion of a compute pool&#39;s or application&#39;s spec, for example v1_19. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFlinkConfigReferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FlinkConfigReference**](FlinkConfigReference.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWhoami

> WhoAmi GetWhoami(ctx).Execute()

Return the authenticated principal's identity.



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.C3Api.GetWhoami(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `C3Api.GetWhoami``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWhoami`: WhoAmi
    fmt.Fprintf(os.Stdout, "Response from `C3Api.GetWhoami`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWhoamiRequest struct via the builder pattern


### Return type

[**WhoAmi**](WhoAmi.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

