/*
Adyen Recurring API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package recurring

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v11/src/common"
)

// checks if the NotifyShopperRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &NotifyShopperRequest{}

// NotifyShopperRequest struct for NotifyShopperRequest
type NotifyShopperRequest struct {
	Amount Amount `json:"amount"`
	// Date on which the subscription amount will be debited from the shopper. In YYYY-MM-DD format
	BillingDate *string `json:"billingDate,omitempty"`
	// Sequence of the debit. Depends on Frequency and Billing Attempts Rule.
	BillingSequenceNumber *string `json:"billingSequenceNumber,omitempty"`
	// Reference of Pre-debit notification that is displayed to the shopper. Optional field. Maps to reference if missing
	DisplayedReference *string `json:"displayedReference,omitempty"`
	// The merchant account identifier with which you want to process the transaction.
	MerchantAccount string `json:"merchantAccount"`
	// This is the `recurringDetailReference` returned in the response when you created the token.
	RecurringDetailReference *string `json:"recurringDetailReference,omitempty"`
	// Pre-debit notification reference sent by the merchant. This is a mandatory field
	Reference string `json:"reference"`
	// The ID that uniquely identifies the shopper.  This `shopperReference` must be the same as the `shopperReference` used in the initial payment.
	ShopperReference string `json:"shopperReference"`
	// This is the `recurringDetailReference` returned in the response when you created the token.
	StoredPaymentMethodId *string `json:"storedPaymentMethodId,omitempty"`
}

// NewNotifyShopperRequest instantiates a new NotifyShopperRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotifyShopperRequest(amount Amount, merchantAccount string, reference string, shopperReference string) *NotifyShopperRequest {
	this := NotifyShopperRequest{}
	this.Amount = amount
	this.MerchantAccount = merchantAccount
	this.Reference = reference
	this.ShopperReference = shopperReference
	return &this
}

// NewNotifyShopperRequestWithDefaults instantiates a new NotifyShopperRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotifyShopperRequestWithDefaults() *NotifyShopperRequest {
	this := NotifyShopperRequest{}
	return &this
}

// GetAmount returns the Amount field value
func (o *NotifyShopperRequest) GetAmount() Amount {
	if o == nil {
		var ret Amount
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *NotifyShopperRequest) GetAmountOk() (*Amount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *NotifyShopperRequest) SetAmount(v Amount) {
	o.Amount = v
}

// GetBillingDate returns the BillingDate field value if set, zero value otherwise.
func (o *NotifyShopperRequest) GetBillingDate() string {
	if o == nil || common.IsNil(o.BillingDate) {
		var ret string
		return ret
	}
	return *o.BillingDate
}

// GetBillingDateOk returns a tuple with the BillingDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyShopperRequest) GetBillingDateOk() (*string, bool) {
	if o == nil || common.IsNil(o.BillingDate) {
		return nil, false
	}
	return o.BillingDate, true
}

// HasBillingDate returns a boolean if a field has been set.
func (o *NotifyShopperRequest) HasBillingDate() bool {
	if o != nil && !common.IsNil(o.BillingDate) {
		return true
	}

	return false
}

// SetBillingDate gets a reference to the given string and assigns it to the BillingDate field.
func (o *NotifyShopperRequest) SetBillingDate(v string) {
	o.BillingDate = &v
}

// GetBillingSequenceNumber returns the BillingSequenceNumber field value if set, zero value otherwise.
func (o *NotifyShopperRequest) GetBillingSequenceNumber() string {
	if o == nil || common.IsNil(o.BillingSequenceNumber) {
		var ret string
		return ret
	}
	return *o.BillingSequenceNumber
}

// GetBillingSequenceNumberOk returns a tuple with the BillingSequenceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyShopperRequest) GetBillingSequenceNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.BillingSequenceNumber) {
		return nil, false
	}
	return o.BillingSequenceNumber, true
}

// HasBillingSequenceNumber returns a boolean if a field has been set.
func (o *NotifyShopperRequest) HasBillingSequenceNumber() bool {
	if o != nil && !common.IsNil(o.BillingSequenceNumber) {
		return true
	}

	return false
}

// SetBillingSequenceNumber gets a reference to the given string and assigns it to the BillingSequenceNumber field.
func (o *NotifyShopperRequest) SetBillingSequenceNumber(v string) {
	o.BillingSequenceNumber = &v
}

// GetDisplayedReference returns the DisplayedReference field value if set, zero value otherwise.
func (o *NotifyShopperRequest) GetDisplayedReference() string {
	if o == nil || common.IsNil(o.DisplayedReference) {
		var ret string
		return ret
	}
	return *o.DisplayedReference
}

// GetDisplayedReferenceOk returns a tuple with the DisplayedReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyShopperRequest) GetDisplayedReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.DisplayedReference) {
		return nil, false
	}
	return o.DisplayedReference, true
}

// HasDisplayedReference returns a boolean if a field has been set.
func (o *NotifyShopperRequest) HasDisplayedReference() bool {
	if o != nil && !common.IsNil(o.DisplayedReference) {
		return true
	}

	return false
}

// SetDisplayedReference gets a reference to the given string and assigns it to the DisplayedReference field.
func (o *NotifyShopperRequest) SetDisplayedReference(v string) {
	o.DisplayedReference = &v
}

// GetMerchantAccount returns the MerchantAccount field value
func (o *NotifyShopperRequest) GetMerchantAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value
// and a boolean to check if the value has been set.
func (o *NotifyShopperRequest) GetMerchantAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantAccount, true
}

// SetMerchantAccount sets field value
func (o *NotifyShopperRequest) SetMerchantAccount(v string) {
	o.MerchantAccount = v
}

// GetRecurringDetailReference returns the RecurringDetailReference field value if set, zero value otherwise.
func (o *NotifyShopperRequest) GetRecurringDetailReference() string {
	if o == nil || common.IsNil(o.RecurringDetailReference) {
		var ret string
		return ret
	}
	return *o.RecurringDetailReference
}

// GetRecurringDetailReferenceOk returns a tuple with the RecurringDetailReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyShopperRequest) GetRecurringDetailReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.RecurringDetailReference) {
		return nil, false
	}
	return o.RecurringDetailReference, true
}

// HasRecurringDetailReference returns a boolean if a field has been set.
func (o *NotifyShopperRequest) HasRecurringDetailReference() bool {
	if o != nil && !common.IsNil(o.RecurringDetailReference) {
		return true
	}

	return false
}

// SetRecurringDetailReference gets a reference to the given string and assigns it to the RecurringDetailReference field.
func (o *NotifyShopperRequest) SetRecurringDetailReference(v string) {
	o.RecurringDetailReference = &v
}

// GetReference returns the Reference field value
func (o *NotifyShopperRequest) GetReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
func (o *NotifyShopperRequest) GetReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reference, true
}

// SetReference sets field value
func (o *NotifyShopperRequest) SetReference(v string) {
	o.Reference = v
}

// GetShopperReference returns the ShopperReference field value
func (o *NotifyShopperRequest) GetShopperReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShopperReference
}

// GetShopperReferenceOk returns a tuple with the ShopperReference field value
// and a boolean to check if the value has been set.
func (o *NotifyShopperRequest) GetShopperReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShopperReference, true
}

// SetShopperReference sets field value
func (o *NotifyShopperRequest) SetShopperReference(v string) {
	o.ShopperReference = v
}

// GetStoredPaymentMethodId returns the StoredPaymentMethodId field value if set, zero value otherwise.
func (o *NotifyShopperRequest) GetStoredPaymentMethodId() string {
	if o == nil || common.IsNil(o.StoredPaymentMethodId) {
		var ret string
		return ret
	}
	return *o.StoredPaymentMethodId
}

// GetStoredPaymentMethodIdOk returns a tuple with the StoredPaymentMethodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyShopperRequest) GetStoredPaymentMethodIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.StoredPaymentMethodId) {
		return nil, false
	}
	return o.StoredPaymentMethodId, true
}

// HasStoredPaymentMethodId returns a boolean if a field has been set.
func (o *NotifyShopperRequest) HasStoredPaymentMethodId() bool {
	if o != nil && !common.IsNil(o.StoredPaymentMethodId) {
		return true
	}

	return false
}

// SetStoredPaymentMethodId gets a reference to the given string and assigns it to the StoredPaymentMethodId field.
func (o *NotifyShopperRequest) SetStoredPaymentMethodId(v string) {
	o.StoredPaymentMethodId = &v
}

func (o NotifyShopperRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotifyShopperRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	if !common.IsNil(o.BillingDate) {
		toSerialize["billingDate"] = o.BillingDate
	}
	if !common.IsNil(o.BillingSequenceNumber) {
		toSerialize["billingSequenceNumber"] = o.BillingSequenceNumber
	}
	if !common.IsNil(o.DisplayedReference) {
		toSerialize["displayedReference"] = o.DisplayedReference
	}
	toSerialize["merchantAccount"] = o.MerchantAccount
	if !common.IsNil(o.RecurringDetailReference) {
		toSerialize["recurringDetailReference"] = o.RecurringDetailReference
	}
	toSerialize["reference"] = o.Reference
	toSerialize["shopperReference"] = o.ShopperReference
	if !common.IsNil(o.StoredPaymentMethodId) {
		toSerialize["storedPaymentMethodId"] = o.StoredPaymentMethodId
	}
	return toSerialize, nil
}

type NullableNotifyShopperRequest struct {
	value *NotifyShopperRequest
	isSet bool
}

func (v NullableNotifyShopperRequest) Get() *NotifyShopperRequest {
	return v.value
}

func (v *NullableNotifyShopperRequest) Set(val *NotifyShopperRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNotifyShopperRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNotifyShopperRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotifyShopperRequest(val *NotifyShopperRequest) *NullableNotifyShopperRequest {
	return &NullableNotifyShopperRequest{value: val, isSet: true}
}

func (v NullableNotifyShopperRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotifyShopperRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
