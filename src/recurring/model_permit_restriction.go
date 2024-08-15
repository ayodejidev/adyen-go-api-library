/*
Adyen Recurring API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package recurring

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v12/src/common"
)

// checks if the PermitRestriction type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PermitRestriction{}

// PermitRestriction struct for PermitRestriction
type PermitRestriction struct {
	MaxAmount              *Amount `json:"maxAmount,omitempty"`
	SingleTransactionLimit *Amount `json:"singleTransactionLimit,omitempty"`
	// Only a single payment can be made using this permit if set to true, otherwise multiple payments are allowed.
	SingleUse *bool `json:"singleUse,omitempty"`
}

// NewPermitRestriction instantiates a new PermitRestriction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermitRestriction() *PermitRestriction {
	this := PermitRestriction{}
	return &this
}

// NewPermitRestrictionWithDefaults instantiates a new PermitRestriction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermitRestrictionWithDefaults() *PermitRestriction {
	this := PermitRestriction{}
	return &this
}

// GetMaxAmount returns the MaxAmount field value if set, zero value otherwise.
func (o *PermitRestriction) GetMaxAmount() Amount {
	if o == nil || common.IsNil(o.MaxAmount) {
		var ret Amount
		return ret
	}
	return *o.MaxAmount
}

// GetMaxAmountOk returns a tuple with the MaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermitRestriction) GetMaxAmountOk() (*Amount, bool) {
	if o == nil || common.IsNil(o.MaxAmount) {
		return nil, false
	}
	return o.MaxAmount, true
}

// HasMaxAmount returns a boolean if a field has been set.
func (o *PermitRestriction) HasMaxAmount() bool {
	if o != nil && !common.IsNil(o.MaxAmount) {
		return true
	}

	return false
}

// SetMaxAmount gets a reference to the given Amount and assigns it to the MaxAmount field.
func (o *PermitRestriction) SetMaxAmount(v Amount) {
	o.MaxAmount = &v
}

// GetSingleTransactionLimit returns the SingleTransactionLimit field value if set, zero value otherwise.
func (o *PermitRestriction) GetSingleTransactionLimit() Amount {
	if o == nil || common.IsNil(o.SingleTransactionLimit) {
		var ret Amount
		return ret
	}
	return *o.SingleTransactionLimit
}

// GetSingleTransactionLimitOk returns a tuple with the SingleTransactionLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermitRestriction) GetSingleTransactionLimitOk() (*Amount, bool) {
	if o == nil || common.IsNil(o.SingleTransactionLimit) {
		return nil, false
	}
	return o.SingleTransactionLimit, true
}

// HasSingleTransactionLimit returns a boolean if a field has been set.
func (o *PermitRestriction) HasSingleTransactionLimit() bool {
	if o != nil && !common.IsNil(o.SingleTransactionLimit) {
		return true
	}

	return false
}

// SetSingleTransactionLimit gets a reference to the given Amount and assigns it to the SingleTransactionLimit field.
func (o *PermitRestriction) SetSingleTransactionLimit(v Amount) {
	o.SingleTransactionLimit = &v
}

// GetSingleUse returns the SingleUse field value if set, zero value otherwise.
func (o *PermitRestriction) GetSingleUse() bool {
	if o == nil || common.IsNil(o.SingleUse) {
		var ret bool
		return ret
	}
	return *o.SingleUse
}

// GetSingleUseOk returns a tuple with the SingleUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermitRestriction) GetSingleUseOk() (*bool, bool) {
	if o == nil || common.IsNil(o.SingleUse) {
		return nil, false
	}
	return o.SingleUse, true
}

// HasSingleUse returns a boolean if a field has been set.
func (o *PermitRestriction) HasSingleUse() bool {
	if o != nil && !common.IsNil(o.SingleUse) {
		return true
	}

	return false
}

// SetSingleUse gets a reference to the given bool and assigns it to the SingleUse field.
func (o *PermitRestriction) SetSingleUse(v bool) {
	o.SingleUse = &v
}

func (o PermitRestriction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PermitRestriction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.MaxAmount) {
		toSerialize["maxAmount"] = o.MaxAmount
	}
	if !common.IsNil(o.SingleTransactionLimit) {
		toSerialize["singleTransactionLimit"] = o.SingleTransactionLimit
	}
	if !common.IsNil(o.SingleUse) {
		toSerialize["singleUse"] = o.SingleUse
	}
	return toSerialize, nil
}

type NullablePermitRestriction struct {
	value *PermitRestriction
	isSet bool
}

func (v NullablePermitRestriction) Get() *PermitRestriction {
	return v.value
}

func (v *NullablePermitRestriction) Set(val *PermitRestriction) {
	v.value = val
	v.isSet = true
}

func (v NullablePermitRestriction) IsSet() bool {
	return v.isSet
}

func (v *NullablePermitRestriction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermitRestriction(val *PermitRestriction) *NullablePermitRestriction {
	return &NullablePermitRestriction{value: val, isSet: true}
}

func (v NullablePermitRestriction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermitRestriction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
