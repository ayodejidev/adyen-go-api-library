/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v14/src/common"
)

// checks if the ProcessingTypesRestriction type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ProcessingTypesRestriction{}

// ProcessingTypesRestriction struct for ProcessingTypesRestriction
type ProcessingTypesRestriction struct {
	// Defines how the condition must be evaluated.
	Operation string `json:"operation"`
	// List of processing types.  Possible values: **atmWithdraw**, **balanceInquiry**, **ecommerce**, **moto**, **pos**, **recurring**, **token**.
	Value []string `json:"value,omitempty"`
}

// NewProcessingTypesRestriction instantiates a new ProcessingTypesRestriction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessingTypesRestriction(operation string) *ProcessingTypesRestriction {
	this := ProcessingTypesRestriction{}
	this.Operation = operation
	return &this
}

// NewProcessingTypesRestrictionWithDefaults instantiates a new ProcessingTypesRestriction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessingTypesRestrictionWithDefaults() *ProcessingTypesRestriction {
	this := ProcessingTypesRestriction{}
	return &this
}

// GetOperation returns the Operation field value
func (o *ProcessingTypesRestriction) GetOperation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *ProcessingTypesRestriction) GetOperationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value
func (o *ProcessingTypesRestriction) SetOperation(v string) {
	o.Operation = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ProcessingTypesRestriction) GetValue() []string {
	if o == nil || common.IsNil(o.Value) {
		var ret []string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessingTypesRestriction) GetValueOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ProcessingTypesRestriction) HasValue() bool {
	if o != nil && !common.IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given []string and assigns it to the Value field.
func (o *ProcessingTypesRestriction) SetValue(v []string) {
	o.Value = v
}

func (o ProcessingTypesRestriction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProcessingTypesRestriction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operation"] = o.Operation
	if !common.IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableProcessingTypesRestriction struct {
	value *ProcessingTypesRestriction
	isSet bool
}

func (v NullableProcessingTypesRestriction) Get() *ProcessingTypesRestriction {
	return v.value
}

func (v *NullableProcessingTypesRestriction) Set(val *ProcessingTypesRestriction) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessingTypesRestriction) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessingTypesRestriction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessingTypesRestriction(val *ProcessingTypesRestriction) *NullableProcessingTypesRestriction {
	return &NullableProcessingTypesRestriction{value: val, isSet: true}
}

func (v NullableProcessingTypesRestriction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessingTypesRestriction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
