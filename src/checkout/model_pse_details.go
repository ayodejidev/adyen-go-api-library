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

// checks if the PseDetails type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PseDetails{}

// PseDetails struct for PseDetails
type PseDetails struct {
	// The shopper's bank.
	Bank string `json:"bank"`
	// The checkout attempt identifier.
	CheckoutAttemptId *string `json:"checkoutAttemptId,omitempty"`
	// The client type.
	ClientType string `json:"clientType"`
	// The identification code.
	Identification string `json:"identification"`
	// The identification type.
	IdentificationType string `json:"identificationType"`
	// The payment method type.
	Type *string `json:"type,omitempty"`
}

// NewPseDetails instantiates a new PseDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPseDetails(bank string, clientType string, identification string, identificationType string) *PseDetails {
	this := PseDetails{}
	this.Bank = bank
	this.ClientType = clientType
	this.Identification = identification
	this.IdentificationType = identificationType
	return &this
}

// NewPseDetailsWithDefaults instantiates a new PseDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPseDetailsWithDefaults() *PseDetails {
	this := PseDetails{}
	return &this
}

// GetBank returns the Bank field value
func (o *PseDetails) GetBank() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bank
}

// GetBankOk returns a tuple with the Bank field value
// and a boolean to check if the value has been set.
func (o *PseDetails) GetBankOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bank, true
}

// SetBank sets field value
func (o *PseDetails) SetBank(v string) {
	o.Bank = v
}

// GetCheckoutAttemptId returns the CheckoutAttemptId field value if set, zero value otherwise.
func (o *PseDetails) GetCheckoutAttemptId() string {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		var ret string
		return ret
	}
	return *o.CheckoutAttemptId
}

// GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PseDetails) GetCheckoutAttemptIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		return nil, false
	}
	return o.CheckoutAttemptId, true
}

// HasCheckoutAttemptId returns a boolean if a field has been set.
func (o *PseDetails) HasCheckoutAttemptId() bool {
	if o != nil && !common.IsNil(o.CheckoutAttemptId) {
		return true
	}

	return false
}

// SetCheckoutAttemptId gets a reference to the given string and assigns it to the CheckoutAttemptId field.
func (o *PseDetails) SetCheckoutAttemptId(v string) {
	o.CheckoutAttemptId = &v
}

// GetClientType returns the ClientType field value
func (o *PseDetails) GetClientType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientType
}

// GetClientTypeOk returns a tuple with the ClientType field value
// and a boolean to check if the value has been set.
func (o *PseDetails) GetClientTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientType, true
}

// SetClientType sets field value
func (o *PseDetails) SetClientType(v string) {
	o.ClientType = v
}

// GetIdentification returns the Identification field value
func (o *PseDetails) GetIdentification() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identification
}

// GetIdentificationOk returns a tuple with the Identification field value
// and a boolean to check if the value has been set.
func (o *PseDetails) GetIdentificationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identification, true
}

// SetIdentification sets field value
func (o *PseDetails) SetIdentification(v string) {
	o.Identification = v
}

// GetIdentificationType returns the IdentificationType field value
func (o *PseDetails) GetIdentificationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdentificationType
}

// GetIdentificationTypeOk returns a tuple with the IdentificationType field value
// and a boolean to check if the value has been set.
func (o *PseDetails) GetIdentificationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdentificationType, true
}

// SetIdentificationType sets field value
func (o *PseDetails) SetIdentificationType(v string) {
	o.IdentificationType = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PseDetails) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PseDetails) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PseDetails) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PseDetails) SetType(v string) {
	o.Type = &v
}

func (o PseDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PseDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bank"] = o.Bank
	if !common.IsNil(o.CheckoutAttemptId) {
		toSerialize["checkoutAttemptId"] = o.CheckoutAttemptId
	}
	toSerialize["clientType"] = o.ClientType
	toSerialize["identification"] = o.Identification
	toSerialize["identificationType"] = o.IdentificationType
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullablePseDetails struct {
	value *PseDetails
	isSet bool
}

func (v NullablePseDetails) Get() *PseDetails {
	return v.value
}

func (v *NullablePseDetails) Set(val *PseDetails) {
	v.value = val
	v.isSet = true
}

func (v NullablePseDetails) IsSet() bool {
	return v.isSet
}

func (v *NullablePseDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePseDetails(val *PseDetails) *NullablePseDetails {
	return &NullablePseDetails{value: val, isSet: true}
}

func (v NullablePseDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePseDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *PseDetails) isValidType() bool {
	var allowedEnumValues = []string{"pse_payulatam"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
