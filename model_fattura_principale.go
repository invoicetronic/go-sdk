/*
Italian eInvoice API v1

The [Italian eInvoice API][2] is a RESTful API that allows you to send and receive invoices through the Italian [Servizio di Interscambio (SDI)][1], or Interchange Service. The API is designed by Invoicetronic to be simple and easy to use, abstracting away SDI complexity while providing complete control over the invoice send/receive process. The API also provides advanced features as encryption at rest, invoice validation, multiple upload formats, webhooks, event logging, client SDKs for commonly used languages, and CLI tools.  For more information, see  [Invoicetronic website][2]  [1]: https://www.fatturapa.gov.it/it/sistemainterscambio/cose-il-sdi/ [2]: https://invoicetronic.com/

API version: 1
Contact: support@invoicetronic.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package invoicesdk

import (
	"encoding/json"
	"time"
)

// checks if the FatturaPrincipale type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FatturaPrincipale{}

// FatturaPrincipale struct for FatturaPrincipale
type FatturaPrincipale struct {
	NumeroFatturaPrincipale NullableString `json:"numero_fattura_principale,omitempty"`
	DataFatturaPrincipale NullableTime `json:"data_fattura_principale,omitempty"`
}

// NewFatturaPrincipale instantiates a new FatturaPrincipale object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFatturaPrincipale() *FatturaPrincipale {
	this := FatturaPrincipale{}
	return &this
}

// NewFatturaPrincipaleWithDefaults instantiates a new FatturaPrincipale object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFatturaPrincipaleWithDefaults() *FatturaPrincipale {
	this := FatturaPrincipale{}
	return &this
}

// GetNumeroFatturaPrincipale returns the NumeroFatturaPrincipale field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FatturaPrincipale) GetNumeroFatturaPrincipale() string {
	if o == nil || IsNil(o.NumeroFatturaPrincipale.Get()) {
		var ret string
		return ret
	}
	return *o.NumeroFatturaPrincipale.Get()
}

// GetNumeroFatturaPrincipaleOk returns a tuple with the NumeroFatturaPrincipale field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FatturaPrincipale) GetNumeroFatturaPrincipaleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NumeroFatturaPrincipale.Get(), o.NumeroFatturaPrincipale.IsSet()
}

// HasNumeroFatturaPrincipale returns a boolean if a field has been set.
func (o *FatturaPrincipale) HasNumeroFatturaPrincipale() bool {
	if o != nil && o.NumeroFatturaPrincipale.IsSet() {
		return true
	}

	return false
}

// SetNumeroFatturaPrincipale gets a reference to the given NullableString and assigns it to the NumeroFatturaPrincipale field.
func (o *FatturaPrincipale) SetNumeroFatturaPrincipale(v string) {
	o.NumeroFatturaPrincipale.Set(&v)
}
// SetNumeroFatturaPrincipaleNil sets the value for NumeroFatturaPrincipale to be an explicit nil
func (o *FatturaPrincipale) SetNumeroFatturaPrincipaleNil() {
	o.NumeroFatturaPrincipale.Set(nil)
}

// UnsetNumeroFatturaPrincipale ensures that no value is present for NumeroFatturaPrincipale, not even an explicit nil
func (o *FatturaPrincipale) UnsetNumeroFatturaPrincipale() {
	o.NumeroFatturaPrincipale.Unset()
}

// GetDataFatturaPrincipale returns the DataFatturaPrincipale field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FatturaPrincipale) GetDataFatturaPrincipale() time.Time {
	if o == nil || IsNil(o.DataFatturaPrincipale.Get()) {
		var ret time.Time
		return ret
	}
	return *o.DataFatturaPrincipale.Get()
}

// GetDataFatturaPrincipaleOk returns a tuple with the DataFatturaPrincipale field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FatturaPrincipale) GetDataFatturaPrincipaleOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataFatturaPrincipale.Get(), o.DataFatturaPrincipale.IsSet()
}

// HasDataFatturaPrincipale returns a boolean if a field has been set.
func (o *FatturaPrincipale) HasDataFatturaPrincipale() bool {
	if o != nil && o.DataFatturaPrincipale.IsSet() {
		return true
	}

	return false
}

// SetDataFatturaPrincipale gets a reference to the given NullableTime and assigns it to the DataFatturaPrincipale field.
func (o *FatturaPrincipale) SetDataFatturaPrincipale(v time.Time) {
	o.DataFatturaPrincipale.Set(&v)
}
// SetDataFatturaPrincipaleNil sets the value for DataFatturaPrincipale to be an explicit nil
func (o *FatturaPrincipale) SetDataFatturaPrincipaleNil() {
	o.DataFatturaPrincipale.Set(nil)
}

// UnsetDataFatturaPrincipale ensures that no value is present for DataFatturaPrincipale, not even an explicit nil
func (o *FatturaPrincipale) UnsetDataFatturaPrincipale() {
	o.DataFatturaPrincipale.Unset()
}

func (o FatturaPrincipale) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FatturaPrincipale) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.NumeroFatturaPrincipale.IsSet() {
		toSerialize["numero_fattura_principale"] = o.NumeroFatturaPrincipale.Get()
	}
	if o.DataFatturaPrincipale.IsSet() {
		toSerialize["data_fattura_principale"] = o.DataFatturaPrincipale.Get()
	}
	return toSerialize, nil
}

type NullableFatturaPrincipale struct {
	value *FatturaPrincipale
	isSet bool
}

func (v NullableFatturaPrincipale) Get() *FatturaPrincipale {
	return v.value
}

func (v *NullableFatturaPrincipale) Set(val *FatturaPrincipale) {
	v.value = val
	v.isSet = true
}

func (v NullableFatturaPrincipale) IsSet() bool {
	return v.isSet
}

func (v *NullableFatturaPrincipale) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFatturaPrincipale(val *FatturaPrincipale) *NullableFatturaPrincipale {
	return &NullableFatturaPrincipale{value: val, isSet: true}
}

func (v NullableFatturaPrincipale) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFatturaPrincipale) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


