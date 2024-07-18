/*
Adyen Payout API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package payout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v11/src/common"
)

// checks if the StoreDetailAndSubmitRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &StoreDetailAndSubmitRequest{}

// StoreDetailAndSubmitRequest struct for StoreDetailAndSubmitRequest
type StoreDetailAndSubmitRequest struct {
	// This field contains additional data, which may be required for a particular request.
	AdditionalData *map[string]string `json:"additionalData,omitempty"`
	Amount         Amount             `json:"amount"`
	Bank           *BankAccount       `json:"bank,omitempty"`
	BillingAddress *Address           `json:"billingAddress,omitempty"`
	Card           *Card              `json:"card,omitempty"`
	// The date of birth. Format: [ISO-8601](https://www.w3.org/TR/NOTE-datetime); example: YYYY-MM-DD For Paysafecard it must be the same as used when registering the Paysafecard account. > This field is mandatory for natural persons.
	DateOfBirth string `json:"dateOfBirth"`
	// The type of the entity the payout is processed for.
	EntityType string `json:"entityType"`
	// An integer value that is added to the normal fraud score. The value can be either positive or negative.
	FraudOffset *int32 `json:"fraudOffset,omitempty"`
	// The merchant account identifier, with which you want to process the transaction.
	MerchantAccount string `json:"merchantAccount"`
	// The shopper's nationality.  A valid value is an ISO 2-character country code (e.g. 'NL').
	Nationality string    `json:"nationality"`
	Recurring   Recurring `json:"recurring"`
	// The merchant reference for this payment. This reference will be used in all communication to the merchant about the status of the payout. Although it is a good idea to make sure it is unique, this is not a requirement.
	Reference string `json:"reference"`
	// The name of the brand to make a payout to.  For Paysafecard it must be set to `paysafecard`.
	SelectedBrand *string `json:"selectedBrand,omitempty"`
	// The shopper's email address.
	ShopperEmail string `json:"shopperEmail"`
	ShopperName  *Name  `json:"shopperName,omitempty"`
	// The shopper's reference for the payment transaction.
	ShopperReference string `json:"shopperReference"`
	// The description of this payout. This description is shown on the bank statement of the shopper (if this is supported by the chosen payment method).
	ShopperStatement *string `json:"shopperStatement,omitempty"`
	// The shopper's social security number.
	SocialSecurityNumber *string `json:"socialSecurityNumber,omitempty"`
	// The shopper's phone number.
	TelephoneNumber *string `json:"telephoneNumber,omitempty"`
}

// NewStoreDetailAndSubmitRequest instantiates a new StoreDetailAndSubmitRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoreDetailAndSubmitRequest(amount Amount, dateOfBirth string, entityType string, merchantAccount string, nationality string, recurring Recurring, reference string, shopperEmail string, shopperReference string) *StoreDetailAndSubmitRequest {
	this := StoreDetailAndSubmitRequest{}
	this.Amount = amount
	this.DateOfBirth = dateOfBirth
	this.EntityType = entityType
	this.MerchantAccount = merchantAccount
	this.Nationality = nationality
	this.Recurring = recurring
	this.Reference = reference
	this.ShopperEmail = shopperEmail
	this.ShopperReference = shopperReference
	return &this
}

// NewStoreDetailAndSubmitRequestWithDefaults instantiates a new StoreDetailAndSubmitRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoreDetailAndSubmitRequestWithDefaults() *StoreDetailAndSubmitRequest {
	this := StoreDetailAndSubmitRequest{}
	return &this
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *StoreDetailAndSubmitRequest) GetAdditionalData() map[string]string {
	if o == nil || common.IsNil(o.AdditionalData) {
		var ret map[string]string
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreDetailAndSubmitRequest) GetAdditionalDataOk() (*map[string]string, bool) {
	if o == nil || common.IsNil(o.AdditionalData) {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *StoreDetailAndSubmitRequest) HasAdditionalData() bool {
	if o != nil && !common.IsNil(o.AdditionalData) {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given map[string]string and assigns it to the AdditionalData field.
func (o *StoreDetailAndSubmitRequest) SetAdditionalData(v map[string]string) {
	o.AdditionalData = &v
}

// GetAmount returns the Amount field value
func (o *StoreDetailAndSubmitRequest) GetAmount() Amount {
	if o == nil {
		var ret Amount
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *StoreDetailAndSubmitRequest) GetAmountOk() (*Amount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *StoreDetailAndSubmitRequest) SetAmount(v Amount) {
	o.Amount = v
}

// GetBank returns the Bank field value if set, zero value otherwise.
func (o *StoreDetailAndSubmitRequest) GetBank() BankAccount {
	if o == nil || common.IsNil(o.Bank) {
		var ret BankAccount
		return ret
	}
	return *o.Bank
}

// GetBankOk returns a tuple with the Bank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreDetailAndSubmitRequest) GetBankOk() (*BankAccount, bool) {
	if o == nil || common.IsNil(o.Bank) {
		return nil, false
	}
	return o.Bank, true
}

// HasBank returns a boolean if a field has been set.
func (o *StoreDetailAndSubmitRequest) HasBank() bool {
	if o != nil && !common.IsNil(o.Bank) {
		return true
	}

	return false
}

// SetBank gets a reference to the given BankAccount and assigns it to the Bank field.
func (o *StoreDetailAndSubmitRequest) SetBank(v BankAccount) {
	o.Bank = &v
}

// GetBillingAddress returns the BillingAddress field value if set, zero value otherwise.
func (o *StoreDetailAndSubmitRequest) GetBillingAddress() Address {
	if o == nil || common.IsNil(o.BillingAddress) {
		var ret Address
		return ret
	}
	return *o.BillingAddress
}

// GetBillingAddressOk returns a tuple with the BillingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreDetailAndSubmitRequest) GetBillingAddressOk() (*Address, bool) {
	if o == nil || common.IsNil(o.BillingAddress) {
		return nil, false
	}
	return o.BillingAddress, true
}

// HasBillingAddress returns a boolean if a field has been set.
func (o *StoreDetailAndSubmitRequest) HasBillingAddress() bool {
	if o != nil && !common.IsNil(o.BillingAddress) {
		return true
	}

	return false
}

// SetBillingAddress gets a reference to the given Address and assigns it to the BillingAddress field.
func (o *StoreDetailAndSubmitRequest) SetBillingAddress(v Address) {
	o.BillingAddress = &v
}

// GetCard returns the Card field value if set, zero value otherwise.
func (o *StoreDetailAndSubmitRequest) GetCard() Card {
	if o == nil || common.IsNil(o.Card) {
		var ret Card
		return ret
	}
	return *o.Card
}

// GetCardOk returns a tuple with the Card field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreDetailAndSubmitRequest) GetCardOk() (*Card, bool) {
	if o == nil || common.IsNil(o.Card) {
		return nil, false
	}
	return o.Card, true
}

// HasCard returns a boolean if a field has been set.
func (o *StoreDetailAndSubmitRequest) HasCard() bool {
	if o != nil && !common.IsNil(o.Card) {
		return true
	}

	return false
}

// SetCard gets a reference to the given Card and assigns it to the Card field.
func (o *StoreDetailAndSubmitRequest) SetCard(v Card) {
	o.Card = &v
}

// GetDateOfBirth returns the DateOfBirth field value
func (o *StoreDetailAndSubmitRequest) GetDateOfBirth() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DateOfBirth
}

// GetDateOfBirthOk returns a tuple with the DateOfBirth field value
// and a boolean to check if the value has been set.
func (o *StoreDetailAndSubmitRequest) GetDateOfBirthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DateOfBirth, true
}

// SetDateOfBirth sets field value
func (o *StoreDetailAndSubmitRequest) SetDateOfBirth(v string) {
	o.DateOfBirth = v
}

// GetEntityType returns the EntityType field value
func (o *StoreDetailAndSubmitRequest) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *StoreDetailAndSubmitRequest) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *StoreDetailAndSubmitRequest) SetEntityType(v string) {
	o.EntityType = v
}

// GetFraudOffset returns the FraudOffset field value if set, zero value otherwise.
func (o *StoreDetailAndSubmitRequest) GetFraudOffset() int32 {
	if o == nil || common.IsNil(o.FraudOffset) {
		var ret int32
		return ret
	}
	return *o.FraudOffset
}

// GetFraudOffsetOk returns a tuple with the FraudOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreDetailAndSubmitRequest) GetFraudOffsetOk() (*int32, bool) {
	if o == nil || common.IsNil(o.FraudOffset) {
		return nil, false
	}
	return o.FraudOffset, true
}

// HasFraudOffset returns a boolean if a field has been set.
func (o *StoreDetailAndSubmitRequest) HasFraudOffset() bool {
	if o != nil && !common.IsNil(o.FraudOffset) {
		return true
	}

	return false
}

// SetFraudOffset gets a reference to the given int32 and assigns it to the FraudOffset field.
func (o *StoreDetailAndSubmitRequest) SetFraudOffset(v int32) {
	o.FraudOffset = &v
}

// GetMerchantAccount returns the MerchantAccount field value
func (o *StoreDetailAndSubmitRequest) GetMerchantAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value
// and a boolean to check if the value has been set.
func (o *StoreDetailAndSubmitRequest) GetMerchantAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantAccount, true
}

// SetMerchantAccount sets field value
func (o *StoreDetailAndSubmitRequest) SetMerchantAccount(v string) {
	o.MerchantAccount = v
}

// GetNationality returns the Nationality field value
func (o *StoreDetailAndSubmitRequest) GetNationality() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Nationality
}

// GetNationalityOk returns a tuple with the Nationality field value
// and a boolean to check if the value has been set.
func (o *StoreDetailAndSubmitRequest) GetNationalityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nationality, true
}

// SetNationality sets field value
func (o *StoreDetailAndSubmitRequest) SetNationality(v string) {
	o.Nationality = v
}

// GetRecurring returns the Recurring field value
func (o *StoreDetailAndSubmitRequest) GetRecurring() Recurring {
	if o == nil {
		var ret Recurring
		return ret
	}

	return o.Recurring
}

// GetRecurringOk returns a tuple with the Recurring field value
// and a boolean to check if the value has been set.
func (o *StoreDetailAndSubmitRequest) GetRecurringOk() (*Recurring, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recurring, true
}

// SetRecurring sets field value
func (o *StoreDetailAndSubmitRequest) SetRecurring(v Recurring) {
	o.Recurring = v
}

// GetReference returns the Reference field value
func (o *StoreDetailAndSubmitRequest) GetReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
func (o *StoreDetailAndSubmitRequest) GetReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reference, true
}

// SetReference sets field value
func (o *StoreDetailAndSubmitRequest) SetReference(v string) {
	o.Reference = v
}

// GetSelectedBrand returns the SelectedBrand field value if set, zero value otherwise.
func (o *StoreDetailAndSubmitRequest) GetSelectedBrand() string {
	if o == nil || common.IsNil(o.SelectedBrand) {
		var ret string
		return ret
	}
	return *o.SelectedBrand
}

// GetSelectedBrandOk returns a tuple with the SelectedBrand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreDetailAndSubmitRequest) GetSelectedBrandOk() (*string, bool) {
	if o == nil || common.IsNil(o.SelectedBrand) {
		return nil, false
	}
	return o.SelectedBrand, true
}

// HasSelectedBrand returns a boolean if a field has been set.
func (o *StoreDetailAndSubmitRequest) HasSelectedBrand() bool {
	if o != nil && !common.IsNil(o.SelectedBrand) {
		return true
	}

	return false
}

// SetSelectedBrand gets a reference to the given string and assigns it to the SelectedBrand field.
func (o *StoreDetailAndSubmitRequest) SetSelectedBrand(v string) {
	o.SelectedBrand = &v
}

// GetShopperEmail returns the ShopperEmail field value
func (o *StoreDetailAndSubmitRequest) GetShopperEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShopperEmail
}

// GetShopperEmailOk returns a tuple with the ShopperEmail field value
// and a boolean to check if the value has been set.
func (o *StoreDetailAndSubmitRequest) GetShopperEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShopperEmail, true
}

// SetShopperEmail sets field value
func (o *StoreDetailAndSubmitRequest) SetShopperEmail(v string) {
	o.ShopperEmail = v
}

// GetShopperName returns the ShopperName field value if set, zero value otherwise.
func (o *StoreDetailAndSubmitRequest) GetShopperName() Name {
	if o == nil || common.IsNil(o.ShopperName) {
		var ret Name
		return ret
	}
	return *o.ShopperName
}

// GetShopperNameOk returns a tuple with the ShopperName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreDetailAndSubmitRequest) GetShopperNameOk() (*Name, bool) {
	if o == nil || common.IsNil(o.ShopperName) {
		return nil, false
	}
	return o.ShopperName, true
}

// HasShopperName returns a boolean if a field has been set.
func (o *StoreDetailAndSubmitRequest) HasShopperName() bool {
	if o != nil && !common.IsNil(o.ShopperName) {
		return true
	}

	return false
}

// SetShopperName gets a reference to the given Name and assigns it to the ShopperName field.
func (o *StoreDetailAndSubmitRequest) SetShopperName(v Name) {
	o.ShopperName = &v
}

// GetShopperReference returns the ShopperReference field value
func (o *StoreDetailAndSubmitRequest) GetShopperReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShopperReference
}

// GetShopperReferenceOk returns a tuple with the ShopperReference field value
// and a boolean to check if the value has been set.
func (o *StoreDetailAndSubmitRequest) GetShopperReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShopperReference, true
}

// SetShopperReference sets field value
func (o *StoreDetailAndSubmitRequest) SetShopperReference(v string) {
	o.ShopperReference = v
}

// GetShopperStatement returns the ShopperStatement field value if set, zero value otherwise.
func (o *StoreDetailAndSubmitRequest) GetShopperStatement() string {
	if o == nil || common.IsNil(o.ShopperStatement) {
		var ret string
		return ret
	}
	return *o.ShopperStatement
}

// GetShopperStatementOk returns a tuple with the ShopperStatement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreDetailAndSubmitRequest) GetShopperStatementOk() (*string, bool) {
	if o == nil || common.IsNil(o.ShopperStatement) {
		return nil, false
	}
	return o.ShopperStatement, true
}

// HasShopperStatement returns a boolean if a field has been set.
func (o *StoreDetailAndSubmitRequest) HasShopperStatement() bool {
	if o != nil && !common.IsNil(o.ShopperStatement) {
		return true
	}

	return false
}

// SetShopperStatement gets a reference to the given string and assigns it to the ShopperStatement field.
func (o *StoreDetailAndSubmitRequest) SetShopperStatement(v string) {
	o.ShopperStatement = &v
}

// GetSocialSecurityNumber returns the SocialSecurityNumber field value if set, zero value otherwise.
func (o *StoreDetailAndSubmitRequest) GetSocialSecurityNumber() string {
	if o == nil || common.IsNil(o.SocialSecurityNumber) {
		var ret string
		return ret
	}
	return *o.SocialSecurityNumber
}

// GetSocialSecurityNumberOk returns a tuple with the SocialSecurityNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreDetailAndSubmitRequest) GetSocialSecurityNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.SocialSecurityNumber) {
		return nil, false
	}
	return o.SocialSecurityNumber, true
}

// HasSocialSecurityNumber returns a boolean if a field has been set.
func (o *StoreDetailAndSubmitRequest) HasSocialSecurityNumber() bool {
	if o != nil && !common.IsNil(o.SocialSecurityNumber) {
		return true
	}

	return false
}

// SetSocialSecurityNumber gets a reference to the given string and assigns it to the SocialSecurityNumber field.
func (o *StoreDetailAndSubmitRequest) SetSocialSecurityNumber(v string) {
	o.SocialSecurityNumber = &v
}

// GetTelephoneNumber returns the TelephoneNumber field value if set, zero value otherwise.
func (o *StoreDetailAndSubmitRequest) GetTelephoneNumber() string {
	if o == nil || common.IsNil(o.TelephoneNumber) {
		var ret string
		return ret
	}
	return *o.TelephoneNumber
}

// GetTelephoneNumberOk returns a tuple with the TelephoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreDetailAndSubmitRequest) GetTelephoneNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.TelephoneNumber) {
		return nil, false
	}
	return o.TelephoneNumber, true
}

// HasTelephoneNumber returns a boolean if a field has been set.
func (o *StoreDetailAndSubmitRequest) HasTelephoneNumber() bool {
	if o != nil && !common.IsNil(o.TelephoneNumber) {
		return true
	}

	return false
}

// SetTelephoneNumber gets a reference to the given string and assigns it to the TelephoneNumber field.
func (o *StoreDetailAndSubmitRequest) SetTelephoneNumber(v string) {
	o.TelephoneNumber = &v
}

func (o StoreDetailAndSubmitRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StoreDetailAndSubmitRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AdditionalData) {
		toSerialize["additionalData"] = o.AdditionalData
	}
	toSerialize["amount"] = o.Amount
	if !common.IsNil(o.Bank) {
		toSerialize["bank"] = o.Bank
	}
	if !common.IsNil(o.BillingAddress) {
		toSerialize["billingAddress"] = o.BillingAddress
	}
	if !common.IsNil(o.Card) {
		toSerialize["card"] = o.Card
	}
	toSerialize["dateOfBirth"] = o.DateOfBirth
	toSerialize["entityType"] = o.EntityType
	if !common.IsNil(o.FraudOffset) {
		toSerialize["fraudOffset"] = o.FraudOffset
	}
	toSerialize["merchantAccount"] = o.MerchantAccount
	toSerialize["nationality"] = o.Nationality
	toSerialize["recurring"] = o.Recurring
	toSerialize["reference"] = o.Reference
	if !common.IsNil(o.SelectedBrand) {
		toSerialize["selectedBrand"] = o.SelectedBrand
	}
	toSerialize["shopperEmail"] = o.ShopperEmail
	if !common.IsNil(o.ShopperName) {
		toSerialize["shopperName"] = o.ShopperName
	}
	toSerialize["shopperReference"] = o.ShopperReference
	if !common.IsNil(o.ShopperStatement) {
		toSerialize["shopperStatement"] = o.ShopperStatement
	}
	if !common.IsNil(o.SocialSecurityNumber) {
		toSerialize["socialSecurityNumber"] = o.SocialSecurityNumber
	}
	if !common.IsNil(o.TelephoneNumber) {
		toSerialize["telephoneNumber"] = o.TelephoneNumber
	}
	return toSerialize, nil
}

type NullableStoreDetailAndSubmitRequest struct {
	value *StoreDetailAndSubmitRequest
	isSet bool
}

func (v NullableStoreDetailAndSubmitRequest) Get() *StoreDetailAndSubmitRequest {
	return v.value
}

func (v *NullableStoreDetailAndSubmitRequest) Set(val *StoreDetailAndSubmitRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableStoreDetailAndSubmitRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableStoreDetailAndSubmitRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoreDetailAndSubmitRequest(val *StoreDetailAndSubmitRequest) *NullableStoreDetailAndSubmitRequest {
	return &NullableStoreDetailAndSubmitRequest{value: val, isSet: true}
}

func (v NullableStoreDetailAndSubmitRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoreDetailAndSubmitRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *StoreDetailAndSubmitRequest) isValidEntityType() bool {
	var allowedEnumValues = []string{"NaturalPerson", "Company"}
	for _, allowed := range allowedEnumValues {
		if o.GetEntityType() == allowed {
			return true
		}
	}
	return false
}
