# \DefaultApi

All URIs are relative to *http://localhost:8080/cmf/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrUpdateApplication**](DefaultApi.md#CreateOrUpdateApplication) | **Post** /environments/{name}/applications | Creates a new Flink Application or updates an existing one.
[**CreateOrUpdateEnvironment**](DefaultApi.md#CreateOrUpdateEnvironment) | **Post** /environments | Create or update an Environment
[**DeleteApplication**](DefaultApi.md#DeleteApplication) | **Delete** /environments/{name}/applications/{appName} | Deletes an Application
[**DeleteEnvironment**](DefaultApi.md#DeleteEnvironment) | **Delete** /environments/{name} | 
[**GetApplication**](DefaultApi.md#GetApplication) | **Get** /environments/{name}/applications/{appName} | Get a Application
[**GetApplications**](DefaultApi.md#GetApplications) | **Get** /environments/{name}/applications | Lists Applications
[**GetEnvironments**](DefaultApi.md#GetEnvironments) | **Get** /environments | Lists Environments
[**StartApplication**](DefaultApi.md#StartApplication) | **Post** /environments/{name}/applications/{appName}/start | Starts an earlier submitted Flink Application
[**StopApplication**](DefaultApi.md#StopApplication) | **Post** /environments/{name}/applications/{appName}/stop | Stops an earlier started Flink Application



## CreateOrUpdateApplication

> Application CreateOrUpdateApplication(ctx, name, application)

Creates a new Flink Application or updates an existing one.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Name of the Environment | 
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

> GetEnvironment CreateOrUpdateEnvironment(ctx, postEnvironment)

Create or update an Environment

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**postEnvironment** | [**PostEnvironment**](PostEnvironment.md)|  | 

### Return type

[**GetEnvironment**](GetEnvironment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplication

> DeleteApplication(ctx, name, appName)

Deletes an Application

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Name of the Environment | 
**appName** | **string**| Name of the Application | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEnvironment

> DeleteEnvironment(ctx, name)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Name of the Environment | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplication

> Application GetApplication(ctx, name, appName, optional)

Get a Application

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Name of the Environment | 
**appName** | **string**| Name of the Application | 
 **optional** | ***GetApplicationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetApplicationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fieldSelector** | **optional.String**| Select fields to return | 

### Return type

[**Application**](Application.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplications

> ApplicationsPage GetApplications(ctx, name, optional)

Lists Applications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Name of the Environment | 
 **optional** | ***GetApplicationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetApplicationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Zero-based page index (0..N) | 
 **size** | **optional.Int32**| The size of the page to be returned | 
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property,(asc|desc). Multiple sort criteria are supported. | 
 **watch** | **optional.Bool**| Subscribe to any updates | 
 **fieldSelector** | **optional.String**| Select fields to return | 

### Return type

[**ApplicationsPage**](ApplicationsPage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironments

> EnvironmentsPage GetEnvironments(ctx, optional)

Lists Environments

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
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property,(asc|desc). Multiple sort criteria are supported. | 

### Return type

[**EnvironmentsPage**](EnvironmentsPage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartApplication

> Application StartApplication(ctx, name, appName, inlineObject)

Starts an earlier submitted Flink Application

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Name of the Environment | 
**appName** | **string**| Name of the Application | 
**inlineObject** | [**InlineObject**](InlineObject.md)|  | 

### Return type

[**Application**](Application.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopApplication

> Application StopApplication(ctx, name, appName)

Stops an earlier started Flink Application

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Name of the Environment | 
**appName** | **string**| Name of the Application | 

### Return type

[**Application**](Application.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

