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

// checks if the MatchingValuesRestriction type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MatchingValuesRestriction{}

// MatchingValuesRestriction struct for MatchingValuesRestriction
type MatchingValuesRestriction struct {
	// Defines how the condition must be evaluated.
	Operation string   `json:"operation"`
	Value     []string `json:"value,omitempty"`
}

// NewMatchingValuesRestriction instantiates a new MatchingValuesRestriction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMatchingValuesRestriction(operation string) *MatchingValuesRestriction {
	this := MatchingValuesRestriction{}
	this.Operation = operation
	return &this
}

// NewMatchingValuesRestrictionWithDefaults instantiates a new MatchingValuesRestriction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMatchingValuesRestrictionWithDefaults() *MatchingValuesRestriction {
	this := MatchingValuesRestriction{}
	return &this
}

// GetOperation returns the Operation field value
func (o *MatchingValuesRestriction) GetOperation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *MatchingValuesRestriction) GetOperationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value
func (o *MatchingValuesRestriction) SetOperation(v string) {
	o.Operation = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *MatchingValuesRestriction) GetValue() []string {
	if o == nil || common.IsNil(o.Value) {
		var ret []string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatchingValuesRestriction) GetValueOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *MatchingValuesRestriction) HasValue() bool {
	if o != nil && !common.IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given []string and assigns it to the Value field.
func (o *MatchingValuesRestriction) SetValue(v []string) {
	o.Value = v
}

func (o MatchingValuesRestriction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MatchingValuesRestriction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operation"] = o.Operation
	if !common.IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableMatchingValuesRestriction struct {
	value *MatchingValuesRestriction
	isSet bool
}

func (v NullableMatchingValuesRestriction) Get() *MatchingValuesRestriction {
	return v.value
}

func (v *NullableMatchingValuesRestriction) Set(val *MatchingValuesRestriction) {
	v.value = val
	v.isSet = true
}

func (v NullableMatchingValuesRestriction) IsSet() bool {
	return v.isSet
}

func (v *NullableMatchingValuesRestriction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMatchingValuesRestriction(val *MatchingValuesRestriction) *NullableMatchingValuesRestriction {
	return &NullableMatchingValuesRestriction{value: val, isSet: true}
}

func (v NullableMatchingValuesRestriction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMatchingValuesRestriction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
