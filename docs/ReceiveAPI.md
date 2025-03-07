# \ReceiveAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReceiveGet**](ReceiveAPI.md#ReceiveGet) | **Get** /receive | List incoming invoices
[**ReceiveIdDelete**](ReceiveAPI.md#ReceiveIdDelete) | **Delete** /receive/{id} | Delete an incoming invoice by id
[**ReceiveIdGet**](ReceiveAPI.md#ReceiveIdGet) | **Get** /receive/{id} | Get an incoming invoice by id



## ReceiveGet

> []Receive ReceiveGet(ctx).CompanyId(companyId).Identifier(identifier).Unread(unread).Committente(committente).Prestatore(prestatore).FileName(fileName).LastUpdateFrom(lastUpdateFrom).LastUpdateTo(lastUpdateTo).DateSentFrom(dateSentFrom).DateSentTo(dateSentTo).DocumentDateFrom(documentDateFrom).DocumentDateTo(documentDateTo).DocumentNumber(documentNumber).Page(page).PageSize(pageSize).Sort(sort).Execute()

List incoming invoices



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	invoicetronicsdk "github.com/invoicetronic/invoice-go-sdk"
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
	sort := "sort_example" // string | Sort by field. Prefix with '-' for descending order. (optional)

	configuration := invoicetronicsdk.NewConfiguration()
	apiClient := invoicetronicsdk.NewAPIClient(configuration)
	resp, r, err := apiClient.ReceiveAPI.ReceiveGet(context.Background()).CompanyId(companyId).Identifier(identifier).Unread(unread).Committente(committente).Prestatore(prestatore).FileName(fileName).LastUpdateFrom(lastUpdateFrom).LastUpdateTo(lastUpdateTo).DateSentFrom(dateSentFrom).DateSentTo(dateSentTo).DocumentDateFrom(documentDateFrom).DocumentDateTo(documentDateTo).DocumentNumber(documentNumber).Page(page).PageSize(pageSize).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReceiveAPI.ReceiveGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReceiveGet`: []Receive
	fmt.Fprintf(os.Stdout, "Response from `ReceiveAPI.ReceiveGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReceiveGetRequest struct via the builder pattern


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
 **sort** | **string** | Sort by field. Prefix with &#39;-&#39; for descending order. | 

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


## ReceiveIdDelete

> Receive ReceiveIdDelete(ctx, id).Execute()

Delete an incoming invoice by id



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	invoicetronicsdk "github.com/invoicetronic/invoice-go-sdk"
)

func main() {
	id := int32(56) // int32 | Item id

	configuration := invoicetronicsdk.NewConfiguration()
	apiClient := invoicetronicsdk.NewAPIClient(configuration)
	resp, r, err := apiClient.ReceiveAPI.ReceiveIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReceiveAPI.ReceiveIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReceiveIdDelete`: Receive
	fmt.Fprintf(os.Stdout, "Response from `ReceiveAPI.ReceiveIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Item id | 

### Other Parameters

Other parameters are passed through a pointer to a apiReceiveIdDeleteRequest struct via the builder pattern


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


## ReceiveIdGet

> Receive ReceiveIdGet(ctx, id).Execute()

Get an incoming invoice by id



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	invoicetronicsdk "github.com/invoicetronic/invoice-go-sdk"
)

func main() {
	id := int32(56) // int32 | Item id

	configuration := invoicetronicsdk.NewConfiguration()
	apiClient := invoicetronicsdk.NewAPIClient(configuration)
	resp, r, err := apiClient.ReceiveAPI.ReceiveIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReceiveAPI.ReceiveIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReceiveIdGet`: Receive
	fmt.Fprintf(os.Stdout, "Response from `ReceiveAPI.ReceiveIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Item id | 

### Other Parameters

Other parameters are passed through a pointer to a apiReceiveIdGetRequest struct via the builder pattern


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

