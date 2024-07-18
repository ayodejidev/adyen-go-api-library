/*
Transfers API

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transfers

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v11/src/common"
)

// checks if the CapitalBalance type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CapitalBalance{}

// CapitalBalance struct for CapitalBalance
type CapitalBalance struct {
	// The three-character [ISO currency code](https://docs.adyen.com/development-resources/currency-codes).
	Currency string `json:"currency"`
	// Fee amount.
	Fee int64 `json:"fee"`
	// Principal amount.
	Principal int64 `json:"principal"`
	// Total amount. A sum of principal amount and fee amount.
	Total int64 `json:"total"`
}

// NewCapitalBalance instantiates a new CapitalBalance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapitalBalance(currency string, fee int64, principal int64, total int64) *CapitalBalance {
	this := CapitalBalance{}
	this.Currency = currency
	this.Fee = fee
	this.Principal = principal
	this.Total = total
	return &this
}

// NewCapitalBalanceWithDefaults instantiates a new CapitalBalance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapitalBalanceWithDefaults() *CapitalBalance {
	this := CapitalBalance{}
	return &this
}

// GetCurrency returns the Currency field value
func (o *CapitalBalance) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *CapitalBalance) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *CapitalBalance) SetCurrency(v string) {
	o.Currency = v
}

// GetFee returns the Fee field value
func (o *CapitalBalance) GetFee() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Fee
}

// GetFeeOk returns a tuple with the Fee field value
// and a boolean to check if the value has been set.
func (o *CapitalBalance) GetFeeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fee, true
}

// SetFee sets field value
func (o *CapitalBalance) SetFee(v int64) {
	o.Fee = v
}

// GetPrincipal returns the Principal field value
func (o *CapitalBalance) GetPrincipal() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field value
// and a boolean to check if the value has been set.
func (o *CapitalBalance) GetPrincipalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Principal, true
}

// SetPrincipal sets field value
func (o *CapitalBalance) SetPrincipal(v int64) {
	o.Principal = v
}

// GetTotal returns the Total field value
func (o *CapitalBalance) GetTotal() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *CapitalBalance) GetTotalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *CapitalBalance) SetTotal(v int64) {
	o.Total = v
}

func (o CapitalBalance) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CapitalBalance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["currency"] = o.Currency
	toSerialize["fee"] = o.Fee
	toSerialize["principal"] = o.Principal
	toSerialize["total"] = o.Total
	return toSerialize, nil
}

type NullableCapitalBalance struct {
	value *CapitalBalance
	isSet bool
}

func (v NullableCapitalBalance) Get() *CapitalBalance {
	return v.value
}

func (v *NullableCapitalBalance) Set(val *CapitalBalance) {
	v.value = val
	v.isSet = true
}

func (v NullableCapitalBalance) IsSet() bool {
	return v.isSet
}

func (v *NullableCapitalBalance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapitalBalance(val *CapitalBalance) *NullableCapitalBalance {
	return &NullableCapitalBalance{value: val, isSet: true}
}

func (v NullableCapitalBalance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapitalBalance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
