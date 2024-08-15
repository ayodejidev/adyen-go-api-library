/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v12/src/common"
)

// checks if the TicketInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TicketInfo{}

// TicketInfo struct for TicketInfo
type TicketInfo struct {
	// Ticket requestorId
	RequestorId *string `json:"requestorId,omitempty"`
}

// NewTicketInfo instantiates a new TicketInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTicketInfo() *TicketInfo {
	this := TicketInfo{}
	return &this
}

// NewTicketInfoWithDefaults instantiates a new TicketInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTicketInfoWithDefaults() *TicketInfo {
	this := TicketInfo{}
	return &this
}

// GetRequestorId returns the RequestorId field value if set, zero value otherwise.
func (o *TicketInfo) GetRequestorId() string {
	if o == nil || common.IsNil(o.RequestorId) {
		var ret string
		return ret
	}
	return *o.RequestorId
}

// GetRequestorIdOk returns a tuple with the RequestorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TicketInfo) GetRequestorIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.RequestorId) {
		return nil, false
	}
	return o.RequestorId, true
}

// HasRequestorId returns a boolean if a field has been set.
func (o *TicketInfo) HasRequestorId() bool {
	if o != nil && !common.IsNil(o.RequestorId) {
		return true
	}

	return false
}

// SetRequestorId gets a reference to the given string and assigns it to the RequestorId field.
func (o *TicketInfo) SetRequestorId(v string) {
	o.RequestorId = &v
}

func (o TicketInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TicketInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.RequestorId) {
		toSerialize["requestorId"] = o.RequestorId
	}
	return toSerialize, nil
}

type NullableTicketInfo struct {
	value *TicketInfo
	isSet bool
}

func (v NullableTicketInfo) Get() *TicketInfo {
	return v.value
}

func (v *NullableTicketInfo) Set(val *TicketInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTicketInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTicketInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTicketInfo(val *TicketInfo) *NullableTicketInfo {
	return &NullableTicketInfo{value: val, isSet: true}
}

func (v NullableTicketInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTicketInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
