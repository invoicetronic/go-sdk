# SubKey

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

## Methods

### NewSubKey

`func NewSubKey() *SubKey`

NewSubKey instantiates a new SubKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubKeyWithDefaults

`func NewSubKeyWithDefaults() *SubKey`

NewSubKeyWithDefaults instantiates a new SubKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubKey) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubKey) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubKey) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SubKey) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreated

`func (o *SubKey) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SubKey) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SubKey) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *SubKey) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetVersion

`func (o *SubKey) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SubKey) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SubKey) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SubKey) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDescription

`func (o *SubKey) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SubKey) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SubKey) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SubKey) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SubKey) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SubKey) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetActive

`func (o *SubKey) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *SubKey) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *SubKey) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *SubKey) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetPermissions

`func (o *SubKey) GetPermissions() Permissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *SubKey) GetPermissionsOk() (*Permissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *SubKey) SetPermissions(v Permissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *SubKey) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetCompanyIds

`func (o *SubKey) GetCompanyIds() []int32`

GetCompanyIds returns the CompanyIds field if non-nil, zero value otherwise.

### GetCompanyIdsOk

`func (o *SubKey) GetCompanyIdsOk() (*[]int32, bool)`

GetCompanyIdsOk returns a tuple with the CompanyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyIds

`func (o *SubKey) SetCompanyIds(v []int32)`

SetCompanyIds sets CompanyIds field to given value.

### HasCompanyIds

`func (o *SubKey) HasCompanyIds() bool`

HasCompanyIds returns a boolean if a field has been set.

### SetCompanyIdsNil

`func (o *SubKey) SetCompanyIdsNil(b bool)`

 SetCompanyIdsNil sets the value for CompanyIds to be an explicit nil

### UnsetCompanyIds
`func (o *SubKey) UnsetCompanyIds()`

UnsetCompanyIds ensures that no value is present for CompanyIds, not even an explicit nil
### GetCorsOrigins

`func (o *SubKey) GetCorsOrigins() []string`

GetCorsOrigins returns the CorsOrigins field if non-nil, zero value otherwise.

### GetCorsOriginsOk

`func (o *SubKey) GetCorsOriginsOk() (*[]string, bool)`

GetCorsOriginsOk returns a tuple with the CorsOrigins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorsOrigins

`func (o *SubKey) SetCorsOrigins(v []string)`

SetCorsOrigins sets CorsOrigins field to given value.

### HasCorsOrigins

`func (o *SubKey) HasCorsOrigins() bool`

HasCorsOrigins returns a boolean if a field has been set.

### SetCorsOriginsNil

`func (o *SubKey) SetCorsOriginsNil(b bool)`

 SetCorsOriginsNil sets the value for CorsOrigins to be an explicit nil

### UnsetCorsOrigins
`func (o *SubKey) UnsetCorsOrigins()`

UnsetCorsOrigins ensures that no value is present for CorsOrigins, not even an explicit nil
### GetPreviousKeyExpiresAt

`func (o *SubKey) GetPreviousKeyExpiresAt() time.Time`

GetPreviousKeyExpiresAt returns the PreviousKeyExpiresAt field if non-nil, zero value otherwise.

### GetPreviousKeyExpiresAtOk

`func (o *SubKey) GetPreviousKeyExpiresAtOk() (*time.Time, bool)`

GetPreviousKeyExpiresAtOk returns a tuple with the PreviousKeyExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousKeyExpiresAt

`func (o *SubKey) SetPreviousKeyExpiresAt(v time.Time)`

SetPreviousKeyExpiresAt sets PreviousKeyExpiresAt field to given value.

### HasPreviousKeyExpiresAt

`func (o *SubKey) HasPreviousKeyExpiresAt() bool`

HasPreviousKeyExpiresAt returns a boolean if a field has been set.

### SetPreviousKeyExpiresAtNil

`func (o *SubKey) SetPreviousKeyExpiresAtNil(b bool)`

 SetPreviousKeyExpiresAtNil sets the value for PreviousKeyExpiresAt to be an explicit nil

### UnsetPreviousKeyExpiresAt
`func (o *SubKey) UnsetPreviousKeyExpiresAt()`

UnsetPreviousKeyExpiresAt ensures that no value is present for PreviousKeyExpiresAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


