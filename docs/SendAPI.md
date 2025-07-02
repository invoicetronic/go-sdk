# \SendAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SendFilePost**](SendAPI.md#SendFilePost) | **Post** /send/file | Add an invoice by file
[**SendGet**](SendAPI.md#SendGet) | **Get** /send | List invoices
[**SendIdGet**](SendAPI.md#SendIdGet) | **Get** /send/{id} | Get a invoice by id
[**SendJsonPost**](SendAPI.md#SendJsonPost) | **Post** /send/json | Add an invoice by json
[**SendPost**](SendAPI.md#SendPost) | **Post** /send | Add an invoice
[**SendValidateFilePost**](SendAPI.md#SendValidateFilePost) | **Post** /send/validate/file | Validate an invoice file
[**SendValidateJsonPost**](SendAPI.md#SendValidateJsonPost) | **Post** /send/validate/json | Validate an invoice by json
[**SendValidatePost**](SendAPI.md#SendValidatePost) | **Post** /send/validate | Validate an invoice
[**SendValidateXmlPost**](SendAPI.md#SendValidateXmlPost) | **Post** /send/validate/xml | Validate an invoice by xml
[**SendXmlPost**](SendAPI.md#SendXmlPost) | **Post** /send/xml | Add an invoice by xml



## SendFilePost

> Send SendFilePost(ctx).File(file).Validate(validate).Signature(signature).Execute()

Add an invoice by file



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
	file := os.NewFile(1234, "some_file") // *os.File | 
	validate := true // bool | Validate the document first, and reject it on failure. (optional) (default to false)
	signature := "signature_example" // string | Whether to digitally sign the document. (optional) (default to "Auto")

	configuration := invoicetronicsdk.NewConfiguration()
	apiClient := invoicetronicsdk.NewAPIClient(configuration)
	resp, r, err := apiClient.SendAPI.SendFilePost(context.Background()).File(file).Validate(validate).Signature(signature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendAPI.SendFilePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendFilePost`: Send
	fmt.Fprintf(os.Stdout, "Response from `SendAPI.SendFilePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendFilePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 
 **validate** | **bool** | Validate the document first, and reject it on failure. | [default to false]
 **signature** | **string** | Whether to digitally sign the document. | [default to &quot;Auto&quot;]

### Return type

[**Send**](Send.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendGet

> []Send SendGet(ctx).CompanyId(companyId).Identifier(identifier).Committente(committente).Prestatore(prestatore).FileName(fileName).LastUpdateFrom(lastUpdateFrom).LastUpdateTo(lastUpdateTo).DateSentFrom(dateSentFrom).DateSentTo(dateSentTo).DocumentDateFrom(documentDateFrom).DocumentDateTo(documentDateTo).DocumentNumber(documentNumber).IncludePayload(includePayload).Page(page).PageSize(pageSize).Sort(sort).Execute()

List invoices



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
	includePayload := true // bool | Include payload in the response. Defaults to false. (optional)
	page := int32(56) // int32 | Page number. (optional) (default to 1)
	pageSize := int32(56) // int32 | Items per page. Cannot be greater than 200. (optional) (default to 100)
	sort := "sort_example" // string | Sort by field. Prefix with '-' for descending order. (optional)

	configuration := invoicetronicsdk.NewConfiguration()
	apiClient := invoicetronicsdk.NewAPIClient(configuration)
	resp, r, err := apiClient.SendAPI.SendGet(context.Background()).CompanyId(companyId).Identifier(identifier).Committente(committente).Prestatore(prestatore).FileName(fileName).LastUpdateFrom(lastUpdateFrom).LastUpdateTo(lastUpdateTo).DateSentFrom(dateSentFrom).DateSentTo(dateSentTo).DocumentDateFrom(documentDateFrom).DocumentDateTo(documentDateTo).DocumentNumber(documentNumber).IncludePayload(includePayload).Page(page).PageSize(pageSize).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendAPI.SendGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendGet`: []Send
	fmt.Fprintf(os.Stdout, "Response from `SendAPI.SendGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | Company id | 
 **identifier** | **string** | SDI identifier. | 
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
 **includePayload** | **bool** | Include payload in the response. Defaults to false. | 
 **page** | **int32** | Page number. | [default to 1]
 **pageSize** | **int32** | Items per page. Cannot be greater than 200. | [default to 100]
 **sort** | **string** | Sort by field. Prefix with &#39;-&#39; for descending order. | 

### Return type

[**[]Send**](Send.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendIdGet

> Send SendIdGet(ctx, id).Execute()

Get a invoice by id



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
	resp, r, err := apiClient.SendAPI.SendIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendAPI.SendIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendIdGet`: Send
	fmt.Fprintf(os.Stdout, "Response from `SendAPI.SendIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Item id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Send**](Send.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendJsonPost

> Send SendJsonPost(ctx).FatturaOrdinaria(fatturaOrdinaria).Validate(validate).Signature(signature).Execute()

Add an invoice by json



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
	fatturaOrdinaria := *invoicetronicsdk.NewFatturaOrdinaria() // FatturaOrdinaria | 
	validate := true // bool | Validate the document first, and reject it on failure. (optional) (default to false)
	signature := "signature_example" // string | Whether to digitally sign the document. (optional) (default to "Auto")

	configuration := invoicetronicsdk.NewConfiguration()
	apiClient := invoicetronicsdk.NewAPIClient(configuration)
	resp, r, err := apiClient.SendAPI.SendJsonPost(context.Background()).FatturaOrdinaria(fatturaOrdinaria).Validate(validate).Signature(signature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendAPI.SendJsonPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendJsonPost`: Send
	fmt.Fprintf(os.Stdout, "Response from `SendAPI.SendJsonPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendJsonPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fatturaOrdinaria** | [**FatturaOrdinaria**](FatturaOrdinaria.md) |  | 
 **validate** | **bool** | Validate the document first, and reject it on failure. | [default to false]
 **signature** | **string** | Whether to digitally sign the document. | [default to &quot;Auto&quot;]

### Return type

[**Send**](Send.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendPost

> Send SendPost(ctx).Send(send).Validate(validate).Signature(signature).Execute()

Add an invoice



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
	send := *invoicetronicsdk.NewSend() // Send | 
	validate := true // bool | Validate the document first, and reject it on failure. (optional) (default to false)
	signature := "signature_example" // string | Whether to digitally sign the document. (optional) (default to "Auto")

	configuration := invoicetronicsdk.NewConfiguration()
	apiClient := invoicetronicsdk.NewAPIClient(configuration)
	resp, r, err := apiClient.SendAPI.SendPost(context.Background()).Send(send).Validate(validate).Signature(signature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendAPI.SendPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendPost`: Send
	fmt.Fprintf(os.Stdout, "Response from `SendAPI.SendPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **send** | [**Send**](Send.md) |  | 
 **validate** | **bool** | Validate the document first, and reject it on failure. | [default to false]
 **signature** | **string** | Whether to digitally sign the document. | [default to &quot;Auto&quot;]

### Return type

[**Send**](Send.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendValidateFilePost

> SendValidateFilePost(ctx).File(file).Execute()

Validate an invoice file



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
	file := os.NewFile(1234, "some_file") // *os.File | 

	configuration := invoicetronicsdk.NewConfiguration()
	apiClient := invoicetronicsdk.NewAPIClient(configuration)
	r, err := apiClient.SendAPI.SendValidateFilePost(context.Background()).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendAPI.SendValidateFilePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendValidateFilePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendValidateJsonPost

> SendValidateJsonPost(ctx).FatturaOrdinaria(fatturaOrdinaria).Execute()

Validate an invoice by json



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
	fatturaOrdinaria := *invoicetronicsdk.NewFatturaOrdinaria() // FatturaOrdinaria | 

	configuration := invoicetronicsdk.NewConfiguration()
	apiClient := invoicetronicsdk.NewAPIClient(configuration)
	r, err := apiClient.SendAPI.SendValidateJsonPost(context.Background()).FatturaOrdinaria(fatturaOrdinaria).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendAPI.SendValidateJsonPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendValidateJsonPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fatturaOrdinaria** | [**FatturaOrdinaria**](FatturaOrdinaria.md) |  | 

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendValidatePost

> SendValidatePost(ctx).Send(send).Execute()

Validate an invoice



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
	send := *invoicetronicsdk.NewSend() // Send | 

	configuration := invoicetronicsdk.NewConfiguration()
	apiClient := invoicetronicsdk.NewAPIClient(configuration)
	r, err := apiClient.SendAPI.SendValidatePost(context.Background()).Send(send).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendAPI.SendValidatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendValidatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **send** | [**Send**](Send.md) |  | 

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendValidateXmlPost

> SendValidateXmlPost(ctx).FatturaOrdinaria(fatturaOrdinaria).Execute()

Validate an invoice by xml



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
	fatturaOrdinaria := *invoicetronicsdk.NewFatturaOrdinaria() // FatturaOrdinaria | 

	configuration := invoicetronicsdk.NewConfiguration()
	apiClient := invoicetronicsdk.NewAPIClient(configuration)
	r, err := apiClient.SendAPI.SendValidateXmlPost(context.Background()).FatturaOrdinaria(fatturaOrdinaria).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendAPI.SendValidateXmlPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendValidateXmlPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fatturaOrdinaria** | [**FatturaOrdinaria**](FatturaOrdinaria.md) |  | 

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/xml
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendXmlPost

> Send SendXmlPost(ctx).FatturaOrdinaria(fatturaOrdinaria).Validate(validate).Signature(signature).Execute()

Add an invoice by xml



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
	fatturaOrdinaria := *invoicetronicsdk.NewFatturaOrdinaria() // FatturaOrdinaria | 
	validate := true // bool | Validate the document first, and reject it on failure. (optional) (default to false)
	signature := "signature_example" // string | Whether to digitally sign the document. (optional) (default to "Auto")

	configuration := invoicetronicsdk.NewConfiguration()
	apiClient := invoicetronicsdk.NewAPIClient(configuration)
	resp, r, err := apiClient.SendAPI.SendXmlPost(context.Background()).FatturaOrdinaria(fatturaOrdinaria).Validate(validate).Signature(signature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendAPI.SendXmlPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendXmlPost`: Send
	fmt.Fprintf(os.Stdout, "Response from `SendAPI.SendXmlPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendXmlPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fatturaOrdinaria** | [**FatturaOrdinaria**](FatturaOrdinaria.md) |  | 
 **validate** | **bool** | Validate the document first, and reject it on failure. | [default to false]
 **signature** | **string** | Whether to digitally sign the document. | [default to &quot;Auto&quot;]

### Return type

[**Send**](Send.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/xml
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

