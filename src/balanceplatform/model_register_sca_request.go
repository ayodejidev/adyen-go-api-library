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

// checks if the RegisterSCARequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &RegisterSCARequest{}

// RegisterSCARequest struct for RegisterSCARequest
type RegisterSCARequest struct {
	// The unique identifier of the payment instrument for which you are registering the SCA device.
	PaymentInstrumentId          string                      `json:"paymentInstrumentId"`
	StrongCustomerAuthentication DelegatedAuthenticationData `json:"strongCustomerAuthentication"`
}

// NewRegisterSCARequest instantiates a new RegisterSCARequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegisterSCARequest(paymentInstrumentId string, strongCustomerAuthentication DelegatedAuthenticationData) *RegisterSCARequest {
	this := RegisterSCARequest{}
	this.PaymentInstrumentId = paymentInstrumentId
	this.StrongCustomerAuthentication = strongCustomerAuthentication
	return &this
}

// NewRegisterSCARequestWithDefaults instantiates a new RegisterSCARequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegisterSCARequestWithDefaults() *RegisterSCARequest {
	this := RegisterSCARequest{}
	return &this
}

// GetPaymentInstrumentId returns the PaymentInstrumentId field value
func (o *RegisterSCARequest) GetPaymentInstrumentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentInstrumentId
}

// GetPaymentInstrumentIdOk returns a tuple with the PaymentInstrumentId field value
// and a boolean to check if the value has been set.
func (o *RegisterSCARequest) GetPaymentInstrumentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentInstrumentId, true
}

// SetPaymentInstrumentId sets field value
func (o *RegisterSCARequest) SetPaymentInstrumentId(v string) {
	o.PaymentInstrumentId = v
}

// GetStrongCustomerAuthentication returns the StrongCustomerAuthentication field value
func (o *RegisterSCARequest) GetStrongCustomerAuthentication() DelegatedAuthenticationData {
	if o == nil {
		var ret DelegatedAuthenticationData
		return ret
	}

	return o.StrongCustomerAuthentication
}

// GetStrongCustomerAuthenticationOk returns a tuple with the StrongCustomerAuthentication field value
// and a boolean to check if the value has been set.
func (o *RegisterSCARequest) GetStrongCustomerAuthenticationOk() (*DelegatedAuthenticationData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StrongCustomerAuthentication, true
}

// SetStrongCustomerAuthentication sets field value
func (o *RegisterSCARequest) SetStrongCustomerAuthentication(v DelegatedAuthenticationData) {
	o.StrongCustomerAuthentication = v
}

func (o RegisterSCARequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegisterSCARequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["paymentInstrumentId"] = o.PaymentInstrumentId
	toSerialize["strongCustomerAuthentication"] = o.StrongCustomerAuthentication
	return toSerialize, nil
}

type NullableRegisterSCARequest struct {
	value *RegisterSCARequest
	isSet bool
}

func (v NullableRegisterSCARequest) Get() *RegisterSCARequest {
	return v.value
}

func (v *NullableRegisterSCARequest) Set(val *RegisterSCARequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRegisterSCARequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRegisterSCARequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegisterSCARequest(val *RegisterSCARequest) *NullableRegisterSCARequest {
	return &NullableRegisterSCARequest{value: val, isSet: true}
}

func (v NullableRegisterSCARequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegisterSCARequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
