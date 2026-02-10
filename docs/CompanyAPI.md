# \CompanyAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompanyGet**](CompanyAPI.md#CompanyGet) | **Get** /company | List companies
[**CompanyIdDelete**](CompanyAPI.md#CompanyIdDelete) | **Delete** /company/{id} | Delete a company
[**CompanyIdGet**](CompanyAPI.md#CompanyIdGet) | **Get** /company/{id} | Get a company by id
[**CompanyPost**](CompanyAPI.md#CompanyPost) | **Post** /company | Add a company
[**CompanyPut**](CompanyAPI.md#CompanyPut) | **Put** /company | Update a company



## CompanyGet

> []Company CompanyGet(ctx).Page(page).PageSize(pageSize).Sort(sort).Execute()

List companies



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
	page := int32(56) // int32 | Page number. (optional) (default to 1)
	pageSize := int32(56) // int32 | Items per page. Cannot be greater than 200. (optional) (default to 100)
	sort := "sort_example" // string | Sort by field. Prefix with '-' for descending order. (optional)

	configuration := invoicetronicsdk.NewConfiguration()
	apiClient := invoicetronicsdk.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyAPI.CompanyGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.CompanyGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompanyGet`: []Company
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.CompanyGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCompanyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number. | [default to 1]
 **pageSize** | **int32** | Items per page. Cannot be greater than 200. | [default to 100]
 **sort** | **string** | Sort by field. Prefix with &#39;-&#39; for descending order. | 

### Return type

[**[]Company**](Company.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompanyIdDelete

> Company CompanyIdDelete(ctx, id).Force(force).Execute()

Delete a company



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
	force := true // bool | Force delete including all related data. (optional) (default to false)

	configuration := invoicetronicsdk.NewConfiguration()
	apiClient := invoicetronicsdk.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyAPI.CompanyIdDelete(context.Background(), id).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.CompanyIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompanyIdDelete`: Company
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.CompanyIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Item id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompanyIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | Force delete including all related data. | [default to false]

### Return type

[**Company**](Company.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompanyIdGet

> Company CompanyIdGet(ctx, id).Execute()

Get a company by id



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
	resp, r, err := apiClient.CompanyAPI.CompanyIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.CompanyIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompanyIdGet`: Company
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.CompanyIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Item id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompanyIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Company**](Company.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompanyPost

> Company CompanyPost(ctx).Company(company).Execute()

Add a company



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
	company := *invoicetronicsdk.NewCompany("IT01234567891", "RSSMRA70A01F205V", "Mario Rossi Srl") // Company | 

	configuration := invoicetronicsdk.NewConfiguration()
	apiClient := invoicetronicsdk.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyAPI.CompanyPost(context.Background()).Company(company).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.CompanyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompanyPost`: Company
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.CompanyPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCompanyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **company** | [**Company**](Company.md) |  | 

### Return type

[**Company**](Company.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompanyPut

> Company CompanyPut(ctx).Company(company).Execute()

Update a company



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
	company := *invoicetronicsdk.NewCompany("IT01234567891", "RSSMRA70A01F205V", "Mario Rossi Srl") // Company | 

	configuration := invoicetronicsdk.NewConfiguration()
	apiClient := invoicetronicsdk.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyAPI.CompanyPut(context.Background()).Company(company).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.CompanyPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompanyPut`: Company
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.CompanyPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCompanyPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **company** | [**Company**](Company.md) |  | 

### Return type

[**Company**](Company.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

