# \C3Api

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetC3LicenseInformation**](C3Api.md#GetC3LicenseInformation) | **Get** /cmf/api/v1/c3/license | Retrieve license information for C3 integration.



## GetC3LicenseInformation

> C3LicenseInformation GetC3LicenseInformation(ctx).Execute()

Retrieve license information for C3 integration.

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

