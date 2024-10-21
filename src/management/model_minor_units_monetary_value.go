/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v14/src/common"
)

// checks if the MinorUnitsMonetaryValue type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MinorUnitsMonetaryValue{}

// MinorUnitsMonetaryValue struct for MinorUnitsMonetaryValue
type MinorUnitsMonetaryValue struct {
	// The transaction amount, in [minor units](https://docs.adyen.com/development-resources/currency-codes).
	Amount *int32 `json:"amount,omitempty"`
	// The three-character [ISO currency code](https://docs.adyen.com/development-resources/currency-codes).
	CurrencyCode *string `json:"currencyCode,omitempty"`
}

// NewMinorUnitsMonetaryValue instantiates a new MinorUnitsMonetaryValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMinorUnitsMonetaryValue() *MinorUnitsMonetaryValue {
	this := MinorUnitsMonetaryValue{}
	return &this
}

// NewMinorUnitsMonetaryValueWithDefaults instantiates a new MinorUnitsMonetaryValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMinorUnitsMonetaryValueWithDefaults() *MinorUnitsMonetaryValue {
	this := MinorUnitsMonetaryValue{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *MinorUnitsMonetaryValue) GetAmount() int32 {
	if o == nil || common.IsNil(o.Amount) {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinorUnitsMonetaryValue) GetAmountOk() (*int32, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *MinorUnitsMonetaryValue) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *MinorUnitsMonetaryValue) SetAmount(v int32) {
	o.Amount = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *MinorUnitsMonetaryValue) GetCurrencyCode() string {
	if o == nil || common.IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinorUnitsMonetaryValue) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *MinorUnitsMonetaryValue) HasCurrencyCode() bool {
	if o != nil && !common.IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *MinorUnitsMonetaryValue) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

func (o MinorUnitsMonetaryValue) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MinorUnitsMonetaryValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.CurrencyCode) {
		toSerialize["currencyCode"] = o.CurrencyCode
	}
	return toSerialize, nil
}

type NullableMinorUnitsMonetaryValue struct {
	value *MinorUnitsMonetaryValue
	isSet bool
}

func (v NullableMinorUnitsMonetaryValue) Get() *MinorUnitsMonetaryValue {
	return v.value
}

func (v *NullableMinorUnitsMonetaryValue) Set(val *MinorUnitsMonetaryValue) {
	v.value = val
	v.isSet = true
}

func (v NullableMinorUnitsMonetaryValue) IsSet() bool {
	return v.isSet
}

func (v *NullableMinorUnitsMonetaryValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMinorUnitsMonetaryValue(val *MinorUnitsMonetaryValue) *NullableMinorUnitsMonetaryValue {
	return &NullableMinorUnitsMonetaryValue{value: val, isSet: true}
}

func (v NullableMinorUnitsMonetaryValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMinorUnitsMonetaryValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
