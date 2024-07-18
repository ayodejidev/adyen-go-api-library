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

// checks if the PaymentResult type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PaymentResult{}

// PaymentResult struct for PaymentResult
type PaymentResult struct {
	// Contains additional information about the payment. Some data fields are included only if you select them first: Go to **Customer Area** > **Developers** > **Additional data**.
	AdditionalData *map[string]string `json:"additionalData,omitempty"`
	// Authorisation code: * When the payment is authorised successfully, this field holds the authorisation code for the payment. * When the payment is not authorised, this field is empty.
	AuthCode  *string `json:"authCode,omitempty"`
	DccAmount *Amount `json:"dccAmount,omitempty"`
	// Cryptographic signature used to verify `dccQuote`. > This value only applies if you have implemented Dynamic Currency Conversion. For more information, [contact Support](https://www.adyen.help/hc/en-us/requests/new).
	DccSignature *string      `json:"dccSignature,omitempty"`
	FraudResult  *FraudResult `json:"fraudResult,omitempty"`
	// The URL to direct the shopper to. > In case of SecurePlus, do not redirect a shopper to this URL.
	IssuerUrl *string `json:"issuerUrl,omitempty"`
	// The payment session.
	Md *string `json:"md,omitempty"`
	// The 3D request data for the issuer.  If the value is **CUPSecurePlus-CollectSMSVerificationCode**, collect an SMS code from the shopper and pass it in the `/authorise3D` request. For more information, see [3D Secure](https://docs.adyen.com/classic-integration/3d-secure).
	PaRequest *string `json:"paRequest,omitempty"`
	// Adyen's 16-character reference associated with the transaction/request. This value is globally unique; quote it when communicating with us about this request.
	PspReference *string `json:"pspReference,omitempty"`
	// If the payment's authorisation is refused or an error occurs during authorisation, this field holds Adyen's mapped reason for the refusal or a description of the error. When a transaction fails, the authorisation response includes `resultCode` and `refusalReason` values.  For more information, see [Refusal reasons](https://docs.adyen.com/development-resources/refusal-reasons).
	RefusalReason *string `json:"refusalReason,omitempty"`
	// The result of the payment. For more information, see [Result codes](https://docs.adyen.com/online-payments/payment-result-codes).  Possible values:  * **AuthenticationFinished** – The payment has been successfully authenticated with 3D Secure 2. Returned for 3D Secure 2 authentication-only transactions. * **AuthenticationNotRequired** – The transaction does not require 3D Secure authentication. Returned for [standalone authentication-only integrations](https://docs.adyen.com/online-payments/3d-secure/other-3ds-flows/authentication-only). * **Authorised** – The payment was successfully authorised. This state serves as an indicator to proceed with the delivery of goods and services. This is a final state. * **Cancelled** – Indicates the payment has been cancelled (either by the shopper or the merchant) before processing was completed. This is a final state. * **ChallengeShopper** – The issuer requires further shopper interaction before the payment can be authenticated. Returned for 3D Secure 2 transactions. * **Error** – There was an error when the payment was being processed. The reason is given in the `refusalReason` field. This is a final state. * **IdentifyShopper** – The issuer requires the shopper's device fingerprint before the payment can be authenticated. Returned for 3D Secure 2 transactions. * **PartiallyAuthorised** – The payment has been authorised for a partial amount. This happens for card payments when the merchant supports Partial Authorisations and the cardholder has insufficient funds. * **Pending** – Indicates that it is not possible to obtain the final status of the payment. This can happen if the systems providing final status information for the payment are unavailable, or if the shopper needs to take further action to complete the payment. * **PresentToShopper** – Indicates that the response contains additional information that you need to present to a shopper, so that they can use it to complete a payment. * **Received** – Indicates the payment has successfully been received by Adyen, and will be processed. This is the initial state for all payments. * **RedirectShopper** – Indicates the shopper should be redirected to an external web page or app to complete the authorisation. * **Refused** – Indicates the payment was refused. The reason is given in the `refusalReason` field. This is a final state.
	ResultCode *string `json:"resultCode,omitempty"`
}

// NewPaymentResult instantiates a new PaymentResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentResult() *PaymentResult {
	this := PaymentResult{}
	return &this
}

// NewPaymentResultWithDefaults instantiates a new PaymentResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentResultWithDefaults() *PaymentResult {
	this := PaymentResult{}
	return &this
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *PaymentResult) GetAdditionalData() map[string]string {
	if o == nil || common.IsNil(o.AdditionalData) {
		var ret map[string]string
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResult) GetAdditionalDataOk() (*map[string]string, bool) {
	if o == nil || common.IsNil(o.AdditionalData) {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *PaymentResult) HasAdditionalData() bool {
	if o != nil && !common.IsNil(o.AdditionalData) {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given map[string]string and assigns it to the AdditionalData field.
func (o *PaymentResult) SetAdditionalData(v map[string]string) {
	o.AdditionalData = &v
}

// GetAuthCode returns the AuthCode field value if set, zero value otherwise.
func (o *PaymentResult) GetAuthCode() string {
	if o == nil || common.IsNil(o.AuthCode) {
		var ret string
		return ret
	}
	return *o.AuthCode
}

// GetAuthCodeOk returns a tuple with the AuthCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResult) GetAuthCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.AuthCode) {
		return nil, false
	}
	return o.AuthCode, true
}

// HasAuthCode returns a boolean if a field has been set.
func (o *PaymentResult) HasAuthCode() bool {
	if o != nil && !common.IsNil(o.AuthCode) {
		return true
	}

	return false
}

// SetAuthCode gets a reference to the given string and assigns it to the AuthCode field.
func (o *PaymentResult) SetAuthCode(v string) {
	o.AuthCode = &v
}

// GetDccAmount returns the DccAmount field value if set, zero value otherwise.
func (o *PaymentResult) GetDccAmount() Amount {
	if o == nil || common.IsNil(o.DccAmount) {
		var ret Amount
		return ret
	}
	return *o.DccAmount
}

// GetDccAmountOk returns a tuple with the DccAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResult) GetDccAmountOk() (*Amount, bool) {
	if o == nil || common.IsNil(o.DccAmount) {
		return nil, false
	}
	return o.DccAmount, true
}

// HasDccAmount returns a boolean if a field has been set.
func (o *PaymentResult) HasDccAmount() bool {
	if o != nil && !common.IsNil(o.DccAmount) {
		return true
	}

	return false
}

// SetDccAmount gets a reference to the given Amount and assigns it to the DccAmount field.
func (o *PaymentResult) SetDccAmount(v Amount) {
	o.DccAmount = &v
}

// GetDccSignature returns the DccSignature field value if set, zero value otherwise.
func (o *PaymentResult) GetDccSignature() string {
	if o == nil || common.IsNil(o.DccSignature) {
		var ret string
		return ret
	}
	return *o.DccSignature
}

// GetDccSignatureOk returns a tuple with the DccSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResult) GetDccSignatureOk() (*string, bool) {
	if o == nil || common.IsNil(o.DccSignature) {
		return nil, false
	}
	return o.DccSignature, true
}

// HasDccSignature returns a boolean if a field has been set.
func (o *PaymentResult) HasDccSignature() bool {
	if o != nil && !common.IsNil(o.DccSignature) {
		return true
	}

	return false
}

// SetDccSignature gets a reference to the given string and assigns it to the DccSignature field.
func (o *PaymentResult) SetDccSignature(v string) {
	o.DccSignature = &v
}

// GetFraudResult returns the FraudResult field value if set, zero value otherwise.
func (o *PaymentResult) GetFraudResult() FraudResult {
	if o == nil || common.IsNil(o.FraudResult) {
		var ret FraudResult
		return ret
	}
	return *o.FraudResult
}

// GetFraudResultOk returns a tuple with the FraudResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResult) GetFraudResultOk() (*FraudResult, bool) {
	if o == nil || common.IsNil(o.FraudResult) {
		return nil, false
	}
	return o.FraudResult, true
}

// HasFraudResult returns a boolean if a field has been set.
func (o *PaymentResult) HasFraudResult() bool {
	if o != nil && !common.IsNil(o.FraudResult) {
		return true
	}

	return false
}

// SetFraudResult gets a reference to the given FraudResult and assigns it to the FraudResult field.
func (o *PaymentResult) SetFraudResult(v FraudResult) {
	o.FraudResult = &v
}

// GetIssuerUrl returns the IssuerUrl field value if set, zero value otherwise.
func (o *PaymentResult) GetIssuerUrl() string {
	if o == nil || common.IsNil(o.IssuerUrl) {
		var ret string
		return ret
	}
	return *o.IssuerUrl
}

// GetIssuerUrlOk returns a tuple with the IssuerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResult) GetIssuerUrlOk() (*string, bool) {
	if o == nil || common.IsNil(o.IssuerUrl) {
		return nil, false
	}
	return o.IssuerUrl, true
}

// HasIssuerUrl returns a boolean if a field has been set.
func (o *PaymentResult) HasIssuerUrl() bool {
	if o != nil && !common.IsNil(o.IssuerUrl) {
		return true
	}

	return false
}

// SetIssuerUrl gets a reference to the given string and assigns it to the IssuerUrl field.
func (o *PaymentResult) SetIssuerUrl(v string) {
	o.IssuerUrl = &v
}

// GetMd returns the Md field value if set, zero value otherwise.
func (o *PaymentResult) GetMd() string {
	if o == nil || common.IsNil(o.Md) {
		var ret string
		return ret
	}
	return *o.Md
}

// GetMdOk returns a tuple with the Md field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResult) GetMdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Md) {
		return nil, false
	}
	return o.Md, true
}

// HasMd returns a boolean if a field has been set.
func (o *PaymentResult) HasMd() bool {
	if o != nil && !common.IsNil(o.Md) {
		return true
	}

	return false
}

// SetMd gets a reference to the given string and assigns it to the Md field.
func (o *PaymentResult) SetMd(v string) {
	o.Md = &v
}

// GetPaRequest returns the PaRequest field value if set, zero value otherwise.
func (o *PaymentResult) GetPaRequest() string {
	if o == nil || common.IsNil(o.PaRequest) {
		var ret string
		return ret
	}
	return *o.PaRequest
}

// GetPaRequestOk returns a tuple with the PaRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResult) GetPaRequestOk() (*string, bool) {
	if o == nil || common.IsNil(o.PaRequest) {
		return nil, false
	}
	return o.PaRequest, true
}

// HasPaRequest returns a boolean if a field has been set.
func (o *PaymentResult) HasPaRequest() bool {
	if o != nil && !common.IsNil(o.PaRequest) {
		return true
	}

	return false
}

// SetPaRequest gets a reference to the given string and assigns it to the PaRequest field.
func (o *PaymentResult) SetPaRequest(v string) {
	o.PaRequest = &v
}

// GetPspReference returns the PspReference field value if set, zero value otherwise.
func (o *PaymentResult) GetPspReference() string {
	if o == nil || common.IsNil(o.PspReference) {
		var ret string
		return ret
	}
	return *o.PspReference
}

// GetPspReferenceOk returns a tuple with the PspReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResult) GetPspReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.PspReference) {
		return nil, false
	}
	return o.PspReference, true
}

// HasPspReference returns a boolean if a field has been set.
func (o *PaymentResult) HasPspReference() bool {
	if o != nil && !common.IsNil(o.PspReference) {
		return true
	}

	return false
}

// SetPspReference gets a reference to the given string and assigns it to the PspReference field.
func (o *PaymentResult) SetPspReference(v string) {
	o.PspReference = &v
}

// GetRefusalReason returns the RefusalReason field value if set, zero value otherwise.
func (o *PaymentResult) GetRefusalReason() string {
	if o == nil || common.IsNil(o.RefusalReason) {
		var ret string
		return ret
	}
	return *o.RefusalReason
}

// GetRefusalReasonOk returns a tuple with the RefusalReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResult) GetRefusalReasonOk() (*string, bool) {
	if o == nil || common.IsNil(o.RefusalReason) {
		return nil, false
	}
	return o.RefusalReason, true
}

// HasRefusalReason returns a boolean if a field has been set.
func (o *PaymentResult) HasRefusalReason() bool {
	if o != nil && !common.IsNil(o.RefusalReason) {
		return true
	}

	return false
}

// SetRefusalReason gets a reference to the given string and assigns it to the RefusalReason field.
func (o *PaymentResult) SetRefusalReason(v string) {
	o.RefusalReason = &v
}

// GetResultCode returns the ResultCode field value if set, zero value otherwise.
func (o *PaymentResult) GetResultCode() string {
	if o == nil || common.IsNil(o.ResultCode) {
		var ret string
		return ret
	}
	return *o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResult) GetResultCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ResultCode) {
		return nil, false
	}
	return o.ResultCode, true
}

// HasResultCode returns a boolean if a field has been set.
func (o *PaymentResult) HasResultCode() bool {
	if o != nil && !common.IsNil(o.ResultCode) {
		return true
	}

	return false
}

// SetResultCode gets a reference to the given string and assigns it to the ResultCode field.
func (o *PaymentResult) SetResultCode(v string) {
	o.ResultCode = &v
}

func (o PaymentResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AdditionalData) {
		toSerialize["additionalData"] = o.AdditionalData
	}
	if !common.IsNil(o.AuthCode) {
		toSerialize["authCode"] = o.AuthCode
	}
	if !common.IsNil(o.DccAmount) {
		toSerialize["dccAmount"] = o.DccAmount
	}
	if !common.IsNil(o.DccSignature) {
		toSerialize["dccSignature"] = o.DccSignature
	}
	if !common.IsNil(o.FraudResult) {
		toSerialize["fraudResult"] = o.FraudResult
	}
	if !common.IsNil(o.IssuerUrl) {
		toSerialize["issuerUrl"] = o.IssuerUrl
	}
	if !common.IsNil(o.Md) {
		toSerialize["md"] = o.Md
	}
	if !common.IsNil(o.PaRequest) {
		toSerialize["paRequest"] = o.PaRequest
	}
	if !common.IsNil(o.PspReference) {
		toSerialize["pspReference"] = o.PspReference
	}
	if !common.IsNil(o.RefusalReason) {
		toSerialize["refusalReason"] = o.RefusalReason
	}
	if !common.IsNil(o.ResultCode) {
		toSerialize["resultCode"] = o.ResultCode
	}
	return toSerialize, nil
}

type NullablePaymentResult struct {
	value *PaymentResult
	isSet bool
}

func (v NullablePaymentResult) Get() *PaymentResult {
	return v.value
}

func (v *NullablePaymentResult) Set(val *PaymentResult) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentResult) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentResult(val *PaymentResult) *NullablePaymentResult {
	return &NullablePaymentResult{value: val, isSet: true}
}

func (v NullablePaymentResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *PaymentResult) isValidResultCode() bool {
	var allowedEnumValues = []string{"AuthenticationFinished", "AuthenticationNotRequired", "Authorised", "Cancelled", "ChallengeShopper", "Error", "IdentifyShopper", "PartiallyAuthorised", "Pending", "PresentToShopper", "Received", "RedirectShopper", "Refused", "Success"}
	for _, allowed := range allowedEnumValues {
		if o.GetResultCode() == allowed {
			return true
		}
	}
	return false
}
