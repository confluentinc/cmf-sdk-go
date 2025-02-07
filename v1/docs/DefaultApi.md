# \DefaultApi

All URIs are relative to *http://localhost:8080/cmf/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrUpdateApplication**](DefaultApi.md#CreateOrUpdateApplication) | **Post** /environments/{envName}/applications | Creates a new Flink Application or updates an existing one in the given Environment.
[**CreateOrUpdateEnvironment**](DefaultApi.md#CreateOrUpdateEnvironment) | **Post** /environments | Create or update an Environment
[**DeleteApplication**](DefaultApi.md#DeleteApplication) | **Delete** /environments/{envName}/applications/{appName} | Deletes an Application of the given name in the given Environment.
[**DeleteEnvironment**](DefaultApi.md#DeleteEnvironment) | **Delete** /environments/{envName} | 
[**GetApplication**](DefaultApi.md#GetApplication) | **Get** /environments/{envName}/applications/{appName} | Retrieve an Application of the given name in the given Environment.
[**GetApplications**](DefaultApi.md#GetApplications) | **Get** /environments/{envName}/applications | Retrieve a paginated list of all applications in the given Environment.
[**GetEnvironment**](DefaultApi.md#GetEnvironment) | **Get** /environments/{envName} | Get/Describe an environment with the given name.
[**GetEnvironments**](DefaultApi.md#GetEnvironments) | **Get** /environments | Retrieve a paginated list of all environments.
[**StartApplication**](DefaultApi.md#StartApplication) | **Post** /environments/{envName}/applications/{appName}/start | Starts an earlier submitted Flink Application
[**SuspendApplication**](DefaultApi.md#SuspendApplication) | **Post** /environments/{envName}/applications/{appName}/suspend | Suspends an earlier started Flink Application



## CreateOrUpdateApplication

> Application CreateOrUpdateApplication(ctx, envName, application)

Creates a new Flink Application or updates an existing one in the given Environment.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string**| Name of the Environment | 
**application** | [**Application**](Application.md)|  | 

### Return type

[**Application**](Application.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrUpdateEnvironment

> Environment CreateOrUpdateEnvironment(ctx, postEnvironment)

Create or update an Environment

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**postEnvironment** | [**PostEnvironment**](PostEnvironment.md)|  | 

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


## DeleteApplication

> DeleteApplication(ctx, envName, appName)

Deletes an Application of the given name in the given Environment.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string**| Name of the Environment | 
**appName** | **string**| Name of the Application | 

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

> DeleteEnvironment(ctx, envName)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string**| Name of the Environment to be deleted. | 

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

> Application GetApplication(ctx, envName, appName)

Retrieve an Application of the given name in the given Environment.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string**| Name of the Environment | 
**appName** | **string**| Name of the Application | 

### Return type

[**Application**](Application.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplications

> ApplicationsPage GetApplications(ctx, envName, optional)

Retrieve a paginated list of all applications in the given Environment.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string**| Name of the Environment | 
 **optional** | ***GetApplicationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetApplicationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Zero-based page index (0..N) | 
 **size** | **optional.Int32**| The size of the page to be returned | 
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

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


## GetEnvironment

> Environment GetEnvironment(ctx, envName)

Get/Describe an environment with the given name.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string**| Name of the Environment to be retrieved. | 

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


## GetEnvironments

> EnvironmentsPage GetEnvironments(ctx, optional)

Retrieve a paginated list of all environments.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetEnvironmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEnvironmentsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Zero-based page index (0..N) | 
 **size** | **optional.Int32**| The size of the page to be returned | 
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

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


## StartApplication

> Application StartApplication(ctx, envName, appName)

Starts an earlier submitted Flink Application

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string**| Name of the Environment | 
**appName** | **string**| Name of the Application | 

### Return type

[**Application**](Application.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuspendApplication

> Application SuspendApplication(ctx, envName, appName)

Suspends an earlier started Flink Application

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envName** | **string**| Name of the Environment | 
**appName** | **string**| Name of the Application | 

### Return type

[**Application**](Application.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

