/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v11/src/common"
)

// checks if the CheckoutSDKAction type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CheckoutSDKAction{}

// CheckoutSDKAction struct for CheckoutSDKAction
type CheckoutSDKAction struct {
	// Encoded payment data.
	PaymentData *string `json:"paymentData,omitempty"`
	// Specifies the payment method.
	PaymentMethodType *string `json:"paymentMethodType,omitempty"`
	// The data to pass to the SDK.
	SdkData *map[string]string `json:"sdkData,omitempty"`
	// The type of the action.
	Type string `json:"type"`
	// Specifies the URL to redirect to.
	Url *string `json:"url,omitempty"`
}

// NewCheckoutSDKAction instantiates a new CheckoutSDKAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutSDKAction(type_ string) *CheckoutSDKAction {
	this := CheckoutSDKAction{}
	this.Type = type_
	return &this
}

// NewCheckoutSDKActionWithDefaults instantiates a new CheckoutSDKAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutSDKActionWithDefaults() *CheckoutSDKAction {
	this := CheckoutSDKAction{}
	return &this
}

// GetPaymentData returns the PaymentData field value if set, zero value otherwise.
func (o *CheckoutSDKAction) GetPaymentData() string {
	if o == nil || common.IsNil(o.PaymentData) {
		var ret string
		return ret
	}
	return *o.PaymentData
}

// GetPaymentDataOk returns a tuple with the PaymentData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSDKAction) GetPaymentDataOk() (*string, bool) {
	if o == nil || common.IsNil(o.PaymentData) {
		return nil, false
	}
	return o.PaymentData, true
}

// HasPaymentData returns a boolean if a field has been set.
func (o *CheckoutSDKAction) HasPaymentData() bool {
	if o != nil && !common.IsNil(o.PaymentData) {
		return true
	}

	return false
}

// SetPaymentData gets a reference to the given string and assigns it to the PaymentData field.
func (o *CheckoutSDKAction) SetPaymentData(v string) {
	o.PaymentData = &v
}

// GetPaymentMethodType returns the PaymentMethodType field value if set, zero value otherwise.
func (o *CheckoutSDKAction) GetPaymentMethodType() string {
	if o == nil || common.IsNil(o.PaymentMethodType) {
		var ret string
		return ret
	}
	return *o.PaymentMethodType
}

// GetPaymentMethodTypeOk returns a tuple with the PaymentMethodType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSDKAction) GetPaymentMethodTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.PaymentMethodType) {
		return nil, false
	}
	return o.PaymentMethodType, true
}

// HasPaymentMethodType returns a boolean if a field has been set.
func (o *CheckoutSDKAction) HasPaymentMethodType() bool {
	if o != nil && !common.IsNil(o.PaymentMethodType) {
		return true
	}

	return false
}

// SetPaymentMethodType gets a reference to the given string and assigns it to the PaymentMethodType field.
func (o *CheckoutSDKAction) SetPaymentMethodType(v string) {
	o.PaymentMethodType = &v
}

// GetSdkData returns the SdkData field value if set, zero value otherwise.
func (o *CheckoutSDKAction) GetSdkData() map[string]string {
	if o == nil || common.IsNil(o.SdkData) {
		var ret map[string]string
		return ret
	}
	return *o.SdkData
}

// GetSdkDataOk returns a tuple with the SdkData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSDKAction) GetSdkDataOk() (*map[string]string, bool) {
	if o == nil || common.IsNil(o.SdkData) {
		return nil, false
	}
	return o.SdkData, true
}

// HasSdkData returns a boolean if a field has been set.
func (o *CheckoutSDKAction) HasSdkData() bool {
	if o != nil && !common.IsNil(o.SdkData) {
		return true
	}

	return false
}

// SetSdkData gets a reference to the given map[string]string and assigns it to the SdkData field.
func (o *CheckoutSDKAction) SetSdkData(v map[string]string) {
	o.SdkData = &v
}

// GetType returns the Type field value
func (o *CheckoutSDKAction) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CheckoutSDKAction) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CheckoutSDKAction) SetType(v string) {
	o.Type = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *CheckoutSDKAction) GetUrl() string {
	if o == nil || common.IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSDKAction) GetUrlOk() (*string, bool) {
	if o == nil || common.IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *CheckoutSDKAction) HasUrl() bool {
	if o != nil && !common.IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *CheckoutSDKAction) SetUrl(v string) {
	o.Url = &v
}

func (o CheckoutSDKAction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckoutSDKAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.PaymentData) {
		toSerialize["paymentData"] = o.PaymentData
	}
	if !common.IsNil(o.PaymentMethodType) {
		toSerialize["paymentMethodType"] = o.PaymentMethodType
	}
	if !common.IsNil(o.SdkData) {
		toSerialize["sdkData"] = o.SdkData
	}
	toSerialize["type"] = o.Type
	if !common.IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableCheckoutSDKAction struct {
	value *CheckoutSDKAction
	isSet bool
}

func (v NullableCheckoutSDKAction) Get() *CheckoutSDKAction {
	return v.value
}

func (v *NullableCheckoutSDKAction) Set(val *CheckoutSDKAction) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutSDKAction) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutSDKAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutSDKAction(val *CheckoutSDKAction) *NullableCheckoutSDKAction {
	return &NullableCheckoutSDKAction{value: val, isSet: true}
}

func (v NullableCheckoutSDKAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutSDKAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *CheckoutSDKAction) isValidType() bool {
	var allowedEnumValues = []string{"sdk", "wechatpaySDK"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
