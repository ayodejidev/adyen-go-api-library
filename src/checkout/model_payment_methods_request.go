/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v11/src/common"
)

// checks if the PaymentMethodsRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PaymentMethodsRequest{}

// PaymentMethodsRequest struct for PaymentMethodsRequest
type PaymentMethodsRequest struct {
	// This field contains additional data, which may be required for a particular payment request.  The `additionalData` object consists of entries, each of which includes the key and value.
	AdditionalData *map[string]string `json:"additionalData,omitempty"`
	// List of payment methods to be presented to the shopper. To refer to payment methods, use their [payment method type](https://docs.adyen.com/payment-methods/payment-method-types).  Example: `\"allowedPaymentMethods\":[\"ideal\",\"giropay\"]`
	AllowedPaymentMethods []string `json:"allowedPaymentMethods,omitempty"`
	Amount                *Amount  `json:"amount,omitempty"`
	// List of payment methods to be hidden from the shopper. To refer to payment methods, use their [payment method type](https://docs.adyen.com/payment-methods/payment-method-types).  Example: `\"blockedPaymentMethods\":[\"ideal\",\"giropay\"]`
	BlockedPaymentMethods []string `json:"blockedPaymentMethods,omitempty"`
	// The platform where a payment transaction takes place. This field can be used for filtering out payment methods that are only available on specific platforms. Possible values: * iOS * Android * Web
	Channel *string `json:"channel,omitempty"`
	// The shopper's country code.
	CountryCode *string `json:"countryCode,omitempty"`
	// The merchant account identifier, with which you want to process the transaction.
	MerchantAccount string              `json:"merchantAccount"`
	Order           *EncryptedOrderData `json:"order,omitempty"`
	// The combination of a language code and a country code to specify the language to be used in the payment.
	ShopperLocale *string `json:"shopperLocale,omitempty"`
	// Required for recurring payments.  Your reference to uniquely identify this shopper, for example user ID or account ID. Minimum length: 3 characters. > Your reference must not include personally identifiable information (PII), for example name or email address.
	ShopperReference *string `json:"shopperReference,omitempty"`
	// Boolean value indicating whether the card payment method should be split into separate debit and credit options.
	SplitCardFundingSources *bool `json:"splitCardFundingSources,omitempty"`
	// Required for Adyen for Platforms integrations if you are a platform model. This is your [reference](https://docs.adyen.com/api-explorer/Management/3/post/merchants/(merchantId)/stores#request-reference) (on [balance platform](https://docs.adyen.com/platforms)) or the [storeReference](https://docs.adyen.com/api-explorer/Account/latest/post/updateAccountHolder#request-accountHolderDetails-storeDetails-storeReference) (in the [classic integration](https://docs.adyen.com/classic-platforms/processing-payments/route-payment-to-store/#route-a-payment-to-a-store)) for the ecommerce or point-of-sale store that is processing the payment.
	Store *string `json:"store,omitempty"`
	// Specifies how payment methods should be filtered based on the 'store' parameter:   - 'exclusive': Only payment methods belonging to the specified 'store' are returned.   - 'inclusive': Payment methods from the 'store' and those not associated with any other store are returned.   - 'skipFilter': All payment methods are returned, regardless of store association.
	StoreFiltrationMode *string `json:"storeFiltrationMode,omitempty"`
}

// NewPaymentMethodsRequest instantiates a new PaymentMethodsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodsRequest(merchantAccount string) *PaymentMethodsRequest {
	this := PaymentMethodsRequest{}
	this.MerchantAccount = merchantAccount
	var splitCardFundingSources bool = false
	this.SplitCardFundingSources = &splitCardFundingSources
	return &this
}

// NewPaymentMethodsRequestWithDefaults instantiates a new PaymentMethodsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodsRequestWithDefaults() *PaymentMethodsRequest {
	this := PaymentMethodsRequest{}
	var splitCardFundingSources bool = false
	this.SplitCardFundingSources = &splitCardFundingSources
	return &this
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *PaymentMethodsRequest) GetAdditionalData() map[string]string {
	if o == nil || common.IsNil(o.AdditionalData) {
		var ret map[string]string
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodsRequest) GetAdditionalDataOk() (*map[string]string, bool) {
	if o == nil || common.IsNil(o.AdditionalData) {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *PaymentMethodsRequest) HasAdditionalData() bool {
	if o != nil && !common.IsNil(o.AdditionalData) {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given map[string]string and assigns it to the AdditionalData field.
func (o *PaymentMethodsRequest) SetAdditionalData(v map[string]string) {
	o.AdditionalData = &v
}

// GetAllowedPaymentMethods returns the AllowedPaymentMethods field value if set, zero value otherwise.
func (o *PaymentMethodsRequest) GetAllowedPaymentMethods() []string {
	if o == nil || common.IsNil(o.AllowedPaymentMethods) {
		var ret []string
		return ret
	}
	return o.AllowedPaymentMethods
}

// GetAllowedPaymentMethodsOk returns a tuple with the AllowedPaymentMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodsRequest) GetAllowedPaymentMethodsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.AllowedPaymentMethods) {
		return nil, false
	}
	return o.AllowedPaymentMethods, true
}

// HasAllowedPaymentMethods returns a boolean if a field has been set.
func (o *PaymentMethodsRequest) HasAllowedPaymentMethods() bool {
	if o != nil && !common.IsNil(o.AllowedPaymentMethods) {
		return true
	}

	return false
}

// SetAllowedPaymentMethods gets a reference to the given []string and assigns it to the AllowedPaymentMethods field.
func (o *PaymentMethodsRequest) SetAllowedPaymentMethods(v []string) {
	o.AllowedPaymentMethods = v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *PaymentMethodsRequest) GetAmount() Amount {
	if o == nil || common.IsNil(o.Amount) {
		var ret Amount
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodsRequest) GetAmountOk() (*Amount, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *PaymentMethodsRequest) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given Amount and assigns it to the Amount field.
func (o *PaymentMethodsRequest) SetAmount(v Amount) {
	o.Amount = &v
}

// GetBlockedPaymentMethods returns the BlockedPaymentMethods field value if set, zero value otherwise.
func (o *PaymentMethodsRequest) GetBlockedPaymentMethods() []string {
	if o == nil || common.IsNil(o.BlockedPaymentMethods) {
		var ret []string
		return ret
	}
	return o.BlockedPaymentMethods
}

// GetBlockedPaymentMethodsOk returns a tuple with the BlockedPaymentMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodsRequest) GetBlockedPaymentMethodsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.BlockedPaymentMethods) {
		return nil, false
	}
	return o.BlockedPaymentMethods, true
}

// HasBlockedPaymentMethods returns a boolean if a field has been set.
func (o *PaymentMethodsRequest) HasBlockedPaymentMethods() bool {
	if o != nil && !common.IsNil(o.BlockedPaymentMethods) {
		return true
	}

	return false
}

// SetBlockedPaymentMethods gets a reference to the given []string and assigns it to the BlockedPaymentMethods field.
func (o *PaymentMethodsRequest) SetBlockedPaymentMethods(v []string) {
	o.BlockedPaymentMethods = v
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *PaymentMethodsRequest) GetChannel() string {
	if o == nil || common.IsNil(o.Channel) {
		var ret string
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodsRequest) GetChannelOk() (*string, bool) {
	if o == nil || common.IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *PaymentMethodsRequest) HasChannel() bool {
	if o != nil && !common.IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given string and assigns it to the Channel field.
func (o *PaymentMethodsRequest) SetChannel(v string) {
	o.Channel = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *PaymentMethodsRequest) GetCountryCode() string {
	if o == nil || common.IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodsRequest) GetCountryCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *PaymentMethodsRequest) HasCountryCode() bool {
	if o != nil && !common.IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *PaymentMethodsRequest) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetMerchantAccount returns the MerchantAccount field value
func (o *PaymentMethodsRequest) GetMerchantAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodsRequest) GetMerchantAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantAccount, true
}

// SetMerchantAccount sets field value
func (o *PaymentMethodsRequest) SetMerchantAccount(v string) {
	o.MerchantAccount = v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *PaymentMethodsRequest) GetOrder() EncryptedOrderData {
	if o == nil || common.IsNil(o.Order) {
		var ret EncryptedOrderData
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodsRequest) GetOrderOk() (*EncryptedOrderData, bool) {
	if o == nil || common.IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *PaymentMethodsRequest) HasOrder() bool {
	if o != nil && !common.IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given EncryptedOrderData and assigns it to the Order field.
func (o *PaymentMethodsRequest) SetOrder(v EncryptedOrderData) {
	o.Order = &v
}

// GetShopperLocale returns the ShopperLocale field value if set, zero value otherwise.
func (o *PaymentMethodsRequest) GetShopperLocale() string {
	if o == nil || common.IsNil(o.ShopperLocale) {
		var ret string
		return ret
	}
	return *o.ShopperLocale
}

// GetShopperLocaleOk returns a tuple with the ShopperLocale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodsRequest) GetShopperLocaleOk() (*string, bool) {
	if o == nil || common.IsNil(o.ShopperLocale) {
		return nil, false
	}
	return o.ShopperLocale, true
}

// HasShopperLocale returns a boolean if a field has been set.
func (o *PaymentMethodsRequest) HasShopperLocale() bool {
	if o != nil && !common.IsNil(o.ShopperLocale) {
		return true
	}

	return false
}

// SetShopperLocale gets a reference to the given string and assigns it to the ShopperLocale field.
func (o *PaymentMethodsRequest) SetShopperLocale(v string) {
	o.ShopperLocale = &v
}

// GetShopperReference returns the ShopperReference field value if set, zero value otherwise.
func (o *PaymentMethodsRequest) GetShopperReference() string {
	if o == nil || common.IsNil(o.ShopperReference) {
		var ret string
		return ret
	}
	return *o.ShopperReference
}

// GetShopperReferenceOk returns a tuple with the ShopperReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodsRequest) GetShopperReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.ShopperReference) {
		return nil, false
	}
	return o.ShopperReference, true
}

// HasShopperReference returns a boolean if a field has been set.
func (o *PaymentMethodsRequest) HasShopperReference() bool {
	if o != nil && !common.IsNil(o.ShopperReference) {
		return true
	}

	return false
}

// SetShopperReference gets a reference to the given string and assigns it to the ShopperReference field.
func (o *PaymentMethodsRequest) SetShopperReference(v string) {
	o.ShopperReference = &v
}

// GetSplitCardFundingSources returns the SplitCardFundingSources field value if set, zero value otherwise.
func (o *PaymentMethodsRequest) GetSplitCardFundingSources() bool {
	if o == nil || common.IsNil(o.SplitCardFundingSources) {
		var ret bool
		return ret
	}
	return *o.SplitCardFundingSources
}

// GetSplitCardFundingSourcesOk returns a tuple with the SplitCardFundingSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodsRequest) GetSplitCardFundingSourcesOk() (*bool, bool) {
	if o == nil || common.IsNil(o.SplitCardFundingSources) {
		return nil, false
	}
	return o.SplitCardFundingSources, true
}

// HasSplitCardFundingSources returns a boolean if a field has been set.
func (o *PaymentMethodsRequest) HasSplitCardFundingSources() bool {
	if o != nil && !common.IsNil(o.SplitCardFundingSources) {
		return true
	}

	return false
}

// SetSplitCardFundingSources gets a reference to the given bool and assigns it to the SplitCardFundingSources field.
func (o *PaymentMethodsRequest) SetSplitCardFundingSources(v bool) {
	o.SplitCardFundingSources = &v
}

// GetStore returns the Store field value if set, zero value otherwise.
func (o *PaymentMethodsRequest) GetStore() string {
	if o == nil || common.IsNil(o.Store) {
		var ret string
		return ret
	}
	return *o.Store
}

// GetStoreOk returns a tuple with the Store field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodsRequest) GetStoreOk() (*string, bool) {
	if o == nil || common.IsNil(o.Store) {
		return nil, false
	}
	return o.Store, true
}

// HasStore returns a boolean if a field has been set.
func (o *PaymentMethodsRequest) HasStore() bool {
	if o != nil && !common.IsNil(o.Store) {
		return true
	}

	return false
}

// SetStore gets a reference to the given string and assigns it to the Store field.
func (o *PaymentMethodsRequest) SetStore(v string) {
	o.Store = &v
}

// GetStoreFiltrationMode returns the StoreFiltrationMode field value if set, zero value otherwise.
func (o *PaymentMethodsRequest) GetStoreFiltrationMode() string {
	if o == nil || common.IsNil(o.StoreFiltrationMode) {
		var ret string
		return ret
	}
	return *o.StoreFiltrationMode
}

// GetStoreFiltrationModeOk returns a tuple with the StoreFiltrationMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodsRequest) GetStoreFiltrationModeOk() (*string, bool) {
	if o == nil || common.IsNil(o.StoreFiltrationMode) {
		return nil, false
	}
	return o.StoreFiltrationMode, true
}

// HasStoreFiltrationMode returns a boolean if a field has been set.
func (o *PaymentMethodsRequest) HasStoreFiltrationMode() bool {
	if o != nil && !common.IsNil(o.StoreFiltrationMode) {
		return true
	}

	return false
}

// SetStoreFiltrationMode gets a reference to the given string and assigns it to the StoreFiltrationMode field.
func (o *PaymentMethodsRequest) SetStoreFiltrationMode(v string) {
	o.StoreFiltrationMode = &v
}

func (o PaymentMethodsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AdditionalData) {
		toSerialize["additionalData"] = o.AdditionalData
	}
	if !common.IsNil(o.AllowedPaymentMethods) {
		toSerialize["allowedPaymentMethods"] = o.AllowedPaymentMethods
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.BlockedPaymentMethods) {
		toSerialize["blockedPaymentMethods"] = o.BlockedPaymentMethods
	}
	if !common.IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if !common.IsNil(o.CountryCode) {
		toSerialize["countryCode"] = o.CountryCode
	}
	toSerialize["merchantAccount"] = o.MerchantAccount
	if !common.IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !common.IsNil(o.ShopperLocale) {
		toSerialize["shopperLocale"] = o.ShopperLocale
	}
	if !common.IsNil(o.ShopperReference) {
		toSerialize["shopperReference"] = o.ShopperReference
	}
	if !common.IsNil(o.SplitCardFundingSources) {
		toSerialize["splitCardFundingSources"] = o.SplitCardFundingSources
	}
	if !common.IsNil(o.Store) {
		toSerialize["store"] = o.Store
	}
	if !common.IsNil(o.StoreFiltrationMode) {
		toSerialize["storeFiltrationMode"] = o.StoreFiltrationMode
	}
	return toSerialize, nil
}

type NullablePaymentMethodsRequest struct {
	value *PaymentMethodsRequest
	isSet bool
}

func (v NullablePaymentMethodsRequest) Get() *PaymentMethodsRequest {
	return v.value
}

func (v *NullablePaymentMethodsRequest) Set(val *PaymentMethodsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodsRequest(val *PaymentMethodsRequest) *NullablePaymentMethodsRequest {
	return &NullablePaymentMethodsRequest{value: val, isSet: true}
}

func (v NullablePaymentMethodsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *PaymentMethodsRequest) isValidChannel() bool {
	var allowedEnumValues = []string{"iOS", "Android", "Web"}
	for _, allowed := range allowedEnumValues {
		if o.GetChannel() == allowed {
			return true
		}
	}
	return false
}
func (o *PaymentMethodsRequest) isValidStoreFiltrationMode() bool {
	var allowedEnumValues = []string{"exclusive", "inclusive", "skipFilter"}
	for _, allowed := range allowedEnumValues {
		if o.GetStoreFiltrationMode() == allowed {
			return true
		}
	}
	return false
}
