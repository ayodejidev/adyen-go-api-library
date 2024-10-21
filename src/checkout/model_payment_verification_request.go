/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v14/src/common"
)

// checks if the PaymentVerificationRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PaymentVerificationRequest{}

// PaymentVerificationRequest struct for PaymentVerificationRequest
type PaymentVerificationRequest struct {
	// Encrypted and signed payment result data. You should receive this value from the Checkout SDK after the shopper completes the payment.
	Payload string `json:"payload"`
}

// NewPaymentVerificationRequest instantiates a new PaymentVerificationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentVerificationRequest(payload string) *PaymentVerificationRequest {
	this := PaymentVerificationRequest{}
	this.Payload = payload
	return &this
}

// NewPaymentVerificationRequestWithDefaults instantiates a new PaymentVerificationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentVerificationRequestWithDefaults() *PaymentVerificationRequest {
	this := PaymentVerificationRequest{}
	return &this
}

// GetPayload returns the Payload field value
func (o *PaymentVerificationRequest) GetPayload() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *PaymentVerificationRequest) GetPayloadOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Payload, true
}

// SetPayload sets field value
func (o *PaymentVerificationRequest) SetPayload(v string) {
	o.Payload = v
}

func (o PaymentVerificationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentVerificationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["payload"] = o.Payload
	return toSerialize, nil
}

type NullablePaymentVerificationRequest struct {
	value *PaymentVerificationRequest
	isSet bool
}

func (v NullablePaymentVerificationRequest) Get() *PaymentVerificationRequest {
	return v.value
}

func (v *NullablePaymentVerificationRequest) Set(val *PaymentVerificationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentVerificationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentVerificationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentVerificationRequest(val *PaymentVerificationRequest) *NullablePaymentVerificationRequest {
	return &NullablePaymentVerificationRequest{value: val, isSet: true}
}

func (v NullablePaymentVerificationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentVerificationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
