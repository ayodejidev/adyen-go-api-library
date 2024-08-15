/*
POS Terminal Management API

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package posterminalmanagement

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v12/src/common"
)

// checks if the GetTerminalsUnderAccountRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetTerminalsUnderAccountRequest{}

// GetTerminalsUnderAccountRequest struct for GetTerminalsUnderAccountRequest
type GetTerminalsUnderAccountRequest struct {
	// Your company account. If you only specify this parameter, the response includes all terminals at all account levels.
	CompanyAccount string `json:"companyAccount"`
	// The merchant account. This is required if you are retrieving the terminals assigned to a store.If you don't specify a `store` the response includes the terminals assigned to the specified merchant account and the terminals assigned to the stores under this merchant account.
	MerchantAccount *string `json:"merchantAccount,omitempty"`
	// The store code of the store. With this parameter, the response only includes the terminals assigned to the specified store.
	Store *string `json:"store,omitempty"`
}

// NewGetTerminalsUnderAccountRequest instantiates a new GetTerminalsUnderAccountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTerminalsUnderAccountRequest(companyAccount string) *GetTerminalsUnderAccountRequest {
	this := GetTerminalsUnderAccountRequest{}
	this.CompanyAccount = companyAccount
	return &this
}

// NewGetTerminalsUnderAccountRequestWithDefaults instantiates a new GetTerminalsUnderAccountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTerminalsUnderAccountRequestWithDefaults() *GetTerminalsUnderAccountRequest {
	this := GetTerminalsUnderAccountRequest{}
	return &this
}

// GetCompanyAccount returns the CompanyAccount field value
func (o *GetTerminalsUnderAccountRequest) GetCompanyAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CompanyAccount
}

// GetCompanyAccountOk returns a tuple with the CompanyAccount field value
// and a boolean to check if the value has been set.
func (o *GetTerminalsUnderAccountRequest) GetCompanyAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompanyAccount, true
}

// SetCompanyAccount sets field value
func (o *GetTerminalsUnderAccountRequest) SetCompanyAccount(v string) {
	o.CompanyAccount = v
}

// GetMerchantAccount returns the MerchantAccount field value if set, zero value otherwise.
func (o *GetTerminalsUnderAccountRequest) GetMerchantAccount() string {
	if o == nil || common.IsNil(o.MerchantAccount) {
		var ret string
		return ret
	}
	return *o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTerminalsUnderAccountRequest) GetMerchantAccountOk() (*string, bool) {
	if o == nil || common.IsNil(o.MerchantAccount) {
		return nil, false
	}
	return o.MerchantAccount, true
}

// HasMerchantAccount returns a boolean if a field has been set.
func (o *GetTerminalsUnderAccountRequest) HasMerchantAccount() bool {
	if o != nil && !common.IsNil(o.MerchantAccount) {
		return true
	}

	return false
}

// SetMerchantAccount gets a reference to the given string and assigns it to the MerchantAccount field.
func (o *GetTerminalsUnderAccountRequest) SetMerchantAccount(v string) {
	o.MerchantAccount = &v
}

// GetStore returns the Store field value if set, zero value otherwise.
func (o *GetTerminalsUnderAccountRequest) GetStore() string {
	if o == nil || common.IsNil(o.Store) {
		var ret string
		return ret
	}
	return *o.Store
}

// GetStoreOk returns a tuple with the Store field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTerminalsUnderAccountRequest) GetStoreOk() (*string, bool) {
	if o == nil || common.IsNil(o.Store) {
		return nil, false
	}
	return o.Store, true
}

// HasStore returns a boolean if a field has been set.
func (o *GetTerminalsUnderAccountRequest) HasStore() bool {
	if o != nil && !common.IsNil(o.Store) {
		return true
	}

	return false
}

// SetStore gets a reference to the given string and assigns it to the Store field.
func (o *GetTerminalsUnderAccountRequest) SetStore(v string) {
	o.Store = &v
}

func (o GetTerminalsUnderAccountRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTerminalsUnderAccountRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["companyAccount"] = o.CompanyAccount
	if !common.IsNil(o.MerchantAccount) {
		toSerialize["merchantAccount"] = o.MerchantAccount
	}
	if !common.IsNil(o.Store) {
		toSerialize["store"] = o.Store
	}
	return toSerialize, nil
}

type NullableGetTerminalsUnderAccountRequest struct {
	value *GetTerminalsUnderAccountRequest
	isSet bool
}

func (v NullableGetTerminalsUnderAccountRequest) Get() *GetTerminalsUnderAccountRequest {
	return v.value
}

func (v *NullableGetTerminalsUnderAccountRequest) Set(val *GetTerminalsUnderAccountRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTerminalsUnderAccountRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTerminalsUnderAccountRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTerminalsUnderAccountRequest(val *GetTerminalsUnderAccountRequest) *NullableGetTerminalsUnderAccountRequest {
	return &NullableGetTerminalsUnderAccountRequest{value: val, isSet: true}
}

func (v NullableGetTerminalsUnderAccountRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTerminalsUnderAccountRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
