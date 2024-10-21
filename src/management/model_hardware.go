/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v14/src/common"
)

// checks if the Hardware type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Hardware{}

// Hardware struct for Hardware
type Hardware struct {
	// The brightness of the display when the terminal is being used, expressed as a percentage.
	DisplayMaximumBackLight *int32 `json:"displayMaximumBackLight,omitempty"`
	// The hour of the day when the terminal is set to reset the Totals report. By default, the reset hour is at 6:00 AM in the timezone of the terminal. Minimum value: 0, maximum value: 23.
	ResetTotalsHour *int32 `json:"resetTotalsHour,omitempty"`
	// The hour of the day when the terminal is set to reboot to apply the configuration and software updates. By default, the restart hour is at 6:00 AM in the timezone of the terminal. Minimum value: 0, maximum value: 23.
	RestartHour *int32 `json:"restartHour,omitempty"`
}

// NewHardware instantiates a new Hardware object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHardware() *Hardware {
	this := Hardware{}
	return &this
}

// NewHardwareWithDefaults instantiates a new Hardware object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHardwareWithDefaults() *Hardware {
	this := Hardware{}
	return &this
}

// GetDisplayMaximumBackLight returns the DisplayMaximumBackLight field value if set, zero value otherwise.
func (o *Hardware) GetDisplayMaximumBackLight() int32 {
	if o == nil || common.IsNil(o.DisplayMaximumBackLight) {
		var ret int32
		return ret
	}
	return *o.DisplayMaximumBackLight
}

// GetDisplayMaximumBackLightOk returns a tuple with the DisplayMaximumBackLight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hardware) GetDisplayMaximumBackLightOk() (*int32, bool) {
	if o == nil || common.IsNil(o.DisplayMaximumBackLight) {
		return nil, false
	}
	return o.DisplayMaximumBackLight, true
}

// HasDisplayMaximumBackLight returns a boolean if a field has been set.
func (o *Hardware) HasDisplayMaximumBackLight() bool {
	if o != nil && !common.IsNil(o.DisplayMaximumBackLight) {
		return true
	}

	return false
}

// SetDisplayMaximumBackLight gets a reference to the given int32 and assigns it to the DisplayMaximumBackLight field.
func (o *Hardware) SetDisplayMaximumBackLight(v int32) {
	o.DisplayMaximumBackLight = &v
}

// GetResetTotalsHour returns the ResetTotalsHour field value if set, zero value otherwise.
func (o *Hardware) GetResetTotalsHour() int32 {
	if o == nil || common.IsNil(o.ResetTotalsHour) {
		var ret int32
		return ret
	}
	return *o.ResetTotalsHour
}

// GetResetTotalsHourOk returns a tuple with the ResetTotalsHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hardware) GetResetTotalsHourOk() (*int32, bool) {
	if o == nil || common.IsNil(o.ResetTotalsHour) {
		return nil, false
	}
	return o.ResetTotalsHour, true
}

// HasResetTotalsHour returns a boolean if a field has been set.
func (o *Hardware) HasResetTotalsHour() bool {
	if o != nil && !common.IsNil(o.ResetTotalsHour) {
		return true
	}

	return false
}

// SetResetTotalsHour gets a reference to the given int32 and assigns it to the ResetTotalsHour field.
func (o *Hardware) SetResetTotalsHour(v int32) {
	o.ResetTotalsHour = &v
}

// GetRestartHour returns the RestartHour field value if set, zero value otherwise.
func (o *Hardware) GetRestartHour() int32 {
	if o == nil || common.IsNil(o.RestartHour) {
		var ret int32
		return ret
	}
	return *o.RestartHour
}

// GetRestartHourOk returns a tuple with the RestartHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hardware) GetRestartHourOk() (*int32, bool) {
	if o == nil || common.IsNil(o.RestartHour) {
		return nil, false
	}
	return o.RestartHour, true
}

// HasRestartHour returns a boolean if a field has been set.
func (o *Hardware) HasRestartHour() bool {
	if o != nil && !common.IsNil(o.RestartHour) {
		return true
	}

	return false
}

// SetRestartHour gets a reference to the given int32 and assigns it to the RestartHour field.
func (o *Hardware) SetRestartHour(v int32) {
	o.RestartHour = &v
}

func (o Hardware) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Hardware) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.DisplayMaximumBackLight) {
		toSerialize["displayMaximumBackLight"] = o.DisplayMaximumBackLight
	}
	if !common.IsNil(o.ResetTotalsHour) {
		toSerialize["resetTotalsHour"] = o.ResetTotalsHour
	}
	if !common.IsNil(o.RestartHour) {
		toSerialize["restartHour"] = o.RestartHour
	}
	return toSerialize, nil
}

type NullableHardware struct {
	value *Hardware
	isSet bool
}

func (v NullableHardware) Get() *Hardware {
	return v.value
}

func (v *NullableHardware) Set(val *Hardware) {
	v.value = val
	v.isSet = true
}

func (v NullableHardware) IsSet() bool {
	return v.isSet
}

func (v *NullableHardware) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHardware(val *Hardware) *NullableHardware {
	return &NullableHardware{value: val, isSet: true}
}

func (v NullableHardware) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHardware) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
