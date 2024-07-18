/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v11/src/common"
)

// checks if the Counterparty type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Counterparty{}

// Counterparty struct for Counterparty
type Counterparty struct {
	BankAccount *BankAccount `json:"bankAccount,omitempty"`
	// The unique identifier of the [transfer instrument](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/transferInstruments__resParam_id).
	TransferInstrumentId *string `json:"transferInstrumentId,omitempty"`
}

// NewCounterparty instantiates a new Counterparty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCounterparty() *Counterparty {
	this := Counterparty{}
	return &this
}

// NewCounterpartyWithDefaults instantiates a new Counterparty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCounterpartyWithDefaults() *Counterparty {
	this := Counterparty{}
	return &this
}

// GetBankAccount returns the BankAccount field value if set, zero value otherwise.
func (o *Counterparty) GetBankAccount() BankAccount {
	if o == nil || common.IsNil(o.BankAccount) {
		var ret BankAccount
		return ret
	}
	return *o.BankAccount
}

// GetBankAccountOk returns a tuple with the BankAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Counterparty) GetBankAccountOk() (*BankAccount, bool) {
	if o == nil || common.IsNil(o.BankAccount) {
		return nil, false
	}
	return o.BankAccount, true
}

// HasBankAccount returns a boolean if a field has been set.
func (o *Counterparty) HasBankAccount() bool {
	if o != nil && !common.IsNil(o.BankAccount) {
		return true
	}

	return false
}

// SetBankAccount gets a reference to the given BankAccount and assigns it to the BankAccount field.
func (o *Counterparty) SetBankAccount(v BankAccount) {
	o.BankAccount = &v
}

// GetTransferInstrumentId returns the TransferInstrumentId field value if set, zero value otherwise.
func (o *Counterparty) GetTransferInstrumentId() string {
	if o == nil || common.IsNil(o.TransferInstrumentId) {
		var ret string
		return ret
	}
	return *o.TransferInstrumentId
}

// GetTransferInstrumentIdOk returns a tuple with the TransferInstrumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Counterparty) GetTransferInstrumentIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TransferInstrumentId) {
		return nil, false
	}
	return o.TransferInstrumentId, true
}

// HasTransferInstrumentId returns a boolean if a field has been set.
func (o *Counterparty) HasTransferInstrumentId() bool {
	if o != nil && !common.IsNil(o.TransferInstrumentId) {
		return true
	}

	return false
}

// SetTransferInstrumentId gets a reference to the given string and assigns it to the TransferInstrumentId field.
func (o *Counterparty) SetTransferInstrumentId(v string) {
	o.TransferInstrumentId = &v
}

func (o Counterparty) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Counterparty) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.BankAccount) {
		toSerialize["bankAccount"] = o.BankAccount
	}
	if !common.IsNil(o.TransferInstrumentId) {
		toSerialize["transferInstrumentId"] = o.TransferInstrumentId
	}
	return toSerialize, nil
}

type NullableCounterparty struct {
	value *Counterparty
	isSet bool
}

func (v NullableCounterparty) Get() *Counterparty {
	return v.value
}

func (v *NullableCounterparty) Set(val *Counterparty) {
	v.value = val
	v.isSet = true
}

func (v NullableCounterparty) IsSet() bool {
	return v.isSet
}

func (v *NullableCounterparty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCounterparty(val *Counterparty) *NullableCounterparty {
	return &NullableCounterparty{value: val, isSet: true}
}

func (v NullableCounterparty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCounterparty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
