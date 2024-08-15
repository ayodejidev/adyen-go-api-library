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

// checks if the MerchantNamesRestriction type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MerchantNamesRestriction{}

// MerchantNamesRestriction struct for MerchantNamesRestriction
type MerchantNamesRestriction struct {
	// Defines how the condition must be evaluated.
	Operation string        `json:"operation"`
	Value     []StringMatch `json:"value,omitempty"`
}

// NewMerchantNamesRestriction instantiates a new MerchantNamesRestriction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantNamesRestriction(operation string) *MerchantNamesRestriction {
	this := MerchantNamesRestriction{}
	this.Operation = operation
	return &this
}

// NewMerchantNamesRestrictionWithDefaults instantiates a new MerchantNamesRestriction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantNamesRestrictionWithDefaults() *MerchantNamesRestriction {
	this := MerchantNamesRestriction{}
	return &this
}

// GetOperation returns the Operation field value
func (o *MerchantNamesRestriction) GetOperation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *MerchantNamesRestriction) GetOperationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value
func (o *MerchantNamesRestriction) SetOperation(v string) {
	o.Operation = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *MerchantNamesRestriction) GetValue() []StringMatch {
	if o == nil || common.IsNil(o.Value) {
		var ret []StringMatch
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantNamesRestriction) GetValueOk() ([]StringMatch, bool) {
	if o == nil || common.IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *MerchantNamesRestriction) HasValue() bool {
	if o != nil && !common.IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given []StringMatch and assigns it to the Value field.
func (o *MerchantNamesRestriction) SetValue(v []StringMatch) {
	o.Value = v
}

func (o MerchantNamesRestriction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantNamesRestriction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operation"] = o.Operation
	if !common.IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableMerchantNamesRestriction struct {
	value *MerchantNamesRestriction
	isSet bool
}

func (v NullableMerchantNamesRestriction) Get() *MerchantNamesRestriction {
	return v.value
}

func (v *NullableMerchantNamesRestriction) Set(val *MerchantNamesRestriction) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantNamesRestriction) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantNamesRestriction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantNamesRestriction(val *MerchantNamesRestriction) *NullableMerchantNamesRestriction {
	return &NullableMerchantNamesRestriction{value: val, isSet: true}
}

func (v NullableMerchantNamesRestriction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantNamesRestriction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
