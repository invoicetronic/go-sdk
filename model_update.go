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

// checks if the Update type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Update{}

// Update struct for Update
type Update struct {
	// Unique identifier. Leave it at 0 for new records as it will be set automatically.
	Id *int32 `json:"id,omitempty"`
	// Creation date. It is set automatically.
	Created *time.Time `json:"created,omitempty"`
	// Row version, for optimistic concurrency. It is set automatically.
	Version *int32 `json:"version,omitempty"`
	// User id.
	UserId *int32 `json:"user_id,omitempty"`
	// Company id.
	CompanyId *int32 `json:"company_id,omitempty"`
	// Send id. This is the id of the sent invoice to which this update refers to.
	SendId *int32 `json:"send_id,omitempty"`
	// When the document was sent to the SDI.
	DateSent NullableTime `json:"date_sent,omitempty"`
	// Last update from SDI.
	LastUpdate *time.Time `json:"last_update,omitempty"`
	// SDI identifier. This is set by the SDI and it is unique within the SDI system.
	Identifier NullableString `json:"identifier,omitempty"`
	// State of the document. Theses are the possible values, as per the SDI documentation:
	State *string `json:"state,omitempty"`
	// Description for the state.
	Description NullableString `json:"description,omitempty"`
	// SDI message id.
	MessageId NullableString `json:"message_id,omitempty"`
	// SDI errors, if any.
	Errors []Error `json:"errors,omitempty"`
	// Wether the item has been read at least once.
	IsRead *bool `json:"is_read,omitempty"`
}

// NewUpdate instantiates a new Update object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdate() *Update {
	this := Update{}
	return &this
}

// NewUpdateWithDefaults instantiates a new Update object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateWithDefaults() *Update {
	this := Update{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Update) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Update) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Update) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Update) SetId(v int32) {
	o.Id = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Update) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Update) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Update) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *Update) SetCreated(v time.Time) {
	o.Created = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Update) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Update) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Update) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *Update) SetVersion(v int32) {
	o.Version = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *Update) GetUserId() int32 {
	if o == nil || IsNil(o.UserId) {
		var ret int32
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Update) GetUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *Update) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int32 and assigns it to the UserId field.
func (o *Update) SetUserId(v int32) {
	o.UserId = &v
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise.
func (o *Update) GetCompanyId() int32 {
	if o == nil || IsNil(o.CompanyId) {
		var ret int32
		return ret
	}
	return *o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Update) GetCompanyIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CompanyId) {
		return nil, false
	}
	return o.CompanyId, true
}

// HasCompanyId returns a boolean if a field has been set.
func (o *Update) HasCompanyId() bool {
	if o != nil && !IsNil(o.CompanyId) {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given int32 and assigns it to the CompanyId field.
func (o *Update) SetCompanyId(v int32) {
	o.CompanyId = &v
}

// GetSendId returns the SendId field value if set, zero value otherwise.
func (o *Update) GetSendId() int32 {
	if o == nil || IsNil(o.SendId) {
		var ret int32
		return ret
	}
	return *o.SendId
}

// GetSendIdOk returns a tuple with the SendId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Update) GetSendIdOk() (*int32, bool) {
	if o == nil || IsNil(o.SendId) {
		return nil, false
	}
	return o.SendId, true
}

// HasSendId returns a boolean if a field has been set.
func (o *Update) HasSendId() bool {
	if o != nil && !IsNil(o.SendId) {
		return true
	}

	return false
}

// SetSendId gets a reference to the given int32 and assigns it to the SendId field.
func (o *Update) SetSendId(v int32) {
	o.SendId = &v
}

// GetDateSent returns the DateSent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Update) GetDateSent() time.Time {
	if o == nil || IsNil(o.DateSent.Get()) {
		var ret time.Time
		return ret
	}
	return *o.DateSent.Get()
}

// GetDateSentOk returns a tuple with the DateSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Update) GetDateSentOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateSent.Get(), o.DateSent.IsSet()
}

// HasDateSent returns a boolean if a field has been set.
func (o *Update) HasDateSent() bool {
	if o != nil && o.DateSent.IsSet() {
		return true
	}

	return false
}

// SetDateSent gets a reference to the given NullableTime and assigns it to the DateSent field.
func (o *Update) SetDateSent(v time.Time) {
	o.DateSent.Set(&v)
}
// SetDateSentNil sets the value for DateSent to be an explicit nil
func (o *Update) SetDateSentNil() {
	o.DateSent.Set(nil)
}

// UnsetDateSent ensures that no value is present for DateSent, not even an explicit nil
func (o *Update) UnsetDateSent() {
	o.DateSent.Unset()
}

// GetLastUpdate returns the LastUpdate field value if set, zero value otherwise.
func (o *Update) GetLastUpdate() time.Time {
	if o == nil || IsNil(o.LastUpdate) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdate
}

// GetLastUpdateOk returns a tuple with the LastUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Update) GetLastUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdate) {
		return nil, false
	}
	return o.LastUpdate, true
}

// HasLastUpdate returns a boolean if a field has been set.
func (o *Update) HasLastUpdate() bool {
	if o != nil && !IsNil(o.LastUpdate) {
		return true
	}

	return false
}

// SetLastUpdate gets a reference to the given time.Time and assigns it to the LastUpdate field.
func (o *Update) SetLastUpdate(v time.Time) {
	o.LastUpdate = &v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Update) GetIdentifier() string {
	if o == nil || IsNil(o.Identifier.Get()) {
		var ret string
		return ret
	}
	return *o.Identifier.Get()
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Update) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Identifier.Get(), o.Identifier.IsSet()
}

// HasIdentifier returns a boolean if a field has been set.
func (o *Update) HasIdentifier() bool {
	if o != nil && o.Identifier.IsSet() {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given NullableString and assigns it to the Identifier field.
func (o *Update) SetIdentifier(v string) {
	o.Identifier.Set(&v)
}
// SetIdentifierNil sets the value for Identifier to be an explicit nil
func (o *Update) SetIdentifierNil() {
	o.Identifier.Set(nil)
}

// UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
func (o *Update) UnsetIdentifier() {
	o.Identifier.Unset()
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Update) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Update) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Update) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Update) SetState(v string) {
	o.State = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Update) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Update) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Update) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Update) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Update) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Update) UnsetDescription() {
	o.Description.Unset()
}

// GetMessageId returns the MessageId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Update) GetMessageId() string {
	if o == nil || IsNil(o.MessageId.Get()) {
		var ret string
		return ret
	}
	return *o.MessageId.Get()
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Update) GetMessageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MessageId.Get(), o.MessageId.IsSet()
}

// HasMessageId returns a boolean if a field has been set.
func (o *Update) HasMessageId() bool {
	if o != nil && o.MessageId.IsSet() {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given NullableString and assigns it to the MessageId field.
func (o *Update) SetMessageId(v string) {
	o.MessageId.Set(&v)
}
// SetMessageIdNil sets the value for MessageId to be an explicit nil
func (o *Update) SetMessageIdNil() {
	o.MessageId.Set(nil)
}

// UnsetMessageId ensures that no value is present for MessageId, not even an explicit nil
func (o *Update) UnsetMessageId() {
	o.MessageId.Unset()
}

// GetErrors returns the Errors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Update) GetErrors() []Error {
	if o == nil {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Update) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *Update) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *Update) SetErrors(v []Error) {
	o.Errors = v
}

// GetIsRead returns the IsRead field value if set, zero value otherwise.
func (o *Update) GetIsRead() bool {
	if o == nil || IsNil(o.IsRead) {
		var ret bool
		return ret
	}
	return *o.IsRead
}

// GetIsReadOk returns a tuple with the IsRead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Update) GetIsReadOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRead) {
		return nil, false
	}
	return o.IsRead, true
}

// HasIsRead returns a boolean if a field has been set.
func (o *Update) HasIsRead() bool {
	if o != nil && !IsNil(o.IsRead) {
		return true
	}

	return false
}

// SetIsRead gets a reference to the given bool and assigns it to the IsRead field.
func (o *Update) SetIsRead(v bool) {
	o.IsRead = &v
}

func (o Update) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Update) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.SendId) {
		toSerialize["send_id"] = o.SendId
	}
	if o.DateSent.IsSet() {
		toSerialize["date_sent"] = o.DateSent.Get()
	}
	if !IsNil(o.LastUpdate) {
		toSerialize["last_update"] = o.LastUpdate
	}
	if o.Identifier.IsSet() {
		toSerialize["identifier"] = o.Identifier.Get()
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.MessageId.IsSet() {
		toSerialize["message_id"] = o.MessageId.Get()
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.IsRead) {
		toSerialize["is_read"] = o.IsRead
	}
	return toSerialize, nil
}

type NullableUpdate struct {
	value *Update
	isSet bool
}

func (v NullableUpdate) Get() *Update {
	return v.value
}

func (v *NullableUpdate) Set(val *Update) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdate(val *Update) *NullableUpdate {
	return &NullableUpdate{value: val, isSet: true}
}

func (v NullableUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


