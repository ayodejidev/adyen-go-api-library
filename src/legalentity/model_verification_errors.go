/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v11/src/common"
)

// checks if the VerificationErrors type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &VerificationErrors{}

// VerificationErrors struct for VerificationErrors
type VerificationErrors struct {
	// List of the verification errors.
	Problems []CapabilityProblem `json:"problems,omitempty"`
}

// NewVerificationErrors instantiates a new VerificationErrors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerificationErrors() *VerificationErrors {
	this := VerificationErrors{}
	return &this
}

// NewVerificationErrorsWithDefaults instantiates a new VerificationErrors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerificationErrorsWithDefaults() *VerificationErrors {
	this := VerificationErrors{}
	return &this
}

// GetProblems returns the Problems field value if set, zero value otherwise.
func (o *VerificationErrors) GetProblems() []CapabilityProblem {
	if o == nil || common.IsNil(o.Problems) {
		var ret []CapabilityProblem
		return ret
	}
	return o.Problems
}

// GetProblemsOk returns a tuple with the Problems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationErrors) GetProblemsOk() ([]CapabilityProblem, bool) {
	if o == nil || common.IsNil(o.Problems) {
		return nil, false
	}
	return o.Problems, true
}

// HasProblems returns a boolean if a field has been set.
func (o *VerificationErrors) HasProblems() bool {
	if o != nil && !common.IsNil(o.Problems) {
		return true
	}

	return false
}

// SetProblems gets a reference to the given []CapabilityProblem and assigns it to the Problems field.
func (o *VerificationErrors) SetProblems(v []CapabilityProblem) {
	o.Problems = v
}

func (o VerificationErrors) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VerificationErrors) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Problems) {
		toSerialize["problems"] = o.Problems
	}
	return toSerialize, nil
}

type NullableVerificationErrors struct {
	value *VerificationErrors
	isSet bool
}

func (v NullableVerificationErrors) Get() *VerificationErrors {
	return v.value
}

func (v *NullableVerificationErrors) Set(val *VerificationErrors) {
	v.value = val
	v.isSet = true
}

func (v NullableVerificationErrors) IsSet() bool {
	return v.isSet
}

func (v *NullableVerificationErrors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerificationErrors(val *VerificationErrors) *NullableVerificationErrors {
	return &NullableVerificationErrors{value: val, isSet: true}
}

func (v NullableVerificationErrors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerificationErrors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
