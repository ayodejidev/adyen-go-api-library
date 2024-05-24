/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v9/src/common"
)

// checks if the PaymentMethodToStore type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PaymentMethodToStore{}

// PaymentMethodToStore struct for PaymentMethodToStore
type PaymentMethodToStore struct {
	// Secondary brand of the card. For example: **plastix**, **hmclub**.
	Brand *string `json:"brand,omitempty"`
	// The card verification code. Only collect raw card data if you are [fully PCI compliant](https://docs.adyen.com/development-resources/pci-dss-compliance-guide).
	Cvc *string `json:"cvc,omitempty"`
	// The encrypted card number.
	EncryptedCardNumber *string `json:"encryptedCardNumber,omitempty"`
	// The encrypted card expiry month.
	EncryptedExpiryMonth *string `json:"encryptedExpiryMonth,omitempty"`
	// The encrypted card expiry year.
	EncryptedExpiryYear *string `json:"encryptedExpiryYear,omitempty"`
	// The encrypted card verification code.
	EncryptedSecurityCode *string `json:"encryptedSecurityCode,omitempty"`
	// The card expiry month. Only collect raw card data if you are [fully PCI compliant](https://docs.adyen.com/development-resources/pci-dss-compliance-guide).
	ExpiryMonth *string `json:"expiryMonth,omitempty"`
	// The card expiry year. Only collect raw card data if you are [fully PCI compliant](https://docs.adyen.com/development-resources/pci-dss-compliance-guide).
	ExpiryYear *string `json:"expiryYear,omitempty"`
	// The name of the card holder.
	HolderName *string `json:"holderName,omitempty"`
	// The card number. Only collect raw card data if you are [fully PCI compliant](https://docs.adyen.com/development-resources/pci-dss-compliance-guide).
	Number *string `json:"number,omitempty"`
	// Set to **scheme**.
	Type *string `json:"type,omitempty"`
}

// NewPaymentMethodToStore instantiates a new PaymentMethodToStore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodToStore() *PaymentMethodToStore {
	this := PaymentMethodToStore{}
	return &this
}

// NewPaymentMethodToStoreWithDefaults instantiates a new PaymentMethodToStore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodToStoreWithDefaults() *PaymentMethodToStore {
	this := PaymentMethodToStore{}
	return &this
}

// GetBrand returns the Brand field value if set, zero value otherwise.
func (o *PaymentMethodToStore) GetBrand() string {
	if o == nil || common.IsNil(o.Brand) {
		var ret string
		return ret
	}
	return *o.Brand
}

// GetBrandOk returns a tuple with the Brand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodToStore) GetBrandOk() (*string, bool) {
	if o == nil || common.IsNil(o.Brand) {
		return nil, false
	}
	return o.Brand, true
}

// HasBrand returns a boolean if a field has been set.
func (o *PaymentMethodToStore) HasBrand() bool {
	if o != nil && !common.IsNil(o.Brand) {
		return true
	}

	return false
}

// SetBrand gets a reference to the given string and assigns it to the Brand field.
func (o *PaymentMethodToStore) SetBrand(v string) {
	o.Brand = &v
}

// GetCvc returns the Cvc field value if set, zero value otherwise.
func (o *PaymentMethodToStore) GetCvc() string {
	if o == nil || common.IsNil(o.Cvc) {
		var ret string
		return ret
	}
	return *o.Cvc
}

// GetCvcOk returns a tuple with the Cvc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodToStore) GetCvcOk() (*string, bool) {
	if o == nil || common.IsNil(o.Cvc) {
		return nil, false
	}
	return o.Cvc, true
}

// HasCvc returns a boolean if a field has been set.
func (o *PaymentMethodToStore) HasCvc() bool {
	if o != nil && !common.IsNil(o.Cvc) {
		return true
	}

	return false
}

// SetCvc gets a reference to the given string and assigns it to the Cvc field.
func (o *PaymentMethodToStore) SetCvc(v string) {
	o.Cvc = &v
}

// GetEncryptedCardNumber returns the EncryptedCardNumber field value if set, zero value otherwise.
func (o *PaymentMethodToStore) GetEncryptedCardNumber() string {
	if o == nil || common.IsNil(o.EncryptedCardNumber) {
		var ret string
		return ret
	}
	return *o.EncryptedCardNumber
}

// GetEncryptedCardNumberOk returns a tuple with the EncryptedCardNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodToStore) GetEncryptedCardNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.EncryptedCardNumber) {
		return nil, false
	}
	return o.EncryptedCardNumber, true
}

// HasEncryptedCardNumber returns a boolean if a field has been set.
func (o *PaymentMethodToStore) HasEncryptedCardNumber() bool {
	if o != nil && !common.IsNil(o.EncryptedCardNumber) {
		return true
	}

	return false
}

// SetEncryptedCardNumber gets a reference to the given string and assigns it to the EncryptedCardNumber field.
func (o *PaymentMethodToStore) SetEncryptedCardNumber(v string) {
	o.EncryptedCardNumber = &v
}

// GetEncryptedExpiryMonth returns the EncryptedExpiryMonth field value if set, zero value otherwise.
func (o *PaymentMethodToStore) GetEncryptedExpiryMonth() string {
	if o == nil || common.IsNil(o.EncryptedExpiryMonth) {
		var ret string
		return ret
	}
	return *o.EncryptedExpiryMonth
}

// GetEncryptedExpiryMonthOk returns a tuple with the EncryptedExpiryMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodToStore) GetEncryptedExpiryMonthOk() (*string, bool) {
	if o == nil || common.IsNil(o.EncryptedExpiryMonth) {
		return nil, false
	}
	return o.EncryptedExpiryMonth, true
}

// HasEncryptedExpiryMonth returns a boolean if a field has been set.
func (o *PaymentMethodToStore) HasEncryptedExpiryMonth() bool {
	if o != nil && !common.IsNil(o.EncryptedExpiryMonth) {
		return true
	}

	return false
}

// SetEncryptedExpiryMonth gets a reference to the given string and assigns it to the EncryptedExpiryMonth field.
func (o *PaymentMethodToStore) SetEncryptedExpiryMonth(v string) {
	o.EncryptedExpiryMonth = &v
}

// GetEncryptedExpiryYear returns the EncryptedExpiryYear field value if set, zero value otherwise.
func (o *PaymentMethodToStore) GetEncryptedExpiryYear() string {
	if o == nil || common.IsNil(o.EncryptedExpiryYear) {
		var ret string
		return ret
	}
	return *o.EncryptedExpiryYear
}

// GetEncryptedExpiryYearOk returns a tuple with the EncryptedExpiryYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodToStore) GetEncryptedExpiryYearOk() (*string, bool) {
	if o == nil || common.IsNil(o.EncryptedExpiryYear) {
		return nil, false
	}
	return o.EncryptedExpiryYear, true
}

// HasEncryptedExpiryYear returns a boolean if a field has been set.
func (o *PaymentMethodToStore) HasEncryptedExpiryYear() bool {
	if o != nil && !common.IsNil(o.EncryptedExpiryYear) {
		return true
	}

	return false
}

// SetEncryptedExpiryYear gets a reference to the given string and assigns it to the EncryptedExpiryYear field.
func (o *PaymentMethodToStore) SetEncryptedExpiryYear(v string) {
	o.EncryptedExpiryYear = &v
}

// GetEncryptedSecurityCode returns the EncryptedSecurityCode field value if set, zero value otherwise.
func (o *PaymentMethodToStore) GetEncryptedSecurityCode() string {
	if o == nil || common.IsNil(o.EncryptedSecurityCode) {
		var ret string
		return ret
	}
	return *o.EncryptedSecurityCode
}

// GetEncryptedSecurityCodeOk returns a tuple with the EncryptedSecurityCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodToStore) GetEncryptedSecurityCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.EncryptedSecurityCode) {
		return nil, false
	}
	return o.EncryptedSecurityCode, true
}

// HasEncryptedSecurityCode returns a boolean if a field has been set.
func (o *PaymentMethodToStore) HasEncryptedSecurityCode() bool {
	if o != nil && !common.IsNil(o.EncryptedSecurityCode) {
		return true
	}

	return false
}

// SetEncryptedSecurityCode gets a reference to the given string and assigns it to the EncryptedSecurityCode field.
func (o *PaymentMethodToStore) SetEncryptedSecurityCode(v string) {
	o.EncryptedSecurityCode = &v
}

// GetExpiryMonth returns the ExpiryMonth field value if set, zero value otherwise.
func (o *PaymentMethodToStore) GetExpiryMonth() string {
	if o == nil || common.IsNil(o.ExpiryMonth) {
		var ret string
		return ret
	}
	return *o.ExpiryMonth
}

// GetExpiryMonthOk returns a tuple with the ExpiryMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodToStore) GetExpiryMonthOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExpiryMonth) {
		return nil, false
	}
	return o.ExpiryMonth, true
}

// HasExpiryMonth returns a boolean if a field has been set.
func (o *PaymentMethodToStore) HasExpiryMonth() bool {
	if o != nil && !common.IsNil(o.ExpiryMonth) {
		return true
	}

	return false
}

// SetExpiryMonth gets a reference to the given string and assigns it to the ExpiryMonth field.
func (o *PaymentMethodToStore) SetExpiryMonth(v string) {
	o.ExpiryMonth = &v
}

// GetExpiryYear returns the ExpiryYear field value if set, zero value otherwise.
func (o *PaymentMethodToStore) GetExpiryYear() string {
	if o == nil || common.IsNil(o.ExpiryYear) {
		var ret string
		return ret
	}
	return *o.ExpiryYear
}

// GetExpiryYearOk returns a tuple with the ExpiryYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodToStore) GetExpiryYearOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExpiryYear) {
		return nil, false
	}
	return o.ExpiryYear, true
}

// HasExpiryYear returns a boolean if a field has been set.
func (o *PaymentMethodToStore) HasExpiryYear() bool {
	if o != nil && !common.IsNil(o.ExpiryYear) {
		return true
	}

	return false
}

// SetExpiryYear gets a reference to the given string and assigns it to the ExpiryYear field.
func (o *PaymentMethodToStore) SetExpiryYear(v string) {
	o.ExpiryYear = &v
}

// GetHolderName returns the HolderName field value if set, zero value otherwise.
func (o *PaymentMethodToStore) GetHolderName() string {
	if o == nil || common.IsNil(o.HolderName) {
		var ret string
		return ret
	}
	return *o.HolderName
}

// GetHolderNameOk returns a tuple with the HolderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodToStore) GetHolderNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.HolderName) {
		return nil, false
	}
	return o.HolderName, true
}

// HasHolderName returns a boolean if a field has been set.
func (o *PaymentMethodToStore) HasHolderName() bool {
	if o != nil && !common.IsNil(o.HolderName) {
		return true
	}

	return false
}

// SetHolderName gets a reference to the given string and assigns it to the HolderName field.
func (o *PaymentMethodToStore) SetHolderName(v string) {
	o.HolderName = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *PaymentMethodToStore) GetNumber() string {
	if o == nil || common.IsNil(o.Number) {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodToStore) GetNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *PaymentMethodToStore) HasNumber() bool {
	if o != nil && !common.IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *PaymentMethodToStore) SetNumber(v string) {
	o.Number = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PaymentMethodToStore) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodToStore) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PaymentMethodToStore) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PaymentMethodToStore) SetType(v string) {
	o.Type = &v
}

func (o PaymentMethodToStore) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodToStore) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Brand) {
		toSerialize["brand"] = o.Brand
	}
	if !common.IsNil(o.Cvc) {
		toSerialize["cvc"] = o.Cvc
	}
	if !common.IsNil(o.EncryptedCardNumber) {
		toSerialize["encryptedCardNumber"] = o.EncryptedCardNumber
	}
	if !common.IsNil(o.EncryptedExpiryMonth) {
		toSerialize["encryptedExpiryMonth"] = o.EncryptedExpiryMonth
	}
	if !common.IsNil(o.EncryptedExpiryYear) {
		toSerialize["encryptedExpiryYear"] = o.EncryptedExpiryYear
	}
	if !common.IsNil(o.EncryptedSecurityCode) {
		toSerialize["encryptedSecurityCode"] = o.EncryptedSecurityCode
	}
	if !common.IsNil(o.ExpiryMonth) {
		toSerialize["expiryMonth"] = o.ExpiryMonth
	}
	if !common.IsNil(o.ExpiryYear) {
		toSerialize["expiryYear"] = o.ExpiryYear
	}
	if !common.IsNil(o.HolderName) {
		toSerialize["holderName"] = o.HolderName
	}
	if !common.IsNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullablePaymentMethodToStore struct {
	value *PaymentMethodToStore
	isSet bool
}

func (v NullablePaymentMethodToStore) Get() *PaymentMethodToStore {
	return v.value
}

func (v *NullablePaymentMethodToStore) Set(val *PaymentMethodToStore) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodToStore) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodToStore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodToStore(val *PaymentMethodToStore) *NullablePaymentMethodToStore {
	return &NullablePaymentMethodToStore{value: val, isSet: true}
}

func (v NullablePaymentMethodToStore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodToStore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
