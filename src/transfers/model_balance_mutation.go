/*
Transfers API

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transfers

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v15/src/common"
)

// checks if the BalanceMutation type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BalanceMutation{}

// BalanceMutation struct for BalanceMutation
type BalanceMutation struct {
	// The amount in the payment's currency that is debited or credited on the balance accounting register.
	Balance *int64 `json:"balance,omitempty"`
	// The three-character [ISO currency code](https://docs.adyen.com/development-resources/currency-codes).
	Currency *string `json:"currency,omitempty"`
	// The amount in the payment's currency that is debited or credited on the received accounting register.
	Received *int64 `json:"received,omitempty"`
	// The amount in the payment's currency that is debited or credited on the reserved accounting register.
	Reserved *int64 `json:"reserved,omitempty"`
}

// NewBalanceMutation instantiates a new BalanceMutation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBalanceMutation() *BalanceMutation {
	this := BalanceMutation{}
	return &this
}

// NewBalanceMutationWithDefaults instantiates a new BalanceMutation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBalanceMutationWithDefaults() *BalanceMutation {
	this := BalanceMutation{}
	return &this
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *BalanceMutation) GetBalance() int64 {
	if o == nil || common.IsNil(o.Balance) {
		var ret int64
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceMutation) GetBalanceOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Balance) {
		return nil, false
	}
	return o.Balance, true
}

// HasBalance returns a boolean if a field has been set.
func (o *BalanceMutation) HasBalance() bool {
	if o != nil && !common.IsNil(o.Balance) {
		return true
	}

	return false
}

// SetBalance gets a reference to the given int64 and assigns it to the Balance field.
func (o *BalanceMutation) SetBalance(v int64) {
	o.Balance = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *BalanceMutation) GetCurrency() string {
	if o == nil || common.IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceMutation) GetCurrencyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *BalanceMutation) HasCurrency() bool {
	if o != nil && !common.IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *BalanceMutation) SetCurrency(v string) {
	o.Currency = &v
}

// GetReceived returns the Received field value if set, zero value otherwise.
func (o *BalanceMutation) GetReceived() int64 {
	if o == nil || common.IsNil(o.Received) {
		var ret int64
		return ret
	}
	return *o.Received
}

// GetReceivedOk returns a tuple with the Received field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceMutation) GetReceivedOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Received) {
		return nil, false
	}
	return o.Received, true
}

// HasReceived returns a boolean if a field has been set.
func (o *BalanceMutation) HasReceived() bool {
	if o != nil && !common.IsNil(o.Received) {
		return true
	}

	return false
}

// SetReceived gets a reference to the given int64 and assigns it to the Received field.
func (o *BalanceMutation) SetReceived(v int64) {
	o.Received = &v
}

// GetReserved returns the Reserved field value if set, zero value otherwise.
func (o *BalanceMutation) GetReserved() int64 {
	if o == nil || common.IsNil(o.Reserved) {
		var ret int64
		return ret
	}
	return *o.Reserved
}

// GetReservedOk returns a tuple with the Reserved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceMutation) GetReservedOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Reserved) {
		return nil, false
	}
	return o.Reserved, true
}

// HasReserved returns a boolean if a field has been set.
func (o *BalanceMutation) HasReserved() bool {
	if o != nil && !common.IsNil(o.Reserved) {
		return true
	}

	return false
}

// SetReserved gets a reference to the given int64 and assigns it to the Reserved field.
func (o *BalanceMutation) SetReserved(v int64) {
	o.Reserved = &v
}

func (o BalanceMutation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BalanceMutation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Balance) {
		toSerialize["balance"] = o.Balance
	}
	if !common.IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !common.IsNil(o.Received) {
		toSerialize["received"] = o.Received
	}
	if !common.IsNil(o.Reserved) {
		toSerialize["reserved"] = o.Reserved
	}
	return toSerialize, nil
}

type NullableBalanceMutation struct {
	value *BalanceMutation
	isSet bool
}

func (v NullableBalanceMutation) Get() *BalanceMutation {
	return v.value
}

func (v *NullableBalanceMutation) Set(val *BalanceMutation) {
	v.value = val
	v.isSet = true
}

func (v NullableBalanceMutation) IsSet() bool {
	return v.isSet
}

func (v *NullableBalanceMutation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBalanceMutation(val *BalanceMutation) *NullableBalanceMutation {
	return &NullableBalanceMutation{value: val, isSet: true}
}

func (v NullableBalanceMutation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBalanceMutation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
