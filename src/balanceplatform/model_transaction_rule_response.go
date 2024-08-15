/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v12/src/common"
)

// checks if the TransactionRuleResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TransactionRuleResponse{}

// TransactionRuleResponse struct for TransactionRuleResponse
type TransactionRuleResponse struct {
	TransactionRule *TransactionRule `json:"transactionRule,omitempty"`
}

// NewTransactionRuleResponse instantiates a new TransactionRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionRuleResponse() *TransactionRuleResponse {
	this := TransactionRuleResponse{}
	return &this
}

// NewTransactionRuleResponseWithDefaults instantiates a new TransactionRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionRuleResponseWithDefaults() *TransactionRuleResponse {
	this := TransactionRuleResponse{}
	return &this
}

// GetTransactionRule returns the TransactionRule field value if set, zero value otherwise.
func (o *TransactionRuleResponse) GetTransactionRule() TransactionRule {
	if o == nil || common.IsNil(o.TransactionRule) {
		var ret TransactionRule
		return ret
	}
	return *o.TransactionRule
}

// GetTransactionRuleOk returns a tuple with the TransactionRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRuleResponse) GetTransactionRuleOk() (*TransactionRule, bool) {
	if o == nil || common.IsNil(o.TransactionRule) {
		return nil, false
	}
	return o.TransactionRule, true
}

// HasTransactionRule returns a boolean if a field has been set.
func (o *TransactionRuleResponse) HasTransactionRule() bool {
	if o != nil && !common.IsNil(o.TransactionRule) {
		return true
	}

	return false
}

// SetTransactionRule gets a reference to the given TransactionRule and assigns it to the TransactionRule field.
func (o *TransactionRuleResponse) SetTransactionRule(v TransactionRule) {
	o.TransactionRule = &v
}

func (o TransactionRuleResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionRuleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.TransactionRule) {
		toSerialize["transactionRule"] = o.TransactionRule
	}
	return toSerialize, nil
}

type NullableTransactionRuleResponse struct {
	value *TransactionRuleResponse
	isSet bool
}

func (v NullableTransactionRuleResponse) Get() *TransactionRuleResponse {
	return v.value
}

func (v *NullableTransactionRuleResponse) Set(val *TransactionRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionRuleResponse(val *TransactionRuleResponse) *NullableTransactionRuleResponse {
	return &NullableTransactionRuleResponse{value: val, isSet: true}
}

func (v NullableTransactionRuleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
