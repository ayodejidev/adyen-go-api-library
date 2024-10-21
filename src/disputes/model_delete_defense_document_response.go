/*
Disputes API

API version: 30
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package disputes

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v14/src/common"
)

// checks if the DeleteDefenseDocumentResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DeleteDefenseDocumentResponse{}

// DeleteDefenseDocumentResponse struct for DeleteDefenseDocumentResponse
type DeleteDefenseDocumentResponse struct {
	DisputeServiceResult DisputeServiceResult `json:"disputeServiceResult"`
}

// NewDeleteDefenseDocumentResponse instantiates a new DeleteDefenseDocumentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteDefenseDocumentResponse(disputeServiceResult DisputeServiceResult) *DeleteDefenseDocumentResponse {
	this := DeleteDefenseDocumentResponse{}
	this.DisputeServiceResult = disputeServiceResult
	return &this
}

// NewDeleteDefenseDocumentResponseWithDefaults instantiates a new DeleteDefenseDocumentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteDefenseDocumentResponseWithDefaults() *DeleteDefenseDocumentResponse {
	this := DeleteDefenseDocumentResponse{}
	return &this
}

// GetDisputeServiceResult returns the DisputeServiceResult field value
func (o *DeleteDefenseDocumentResponse) GetDisputeServiceResult() DisputeServiceResult {
	if o == nil {
		var ret DisputeServiceResult
		return ret
	}

	return o.DisputeServiceResult
}

// GetDisputeServiceResultOk returns a tuple with the DisputeServiceResult field value
// and a boolean to check if the value has been set.
func (o *DeleteDefenseDocumentResponse) GetDisputeServiceResultOk() (*DisputeServiceResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisputeServiceResult, true
}

// SetDisputeServiceResult sets field value
func (o *DeleteDefenseDocumentResponse) SetDisputeServiceResult(v DisputeServiceResult) {
	o.DisputeServiceResult = v
}

func (o DeleteDefenseDocumentResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteDefenseDocumentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["disputeServiceResult"] = o.DisputeServiceResult
	return toSerialize, nil
}

type NullableDeleteDefenseDocumentResponse struct {
	value *DeleteDefenseDocumentResponse
	isSet bool
}

func (v NullableDeleteDefenseDocumentResponse) Get() *DeleteDefenseDocumentResponse {
	return v.value
}

func (v *NullableDeleteDefenseDocumentResponse) Set(val *DeleteDefenseDocumentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteDefenseDocumentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteDefenseDocumentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteDefenseDocumentResponse(val *DeleteDefenseDocumentResponse) *NullableDeleteDefenseDocumentResponse {
	return &NullableDeleteDefenseDocumentResponse{value: val, isSet: true}
}

func (v NullableDeleteDefenseDocumentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteDefenseDocumentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
