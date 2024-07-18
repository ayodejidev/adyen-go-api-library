/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v11/src/common"
)

// checks if the TransactionRuleRestrictions type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TransactionRuleRestrictions{}

// TransactionRuleRestrictions struct for TransactionRuleRestrictions
type TransactionRuleRestrictions struct {
	ActiveNetworkTokens         *ActiveNetworkTokensRestriction      `json:"activeNetworkTokens,omitempty"`
	BrandVariants               *BrandVariantsRestriction            `json:"brandVariants,omitempty"`
	CounterpartyBank            *CounterpartyBankRestriction         `json:"counterpartyBank,omitempty"`
	Countries                   *CountriesRestriction                `json:"countries,omitempty"`
	DayOfWeek                   *DayOfWeekRestriction                `json:"dayOfWeek,omitempty"`
	DifferentCurrencies         *DifferentCurrenciesRestriction      `json:"differentCurrencies,omitempty"`
	EntryModes                  *EntryModesRestriction               `json:"entryModes,omitempty"`
	InternationalTransaction    *InternationalTransactionRestriction `json:"internationalTransaction,omitempty"`
	MatchingTransactions        *MatchingTransactionsRestriction     `json:"matchingTransactions,omitempty"`
	MatchingValues              *MatchingValuesRestriction           `json:"matchingValues,omitempty"`
	Mccs                        *MccsRestriction                     `json:"mccs,omitempty"`
	MerchantNames               *MerchantNamesRestriction            `json:"merchantNames,omitempty"`
	Merchants                   *MerchantsRestriction                `json:"merchants,omitempty"`
	ProcessingTypes             *ProcessingTypesRestriction          `json:"processingTypes,omitempty"`
	RiskScores                  *RiskScoresRestriction               `json:"riskScores,omitempty"`
	SameAmountRestriction       *SameAmountRestriction               `json:"sameAmountRestriction,omitempty"`
	SameCounterpartyRestriction *SameCounterpartyRestriction         `json:"sameCounterpartyRestriction,omitempty"`
	TimeOfDay                   *TimeOfDayRestriction                `json:"timeOfDay,omitempty"`
	TotalAmount                 *TotalAmountRestriction              `json:"totalAmount,omitempty"`
}

// NewTransactionRuleRestrictions instantiates a new TransactionRuleRestrictions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionRuleRestrictions() *TransactionRuleRestrictions {
	this := TransactionRuleRestrictions{}
	return &this
}

// NewTransactionRuleRestrictionsWithDefaults instantiates a new TransactionRuleRestrictions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionRuleRestrictionsWithDefaults() *TransactionRuleRestrictions {
	this := TransactionRuleRestrictions{}
	return &this
}

// GetActiveNetworkTokens returns the ActiveNetworkTokens field value if set, zero value otherwise.
func (o *TransactionRuleRestrictions) GetActiveNetworkTokens() ActiveNetworkTokensRestriction {
	if o == nil || common.IsNil(o.ActiveNetworkTokens) {
		var ret ActiveNetworkTokensRestriction
		return ret
	}
	return *o.ActiveNetworkTokens
}

// GetActiveNetworkTokensOk returns a tuple with the ActiveNetworkTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRuleRestrictions) GetActiveNetworkTokensOk() (*ActiveNetworkTokensRestriction, bool) {
	if o == nil || common.IsNil(o.ActiveNetworkTokens) {
		return nil, false
	}
	return o.ActiveNetworkTokens, true
}

// HasActiveNetworkTokens returns a boolean if a field has been set.
func (o *TransactionRuleRestrictions) HasActiveNetworkTokens() bool {
	if o != nil && !common.IsNil(o.ActiveNetworkTokens) {
		return true
	}

	return false
}

// SetActiveNetworkTokens gets a reference to the given ActiveNetworkTokensRestriction and assigns it to the ActiveNetworkTokens field.
func (o *TransactionRuleRestrictions) SetActiveNetworkTokens(v ActiveNetworkTokensRestriction) {
	o.ActiveNetworkTokens = &v
}

// GetBrandVariants returns the BrandVariants field value if set, zero value otherwise.
func (o *TransactionRuleRestrictions) GetBrandVariants() BrandVariantsRestriction {
	if o == nil || common.IsNil(o.BrandVariants) {
		var ret BrandVariantsRestriction
		return ret
	}
	return *o.BrandVariants
}

// GetBrandVariantsOk returns a tuple with the BrandVariants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRuleRestrictions) GetBrandVariantsOk() (*BrandVariantsRestriction, bool) {
	if o == nil || common.IsNil(o.BrandVariants) {
		return nil, false
	}
	return o.BrandVariants, true
}

// HasBrandVariants returns a boolean if a field has been set.
func (o *TransactionRuleRestrictions) HasBrandVariants() bool {
	if o != nil && !common.IsNil(o.BrandVariants) {
		return true
	}

	return false
}

// SetBrandVariants gets a reference to the given BrandVariantsRestriction and assigns it to the BrandVariants field.
func (o *TransactionRuleRestrictions) SetBrandVariants(v BrandVariantsRestriction) {
	o.BrandVariants = &v
}

// GetCounterpartyBank returns the CounterpartyBank field value if set, zero value otherwise.
func (o *TransactionRuleRestrictions) GetCounterpartyBank() CounterpartyBankRestriction {
	if o == nil || common.IsNil(o.CounterpartyBank) {
		var ret CounterpartyBankRestriction
		return ret
	}
	return *o.CounterpartyBank
}

// GetCounterpartyBankOk returns a tuple with the CounterpartyBank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRuleRestrictions) GetCounterpartyBankOk() (*CounterpartyBankRestriction, bool) {
	if o == nil || common.IsNil(o.CounterpartyBank) {
		return nil, false
	}
	return o.CounterpartyBank, true
}

// HasCounterpartyBank returns a boolean if a field has been set.
func (o *TransactionRuleRestrictions) HasCounterpartyBank() bool {
	if o != nil && !common.IsNil(o.CounterpartyBank) {
		return true
	}

	return false
}

// SetCounterpartyBank gets a reference to the given CounterpartyBankRestriction and assigns it to the CounterpartyBank field.
func (o *TransactionRuleRestrictions) SetCounterpartyBank(v CounterpartyBankRestriction) {
	o.CounterpartyBank = &v
}

// GetCountries returns the Countries field value if set, zero value otherwise.
func (o *TransactionRuleRestrictions) GetCountries() CountriesRestriction {
	if o == nil || common.IsNil(o.Countries) {
		var ret CountriesRestriction
		return ret
	}
	return *o.Countries
}

// GetCountriesOk returns a tuple with the Countries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRuleRestrictions) GetCountriesOk() (*CountriesRestriction, bool) {
	if o == nil || common.IsNil(o.Countries) {
		return nil, false
	}
	return o.Countries, true
}

// HasCountries returns a boolean if a field has been set.
func (o *TransactionRuleRestrictions) HasCountries() bool {
	if o != nil && !common.IsNil(o.Countries) {
		return true
	}

	return false
}

// SetCountries gets a reference to the given CountriesRestriction and assigns it to the Countries field.
func (o *TransactionRuleRestrictions) SetCountries(v CountriesRestriction) {
	o.Countries = &v
}

// GetDayOfWeek returns the DayOfWeek field value if set, zero value otherwise.
func (o *TransactionRuleRestrictions) GetDayOfWeek() DayOfWeekRestriction {
	if o == nil || common.IsNil(o.DayOfWeek) {
		var ret DayOfWeekRestriction
		return ret
	}
	return *o.DayOfWeek
}

// GetDayOfWeekOk returns a tuple with the DayOfWeek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRuleRestrictions) GetDayOfWeekOk() (*DayOfWeekRestriction, bool) {
	if o == nil || common.IsNil(o.DayOfWeek) {
		return nil, false
	}
	return o.DayOfWeek, true
}

// HasDayOfWeek returns a boolean if a field has been set.
func (o *TransactionRuleRestrictions) HasDayOfWeek() bool {
	if o != nil && !common.IsNil(o.DayOfWeek) {
		return true
	}

	return false
}

// SetDayOfWeek gets a reference to the given DayOfWeekRestriction and assigns it to the DayOfWeek field.
func (o *TransactionRuleRestrictions) SetDayOfWeek(v DayOfWeekRestriction) {
	o.DayOfWeek = &v
}

// GetDifferentCurrencies returns the DifferentCurrencies field value if set, zero value otherwise.
func (o *TransactionRuleRestrictions) GetDifferentCurrencies() DifferentCurrenciesRestriction {
	if o == nil || common.IsNil(o.DifferentCurrencies) {
		var ret DifferentCurrenciesRestriction
		return ret
	}
	return *o.DifferentCurrencies
}

// GetDifferentCurrenciesOk returns a tuple with the DifferentCurrencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRuleRestrictions) GetDifferentCurrenciesOk() (*DifferentCurrenciesRestriction, bool) {
	if o == nil || common.IsNil(o.DifferentCurrencies) {
		return nil, false
	}
	return o.DifferentCurrencies, true
}

// HasDifferentCurrencies returns a boolean if a field has been set.
func (o *TransactionRuleRestrictions) HasDifferentCurrencies() bool {
	if o != nil && !common.IsNil(o.DifferentCurrencies) {
		return true
	}

	return false
}

// SetDifferentCurrencies gets a reference to the given DifferentCurrenciesRestriction and assigns it to the DifferentCurrencies field.
func (o *TransactionRuleRestrictions) SetDifferentCurrencies(v DifferentCurrenciesRestriction) {
	o.DifferentCurrencies = &v
}

// GetEntryModes returns the EntryModes field value if set, zero value otherwise.
func (o *TransactionRuleRestrictions) GetEntryModes() EntryModesRestriction {
	if o == nil || common.IsNil(o.EntryModes) {
		var ret EntryModesRestriction
		return ret
	}
	return *o.EntryModes
}

// GetEntryModesOk returns a tuple with the EntryModes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRuleRestrictions) GetEntryModesOk() (*EntryModesRestriction, bool) {
	if o == nil || common.IsNil(o.EntryModes) {
		return nil, false
	}
	return o.EntryModes, true
}

// HasEntryModes returns a boolean if a field has been set.
func (o *TransactionRuleRestrictions) HasEntryModes() bool {
	if o != nil && !common.IsNil(o.EntryModes) {
		return true
	}

	return false
}

// SetEntryModes gets a reference to the given EntryModesRestriction and assigns it to the EntryModes field.
func (o *TransactionRuleRestrictions) SetEntryModes(v EntryModesRestriction) {
	o.EntryModes = &v
}

// GetInternationalTransaction returns the InternationalTransaction field value if set, zero value otherwise.
func (o *TransactionRuleRestrictions) GetInternationalTransaction() InternationalTransactionRestriction {
	if o == nil || common.IsNil(o.InternationalTransaction) {
		var ret InternationalTransactionRestriction
		return ret
	}
	return *o.InternationalTransaction
}

// GetInternationalTransactionOk returns a tuple with the InternationalTransaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRuleRestrictions) GetInternationalTransactionOk() (*InternationalTransactionRestriction, bool) {
	if o == nil || common.IsNil(o.InternationalTransaction) {
		return nil, false
	}
	return o.InternationalTransaction, true
}

// HasInternationalTransaction returns a boolean if a field has been set.
func (o *TransactionRuleRestrictions) HasInternationalTransaction() bool {
	if o != nil && !common.IsNil(o.InternationalTransaction) {
		return true
	}

	return false
}

// SetInternationalTransaction gets a reference to the given InternationalTransactionRestriction and assigns it to the InternationalTransaction field.
func (o *TransactionRuleRestrictions) SetInternationalTransaction(v InternationalTransactionRestriction) {
	o.InternationalTransaction = &v
}

// GetMatchingTransactions returns the MatchingTransactions field value if set, zero value otherwise.
func (o *TransactionRuleRestrictions) GetMatchingTransactions() MatchingTransactionsRestriction {
	if o == nil || common.IsNil(o.MatchingTransactions) {
		var ret MatchingTransactionsRestriction
		return ret
	}
	return *o.MatchingTransactions
}

// GetMatchingTransactionsOk returns a tuple with the MatchingTransactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRuleRestrictions) GetMatchingTransactionsOk() (*MatchingTransactionsRestriction, bool) {
	if o == nil || common.IsNil(o.MatchingTransactions) {
		return nil, false
	}
	return o.MatchingTransactions, true
}

// HasMatchingTransactions returns a boolean if a field has been set.
func (o *TransactionRuleRestrictions) HasMatchingTransactions() bool {
	if o != nil && !common.IsNil(o.MatchingTransactions) {
		return true
	}

	return false
}

// SetMatchingTransactions gets a reference to the given MatchingTransactionsRestriction and assigns it to the MatchingTransactions field.
func (o *TransactionRuleRestrictions) SetMatchingTransactions(v MatchingTransactionsRestriction) {
	o.MatchingTransactions = &v
}

// GetMatchingValues returns the MatchingValues field value if set, zero value otherwise.
func (o *TransactionRuleRestrictions) GetMatchingValues() MatchingValuesRestriction {
	if o == nil || common.IsNil(o.MatchingValues) {
		var ret MatchingValuesRestriction
		return ret
	}
	return *o.MatchingValues
}

// GetMatchingValuesOk returns a tuple with the MatchingValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRuleRestrictions) GetMatchingValuesOk() (*MatchingValuesRestriction, bool) {
	if o == nil || common.IsNil(o.MatchingValues) {
		return nil, false
	}
	return o.MatchingValues, true
}

// HasMatchingValues returns a boolean if a field has been set.
func (o *TransactionRuleRestrictions) HasMatchingValues() bool {
	if o != nil && !common.IsNil(o.MatchingValues) {
		return true
	}

	return false
}

// SetMatchingValues gets a reference to the given MatchingValuesRestriction and assigns it to the MatchingValues field.
func (o *TransactionRuleRestrictions) SetMatchingValues(v MatchingValuesRestriction) {
	o.MatchingValues = &v
}

// GetMccs returns the Mccs field value if set, zero value otherwise.
func (o *TransactionRuleRestrictions) GetMccs() MccsRestriction {
	if o == nil || common.IsNil(o.Mccs) {
		var ret MccsRestriction
		return ret
	}
	return *o.Mccs
}

// GetMccsOk returns a tuple with the Mccs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRuleRestrictions) GetMccsOk() (*MccsRestriction, bool) {
	if o == nil || common.IsNil(o.Mccs) {
		return nil, false
	}
	return o.Mccs, true
}

// HasMccs returns a boolean if a field has been set.
func (o *TransactionRuleRestrictions) HasMccs() bool {
	if o != nil && !common.IsNil(o.Mccs) {
		return true
	}

	return false
}

// SetMccs gets a reference to the given MccsRestriction and assigns it to the Mccs field.
func (o *TransactionRuleRestrictions) SetMccs(v MccsRestriction) {
	o.Mccs = &v
}

// GetMerchantNames returns the MerchantNames field value if set, zero value otherwise.
func (o *TransactionRuleRestrictions) GetMerchantNames() MerchantNamesRestriction {
	if o == nil || common.IsNil(o.MerchantNames) {
		var ret MerchantNamesRestriction
		return ret
	}
	return *o.MerchantNames
}

// GetMerchantNamesOk returns a tuple with the MerchantNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRuleRestrictions) GetMerchantNamesOk() (*MerchantNamesRestriction, bool) {
	if o == nil || common.IsNil(o.MerchantNames) {
		return nil, false
	}
	return o.MerchantNames, true
}

// HasMerchantNames returns a boolean if a field has been set.
func (o *TransactionRuleRestrictions) HasMerchantNames() bool {
	if o != nil && !common.IsNil(o.MerchantNames) {
		return true
	}

	return false
}

// SetMerchantNames gets a reference to the given MerchantNamesRestriction and assigns it to the MerchantNames field.
func (o *TransactionRuleRestrictions) SetMerchantNames(v MerchantNamesRestriction) {
	o.MerchantNames = &v
}

// GetMerchants returns the Merchants field value if set, zero value otherwise.
func (o *TransactionRuleRestrictions) GetMerchants() MerchantsRestriction {
	if o == nil || common.IsNil(o.Merchants) {
		var ret MerchantsRestriction
		return ret
	}
	return *o.Merchants
}

// GetMerchantsOk returns a tuple with the Merchants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRuleRestrictions) GetMerchantsOk() (*MerchantsRestriction, bool) {
	if o == nil || common.IsNil(o.Merchants) {
		return nil, false
	}
	return o.Merchants, true
}

// HasMerchants returns a boolean if a field has been set.
func (o *TransactionRuleRestrictions) HasMerchants() bool {
	if o != nil && !common.IsNil(o.Merchants) {
		return true
	}

	return false
}

// SetMerchants gets a reference to the given MerchantsRestriction and assigns it to the Merchants field.
func (o *TransactionRuleRestrictions) SetMerchants(v MerchantsRestriction) {
	o.Merchants = &v
}

// GetProcessingTypes returns the ProcessingTypes field value if set, zero value otherwise.
func (o *TransactionRuleRestrictions) GetProcessingTypes() ProcessingTypesRestriction {
	if o == nil || common.IsNil(o.ProcessingTypes) {
		var ret ProcessingTypesRestriction
		return ret
	}
	return *o.ProcessingTypes
}

// GetProcessingTypesOk returns a tuple with the ProcessingTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRuleRestrictions) GetProcessingTypesOk() (*ProcessingTypesRestriction, bool) {
	if o == nil || common.IsNil(o.ProcessingTypes) {
		return nil, false
	}
	return o.ProcessingTypes, true
}

// HasProcessingTypes returns a boolean if a field has been set.
func (o *TransactionRuleRestrictions) HasProcessingTypes() bool {
	if o != nil && !common.IsNil(o.ProcessingTypes) {
		return true
	}

	return false
}

// SetProcessingTypes gets a reference to the given ProcessingTypesRestriction and assigns it to the ProcessingTypes field.
func (o *TransactionRuleRestrictions) SetProcessingTypes(v ProcessingTypesRestriction) {
	o.ProcessingTypes = &v
}

// GetRiskScores returns the RiskScores field value if set, zero value otherwise.
func (o *TransactionRuleRestrictions) GetRiskScores() RiskScoresRestriction {
	if o == nil || common.IsNil(o.RiskScores) {
		var ret RiskScoresRestriction
		return ret
	}
	return *o.RiskScores
}

// GetRiskScoresOk returns a tuple with the RiskScores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRuleRestrictions) GetRiskScoresOk() (*RiskScoresRestriction, bool) {
	if o == nil || common.IsNil(o.RiskScores) {
		return nil, false
	}
	return o.RiskScores, true
}

// HasRiskScores returns a boolean if a field has been set.
func (o *TransactionRuleRestrictions) HasRiskScores() bool {
	if o != nil && !common.IsNil(o.RiskScores) {
		return true
	}

	return false
}

// SetRiskScores gets a reference to the given RiskScoresRestriction and assigns it to the RiskScores field.
func (o *TransactionRuleRestrictions) SetRiskScores(v RiskScoresRestriction) {
	o.RiskScores = &v
}

// GetSameAmountRestriction returns the SameAmountRestriction field value if set, zero value otherwise.
func (o *TransactionRuleRestrictions) GetSameAmountRestriction() SameAmountRestriction {
	if o == nil || common.IsNil(o.SameAmountRestriction) {
		var ret SameAmountRestriction
		return ret
	}
	return *o.SameAmountRestriction
}

// GetSameAmountRestrictionOk returns a tuple with the SameAmountRestriction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRuleRestrictions) GetSameAmountRestrictionOk() (*SameAmountRestriction, bool) {
	if o == nil || common.IsNil(o.SameAmountRestriction) {
		return nil, false
	}
	return o.SameAmountRestriction, true
}

// HasSameAmountRestriction returns a boolean if a field has been set.
func (o *TransactionRuleRestrictions) HasSameAmountRestriction() bool {
	if o != nil && !common.IsNil(o.SameAmountRestriction) {
		return true
	}

	return false
}

// SetSameAmountRestriction gets a reference to the given SameAmountRestriction and assigns it to the SameAmountRestriction field.
func (o *TransactionRuleRestrictions) SetSameAmountRestriction(v SameAmountRestriction) {
	o.SameAmountRestriction = &v
}

// GetSameCounterpartyRestriction returns the SameCounterpartyRestriction field value if set, zero value otherwise.
func (o *TransactionRuleRestrictions) GetSameCounterpartyRestriction() SameCounterpartyRestriction {
	if o == nil || common.IsNil(o.SameCounterpartyRestriction) {
		var ret SameCounterpartyRestriction
		return ret
	}
	return *o.SameCounterpartyRestriction
}

// GetSameCounterpartyRestrictionOk returns a tuple with the SameCounterpartyRestriction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRuleRestrictions) GetSameCounterpartyRestrictionOk() (*SameCounterpartyRestriction, bool) {
	if o == nil || common.IsNil(o.SameCounterpartyRestriction) {
		return nil, false
	}
	return o.SameCounterpartyRestriction, true
}

// HasSameCounterpartyRestriction returns a boolean if a field has been set.
func (o *TransactionRuleRestrictions) HasSameCounterpartyRestriction() bool {
	if o != nil && !common.IsNil(o.SameCounterpartyRestriction) {
		return true
	}

	return false
}

// SetSameCounterpartyRestriction gets a reference to the given SameCounterpartyRestriction and assigns it to the SameCounterpartyRestriction field.
func (o *TransactionRuleRestrictions) SetSameCounterpartyRestriction(v SameCounterpartyRestriction) {
	o.SameCounterpartyRestriction = &v
}

// GetTimeOfDay returns the TimeOfDay field value if set, zero value otherwise.
func (o *TransactionRuleRestrictions) GetTimeOfDay() TimeOfDayRestriction {
	if o == nil || common.IsNil(o.TimeOfDay) {
		var ret TimeOfDayRestriction
		return ret
	}
	return *o.TimeOfDay
}

// GetTimeOfDayOk returns a tuple with the TimeOfDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRuleRestrictions) GetTimeOfDayOk() (*TimeOfDayRestriction, bool) {
	if o == nil || common.IsNil(o.TimeOfDay) {
		return nil, false
	}
	return o.TimeOfDay, true
}

// HasTimeOfDay returns a boolean if a field has been set.
func (o *TransactionRuleRestrictions) HasTimeOfDay() bool {
	if o != nil && !common.IsNil(o.TimeOfDay) {
		return true
	}

	return false
}

// SetTimeOfDay gets a reference to the given TimeOfDayRestriction and assigns it to the TimeOfDay field.
func (o *TransactionRuleRestrictions) SetTimeOfDay(v TimeOfDayRestriction) {
	o.TimeOfDay = &v
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise.
func (o *TransactionRuleRestrictions) GetTotalAmount() TotalAmountRestriction {
	if o == nil || common.IsNil(o.TotalAmount) {
		var ret TotalAmountRestriction
		return ret
	}
	return *o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRuleRestrictions) GetTotalAmountOk() (*TotalAmountRestriction, bool) {
	if o == nil || common.IsNil(o.TotalAmount) {
		return nil, false
	}
	return o.TotalAmount, true
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *TransactionRuleRestrictions) HasTotalAmount() bool {
	if o != nil && !common.IsNil(o.TotalAmount) {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given TotalAmountRestriction and assigns it to the TotalAmount field.
func (o *TransactionRuleRestrictions) SetTotalAmount(v TotalAmountRestriction) {
	o.TotalAmount = &v
}

func (o TransactionRuleRestrictions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionRuleRestrictions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ActiveNetworkTokens) {
		toSerialize["activeNetworkTokens"] = o.ActiveNetworkTokens
	}
	if !common.IsNil(o.BrandVariants) {
		toSerialize["brandVariants"] = o.BrandVariants
	}
	if !common.IsNil(o.CounterpartyBank) {
		toSerialize["counterpartyBank"] = o.CounterpartyBank
	}
	if !common.IsNil(o.Countries) {
		toSerialize["countries"] = o.Countries
	}
	if !common.IsNil(o.DayOfWeek) {
		toSerialize["dayOfWeek"] = o.DayOfWeek
	}
	if !common.IsNil(o.DifferentCurrencies) {
		toSerialize["differentCurrencies"] = o.DifferentCurrencies
	}
	if !common.IsNil(o.EntryModes) {
		toSerialize["entryModes"] = o.EntryModes
	}
	if !common.IsNil(o.InternationalTransaction) {
		toSerialize["internationalTransaction"] = o.InternationalTransaction
	}
	if !common.IsNil(o.MatchingTransactions) {
		toSerialize["matchingTransactions"] = o.MatchingTransactions
	}
	if !common.IsNil(o.MatchingValues) {
		toSerialize["matchingValues"] = o.MatchingValues
	}
	if !common.IsNil(o.Mccs) {
		toSerialize["mccs"] = o.Mccs
	}
	if !common.IsNil(o.MerchantNames) {
		toSerialize["merchantNames"] = o.MerchantNames
	}
	if !common.IsNil(o.Merchants) {
		toSerialize["merchants"] = o.Merchants
	}
	if !common.IsNil(o.ProcessingTypes) {
		toSerialize["processingTypes"] = o.ProcessingTypes
	}
	if !common.IsNil(o.RiskScores) {
		toSerialize["riskScores"] = o.RiskScores
	}
	if !common.IsNil(o.SameAmountRestriction) {
		toSerialize["sameAmountRestriction"] = o.SameAmountRestriction
	}
	if !common.IsNil(o.SameCounterpartyRestriction) {
		toSerialize["sameCounterpartyRestriction"] = o.SameCounterpartyRestriction
	}
	if !common.IsNil(o.TimeOfDay) {
		toSerialize["timeOfDay"] = o.TimeOfDay
	}
	if !common.IsNil(o.TotalAmount) {
		toSerialize["totalAmount"] = o.TotalAmount
	}
	return toSerialize, nil
}

type NullableTransactionRuleRestrictions struct {
	value *TransactionRuleRestrictions
	isSet bool
}

func (v NullableTransactionRuleRestrictions) Get() *TransactionRuleRestrictions {
	return v.value
}

func (v *NullableTransactionRuleRestrictions) Set(val *TransactionRuleRestrictions) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionRuleRestrictions) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionRuleRestrictions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionRuleRestrictions(val *TransactionRuleRestrictions) *NullableTransactionRuleRestrictions {
	return &NullableTransactionRuleRestrictions{value: val, isSet: true}
}

func (v NullableTransactionRuleRestrictions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionRuleRestrictions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
