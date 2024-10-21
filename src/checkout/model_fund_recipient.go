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

// checks if the FundRecipient type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &FundRecipient{}

// FundRecipient struct for FundRecipient
type FundRecipient struct {
	// Fund Recipient Iban for C2C payments
	IBAN           *string      `json:"IBAN,omitempty"`
	BillingAddress *Address     `json:"billingAddress,omitempty"`
	PaymentMethod  *CardDetails `json:"paymentMethod,omitempty"`
	// The email address of the shopper.
	ShopperEmail *string `json:"shopperEmail,omitempty"`
	ShopperName  *Name   `json:"shopperName,omitempty"`
	// Required for recurring payments.  Your reference to uniquely identify this shopper, for example user ID or account ID. Minimum length: 3 characters. > Your reference must not include personally identifiable information (PII), for example name or email address.
	ShopperReference *string `json:"shopperReference,omitempty"`
	// This is the `recurringDetailReference` returned in the response when you created the token.
	StoredPaymentMethodId *string      `json:"storedPaymentMethodId,omitempty"`
	SubMerchant           *SubMerchant `json:"subMerchant,omitempty"`
	// The telephone number of the shopper.
	TelephoneNumber *string `json:"telephoneNumber,omitempty"`
	// Indicates where the money is going.
	WalletIdentifier *string `json:"walletIdentifier,omitempty"`
	// Indicates the tax identifier of the fund recipient
	WalletOwnerTaxId *string `json:"walletOwnerTaxId,omitempty"`
	// The purpose of a digital wallet transaction
	WalletPurpose *string `json:"walletPurpose,omitempty"`
}

// NewFundRecipient instantiates a new FundRecipient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFundRecipient() *FundRecipient {
	this := FundRecipient{}
	return &this
}

// NewFundRecipientWithDefaults instantiates a new FundRecipient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFundRecipientWithDefaults() *FundRecipient {
	this := FundRecipient{}
	return &this
}

// GetIBAN returns the IBAN field value if set, zero value otherwise.
func (o *FundRecipient) GetIBAN() string {
	if o == nil || common.IsNil(o.IBAN) {
		var ret string
		return ret
	}
	return *o.IBAN
}

// GetIBANOk returns a tuple with the IBAN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundRecipient) GetIBANOk() (*string, bool) {
	if o == nil || common.IsNil(o.IBAN) {
		return nil, false
	}
	return o.IBAN, true
}

// HasIBAN returns a boolean if a field has been set.
func (o *FundRecipient) HasIBAN() bool {
	if o != nil && !common.IsNil(o.IBAN) {
		return true
	}

	return false
}

// SetIBAN gets a reference to the given string and assigns it to the IBAN field.
func (o *FundRecipient) SetIBAN(v string) {
	o.IBAN = &v
}

// GetBillingAddress returns the BillingAddress field value if set, zero value otherwise.
func (o *FundRecipient) GetBillingAddress() Address {
	if o == nil || common.IsNil(o.BillingAddress) {
		var ret Address
		return ret
	}
	return *o.BillingAddress
}

// GetBillingAddressOk returns a tuple with the BillingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundRecipient) GetBillingAddressOk() (*Address, bool) {
	if o == nil || common.IsNil(o.BillingAddress) {
		return nil, false
	}
	return o.BillingAddress, true
}

// HasBillingAddress returns a boolean if a field has been set.
func (o *FundRecipient) HasBillingAddress() bool {
	if o != nil && !common.IsNil(o.BillingAddress) {
		return true
	}

	return false
}

// SetBillingAddress gets a reference to the given Address and assigns it to the BillingAddress field.
func (o *FundRecipient) SetBillingAddress(v Address) {
	o.BillingAddress = &v
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *FundRecipient) GetPaymentMethod() CardDetails {
	if o == nil || common.IsNil(o.PaymentMethod) {
		var ret CardDetails
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundRecipient) GetPaymentMethodOk() (*CardDetails, bool) {
	if o == nil || common.IsNil(o.PaymentMethod) {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *FundRecipient) HasPaymentMethod() bool {
	if o != nil && !common.IsNil(o.PaymentMethod) {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given CardDetails and assigns it to the PaymentMethod field.
func (o *FundRecipient) SetPaymentMethod(v CardDetails) {
	o.PaymentMethod = &v
}

// GetShopperEmail returns the ShopperEmail field value if set, zero value otherwise.
func (o *FundRecipient) GetShopperEmail() string {
	if o == nil || common.IsNil(o.ShopperEmail) {
		var ret string
		return ret
	}
	return *o.ShopperEmail
}

// GetShopperEmailOk returns a tuple with the ShopperEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundRecipient) GetShopperEmailOk() (*string, bool) {
	if o == nil || common.IsNil(o.ShopperEmail) {
		return nil, false
	}
	return o.ShopperEmail, true
}

// HasShopperEmail returns a boolean if a field has been set.
func (o *FundRecipient) HasShopperEmail() bool {
	if o != nil && !common.IsNil(o.ShopperEmail) {
		return true
	}

	return false
}

// SetShopperEmail gets a reference to the given string and assigns it to the ShopperEmail field.
func (o *FundRecipient) SetShopperEmail(v string) {
	o.ShopperEmail = &v
}

// GetShopperName returns the ShopperName field value if set, zero value otherwise.
func (o *FundRecipient) GetShopperName() Name {
	if o == nil || common.IsNil(o.ShopperName) {
		var ret Name
		return ret
	}
	return *o.ShopperName
}

// GetShopperNameOk returns a tuple with the ShopperName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundRecipient) GetShopperNameOk() (*Name, bool) {
	if o == nil || common.IsNil(o.ShopperName) {
		return nil, false
	}
	return o.ShopperName, true
}

// HasShopperName returns a boolean if a field has been set.
func (o *FundRecipient) HasShopperName() bool {
	if o != nil && !common.IsNil(o.ShopperName) {
		return true
	}

	return false
}

// SetShopperName gets a reference to the given Name and assigns it to the ShopperName field.
func (o *FundRecipient) SetShopperName(v Name) {
	o.ShopperName = &v
}

// GetShopperReference returns the ShopperReference field value if set, zero value otherwise.
func (o *FundRecipient) GetShopperReference() string {
	if o == nil || common.IsNil(o.ShopperReference) {
		var ret string
		return ret
	}
	return *o.ShopperReference
}

// GetShopperReferenceOk returns a tuple with the ShopperReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundRecipient) GetShopperReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.ShopperReference) {
		return nil, false
	}
	return o.ShopperReference, true
}

// HasShopperReference returns a boolean if a field has been set.
func (o *FundRecipient) HasShopperReference() bool {
	if o != nil && !common.IsNil(o.ShopperReference) {
		return true
	}

	return false
}

// SetShopperReference gets a reference to the given string and assigns it to the ShopperReference field.
func (o *FundRecipient) SetShopperReference(v string) {
	o.ShopperReference = &v
}

// GetStoredPaymentMethodId returns the StoredPaymentMethodId field value if set, zero value otherwise.
func (o *FundRecipient) GetStoredPaymentMethodId() string {
	if o == nil || common.IsNil(o.StoredPaymentMethodId) {
		var ret string
		return ret
	}
	return *o.StoredPaymentMethodId
}

// GetStoredPaymentMethodIdOk returns a tuple with the StoredPaymentMethodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundRecipient) GetStoredPaymentMethodIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.StoredPaymentMethodId) {
		return nil, false
	}
	return o.StoredPaymentMethodId, true
}

// HasStoredPaymentMethodId returns a boolean if a field has been set.
func (o *FundRecipient) HasStoredPaymentMethodId() bool {
	if o != nil && !common.IsNil(o.StoredPaymentMethodId) {
		return true
	}

	return false
}

// SetStoredPaymentMethodId gets a reference to the given string and assigns it to the StoredPaymentMethodId field.
func (o *FundRecipient) SetStoredPaymentMethodId(v string) {
	o.StoredPaymentMethodId = &v
}

// GetSubMerchant returns the SubMerchant field value if set, zero value otherwise.
func (o *FundRecipient) GetSubMerchant() SubMerchant {
	if o == nil || common.IsNil(o.SubMerchant) {
		var ret SubMerchant
		return ret
	}
	return *o.SubMerchant
}

// GetSubMerchantOk returns a tuple with the SubMerchant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundRecipient) GetSubMerchantOk() (*SubMerchant, bool) {
	if o == nil || common.IsNil(o.SubMerchant) {
		return nil, false
	}
	return o.SubMerchant, true
}

// HasSubMerchant returns a boolean if a field has been set.
func (o *FundRecipient) HasSubMerchant() bool {
	if o != nil && !common.IsNil(o.SubMerchant) {
		return true
	}

	return false
}

// SetSubMerchant gets a reference to the given SubMerchant and assigns it to the SubMerchant field.
func (o *FundRecipient) SetSubMerchant(v SubMerchant) {
	o.SubMerchant = &v
}

// GetTelephoneNumber returns the TelephoneNumber field value if set, zero value otherwise.
func (o *FundRecipient) GetTelephoneNumber() string {
	if o == nil || common.IsNil(o.TelephoneNumber) {
		var ret string
		return ret
	}
	return *o.TelephoneNumber
}

// GetTelephoneNumberOk returns a tuple with the TelephoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundRecipient) GetTelephoneNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.TelephoneNumber) {
		return nil, false
	}
	return o.TelephoneNumber, true
}

// HasTelephoneNumber returns a boolean if a field has been set.
func (o *FundRecipient) HasTelephoneNumber() bool {
	if o != nil && !common.IsNil(o.TelephoneNumber) {
		return true
	}

	return false
}

// SetTelephoneNumber gets a reference to the given string and assigns it to the TelephoneNumber field.
func (o *FundRecipient) SetTelephoneNumber(v string) {
	o.TelephoneNumber = &v
}

// GetWalletIdentifier returns the WalletIdentifier field value if set, zero value otherwise.
func (o *FundRecipient) GetWalletIdentifier() string {
	if o == nil || common.IsNil(o.WalletIdentifier) {
		var ret string
		return ret
	}
	return *o.WalletIdentifier
}

// GetWalletIdentifierOk returns a tuple with the WalletIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundRecipient) GetWalletIdentifierOk() (*string, bool) {
	if o == nil || common.IsNil(o.WalletIdentifier) {
		return nil, false
	}
	return o.WalletIdentifier, true
}

// HasWalletIdentifier returns a boolean if a field has been set.
func (o *FundRecipient) HasWalletIdentifier() bool {
	if o != nil && !common.IsNil(o.WalletIdentifier) {
		return true
	}

	return false
}

// SetWalletIdentifier gets a reference to the given string and assigns it to the WalletIdentifier field.
func (o *FundRecipient) SetWalletIdentifier(v string) {
	o.WalletIdentifier = &v
}

// GetWalletOwnerTaxId returns the WalletOwnerTaxId field value if set, zero value otherwise.
func (o *FundRecipient) GetWalletOwnerTaxId() string {
	if o == nil || common.IsNil(o.WalletOwnerTaxId) {
		var ret string
		return ret
	}
	return *o.WalletOwnerTaxId
}

// GetWalletOwnerTaxIdOk returns a tuple with the WalletOwnerTaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundRecipient) GetWalletOwnerTaxIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.WalletOwnerTaxId) {
		return nil, false
	}
	return o.WalletOwnerTaxId, true
}

// HasWalletOwnerTaxId returns a boolean if a field has been set.
func (o *FundRecipient) HasWalletOwnerTaxId() bool {
	if o != nil && !common.IsNil(o.WalletOwnerTaxId) {
		return true
	}

	return false
}

// SetWalletOwnerTaxId gets a reference to the given string and assigns it to the WalletOwnerTaxId field.
func (o *FundRecipient) SetWalletOwnerTaxId(v string) {
	o.WalletOwnerTaxId = &v
}

// GetWalletPurpose returns the WalletPurpose field value if set, zero value otherwise.
func (o *FundRecipient) GetWalletPurpose() string {
	if o == nil || common.IsNil(o.WalletPurpose) {
		var ret string
		return ret
	}
	return *o.WalletPurpose
}

// GetWalletPurposeOk returns a tuple with the WalletPurpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundRecipient) GetWalletPurposeOk() (*string, bool) {
	if o == nil || common.IsNil(o.WalletPurpose) {
		return nil, false
	}
	return o.WalletPurpose, true
}

// HasWalletPurpose returns a boolean if a field has been set.
func (o *FundRecipient) HasWalletPurpose() bool {
	if o != nil && !common.IsNil(o.WalletPurpose) {
		return true
	}

	return false
}

// SetWalletPurpose gets a reference to the given string and assigns it to the WalletPurpose field.
func (o *FundRecipient) SetWalletPurpose(v string) {
	o.WalletPurpose = &v
}

func (o FundRecipient) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FundRecipient) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.IBAN) {
		toSerialize["IBAN"] = o.IBAN
	}
	if !common.IsNil(o.BillingAddress) {
		toSerialize["billingAddress"] = o.BillingAddress
	}
	if !common.IsNil(o.PaymentMethod) {
		toSerialize["paymentMethod"] = o.PaymentMethod
	}
	if !common.IsNil(o.ShopperEmail) {
		toSerialize["shopperEmail"] = o.ShopperEmail
	}
	if !common.IsNil(o.ShopperName) {
		toSerialize["shopperName"] = o.ShopperName
	}
	if !common.IsNil(o.ShopperReference) {
		toSerialize["shopperReference"] = o.ShopperReference
	}
	if !common.IsNil(o.StoredPaymentMethodId) {
		toSerialize["storedPaymentMethodId"] = o.StoredPaymentMethodId
	}
	if !common.IsNil(o.SubMerchant) {
		toSerialize["subMerchant"] = o.SubMerchant
	}
	if !common.IsNil(o.TelephoneNumber) {
		toSerialize["telephoneNumber"] = o.TelephoneNumber
	}
	if !common.IsNil(o.WalletIdentifier) {
		toSerialize["walletIdentifier"] = o.WalletIdentifier
	}
	if !common.IsNil(o.WalletOwnerTaxId) {
		toSerialize["walletOwnerTaxId"] = o.WalletOwnerTaxId
	}
	if !common.IsNil(o.WalletPurpose) {
		toSerialize["walletPurpose"] = o.WalletPurpose
	}
	return toSerialize, nil
}

type NullableFundRecipient struct {
	value *FundRecipient
	isSet bool
}

func (v NullableFundRecipient) Get() *FundRecipient {
	return v.value
}

func (v *NullableFundRecipient) Set(val *FundRecipient) {
	v.value = val
	v.isSet = true
}

func (v NullableFundRecipient) IsSet() bool {
	return v.isSet
}

func (v *NullableFundRecipient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFundRecipient(val *FundRecipient) *NullableFundRecipient {
	return &NullableFundRecipient{value: val, isSet: true}
}

func (v NullableFundRecipient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFundRecipient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *FundRecipient) isValidWalletPurpose() bool {
	var allowedEnumValues = []string{"identifiedBoleto", "transferDifferentWallet", "transferOwnWallet", "transferSameWallet", "unidentifiedBoleto"}
	for _, allowed := range allowedEnumValues {
		if o.GetWalletPurpose() == allowed {
			return true
		}
	}
	return false
}
