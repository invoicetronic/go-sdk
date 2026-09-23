# Permissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Company** | Pointer to **NullableString** | Companies: &#x60;Read&#x60; lists and reads them, &#x60;Write&#x60; also creates, updates and deletes them. | [optional] 
**Send** | Pointer to **NullableString** | Outgoing invoices: &#x60;Read&#x60; lists and reads them, &#x60;Write&#x60; also sends and validates invoices. | [optional] 
**Receive** | Pointer to **NullableString** | Incoming invoices: &#x60;Read&#x60; lists and reads them, &#x60;Write&#x60; also deletes them. | [optional] 
**Webhook** | Pointer to **NullableString** | Webhooks: &#x60;Read&#x60; lists and reads them, &#x60;Write&#x60; also creates, updates and deletes them. | [optional] 
**Update** | Pointer to **NullableString** | SDI status updates of outgoing invoices. | [optional] 
**Log** | Pointer to **NullableString** | Event log. | [optional] 
**Webhookhistory** | Pointer to **NullableString** | Webhook delivery history. | [optional] 
**Export** | Pointer to **NullableString** | Invoice export. | [optional] 
**Status** | Pointer to **NullableString** | Account status (remaining operations and signatures). | [optional] 

## Methods

### NewPermissions

`func NewPermissions() *Permissions`

NewPermissions instantiates a new Permissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionsWithDefaults

`func NewPermissionsWithDefaults() *Permissions`

NewPermissionsWithDefaults instantiates a new Permissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompany

`func (o *Permissions) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *Permissions) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *Permissions) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *Permissions) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *Permissions) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *Permissions) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetSend

`func (o *Permissions) GetSend() string`

GetSend returns the Send field if non-nil, zero value otherwise.

### GetSendOk

`func (o *Permissions) GetSendOk() (*string, bool)`

GetSendOk returns a tuple with the Send field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSend

`func (o *Permissions) SetSend(v string)`

SetSend sets Send field to given value.

### HasSend

`func (o *Permissions) HasSend() bool`

HasSend returns a boolean if a field has been set.

### SetSendNil

`func (o *Permissions) SetSendNil(b bool)`

 SetSendNil sets the value for Send to be an explicit nil

### UnsetSend
`func (o *Permissions) UnsetSend()`

UnsetSend ensures that no value is present for Send, not even an explicit nil
### GetReceive

`func (o *Permissions) GetReceive() string`

GetReceive returns the Receive field if non-nil, zero value otherwise.

### GetReceiveOk

`func (o *Permissions) GetReceiveOk() (*string, bool)`

GetReceiveOk returns a tuple with the Receive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceive

`func (o *Permissions) SetReceive(v string)`

SetReceive sets Receive field to given value.

### HasReceive

`func (o *Permissions) HasReceive() bool`

HasReceive returns a boolean if a field has been set.

### SetReceiveNil

`func (o *Permissions) SetReceiveNil(b bool)`

 SetReceiveNil sets the value for Receive to be an explicit nil

### UnsetReceive
`func (o *Permissions) UnsetReceive()`

UnsetReceive ensures that no value is present for Receive, not even an explicit nil
### GetWebhook

`func (o *Permissions) GetWebhook() string`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *Permissions) GetWebhookOk() (*string, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *Permissions) SetWebhook(v string)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *Permissions) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.

### SetWebhookNil

`func (o *Permissions) SetWebhookNil(b bool)`

 SetWebhookNil sets the value for Webhook to be an explicit nil

### UnsetWebhook
`func (o *Permissions) UnsetWebhook()`

UnsetWebhook ensures that no value is present for Webhook, not even an explicit nil
### GetUpdate

`func (o *Permissions) GetUpdate() string`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *Permissions) GetUpdateOk() (*string, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *Permissions) SetUpdate(v string)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *Permissions) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### SetUpdateNil

`func (o *Permissions) SetUpdateNil(b bool)`

 SetUpdateNil sets the value for Update to be an explicit nil

### UnsetUpdate
`func (o *Permissions) UnsetUpdate()`

UnsetUpdate ensures that no value is present for Update, not even an explicit nil
### GetLog

`func (o *Permissions) GetLog() string`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *Permissions) GetLogOk() (*string, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *Permissions) SetLog(v string)`

SetLog sets Log field to given value.

### HasLog

`func (o *Permissions) HasLog() bool`

HasLog returns a boolean if a field has been set.

### SetLogNil

`func (o *Permissions) SetLogNil(b bool)`

 SetLogNil sets the value for Log to be an explicit nil

### UnsetLog
`func (o *Permissions) UnsetLog()`

UnsetLog ensures that no value is present for Log, not even an explicit nil
### GetWebhookhistory

`func (o *Permissions) GetWebhookhistory() string`

GetWebhookhistory returns the Webhookhistory field if non-nil, zero value otherwise.

### GetWebhookhistoryOk

`func (o *Permissions) GetWebhookhistoryOk() (*string, bool)`

GetWebhookhistoryOk returns a tuple with the Webhookhistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookhistory

`func (o *Permissions) SetWebhookhistory(v string)`

SetWebhookhistory sets Webhookhistory field to given value.

### HasWebhookhistory

`func (o *Permissions) HasWebhookhistory() bool`

HasWebhookhistory returns a boolean if a field has been set.

### SetWebhookhistoryNil

`func (o *Permissions) SetWebhookhistoryNil(b bool)`

 SetWebhookhistoryNil sets the value for Webhookhistory to be an explicit nil

### UnsetWebhookhistory
`func (o *Permissions) UnsetWebhookhistory()`

UnsetWebhookhistory ensures that no value is present for Webhookhistory, not even an explicit nil
### GetExport

`func (o *Permissions) GetExport() string`

GetExport returns the Export field if non-nil, zero value otherwise.

### GetExportOk

`func (o *Permissions) GetExportOk() (*string, bool)`

GetExportOk returns a tuple with the Export field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExport

`func (o *Permissions) SetExport(v string)`

SetExport sets Export field to given value.

### HasExport

`func (o *Permissions) HasExport() bool`

HasExport returns a boolean if a field has been set.

### SetExportNil

`func (o *Permissions) SetExportNil(b bool)`

 SetExportNil sets the value for Export to be an explicit nil

### UnsetExport
`func (o *Permissions) UnsetExport()`

UnsetExport ensures that no value is present for Export, not even an explicit nil
### GetStatus

`func (o *Permissions) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Permissions) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Permissions) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Permissions) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *Permissions) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *Permissions) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


