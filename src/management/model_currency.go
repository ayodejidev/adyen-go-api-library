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

// checks if the Currency type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Currency{}

// Currency struct for Currency
type Currency struct {
	// Surcharge amount per transaction, in [minor units](https://docs.adyen.com/development-resources/currency-codes).
	Amount *int32 `json:"amount,omitempty"`
	// Three-character [ISO currency code](https://docs.adyen.com/development-resources/currency-codes). For example, **AUD**.
	CurrencyCode string `json:"currencyCode"`
	// Surcharge percentage per transaction. The maximum number of decimal places is two. For example, **1%** or **2.27%**.
	Percentage *float64 `json:"percentage,omitempty"`
}

// NewCurrency instantiates a new Currency object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCurrency(currencyCode string) *Currency {
	this := Currency{}
	this.CurrencyCode = currencyCode
	return &this
}

// NewCurrencyWithDefaults instantiates a new Currency object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCurrencyWithDefaults() *Currency {
	this := Currency{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *Currency) GetAmount() int32 {
	if o == nil || common.IsNil(o.Amount) {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Currency) GetAmountOk() (*int32, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *Currency) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *Currency) SetAmount(v int32) {
	o.Amount = &v
}

// GetCurrencyCode returns the CurrencyCode field value
func (o *Currency) GetCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value
// and a boolean to check if the value has been set.
func (o *Currency) GetCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// SetCurrencyCode sets field value
func (o *Currency) SetCurrencyCode(v string) {
	o.CurrencyCode = v
}

// GetPercentage returns the Percentage field value if set, zero value otherwise.
func (o *Currency) GetPercentage() float64 {
	if o == nil || common.IsNil(o.Percentage) {
		var ret float64
		return ret
	}
	return *o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Currency) GetPercentageOk() (*float64, bool) {
	if o == nil || common.IsNil(o.Percentage) {
		return nil, false
	}
	return o.Percentage, true
}

// HasPercentage returns a boolean if a field has been set.
func (o *Currency) HasPercentage() bool {
	if o != nil && !common.IsNil(o.Percentage) {
		return true
	}

	return false
}

// SetPercentage gets a reference to the given float64 and assigns it to the Percentage field.
func (o *Currency) SetPercentage(v float64) {
	o.Percentage = &v
}

func (o Currency) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Currency) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	toSerialize["currencyCode"] = o.CurrencyCode
	if !common.IsNil(o.Percentage) {
		toSerialize["percentage"] = o.Percentage
	}
	return toSerialize, nil
}

type NullableCurrency struct {
	value *Currency
	isSet bool
}

func (v NullableCurrency) Get() *Currency {
	return v.value
}

func (v *NullableCurrency) Set(val *Currency) {
	v.value = val
	v.isSet = true
}

func (v NullableCurrency) IsSet() bool {
	return v.isSet
}

func (v *NullableCurrency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCurrency(val *Currency) *NullableCurrency {
	return &NullableCurrency{value: val, isSet: true}
}

func (v NullableCurrency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCurrency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
