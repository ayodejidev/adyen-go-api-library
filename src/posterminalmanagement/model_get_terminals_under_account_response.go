/*
POS Terminal Management API

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package posterminalmanagement

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v14/src/common"
)

// checks if the GetTerminalsUnderAccountResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetTerminalsUnderAccountResponse{}

// GetTerminalsUnderAccountResponse struct for GetTerminalsUnderAccountResponse
type GetTerminalsUnderAccountResponse struct {
	// Your company account.
	CompanyAccount string `json:"companyAccount"`
	// Array that returns a list of all terminals that are in the inventory of the company account.
	InventoryTerminals []string `json:"inventoryTerminals,omitempty"`
	// Array that returns a list of all merchant accounts belonging to the company account.
	MerchantAccounts []MerchantAccount `json:"merchantAccounts,omitempty"`
}

// NewGetTerminalsUnderAccountResponse instantiates a new GetTerminalsUnderAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTerminalsUnderAccountResponse(companyAccount string) *GetTerminalsUnderAccountResponse {
	this := GetTerminalsUnderAccountResponse{}
	this.CompanyAccount = companyAccount
	return &this
}

// NewGetTerminalsUnderAccountResponseWithDefaults instantiates a new GetTerminalsUnderAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTerminalsUnderAccountResponseWithDefaults() *GetTerminalsUnderAccountResponse {
	this := GetTerminalsUnderAccountResponse{}
	return &this
}

// GetCompanyAccount returns the CompanyAccount field value
func (o *GetTerminalsUnderAccountResponse) GetCompanyAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CompanyAccount
}

// GetCompanyAccountOk returns a tuple with the CompanyAccount field value
// and a boolean to check if the value has been set.
func (o *GetTerminalsUnderAccountResponse) GetCompanyAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompanyAccount, true
}

// SetCompanyAccount sets field value
func (o *GetTerminalsUnderAccountResponse) SetCompanyAccount(v string) {
	o.CompanyAccount = v
}

// GetInventoryTerminals returns the InventoryTerminals field value if set, zero value otherwise.
func (o *GetTerminalsUnderAccountResponse) GetInventoryTerminals() []string {
	if o == nil || common.IsNil(o.InventoryTerminals) {
		var ret []string
		return ret
	}
	return o.InventoryTerminals
}

// GetInventoryTerminalsOk returns a tuple with the InventoryTerminals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTerminalsUnderAccountResponse) GetInventoryTerminalsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.InventoryTerminals) {
		return nil, false
	}
	return o.InventoryTerminals, true
}

// HasInventoryTerminals returns a boolean if a field has been set.
func (o *GetTerminalsUnderAccountResponse) HasInventoryTerminals() bool {
	if o != nil && !common.IsNil(o.InventoryTerminals) {
		return true
	}

	return false
}

// SetInventoryTerminals gets a reference to the given []string and assigns it to the InventoryTerminals field.
func (o *GetTerminalsUnderAccountResponse) SetInventoryTerminals(v []string) {
	o.InventoryTerminals = v
}

// GetMerchantAccounts returns the MerchantAccounts field value if set, zero value otherwise.
func (o *GetTerminalsUnderAccountResponse) GetMerchantAccounts() []MerchantAccount {
	if o == nil || common.IsNil(o.MerchantAccounts) {
		var ret []MerchantAccount
		return ret
	}
	return o.MerchantAccounts
}

// GetMerchantAccountsOk returns a tuple with the MerchantAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTerminalsUnderAccountResponse) GetMerchantAccountsOk() ([]MerchantAccount, bool) {
	if o == nil || common.IsNil(o.MerchantAccounts) {
		return nil, false
	}
	return o.MerchantAccounts, true
}

// HasMerchantAccounts returns a boolean if a field has been set.
func (o *GetTerminalsUnderAccountResponse) HasMerchantAccounts() bool {
	if o != nil && !common.IsNil(o.MerchantAccounts) {
		return true
	}

	return false
}

// SetMerchantAccounts gets a reference to the given []MerchantAccount and assigns it to the MerchantAccounts field.
func (o *GetTerminalsUnderAccountResponse) SetMerchantAccounts(v []MerchantAccount) {
	o.MerchantAccounts = v
}

func (o GetTerminalsUnderAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTerminalsUnderAccountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["companyAccount"] = o.CompanyAccount
	if !common.IsNil(o.InventoryTerminals) {
		toSerialize["inventoryTerminals"] = o.InventoryTerminals
	}
	if !common.IsNil(o.MerchantAccounts) {
		toSerialize["merchantAccounts"] = o.MerchantAccounts
	}
	return toSerialize, nil
}

type NullableGetTerminalsUnderAccountResponse struct {
	value *GetTerminalsUnderAccountResponse
	isSet bool
}

func (v NullableGetTerminalsUnderAccountResponse) Get() *GetTerminalsUnderAccountResponse {
	return v.value
}

func (v *NullableGetTerminalsUnderAccountResponse) Set(val *GetTerminalsUnderAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTerminalsUnderAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTerminalsUnderAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTerminalsUnderAccountResponse(val *GetTerminalsUnderAccountResponse) *NullableGetTerminalsUnderAccountResponse {
	return &NullableGetTerminalsUnderAccountResponse{value: val, isSet: true}
}

func (v NullableGetTerminalsUnderAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTerminalsUnderAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
