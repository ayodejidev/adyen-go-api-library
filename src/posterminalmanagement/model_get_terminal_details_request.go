/*
POS Terminal Management API

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package posterminalmanagement

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v12/src/common"
)

// checks if the GetTerminalDetailsRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetTerminalDetailsRequest{}

// GetTerminalDetailsRequest struct for GetTerminalDetailsRequest
type GetTerminalDetailsRequest struct {
	// The unique terminal ID in the format `[Device model]-[Serial number]`.   For example, **V400m-324689776**.
	Terminal string `json:"terminal"`
}

// NewGetTerminalDetailsRequest instantiates a new GetTerminalDetailsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTerminalDetailsRequest(terminal string) *GetTerminalDetailsRequest {
	this := GetTerminalDetailsRequest{}
	this.Terminal = terminal
	return &this
}

// NewGetTerminalDetailsRequestWithDefaults instantiates a new GetTerminalDetailsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTerminalDetailsRequestWithDefaults() *GetTerminalDetailsRequest {
	this := GetTerminalDetailsRequest{}
	return &this
}

// GetTerminal returns the Terminal field value
func (o *GetTerminalDetailsRequest) GetTerminal() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Terminal
}

// GetTerminalOk returns a tuple with the Terminal field value
// and a boolean to check if the value has been set.
func (o *GetTerminalDetailsRequest) GetTerminalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Terminal, true
}

// SetTerminal sets field value
func (o *GetTerminalDetailsRequest) SetTerminal(v string) {
	o.Terminal = v
}

func (o GetTerminalDetailsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTerminalDetailsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["terminal"] = o.Terminal
	return toSerialize, nil
}

type NullableGetTerminalDetailsRequest struct {
	value *GetTerminalDetailsRequest
	isSet bool
}

func (v NullableGetTerminalDetailsRequest) Get() *GetTerminalDetailsRequest {
	return v.value
}

func (v *NullableGetTerminalDetailsRequest) Set(val *GetTerminalDetailsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTerminalDetailsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTerminalDetailsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTerminalDetailsRequest(val *GetTerminalDetailsRequest) *NullableGetTerminalDetailsRequest {
	return &NullableGetTerminalDetailsRequest{value: val, isSet: true}
}

func (v NullableGetTerminalDetailsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTerminalDetailsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
