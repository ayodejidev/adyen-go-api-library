/*
Transfer webhooks

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transferwebhook

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// checks if the AmountAdjustment type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AmountAdjustment{}

// AmountAdjustment struct for AmountAdjustment
type AmountAdjustment struct {
	Amount *Amount `json:"amount,omitempty"`
	// The type of markup that is applied to an authorised payment.  Possible values: **exchange**, **forexMarkup**, **authHoldReserve**, **atmMarkup**.
	AmountAdjustmentType *string `json:"amountAdjustmentType,omitempty"`
	// The basepoints associated with the applied markup.
	Basepoints *int32 `json:"basepoints,omitempty"`
}

// NewAmountAdjustment instantiates a new AmountAdjustment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmountAdjustment() *AmountAdjustment {
	this := AmountAdjustment{}
	return &this
}

// NewAmountAdjustmentWithDefaults instantiates a new AmountAdjustment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmountAdjustmentWithDefaults() *AmountAdjustment {
	this := AmountAdjustment{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *AmountAdjustment) GetAmount() Amount {
	if o == nil || common.IsNil(o.Amount) {
		var ret Amount
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmountAdjustment) GetAmountOk() (*Amount, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *AmountAdjustment) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given Amount and assigns it to the Amount field.
func (o *AmountAdjustment) SetAmount(v Amount) {
	o.Amount = &v
}

// GetAmountAdjustmentType returns the AmountAdjustmentType field value if set, zero value otherwise.
func (o *AmountAdjustment) GetAmountAdjustmentType() string {
	if o == nil || common.IsNil(o.AmountAdjustmentType) {
		var ret string
		return ret
	}
	return *o.AmountAdjustmentType
}

// GetAmountAdjustmentTypeOk returns a tuple with the AmountAdjustmentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmountAdjustment) GetAmountAdjustmentTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.AmountAdjustmentType) {
		return nil, false
	}
	return o.AmountAdjustmentType, true
}

// HasAmountAdjustmentType returns a boolean if a field has been set.
func (o *AmountAdjustment) HasAmountAdjustmentType() bool {
	if o != nil && !common.IsNil(o.AmountAdjustmentType) {
		return true
	}

	return false
}

// SetAmountAdjustmentType gets a reference to the given string and assigns it to the AmountAdjustmentType field.
func (o *AmountAdjustment) SetAmountAdjustmentType(v string) {
	o.AmountAdjustmentType = &v
}

// GetBasepoints returns the Basepoints field value if set, zero value otherwise.
func (o *AmountAdjustment) GetBasepoints() int32 {
	if o == nil || common.IsNil(o.Basepoints) {
		var ret int32
		return ret
	}
	return *o.Basepoints
}

// GetBasepointsOk returns a tuple with the Basepoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmountAdjustment) GetBasepointsOk() (*int32, bool) {
	if o == nil || common.IsNil(o.Basepoints) {
		return nil, false
	}
	return o.Basepoints, true
}

// HasBasepoints returns a boolean if a field has been set.
func (o *AmountAdjustment) HasBasepoints() bool {
	if o != nil && !common.IsNil(o.Basepoints) {
		return true
	}

	return false
}

// SetBasepoints gets a reference to the given int32 and assigns it to the Basepoints field.
func (o *AmountAdjustment) SetBasepoints(v int32) {
	o.Basepoints = &v
}

func (o AmountAdjustment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AmountAdjustment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.AmountAdjustmentType) {
		toSerialize["amountAdjustmentType"] = o.AmountAdjustmentType
	}
	if !common.IsNil(o.Basepoints) {
		toSerialize["basepoints"] = o.Basepoints
	}
	return toSerialize, nil
}

type NullableAmountAdjustment struct {
	value *AmountAdjustment
	isSet bool
}

func (v NullableAmountAdjustment) Get() *AmountAdjustment {
	return v.value
}

func (v *NullableAmountAdjustment) Set(val *AmountAdjustment) {
	v.value = val
	v.isSet = true
}

func (v NullableAmountAdjustment) IsSet() bool {
	return v.isSet
}

func (v *NullableAmountAdjustment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmountAdjustment(val *AmountAdjustment) *NullableAmountAdjustment {
	return &NullableAmountAdjustment{value: val, isSet: true}
}

func (v NullableAmountAdjustment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmountAdjustment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *AmountAdjustment) isValidAmountAdjustmentType() bool {
	var allowedEnumValues = []string{"atmMarkup", "authHoldReserve", "exchange", "forexMarkup"}
	for _, allowed := range allowedEnumValues {
		if o.GetAmountAdjustmentType() == allowed {
			return true
		}
	}
	return false
}
