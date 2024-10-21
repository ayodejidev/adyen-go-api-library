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

// checks if the CheckoutQrCodeAction type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CheckoutQrCodeAction{}

// CheckoutQrCodeAction struct for CheckoutQrCodeAction
type CheckoutQrCodeAction struct {
	// Expiry time of the QR code.
	ExpiresAt *string `json:"expiresAt,omitempty"`
	// Encoded payment data.
	PaymentData *string `json:"paymentData,omitempty"`
	// Specifies the payment method.
	PaymentMethodType *string `json:"paymentMethodType,omitempty"`
	// The contents of the QR code as a UTF8 string.
	QrCodeData *string `json:"qrCodeData,omitempty"`
	// **qrCode**
	Type string `json:"type"`
	// Specifies the URL to redirect to.
	Url *string `json:"url,omitempty"`
}

// NewCheckoutQrCodeAction instantiates a new CheckoutQrCodeAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutQrCodeAction(type_ string) *CheckoutQrCodeAction {
	this := CheckoutQrCodeAction{}
	this.Type = type_
	return &this
}

// NewCheckoutQrCodeActionWithDefaults instantiates a new CheckoutQrCodeAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutQrCodeActionWithDefaults() *CheckoutQrCodeAction {
	this := CheckoutQrCodeAction{}
	return &this
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *CheckoutQrCodeAction) GetExpiresAt() string {
	if o == nil || common.IsNil(o.ExpiresAt) {
		var ret string
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutQrCodeAction) GetExpiresAtOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *CheckoutQrCodeAction) HasExpiresAt() bool {
	if o != nil && !common.IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given string and assigns it to the ExpiresAt field.
func (o *CheckoutQrCodeAction) SetExpiresAt(v string) {
	o.ExpiresAt = &v
}

// GetPaymentData returns the PaymentData field value if set, zero value otherwise.
func (o *CheckoutQrCodeAction) GetPaymentData() string {
	if o == nil || common.IsNil(o.PaymentData) {
		var ret string
		return ret
	}
	return *o.PaymentData
}

// GetPaymentDataOk returns a tuple with the PaymentData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutQrCodeAction) GetPaymentDataOk() (*string, bool) {
	if o == nil || common.IsNil(o.PaymentData) {
		return nil, false
	}
	return o.PaymentData, true
}

// HasPaymentData returns a boolean if a field has been set.
func (o *CheckoutQrCodeAction) HasPaymentData() bool {
	if o != nil && !common.IsNil(o.PaymentData) {
		return true
	}

	return false
}

// SetPaymentData gets a reference to the given string and assigns it to the PaymentData field.
func (o *CheckoutQrCodeAction) SetPaymentData(v string) {
	o.PaymentData = &v
}

// GetPaymentMethodType returns the PaymentMethodType field value if set, zero value otherwise.
func (o *CheckoutQrCodeAction) GetPaymentMethodType() string {
	if o == nil || common.IsNil(o.PaymentMethodType) {
		var ret string
		return ret
	}
	return *o.PaymentMethodType
}

// GetPaymentMethodTypeOk returns a tuple with the PaymentMethodType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutQrCodeAction) GetPaymentMethodTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.PaymentMethodType) {
		return nil, false
	}
	return o.PaymentMethodType, true
}

// HasPaymentMethodType returns a boolean if a field has been set.
func (o *CheckoutQrCodeAction) HasPaymentMethodType() bool {
	if o != nil && !common.IsNil(o.PaymentMethodType) {
		return true
	}

	return false
}

// SetPaymentMethodType gets a reference to the given string and assigns it to the PaymentMethodType field.
func (o *CheckoutQrCodeAction) SetPaymentMethodType(v string) {
	o.PaymentMethodType = &v
}

// GetQrCodeData returns the QrCodeData field value if set, zero value otherwise.
func (o *CheckoutQrCodeAction) GetQrCodeData() string {
	if o == nil || common.IsNil(o.QrCodeData) {
		var ret string
		return ret
	}
	return *o.QrCodeData
}

// GetQrCodeDataOk returns a tuple with the QrCodeData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutQrCodeAction) GetQrCodeDataOk() (*string, bool) {
	if o == nil || common.IsNil(o.QrCodeData) {
		return nil, false
	}
	return o.QrCodeData, true
}

// HasQrCodeData returns a boolean if a field has been set.
func (o *CheckoutQrCodeAction) HasQrCodeData() bool {
	if o != nil && !common.IsNil(o.QrCodeData) {
		return true
	}

	return false
}

// SetQrCodeData gets a reference to the given string and assigns it to the QrCodeData field.
func (o *CheckoutQrCodeAction) SetQrCodeData(v string) {
	o.QrCodeData = &v
}

// GetType returns the Type field value
func (o *CheckoutQrCodeAction) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CheckoutQrCodeAction) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CheckoutQrCodeAction) SetType(v string) {
	o.Type = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *CheckoutQrCodeAction) GetUrl() string {
	if o == nil || common.IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutQrCodeAction) GetUrlOk() (*string, bool) {
	if o == nil || common.IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *CheckoutQrCodeAction) HasUrl() bool {
	if o != nil && !common.IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *CheckoutQrCodeAction) SetUrl(v string) {
	o.Url = &v
}

func (o CheckoutQrCodeAction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckoutQrCodeAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ExpiresAt) {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	if !common.IsNil(o.PaymentData) {
		toSerialize["paymentData"] = o.PaymentData
	}
	if !common.IsNil(o.PaymentMethodType) {
		toSerialize["paymentMethodType"] = o.PaymentMethodType
	}
	if !common.IsNil(o.QrCodeData) {
		toSerialize["qrCodeData"] = o.QrCodeData
	}
	toSerialize["type"] = o.Type
	if !common.IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableCheckoutQrCodeAction struct {
	value *CheckoutQrCodeAction
	isSet bool
}

func (v NullableCheckoutQrCodeAction) Get() *CheckoutQrCodeAction {
	return v.value
}

func (v *NullableCheckoutQrCodeAction) Set(val *CheckoutQrCodeAction) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutQrCodeAction) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutQrCodeAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutQrCodeAction(val *CheckoutQrCodeAction) *NullableCheckoutQrCodeAction {
	return &NullableCheckoutQrCodeAction{value: val, isSet: true}
}

func (v NullableCheckoutQrCodeAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutQrCodeAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *CheckoutQrCodeAction) isValidType() bool {
	var allowedEnumValues = []string{"qrCode"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
