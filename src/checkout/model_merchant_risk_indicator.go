/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"
	"time"

	"github.com/adyen/adyen-go-api-library/v11/src/common"
)

// checks if the MerchantRiskIndicator type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MerchantRiskIndicator{}

// MerchantRiskIndicator struct for MerchantRiskIndicator
type MerchantRiskIndicator struct {
	// Whether the chosen delivery address is identical to the billing address.
	AddressMatch *bool `json:"addressMatch,omitempty"`
	// Indicator regarding the delivery address. Allowed values: * `shipToBillingAddress` * `shipToVerifiedAddress` * `shipToNewAddress` * `shipToStore` * `digitalGoods` * `goodsNotShipped` * `other`
	DeliveryAddressIndicator *string `json:"deliveryAddressIndicator,omitempty"`
	// The delivery email address (for digital goods).
	// Deprecated
	DeliveryEmail *string `json:"deliveryEmail,omitempty"`
	// For Electronic delivery, the email address to which the merchandise was delivered. Maximum length: 254 characters.
	DeliveryEmailAddress *string `json:"deliveryEmailAddress,omitempty"`
	// The estimated delivery time for the shopper to receive the goods. Allowed values: * `electronicDelivery` * `sameDayShipping` * `overnightShipping` * `twoOrMoreDaysShipping`
	DeliveryTimeframe *string `json:"deliveryTimeframe,omitempty"`
	GiftCardAmount    *Amount `json:"giftCardAmount,omitempty"`
	// For prepaid or gift card purchase, total count of individual prepaid or gift cards/codes purchased.
	GiftCardCount *int32 `json:"giftCardCount,omitempty"`
	// For prepaid or gift card purchase, [ISO 4217](https://www.iso.org/iso-4217-currency-codes.html) three-digit currency code of the gift card, other than those listed in Table A.5 of the EMVCo 3D Secure Protocol and Core Functions Specification.
	GiftCardCurr *string `json:"giftCardCurr,omitempty"`
	// For pre-order purchases, the expected date this product will be available to the shopper.
	PreOrderDate *time.Time `json:"preOrderDate,omitempty"`
	// Indicator for whether this transaction is for pre-ordering a product.
	PreOrderPurchase *bool `json:"preOrderPurchase,omitempty"`
	// Indicates whether Cardholder is placing an order for merchandise with a future availability or release date.
	PreOrderPurchaseInd *string `json:"preOrderPurchaseInd,omitempty"`
	// Indicator for whether the shopper has already purchased the same items in the past.
	ReorderItems *bool `json:"reorderItems,omitempty"`
	// Indicates whether the cardholder is reordering previously purchased merchandise.
	ReorderItemsInd *string `json:"reorderItemsInd,omitempty"`
	// Indicates shipping method chosen for the transaction.
	ShipIndicator *string `json:"shipIndicator,omitempty"`
}

// NewMerchantRiskIndicator instantiates a new MerchantRiskIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantRiskIndicator() *MerchantRiskIndicator {
	this := MerchantRiskIndicator{}
	return &this
}

// NewMerchantRiskIndicatorWithDefaults instantiates a new MerchantRiskIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantRiskIndicatorWithDefaults() *MerchantRiskIndicator {
	this := MerchantRiskIndicator{}
	return &this
}

// GetAddressMatch returns the AddressMatch field value if set, zero value otherwise.
func (o *MerchantRiskIndicator) GetAddressMatch() bool {
	if o == nil || common.IsNil(o.AddressMatch) {
		var ret bool
		return ret
	}
	return *o.AddressMatch
}

// GetAddressMatchOk returns a tuple with the AddressMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantRiskIndicator) GetAddressMatchOk() (*bool, bool) {
	if o == nil || common.IsNil(o.AddressMatch) {
		return nil, false
	}
	return o.AddressMatch, true
}

// HasAddressMatch returns a boolean if a field has been set.
func (o *MerchantRiskIndicator) HasAddressMatch() bool {
	if o != nil && !common.IsNil(o.AddressMatch) {
		return true
	}

	return false
}

// SetAddressMatch gets a reference to the given bool and assigns it to the AddressMatch field.
func (o *MerchantRiskIndicator) SetAddressMatch(v bool) {
	o.AddressMatch = &v
}

// GetDeliveryAddressIndicator returns the DeliveryAddressIndicator field value if set, zero value otherwise.
func (o *MerchantRiskIndicator) GetDeliveryAddressIndicator() string {
	if o == nil || common.IsNil(o.DeliveryAddressIndicator) {
		var ret string
		return ret
	}
	return *o.DeliveryAddressIndicator
}

// GetDeliveryAddressIndicatorOk returns a tuple with the DeliveryAddressIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantRiskIndicator) GetDeliveryAddressIndicatorOk() (*string, bool) {
	if o == nil || common.IsNil(o.DeliveryAddressIndicator) {
		return nil, false
	}
	return o.DeliveryAddressIndicator, true
}

// HasDeliveryAddressIndicator returns a boolean if a field has been set.
func (o *MerchantRiskIndicator) HasDeliveryAddressIndicator() bool {
	if o != nil && !common.IsNil(o.DeliveryAddressIndicator) {
		return true
	}

	return false
}

// SetDeliveryAddressIndicator gets a reference to the given string and assigns it to the DeliveryAddressIndicator field.
func (o *MerchantRiskIndicator) SetDeliveryAddressIndicator(v string) {
	o.DeliveryAddressIndicator = &v
}

// GetDeliveryEmail returns the DeliveryEmail field value if set, zero value otherwise.
// Deprecated
func (o *MerchantRiskIndicator) GetDeliveryEmail() string {
	if o == nil || common.IsNil(o.DeliveryEmail) {
		var ret string
		return ret
	}
	return *o.DeliveryEmail
}

// GetDeliveryEmailOk returns a tuple with the DeliveryEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *MerchantRiskIndicator) GetDeliveryEmailOk() (*string, bool) {
	if o == nil || common.IsNil(o.DeliveryEmail) {
		return nil, false
	}
	return o.DeliveryEmail, true
}

// HasDeliveryEmail returns a boolean if a field has been set.
func (o *MerchantRiskIndicator) HasDeliveryEmail() bool {
	if o != nil && !common.IsNil(o.DeliveryEmail) {
		return true
	}

	return false
}

// SetDeliveryEmail gets a reference to the given string and assigns it to the DeliveryEmail field.
// Deprecated
func (o *MerchantRiskIndicator) SetDeliveryEmail(v string) {
	o.DeliveryEmail = &v
}

// GetDeliveryEmailAddress returns the DeliveryEmailAddress field value if set, zero value otherwise.
func (o *MerchantRiskIndicator) GetDeliveryEmailAddress() string {
	if o == nil || common.IsNil(o.DeliveryEmailAddress) {
		var ret string
		return ret
	}
	return *o.DeliveryEmailAddress
}

// GetDeliveryEmailAddressOk returns a tuple with the DeliveryEmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantRiskIndicator) GetDeliveryEmailAddressOk() (*string, bool) {
	if o == nil || common.IsNil(o.DeliveryEmailAddress) {
		return nil, false
	}
	return o.DeliveryEmailAddress, true
}

// HasDeliveryEmailAddress returns a boolean if a field has been set.
func (o *MerchantRiskIndicator) HasDeliveryEmailAddress() bool {
	if o != nil && !common.IsNil(o.DeliveryEmailAddress) {
		return true
	}

	return false
}

// SetDeliveryEmailAddress gets a reference to the given string and assigns it to the DeliveryEmailAddress field.
func (o *MerchantRiskIndicator) SetDeliveryEmailAddress(v string) {
	o.DeliveryEmailAddress = &v
}

// GetDeliveryTimeframe returns the DeliveryTimeframe field value if set, zero value otherwise.
func (o *MerchantRiskIndicator) GetDeliveryTimeframe() string {
	if o == nil || common.IsNil(o.DeliveryTimeframe) {
		var ret string
		return ret
	}
	return *o.DeliveryTimeframe
}

// GetDeliveryTimeframeOk returns a tuple with the DeliveryTimeframe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantRiskIndicator) GetDeliveryTimeframeOk() (*string, bool) {
	if o == nil || common.IsNil(o.DeliveryTimeframe) {
		return nil, false
	}
	return o.DeliveryTimeframe, true
}

// HasDeliveryTimeframe returns a boolean if a field has been set.
func (o *MerchantRiskIndicator) HasDeliveryTimeframe() bool {
	if o != nil && !common.IsNil(o.DeliveryTimeframe) {
		return true
	}

	return false
}

// SetDeliveryTimeframe gets a reference to the given string and assigns it to the DeliveryTimeframe field.
func (o *MerchantRiskIndicator) SetDeliveryTimeframe(v string) {
	o.DeliveryTimeframe = &v
}

// GetGiftCardAmount returns the GiftCardAmount field value if set, zero value otherwise.
func (o *MerchantRiskIndicator) GetGiftCardAmount() Amount {
	if o == nil || common.IsNil(o.GiftCardAmount) {
		var ret Amount
		return ret
	}
	return *o.GiftCardAmount
}

// GetGiftCardAmountOk returns a tuple with the GiftCardAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantRiskIndicator) GetGiftCardAmountOk() (*Amount, bool) {
	if o == nil || common.IsNil(o.GiftCardAmount) {
		return nil, false
	}
	return o.GiftCardAmount, true
}

// HasGiftCardAmount returns a boolean if a field has been set.
func (o *MerchantRiskIndicator) HasGiftCardAmount() bool {
	if o != nil && !common.IsNil(o.GiftCardAmount) {
		return true
	}

	return false
}

// SetGiftCardAmount gets a reference to the given Amount and assigns it to the GiftCardAmount field.
func (o *MerchantRiskIndicator) SetGiftCardAmount(v Amount) {
	o.GiftCardAmount = &v
}

// GetGiftCardCount returns the GiftCardCount field value if set, zero value otherwise.
func (o *MerchantRiskIndicator) GetGiftCardCount() int32 {
	if o == nil || common.IsNil(o.GiftCardCount) {
		var ret int32
		return ret
	}
	return *o.GiftCardCount
}

// GetGiftCardCountOk returns a tuple with the GiftCardCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantRiskIndicator) GetGiftCardCountOk() (*int32, bool) {
	if o == nil || common.IsNil(o.GiftCardCount) {
		return nil, false
	}
	return o.GiftCardCount, true
}

// HasGiftCardCount returns a boolean if a field has been set.
func (o *MerchantRiskIndicator) HasGiftCardCount() bool {
	if o != nil && !common.IsNil(o.GiftCardCount) {
		return true
	}

	return false
}

// SetGiftCardCount gets a reference to the given int32 and assigns it to the GiftCardCount field.
func (o *MerchantRiskIndicator) SetGiftCardCount(v int32) {
	o.GiftCardCount = &v
}

// GetGiftCardCurr returns the GiftCardCurr field value if set, zero value otherwise.
func (o *MerchantRiskIndicator) GetGiftCardCurr() string {
	if o == nil || common.IsNil(o.GiftCardCurr) {
		var ret string
		return ret
	}
	return *o.GiftCardCurr
}

// GetGiftCardCurrOk returns a tuple with the GiftCardCurr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantRiskIndicator) GetGiftCardCurrOk() (*string, bool) {
	if o == nil || common.IsNil(o.GiftCardCurr) {
		return nil, false
	}
	return o.GiftCardCurr, true
}

// HasGiftCardCurr returns a boolean if a field has been set.
func (o *MerchantRiskIndicator) HasGiftCardCurr() bool {
	if o != nil && !common.IsNil(o.GiftCardCurr) {
		return true
	}

	return false
}

// SetGiftCardCurr gets a reference to the given string and assigns it to the GiftCardCurr field.
func (o *MerchantRiskIndicator) SetGiftCardCurr(v string) {
	o.GiftCardCurr = &v
}

// GetPreOrderDate returns the PreOrderDate field value if set, zero value otherwise.
func (o *MerchantRiskIndicator) GetPreOrderDate() time.Time {
	if o == nil || common.IsNil(o.PreOrderDate) {
		var ret time.Time
		return ret
	}
	return *o.PreOrderDate
}

// GetPreOrderDateOk returns a tuple with the PreOrderDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantRiskIndicator) GetPreOrderDateOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.PreOrderDate) {
		return nil, false
	}
	return o.PreOrderDate, true
}

// HasPreOrderDate returns a boolean if a field has been set.
func (o *MerchantRiskIndicator) HasPreOrderDate() bool {
	if o != nil && !common.IsNil(o.PreOrderDate) {
		return true
	}

	return false
}

// SetPreOrderDate gets a reference to the given time.Time and assigns it to the PreOrderDate field.
func (o *MerchantRiskIndicator) SetPreOrderDate(v time.Time) {
	o.PreOrderDate = &v
}

// GetPreOrderPurchase returns the PreOrderPurchase field value if set, zero value otherwise.
func (o *MerchantRiskIndicator) GetPreOrderPurchase() bool {
	if o == nil || common.IsNil(o.PreOrderPurchase) {
		var ret bool
		return ret
	}
	return *o.PreOrderPurchase
}

// GetPreOrderPurchaseOk returns a tuple with the PreOrderPurchase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantRiskIndicator) GetPreOrderPurchaseOk() (*bool, bool) {
	if o == nil || common.IsNil(o.PreOrderPurchase) {
		return nil, false
	}
	return o.PreOrderPurchase, true
}

// HasPreOrderPurchase returns a boolean if a field has been set.
func (o *MerchantRiskIndicator) HasPreOrderPurchase() bool {
	if o != nil && !common.IsNil(o.PreOrderPurchase) {
		return true
	}

	return false
}

// SetPreOrderPurchase gets a reference to the given bool and assigns it to the PreOrderPurchase field.
func (o *MerchantRiskIndicator) SetPreOrderPurchase(v bool) {
	o.PreOrderPurchase = &v
}

// GetPreOrderPurchaseInd returns the PreOrderPurchaseInd field value if set, zero value otherwise.
func (o *MerchantRiskIndicator) GetPreOrderPurchaseInd() string {
	if o == nil || common.IsNil(o.PreOrderPurchaseInd) {
		var ret string
		return ret
	}
	return *o.PreOrderPurchaseInd
}

// GetPreOrderPurchaseIndOk returns a tuple with the PreOrderPurchaseInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantRiskIndicator) GetPreOrderPurchaseIndOk() (*string, bool) {
	if o == nil || common.IsNil(o.PreOrderPurchaseInd) {
		return nil, false
	}
	return o.PreOrderPurchaseInd, true
}

// HasPreOrderPurchaseInd returns a boolean if a field has been set.
func (o *MerchantRiskIndicator) HasPreOrderPurchaseInd() bool {
	if o != nil && !common.IsNil(o.PreOrderPurchaseInd) {
		return true
	}

	return false
}

// SetPreOrderPurchaseInd gets a reference to the given string and assigns it to the PreOrderPurchaseInd field.
func (o *MerchantRiskIndicator) SetPreOrderPurchaseInd(v string) {
	o.PreOrderPurchaseInd = &v
}

// GetReorderItems returns the ReorderItems field value if set, zero value otherwise.
func (o *MerchantRiskIndicator) GetReorderItems() bool {
	if o == nil || common.IsNil(o.ReorderItems) {
		var ret bool
		return ret
	}
	return *o.ReorderItems
}

// GetReorderItemsOk returns a tuple with the ReorderItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantRiskIndicator) GetReorderItemsOk() (*bool, bool) {
	if o == nil || common.IsNil(o.ReorderItems) {
		return nil, false
	}
	return o.ReorderItems, true
}

// HasReorderItems returns a boolean if a field has been set.
func (o *MerchantRiskIndicator) HasReorderItems() bool {
	if o != nil && !common.IsNil(o.ReorderItems) {
		return true
	}

	return false
}

// SetReorderItems gets a reference to the given bool and assigns it to the ReorderItems field.
func (o *MerchantRiskIndicator) SetReorderItems(v bool) {
	o.ReorderItems = &v
}

// GetReorderItemsInd returns the ReorderItemsInd field value if set, zero value otherwise.
func (o *MerchantRiskIndicator) GetReorderItemsInd() string {
	if o == nil || common.IsNil(o.ReorderItemsInd) {
		var ret string
		return ret
	}
	return *o.ReorderItemsInd
}

// GetReorderItemsIndOk returns a tuple with the ReorderItemsInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantRiskIndicator) GetReorderItemsIndOk() (*string, bool) {
	if o == nil || common.IsNil(o.ReorderItemsInd) {
		return nil, false
	}
	return o.ReorderItemsInd, true
}

// HasReorderItemsInd returns a boolean if a field has been set.
func (o *MerchantRiskIndicator) HasReorderItemsInd() bool {
	if o != nil && !common.IsNil(o.ReorderItemsInd) {
		return true
	}

	return false
}

// SetReorderItemsInd gets a reference to the given string and assigns it to the ReorderItemsInd field.
func (o *MerchantRiskIndicator) SetReorderItemsInd(v string) {
	o.ReorderItemsInd = &v
}

// GetShipIndicator returns the ShipIndicator field value if set, zero value otherwise.
func (o *MerchantRiskIndicator) GetShipIndicator() string {
	if o == nil || common.IsNil(o.ShipIndicator) {
		var ret string
		return ret
	}
	return *o.ShipIndicator
}

// GetShipIndicatorOk returns a tuple with the ShipIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantRiskIndicator) GetShipIndicatorOk() (*string, bool) {
	if o == nil || common.IsNil(o.ShipIndicator) {
		return nil, false
	}
	return o.ShipIndicator, true
}

// HasShipIndicator returns a boolean if a field has been set.
func (o *MerchantRiskIndicator) HasShipIndicator() bool {
	if o != nil && !common.IsNil(o.ShipIndicator) {
		return true
	}

	return false
}

// SetShipIndicator gets a reference to the given string and assigns it to the ShipIndicator field.
func (o *MerchantRiskIndicator) SetShipIndicator(v string) {
	o.ShipIndicator = &v
}

func (o MerchantRiskIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantRiskIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AddressMatch) {
		toSerialize["addressMatch"] = o.AddressMatch
	}
	if !common.IsNil(o.DeliveryAddressIndicator) {
		toSerialize["deliveryAddressIndicator"] = o.DeliveryAddressIndicator
	}
	if !common.IsNil(o.DeliveryEmail) {
		toSerialize["deliveryEmail"] = o.DeliveryEmail
	}
	if !common.IsNil(o.DeliveryEmailAddress) {
		toSerialize["deliveryEmailAddress"] = o.DeliveryEmailAddress
	}
	if !common.IsNil(o.DeliveryTimeframe) {
		toSerialize["deliveryTimeframe"] = o.DeliveryTimeframe
	}
	if !common.IsNil(o.GiftCardAmount) {
		toSerialize["giftCardAmount"] = o.GiftCardAmount
	}
	if !common.IsNil(o.GiftCardCount) {
		toSerialize["giftCardCount"] = o.GiftCardCount
	}
	if !common.IsNil(o.GiftCardCurr) {
		toSerialize["giftCardCurr"] = o.GiftCardCurr
	}
	if !common.IsNil(o.PreOrderDate) {
		toSerialize["preOrderDate"] = o.PreOrderDate
	}
	if !common.IsNil(o.PreOrderPurchase) {
		toSerialize["preOrderPurchase"] = o.PreOrderPurchase
	}
	if !common.IsNil(o.PreOrderPurchaseInd) {
		toSerialize["preOrderPurchaseInd"] = o.PreOrderPurchaseInd
	}
	if !common.IsNil(o.ReorderItems) {
		toSerialize["reorderItems"] = o.ReorderItems
	}
	if !common.IsNil(o.ReorderItemsInd) {
		toSerialize["reorderItemsInd"] = o.ReorderItemsInd
	}
	if !common.IsNil(o.ShipIndicator) {
		toSerialize["shipIndicator"] = o.ShipIndicator
	}
	return toSerialize, nil
}

type NullableMerchantRiskIndicator struct {
	value *MerchantRiskIndicator
	isSet bool
}

func (v NullableMerchantRiskIndicator) Get() *MerchantRiskIndicator {
	return v.value
}

func (v *NullableMerchantRiskIndicator) Set(val *MerchantRiskIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantRiskIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantRiskIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantRiskIndicator(val *MerchantRiskIndicator) *NullableMerchantRiskIndicator {
	return &NullableMerchantRiskIndicator{value: val, isSet: true}
}

func (v NullableMerchantRiskIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantRiskIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *MerchantRiskIndicator) isValidDeliveryAddressIndicator() bool {
	var allowedEnumValues = []string{"shipToBillingAddress", "shipToVerifiedAddress", "shipToNewAddress", "shipToStore", "digitalGoods", "goodsNotShipped", "other"}
	for _, allowed := range allowedEnumValues {
		if o.GetDeliveryAddressIndicator() == allowed {
			return true
		}
	}
	return false
}
func (o *MerchantRiskIndicator) isValidDeliveryTimeframe() bool {
	var allowedEnumValues = []string{"electronicDelivery", "sameDayShipping", "overnightShipping", "twoOrMoreDaysShipping"}
	for _, allowed := range allowedEnumValues {
		if o.GetDeliveryTimeframe() == allowed {
			return true
		}
	}
	return false
}
