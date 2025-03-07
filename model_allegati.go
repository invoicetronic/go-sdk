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
)

// checks if the Allegati type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Allegati{}

// Allegati struct for Allegati
type Allegati struct {
	NomeAttachment NullableString `json:"nome_attachment,omitempty"`
	AlgoritmoCompressione NullableString `json:"algoritmo_compressione,omitempty"`
	FormatoAttachment NullableString `json:"formato_attachment,omitempty"`
	DescrizioneAttachment NullableString `json:"descrizione_attachment,omitempty"`
	Attachment NullableString `json:"attachment,omitempty"`
}

// NewAllegati instantiates a new Allegati object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllegati() *Allegati {
	this := Allegati{}
	return &this
}

// NewAllegatiWithDefaults instantiates a new Allegati object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllegatiWithDefaults() *Allegati {
	this := Allegati{}
	return &this
}

// GetNomeAttachment returns the NomeAttachment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Allegati) GetNomeAttachment() string {
	if o == nil || IsNil(o.NomeAttachment.Get()) {
		var ret string
		return ret
	}
	return *o.NomeAttachment.Get()
}

// GetNomeAttachmentOk returns a tuple with the NomeAttachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Allegati) GetNomeAttachmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NomeAttachment.Get(), o.NomeAttachment.IsSet()
}

// HasNomeAttachment returns a boolean if a field has been set.
func (o *Allegati) HasNomeAttachment() bool {
	if o != nil && o.NomeAttachment.IsSet() {
		return true
	}

	return false
}

// SetNomeAttachment gets a reference to the given NullableString and assigns it to the NomeAttachment field.
func (o *Allegati) SetNomeAttachment(v string) {
	o.NomeAttachment.Set(&v)
}
// SetNomeAttachmentNil sets the value for NomeAttachment to be an explicit nil
func (o *Allegati) SetNomeAttachmentNil() {
	o.NomeAttachment.Set(nil)
}

// UnsetNomeAttachment ensures that no value is present for NomeAttachment, not even an explicit nil
func (o *Allegati) UnsetNomeAttachment() {
	o.NomeAttachment.Unset()
}

// GetAlgoritmoCompressione returns the AlgoritmoCompressione field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Allegati) GetAlgoritmoCompressione() string {
	if o == nil || IsNil(o.AlgoritmoCompressione.Get()) {
		var ret string
		return ret
	}
	return *o.AlgoritmoCompressione.Get()
}

// GetAlgoritmoCompressioneOk returns a tuple with the AlgoritmoCompressione field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Allegati) GetAlgoritmoCompressioneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AlgoritmoCompressione.Get(), o.AlgoritmoCompressione.IsSet()
}

// HasAlgoritmoCompressione returns a boolean if a field has been set.
func (o *Allegati) HasAlgoritmoCompressione() bool {
	if o != nil && o.AlgoritmoCompressione.IsSet() {
		return true
	}

	return false
}

// SetAlgoritmoCompressione gets a reference to the given NullableString and assigns it to the AlgoritmoCompressione field.
func (o *Allegati) SetAlgoritmoCompressione(v string) {
	o.AlgoritmoCompressione.Set(&v)
}
// SetAlgoritmoCompressioneNil sets the value for AlgoritmoCompressione to be an explicit nil
func (o *Allegati) SetAlgoritmoCompressioneNil() {
	o.AlgoritmoCompressione.Set(nil)
}

// UnsetAlgoritmoCompressione ensures that no value is present for AlgoritmoCompressione, not even an explicit nil
func (o *Allegati) UnsetAlgoritmoCompressione() {
	o.AlgoritmoCompressione.Unset()
}

// GetFormatoAttachment returns the FormatoAttachment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Allegati) GetFormatoAttachment() string {
	if o == nil || IsNil(o.FormatoAttachment.Get()) {
		var ret string
		return ret
	}
	return *o.FormatoAttachment.Get()
}

// GetFormatoAttachmentOk returns a tuple with the FormatoAttachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Allegati) GetFormatoAttachmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FormatoAttachment.Get(), o.FormatoAttachment.IsSet()
}

// HasFormatoAttachment returns a boolean if a field has been set.
func (o *Allegati) HasFormatoAttachment() bool {
	if o != nil && o.FormatoAttachment.IsSet() {
		return true
	}

	return false
}

// SetFormatoAttachment gets a reference to the given NullableString and assigns it to the FormatoAttachment field.
func (o *Allegati) SetFormatoAttachment(v string) {
	o.FormatoAttachment.Set(&v)
}
// SetFormatoAttachmentNil sets the value for FormatoAttachment to be an explicit nil
func (o *Allegati) SetFormatoAttachmentNil() {
	o.FormatoAttachment.Set(nil)
}

// UnsetFormatoAttachment ensures that no value is present for FormatoAttachment, not even an explicit nil
func (o *Allegati) UnsetFormatoAttachment() {
	o.FormatoAttachment.Unset()
}

// GetDescrizioneAttachment returns the DescrizioneAttachment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Allegati) GetDescrizioneAttachment() string {
	if o == nil || IsNil(o.DescrizioneAttachment.Get()) {
		var ret string
		return ret
	}
	return *o.DescrizioneAttachment.Get()
}

// GetDescrizioneAttachmentOk returns a tuple with the DescrizioneAttachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Allegati) GetDescrizioneAttachmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DescrizioneAttachment.Get(), o.DescrizioneAttachment.IsSet()
}

// HasDescrizioneAttachment returns a boolean if a field has been set.
func (o *Allegati) HasDescrizioneAttachment() bool {
	if o != nil && o.DescrizioneAttachment.IsSet() {
		return true
	}

	return false
}

// SetDescrizioneAttachment gets a reference to the given NullableString and assigns it to the DescrizioneAttachment field.
func (o *Allegati) SetDescrizioneAttachment(v string) {
	o.DescrizioneAttachment.Set(&v)
}
// SetDescrizioneAttachmentNil sets the value for DescrizioneAttachment to be an explicit nil
func (o *Allegati) SetDescrizioneAttachmentNil() {
	o.DescrizioneAttachment.Set(nil)
}

// UnsetDescrizioneAttachment ensures that no value is present for DescrizioneAttachment, not even an explicit nil
func (o *Allegati) UnsetDescrizioneAttachment() {
	o.DescrizioneAttachment.Unset()
}

// GetAttachment returns the Attachment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Allegati) GetAttachment() string {
	if o == nil || IsNil(o.Attachment.Get()) {
		var ret string
		return ret
	}
	return *o.Attachment.Get()
}

// GetAttachmentOk returns a tuple with the Attachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Allegati) GetAttachmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attachment.Get(), o.Attachment.IsSet()
}

// HasAttachment returns a boolean if a field has been set.
func (o *Allegati) HasAttachment() bool {
	if o != nil && o.Attachment.IsSet() {
		return true
	}

	return false
}

// SetAttachment gets a reference to the given NullableString and assigns it to the Attachment field.
func (o *Allegati) SetAttachment(v string) {
	o.Attachment.Set(&v)
}
// SetAttachmentNil sets the value for Attachment to be an explicit nil
func (o *Allegati) SetAttachmentNil() {
	o.Attachment.Set(nil)
}

// UnsetAttachment ensures that no value is present for Attachment, not even an explicit nil
func (o *Allegati) UnsetAttachment() {
	o.Attachment.Unset()
}

func (o Allegati) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Allegati) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.NomeAttachment.IsSet() {
		toSerialize["nome_attachment"] = o.NomeAttachment.Get()
	}
	if o.AlgoritmoCompressione.IsSet() {
		toSerialize["algoritmo_compressione"] = o.AlgoritmoCompressione.Get()
	}
	if o.FormatoAttachment.IsSet() {
		toSerialize["formato_attachment"] = o.FormatoAttachment.Get()
	}
	if o.DescrizioneAttachment.IsSet() {
		toSerialize["descrizione_attachment"] = o.DescrizioneAttachment.Get()
	}
	if o.Attachment.IsSet() {
		toSerialize["attachment"] = o.Attachment.Get()
	}
	return toSerialize, nil
}

type NullableAllegati struct {
	value *Allegati
	isSet bool
}

func (v NullableAllegati) Get() *Allegati {
	return v.value
}

func (v *NullableAllegati) Set(val *Allegati) {
	v.value = val
	v.isSet = true
}

func (v NullableAllegati) IsSet() bool {
	return v.isSet
}

func (v *NullableAllegati) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllegati(val *Allegati) *NullableAllegati {
	return &NullableAllegati{value: val, isSet: true}
}

func (v NullableAllegati) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllegati) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


