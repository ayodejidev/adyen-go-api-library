/*
Transfers API

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transfers

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v9/src/common"
)

// checks if the TransactionRulesResult type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TransactionRulesResult{}

// TransactionRulesResult struct for TransactionRulesResult
type TransactionRulesResult struct {
	// The advice given by the Risk analysis.
	Advice *string `json:"advice,omitempty"`
	// Indicates whether the transaction passed the evaluation for all hardblock rules
	AllHardBlockRulesPassed *bool `json:"allHardBlockRulesPassed,omitempty"`
	// The score of the Risk analysis.
	Score *int32 `json:"score,omitempty"`
	// Array containing all the transaction rules that the transaction triggered.
	TriggeredTransactionRules []TransactionEventViolation `json:"triggeredTransactionRules,omitempty"`
}

// NewTransactionRulesResult instantiates a new TransactionRulesResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionRulesResult() *TransactionRulesResult {
	this := TransactionRulesResult{}
	return &this
}

// NewTransactionRulesResultWithDefaults instantiates a new TransactionRulesResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionRulesResultWithDefaults() *TransactionRulesResult {
	this := TransactionRulesResult{}
	return &this
}

// GetAdvice returns the Advice field value if set, zero value otherwise.
func (o *TransactionRulesResult) GetAdvice() string {
	if o == nil || common.IsNil(o.Advice) {
		var ret string
		return ret
	}
	return *o.Advice
}

// GetAdviceOk returns a tuple with the Advice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRulesResult) GetAdviceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Advice) {
		return nil, false
	}
	return o.Advice, true
}

// HasAdvice returns a boolean if a field has been set.
func (o *TransactionRulesResult) HasAdvice() bool {
	if o != nil && !common.IsNil(o.Advice) {
		return true
	}

	return false
}

// SetAdvice gets a reference to the given string and assigns it to the Advice field.
func (o *TransactionRulesResult) SetAdvice(v string) {
	o.Advice = &v
}

// GetAllHardBlockRulesPassed returns the AllHardBlockRulesPassed field value if set, zero value otherwise.
func (o *TransactionRulesResult) GetAllHardBlockRulesPassed() bool {
	if o == nil || common.IsNil(o.AllHardBlockRulesPassed) {
		var ret bool
		return ret
	}
	return *o.AllHardBlockRulesPassed
}

// GetAllHardBlockRulesPassedOk returns a tuple with the AllHardBlockRulesPassed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRulesResult) GetAllHardBlockRulesPassedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.AllHardBlockRulesPassed) {
		return nil, false
	}
	return o.AllHardBlockRulesPassed, true
}

// HasAllHardBlockRulesPassed returns a boolean if a field has been set.
func (o *TransactionRulesResult) HasAllHardBlockRulesPassed() bool {
	if o != nil && !common.IsNil(o.AllHardBlockRulesPassed) {
		return true
	}

	return false
}

// SetAllHardBlockRulesPassed gets a reference to the given bool and assigns it to the AllHardBlockRulesPassed field.
func (o *TransactionRulesResult) SetAllHardBlockRulesPassed(v bool) {
	o.AllHardBlockRulesPassed = &v
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *TransactionRulesResult) GetScore() int32 {
	if o == nil || common.IsNil(o.Score) {
		var ret int32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRulesResult) GetScoreOk() (*int32, bool) {
	if o == nil || common.IsNil(o.Score) {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *TransactionRulesResult) HasScore() bool {
	if o != nil && !common.IsNil(o.Score) {
		return true
	}

	return false
}

// SetScore gets a reference to the given int32 and assigns it to the Score field.
func (o *TransactionRulesResult) SetScore(v int32) {
	o.Score = &v
}

// GetTriggeredTransactionRules returns the TriggeredTransactionRules field value if set, zero value otherwise.
func (o *TransactionRulesResult) GetTriggeredTransactionRules() []TransactionEventViolation {
	if o == nil || common.IsNil(o.TriggeredTransactionRules) {
		var ret []TransactionEventViolation
		return ret
	}
	return o.TriggeredTransactionRules
}

// GetTriggeredTransactionRulesOk returns a tuple with the TriggeredTransactionRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRulesResult) GetTriggeredTransactionRulesOk() ([]TransactionEventViolation, bool) {
	if o == nil || common.IsNil(o.TriggeredTransactionRules) {
		return nil, false
	}
	return o.TriggeredTransactionRules, true
}

// HasTriggeredTransactionRules returns a boolean if a field has been set.
func (o *TransactionRulesResult) HasTriggeredTransactionRules() bool {
	if o != nil && !common.IsNil(o.TriggeredTransactionRules) {
		return true
	}

	return false
}

// SetTriggeredTransactionRules gets a reference to the given []TransactionEventViolation and assigns it to the TriggeredTransactionRules field.
func (o *TransactionRulesResult) SetTriggeredTransactionRules(v []TransactionEventViolation) {
	o.TriggeredTransactionRules = v
}

func (o TransactionRulesResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionRulesResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Advice) {
		toSerialize["advice"] = o.Advice
	}
	if !common.IsNil(o.AllHardBlockRulesPassed) {
		toSerialize["allHardBlockRulesPassed"] = o.AllHardBlockRulesPassed
	}
	if !common.IsNil(o.Score) {
		toSerialize["score"] = o.Score
	}
	if !common.IsNil(o.TriggeredTransactionRules) {
		toSerialize["triggeredTransactionRules"] = o.TriggeredTransactionRules
	}
	return toSerialize, nil
}

type NullableTransactionRulesResult struct {
	value *TransactionRulesResult
	isSet bool
}

func (v NullableTransactionRulesResult) Get() *TransactionRulesResult {
	return v.value
}

func (v *NullableTransactionRulesResult) Set(val *TransactionRulesResult) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionRulesResult) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionRulesResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionRulesResult(val *TransactionRulesResult) *NullableTransactionRulesResult {
	return &NullableTransactionRulesResult{value: val, isSet: true}
}

func (v NullableTransactionRulesResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionRulesResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
