/*
Transfers API

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transfers

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v9/src/common"
)

// checks if the Modification type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Modification{}

// Modification struct for Modification
type Modification struct {
	// The direction of the money movement.
	Direction *string `json:"direction,omitempty"`
	// Our reference for the modification.
	Id *string `json:"id,omitempty"`
	// Your reference for the modification, used internally within your platform.
	Reference *string `json:"reference,omitempty"`
	// The status of the transfer event.
	Status *string `json:"status,omitempty"`
	// The type of transfer modification.
	Type *string `json:"type,omitempty"`
}

// NewModification instantiates a new Modification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModification() *Modification {
	this := Modification{}
	return &this
}

// NewModificationWithDefaults instantiates a new Modification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModificationWithDefaults() *Modification {
	this := Modification{}
	return &this
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *Modification) GetDirection() string {
	if o == nil || common.IsNil(o.Direction) {
		var ret string
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Modification) GetDirectionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *Modification) HasDirection() bool {
	if o != nil && !common.IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given string and assigns it to the Direction field.
func (o *Modification) SetDirection(v string) {
	o.Direction = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Modification) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Modification) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Modification) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Modification) SetId(v string) {
	o.Id = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *Modification) GetReference() string {
	if o == nil || common.IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Modification) GetReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *Modification) HasReference() bool {
	if o != nil && !common.IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *Modification) SetReference(v string) {
	o.Reference = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Modification) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Modification) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Modification) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Modification) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Modification) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Modification) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Modification) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Modification) SetType(v string) {
	o.Type = &v
}

func (o Modification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Modification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Direction) {
		toSerialize["direction"] = o.Direction
	}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableModification struct {
	value *Modification
	isSet bool
}

func (v NullableModification) Get() *Modification {
	return v.value
}

func (v *NullableModification) Set(val *Modification) {
	v.value = val
	v.isSet = true
}

func (v NullableModification) IsSet() bool {
	return v.isSet
}

func (v *NullableModification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModification(val *Modification) *NullableModification {
	return &NullableModification{value: val, isSet: true}
}

func (v NullableModification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *Modification) isValidStatus() bool {
	var allowedEnumValues = []string{"approvalPending", "atmWithdrawal", "atmWithdrawalReversalPending", "atmWithdrawalReversed", "authAdjustmentAuthorised", "authAdjustmentError", "authAdjustmentRefused", "authorised", "bankTransfer", "bankTransferPending", "booked", "bookingPending", "cancelled", "capturePending", "captureReversalPending", "captureReversed", "captured", "capturedExternally", "chargeback", "chargebackExternally", "chargebackPending", "chargebackReversalPending", "chargebackReversed", "credited", "depositCorrection", "depositCorrectionPending", "dispute", "disputeClosed", "disputeExpired", "disputeNeedsReview", "error", "expired", "failed", "fee", "feePending", "internalTransfer", "internalTransferPending", "invoiceDeduction", "invoiceDeductionPending", "manualCorrectionPending", "manuallyCorrected", "matchedStatement", "matchedStatementPending", "merchantPayin", "merchantPayinPending", "merchantPayinReversed", "merchantPayinReversedPending", "miscCost", "miscCostPending", "paymentCost", "paymentCostPending", "received", "refundPending", "refundReversalPending", "refundReversed", "refunded", "refundedExternally", "refused", "rejected", "reserveAdjustment", "reserveAdjustmentPending", "returned", "secondChargeback", "secondChargebackPending", "undefined"}
	for _, allowed := range allowedEnumValues {
		if o.GetStatus() == allowed {
			return true
		}
	}
	return false
}
