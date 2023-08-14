/*
Transfer webhooks

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transferwebhook

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// checks if the BRLocalAccountIdentification type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BRLocalAccountIdentification{}

// BRLocalAccountIdentification struct for BRLocalAccountIdentification
type BRLocalAccountIdentification struct {
	// The bank account number, without separators or whitespace.
	AccountNumber string `json:"accountNumber"`
	// The 3-digit bank code, with leading zeros.
	BankCode string `json:"bankCode"`
	// The bank account branch number, without separators or whitespace.
	BranchNumber string `json:"branchNumber"`
	// **brLocal**
	Type string `json:"type"`
}

// NewBRLocalAccountIdentification instantiates a new BRLocalAccountIdentification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBRLocalAccountIdentification(accountNumber string, bankCode string, branchNumber string, type_ string) *BRLocalAccountIdentification {
	this := BRLocalAccountIdentification{}
	this.AccountNumber = accountNumber
	this.BankCode = bankCode
	this.BranchNumber = branchNumber
	this.Type = type_
	return &this
}

// NewBRLocalAccountIdentificationWithDefaults instantiates a new BRLocalAccountIdentification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBRLocalAccountIdentificationWithDefaults() *BRLocalAccountIdentification {
	this := BRLocalAccountIdentification{}
	var type_ string = "brLocal"
	this.Type = type_
	return &this
}

// GetAccountNumber returns the AccountNumber field value
func (o *BRLocalAccountIdentification) GetAccountNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value
// and a boolean to check if the value has been set.
func (o *BRLocalAccountIdentification) GetAccountNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountNumber, true
}

// SetAccountNumber sets field value
func (o *BRLocalAccountIdentification) SetAccountNumber(v string) {
	o.AccountNumber = v
}

// GetBankCode returns the BankCode field value
func (o *BRLocalAccountIdentification) GetBankCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BankCode
}

// GetBankCodeOk returns a tuple with the BankCode field value
// and a boolean to check if the value has been set.
func (o *BRLocalAccountIdentification) GetBankCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BankCode, true
}

// SetBankCode sets field value
func (o *BRLocalAccountIdentification) SetBankCode(v string) {
	o.BankCode = v
}

// GetBranchNumber returns the BranchNumber field value
func (o *BRLocalAccountIdentification) GetBranchNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BranchNumber
}

// GetBranchNumberOk returns a tuple with the BranchNumber field value
// and a boolean to check if the value has been set.
func (o *BRLocalAccountIdentification) GetBranchNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BranchNumber, true
}

// SetBranchNumber sets field value
func (o *BRLocalAccountIdentification) SetBranchNumber(v string) {
	o.BranchNumber = v
}

// GetType returns the Type field value
func (o *BRLocalAccountIdentification) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BRLocalAccountIdentification) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BRLocalAccountIdentification) SetType(v string) {
	o.Type = v
}

func (o BRLocalAccountIdentification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BRLocalAccountIdentification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountNumber"] = o.AccountNumber
	toSerialize["bankCode"] = o.BankCode
	toSerialize["branchNumber"] = o.BranchNumber
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableBRLocalAccountIdentification struct {
	value *BRLocalAccountIdentification
	isSet bool
}

func (v NullableBRLocalAccountIdentification) Get() *BRLocalAccountIdentification {
	return v.value
}

func (v *NullableBRLocalAccountIdentification) Set(val *BRLocalAccountIdentification) {
	v.value = val
	v.isSet = true
}

func (v NullableBRLocalAccountIdentification) IsSet() bool {
	return v.isSet
}

func (v *NullableBRLocalAccountIdentification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBRLocalAccountIdentification(val *BRLocalAccountIdentification) *NullableBRLocalAccountIdentification {
	return &NullableBRLocalAccountIdentification{value: val, isSet: true}
}

func (v NullableBRLocalAccountIdentification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBRLocalAccountIdentification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *BRLocalAccountIdentification) isValidType() bool {
	var allowedEnumValues = []string{"brLocal"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
