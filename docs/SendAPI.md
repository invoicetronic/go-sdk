# \SendAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InvoiceV1SendFilesPost**](SendAPI.md#InvoiceV1SendFilesPost) | **Post** /invoice/v1/send/files | Add an invoice by file
[**InvoiceV1SendGet**](SendAPI.md#InvoiceV1SendGet) | **Get** /invoice/v1/send | List invoices
[**InvoiceV1SendIdGet**](SendAPI.md#InvoiceV1SendIdGet) | **Get** /invoice/v1/send/{id} | Get a invoice by id
[**InvoiceV1SendJsonPost**](SendAPI.md#InvoiceV1SendJsonPost) | **Post** /invoice/v1/send/json | Add an invoice by json
[**InvoiceV1SendPost**](SendAPI.md#InvoiceV1SendPost) | **Post** /invoice/v1/send | Add an invoice
[**InvoiceV1SendValidateFilesPost**](SendAPI.md#InvoiceV1SendValidateFilesPost) | **Post** /invoice/v1/send/validate/files | Validate an invoice by file
[**InvoiceV1SendValidateJsonPost**](SendAPI.md#InvoiceV1SendValidateJsonPost) | **Post** /invoice/v1/send/validate/json | Validate an invoice by json
[**InvoiceV1SendValidatePost**](SendAPI.md#InvoiceV1SendValidatePost) | **Post** /invoice/v1/send/validate | Validate an invoice
[**InvoiceV1SendValidateXmlPost**](SendAPI.md#InvoiceV1SendValidateXmlPost) | **Post** /invoice/v1/send/validate/xml | Validate an invoice by xml
[**InvoiceV1SendXmlPost**](SendAPI.md#InvoiceV1SendXmlPost) | **Post** /invoice/v1/send/xml | Add an invoice by xml



## InvoiceV1SendFilesPost

> Send InvoiceV1SendFilesPost(ctx).Files(files).Validate(validate).Signature(signature).Execute()

Add an invoice by file



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
	files := []*os.File{"TODO"} // []*os.File | 
	validate := true // bool | Validate the document first, and reject it on failure. (optional) (default to false)
	signature := "signature_example" // string | Whether to digitally sign the document. (optional) (default to "Auto")

	configuration := invoicesdk.NewConfiguration()
	apiClient := invoicesdk.NewAPIClient(configuration)
	resp, r, err := apiClient.SendAPI.InvoiceV1SendFilesPost(context.Background()).Files(files).Validate(validate).Signature(signature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendAPI.InvoiceV1SendFilesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceV1SendFilesPost`: Send
	fmt.Fprintf(os.Stdout, "Response from `SendAPI.InvoiceV1SendFilesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceV1SendFilesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **files** | **[]*os.File** |  | 
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


## InvoiceV1SendGet

> []Send InvoiceV1SendGet(ctx).CompanyId(companyId).Identifier(identifier).Committente(committente).Prestatore(prestatore).FileName(fileName).LastUpdateFrom(lastUpdateFrom).LastUpdateTo(lastUpdateTo).DateSentFrom(dateSentFrom).DateSentTo(dateSentTo).DocumentDateFrom(documentDateFrom).DocumentDateTo(documentDateTo).DocumentNumber(documentNumber).Page(page).PageSize(pageSize).Execute()

List invoices



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
	resp, r, err := apiClient.SendAPI.InvoiceV1SendGet(context.Background()).CompanyId(companyId).Identifier(identifier).Committente(committente).Prestatore(prestatore).FileName(fileName).LastUpdateFrom(lastUpdateFrom).LastUpdateTo(lastUpdateTo).DateSentFrom(dateSentFrom).DateSentTo(dateSentTo).DocumentDateFrom(documentDateFrom).DocumentDateTo(documentDateTo).DocumentNumber(documentNumber).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendAPI.InvoiceV1SendGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceV1SendGet`: []Send
	fmt.Fprintf(os.Stdout, "Response from `SendAPI.InvoiceV1SendGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceV1SendGetRequest struct via the builder pattern


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
 **page** | **int32** | Page number. Defaults to 1. | [default to 1]
 **pageSize** | **int32** | Items per page. Defaults to 50. Cannot be greater than 200. | [default to 100]

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


## InvoiceV1SendIdGet

> Send InvoiceV1SendIdGet(ctx, id).Execute()

Get a invoice by id



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
	resp, r, err := apiClient.SendAPI.InvoiceV1SendIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendAPI.InvoiceV1SendIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceV1SendIdGet`: Send
	fmt.Fprintf(os.Stdout, "Response from `SendAPI.InvoiceV1SendIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Item id | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceV1SendIdGetRequest struct via the builder pattern


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


## InvoiceV1SendJsonPost

> Send InvoiceV1SendJsonPost(ctx).FatturaOrdinaria(fatturaOrdinaria).Validate(validate).Signature(signature).Execute()

Add an invoice by json



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
	fatturaOrdinaria := *invoicesdk.NewFatturaOrdinaria() // FatturaOrdinaria | 
	validate := true // bool | Validate the document first, and reject it on failure. (optional) (default to false)
	signature := "signature_example" // string | Whether to digitally sign the document. (optional) (default to "Auto")

	configuration := invoicesdk.NewConfiguration()
	apiClient := invoicesdk.NewAPIClient(configuration)
	resp, r, err := apiClient.SendAPI.InvoiceV1SendJsonPost(context.Background()).FatturaOrdinaria(fatturaOrdinaria).Validate(validate).Signature(signature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendAPI.InvoiceV1SendJsonPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceV1SendJsonPost`: Send
	fmt.Fprintf(os.Stdout, "Response from `SendAPI.InvoiceV1SendJsonPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceV1SendJsonPostRequest struct via the builder pattern


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


## InvoiceV1SendPost

> Send InvoiceV1SendPost(ctx).Send(send).Validate(validate).Signature(signature).Execute()

Add an invoice



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
	send := *invoicesdk.NewSend() // Send | 
	validate := true // bool | Validate the document first, and reject it on failure. (optional) (default to false)
	signature := "signature_example" // string | Whether to digitally sign the document. (optional) (default to "Auto")

	configuration := invoicesdk.NewConfiguration()
	apiClient := invoicesdk.NewAPIClient(configuration)
	resp, r, err := apiClient.SendAPI.InvoiceV1SendPost(context.Background()).Send(send).Validate(validate).Signature(signature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendAPI.InvoiceV1SendPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceV1SendPost`: Send
	fmt.Fprintf(os.Stdout, "Response from `SendAPI.InvoiceV1SendPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceV1SendPostRequest struct via the builder pattern


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


## InvoiceV1SendValidateFilesPost

> InvoiceV1SendValidateFilesPost(ctx).Files(files).Execute()

Validate an invoice by file



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
	files := []*os.File{"TODO"} // []*os.File | 

	configuration := invoicesdk.NewConfiguration()
	apiClient := invoicesdk.NewAPIClient(configuration)
	r, err := apiClient.SendAPI.InvoiceV1SendValidateFilesPost(context.Background()).Files(files).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendAPI.InvoiceV1SendValidateFilesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceV1SendValidateFilesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **files** | **[]*os.File** |  | 

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


## InvoiceV1SendValidateJsonPost

> InvoiceV1SendValidateJsonPost(ctx).FatturaOrdinaria(fatturaOrdinaria).Execute()

Validate an invoice by json



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
	fatturaOrdinaria := *invoicesdk.NewFatturaOrdinaria() // FatturaOrdinaria | 

	configuration := invoicesdk.NewConfiguration()
	apiClient := invoicesdk.NewAPIClient(configuration)
	r, err := apiClient.SendAPI.InvoiceV1SendValidateJsonPost(context.Background()).FatturaOrdinaria(fatturaOrdinaria).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendAPI.InvoiceV1SendValidateJsonPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceV1SendValidateJsonPostRequest struct via the builder pattern


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


## InvoiceV1SendValidatePost

> InvoiceV1SendValidatePost(ctx).Send(send).Execute()

Validate an invoice



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
	send := *invoicesdk.NewSend() // Send | 

	configuration := invoicesdk.NewConfiguration()
	apiClient := invoicesdk.NewAPIClient(configuration)
	r, err := apiClient.SendAPI.InvoiceV1SendValidatePost(context.Background()).Send(send).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendAPI.InvoiceV1SendValidatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceV1SendValidatePostRequest struct via the builder pattern


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


## InvoiceV1SendValidateXmlPost

> InvoiceV1SendValidateXmlPost(ctx).FatturaOrdinaria(fatturaOrdinaria).Execute()

Validate an invoice by xml



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
	fatturaOrdinaria := *invoicesdk.NewFatturaOrdinaria() // FatturaOrdinaria | 

	configuration := invoicesdk.NewConfiguration()
	apiClient := invoicesdk.NewAPIClient(configuration)
	r, err := apiClient.SendAPI.InvoiceV1SendValidateXmlPost(context.Background()).FatturaOrdinaria(fatturaOrdinaria).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendAPI.InvoiceV1SendValidateXmlPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceV1SendValidateXmlPostRequest struct via the builder pattern


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


## InvoiceV1SendXmlPost

> Send InvoiceV1SendXmlPost(ctx).FatturaOrdinaria(fatturaOrdinaria).Validate(validate).Signature(signature).Execute()

Add an invoice by xml



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
	fatturaOrdinaria := *invoicesdk.NewFatturaOrdinaria() // FatturaOrdinaria | 
	validate := true // bool | Validate the document first, and reject it on failure. (optional) (default to false)
	signature := "signature_example" // string | Whether to digitally sign the document. (optional) (default to "Auto")

	configuration := invoicesdk.NewConfiguration()
	apiClient := invoicesdk.NewAPIClient(configuration)
	resp, r, err := apiClient.SendAPI.InvoiceV1SendXmlPost(context.Background()).FatturaOrdinaria(fatturaOrdinaria).Validate(validate).Signature(signature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendAPI.InvoiceV1SendXmlPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceV1SendXmlPost`: Send
	fmt.Fprintf(os.Stdout, "Response from `SendAPI.InvoiceV1SendXmlPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceV1SendXmlPostRequest struct via the builder pattern


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

