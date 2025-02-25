/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v19/src/common"
)

// checks if the PaymentAmountUpdateResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PaymentAmountUpdateResponse{}

// PaymentAmountUpdateResponse struct for PaymentAmountUpdateResponse
type PaymentAmountUpdateResponse struct {
	Amount Amount `json:"amount"`
	// The reason for the amount update. Possible values:  * **delayedCharge**  * **noShow**  * **installment**
	IndustryUsage *string `json:"industryUsage,omitempty"`
	// Price and product information of the refunded items, required for [partial refunds](https://docs.adyen.com/online-payments/refund#refund-a-payment). > This field is required for partial refunds with 3x 4x Oney, Affirm, Afterpay, Atome, Clearpay, Klarna, Ratepay, Walley, and Zip.
	LineItems []LineItem `json:"lineItems,omitempty"`
	// The merchant account that is used to process the payment.
	MerchantAccount string `json:"merchantAccount"`
	// The [`pspReference`](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments__resParam_pspReference) of the payment to update.
	PaymentPspReference string `json:"paymentPspReference"`
	// Adyen's 16-character reference associated with the amount update request.
	PspReference string `json:"pspReference"`
	// Your reference for the amount update request. Maximum length: 80 characters.
	Reference string `json:"reference"`
	// An array of objects specifying how the amount should be split between accounts when using Adyen for Platforms. For more information, see how to process payments for [marketplaces](https://docs.adyen.com/marketplaces/process-payments) or [platforms](https://docs.adyen.com/platforms/process-payments).
	Splits []Split `json:"splits,omitempty"`
	// The status of your request. This will always have the value **received**.
	Status string `json:"status"`
}

// NewPaymentAmountUpdateResponse instantiates a new PaymentAmountUpdateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentAmountUpdateResponse(amount Amount, merchantAccount string, paymentPspReference string, pspReference string, reference string, status string) *PaymentAmountUpdateResponse {
	this := PaymentAmountUpdateResponse{}
	this.Amount = amount
	this.MerchantAccount = merchantAccount
	this.PaymentPspReference = paymentPspReference
	this.PspReference = pspReference
	this.Reference = reference
	this.Status = status
	return &this
}

// NewPaymentAmountUpdateResponseWithDefaults instantiates a new PaymentAmountUpdateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentAmountUpdateResponseWithDefaults() *PaymentAmountUpdateResponse {
	this := PaymentAmountUpdateResponse{}
	return &this
}

// GetAmount returns the Amount field value
func (o *PaymentAmountUpdateResponse) GetAmount() Amount {
	if o == nil {
		var ret Amount
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *PaymentAmountUpdateResponse) GetAmountOk() (*Amount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *PaymentAmountUpdateResponse) SetAmount(v Amount) {
	o.Amount = v
}

// GetIndustryUsage returns the IndustryUsage field value if set, zero value otherwise.
func (o *PaymentAmountUpdateResponse) GetIndustryUsage() string {
	if o == nil || common.IsNil(o.IndustryUsage) {
		var ret string
		return ret
	}
	return *o.IndustryUsage
}

// GetIndustryUsageOk returns a tuple with the IndustryUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentAmountUpdateResponse) GetIndustryUsageOk() (*string, bool) {
	if o == nil || common.IsNil(o.IndustryUsage) {
		return nil, false
	}
	return o.IndustryUsage, true
}

// HasIndustryUsage returns a boolean if a field has been set.
func (o *PaymentAmountUpdateResponse) HasIndustryUsage() bool {
	if o != nil && !common.IsNil(o.IndustryUsage) {
		return true
	}

	return false
}

// SetIndustryUsage gets a reference to the given string and assigns it to the IndustryUsage field.
func (o *PaymentAmountUpdateResponse) SetIndustryUsage(v string) {
	o.IndustryUsage = &v
}

// GetLineItems returns the LineItems field value if set, zero value otherwise.
func (o *PaymentAmountUpdateResponse) GetLineItems() []LineItem {
	if o == nil || common.IsNil(o.LineItems) {
		var ret []LineItem
		return ret
	}
	return o.LineItems
}

// GetLineItemsOk returns a tuple with the LineItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentAmountUpdateResponse) GetLineItemsOk() ([]LineItem, bool) {
	if o == nil || common.IsNil(o.LineItems) {
		return nil, false
	}
	return o.LineItems, true
}

// HasLineItems returns a boolean if a field has been set.
func (o *PaymentAmountUpdateResponse) HasLineItems() bool {
	if o != nil && !common.IsNil(o.LineItems) {
		return true
	}

	return false
}

// SetLineItems gets a reference to the given []LineItem and assigns it to the LineItems field.
func (o *PaymentAmountUpdateResponse) SetLineItems(v []LineItem) {
	o.LineItems = v
}

// GetMerchantAccount returns the MerchantAccount field value
func (o *PaymentAmountUpdateResponse) GetMerchantAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value
// and a boolean to check if the value has been set.
func (o *PaymentAmountUpdateResponse) GetMerchantAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantAccount, true
}

// SetMerchantAccount sets field value
func (o *PaymentAmountUpdateResponse) SetMerchantAccount(v string) {
	o.MerchantAccount = v
}

// GetPaymentPspReference returns the PaymentPspReference field value
func (o *PaymentAmountUpdateResponse) GetPaymentPspReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentPspReference
}

// GetPaymentPspReferenceOk returns a tuple with the PaymentPspReference field value
// and a boolean to check if the value has been set.
func (o *PaymentAmountUpdateResponse) GetPaymentPspReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentPspReference, true
}

// SetPaymentPspReference sets field value
func (o *PaymentAmountUpdateResponse) SetPaymentPspReference(v string) {
	o.PaymentPspReference = v
}

// GetPspReference returns the PspReference field value
func (o *PaymentAmountUpdateResponse) GetPspReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PspReference
}

// GetPspReferenceOk returns a tuple with the PspReference field value
// and a boolean to check if the value has been set.
func (o *PaymentAmountUpdateResponse) GetPspReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PspReference, true
}

// SetPspReference sets field value
func (o *PaymentAmountUpdateResponse) SetPspReference(v string) {
	o.PspReference = v
}

// GetReference returns the Reference field value
func (o *PaymentAmountUpdateResponse) GetReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
func (o *PaymentAmountUpdateResponse) GetReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reference, true
}

// SetReference sets field value
func (o *PaymentAmountUpdateResponse) SetReference(v string) {
	o.Reference = v
}

// GetSplits returns the Splits field value if set, zero value otherwise.
func (o *PaymentAmountUpdateResponse) GetSplits() []Split {
	if o == nil || common.IsNil(o.Splits) {
		var ret []Split
		return ret
	}
	return o.Splits
}

// GetSplitsOk returns a tuple with the Splits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentAmountUpdateResponse) GetSplitsOk() ([]Split, bool) {
	if o == nil || common.IsNil(o.Splits) {
		return nil, false
	}
	return o.Splits, true
}

// HasSplits returns a boolean if a field has been set.
func (o *PaymentAmountUpdateResponse) HasSplits() bool {
	if o != nil && !common.IsNil(o.Splits) {
		return true
	}

	return false
}

// SetSplits gets a reference to the given []Split and assigns it to the Splits field.
func (o *PaymentAmountUpdateResponse) SetSplits(v []Split) {
	o.Splits = v
}

// GetStatus returns the Status field value
func (o *PaymentAmountUpdateResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PaymentAmountUpdateResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PaymentAmountUpdateResponse) SetStatus(v string) {
	o.Status = v
}

func (o PaymentAmountUpdateResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentAmountUpdateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	if !common.IsNil(o.IndustryUsage) {
		toSerialize["industryUsage"] = o.IndustryUsage
	}
	if !common.IsNil(o.LineItems) {
		toSerialize["lineItems"] = o.LineItems
	}
	toSerialize["merchantAccount"] = o.MerchantAccount
	toSerialize["paymentPspReference"] = o.PaymentPspReference
	toSerialize["pspReference"] = o.PspReference
	toSerialize["reference"] = o.Reference
	if !common.IsNil(o.Splits) {
		toSerialize["splits"] = o.Splits
	}
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

type NullablePaymentAmountUpdateResponse struct {
	value *PaymentAmountUpdateResponse
	isSet bool
}

func (v NullablePaymentAmountUpdateResponse) Get() *PaymentAmountUpdateResponse {
	return v.value
}

func (v *NullablePaymentAmountUpdateResponse) Set(val *PaymentAmountUpdateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentAmountUpdateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentAmountUpdateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentAmountUpdateResponse(val *PaymentAmountUpdateResponse) *NullablePaymentAmountUpdateResponse {
	return &NullablePaymentAmountUpdateResponse{value: val, isSet: true}
}

func (v NullablePaymentAmountUpdateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentAmountUpdateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *PaymentAmountUpdateResponse) isValidIndustryUsage() bool {
	var allowedEnumValues = []string{"delayedCharge", "installment", "noShow"}
	for _, allowed := range allowedEnumValues {
		if o.GetIndustryUsage() == allowed {
			return true
		}
	}
	return false
}
func (o *PaymentAmountUpdateResponse) isValidStatus() bool {
	var allowedEnumValues = []string{"received"}
	for _, allowed := range allowedEnumValues {
		if o.GetStatus() == allowed {
			return true
		}
	}
	return false
}
