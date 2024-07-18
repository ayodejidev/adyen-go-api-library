/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v11/src/common"
)

// checks if the TerminalActionScheduleDetail type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TerminalActionScheduleDetail{}

// TerminalActionScheduleDetail struct for TerminalActionScheduleDetail
type TerminalActionScheduleDetail struct {
	// The ID of the action on the specified terminal.
	Id *string `json:"id,omitempty"`
	// The unique ID of the terminal that the action applies to.
	TerminalId *string `json:"terminalId,omitempty"`
}

// NewTerminalActionScheduleDetail instantiates a new TerminalActionScheduleDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerminalActionScheduleDetail() *TerminalActionScheduleDetail {
	this := TerminalActionScheduleDetail{}
	return &this
}

// NewTerminalActionScheduleDetailWithDefaults instantiates a new TerminalActionScheduleDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerminalActionScheduleDetailWithDefaults() *TerminalActionScheduleDetail {
	this := TerminalActionScheduleDetail{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TerminalActionScheduleDetail) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalActionScheduleDetail) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TerminalActionScheduleDetail) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TerminalActionScheduleDetail) SetId(v string) {
	o.Id = &v
}

// GetTerminalId returns the TerminalId field value if set, zero value otherwise.
func (o *TerminalActionScheduleDetail) GetTerminalId() string {
	if o == nil || common.IsNil(o.TerminalId) {
		var ret string
		return ret
	}
	return *o.TerminalId
}

// GetTerminalIdOk returns a tuple with the TerminalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalActionScheduleDetail) GetTerminalIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TerminalId) {
		return nil, false
	}
	return o.TerminalId, true
}

// HasTerminalId returns a boolean if a field has been set.
func (o *TerminalActionScheduleDetail) HasTerminalId() bool {
	if o != nil && !common.IsNil(o.TerminalId) {
		return true
	}

	return false
}

// SetTerminalId gets a reference to the given string and assigns it to the TerminalId field.
func (o *TerminalActionScheduleDetail) SetTerminalId(v string) {
	o.TerminalId = &v
}

func (o TerminalActionScheduleDetail) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerminalActionScheduleDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.TerminalId) {
		toSerialize["terminalId"] = o.TerminalId
	}
	return toSerialize, nil
}

type NullableTerminalActionScheduleDetail struct {
	value *TerminalActionScheduleDetail
	isSet bool
}

func (v NullableTerminalActionScheduleDetail) Get() *TerminalActionScheduleDetail {
	return v.value
}

func (v *NullableTerminalActionScheduleDetail) Set(val *TerminalActionScheduleDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminalActionScheduleDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminalActionScheduleDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminalActionScheduleDetail(val *TerminalActionScheduleDetail) *NullableTerminalActionScheduleDetail {
	return &NullableTerminalActionScheduleDetail{value: val, isSet: true}
}

func (v NullableTerminalActionScheduleDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminalActionScheduleDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
