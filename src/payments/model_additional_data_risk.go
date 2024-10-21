/*
Adyen Payment API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package payments

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v14/src/common"
)

// checks if the AdditionalDataRisk type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AdditionalDataRisk{}

// AdditionalDataRisk struct for AdditionalDataRisk
type AdditionalDataRisk struct {
	// The data for your custom risk field. For more information, refer to [Create custom risk fields](https://docs.adyen.com/risk-management/configure-custom-risk-rules#step-1-create-custom-risk-fields).
	RiskdataCustomFieldName *string `json:"riskdata.[customFieldName],omitempty"`
	// The price of item in the basket, represented in [minor units](https://docs.adyen.com/development-resources/currency-codes).
	RiskdataBasketItemItemNrAmountPerItem *string `json:"riskdata.basket.item[itemNr].amountPerItem,omitempty"`
	// Brand of the item.
	RiskdataBasketItemItemNrBrand *string `json:"riskdata.basket.item[itemNr].brand,omitempty"`
	// Category of the item.
	RiskdataBasketItemItemNrCategory *string `json:"riskdata.basket.item[itemNr].category,omitempty"`
	// Color of the item.
	RiskdataBasketItemItemNrColor *string `json:"riskdata.basket.item[itemNr].color,omitempty"`
	// The three-character [ISO currency code](https://en.wikipedia.org/wiki/ISO_4217).
	RiskdataBasketItemItemNrCurrency *string `json:"riskdata.basket.item[itemNr].currency,omitempty"`
	// ID of the item.
	RiskdataBasketItemItemNrItemID *string `json:"riskdata.basket.item[itemNr].itemID,omitempty"`
	// Manufacturer of the item.
	RiskdataBasketItemItemNrManufacturer *string `json:"riskdata.basket.item[itemNr].manufacturer,omitempty"`
	// A text description of the product the invoice line refers to.
	RiskdataBasketItemItemNrProductTitle *string `json:"riskdata.basket.item[itemNr].productTitle,omitempty"`
	// Quantity of the item purchased.
	RiskdataBasketItemItemNrQuantity *string `json:"riskdata.basket.item[itemNr].quantity,omitempty"`
	// Email associated with the given product in the basket (usually in electronic gift cards).
	RiskdataBasketItemItemNrReceiverEmail *string `json:"riskdata.basket.item[itemNr].receiverEmail,omitempty"`
	// Size of the item.
	RiskdataBasketItemItemNrSize *string `json:"riskdata.basket.item[itemNr].size,omitempty"`
	// [Stock keeping unit](https://en.wikipedia.org/wiki/Stock_keeping_unit).
	RiskdataBasketItemItemNrSku *string `json:"riskdata.basket.item[itemNr].sku,omitempty"`
	// [Universal Product Code](https://en.wikipedia.org/wiki/Universal_Product_Code).
	RiskdataBasketItemItemNrUpc *string `json:"riskdata.basket.item[itemNr].upc,omitempty"`
	// Code of the promotion.
	RiskdataPromotionsPromotionItemNrPromotionCode *string `json:"riskdata.promotions.promotion[itemNr].promotionCode,omitempty"`
	// The discount amount of the promotion, represented in [minor units](https://docs.adyen.com/development-resources/currency-codes).
	RiskdataPromotionsPromotionItemNrPromotionDiscountAmount *string `json:"riskdata.promotions.promotion[itemNr].promotionDiscountAmount,omitempty"`
	// The three-character [ISO currency code](https://en.wikipedia.org/wiki/ISO_4217).
	RiskdataPromotionsPromotionItemNrPromotionDiscountCurrency *string `json:"riskdata.promotions.promotion[itemNr].promotionDiscountCurrency,omitempty"`
	// Promotion's percentage discount. It is represented in percentage value and there is no need to include the '%' sign.  e.g. for a promotion discount of 30%, the value of the field should be 30.
	RiskdataPromotionsPromotionItemNrPromotionDiscountPercentage *string `json:"riskdata.promotions.promotion[itemNr].promotionDiscountPercentage,omitempty"`
	// Name of the promotion.
	RiskdataPromotionsPromotionItemNrPromotionName *string `json:"riskdata.promotions.promotion[itemNr].promotionName,omitempty"`
	// Reference number of the risk profile that you want to apply to the payment. If not provided or left blank, the merchant-level account's default risk profile will be applied to the payment. For more information, see [dynamically assign a risk profile to a payment](https://docs.adyen.com/risk-management/create-and-use-risk-profiles#dynamically-assign-a-risk-profile-to-a-payment).
	RiskdataRiskProfileReference *string `json:"riskdata.riskProfileReference,omitempty"`
	// If this parameter is provided with the value **true**, risk checks for the payment request are skipped and the transaction will not get a risk score.
	RiskdataSkipRisk *string `json:"riskdata.skipRisk,omitempty"`
}

// NewAdditionalDataRisk instantiates a new AdditionalDataRisk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalDataRisk() *AdditionalDataRisk {
	this := AdditionalDataRisk{}
	return &this
}

// NewAdditionalDataRiskWithDefaults instantiates a new AdditionalDataRisk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalDataRiskWithDefaults() *AdditionalDataRisk {
	this := AdditionalDataRisk{}
	return &this
}

// GetRiskdataCustomFieldName returns the RiskdataCustomFieldName field value if set, zero value otherwise.
func (o *AdditionalDataRisk) GetRiskdataCustomFieldName() string {
	if o == nil || common.IsNil(o.RiskdataCustomFieldName) {
		var ret string
		return ret
	}
	return *o.RiskdataCustomFieldName
}

// GetRiskdataCustomFieldNameOk returns a tuple with the RiskdataCustomFieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataRisk) GetRiskdataCustomFieldNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.RiskdataCustomFieldName) {
		return nil, false
	}
	return o.RiskdataCustomFieldName, true
}

// HasRiskdataCustomFieldName returns a boolean if a field has been set.
func (o *AdditionalDataRisk) HasRiskdataCustomFieldName() bool {
	if o != nil && !common.IsNil(o.RiskdataCustomFieldName) {
		return true
	}

	return false
}

// SetRiskdataCustomFieldName gets a reference to the given string and assigns it to the RiskdataCustomFieldName field.
func (o *AdditionalDataRisk) SetRiskdataCustomFieldName(v string) {
	o.RiskdataCustomFieldName = &v
}

// GetRiskdataBasketItemItemNrAmountPerItem returns the RiskdataBasketItemItemNrAmountPerItem field value if set, zero value otherwise.
func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrAmountPerItem() string {
	if o == nil || common.IsNil(o.RiskdataBasketItemItemNrAmountPerItem) {
		var ret string
		return ret
	}
	return *o.RiskdataBasketItemItemNrAmountPerItem
}

// GetRiskdataBasketItemItemNrAmountPerItemOk returns a tuple with the RiskdataBasketItemItemNrAmountPerItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrAmountPerItemOk() (*string, bool) {
	if o == nil || common.IsNil(o.RiskdataBasketItemItemNrAmountPerItem) {
		return nil, false
	}
	return o.RiskdataBasketItemItemNrAmountPerItem, true
}

// HasRiskdataBasketItemItemNrAmountPerItem returns a boolean if a field has been set.
func (o *AdditionalDataRisk) HasRiskdataBasketItemItemNrAmountPerItem() bool {
	if o != nil && !common.IsNil(o.RiskdataBasketItemItemNrAmountPerItem) {
		return true
	}

	return false
}

// SetRiskdataBasketItemItemNrAmountPerItem gets a reference to the given string and assigns it to the RiskdataBasketItemItemNrAmountPerItem field.
func (o *AdditionalDataRisk) SetRiskdataBasketItemItemNrAmountPerItem(v string) {
	o.RiskdataBasketItemItemNrAmountPerItem = &v
}

// GetRiskdataBasketItemItemNrBrand returns the RiskdataBasketItemItemNrBrand field value if set, zero value otherwise.
func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrBrand() string {
	if o == nil || common.IsNil(o.RiskdataBasketItemItemNrBrand) {
		var ret string
		return ret
	}
	return *o.RiskdataBasketItemItemNrBrand
}

// GetRiskdataBasketItemItemNrBrandOk returns a tuple with the RiskdataBasketItemItemNrBrand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrBrandOk() (*string, bool) {
	if o == nil || common.IsNil(o.RiskdataBasketItemItemNrBrand) {
		return nil, false
	}
	return o.RiskdataBasketItemItemNrBrand, true
}

// HasRiskdataBasketItemItemNrBrand returns a boolean if a field has been set.
func (o *AdditionalDataRisk) HasRiskdataBasketItemItemNrBrand() bool {
	if o != nil && !common.IsNil(o.RiskdataBasketItemItemNrBrand) {
		return true
	}

	return false
}

// SetRiskdataBasketItemItemNrBrand gets a reference to the given string and assigns it to the RiskdataBasketItemItemNrBrand field.
func (o *AdditionalDataRisk) SetRiskdataBasketItemItemNrBrand(v string) {
	o.RiskdataBasketItemItemNrBrand = &v
}

// GetRiskdataBasketItemItemNrCategory returns the RiskdataBasketItemItemNrCategory field value if set, zero value otherwise.
func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrCategory() string {
	if o == nil || common.IsNil(o.RiskdataBasketItemItemNrCategory) {
		var ret string
		return ret
	}
	return *o.RiskdataBasketItemItemNrCategory
}

// GetRiskdataBasketItemItemNrCategoryOk returns a tuple with the RiskdataBasketItemItemNrCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrCategoryOk() (*string, bool) {
	if o == nil || common.IsNil(o.RiskdataBasketItemItemNrCategory) {
		return nil, false
	}
	return o.RiskdataBasketItemItemNrCategory, true
}

// HasRiskdataBasketItemItemNrCategory returns a boolean if a field has been set.
func (o *AdditionalDataRisk) HasRiskdataBasketItemItemNrCategory() bool {
	if o != nil && !common.IsNil(o.RiskdataBasketItemItemNrCategory) {
		return true
	}

	return false
}

// SetRiskdataBasketItemItemNrCategory gets a reference to the given string and assigns it to the RiskdataBasketItemItemNrCategory field.
func (o *AdditionalDataRisk) SetRiskdataBasketItemItemNrCategory(v string) {
	o.RiskdataBasketItemItemNrCategory = &v
}

// GetRiskdataBasketItemItemNrColor returns the RiskdataBasketItemItemNrColor field value if set, zero value otherwise.
func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrColor() string {
	if o == nil || common.IsNil(o.RiskdataBasketItemItemNrColor) {
		var ret string
		return ret
	}
	return *o.RiskdataBasketItemItemNrColor
}

// GetRiskdataBasketItemItemNrColorOk returns a tuple with the RiskdataBasketItemItemNrColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrColorOk() (*string, bool) {
	if o == nil || common.IsNil(o.RiskdataBasketItemItemNrColor) {
		return nil, false
	}
	return o.RiskdataBasketItemItemNrColor, true
}

// HasRiskdataBasketItemItemNrColor returns a boolean if a field has been set.
func (o *AdditionalDataRisk) HasRiskdataBasketItemItemNrColor() bool {
	if o != nil && !common.IsNil(o.RiskdataBasketItemItemNrColor) {
		return true
	}

	return false
}

// SetRiskdataBasketItemItemNrColor gets a reference to the given string and assigns it to the RiskdataBasketItemItemNrColor field.
func (o *AdditionalDataRisk) SetRiskdataBasketItemItemNrColor(v string) {
	o.RiskdataBasketItemItemNrColor = &v
}

// GetRiskdataBasketItemItemNrCurrency returns the RiskdataBasketItemItemNrCurrency field value if set, zero value otherwise.
func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrCurrency() string {
	if o == nil || common.IsNil(o.RiskdataBasketItemItemNrCurrency) {
		var ret string
		return ret
	}
	return *o.RiskdataBasketItemItemNrCurrency
}

// GetRiskdataBasketItemItemNrCurrencyOk returns a tuple with the RiskdataBasketItemItemNrCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrCurrencyOk() (*string, bool) {
	if o == nil || common.IsNil(o.RiskdataBasketItemItemNrCurrency) {
		return nil, false
	}
	return o.RiskdataBasketItemItemNrCurrency, true
}

// HasRiskdataBasketItemItemNrCurrency returns a boolean if a field has been set.
func (o *AdditionalDataRisk) HasRiskdataBasketItemItemNrCurrency() bool {
	if o != nil && !common.IsNil(o.RiskdataBasketItemItemNrCurrency) {
		return true
	}

	return false
}

// SetRiskdataBasketItemItemNrCurrency gets a reference to the given string and assigns it to the RiskdataBasketItemItemNrCurrency field.
func (o *AdditionalDataRisk) SetRiskdataBasketItemItemNrCurrency(v string) {
	o.RiskdataBasketItemItemNrCurrency = &v
}

// GetRiskdataBasketItemItemNrItemID returns the RiskdataBasketItemItemNrItemID field value if set, zero value otherwise.
func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrItemID() string {
	if o == nil || common.IsNil(o.RiskdataBasketItemItemNrItemID) {
		var ret string
		return ret
	}
	return *o.RiskdataBasketItemItemNrItemID
}

// GetRiskdataBasketItemItemNrItemIDOk returns a tuple with the RiskdataBasketItemItemNrItemID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrItemIDOk() (*string, bool) {
	if o == nil || common.IsNil(o.RiskdataBasketItemItemNrItemID) {
		return nil, false
	}
	return o.RiskdataBasketItemItemNrItemID, true
}

// HasRiskdataBasketItemItemNrItemID returns a boolean if a field has been set.
func (o *AdditionalDataRisk) HasRiskdataBasketItemItemNrItemID() bool {
	if o != nil && !common.IsNil(o.RiskdataBasketItemItemNrItemID) {
		return true
	}

	return false
}

// SetRiskdataBasketItemItemNrItemID gets a reference to the given string and assigns it to the RiskdataBasketItemItemNrItemID field.
func (o *AdditionalDataRisk) SetRiskdataBasketItemItemNrItemID(v string) {
	o.RiskdataBasketItemItemNrItemID = &v
}

// GetRiskdataBasketItemItemNrManufacturer returns the RiskdataBasketItemItemNrManufacturer field value if set, zero value otherwise.
func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrManufacturer() string {
	if o == nil || common.IsNil(o.RiskdataBasketItemItemNrManufacturer) {
		var ret string
		return ret
	}
	return *o.RiskdataBasketItemItemNrManufacturer
}

// GetRiskdataBasketItemItemNrManufacturerOk returns a tuple with the RiskdataBasketItemItemNrManufacturer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrManufacturerOk() (*string, bool) {
	if o == nil || common.IsNil(o.RiskdataBasketItemItemNrManufacturer) {
		return nil, false
	}
	return o.RiskdataBasketItemItemNrManufacturer, true
}

// HasRiskdataBasketItemItemNrManufacturer returns a boolean if a field has been set.
func (o *AdditionalDataRisk) HasRiskdataBasketItemItemNrManufacturer() bool {
	if o != nil && !common.IsNil(o.RiskdataBasketItemItemNrManufacturer) {
		return true
	}

	return false
}

// SetRiskdataBasketItemItemNrManufacturer gets a reference to the given string and assigns it to the RiskdataBasketItemItemNrManufacturer field.
func (o *AdditionalDataRisk) SetRiskdataBasketItemItemNrManufacturer(v string) {
	o.RiskdataBasketItemItemNrManufacturer = &v
}

// GetRiskdataBasketItemItemNrProductTitle returns the RiskdataBasketItemItemNrProductTitle field value if set, zero value otherwise.
func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrProductTitle() string {
	if o == nil || common.IsNil(o.RiskdataBasketItemItemNrProductTitle) {
		var ret string
		return ret
	}
	return *o.RiskdataBasketItemItemNrProductTitle
}

// GetRiskdataBasketItemItemNrProductTitleOk returns a tuple with the RiskdataBasketItemItemNrProductTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrProductTitleOk() (*string, bool) {
	if o == nil || common.IsNil(o.RiskdataBasketItemItemNrProductTitle) {
		return nil, false
	}
	return o.RiskdataBasketItemItemNrProductTitle, true
}

// HasRiskdataBasketItemItemNrProductTitle returns a boolean if a field has been set.
func (o *AdditionalDataRisk) HasRiskdataBasketItemItemNrProductTitle() bool {
	if o != nil && !common.IsNil(o.RiskdataBasketItemItemNrProductTitle) {
		return true
	}

	return false
}

// SetRiskdataBasketItemItemNrProductTitle gets a reference to the given string and assigns it to the RiskdataBasketItemItemNrProductTitle field.
func (o *AdditionalDataRisk) SetRiskdataBasketItemItemNrProductTitle(v string) {
	o.RiskdataBasketItemItemNrProductTitle = &v
}

// GetRiskdataBasketItemItemNrQuantity returns the RiskdataBasketItemItemNrQuantity field value if set, zero value otherwise.
func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrQuantity() string {
	if o == nil || common.IsNil(o.RiskdataBasketItemItemNrQuantity) {
		var ret string
		return ret
	}
	return *o.RiskdataBasketItemItemNrQuantity
}

// GetRiskdataBasketItemItemNrQuantityOk returns a tuple with the RiskdataBasketItemItemNrQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrQuantityOk() (*string, bool) {
	if o == nil || common.IsNil(o.RiskdataBasketItemItemNrQuantity) {
		return nil, false
	}
	return o.RiskdataBasketItemItemNrQuantity, true
}

// HasRiskdataBasketItemItemNrQuantity returns a boolean if a field has been set.
func (o *AdditionalDataRisk) HasRiskdataBasketItemItemNrQuantity() bool {
	if o != nil && !common.IsNil(o.RiskdataBasketItemItemNrQuantity) {
		return true
	}

	return false
}

// SetRiskdataBasketItemItemNrQuantity gets a reference to the given string and assigns it to the RiskdataBasketItemItemNrQuantity field.
func (o *AdditionalDataRisk) SetRiskdataBasketItemItemNrQuantity(v string) {
	o.RiskdataBasketItemItemNrQuantity = &v
}

// GetRiskdataBasketItemItemNrReceiverEmail returns the RiskdataBasketItemItemNrReceiverEmail field value if set, zero value otherwise.
func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrReceiverEmail() string {
	if o == nil || common.IsNil(o.RiskdataBasketItemItemNrReceiverEmail) {
		var ret string
		return ret
	}
	return *o.RiskdataBasketItemItemNrReceiverEmail
}

// GetRiskdataBasketItemItemNrReceiverEmailOk returns a tuple with the RiskdataBasketItemItemNrReceiverEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrReceiverEmailOk() (*string, bool) {
	if o == nil || common.IsNil(o.RiskdataBasketItemItemNrReceiverEmail) {
		return nil, false
	}
	return o.RiskdataBasketItemItemNrReceiverEmail, true
}

// HasRiskdataBasketItemItemNrReceiverEmail returns a boolean if a field has been set.
func (o *AdditionalDataRisk) HasRiskdataBasketItemItemNrReceiverEmail() bool {
	if o != nil && !common.IsNil(o.RiskdataBasketItemItemNrReceiverEmail) {
		return true
	}

	return false
}

// SetRiskdataBasketItemItemNrReceiverEmail gets a reference to the given string and assigns it to the RiskdataBasketItemItemNrReceiverEmail field.
func (o *AdditionalDataRisk) SetRiskdataBasketItemItemNrReceiverEmail(v string) {
	o.RiskdataBasketItemItemNrReceiverEmail = &v
}

// GetRiskdataBasketItemItemNrSize returns the RiskdataBasketItemItemNrSize field value if set, zero value otherwise.
func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrSize() string {
	if o == nil || common.IsNil(o.RiskdataBasketItemItemNrSize) {
		var ret string
		return ret
	}
	return *o.RiskdataBasketItemItemNrSize
}

// GetRiskdataBasketItemItemNrSizeOk returns a tuple with the RiskdataBasketItemItemNrSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrSizeOk() (*string, bool) {
	if o == nil || common.IsNil(o.RiskdataBasketItemItemNrSize) {
		return nil, false
	}
	return o.RiskdataBasketItemItemNrSize, true
}

// HasRiskdataBasketItemItemNrSize returns a boolean if a field has been set.
func (o *AdditionalDataRisk) HasRiskdataBasketItemItemNrSize() bool {
	if o != nil && !common.IsNil(o.RiskdataBasketItemItemNrSize) {
		return true
	}

	return false
}

// SetRiskdataBasketItemItemNrSize gets a reference to the given string and assigns it to the RiskdataBasketItemItemNrSize field.
func (o *AdditionalDataRisk) SetRiskdataBasketItemItemNrSize(v string) {
	o.RiskdataBasketItemItemNrSize = &v
}

// GetRiskdataBasketItemItemNrSku returns the RiskdataBasketItemItemNrSku field value if set, zero value otherwise.
func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrSku() string {
	if o == nil || common.IsNil(o.RiskdataBasketItemItemNrSku) {
		var ret string
		return ret
	}
	return *o.RiskdataBasketItemItemNrSku
}

// GetRiskdataBasketItemItemNrSkuOk returns a tuple with the RiskdataBasketItemItemNrSku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrSkuOk() (*string, bool) {
	if o == nil || common.IsNil(o.RiskdataBasketItemItemNrSku) {
		return nil, false
	}
	return o.RiskdataBasketItemItemNrSku, true
}

// HasRiskdataBasketItemItemNrSku returns a boolean if a field has been set.
func (o *AdditionalDataRisk) HasRiskdataBasketItemItemNrSku() bool {
	if o != nil && !common.IsNil(o.RiskdataBasketItemItemNrSku) {
		return true
	}

	return false
}

// SetRiskdataBasketItemItemNrSku gets a reference to the given string and assigns it to the RiskdataBasketItemItemNrSku field.
func (o *AdditionalDataRisk) SetRiskdataBasketItemItemNrSku(v string) {
	o.RiskdataBasketItemItemNrSku = &v
}

// GetRiskdataBasketItemItemNrUpc returns the RiskdataBasketItemItemNrUpc field value if set, zero value otherwise.
func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrUpc() string {
	if o == nil || common.IsNil(o.RiskdataBasketItemItemNrUpc) {
		var ret string
		return ret
	}
	return *o.RiskdataBasketItemItemNrUpc
}

// GetRiskdataBasketItemItemNrUpcOk returns a tuple with the RiskdataBasketItemItemNrUpc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataRisk) GetRiskdataBasketItemItemNrUpcOk() (*string, bool) {
	if o == nil || common.IsNil(o.RiskdataBasketItemItemNrUpc) {
		return nil, false
	}
	return o.RiskdataBasketItemItemNrUpc, true
}

// HasRiskdataBasketItemItemNrUpc returns a boolean if a field has been set.
func (o *AdditionalDataRisk) HasRiskdataBasketItemItemNrUpc() bool {
	if o != nil && !common.IsNil(o.RiskdataBasketItemItemNrUpc) {
		return true
	}

	return false
}

// SetRiskdataBasketItemItemNrUpc gets a reference to the given string and assigns it to the RiskdataBasketItemItemNrUpc field.
func (o *AdditionalDataRisk) SetRiskdataBasketItemItemNrUpc(v string) {
	o.RiskdataBasketItemItemNrUpc = &v
}

// GetRiskdataPromotionsPromotionItemNrPromotionCode returns the RiskdataPromotionsPromotionItemNrPromotionCode field value if set, zero value otherwise.
func (o *AdditionalDataRisk) GetRiskdataPromotionsPromotionItemNrPromotionCode() string {
	if o == nil || common.IsNil(o.RiskdataPromotionsPromotionItemNrPromotionCode) {
		var ret string
		return ret
	}
	return *o.RiskdataPromotionsPromotionItemNrPromotionCode
}

// GetRiskdataPromotionsPromotionItemNrPromotionCodeOk returns a tuple with the RiskdataPromotionsPromotionItemNrPromotionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataRisk) GetRiskdataPromotionsPromotionItemNrPromotionCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.RiskdataPromotionsPromotionItemNrPromotionCode) {
		return nil, false
	}
	return o.RiskdataPromotionsPromotionItemNrPromotionCode, true
}

// HasRiskdataPromotionsPromotionItemNrPromotionCode returns a boolean if a field has been set.
func (o *AdditionalDataRisk) HasRiskdataPromotionsPromotionItemNrPromotionCode() bool {
	if o != nil && !common.IsNil(o.RiskdataPromotionsPromotionItemNrPromotionCode) {
		return true
	}

	return false
}

// SetRiskdataPromotionsPromotionItemNrPromotionCode gets a reference to the given string and assigns it to the RiskdataPromotionsPromotionItemNrPromotionCode field.
func (o *AdditionalDataRisk) SetRiskdataPromotionsPromotionItemNrPromotionCode(v string) {
	o.RiskdataPromotionsPromotionItemNrPromotionCode = &v
}

// GetRiskdataPromotionsPromotionItemNrPromotionDiscountAmount returns the RiskdataPromotionsPromotionItemNrPromotionDiscountAmount field value if set, zero value otherwise.
func (o *AdditionalDataRisk) GetRiskdataPromotionsPromotionItemNrPromotionDiscountAmount() string {
	if o == nil || common.IsNil(o.RiskdataPromotionsPromotionItemNrPromotionDiscountAmount) {
		var ret string
		return ret
	}
	return *o.RiskdataPromotionsPromotionItemNrPromotionDiscountAmount
}

// GetRiskdataPromotionsPromotionItemNrPromotionDiscountAmountOk returns a tuple with the RiskdataPromotionsPromotionItemNrPromotionDiscountAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataRisk) GetRiskdataPromotionsPromotionItemNrPromotionDiscountAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.RiskdataPromotionsPromotionItemNrPromotionDiscountAmount) {
		return nil, false
	}
	return o.RiskdataPromotionsPromotionItemNrPromotionDiscountAmount, true
}

// HasRiskdataPromotionsPromotionItemNrPromotionDiscountAmount returns a boolean if a field has been set.
func (o *AdditionalDataRisk) HasRiskdataPromotionsPromotionItemNrPromotionDiscountAmount() bool {
	if o != nil && !common.IsNil(o.RiskdataPromotionsPromotionItemNrPromotionDiscountAmount) {
		return true
	}

	return false
}

// SetRiskdataPromotionsPromotionItemNrPromotionDiscountAmount gets a reference to the given string and assigns it to the RiskdataPromotionsPromotionItemNrPromotionDiscountAmount field.
func (o *AdditionalDataRisk) SetRiskdataPromotionsPromotionItemNrPromotionDiscountAmount(v string) {
	o.RiskdataPromotionsPromotionItemNrPromotionDiscountAmount = &v
}

// GetRiskdataPromotionsPromotionItemNrPromotionDiscountCurrency returns the RiskdataPromotionsPromotionItemNrPromotionDiscountCurrency field value if set, zero value otherwise.
func (o *AdditionalDataRisk) GetRiskdataPromotionsPromotionItemNrPromotionDiscountCurrency() string {
	if o == nil || common.IsNil(o.RiskdataPromotionsPromotionItemNrPromotionDiscountCurrency) {
		var ret string
		return ret
	}
	return *o.RiskdataPromotionsPromotionItemNrPromotionDiscountCurrency
}

// GetRiskdataPromotionsPromotionItemNrPromotionDiscountCurrencyOk returns a tuple with the RiskdataPromotionsPromotionItemNrPromotionDiscountCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataRisk) GetRiskdataPromotionsPromotionItemNrPromotionDiscountCurrencyOk() (*string, bool) {
	if o == nil || common.IsNil(o.RiskdataPromotionsPromotionItemNrPromotionDiscountCurrency) {
		return nil, false
	}
	return o.RiskdataPromotionsPromotionItemNrPromotionDiscountCurrency, true
}

// HasRiskdataPromotionsPromotionItemNrPromotionDiscountCurrency returns a boolean if a field has been set.
func (o *AdditionalDataRisk) HasRiskdataPromotionsPromotionItemNrPromotionDiscountCurrency() bool {
	if o != nil && !common.IsNil(o.RiskdataPromotionsPromotionItemNrPromotionDiscountCurrency) {
		return true
	}

	return false
}

// SetRiskdataPromotionsPromotionItemNrPromotionDiscountCurrency gets a reference to the given string and assigns it to the RiskdataPromotionsPromotionItemNrPromotionDiscountCurrency field.
func (o *AdditionalDataRisk) SetRiskdataPromotionsPromotionItemNrPromotionDiscountCurrency(v string) {
	o.RiskdataPromotionsPromotionItemNrPromotionDiscountCurrency = &v
}

// GetRiskdataPromotionsPromotionItemNrPromotionDiscountPercentage returns the RiskdataPromotionsPromotionItemNrPromotionDiscountPercentage field value if set, zero value otherwise.
func (o *AdditionalDataRisk) GetRiskdataPromotionsPromotionItemNrPromotionDiscountPercentage() string {
	if o == nil || common.IsNil(o.RiskdataPromotionsPromotionItemNrPromotionDiscountPercentage) {
		var ret string
		return ret
	}
	return *o.RiskdataPromotionsPromotionItemNrPromotionDiscountPercentage
}

// GetRiskdataPromotionsPromotionItemNrPromotionDiscountPercentageOk returns a tuple with the RiskdataPromotionsPromotionItemNrPromotionDiscountPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataRisk) GetRiskdataPromotionsPromotionItemNrPromotionDiscountPercentageOk() (*string, bool) {
	if o == nil || common.IsNil(o.RiskdataPromotionsPromotionItemNrPromotionDiscountPercentage) {
		return nil, false
	}
	return o.RiskdataPromotionsPromotionItemNrPromotionDiscountPercentage, true
}

// HasRiskdataPromotionsPromotionItemNrPromotionDiscountPercentage returns a boolean if a field has been set.
func (o *AdditionalDataRisk) HasRiskdataPromotionsPromotionItemNrPromotionDiscountPercentage() bool {
	if o != nil && !common.IsNil(o.RiskdataPromotionsPromotionItemNrPromotionDiscountPercentage) {
		return true
	}

	return false
}

// SetRiskdataPromotionsPromotionItemNrPromotionDiscountPercentage gets a reference to the given string and assigns it to the RiskdataPromotionsPromotionItemNrPromotionDiscountPercentage field.
func (o *AdditionalDataRisk) SetRiskdataPromotionsPromotionItemNrPromotionDiscountPercentage(v string) {
	o.RiskdataPromotionsPromotionItemNrPromotionDiscountPercentage = &v
}

// GetRiskdataPromotionsPromotionItemNrPromotionName returns the RiskdataPromotionsPromotionItemNrPromotionName field value if set, zero value otherwise.
func (o *AdditionalDataRisk) GetRiskdataPromotionsPromotionItemNrPromotionName() string {
	if o == nil || common.IsNil(o.RiskdataPromotionsPromotionItemNrPromotionName) {
		var ret string
		return ret
	}
	return *o.RiskdataPromotionsPromotionItemNrPromotionName
}

// GetRiskdataPromotionsPromotionItemNrPromotionNameOk returns a tuple with the RiskdataPromotionsPromotionItemNrPromotionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataRisk) GetRiskdataPromotionsPromotionItemNrPromotionNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.RiskdataPromotionsPromotionItemNrPromotionName) {
		return nil, false
	}
	return o.RiskdataPromotionsPromotionItemNrPromotionName, true
}

// HasRiskdataPromotionsPromotionItemNrPromotionName returns a boolean if a field has been set.
func (o *AdditionalDataRisk) HasRiskdataPromotionsPromotionItemNrPromotionName() bool {
	if o != nil && !common.IsNil(o.RiskdataPromotionsPromotionItemNrPromotionName) {
		return true
	}

	return false
}

// SetRiskdataPromotionsPromotionItemNrPromotionName gets a reference to the given string and assigns it to the RiskdataPromotionsPromotionItemNrPromotionName field.
func (o *AdditionalDataRisk) SetRiskdataPromotionsPromotionItemNrPromotionName(v string) {
	o.RiskdataPromotionsPromotionItemNrPromotionName = &v
}

// GetRiskdataRiskProfileReference returns the RiskdataRiskProfileReference field value if set, zero value otherwise.
func (o *AdditionalDataRisk) GetRiskdataRiskProfileReference() string {
	if o == nil || common.IsNil(o.RiskdataRiskProfileReference) {
		var ret string
		return ret
	}
	return *o.RiskdataRiskProfileReference
}

// GetRiskdataRiskProfileReferenceOk returns a tuple with the RiskdataRiskProfileReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataRisk) GetRiskdataRiskProfileReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.RiskdataRiskProfileReference) {
		return nil, false
	}
	return o.RiskdataRiskProfileReference, true
}

// HasRiskdataRiskProfileReference returns a boolean if a field has been set.
func (o *AdditionalDataRisk) HasRiskdataRiskProfileReference() bool {
	if o != nil && !common.IsNil(o.RiskdataRiskProfileReference) {
		return true
	}

	return false
}

// SetRiskdataRiskProfileReference gets a reference to the given string and assigns it to the RiskdataRiskProfileReference field.
func (o *AdditionalDataRisk) SetRiskdataRiskProfileReference(v string) {
	o.RiskdataRiskProfileReference = &v
}

// GetRiskdataSkipRisk returns the RiskdataSkipRisk field value if set, zero value otherwise.
func (o *AdditionalDataRisk) GetRiskdataSkipRisk() string {
	if o == nil || common.IsNil(o.RiskdataSkipRisk) {
		var ret string
		return ret
	}
	return *o.RiskdataSkipRisk
}

// GetRiskdataSkipRiskOk returns a tuple with the RiskdataSkipRisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataRisk) GetRiskdataSkipRiskOk() (*string, bool) {
	if o == nil || common.IsNil(o.RiskdataSkipRisk) {
		return nil, false
	}
	return o.RiskdataSkipRisk, true
}

// HasRiskdataSkipRisk returns a boolean if a field has been set.
func (o *AdditionalDataRisk) HasRiskdataSkipRisk() bool {
	if o != nil && !common.IsNil(o.RiskdataSkipRisk) {
		return true
	}

	return false
}

// SetRiskdataSkipRisk gets a reference to the given string and assigns it to the RiskdataSkipRisk field.
func (o *AdditionalDataRisk) SetRiskdataSkipRisk(v string) {
	o.RiskdataSkipRisk = &v
}

func (o AdditionalDataRisk) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdditionalDataRisk) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.RiskdataCustomFieldName) {
		toSerialize["riskdata.[customFieldName]"] = o.RiskdataCustomFieldName
	}
	if !common.IsNil(o.RiskdataBasketItemItemNrAmountPerItem) {
		toSerialize["riskdata.basket.item[itemNr].amountPerItem"] = o.RiskdataBasketItemItemNrAmountPerItem
	}
	if !common.IsNil(o.RiskdataBasketItemItemNrBrand) {
		toSerialize["riskdata.basket.item[itemNr].brand"] = o.RiskdataBasketItemItemNrBrand
	}
	if !common.IsNil(o.RiskdataBasketItemItemNrCategory) {
		toSerialize["riskdata.basket.item[itemNr].category"] = o.RiskdataBasketItemItemNrCategory
	}
	if !common.IsNil(o.RiskdataBasketItemItemNrColor) {
		toSerialize["riskdata.basket.item[itemNr].color"] = o.RiskdataBasketItemItemNrColor
	}
	if !common.IsNil(o.RiskdataBasketItemItemNrCurrency) {
		toSerialize["riskdata.basket.item[itemNr].currency"] = o.RiskdataBasketItemItemNrCurrency
	}
	if !common.IsNil(o.RiskdataBasketItemItemNrItemID) {
		toSerialize["riskdata.basket.item[itemNr].itemID"] = o.RiskdataBasketItemItemNrItemID
	}
	if !common.IsNil(o.RiskdataBasketItemItemNrManufacturer) {
		toSerialize["riskdata.basket.item[itemNr].manufacturer"] = o.RiskdataBasketItemItemNrManufacturer
	}
	if !common.IsNil(o.RiskdataBasketItemItemNrProductTitle) {
		toSerialize["riskdata.basket.item[itemNr].productTitle"] = o.RiskdataBasketItemItemNrProductTitle
	}
	if !common.IsNil(o.RiskdataBasketItemItemNrQuantity) {
		toSerialize["riskdata.basket.item[itemNr].quantity"] = o.RiskdataBasketItemItemNrQuantity
	}
	if !common.IsNil(o.RiskdataBasketItemItemNrReceiverEmail) {
		toSerialize["riskdata.basket.item[itemNr].receiverEmail"] = o.RiskdataBasketItemItemNrReceiverEmail
	}
	if !common.IsNil(o.RiskdataBasketItemItemNrSize) {
		toSerialize["riskdata.basket.item[itemNr].size"] = o.RiskdataBasketItemItemNrSize
	}
	if !common.IsNil(o.RiskdataBasketItemItemNrSku) {
		toSerialize["riskdata.basket.item[itemNr].sku"] = o.RiskdataBasketItemItemNrSku
	}
	if !common.IsNil(o.RiskdataBasketItemItemNrUpc) {
		toSerialize["riskdata.basket.item[itemNr].upc"] = o.RiskdataBasketItemItemNrUpc
	}
	if !common.IsNil(o.RiskdataPromotionsPromotionItemNrPromotionCode) {
		toSerialize["riskdata.promotions.promotion[itemNr].promotionCode"] = o.RiskdataPromotionsPromotionItemNrPromotionCode
	}
	if !common.IsNil(o.RiskdataPromotionsPromotionItemNrPromotionDiscountAmount) {
		toSerialize["riskdata.promotions.promotion[itemNr].promotionDiscountAmount"] = o.RiskdataPromotionsPromotionItemNrPromotionDiscountAmount
	}
	if !common.IsNil(o.RiskdataPromotionsPromotionItemNrPromotionDiscountCurrency) {
		toSerialize["riskdata.promotions.promotion[itemNr].promotionDiscountCurrency"] = o.RiskdataPromotionsPromotionItemNrPromotionDiscountCurrency
	}
	if !common.IsNil(o.RiskdataPromotionsPromotionItemNrPromotionDiscountPercentage) {
		toSerialize["riskdata.promotions.promotion[itemNr].promotionDiscountPercentage"] = o.RiskdataPromotionsPromotionItemNrPromotionDiscountPercentage
	}
	if !common.IsNil(o.RiskdataPromotionsPromotionItemNrPromotionName) {
		toSerialize["riskdata.promotions.promotion[itemNr].promotionName"] = o.RiskdataPromotionsPromotionItemNrPromotionName
	}
	if !common.IsNil(o.RiskdataRiskProfileReference) {
		toSerialize["riskdata.riskProfileReference"] = o.RiskdataRiskProfileReference
	}
	if !common.IsNil(o.RiskdataSkipRisk) {
		toSerialize["riskdata.skipRisk"] = o.RiskdataSkipRisk
	}
	return toSerialize, nil
}

type NullableAdditionalDataRisk struct {
	value *AdditionalDataRisk
	isSet bool
}

func (v NullableAdditionalDataRisk) Get() *AdditionalDataRisk {
	return v.value
}

func (v *NullableAdditionalDataRisk) Set(val *AdditionalDataRisk) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalDataRisk) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalDataRisk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalDataRisk(val *AdditionalDataRisk) *NullableAdditionalDataRisk {
	return &NullableAdditionalDataRisk{value: val, isSet: true}
}

func (v NullableAdditionalDataRisk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalDataRisk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
