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

// checks if the BalanceCheckResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BalanceCheckResponse{}

// BalanceCheckResponse struct for BalanceCheckResponse
type BalanceCheckResponse struct {
	// Contains additional information about the payment. Some data fields are included only if you select them first: Go to **Customer Area** > **Developers** > **Additional data**.
	AdditionalData *map[string]string `json:"additionalData,omitempty"`
	Balance        Amount             `json:"balance"`
	FraudResult    *FraudResult       `json:"fraudResult,omitempty"`
	// Adyen's 16-character reference associated with the transaction/request. This value is globally unique; quote it when communicating with us about this request.
	PspReference *string `json:"pspReference,omitempty"`
	// If the payment's authorisation is refused or an error occurs during authorisation, this field holds Adyen's mapped reason for the refusal or a description of the error. When a transaction fails, the authorisation response includes `resultCode` and `refusalReason` values.  For more information, see [Refusal reasons](https://docs.adyen.com/development-resources/refusal-reasons).
	RefusalReason *string `json:"refusalReason,omitempty"`
	// The result of the cancellation request.  Possible values:  * **Success** – Indicates that the balance check was successful. * **NotEnoughBalance** – Commonly indicates that the card did not have enough balance to pay the amount in the request, or that the currency of the balance on the card did not match the currency of the requested amount. * **Failed** – Indicates that the balance check failed.
	ResultCode       string  `json:"resultCode"`
	TransactionLimit *Amount `json:"transactionLimit,omitempty"`
}

// NewBalanceCheckResponse instantiates a new BalanceCheckResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBalanceCheckResponse(balance Amount, resultCode string) *BalanceCheckResponse {
	this := BalanceCheckResponse{}
	this.Balance = balance
	this.ResultCode = resultCode
	return &this
}

// NewBalanceCheckResponseWithDefaults instantiates a new BalanceCheckResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBalanceCheckResponseWithDefaults() *BalanceCheckResponse {
	this := BalanceCheckResponse{}
	return &this
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *BalanceCheckResponse) GetAdditionalData() map[string]string {
	if o == nil || common.IsNil(o.AdditionalData) {
		var ret map[string]string
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceCheckResponse) GetAdditionalDataOk() (*map[string]string, bool) {
	if o == nil || common.IsNil(o.AdditionalData) {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *BalanceCheckResponse) HasAdditionalData() bool {
	if o != nil && !common.IsNil(o.AdditionalData) {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given map[string]string and assigns it to the AdditionalData field.
func (o *BalanceCheckResponse) SetAdditionalData(v map[string]string) {
	o.AdditionalData = &v
}

// GetBalance returns the Balance field value
func (o *BalanceCheckResponse) GetBalance() Amount {
	if o == nil {
		var ret Amount
		return ret
	}

	return o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value
// and a boolean to check if the value has been set.
func (o *BalanceCheckResponse) GetBalanceOk() (*Amount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Balance, true
}

// SetBalance sets field value
func (o *BalanceCheckResponse) SetBalance(v Amount) {
	o.Balance = v
}

// GetFraudResult returns the FraudResult field value if set, zero value otherwise.
func (o *BalanceCheckResponse) GetFraudResult() FraudResult {
	if o == nil || common.IsNil(o.FraudResult) {
		var ret FraudResult
		return ret
	}
	return *o.FraudResult
}

// GetFraudResultOk returns a tuple with the FraudResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceCheckResponse) GetFraudResultOk() (*FraudResult, bool) {
	if o == nil || common.IsNil(o.FraudResult) {
		return nil, false
	}
	return o.FraudResult, true
}

// HasFraudResult returns a boolean if a field has been set.
func (o *BalanceCheckResponse) HasFraudResult() bool {
	if o != nil && !common.IsNil(o.FraudResult) {
		return true
	}

	return false
}

// SetFraudResult gets a reference to the given FraudResult and assigns it to the FraudResult field.
func (o *BalanceCheckResponse) SetFraudResult(v FraudResult) {
	o.FraudResult = &v
}

// GetPspReference returns the PspReference field value if set, zero value otherwise.
func (o *BalanceCheckResponse) GetPspReference() string {
	if o == nil || common.IsNil(o.PspReference) {
		var ret string
		return ret
	}
	return *o.PspReference
}

// GetPspReferenceOk returns a tuple with the PspReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceCheckResponse) GetPspReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.PspReference) {
		return nil, false
	}
	return o.PspReference, true
}

// HasPspReference returns a boolean if a field has been set.
func (o *BalanceCheckResponse) HasPspReference() bool {
	if o != nil && !common.IsNil(o.PspReference) {
		return true
	}

	return false
}

// SetPspReference gets a reference to the given string and assigns it to the PspReference field.
func (o *BalanceCheckResponse) SetPspReference(v string) {
	o.PspReference = &v
}

// GetRefusalReason returns the RefusalReason field value if set, zero value otherwise.
func (o *BalanceCheckResponse) GetRefusalReason() string {
	if o == nil || common.IsNil(o.RefusalReason) {
		var ret string
		return ret
	}
	return *o.RefusalReason
}

// GetRefusalReasonOk returns a tuple with the RefusalReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceCheckResponse) GetRefusalReasonOk() (*string, bool) {
	if o == nil || common.IsNil(o.RefusalReason) {
		return nil, false
	}
	return o.RefusalReason, true
}

// HasRefusalReason returns a boolean if a field has been set.
func (o *BalanceCheckResponse) HasRefusalReason() bool {
	if o != nil && !common.IsNil(o.RefusalReason) {
		return true
	}

	return false
}

// SetRefusalReason gets a reference to the given string and assigns it to the RefusalReason field.
func (o *BalanceCheckResponse) SetRefusalReason(v string) {
	o.RefusalReason = &v
}

// GetResultCode returns the ResultCode field value
func (o *BalanceCheckResponse) GetResultCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value
// and a boolean to check if the value has been set.
func (o *BalanceCheckResponse) GetResultCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResultCode, true
}

// SetResultCode sets field value
func (o *BalanceCheckResponse) SetResultCode(v string) {
	o.ResultCode = v
}

// GetTransactionLimit returns the TransactionLimit field value if set, zero value otherwise.
func (o *BalanceCheckResponse) GetTransactionLimit() Amount {
	if o == nil || common.IsNil(o.TransactionLimit) {
		var ret Amount
		return ret
	}
	return *o.TransactionLimit
}

// GetTransactionLimitOk returns a tuple with the TransactionLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceCheckResponse) GetTransactionLimitOk() (*Amount, bool) {
	if o == nil || common.IsNil(o.TransactionLimit) {
		return nil, false
	}
	return o.TransactionLimit, true
}

// HasTransactionLimit returns a boolean if a field has been set.
func (o *BalanceCheckResponse) HasTransactionLimit() bool {
	if o != nil && !common.IsNil(o.TransactionLimit) {
		return true
	}

	return false
}

// SetTransactionLimit gets a reference to the given Amount and assigns it to the TransactionLimit field.
func (o *BalanceCheckResponse) SetTransactionLimit(v Amount) {
	o.TransactionLimit = &v
}

func (o BalanceCheckResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BalanceCheckResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AdditionalData) {
		toSerialize["additionalData"] = o.AdditionalData
	}
	toSerialize["balance"] = o.Balance
	if !common.IsNil(o.FraudResult) {
		toSerialize["fraudResult"] = o.FraudResult
	}
	if !common.IsNil(o.PspReference) {
		toSerialize["pspReference"] = o.PspReference
	}
	if !common.IsNil(o.RefusalReason) {
		toSerialize["refusalReason"] = o.RefusalReason
	}
	toSerialize["resultCode"] = o.ResultCode
	if !common.IsNil(o.TransactionLimit) {
		toSerialize["transactionLimit"] = o.TransactionLimit
	}
	return toSerialize, nil
}

type NullableBalanceCheckResponse struct {
	value *BalanceCheckResponse
	isSet bool
}

func (v NullableBalanceCheckResponse) Get() *BalanceCheckResponse {
	return v.value
}

func (v *NullableBalanceCheckResponse) Set(val *BalanceCheckResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBalanceCheckResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBalanceCheckResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBalanceCheckResponse(val *BalanceCheckResponse) *NullableBalanceCheckResponse {
	return &NullableBalanceCheckResponse{value: val, isSet: true}
}

func (v NullableBalanceCheckResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBalanceCheckResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *BalanceCheckResponse) isValidResultCode() bool {
	var allowedEnumValues = []string{"Success", "NotEnoughBalance", "Failed"}
	for _, allowed := range allowedEnumValues {
		if o.GetResultCode() == allowed {
			return true
		}
	}
	return false
}
