/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v12/src/common"
)

// checks if the OrderItem type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OrderItem{}

// OrderItem struct for OrderItem
type OrderItem struct {
	// The unique identifier of the product.
	Id *string `json:"id,omitempty"`
	// The number of installments for the specified product `id`.
	Installments *int64 `json:"installments,omitempty"`
	// The name of the product.
	Name *string `json:"name,omitempty"`
	// The number of items with the specified product `id` included in the order.
	Quantity *int32 `json:"quantity,omitempty"`
}

// NewOrderItem instantiates a new OrderItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderItem() *OrderItem {
	this := OrderItem{}
	return &this
}

// NewOrderItemWithDefaults instantiates a new OrderItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderItemWithDefaults() *OrderItem {
	this := OrderItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrderItem) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItem) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrderItem) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrderItem) SetId(v string) {
	o.Id = &v
}

// GetInstallments returns the Installments field value if set, zero value otherwise.
func (o *OrderItem) GetInstallments() int64 {
	if o == nil || common.IsNil(o.Installments) {
		var ret int64
		return ret
	}
	return *o.Installments
}

// GetInstallmentsOk returns a tuple with the Installments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItem) GetInstallmentsOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Installments) {
		return nil, false
	}
	return o.Installments, true
}

// HasInstallments returns a boolean if a field has been set.
func (o *OrderItem) HasInstallments() bool {
	if o != nil && !common.IsNil(o.Installments) {
		return true
	}

	return false
}

// SetInstallments gets a reference to the given int64 and assigns it to the Installments field.
func (o *OrderItem) SetInstallments(v int64) {
	o.Installments = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrderItem) GetName() string {
	if o == nil || common.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItem) GetNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrderItem) HasName() bool {
	if o != nil && !common.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrderItem) SetName(v string) {
	o.Name = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *OrderItem) GetQuantity() int32 {
	if o == nil || common.IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItem) GetQuantityOk() (*int32, bool) {
	if o == nil || common.IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *OrderItem) HasQuantity() bool {
	if o != nil && !common.IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *OrderItem) SetQuantity(v int32) {
	o.Quantity = &v
}

func (o OrderItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.Installments) {
		toSerialize["installments"] = o.Installments
	}
	if !common.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !common.IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	return toSerialize, nil
}

type NullableOrderItem struct {
	value *OrderItem
	isSet bool
}

func (v NullableOrderItem) Get() *OrderItem {
	return v.value
}

func (v *NullableOrderItem) Set(val *OrderItem) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderItem) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderItem(val *OrderItem) *NullableOrderItem {
	return &NullableOrderItem{value: val, isSet: true}
}

func (v NullableOrderItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
