# \LogAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LogGet**](LogAPI.md#LogGet) | **Get** /log | List events
[**LogIdGet**](LogAPI.md#LogIdGet) | **Get** /log/{id} | Get an event by id



## LogGet

> []Event LogGet(ctx).CompanyId(companyId).Endpoint(endpoint).Method(method).ApiVerion(apiVerion).StatusCode(statusCode).DateCreatedFrom(dateCreatedFrom).DateCreatedTo(dateCreatedTo).Page(page).PageSize(pageSize).Sort(sort).Query(query).Success(success).DateTimeFrom(dateTimeFrom).DateTimeTo(dateTimeTo).Execute()

List events



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
	endpoint := "endpoint_example" // string |  (optional)
	method := "method_example" // string |  (optional)
	apiVerion := int32(56) // int32 | Api version (optional)
	statusCode := int32(56) // int32 | Response status code (optional)
	dateCreatedFrom := time.Now() // time.Time | UTC ISO 8601 (2024-11-29T12:34:56Z) (optional)
	dateCreatedTo := time.Now() // time.Time | UTC ISO 8601 (2024-11-29T12:34:56Z) (optional)
	page := int32(56) // int32 | Page number. Defaults to 1. (optional) (default to 1)
	pageSize := int32(56) // int32 | Items per page. Defaults to 50. Cannot be greater than 200. (optional) (default to 100)
	sort := "sort_example" // string | Sort by field. Prefix with '-' for descending order. (optional)
	query := "query_example" // string |  (optional)
	success := true // bool |  (optional)
	dateTimeFrom := time.Now() // time.Time | Date and time of the event (optional)
	dateTimeTo := time.Now() // time.Time | Date and time of the event (optional)

	configuration := invoicetronicsdk.NewConfiguration()
	apiClient := invoicetronicsdk.NewAPIClient(configuration)
	resp, r, err := apiClient.LogAPI.LogGet(context.Background()).CompanyId(companyId).Endpoint(endpoint).Method(method).ApiVerion(apiVerion).StatusCode(statusCode).DateCreatedFrom(dateCreatedFrom).DateCreatedTo(dateCreatedTo).Page(page).PageSize(pageSize).Sort(sort).Query(query).Success(success).DateTimeFrom(dateTimeFrom).DateTimeTo(dateTimeTo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogAPI.LogGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LogGet`: []Event
	fmt.Fprintf(os.Stdout, "Response from `LogAPI.LogGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLogGetRequest struct via the builder pattern


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
 **sort** | **string** | Sort by field. Prefix with &#39;-&#39; for descending order. | 
 **query** | **string** |  | 
 **success** | **bool** |  | 
 **dateTimeFrom** | **time.Time** | Date and time of the event | 
 **dateTimeTo** | **time.Time** | Date and time of the event | 

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


## LogIdGet

> Event LogIdGet(ctx, id).Execute()

Get an event by id



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
	resp, r, err := apiClient.LogAPI.LogIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogAPI.LogIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LogIdGet`: Event
	fmt.Fprintf(os.Stdout, "Response from `LogAPI.LogIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Item id | 

### Other Parameters

Other parameters are passed through a pointer to a apiLogIdGetRequest struct via the builder pattern


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

