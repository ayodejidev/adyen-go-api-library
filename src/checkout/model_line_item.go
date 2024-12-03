/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v15/src/common"
)

// checks if the LineItem type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &LineItem{}

// LineItem struct for LineItem
type LineItem struct {
	// Item amount excluding the tax, in minor units.
	AmountExcludingTax *int64 `json:"amountExcludingTax,omitempty"`
	// Item amount including the tax, in minor units.
	AmountIncludingTax *int64 `json:"amountIncludingTax,omitempty"`
	// Brand of the item.
	Brand *string `json:"brand,omitempty"`
	// Color of the item.
	Color *string `json:"color,omitempty"`
	// Description of the line item.
	Description *string `json:"description,omitempty"`
	// ID of the line item.
	Id *string `json:"id,omitempty"`
	// Link to the picture of the purchased item.
	ImageUrl *string `json:"imageUrl,omitempty"`
	// Item category, used by the payment methods PayPal and Ratepay.
	ItemCategory *string `json:"itemCategory,omitempty"`
	// Manufacturer of the item.
	Manufacturer *string `json:"manufacturer,omitempty"`
	// Marketplace seller id.
	MarketplaceSellerId *string `json:"marketplaceSellerId,omitempty"`
	// Link to the purchased item.
	ProductUrl *string `json:"productUrl,omitempty"`
	// Number of items.
	Quantity *int64 `json:"quantity,omitempty"`
	// Email associated with the given product in the basket (usually in electronic gift cards).
	ReceiverEmail *string `json:"receiverEmail,omitempty"`
	// Size of the item.
	Size *string `json:"size,omitempty"`
	// Stock keeping unit.
	Sku *string `json:"sku,omitempty"`
	// Tax amount, in minor units.
	TaxAmount *int64 `json:"taxAmount,omitempty"`
	// Tax percentage, in minor units.
	TaxPercentage *int64 `json:"taxPercentage,omitempty"`
	// Universal Product Code.
	Upc *string `json:"upc,omitempty"`
}

// NewLineItem instantiates a new LineItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLineItem() *LineItem {
	this := LineItem{}
	return &this
}

// NewLineItemWithDefaults instantiates a new LineItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineItemWithDefaults() *LineItem {
	this := LineItem{}
	return &this
}

// GetAmountExcludingTax returns the AmountExcludingTax field value if set, zero value otherwise.
func (o *LineItem) GetAmountExcludingTax() int64 {
	if o == nil || common.IsNil(o.AmountExcludingTax) {
		var ret int64
		return ret
	}
	return *o.AmountExcludingTax
}

// GetAmountExcludingTaxOk returns a tuple with the AmountExcludingTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetAmountExcludingTaxOk() (*int64, bool) {
	if o == nil || common.IsNil(o.AmountExcludingTax) {
		return nil, false
	}
	return o.AmountExcludingTax, true
}

// HasAmountExcludingTax returns a boolean if a field has been set.
func (o *LineItem) HasAmountExcludingTax() bool {
	if o != nil && !common.IsNil(o.AmountExcludingTax) {
		return true
	}

	return false
}

// SetAmountExcludingTax gets a reference to the given int64 and assigns it to the AmountExcludingTax field.
func (o *LineItem) SetAmountExcludingTax(v int64) {
	o.AmountExcludingTax = &v
}

// GetAmountIncludingTax returns the AmountIncludingTax field value if set, zero value otherwise.
func (o *LineItem) GetAmountIncludingTax() int64 {
	if o == nil || common.IsNil(o.AmountIncludingTax) {
		var ret int64
		return ret
	}
	return *o.AmountIncludingTax
}

// GetAmountIncludingTaxOk returns a tuple with the AmountIncludingTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetAmountIncludingTaxOk() (*int64, bool) {
	if o == nil || common.IsNil(o.AmountIncludingTax) {
		return nil, false
	}
	return o.AmountIncludingTax, true
}

// HasAmountIncludingTax returns a boolean if a field has been set.
func (o *LineItem) HasAmountIncludingTax() bool {
	if o != nil && !common.IsNil(o.AmountIncludingTax) {
		return true
	}

	return false
}

// SetAmountIncludingTax gets a reference to the given int64 and assigns it to the AmountIncludingTax field.
func (o *LineItem) SetAmountIncludingTax(v int64) {
	o.AmountIncludingTax = &v
}

// GetBrand returns the Brand field value if set, zero value otherwise.
func (o *LineItem) GetBrand() string {
	if o == nil || common.IsNil(o.Brand) {
		var ret string
		return ret
	}
	return *o.Brand
}

// GetBrandOk returns a tuple with the Brand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetBrandOk() (*string, bool) {
	if o == nil || common.IsNil(o.Brand) {
		return nil, false
	}
	return o.Brand, true
}

// HasBrand returns a boolean if a field has been set.
func (o *LineItem) HasBrand() bool {
	if o != nil && !common.IsNil(o.Brand) {
		return true
	}

	return false
}

// SetBrand gets a reference to the given string and assigns it to the Brand field.
func (o *LineItem) SetBrand(v string) {
	o.Brand = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *LineItem) GetColor() string {
	if o == nil || common.IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetColorOk() (*string, bool) {
	if o == nil || common.IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *LineItem) HasColor() bool {
	if o != nil && !common.IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *LineItem) SetColor(v string) {
	o.Color = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LineItem) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LineItem) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LineItem) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LineItem) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LineItem) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LineItem) SetId(v string) {
	o.Id = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *LineItem) GetImageUrl() string {
	if o == nil || common.IsNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetImageUrlOk() (*string, bool) {
	if o == nil || common.IsNil(o.ImageUrl) {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *LineItem) HasImageUrl() bool {
	if o != nil && !common.IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *LineItem) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetItemCategory returns the ItemCategory field value if set, zero value otherwise.
func (o *LineItem) GetItemCategory() string {
	if o == nil || common.IsNil(o.ItemCategory) {
		var ret string
		return ret
	}
	return *o.ItemCategory
}

// GetItemCategoryOk returns a tuple with the ItemCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetItemCategoryOk() (*string, bool) {
	if o == nil || common.IsNil(o.ItemCategory) {
		return nil, false
	}
	return o.ItemCategory, true
}

// HasItemCategory returns a boolean if a field has been set.
func (o *LineItem) HasItemCategory() bool {
	if o != nil && !common.IsNil(o.ItemCategory) {
		return true
	}

	return false
}

// SetItemCategory gets a reference to the given string and assigns it to the ItemCategory field.
func (o *LineItem) SetItemCategory(v string) {
	o.ItemCategory = &v
}

// GetManufacturer returns the Manufacturer field value if set, zero value otherwise.
func (o *LineItem) GetManufacturer() string {
	if o == nil || common.IsNil(o.Manufacturer) {
		var ret string
		return ret
	}
	return *o.Manufacturer
}

// GetManufacturerOk returns a tuple with the Manufacturer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetManufacturerOk() (*string, bool) {
	if o == nil || common.IsNil(o.Manufacturer) {
		return nil, false
	}
	return o.Manufacturer, true
}

// HasManufacturer returns a boolean if a field has been set.
func (o *LineItem) HasManufacturer() bool {
	if o != nil && !common.IsNil(o.Manufacturer) {
		return true
	}

	return false
}

// SetManufacturer gets a reference to the given string and assigns it to the Manufacturer field.
func (o *LineItem) SetManufacturer(v string) {
	o.Manufacturer = &v
}

// GetMarketplaceSellerId returns the MarketplaceSellerId field value if set, zero value otherwise.
func (o *LineItem) GetMarketplaceSellerId() string {
	if o == nil || common.IsNil(o.MarketplaceSellerId) {
		var ret string
		return ret
	}
	return *o.MarketplaceSellerId
}

// GetMarketplaceSellerIdOk returns a tuple with the MarketplaceSellerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetMarketplaceSellerIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarketplaceSellerId) {
		return nil, false
	}
	return o.MarketplaceSellerId, true
}

// HasMarketplaceSellerId returns a boolean if a field has been set.
func (o *LineItem) HasMarketplaceSellerId() bool {
	if o != nil && !common.IsNil(o.MarketplaceSellerId) {
		return true
	}

	return false
}

// SetMarketplaceSellerId gets a reference to the given string and assigns it to the MarketplaceSellerId field.
func (o *LineItem) SetMarketplaceSellerId(v string) {
	o.MarketplaceSellerId = &v
}

// GetProductUrl returns the ProductUrl field value if set, zero value otherwise.
func (o *LineItem) GetProductUrl() string {
	if o == nil || common.IsNil(o.ProductUrl) {
		var ret string
		return ret
	}
	return *o.ProductUrl
}

// GetProductUrlOk returns a tuple with the ProductUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetProductUrlOk() (*string, bool) {
	if o == nil || common.IsNil(o.ProductUrl) {
		return nil, false
	}
	return o.ProductUrl, true
}

// HasProductUrl returns a boolean if a field has been set.
func (o *LineItem) HasProductUrl() bool {
	if o != nil && !common.IsNil(o.ProductUrl) {
		return true
	}

	return false
}

// SetProductUrl gets a reference to the given string and assigns it to the ProductUrl field.
func (o *LineItem) SetProductUrl(v string) {
	o.ProductUrl = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *LineItem) GetQuantity() int64 {
	if o == nil || common.IsNil(o.Quantity) {
		var ret int64
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetQuantityOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *LineItem) HasQuantity() bool {
	if o != nil && !common.IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int64 and assigns it to the Quantity field.
func (o *LineItem) SetQuantity(v int64) {
	o.Quantity = &v
}

// GetReceiverEmail returns the ReceiverEmail field value if set, zero value otherwise.
func (o *LineItem) GetReceiverEmail() string {
	if o == nil || common.IsNil(o.ReceiverEmail) {
		var ret string
		return ret
	}
	return *o.ReceiverEmail
}

// GetReceiverEmailOk returns a tuple with the ReceiverEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetReceiverEmailOk() (*string, bool) {
	if o == nil || common.IsNil(o.ReceiverEmail) {
		return nil, false
	}
	return o.ReceiverEmail, true
}

// HasReceiverEmail returns a boolean if a field has been set.
func (o *LineItem) HasReceiverEmail() bool {
	if o != nil && !common.IsNil(o.ReceiverEmail) {
		return true
	}

	return false
}

// SetReceiverEmail gets a reference to the given string and assigns it to the ReceiverEmail field.
func (o *LineItem) SetReceiverEmail(v string) {
	o.ReceiverEmail = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *LineItem) GetSize() string {
	if o == nil || common.IsNil(o.Size) {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetSizeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *LineItem) HasSize() bool {
	if o != nil && !common.IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *LineItem) SetSize(v string) {
	o.Size = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *LineItem) GetSku() string {
	if o == nil || common.IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetSkuOk() (*string, bool) {
	if o == nil || common.IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *LineItem) HasSku() bool {
	if o != nil && !common.IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *LineItem) SetSku(v string) {
	o.Sku = &v
}

// GetTaxAmount returns the TaxAmount field value if set, zero value otherwise.
func (o *LineItem) GetTaxAmount() int64 {
	if o == nil || common.IsNil(o.TaxAmount) {
		var ret int64
		return ret
	}
	return *o.TaxAmount
}

// GetTaxAmountOk returns a tuple with the TaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetTaxAmountOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TaxAmount) {
		return nil, false
	}
	return o.TaxAmount, true
}

// HasTaxAmount returns a boolean if a field has been set.
func (o *LineItem) HasTaxAmount() bool {
	if o != nil && !common.IsNil(o.TaxAmount) {
		return true
	}

	return false
}

// SetTaxAmount gets a reference to the given int64 and assigns it to the TaxAmount field.
func (o *LineItem) SetTaxAmount(v int64) {
	o.TaxAmount = &v
}

// GetTaxPercentage returns the TaxPercentage field value if set, zero value otherwise.
func (o *LineItem) GetTaxPercentage() int64 {
	if o == nil || common.IsNil(o.TaxPercentage) {
		var ret int64
		return ret
	}
	return *o.TaxPercentage
}

// GetTaxPercentageOk returns a tuple with the TaxPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetTaxPercentageOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TaxPercentage) {
		return nil, false
	}
	return o.TaxPercentage, true
}

// HasTaxPercentage returns a boolean if a field has been set.
func (o *LineItem) HasTaxPercentage() bool {
	if o != nil && !common.IsNil(o.TaxPercentage) {
		return true
	}

	return false
}

// SetTaxPercentage gets a reference to the given int64 and assigns it to the TaxPercentage field.
func (o *LineItem) SetTaxPercentage(v int64) {
	o.TaxPercentage = &v
}

// GetUpc returns the Upc field value if set, zero value otherwise.
func (o *LineItem) GetUpc() string {
	if o == nil || common.IsNil(o.Upc) {
		var ret string
		return ret
	}
	return *o.Upc
}

// GetUpcOk returns a tuple with the Upc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItem) GetUpcOk() (*string, bool) {
	if o == nil || common.IsNil(o.Upc) {
		return nil, false
	}
	return o.Upc, true
}

// HasUpc returns a boolean if a field has been set.
func (o *LineItem) HasUpc() bool {
	if o != nil && !common.IsNil(o.Upc) {
		return true
	}

	return false
}

// SetUpc gets a reference to the given string and assigns it to the Upc field.
func (o *LineItem) SetUpc(v string) {
	o.Upc = &v
}

func (o LineItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LineItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AmountExcludingTax) {
		toSerialize["amountExcludingTax"] = o.AmountExcludingTax
	}
	if !common.IsNil(o.AmountIncludingTax) {
		toSerialize["amountIncludingTax"] = o.AmountIncludingTax
	}
	if !common.IsNil(o.Brand) {
		toSerialize["brand"] = o.Brand
	}
	if !common.IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	if !common.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.ImageUrl) {
		toSerialize["imageUrl"] = o.ImageUrl
	}
	if !common.IsNil(o.ItemCategory) {
		toSerialize["itemCategory"] = o.ItemCategory
	}
	if !common.IsNil(o.Manufacturer) {
		toSerialize["manufacturer"] = o.Manufacturer
	}
	if !common.IsNil(o.MarketplaceSellerId) {
		toSerialize["marketplaceSellerId"] = o.MarketplaceSellerId
	}
	if !common.IsNil(o.ProductUrl) {
		toSerialize["productUrl"] = o.ProductUrl
	}
	if !common.IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !common.IsNil(o.ReceiverEmail) {
		toSerialize["receiverEmail"] = o.ReceiverEmail
	}
	if !common.IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !common.IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	if !common.IsNil(o.TaxAmount) {
		toSerialize["taxAmount"] = o.TaxAmount
	}
	if !common.IsNil(o.TaxPercentage) {
		toSerialize["taxPercentage"] = o.TaxPercentage
	}
	if !common.IsNil(o.Upc) {
		toSerialize["upc"] = o.Upc
	}
	return toSerialize, nil
}

type NullableLineItem struct {
	value *LineItem
	isSet bool
}

func (v NullableLineItem) Get() *LineItem {
	return v.value
}

func (v *NullableLineItem) Set(val *LineItem) {
	v.value = val
	v.isSet = true
}

func (v NullableLineItem) IsSet() bool {
	return v.isSet
}

func (v *NullableLineItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineItem(val *LineItem) *NullableLineItem {
	return &NullableLineItem{value: val, isSet: true}
}

func (v NullableLineItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
