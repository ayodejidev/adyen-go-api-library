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

// checks if the PaymentDetails type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PaymentDetails{}

// PaymentDetails struct for PaymentDetails
type PaymentDetails struct {
	// The checkout attempt identifier.
	CheckoutAttemptId *string `json:"checkoutAttemptId,omitempty"`
	// The payment method type.
	Type *string `json:"type,omitempty"`
}

// NewPaymentDetails instantiates a new PaymentDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentDetails() *PaymentDetails {
	this := PaymentDetails{}
	return &this
}

// NewPaymentDetailsWithDefaults instantiates a new PaymentDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentDetailsWithDefaults() *PaymentDetails {
	this := PaymentDetails{}
	return &this
}

// GetCheckoutAttemptId returns the CheckoutAttemptId field value if set, zero value otherwise.
func (o *PaymentDetails) GetCheckoutAttemptId() string {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		var ret string
		return ret
	}
	return *o.CheckoutAttemptId
}

// GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentDetails) GetCheckoutAttemptIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		return nil, false
	}
	return o.CheckoutAttemptId, true
}

// HasCheckoutAttemptId returns a boolean if a field has been set.
func (o *PaymentDetails) HasCheckoutAttemptId() bool {
	if o != nil && !common.IsNil(o.CheckoutAttemptId) {
		return true
	}

	return false
}

// SetCheckoutAttemptId gets a reference to the given string and assigns it to the CheckoutAttemptId field.
func (o *PaymentDetails) SetCheckoutAttemptId(v string) {
	o.CheckoutAttemptId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PaymentDetails) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentDetails) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PaymentDetails) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PaymentDetails) SetType(v string) {
	o.Type = &v
}

func (o PaymentDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CheckoutAttemptId) {
		toSerialize["checkoutAttemptId"] = o.CheckoutAttemptId
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullablePaymentDetails struct {
	value *PaymentDetails
	isSet bool
}

func (v NullablePaymentDetails) Get() *PaymentDetails {
	return v.value
}

func (v *NullablePaymentDetails) Set(val *PaymentDetails) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentDetails) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentDetails(val *PaymentDetails) *NullablePaymentDetails {
	return &NullablePaymentDetails{value: val, isSet: true}
}

func (v NullablePaymentDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *PaymentDetails) isValidType() bool {
	var allowedEnumValues = []string{"alipay", "multibanco", "bankTransfer_IBAN", "paybright", "paynow", "affirm", "affirm_pos", "trustly", "trustlyvector", "oney", "facilypay", "facilypay_3x", "facilypay_4x", "facilypay_6x", "facilypay_10x", "facilypay_12x", "unionpay", "kcp_banktransfer", "kcp_payco", "kcp_creditcard", "wechatpaySDK", "wechatpayQR", "wechatpayWeb", "molpay_boost", "wallet_IN", "payu_IN_cashcard", "payu_IN_nb", "upi_qr", "paytm", "molpay_ebanking_VN", "molpay_ebanking_MY", "molpay_ebanking_direct_MY", "swish", "pix", "bizum", "walley", "walley_b2b", "alma", "paypo", "scalapay", "scalapay_3x", "scalapay_4x", "molpay_fpx", "konbini", "directEbanking", "boletobancario", "neteller", "paysafecard", "cashticket", "ikano", "karenmillen", "oasis", "warehouse", "primeiropay_boleto", "mada", "benefit", "knet", "omannet", "gopay_wallet", "kcp_naverpay", "onlinebanking_IN", "fawry", "atome", "moneybookers", "naps", "nordea", "boletobancario_bradesco", "boletobancario_itau", "boletobancario_santander", "boletobancario_bancodobrasil", "boletobancario_hsbc", "molpay_maybank2u", "molpay_cimb", "molpay_rhb", "molpay_amb", "molpay_hlb", "molpay_affin_epg", "molpay_bankislam", "molpay_publicbank", "fpx_agrobank", "touchngo", "maybank2u_mae", "duitnow", "promptpay", "twint_pos", "alipay_hk", "alipay_hk_web", "alipay_hk_wap", "alipay_wap", "balanceplatform"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
