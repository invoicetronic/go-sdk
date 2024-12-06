# \UpdateAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InvoiceV1UpdateGet**](UpdateAPI.md#InvoiceV1UpdateGet) | **Get** /invoice/v1/update | List updates
[**InvoiceV1UpdateIdGet**](UpdateAPI.md#InvoiceV1UpdateIdGet) | **Get** /invoice/v1/update/{id} | Get an update by id



## InvoiceV1UpdateGet

> []Update InvoiceV1UpdateGet(ctx).CompanyId(companyId).Identifier(identifier).Unread(unread).SendId(sendId).State(state).LastUpdateFrom(lastUpdateFrom).LastUpdateTo(lastUpdateTo).DateSentFrom(dateSentFrom).DateSentTo(dateSentTo).Page(page).PageSize(pageSize).Execute()

List updates



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
	companyId := int32(56) // int32 | Company id. (optional)
	identifier := "identifier_example" // string | SDI identifier. (optional)
	unread := true // bool | Only unread items. (optional)
	sendId := int32(56) // int32 | Send item's id. (optional)
	state := "state_example" // string | SDI state (optional)
	lastUpdateFrom := time.Now() // time.Time | UTC ISO 8601 (2024-11-29T12:34:56Z) (optional)
	lastUpdateTo := time.Now() // time.Time | UTC ISO 8601 (2024-11-29T12:34:56Z) (optional)
	dateSentFrom := time.Now() // time.Time | UTC ISO 8601 (2024-11-29T12:34:56Z) (optional)
	dateSentTo := time.Now() // time.Time | UTC ISO 8601 format (2024-11-29T12:34:56Z) (optional)
	page := int32(56) // int32 | Page number. (optional) (default to 1)
	pageSize := int32(56) // int32 | Items per page. (optional) (default to 100)

	configuration := invoicesdk.NewConfiguration()
	apiClient := invoicesdk.NewAPIClient(configuration)
	resp, r, err := apiClient.UpdateAPI.InvoiceV1UpdateGet(context.Background()).CompanyId(companyId).Identifier(identifier).Unread(unread).SendId(sendId).State(state).LastUpdateFrom(lastUpdateFrom).LastUpdateTo(lastUpdateTo).DateSentFrom(dateSentFrom).DateSentTo(dateSentTo).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpdateAPI.InvoiceV1UpdateGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceV1UpdateGet`: []Update
	fmt.Fprintf(os.Stdout, "Response from `UpdateAPI.InvoiceV1UpdateGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceV1UpdateGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | Company id. | 
 **identifier** | **string** | SDI identifier. | 
 **unread** | **bool** | Only unread items. | 
 **sendId** | **int32** | Send item&#39;s id. | 
 **state** | **string** | SDI state | 
 **lastUpdateFrom** | **time.Time** | UTC ISO 8601 (2024-11-29T12:34:56Z) | 
 **lastUpdateTo** | **time.Time** | UTC ISO 8601 (2024-11-29T12:34:56Z) | 
 **dateSentFrom** | **time.Time** | UTC ISO 8601 (2024-11-29T12:34:56Z) | 
 **dateSentTo** | **time.Time** | UTC ISO 8601 format (2024-11-29T12:34:56Z) | 
 **page** | **int32** | Page number. | [default to 1]
 **pageSize** | **int32** | Items per page. | [default to 100]

### Return type

[**[]Update**](Update.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoiceV1UpdateIdGet

> Update InvoiceV1UpdateIdGet(ctx, id).Execute()

Get an update by id



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
	resp, r, err := apiClient.UpdateAPI.InvoiceV1UpdateIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpdateAPI.InvoiceV1UpdateIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceV1UpdateIdGet`: Update
	fmt.Fprintf(os.Stdout, "Response from `UpdateAPI.InvoiceV1UpdateIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceV1UpdateIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Update**](Update.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

