/*
Transfers API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transfers

import (
	"encoding/json"
	"time"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// checks if the Transfer type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Transfer{}

// Transfer struct for Transfer
type Transfer struct {
	AccountHolder  *ResourceReference `json:"accountHolder,omitempty"`
	Amount         Amount             `json:"amount"`
	BalanceAccount *ResourceReference `json:"balanceAccount,omitempty"`
	// The unique identifier of the source [balance account](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/balanceAccounts__resParam_id).
	// Deprecated
	BalanceAccountId *string `json:"balanceAccountId,omitempty"`
	// The type of transfer.  Possible values:   - **bank**: Transfer to a [transfer instrument](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/transferInstruments__resParam_id) or a bank account.  - **internal**: Transfer to another [balance account](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/balanceAccounts__resParam_id) within your platform.  - **issuedCard**: Transfer initiated by a Adyen-issued card.  - **platformPayment**: Fund movements related to payments that are acquired for your users.
	Category     string         `json:"category"`
	Counterparty CounterpartyV3 `json:"counterparty"`
	// The date and time when the event was triggered, in ISO 8601 extended format. For example, **2020-12-18T10:15:30+01:00**.
	CreationDate *time.Time `json:"creationDate,omitempty"`
	// Your description for the transfer. It is used by most banks as the transfer description. We recommend sending a maximum of 140 characters, otherwise the description may be truncated.  Supported characters: **[a-z] [A-Z] [0-9] / - ?** **: ( ) . , ' + Space**  Supported characters for **regular** and **fast** transfers to a US counterparty: **[a-z] [A-Z] [0-9] & $ % # @** **~ = + - _ ' \" ! ?**
	Description *string `json:"description,omitempty"`
	// The direction of the transfer.  Possible values: **incoming**, **outgoing**.
	Direction *string `json:"direction,omitempty"`
	// The ID of the resource.
	Id                *string            `json:"id,omitempty"`
	PaymentInstrument *PaymentInstrument `json:"paymentInstrument,omitempty"`
	// The unique identifier of the [payment instrument](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/balanceAccounts__resParam_id) used in the transfer.
	// Deprecated
	PaymentInstrumentId *string `json:"paymentInstrumentId,omitempty"`
	// The priority for the bank transfer. This sets the speed at which the transfer is sent and the fees that you have to pay. Required for transfers with `category` **bank**.  Possible values:  * **regular**: For normal, low-value transactions.  * **fast**: Faster way to transfer funds but has higher fees. Recommended for high-priority, low-value transactions.  * **wire**: Fastest way to transfer funds but has the highest fees. Recommended for high-priority, high-value transactions.  * **instant**: Instant way to transfer funds in [SEPA countries](https://www.ecb.europa.eu/paym/integration/retail/sepa/html/index.en.html).  * **crossBorder**: High-value transfer to a recipient in a different country.  * **internal**: Transfer to an Adyen-issued business bank account (by bank account number/IBAN).
	Priority *string `json:"priority,omitempty"`
	// Additional information about the status of the transfer.
	Reason *string `json:"reason,omitempty"`
	// Your reference for the transfer, used internally within your platform. If you don't provide this in the request, Adyen generates a unique reference.
	Reference *string `json:"reference,omitempty"`
	//  A reference that is sent to the recipient. This reference is also sent in all webhooks related to the transfer, so you can use it to track statuses for both the source and recipient of funds.   Supported characters: **a-z**, **A-Z**, **0-9**. The maximum length depends on the `category`.  - **internal**: 80 characters  - **bank**: 35 characters when transferring to an IBAN, 15 characters for others.
	ReferenceForBeneficiary *string `json:"referenceForBeneficiary,omitempty"`
	// The result of the transfer.   For example, **authorised**, **refused**, or **error**.
	Status string `json:"status"`
}

// NewTransfer instantiates a new Transfer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransfer(amount Amount, category string, counterparty CounterpartyV3, status string) *Transfer {
	this := Transfer{}
	this.Amount = amount
	this.Category = category
	this.Counterparty = counterparty
	this.Status = status
	return &this
}

// NewTransferWithDefaults instantiates a new Transfer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferWithDefaults() *Transfer {
	this := Transfer{}
	return &this
}

// GetAccountHolder returns the AccountHolder field value if set, zero value otherwise.
func (o *Transfer) GetAccountHolder() ResourceReference {
	if o == nil || common.IsNil(o.AccountHolder) {
		var ret ResourceReference
		return ret
	}
	return *o.AccountHolder
}

// GetAccountHolderOk returns a tuple with the AccountHolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetAccountHolderOk() (*ResourceReference, bool) {
	if o == nil || common.IsNil(o.AccountHolder) {
		return nil, false
	}
	return o.AccountHolder, true
}

// HasAccountHolder returns a boolean if a field has been set.
func (o *Transfer) HasAccountHolder() bool {
	if o != nil && !common.IsNil(o.AccountHolder) {
		return true
	}

	return false
}

// SetAccountHolder gets a reference to the given ResourceReference and assigns it to the AccountHolder field.
func (o *Transfer) SetAccountHolder(v ResourceReference) {
	o.AccountHolder = &v
}

// GetAmount returns the Amount field value
func (o *Transfer) GetAmount() Amount {
	if o == nil {
		var ret Amount
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *Transfer) GetAmountOk() (*Amount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *Transfer) SetAmount(v Amount) {
	o.Amount = v
}

// GetBalanceAccount returns the BalanceAccount field value if set, zero value otherwise.
func (o *Transfer) GetBalanceAccount() ResourceReference {
	if o == nil || common.IsNil(o.BalanceAccount) {
		var ret ResourceReference
		return ret
	}
	return *o.BalanceAccount
}

// GetBalanceAccountOk returns a tuple with the BalanceAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetBalanceAccountOk() (*ResourceReference, bool) {
	if o == nil || common.IsNil(o.BalanceAccount) {
		return nil, false
	}
	return o.BalanceAccount, true
}

// HasBalanceAccount returns a boolean if a field has been set.
func (o *Transfer) HasBalanceAccount() bool {
	if o != nil && !common.IsNil(o.BalanceAccount) {
		return true
	}

	return false
}

// SetBalanceAccount gets a reference to the given ResourceReference and assigns it to the BalanceAccount field.
func (o *Transfer) SetBalanceAccount(v ResourceReference) {
	o.BalanceAccount = &v
}

// GetBalanceAccountId returns the BalanceAccountId field value if set, zero value otherwise.
// Deprecated
func (o *Transfer) GetBalanceAccountId() string {
	if o == nil || common.IsNil(o.BalanceAccountId) {
		var ret string
		return ret
	}
	return *o.BalanceAccountId
}

// GetBalanceAccountIdOk returns a tuple with the BalanceAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *Transfer) GetBalanceAccountIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.BalanceAccountId) {
		return nil, false
	}
	return o.BalanceAccountId, true
}

// HasBalanceAccountId returns a boolean if a field has been set.
func (o *Transfer) HasBalanceAccountId() bool {
	if o != nil && !common.IsNil(o.BalanceAccountId) {
		return true
	}

	return false
}

// SetBalanceAccountId gets a reference to the given string and assigns it to the BalanceAccountId field.
// Deprecated
func (o *Transfer) SetBalanceAccountId(v string) {
	o.BalanceAccountId = &v
}

// GetCategory returns the Category field value
func (o *Transfer) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *Transfer) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *Transfer) SetCategory(v string) {
	o.Category = v
}

// GetCounterparty returns the Counterparty field value
func (o *Transfer) GetCounterparty() CounterpartyV3 {
	if o == nil {
		var ret CounterpartyV3
		return ret
	}

	return o.Counterparty
}

// GetCounterpartyOk returns a tuple with the Counterparty field value
// and a boolean to check if the value has been set.
func (o *Transfer) GetCounterpartyOk() (*CounterpartyV3, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Counterparty, true
}

// SetCounterparty sets field value
func (o *Transfer) SetCounterparty(v CounterpartyV3) {
	o.Counterparty = v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *Transfer) GetCreationDate() time.Time {
	if o == nil || common.IsNil(o.CreationDate) {
		var ret time.Time
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetCreationDateOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.CreationDate) {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *Transfer) HasCreationDate() bool {
	if o != nil && !common.IsNil(o.CreationDate) {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given time.Time and assigns it to the CreationDate field.
func (o *Transfer) SetCreationDate(v time.Time) {
	o.CreationDate = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Transfer) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Transfer) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Transfer) SetDescription(v string) {
	o.Description = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *Transfer) GetDirection() string {
	if o == nil || common.IsNil(o.Direction) {
		var ret string
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetDirectionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *Transfer) HasDirection() bool {
	if o != nil && !common.IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given string and assigns it to the Direction field.
func (o *Transfer) SetDirection(v string) {
	o.Direction = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Transfer) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Transfer) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Transfer) SetId(v string) {
	o.Id = &v
}

// GetPaymentInstrument returns the PaymentInstrument field value if set, zero value otherwise.
func (o *Transfer) GetPaymentInstrument() PaymentInstrument {
	if o == nil || common.IsNil(o.PaymentInstrument) {
		var ret PaymentInstrument
		return ret
	}
	return *o.PaymentInstrument
}

// GetPaymentInstrumentOk returns a tuple with the PaymentInstrument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetPaymentInstrumentOk() (*PaymentInstrument, bool) {
	if o == nil || common.IsNil(o.PaymentInstrument) {
		return nil, false
	}
	return o.PaymentInstrument, true
}

// HasPaymentInstrument returns a boolean if a field has been set.
func (o *Transfer) HasPaymentInstrument() bool {
	if o != nil && !common.IsNil(o.PaymentInstrument) {
		return true
	}

	return false
}

// SetPaymentInstrument gets a reference to the given PaymentInstrument and assigns it to the PaymentInstrument field.
func (o *Transfer) SetPaymentInstrument(v PaymentInstrument) {
	o.PaymentInstrument = &v
}

// GetPaymentInstrumentId returns the PaymentInstrumentId field value if set, zero value otherwise.
// Deprecated
func (o *Transfer) GetPaymentInstrumentId() string {
	if o == nil || common.IsNil(o.PaymentInstrumentId) {
		var ret string
		return ret
	}
	return *o.PaymentInstrumentId
}

// GetPaymentInstrumentIdOk returns a tuple with the PaymentInstrumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *Transfer) GetPaymentInstrumentIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.PaymentInstrumentId) {
		return nil, false
	}
	return o.PaymentInstrumentId, true
}

// HasPaymentInstrumentId returns a boolean if a field has been set.
func (o *Transfer) HasPaymentInstrumentId() bool {
	if o != nil && !common.IsNil(o.PaymentInstrumentId) {
		return true
	}

	return false
}

// SetPaymentInstrumentId gets a reference to the given string and assigns it to the PaymentInstrumentId field.
// Deprecated
func (o *Transfer) SetPaymentInstrumentId(v string) {
	o.PaymentInstrumentId = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *Transfer) GetPriority() string {
	if o == nil || common.IsNil(o.Priority) {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetPriorityOk() (*string, bool) {
	if o == nil || common.IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *Transfer) HasPriority() bool {
	if o != nil && !common.IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *Transfer) SetPriority(v string) {
	o.Priority = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *Transfer) GetReason() string {
	if o == nil || common.IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetReasonOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *Transfer) HasReason() bool {
	if o != nil && !common.IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *Transfer) SetReason(v string) {
	o.Reason = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *Transfer) GetReference() string {
	if o == nil || common.IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *Transfer) HasReference() bool {
	if o != nil && !common.IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *Transfer) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceForBeneficiary returns the ReferenceForBeneficiary field value if set, zero value otherwise.
func (o *Transfer) GetReferenceForBeneficiary() string {
	if o == nil || common.IsNil(o.ReferenceForBeneficiary) {
		var ret string
		return ret
	}
	return *o.ReferenceForBeneficiary
}

// GetReferenceForBeneficiaryOk returns a tuple with the ReferenceForBeneficiary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetReferenceForBeneficiaryOk() (*string, bool) {
	if o == nil || common.IsNil(o.ReferenceForBeneficiary) {
		return nil, false
	}
	return o.ReferenceForBeneficiary, true
}

// HasReferenceForBeneficiary returns a boolean if a field has been set.
func (o *Transfer) HasReferenceForBeneficiary() bool {
	if o != nil && !common.IsNil(o.ReferenceForBeneficiary) {
		return true
	}

	return false
}

// SetReferenceForBeneficiary gets a reference to the given string and assigns it to the ReferenceForBeneficiary field.
func (o *Transfer) SetReferenceForBeneficiary(v string) {
	o.ReferenceForBeneficiary = &v
}

// GetStatus returns the Status field value
func (o *Transfer) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Transfer) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Transfer) SetStatus(v string) {
	o.Status = v
}

func (o Transfer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Transfer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AccountHolder) {
		toSerialize["accountHolder"] = o.AccountHolder
	}
	toSerialize["amount"] = o.Amount
	if !common.IsNil(o.BalanceAccount) {
		toSerialize["balanceAccount"] = o.BalanceAccount
	}
	if !common.IsNil(o.BalanceAccountId) {
		toSerialize["balanceAccountId"] = o.BalanceAccountId
	}
	toSerialize["category"] = o.Category
	toSerialize["counterparty"] = o.Counterparty
	if !common.IsNil(o.CreationDate) {
		toSerialize["creationDate"] = o.CreationDate
	}
	if !common.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !common.IsNil(o.Direction) {
		toSerialize["direction"] = o.Direction
	}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.PaymentInstrument) {
		toSerialize["paymentInstrument"] = o.PaymentInstrument
	}
	if !common.IsNil(o.PaymentInstrumentId) {
		toSerialize["paymentInstrumentId"] = o.PaymentInstrumentId
	}
	if !common.IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !common.IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !common.IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !common.IsNil(o.ReferenceForBeneficiary) {
		toSerialize["referenceForBeneficiary"] = o.ReferenceForBeneficiary
	}
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

type NullableTransfer struct {
	value *Transfer
	isSet bool
}

func (v NullableTransfer) Get() *Transfer {
	return v.value
}

func (v *NullableTransfer) Set(val *Transfer) {
	v.value = val
	v.isSet = true
}

func (v NullableTransfer) IsSet() bool {
	return v.isSet
}

func (v *NullableTransfer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransfer(val *Transfer) *NullableTransfer {
	return &NullableTransfer{value: val, isSet: true}
}

func (v NullableTransfer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransfer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *Transfer) isValidCategory() bool {
	var allowedEnumValues = []string{"bank", "internal", "issuedCard", "platformPayment"}
	for _, allowed := range allowedEnumValues {
		if o.GetCategory() == allowed {
			return true
		}
	}
	return false
}
func (o *Transfer) isValidDirection() bool {
	var allowedEnumValues = []string{"incoming", "outgoing"}
	for _, allowed := range allowedEnumValues {
		if o.GetDirection() == allowed {
			return true
		}
	}
	return false
}
func (o *Transfer) isValidPriority() bool {
	var allowedEnumValues = []string{"crossBorder", "directDebit", "fast", "instant", "internal", "regular", "wire"}
	for _, allowed := range allowedEnumValues {
		if o.GetPriority() == allowed {
			return true
		}
	}
	return false
}
func (o *Transfer) isValidReason() bool {
	var allowedEnumValues = []string{"amountLimitExceeded", "approved", "counterpartyAccountBlocked", "counterpartyAccountClosed", "counterpartyAccountNotFound", "counterpartyAddressRequired", "counterpartyBankTimedOut", "counterpartyBankUnavailable", "error", "notEnoughBalance", "refusedByCounterpartyBank", "routeNotFound", "unknown"}
	for _, allowed := range allowedEnumValues {
		if o.GetReason() == allowed {
			return true
		}
	}
	return false
}
func (o *Transfer) isValidStatus() bool {
	var allowedEnumValues = []string{"approvalPending", "atmWithdrawal", "atmWithdrawalReversalPending", "atmWithdrawalReversed", "authAdjustmentAuthorised", "authAdjustmentError", "authAdjustmentRefused", "authorised", "bankTransfer", "bankTransferPending", "booked", "bookingPending", "cancelled", "capturePending", "captureReversalPending", "captureReversed", "captured", "capturedExternally", "chargeback", "chargebackExternally", "chargebackPending", "chargebackReversalPending", "chargebackReversed", "credited", "depositCorrection", "depositCorrectionPending", "dispute", "disputeClosed", "disputeExpired", "disputeNeedsReview", "error", "expired", "failed", "fee", "feePending", "internalTransfer", "internalTransferPending", "invoiceDeduction", "invoiceDeductionPending", "manualCorrectionPending", "manuallyCorrected", "matchedStatement", "matchedStatementPending", "merchantPayin", "merchantPayinPending", "merchantPayinReversed", "merchantPayinReversedPending", "miscCost", "miscCostPending", "operationAuthorized", "operationBooked", "operationPending", "operationReceived", "paymentCost", "paymentCostPending", "received", "refundPending", "refundReversalPending", "refundReversed", "refunded", "refundedExternally", "refused", "reserveAdjustment", "reserveAdjustmentPending", "returned", "secondChargeback", "secondChargebackPending", "undefined"}
	for _, allowed := range allowedEnumValues {
		if o.GetStatus() == allowed {
			return true
		}
	}
	return false
}
