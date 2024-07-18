/*
Adyen Payment API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package payments

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v11/src/common"
)

// checks if the AdditionalDataOpenInvoice type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AdditionalDataOpenInvoice{}

// AdditionalDataOpenInvoice struct for AdditionalDataOpenInvoice
type AdditionalDataOpenInvoice struct {
	// Holds different merchant data points like product, purchase, customer, and so on. It takes data in a Base64 encoded string.  The `merchantData` parameter needs to be added to the `openinvoicedata` signature at the end.  Since the field is optional, if it's not included it does not impact computing the merchant signature.  Applies only to Klarna.  You can contact Klarna for the format and structure of the string.
	OpeninvoicedataMerchantData *string `json:"openinvoicedata.merchantData,omitempty"`
	// The number of invoice lines included in `openinvoicedata`.  There needs to be at least one line, so `numberOfLines` needs to be at least 1.
	OpeninvoicedataNumberOfLines *string `json:"openinvoicedata.numberOfLines,omitempty"`
	// First name of the recipient. If the delivery address and the billing address are different, specify the `recipientFirstName` and `recipientLastName` to share the delivery address with Klarna. Otherwise, only the billing address is shared with Klarna.
	OpeninvoicedataRecipientFirstName *string `json:"openinvoicedata.recipientFirstName,omitempty"`
	// Last name of the recipient. If the delivery address and the billing address are different, specify the `recipientFirstName` and `recipientLastName` to share the delivery address with Klarna. Otherwise, only the billing address is shared with Klarna.
	OpeninvoicedataRecipientLastName *string `json:"openinvoicedata.recipientLastName,omitempty"`
	// The three-character ISO currency code.
	OpeninvoicedataLineItemNrCurrencyCode *string `json:"openinvoicedataLine[itemNr].currencyCode,omitempty"`
	// A text description of the product the invoice line refers to.
	OpeninvoicedataLineItemNrDescription *string `json:"openinvoicedataLine[itemNr].description,omitempty"`
	// The price for one item in the invoice line, represented in minor units.  The due amount for the item, VAT excluded.
	OpeninvoicedataLineItemNrItemAmount *string `json:"openinvoicedataLine[itemNr].itemAmount,omitempty"`
	// A unique id for this item. Required for RatePay if the description of each item is not unique.
	OpeninvoicedataLineItemNrItemId *string `json:"openinvoicedataLine[itemNr].itemId,omitempty"`
	// The VAT due for one item in the invoice line, represented in minor units.
	OpeninvoicedataLineItemNrItemVatAmount *string `json:"openinvoicedataLine[itemNr].itemVatAmount,omitempty"`
	// The VAT percentage for one item in the invoice line, represented in minor units.  For example, 19% VAT is specified as 1900.
	OpeninvoicedataLineItemNrItemVatPercentage *string `json:"openinvoicedataLine[itemNr].itemVatPercentage,omitempty"`
	// The number of units purchased of a specific product.
	OpeninvoicedataLineItemNrNumberOfItems *string `json:"openinvoicedataLine[itemNr].numberOfItems,omitempty"`
	// Name of the shipping company handling the the return shipment.
	OpeninvoicedataLineItemNrReturnShippingCompany *string `json:"openinvoicedataLine[itemNr].returnShippingCompany,omitempty"`
	// The tracking number for the return of the shipment.
	OpeninvoicedataLineItemNrReturnTrackingNumber *string `json:"openinvoicedataLine[itemNr].returnTrackingNumber,omitempty"`
	// URI where the customer can track the return of their shipment.
	OpeninvoicedataLineItemNrReturnTrackingUri *string `json:"openinvoicedataLine[itemNr].returnTrackingUri,omitempty"`
	// Name of the shipping company handling the delivery.
	OpeninvoicedataLineItemNrShippingCompany *string `json:"openinvoicedataLine[itemNr].shippingCompany,omitempty"`
	// Shipping method.
	OpeninvoicedataLineItemNrShippingMethod *string `json:"openinvoicedataLine[itemNr].shippingMethod,omitempty"`
	// The tracking number for the shipment.
	OpeninvoicedataLineItemNrTrackingNumber *string `json:"openinvoicedataLine[itemNr].trackingNumber,omitempty"`
	// URI where the customer can track their shipment.
	OpeninvoicedataLineItemNrTrackingUri *string `json:"openinvoicedataLine[itemNr].trackingUri,omitempty"`
}

// NewAdditionalDataOpenInvoice instantiates a new AdditionalDataOpenInvoice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalDataOpenInvoice() *AdditionalDataOpenInvoice {
	this := AdditionalDataOpenInvoice{}
	return &this
}

// NewAdditionalDataOpenInvoiceWithDefaults instantiates a new AdditionalDataOpenInvoice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalDataOpenInvoiceWithDefaults() *AdditionalDataOpenInvoice {
	this := AdditionalDataOpenInvoice{}
	return &this
}

// GetOpeninvoicedataMerchantData returns the OpeninvoicedataMerchantData field value if set, zero value otherwise.
func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataMerchantData() string {
	if o == nil || common.IsNil(o.OpeninvoicedataMerchantData) {
		var ret string
		return ret
	}
	return *o.OpeninvoicedataMerchantData
}

// GetOpeninvoicedataMerchantDataOk returns a tuple with the OpeninvoicedataMerchantData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataMerchantDataOk() (*string, bool) {
	if o == nil || common.IsNil(o.OpeninvoicedataMerchantData) {
		return nil, false
	}
	return o.OpeninvoicedataMerchantData, true
}

// HasOpeninvoicedataMerchantData returns a boolean if a field has been set.
func (o *AdditionalDataOpenInvoice) HasOpeninvoicedataMerchantData() bool {
	if o != nil && !common.IsNil(o.OpeninvoicedataMerchantData) {
		return true
	}

	return false
}

// SetOpeninvoicedataMerchantData gets a reference to the given string and assigns it to the OpeninvoicedataMerchantData field.
func (o *AdditionalDataOpenInvoice) SetOpeninvoicedataMerchantData(v string) {
	o.OpeninvoicedataMerchantData = &v
}

// GetOpeninvoicedataNumberOfLines returns the OpeninvoicedataNumberOfLines field value if set, zero value otherwise.
func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataNumberOfLines() string {
	if o == nil || common.IsNil(o.OpeninvoicedataNumberOfLines) {
		var ret string
		return ret
	}
	return *o.OpeninvoicedataNumberOfLines
}

// GetOpeninvoicedataNumberOfLinesOk returns a tuple with the OpeninvoicedataNumberOfLines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataNumberOfLinesOk() (*string, bool) {
	if o == nil || common.IsNil(o.OpeninvoicedataNumberOfLines) {
		return nil, false
	}
	return o.OpeninvoicedataNumberOfLines, true
}

// HasOpeninvoicedataNumberOfLines returns a boolean if a field has been set.
func (o *AdditionalDataOpenInvoice) HasOpeninvoicedataNumberOfLines() bool {
	if o != nil && !common.IsNil(o.OpeninvoicedataNumberOfLines) {
		return true
	}

	return false
}

// SetOpeninvoicedataNumberOfLines gets a reference to the given string and assigns it to the OpeninvoicedataNumberOfLines field.
func (o *AdditionalDataOpenInvoice) SetOpeninvoicedataNumberOfLines(v string) {
	o.OpeninvoicedataNumberOfLines = &v
}

// GetOpeninvoicedataRecipientFirstName returns the OpeninvoicedataRecipientFirstName field value if set, zero value otherwise.
func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataRecipientFirstName() string {
	if o == nil || common.IsNil(o.OpeninvoicedataRecipientFirstName) {
		var ret string
		return ret
	}
	return *o.OpeninvoicedataRecipientFirstName
}

// GetOpeninvoicedataRecipientFirstNameOk returns a tuple with the OpeninvoicedataRecipientFirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataRecipientFirstNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.OpeninvoicedataRecipientFirstName) {
		return nil, false
	}
	return o.OpeninvoicedataRecipientFirstName, true
}

// HasOpeninvoicedataRecipientFirstName returns a boolean if a field has been set.
func (o *AdditionalDataOpenInvoice) HasOpeninvoicedataRecipientFirstName() bool {
	if o != nil && !common.IsNil(o.OpeninvoicedataRecipientFirstName) {
		return true
	}

	return false
}

// SetOpeninvoicedataRecipientFirstName gets a reference to the given string and assigns it to the OpeninvoicedataRecipientFirstName field.
func (o *AdditionalDataOpenInvoice) SetOpeninvoicedataRecipientFirstName(v string) {
	o.OpeninvoicedataRecipientFirstName = &v
}

// GetOpeninvoicedataRecipientLastName returns the OpeninvoicedataRecipientLastName field value if set, zero value otherwise.
func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataRecipientLastName() string {
	if o == nil || common.IsNil(o.OpeninvoicedataRecipientLastName) {
		var ret string
		return ret
	}
	return *o.OpeninvoicedataRecipientLastName
}

// GetOpeninvoicedataRecipientLastNameOk returns a tuple with the OpeninvoicedataRecipientLastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataRecipientLastNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.OpeninvoicedataRecipientLastName) {
		return nil, false
	}
	return o.OpeninvoicedataRecipientLastName, true
}

// HasOpeninvoicedataRecipientLastName returns a boolean if a field has been set.
func (o *AdditionalDataOpenInvoice) HasOpeninvoicedataRecipientLastName() bool {
	if o != nil && !common.IsNil(o.OpeninvoicedataRecipientLastName) {
		return true
	}

	return false
}

// SetOpeninvoicedataRecipientLastName gets a reference to the given string and assigns it to the OpeninvoicedataRecipientLastName field.
func (o *AdditionalDataOpenInvoice) SetOpeninvoicedataRecipientLastName(v string) {
	o.OpeninvoicedataRecipientLastName = &v
}

// GetOpeninvoicedataLineItemNrCurrencyCode returns the OpeninvoicedataLineItemNrCurrencyCode field value if set, zero value otherwise.
func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrCurrencyCode() string {
	if o == nil || common.IsNil(o.OpeninvoicedataLineItemNrCurrencyCode) {
		var ret string
		return ret
	}
	return *o.OpeninvoicedataLineItemNrCurrencyCode
}

// GetOpeninvoicedataLineItemNrCurrencyCodeOk returns a tuple with the OpeninvoicedataLineItemNrCurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrCurrencyCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.OpeninvoicedataLineItemNrCurrencyCode) {
		return nil, false
	}
	return o.OpeninvoicedataLineItemNrCurrencyCode, true
}

// HasOpeninvoicedataLineItemNrCurrencyCode returns a boolean if a field has been set.
func (o *AdditionalDataOpenInvoice) HasOpeninvoicedataLineItemNrCurrencyCode() bool {
	if o != nil && !common.IsNil(o.OpeninvoicedataLineItemNrCurrencyCode) {
		return true
	}

	return false
}

// SetOpeninvoicedataLineItemNrCurrencyCode gets a reference to the given string and assigns it to the OpeninvoicedataLineItemNrCurrencyCode field.
func (o *AdditionalDataOpenInvoice) SetOpeninvoicedataLineItemNrCurrencyCode(v string) {
	o.OpeninvoicedataLineItemNrCurrencyCode = &v
}

// GetOpeninvoicedataLineItemNrDescription returns the OpeninvoicedataLineItemNrDescription field value if set, zero value otherwise.
func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrDescription() string {
	if o == nil || common.IsNil(o.OpeninvoicedataLineItemNrDescription) {
		var ret string
		return ret
	}
	return *o.OpeninvoicedataLineItemNrDescription
}

// GetOpeninvoicedataLineItemNrDescriptionOk returns a tuple with the OpeninvoicedataLineItemNrDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.OpeninvoicedataLineItemNrDescription) {
		return nil, false
	}
	return o.OpeninvoicedataLineItemNrDescription, true
}

// HasOpeninvoicedataLineItemNrDescription returns a boolean if a field has been set.
func (o *AdditionalDataOpenInvoice) HasOpeninvoicedataLineItemNrDescription() bool {
	if o != nil && !common.IsNil(o.OpeninvoicedataLineItemNrDescription) {
		return true
	}

	return false
}

// SetOpeninvoicedataLineItemNrDescription gets a reference to the given string and assigns it to the OpeninvoicedataLineItemNrDescription field.
func (o *AdditionalDataOpenInvoice) SetOpeninvoicedataLineItemNrDescription(v string) {
	o.OpeninvoicedataLineItemNrDescription = &v
}

// GetOpeninvoicedataLineItemNrItemAmount returns the OpeninvoicedataLineItemNrItemAmount field value if set, zero value otherwise.
func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrItemAmount() string {
	if o == nil || common.IsNil(o.OpeninvoicedataLineItemNrItemAmount) {
		var ret string
		return ret
	}
	return *o.OpeninvoicedataLineItemNrItemAmount
}

// GetOpeninvoicedataLineItemNrItemAmountOk returns a tuple with the OpeninvoicedataLineItemNrItemAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrItemAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.OpeninvoicedataLineItemNrItemAmount) {
		return nil, false
	}
	return o.OpeninvoicedataLineItemNrItemAmount, true
}

// HasOpeninvoicedataLineItemNrItemAmount returns a boolean if a field has been set.
func (o *AdditionalDataOpenInvoice) HasOpeninvoicedataLineItemNrItemAmount() bool {
	if o != nil && !common.IsNil(o.OpeninvoicedataLineItemNrItemAmount) {
		return true
	}

	return false
}

// SetOpeninvoicedataLineItemNrItemAmount gets a reference to the given string and assigns it to the OpeninvoicedataLineItemNrItemAmount field.
func (o *AdditionalDataOpenInvoice) SetOpeninvoicedataLineItemNrItemAmount(v string) {
	o.OpeninvoicedataLineItemNrItemAmount = &v
}

// GetOpeninvoicedataLineItemNrItemId returns the OpeninvoicedataLineItemNrItemId field value if set, zero value otherwise.
func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrItemId() string {
	if o == nil || common.IsNil(o.OpeninvoicedataLineItemNrItemId) {
		var ret string
		return ret
	}
	return *o.OpeninvoicedataLineItemNrItemId
}

// GetOpeninvoicedataLineItemNrItemIdOk returns a tuple with the OpeninvoicedataLineItemNrItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrItemIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.OpeninvoicedataLineItemNrItemId) {
		return nil, false
	}
	return o.OpeninvoicedataLineItemNrItemId, true
}

// HasOpeninvoicedataLineItemNrItemId returns a boolean if a field has been set.
func (o *AdditionalDataOpenInvoice) HasOpeninvoicedataLineItemNrItemId() bool {
	if o != nil && !common.IsNil(o.OpeninvoicedataLineItemNrItemId) {
		return true
	}

	return false
}

// SetOpeninvoicedataLineItemNrItemId gets a reference to the given string and assigns it to the OpeninvoicedataLineItemNrItemId field.
func (o *AdditionalDataOpenInvoice) SetOpeninvoicedataLineItemNrItemId(v string) {
	o.OpeninvoicedataLineItemNrItemId = &v
}

// GetOpeninvoicedataLineItemNrItemVatAmount returns the OpeninvoicedataLineItemNrItemVatAmount field value if set, zero value otherwise.
func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrItemVatAmount() string {
	if o == nil || common.IsNil(o.OpeninvoicedataLineItemNrItemVatAmount) {
		var ret string
		return ret
	}
	return *o.OpeninvoicedataLineItemNrItemVatAmount
}

// GetOpeninvoicedataLineItemNrItemVatAmountOk returns a tuple with the OpeninvoicedataLineItemNrItemVatAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrItemVatAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.OpeninvoicedataLineItemNrItemVatAmount) {
		return nil, false
	}
	return o.OpeninvoicedataLineItemNrItemVatAmount, true
}

// HasOpeninvoicedataLineItemNrItemVatAmount returns a boolean if a field has been set.
func (o *AdditionalDataOpenInvoice) HasOpeninvoicedataLineItemNrItemVatAmount() bool {
	if o != nil && !common.IsNil(o.OpeninvoicedataLineItemNrItemVatAmount) {
		return true
	}

	return false
}

// SetOpeninvoicedataLineItemNrItemVatAmount gets a reference to the given string and assigns it to the OpeninvoicedataLineItemNrItemVatAmount field.
func (o *AdditionalDataOpenInvoice) SetOpeninvoicedataLineItemNrItemVatAmount(v string) {
	o.OpeninvoicedataLineItemNrItemVatAmount = &v
}

// GetOpeninvoicedataLineItemNrItemVatPercentage returns the OpeninvoicedataLineItemNrItemVatPercentage field value if set, zero value otherwise.
func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrItemVatPercentage() string {
	if o == nil || common.IsNil(o.OpeninvoicedataLineItemNrItemVatPercentage) {
		var ret string
		return ret
	}
	return *o.OpeninvoicedataLineItemNrItemVatPercentage
}

// GetOpeninvoicedataLineItemNrItemVatPercentageOk returns a tuple with the OpeninvoicedataLineItemNrItemVatPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrItemVatPercentageOk() (*string, bool) {
	if o == nil || common.IsNil(o.OpeninvoicedataLineItemNrItemVatPercentage) {
		return nil, false
	}
	return o.OpeninvoicedataLineItemNrItemVatPercentage, true
}

// HasOpeninvoicedataLineItemNrItemVatPercentage returns a boolean if a field has been set.
func (o *AdditionalDataOpenInvoice) HasOpeninvoicedataLineItemNrItemVatPercentage() bool {
	if o != nil && !common.IsNil(o.OpeninvoicedataLineItemNrItemVatPercentage) {
		return true
	}

	return false
}

// SetOpeninvoicedataLineItemNrItemVatPercentage gets a reference to the given string and assigns it to the OpeninvoicedataLineItemNrItemVatPercentage field.
func (o *AdditionalDataOpenInvoice) SetOpeninvoicedataLineItemNrItemVatPercentage(v string) {
	o.OpeninvoicedataLineItemNrItemVatPercentage = &v
}

// GetOpeninvoicedataLineItemNrNumberOfItems returns the OpeninvoicedataLineItemNrNumberOfItems field value if set, zero value otherwise.
func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrNumberOfItems() string {
	if o == nil || common.IsNil(o.OpeninvoicedataLineItemNrNumberOfItems) {
		var ret string
		return ret
	}
	return *o.OpeninvoicedataLineItemNrNumberOfItems
}

// GetOpeninvoicedataLineItemNrNumberOfItemsOk returns a tuple with the OpeninvoicedataLineItemNrNumberOfItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrNumberOfItemsOk() (*string, bool) {
	if o == nil || common.IsNil(o.OpeninvoicedataLineItemNrNumberOfItems) {
		return nil, false
	}
	return o.OpeninvoicedataLineItemNrNumberOfItems, true
}

// HasOpeninvoicedataLineItemNrNumberOfItems returns a boolean if a field has been set.
func (o *AdditionalDataOpenInvoice) HasOpeninvoicedataLineItemNrNumberOfItems() bool {
	if o != nil && !common.IsNil(o.OpeninvoicedataLineItemNrNumberOfItems) {
		return true
	}

	return false
}

// SetOpeninvoicedataLineItemNrNumberOfItems gets a reference to the given string and assigns it to the OpeninvoicedataLineItemNrNumberOfItems field.
func (o *AdditionalDataOpenInvoice) SetOpeninvoicedataLineItemNrNumberOfItems(v string) {
	o.OpeninvoicedataLineItemNrNumberOfItems = &v
}

// GetOpeninvoicedataLineItemNrReturnShippingCompany returns the OpeninvoicedataLineItemNrReturnShippingCompany field value if set, zero value otherwise.
func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrReturnShippingCompany() string {
	if o == nil || common.IsNil(o.OpeninvoicedataLineItemNrReturnShippingCompany) {
		var ret string
		return ret
	}
	return *o.OpeninvoicedataLineItemNrReturnShippingCompany
}

// GetOpeninvoicedataLineItemNrReturnShippingCompanyOk returns a tuple with the OpeninvoicedataLineItemNrReturnShippingCompany field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrReturnShippingCompanyOk() (*string, bool) {
	if o == nil || common.IsNil(o.OpeninvoicedataLineItemNrReturnShippingCompany) {
		return nil, false
	}
	return o.OpeninvoicedataLineItemNrReturnShippingCompany, true
}

// HasOpeninvoicedataLineItemNrReturnShippingCompany returns a boolean if a field has been set.
func (o *AdditionalDataOpenInvoice) HasOpeninvoicedataLineItemNrReturnShippingCompany() bool {
	if o != nil && !common.IsNil(o.OpeninvoicedataLineItemNrReturnShippingCompany) {
		return true
	}

	return false
}

// SetOpeninvoicedataLineItemNrReturnShippingCompany gets a reference to the given string and assigns it to the OpeninvoicedataLineItemNrReturnShippingCompany field.
func (o *AdditionalDataOpenInvoice) SetOpeninvoicedataLineItemNrReturnShippingCompany(v string) {
	o.OpeninvoicedataLineItemNrReturnShippingCompany = &v
}

// GetOpeninvoicedataLineItemNrReturnTrackingNumber returns the OpeninvoicedataLineItemNrReturnTrackingNumber field value if set, zero value otherwise.
func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrReturnTrackingNumber() string {
	if o == nil || common.IsNil(o.OpeninvoicedataLineItemNrReturnTrackingNumber) {
		var ret string
		return ret
	}
	return *o.OpeninvoicedataLineItemNrReturnTrackingNumber
}

// GetOpeninvoicedataLineItemNrReturnTrackingNumberOk returns a tuple with the OpeninvoicedataLineItemNrReturnTrackingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrReturnTrackingNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.OpeninvoicedataLineItemNrReturnTrackingNumber) {
		return nil, false
	}
	return o.OpeninvoicedataLineItemNrReturnTrackingNumber, true
}

// HasOpeninvoicedataLineItemNrReturnTrackingNumber returns a boolean if a field has been set.
func (o *AdditionalDataOpenInvoice) HasOpeninvoicedataLineItemNrReturnTrackingNumber() bool {
	if o != nil && !common.IsNil(o.OpeninvoicedataLineItemNrReturnTrackingNumber) {
		return true
	}

	return false
}

// SetOpeninvoicedataLineItemNrReturnTrackingNumber gets a reference to the given string and assigns it to the OpeninvoicedataLineItemNrReturnTrackingNumber field.
func (o *AdditionalDataOpenInvoice) SetOpeninvoicedataLineItemNrReturnTrackingNumber(v string) {
	o.OpeninvoicedataLineItemNrReturnTrackingNumber = &v
}

// GetOpeninvoicedataLineItemNrReturnTrackingUri returns the OpeninvoicedataLineItemNrReturnTrackingUri field value if set, zero value otherwise.
func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrReturnTrackingUri() string {
	if o == nil || common.IsNil(o.OpeninvoicedataLineItemNrReturnTrackingUri) {
		var ret string
		return ret
	}
	return *o.OpeninvoicedataLineItemNrReturnTrackingUri
}

// GetOpeninvoicedataLineItemNrReturnTrackingUriOk returns a tuple with the OpeninvoicedataLineItemNrReturnTrackingUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrReturnTrackingUriOk() (*string, bool) {
	if o == nil || common.IsNil(o.OpeninvoicedataLineItemNrReturnTrackingUri) {
		return nil, false
	}
	return o.OpeninvoicedataLineItemNrReturnTrackingUri, true
}

// HasOpeninvoicedataLineItemNrReturnTrackingUri returns a boolean if a field has been set.
func (o *AdditionalDataOpenInvoice) HasOpeninvoicedataLineItemNrReturnTrackingUri() bool {
	if o != nil && !common.IsNil(o.OpeninvoicedataLineItemNrReturnTrackingUri) {
		return true
	}

	return false
}

// SetOpeninvoicedataLineItemNrReturnTrackingUri gets a reference to the given string and assigns it to the OpeninvoicedataLineItemNrReturnTrackingUri field.
func (o *AdditionalDataOpenInvoice) SetOpeninvoicedataLineItemNrReturnTrackingUri(v string) {
	o.OpeninvoicedataLineItemNrReturnTrackingUri = &v
}

// GetOpeninvoicedataLineItemNrShippingCompany returns the OpeninvoicedataLineItemNrShippingCompany field value if set, zero value otherwise.
func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrShippingCompany() string {
	if o == nil || common.IsNil(o.OpeninvoicedataLineItemNrShippingCompany) {
		var ret string
		return ret
	}
	return *o.OpeninvoicedataLineItemNrShippingCompany
}

// GetOpeninvoicedataLineItemNrShippingCompanyOk returns a tuple with the OpeninvoicedataLineItemNrShippingCompany field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrShippingCompanyOk() (*string, bool) {
	if o == nil || common.IsNil(o.OpeninvoicedataLineItemNrShippingCompany) {
		return nil, false
	}
	return o.OpeninvoicedataLineItemNrShippingCompany, true
}

// HasOpeninvoicedataLineItemNrShippingCompany returns a boolean if a field has been set.
func (o *AdditionalDataOpenInvoice) HasOpeninvoicedataLineItemNrShippingCompany() bool {
	if o != nil && !common.IsNil(o.OpeninvoicedataLineItemNrShippingCompany) {
		return true
	}

	return false
}

// SetOpeninvoicedataLineItemNrShippingCompany gets a reference to the given string and assigns it to the OpeninvoicedataLineItemNrShippingCompany field.
func (o *AdditionalDataOpenInvoice) SetOpeninvoicedataLineItemNrShippingCompany(v string) {
	o.OpeninvoicedataLineItemNrShippingCompany = &v
}

// GetOpeninvoicedataLineItemNrShippingMethod returns the OpeninvoicedataLineItemNrShippingMethod field value if set, zero value otherwise.
func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrShippingMethod() string {
	if o == nil || common.IsNil(o.OpeninvoicedataLineItemNrShippingMethod) {
		var ret string
		return ret
	}
	return *o.OpeninvoicedataLineItemNrShippingMethod
}

// GetOpeninvoicedataLineItemNrShippingMethodOk returns a tuple with the OpeninvoicedataLineItemNrShippingMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrShippingMethodOk() (*string, bool) {
	if o == nil || common.IsNil(o.OpeninvoicedataLineItemNrShippingMethod) {
		return nil, false
	}
	return o.OpeninvoicedataLineItemNrShippingMethod, true
}

// HasOpeninvoicedataLineItemNrShippingMethod returns a boolean if a field has been set.
func (o *AdditionalDataOpenInvoice) HasOpeninvoicedataLineItemNrShippingMethod() bool {
	if o != nil && !common.IsNil(o.OpeninvoicedataLineItemNrShippingMethod) {
		return true
	}

	return false
}

// SetOpeninvoicedataLineItemNrShippingMethod gets a reference to the given string and assigns it to the OpeninvoicedataLineItemNrShippingMethod field.
func (o *AdditionalDataOpenInvoice) SetOpeninvoicedataLineItemNrShippingMethod(v string) {
	o.OpeninvoicedataLineItemNrShippingMethod = &v
}

// GetOpeninvoicedataLineItemNrTrackingNumber returns the OpeninvoicedataLineItemNrTrackingNumber field value if set, zero value otherwise.
func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrTrackingNumber() string {
	if o == nil || common.IsNil(o.OpeninvoicedataLineItemNrTrackingNumber) {
		var ret string
		return ret
	}
	return *o.OpeninvoicedataLineItemNrTrackingNumber
}

// GetOpeninvoicedataLineItemNrTrackingNumberOk returns a tuple with the OpeninvoicedataLineItemNrTrackingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrTrackingNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.OpeninvoicedataLineItemNrTrackingNumber) {
		return nil, false
	}
	return o.OpeninvoicedataLineItemNrTrackingNumber, true
}

// HasOpeninvoicedataLineItemNrTrackingNumber returns a boolean if a field has been set.
func (o *AdditionalDataOpenInvoice) HasOpeninvoicedataLineItemNrTrackingNumber() bool {
	if o != nil && !common.IsNil(o.OpeninvoicedataLineItemNrTrackingNumber) {
		return true
	}

	return false
}

// SetOpeninvoicedataLineItemNrTrackingNumber gets a reference to the given string and assigns it to the OpeninvoicedataLineItemNrTrackingNumber field.
func (o *AdditionalDataOpenInvoice) SetOpeninvoicedataLineItemNrTrackingNumber(v string) {
	o.OpeninvoicedataLineItemNrTrackingNumber = &v
}

// GetOpeninvoicedataLineItemNrTrackingUri returns the OpeninvoicedataLineItemNrTrackingUri field value if set, zero value otherwise.
func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrTrackingUri() string {
	if o == nil || common.IsNil(o.OpeninvoicedataLineItemNrTrackingUri) {
		var ret string
		return ret
	}
	return *o.OpeninvoicedataLineItemNrTrackingUri
}

// GetOpeninvoicedataLineItemNrTrackingUriOk returns a tuple with the OpeninvoicedataLineItemNrTrackingUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalDataOpenInvoice) GetOpeninvoicedataLineItemNrTrackingUriOk() (*string, bool) {
	if o == nil || common.IsNil(o.OpeninvoicedataLineItemNrTrackingUri) {
		return nil, false
	}
	return o.OpeninvoicedataLineItemNrTrackingUri, true
}

// HasOpeninvoicedataLineItemNrTrackingUri returns a boolean if a field has been set.
func (o *AdditionalDataOpenInvoice) HasOpeninvoicedataLineItemNrTrackingUri() bool {
	if o != nil && !common.IsNil(o.OpeninvoicedataLineItemNrTrackingUri) {
		return true
	}

	return false
}

// SetOpeninvoicedataLineItemNrTrackingUri gets a reference to the given string and assigns it to the OpeninvoicedataLineItemNrTrackingUri field.
func (o *AdditionalDataOpenInvoice) SetOpeninvoicedataLineItemNrTrackingUri(v string) {
	o.OpeninvoicedataLineItemNrTrackingUri = &v
}

func (o AdditionalDataOpenInvoice) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdditionalDataOpenInvoice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.OpeninvoicedataMerchantData) {
		toSerialize["openinvoicedata.merchantData"] = o.OpeninvoicedataMerchantData
	}
	if !common.IsNil(o.OpeninvoicedataNumberOfLines) {
		toSerialize["openinvoicedata.numberOfLines"] = o.OpeninvoicedataNumberOfLines
	}
	if !common.IsNil(o.OpeninvoicedataRecipientFirstName) {
		toSerialize["openinvoicedata.recipientFirstName"] = o.OpeninvoicedataRecipientFirstName
	}
	if !common.IsNil(o.OpeninvoicedataRecipientLastName) {
		toSerialize["openinvoicedata.recipientLastName"] = o.OpeninvoicedataRecipientLastName
	}
	if !common.IsNil(o.OpeninvoicedataLineItemNrCurrencyCode) {
		toSerialize["openinvoicedataLine[itemNr].currencyCode"] = o.OpeninvoicedataLineItemNrCurrencyCode
	}
	if !common.IsNil(o.OpeninvoicedataLineItemNrDescription) {
		toSerialize["openinvoicedataLine[itemNr].description"] = o.OpeninvoicedataLineItemNrDescription
	}
	if !common.IsNil(o.OpeninvoicedataLineItemNrItemAmount) {
		toSerialize["openinvoicedataLine[itemNr].itemAmount"] = o.OpeninvoicedataLineItemNrItemAmount
	}
	if !common.IsNil(o.OpeninvoicedataLineItemNrItemId) {
		toSerialize["openinvoicedataLine[itemNr].itemId"] = o.OpeninvoicedataLineItemNrItemId
	}
	if !common.IsNil(o.OpeninvoicedataLineItemNrItemVatAmount) {
		toSerialize["openinvoicedataLine[itemNr].itemVatAmount"] = o.OpeninvoicedataLineItemNrItemVatAmount
	}
	if !common.IsNil(o.OpeninvoicedataLineItemNrItemVatPercentage) {
		toSerialize["openinvoicedataLine[itemNr].itemVatPercentage"] = o.OpeninvoicedataLineItemNrItemVatPercentage
	}
	if !common.IsNil(o.OpeninvoicedataLineItemNrNumberOfItems) {
		toSerialize["openinvoicedataLine[itemNr].numberOfItems"] = o.OpeninvoicedataLineItemNrNumberOfItems
	}
	if !common.IsNil(o.OpeninvoicedataLineItemNrReturnShippingCompany) {
		toSerialize["openinvoicedataLine[itemNr].returnShippingCompany"] = o.OpeninvoicedataLineItemNrReturnShippingCompany
	}
	if !common.IsNil(o.OpeninvoicedataLineItemNrReturnTrackingNumber) {
		toSerialize["openinvoicedataLine[itemNr].returnTrackingNumber"] = o.OpeninvoicedataLineItemNrReturnTrackingNumber
	}
	if !common.IsNil(o.OpeninvoicedataLineItemNrReturnTrackingUri) {
		toSerialize["openinvoicedataLine[itemNr].returnTrackingUri"] = o.OpeninvoicedataLineItemNrReturnTrackingUri
	}
	if !common.IsNil(o.OpeninvoicedataLineItemNrShippingCompany) {
		toSerialize["openinvoicedataLine[itemNr].shippingCompany"] = o.OpeninvoicedataLineItemNrShippingCompany
	}
	if !common.IsNil(o.OpeninvoicedataLineItemNrShippingMethod) {
		toSerialize["openinvoicedataLine[itemNr].shippingMethod"] = o.OpeninvoicedataLineItemNrShippingMethod
	}
	if !common.IsNil(o.OpeninvoicedataLineItemNrTrackingNumber) {
		toSerialize["openinvoicedataLine[itemNr].trackingNumber"] = o.OpeninvoicedataLineItemNrTrackingNumber
	}
	if !common.IsNil(o.OpeninvoicedataLineItemNrTrackingUri) {
		toSerialize["openinvoicedataLine[itemNr].trackingUri"] = o.OpeninvoicedataLineItemNrTrackingUri
	}
	return toSerialize, nil
}

type NullableAdditionalDataOpenInvoice struct {
	value *AdditionalDataOpenInvoice
	isSet bool
}

func (v NullableAdditionalDataOpenInvoice) Get() *AdditionalDataOpenInvoice {
	return v.value
}

func (v *NullableAdditionalDataOpenInvoice) Set(val *AdditionalDataOpenInvoice) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalDataOpenInvoice) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalDataOpenInvoice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalDataOpenInvoice(val *AdditionalDataOpenInvoice) *NullableAdditionalDataOpenInvoice {
	return &NullableAdditionalDataOpenInvoice{value: val, isSet: true}
}

func (v NullableAdditionalDataOpenInvoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalDataOpenInvoice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
