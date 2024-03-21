/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v9/src/common"
)

// checks if the PaymentMethodSetupInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PaymentMethodSetupInfo{}

// PaymentMethodSetupInfo struct for PaymentMethodSetupInfo
type PaymentMethodSetupInfo struct {
	AfterpayTouch *AfterpayTouchInfo `json:"afterpayTouch,omitempty"`
	ApplePay      *ApplePayInfo      `json:"applePay,omitempty"`
	Bcmc          *BcmcInfo          `json:"bcmc,omitempty"`
	// The unique identifier of the business line. Required if you are a [platform model](https://docs.adyen.com/platforms).
	BusinessLineId  *string              `json:"businessLineId,omitempty"`
	CartesBancaires *CartesBancairesInfo `json:"cartesBancaires,omitempty"`
	Clearpay        *ClearpayInfo        `json:"clearpay,omitempty"`
	// The list of countries where a payment method is available. By default, all countries supported by the payment method.
	Countries []string              `json:"countries,omitempty"`
	Cup       *GenericPmWithTdiInfo `json:"cup,omitempty"`
	// The list of currencies that a payment method supports. By default, all currencies supported by the payment method.
	Currencies []string `json:"currencies,omitempty"`
	// The list of custom routing flags to route payment to the intended acquirer.
	CustomRoutingFlags []string              `json:"customRoutingFlags,omitempty"`
	Diners             *GenericPmWithTdiInfo `json:"diners,omitempty"`
	Discover           *GenericPmWithTdiInfo `json:"discover,omitempty"`
	EftposAustralia    *GenericPmWithTdiInfo `json:"eftpos_australia,omitempty"`
	GiroPay            *GiroPayInfo          `json:"giroPay,omitempty"`
	Girocard           *GenericPmWithTdiInfo `json:"girocard,omitempty"`
	GooglePay          *GooglePayInfo        `json:"googlePay,omitempty"`
	Ideal              *GenericPmWithTdiInfo `json:"ideal,omitempty"`
	InteracCard        *GenericPmWithTdiInfo `json:"interac_card,omitempty"`
	Jcb                *GenericPmWithTdiInfo `json:"jcb,omitempty"`
	Klarna             *KlarnaInfo           `json:"klarna,omitempty"`
	Maestro            *GenericPmWithTdiInfo `json:"maestro,omitempty"`
	Mc                 *GenericPmWithTdiInfo `json:"mc,omitempty"`
	MealVoucherFR      *MealVoucherFRInfo    `json:"mealVoucher_FR,omitempty"`
	Paypal             *PayPalInfo           `json:"paypal,omitempty"`
	// Your reference for the payment method. Supported characters a-z, A-Z, 0-9.
	Reference *string `json:"reference,omitempty"`
	// The sales channel. Required if the merchant account does not have a sales channel. When you provide this field, it overrides the default sales channel set on the merchant account.  Possible values: **eCommerce**, **pos**, **contAuth**, and **moto**.
	ShopperInteraction *string     `json:"shopperInteraction,omitempty"`
	Sofort             *SofortInfo `json:"sofort,omitempty"`
	// The unique identifier of the store for which to configure the payment method, if any.
	StoreIds []string   `json:"storeIds,omitempty"`
	Swish    *SwishInfo `json:"swish,omitempty"`
	Twint    *TwintInfo `json:"twint,omitempty"`
	// Payment method [variant](https://docs.adyen.com/development-resources/paymentmethodvariant#management-api).
	Type  string                `json:"type"`
	Vipps *VippsInfo            `json:"vipps,omitempty"`
	Visa  *GenericPmWithTdiInfo `json:"visa,omitempty"`
}

// NewPaymentMethodSetupInfo instantiates a new PaymentMethodSetupInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodSetupInfo(type_ string) *PaymentMethodSetupInfo {
	this := PaymentMethodSetupInfo{}
	this.Type = type_
	return &this
}

// NewPaymentMethodSetupInfoWithDefaults instantiates a new PaymentMethodSetupInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodSetupInfoWithDefaults() *PaymentMethodSetupInfo {
	this := PaymentMethodSetupInfo{}
	return &this
}

// GetAfterpayTouch returns the AfterpayTouch field value if set, zero value otherwise.
func (o *PaymentMethodSetupInfo) GetAfterpayTouch() AfterpayTouchInfo {
	if o == nil || common.IsNil(o.AfterpayTouch) {
		var ret AfterpayTouchInfo
		return ret
	}
	return *o.AfterpayTouch
}

// GetAfterpayTouchOk returns a tuple with the AfterpayTouch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodSetupInfo) GetAfterpayTouchOk() (*AfterpayTouchInfo, bool) {
	if o == nil || common.IsNil(o.AfterpayTouch) {
		return nil, false
	}
	return o.AfterpayTouch, true
}

// HasAfterpayTouch returns a boolean if a field has been set.
func (o *PaymentMethodSetupInfo) HasAfterpayTouch() bool {
	if o != nil && !common.IsNil(o.AfterpayTouch) {
		return true
	}

	return false
}

// SetAfterpayTouch gets a reference to the given AfterpayTouchInfo and assigns it to the AfterpayTouch field.
func (o *PaymentMethodSetupInfo) SetAfterpayTouch(v AfterpayTouchInfo) {
	o.AfterpayTouch = &v
}

// GetApplePay returns the ApplePay field value if set, zero value otherwise.
func (o *PaymentMethodSetupInfo) GetApplePay() ApplePayInfo {
	if o == nil || common.IsNil(o.ApplePay) {
		var ret ApplePayInfo
		return ret
	}
	return *o.ApplePay
}

// GetApplePayOk returns a tuple with the ApplePay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodSetupInfo) GetApplePayOk() (*ApplePayInfo, bool) {
	if o == nil || common.IsNil(o.ApplePay) {
		return nil, false
	}
	return o.ApplePay, true
}

// HasApplePay returns a boolean if a field has been set.
func (o *PaymentMethodSetupInfo) HasApplePay() bool {
	if o != nil && !common.IsNil(o.ApplePay) {
		return true
	}

	return false
}

// SetApplePay gets a reference to the given ApplePayInfo and assigns it to the ApplePay field.
func (o *PaymentMethodSetupInfo) SetApplePay(v ApplePayInfo) {
	o.ApplePay = &v
}

// GetBcmc returns the Bcmc field value if set, zero value otherwise.
func (o *PaymentMethodSetupInfo) GetBcmc() BcmcInfo {
	if o == nil || common.IsNil(o.Bcmc) {
		var ret BcmcInfo
		return ret
	}
	return *o.Bcmc
}

// GetBcmcOk returns a tuple with the Bcmc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodSetupInfo) GetBcmcOk() (*BcmcInfo, bool) {
	if o == nil || common.IsNil(o.Bcmc) {
		return nil, false
	}
	return o.Bcmc, true
}

// HasBcmc returns a boolean if a field has been set.
func (o *PaymentMethodSetupInfo) HasBcmc() bool {
	if o != nil && !common.IsNil(o.Bcmc) {
		return true
	}

	return false
}

// SetBcmc gets a reference to the given BcmcInfo and assigns it to the Bcmc field.
func (o *PaymentMethodSetupInfo) SetBcmc(v BcmcInfo) {
	o.Bcmc = &v
}

// GetBusinessLineId returns the BusinessLineId field value if set, zero value otherwise.
func (o *PaymentMethodSetupInfo) GetBusinessLineId() string {
	if o == nil || common.IsNil(o.BusinessLineId) {
		var ret string
		return ret
	}
	return *o.BusinessLineId
}

// GetBusinessLineIdOk returns a tuple with the BusinessLineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodSetupInfo) GetBusinessLineIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.BusinessLineId) {
		return nil, false
	}
	return o.BusinessLineId, true
}

// HasBusinessLineId returns a boolean if a field has been set.
func (o *PaymentMethodSetupInfo) HasBusinessLineId() bool {
	if o != nil && !common.IsNil(o.BusinessLineId) {
		return true
	}

	return false
}

// SetBusinessLineId gets a reference to the given string and assigns it to the BusinessLineId field.
func (o *PaymentMethodSetupInfo) SetBusinessLineId(v string) {
	o.BusinessLineId = &v
}

// GetCartesBancaires returns the CartesBancaires field value if set, zero value otherwise.
func (o *PaymentMethodSetupInfo) GetCartesBancaires() CartesBancairesInfo {
	if o == nil || common.IsNil(o.CartesBancaires) {
		var ret CartesBancairesInfo
		return ret
	}
	return *o.CartesBancaires
}

// GetCartesBancairesOk returns a tuple with the CartesBancaires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodSetupInfo) GetCartesBancairesOk() (*CartesBancairesInfo, bool) {
	if o == nil || common.IsNil(o.CartesBancaires) {
		return nil, false
	}
	return o.CartesBancaires, true
}

// HasCartesBancaires returns a boolean if a field has been set.
func (o *PaymentMethodSetupInfo) HasCartesBancaires() bool {
	if o != nil && !common.IsNil(o.CartesBancaires) {
		return true
	}

	return false
}

// SetCartesBancaires gets a reference to the given CartesBancairesInfo and assigns it to the CartesBancaires field.
func (o *PaymentMethodSetupInfo) SetCartesBancaires(v CartesBancairesInfo) {
	o.CartesBancaires = &v
}

// GetClearpay returns the Clearpay field value if set, zero value otherwise.
func (o *PaymentMethodSetupInfo) GetClearpay() ClearpayInfo {
	if o == nil || common.IsNil(o.Clearpay) {
		var ret ClearpayInfo
		return ret
	}
	return *o.Clearpay
}

// GetClearpayOk returns a tuple with the Clearpay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodSetupInfo) GetClearpayOk() (*ClearpayInfo, bool) {
	if o == nil || common.IsNil(o.Clearpay) {
		return nil, false
	}
	return o.Clearpay, true
}

// HasClearpay returns a boolean if a field has been set.
func (o *PaymentMethodSetupInfo) HasClearpay() bool {
	if o != nil && !common.IsNil(o.Clearpay) {
		return true
	}

	return false
}

// SetClearpay gets a reference to the given ClearpayInfo and assigns it to the Clearpay field.
func (o *PaymentMethodSetupInfo) SetClearpay(v ClearpayInfo) {
	o.Clearpay = &v
}

// GetCountries returns the Countries field value if set, zero value otherwise.
func (o *PaymentMethodSetupInfo) GetCountries() []string {
	if o == nil || common.IsNil(o.Countries) {
		var ret []string
		return ret
	}
	return o.Countries
}

// GetCountriesOk returns a tuple with the Countries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodSetupInfo) GetCountriesOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Countries) {
		return nil, false
	}
	return o.Countries, true
}

// HasCountries returns a boolean if a field has been set.
func (o *PaymentMethodSetupInfo) HasCountries() bool {
	if o != nil && !common.IsNil(o.Countries) {
		return true
	}

	return false
}

// SetCountries gets a reference to the given []string and assigns it to the Countries field.
func (o *PaymentMethodSetupInfo) SetCountries(v []string) {
	o.Countries = v
}

// GetCup returns the Cup field value if set, zero value otherwise.
func (o *PaymentMethodSetupInfo) GetCup() GenericPmWithTdiInfo {
	if o == nil || common.IsNil(o.Cup) {
		var ret GenericPmWithTdiInfo
		return ret
	}
	return *o.Cup
}

// GetCupOk returns a tuple with the Cup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodSetupInfo) GetCupOk() (*GenericPmWithTdiInfo, bool) {
	if o == nil || common.IsNil(o.Cup) {
		return nil, false
	}
	return o.Cup, true
}

// HasCup returns a boolean if a field has been set.
func (o *PaymentMethodSetupInfo) HasCup() bool {
	if o != nil && !common.IsNil(o.Cup) {
		return true
	}

	return false
}

// SetCup gets a reference to the given GenericPmWithTdiInfo and assigns it to the Cup field.
func (o *PaymentMethodSetupInfo) SetCup(v GenericPmWithTdiInfo) {
	o.Cup = &v
}

// GetCurrencies returns the Currencies field value if set, zero value otherwise.
func (o *PaymentMethodSetupInfo) GetCurrencies() []string {
	if o == nil || common.IsNil(o.Currencies) {
		var ret []string
		return ret
	}
	return o.Currencies
}

// GetCurrenciesOk returns a tuple with the Currencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodSetupInfo) GetCurrenciesOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Currencies) {
		return nil, false
	}
	return o.Currencies, true
}

// HasCurrencies returns a boolean if a field has been set.
func (o *PaymentMethodSetupInfo) HasCurrencies() bool {
	if o != nil && !common.IsNil(o.Currencies) {
		return true
	}

	return false
}

// SetCurrencies gets a reference to the given []string and assigns it to the Currencies field.
func (o *PaymentMethodSetupInfo) SetCurrencies(v []string) {
	o.Currencies = v
}

// GetCustomRoutingFlags returns the CustomRoutingFlags field value if set, zero value otherwise.
func (o *PaymentMethodSetupInfo) GetCustomRoutingFlags() []string {
	if o == nil || common.IsNil(o.CustomRoutingFlags) {
		var ret []string
		return ret
	}
	return o.CustomRoutingFlags
}

// GetCustomRoutingFlagsOk returns a tuple with the CustomRoutingFlags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodSetupInfo) GetCustomRoutingFlagsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.CustomRoutingFlags) {
		return nil, false
	}
	return o.CustomRoutingFlags, true
}

// HasCustomRoutingFlags returns a boolean if a field has been set.
func (o *PaymentMethodSetupInfo) HasCustomRoutingFlags() bool {
	if o != nil && !common.IsNil(o.CustomRoutingFlags) {
		return true
	}

	return false
}

// SetCustomRoutingFlags gets a reference to the given []string and assigns it to the CustomRoutingFlags field.
func (o *PaymentMethodSetupInfo) SetCustomRoutingFlags(v []string) {
	o.CustomRoutingFlags = v
}

// GetDiners returns the Diners field value if set, zero value otherwise.
func (o *PaymentMethodSetupInfo) GetDiners() GenericPmWithTdiInfo {
	if o == nil || common.IsNil(o.Diners) {
		var ret GenericPmWithTdiInfo
		return ret
	}
	return *o.Diners
}

// GetDinersOk returns a tuple with the Diners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodSetupInfo) GetDinersOk() (*GenericPmWithTdiInfo, bool) {
	if o == nil || common.IsNil(o.Diners) {
		return nil, false
	}
	return o.Diners, true
}

// HasDiners returns a boolean if a field has been set.
func (o *PaymentMethodSetupInfo) HasDiners() bool {
	if o != nil && !common.IsNil(o.Diners) {
		return true
	}

	return false
}

// SetDiners gets a reference to the given GenericPmWithTdiInfo and assigns it to the Diners field.
func (o *PaymentMethodSetupInfo) SetDiners(v GenericPmWithTdiInfo) {
	o.Diners = &v
}

// GetDiscover returns the Discover field value if set, zero value otherwise.
func (o *PaymentMethodSetupInfo) GetDiscover() GenericPmWithTdiInfo {
	if o == nil || common.IsNil(o.Discover) {
		var ret GenericPmWithTdiInfo
		return ret
	}
	return *o.Discover
}

// GetDiscoverOk returns a tuple with the Discover field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodSetupInfo) GetDiscoverOk() (*GenericPmWithTdiInfo, bool) {
	if o == nil || common.IsNil(o.Discover) {
		return nil, false
	}
	return o.Discover, true
}

// HasDiscover returns a boolean if a field has been set.
func (o *PaymentMethodSetupInfo) HasDiscover() bool {
	if o != nil && !common.IsNil(o.Discover) {
		return true
	}

	return false
}

// SetDiscover gets a reference to the given GenericPmWithTdiInfo and assigns it to the Discover field.
func (o *PaymentMethodSetupInfo) SetDiscover(v GenericPmWithTdiInfo) {
	o.Discover = &v
}

// GetEftposAustralia returns the EftposAustralia field value if set, zero value otherwise.
func (o *PaymentMethodSetupInfo) GetEftposAustralia() GenericPmWithTdiInfo {
	if o == nil || common.IsNil(o.EftposAustralia) {
		var ret GenericPmWithTdiInfo
		return ret
	}
	return *o.EftposAustralia
}

// GetEftposAustraliaOk returns a tuple with the EftposAustralia field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodSetupInfo) GetEftposAustraliaOk() (*GenericPmWithTdiInfo, bool) {
	if o == nil || common.IsNil(o.EftposAustralia) {
		return nil, false
	}
	return o.EftposAustralia, true
}

// HasEftposAustralia returns a boolean if a field has been set.
func (o *PaymentMethodSetupInfo) HasEftposAustralia() bool {
	if o != nil && !common.IsNil(o.EftposAustralia) {
		return true
	}

	return false
}

// SetEftposAustralia gets a reference to the given GenericPmWithTdiInfo and assigns it to the EftposAustralia field.
func (o *PaymentMethodSetupInfo) SetEftposAustralia(v GenericPmWithTdiInfo) {
	o.EftposAustralia = &v
}

// GetGiroPay returns the GiroPay field value if set, zero value otherwise.
func (o *PaymentMethodSetupInfo) GetGiroPay() GiroPayInfo {
	if o == nil || common.IsNil(o.GiroPay) {
		var ret GiroPayInfo
		return ret
	}
	return *o.GiroPay
}

// GetGiroPayOk returns a tuple with the GiroPay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodSetupInfo) GetGiroPayOk() (*GiroPayInfo, bool) {
	if o == nil || common.IsNil(o.GiroPay) {
		return nil, false
	}
	return o.GiroPay, true
}

// HasGiroPay returns a boolean if a field has been set.
func (o *PaymentMethodSetupInfo) HasGiroPay() bool {
	if o != nil && !common.IsNil(o.GiroPay) {
		return true
	}

	return false
}

// SetGiroPay gets a reference to the given GiroPayInfo and assigns it to the GiroPay field.
func (o *PaymentMethodSetupInfo) SetGiroPay(v GiroPayInfo) {
	o.GiroPay = &v
}

// GetGirocard returns the Girocard field value if set, zero value otherwise.
func (o *PaymentMethodSetupInfo) GetGirocard() GenericPmWithTdiInfo {
	if o == nil || common.IsNil(o.Girocard) {
		var ret GenericPmWithTdiInfo
		return ret
	}
	return *o.Girocard
}

// GetGirocardOk returns a tuple with the Girocard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodSetupInfo) GetGirocardOk() (*GenericPmWithTdiInfo, bool) {
	if o == nil || common.IsNil(o.Girocard) {
		return nil, false
	}
	return o.Girocard, true
}

// HasGirocard returns a boolean if a field has been set.
func (o *PaymentMethodSetupInfo) HasGirocard() bool {
	if o != nil && !common.IsNil(o.Girocard) {
		return true
	}

	return false
}

// SetGirocard gets a reference to the given GenericPmWithTdiInfo and assigns it to the Girocard field.
func (o *PaymentMethodSetupInfo) SetGirocard(v GenericPmWithTdiInfo) {
	o.Girocard = &v
}

// GetGooglePay returns the GooglePay field value if set, zero value otherwise.
func (o *PaymentMethodSetupInfo) GetGooglePay() GooglePayInfo {
	if o == nil || common.IsNil(o.GooglePay) {
		var ret GooglePayInfo
		return ret
	}
	return *o.GooglePay
}

// GetGooglePayOk returns a tuple with the GooglePay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodSetupInfo) GetGooglePayOk() (*GooglePayInfo, bool) {
	if o == nil || common.IsNil(o.GooglePay) {
		return nil, false
	}
	return o.GooglePay, true
}

// HasGooglePay returns a boolean if a field has been set.
func (o *PaymentMethodSetupInfo) HasGooglePay() bool {
	if o != nil && !common.IsNil(o.GooglePay) {
		return true
	}

	return false
}

// SetGooglePay gets a reference to the given GooglePayInfo and assigns it to the GooglePay field.
func (o *PaymentMethodSetupInfo) SetGooglePay(v GooglePayInfo) {
	o.GooglePay = &v
}

// GetIdeal returns the Ideal field value if set, zero value otherwise.
func (o *PaymentMethodSetupInfo) GetIdeal() GenericPmWithTdiInfo {
	if o == nil || common.IsNil(o.Ideal) {
		var ret GenericPmWithTdiInfo
		return ret
	}
	return *o.Ideal
}

// GetIdealOk returns a tuple with the Ideal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodSetupInfo) GetIdealOk() (*GenericPmWithTdiInfo, bool) {
	if o == nil || common.IsNil(o.Ideal) {
		return nil, false
	}
	return o.Ideal, true
}

// HasIdeal returns a boolean if a field has been set.
func (o *PaymentMethodSetupInfo) HasIdeal() bool {
	if o != nil && !common.IsNil(o.Ideal) {
		return true
	}

	return false
}

// SetIdeal gets a reference to the given GenericPmWithTdiInfo and assigns it to the Ideal field.
func (o *PaymentMethodSetupInfo) SetIdeal(v GenericPmWithTdiInfo) {
	o.Ideal = &v
}

// GetInteracCard returns the InteracCard field value if set, zero value otherwise.
func (o *PaymentMethodSetupInfo) GetInteracCard() GenericPmWithTdiInfo {
	if o == nil || common.IsNil(o.InteracCard) {
		var ret GenericPmWithTdiInfo
		return ret
	}
	return *o.InteracCard
}

// GetInteracCardOk returns a tuple with the InteracCard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodSetupInfo) GetInteracCardOk() (*GenericPmWithTdiInfo, bool) {
	if o == nil || common.IsNil(o.InteracCard) {
		return nil, false
	}
	return o.InteracCard, true
}

// HasInteracCard returns a boolean if a field has been set.
func (o *PaymentMethodSetupInfo) HasInteracCard() bool {
	if o != nil && !common.IsNil(o.InteracCard) {
		return true
	}

	return false
}

// SetInteracCard gets a reference to the given GenericPmWithTdiInfo and assigns it to the InteracCard field.
func (o *PaymentMethodSetupInfo) SetInteracCard(v GenericPmWithTdiInfo) {
	o.InteracCard = &v
}

// GetJcb returns the Jcb field value if set, zero value otherwise.
func (o *PaymentMethodSetupInfo) GetJcb() GenericPmWithTdiInfo {
	if o == nil || common.IsNil(o.Jcb) {
		var ret GenericPmWithTdiInfo
		return ret
	}
	return *o.Jcb
}

// GetJcbOk returns a tuple with the Jcb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodSetupInfo) GetJcbOk() (*GenericPmWithTdiInfo, bool) {
	if o == nil || common.IsNil(o.Jcb) {
		return nil, false
	}
	return o.Jcb, true
}

// HasJcb returns a boolean if a field has been set.
func (o *PaymentMethodSetupInfo) HasJcb() bool {
	if o != nil && !common.IsNil(o.Jcb) {
		return true
	}

	return false
}

// SetJcb gets a reference to the given GenericPmWithTdiInfo and assigns it to the Jcb field.
func (o *PaymentMethodSetupInfo) SetJcb(v GenericPmWithTdiInfo) {
	o.Jcb = &v
}

// GetKlarna returns the Klarna field value if set, zero value otherwise.
func (o *PaymentMethodSetupInfo) GetKlarna() KlarnaInfo {
	if o == nil || common.IsNil(o.Klarna) {
		var ret KlarnaInfo
		return ret
	}
	return *o.Klarna
}

// GetKlarnaOk returns a tuple with the Klarna field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodSetupInfo) GetKlarnaOk() (*KlarnaInfo, bool) {
	if o == nil || common.IsNil(o.Klarna) {
		return nil, false
	}
	return o.Klarna, true
}

// HasKlarna returns a boolean if a field has been set.
func (o *PaymentMethodSetupInfo) HasKlarna() bool {
	if o != nil && !common.IsNil(o.Klarna) {
		return true
	}

	return false
}

// SetKlarna gets a reference to the given KlarnaInfo and assigns it to the Klarna field.
func (o *PaymentMethodSetupInfo) SetKlarna(v KlarnaInfo) {
	o.Klarna = &v
}

// GetMaestro returns the Maestro field value if set, zero value otherwise.
func (o *PaymentMethodSetupInfo) GetMaestro() GenericPmWithTdiInfo {
	if o == nil || common.IsNil(o.Maestro) {
		var ret GenericPmWithTdiInfo
		return ret
	}
	return *o.Maestro
}

// GetMaestroOk returns a tuple with the Maestro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodSetupInfo) GetMaestroOk() (*GenericPmWithTdiInfo, bool) {
	if o == nil || common.IsNil(o.Maestro) {
		return nil, false
	}
	return o.Maestro, true
}

// HasMaestro returns a boolean if a field has been set.
func (o *PaymentMethodSetupInfo) HasMaestro() bool {
	if o != nil && !common.IsNil(o.Maestro) {
		return true
	}

	return false
}

// SetMaestro gets a reference to the given GenericPmWithTdiInfo and assigns it to the Maestro field.
func (o *PaymentMethodSetupInfo) SetMaestro(v GenericPmWithTdiInfo) {
	o.Maestro = &v
}

// GetMc returns the Mc field value if set, zero value otherwise.
func (o *PaymentMethodSetupInfo) GetMc() GenericPmWithTdiInfo {
	if o == nil || common.IsNil(o.Mc) {
		var ret GenericPmWithTdiInfo
		return ret
	}
	return *o.Mc
}

// GetMcOk returns a tuple with the Mc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodSetupInfo) GetMcOk() (*GenericPmWithTdiInfo, bool) {
	if o == nil || common.IsNil(o.Mc) {
		return nil, false
	}
	return o.Mc, true
}

// HasMc returns a boolean if a field has been set.
func (o *PaymentMethodSetupInfo) HasMc() bool {
	if o != nil && !common.IsNil(o.Mc) {
		return true
	}

	return false
}

// SetMc gets a reference to the given GenericPmWithTdiInfo and assigns it to the Mc field.
func (o *PaymentMethodSetupInfo) SetMc(v GenericPmWithTdiInfo) {
	o.Mc = &v
}

// GetMealVoucherFR returns the MealVoucherFR field value if set, zero value otherwise.
func (o *PaymentMethodSetupInfo) GetMealVoucherFR() MealVoucherFRInfo {
	if o == nil || common.IsNil(o.MealVoucherFR) {
		var ret MealVoucherFRInfo
		return ret
	}
	return *o.MealVoucherFR
}

// GetMealVoucherFROk returns a tuple with the MealVoucherFR field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodSetupInfo) GetMealVoucherFROk() (*MealVoucherFRInfo, bool) {
	if o == nil || common.IsNil(o.MealVoucherFR) {
		return nil, false
	}
	return o.MealVoucherFR, true
}

// HasMealVoucherFR returns a boolean if a field has been set.
func (o *PaymentMethodSetupInfo) HasMealVoucherFR() bool {
	if o != nil && !common.IsNil(o.MealVoucherFR) {
		return true
	}

	return false
}

// SetMealVoucherFR gets a reference to the given MealVoucherFRInfo and assigns it to the MealVoucherFR field.
func (o *PaymentMethodSetupInfo) SetMealVoucherFR(v MealVoucherFRInfo) {
	o.MealVoucherFR = &v
}

// GetPaypal returns the Paypal field value if set, zero value otherwise.
func (o *PaymentMethodSetupInfo) GetPaypal() PayPalInfo {
	if o == nil || common.IsNil(o.Paypal) {
		var ret PayPalInfo
		return ret
	}
	return *o.Paypal
}

// GetPaypalOk returns a tuple with the Paypal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodSetupInfo) GetPaypalOk() (*PayPalInfo, bool) {
	if o == nil || common.IsNil(o.Paypal) {
		return nil, false
	}
	return o.Paypal, true
}

// HasPaypal returns a boolean if a field has been set.
func (o *PaymentMethodSetupInfo) HasPaypal() bool {
	if o != nil && !common.IsNil(o.Paypal) {
		return true
	}

	return false
}

// SetPaypal gets a reference to the given PayPalInfo and assigns it to the Paypal field.
func (o *PaymentMethodSetupInfo) SetPaypal(v PayPalInfo) {
	o.Paypal = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *PaymentMethodSetupInfo) GetReference() string {
	if o == nil || common.IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodSetupInfo) GetReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PaymentMethodSetupInfo) HasReference() bool {
	if o != nil && !common.IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *PaymentMethodSetupInfo) SetReference(v string) {
	o.Reference = &v
}

// GetShopperInteraction returns the ShopperInteraction field value if set, zero value otherwise.
func (o *PaymentMethodSetupInfo) GetShopperInteraction() string {
	if o == nil || common.IsNil(o.ShopperInteraction) {
		var ret string
		return ret
	}
	return *o.ShopperInteraction
}

// GetShopperInteractionOk returns a tuple with the ShopperInteraction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodSetupInfo) GetShopperInteractionOk() (*string, bool) {
	if o == nil || common.IsNil(o.ShopperInteraction) {
		return nil, false
	}
	return o.ShopperInteraction, true
}

// HasShopperInteraction returns a boolean if a field has been set.
func (o *PaymentMethodSetupInfo) HasShopperInteraction() bool {
	if o != nil && !common.IsNil(o.ShopperInteraction) {
		return true
	}

	return false
}

// SetShopperInteraction gets a reference to the given string and assigns it to the ShopperInteraction field.
func (o *PaymentMethodSetupInfo) SetShopperInteraction(v string) {
	o.ShopperInteraction = &v
}

// GetSofort returns the Sofort field value if set, zero value otherwise.
func (o *PaymentMethodSetupInfo) GetSofort() SofortInfo {
	if o == nil || common.IsNil(o.Sofort) {
		var ret SofortInfo
		return ret
	}
	return *o.Sofort
}

// GetSofortOk returns a tuple with the Sofort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodSetupInfo) GetSofortOk() (*SofortInfo, bool) {
	if o == nil || common.IsNil(o.Sofort) {
		return nil, false
	}
	return o.Sofort, true
}

// HasSofort returns a boolean if a field has been set.
func (o *PaymentMethodSetupInfo) HasSofort() bool {
	if o != nil && !common.IsNil(o.Sofort) {
		return true
	}

	return false
}

// SetSofort gets a reference to the given SofortInfo and assigns it to the Sofort field.
func (o *PaymentMethodSetupInfo) SetSofort(v SofortInfo) {
	o.Sofort = &v
}

// GetStoreIds returns the StoreIds field value if set, zero value otherwise.
func (o *PaymentMethodSetupInfo) GetStoreIds() []string {
	if o == nil || common.IsNil(o.StoreIds) {
		var ret []string
		return ret
	}
	return o.StoreIds
}

// GetStoreIdsOk returns a tuple with the StoreIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodSetupInfo) GetStoreIdsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.StoreIds) {
		return nil, false
	}
	return o.StoreIds, true
}

// HasStoreIds returns a boolean if a field has been set.
func (o *PaymentMethodSetupInfo) HasStoreIds() bool {
	if o != nil && !common.IsNil(o.StoreIds) {
		return true
	}

	return false
}

// SetStoreIds gets a reference to the given []string and assigns it to the StoreIds field.
func (o *PaymentMethodSetupInfo) SetStoreIds(v []string) {
	o.StoreIds = v
}

// GetSwish returns the Swish field value if set, zero value otherwise.
func (o *PaymentMethodSetupInfo) GetSwish() SwishInfo {
	if o == nil || common.IsNil(o.Swish) {
		var ret SwishInfo
		return ret
	}
	return *o.Swish
}

// GetSwishOk returns a tuple with the Swish field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodSetupInfo) GetSwishOk() (*SwishInfo, bool) {
	if o == nil || common.IsNil(o.Swish) {
		return nil, false
	}
	return o.Swish, true
}

// HasSwish returns a boolean if a field has been set.
func (o *PaymentMethodSetupInfo) HasSwish() bool {
	if o != nil && !common.IsNil(o.Swish) {
		return true
	}

	return false
}

// SetSwish gets a reference to the given SwishInfo and assigns it to the Swish field.
func (o *PaymentMethodSetupInfo) SetSwish(v SwishInfo) {
	o.Swish = &v
}

// GetTwint returns the Twint field value if set, zero value otherwise.
func (o *PaymentMethodSetupInfo) GetTwint() TwintInfo {
	if o == nil || common.IsNil(o.Twint) {
		var ret TwintInfo
		return ret
	}
	return *o.Twint
}

// GetTwintOk returns a tuple with the Twint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodSetupInfo) GetTwintOk() (*TwintInfo, bool) {
	if o == nil || common.IsNil(o.Twint) {
		return nil, false
	}
	return o.Twint, true
}

// HasTwint returns a boolean if a field has been set.
func (o *PaymentMethodSetupInfo) HasTwint() bool {
	if o != nil && !common.IsNil(o.Twint) {
		return true
	}

	return false
}

// SetTwint gets a reference to the given TwintInfo and assigns it to the Twint field.
func (o *PaymentMethodSetupInfo) SetTwint(v TwintInfo) {
	o.Twint = &v
}

// GetType returns the Type field value
func (o *PaymentMethodSetupInfo) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodSetupInfo) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PaymentMethodSetupInfo) SetType(v string) {
	o.Type = v
}

// GetVipps returns the Vipps field value if set, zero value otherwise.
func (o *PaymentMethodSetupInfo) GetVipps() VippsInfo {
	if o == nil || common.IsNil(o.Vipps) {
		var ret VippsInfo
		return ret
	}
	return *o.Vipps
}

// GetVippsOk returns a tuple with the Vipps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodSetupInfo) GetVippsOk() (*VippsInfo, bool) {
	if o == nil || common.IsNil(o.Vipps) {
		return nil, false
	}
	return o.Vipps, true
}

// HasVipps returns a boolean if a field has been set.
func (o *PaymentMethodSetupInfo) HasVipps() bool {
	if o != nil && !common.IsNil(o.Vipps) {
		return true
	}

	return false
}

// SetVipps gets a reference to the given VippsInfo and assigns it to the Vipps field.
func (o *PaymentMethodSetupInfo) SetVipps(v VippsInfo) {
	o.Vipps = &v
}

// GetVisa returns the Visa field value if set, zero value otherwise.
func (o *PaymentMethodSetupInfo) GetVisa() GenericPmWithTdiInfo {
	if o == nil || common.IsNil(o.Visa) {
		var ret GenericPmWithTdiInfo
		return ret
	}
	return *o.Visa
}

// GetVisaOk returns a tuple with the Visa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodSetupInfo) GetVisaOk() (*GenericPmWithTdiInfo, bool) {
	if o == nil || common.IsNil(o.Visa) {
		return nil, false
	}
	return o.Visa, true
}

// HasVisa returns a boolean if a field has been set.
func (o *PaymentMethodSetupInfo) HasVisa() bool {
	if o != nil && !common.IsNil(o.Visa) {
		return true
	}

	return false
}

// SetVisa gets a reference to the given GenericPmWithTdiInfo and assigns it to the Visa field.
func (o *PaymentMethodSetupInfo) SetVisa(v GenericPmWithTdiInfo) {
	o.Visa = &v
}

func (o PaymentMethodSetupInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodSetupInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AfterpayTouch) {
		toSerialize["afterpayTouch"] = o.AfterpayTouch
	}
	if !common.IsNil(o.ApplePay) {
		toSerialize["applePay"] = o.ApplePay
	}
	if !common.IsNil(o.Bcmc) {
		toSerialize["bcmc"] = o.Bcmc
	}
	if !common.IsNil(o.BusinessLineId) {
		toSerialize["businessLineId"] = o.BusinessLineId
	}
	if !common.IsNil(o.CartesBancaires) {
		toSerialize["cartesBancaires"] = o.CartesBancaires
	}
	if !common.IsNil(o.Clearpay) {
		toSerialize["clearpay"] = o.Clearpay
	}
	if !common.IsNil(o.Countries) {
		toSerialize["countries"] = o.Countries
	}
	if !common.IsNil(o.Cup) {
		toSerialize["cup"] = o.Cup
	}
	if !common.IsNil(o.Currencies) {
		toSerialize["currencies"] = o.Currencies
	}
	if !common.IsNil(o.CustomRoutingFlags) {
		toSerialize["customRoutingFlags"] = o.CustomRoutingFlags
	}
	if !common.IsNil(o.Diners) {
		toSerialize["diners"] = o.Diners
	}
	if !common.IsNil(o.Discover) {
		toSerialize["discover"] = o.Discover
	}
	if !common.IsNil(o.EftposAustralia) {
		toSerialize["eftpos_australia"] = o.EftposAustralia
	}
	if !common.IsNil(o.GiroPay) {
		toSerialize["giroPay"] = o.GiroPay
	}
	if !common.IsNil(o.Girocard) {
		toSerialize["girocard"] = o.Girocard
	}
	if !common.IsNil(o.GooglePay) {
		toSerialize["googlePay"] = o.GooglePay
	}
	if !common.IsNil(o.Ideal) {
		toSerialize["ideal"] = o.Ideal
	}
	if !common.IsNil(o.InteracCard) {
		toSerialize["interac_card"] = o.InteracCard
	}
	if !common.IsNil(o.Jcb) {
		toSerialize["jcb"] = o.Jcb
	}
	if !common.IsNil(o.Klarna) {
		toSerialize["klarna"] = o.Klarna
	}
	if !common.IsNil(o.Maestro) {
		toSerialize["maestro"] = o.Maestro
	}
	if !common.IsNil(o.Mc) {
		toSerialize["mc"] = o.Mc
	}
	if !common.IsNil(o.MealVoucherFR) {
		toSerialize["mealVoucher_FR"] = o.MealVoucherFR
	}
	if !common.IsNil(o.Paypal) {
		toSerialize["paypal"] = o.Paypal
	}
	if !common.IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !common.IsNil(o.ShopperInteraction) {
		toSerialize["shopperInteraction"] = o.ShopperInteraction
	}
	if !common.IsNil(o.Sofort) {
		toSerialize["sofort"] = o.Sofort
	}
	if !common.IsNil(o.StoreIds) {
		toSerialize["storeIds"] = o.StoreIds
	}
	if !common.IsNil(o.Swish) {
		toSerialize["swish"] = o.Swish
	}
	if !common.IsNil(o.Twint) {
		toSerialize["twint"] = o.Twint
	}
	toSerialize["type"] = o.Type
	if !common.IsNil(o.Vipps) {
		toSerialize["vipps"] = o.Vipps
	}
	if !common.IsNil(o.Visa) {
		toSerialize["visa"] = o.Visa
	}
	return toSerialize, nil
}

type NullablePaymentMethodSetupInfo struct {
	value *PaymentMethodSetupInfo
	isSet bool
}

func (v NullablePaymentMethodSetupInfo) Get() *PaymentMethodSetupInfo {
	return v.value
}

func (v *NullablePaymentMethodSetupInfo) Set(val *PaymentMethodSetupInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodSetupInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodSetupInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodSetupInfo(val *PaymentMethodSetupInfo) *NullablePaymentMethodSetupInfo {
	return &NullablePaymentMethodSetupInfo{value: val, isSet: true}
}

func (v NullablePaymentMethodSetupInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodSetupInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *PaymentMethodSetupInfo) isValidShopperInteraction() bool {
	var allowedEnumValues = []string{"eCommerce", "pos", "moto", "contAuth"}
	for _, allowed := range allowedEnumValues {
		if o.GetShopperInteraction() == allowed {
			return true
		}
	}
	return false
}
func (o *PaymentMethodSetupInfo) isValidType() bool {
	var allowedEnumValues = []string{"afterpaytouch", "alipay", "alipay_hk", "amex", "applepay", "bcmc", "blik", "cartebancaire", "clearpay", "cup", "diners", "directdebit_GB", "discover", "ebanking_FI", "eftpos_australia", "elo", "elocredit", "elodebit", "girocard", "googlepay", "hiper", "hipercard", "ideal", "interac_card", "jcb", "klarna", "klarna_account", "klarna_paynow", "maestro", "mbway", "mc", "mcdebit", "mealVoucher_FR", "mobilepay", "multibanco", "onlineBanking_PL", "paybybank", "paypal", "payshop", "swish", "trustly", "twint", "twint_pos", "vipps", "visa", "visadebit", "vpay", "wechatpay", "wechatpay_pos"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
