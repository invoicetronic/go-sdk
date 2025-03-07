# \UpdateAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateGet**](UpdateAPI.md#UpdateGet) | **Get** /update | List updates
[**UpdateIdGet**](UpdateAPI.md#UpdateIdGet) | **Get** /update/{id} | Get an update by id



## UpdateGet

> []Update UpdateGet(ctx).CompanyId(companyId).Identifier(identifier).Unread(unread).SendId(sendId).State(state).LastUpdateFrom(lastUpdateFrom).LastUpdateTo(lastUpdateTo).DateSentFrom(dateSentFrom).DateSentTo(dateSentTo).Page(page).PageSize(pageSize).Sort(sort).Execute()

List updates



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
	companyId := int32(56) // int32 | Company id (optional)
	identifier := "identifier_example" // string | SDI identifier. (optional)
	unread := true // bool | Unread items only. (optional)
	sendId := int32(56) // int32 | Send item's id. (optional)
	state := "state_example" // string | SDI state (optional)
	lastUpdateFrom := time.Now() // time.Time | UTC ISO 8601 (2024-11-29T12:34:56Z) (optional)
	lastUpdateTo := time.Now() // time.Time | UTC ISO 8601 (2024-11-29T12:34:56Z) (optional)
	dateSentFrom := time.Now() // time.Time | UTC ISO 8601 (2024-11-29T12:34:56Z) (optional)
	dateSentTo := time.Now() // time.Time | UTC ISO 8601 (2024-11-29T12:34:56Z) (optional)
	page := int32(56) // int32 | Page number. Defaults to 1. (optional) (default to 1)
	pageSize := int32(56) // int32 | Items per page. Defaults to 50. Cannot be greater than 200. (optional) (default to 100)
	sort := "sort_example" // string | Sort by field. Prefix with '-' for descending order. (optional)

	configuration := invoicetronicsdk.NewConfiguration()
	apiClient := invoicetronicsdk.NewAPIClient(configuration)
	resp, r, err := apiClient.UpdateAPI.UpdateGet(context.Background()).CompanyId(companyId).Identifier(identifier).Unread(unread).SendId(sendId).State(state).LastUpdateFrom(lastUpdateFrom).LastUpdateTo(lastUpdateTo).DateSentFrom(dateSentFrom).DateSentTo(dateSentTo).Page(page).PageSize(pageSize).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpdateAPI.UpdateGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGet`: []Update
	fmt.Fprintf(os.Stdout, "Response from `UpdateAPI.UpdateGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | Company id | 
 **identifier** | **string** | SDI identifier. | 
 **unread** | **bool** | Unread items only. | 
 **sendId** | **int32** | Send item&#39;s id. | 
 **state** | **string** | SDI state | 
 **lastUpdateFrom** | **time.Time** | UTC ISO 8601 (2024-11-29T12:34:56Z) | 
 **lastUpdateTo** | **time.Time** | UTC ISO 8601 (2024-11-29T12:34:56Z) | 
 **dateSentFrom** | **time.Time** | UTC ISO 8601 (2024-11-29T12:34:56Z) | 
 **dateSentTo** | **time.Time** | UTC ISO 8601 (2024-11-29T12:34:56Z) | 
 **page** | **int32** | Page number. Defaults to 1. | [default to 1]
 **pageSize** | **int32** | Items per page. Defaults to 50. Cannot be greater than 200. | [default to 100]
 **sort** | **string** | Sort by field. Prefix with &#39;-&#39; for descending order. | 

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


## UpdateIdGet

> Update UpdateIdGet(ctx, id).Execute()

Get an update by id



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
	resp, r, err := apiClient.UpdateAPI.UpdateIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpdateAPI.UpdateIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIdGet`: Update
	fmt.Fprintf(os.Stdout, "Response from `UpdateAPI.UpdateIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Item id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIdGetRequest struct via the builder pattern


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

