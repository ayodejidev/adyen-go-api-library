/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v12/src/common"
)

// checks if the TaxTotal type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TaxTotal{}

// TaxTotal struct for TaxTotal
type TaxTotal struct {
	Amount *Amount `json:"amount,omitempty"`
}

// NewTaxTotal instantiates a new TaxTotal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxTotal() *TaxTotal {
	this := TaxTotal{}
	return &this
}

// NewTaxTotalWithDefaults instantiates a new TaxTotal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxTotalWithDefaults() *TaxTotal {
	this := TaxTotal{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *TaxTotal) GetAmount() Amount {
	if o == nil || common.IsNil(o.Amount) {
		var ret Amount
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxTotal) GetAmountOk() (*Amount, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *TaxTotal) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given Amount and assigns it to the Amount field.
func (o *TaxTotal) SetAmount(v Amount) {
	o.Amount = &v
}

func (o TaxTotal) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaxTotal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	return toSerialize, nil
}

type NullableTaxTotal struct {
	value *TaxTotal
	isSet bool
}

func (v NullableTaxTotal) Get() *TaxTotal {
	return v.value
}

func (v *NullableTaxTotal) Set(val *TaxTotal) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxTotal) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxTotal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxTotal(val *TaxTotal) *NullableTaxTotal {
	return &NullableTaxTotal{value: val, isSet: true}
}

func (v NullableTaxTotal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxTotal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
