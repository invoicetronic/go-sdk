# \ReceiveAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InvoiceV1ReceiveGet**](ReceiveAPI.md#InvoiceV1ReceiveGet) | **Get** /invoice/v1/receive | List incoming invoices
[**InvoiceV1ReceiveIdDelete**](ReceiveAPI.md#InvoiceV1ReceiveIdDelete) | **Delete** /invoice/v1/receive/{id} | Delete an incoming invoice by id
[**InvoiceV1ReceiveIdGet**](ReceiveAPI.md#InvoiceV1ReceiveIdGet) | **Get** /invoice/v1/receive/{id} | Get an incoming invoice by id



## InvoiceV1ReceiveGet

> []Receive InvoiceV1ReceiveGet(ctx).CompanyId(companyId).Identifier(identifier).Unread(unread).Committente(committente).Prestatore(prestatore).FileName(fileName).LastUpdateFrom(lastUpdateFrom).LastUpdateTo(lastUpdateTo).DateSentFrom(dateSentFrom).DateSentTo(dateSentTo).DocumentDateFrom(documentDateFrom).DocumentDateTo(documentDateTo).DocumentNumber(documentNumber).Page(page).PageSize(pageSize).Execute()

List incoming invoices



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
	identifier := "identifier_example" // string | SDI identifier. (optional)
	unread := true // bool | Unread items only. (optional)
	committente := "committente_example" // string | Vat number or fiscal code. (optional)
	prestatore := "prestatore_example" // string | Vat number or fiscal code. (optional)
	fileName := "fileName_example" // string | File name. (optional)
	lastUpdateFrom := time.Now() // time.Time | UTC ISO 8601 (2024-11-29T12:34:56Z) (optional)
	lastUpdateTo := time.Now() // time.Time | UTC ISO 8601 (2024-11-29T12:34:56Z) (optional)
	dateSentFrom := time.Now() // time.Time | UTC ISO 8601 (2024-11-29T12:34:56Z) (optional)
	dateSentTo := time.Now() // time.Time | UTC ISO 8601 (2024-11-29T12:34:56Z) (optional)
	documentDateFrom := time.Now() // time.Time | UTC ISO 8601 (2024-11-29T12:34:56Z) (optional)
	documentDateTo := time.Now() // time.Time | UTC ISO 8601 (2024-11-29T12:34:56Z) (optional)
	documentNumber := "documentNumber_example" // string | Document number. (optional)
	page := int32(56) // int32 | Page number. Defaults to 1. (optional) (default to 1)
	pageSize := int32(56) // int32 | Items per page. Defaults to 50. Cannot be greater than 200. (optional) (default to 100)

	configuration := invoicesdk.NewConfiguration()
	apiClient := invoicesdk.NewAPIClient(configuration)
	resp, r, err := apiClient.ReceiveAPI.InvoiceV1ReceiveGet(context.Background()).CompanyId(companyId).Identifier(identifier).Unread(unread).Committente(committente).Prestatore(prestatore).FileName(fileName).LastUpdateFrom(lastUpdateFrom).LastUpdateTo(lastUpdateTo).DateSentFrom(dateSentFrom).DateSentTo(dateSentTo).DocumentDateFrom(documentDateFrom).DocumentDateTo(documentDateTo).DocumentNumber(documentNumber).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReceiveAPI.InvoiceV1ReceiveGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceV1ReceiveGet`: []Receive
	fmt.Fprintf(os.Stdout, "Response from `ReceiveAPI.InvoiceV1ReceiveGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceV1ReceiveGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | Company id | 
 **identifier** | **string** | SDI identifier. | 
 **unread** | **bool** | Unread items only. | 
 **committente** | **string** | Vat number or fiscal code. | 
 **prestatore** | **string** | Vat number or fiscal code. | 
 **fileName** | **string** | File name. | 
 **lastUpdateFrom** | **time.Time** | UTC ISO 8601 (2024-11-29T12:34:56Z) | 
 **lastUpdateTo** | **time.Time** | UTC ISO 8601 (2024-11-29T12:34:56Z) | 
 **dateSentFrom** | **time.Time** | UTC ISO 8601 (2024-11-29T12:34:56Z) | 
 **dateSentTo** | **time.Time** | UTC ISO 8601 (2024-11-29T12:34:56Z) | 
 **documentDateFrom** | **time.Time** | UTC ISO 8601 (2024-11-29T12:34:56Z) | 
 **documentDateTo** | **time.Time** | UTC ISO 8601 (2024-11-29T12:34:56Z) | 
 **documentNumber** | **string** | Document number. | 
 **page** | **int32** | Page number. Defaults to 1. | [default to 1]
 **pageSize** | **int32** | Items per page. Defaults to 50. Cannot be greater than 200. | [default to 100]

### Return type

[**[]Receive**](Receive.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoiceV1ReceiveIdDelete

> Receive InvoiceV1ReceiveIdDelete(ctx, id).Execute()

Delete an incoming invoice by id



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
	resp, r, err := apiClient.ReceiveAPI.InvoiceV1ReceiveIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReceiveAPI.InvoiceV1ReceiveIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceV1ReceiveIdDelete`: Receive
	fmt.Fprintf(os.Stdout, "Response from `ReceiveAPI.InvoiceV1ReceiveIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Item id | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceV1ReceiveIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Receive**](Receive.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoiceV1ReceiveIdGet

> Receive InvoiceV1ReceiveIdGet(ctx, id).Execute()

Get an incoming invoice by id



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
	resp, r, err := apiClient.ReceiveAPI.InvoiceV1ReceiveIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReceiveAPI.InvoiceV1ReceiveIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceV1ReceiveIdGet`: Receive
	fmt.Fprintf(os.Stdout, "Response from `ReceiveAPI.InvoiceV1ReceiveIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Item id | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceV1ReceiveIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Receive**](Receive.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

