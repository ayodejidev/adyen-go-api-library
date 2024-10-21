/*
Transfers API

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transfers

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v14/src/common"
)

// checks if the TransferReview type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TransferReview{}

// TransferReview struct for TransferReview
type TransferReview struct {
	// Shows the number of approvals completed for the transfer.
	NumberOfApprovalsCompleted *int32 `json:"numberOfApprovalsCompleted,omitempty"`
	// Shows the number of [approvals](https://docs.adyen.com/api-explorer/transfers/latest/post/transfers/approve) required to process the transfer.
	NumberOfApprovalsRequired *int32 `json:"numberOfApprovalsRequired,omitempty"`
	// Shows the status of the Strong Customer Authentication (SCA) process.  Possible values: **required**, **completed**, **notApplicable**.
	ScaOnApproval *string `json:"scaOnApproval,omitempty"`
}

// NewTransferReview instantiates a new TransferReview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferReview() *TransferReview {
	this := TransferReview{}
	return &this
}

// NewTransferReviewWithDefaults instantiates a new TransferReview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferReviewWithDefaults() *TransferReview {
	this := TransferReview{}
	return &this
}

// GetNumberOfApprovalsCompleted returns the NumberOfApprovalsCompleted field value if set, zero value otherwise.
func (o *TransferReview) GetNumberOfApprovalsCompleted() int32 {
	if o == nil || common.IsNil(o.NumberOfApprovalsCompleted) {
		var ret int32
		return ret
	}
	return *o.NumberOfApprovalsCompleted
}

// GetNumberOfApprovalsCompletedOk returns a tuple with the NumberOfApprovalsCompleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferReview) GetNumberOfApprovalsCompletedOk() (*int32, bool) {
	if o == nil || common.IsNil(o.NumberOfApprovalsCompleted) {
		return nil, false
	}
	return o.NumberOfApprovalsCompleted, true
}

// HasNumberOfApprovalsCompleted returns a boolean if a field has been set.
func (o *TransferReview) HasNumberOfApprovalsCompleted() bool {
	if o != nil && !common.IsNil(o.NumberOfApprovalsCompleted) {
		return true
	}

	return false
}

// SetNumberOfApprovalsCompleted gets a reference to the given int32 and assigns it to the NumberOfApprovalsCompleted field.
func (o *TransferReview) SetNumberOfApprovalsCompleted(v int32) {
	o.NumberOfApprovalsCompleted = &v
}

// GetNumberOfApprovalsRequired returns the NumberOfApprovalsRequired field value if set, zero value otherwise.
func (o *TransferReview) GetNumberOfApprovalsRequired() int32 {
	if o == nil || common.IsNil(o.NumberOfApprovalsRequired) {
		var ret int32
		return ret
	}
	return *o.NumberOfApprovalsRequired
}

// GetNumberOfApprovalsRequiredOk returns a tuple with the NumberOfApprovalsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferReview) GetNumberOfApprovalsRequiredOk() (*int32, bool) {
	if o == nil || common.IsNil(o.NumberOfApprovalsRequired) {
		return nil, false
	}
	return o.NumberOfApprovalsRequired, true
}

// HasNumberOfApprovalsRequired returns a boolean if a field has been set.
func (o *TransferReview) HasNumberOfApprovalsRequired() bool {
	if o != nil && !common.IsNil(o.NumberOfApprovalsRequired) {
		return true
	}

	return false
}

// SetNumberOfApprovalsRequired gets a reference to the given int32 and assigns it to the NumberOfApprovalsRequired field.
func (o *TransferReview) SetNumberOfApprovalsRequired(v int32) {
	o.NumberOfApprovalsRequired = &v
}

// GetScaOnApproval returns the ScaOnApproval field value if set, zero value otherwise.
func (o *TransferReview) GetScaOnApproval() string {
	if o == nil || common.IsNil(o.ScaOnApproval) {
		var ret string
		return ret
	}
	return *o.ScaOnApproval
}

// GetScaOnApprovalOk returns a tuple with the ScaOnApproval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferReview) GetScaOnApprovalOk() (*string, bool) {
	if o == nil || common.IsNil(o.ScaOnApproval) {
		return nil, false
	}
	return o.ScaOnApproval, true
}

// HasScaOnApproval returns a boolean if a field has been set.
func (o *TransferReview) HasScaOnApproval() bool {
	if o != nil && !common.IsNil(o.ScaOnApproval) {
		return true
	}

	return false
}

// SetScaOnApproval gets a reference to the given string and assigns it to the ScaOnApproval field.
func (o *TransferReview) SetScaOnApproval(v string) {
	o.ScaOnApproval = &v
}

func (o TransferReview) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferReview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.NumberOfApprovalsCompleted) {
		toSerialize["numberOfApprovalsCompleted"] = o.NumberOfApprovalsCompleted
	}
	if !common.IsNil(o.NumberOfApprovalsRequired) {
		toSerialize["numberOfApprovalsRequired"] = o.NumberOfApprovalsRequired
	}
	if !common.IsNil(o.ScaOnApproval) {
		toSerialize["scaOnApproval"] = o.ScaOnApproval
	}
	return toSerialize, nil
}

type NullableTransferReview struct {
	value *TransferReview
	isSet bool
}

func (v NullableTransferReview) Get() *TransferReview {
	return v.value
}

func (v *NullableTransferReview) Set(val *TransferReview) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferReview) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferReview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferReview(val *TransferReview) *NullableTransferReview {
	return &NullableTransferReview{value: val, isSet: true}
}

func (v NullableTransferReview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferReview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *TransferReview) isValidScaOnApproval() bool {
	var allowedEnumValues = []string{"completed", "notApplicable", "required"}
	for _, allowed := range allowedEnumValues {
		if o.GetScaOnApproval() == allowed {
			return true
		}
	}
	return false
}
