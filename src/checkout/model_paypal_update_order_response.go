/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v12/src/common"
)

// checks if the PaypalUpdateOrderResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PaypalUpdateOrderResponse{}

// PaypalUpdateOrderResponse struct for PaypalUpdateOrderResponse
type PaypalUpdateOrderResponse struct {
	// The updated paymentData.
	PaymentData string `json:"paymentData"`
	// The status of the request. This indicates whether the order was successfully updated with PayPal.
	Status string `json:"status"`
}

// NewPaypalUpdateOrderResponse instantiates a new PaypalUpdateOrderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaypalUpdateOrderResponse(paymentData string, status string) *PaypalUpdateOrderResponse {
	this := PaypalUpdateOrderResponse{}
	this.PaymentData = paymentData
	this.Status = status
	return &this
}

// NewPaypalUpdateOrderResponseWithDefaults instantiates a new PaypalUpdateOrderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaypalUpdateOrderResponseWithDefaults() *PaypalUpdateOrderResponse {
	this := PaypalUpdateOrderResponse{}
	return &this
}

// GetPaymentData returns the PaymentData field value
func (o *PaypalUpdateOrderResponse) GetPaymentData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentData
}

// GetPaymentDataOk returns a tuple with the PaymentData field value
// and a boolean to check if the value has been set.
func (o *PaypalUpdateOrderResponse) GetPaymentDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentData, true
}

// SetPaymentData sets field value
func (o *PaypalUpdateOrderResponse) SetPaymentData(v string) {
	o.PaymentData = v
}

// GetStatus returns the Status field value
func (o *PaypalUpdateOrderResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PaypalUpdateOrderResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PaypalUpdateOrderResponse) SetStatus(v string) {
	o.Status = v
}

func (o PaypalUpdateOrderResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaypalUpdateOrderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["paymentData"] = o.PaymentData
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

type NullablePaypalUpdateOrderResponse struct {
	value *PaypalUpdateOrderResponse
	isSet bool
}

func (v NullablePaypalUpdateOrderResponse) Get() *PaypalUpdateOrderResponse {
	return v.value
}

func (v *NullablePaypalUpdateOrderResponse) Set(val *PaypalUpdateOrderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaypalUpdateOrderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaypalUpdateOrderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaypalUpdateOrderResponse(val *PaypalUpdateOrderResponse) *NullablePaypalUpdateOrderResponse {
	return &NullablePaypalUpdateOrderResponse{value: val, isSet: true}
}

func (v NullablePaypalUpdateOrderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaypalUpdateOrderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *PaypalUpdateOrderResponse) isValidStatus() bool {
	var allowedEnumValues = []string{"error", "success"}
	for _, allowed := range allowedEnumValues {
		if o.GetStatus() == allowed {
			return true
		}
	}
	return false
}
