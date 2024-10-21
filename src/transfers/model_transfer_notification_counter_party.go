/*
Transfers API

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transfers

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v14/src/common"
)

// checks if the TransferNotificationCounterParty type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TransferNotificationCounterParty{}

// TransferNotificationCounterParty struct for TransferNotificationCounterParty
type TransferNotificationCounterParty struct {
	// The unique identifier of the counterparty [balance account](https://docs.adyen.com/api-explorer/balanceplatform/latest/post/balanceAccounts#responses-200-id).
	BalanceAccountId *string                           `json:"balanceAccountId,omitempty"`
	BankAccount      *BankAccountV3                    `json:"bankAccount,omitempty"`
	Card             *Card                             `json:"card,omitempty"`
	Merchant         *TransferNotificationMerchantData `json:"merchant,omitempty"`
	// The unique identifier of the counterparty [transfer instrument](https://docs.adyen.com/api-explorer/legalentity/latest/post/transferInstruments#responses-200-id).
	TransferInstrumentId *string `json:"transferInstrumentId,omitempty"`
}

// NewTransferNotificationCounterParty instantiates a new TransferNotificationCounterParty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferNotificationCounterParty() *TransferNotificationCounterParty {
	this := TransferNotificationCounterParty{}
	return &this
}

// NewTransferNotificationCounterPartyWithDefaults instantiates a new TransferNotificationCounterParty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferNotificationCounterPartyWithDefaults() *TransferNotificationCounterParty {
	this := TransferNotificationCounterParty{}
	return &this
}

// GetBalanceAccountId returns the BalanceAccountId field value if set, zero value otherwise.
func (o *TransferNotificationCounterParty) GetBalanceAccountId() string {
	if o == nil || common.IsNil(o.BalanceAccountId) {
		var ret string
		return ret
	}
	return *o.BalanceAccountId
}

// GetBalanceAccountIdOk returns a tuple with the BalanceAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferNotificationCounterParty) GetBalanceAccountIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.BalanceAccountId) {
		return nil, false
	}
	return o.BalanceAccountId, true
}

// HasBalanceAccountId returns a boolean if a field has been set.
func (o *TransferNotificationCounterParty) HasBalanceAccountId() bool {
	if o != nil && !common.IsNil(o.BalanceAccountId) {
		return true
	}

	return false
}

// SetBalanceAccountId gets a reference to the given string and assigns it to the BalanceAccountId field.
func (o *TransferNotificationCounterParty) SetBalanceAccountId(v string) {
	o.BalanceAccountId = &v
}

// GetBankAccount returns the BankAccount field value if set, zero value otherwise.
func (o *TransferNotificationCounterParty) GetBankAccount() BankAccountV3 {
	if o == nil || common.IsNil(o.BankAccount) {
		var ret BankAccountV3
		return ret
	}
	return *o.BankAccount
}

// GetBankAccountOk returns a tuple with the BankAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferNotificationCounterParty) GetBankAccountOk() (*BankAccountV3, bool) {
	if o == nil || common.IsNil(o.BankAccount) {
		return nil, false
	}
	return o.BankAccount, true
}

// HasBankAccount returns a boolean if a field has been set.
func (o *TransferNotificationCounterParty) HasBankAccount() bool {
	if o != nil && !common.IsNil(o.BankAccount) {
		return true
	}

	return false
}

// SetBankAccount gets a reference to the given BankAccountV3 and assigns it to the BankAccount field.
func (o *TransferNotificationCounterParty) SetBankAccount(v BankAccountV3) {
	o.BankAccount = &v
}

// GetCard returns the Card field value if set, zero value otherwise.
func (o *TransferNotificationCounterParty) GetCard() Card {
	if o == nil || common.IsNil(o.Card) {
		var ret Card
		return ret
	}
	return *o.Card
}

// GetCardOk returns a tuple with the Card field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferNotificationCounterParty) GetCardOk() (*Card, bool) {
	if o == nil || common.IsNil(o.Card) {
		return nil, false
	}
	return o.Card, true
}

// HasCard returns a boolean if a field has been set.
func (o *TransferNotificationCounterParty) HasCard() bool {
	if o != nil && !common.IsNil(o.Card) {
		return true
	}

	return false
}

// SetCard gets a reference to the given Card and assigns it to the Card field.
func (o *TransferNotificationCounterParty) SetCard(v Card) {
	o.Card = &v
}

// GetMerchant returns the Merchant field value if set, zero value otherwise.
func (o *TransferNotificationCounterParty) GetMerchant() TransferNotificationMerchantData {
	if o == nil || common.IsNil(o.Merchant) {
		var ret TransferNotificationMerchantData
		return ret
	}
	return *o.Merchant
}

// GetMerchantOk returns a tuple with the Merchant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferNotificationCounterParty) GetMerchantOk() (*TransferNotificationMerchantData, bool) {
	if o == nil || common.IsNil(o.Merchant) {
		return nil, false
	}
	return o.Merchant, true
}

// HasMerchant returns a boolean if a field has been set.
func (o *TransferNotificationCounterParty) HasMerchant() bool {
	if o != nil && !common.IsNil(o.Merchant) {
		return true
	}

	return false
}

// SetMerchant gets a reference to the given TransferNotificationMerchantData and assigns it to the Merchant field.
func (o *TransferNotificationCounterParty) SetMerchant(v TransferNotificationMerchantData) {
	o.Merchant = &v
}

// GetTransferInstrumentId returns the TransferInstrumentId field value if set, zero value otherwise.
func (o *TransferNotificationCounterParty) GetTransferInstrumentId() string {
	if o == nil || common.IsNil(o.TransferInstrumentId) {
		var ret string
		return ret
	}
	return *o.TransferInstrumentId
}

// GetTransferInstrumentIdOk returns a tuple with the TransferInstrumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferNotificationCounterParty) GetTransferInstrumentIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TransferInstrumentId) {
		return nil, false
	}
	return o.TransferInstrumentId, true
}

// HasTransferInstrumentId returns a boolean if a field has been set.
func (o *TransferNotificationCounterParty) HasTransferInstrumentId() bool {
	if o != nil && !common.IsNil(o.TransferInstrumentId) {
		return true
	}

	return false
}

// SetTransferInstrumentId gets a reference to the given string and assigns it to the TransferInstrumentId field.
func (o *TransferNotificationCounterParty) SetTransferInstrumentId(v string) {
	o.TransferInstrumentId = &v
}

func (o TransferNotificationCounterParty) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferNotificationCounterParty) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.BalanceAccountId) {
		toSerialize["balanceAccountId"] = o.BalanceAccountId
	}
	if !common.IsNil(o.BankAccount) {
		toSerialize["bankAccount"] = o.BankAccount
	}
	if !common.IsNil(o.Card) {
		toSerialize["card"] = o.Card
	}
	if !common.IsNil(o.Merchant) {
		toSerialize["merchant"] = o.Merchant
	}
	if !common.IsNil(o.TransferInstrumentId) {
		toSerialize["transferInstrumentId"] = o.TransferInstrumentId
	}
	return toSerialize, nil
}

type NullableTransferNotificationCounterParty struct {
	value *TransferNotificationCounterParty
	isSet bool
}

func (v NullableTransferNotificationCounterParty) Get() *TransferNotificationCounterParty {
	return v.value
}

func (v *NullableTransferNotificationCounterParty) Set(val *TransferNotificationCounterParty) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferNotificationCounterParty) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferNotificationCounterParty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferNotificationCounterParty(val *TransferNotificationCounterParty) *NullableTransferNotificationCounterParty {
	return &NullableTransferNotificationCounterParty{value: val, isSet: true}
}

func (v NullableTransferNotificationCounterParty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferNotificationCounterParty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
