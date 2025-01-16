# \LogAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InvoiceV1LogGet**](LogAPI.md#InvoiceV1LogGet) | **Get** /invoice/v1/log | List events
[**InvoiceV1LogIdGet**](LogAPI.md#InvoiceV1LogIdGet) | **Get** /invoice/v1/log/{id} | Get an event by id



## InvoiceV1LogGet

> []Event InvoiceV1LogGet(ctx).CompanyId(companyId).Endpoint(endpoint).Method(method).ApiVerion(apiVerion).StatusCode(statusCode).DateCreatedFrom(dateCreatedFrom).DateCreatedTo(dateCreatedTo).Page(page).PageSize(pageSize).Execute()

List events



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	invoicesdk "github.com/invoicetronic/invoice-go-sdk"
)

func main() {
	companyId := int32(56) // int32 | Company id (optional)
	endpoint := "endpoint_example" // string |  (optional)
	method := "method_example" // string |  (optional)
	apiVerion := int32(56) // int32 | Api version (optional)
	statusCode := int32(56) // int32 | Response status code (optional)
	dateCreatedFrom := time.Now() // time.Time | UTC ISO 8601 (2024-11-29T12:34:56Z) (optional)
	dateCreatedTo := time.Now() // time.Time | UTC ISO 8601 (2024-11-29T12:34:56Z) (optional)
	page := int32(56) // int32 | Page number. Defaults to 1. (optional) (default to 1)
	pageSize := int32(56) // int32 | Items per page. Defaults to 50. Cannot be greater than 200. (optional) (default to 100)

	configuration := invoicesdk.NewConfiguration()
	apiClient := invoicesdk.NewAPIClient(configuration)
	resp, r, err := apiClient.LogAPI.InvoiceV1LogGet(context.Background()).CompanyId(companyId).Endpoint(endpoint).Method(method).ApiVerion(apiVerion).StatusCode(statusCode).DateCreatedFrom(dateCreatedFrom).DateCreatedTo(dateCreatedTo).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogAPI.InvoiceV1LogGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceV1LogGet`: []Event
	fmt.Fprintf(os.Stdout, "Response from `LogAPI.InvoiceV1LogGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceV1LogGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | Company id | 
 **endpoint** | **string** |  | 
 **method** | **string** |  | 
 **apiVerion** | **int32** | Api version | 
 **statusCode** | **int32** | Response status code | 
 **dateCreatedFrom** | **time.Time** | UTC ISO 8601 (2024-11-29T12:34:56Z) | 
 **dateCreatedTo** | **time.Time** | UTC ISO 8601 (2024-11-29T12:34:56Z) | 
 **page** | **int32** | Page number. Defaults to 1. | [default to 1]
 **pageSize** | **int32** | Items per page. Defaults to 50. Cannot be greater than 200. | [default to 100]

### Return type

[**[]Event**](Event.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoiceV1LogIdGet

> Event InvoiceV1LogIdGet(ctx, id).Execute()

Get an event by id



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
	id := int32(56) // int32 | Item id

	configuration := invoicesdk.NewConfiguration()
	apiClient := invoicesdk.NewAPIClient(configuration)
	resp, r, err := apiClient.LogAPI.InvoiceV1LogIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogAPI.InvoiceV1LogIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceV1LogIdGet`: Event
	fmt.Fprintf(os.Stdout, "Response from `LogAPI.InvoiceV1LogIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Item id | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceV1LogIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Event**](Event.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

