# \SubkeyAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubkeyGet**](SubkeyAPI.md#SubkeyGet) | **Get** /subkey | List restricted keys
[**SubkeyIdDelete**](SubkeyAPI.md#SubkeyIdDelete) | **Delete** /subkey/{id} | Delete a restricted key
[**SubkeyIdGet**](SubkeyAPI.md#SubkeyIdGet) | **Get** /subkey/{id} | Get a restricted key by id
[**SubkeyIdRollPost**](SubkeyAPI.md#SubkeyIdRollPost) | **Post** /subkey/{id}/roll | Roll the secrets of a restricted key
[**SubkeyPost**](SubkeyAPI.md#SubkeyPost) | **Post** /subkey | Add a restricted key
[**SubkeyPut**](SubkeyAPI.md#SubkeyPut) | **Put** /subkey | Update a restricted key



## SubkeyGet

> []SubKey SubkeyGet(ctx).Page(page).PageSize(pageSize).CompanyId(companyId).Active(active).Q(q).Execute()

List restricted keys



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
	companyId := int32(56) // int32 | Company id (optional)
	active := true // bool | Active keys only (true) or inactive only (false). (optional)
	q := "q_example" // string | Human-readable label: free-text search. (optional)

	configuration := invoicetronicsdk.NewConfiguration()
	apiClient := invoicetronicsdk.NewAPIClient(configuration)
	resp, r, err := apiClient.SubkeyAPI.SubkeyGet(context.Background()).Page(page).PageSize(pageSize).CompanyId(companyId).Active(active).Q(q).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubkeyAPI.SubkeyGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubkeyGet`: []SubKey
	fmt.Fprintf(os.Stdout, "Response from `SubkeyAPI.SubkeyGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubkeyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number. | [default to 1]
 **pageSize** | **int32** | Items per page. Cannot be greater than 200. | [default to 100]
 **companyId** | **int32** | Company id | 
 **active** | **bool** | Active keys only (true) or inactive only (false). | 
 **q** | **string** | Human-readable label: free-text search. | 

### Return type

[**[]SubKey**](SubKey.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubkeyIdDelete

> SubKey SubkeyIdDelete(ctx, id).Execute()

Delete a restricted key



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
	resp, r, err := apiClient.SubkeyAPI.SubkeyIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubkeyAPI.SubkeyIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubkeyIdDelete`: SubKey
	fmt.Fprintf(os.Stdout, "Response from `SubkeyAPI.SubkeyIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Item id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubkeyIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubKey**](SubKey.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubkeyIdGet

> SubKey SubkeyIdGet(ctx, id).Execute()

Get a restricted key by id



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
	resp, r, err := apiClient.SubkeyAPI.SubkeyIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubkeyAPI.SubkeyIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubkeyIdGet`: SubKey
	fmt.Fprintf(os.Stdout, "Response from `SubkeyAPI.SubkeyIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Item id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubkeyIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubKey**](SubKey.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubkeyIdRollPost

> SubKeyWithSecrets SubkeyIdRollPost(ctx, id).ExpiresInHours(expiresInHours).Execute()

Roll the secrets of a restricted key



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
	expiresInHours := int32(56) // int32 | Hours the replaced secrets keep working, from 1 to 168. When omitted, they stop working at once. (optional)

	configuration := invoicetronicsdk.NewConfiguration()
	apiClient := invoicetronicsdk.NewAPIClient(configuration)
	resp, r, err := apiClient.SubkeyAPI.SubkeyIdRollPost(context.Background(), id).ExpiresInHours(expiresInHours).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubkeyAPI.SubkeyIdRollPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubkeyIdRollPost`: SubKeyWithSecrets
	fmt.Fprintf(os.Stdout, "Response from `SubkeyAPI.SubkeyIdRollPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Item id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubkeyIdRollPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expiresInHours** | **int32** | Hours the replaced secrets keep working, from 1 to 168. When omitted, they stop working at once. | 

### Return type

[**SubKeyWithSecrets**](SubKeyWithSecrets.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubkeyPost

> SubKeyWithSecrets SubkeyPost(ctx).SubKeyRequest(subKeyRequest).Execute()

Add a restricted key



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
	subKeyRequest := *invoicetronicsdk.NewSubKeyRequest("Studio Rossi Srl") // SubKeyRequest | 

	configuration := invoicetronicsdk.NewConfiguration()
	apiClient := invoicetronicsdk.NewAPIClient(configuration)
	resp, r, err := apiClient.SubkeyAPI.SubkeyPost(context.Background()).SubKeyRequest(subKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubkeyAPI.SubkeyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubkeyPost`: SubKeyWithSecrets
	fmt.Fprintf(os.Stdout, "Response from `SubkeyAPI.SubkeyPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubkeyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subKeyRequest** | [**SubKeyRequest**](SubKeyRequest.md) |  | 

### Return type

[**SubKeyWithSecrets**](SubKeyWithSecrets.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubkeyPut

> SubKey SubkeyPut(ctx).SubKeyUpdate(subKeyUpdate).Execute()

Update a restricted key



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
	subKeyUpdate := *invoicetronicsdk.NewSubKeyUpdate("Studio Rossi Srl") // SubKeyUpdate | 

	configuration := invoicetronicsdk.NewConfiguration()
	apiClient := invoicetronicsdk.NewAPIClient(configuration)
	resp, r, err := apiClient.SubkeyAPI.SubkeyPut(context.Background()).SubKeyUpdate(subKeyUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubkeyAPI.SubkeyPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubkeyPut`: SubKey
	fmt.Fprintf(os.Stdout, "Response from `SubkeyAPI.SubkeyPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubkeyPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subKeyUpdate** | [**SubKeyUpdate**](SubKeyUpdate.md) |  | 

### Return type

[**SubKey**](SubKey.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

