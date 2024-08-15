/*
Adyen Balance Control API

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balancecontrol

import (
	"encoding/json"
	"time"

	"github.com/adyen/adyen-go-api-library/v12/src/common"
)

// checks if the BalanceTransferResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BalanceTransferResponse{}

// BalanceTransferResponse struct for BalanceTransferResponse
type BalanceTransferResponse struct {
	Amount Amount `json:"amount"`
	// The date when the balance transfer was requested.
	CreatedAt time.Time `json:"createdAt"`
	// A human-readable description for the transfer. You can use alphanumeric characters and hyphens. We recommend sending a maximum of 140 characters, otherwise the description may be truncated.
	Description *string `json:"description,omitempty"`
	// The unique identifier of the source merchant account from which funds are deducted.
	FromMerchant string `json:"fromMerchant"`
	// Adyen's 16-character string reference associated with the balance transfer.
	PspReference string `json:"pspReference"`
	// A reference for the balance transfer. If you don't provide this in the request, Adyen generates a unique reference. Maximum length: 80 characters.
	Reference *string `json:"reference,omitempty"`
	// The status of the balance transfer. Possible values: **transferred**, **failed**, **error**, and **notEnoughBalance**.
	Status string `json:"status"`
	// The unique identifier of the destination merchant account from which funds are transferred.
	ToMerchant string `json:"toMerchant"`
	// The type of balance transfer. Possible values: **tax**, **fee**, **terminalSale**, **credit**, **debit**, and **adjustment**.
	Type string `json:"type"`
}

// NewBalanceTransferResponse instantiates a new BalanceTransferResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBalanceTransferResponse(amount Amount, createdAt time.Time, fromMerchant string, pspReference string, status string, toMerchant string, type_ string) *BalanceTransferResponse {
	this := BalanceTransferResponse{}
	this.Amount = amount
	this.CreatedAt = createdAt
	this.FromMerchant = fromMerchant
	this.PspReference = pspReference
	this.Status = status
	this.ToMerchant = toMerchant
	this.Type = type_
	return &this
}

// NewBalanceTransferResponseWithDefaults instantiates a new BalanceTransferResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBalanceTransferResponseWithDefaults() *BalanceTransferResponse {
	this := BalanceTransferResponse{}
	return &this
}

// GetAmount returns the Amount field value
func (o *BalanceTransferResponse) GetAmount() Amount {
	if o == nil {
		var ret Amount
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *BalanceTransferResponse) GetAmountOk() (*Amount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *BalanceTransferResponse) SetAmount(v Amount) {
	o.Amount = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *BalanceTransferResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *BalanceTransferResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *BalanceTransferResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BalanceTransferResponse) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceTransferResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BalanceTransferResponse) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BalanceTransferResponse) SetDescription(v string) {
	o.Description = &v
}

// GetFromMerchant returns the FromMerchant field value
func (o *BalanceTransferResponse) GetFromMerchant() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FromMerchant
}

// GetFromMerchantOk returns a tuple with the FromMerchant field value
// and a boolean to check if the value has been set.
func (o *BalanceTransferResponse) GetFromMerchantOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromMerchant, true
}

// SetFromMerchant sets field value
func (o *BalanceTransferResponse) SetFromMerchant(v string) {
	o.FromMerchant = v
}

// GetPspReference returns the PspReference field value
func (o *BalanceTransferResponse) GetPspReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PspReference
}

// GetPspReferenceOk returns a tuple with the PspReference field value
// and a boolean to check if the value has been set.
func (o *BalanceTransferResponse) GetPspReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PspReference, true
}

// SetPspReference sets field value
func (o *BalanceTransferResponse) SetPspReference(v string) {
	o.PspReference = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *BalanceTransferResponse) GetReference() string {
	if o == nil || common.IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceTransferResponse) GetReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *BalanceTransferResponse) HasReference() bool {
	if o != nil && !common.IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *BalanceTransferResponse) SetReference(v string) {
	o.Reference = &v
}

// GetStatus returns the Status field value
func (o *BalanceTransferResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BalanceTransferResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BalanceTransferResponse) SetStatus(v string) {
	o.Status = v
}

// GetToMerchant returns the ToMerchant field value
func (o *BalanceTransferResponse) GetToMerchant() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToMerchant
}

// GetToMerchantOk returns a tuple with the ToMerchant field value
// and a boolean to check if the value has been set.
func (o *BalanceTransferResponse) GetToMerchantOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToMerchant, true
}

// SetToMerchant sets field value
func (o *BalanceTransferResponse) SetToMerchant(v string) {
	o.ToMerchant = v
}

// GetType returns the Type field value
func (o *BalanceTransferResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BalanceTransferResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BalanceTransferResponse) SetType(v string) {
	o.Type = v
}

func (o BalanceTransferResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BalanceTransferResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	toSerialize["createdAt"] = o.CreatedAt
	if !common.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["fromMerchant"] = o.FromMerchant
	toSerialize["pspReference"] = o.PspReference
	if !common.IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	toSerialize["status"] = o.Status
	toSerialize["toMerchant"] = o.ToMerchant
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableBalanceTransferResponse struct {
	value *BalanceTransferResponse
	isSet bool
}

func (v NullableBalanceTransferResponse) Get() *BalanceTransferResponse {
	return v.value
}

func (v *NullableBalanceTransferResponse) Set(val *BalanceTransferResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBalanceTransferResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBalanceTransferResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBalanceTransferResponse(val *BalanceTransferResponse) *NullableBalanceTransferResponse {
	return &NullableBalanceTransferResponse{value: val, isSet: true}
}

func (v NullableBalanceTransferResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBalanceTransferResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *BalanceTransferResponse) isValidStatus() bool {
	var allowedEnumValues = []string{"error", "failed", "notEnoughBalance", "transferred"}
	for _, allowed := range allowedEnumValues {
		if o.GetStatus() == allowed {
			return true
		}
	}
	return false
}
func (o *BalanceTransferResponse) isValidType() bool {
	var allowedEnumValues = []string{"tax", "fee", "terminalSale", "credit", "debit", "adjustment"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
