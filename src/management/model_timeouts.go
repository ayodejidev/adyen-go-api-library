/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v15/src/common"
)

// checks if the Timeouts type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Timeouts{}

// Timeouts struct for Timeouts
type Timeouts struct {
	// Indicates the number of seconds of inactivity after which the terminal display goes into sleep mode.
	FromActiveToSleep *int32 `json:"fromActiveToSleep,omitempty"`
}

// NewTimeouts instantiates a new Timeouts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeouts() *Timeouts {
	this := Timeouts{}
	return &this
}

// NewTimeoutsWithDefaults instantiates a new Timeouts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeoutsWithDefaults() *Timeouts {
	this := Timeouts{}
	return &this
}

// GetFromActiveToSleep returns the FromActiveToSleep field value if set, zero value otherwise.
func (o *Timeouts) GetFromActiveToSleep() int32 {
	if o == nil || common.IsNil(o.FromActiveToSleep) {
		var ret int32
		return ret
	}
	return *o.FromActiveToSleep
}

// GetFromActiveToSleepOk returns a tuple with the FromActiveToSleep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Timeouts) GetFromActiveToSleepOk() (*int32, bool) {
	if o == nil || common.IsNil(o.FromActiveToSleep) {
		return nil, false
	}
	return o.FromActiveToSleep, true
}

// HasFromActiveToSleep returns a boolean if a field has been set.
func (o *Timeouts) HasFromActiveToSleep() bool {
	if o != nil && !common.IsNil(o.FromActiveToSleep) {
		return true
	}

	return false
}

// SetFromActiveToSleep gets a reference to the given int32 and assigns it to the FromActiveToSleep field.
func (o *Timeouts) SetFromActiveToSleep(v int32) {
	o.FromActiveToSleep = &v
}

func (o Timeouts) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Timeouts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.FromActiveToSleep) {
		toSerialize["fromActiveToSleep"] = o.FromActiveToSleep
	}
	return toSerialize, nil
}

type NullableTimeouts struct {
	value *Timeouts
	isSet bool
}

func (v NullableTimeouts) Get() *Timeouts {
	return v.value
}

func (v *NullableTimeouts) Set(val *Timeouts) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeouts) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeouts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeouts(val *Timeouts) *NullableTimeouts {
	return &NullableTimeouts{value: val, isSet: true}
}

func (v NullableTimeouts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeouts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
