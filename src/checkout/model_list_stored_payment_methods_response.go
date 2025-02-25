/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v19/src/common"
)

// checks if the ListStoredPaymentMethodsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ListStoredPaymentMethodsResponse{}

// ListStoredPaymentMethodsResponse struct for ListStoredPaymentMethodsResponse
type ListStoredPaymentMethodsResponse struct {
	// Your merchant account.
	MerchantAccount *string `json:"merchantAccount,omitempty"`
	// Your reference to uniquely identify this shopper, for example user ID or account ID. Minimum length: 3 characters. > Your reference must not include personally identifiable information (PII), for example name or email address.
	ShopperReference *string `json:"shopperReference,omitempty"`
	// List of all stored payment methods.
	StoredPaymentMethods []StoredPaymentMethodResource `json:"storedPaymentMethods,omitempty"`
}

// NewListStoredPaymentMethodsResponse instantiates a new ListStoredPaymentMethodsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListStoredPaymentMethodsResponse() *ListStoredPaymentMethodsResponse {
	this := ListStoredPaymentMethodsResponse{}
	return &this
}

// NewListStoredPaymentMethodsResponseWithDefaults instantiates a new ListStoredPaymentMethodsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListStoredPaymentMethodsResponseWithDefaults() *ListStoredPaymentMethodsResponse {
	this := ListStoredPaymentMethodsResponse{}
	return &this
}

// GetMerchantAccount returns the MerchantAccount field value if set, zero value otherwise.
func (o *ListStoredPaymentMethodsResponse) GetMerchantAccount() string {
	if o == nil || common.IsNil(o.MerchantAccount) {
		var ret string
		return ret
	}
	return *o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStoredPaymentMethodsResponse) GetMerchantAccountOk() (*string, bool) {
	if o == nil || common.IsNil(o.MerchantAccount) {
		return nil, false
	}
	return o.MerchantAccount, true
}

// HasMerchantAccount returns a boolean if a field has been set.
func (o *ListStoredPaymentMethodsResponse) HasMerchantAccount() bool {
	if o != nil && !common.IsNil(o.MerchantAccount) {
		return true
	}

	return false
}

// SetMerchantAccount gets a reference to the given string and assigns it to the MerchantAccount field.
func (o *ListStoredPaymentMethodsResponse) SetMerchantAccount(v string) {
	o.MerchantAccount = &v
}

// GetShopperReference returns the ShopperReference field value if set, zero value otherwise.
func (o *ListStoredPaymentMethodsResponse) GetShopperReference() string {
	if o == nil || common.IsNil(o.ShopperReference) {
		var ret string
		return ret
	}
	return *o.ShopperReference
}

// GetShopperReferenceOk returns a tuple with the ShopperReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStoredPaymentMethodsResponse) GetShopperReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.ShopperReference) {
		return nil, false
	}
	return o.ShopperReference, true
}

// HasShopperReference returns a boolean if a field has been set.
func (o *ListStoredPaymentMethodsResponse) HasShopperReference() bool {
	if o != nil && !common.IsNil(o.ShopperReference) {
		return true
	}

	return false
}

// SetShopperReference gets a reference to the given string and assigns it to the ShopperReference field.
func (o *ListStoredPaymentMethodsResponse) SetShopperReference(v string) {
	o.ShopperReference = &v
}

// GetStoredPaymentMethods returns the StoredPaymentMethods field value if set, zero value otherwise.
func (o *ListStoredPaymentMethodsResponse) GetStoredPaymentMethods() []StoredPaymentMethodResource {
	if o == nil || common.IsNil(o.StoredPaymentMethods) {
		var ret []StoredPaymentMethodResource
		return ret
	}
	return o.StoredPaymentMethods
}

// GetStoredPaymentMethodsOk returns a tuple with the StoredPaymentMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStoredPaymentMethodsResponse) GetStoredPaymentMethodsOk() ([]StoredPaymentMethodResource, bool) {
	if o == nil || common.IsNil(o.StoredPaymentMethods) {
		return nil, false
	}
	return o.StoredPaymentMethods, true
}

// HasStoredPaymentMethods returns a boolean if a field has been set.
func (o *ListStoredPaymentMethodsResponse) HasStoredPaymentMethods() bool {
	if o != nil && !common.IsNil(o.StoredPaymentMethods) {
		return true
	}

	return false
}

// SetStoredPaymentMethods gets a reference to the given []StoredPaymentMethodResource and assigns it to the StoredPaymentMethods field.
func (o *ListStoredPaymentMethodsResponse) SetStoredPaymentMethods(v []StoredPaymentMethodResource) {
	o.StoredPaymentMethods = v
}

func (o ListStoredPaymentMethodsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListStoredPaymentMethodsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.MerchantAccount) {
		toSerialize["merchantAccount"] = o.MerchantAccount
	}
	if !common.IsNil(o.ShopperReference) {
		toSerialize["shopperReference"] = o.ShopperReference
	}
	if !common.IsNil(o.StoredPaymentMethods) {
		toSerialize["storedPaymentMethods"] = o.StoredPaymentMethods
	}
	return toSerialize, nil
}

type NullableListStoredPaymentMethodsResponse struct {
	value *ListStoredPaymentMethodsResponse
	isSet bool
}

func (v NullableListStoredPaymentMethodsResponse) Get() *ListStoredPaymentMethodsResponse {
	return v.value
}

func (v *NullableListStoredPaymentMethodsResponse) Set(val *ListStoredPaymentMethodsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListStoredPaymentMethodsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListStoredPaymentMethodsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListStoredPaymentMethodsResponse(val *ListStoredPaymentMethodsResponse) *NullableListStoredPaymentMethodsResponse {
	return &NullableListStoredPaymentMethodsResponse{value: val, isSet: true}
}

func (v NullableListStoredPaymentMethodsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListStoredPaymentMethodsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
