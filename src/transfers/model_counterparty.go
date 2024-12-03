/*
Transfers API

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transfers

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v15/src/common"
)

// checks if the Counterparty type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Counterparty{}

// Counterparty struct for Counterparty
type Counterparty struct {
	// The identifier of the receiving account holder. The payout will default to the primary balance account of this account holder if no `balanceAccountId` is provided.
	AccountHolderId *string `json:"accountHolderId,omitempty"`
	// The identifier of the balance account that belongs to the receiving account holder.
	BalanceAccountId *string `json:"balanceAccountId,omitempty"`
	// The identifier of the transfer instrument that belongs to the legal entity of the account holder.
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

// GetAccountHolderId returns the AccountHolderId field value if set, zero value otherwise.
func (o *Counterparty) GetAccountHolderId() string {
	if o == nil || common.IsNil(o.AccountHolderId) {
		var ret string
		return ret
	}
	return *o.AccountHolderId
}

// GetAccountHolderIdOk returns a tuple with the AccountHolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Counterparty) GetAccountHolderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.AccountHolderId) {
		return nil, false
	}
	return o.AccountHolderId, true
}

// HasAccountHolderId returns a boolean if a field has been set.
func (o *Counterparty) HasAccountHolderId() bool {
	if o != nil && !common.IsNil(o.AccountHolderId) {
		return true
	}

	return false
}

// SetAccountHolderId gets a reference to the given string and assigns it to the AccountHolderId field.
func (o *Counterparty) SetAccountHolderId(v string) {
	o.AccountHolderId = &v
}

// GetBalanceAccountId returns the BalanceAccountId field value if set, zero value otherwise.
func (o *Counterparty) GetBalanceAccountId() string {
	if o == nil || common.IsNil(o.BalanceAccountId) {
		var ret string
		return ret
	}
	return *o.BalanceAccountId
}

// GetBalanceAccountIdOk returns a tuple with the BalanceAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Counterparty) GetBalanceAccountIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.BalanceAccountId) {
		return nil, false
	}
	return o.BalanceAccountId, true
}

// HasBalanceAccountId returns a boolean if a field has been set.
func (o *Counterparty) HasBalanceAccountId() bool {
	if o != nil && !common.IsNil(o.BalanceAccountId) {
		return true
	}

	return false
}

// SetBalanceAccountId gets a reference to the given string and assigns it to the BalanceAccountId field.
func (o *Counterparty) SetBalanceAccountId(v string) {
	o.BalanceAccountId = &v
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
	if !common.IsNil(o.AccountHolderId) {
		toSerialize["accountHolderId"] = o.AccountHolderId
	}
	if !common.IsNil(o.BalanceAccountId) {
		toSerialize["balanceAccountId"] = o.BalanceAccountId
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
