# \WebhookAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InvoiceV1WebhookGet**](WebhookAPI.md#InvoiceV1WebhookGet) | **Get** /invoice/v1/webhook | List webhooks
[**InvoiceV1WebhookIdDelete**](WebhookAPI.md#InvoiceV1WebhookIdDelete) | **Delete** /invoice/v1/webhook/{id} | Delete a webhook by id
[**InvoiceV1WebhookIdGet**](WebhookAPI.md#InvoiceV1WebhookIdGet) | **Get** /invoice/v1/webhook/{id} | Get a webhook by id
[**InvoiceV1WebhookPost**](WebhookAPI.md#InvoiceV1WebhookPost) | **Post** /invoice/v1/webhook | Add a webhook
[**InvoiceV1WebhookPut**](WebhookAPI.md#InvoiceV1WebhookPut) | **Put** /invoice/v1/webhook | Update a webhook
[**InvoiceV1WebhookhistoryGet**](WebhookAPI.md#InvoiceV1WebhookhistoryGet) | **Get** /invoice/v1/webhookhistory | List webhook history items
[**InvoiceV1WebhookhistoryIdGet**](WebhookAPI.md#InvoiceV1WebhookhistoryIdGet) | **Get** /invoice/v1/webhookhistory/{id} | Get a webhook history item by id



## InvoiceV1WebhookGet

> []WebHook InvoiceV1WebhookGet(ctx).Page(page).PageSize(pageSize).Execute()

List webhooks



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	invoicesdk "github.com/invoicetronic/invoice-go-sdk"
)

func main() {
	page := int32(56) // int32 | Page number. (optional) (default to 1)
	pageSize := int32(56) // int32 | Items per page. (optional) (default to 100)

	configuration := invoicesdk.NewConfiguration()
	apiClient := invoicesdk.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.InvoiceV1WebhookGet(context.Background()).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.InvoiceV1WebhookGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceV1WebhookGet`: []WebHook
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.InvoiceV1WebhookGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceV1WebhookGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number. | [default to 1]
 **pageSize** | **int32** | Items per page. | [default to 100]

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


## InvoiceV1WebhookIdDelete

> WebHook InvoiceV1WebhookIdDelete(ctx, id).Execute()

Delete a webhook by id



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	invoicesdk "github.com/invoicetronic/invoice-go-sdk"
)

func main() {
	id := int32(56) // int32 | Item id.

	configuration := invoicesdk.NewConfiguration()
	apiClient := invoicesdk.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.InvoiceV1WebhookIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.InvoiceV1WebhookIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceV1WebhookIdDelete`: WebHook
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.InvoiceV1WebhookIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceV1WebhookIdDeleteRequest struct via the builder pattern


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


## InvoiceV1WebhookIdGet

> WebHook InvoiceV1WebhookIdGet(ctx, id).Execute()

Get a webhook by id



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	invoicesdk "github.com/invoicetronic/invoice-go-sdk"
)

func main() {
	id := int32(56) // int32 | Item id.

	configuration := invoicesdk.NewConfiguration()
	apiClient := invoicesdk.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.InvoiceV1WebhookIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.InvoiceV1WebhookIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceV1WebhookIdGet`: WebHook
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.InvoiceV1WebhookIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceV1WebhookIdGetRequest struct via the builder pattern


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


## InvoiceV1WebhookPost

> WebHook InvoiceV1WebhookPost(ctx).WebHook(webHook).Execute()

Add a webhook



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	invoicesdk "github.com/invoicetronic/invoice-go-sdk"
)

func main() {
	webHook := *invoicesdk.NewWebHook() // WebHook | 

	configuration := invoicesdk.NewConfiguration()
	apiClient := invoicesdk.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.InvoiceV1WebhookPost(context.Background()).WebHook(webHook).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.InvoiceV1WebhookPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceV1WebhookPost`: WebHook
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.InvoiceV1WebhookPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceV1WebhookPostRequest struct via the builder pattern


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


## InvoiceV1WebhookPut

> WebHook InvoiceV1WebhookPut(ctx).WebHook(webHook).Execute()

Update a webhook



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	invoicesdk "github.com/invoicetronic/invoice-go-sdk"
)

func main() {
	webHook := *invoicesdk.NewWebHook() // WebHook | 

	configuration := invoicesdk.NewConfiguration()
	apiClient := invoicesdk.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.InvoiceV1WebhookPut(context.Background()).WebHook(webHook).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.InvoiceV1WebhookPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceV1WebhookPut`: WebHook
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.InvoiceV1WebhookPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceV1WebhookPutRequest struct via the builder pattern


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


## InvoiceV1WebhookhistoryGet

> []WebHookHistory InvoiceV1WebhookhistoryGet(ctx).Page(page).PageSize(pageSize).Execute()

List webhook history items

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	invoicesdk "github.com/invoicetronic/invoice-go-sdk"
)

func main() {
	page := int32(56) // int32 | Page number. (optional) (default to 1)
	pageSize := int32(56) // int32 | Items per page. (optional) (default to 100)

	configuration := invoicesdk.NewConfiguration()
	apiClient := invoicesdk.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.InvoiceV1WebhookhistoryGet(context.Background()).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.InvoiceV1WebhookhistoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceV1WebhookhistoryGet`: []WebHookHistory
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.InvoiceV1WebhookhistoryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceV1WebhookhistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number. | [default to 1]
 **pageSize** | **int32** | Items per page. | [default to 100]

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


## InvoiceV1WebhookhistoryIdGet

> WebHookHistory InvoiceV1WebhookhistoryIdGet(ctx, id).Execute()

Get a webhook history item by id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	invoicesdk "github.com/invoicetronic/invoice-go-sdk"
)

func main() {
	id := int32(56) // int32 | Item id.

	configuration := invoicesdk.NewConfiguration()
	apiClient := invoicesdk.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.InvoiceV1WebhookhistoryIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.InvoiceV1WebhookhistoryIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceV1WebhookhistoryIdGet`: WebHookHistory
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.InvoiceV1WebhookhistoryIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceV1WebhookhistoryIdGetRequest struct via the builder pattern


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

