# SubKeyWithSecrets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique identifier. | [optional] 
**Created** | Pointer to **time.Time** | Creation date. | [optional] [readonly] 
**Version** | Pointer to **int32** | Row version, for optimistic concurrency. | [optional] [readonly] 
**Description** | Pointer to **NullableString** | Human-readable label. | [optional] 
**Active** | Pointer to **bool** | Whether the key can authenticate. | [optional] 
**Permissions** | Pointer to [**Permissions**](Permissions.md) |  | [optional] 
**CompanyIds** | Pointer to **[]int32** | Companies the key can access. Empty means all the companies of the account. | [optional] 
**CorsOrigins** | Pointer to **[]string** | Browser origins allowed to call the API with this key (CORS). | [optional] 
**PreviousKeyExpiresAt** | Pointer to **NullableTime** | When the secrets replaced by the last roll stop working; null when there are none still valid. | [optional] 
**TestKey** | Pointer to **NullableString** | Sandbox secret. | [optional] 
**LiveKey** | Pointer to **NullableString** | Production secret. | [optional] 

## Methods

### NewSubKeyWithSecrets

`func NewSubKeyWithSecrets() *SubKeyWithSecrets`

NewSubKeyWithSecrets instantiates a new SubKeyWithSecrets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubKeyWithSecretsWithDefaults

`func NewSubKeyWithSecretsWithDefaults() *SubKeyWithSecrets`

NewSubKeyWithSecretsWithDefaults instantiates a new SubKeyWithSecrets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubKeyWithSecrets) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubKeyWithSecrets) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubKeyWithSecrets) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SubKeyWithSecrets) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreated

`func (o *SubKeyWithSecrets) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SubKeyWithSecrets) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SubKeyWithSecrets) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *SubKeyWithSecrets) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetVersion

`func (o *SubKeyWithSecrets) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SubKeyWithSecrets) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SubKeyWithSecrets) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SubKeyWithSecrets) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDescription

`func (o *SubKeyWithSecrets) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SubKeyWithSecrets) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SubKeyWithSecrets) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SubKeyWithSecrets) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SubKeyWithSecrets) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SubKeyWithSecrets) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetActive

`func (o *SubKeyWithSecrets) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *SubKeyWithSecrets) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *SubKeyWithSecrets) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *SubKeyWithSecrets) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetPermissions

`func (o *SubKeyWithSecrets) GetPermissions() Permissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *SubKeyWithSecrets) GetPermissionsOk() (*Permissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *SubKeyWithSecrets) SetPermissions(v Permissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *SubKeyWithSecrets) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetCompanyIds

`func (o *SubKeyWithSecrets) GetCompanyIds() []int32`

GetCompanyIds returns the CompanyIds field if non-nil, zero value otherwise.

### GetCompanyIdsOk

`func (o *SubKeyWithSecrets) GetCompanyIdsOk() (*[]int32, bool)`

GetCompanyIdsOk returns a tuple with the CompanyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyIds

`func (o *SubKeyWithSecrets) SetCompanyIds(v []int32)`

SetCompanyIds sets CompanyIds field to given value.

### HasCompanyIds

`func (o *SubKeyWithSecrets) HasCompanyIds() bool`

HasCompanyIds returns a boolean if a field has been set.

### SetCompanyIdsNil

`func (o *SubKeyWithSecrets) SetCompanyIdsNil(b bool)`

 SetCompanyIdsNil sets the value for CompanyIds to be an explicit nil

### UnsetCompanyIds
`func (o *SubKeyWithSecrets) UnsetCompanyIds()`

UnsetCompanyIds ensures that no value is present for CompanyIds, not even an explicit nil
### GetCorsOrigins

`func (o *SubKeyWithSecrets) GetCorsOrigins() []string`

GetCorsOrigins returns the CorsOrigins field if non-nil, zero value otherwise.

### GetCorsOriginsOk

`func (o *SubKeyWithSecrets) GetCorsOriginsOk() (*[]string, bool)`

GetCorsOriginsOk returns a tuple with the CorsOrigins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorsOrigins

`func (o *SubKeyWithSecrets) SetCorsOrigins(v []string)`

SetCorsOrigins sets CorsOrigins field to given value.

### HasCorsOrigins

`func (o *SubKeyWithSecrets) HasCorsOrigins() bool`

HasCorsOrigins returns a boolean if a field has been set.

### SetCorsOriginsNil

`func (o *SubKeyWithSecrets) SetCorsOriginsNil(b bool)`

 SetCorsOriginsNil sets the value for CorsOrigins to be an explicit nil

### UnsetCorsOrigins
`func (o *SubKeyWithSecrets) UnsetCorsOrigins()`

UnsetCorsOrigins ensures that no value is present for CorsOrigins, not even an explicit nil
### GetPreviousKeyExpiresAt

`func (o *SubKeyWithSecrets) GetPreviousKeyExpiresAt() time.Time`

GetPreviousKeyExpiresAt returns the PreviousKeyExpiresAt field if non-nil, zero value otherwise.

### GetPreviousKeyExpiresAtOk

`func (o *SubKeyWithSecrets) GetPreviousKeyExpiresAtOk() (*time.Time, bool)`

GetPreviousKeyExpiresAtOk returns a tuple with the PreviousKeyExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousKeyExpiresAt

`func (o *SubKeyWithSecrets) SetPreviousKeyExpiresAt(v time.Time)`

SetPreviousKeyExpiresAt sets PreviousKeyExpiresAt field to given value.

### HasPreviousKeyExpiresAt

`func (o *SubKeyWithSecrets) HasPreviousKeyExpiresAt() bool`

HasPreviousKeyExpiresAt returns a boolean if a field has been set.

### SetPreviousKeyExpiresAtNil

`func (o *SubKeyWithSecrets) SetPreviousKeyExpiresAtNil(b bool)`

 SetPreviousKeyExpiresAtNil sets the value for PreviousKeyExpiresAt to be an explicit nil

### UnsetPreviousKeyExpiresAt
`func (o *SubKeyWithSecrets) UnsetPreviousKeyExpiresAt()`

UnsetPreviousKeyExpiresAt ensures that no value is present for PreviousKeyExpiresAt, not even an explicit nil
### GetTestKey

`func (o *SubKeyWithSecrets) GetTestKey() string`

GetTestKey returns the TestKey field if non-nil, zero value otherwise.

### GetTestKeyOk

`func (o *SubKeyWithSecrets) GetTestKeyOk() (*string, bool)`

GetTestKeyOk returns a tuple with the TestKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestKey

`func (o *SubKeyWithSecrets) SetTestKey(v string)`

SetTestKey sets TestKey field to given value.

### HasTestKey

`func (o *SubKeyWithSecrets) HasTestKey() bool`

HasTestKey returns a boolean if a field has been set.

### SetTestKeyNil

`func (o *SubKeyWithSecrets) SetTestKeyNil(b bool)`

 SetTestKeyNil sets the value for TestKey to be an explicit nil

### UnsetTestKey
`func (o *SubKeyWithSecrets) UnsetTestKey()`

UnsetTestKey ensures that no value is present for TestKey, not even an explicit nil
### GetLiveKey

`func (o *SubKeyWithSecrets) GetLiveKey() string`

GetLiveKey returns the LiveKey field if non-nil, zero value otherwise.

### GetLiveKeyOk

`func (o *SubKeyWithSecrets) GetLiveKeyOk() (*string, bool)`

GetLiveKeyOk returns a tuple with the LiveKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveKey

`func (o *SubKeyWithSecrets) SetLiveKey(v string)`

SetLiveKey sets LiveKey field to given value.

### HasLiveKey

`func (o *SubKeyWithSecrets) HasLiveKey() bool`

HasLiveKey returns a boolean if a field has been set.

### SetLiveKeyNil

`func (o *SubKeyWithSecrets) SetLiveKeyNil(b bool)`

 SetLiveKeyNil sets the value for LiveKey to be an explicit nil

### UnsetLiveKey
`func (o *SubKeyWithSecrets) UnsetLiveKey()`

UnsetLiveKey ensures that no value is present for LiveKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


