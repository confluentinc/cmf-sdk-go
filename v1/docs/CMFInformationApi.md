# \CMFInformationApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetResourceInformation**](CMFInformationApi.md#GetResourceInformation) | **Get** /cmf/api/v1/resource-information | Retrieve resource information about the CMF deployment.
[**GetSystemInformation**](CMFInformationApi.md#GetSystemInformation) | **Get** /cmf/api/v1/system-information | Retrieve system information about the CMF deployment.



## GetResourceInformation

> ResourceInformation GetResourceInformation(ctx).Execute()

Retrieve resource information about the CMF deployment.

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
    resp, r, err := api_client.CMFInformationApi.GetResourceInformation(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CMFInformationApi.GetResourceInformation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetResourceInformation`: ResourceInformation
    fmt.Fprintf(os.Stdout, "Response from `CMFInformationApi.GetResourceInformation`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetResourceInformationRequest struct via the builder pattern


### Return type

[**ResourceInformation**](ResourceInformation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemInformation

> SystemInformation GetSystemInformation(ctx).Execute()

Retrieve system information about the CMF deployment.

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
    resp, r, err := api_client.CMFInformationApi.GetSystemInformation(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CMFInformationApi.GetSystemInformation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSystemInformation`: SystemInformation
    fmt.Fprintf(os.Stdout, "Response from `CMFInformationApi.GetSystemInformation`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemInformationRequest struct via the builder pattern


### Return type

[**SystemInformation**](SystemInformation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

