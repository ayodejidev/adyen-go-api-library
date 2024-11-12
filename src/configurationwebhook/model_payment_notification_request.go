/*
Configuration webhooks

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationwebhook

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v14/src/common"
)

// checks if the PaymentNotificationRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PaymentNotificationRequest{}

// PaymentNotificationRequest struct for PaymentNotificationRequest
type PaymentNotificationRequest struct {
	Data PaymentInstrumentNotificationData `json:"data"`
	// The environment from which the webhook originated.  Possible values: **test**, **live**.
	Environment string `json:"environment"`
	// Type of webhook.
	Type string `json:"type"`
}

// NewPaymentNotificationRequest instantiates a new PaymentNotificationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentNotificationRequest(data PaymentInstrumentNotificationData, environment string, type_ string) *PaymentNotificationRequest {
	this := PaymentNotificationRequest{}
	this.Data = data
	this.Environment = environment
	this.Type = type_
	return &this
}

// NewPaymentNotificationRequestWithDefaults instantiates a new PaymentNotificationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentNotificationRequestWithDefaults() *PaymentNotificationRequest {
	this := PaymentNotificationRequest{}
	return &this
}

// GetData returns the Data field value
func (o *PaymentNotificationRequest) GetData() PaymentInstrumentNotificationData {
	if o == nil {
		var ret PaymentInstrumentNotificationData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PaymentNotificationRequest) GetDataOk() (*PaymentInstrumentNotificationData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PaymentNotificationRequest) SetData(v PaymentInstrumentNotificationData) {
	o.Data = v
}

// GetEnvironment returns the Environment field value
func (o *PaymentNotificationRequest) GetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *PaymentNotificationRequest) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *PaymentNotificationRequest) SetEnvironment(v string) {
	o.Environment = v
}

// GetType returns the Type field value
func (o *PaymentNotificationRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PaymentNotificationRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PaymentNotificationRequest) SetType(v string) {
	o.Type = v
}

func (o PaymentNotificationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentNotificationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["environment"] = o.Environment
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullablePaymentNotificationRequest struct {
	value *PaymentNotificationRequest
	isSet bool
}

func (v NullablePaymentNotificationRequest) Get() *PaymentNotificationRequest {
	return v.value
}

func (v *NullablePaymentNotificationRequest) Set(val *PaymentNotificationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentNotificationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentNotificationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentNotificationRequest(val *PaymentNotificationRequest) *NullablePaymentNotificationRequest {
	return &NullablePaymentNotificationRequest{value: val, isSet: true}
}

func (v NullablePaymentNotificationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentNotificationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *PaymentNotificationRequest) isValidType() bool {
	var allowedEnumValues = []string{"balancePlatform.paymentInstrument.created", "balancePlatform.paymentInstrument.updated"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
