/*
Invoicetronic API

The [Invoicetronic API][2] is a RESTful service that allows you to send and receive invoices through the Italian [Servizio di Interscambio (SDI)][1], or Interchange Service. The API is designed to be simple and easy to use, abstracting away SDI complexity while providing complete control over the invoice send/receive process. It provides advanced features as encryption at rest, multi-language pre-flight invoice validation, multiple upload formats, webhooks, event logging, client SDKs, and CLI tools.  For more information, see  [Invoicetronic website][2]  [1]: https://www.fatturapa.gov.it/it/sistemainterscambio/cose-il-sdi/ [2]: https://invoicetronic.com/

API version: 1
Contact: support@invoicetronic.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package invoicetronicsdk

import (
	"encoding/json"
	"time"
)

// checks if the DatiContratto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatiContratto{}

// DatiContratto struct for DatiContratto
type DatiContratto struct {
	RiferimentoNumeroLinea []int32 `json:"riferimento_numero_linea,omitempty"`
	IdDocumento NullableString `json:"id_documento,omitempty"`
	Data NullableTime `json:"data,omitempty"`
	NumItem NullableString `json:"num_item,omitempty"`
	CodiceCommessaConvenzione NullableString `json:"codice_commessa_convenzione,omitempty"`
	CodiceCup NullableString `json:"codice_cup,omitempty"`
	CodiceCig NullableString `json:"codice_cig,omitempty"`
}

// NewDatiContratto instantiates a new DatiContratto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatiContratto() *DatiContratto {
	this := DatiContratto{}
	return &this
}

// NewDatiContrattoWithDefaults instantiates a new DatiContratto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatiContrattoWithDefaults() *DatiContratto {
	this := DatiContratto{}
	return &this
}

// GetRiferimentoNumeroLinea returns the RiferimentoNumeroLinea field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatiContratto) GetRiferimentoNumeroLinea() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.RiferimentoNumeroLinea
}

// GetRiferimentoNumeroLineaOk returns a tuple with the RiferimentoNumeroLinea field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatiContratto) GetRiferimentoNumeroLineaOk() ([]int32, bool) {
	if o == nil || IsNil(o.RiferimentoNumeroLinea) {
		return nil, false
	}
	return o.RiferimentoNumeroLinea, true
}

// HasRiferimentoNumeroLinea returns a boolean if a field has been set.
func (o *DatiContratto) HasRiferimentoNumeroLinea() bool {
	if o != nil && !IsNil(o.RiferimentoNumeroLinea) {
		return true
	}

	return false
}

// SetRiferimentoNumeroLinea gets a reference to the given []int32 and assigns it to the RiferimentoNumeroLinea field.
func (o *DatiContratto) SetRiferimentoNumeroLinea(v []int32) {
	o.RiferimentoNumeroLinea = v
}

// GetIdDocumento returns the IdDocumento field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatiContratto) GetIdDocumento() string {
	if o == nil || IsNil(o.IdDocumento.Get()) {
		var ret string
		return ret
	}
	return *o.IdDocumento.Get()
}

// GetIdDocumentoOk returns a tuple with the IdDocumento field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatiContratto) GetIdDocumentoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IdDocumento.Get(), o.IdDocumento.IsSet()
}

// HasIdDocumento returns a boolean if a field has been set.
func (o *DatiContratto) HasIdDocumento() bool {
	if o != nil && o.IdDocumento.IsSet() {
		return true
	}

	return false
}

// SetIdDocumento gets a reference to the given NullableString and assigns it to the IdDocumento field.
func (o *DatiContratto) SetIdDocumento(v string) {
	o.IdDocumento.Set(&v)
}
// SetIdDocumentoNil sets the value for IdDocumento to be an explicit nil
func (o *DatiContratto) SetIdDocumentoNil() {
	o.IdDocumento.Set(nil)
}

// UnsetIdDocumento ensures that no value is present for IdDocumento, not even an explicit nil
func (o *DatiContratto) UnsetIdDocumento() {
	o.IdDocumento.Unset()
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatiContratto) GetData() time.Time {
	if o == nil || IsNil(o.Data.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatiContratto) GetDataOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *DatiContratto) HasData() bool {
	if o != nil && o.Data.IsSet() {
		return true
	}

	return false
}

// SetData gets a reference to the given NullableTime and assigns it to the Data field.
func (o *DatiContratto) SetData(v time.Time) {
	o.Data.Set(&v)
}
// SetDataNil sets the value for Data to be an explicit nil
func (o *DatiContratto) SetDataNil() {
	o.Data.Set(nil)
}

// UnsetData ensures that no value is present for Data, not even an explicit nil
func (o *DatiContratto) UnsetData() {
	o.Data.Unset()
}

// GetNumItem returns the NumItem field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatiContratto) GetNumItem() string {
	if o == nil || IsNil(o.NumItem.Get()) {
		var ret string
		return ret
	}
	return *o.NumItem.Get()
}

// GetNumItemOk returns a tuple with the NumItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatiContratto) GetNumItemOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NumItem.Get(), o.NumItem.IsSet()
}

// HasNumItem returns a boolean if a field has been set.
func (o *DatiContratto) HasNumItem() bool {
	if o != nil && o.NumItem.IsSet() {
		return true
	}

	return false
}

// SetNumItem gets a reference to the given NullableString and assigns it to the NumItem field.
func (o *DatiContratto) SetNumItem(v string) {
	o.NumItem.Set(&v)
}
// SetNumItemNil sets the value for NumItem to be an explicit nil
func (o *DatiContratto) SetNumItemNil() {
	o.NumItem.Set(nil)
}

// UnsetNumItem ensures that no value is present for NumItem, not even an explicit nil
func (o *DatiContratto) UnsetNumItem() {
	o.NumItem.Unset()
}

// GetCodiceCommessaConvenzione returns the CodiceCommessaConvenzione field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatiContratto) GetCodiceCommessaConvenzione() string {
	if o == nil || IsNil(o.CodiceCommessaConvenzione.Get()) {
		var ret string
		return ret
	}
	return *o.CodiceCommessaConvenzione.Get()
}

// GetCodiceCommessaConvenzioneOk returns a tuple with the CodiceCommessaConvenzione field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatiContratto) GetCodiceCommessaConvenzioneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CodiceCommessaConvenzione.Get(), o.CodiceCommessaConvenzione.IsSet()
}

// HasCodiceCommessaConvenzione returns a boolean if a field has been set.
func (o *DatiContratto) HasCodiceCommessaConvenzione() bool {
	if o != nil && o.CodiceCommessaConvenzione.IsSet() {
		return true
	}

	return false
}

// SetCodiceCommessaConvenzione gets a reference to the given NullableString and assigns it to the CodiceCommessaConvenzione field.
func (o *DatiContratto) SetCodiceCommessaConvenzione(v string) {
	o.CodiceCommessaConvenzione.Set(&v)
}
// SetCodiceCommessaConvenzioneNil sets the value for CodiceCommessaConvenzione to be an explicit nil
func (o *DatiContratto) SetCodiceCommessaConvenzioneNil() {
	o.CodiceCommessaConvenzione.Set(nil)
}

// UnsetCodiceCommessaConvenzione ensures that no value is present for CodiceCommessaConvenzione, not even an explicit nil
func (o *DatiContratto) UnsetCodiceCommessaConvenzione() {
	o.CodiceCommessaConvenzione.Unset()
}

// GetCodiceCup returns the CodiceCup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatiContratto) GetCodiceCup() string {
	if o == nil || IsNil(o.CodiceCup.Get()) {
		var ret string
		return ret
	}
	return *o.CodiceCup.Get()
}

// GetCodiceCupOk returns a tuple with the CodiceCup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatiContratto) GetCodiceCupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CodiceCup.Get(), o.CodiceCup.IsSet()
}

// HasCodiceCup returns a boolean if a field has been set.
func (o *DatiContratto) HasCodiceCup() bool {
	if o != nil && o.CodiceCup.IsSet() {
		return true
	}

	return false
}

// SetCodiceCup gets a reference to the given NullableString and assigns it to the CodiceCup field.
func (o *DatiContratto) SetCodiceCup(v string) {
	o.CodiceCup.Set(&v)
}
// SetCodiceCupNil sets the value for CodiceCup to be an explicit nil
func (o *DatiContratto) SetCodiceCupNil() {
	o.CodiceCup.Set(nil)
}

// UnsetCodiceCup ensures that no value is present for CodiceCup, not even an explicit nil
func (o *DatiContratto) UnsetCodiceCup() {
	o.CodiceCup.Unset()
}

// GetCodiceCig returns the CodiceCig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatiContratto) GetCodiceCig() string {
	if o == nil || IsNil(o.CodiceCig.Get()) {
		var ret string
		return ret
	}
	return *o.CodiceCig.Get()
}

// GetCodiceCigOk returns a tuple with the CodiceCig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatiContratto) GetCodiceCigOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CodiceCig.Get(), o.CodiceCig.IsSet()
}

// HasCodiceCig returns a boolean if a field has been set.
func (o *DatiContratto) HasCodiceCig() bool {
	if o != nil && o.CodiceCig.IsSet() {
		return true
	}

	return false
}

// SetCodiceCig gets a reference to the given NullableString and assigns it to the CodiceCig field.
func (o *DatiContratto) SetCodiceCig(v string) {
	o.CodiceCig.Set(&v)
}
// SetCodiceCigNil sets the value for CodiceCig to be an explicit nil
func (o *DatiContratto) SetCodiceCigNil() {
	o.CodiceCig.Set(nil)
}

// UnsetCodiceCig ensures that no value is present for CodiceCig, not even an explicit nil
func (o *DatiContratto) UnsetCodiceCig() {
	o.CodiceCig.Unset()
}

func (o DatiContratto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatiContratto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.RiferimentoNumeroLinea != nil {
		toSerialize["riferimento_numero_linea"] = o.RiferimentoNumeroLinea
	}
	if o.IdDocumento.IsSet() {
		toSerialize["id_documento"] = o.IdDocumento.Get()
	}
	if o.Data.IsSet() {
		toSerialize["data"] = o.Data.Get()
	}
	if o.NumItem.IsSet() {
		toSerialize["num_item"] = o.NumItem.Get()
	}
	if o.CodiceCommessaConvenzione.IsSet() {
		toSerialize["codice_commessa_convenzione"] = o.CodiceCommessaConvenzione.Get()
	}
	if o.CodiceCup.IsSet() {
		toSerialize["codice_cup"] = o.CodiceCup.Get()
	}
	if o.CodiceCig.IsSet() {
		toSerialize["codice_cig"] = o.CodiceCig.Get()
	}
	return toSerialize, nil
}

type NullableDatiContratto struct {
	value *DatiContratto
	isSet bool
}

func (v NullableDatiContratto) Get() *DatiContratto {
	return v.value
}

func (v *NullableDatiContratto) Set(val *DatiContratto) {
	v.value = val
	v.isSet = true
}

func (v NullableDatiContratto) IsSet() bool {
	return v.isSet
}

func (v *NullableDatiContratto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatiContratto(val *DatiContratto) *NullableDatiContratto {
	return &NullableDatiContratto{value: val, isSet: true}
}

func (v NullableDatiContratto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatiContratto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


