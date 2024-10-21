/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v14/src/common"
)

// checks if the CardDetails type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CardDetails{}

// CardDetails struct for CardDetails
type CardDetails struct {
	// Secondary brand of the card. For example: **plastix**, **hmclub**.
	Brand *string `json:"brand,omitempty"`
	// The checkout attempt identifier.
	CheckoutAttemptId *string `json:"checkoutAttemptId,omitempty"`
	// Deprecated
	CupsecureplusSmscode *string `json:"cupsecureplus.smscode,omitempty"`
	// The card verification code. Only collect raw card data if you are [fully PCI compliant](https://docs.adyen.com/development-resources/pci-dss-compliance-guide).
	Cvc *string `json:"cvc,omitempty"`
	// The encrypted card number.
	EncryptedCardNumber *string `json:"encryptedCardNumber,omitempty"`
	// The encrypted card expiry month.
	EncryptedExpiryMonth *string `json:"encryptedExpiryMonth,omitempty"`
	// The encrypted card expiry year.
	EncryptedExpiryYear *string `json:"encryptedExpiryYear,omitempty"`
	// The encrypted card verification code.
	EncryptedSecurityCode *string `json:"encryptedSecurityCode,omitempty"`
	// The card expiry month. Only collect raw card data if you are [fully PCI compliant](https://docs.adyen.com/development-resources/pci-dss-compliance-guide).
	ExpiryMonth *string `json:"expiryMonth,omitempty"`
	// The card expiry year. Only collect raw card data if you are [fully PCI compliant](https://docs.adyen.com/development-resources/pci-dss-compliance-guide).
	ExpiryYear *string `json:"expiryYear,omitempty"`
	// The funding source that should be used when multiple sources are available. For Brazilian combo cards, by default the funding source is credit. To use debit, set this value to **debit**.
	FundingSource *string `json:"fundingSource,omitempty"`
	// The name of the card holder.
	HolderName *string `json:"holderName,omitempty"`
	// The transaction identifier from card schemes. This is the [`networkTxReference`](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments__resParam_additionalData-ResponseAdditionalDataCommon-networkTxReference) from the response to the first payment.
	NetworkPaymentReference *string `json:"networkPaymentReference,omitempty"`
	// The card number. Only collect raw card data if you are [fully PCI compliant](https://docs.adyen.com/development-resources/pci-dss-compliance-guide).
	Number *string `json:"number,omitempty"`
	// This is the `recurringDetailReference` returned in the response when you created the token.
	// Deprecated since Adyen Checkout API v49
	// Use `storedPaymentMethodId` instead.
	RecurringDetailReference *string `json:"recurringDetailReference,omitempty"`
	// The `shopperNotificationReference` returned in the response when you requested to notify the shopper. Used only for recurring payments in India.
	ShopperNotificationReference *string `json:"shopperNotificationReference,omitempty"`
	// An identifier used for the Click to Pay transaction.
	SrcCorrelationId *string `json:"srcCorrelationId,omitempty"`
	// The SRC reference for the Click to Pay token.
	SrcDigitalCardId *string `json:"srcDigitalCardId,omitempty"`
	// The scheme that is being used for Click to Pay.
	SrcScheme *string `json:"srcScheme,omitempty"`
	// The reference for the Click to Pay token.
	SrcTokenReference *string `json:"srcTokenReference,omitempty"`
	// This is the `recurringDetailReference` returned in the response when you created the token.
	StoredPaymentMethodId *string `json:"storedPaymentMethodId,omitempty"`
	// Required for mobile integrations. Version of the 3D Secure 2 mobile SDK.
	ThreeDS2SdkVersion *string `json:"threeDS2SdkVersion,omitempty"`
	// Default payment method details. Common for scheme payment methods, and for simple payment method details.
	Type *string `json:"type,omitempty"`
}

// NewCardDetails instantiates a new CardDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardDetails() *CardDetails {
	this := CardDetails{}
	var type_ string = "scheme"
	this.Type = &type_
	return &this
}

// NewCardDetailsWithDefaults instantiates a new CardDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardDetailsWithDefaults() *CardDetails {
	this := CardDetails{}
	var type_ string = "scheme"
	this.Type = &type_
	return &this
}

// GetBrand returns the Brand field value if set, zero value otherwise.
func (o *CardDetails) GetBrand() string {
	if o == nil || common.IsNil(o.Brand) {
		var ret string
		return ret
	}
	return *o.Brand
}

// GetBrandOk returns a tuple with the Brand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardDetails) GetBrandOk() (*string, bool) {
	if o == nil || common.IsNil(o.Brand) {
		return nil, false
	}
	return o.Brand, true
}

// HasBrand returns a boolean if a field has been set.
func (o *CardDetails) HasBrand() bool {
	if o != nil && !common.IsNil(o.Brand) {
		return true
	}

	return false
}

// SetBrand gets a reference to the given string and assigns it to the Brand field.
func (o *CardDetails) SetBrand(v string) {
	o.Brand = &v
}

// GetCheckoutAttemptId returns the CheckoutAttemptId field value if set, zero value otherwise.
func (o *CardDetails) GetCheckoutAttemptId() string {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		var ret string
		return ret
	}
	return *o.CheckoutAttemptId
}

// GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardDetails) GetCheckoutAttemptIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		return nil, false
	}
	return o.CheckoutAttemptId, true
}

// HasCheckoutAttemptId returns a boolean if a field has been set.
func (o *CardDetails) HasCheckoutAttemptId() bool {
	if o != nil && !common.IsNil(o.CheckoutAttemptId) {
		return true
	}

	return false
}

// SetCheckoutAttemptId gets a reference to the given string and assigns it to the CheckoutAttemptId field.
func (o *CardDetails) SetCheckoutAttemptId(v string) {
	o.CheckoutAttemptId = &v
}

// GetCupsecureplusSmscode returns the CupsecureplusSmscode field value if set, zero value otherwise.
// Deprecated
func (o *CardDetails) GetCupsecureplusSmscode() string {
	if o == nil || common.IsNil(o.CupsecureplusSmscode) {
		var ret string
		return ret
	}
	return *o.CupsecureplusSmscode
}

// GetCupsecureplusSmscodeOk returns a tuple with the CupsecureplusSmscode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CardDetails) GetCupsecureplusSmscodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.CupsecureplusSmscode) {
		return nil, false
	}
	return o.CupsecureplusSmscode, true
}

// HasCupsecureplusSmscode returns a boolean if a field has been set.
func (o *CardDetails) HasCupsecureplusSmscode() bool {
	if o != nil && !common.IsNil(o.CupsecureplusSmscode) {
		return true
	}

	return false
}

// SetCupsecureplusSmscode gets a reference to the given string and assigns it to the CupsecureplusSmscode field.
// Deprecated
func (o *CardDetails) SetCupsecureplusSmscode(v string) {
	o.CupsecureplusSmscode = &v
}

// GetCvc returns the Cvc field value if set, zero value otherwise.
func (o *CardDetails) GetCvc() string {
	if o == nil || common.IsNil(o.Cvc) {
		var ret string
		return ret
	}
	return *o.Cvc
}

// GetCvcOk returns a tuple with the Cvc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardDetails) GetCvcOk() (*string, bool) {
	if o == nil || common.IsNil(o.Cvc) {
		return nil, false
	}
	return o.Cvc, true
}

// HasCvc returns a boolean if a field has been set.
func (o *CardDetails) HasCvc() bool {
	if o != nil && !common.IsNil(o.Cvc) {
		return true
	}

	return false
}

// SetCvc gets a reference to the given string and assigns it to the Cvc field.
func (o *CardDetails) SetCvc(v string) {
	o.Cvc = &v
}

// GetEncryptedCardNumber returns the EncryptedCardNumber field value if set, zero value otherwise.
func (o *CardDetails) GetEncryptedCardNumber() string {
	if o == nil || common.IsNil(o.EncryptedCardNumber) {
		var ret string
		return ret
	}
	return *o.EncryptedCardNumber
}

// GetEncryptedCardNumberOk returns a tuple with the EncryptedCardNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardDetails) GetEncryptedCardNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.EncryptedCardNumber) {
		return nil, false
	}
	return o.EncryptedCardNumber, true
}

// HasEncryptedCardNumber returns a boolean if a field has been set.
func (o *CardDetails) HasEncryptedCardNumber() bool {
	if o != nil && !common.IsNil(o.EncryptedCardNumber) {
		return true
	}

	return false
}

// SetEncryptedCardNumber gets a reference to the given string and assigns it to the EncryptedCardNumber field.
func (o *CardDetails) SetEncryptedCardNumber(v string) {
	o.EncryptedCardNumber = &v
}

// GetEncryptedExpiryMonth returns the EncryptedExpiryMonth field value if set, zero value otherwise.
func (o *CardDetails) GetEncryptedExpiryMonth() string {
	if o == nil || common.IsNil(o.EncryptedExpiryMonth) {
		var ret string
		return ret
	}
	return *o.EncryptedExpiryMonth
}

// GetEncryptedExpiryMonthOk returns a tuple with the EncryptedExpiryMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardDetails) GetEncryptedExpiryMonthOk() (*string, bool) {
	if o == nil || common.IsNil(o.EncryptedExpiryMonth) {
		return nil, false
	}
	return o.EncryptedExpiryMonth, true
}

// HasEncryptedExpiryMonth returns a boolean if a field has been set.
func (o *CardDetails) HasEncryptedExpiryMonth() bool {
	if o != nil && !common.IsNil(o.EncryptedExpiryMonth) {
		return true
	}

	return false
}

// SetEncryptedExpiryMonth gets a reference to the given string and assigns it to the EncryptedExpiryMonth field.
func (o *CardDetails) SetEncryptedExpiryMonth(v string) {
	o.EncryptedExpiryMonth = &v
}

// GetEncryptedExpiryYear returns the EncryptedExpiryYear field value if set, zero value otherwise.
func (o *CardDetails) GetEncryptedExpiryYear() string {
	if o == nil || common.IsNil(o.EncryptedExpiryYear) {
		var ret string
		return ret
	}
	return *o.EncryptedExpiryYear
}

// GetEncryptedExpiryYearOk returns a tuple with the EncryptedExpiryYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardDetails) GetEncryptedExpiryYearOk() (*string, bool) {
	if o == nil || common.IsNil(o.EncryptedExpiryYear) {
		return nil, false
	}
	return o.EncryptedExpiryYear, true
}

// HasEncryptedExpiryYear returns a boolean if a field has been set.
func (o *CardDetails) HasEncryptedExpiryYear() bool {
	if o != nil && !common.IsNil(o.EncryptedExpiryYear) {
		return true
	}

	return false
}

// SetEncryptedExpiryYear gets a reference to the given string and assigns it to the EncryptedExpiryYear field.
func (o *CardDetails) SetEncryptedExpiryYear(v string) {
	o.EncryptedExpiryYear = &v
}

// GetEncryptedSecurityCode returns the EncryptedSecurityCode field value if set, zero value otherwise.
func (o *CardDetails) GetEncryptedSecurityCode() string {
	if o == nil || common.IsNil(o.EncryptedSecurityCode) {
		var ret string
		return ret
	}
	return *o.EncryptedSecurityCode
}

// GetEncryptedSecurityCodeOk returns a tuple with the EncryptedSecurityCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardDetails) GetEncryptedSecurityCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.EncryptedSecurityCode) {
		return nil, false
	}
	return o.EncryptedSecurityCode, true
}

// HasEncryptedSecurityCode returns a boolean if a field has been set.
func (o *CardDetails) HasEncryptedSecurityCode() bool {
	if o != nil && !common.IsNil(o.EncryptedSecurityCode) {
		return true
	}

	return false
}

// SetEncryptedSecurityCode gets a reference to the given string and assigns it to the EncryptedSecurityCode field.
func (o *CardDetails) SetEncryptedSecurityCode(v string) {
	o.EncryptedSecurityCode = &v
}

// GetExpiryMonth returns the ExpiryMonth field value if set, zero value otherwise.
func (o *CardDetails) GetExpiryMonth() string {
	if o == nil || common.IsNil(o.ExpiryMonth) {
		var ret string
		return ret
	}
	return *o.ExpiryMonth
}

// GetExpiryMonthOk returns a tuple with the ExpiryMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardDetails) GetExpiryMonthOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExpiryMonth) {
		return nil, false
	}
	return o.ExpiryMonth, true
}

// HasExpiryMonth returns a boolean if a field has been set.
func (o *CardDetails) HasExpiryMonth() bool {
	if o != nil && !common.IsNil(o.ExpiryMonth) {
		return true
	}

	return false
}

// SetExpiryMonth gets a reference to the given string and assigns it to the ExpiryMonth field.
func (o *CardDetails) SetExpiryMonth(v string) {
	o.ExpiryMonth = &v
}

// GetExpiryYear returns the ExpiryYear field value if set, zero value otherwise.
func (o *CardDetails) GetExpiryYear() string {
	if o == nil || common.IsNil(o.ExpiryYear) {
		var ret string
		return ret
	}
	return *o.ExpiryYear
}

// GetExpiryYearOk returns a tuple with the ExpiryYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardDetails) GetExpiryYearOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExpiryYear) {
		return nil, false
	}
	return o.ExpiryYear, true
}

// HasExpiryYear returns a boolean if a field has been set.
func (o *CardDetails) HasExpiryYear() bool {
	if o != nil && !common.IsNil(o.ExpiryYear) {
		return true
	}

	return false
}

// SetExpiryYear gets a reference to the given string and assigns it to the ExpiryYear field.
func (o *CardDetails) SetExpiryYear(v string) {
	o.ExpiryYear = &v
}

// GetFundingSource returns the FundingSource field value if set, zero value otherwise.
func (o *CardDetails) GetFundingSource() string {
	if o == nil || common.IsNil(o.FundingSource) {
		var ret string
		return ret
	}
	return *o.FundingSource
}

// GetFundingSourceOk returns a tuple with the FundingSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardDetails) GetFundingSourceOk() (*string, bool) {
	if o == nil || common.IsNil(o.FundingSource) {
		return nil, false
	}
	return o.FundingSource, true
}

// HasFundingSource returns a boolean if a field has been set.
func (o *CardDetails) HasFundingSource() bool {
	if o != nil && !common.IsNil(o.FundingSource) {
		return true
	}

	return false
}

// SetFundingSource gets a reference to the given string and assigns it to the FundingSource field.
func (o *CardDetails) SetFundingSource(v string) {
	o.FundingSource = &v
}

// GetHolderName returns the HolderName field value if set, zero value otherwise.
func (o *CardDetails) GetHolderName() string {
	if o == nil || common.IsNil(o.HolderName) {
		var ret string
		return ret
	}
	return *o.HolderName
}

// GetHolderNameOk returns a tuple with the HolderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardDetails) GetHolderNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.HolderName) {
		return nil, false
	}
	return o.HolderName, true
}

// HasHolderName returns a boolean if a field has been set.
func (o *CardDetails) HasHolderName() bool {
	if o != nil && !common.IsNil(o.HolderName) {
		return true
	}

	return false
}

// SetHolderName gets a reference to the given string and assigns it to the HolderName field.
func (o *CardDetails) SetHolderName(v string) {
	o.HolderName = &v
}

// GetNetworkPaymentReference returns the NetworkPaymentReference field value if set, zero value otherwise.
func (o *CardDetails) GetNetworkPaymentReference() string {
	if o == nil || common.IsNil(o.NetworkPaymentReference) {
		var ret string
		return ret
	}
	return *o.NetworkPaymentReference
}

// GetNetworkPaymentReferenceOk returns a tuple with the NetworkPaymentReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardDetails) GetNetworkPaymentReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.NetworkPaymentReference) {
		return nil, false
	}
	return o.NetworkPaymentReference, true
}

// HasNetworkPaymentReference returns a boolean if a field has been set.
func (o *CardDetails) HasNetworkPaymentReference() bool {
	if o != nil && !common.IsNil(o.NetworkPaymentReference) {
		return true
	}

	return false
}

// SetNetworkPaymentReference gets a reference to the given string and assigns it to the NetworkPaymentReference field.
func (o *CardDetails) SetNetworkPaymentReference(v string) {
	o.NetworkPaymentReference = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *CardDetails) GetNumber() string {
	if o == nil || common.IsNil(o.Number) {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardDetails) GetNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *CardDetails) HasNumber() bool {
	if o != nil && !common.IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *CardDetails) SetNumber(v string) {
	o.Number = &v
}

// GetRecurringDetailReference returns the RecurringDetailReference field value if set, zero value otherwise.
// Deprecated since Adyen Checkout API v49
// Use `storedPaymentMethodId` instead.
func (o *CardDetails) GetRecurringDetailReference() string {
	if o == nil || common.IsNil(o.RecurringDetailReference) {
		var ret string
		return ret
	}
	return *o.RecurringDetailReference
}

// GetRecurringDetailReferenceOk returns a tuple with the RecurringDetailReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated since Adyen Checkout API v49
// Use `storedPaymentMethodId` instead.
func (o *CardDetails) GetRecurringDetailReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.RecurringDetailReference) {
		return nil, false
	}
	return o.RecurringDetailReference, true
}

// HasRecurringDetailReference returns a boolean if a field has been set.
func (o *CardDetails) HasRecurringDetailReference() bool {
	if o != nil && !common.IsNil(o.RecurringDetailReference) {
		return true
	}

	return false
}

// SetRecurringDetailReference gets a reference to the given string and assigns it to the RecurringDetailReference field.
// Deprecated since Adyen Checkout API v49
// Use `storedPaymentMethodId` instead.
func (o *CardDetails) SetRecurringDetailReference(v string) {
	o.RecurringDetailReference = &v
}

// GetShopperNotificationReference returns the ShopperNotificationReference field value if set, zero value otherwise.
func (o *CardDetails) GetShopperNotificationReference() string {
	if o == nil || common.IsNil(o.ShopperNotificationReference) {
		var ret string
		return ret
	}
	return *o.ShopperNotificationReference
}

// GetShopperNotificationReferenceOk returns a tuple with the ShopperNotificationReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardDetails) GetShopperNotificationReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.ShopperNotificationReference) {
		return nil, false
	}
	return o.ShopperNotificationReference, true
}

// HasShopperNotificationReference returns a boolean if a field has been set.
func (o *CardDetails) HasShopperNotificationReference() bool {
	if o != nil && !common.IsNil(o.ShopperNotificationReference) {
		return true
	}

	return false
}

// SetShopperNotificationReference gets a reference to the given string and assigns it to the ShopperNotificationReference field.
func (o *CardDetails) SetShopperNotificationReference(v string) {
	o.ShopperNotificationReference = &v
}

// GetSrcCorrelationId returns the SrcCorrelationId field value if set, zero value otherwise.
func (o *CardDetails) GetSrcCorrelationId() string {
	if o == nil || common.IsNil(o.SrcCorrelationId) {
		var ret string
		return ret
	}
	return *o.SrcCorrelationId
}

// GetSrcCorrelationIdOk returns a tuple with the SrcCorrelationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardDetails) GetSrcCorrelationIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.SrcCorrelationId) {
		return nil, false
	}
	return o.SrcCorrelationId, true
}

// HasSrcCorrelationId returns a boolean if a field has been set.
func (o *CardDetails) HasSrcCorrelationId() bool {
	if o != nil && !common.IsNil(o.SrcCorrelationId) {
		return true
	}

	return false
}

// SetSrcCorrelationId gets a reference to the given string and assigns it to the SrcCorrelationId field.
func (o *CardDetails) SetSrcCorrelationId(v string) {
	o.SrcCorrelationId = &v
}

// GetSrcDigitalCardId returns the SrcDigitalCardId field value if set, zero value otherwise.
func (o *CardDetails) GetSrcDigitalCardId() string {
	if o == nil || common.IsNil(o.SrcDigitalCardId) {
		var ret string
		return ret
	}
	return *o.SrcDigitalCardId
}

// GetSrcDigitalCardIdOk returns a tuple with the SrcDigitalCardId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardDetails) GetSrcDigitalCardIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.SrcDigitalCardId) {
		return nil, false
	}
	return o.SrcDigitalCardId, true
}

// HasSrcDigitalCardId returns a boolean if a field has been set.
func (o *CardDetails) HasSrcDigitalCardId() bool {
	if o != nil && !common.IsNil(o.SrcDigitalCardId) {
		return true
	}

	return false
}

// SetSrcDigitalCardId gets a reference to the given string and assigns it to the SrcDigitalCardId field.
func (o *CardDetails) SetSrcDigitalCardId(v string) {
	o.SrcDigitalCardId = &v
}

// GetSrcScheme returns the SrcScheme field value if set, zero value otherwise.
func (o *CardDetails) GetSrcScheme() string {
	if o == nil || common.IsNil(o.SrcScheme) {
		var ret string
		return ret
	}
	return *o.SrcScheme
}

// GetSrcSchemeOk returns a tuple with the SrcScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardDetails) GetSrcSchemeOk() (*string, bool) {
	if o == nil || common.IsNil(o.SrcScheme) {
		return nil, false
	}
	return o.SrcScheme, true
}

// HasSrcScheme returns a boolean if a field has been set.
func (o *CardDetails) HasSrcScheme() bool {
	if o != nil && !common.IsNil(o.SrcScheme) {
		return true
	}

	return false
}

// SetSrcScheme gets a reference to the given string and assigns it to the SrcScheme field.
func (o *CardDetails) SetSrcScheme(v string) {
	o.SrcScheme = &v
}

// GetSrcTokenReference returns the SrcTokenReference field value if set, zero value otherwise.
func (o *CardDetails) GetSrcTokenReference() string {
	if o == nil || common.IsNil(o.SrcTokenReference) {
		var ret string
		return ret
	}
	return *o.SrcTokenReference
}

// GetSrcTokenReferenceOk returns a tuple with the SrcTokenReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardDetails) GetSrcTokenReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.SrcTokenReference) {
		return nil, false
	}
	return o.SrcTokenReference, true
}

// HasSrcTokenReference returns a boolean if a field has been set.
func (o *CardDetails) HasSrcTokenReference() bool {
	if o != nil && !common.IsNil(o.SrcTokenReference) {
		return true
	}

	return false
}

// SetSrcTokenReference gets a reference to the given string and assigns it to the SrcTokenReference field.
func (o *CardDetails) SetSrcTokenReference(v string) {
	o.SrcTokenReference = &v
}

// GetStoredPaymentMethodId returns the StoredPaymentMethodId field value if set, zero value otherwise.
func (o *CardDetails) GetStoredPaymentMethodId() string {
	if o == nil || common.IsNil(o.StoredPaymentMethodId) {
		var ret string
		return ret
	}
	return *o.StoredPaymentMethodId
}

// GetStoredPaymentMethodIdOk returns a tuple with the StoredPaymentMethodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardDetails) GetStoredPaymentMethodIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.StoredPaymentMethodId) {
		return nil, false
	}
	return o.StoredPaymentMethodId, true
}

// HasStoredPaymentMethodId returns a boolean if a field has been set.
func (o *CardDetails) HasStoredPaymentMethodId() bool {
	if o != nil && !common.IsNil(o.StoredPaymentMethodId) {
		return true
	}

	return false
}

// SetStoredPaymentMethodId gets a reference to the given string and assigns it to the StoredPaymentMethodId field.
func (o *CardDetails) SetStoredPaymentMethodId(v string) {
	o.StoredPaymentMethodId = &v
}

// GetThreeDS2SdkVersion returns the ThreeDS2SdkVersion field value if set, zero value otherwise.
func (o *CardDetails) GetThreeDS2SdkVersion() string {
	if o == nil || common.IsNil(o.ThreeDS2SdkVersion) {
		var ret string
		return ret
	}
	return *o.ThreeDS2SdkVersion
}

// GetThreeDS2SdkVersionOk returns a tuple with the ThreeDS2SdkVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardDetails) GetThreeDS2SdkVersionOk() (*string, bool) {
	if o == nil || common.IsNil(o.ThreeDS2SdkVersion) {
		return nil, false
	}
	return o.ThreeDS2SdkVersion, true
}

// HasThreeDS2SdkVersion returns a boolean if a field has been set.
func (o *CardDetails) HasThreeDS2SdkVersion() bool {
	if o != nil && !common.IsNil(o.ThreeDS2SdkVersion) {
		return true
	}

	return false
}

// SetThreeDS2SdkVersion gets a reference to the given string and assigns it to the ThreeDS2SdkVersion field.
func (o *CardDetails) SetThreeDS2SdkVersion(v string) {
	o.ThreeDS2SdkVersion = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CardDetails) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardDetails) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CardDetails) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CardDetails) SetType(v string) {
	o.Type = &v
}

func (o CardDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Brand) {
		toSerialize["brand"] = o.Brand
	}
	if !common.IsNil(o.CheckoutAttemptId) {
		toSerialize["checkoutAttemptId"] = o.CheckoutAttemptId
	}
	if !common.IsNil(o.CupsecureplusSmscode) {
		toSerialize["cupsecureplus.smscode"] = o.CupsecureplusSmscode
	}
	if !common.IsNil(o.Cvc) {
		toSerialize["cvc"] = o.Cvc
	}
	if !common.IsNil(o.EncryptedCardNumber) {
		toSerialize["encryptedCardNumber"] = o.EncryptedCardNumber
	}
	if !common.IsNil(o.EncryptedExpiryMonth) {
		toSerialize["encryptedExpiryMonth"] = o.EncryptedExpiryMonth
	}
	if !common.IsNil(o.EncryptedExpiryYear) {
		toSerialize["encryptedExpiryYear"] = o.EncryptedExpiryYear
	}
	if !common.IsNil(o.EncryptedSecurityCode) {
		toSerialize["encryptedSecurityCode"] = o.EncryptedSecurityCode
	}
	if !common.IsNil(o.ExpiryMonth) {
		toSerialize["expiryMonth"] = o.ExpiryMonth
	}
	if !common.IsNil(o.ExpiryYear) {
		toSerialize["expiryYear"] = o.ExpiryYear
	}
	if !common.IsNil(o.FundingSource) {
		toSerialize["fundingSource"] = o.FundingSource
	}
	if !common.IsNil(o.HolderName) {
		toSerialize["holderName"] = o.HolderName
	}
	if !common.IsNil(o.NetworkPaymentReference) {
		toSerialize["networkPaymentReference"] = o.NetworkPaymentReference
	}
	if !common.IsNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	if !common.IsNil(o.RecurringDetailReference) {
		toSerialize["recurringDetailReference"] = o.RecurringDetailReference
	}
	if !common.IsNil(o.ShopperNotificationReference) {
		toSerialize["shopperNotificationReference"] = o.ShopperNotificationReference
	}
	if !common.IsNil(o.SrcCorrelationId) {
		toSerialize["srcCorrelationId"] = o.SrcCorrelationId
	}
	if !common.IsNil(o.SrcDigitalCardId) {
		toSerialize["srcDigitalCardId"] = o.SrcDigitalCardId
	}
	if !common.IsNil(o.SrcScheme) {
		toSerialize["srcScheme"] = o.SrcScheme
	}
	if !common.IsNil(o.SrcTokenReference) {
		toSerialize["srcTokenReference"] = o.SrcTokenReference
	}
	if !common.IsNil(o.StoredPaymentMethodId) {
		toSerialize["storedPaymentMethodId"] = o.StoredPaymentMethodId
	}
	if !common.IsNil(o.ThreeDS2SdkVersion) {
		toSerialize["threeDS2SdkVersion"] = o.ThreeDS2SdkVersion
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableCardDetails struct {
	value *CardDetails
	isSet bool
}

func (v NullableCardDetails) Get() *CardDetails {
	return v.value
}

func (v *NullableCardDetails) Set(val *CardDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableCardDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableCardDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardDetails(val *CardDetails) *NullableCardDetails {
	return &NullableCardDetails{value: val, isSet: true}
}

func (v NullableCardDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *CardDetails) isValidFundingSource() bool {
	var allowedEnumValues = []string{"credit", "debit"}
	for _, allowed := range allowedEnumValues {
		if o.GetFundingSource() == allowed {
			return true
		}
	}
	return false
}
func (o *CardDetails) isValidType() bool {
	var allowedEnumValues = []string{"bcmc", "scheme", "networkToken", "giftcard", "card"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
