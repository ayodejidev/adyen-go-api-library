/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// checks if the AdditionalCommission type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AdditionalCommission{}

// AdditionalCommission struct for AdditionalCommission
type AdditionalCommission struct {
	// Unique identifier of the balance account to which the additional commission is booked.
	BalanceAccountId *string `json:"balanceAccountId,omitempty"`
	// A fixed commission fee, in minor units.
	FixedAmount *int64 `json:"fixedAmount,omitempty"`
	// A variable commission fee, in basis points.
	VariablePercentage *int64 `json:"variablePercentage,omitempty"`
}

// NewAdditionalCommission instantiates a new AdditionalCommission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalCommission() *AdditionalCommission {
	this := AdditionalCommission{}
	return &this
}

// NewAdditionalCommissionWithDefaults instantiates a new AdditionalCommission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalCommissionWithDefaults() *AdditionalCommission {
	this := AdditionalCommission{}
	return &this
}

// GetBalanceAccountId returns the BalanceAccountId field value if set, zero value otherwise.
func (o *AdditionalCommission) GetBalanceAccountId() string {
	if o == nil || common.IsNil(o.BalanceAccountId) {
		var ret string
		return ret
	}
	return *o.BalanceAccountId
}

// GetBalanceAccountIdOk returns a tuple with the BalanceAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalCommission) GetBalanceAccountIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.BalanceAccountId) {
		return nil, false
	}
	return o.BalanceAccountId, true
}

// HasBalanceAccountId returns a boolean if a field has been set.
func (o *AdditionalCommission) HasBalanceAccountId() bool {
	if o != nil && !common.IsNil(o.BalanceAccountId) {
		return true
	}

	return false
}

// SetBalanceAccountId gets a reference to the given string and assigns it to the BalanceAccountId field.
func (o *AdditionalCommission) SetBalanceAccountId(v string) {
	o.BalanceAccountId = &v
}

// GetFixedAmount returns the FixedAmount field value if set, zero value otherwise.
func (o *AdditionalCommission) GetFixedAmount() int64 {
	if o == nil || common.IsNil(o.FixedAmount) {
		var ret int64
		return ret
	}
	return *o.FixedAmount
}

// GetFixedAmountOk returns a tuple with the FixedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalCommission) GetFixedAmountOk() (*int64, bool) {
	if o == nil || common.IsNil(o.FixedAmount) {
		return nil, false
	}
	return o.FixedAmount, true
}

// HasFixedAmount returns a boolean if a field has been set.
func (o *AdditionalCommission) HasFixedAmount() bool {
	if o != nil && !common.IsNil(o.FixedAmount) {
		return true
	}

	return false
}

// SetFixedAmount gets a reference to the given int64 and assigns it to the FixedAmount field.
func (o *AdditionalCommission) SetFixedAmount(v int64) {
	o.FixedAmount = &v
}

// GetVariablePercentage returns the VariablePercentage field value if set, zero value otherwise.
func (o *AdditionalCommission) GetVariablePercentage() int64 {
	if o == nil || common.IsNil(o.VariablePercentage) {
		var ret int64
		return ret
	}
	return *o.VariablePercentage
}

// GetVariablePercentageOk returns a tuple with the VariablePercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalCommission) GetVariablePercentageOk() (*int64, bool) {
	if o == nil || common.IsNil(o.VariablePercentage) {
		return nil, false
	}
	return o.VariablePercentage, true
}

// HasVariablePercentage returns a boolean if a field has been set.
func (o *AdditionalCommission) HasVariablePercentage() bool {
	if o != nil && !common.IsNil(o.VariablePercentage) {
		return true
	}

	return false
}

// SetVariablePercentage gets a reference to the given int64 and assigns it to the VariablePercentage field.
func (o *AdditionalCommission) SetVariablePercentage(v int64) {
	o.VariablePercentage = &v
}

func (o AdditionalCommission) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdditionalCommission) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.BalanceAccountId) {
		toSerialize["balanceAccountId"] = o.BalanceAccountId
	}
	if !common.IsNil(o.FixedAmount) {
		toSerialize["fixedAmount"] = o.FixedAmount
	}
	if !common.IsNil(o.VariablePercentage) {
		toSerialize["variablePercentage"] = o.VariablePercentage
	}
	return toSerialize, nil
}

type NullableAdditionalCommission struct {
	value *AdditionalCommission
	isSet bool
}

func (v NullableAdditionalCommission) Get() *AdditionalCommission {
	return v.value
}

func (v *NullableAdditionalCommission) Set(val *AdditionalCommission) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalCommission) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalCommission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalCommission(val *AdditionalCommission) *NullableAdditionalCommission {
	return &NullableAdditionalCommission{value: val, isSet: true}
}

func (v NullableAdditionalCommission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalCommission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
