/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v12/src/common"
)

// checks if the DragonpayDetails type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DragonpayDetails{}

// DragonpayDetails struct for DragonpayDetails
type DragonpayDetails struct {
	// The checkout attempt identifier.
	CheckoutAttemptId *string `json:"checkoutAttemptId,omitempty"`
	// The Dragonpay issuer value of the shopper's selected bank. Set this to an **id** of a Dragonpay issuer to preselect it.
	Issuer string `json:"issuer"`
	// The shopper’s email address.
	ShopperEmail *string `json:"shopperEmail,omitempty"`
	// **dragonpay**
	Type string `json:"type"`
}

// NewDragonpayDetails instantiates a new DragonpayDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDragonpayDetails(issuer string, type_ string) *DragonpayDetails {
	this := DragonpayDetails{}
	this.Issuer = issuer
	this.Type = type_
	return &this
}

// NewDragonpayDetailsWithDefaults instantiates a new DragonpayDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDragonpayDetailsWithDefaults() *DragonpayDetails {
	this := DragonpayDetails{}
	return &this
}

// GetCheckoutAttemptId returns the CheckoutAttemptId field value if set, zero value otherwise.
func (o *DragonpayDetails) GetCheckoutAttemptId() string {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		var ret string
		return ret
	}
	return *o.CheckoutAttemptId
}

// GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DragonpayDetails) GetCheckoutAttemptIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		return nil, false
	}
	return o.CheckoutAttemptId, true
}

// HasCheckoutAttemptId returns a boolean if a field has been set.
func (o *DragonpayDetails) HasCheckoutAttemptId() bool {
	if o != nil && !common.IsNil(o.CheckoutAttemptId) {
		return true
	}

	return false
}

// SetCheckoutAttemptId gets a reference to the given string and assigns it to the CheckoutAttemptId field.
func (o *DragonpayDetails) SetCheckoutAttemptId(v string) {
	o.CheckoutAttemptId = &v
}

// GetIssuer returns the Issuer field value
func (o *DragonpayDetails) GetIssuer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value
// and a boolean to check if the value has been set.
func (o *DragonpayDetails) GetIssuerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Issuer, true
}

// SetIssuer sets field value
func (o *DragonpayDetails) SetIssuer(v string) {
	o.Issuer = v
}

// GetShopperEmail returns the ShopperEmail field value if set, zero value otherwise.
func (o *DragonpayDetails) GetShopperEmail() string {
	if o == nil || common.IsNil(o.ShopperEmail) {
		var ret string
		return ret
	}
	return *o.ShopperEmail
}

// GetShopperEmailOk returns a tuple with the ShopperEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DragonpayDetails) GetShopperEmailOk() (*string, bool) {
	if o == nil || common.IsNil(o.ShopperEmail) {
		return nil, false
	}
	return o.ShopperEmail, true
}

// HasShopperEmail returns a boolean if a field has been set.
func (o *DragonpayDetails) HasShopperEmail() bool {
	if o != nil && !common.IsNil(o.ShopperEmail) {
		return true
	}

	return false
}

// SetShopperEmail gets a reference to the given string and assigns it to the ShopperEmail field.
func (o *DragonpayDetails) SetShopperEmail(v string) {
	o.ShopperEmail = &v
}

// GetType returns the Type field value
func (o *DragonpayDetails) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DragonpayDetails) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DragonpayDetails) SetType(v string) {
	o.Type = v
}

func (o DragonpayDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DragonpayDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CheckoutAttemptId) {
		toSerialize["checkoutAttemptId"] = o.CheckoutAttemptId
	}
	toSerialize["issuer"] = o.Issuer
	if !common.IsNil(o.ShopperEmail) {
		toSerialize["shopperEmail"] = o.ShopperEmail
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableDragonpayDetails struct {
	value *DragonpayDetails
	isSet bool
}

func (v NullableDragonpayDetails) Get() *DragonpayDetails {
	return v.value
}

func (v *NullableDragonpayDetails) Set(val *DragonpayDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableDragonpayDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableDragonpayDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDragonpayDetails(val *DragonpayDetails) *NullableDragonpayDetails {
	return &NullableDragonpayDetails{value: val, isSet: true}
}

func (v NullableDragonpayDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDragonpayDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *DragonpayDetails) isValidType() bool {
	var allowedEnumValues = []string{"dragonpay_ebanking", "dragonpay_otc_banking", "dragonpay_otc_non_banking", "dragonpay_otc_philippines"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
