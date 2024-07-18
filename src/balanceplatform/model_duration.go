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

// checks if the Duration type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Duration{}

// Duration struct for Duration
type Duration struct {
	// The unit of time. You can only use **minutes** and **hours** if the `interval.type` is **sliding**.  Possible values: **minutes**, **hours**, **days**, **weeks**, or **months**
	Unit *string `json:"unit,omitempty"`
	// The length of time by the unit. For example, 5 days.  The maximum duration is 90 days or an equivalent in other units. For example, 3 months.
	Value *int32 `json:"value,omitempty"`
}

// NewDuration instantiates a new Duration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDuration() *Duration {
	this := Duration{}
	return &this
}

// NewDurationWithDefaults instantiates a new Duration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDurationWithDefaults() *Duration {
	this := Duration{}
	return &this
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *Duration) GetUnit() string {
	if o == nil || common.IsNil(o.Unit) {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Duration) GetUnitOk() (*string, bool) {
	if o == nil || common.IsNil(o.Unit) {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *Duration) HasUnit() bool {
	if o != nil && !common.IsNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *Duration) SetUnit(v string) {
	o.Unit = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Duration) GetValue() int32 {
	if o == nil || common.IsNil(o.Value) {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Duration) GetValueOk() (*int32, bool) {
	if o == nil || common.IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Duration) HasValue() bool {
	if o != nil && !common.IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *Duration) SetValue(v int32) {
	o.Value = &v
}

func (o Duration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Duration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}
	if !common.IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableDuration struct {
	value *Duration
	isSet bool
}

func (v NullableDuration) Get() *Duration {
	return v.value
}

func (v *NullableDuration) Set(val *Duration) {
	v.value = val
	v.isSet = true
}

func (v NullableDuration) IsSet() bool {
	return v.isSet
}

func (v *NullableDuration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDuration(val *Duration) *NullableDuration {
	return &NullableDuration{value: val, isSet: true}
}

func (v NullableDuration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDuration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *Duration) isValidUnit() bool {
	var allowedEnumValues = []string{"days", "hours", "minutes", "months", "weeks"}
	for _, allowed := range allowedEnumValues {
		if o.GetUnit() == allowed {
			return true
		}
	}
	return false
}
