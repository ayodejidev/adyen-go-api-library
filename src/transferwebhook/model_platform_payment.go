/*
Transfer webhooks

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transferwebhook

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v11/src/common"
)

// checks if the PlatformPayment type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PlatformPayment{}

// PlatformPayment struct for PlatformPayment
type PlatformPayment struct {
	// The capture's merchant reference included in the transfer.
	ModificationMerchantReference *string `json:"modificationMerchantReference,omitempty"`
	// The capture reference included in the transfer.
	ModificationPspReference *string `json:"modificationPspReference,omitempty"`
	// The payment's merchant reference included in the transfer.
	PaymentMerchantReference *string `json:"paymentMerchantReference,omitempty"`
	// The type of the related split.
	PlatformPaymentType *string `json:"platformPaymentType,omitempty"`
	// The payment reference included in the transfer.
	PspPaymentReference *string `json:"pspPaymentReference,omitempty"`
	// **platformPayment**
	Type *string `json:"type,omitempty"`
}

// NewPlatformPayment instantiates a new PlatformPayment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlatformPayment() *PlatformPayment {
	this := PlatformPayment{}
	var type_ string = "platformPayment"
	this.Type = &type_
	return &this
}

// NewPlatformPaymentWithDefaults instantiates a new PlatformPayment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlatformPaymentWithDefaults() *PlatformPayment {
	this := PlatformPayment{}
	var type_ string = "platformPayment"
	this.Type = &type_
	return &this
}

// GetModificationMerchantReference returns the ModificationMerchantReference field value if set, zero value otherwise.
func (o *PlatformPayment) GetModificationMerchantReference() string {
	if o == nil || common.IsNil(o.ModificationMerchantReference) {
		var ret string
		return ret
	}
	return *o.ModificationMerchantReference
}

// GetModificationMerchantReferenceOk returns a tuple with the ModificationMerchantReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformPayment) GetModificationMerchantReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.ModificationMerchantReference) {
		return nil, false
	}
	return o.ModificationMerchantReference, true
}

// HasModificationMerchantReference returns a boolean if a field has been set.
func (o *PlatformPayment) HasModificationMerchantReference() bool {
	if o != nil && !common.IsNil(o.ModificationMerchantReference) {
		return true
	}

	return false
}

// SetModificationMerchantReference gets a reference to the given string and assigns it to the ModificationMerchantReference field.
func (o *PlatformPayment) SetModificationMerchantReference(v string) {
	o.ModificationMerchantReference = &v
}

// GetModificationPspReference returns the ModificationPspReference field value if set, zero value otherwise.
func (o *PlatformPayment) GetModificationPspReference() string {
	if o == nil || common.IsNil(o.ModificationPspReference) {
		var ret string
		return ret
	}
	return *o.ModificationPspReference
}

// GetModificationPspReferenceOk returns a tuple with the ModificationPspReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformPayment) GetModificationPspReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.ModificationPspReference) {
		return nil, false
	}
	return o.ModificationPspReference, true
}

// HasModificationPspReference returns a boolean if a field has been set.
func (o *PlatformPayment) HasModificationPspReference() bool {
	if o != nil && !common.IsNil(o.ModificationPspReference) {
		return true
	}

	return false
}

// SetModificationPspReference gets a reference to the given string and assigns it to the ModificationPspReference field.
func (o *PlatformPayment) SetModificationPspReference(v string) {
	o.ModificationPspReference = &v
}

// GetPaymentMerchantReference returns the PaymentMerchantReference field value if set, zero value otherwise.
func (o *PlatformPayment) GetPaymentMerchantReference() string {
	if o == nil || common.IsNil(o.PaymentMerchantReference) {
		var ret string
		return ret
	}
	return *o.PaymentMerchantReference
}

// GetPaymentMerchantReferenceOk returns a tuple with the PaymentMerchantReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformPayment) GetPaymentMerchantReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.PaymentMerchantReference) {
		return nil, false
	}
	return o.PaymentMerchantReference, true
}

// HasPaymentMerchantReference returns a boolean if a field has been set.
func (o *PlatformPayment) HasPaymentMerchantReference() bool {
	if o != nil && !common.IsNil(o.PaymentMerchantReference) {
		return true
	}

	return false
}

// SetPaymentMerchantReference gets a reference to the given string and assigns it to the PaymentMerchantReference field.
func (o *PlatformPayment) SetPaymentMerchantReference(v string) {
	o.PaymentMerchantReference = &v
}

// GetPlatformPaymentType returns the PlatformPaymentType field value if set, zero value otherwise.
func (o *PlatformPayment) GetPlatformPaymentType() string {
	if o == nil || common.IsNil(o.PlatformPaymentType) {
		var ret string
		return ret
	}
	return *o.PlatformPaymentType
}

// GetPlatformPaymentTypeOk returns a tuple with the PlatformPaymentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformPayment) GetPlatformPaymentTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.PlatformPaymentType) {
		return nil, false
	}
	return o.PlatformPaymentType, true
}

// HasPlatformPaymentType returns a boolean if a field has been set.
func (o *PlatformPayment) HasPlatformPaymentType() bool {
	if o != nil && !common.IsNil(o.PlatformPaymentType) {
		return true
	}

	return false
}

// SetPlatformPaymentType gets a reference to the given string and assigns it to the PlatformPaymentType field.
func (o *PlatformPayment) SetPlatformPaymentType(v string) {
	o.PlatformPaymentType = &v
}

// GetPspPaymentReference returns the PspPaymentReference field value if set, zero value otherwise.
func (o *PlatformPayment) GetPspPaymentReference() string {
	if o == nil || common.IsNil(o.PspPaymentReference) {
		var ret string
		return ret
	}
	return *o.PspPaymentReference
}

// GetPspPaymentReferenceOk returns a tuple with the PspPaymentReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformPayment) GetPspPaymentReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.PspPaymentReference) {
		return nil, false
	}
	return o.PspPaymentReference, true
}

// HasPspPaymentReference returns a boolean if a field has been set.
func (o *PlatformPayment) HasPspPaymentReference() bool {
	if o != nil && !common.IsNil(o.PspPaymentReference) {
		return true
	}

	return false
}

// SetPspPaymentReference gets a reference to the given string and assigns it to the PspPaymentReference field.
func (o *PlatformPayment) SetPspPaymentReference(v string) {
	o.PspPaymentReference = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PlatformPayment) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformPayment) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PlatformPayment) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PlatformPayment) SetType(v string) {
	o.Type = &v
}

func (o PlatformPayment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlatformPayment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ModificationMerchantReference) {
		toSerialize["modificationMerchantReference"] = o.ModificationMerchantReference
	}
	if !common.IsNil(o.ModificationPspReference) {
		toSerialize["modificationPspReference"] = o.ModificationPspReference
	}
	if !common.IsNil(o.PaymentMerchantReference) {
		toSerialize["paymentMerchantReference"] = o.PaymentMerchantReference
	}
	if !common.IsNil(o.PlatformPaymentType) {
		toSerialize["platformPaymentType"] = o.PlatformPaymentType
	}
	if !common.IsNil(o.PspPaymentReference) {
		toSerialize["pspPaymentReference"] = o.PspPaymentReference
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullablePlatformPayment struct {
	value *PlatformPayment
	isSet bool
}

func (v NullablePlatformPayment) Get() *PlatformPayment {
	return v.value
}

func (v *NullablePlatformPayment) Set(val *PlatformPayment) {
	v.value = val
	v.isSet = true
}

func (v NullablePlatformPayment) IsSet() bool {
	return v.isSet
}

func (v *NullablePlatformPayment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlatformPayment(val *PlatformPayment) *NullablePlatformPayment {
	return &NullablePlatformPayment{value: val, isSet: true}
}

func (v NullablePlatformPayment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlatformPayment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *PlatformPayment) isValidPlatformPaymentType() bool {
	var allowedEnumValues = []string{"AcquiringFees", "AdyenCommission", "AdyenFees", "AdyenMarkup", "BalanceAccount", "Commission", "Default", "Interchange", "PaymentFee", "Remainder", "SchemeFee", "Surcharge", "Tip", "TopUp", "VAT"}
	for _, allowed := range allowedEnumValues {
		if o.GetPlatformPaymentType() == allowed {
			return true
		}
	}
	return false
}
func (o *PlatformPayment) isValidType() bool {
	var allowedEnumValues = []string{"platformPayment"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
