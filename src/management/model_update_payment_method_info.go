/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v11/src/common"
)

// checks if the UpdatePaymentMethodInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UpdatePaymentMethodInfo{}

// UpdatePaymentMethodInfo struct for UpdatePaymentMethodInfo
type UpdatePaymentMethodInfo struct {
	Bcmc            *BcmcInfo            `json:"bcmc,omitempty"`
	CartesBancaires *CartesBancairesInfo `json:"cartesBancaires,omitempty"`
	// The list of countries where a payment method is available. By default, all countries supported by the payment method.
	Countries []string              `json:"countries,omitempty"`
	Cup       *GenericPmWithTdiInfo `json:"cup,omitempty"`
	// The list of currencies that a payment method supports. By default, all currencies supported by the payment method.
	Currencies []string `json:"currencies,omitempty"`
	// Custom routing flags for acquirer routing.
	CustomRoutingFlags []string              `json:"customRoutingFlags,omitempty"`
	Diners             *GenericPmWithTdiInfo `json:"diners,omitempty"`
	Discover           *GenericPmWithTdiInfo `json:"discover,omitempty"`
	EftposAustralia    *GenericPmWithTdiInfo `json:"eftpos_australia,omitempty"`
	// Indicates whether the payment method is enabled (**true**) or disabled (**false**).
	Enabled     *bool                 `json:"enabled,omitempty"`
	Girocard    *GenericPmWithTdiInfo `json:"girocard,omitempty"`
	Ideal       *GenericPmWithTdiInfo `json:"ideal,omitempty"`
	InteracCard *GenericPmWithTdiInfo `json:"interac_card,omitempty"`
	Jcb         *GenericPmWithTdiInfo `json:"jcb,omitempty"`
	Maestro     *GenericPmWithTdiInfo `json:"maestro,omitempty"`
	Mc          *GenericPmWithTdiInfo `json:"mc,omitempty"`
	// The list of stores for this payment method
	StoreIds []string              `json:"storeIds,omitempty"`
	Visa     *GenericPmWithTdiInfo `json:"visa,omitempty"`
}

// NewUpdatePaymentMethodInfo instantiates a new UpdatePaymentMethodInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePaymentMethodInfo() *UpdatePaymentMethodInfo {
	this := UpdatePaymentMethodInfo{}
	return &this
}

// NewUpdatePaymentMethodInfoWithDefaults instantiates a new UpdatePaymentMethodInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePaymentMethodInfoWithDefaults() *UpdatePaymentMethodInfo {
	this := UpdatePaymentMethodInfo{}
	return &this
}

// GetBcmc returns the Bcmc field value if set, zero value otherwise.
func (o *UpdatePaymentMethodInfo) GetBcmc() BcmcInfo {
	if o == nil || common.IsNil(o.Bcmc) {
		var ret BcmcInfo
		return ret
	}
	return *o.Bcmc
}

// GetBcmcOk returns a tuple with the Bcmc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePaymentMethodInfo) GetBcmcOk() (*BcmcInfo, bool) {
	if o == nil || common.IsNil(o.Bcmc) {
		return nil, false
	}
	return o.Bcmc, true
}

// HasBcmc returns a boolean if a field has been set.
func (o *UpdatePaymentMethodInfo) HasBcmc() bool {
	if o != nil && !common.IsNil(o.Bcmc) {
		return true
	}

	return false
}

// SetBcmc gets a reference to the given BcmcInfo and assigns it to the Bcmc field.
func (o *UpdatePaymentMethodInfo) SetBcmc(v BcmcInfo) {
	o.Bcmc = &v
}

// GetCartesBancaires returns the CartesBancaires field value if set, zero value otherwise.
func (o *UpdatePaymentMethodInfo) GetCartesBancaires() CartesBancairesInfo {
	if o == nil || common.IsNil(o.CartesBancaires) {
		var ret CartesBancairesInfo
		return ret
	}
	return *o.CartesBancaires
}

// GetCartesBancairesOk returns a tuple with the CartesBancaires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePaymentMethodInfo) GetCartesBancairesOk() (*CartesBancairesInfo, bool) {
	if o == nil || common.IsNil(o.CartesBancaires) {
		return nil, false
	}
	return o.CartesBancaires, true
}

// HasCartesBancaires returns a boolean if a field has been set.
func (o *UpdatePaymentMethodInfo) HasCartesBancaires() bool {
	if o != nil && !common.IsNil(o.CartesBancaires) {
		return true
	}

	return false
}

// SetCartesBancaires gets a reference to the given CartesBancairesInfo and assigns it to the CartesBancaires field.
func (o *UpdatePaymentMethodInfo) SetCartesBancaires(v CartesBancairesInfo) {
	o.CartesBancaires = &v
}

// GetCountries returns the Countries field value if set, zero value otherwise.
func (o *UpdatePaymentMethodInfo) GetCountries() []string {
	if o == nil || common.IsNil(o.Countries) {
		var ret []string
		return ret
	}
	return o.Countries
}

// GetCountriesOk returns a tuple with the Countries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePaymentMethodInfo) GetCountriesOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Countries) {
		return nil, false
	}
	return o.Countries, true
}

// HasCountries returns a boolean if a field has been set.
func (o *UpdatePaymentMethodInfo) HasCountries() bool {
	if o != nil && !common.IsNil(o.Countries) {
		return true
	}

	return false
}

// SetCountries gets a reference to the given []string and assigns it to the Countries field.
func (o *UpdatePaymentMethodInfo) SetCountries(v []string) {
	o.Countries = v
}

// GetCup returns the Cup field value if set, zero value otherwise.
func (o *UpdatePaymentMethodInfo) GetCup() GenericPmWithTdiInfo {
	if o == nil || common.IsNil(o.Cup) {
		var ret GenericPmWithTdiInfo
		return ret
	}
	return *o.Cup
}

// GetCupOk returns a tuple with the Cup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePaymentMethodInfo) GetCupOk() (*GenericPmWithTdiInfo, bool) {
	if o == nil || common.IsNil(o.Cup) {
		return nil, false
	}
	return o.Cup, true
}

// HasCup returns a boolean if a field has been set.
func (o *UpdatePaymentMethodInfo) HasCup() bool {
	if o != nil && !common.IsNil(o.Cup) {
		return true
	}

	return false
}

// SetCup gets a reference to the given GenericPmWithTdiInfo and assigns it to the Cup field.
func (o *UpdatePaymentMethodInfo) SetCup(v GenericPmWithTdiInfo) {
	o.Cup = &v
}

// GetCurrencies returns the Currencies field value if set, zero value otherwise.
func (o *UpdatePaymentMethodInfo) GetCurrencies() []string {
	if o == nil || common.IsNil(o.Currencies) {
		var ret []string
		return ret
	}
	return o.Currencies
}

// GetCurrenciesOk returns a tuple with the Currencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePaymentMethodInfo) GetCurrenciesOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Currencies) {
		return nil, false
	}
	return o.Currencies, true
}

// HasCurrencies returns a boolean if a field has been set.
func (o *UpdatePaymentMethodInfo) HasCurrencies() bool {
	if o != nil && !common.IsNil(o.Currencies) {
		return true
	}

	return false
}

// SetCurrencies gets a reference to the given []string and assigns it to the Currencies field.
func (o *UpdatePaymentMethodInfo) SetCurrencies(v []string) {
	o.Currencies = v
}

// GetCustomRoutingFlags returns the CustomRoutingFlags field value if set, zero value otherwise.
func (o *UpdatePaymentMethodInfo) GetCustomRoutingFlags() []string {
	if o == nil || common.IsNil(o.CustomRoutingFlags) {
		var ret []string
		return ret
	}
	return o.CustomRoutingFlags
}

// GetCustomRoutingFlagsOk returns a tuple with the CustomRoutingFlags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePaymentMethodInfo) GetCustomRoutingFlagsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.CustomRoutingFlags) {
		return nil, false
	}
	return o.CustomRoutingFlags, true
}

// HasCustomRoutingFlags returns a boolean if a field has been set.
func (o *UpdatePaymentMethodInfo) HasCustomRoutingFlags() bool {
	if o != nil && !common.IsNil(o.CustomRoutingFlags) {
		return true
	}

	return false
}

// SetCustomRoutingFlags gets a reference to the given []string and assigns it to the CustomRoutingFlags field.
func (o *UpdatePaymentMethodInfo) SetCustomRoutingFlags(v []string) {
	o.CustomRoutingFlags = v
}

// GetDiners returns the Diners field value if set, zero value otherwise.
func (o *UpdatePaymentMethodInfo) GetDiners() GenericPmWithTdiInfo {
	if o == nil || common.IsNil(o.Diners) {
		var ret GenericPmWithTdiInfo
		return ret
	}
	return *o.Diners
}

// GetDinersOk returns a tuple with the Diners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePaymentMethodInfo) GetDinersOk() (*GenericPmWithTdiInfo, bool) {
	if o == nil || common.IsNil(o.Diners) {
		return nil, false
	}
	return o.Diners, true
}

// HasDiners returns a boolean if a field has been set.
func (o *UpdatePaymentMethodInfo) HasDiners() bool {
	if o != nil && !common.IsNil(o.Diners) {
		return true
	}

	return false
}

// SetDiners gets a reference to the given GenericPmWithTdiInfo and assigns it to the Diners field.
func (o *UpdatePaymentMethodInfo) SetDiners(v GenericPmWithTdiInfo) {
	o.Diners = &v
}

// GetDiscover returns the Discover field value if set, zero value otherwise.
func (o *UpdatePaymentMethodInfo) GetDiscover() GenericPmWithTdiInfo {
	if o == nil || common.IsNil(o.Discover) {
		var ret GenericPmWithTdiInfo
		return ret
	}
	return *o.Discover
}

// GetDiscoverOk returns a tuple with the Discover field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePaymentMethodInfo) GetDiscoverOk() (*GenericPmWithTdiInfo, bool) {
	if o == nil || common.IsNil(o.Discover) {
		return nil, false
	}
	return o.Discover, true
}

// HasDiscover returns a boolean if a field has been set.
func (o *UpdatePaymentMethodInfo) HasDiscover() bool {
	if o != nil && !common.IsNil(o.Discover) {
		return true
	}

	return false
}

// SetDiscover gets a reference to the given GenericPmWithTdiInfo and assigns it to the Discover field.
func (o *UpdatePaymentMethodInfo) SetDiscover(v GenericPmWithTdiInfo) {
	o.Discover = &v
}

// GetEftposAustralia returns the EftposAustralia field value if set, zero value otherwise.
func (o *UpdatePaymentMethodInfo) GetEftposAustralia() GenericPmWithTdiInfo {
	if o == nil || common.IsNil(o.EftposAustralia) {
		var ret GenericPmWithTdiInfo
		return ret
	}
	return *o.EftposAustralia
}

// GetEftposAustraliaOk returns a tuple with the EftposAustralia field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePaymentMethodInfo) GetEftposAustraliaOk() (*GenericPmWithTdiInfo, bool) {
	if o == nil || common.IsNil(o.EftposAustralia) {
		return nil, false
	}
	return o.EftposAustralia, true
}

// HasEftposAustralia returns a boolean if a field has been set.
func (o *UpdatePaymentMethodInfo) HasEftposAustralia() bool {
	if o != nil && !common.IsNil(o.EftposAustralia) {
		return true
	}

	return false
}

// SetEftposAustralia gets a reference to the given GenericPmWithTdiInfo and assigns it to the EftposAustralia field.
func (o *UpdatePaymentMethodInfo) SetEftposAustralia(v GenericPmWithTdiInfo) {
	o.EftposAustralia = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdatePaymentMethodInfo) GetEnabled() bool {
	if o == nil || common.IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePaymentMethodInfo) GetEnabledOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdatePaymentMethodInfo) HasEnabled() bool {
	if o != nil && !common.IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdatePaymentMethodInfo) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetGirocard returns the Girocard field value if set, zero value otherwise.
func (o *UpdatePaymentMethodInfo) GetGirocard() GenericPmWithTdiInfo {
	if o == nil || common.IsNil(o.Girocard) {
		var ret GenericPmWithTdiInfo
		return ret
	}
	return *o.Girocard
}

// GetGirocardOk returns a tuple with the Girocard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePaymentMethodInfo) GetGirocardOk() (*GenericPmWithTdiInfo, bool) {
	if o == nil || common.IsNil(o.Girocard) {
		return nil, false
	}
	return o.Girocard, true
}

// HasGirocard returns a boolean if a field has been set.
func (o *UpdatePaymentMethodInfo) HasGirocard() bool {
	if o != nil && !common.IsNil(o.Girocard) {
		return true
	}

	return false
}

// SetGirocard gets a reference to the given GenericPmWithTdiInfo and assigns it to the Girocard field.
func (o *UpdatePaymentMethodInfo) SetGirocard(v GenericPmWithTdiInfo) {
	o.Girocard = &v
}

// GetIdeal returns the Ideal field value if set, zero value otherwise.
func (o *UpdatePaymentMethodInfo) GetIdeal() GenericPmWithTdiInfo {
	if o == nil || common.IsNil(o.Ideal) {
		var ret GenericPmWithTdiInfo
		return ret
	}
	return *o.Ideal
}

// GetIdealOk returns a tuple with the Ideal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePaymentMethodInfo) GetIdealOk() (*GenericPmWithTdiInfo, bool) {
	if o == nil || common.IsNil(o.Ideal) {
		return nil, false
	}
	return o.Ideal, true
}

// HasIdeal returns a boolean if a field has been set.
func (o *UpdatePaymentMethodInfo) HasIdeal() bool {
	if o != nil && !common.IsNil(o.Ideal) {
		return true
	}

	return false
}

// SetIdeal gets a reference to the given GenericPmWithTdiInfo and assigns it to the Ideal field.
func (o *UpdatePaymentMethodInfo) SetIdeal(v GenericPmWithTdiInfo) {
	o.Ideal = &v
}

// GetInteracCard returns the InteracCard field value if set, zero value otherwise.
func (o *UpdatePaymentMethodInfo) GetInteracCard() GenericPmWithTdiInfo {
	if o == nil || common.IsNil(o.InteracCard) {
		var ret GenericPmWithTdiInfo
		return ret
	}
	return *o.InteracCard
}

// GetInteracCardOk returns a tuple with the InteracCard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePaymentMethodInfo) GetInteracCardOk() (*GenericPmWithTdiInfo, bool) {
	if o == nil || common.IsNil(o.InteracCard) {
		return nil, false
	}
	return o.InteracCard, true
}

// HasInteracCard returns a boolean if a field has been set.
func (o *UpdatePaymentMethodInfo) HasInteracCard() bool {
	if o != nil && !common.IsNil(o.InteracCard) {
		return true
	}

	return false
}

// SetInteracCard gets a reference to the given GenericPmWithTdiInfo and assigns it to the InteracCard field.
func (o *UpdatePaymentMethodInfo) SetInteracCard(v GenericPmWithTdiInfo) {
	o.InteracCard = &v
}

// GetJcb returns the Jcb field value if set, zero value otherwise.
func (o *UpdatePaymentMethodInfo) GetJcb() GenericPmWithTdiInfo {
	if o == nil || common.IsNil(o.Jcb) {
		var ret GenericPmWithTdiInfo
		return ret
	}
	return *o.Jcb
}

// GetJcbOk returns a tuple with the Jcb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePaymentMethodInfo) GetJcbOk() (*GenericPmWithTdiInfo, bool) {
	if o == nil || common.IsNil(o.Jcb) {
		return nil, false
	}
	return o.Jcb, true
}

// HasJcb returns a boolean if a field has been set.
func (o *UpdatePaymentMethodInfo) HasJcb() bool {
	if o != nil && !common.IsNil(o.Jcb) {
		return true
	}

	return false
}

// SetJcb gets a reference to the given GenericPmWithTdiInfo and assigns it to the Jcb field.
func (o *UpdatePaymentMethodInfo) SetJcb(v GenericPmWithTdiInfo) {
	o.Jcb = &v
}

// GetMaestro returns the Maestro field value if set, zero value otherwise.
func (o *UpdatePaymentMethodInfo) GetMaestro() GenericPmWithTdiInfo {
	if o == nil || common.IsNil(o.Maestro) {
		var ret GenericPmWithTdiInfo
		return ret
	}
	return *o.Maestro
}

// GetMaestroOk returns a tuple with the Maestro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePaymentMethodInfo) GetMaestroOk() (*GenericPmWithTdiInfo, bool) {
	if o == nil || common.IsNil(o.Maestro) {
		return nil, false
	}
	return o.Maestro, true
}

// HasMaestro returns a boolean if a field has been set.
func (o *UpdatePaymentMethodInfo) HasMaestro() bool {
	if o != nil && !common.IsNil(o.Maestro) {
		return true
	}

	return false
}

// SetMaestro gets a reference to the given GenericPmWithTdiInfo and assigns it to the Maestro field.
func (o *UpdatePaymentMethodInfo) SetMaestro(v GenericPmWithTdiInfo) {
	o.Maestro = &v
}

// GetMc returns the Mc field value if set, zero value otherwise.
func (o *UpdatePaymentMethodInfo) GetMc() GenericPmWithTdiInfo {
	if o == nil || common.IsNil(o.Mc) {
		var ret GenericPmWithTdiInfo
		return ret
	}
	return *o.Mc
}

// GetMcOk returns a tuple with the Mc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePaymentMethodInfo) GetMcOk() (*GenericPmWithTdiInfo, bool) {
	if o == nil || common.IsNil(o.Mc) {
		return nil, false
	}
	return o.Mc, true
}

// HasMc returns a boolean if a field has been set.
func (o *UpdatePaymentMethodInfo) HasMc() bool {
	if o != nil && !common.IsNil(o.Mc) {
		return true
	}

	return false
}

// SetMc gets a reference to the given GenericPmWithTdiInfo and assigns it to the Mc field.
func (o *UpdatePaymentMethodInfo) SetMc(v GenericPmWithTdiInfo) {
	o.Mc = &v
}

// GetStoreIds returns the StoreIds field value if set, zero value otherwise.
func (o *UpdatePaymentMethodInfo) GetStoreIds() []string {
	if o == nil || common.IsNil(o.StoreIds) {
		var ret []string
		return ret
	}
	return o.StoreIds
}

// GetStoreIdsOk returns a tuple with the StoreIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePaymentMethodInfo) GetStoreIdsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.StoreIds) {
		return nil, false
	}
	return o.StoreIds, true
}

// HasStoreIds returns a boolean if a field has been set.
func (o *UpdatePaymentMethodInfo) HasStoreIds() bool {
	if o != nil && !common.IsNil(o.StoreIds) {
		return true
	}

	return false
}

// SetStoreIds gets a reference to the given []string and assigns it to the StoreIds field.
func (o *UpdatePaymentMethodInfo) SetStoreIds(v []string) {
	o.StoreIds = v
}

// GetVisa returns the Visa field value if set, zero value otherwise.
func (o *UpdatePaymentMethodInfo) GetVisa() GenericPmWithTdiInfo {
	if o == nil || common.IsNil(o.Visa) {
		var ret GenericPmWithTdiInfo
		return ret
	}
	return *o.Visa
}

// GetVisaOk returns a tuple with the Visa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePaymentMethodInfo) GetVisaOk() (*GenericPmWithTdiInfo, bool) {
	if o == nil || common.IsNil(o.Visa) {
		return nil, false
	}
	return o.Visa, true
}

// HasVisa returns a boolean if a field has been set.
func (o *UpdatePaymentMethodInfo) HasVisa() bool {
	if o != nil && !common.IsNil(o.Visa) {
		return true
	}

	return false
}

// SetVisa gets a reference to the given GenericPmWithTdiInfo and assigns it to the Visa field.
func (o *UpdatePaymentMethodInfo) SetVisa(v GenericPmWithTdiInfo) {
	o.Visa = &v
}

func (o UpdatePaymentMethodInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdatePaymentMethodInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Bcmc) {
		toSerialize["bcmc"] = o.Bcmc
	}
	if !common.IsNil(o.CartesBancaires) {
		toSerialize["cartesBancaires"] = o.CartesBancaires
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
	if !common.IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !common.IsNil(o.Girocard) {
		toSerialize["girocard"] = o.Girocard
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
	if !common.IsNil(o.Maestro) {
		toSerialize["maestro"] = o.Maestro
	}
	if !common.IsNil(o.Mc) {
		toSerialize["mc"] = o.Mc
	}
	if !common.IsNil(o.StoreIds) {
		toSerialize["storeIds"] = o.StoreIds
	}
	if !common.IsNil(o.Visa) {
		toSerialize["visa"] = o.Visa
	}
	return toSerialize, nil
}

type NullableUpdatePaymentMethodInfo struct {
	value *UpdatePaymentMethodInfo
	isSet bool
}

func (v NullableUpdatePaymentMethodInfo) Get() *UpdatePaymentMethodInfo {
	return v.value
}

func (v *NullableUpdatePaymentMethodInfo) Set(val *UpdatePaymentMethodInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePaymentMethodInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePaymentMethodInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePaymentMethodInfo(val *UpdatePaymentMethodInfo) *NullableUpdatePaymentMethodInfo {
	return &NullableUpdatePaymentMethodInfo{value: val, isSet: true}
}

func (v NullableUpdatePaymentMethodInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatePaymentMethodInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
