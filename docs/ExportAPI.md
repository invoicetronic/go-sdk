# \ExportAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportGet**](ExportAPI.md#ExportGet) | **Get** /export | Export invoices as a ZIP archive



## ExportGet

> ExportGet(ctx).Type_(type_).CompanyId(companyId).Year(year).Month(month).Quarter(quarter).DocumentDateFrom(documentDateFrom).DocumentDateTo(documentDateTo).Execute()

Export invoices as a ZIP archive



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	invoicetronicsdk "github.com/invoicetronic/go-sdk"
)

func main() {
	type_ := "type__example" // string |  (optional) (default to "Both")
	companyId := int32(56) // int32 | Company id (optional)
	year := int32(56) // int32 |  (optional)
	month := int32(56) // int32 |  (optional)
	quarter := int32(56) // int32 |  (optional)
	documentDateFrom := time.Now() // time.Time | UTC ISO 8601 (2024-11-29T12:34:56Z) (optional)
	documentDateTo := time.Now() // time.Time | UTC ISO 8601 (2024-11-29T12:34:56Z) (optional)

	configuration := invoicetronicsdk.NewConfiguration()
	apiClient := invoicetronicsdk.NewAPIClient(configuration)
	r, err := apiClient.ExportAPI.ExportGet(context.Background()).Type_(type_).CompanyId(companyId).Year(year).Month(month).Quarter(quarter).DocumentDateFrom(documentDateFrom).DocumentDateTo(documentDateTo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.ExportGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** |  | [default to &quot;Both&quot;]
 **companyId** | **int32** | Company id | 
 **year** | **int32** |  | 
 **month** | **int32** |  | 
 **quarter** | **int32** |  | 
 **documentDateFrom** | **time.Time** | UTC ISO 8601 (2024-11-29T12:34:56Z) | 
 **documentDateTo** | **time.Time** | UTC ISO 8601 (2024-11-29T12:34:56Z) | 

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

