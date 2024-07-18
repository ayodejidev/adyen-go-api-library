/*
POS Terminal Management API

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package posterminalmanagement

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v11/src/common"
)

// checks if the AssignTerminalsRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AssignTerminalsRequest{}

// AssignTerminalsRequest struct for AssignTerminalsRequest
type AssignTerminalsRequest struct {
	// Your company account. To return terminals to the company inventory, specify only this parameter and the `terminals`.
	CompanyAccount string `json:"companyAccount"`
	// Name of the merchant account. Specify this parameter to assign terminals to this merchant account or to a store under this merchant account.
	MerchantAccount *string `json:"merchantAccount,omitempty"`
	// Boolean that indicates if you are assigning the terminals to the merchant inventory. Do not use when assigning terminals to a store. Required when assigning the terminal to a merchant account.  - Set this to **true** to assign the terminals to the merchant inventory. This also means that the terminals cannot be boarded.  - Set this to **false** to assign the terminals to the merchant account as in-store terminals. This makes the terminals ready to be boarded and to process payments through the specified merchant account.
	MerchantInventory *bool `json:"merchantInventory,omitempty"`
	// The store code of the store that you want to assign the terminals to.
	Store *string `json:"store,omitempty"`
	// Array containing a list of terminal IDs that you want to assign or reassign to the merchant account or store, or that you want to return to the company inventory.  For example, `[\"V400m-324689776\",\"P400Plus-329127412\"]`.
	Terminals []string `json:"terminals"`
}

// NewAssignTerminalsRequest instantiates a new AssignTerminalsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssignTerminalsRequest(companyAccount string, terminals []string) *AssignTerminalsRequest {
	this := AssignTerminalsRequest{}
	this.CompanyAccount = companyAccount
	this.Terminals = terminals
	return &this
}

// NewAssignTerminalsRequestWithDefaults instantiates a new AssignTerminalsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssignTerminalsRequestWithDefaults() *AssignTerminalsRequest {
	this := AssignTerminalsRequest{}
	return &this
}

// GetCompanyAccount returns the CompanyAccount field value
func (o *AssignTerminalsRequest) GetCompanyAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CompanyAccount
}

// GetCompanyAccountOk returns a tuple with the CompanyAccount field value
// and a boolean to check if the value has been set.
func (o *AssignTerminalsRequest) GetCompanyAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompanyAccount, true
}

// SetCompanyAccount sets field value
func (o *AssignTerminalsRequest) SetCompanyAccount(v string) {
	o.CompanyAccount = v
}

// GetMerchantAccount returns the MerchantAccount field value if set, zero value otherwise.
func (o *AssignTerminalsRequest) GetMerchantAccount() string {
	if o == nil || common.IsNil(o.MerchantAccount) {
		var ret string
		return ret
	}
	return *o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignTerminalsRequest) GetMerchantAccountOk() (*string, bool) {
	if o == nil || common.IsNil(o.MerchantAccount) {
		return nil, false
	}
	return o.MerchantAccount, true
}

// HasMerchantAccount returns a boolean if a field has been set.
func (o *AssignTerminalsRequest) HasMerchantAccount() bool {
	if o != nil && !common.IsNil(o.MerchantAccount) {
		return true
	}

	return false
}

// SetMerchantAccount gets a reference to the given string and assigns it to the MerchantAccount field.
func (o *AssignTerminalsRequest) SetMerchantAccount(v string) {
	o.MerchantAccount = &v
}

// GetMerchantInventory returns the MerchantInventory field value if set, zero value otherwise.
func (o *AssignTerminalsRequest) GetMerchantInventory() bool {
	if o == nil || common.IsNil(o.MerchantInventory) {
		var ret bool
		return ret
	}
	return *o.MerchantInventory
}

// GetMerchantInventoryOk returns a tuple with the MerchantInventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignTerminalsRequest) GetMerchantInventoryOk() (*bool, bool) {
	if o == nil || common.IsNil(o.MerchantInventory) {
		return nil, false
	}
	return o.MerchantInventory, true
}

// HasMerchantInventory returns a boolean if a field has been set.
func (o *AssignTerminalsRequest) HasMerchantInventory() bool {
	if o != nil && !common.IsNil(o.MerchantInventory) {
		return true
	}

	return false
}

// SetMerchantInventory gets a reference to the given bool and assigns it to the MerchantInventory field.
func (o *AssignTerminalsRequest) SetMerchantInventory(v bool) {
	o.MerchantInventory = &v
}

// GetStore returns the Store field value if set, zero value otherwise.
func (o *AssignTerminalsRequest) GetStore() string {
	if o == nil || common.IsNil(o.Store) {
		var ret string
		return ret
	}
	return *o.Store
}

// GetStoreOk returns a tuple with the Store field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignTerminalsRequest) GetStoreOk() (*string, bool) {
	if o == nil || common.IsNil(o.Store) {
		return nil, false
	}
	return o.Store, true
}

// HasStore returns a boolean if a field has been set.
func (o *AssignTerminalsRequest) HasStore() bool {
	if o != nil && !common.IsNil(o.Store) {
		return true
	}

	return false
}

// SetStore gets a reference to the given string and assigns it to the Store field.
func (o *AssignTerminalsRequest) SetStore(v string) {
	o.Store = &v
}

// GetTerminals returns the Terminals field value
func (o *AssignTerminalsRequest) GetTerminals() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Terminals
}

// GetTerminalsOk returns a tuple with the Terminals field value
// and a boolean to check if the value has been set.
func (o *AssignTerminalsRequest) GetTerminalsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Terminals, true
}

// SetTerminals sets field value
func (o *AssignTerminalsRequest) SetTerminals(v []string) {
	o.Terminals = v
}

func (o AssignTerminalsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssignTerminalsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["companyAccount"] = o.CompanyAccount
	if !common.IsNil(o.MerchantAccount) {
		toSerialize["merchantAccount"] = o.MerchantAccount
	}
	if !common.IsNil(o.MerchantInventory) {
		toSerialize["merchantInventory"] = o.MerchantInventory
	}
	if !common.IsNil(o.Store) {
		toSerialize["store"] = o.Store
	}
	toSerialize["terminals"] = o.Terminals
	return toSerialize, nil
}

type NullableAssignTerminalsRequest struct {
	value *AssignTerminalsRequest
	isSet bool
}

func (v NullableAssignTerminalsRequest) Get() *AssignTerminalsRequest {
	return v.value
}

func (v *NullableAssignTerminalsRequest) Set(val *AssignTerminalsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAssignTerminalsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAssignTerminalsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssignTerminalsRequest(val *AssignTerminalsRequest) *NullableAssignTerminalsRequest {
	return &NullableAssignTerminalsRequest{value: val, isSet: true}
}

func (v NullableAssignTerminalsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssignTerminalsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
