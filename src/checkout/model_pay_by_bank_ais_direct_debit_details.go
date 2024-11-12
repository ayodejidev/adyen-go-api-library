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

// checks if the PayByBankAISDirectDebitDetails type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PayByBankAISDirectDebitDetails{}

// PayByBankAISDirectDebitDetails struct for PayByBankAISDirectDebitDetails
type PayByBankAISDirectDebitDetails struct {
	// The checkout attempt identifier.
	CheckoutAttemptId *string `json:"checkoutAttemptId,omitempty"`
	// This is the `recurringDetailReference` returned in the response when you created the token.
	// Deprecated since Adyen Checkout API v49
	// Use `storedPaymentMethodId` instead.
	RecurringDetailReference *string `json:"recurringDetailReference,omitempty"`
	// This is the `recurringDetailReference` returned in the response when you created the token.
	StoredPaymentMethodId *string `json:"storedPaymentMethodId,omitempty"`
	// **paybybank_**
	Type string `json:"type"`
}

// NewPayByBankAISDirectDebitDetails instantiates a new PayByBankAISDirectDebitDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayByBankAISDirectDebitDetails(type_ string) *PayByBankAISDirectDebitDetails {
	this := PayByBankAISDirectDebitDetails{}
	this.Type = type_
	return &this
}

// NewPayByBankAISDirectDebitDetailsWithDefaults instantiates a new PayByBankAISDirectDebitDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayByBankAISDirectDebitDetailsWithDefaults() *PayByBankAISDirectDebitDetails {
	this := PayByBankAISDirectDebitDetails{}
	var type_ string = "paybybank_AIS_DD"
	this.Type = type_
	return &this
}

// GetCheckoutAttemptId returns the CheckoutAttemptId field value if set, zero value otherwise.
func (o *PayByBankAISDirectDebitDetails) GetCheckoutAttemptId() string {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		var ret string
		return ret
	}
	return *o.CheckoutAttemptId
}

// GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayByBankAISDirectDebitDetails) GetCheckoutAttemptIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		return nil, false
	}
	return o.CheckoutAttemptId, true
}

// HasCheckoutAttemptId returns a boolean if a field has been set.
func (o *PayByBankAISDirectDebitDetails) HasCheckoutAttemptId() bool {
	if o != nil && !common.IsNil(o.CheckoutAttemptId) {
		return true
	}

	return false
}

// SetCheckoutAttemptId gets a reference to the given string and assigns it to the CheckoutAttemptId field.
func (o *PayByBankAISDirectDebitDetails) SetCheckoutAttemptId(v string) {
	o.CheckoutAttemptId = &v
}

// GetRecurringDetailReference returns the RecurringDetailReference field value if set, zero value otherwise.
// Deprecated since Adyen Checkout API v49
// Use `storedPaymentMethodId` instead.
func (o *PayByBankAISDirectDebitDetails) GetRecurringDetailReference() string {
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
func (o *PayByBankAISDirectDebitDetails) GetRecurringDetailReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.RecurringDetailReference) {
		return nil, false
	}
	return o.RecurringDetailReference, true
}

// HasRecurringDetailReference returns a boolean if a field has been set.
func (o *PayByBankAISDirectDebitDetails) HasRecurringDetailReference() bool {
	if o != nil && !common.IsNil(o.RecurringDetailReference) {
		return true
	}

	return false
}

// SetRecurringDetailReference gets a reference to the given string and assigns it to the RecurringDetailReference field.
// Deprecated since Adyen Checkout API v49
// Use `storedPaymentMethodId` instead.
func (o *PayByBankAISDirectDebitDetails) SetRecurringDetailReference(v string) {
	o.RecurringDetailReference = &v
}

// GetStoredPaymentMethodId returns the StoredPaymentMethodId field value if set, zero value otherwise.
func (o *PayByBankAISDirectDebitDetails) GetStoredPaymentMethodId() string {
	if o == nil || common.IsNil(o.StoredPaymentMethodId) {
		var ret string
		return ret
	}
	return *o.StoredPaymentMethodId
}

// GetStoredPaymentMethodIdOk returns a tuple with the StoredPaymentMethodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayByBankAISDirectDebitDetails) GetStoredPaymentMethodIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.StoredPaymentMethodId) {
		return nil, false
	}
	return o.StoredPaymentMethodId, true
}

// HasStoredPaymentMethodId returns a boolean if a field has been set.
func (o *PayByBankAISDirectDebitDetails) HasStoredPaymentMethodId() bool {
	if o != nil && !common.IsNil(o.StoredPaymentMethodId) {
		return true
	}

	return false
}

// SetStoredPaymentMethodId gets a reference to the given string and assigns it to the StoredPaymentMethodId field.
func (o *PayByBankAISDirectDebitDetails) SetStoredPaymentMethodId(v string) {
	o.StoredPaymentMethodId = &v
}

// GetType returns the Type field value
func (o *PayByBankAISDirectDebitDetails) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PayByBankAISDirectDebitDetails) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PayByBankAISDirectDebitDetails) SetType(v string) {
	o.Type = v
}

func (o PayByBankAISDirectDebitDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PayByBankAISDirectDebitDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CheckoutAttemptId) {
		toSerialize["checkoutAttemptId"] = o.CheckoutAttemptId
	}
	if !common.IsNil(o.RecurringDetailReference) {
		toSerialize["recurringDetailReference"] = o.RecurringDetailReference
	}
	if !common.IsNil(o.StoredPaymentMethodId) {
		toSerialize["storedPaymentMethodId"] = o.StoredPaymentMethodId
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullablePayByBankAISDirectDebitDetails struct {
	value *PayByBankAISDirectDebitDetails
	isSet bool
}

func (v NullablePayByBankAISDirectDebitDetails) Get() *PayByBankAISDirectDebitDetails {
	return v.value
}

func (v *NullablePayByBankAISDirectDebitDetails) Set(val *PayByBankAISDirectDebitDetails) {
	v.value = val
	v.isSet = true
}

func (v NullablePayByBankAISDirectDebitDetails) IsSet() bool {
	return v.isSet
}

func (v *NullablePayByBankAISDirectDebitDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayByBankAISDirectDebitDetails(val *PayByBankAISDirectDebitDetails) *NullablePayByBankAISDirectDebitDetails {
	return &NullablePayByBankAISDirectDebitDetails{value: val, isSet: true}
}

func (v NullablePayByBankAISDirectDebitDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayByBankAISDirectDebitDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *PayByBankAISDirectDebitDetails) isValidType() bool {
	var allowedEnumValues = []string{"paybybank_AIS_DD"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
