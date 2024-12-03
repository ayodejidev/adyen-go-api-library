/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v15/src/common"
)

// checks if the AchDetails type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AchDetails{}

// AchDetails struct for AchDetails
type AchDetails struct {
	// The bank account number (without separators).
	BankAccountNumber *string `json:"bankAccountNumber,omitempty"`
	// The bank account type (checking, savings...).
	BankAccountType *string `json:"bankAccountType,omitempty"`
	// The bank routing number of the account. The field value is `nil` in most cases.
	BankLocationId *string `json:"bankLocationId,omitempty"`
	// The checkout attempt identifier.
	CheckoutAttemptId *string `json:"checkoutAttemptId,omitempty"`
	// Encrypted bank account number. The bank account number (without separators).
	EncryptedBankAccountNumber *string `json:"encryptedBankAccountNumber,omitempty"`
	// Encrypted location id. The bank routing number of the account. The field value is `nil` in most cases.
	EncryptedBankLocationId *string `json:"encryptedBankLocationId,omitempty"`
	// The name of the bank account holder. If you submit a name with non-Latin characters, we automatically replace some of them with corresponding Latin characters to meet the FATF recommendations. For example: * χ12 is converted to ch12. * üA is converted to euA. * Peter Møller is converted to Peter Mller, because banks don't accept 'ø'. After replacement, the ownerName must have at least three alphanumeric characters (A-Z, a-z, 0-9), and at least one of them must be a valid Latin character (A-Z, a-z). For example: * John17 - allowed. * J17 - allowed. * 171 - not allowed. * John-7 - allowed. > If provided details don't match the required format, the response returns the error message: 203 'Invalid bank account holder name'.
	OwnerName *string `json:"ownerName,omitempty"`
	// This is the `recurringDetailReference` returned in the response when you created the token.
	// Deprecated since Adyen Checkout API v49
	// Use `storedPaymentMethodId` instead.
	RecurringDetailReference *string `json:"recurringDetailReference,omitempty"`
	// This is the `recurringDetailReference` returned in the response when you created the token.
	StoredPaymentMethodId *string `json:"storedPaymentMethodId,omitempty"`
	// The unique identifier of your user's verified transfer instrument, which you can use to top up their balance accounts.
	TransferInstrumentId *string `json:"transferInstrumentId,omitempty"`
	// **ach**
	Type *string `json:"type,omitempty"`
}

// NewAchDetails instantiates a new AchDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAchDetails() *AchDetails {
	this := AchDetails{}
	var type_ string = "ach"
	this.Type = &type_
	return &this
}

// NewAchDetailsWithDefaults instantiates a new AchDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAchDetailsWithDefaults() *AchDetails {
	this := AchDetails{}
	var type_ string = "ach"
	this.Type = &type_
	return &this
}

// GetBankAccountNumber returns the BankAccountNumber field value if set, zero value otherwise.
func (o *AchDetails) GetBankAccountNumber() string {
	if o == nil || common.IsNil(o.BankAccountNumber) {
		var ret string
		return ret
	}
	return *o.BankAccountNumber
}

// GetBankAccountNumberOk returns a tuple with the BankAccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AchDetails) GetBankAccountNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.BankAccountNumber) {
		return nil, false
	}
	return o.BankAccountNumber, true
}

// HasBankAccountNumber returns a boolean if a field has been set.
func (o *AchDetails) HasBankAccountNumber() bool {
	if o != nil && !common.IsNil(o.BankAccountNumber) {
		return true
	}

	return false
}

// SetBankAccountNumber gets a reference to the given string and assigns it to the BankAccountNumber field.
func (o *AchDetails) SetBankAccountNumber(v string) {
	o.BankAccountNumber = &v
}

// GetBankAccountType returns the BankAccountType field value if set, zero value otherwise.
func (o *AchDetails) GetBankAccountType() string {
	if o == nil || common.IsNil(o.BankAccountType) {
		var ret string
		return ret
	}
	return *o.BankAccountType
}

// GetBankAccountTypeOk returns a tuple with the BankAccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AchDetails) GetBankAccountTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.BankAccountType) {
		return nil, false
	}
	return o.BankAccountType, true
}

// HasBankAccountType returns a boolean if a field has been set.
func (o *AchDetails) HasBankAccountType() bool {
	if o != nil && !common.IsNil(o.BankAccountType) {
		return true
	}

	return false
}

// SetBankAccountType gets a reference to the given string and assigns it to the BankAccountType field.
func (o *AchDetails) SetBankAccountType(v string) {
	o.BankAccountType = &v
}

// GetBankLocationId returns the BankLocationId field value if set, zero value otherwise.
func (o *AchDetails) GetBankLocationId() string {
	if o == nil || common.IsNil(o.BankLocationId) {
		var ret string
		return ret
	}
	return *o.BankLocationId
}

// GetBankLocationIdOk returns a tuple with the BankLocationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AchDetails) GetBankLocationIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.BankLocationId) {
		return nil, false
	}
	return o.BankLocationId, true
}

// HasBankLocationId returns a boolean if a field has been set.
func (o *AchDetails) HasBankLocationId() bool {
	if o != nil && !common.IsNil(o.BankLocationId) {
		return true
	}

	return false
}

// SetBankLocationId gets a reference to the given string and assigns it to the BankLocationId field.
func (o *AchDetails) SetBankLocationId(v string) {
	o.BankLocationId = &v
}

// GetCheckoutAttemptId returns the CheckoutAttemptId field value if set, zero value otherwise.
func (o *AchDetails) GetCheckoutAttemptId() string {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		var ret string
		return ret
	}
	return *o.CheckoutAttemptId
}

// GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AchDetails) GetCheckoutAttemptIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		return nil, false
	}
	return o.CheckoutAttemptId, true
}

// HasCheckoutAttemptId returns a boolean if a field has been set.
func (o *AchDetails) HasCheckoutAttemptId() bool {
	if o != nil && !common.IsNil(o.CheckoutAttemptId) {
		return true
	}

	return false
}

// SetCheckoutAttemptId gets a reference to the given string and assigns it to the CheckoutAttemptId field.
func (o *AchDetails) SetCheckoutAttemptId(v string) {
	o.CheckoutAttemptId = &v
}

// GetEncryptedBankAccountNumber returns the EncryptedBankAccountNumber field value if set, zero value otherwise.
func (o *AchDetails) GetEncryptedBankAccountNumber() string {
	if o == nil || common.IsNil(o.EncryptedBankAccountNumber) {
		var ret string
		return ret
	}
	return *o.EncryptedBankAccountNumber
}

// GetEncryptedBankAccountNumberOk returns a tuple with the EncryptedBankAccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AchDetails) GetEncryptedBankAccountNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.EncryptedBankAccountNumber) {
		return nil, false
	}
	return o.EncryptedBankAccountNumber, true
}

// HasEncryptedBankAccountNumber returns a boolean if a field has been set.
func (o *AchDetails) HasEncryptedBankAccountNumber() bool {
	if o != nil && !common.IsNil(o.EncryptedBankAccountNumber) {
		return true
	}

	return false
}

// SetEncryptedBankAccountNumber gets a reference to the given string and assigns it to the EncryptedBankAccountNumber field.
func (o *AchDetails) SetEncryptedBankAccountNumber(v string) {
	o.EncryptedBankAccountNumber = &v
}

// GetEncryptedBankLocationId returns the EncryptedBankLocationId field value if set, zero value otherwise.
func (o *AchDetails) GetEncryptedBankLocationId() string {
	if o == nil || common.IsNil(o.EncryptedBankLocationId) {
		var ret string
		return ret
	}
	return *o.EncryptedBankLocationId
}

// GetEncryptedBankLocationIdOk returns a tuple with the EncryptedBankLocationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AchDetails) GetEncryptedBankLocationIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.EncryptedBankLocationId) {
		return nil, false
	}
	return o.EncryptedBankLocationId, true
}

// HasEncryptedBankLocationId returns a boolean if a field has been set.
func (o *AchDetails) HasEncryptedBankLocationId() bool {
	if o != nil && !common.IsNil(o.EncryptedBankLocationId) {
		return true
	}

	return false
}

// SetEncryptedBankLocationId gets a reference to the given string and assigns it to the EncryptedBankLocationId field.
func (o *AchDetails) SetEncryptedBankLocationId(v string) {
	o.EncryptedBankLocationId = &v
}

// GetOwnerName returns the OwnerName field value if set, zero value otherwise.
func (o *AchDetails) GetOwnerName() string {
	if o == nil || common.IsNil(o.OwnerName) {
		var ret string
		return ret
	}
	return *o.OwnerName
}

// GetOwnerNameOk returns a tuple with the OwnerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AchDetails) GetOwnerNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.OwnerName) {
		return nil, false
	}
	return o.OwnerName, true
}

// HasOwnerName returns a boolean if a field has been set.
func (o *AchDetails) HasOwnerName() bool {
	if o != nil && !common.IsNil(o.OwnerName) {
		return true
	}

	return false
}

// SetOwnerName gets a reference to the given string and assigns it to the OwnerName field.
func (o *AchDetails) SetOwnerName(v string) {
	o.OwnerName = &v
}

// GetRecurringDetailReference returns the RecurringDetailReference field value if set, zero value otherwise.
// Deprecated since Adyen Checkout API v49
// Use `storedPaymentMethodId` instead.
func (o *AchDetails) GetRecurringDetailReference() string {
	if o == nil || common.IsNil(o.RecurringDetailReference) {
		var ret string
		return ret
	}
	return *o.RecurringDetailReference
}

// GetRecurringDetailReferenceOk returns a tuple with the RecurringDetailReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated since Adyen Checkout API v49
// Use `storedPaymentMethodId` instead.
func (o *AchDetails) GetRecurringDetailReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.RecurringDetailReference) {
		return nil, false
	}
	return o.RecurringDetailReference, true
}

// HasRecurringDetailReference returns a boolean if a field has been set.
func (o *AchDetails) HasRecurringDetailReference() bool {
	if o != nil && !common.IsNil(o.RecurringDetailReference) {
		return true
	}

	return false
}

// SetRecurringDetailReference gets a reference to the given string and assigns it to the RecurringDetailReference field.
// Deprecated since Adyen Checkout API v49
// Use `storedPaymentMethodId` instead.
func (o *AchDetails) SetRecurringDetailReference(v string) {
	o.RecurringDetailReference = &v
}

// GetStoredPaymentMethodId returns the StoredPaymentMethodId field value if set, zero value otherwise.
func (o *AchDetails) GetStoredPaymentMethodId() string {
	if o == nil || common.IsNil(o.StoredPaymentMethodId) {
		var ret string
		return ret
	}
	return *o.StoredPaymentMethodId
}

// GetStoredPaymentMethodIdOk returns a tuple with the StoredPaymentMethodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AchDetails) GetStoredPaymentMethodIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.StoredPaymentMethodId) {
		return nil, false
	}
	return o.StoredPaymentMethodId, true
}

// HasStoredPaymentMethodId returns a boolean if a field has been set.
func (o *AchDetails) HasStoredPaymentMethodId() bool {
	if o != nil && !common.IsNil(o.StoredPaymentMethodId) {
		return true
	}

	return false
}

// SetStoredPaymentMethodId gets a reference to the given string and assigns it to the StoredPaymentMethodId field.
func (o *AchDetails) SetStoredPaymentMethodId(v string) {
	o.StoredPaymentMethodId = &v
}

// GetTransferInstrumentId returns the TransferInstrumentId field value if set, zero value otherwise.
func (o *AchDetails) GetTransferInstrumentId() string {
	if o == nil || common.IsNil(o.TransferInstrumentId) {
		var ret string
		return ret
	}
	return *o.TransferInstrumentId
}

// GetTransferInstrumentIdOk returns a tuple with the TransferInstrumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AchDetails) GetTransferInstrumentIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TransferInstrumentId) {
		return nil, false
	}
	return o.TransferInstrumentId, true
}

// HasTransferInstrumentId returns a boolean if a field has been set.
func (o *AchDetails) HasTransferInstrumentId() bool {
	if o != nil && !common.IsNil(o.TransferInstrumentId) {
		return true
	}

	return false
}

// SetTransferInstrumentId gets a reference to the given string and assigns it to the TransferInstrumentId field.
func (o *AchDetails) SetTransferInstrumentId(v string) {
	o.TransferInstrumentId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AchDetails) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AchDetails) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AchDetails) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AchDetails) SetType(v string) {
	o.Type = &v
}

func (o AchDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AchDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.BankAccountNumber) {
		toSerialize["bankAccountNumber"] = o.BankAccountNumber
	}
	if !common.IsNil(o.BankAccountType) {
		toSerialize["bankAccountType"] = o.BankAccountType
	}
	if !common.IsNil(o.BankLocationId) {
		toSerialize["bankLocationId"] = o.BankLocationId
	}
	if !common.IsNil(o.CheckoutAttemptId) {
		toSerialize["checkoutAttemptId"] = o.CheckoutAttemptId
	}
	if !common.IsNil(o.EncryptedBankAccountNumber) {
		toSerialize["encryptedBankAccountNumber"] = o.EncryptedBankAccountNumber
	}
	if !common.IsNil(o.EncryptedBankLocationId) {
		toSerialize["encryptedBankLocationId"] = o.EncryptedBankLocationId
	}
	if !common.IsNil(o.OwnerName) {
		toSerialize["ownerName"] = o.OwnerName
	}
	if !common.IsNil(o.RecurringDetailReference) {
		toSerialize["recurringDetailReference"] = o.RecurringDetailReference
	}
	if !common.IsNil(o.StoredPaymentMethodId) {
		toSerialize["storedPaymentMethodId"] = o.StoredPaymentMethodId
	}
	if !common.IsNil(o.TransferInstrumentId) {
		toSerialize["transferInstrumentId"] = o.TransferInstrumentId
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableAchDetails struct {
	value *AchDetails
	isSet bool
}

func (v NullableAchDetails) Get() *AchDetails {
	return v.value
}

func (v *NullableAchDetails) Set(val *AchDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableAchDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableAchDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAchDetails(val *AchDetails) *NullableAchDetails {
	return &NullableAchDetails{value: val, isSet: true}
}

func (v NullableAchDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAchDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *AchDetails) isValidBankAccountType() bool {
	var allowedEnumValues = []string{"balance", "checking", "deposit", "general", "other", "payment", "savings"}
	for _, allowed := range allowedEnumValues {
		if o.GetBankAccountType() == allowed {
			return true
		}
	}
	return false
}
func (o *AchDetails) isValidType() bool {
	var allowedEnumValues = []string{"ach", "ach_plaid"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
