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

// checks if the ApproveTransfersRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ApproveTransfersRequest{}

// ApproveTransfersRequest struct for ApproveTransfersRequest
type ApproveTransfersRequest struct {
	// Contains the unique identifiers of the transfers that you want to approve.
	TransferIds []string `json:"transferIds,omitempty"`
}

// NewApproveTransfersRequest instantiates a new ApproveTransfersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApproveTransfersRequest() *ApproveTransfersRequest {
	this := ApproveTransfersRequest{}
	return &this
}

// NewApproveTransfersRequestWithDefaults instantiates a new ApproveTransfersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApproveTransfersRequestWithDefaults() *ApproveTransfersRequest {
	this := ApproveTransfersRequest{}
	return &this
}

// GetTransferIds returns the TransferIds field value if set, zero value otherwise.
func (o *ApproveTransfersRequest) GetTransferIds() []string {
	if o == nil || common.IsNil(o.TransferIds) {
		var ret []string
		return ret
	}
	return o.TransferIds
}

// GetTransferIdsOk returns a tuple with the TransferIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApproveTransfersRequest) GetTransferIdsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.TransferIds) {
		return nil, false
	}
	return o.TransferIds, true
}

// HasTransferIds returns a boolean if a field has been set.
func (o *ApproveTransfersRequest) HasTransferIds() bool {
	if o != nil && !common.IsNil(o.TransferIds) {
		return true
	}

	return false
}

// SetTransferIds gets a reference to the given []string and assigns it to the TransferIds field.
func (o *ApproveTransfersRequest) SetTransferIds(v []string) {
	o.TransferIds = v
}

func (o ApproveTransfersRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApproveTransfersRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.TransferIds) {
		toSerialize["transferIds"] = o.TransferIds
	}
	return toSerialize, nil
}

type NullableApproveTransfersRequest struct {
	value *ApproveTransfersRequest
	isSet bool
}

func (v NullableApproveTransfersRequest) Get() *ApproveTransfersRequest {
	return v.value
}

func (v *NullableApproveTransfersRequest) Set(val *ApproveTransfersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApproveTransfersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApproveTransfersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApproveTransfersRequest(val *ApproveTransfersRequest) *NullableApproveTransfersRequest {
	return &NullableApproveTransfersRequest{value: val, isSet: true}
}

func (v NullableApproveTransfersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApproveTransfersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
