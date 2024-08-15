/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v12/src/common"
)

// checks if the PaginatedGetCardOrderItemResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PaginatedGetCardOrderItemResponse{}

// PaginatedGetCardOrderItemResponse struct for PaginatedGetCardOrderItemResponse
type PaginatedGetCardOrderItemResponse struct {
	// List of card order items in the card order batch.
	Data []CardOrderItem `json:"data"`
	// Indicates whether there are more items on the next page.
	HasNext bool `json:"hasNext"`
	// Indicates whether there are more items on the previous page.
	HasPrevious bool `json:"hasPrevious"`
}

// NewPaginatedGetCardOrderItemResponse instantiates a new PaginatedGetCardOrderItemResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedGetCardOrderItemResponse(data []CardOrderItem, hasNext bool, hasPrevious bool) *PaginatedGetCardOrderItemResponse {
	this := PaginatedGetCardOrderItemResponse{}
	this.Data = data
	this.HasNext = hasNext
	this.HasPrevious = hasPrevious
	return &this
}

// NewPaginatedGetCardOrderItemResponseWithDefaults instantiates a new PaginatedGetCardOrderItemResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedGetCardOrderItemResponseWithDefaults() *PaginatedGetCardOrderItemResponse {
	this := PaginatedGetCardOrderItemResponse{}
	return &this
}

// GetData returns the Data field value
func (o *PaginatedGetCardOrderItemResponse) GetData() []CardOrderItem {
	if o == nil {
		var ret []CardOrderItem
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PaginatedGetCardOrderItemResponse) GetDataOk() ([]CardOrderItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *PaginatedGetCardOrderItemResponse) SetData(v []CardOrderItem) {
	o.Data = v
}

// GetHasNext returns the HasNext field value
func (o *PaginatedGetCardOrderItemResponse) GetHasNext() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasNext
}

// GetHasNextOk returns a tuple with the HasNext field value
// and a boolean to check if the value has been set.
func (o *PaginatedGetCardOrderItemResponse) GetHasNextOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasNext, true
}

// SetHasNext sets field value
func (o *PaginatedGetCardOrderItemResponse) SetHasNext(v bool) {
	o.HasNext = v
}

// GetHasPrevious returns the HasPrevious field value
func (o *PaginatedGetCardOrderItemResponse) GetHasPrevious() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasPrevious
}

// GetHasPreviousOk returns a tuple with the HasPrevious field value
// and a boolean to check if the value has been set.
func (o *PaginatedGetCardOrderItemResponse) GetHasPreviousOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasPrevious, true
}

// SetHasPrevious sets field value
func (o *PaginatedGetCardOrderItemResponse) SetHasPrevious(v bool) {
	o.HasPrevious = v
}

func (o PaginatedGetCardOrderItemResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedGetCardOrderItemResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["hasNext"] = o.HasNext
	toSerialize["hasPrevious"] = o.HasPrevious
	return toSerialize, nil
}

type NullablePaginatedGetCardOrderItemResponse struct {
	value *PaginatedGetCardOrderItemResponse
	isSet bool
}

func (v NullablePaginatedGetCardOrderItemResponse) Get() *PaginatedGetCardOrderItemResponse {
	return v.value
}

func (v *NullablePaginatedGetCardOrderItemResponse) Set(val *PaginatedGetCardOrderItemResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedGetCardOrderItemResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedGetCardOrderItemResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedGetCardOrderItemResponse(val *PaginatedGetCardOrderItemResponse) *NullablePaginatedGetCardOrderItemResponse {
	return &NullablePaginatedGetCardOrderItemResponse{value: val, isSet: true}
}

func (v NullablePaginatedGetCardOrderItemResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedGetCardOrderItemResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
