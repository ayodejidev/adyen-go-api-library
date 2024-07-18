/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v11/src/common"
)

// checks if the AmountMinMaxRequirement type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AmountMinMaxRequirement{}

// AmountMinMaxRequirement struct for AmountMinMaxRequirement
type AmountMinMaxRequirement struct {
	// Specifies the eligible amounts for a particular route.
	Description *string `json:"description,omitempty"`
	// Maximum amount.
	Max *int64 `json:"max,omitempty"`
	// Minimum amount.
	Min *int64 `json:"min,omitempty"`
	// **amountMinMaxRequirement**
	Type string `json:"type"`
}

// NewAmountMinMaxRequirement instantiates a new AmountMinMaxRequirement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmountMinMaxRequirement(type_ string) *AmountMinMaxRequirement {
	this := AmountMinMaxRequirement{}
	this.Type = type_
	return &this
}

// NewAmountMinMaxRequirementWithDefaults instantiates a new AmountMinMaxRequirement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmountMinMaxRequirementWithDefaults() *AmountMinMaxRequirement {
	this := AmountMinMaxRequirement{}
	var type_ string = "amountMinMaxRequirement"
	this.Type = type_
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AmountMinMaxRequirement) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmountMinMaxRequirement) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AmountMinMaxRequirement) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AmountMinMaxRequirement) SetDescription(v string) {
	o.Description = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *AmountMinMaxRequirement) GetMax() int64 {
	if o == nil || common.IsNil(o.Max) {
		var ret int64
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmountMinMaxRequirement) GetMaxOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Max) {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *AmountMinMaxRequirement) HasMax() bool {
	if o != nil && !common.IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given int64 and assigns it to the Max field.
func (o *AmountMinMaxRequirement) SetMax(v int64) {
	o.Max = &v
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *AmountMinMaxRequirement) GetMin() int64 {
	if o == nil || common.IsNil(o.Min) {
		var ret int64
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmountMinMaxRequirement) GetMinOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Min) {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *AmountMinMaxRequirement) HasMin() bool {
	if o != nil && !common.IsNil(o.Min) {
		return true
	}

	return false
}

// SetMin gets a reference to the given int64 and assigns it to the Min field.
func (o *AmountMinMaxRequirement) SetMin(v int64) {
	o.Min = &v
}

// GetType returns the Type field value
func (o *AmountMinMaxRequirement) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AmountMinMaxRequirement) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AmountMinMaxRequirement) SetType(v string) {
	o.Type = v
}

func (o AmountMinMaxRequirement) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AmountMinMaxRequirement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !common.IsNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	if !common.IsNil(o.Min) {
		toSerialize["min"] = o.Min
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableAmountMinMaxRequirement struct {
	value *AmountMinMaxRequirement
	isSet bool
}

func (v NullableAmountMinMaxRequirement) Get() *AmountMinMaxRequirement {
	return v.value
}

func (v *NullableAmountMinMaxRequirement) Set(val *AmountMinMaxRequirement) {
	v.value = val
	v.isSet = true
}

func (v NullableAmountMinMaxRequirement) IsSet() bool {
	return v.isSet
}

func (v *NullableAmountMinMaxRequirement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmountMinMaxRequirement(val *AmountMinMaxRequirement) *NullableAmountMinMaxRequirement {
	return &NullableAmountMinMaxRequirement{value: val, isSet: true}
}

func (v NullableAmountMinMaxRequirement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmountMinMaxRequirement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *AmountMinMaxRequirement) isValidType() bool {
	var allowedEnumValues = []string{"amountMinMaxRequirement"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
