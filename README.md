# Go SDK for the Italian eInvoice API

The Italian eInvoice API is a RESTful API that allows you to send and receive invoices through the
Italian [Servizio di Interscambio (SDI)][1], or Interchange Service. The API is designed by Invoicetronic to be simple
and easy to use, abstracting away SDI complexity while still providing complete control over the
invoice send/receive process. The API also provides advanced features and a rich toolchain, such as invoice validation,
multiple upload methods, webhooks, event logs, CORS support, client SDKs for commonly used languages, and CLI tools.

For more information, see  [Invoicetronic website][2]

[1]: https://www.fatturapa.gov.it/it/sistemainterscambio/cose-il-sdi/
[2]: https://invoicetronic.com/

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Generator version: 7.10.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://invoicetronic.com](https://invoicetronic.com)

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import invoicesdk "github.com/invoicetronic/invoice-go-sdk"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `invoicesdk.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), invoicesdk.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `invoicesdk.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), invoicesdk.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `invoicesdk.ContextOperationServerIndices` and `invoicesdk.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), invoicesdk.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), invoicesdk.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*CompanyAPI* | [**InvoiceV1CompanyGet**](docs/CompanyAPI.md#invoicev1companyget) | **Get** /invoice/v1/company | List companies
*CompanyAPI* | [**InvoiceV1CompanyIdDelete**](docs/CompanyAPI.md#invoicev1companyiddelete) | **Delete** /invoice/v1/company/{id} | Delete a company
*CompanyAPI* | [**InvoiceV1CompanyIdGet**](docs/CompanyAPI.md#invoicev1companyidget) | **Get** /invoice/v1/company/{id} | Get a company by id
*CompanyAPI* | [**InvoiceV1CompanyPost**](docs/CompanyAPI.md#invoicev1companypost) | **Post** /invoice/v1/company | Add a company
*CompanyAPI* | [**InvoiceV1CompanyPut**](docs/CompanyAPI.md#invoicev1companyput) | **Put** /invoice/v1/company | Update a company
*LogAPI* | [**InvoiceV1LogGet**](docs/LogAPI.md#invoicev1logget) | **Get** /invoice/v1/log | List events
*LogAPI* | [**InvoiceV1LogIdGet**](docs/LogAPI.md#invoicev1logidget) | **Get** /invoice/v1/log/{id} | Get an event by id
*ReceiveAPI* | [**InvoiceV1ReceiveGet**](docs/ReceiveAPI.md#invoicev1receiveget) | **Get** /invoice/v1/receive | List incoming invoices
*ReceiveAPI* | [**InvoiceV1ReceiveIdDelete**](docs/ReceiveAPI.md#invoicev1receiveiddelete) | **Delete** /invoice/v1/receive/{id} | Delete an incoming invoice by id
*ReceiveAPI* | [**InvoiceV1ReceiveIdGet**](docs/ReceiveAPI.md#invoicev1receiveidget) | **Get** /invoice/v1/receive/{id} | Get an incoming invoice by id
*SendAPI* | [**InvoiceV1SendFilesPost**](docs/SendAPI.md#invoicev1sendfilespost) | **Post** /invoice/v1/send/files | Add a send invoice by file
*SendAPI* | [**InvoiceV1SendGet**](docs/SendAPI.md#invoicev1sendget) | **Get** /invoice/v1/send | List send invoices
*SendAPI* | [**InvoiceV1SendIdGet**](docs/SendAPI.md#invoicev1sendidget) | **Get** /invoice/v1/send/{id} | Get a send invoice by id
*SendAPI* | [**InvoiceV1SendJsonPost**](docs/SendAPI.md#invoicev1sendjsonpost) | **Post** /invoice/v1/send/json | Add a send invoice by json
*SendAPI* | [**InvoiceV1SendPost**](docs/SendAPI.md#invoicev1sendpost) | **Post** /invoice/v1/send | Add a send invoice
*SendAPI* | [**InvoiceV1SendXmlPost**](docs/SendAPI.md#invoicev1sendxmlpost) | **Post** /invoice/v1/send/xml | Add a send invoice by xml
*UpdateAPI* | [**InvoiceV1UpdateGet**](docs/UpdateAPI.md#invoicev1updateget) | **Get** /invoice/v1/update | List updates
*UpdateAPI* | [**InvoiceV1UpdateIdGet**](docs/UpdateAPI.md#invoicev1updateidget) | **Get** /invoice/v1/update/{id} | Get an update by id
*WebhookAPI* | [**InvoiceV1WebhookGet**](docs/WebhookAPI.md#invoicev1webhookget) | **Get** /invoice/v1/webhook | List webhooks
*WebhookAPI* | [**InvoiceV1WebhookIdDelete**](docs/WebhookAPI.md#invoicev1webhookiddelete) | **Delete** /invoice/v1/webhook/{id} | Delete a webhook by id
*WebhookAPI* | [**InvoiceV1WebhookIdGet**](docs/WebhookAPI.md#invoicev1webhookidget) | **Get** /invoice/v1/webhook/{id} | Get a webhook by id
*WebhookAPI* | [**InvoiceV1WebhookPost**](docs/WebhookAPI.md#invoicev1webhookpost) | **Post** /invoice/v1/webhook | Add a webhook
*WebhookAPI* | [**InvoiceV1WebhookPut**](docs/WebhookAPI.md#invoicev1webhookput) | **Put** /invoice/v1/webhook | Update a webhook
*WebhookAPI* | [**InvoiceV1WebhookhistoryGet**](docs/WebhookAPI.md#invoicev1webhookhistoryget) | **Get** /invoice/v1/webhookhistory | List webhook history items
*WebhookAPI* | [**InvoiceV1WebhookhistoryIdGet**](docs/WebhookAPI.md#invoicev1webhookhistoryidget) | **Get** /invoice/v1/webhookhistory/{id} | Get a webhook history item by id


## Documentation For Models

 - [Allegati](docs/Allegati.md)
 - [AltriDatiGestionali](docs/AltriDatiGestionali.md)
 - [Anagrafica](docs/Anagrafica.md)
 - [CedentePrestatore](docs/CedentePrestatore.md)
 - [CessionarioCommittente](docs/CessionarioCommittente.md)
 - [CodiceArticolo](docs/CodiceArticolo.md)
 - [Company](docs/Company.md)
 - [Contatti](docs/Contatti.md)
 - [ContattiTrasmittente](docs/ContattiTrasmittente.md)
 - [DatiAnagrafici](docs/DatiAnagrafici.md)
 - [DatiAnagraficiCedentePrestatore](docs/DatiAnagraficiCedentePrestatore.md)
 - [DatiAnagraficiCessionarioCommittente](docs/DatiAnagraficiCessionarioCommittente.md)
 - [DatiAnagraficiVettore](docs/DatiAnagraficiVettore.md)
 - [DatiBeniServizi](docs/DatiBeniServizi.md)
 - [DatiBollo](docs/DatiBollo.md)
 - [DatiCassaPrevidenziale](docs/DatiCassaPrevidenziale.md)
 - [DatiContratto](docs/DatiContratto.md)
 - [DatiConvenzione](docs/DatiConvenzione.md)
 - [DatiDDT](docs/DatiDDT.md)
 - [DatiFattureCollegate](docs/DatiFattureCollegate.md)
 - [DatiGenerali](docs/DatiGenerali.md)
 - [DatiGeneraliDocumento](docs/DatiGeneraliDocumento.md)
 - [DatiOrdineAcquisto](docs/DatiOrdineAcquisto.md)
 - [DatiPagamento](docs/DatiPagamento.md)
 - [DatiRicezione](docs/DatiRicezione.md)
 - [DatiRiepilogo](docs/DatiRiepilogo.md)
 - [DatiRitenuta](docs/DatiRitenuta.md)
 - [DatiSAL](docs/DatiSAL.md)
 - [DatiTrasmissione](docs/DatiTrasmissione.md)
 - [DatiTrasporto](docs/DatiTrasporto.md)
 - [DatiVeicoli](docs/DatiVeicoli.md)
 - [DettaglioLinee](docs/DettaglioLinee.md)
 - [DettaglioPagamento](docs/DettaglioPagamento.md)
 - [DocumentData](docs/DocumentData.md)
 - [Error](docs/Error.md)
 - [Event](docs/Event.md)
 - [FatturaElettronicaBody](docs/FatturaElettronicaBody.md)
 - [FatturaElettronicaHeader](docs/FatturaElettronicaHeader.md)
 - [FatturaOrdinaria](docs/FatturaOrdinaria.md)
 - [FatturaPrincipale](docs/FatturaPrincipale.md)
 - [IdFiscaleIVA](docs/IdFiscaleIVA.md)
 - [IdTrasmittente](docs/IdTrasmittente.md)
 - [IndirizzoResa](docs/IndirizzoResa.md)
 - [IscrizioneREA](docs/IscrizioneREA.md)
 - [ProblemDetails](docs/ProblemDetails.md)
 - [RappresentanteFiscale](docs/RappresentanteFiscale.md)
 - [RappresentanteFiscaleCessionarioCommittente](docs/RappresentanteFiscaleCessionarioCommittente.md)
 - [Receive](docs/Receive.md)
 - [ScontoMaggiorazione](docs/ScontoMaggiorazione.md)
 - [SedeCedentePrestatore](docs/SedeCedentePrestatore.md)
 - [SedeCessionarioCommittente](docs/SedeCessionarioCommittente.md)
 - [Send](docs/Send.md)
 - [StabileOrganizzazione](docs/StabileOrganizzazione.md)
 - [TerzoIntermediarioOSoggettoEmittente](docs/TerzoIntermediarioOSoggettoEmittente.md)
 - [Update](docs/Update.md)
 - [WebHook](docs/WebHook.md)
 - [WebHookHistory](docs/WebHookHistory.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### Basic

- **Type**: HTTP basic authentication

Example

```go
auth := context.WithValue(context.Background(), invoicesdk.ContextBasicAuth, invoicesdk.BasicAuth{
	UserName: "username",
	Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support@invoicetronic.com

