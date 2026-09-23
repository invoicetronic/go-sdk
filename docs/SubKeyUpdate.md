# SubKeyUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Human-readable label, e.g. the name of the tenant the key is for. | 
**Active** | Pointer to **bool** | Whether the key can authenticate. Defaults to true. | [optional] 
**Permissions** | Pointer to [**Permissions**](Permissions.md) |  | [optional] 
**CompanyIds** | Pointer to **[]int32** | Companies the key can access. When omitted or empty, the key can access all the companies of the account, including the ones created later. | [optional] 
**CorsOrigins** | Pointer to **[]string** | Browser origins allowed to call the API with this key (CORS), e.g. &#x60;https://app.example.com&#x60; or &#x60;*.example.com&#x60;. A key used from a browser is public: keep its permissions and companies minimal. | [optional] 
**Id** | Pointer to **int32** | Id of the restricted key to update. | [optional] 
**Version** | Pointer to **int32** | Row version read with the key, for optimistic concurrency: a stale version fails with 422. | [optional] 

## Methods

### NewSubKeyUpdate

`func NewSubKeyUpdate(description string, ) *SubKeyUpdate`

NewSubKeyUpdate instantiates a new SubKeyUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubKeyUpdateWithDefaults

`func NewSubKeyUpdateWithDefaults() *SubKeyUpdate`

NewSubKeyUpdateWithDefaults instantiates a new SubKeyUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SubKeyUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SubKeyUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SubKeyUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetActive

`func (o *SubKeyUpdate) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *SubKeyUpdate) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *SubKeyUpdate) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *SubKeyUpdate) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetPermissions

`func (o *SubKeyUpdate) GetPermissions() Permissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *SubKeyUpdate) GetPermissionsOk() (*Permissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *SubKeyUpdate) SetPermissions(v Permissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *SubKeyUpdate) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetCompanyIds

`func (o *SubKeyUpdate) GetCompanyIds() []int32`

GetCompanyIds returns the CompanyIds field if non-nil, zero value otherwise.

### GetCompanyIdsOk

`func (o *SubKeyUpdate) GetCompanyIdsOk() (*[]int32, bool)`

GetCompanyIdsOk returns a tuple with the CompanyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyIds

`func (o *SubKeyUpdate) SetCompanyIds(v []int32)`

SetCompanyIds sets CompanyIds field to given value.

### HasCompanyIds

`func (o *SubKeyUpdate) HasCompanyIds() bool`

HasCompanyIds returns a boolean if a field has been set.

### SetCompanyIdsNil

`func (o *SubKeyUpdate) SetCompanyIdsNil(b bool)`

 SetCompanyIdsNil sets the value for CompanyIds to be an explicit nil

### UnsetCompanyIds
`func (o *SubKeyUpdate) UnsetCompanyIds()`

UnsetCompanyIds ensures that no value is present for CompanyIds, not even an explicit nil
### GetCorsOrigins

`func (o *SubKeyUpdate) GetCorsOrigins() []string`

GetCorsOrigins returns the CorsOrigins field if non-nil, zero value otherwise.

### GetCorsOriginsOk

`func (o *SubKeyUpdate) GetCorsOriginsOk() (*[]string, bool)`

GetCorsOriginsOk returns a tuple with the CorsOrigins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorsOrigins

`func (o *SubKeyUpdate) SetCorsOrigins(v []string)`

SetCorsOrigins sets CorsOrigins field to given value.

### HasCorsOrigins

`func (o *SubKeyUpdate) HasCorsOrigins() bool`

HasCorsOrigins returns a boolean if a field has been set.

### SetCorsOriginsNil

`func (o *SubKeyUpdate) SetCorsOriginsNil(b bool)`

 SetCorsOriginsNil sets the value for CorsOrigins to be an explicit nil

### UnsetCorsOrigins
`func (o *SubKeyUpdate) UnsetCorsOrigins()`

UnsetCorsOrigins ensures that no value is present for CorsOrigins, not even an explicit nil
### GetId

`func (o *SubKeyUpdate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubKeyUpdate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubKeyUpdate) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SubKeyUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVersion

`func (o *SubKeyUpdate) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SubKeyUpdate) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SubKeyUpdate) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SubKeyUpdate) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


