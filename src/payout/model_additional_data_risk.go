/*
 * Adyen Payout API
 *
 * A set of API endpoints that allow you to store payout details, confirm, or decline a payout.  For more information, refer to [Online payouts](https://docs.adyen.com/checkout/online-payouts).
 *
 * API version: 52
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package payout

// AdditionalDataRisk struct for AdditionalDataRisk
type AdditionalDataRisk struct {
	// The data for your custom risk field. For more information, refer to [Create custom risk fields](https://docs.adyen.com/risk-management/configure-custom-risk-rules#step-1-create-custom-risk-fields).
	RiskdataCustomFieldName string `json:"riskdata.[customFieldName],omitempty"`
	// ID of the item.
	RiskdataBasketItemItemNrItemID string `json:"riskdata.basket.item[itemNr].itemID,omitempty"`
	// A text description of the product the invoice line refers to.
	RiskdataBasketItemItemNrProductTitle string `json:"riskdata.basket.item[itemNr].productTitle,omitempty"`
	// The price of item in the basket, represented in [minor units](https://docs.adyen.com/development-resources/currency-codes).
	RiskdataBasketItemItemNrAmountPerItem string `json:"riskdata.basket.item[itemNr].amountPerItem,omitempty"`
	// The three-character [ISO currency code](https://en.wikipedia.org/wiki/ISO_4217).
	RiskdataBasketItemItemNrCurrency string `json:"riskdata.basket.item[itemNr].currency,omitempty"`
	// [Universal Product Code](https://en.wikipedia.org/wiki/Universal_Product_Code).
	RiskdataBasketItemItemNrUpc string `json:"riskdata.basket.item[itemNr].upc,omitempty"`
	// [Stock keeping unit](https://en.wikipedia.org/wiki/Stock_keeping_unit).
	RiskdataBasketItemItemNrSku string `json:"riskdata.basket.item[itemNr].sku,omitempty"`
	// Brand of the item.
	RiskdataBasketItemItemNrBrand string `json:"riskdata.basket.item[itemNr].brand,omitempty"`
	// Manufacturer of the item.
	RiskdataBasketItemItemNrManufacturer string `json:"riskdata.basket.item[itemNr].manufacturer,omitempty"`
	// Category of the item.
	RiskdataBasketItemItemNrCategory string `json:"riskdata.basket.item[itemNr].category,omitempty"`
	// Color of the item.
	RiskdataBasketItemItemNrColor string `json:"riskdata.basket.item[itemNr].color,omitempty"`
	// Size of the item.
	RiskdataBasketItemItemNrSize string `json:"riskdata.basket.item[itemNr].size,omitempty"`
	// Quantity of the item purchased.
	RiskdataBasketItemItemNrQuantity string `json:"riskdata.basket.item[itemNr].quantity,omitempty"`
	// Email associated with the given product in the basket (usually in electronic gift cards).
	RiskdataBasketItemItemNrReceiverEmail string `json:"riskdata.basket.item[itemNr].receiverEmail,omitempty"`
	// Code of the promotion.
	RiskdataPromotionsPromotionItemNrPromotionCode string `json:"riskdata.promotions.promotion[itemNr].promotionCode,omitempty"`
	// Name of the promotion.
	RiskdataPromotionsPromotionItemNrPromotionName string `json:"riskdata.promotions.promotion[itemNr].promotionName,omitempty"`
	// The discount amount of the promotion, represented in [minor units](https://docs.adyen.com/development-resources/currency-codes).
	RiskdataPromotionsPromotionItemNrPromotionDiscountAmount string `json:"riskdata.promotions.promotion[itemNr].promotionDiscountAmount,omitempty"`
	// The three-character [ISO currency code](https://en.wikipedia.org/wiki/ISO_4217).
	RiskdataPromotionsPromotionItemNrPromotionDiscountCurrency string `json:"riskdata.promotions.promotion[itemNr].promotionDiscountCurrency,omitempty"`
	// Promotion's percentage discount. It is represented in percentage value and there is no need to include the '%' sign.  e.g. for a promotion discount of 30%, the value of the field should be 30.
	RiskdataPromotionsPromotionItemNrPromotionDiscountPercentage string `json:"riskdata.promotions.promotion[itemNr].promotionDiscountPercentage,omitempty"`
}
