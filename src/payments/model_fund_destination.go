/*
Adyen Payment API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package payments

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v19/src/common"
)

// checks if the FundDestination type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &FundDestination{}

// FundDestination struct for FundDestination
type FundDestination struct {
	// Bank Account Number of the recipient
	IBAN *string `json:"IBAN,omitempty"`
	// a map of name/value pairs for passing in additional/industry-specific data
	AdditionalData *map[string]string `json:"additionalData,omitempty"`
	BillingAddress *Address           `json:"billingAddress,omitempty"`
	Card           *Card              `json:"card,omitempty"`
	// The `recurringDetailReference` you want to use for this payment. The value `LATEST` can be used to select the most recently stored recurring detail.
	SelectedRecurringDetailReference *string `json:"selectedRecurringDetailReference,omitempty"`
	// the email address of the person
	ShopperEmail *string `json:"shopperEmail,omitempty"`
	ShopperName  *Name   `json:"shopperName,omitempty"`
	// Required for recurring payments.  Your reference to uniquely identify this shopper, for example user ID or account ID. Minimum length: 3 characters. > Your reference must not include personally identifiable information (PII), for example name or email address.
	ShopperReference *string      `json:"shopperReference,omitempty"`
	SubMerchant      *SubMerchant `json:"subMerchant,omitempty"`
	// the telephone number of the person
	TelephoneNumber *string `json:"telephoneNumber,omitempty"`
	// The purpose of a digital wallet transaction.
	WalletPurpose *string `json:"walletPurpose,omitempty"`
}

// NewFundDestination instantiates a new FundDestination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFundDestination() *FundDestination {
	this := FundDestination{}
	return &this
}

// NewFundDestinationWithDefaults instantiates a new FundDestination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFundDestinationWithDefaults() *FundDestination {
	this := FundDestination{}
	return &this
}

// GetIBAN returns the IBAN field value if set, zero value otherwise.
func (o *FundDestination) GetIBAN() string {
	if o == nil || common.IsNil(o.IBAN) {
		var ret string
		return ret
	}
	return *o.IBAN
}

// GetIBANOk returns a tuple with the IBAN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundDestination) GetIBANOk() (*string, bool) {
	if o == nil || common.IsNil(o.IBAN) {
		return nil, false
	}
	return o.IBAN, true
}

// HasIBAN returns a boolean if a field has been set.
func (o *FundDestination) HasIBAN() bool {
	if o != nil && !common.IsNil(o.IBAN) {
		return true
	}

	return false
}

// SetIBAN gets a reference to the given string and assigns it to the IBAN field.
func (o *FundDestination) SetIBAN(v string) {
	o.IBAN = &v
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *FundDestination) GetAdditionalData() map[string]string {
	if o == nil || common.IsNil(o.AdditionalData) {
		var ret map[string]string
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundDestination) GetAdditionalDataOk() (*map[string]string, bool) {
	if o == nil || common.IsNil(o.AdditionalData) {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *FundDestination) HasAdditionalData() bool {
	if o != nil && !common.IsNil(o.AdditionalData) {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given map[string]string and assigns it to the AdditionalData field.
func (o *FundDestination) SetAdditionalData(v map[string]string) {
	o.AdditionalData = &v
}

// GetBillingAddress returns the BillingAddress field value if set, zero value otherwise.
func (o *FundDestination) GetBillingAddress() Address {
	if o == nil || common.IsNil(o.BillingAddress) {
		var ret Address
		return ret
	}
	return *o.BillingAddress
}

// GetBillingAddressOk returns a tuple with the BillingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundDestination) GetBillingAddressOk() (*Address, bool) {
	if o == nil || common.IsNil(o.BillingAddress) {
		return nil, false
	}
	return o.BillingAddress, true
}

// HasBillingAddress returns a boolean if a field has been set.
func (o *FundDestination) HasBillingAddress() bool {
	if o != nil && !common.IsNil(o.BillingAddress) {
		return true
	}

	return false
}

// SetBillingAddress gets a reference to the given Address and assigns it to the BillingAddress field.
func (o *FundDestination) SetBillingAddress(v Address) {
	o.BillingAddress = &v
}

// GetCard returns the Card field value if set, zero value otherwise.
func (o *FundDestination) GetCard() Card {
	if o == nil || common.IsNil(o.Card) {
		var ret Card
		return ret
	}
	return *o.Card
}

// GetCardOk returns a tuple with the Card field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundDestination) GetCardOk() (*Card, bool) {
	if o == nil || common.IsNil(o.Card) {
		return nil, false
	}
	return o.Card, true
}

// HasCard returns a boolean if a field has been set.
func (o *FundDestination) HasCard() bool {
	if o != nil && !common.IsNil(o.Card) {
		return true
	}

	return false
}

// SetCard gets a reference to the given Card and assigns it to the Card field.
func (o *FundDestination) SetCard(v Card) {
	o.Card = &v
}

// GetSelectedRecurringDetailReference returns the SelectedRecurringDetailReference field value if set, zero value otherwise.
func (o *FundDestination) GetSelectedRecurringDetailReference() string {
	if o == nil || common.IsNil(o.SelectedRecurringDetailReference) {
		var ret string
		return ret
	}
	return *o.SelectedRecurringDetailReference
}

// GetSelectedRecurringDetailReferenceOk returns a tuple with the SelectedRecurringDetailReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundDestination) GetSelectedRecurringDetailReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.SelectedRecurringDetailReference) {
		return nil, false
	}
	return o.SelectedRecurringDetailReference, true
}

// HasSelectedRecurringDetailReference returns a boolean if a field has been set.
func (o *FundDestination) HasSelectedRecurringDetailReference() bool {
	if o != nil && !common.IsNil(o.SelectedRecurringDetailReference) {
		return true
	}

	return false
}

// SetSelectedRecurringDetailReference gets a reference to the given string and assigns it to the SelectedRecurringDetailReference field.
func (o *FundDestination) SetSelectedRecurringDetailReference(v string) {
	o.SelectedRecurringDetailReference = &v
}

// GetShopperEmail returns the ShopperEmail field value if set, zero value otherwise.
func (o *FundDestination) GetShopperEmail() string {
	if o == nil || common.IsNil(o.ShopperEmail) {
		var ret string
		return ret
	}
	return *o.ShopperEmail
}

// GetShopperEmailOk returns a tuple with the ShopperEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundDestination) GetShopperEmailOk() (*string, bool) {
	if o == nil || common.IsNil(o.ShopperEmail) {
		return nil, false
	}
	return o.ShopperEmail, true
}

// HasShopperEmail returns a boolean if a field has been set.
func (o *FundDestination) HasShopperEmail() bool {
	if o != nil && !common.IsNil(o.ShopperEmail) {
		return true
	}

	return false
}

// SetShopperEmail gets a reference to the given string and assigns it to the ShopperEmail field.
func (o *FundDestination) SetShopperEmail(v string) {
	o.ShopperEmail = &v
}

// GetShopperName returns the ShopperName field value if set, zero value otherwise.
func (o *FundDestination) GetShopperName() Name {
	if o == nil || common.IsNil(o.ShopperName) {
		var ret Name
		return ret
	}
	return *o.ShopperName
}

// GetShopperNameOk returns a tuple with the ShopperName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundDestination) GetShopperNameOk() (*Name, bool) {
	if o == nil || common.IsNil(o.ShopperName) {
		return nil, false
	}
	return o.ShopperName, true
}

// HasShopperName returns a boolean if a field has been set.
func (o *FundDestination) HasShopperName() bool {
	if o != nil && !common.IsNil(o.ShopperName) {
		return true
	}

	return false
}

// SetShopperName gets a reference to the given Name and assigns it to the ShopperName field.
func (o *FundDestination) SetShopperName(v Name) {
	o.ShopperName = &v
}

// GetShopperReference returns the ShopperReference field value if set, zero value otherwise.
func (o *FundDestination) GetShopperReference() string {
	if o == nil || common.IsNil(o.ShopperReference) {
		var ret string
		return ret
	}
	return *o.ShopperReference
}

// GetShopperReferenceOk returns a tuple with the ShopperReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundDestination) GetShopperReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.ShopperReference) {
		return nil, false
	}
	return o.ShopperReference, true
}

// HasShopperReference returns a boolean if a field has been set.
func (o *FundDestination) HasShopperReference() bool {
	if o != nil && !common.IsNil(o.ShopperReference) {
		return true
	}

	return false
}

// SetShopperReference gets a reference to the given string and assigns it to the ShopperReference field.
func (o *FundDestination) SetShopperReference(v string) {
	o.ShopperReference = &v
}

// GetSubMerchant returns the SubMerchant field value if set, zero value otherwise.
func (o *FundDestination) GetSubMerchant() SubMerchant {
	if o == nil || common.IsNil(o.SubMerchant) {
		var ret SubMerchant
		return ret
	}
	return *o.SubMerchant
}

// GetSubMerchantOk returns a tuple with the SubMerchant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundDestination) GetSubMerchantOk() (*SubMerchant, bool) {
	if o == nil || common.IsNil(o.SubMerchant) {
		return nil, false
	}
	return o.SubMerchant, true
}

// HasSubMerchant returns a boolean if a field has been set.
func (o *FundDestination) HasSubMerchant() bool {
	if o != nil && !common.IsNil(o.SubMerchant) {
		return true
	}

	return false
}

// SetSubMerchant gets a reference to the given SubMerchant and assigns it to the SubMerchant field.
func (o *FundDestination) SetSubMerchant(v SubMerchant) {
	o.SubMerchant = &v
}

// GetTelephoneNumber returns the TelephoneNumber field value if set, zero value otherwise.
func (o *FundDestination) GetTelephoneNumber() string {
	if o == nil || common.IsNil(o.TelephoneNumber) {
		var ret string
		return ret
	}
	return *o.TelephoneNumber
}

// GetTelephoneNumberOk returns a tuple with the TelephoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundDestination) GetTelephoneNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.TelephoneNumber) {
		return nil, false
	}
	return o.TelephoneNumber, true
}

// HasTelephoneNumber returns a boolean if a field has been set.
func (o *FundDestination) HasTelephoneNumber() bool {
	if o != nil && !common.IsNil(o.TelephoneNumber) {
		return true
	}

	return false
}

// SetTelephoneNumber gets a reference to the given string and assigns it to the TelephoneNumber field.
func (o *FundDestination) SetTelephoneNumber(v string) {
	o.TelephoneNumber = &v
}

// GetWalletPurpose returns the WalletPurpose field value if set, zero value otherwise.
func (o *FundDestination) GetWalletPurpose() string {
	if o == nil || common.IsNil(o.WalletPurpose) {
		var ret string
		return ret
	}
	return *o.WalletPurpose
}

// GetWalletPurposeOk returns a tuple with the WalletPurpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundDestination) GetWalletPurposeOk() (*string, bool) {
	if o == nil || common.IsNil(o.WalletPurpose) {
		return nil, false
	}
	return o.WalletPurpose, true
}

// HasWalletPurpose returns a boolean if a field has been set.
func (o *FundDestination) HasWalletPurpose() bool {
	if o != nil && !common.IsNil(o.WalletPurpose) {
		return true
	}

	return false
}

// SetWalletPurpose gets a reference to the given string and assigns it to the WalletPurpose field.
func (o *FundDestination) SetWalletPurpose(v string) {
	o.WalletPurpose = &v
}

func (o FundDestination) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FundDestination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.IBAN) {
		toSerialize["IBAN"] = o.IBAN
	}
	if !common.IsNil(o.AdditionalData) {
		toSerialize["additionalData"] = o.AdditionalData
	}
	if !common.IsNil(o.BillingAddress) {
		toSerialize["billingAddress"] = o.BillingAddress
	}
	if !common.IsNil(o.Card) {
		toSerialize["card"] = o.Card
	}
	if !common.IsNil(o.SelectedRecurringDetailReference) {
		toSerialize["selectedRecurringDetailReference"] = o.SelectedRecurringDetailReference
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
	if !common.IsNil(o.SubMerchant) {
		toSerialize["subMerchant"] = o.SubMerchant
	}
	if !common.IsNil(o.TelephoneNumber) {
		toSerialize["telephoneNumber"] = o.TelephoneNumber
	}
	if !common.IsNil(o.WalletPurpose) {
		toSerialize["walletPurpose"] = o.WalletPurpose
	}
	return toSerialize, nil
}

type NullableFundDestination struct {
	value *FundDestination
	isSet bool
}

func (v NullableFundDestination) Get() *FundDestination {
	return v.value
}

func (v *NullableFundDestination) Set(val *FundDestination) {
	v.value = val
	v.isSet = true
}

func (v NullableFundDestination) IsSet() bool {
	return v.isSet
}

func (v *NullableFundDestination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFundDestination(val *FundDestination) *NullableFundDestination {
	return &NullableFundDestination{value: val, isSet: true}
}

func (v NullableFundDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFundDestination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
