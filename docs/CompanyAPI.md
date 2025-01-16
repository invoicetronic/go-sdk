# \CompanyAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InvoiceV1CompanyGet**](CompanyAPI.md#InvoiceV1CompanyGet) | **Get** /invoice/v1/company | List companies
[**InvoiceV1CompanyIdDelete**](CompanyAPI.md#InvoiceV1CompanyIdDelete) | **Delete** /invoice/v1/company/{id} | Delete a company
[**InvoiceV1CompanyIdGet**](CompanyAPI.md#InvoiceV1CompanyIdGet) | **Get** /invoice/v1/company/{id} | Get a company by id
[**InvoiceV1CompanyPost**](CompanyAPI.md#InvoiceV1CompanyPost) | **Post** /invoice/v1/company | Add a company
[**InvoiceV1CompanyPut**](CompanyAPI.md#InvoiceV1CompanyPut) | **Put** /invoice/v1/company | Update a company



## InvoiceV1CompanyGet

> []Company InvoiceV1CompanyGet(ctx).Page(page).PageSize(pageSize).Execute()

List companies



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
	page := int32(56) // int32 | Page number. Defaults to 1. (optional) (default to 1)
	pageSize := int32(56) // int32 | Items per page. Defaults to 50. Cannot be greater than 200. (optional) (default to 100)

	configuration := invoicesdk.NewConfiguration()
	apiClient := invoicesdk.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyAPI.InvoiceV1CompanyGet(context.Background()).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.InvoiceV1CompanyGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceV1CompanyGet`: []Company
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.InvoiceV1CompanyGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceV1CompanyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number. Defaults to 1. | [default to 1]
 **pageSize** | **int32** | Items per page. Defaults to 50. Cannot be greater than 200. | [default to 100]

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


## InvoiceV1CompanyIdDelete

> Company InvoiceV1CompanyIdDelete(ctx, id).Execute()

Delete a company



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
	resp, r, err := apiClient.CompanyAPI.InvoiceV1CompanyIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.InvoiceV1CompanyIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceV1CompanyIdDelete`: Company
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.InvoiceV1CompanyIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Item id | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceV1CompanyIdDeleteRequest struct via the builder pattern


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


## InvoiceV1CompanyIdGet

> Company InvoiceV1CompanyIdGet(ctx, id).Execute()

Get a company by id



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
	resp, r, err := apiClient.CompanyAPI.InvoiceV1CompanyIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.InvoiceV1CompanyIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceV1CompanyIdGet`: Company
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.InvoiceV1CompanyIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Item id | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceV1CompanyIdGetRequest struct via the builder pattern


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


## InvoiceV1CompanyPost

> Company InvoiceV1CompanyPost(ctx).Company(company).Execute()

Add a company



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
	company := *invoicesdk.NewCompany("IT01234567891", "RSSMRA70A01F205V", "Mario Rossi Srl") // Company | 

	configuration := invoicesdk.NewConfiguration()
	apiClient := invoicesdk.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyAPI.InvoiceV1CompanyPost(context.Background()).Company(company).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.InvoiceV1CompanyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceV1CompanyPost`: Company
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.InvoiceV1CompanyPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceV1CompanyPostRequest struct via the builder pattern


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


## InvoiceV1CompanyPut

> Company InvoiceV1CompanyPut(ctx).Company(company).Execute()

Update a company



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
	company := *invoicesdk.NewCompany("IT01234567891", "RSSMRA70A01F205V", "Mario Rossi Srl") // Company | 

	configuration := invoicesdk.NewConfiguration()
	apiClient := invoicesdk.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyAPI.InvoiceV1CompanyPut(context.Background()).Company(company).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.InvoiceV1CompanyPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceV1CompanyPut`: Company
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.InvoiceV1CompanyPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceV1CompanyPutRequest struct via the builder pattern


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

