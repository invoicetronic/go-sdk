# \WebhookAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WebhookGet**](WebhookAPI.md#WebhookGet) | **Get** /webhook | List webhooks
[**WebhookIdDelete**](WebhookAPI.md#WebhookIdDelete) | **Delete** /webhook/{id} | Delete a webhook by id
[**WebhookIdGet**](WebhookAPI.md#WebhookIdGet) | **Get** /webhook/{id} | Get a webhook by id
[**WebhookPost**](WebhookAPI.md#WebhookPost) | **Post** /webhook | Add a webhook
[**WebhookPut**](WebhookAPI.md#WebhookPut) | **Put** /webhook | Update a webhook
[**WebhookhistoryGet**](WebhookAPI.md#WebhookhistoryGet) | **Get** /webhookhistory | List webhook history items
[**WebhookhistoryIdGet**](WebhookAPI.md#WebhookhistoryIdGet) | **Get** /webhookhistory/{id} | Get a webhook history item by id



## WebhookGet

> []WebHook WebhookGet(ctx).CompanyId(companyId).Page(page).PageSize(pageSize).Sort(sort).Description(description).Enabled(enabled).Events(events).Url(url).Execute()

List webhooks



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	invoicetronicsdk "github.com/invoicetronic/go-sdk"
)

func main() {
	companyId := int32(56) // int32 | Company id (optional)
	page := int32(56) // int32 | Page number. Defaults to 1. (optional) (default to 1)
	pageSize := int32(56) // int32 | Items per page. Defaults to 50. Cannot be greater than 200. (optional) (default to 100)
	sort := "sort_example" // string | Sort by field. Prefix with '-' for descending order. (optional)
	description := "description_example" // string |  (optional)
	enabled := true // bool |  (optional)
	events := "events_example" // string |  (optional)
	url := "url_example" // string |  (optional)

	configuration := invoicetronicsdk.NewConfiguration()
	apiClient := invoicetronicsdk.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.WebhookGet(context.Background()).CompanyId(companyId).Page(page).PageSize(pageSize).Sort(sort).Description(description).Enabled(enabled).Events(events).Url(url).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.WebhookGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookGet`: []WebHook
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.WebhookGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | Company id | 
 **page** | **int32** | Page number. Defaults to 1. | [default to 1]
 **pageSize** | **int32** | Items per page. Defaults to 50. Cannot be greater than 200. | [default to 100]
 **sort** | **string** | Sort by field. Prefix with &#39;-&#39; for descending order. | 
 **description** | **string** |  | 
 **enabled** | **bool** |  | 
 **events** | **string** |  | 
 **url** | **string** |  | 

### Return type

[**[]WebHook**](WebHook.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookIdDelete

> WebHook WebhookIdDelete(ctx, id).Execute()

Delete a webhook by id



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	invoicetronicsdk "github.com/invoicetronic/go-sdk"
)

func main() {
	id := int32(56) // int32 | Item id

	configuration := invoicetronicsdk.NewConfiguration()
	apiClient := invoicetronicsdk.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.WebhookIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.WebhookIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookIdDelete`: WebHook
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.WebhookIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Item id | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebhookIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WebHook**](WebHook.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookIdGet

> WebHook WebhookIdGet(ctx, id).Execute()

Get a webhook by id



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	invoicetronicsdk "github.com/invoicetronic/go-sdk"
)

func main() {
	id := int32(56) // int32 | Item id

	configuration := invoicetronicsdk.NewConfiguration()
	apiClient := invoicetronicsdk.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.WebhookIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.WebhookIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookIdGet`: WebHook
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.WebhookIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Item id | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebhookIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WebHook**](WebHook.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookPost

> WebHook WebhookPost(ctx).WebHook(webHook).Execute()

Add a webhook



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	invoicetronicsdk "github.com/invoicetronic/go-sdk"
)

func main() {
	webHook := *invoicetronicsdk.NewWebHook() // WebHook | 

	configuration := invoicetronicsdk.NewConfiguration()
	apiClient := invoicetronicsdk.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.WebhookPost(context.Background()).WebHook(webHook).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.WebhookPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookPost`: WebHook
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.WebhookPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webHook** | [**WebHook**](WebHook.md) |  | 

### Return type

[**WebHook**](WebHook.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookPut

> WebHook WebhookPut(ctx).WebHook(webHook).Execute()

Update a webhook



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	invoicetronicsdk "github.com/invoicetronic/go-sdk"
)

func main() {
	webHook := *invoicetronicsdk.NewWebHook() // WebHook | 

	configuration := invoicetronicsdk.NewConfiguration()
	apiClient := invoicetronicsdk.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.WebhookPut(context.Background()).WebHook(webHook).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.WebhookPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookPut`: WebHook
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.WebhookPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webHook** | [**WebHook**](WebHook.md) |  | 

### Return type

[**WebHook**](WebHook.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookhistoryGet

> []WebHookHistory WebhookhistoryGet(ctx).Page(page).PageSize(pageSize).Sort(sort).WebhookId(webhookId).Execute()

List webhook history items



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	invoicetronicsdk "github.com/invoicetronic/go-sdk"
)

func main() {
	page := int32(56) // int32 | Page number. Defaults to 1. (optional) (default to 1)
	pageSize := int32(56) // int32 | Items per page. Defaults to 50. Cannot be greater than 200. (optional) (default to 100)
	sort := "sort_example" // string | Sort by field. Prefix with '-' for descending order. (optional)
	webhookId := int32(56) // int32 | WebHook id (optional)

	configuration := invoicetronicsdk.NewConfiguration()
	apiClient := invoicetronicsdk.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.WebhookhistoryGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).WebhookId(webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.WebhookhistoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookhistoryGet`: []WebHookHistory
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.WebhookhistoryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookhistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number. Defaults to 1. | [default to 1]
 **pageSize** | **int32** | Items per page. Defaults to 50. Cannot be greater than 200. | [default to 100]
 **sort** | **string** | Sort by field. Prefix with &#39;-&#39; for descending order. | 
 **webhookId** | **int32** | WebHook id | 

### Return type

[**[]WebHookHistory**](WebHookHistory.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookhistoryIdGet

> WebHookHistory WebhookhistoryIdGet(ctx, id).Execute()

Get a webhook history item by id



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	invoicetronicsdk "github.com/invoicetronic/go-sdk"
)

func main() {
	id := int32(56) // int32 | Item id

	configuration := invoicetronicsdk.NewConfiguration()
	apiClient := invoicetronicsdk.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.WebhookhistoryIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.WebhookhistoryIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookhistoryIdGet`: WebHookHistory
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.WebhookhistoryIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Item id | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebhookhistoryIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WebHookHistory**](WebHookHistory.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

