# AuthConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthnEnabled** | Pointer to **bool** | Whether incoming requests are authenticated by CMF. | [optional] 
**AuthzEnabled** | Pointer to **bool** | Whether CMF enforces fine-grained Flink authorization (RBAC). | [optional] 
**SsoEnabled** | Pointer to **bool** | Whether an OIDC/SSO login flow is available, advertised so the SPA shows the SSO login affordance. Set via cmf.ui.auth.sso-enabled; only meaningful when authnEnabled is true. | [optional] 
**BasicAuthEnabled** | Pointer to **bool** | Whether a Basic username/password login flow is available, advertised so the SPA shows the username/password form. Set via cmf.ui.auth.basic-auth-enabled; only meaningful when authnEnabled is true. | [optional] 
**TokenMaxLifeSeconds** | Pointer to **int64** | Maximum lifetime, in seconds, of a token issued by the login flow, used to schedule client-side refresh. Sourced from cmf.ui.auth.token-max-life-seconds; null until then. | [optional] 

## Methods

### NewAuthConfig

`func NewAuthConfig() *AuthConfig`

NewAuthConfig instantiates a new AuthConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthConfigWithDefaults

`func NewAuthConfigWithDefaults() *AuthConfig`

NewAuthConfigWithDefaults instantiates a new AuthConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthnEnabled

`func (o *AuthConfig) GetAuthnEnabled() bool`

GetAuthnEnabled returns the AuthnEnabled field if non-nil, zero value otherwise.

### GetAuthnEnabledOk

`func (o *AuthConfig) GetAuthnEnabledOk() (*bool, bool)`

GetAuthnEnabledOk returns a tuple with the AuthnEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthnEnabled

`func (o *AuthConfig) SetAuthnEnabled(v bool)`

SetAuthnEnabled sets AuthnEnabled field to given value.

### HasAuthnEnabled

`func (o *AuthConfig) HasAuthnEnabled() bool`

HasAuthnEnabled returns a boolean if a field has been set.

### GetAuthzEnabled

`func (o *AuthConfig) GetAuthzEnabled() bool`

GetAuthzEnabled returns the AuthzEnabled field if non-nil, zero value otherwise.

### GetAuthzEnabledOk

`func (o *AuthConfig) GetAuthzEnabledOk() (*bool, bool)`

GetAuthzEnabledOk returns a tuple with the AuthzEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthzEnabled

`func (o *AuthConfig) SetAuthzEnabled(v bool)`

SetAuthzEnabled sets AuthzEnabled field to given value.

### HasAuthzEnabled

`func (o *AuthConfig) HasAuthzEnabled() bool`

HasAuthzEnabled returns a boolean if a field has been set.

### GetSsoEnabled

`func (o *AuthConfig) GetSsoEnabled() bool`

GetSsoEnabled returns the SsoEnabled field if non-nil, zero value otherwise.

### GetSsoEnabledOk

`func (o *AuthConfig) GetSsoEnabledOk() (*bool, bool)`

GetSsoEnabledOk returns a tuple with the SsoEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoEnabled

`func (o *AuthConfig) SetSsoEnabled(v bool)`

SetSsoEnabled sets SsoEnabled field to given value.

### HasSsoEnabled

`func (o *AuthConfig) HasSsoEnabled() bool`

HasSsoEnabled returns a boolean if a field has been set.

### GetBasicAuthEnabled

`func (o *AuthConfig) GetBasicAuthEnabled() bool`

GetBasicAuthEnabled returns the BasicAuthEnabled field if non-nil, zero value otherwise.

### GetBasicAuthEnabledOk

`func (o *AuthConfig) GetBasicAuthEnabledOk() (*bool, bool)`

GetBasicAuthEnabledOk returns a tuple with the BasicAuthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthEnabled

`func (o *AuthConfig) SetBasicAuthEnabled(v bool)`

SetBasicAuthEnabled sets BasicAuthEnabled field to given value.

### HasBasicAuthEnabled

`func (o *AuthConfig) HasBasicAuthEnabled() bool`

HasBasicAuthEnabled returns a boolean if a field has been set.

### GetTokenMaxLifeSeconds

`func (o *AuthConfig) GetTokenMaxLifeSeconds() int64`

GetTokenMaxLifeSeconds returns the TokenMaxLifeSeconds field if non-nil, zero value otherwise.

### GetTokenMaxLifeSecondsOk

`func (o *AuthConfig) GetTokenMaxLifeSecondsOk() (*int64, bool)`

GetTokenMaxLifeSecondsOk returns a tuple with the TokenMaxLifeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenMaxLifeSeconds

`func (o *AuthConfig) SetTokenMaxLifeSeconds(v int64)`

SetTokenMaxLifeSeconds sets TokenMaxLifeSeconds field to given value.

### HasTokenMaxLifeSeconds

`func (o *AuthConfig) HasTokenMaxLifeSeconds() bool`

HasTokenMaxLifeSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


