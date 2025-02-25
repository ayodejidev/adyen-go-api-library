/*
Transfers API

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transfers

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v19/src/common"
)

// checks if the TransactionSearchResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TransactionSearchResponse{}

// TransactionSearchResponse struct for TransactionSearchResponse
type TransactionSearchResponse struct {
	Links *Links `json:"_links,omitempty"`
	// Contains the transactions that match the query parameters.
	Data []Transaction `json:"data,omitempty"`
}

// NewTransactionSearchResponse instantiates a new TransactionSearchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionSearchResponse() *TransactionSearchResponse {
	this := TransactionSearchResponse{}
	return &this
}

// NewTransactionSearchResponseWithDefaults instantiates a new TransactionSearchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionSearchResponseWithDefaults() *TransactionSearchResponse {
	this := TransactionSearchResponse{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *TransactionSearchResponse) GetLinks() Links {
	if o == nil || common.IsNil(o.Links) {
		var ret Links
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionSearchResponse) GetLinksOk() (*Links, bool) {
	if o == nil || common.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *TransactionSearchResponse) HasLinks() bool {
	if o != nil && !common.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given Links and assigns it to the Links field.
func (o *TransactionSearchResponse) SetLinks(v Links) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *TransactionSearchResponse) GetData() []Transaction {
	if o == nil || common.IsNil(o.Data) {
		var ret []Transaction
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionSearchResponse) GetDataOk() ([]Transaction, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *TransactionSearchResponse) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []Transaction and assigns it to the Data field.
func (o *TransactionSearchResponse) SetData(v []Transaction) {
	o.Data = v
}

func (o TransactionSearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionSearchResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !common.IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableTransactionSearchResponse struct {
	value *TransactionSearchResponse
	isSet bool
}

func (v NullableTransactionSearchResponse) Get() *TransactionSearchResponse {
	return v.value
}

func (v *NullableTransactionSearchResponse) Set(val *TransactionSearchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionSearchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionSearchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionSearchResponse(val *TransactionSearchResponse) *NullableTransactionSearchResponse {
	return &NullableTransactionSearchResponse{value: val, isSet: true}
}

func (v NullableTransactionSearchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionSearchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
