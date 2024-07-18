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

// checks if the TerminalProductPrice type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TerminalProductPrice{}

// TerminalProductPrice struct for TerminalProductPrice
type TerminalProductPrice struct {
	// The three-character [ISO currency code](https://docs.adyen.com/development-resources/currency-codes).
	Currency *string `json:"currency,omitempty"`
	// The price of the item.
	Value *float64 `json:"value,omitempty"`
}

// NewTerminalProductPrice instantiates a new TerminalProductPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerminalProductPrice() *TerminalProductPrice {
	this := TerminalProductPrice{}
	return &this
}

// NewTerminalProductPriceWithDefaults instantiates a new TerminalProductPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerminalProductPriceWithDefaults() *TerminalProductPrice {
	this := TerminalProductPrice{}
	return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *TerminalProductPrice) GetCurrency() string {
	if o == nil || common.IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalProductPrice) GetCurrencyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *TerminalProductPrice) HasCurrency() bool {
	if o != nil && !common.IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *TerminalProductPrice) SetCurrency(v string) {
	o.Currency = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *TerminalProductPrice) GetValue() float64 {
	if o == nil || common.IsNil(o.Value) {
		var ret float64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalProductPrice) GetValueOk() (*float64, bool) {
	if o == nil || common.IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TerminalProductPrice) HasValue() bool {
	if o != nil && !common.IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float64 and assigns it to the Value field.
func (o *TerminalProductPrice) SetValue(v float64) {
	o.Value = &v
}

func (o TerminalProductPrice) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerminalProductPrice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !common.IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableTerminalProductPrice struct {
	value *TerminalProductPrice
	isSet bool
}

func (v NullableTerminalProductPrice) Get() *TerminalProductPrice {
	return v.value
}

func (v *NullableTerminalProductPrice) Set(val *TerminalProductPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminalProductPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminalProductPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminalProductPrice(val *TerminalProductPrice) *NullableTerminalProductPrice {
	return &NullableTerminalProductPrice{value: val, isSet: true}
}

func (v NullableTerminalProductPrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminalProductPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
