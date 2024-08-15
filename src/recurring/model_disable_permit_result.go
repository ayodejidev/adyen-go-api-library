/*
Adyen Recurring API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package recurring

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v12/src/common"
)

// checks if the DisablePermitResult type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DisablePermitResult{}

// DisablePermitResult struct for DisablePermitResult
type DisablePermitResult struct {
	// A unique reference associated with the request. This value is globally unique; quote it when communicating with us about this request.
	PspReference *string `json:"pspReference,omitempty"`
	// Status of the disable request.
	Status *string `json:"status,omitempty"`
}

// NewDisablePermitResult instantiates a new DisablePermitResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDisablePermitResult() *DisablePermitResult {
	this := DisablePermitResult{}
	return &this
}

// NewDisablePermitResultWithDefaults instantiates a new DisablePermitResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDisablePermitResultWithDefaults() *DisablePermitResult {
	this := DisablePermitResult{}
	return &this
}

// GetPspReference returns the PspReference field value if set, zero value otherwise.
func (o *DisablePermitResult) GetPspReference() string {
	if o == nil || common.IsNil(o.PspReference) {
		var ret string
		return ret
	}
	return *o.PspReference
}

// GetPspReferenceOk returns a tuple with the PspReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisablePermitResult) GetPspReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.PspReference) {
		return nil, false
	}
	return o.PspReference, true
}

// HasPspReference returns a boolean if a field has been set.
func (o *DisablePermitResult) HasPspReference() bool {
	if o != nil && !common.IsNil(o.PspReference) {
		return true
	}

	return false
}

// SetPspReference gets a reference to the given string and assigns it to the PspReference field.
func (o *DisablePermitResult) SetPspReference(v string) {
	o.PspReference = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DisablePermitResult) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisablePermitResult) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DisablePermitResult) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DisablePermitResult) SetStatus(v string) {
	o.Status = &v
}

func (o DisablePermitResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DisablePermitResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.PspReference) {
		toSerialize["pspReference"] = o.PspReference
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableDisablePermitResult struct {
	value *DisablePermitResult
	isSet bool
}

func (v NullableDisablePermitResult) Get() *DisablePermitResult {
	return v.value
}

func (v *NullableDisablePermitResult) Set(val *DisablePermitResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDisablePermitResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDisablePermitResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDisablePermitResult(val *DisablePermitResult) *NullableDisablePermitResult {
	return &NullableDisablePermitResult{value: val, isSet: true}
}

func (v NullableDisablePermitResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDisablePermitResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
