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
    resp, r, err := api_client.SecretsApi.GetSecrets(context.Background()).Page(page).Size(size).Sort(sort).Execute()
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

