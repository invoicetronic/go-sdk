# Company

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique identifier. Leave it at 0 for new records as it will be set automatically. | [optional] 
**Created** | Pointer to **time.Time** | Creation date. It is set automatically. | [optional] 
**Version** | Pointer to **int32** | Row version, for optimistic concurrency. It is set automatically. | [optional] 
**UserId** | Pointer to **int32** | User id. | [optional] 
**Vat** | **string** | Vat number. Must include the country code. | 
**FiscalCode** | **string** | Fiscal code. In most cases it&#39;s the same as the vat number. | 
**Name** | **string** | Name | 
**Counter** | Pointer to **int32** | Holds the last unique value used to generate a XML filename. This is automatically updated by the system   when a raw XML file is uploaded. Normally, you do not need or want to change this value. | [optional] 

## Methods

### NewCompany

`func NewCompany(vat string, fiscalCode string, name string, ) *Company`

NewCompany instantiates a new Company object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyWithDefaults

`func NewCompanyWithDefaults() *Company`

NewCompanyWithDefaults instantiates a new Company object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Company) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Company) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Company) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Company) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreated

`func (o *Company) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Company) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Company) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Company) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetVersion

`func (o *Company) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Company) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Company) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Company) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetUserId

`func (o *Company) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Company) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Company) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Company) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetVat

`func (o *Company) GetVat() string`

GetVat returns the Vat field if non-nil, zero value otherwise.

### GetVatOk

`func (o *Company) GetVatOk() (*string, bool)`

GetVatOk returns a tuple with the Vat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVat

`func (o *Company) SetVat(v string)`

SetVat sets Vat field to given value.


### GetFiscalCode

`func (o *Company) GetFiscalCode() string`

GetFiscalCode returns the FiscalCode field if non-nil, zero value otherwise.

### GetFiscalCodeOk

`func (o *Company) GetFiscalCodeOk() (*string, bool)`

GetFiscalCodeOk returns a tuple with the FiscalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalCode

`func (o *Company) SetFiscalCode(v string)`

SetFiscalCode sets FiscalCode field to given value.


### GetName

`func (o *Company) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Company) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Company) SetName(v string)`

SetName sets Name field to given value.


### GetCounter

`func (o *Company) GetCounter() int32`

GetCounter returns the Counter field if non-nil, zero value otherwise.

### GetCounterOk

`func (o *Company) GetCounterOk() (*int32, bool)`

GetCounterOk returns a tuple with the Counter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounter

`func (o *Company) SetCounter(v int32)`

SetCounter sets Counter field to given value.

### HasCounter

`func (o *Company) HasCounter() bool`

HasCounter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


