# WhoAmi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Principal** | Pointer to **NullableString** | The authenticated principal name -- the JWT subject for token/OIDC/Basic logins, or the mTLS client-certificate subject mapped to a principal. Null when the request is not authenticated (authentication disabled). | [optional] 
**AuthMethod** | Pointer to **NullableString** | How the request was authenticated: \&quot;certificate\&quot; when the identity came from the TLS client certificate, \&quot;token\&quot; when it came from an HTTP credential (a bearer token, a session cookie, or basic credentials). Null when the request is not authenticated. In deployments that combine mTLS with token authentication, a request carrying a token authenticates as the token&#39;s principal and reports \&quot;token\&quot;; the client certificate then only secures the transport. | [optional] 

## Methods

### NewWhoAmi

`func NewWhoAmi() *WhoAmi`

NewWhoAmi instantiates a new WhoAmi object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhoAmiWithDefaults

`func NewWhoAmiWithDefaults() *WhoAmi`

NewWhoAmiWithDefaults instantiates a new WhoAmi object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrincipal

`func (o *WhoAmi) GetPrincipal() string`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *WhoAmi) GetPrincipalOk() (*string, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *WhoAmi) SetPrincipal(v string)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *WhoAmi) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### SetPrincipalNil

`func (o *WhoAmi) SetPrincipalNil(b bool)`

 SetPrincipalNil sets the value for Principal to be an explicit nil

### UnsetPrincipal
`func (o *WhoAmi) UnsetPrincipal()`

UnsetPrincipal ensures that no value is present for Principal, not even an explicit nil
### GetAuthMethod

`func (o *WhoAmi) GetAuthMethod() string`

GetAuthMethod returns the AuthMethod field if non-nil, zero value otherwise.

### GetAuthMethodOk

`func (o *WhoAmi) GetAuthMethodOk() (*string, bool)`

GetAuthMethodOk returns a tuple with the AuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethod

`func (o *WhoAmi) SetAuthMethod(v string)`

SetAuthMethod sets AuthMethod field to given value.

### HasAuthMethod

`func (o *WhoAmi) HasAuthMethod() bool`

HasAuthMethod returns a boolean if a field has been set.

### SetAuthMethodNil

`func (o *WhoAmi) SetAuthMethodNil(b bool)`

 SetAuthMethodNil sets the value for AuthMethod to be an explicit nil

### UnsetAuthMethod
`func (o *WhoAmi) UnsetAuthMethod()`

UnsetAuthMethod ensures that no value is present for AuthMethod, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


