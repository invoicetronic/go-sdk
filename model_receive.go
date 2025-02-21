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

// checks if the Receive type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Receive{}

// Receive struct for Receive
type Receive struct {
	// Unique identifier. Leave it at 0 for new records as it will be set automatically.
	Id *int32 `json:"id,omitempty"`
	// Creation date. It is set automatically.
	Created *time.Time `json:"created,omitempty"`
	// Row version, for optimistic concurrency. It is set automatically.
	Version *int32 `json:"version,omitempty"`
	// User id.
	UserId *int32 `json:"user_id,omitempty"`
	// Company id. On send, this is the sender and must be set in advance. On receive, it will be  automatically set based on the recipient's VAT number. If a matching company is not found, the invoice will be rejected until the company is created.
	CompanyId *int32 `json:"company_id,omitempty"`
	// VAT number of the Cessionario/Committente (customer). This is automatically set based on the recipient's VAT number.
	Committente NullableString `json:"committente,omitempty"`
	// VAT number of the Cedente/Prestatore (vendor). This is automatically set based on the sender's VAT number.
	Prestatore NullableString `json:"prestatore,omitempty"`
	// SDI identifier. This is set by the SDI and is guaranted to be unique within the SDI system.
	Identifier NullableString `json:"identifier,omitempty"`
	// Xml file name.
	FileName NullableString `json:"file_name,omitempty"`
	// SDI format (FPA12, FPR12, FSM10, ...)
	Format NullableString `json:"format,omitempty"`
	// Xml payloaad. This is the actual xml content, as string. On send, it can be base64 encoded. If it's not, it will be encoded before sending. It is guaranteed to be cyphered at rest.
	Payload NullableString `json:"payload,omitempty"`
	// Last update from SDI.
	LastUpdate NullableTime `json:"last_update,omitempty"`
	// When the invoice was sent to SDI.
	DateSent NullableTime `json:"date_sent,omitempty"`
	// The invoices included in the payload. This is set by the system, based on the xml content.
	Documents []DocumentData `json:"documents,omitempty"`
	// Whether the payload is Base64 encoded or a plain XML (text).
	Encoding *string `json:"encoding,omitempty"`
	// Wether the invoice has been read at least once.
	IsRead *bool `json:"is_read,omitempty"`
	// SDI message id.
	MessageId NullableString `json:"message_id,omitempty"`
}

// NewReceive instantiates a new Receive object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReceive() *Receive {
	this := Receive{}
	return &this
}

// NewReceiveWithDefaults instantiates a new Receive object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReceiveWithDefaults() *Receive {
	this := Receive{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Receive) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receive) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Receive) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Receive) SetId(v int32) {
	o.Id = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Receive) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receive) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Receive) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *Receive) SetCreated(v time.Time) {
	o.Created = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Receive) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receive) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Receive) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *Receive) SetVersion(v int32) {
	o.Version = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *Receive) GetUserId() int32 {
	if o == nil || IsNil(o.UserId) {
		var ret int32
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receive) GetUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *Receive) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int32 and assigns it to the UserId field.
func (o *Receive) SetUserId(v int32) {
	o.UserId = &v
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise.
func (o *Receive) GetCompanyId() int32 {
	if o == nil || IsNil(o.CompanyId) {
		var ret int32
		return ret
	}
	return *o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receive) GetCompanyIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CompanyId) {
		return nil, false
	}
	return o.CompanyId, true
}

// HasCompanyId returns a boolean if a field has been set.
func (o *Receive) HasCompanyId() bool {
	if o != nil && !IsNil(o.CompanyId) {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given int32 and assigns it to the CompanyId field.
func (o *Receive) SetCompanyId(v int32) {
	o.CompanyId = &v
}

// GetCommittente returns the Committente field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Receive) GetCommittente() string {
	if o == nil || IsNil(o.Committente.Get()) {
		var ret string
		return ret
	}
	return *o.Committente.Get()
}

// GetCommittenteOk returns a tuple with the Committente field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Receive) GetCommittenteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Committente.Get(), o.Committente.IsSet()
}

// HasCommittente returns a boolean if a field has been set.
func (o *Receive) HasCommittente() bool {
	if o != nil && o.Committente.IsSet() {
		return true
	}

	return false
}

// SetCommittente gets a reference to the given NullableString and assigns it to the Committente field.
func (o *Receive) SetCommittente(v string) {
	o.Committente.Set(&v)
}
// SetCommittenteNil sets the value for Committente to be an explicit nil
func (o *Receive) SetCommittenteNil() {
	o.Committente.Set(nil)
}

// UnsetCommittente ensures that no value is present for Committente, not even an explicit nil
func (o *Receive) UnsetCommittente() {
	o.Committente.Unset()
}

// GetPrestatore returns the Prestatore field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Receive) GetPrestatore() string {
	if o == nil || IsNil(o.Prestatore.Get()) {
		var ret string
		return ret
	}
	return *o.Prestatore.Get()
}

// GetPrestatoreOk returns a tuple with the Prestatore field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Receive) GetPrestatoreOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Prestatore.Get(), o.Prestatore.IsSet()
}

// HasPrestatore returns a boolean if a field has been set.
func (o *Receive) HasPrestatore() bool {
	if o != nil && o.Prestatore.IsSet() {
		return true
	}

	return false
}

// SetPrestatore gets a reference to the given NullableString and assigns it to the Prestatore field.
func (o *Receive) SetPrestatore(v string) {
	o.Prestatore.Set(&v)
}
// SetPrestatoreNil sets the value for Prestatore to be an explicit nil
func (o *Receive) SetPrestatoreNil() {
	o.Prestatore.Set(nil)
}

// UnsetPrestatore ensures that no value is present for Prestatore, not even an explicit nil
func (o *Receive) UnsetPrestatore() {
	o.Prestatore.Unset()
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Receive) GetIdentifier() string {
	if o == nil || IsNil(o.Identifier.Get()) {
		var ret string
		return ret
	}
	return *o.Identifier.Get()
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Receive) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Identifier.Get(), o.Identifier.IsSet()
}

// HasIdentifier returns a boolean if a field has been set.
func (o *Receive) HasIdentifier() bool {
	if o != nil && o.Identifier.IsSet() {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given NullableString and assigns it to the Identifier field.
func (o *Receive) SetIdentifier(v string) {
	o.Identifier.Set(&v)
}
// SetIdentifierNil sets the value for Identifier to be an explicit nil
func (o *Receive) SetIdentifierNil() {
	o.Identifier.Set(nil)
}

// UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
func (o *Receive) UnsetIdentifier() {
	o.Identifier.Unset()
}

// GetFileName returns the FileName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Receive) GetFileName() string {
	if o == nil || IsNil(o.FileName.Get()) {
		var ret string
		return ret
	}
	return *o.FileName.Get()
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Receive) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileName.Get(), o.FileName.IsSet()
}

// HasFileName returns a boolean if a field has been set.
func (o *Receive) HasFileName() bool {
	if o != nil && o.FileName.IsSet() {
		return true
	}

	return false
}

// SetFileName gets a reference to the given NullableString and assigns it to the FileName field.
func (o *Receive) SetFileName(v string) {
	o.FileName.Set(&v)
}
// SetFileNameNil sets the value for FileName to be an explicit nil
func (o *Receive) SetFileNameNil() {
	o.FileName.Set(nil)
}

// UnsetFileName ensures that no value is present for FileName, not even an explicit nil
func (o *Receive) UnsetFileName() {
	o.FileName.Unset()
}

// GetFormat returns the Format field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Receive) GetFormat() string {
	if o == nil || IsNil(o.Format.Get()) {
		var ret string
		return ret
	}
	return *o.Format.Get()
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Receive) GetFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Format.Get(), o.Format.IsSet()
}

// HasFormat returns a boolean if a field has been set.
func (o *Receive) HasFormat() bool {
	if o != nil && o.Format.IsSet() {
		return true
	}

	return false
}

// SetFormat gets a reference to the given NullableString and assigns it to the Format field.
func (o *Receive) SetFormat(v string) {
	o.Format.Set(&v)
}
// SetFormatNil sets the value for Format to be an explicit nil
func (o *Receive) SetFormatNil() {
	o.Format.Set(nil)
}

// UnsetFormat ensures that no value is present for Format, not even an explicit nil
func (o *Receive) UnsetFormat() {
	o.Format.Unset()
}

// GetPayload returns the Payload field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Receive) GetPayload() string {
	if o == nil || IsNil(o.Payload.Get()) {
		var ret string
		return ret
	}
	return *o.Payload.Get()
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Receive) GetPayloadOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Payload.Get(), o.Payload.IsSet()
}

// HasPayload returns a boolean if a field has been set.
func (o *Receive) HasPayload() bool {
	if o != nil && o.Payload.IsSet() {
		return true
	}

	return false
}

// SetPayload gets a reference to the given NullableString and assigns it to the Payload field.
func (o *Receive) SetPayload(v string) {
	o.Payload.Set(&v)
}
// SetPayloadNil sets the value for Payload to be an explicit nil
func (o *Receive) SetPayloadNil() {
	o.Payload.Set(nil)
}

// UnsetPayload ensures that no value is present for Payload, not even an explicit nil
func (o *Receive) UnsetPayload() {
	o.Payload.Unset()
}

// GetLastUpdate returns the LastUpdate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Receive) GetLastUpdate() time.Time {
	if o == nil || IsNil(o.LastUpdate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdate.Get()
}

// GetLastUpdateOk returns a tuple with the LastUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Receive) GetLastUpdateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdate.Get(), o.LastUpdate.IsSet()
}

// HasLastUpdate returns a boolean if a field has been set.
func (o *Receive) HasLastUpdate() bool {
	if o != nil && o.LastUpdate.IsSet() {
		return true
	}

	return false
}

// SetLastUpdate gets a reference to the given NullableTime and assigns it to the LastUpdate field.
func (o *Receive) SetLastUpdate(v time.Time) {
	o.LastUpdate.Set(&v)
}
// SetLastUpdateNil sets the value for LastUpdate to be an explicit nil
func (o *Receive) SetLastUpdateNil() {
	o.LastUpdate.Set(nil)
}

// UnsetLastUpdate ensures that no value is present for LastUpdate, not even an explicit nil
func (o *Receive) UnsetLastUpdate() {
	o.LastUpdate.Unset()
}

// GetDateSent returns the DateSent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Receive) GetDateSent() time.Time {
	if o == nil || IsNil(o.DateSent.Get()) {
		var ret time.Time
		return ret
	}
	return *o.DateSent.Get()
}

// GetDateSentOk returns a tuple with the DateSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Receive) GetDateSentOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateSent.Get(), o.DateSent.IsSet()
}

// HasDateSent returns a boolean if a field has been set.
func (o *Receive) HasDateSent() bool {
	if o != nil && o.DateSent.IsSet() {
		return true
	}

	return false
}

// SetDateSent gets a reference to the given NullableTime and assigns it to the DateSent field.
func (o *Receive) SetDateSent(v time.Time) {
	o.DateSent.Set(&v)
}
// SetDateSentNil sets the value for DateSent to be an explicit nil
func (o *Receive) SetDateSentNil() {
	o.DateSent.Set(nil)
}

// UnsetDateSent ensures that no value is present for DateSent, not even an explicit nil
func (o *Receive) UnsetDateSent() {
	o.DateSent.Unset()
}

// GetDocuments returns the Documents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Receive) GetDocuments() []DocumentData {
	if o == nil {
		var ret []DocumentData
		return ret
	}
	return o.Documents
}

// GetDocumentsOk returns a tuple with the Documents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Receive) GetDocumentsOk() ([]DocumentData, bool) {
	if o == nil || IsNil(o.Documents) {
		return nil, false
	}
	return o.Documents, true
}

// HasDocuments returns a boolean if a field has been set.
func (o *Receive) HasDocuments() bool {
	if o != nil && !IsNil(o.Documents) {
		return true
	}

	return false
}

// SetDocuments gets a reference to the given []DocumentData and assigns it to the Documents field.
func (o *Receive) SetDocuments(v []DocumentData) {
	o.Documents = v
}

// GetEncoding returns the Encoding field value if set, zero value otherwise.
func (o *Receive) GetEncoding() string {
	if o == nil || IsNil(o.Encoding) {
		var ret string
		return ret
	}
	return *o.Encoding
}

// GetEncodingOk returns a tuple with the Encoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receive) GetEncodingOk() (*string, bool) {
	if o == nil || IsNil(o.Encoding) {
		return nil, false
	}
	return o.Encoding, true
}

// HasEncoding returns a boolean if a field has been set.
func (o *Receive) HasEncoding() bool {
	if o != nil && !IsNil(o.Encoding) {
		return true
	}

	return false
}

// SetEncoding gets a reference to the given string and assigns it to the Encoding field.
func (o *Receive) SetEncoding(v string) {
	o.Encoding = &v
}

// GetIsRead returns the IsRead field value if set, zero value otherwise.
func (o *Receive) GetIsRead() bool {
	if o == nil || IsNil(o.IsRead) {
		var ret bool
		return ret
	}
	return *o.IsRead
}

// GetIsReadOk returns a tuple with the IsRead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receive) GetIsReadOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRead) {
		return nil, false
	}
	return o.IsRead, true
}

// HasIsRead returns a boolean if a field has been set.
func (o *Receive) HasIsRead() bool {
	if o != nil && !IsNil(o.IsRead) {
		return true
	}

	return false
}

// SetIsRead gets a reference to the given bool and assigns it to the IsRead field.
func (o *Receive) SetIsRead(v bool) {
	o.IsRead = &v
}

// GetMessageId returns the MessageId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Receive) GetMessageId() string {
	if o == nil || IsNil(o.MessageId.Get()) {
		var ret string
		return ret
	}
	return *o.MessageId.Get()
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Receive) GetMessageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MessageId.Get(), o.MessageId.IsSet()
}

// HasMessageId returns a boolean if a field has been set.
func (o *Receive) HasMessageId() bool {
	if o != nil && o.MessageId.IsSet() {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given NullableString and assigns it to the MessageId field.
func (o *Receive) SetMessageId(v string) {
	o.MessageId.Set(&v)
}
// SetMessageIdNil sets the value for MessageId to be an explicit nil
func (o *Receive) SetMessageIdNil() {
	o.MessageId.Set(nil)
}

// UnsetMessageId ensures that no value is present for MessageId, not even an explicit nil
func (o *Receive) UnsetMessageId() {
	o.MessageId.Unset()
}

func (o Receive) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Receive) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.CompanyId) {
		toSerialize["company_id"] = o.CompanyId
	}
	if o.Committente.IsSet() {
		toSerialize["committente"] = o.Committente.Get()
	}
	if o.Prestatore.IsSet() {
		toSerialize["prestatore"] = o.Prestatore.Get()
	}
	if o.Identifier.IsSet() {
		toSerialize["identifier"] = o.Identifier.Get()
	}
	if o.FileName.IsSet() {
		toSerialize["file_name"] = o.FileName.Get()
	}
	if o.Format.IsSet() {
		toSerialize["format"] = o.Format.Get()
	}
	if o.Payload.IsSet() {
		toSerialize["payload"] = o.Payload.Get()
	}
	if o.LastUpdate.IsSet() {
		toSerialize["last_update"] = o.LastUpdate.Get()
	}
	if o.DateSent.IsSet() {
		toSerialize["date_sent"] = o.DateSent.Get()
	}
	if o.Documents != nil {
		toSerialize["documents"] = o.Documents
	}
	if !IsNil(o.Encoding) {
		toSerialize["encoding"] = o.Encoding
	}
	if !IsNil(o.IsRead) {
		toSerialize["is_read"] = o.IsRead
	}
	if o.MessageId.IsSet() {
		toSerialize["message_id"] = o.MessageId.Get()
	}
	return toSerialize, nil
}

type NullableReceive struct {
	value *Receive
	isSet bool
}

func (v NullableReceive) Get() *Receive {
	return v.value
}

func (v *NullableReceive) Set(val *Receive) {
	v.value = val
	v.isSet = true
}

func (v NullableReceive) IsSet() bool {
	return v.isSet
}

func (v *NullableReceive) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReceive(val *Receive) *NullableReceive {
	return &NullableReceive{value: val, isSet: true}
}

func (v NullableReceive) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReceive) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


