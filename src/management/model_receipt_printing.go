/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v19/src/common"
)

// checks if the ReceiptPrinting type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ReceiptPrinting{}

// ReceiptPrinting struct for ReceiptPrinting
type ReceiptPrinting struct {
	// Print a merchant receipt when the payment is approved.
	MerchantApproved *bool `json:"merchantApproved,omitempty"`
	// Print a merchant receipt when the transaction is cancelled.
	MerchantCancelled *bool `json:"merchantCancelled,omitempty"`
	// Print a merchant receipt when capturing the payment is approved.
	MerchantCaptureApproved *bool `json:"merchantCaptureApproved,omitempty"`
	// Print a merchant receipt when capturing the payment is refused.
	MerchantCaptureRefused *bool `json:"merchantCaptureRefused,omitempty"`
	// Print a merchant receipt when the refund is approved.
	MerchantRefundApproved *bool `json:"merchantRefundApproved,omitempty"`
	// Print a merchant receipt when the refund is refused.
	MerchantRefundRefused *bool `json:"merchantRefundRefused,omitempty"`
	// Print a merchant receipt when the payment is refused.
	MerchantRefused *bool `json:"merchantRefused,omitempty"`
	// Print a merchant receipt when a previous transaction is voided.
	MerchantVoid *bool `json:"merchantVoid,omitempty"`
	// Print a shopper receipt when the payment is approved.
	ShopperApproved *bool `json:"shopperApproved,omitempty"`
	// Print a shopper receipt when the transaction is cancelled.
	ShopperCancelled *bool `json:"shopperCancelled,omitempty"`
	// Print a shopper receipt when capturing the payment is approved.
	ShopperCaptureApproved *bool `json:"shopperCaptureApproved,omitempty"`
	// Print a shopper receipt when capturing the payment is refused.
	ShopperCaptureRefused *bool `json:"shopperCaptureRefused,omitempty"`
	// Print a shopper receipt when the refund is approved.
	ShopperRefundApproved *bool `json:"shopperRefundApproved,omitempty"`
	// Print a shopper receipt when the refund is refused.
	ShopperRefundRefused *bool `json:"shopperRefundRefused,omitempty"`
	// Print a shopper receipt when the payment is refused.
	ShopperRefused *bool `json:"shopperRefused,omitempty"`
	// Print a shopper receipt when a previous transaction is voided.
	ShopperVoid *bool `json:"shopperVoid,omitempty"`
}

// NewReceiptPrinting instantiates a new ReceiptPrinting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReceiptPrinting() *ReceiptPrinting {
	this := ReceiptPrinting{}
	return &this
}

// NewReceiptPrintingWithDefaults instantiates a new ReceiptPrinting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReceiptPrintingWithDefaults() *ReceiptPrinting {
	this := ReceiptPrinting{}
	return &this
}

// GetMerchantApproved returns the MerchantApproved field value if set, zero value otherwise.
func (o *ReceiptPrinting) GetMerchantApproved() bool {
	if o == nil || common.IsNil(o.MerchantApproved) {
		var ret bool
		return ret
	}
	return *o.MerchantApproved
}

// GetMerchantApprovedOk returns a tuple with the MerchantApproved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReceiptPrinting) GetMerchantApprovedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.MerchantApproved) {
		return nil, false
	}
	return o.MerchantApproved, true
}

// HasMerchantApproved returns a boolean if a field has been set.
func (o *ReceiptPrinting) HasMerchantApproved() bool {
	if o != nil && !common.IsNil(o.MerchantApproved) {
		return true
	}

	return false
}

// SetMerchantApproved gets a reference to the given bool and assigns it to the MerchantApproved field.
func (o *ReceiptPrinting) SetMerchantApproved(v bool) {
	o.MerchantApproved = &v
}

// GetMerchantCancelled returns the MerchantCancelled field value if set, zero value otherwise.
func (o *ReceiptPrinting) GetMerchantCancelled() bool {
	if o == nil || common.IsNil(o.MerchantCancelled) {
		var ret bool
		return ret
	}
	return *o.MerchantCancelled
}

// GetMerchantCancelledOk returns a tuple with the MerchantCancelled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReceiptPrinting) GetMerchantCancelledOk() (*bool, bool) {
	if o == nil || common.IsNil(o.MerchantCancelled) {
		return nil, false
	}
	return o.MerchantCancelled, true
}

// HasMerchantCancelled returns a boolean if a field has been set.
func (o *ReceiptPrinting) HasMerchantCancelled() bool {
	if o != nil && !common.IsNil(o.MerchantCancelled) {
		return true
	}

	return false
}

// SetMerchantCancelled gets a reference to the given bool and assigns it to the MerchantCancelled field.
func (o *ReceiptPrinting) SetMerchantCancelled(v bool) {
	o.MerchantCancelled = &v
}

// GetMerchantCaptureApproved returns the MerchantCaptureApproved field value if set, zero value otherwise.
func (o *ReceiptPrinting) GetMerchantCaptureApproved() bool {
	if o == nil || common.IsNil(o.MerchantCaptureApproved) {
		var ret bool
		return ret
	}
	return *o.MerchantCaptureApproved
}

// GetMerchantCaptureApprovedOk returns a tuple with the MerchantCaptureApproved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReceiptPrinting) GetMerchantCaptureApprovedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.MerchantCaptureApproved) {
		return nil, false
	}
	return o.MerchantCaptureApproved, true
}

// HasMerchantCaptureApproved returns a boolean if a field has been set.
func (o *ReceiptPrinting) HasMerchantCaptureApproved() bool {
	if o != nil && !common.IsNil(o.MerchantCaptureApproved) {
		return true
	}

	return false
}

// SetMerchantCaptureApproved gets a reference to the given bool and assigns it to the MerchantCaptureApproved field.
func (o *ReceiptPrinting) SetMerchantCaptureApproved(v bool) {
	o.MerchantCaptureApproved = &v
}

// GetMerchantCaptureRefused returns the MerchantCaptureRefused field value if set, zero value otherwise.
func (o *ReceiptPrinting) GetMerchantCaptureRefused() bool {
	if o == nil || common.IsNil(o.MerchantCaptureRefused) {
		var ret bool
		return ret
	}
	return *o.MerchantCaptureRefused
}

// GetMerchantCaptureRefusedOk returns a tuple with the MerchantCaptureRefused field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReceiptPrinting) GetMerchantCaptureRefusedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.MerchantCaptureRefused) {
		return nil, false
	}
	return o.MerchantCaptureRefused, true
}

// HasMerchantCaptureRefused returns a boolean if a field has been set.
func (o *ReceiptPrinting) HasMerchantCaptureRefused() bool {
	if o != nil && !common.IsNil(o.MerchantCaptureRefused) {
		return true
	}

	return false
}

// SetMerchantCaptureRefused gets a reference to the given bool and assigns it to the MerchantCaptureRefused field.
func (o *ReceiptPrinting) SetMerchantCaptureRefused(v bool) {
	o.MerchantCaptureRefused = &v
}

// GetMerchantRefundApproved returns the MerchantRefundApproved field value if set, zero value otherwise.
func (o *ReceiptPrinting) GetMerchantRefundApproved() bool {
	if o == nil || common.IsNil(o.MerchantRefundApproved) {
		var ret bool
		return ret
	}
	return *o.MerchantRefundApproved
}

// GetMerchantRefundApprovedOk returns a tuple with the MerchantRefundApproved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReceiptPrinting) GetMerchantRefundApprovedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.MerchantRefundApproved) {
		return nil, false
	}
	return o.MerchantRefundApproved, true
}

// HasMerchantRefundApproved returns a boolean if a field has been set.
func (o *ReceiptPrinting) HasMerchantRefundApproved() bool {
	if o != nil && !common.IsNil(o.MerchantRefundApproved) {
		return true
	}

	return false
}

// SetMerchantRefundApproved gets a reference to the given bool and assigns it to the MerchantRefundApproved field.
func (o *ReceiptPrinting) SetMerchantRefundApproved(v bool) {
	o.MerchantRefundApproved = &v
}

// GetMerchantRefundRefused returns the MerchantRefundRefused field value if set, zero value otherwise.
func (o *ReceiptPrinting) GetMerchantRefundRefused() bool {
	if o == nil || common.IsNil(o.MerchantRefundRefused) {
		var ret bool
		return ret
	}
	return *o.MerchantRefundRefused
}

// GetMerchantRefundRefusedOk returns a tuple with the MerchantRefundRefused field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReceiptPrinting) GetMerchantRefundRefusedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.MerchantRefundRefused) {
		return nil, false
	}
	return o.MerchantRefundRefused, true
}

// HasMerchantRefundRefused returns a boolean if a field has been set.
func (o *ReceiptPrinting) HasMerchantRefundRefused() bool {
	if o != nil && !common.IsNil(o.MerchantRefundRefused) {
		return true
	}

	return false
}

// SetMerchantRefundRefused gets a reference to the given bool and assigns it to the MerchantRefundRefused field.
func (o *ReceiptPrinting) SetMerchantRefundRefused(v bool) {
	o.MerchantRefundRefused = &v
}

// GetMerchantRefused returns the MerchantRefused field value if set, zero value otherwise.
func (o *ReceiptPrinting) GetMerchantRefused() bool {
	if o == nil || common.IsNil(o.MerchantRefused) {
		var ret bool
		return ret
	}
	return *o.MerchantRefused
}

// GetMerchantRefusedOk returns a tuple with the MerchantRefused field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReceiptPrinting) GetMerchantRefusedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.MerchantRefused) {
		return nil, false
	}
	return o.MerchantRefused, true
}

// HasMerchantRefused returns a boolean if a field has been set.
func (o *ReceiptPrinting) HasMerchantRefused() bool {
	if o != nil && !common.IsNil(o.MerchantRefused) {
		return true
	}

	return false
}

// SetMerchantRefused gets a reference to the given bool and assigns it to the MerchantRefused field.
func (o *ReceiptPrinting) SetMerchantRefused(v bool) {
	o.MerchantRefused = &v
}

// GetMerchantVoid returns the MerchantVoid field value if set, zero value otherwise.
func (o *ReceiptPrinting) GetMerchantVoid() bool {
	if o == nil || common.IsNil(o.MerchantVoid) {
		var ret bool
		return ret
	}
	return *o.MerchantVoid
}

// GetMerchantVoidOk returns a tuple with the MerchantVoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReceiptPrinting) GetMerchantVoidOk() (*bool, bool) {
	if o == nil || common.IsNil(o.MerchantVoid) {
		return nil, false
	}
	return o.MerchantVoid, true
}

// HasMerchantVoid returns a boolean if a field has been set.
func (o *ReceiptPrinting) HasMerchantVoid() bool {
	if o != nil && !common.IsNil(o.MerchantVoid) {
		return true
	}

	return false
}

// SetMerchantVoid gets a reference to the given bool and assigns it to the MerchantVoid field.
func (o *ReceiptPrinting) SetMerchantVoid(v bool) {
	o.MerchantVoid = &v
}

// GetShopperApproved returns the ShopperApproved field value if set, zero value otherwise.
func (o *ReceiptPrinting) GetShopperApproved() bool {
	if o == nil || common.IsNil(o.ShopperApproved) {
		var ret bool
		return ret
	}
	return *o.ShopperApproved
}

// GetShopperApprovedOk returns a tuple with the ShopperApproved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReceiptPrinting) GetShopperApprovedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.ShopperApproved) {
		return nil, false
	}
	return o.ShopperApproved, true
}

// HasShopperApproved returns a boolean if a field has been set.
func (o *ReceiptPrinting) HasShopperApproved() bool {
	if o != nil && !common.IsNil(o.ShopperApproved) {
		return true
	}

	return false
}

// SetShopperApproved gets a reference to the given bool and assigns it to the ShopperApproved field.
func (o *ReceiptPrinting) SetShopperApproved(v bool) {
	o.ShopperApproved = &v
}

// GetShopperCancelled returns the ShopperCancelled field value if set, zero value otherwise.
func (o *ReceiptPrinting) GetShopperCancelled() bool {
	if o == nil || common.IsNil(o.ShopperCancelled) {
		var ret bool
		return ret
	}
	return *o.ShopperCancelled
}

// GetShopperCancelledOk returns a tuple with the ShopperCancelled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReceiptPrinting) GetShopperCancelledOk() (*bool, bool) {
	if o == nil || common.IsNil(o.ShopperCancelled) {
		return nil, false
	}
	return o.ShopperCancelled, true
}

// HasShopperCancelled returns a boolean if a field has been set.
func (o *ReceiptPrinting) HasShopperCancelled() bool {
	if o != nil && !common.IsNil(o.ShopperCancelled) {
		return true
	}

	return false
}

// SetShopperCancelled gets a reference to the given bool and assigns it to the ShopperCancelled field.
func (o *ReceiptPrinting) SetShopperCancelled(v bool) {
	o.ShopperCancelled = &v
}

// GetShopperCaptureApproved returns the ShopperCaptureApproved field value if set, zero value otherwise.
func (o *ReceiptPrinting) GetShopperCaptureApproved() bool {
	if o == nil || common.IsNil(o.ShopperCaptureApproved) {
		var ret bool
		return ret
	}
	return *o.ShopperCaptureApproved
}

// GetShopperCaptureApprovedOk returns a tuple with the ShopperCaptureApproved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReceiptPrinting) GetShopperCaptureApprovedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.ShopperCaptureApproved) {
		return nil, false
	}
	return o.ShopperCaptureApproved, true
}

// HasShopperCaptureApproved returns a boolean if a field has been set.
func (o *ReceiptPrinting) HasShopperCaptureApproved() bool {
	if o != nil && !common.IsNil(o.ShopperCaptureApproved) {
		return true
	}

	return false
}

// SetShopperCaptureApproved gets a reference to the given bool and assigns it to the ShopperCaptureApproved field.
func (o *ReceiptPrinting) SetShopperCaptureApproved(v bool) {
	o.ShopperCaptureApproved = &v
}

// GetShopperCaptureRefused returns the ShopperCaptureRefused field value if set, zero value otherwise.
func (o *ReceiptPrinting) GetShopperCaptureRefused() bool {
	if o == nil || common.IsNil(o.ShopperCaptureRefused) {
		var ret bool
		return ret
	}
	return *o.ShopperCaptureRefused
}

// GetShopperCaptureRefusedOk returns a tuple with the ShopperCaptureRefused field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReceiptPrinting) GetShopperCaptureRefusedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.ShopperCaptureRefused) {
		return nil, false
	}
	return o.ShopperCaptureRefused, true
}

// HasShopperCaptureRefused returns a boolean if a field has been set.
func (o *ReceiptPrinting) HasShopperCaptureRefused() bool {
	if o != nil && !common.IsNil(o.ShopperCaptureRefused) {
		return true
	}

	return false
}

// SetShopperCaptureRefused gets a reference to the given bool and assigns it to the ShopperCaptureRefused field.
func (o *ReceiptPrinting) SetShopperCaptureRefused(v bool) {
	o.ShopperCaptureRefused = &v
}

// GetShopperRefundApproved returns the ShopperRefundApproved field value if set, zero value otherwise.
func (o *ReceiptPrinting) GetShopperRefundApproved() bool {
	if o == nil || common.IsNil(o.ShopperRefundApproved) {
		var ret bool
		return ret
	}
	return *o.ShopperRefundApproved
}

// GetShopperRefundApprovedOk returns a tuple with the ShopperRefundApproved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReceiptPrinting) GetShopperRefundApprovedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.ShopperRefundApproved) {
		return nil, false
	}
	return o.ShopperRefundApproved, true
}

// HasShopperRefundApproved returns a boolean if a field has been set.
func (o *ReceiptPrinting) HasShopperRefundApproved() bool {
	if o != nil && !common.IsNil(o.ShopperRefundApproved) {
		return true
	}

	return false
}

// SetShopperRefundApproved gets a reference to the given bool and assigns it to the ShopperRefundApproved field.
func (o *ReceiptPrinting) SetShopperRefundApproved(v bool) {
	o.ShopperRefundApproved = &v
}

// GetShopperRefundRefused returns the ShopperRefundRefused field value if set, zero value otherwise.
func (o *ReceiptPrinting) GetShopperRefundRefused() bool {
	if o == nil || common.IsNil(o.ShopperRefundRefused) {
		var ret bool
		return ret
	}
	return *o.ShopperRefundRefused
}

// GetShopperRefundRefusedOk returns a tuple with the ShopperRefundRefused field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReceiptPrinting) GetShopperRefundRefusedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.ShopperRefundRefused) {
		return nil, false
	}
	return o.ShopperRefundRefused, true
}

// HasShopperRefundRefused returns a boolean if a field has been set.
func (o *ReceiptPrinting) HasShopperRefundRefused() bool {
	if o != nil && !common.IsNil(o.ShopperRefundRefused) {
		return true
	}

	return false
}

// SetShopperRefundRefused gets a reference to the given bool and assigns it to the ShopperRefundRefused field.
func (o *ReceiptPrinting) SetShopperRefundRefused(v bool) {
	o.ShopperRefundRefused = &v
}

// GetShopperRefused returns the ShopperRefused field value if set, zero value otherwise.
func (o *ReceiptPrinting) GetShopperRefused() bool {
	if o == nil || common.IsNil(o.ShopperRefused) {
		var ret bool
		return ret
	}
	return *o.ShopperRefused
}

// GetShopperRefusedOk returns a tuple with the ShopperRefused field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReceiptPrinting) GetShopperRefusedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.ShopperRefused) {
		return nil, false
	}
	return o.ShopperRefused, true
}

// HasShopperRefused returns a boolean if a field has been set.
func (o *ReceiptPrinting) HasShopperRefused() bool {
	if o != nil && !common.IsNil(o.ShopperRefused) {
		return true
	}

	return false
}

// SetShopperRefused gets a reference to the given bool and assigns it to the ShopperRefused field.
func (o *ReceiptPrinting) SetShopperRefused(v bool) {
	o.ShopperRefused = &v
}

// GetShopperVoid returns the ShopperVoid field value if set, zero value otherwise.
func (o *ReceiptPrinting) GetShopperVoid() bool {
	if o == nil || common.IsNil(o.ShopperVoid) {
		var ret bool
		return ret
	}
	return *o.ShopperVoid
}

// GetShopperVoidOk returns a tuple with the ShopperVoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReceiptPrinting) GetShopperVoidOk() (*bool, bool) {
	if o == nil || common.IsNil(o.ShopperVoid) {
		return nil, false
	}
	return o.ShopperVoid, true
}

// HasShopperVoid returns a boolean if a field has been set.
func (o *ReceiptPrinting) HasShopperVoid() bool {
	if o != nil && !common.IsNil(o.ShopperVoid) {
		return true
	}

	return false
}

// SetShopperVoid gets a reference to the given bool and assigns it to the ShopperVoid field.
func (o *ReceiptPrinting) SetShopperVoid(v bool) {
	o.ShopperVoid = &v
}

func (o ReceiptPrinting) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReceiptPrinting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.MerchantApproved) {
		toSerialize["merchantApproved"] = o.MerchantApproved
	}
	if !common.IsNil(o.MerchantCancelled) {
		toSerialize["merchantCancelled"] = o.MerchantCancelled
	}
	if !common.IsNil(o.MerchantCaptureApproved) {
		toSerialize["merchantCaptureApproved"] = o.MerchantCaptureApproved
	}
	if !common.IsNil(o.MerchantCaptureRefused) {
		toSerialize["merchantCaptureRefused"] = o.MerchantCaptureRefused
	}
	if !common.IsNil(o.MerchantRefundApproved) {
		toSerialize["merchantRefundApproved"] = o.MerchantRefundApproved
	}
	if !common.IsNil(o.MerchantRefundRefused) {
		toSerialize["merchantRefundRefused"] = o.MerchantRefundRefused
	}
	if !common.IsNil(o.MerchantRefused) {
		toSerialize["merchantRefused"] = o.MerchantRefused
	}
	if !common.IsNil(o.MerchantVoid) {
		toSerialize["merchantVoid"] = o.MerchantVoid
	}
	if !common.IsNil(o.ShopperApproved) {
		toSerialize["shopperApproved"] = o.ShopperApproved
	}
	if !common.IsNil(o.ShopperCancelled) {
		toSerialize["shopperCancelled"] = o.ShopperCancelled
	}
	if !common.IsNil(o.ShopperCaptureApproved) {
		toSerialize["shopperCaptureApproved"] = o.ShopperCaptureApproved
	}
	if !common.IsNil(o.ShopperCaptureRefused) {
		toSerialize["shopperCaptureRefused"] = o.ShopperCaptureRefused
	}
	if !common.IsNil(o.ShopperRefundApproved) {
		toSerialize["shopperRefundApproved"] = o.ShopperRefundApproved
	}
	if !common.IsNil(o.ShopperRefundRefused) {
		toSerialize["shopperRefundRefused"] = o.ShopperRefundRefused
	}
	if !common.IsNil(o.ShopperRefused) {
		toSerialize["shopperRefused"] = o.ShopperRefused
	}
	if !common.IsNil(o.ShopperVoid) {
		toSerialize["shopperVoid"] = o.ShopperVoid
	}
	return toSerialize, nil
}

type NullableReceiptPrinting struct {
	value *ReceiptPrinting
	isSet bool
}

func (v NullableReceiptPrinting) Get() *ReceiptPrinting {
	return v.value
}

func (v *NullableReceiptPrinting) Set(val *ReceiptPrinting) {
	v.value = val
	v.isSet = true
}

func (v NullableReceiptPrinting) IsSet() bool {
	return v.isSet
}

func (v *NullableReceiptPrinting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReceiptPrinting(val *ReceiptPrinting) *NullableReceiptPrinting {
	return &NullableReceiptPrinting{value: val, isSet: true}
}

func (v NullableReceiptPrinting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReceiptPrinting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
