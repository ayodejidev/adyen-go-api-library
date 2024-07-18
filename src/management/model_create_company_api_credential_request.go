/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v11/src/common"
)

// checks if the CreateCompanyApiCredentialRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CreateCompanyApiCredentialRequest{}

// CreateCompanyApiCredentialRequest struct for CreateCompanyApiCredentialRequest
type CreateCompanyApiCredentialRequest struct {
	// List of [allowed origins](https://docs.adyen.com/development-resources/client-side-authentication#allowed-origins) for the new API credential.
	AllowedOrigins []string `json:"allowedOrigins,omitempty"`
	// List of merchant accounts that the API credential has access to.
	AssociatedMerchantAccounts []string `json:"associatedMerchantAccounts,omitempty"`
	// Description of the API credential.
	Description *string `json:"description,omitempty"`
	// List of [roles](https://docs.adyen.com/development-resources/api-credentials#roles-1) for the API credential. Only roles assigned to 'ws@Company.<CompanyName>' can be assigned to other API credentials.
	Roles []string `json:"roles,omitempty"`
}

// NewCreateCompanyApiCredentialRequest instantiates a new CreateCompanyApiCredentialRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCompanyApiCredentialRequest() *CreateCompanyApiCredentialRequest {
	this := CreateCompanyApiCredentialRequest{}
	return &this
}

// NewCreateCompanyApiCredentialRequestWithDefaults instantiates a new CreateCompanyApiCredentialRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCompanyApiCredentialRequestWithDefaults() *CreateCompanyApiCredentialRequest {
	this := CreateCompanyApiCredentialRequest{}
	return &this
}

// GetAllowedOrigins returns the AllowedOrigins field value if set, zero value otherwise.
func (o *CreateCompanyApiCredentialRequest) GetAllowedOrigins() []string {
	if o == nil || common.IsNil(o.AllowedOrigins) {
		var ret []string
		return ret
	}
	return o.AllowedOrigins
}

// GetAllowedOriginsOk returns a tuple with the AllowedOrigins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCompanyApiCredentialRequest) GetAllowedOriginsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.AllowedOrigins) {
		return nil, false
	}
	return o.AllowedOrigins, true
}

// HasAllowedOrigins returns a boolean if a field has been set.
func (o *CreateCompanyApiCredentialRequest) HasAllowedOrigins() bool {
	if o != nil && !common.IsNil(o.AllowedOrigins) {
		return true
	}

	return false
}

// SetAllowedOrigins gets a reference to the given []string and assigns it to the AllowedOrigins field.
func (o *CreateCompanyApiCredentialRequest) SetAllowedOrigins(v []string) {
	o.AllowedOrigins = v
}

// GetAssociatedMerchantAccounts returns the AssociatedMerchantAccounts field value if set, zero value otherwise.
func (o *CreateCompanyApiCredentialRequest) GetAssociatedMerchantAccounts() []string {
	if o == nil || common.IsNil(o.AssociatedMerchantAccounts) {
		var ret []string
		return ret
	}
	return o.AssociatedMerchantAccounts
}

// GetAssociatedMerchantAccountsOk returns a tuple with the AssociatedMerchantAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCompanyApiCredentialRequest) GetAssociatedMerchantAccountsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.AssociatedMerchantAccounts) {
		return nil, false
	}
	return o.AssociatedMerchantAccounts, true
}

// HasAssociatedMerchantAccounts returns a boolean if a field has been set.
func (o *CreateCompanyApiCredentialRequest) HasAssociatedMerchantAccounts() bool {
	if o != nil && !common.IsNil(o.AssociatedMerchantAccounts) {
		return true
	}

	return false
}

// SetAssociatedMerchantAccounts gets a reference to the given []string and assigns it to the AssociatedMerchantAccounts field.
func (o *CreateCompanyApiCredentialRequest) SetAssociatedMerchantAccounts(v []string) {
	o.AssociatedMerchantAccounts = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateCompanyApiCredentialRequest) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCompanyApiCredentialRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateCompanyApiCredentialRequest) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateCompanyApiCredentialRequest) SetDescription(v string) {
	o.Description = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *CreateCompanyApiCredentialRequest) GetRoles() []string {
	if o == nil || common.IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCompanyApiCredentialRequest) GetRolesOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *CreateCompanyApiCredentialRequest) HasRoles() bool {
	if o != nil && !common.IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *CreateCompanyApiCredentialRequest) SetRoles(v []string) {
	o.Roles = v
}

func (o CreateCompanyApiCredentialRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCompanyApiCredentialRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AllowedOrigins) {
		toSerialize["allowedOrigins"] = o.AllowedOrigins
	}
	if !common.IsNil(o.AssociatedMerchantAccounts) {
		toSerialize["associatedMerchantAccounts"] = o.AssociatedMerchantAccounts
	}
	if !common.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !common.IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	return toSerialize, nil
}

type NullableCreateCompanyApiCredentialRequest struct {
	value *CreateCompanyApiCredentialRequest
	isSet bool
}

func (v NullableCreateCompanyApiCredentialRequest) Get() *CreateCompanyApiCredentialRequest {
	return v.value
}

func (v *NullableCreateCompanyApiCredentialRequest) Set(val *CreateCompanyApiCredentialRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCompanyApiCredentialRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCompanyApiCredentialRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCompanyApiCredentialRequest(val *CreateCompanyApiCredentialRequest) *NullableCreateCompanyApiCredentialRequest {
	return &NullableCreateCompanyApiCredentialRequest{value: val, isSet: true}
}

func (v NullableCreateCompanyApiCredentialRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCompanyApiCredentialRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
