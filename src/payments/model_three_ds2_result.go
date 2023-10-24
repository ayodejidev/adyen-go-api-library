/*
Adyen Payment API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package payments

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// checks if the ThreeDS2Result type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ThreeDS2Result{}

// ThreeDS2Result struct for ThreeDS2Result
type ThreeDS2Result struct {
	// The `authenticationValue` value as defined in the 3D Secure 2 specification.
	AuthenticationValue *string `json:"authenticationValue,omitempty"`
	// The algorithm used by the ACS to calculate the authentication value, only for Cartes Bancaires integrations.
	CavvAlgorithm *string `json:"cavvAlgorithm,omitempty"`
	// Indicator informing the Access Control Server (ACS) and the Directory Server (DS) that the authentication has been cancelled. For possible values, refer to [3D Secure API reference](https://docs.adyen.com/online-payments/3d-secure/api-reference#mpidata).
	ChallengeCancel *string `json:"challengeCancel,omitempty"`
	// Specifies a preference for receiving a challenge from the issuer. Allowed values: * `noPreference` * `requestNoChallenge` * `requestChallenge` * `requestChallengeAsMandate`
	ChallengeIndicator *string `json:"challengeIndicator,omitempty"`
	// The `dsTransID` value as defined in the 3D Secure 2 specification.
	DsTransID *string `json:"dsTransID,omitempty"`
	// The `eci` value as defined in the 3D Secure 2 specification.
	Eci *string `json:"eci,omitempty"`
	// Indicates the exemption type that was applied by the issuer to the authentication, if exemption applied. Allowed values: * `lowValue` * `secureCorporate` * `trustedBeneficiary` * `transactionRiskAnalysis`
	ExemptionIndicator *string `json:"exemptionIndicator,omitempty"`
	// The `messageVersion` value as defined in the 3D Secure 2 specification.
	MessageVersion *string `json:"messageVersion,omitempty"`
	// Risk score calculated by Cartes Bancaires Directory Server (DS).
	RiskScore *string `json:"riskScore,omitempty"`
	// The `threeDSServerTransID` value as defined in the 3D Secure 2 specification.
	ThreeDSServerTransID *string `json:"threeDSServerTransID,omitempty"`
	// The `timestamp` value of the 3D Secure 2 authentication.
	Timestamp *string `json:"timestamp,omitempty"`
	// The `transStatus` value as defined in the 3D Secure 2 specification.
	TransStatus *string `json:"transStatus,omitempty"`
	// Provides information on why the `transStatus` field has the specified value. For possible values, refer to [our docs](https://docs.adyen.com/online-payments/3d-secure/api-reference#possible-transstatusreason-values).
	TransStatusReason *string `json:"transStatusReason,omitempty"`
	// The `whiteListStatus` value as defined in the 3D Secure 2 specification.
	WhiteListStatus *string `json:"whiteListStatus,omitempty"`
}

// NewThreeDS2Result instantiates a new ThreeDS2Result object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThreeDS2Result() *ThreeDS2Result {
	this := ThreeDS2Result{}
	return &this
}

// NewThreeDS2ResultWithDefaults instantiates a new ThreeDS2Result object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThreeDS2ResultWithDefaults() *ThreeDS2Result {
	this := ThreeDS2Result{}
	return &this
}

// GetAuthenticationValue returns the AuthenticationValue field value if set, zero value otherwise.
func (o *ThreeDS2Result) GetAuthenticationValue() string {
	if o == nil || common.IsNil(o.AuthenticationValue) {
		var ret string
		return ret
	}
	return *o.AuthenticationValue
}

// GetAuthenticationValueOk returns a tuple with the AuthenticationValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDS2Result) GetAuthenticationValueOk() (*string, bool) {
	if o == nil || common.IsNil(o.AuthenticationValue) {
		return nil, false
	}
	return o.AuthenticationValue, true
}

// HasAuthenticationValue returns a boolean if a field has been set.
func (o *ThreeDS2Result) HasAuthenticationValue() bool {
	if o != nil && !common.IsNil(o.AuthenticationValue) {
		return true
	}

	return false
}

// SetAuthenticationValue gets a reference to the given string and assigns it to the AuthenticationValue field.
func (o *ThreeDS2Result) SetAuthenticationValue(v string) {
	o.AuthenticationValue = &v
}

// GetCavvAlgorithm returns the CavvAlgorithm field value if set, zero value otherwise.
func (o *ThreeDS2Result) GetCavvAlgorithm() string {
	if o == nil || common.IsNil(o.CavvAlgorithm) {
		var ret string
		return ret
	}
	return *o.CavvAlgorithm
}

// GetCavvAlgorithmOk returns a tuple with the CavvAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDS2Result) GetCavvAlgorithmOk() (*string, bool) {
	if o == nil || common.IsNil(o.CavvAlgorithm) {
		return nil, false
	}
	return o.CavvAlgorithm, true
}

// HasCavvAlgorithm returns a boolean if a field has been set.
func (o *ThreeDS2Result) HasCavvAlgorithm() bool {
	if o != nil && !common.IsNil(o.CavvAlgorithm) {
		return true
	}

	return false
}

// SetCavvAlgorithm gets a reference to the given string and assigns it to the CavvAlgorithm field.
func (o *ThreeDS2Result) SetCavvAlgorithm(v string) {
	o.CavvAlgorithm = &v
}

// GetChallengeCancel returns the ChallengeCancel field value if set, zero value otherwise.
func (o *ThreeDS2Result) GetChallengeCancel() string {
	if o == nil || common.IsNil(o.ChallengeCancel) {
		var ret string
		return ret
	}
	return *o.ChallengeCancel
}

// GetChallengeCancelOk returns a tuple with the ChallengeCancel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDS2Result) GetChallengeCancelOk() (*string, bool) {
	if o == nil || common.IsNil(o.ChallengeCancel) {
		return nil, false
	}
	return o.ChallengeCancel, true
}

// HasChallengeCancel returns a boolean if a field has been set.
func (o *ThreeDS2Result) HasChallengeCancel() bool {
	if o != nil && !common.IsNil(o.ChallengeCancel) {
		return true
	}

	return false
}

// SetChallengeCancel gets a reference to the given string and assigns it to the ChallengeCancel field.
func (o *ThreeDS2Result) SetChallengeCancel(v string) {
	o.ChallengeCancel = &v
}

// GetChallengeIndicator returns the ChallengeIndicator field value if set, zero value otherwise.
func (o *ThreeDS2Result) GetChallengeIndicator() string {
	if o == nil || common.IsNil(o.ChallengeIndicator) {
		var ret string
		return ret
	}
	return *o.ChallengeIndicator
}

// GetChallengeIndicatorOk returns a tuple with the ChallengeIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDS2Result) GetChallengeIndicatorOk() (*string, bool) {
	if o == nil || common.IsNil(o.ChallengeIndicator) {
		return nil, false
	}
	return o.ChallengeIndicator, true
}

// HasChallengeIndicator returns a boolean if a field has been set.
func (o *ThreeDS2Result) HasChallengeIndicator() bool {
	if o != nil && !common.IsNil(o.ChallengeIndicator) {
		return true
	}

	return false
}

// SetChallengeIndicator gets a reference to the given string and assigns it to the ChallengeIndicator field.
func (o *ThreeDS2Result) SetChallengeIndicator(v string) {
	o.ChallengeIndicator = &v
}

// GetDsTransID returns the DsTransID field value if set, zero value otherwise.
func (o *ThreeDS2Result) GetDsTransID() string {
	if o == nil || common.IsNil(o.DsTransID) {
		var ret string
		return ret
	}
	return *o.DsTransID
}

// GetDsTransIDOk returns a tuple with the DsTransID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDS2Result) GetDsTransIDOk() (*string, bool) {
	if o == nil || common.IsNil(o.DsTransID) {
		return nil, false
	}
	return o.DsTransID, true
}

// HasDsTransID returns a boolean if a field has been set.
func (o *ThreeDS2Result) HasDsTransID() bool {
	if o != nil && !common.IsNil(o.DsTransID) {
		return true
	}

	return false
}

// SetDsTransID gets a reference to the given string and assigns it to the DsTransID field.
func (o *ThreeDS2Result) SetDsTransID(v string) {
	o.DsTransID = &v
}

// GetEci returns the Eci field value if set, zero value otherwise.
func (o *ThreeDS2Result) GetEci() string {
	if o == nil || common.IsNil(o.Eci) {
		var ret string
		return ret
	}
	return *o.Eci
}

// GetEciOk returns a tuple with the Eci field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDS2Result) GetEciOk() (*string, bool) {
	if o == nil || common.IsNil(o.Eci) {
		return nil, false
	}
	return o.Eci, true
}

// HasEci returns a boolean if a field has been set.
func (o *ThreeDS2Result) HasEci() bool {
	if o != nil && !common.IsNil(o.Eci) {
		return true
	}

	return false
}

// SetEci gets a reference to the given string and assigns it to the Eci field.
func (o *ThreeDS2Result) SetEci(v string) {
	o.Eci = &v
}

// GetExemptionIndicator returns the ExemptionIndicator field value if set, zero value otherwise.
func (o *ThreeDS2Result) GetExemptionIndicator() string {
	if o == nil || common.IsNil(o.ExemptionIndicator) {
		var ret string
		return ret
	}
	return *o.ExemptionIndicator
}

// GetExemptionIndicatorOk returns a tuple with the ExemptionIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDS2Result) GetExemptionIndicatorOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExemptionIndicator) {
		return nil, false
	}
	return o.ExemptionIndicator, true
}

// HasExemptionIndicator returns a boolean if a field has been set.
func (o *ThreeDS2Result) HasExemptionIndicator() bool {
	if o != nil && !common.IsNil(o.ExemptionIndicator) {
		return true
	}

	return false
}

// SetExemptionIndicator gets a reference to the given string and assigns it to the ExemptionIndicator field.
func (o *ThreeDS2Result) SetExemptionIndicator(v string) {
	o.ExemptionIndicator = &v
}

// GetMessageVersion returns the MessageVersion field value if set, zero value otherwise.
func (o *ThreeDS2Result) GetMessageVersion() string {
	if o == nil || common.IsNil(o.MessageVersion) {
		var ret string
		return ret
	}
	return *o.MessageVersion
}

// GetMessageVersionOk returns a tuple with the MessageVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDS2Result) GetMessageVersionOk() (*string, bool) {
	if o == nil || common.IsNil(o.MessageVersion) {
		return nil, false
	}
	return o.MessageVersion, true
}

// HasMessageVersion returns a boolean if a field has been set.
func (o *ThreeDS2Result) HasMessageVersion() bool {
	if o != nil && !common.IsNil(o.MessageVersion) {
		return true
	}

	return false
}

// SetMessageVersion gets a reference to the given string and assigns it to the MessageVersion field.
func (o *ThreeDS2Result) SetMessageVersion(v string) {
	o.MessageVersion = &v
}

// GetRiskScore returns the RiskScore field value if set, zero value otherwise.
func (o *ThreeDS2Result) GetRiskScore() string {
	if o == nil || common.IsNil(o.RiskScore) {
		var ret string
		return ret
	}
	return *o.RiskScore
}

// GetRiskScoreOk returns a tuple with the RiskScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDS2Result) GetRiskScoreOk() (*string, bool) {
	if o == nil || common.IsNil(o.RiskScore) {
		return nil, false
	}
	return o.RiskScore, true
}

// HasRiskScore returns a boolean if a field has been set.
func (o *ThreeDS2Result) HasRiskScore() bool {
	if o != nil && !common.IsNil(o.RiskScore) {
		return true
	}

	return false
}

// SetRiskScore gets a reference to the given string and assigns it to the RiskScore field.
func (o *ThreeDS2Result) SetRiskScore(v string) {
	o.RiskScore = &v
}

// GetThreeDSServerTransID returns the ThreeDSServerTransID field value if set, zero value otherwise.
func (o *ThreeDS2Result) GetThreeDSServerTransID() string {
	if o == nil || common.IsNil(o.ThreeDSServerTransID) {
		var ret string
		return ret
	}
	return *o.ThreeDSServerTransID
}

// GetThreeDSServerTransIDOk returns a tuple with the ThreeDSServerTransID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDS2Result) GetThreeDSServerTransIDOk() (*string, bool) {
	if o == nil || common.IsNil(o.ThreeDSServerTransID) {
		return nil, false
	}
	return o.ThreeDSServerTransID, true
}

// HasThreeDSServerTransID returns a boolean if a field has been set.
func (o *ThreeDS2Result) HasThreeDSServerTransID() bool {
	if o != nil && !common.IsNil(o.ThreeDSServerTransID) {
		return true
	}

	return false
}

// SetThreeDSServerTransID gets a reference to the given string and assigns it to the ThreeDSServerTransID field.
func (o *ThreeDS2Result) SetThreeDSServerTransID(v string) {
	o.ThreeDSServerTransID = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ThreeDS2Result) GetTimestamp() string {
	if o == nil || common.IsNil(o.Timestamp) {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDS2Result) GetTimestampOk() (*string, bool) {
	if o == nil || common.IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ThreeDS2Result) HasTimestamp() bool {
	if o != nil && !common.IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *ThreeDS2Result) SetTimestamp(v string) {
	o.Timestamp = &v
}

// GetTransStatus returns the TransStatus field value if set, zero value otherwise.
func (o *ThreeDS2Result) GetTransStatus() string {
	if o == nil || common.IsNil(o.TransStatus) {
		var ret string
		return ret
	}
	return *o.TransStatus
}

// GetTransStatusOk returns a tuple with the TransStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDS2Result) GetTransStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.TransStatus) {
		return nil, false
	}
	return o.TransStatus, true
}

// HasTransStatus returns a boolean if a field has been set.
func (o *ThreeDS2Result) HasTransStatus() bool {
	if o != nil && !common.IsNil(o.TransStatus) {
		return true
	}

	return false
}

// SetTransStatus gets a reference to the given string and assigns it to the TransStatus field.
func (o *ThreeDS2Result) SetTransStatus(v string) {
	o.TransStatus = &v
}

// GetTransStatusReason returns the TransStatusReason field value if set, zero value otherwise.
func (o *ThreeDS2Result) GetTransStatusReason() string {
	if o == nil || common.IsNil(o.TransStatusReason) {
		var ret string
		return ret
	}
	return *o.TransStatusReason
}

// GetTransStatusReasonOk returns a tuple with the TransStatusReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDS2Result) GetTransStatusReasonOk() (*string, bool) {
	if o == nil || common.IsNil(o.TransStatusReason) {
		return nil, false
	}
	return o.TransStatusReason, true
}

// HasTransStatusReason returns a boolean if a field has been set.
func (o *ThreeDS2Result) HasTransStatusReason() bool {
	if o != nil && !common.IsNil(o.TransStatusReason) {
		return true
	}

	return false
}

// SetTransStatusReason gets a reference to the given string and assigns it to the TransStatusReason field.
func (o *ThreeDS2Result) SetTransStatusReason(v string) {
	o.TransStatusReason = &v
}

// GetWhiteListStatus returns the WhiteListStatus field value if set, zero value otherwise.
func (o *ThreeDS2Result) GetWhiteListStatus() string {
	if o == nil || common.IsNil(o.WhiteListStatus) {
		var ret string
		return ret
	}
	return *o.WhiteListStatus
}

// GetWhiteListStatusOk returns a tuple with the WhiteListStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDS2Result) GetWhiteListStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.WhiteListStatus) {
		return nil, false
	}
	return o.WhiteListStatus, true
}

// HasWhiteListStatus returns a boolean if a field has been set.
func (o *ThreeDS2Result) HasWhiteListStatus() bool {
	if o != nil && !common.IsNil(o.WhiteListStatus) {
		return true
	}

	return false
}

// SetWhiteListStatus gets a reference to the given string and assigns it to the WhiteListStatus field.
func (o *ThreeDS2Result) SetWhiteListStatus(v string) {
	o.WhiteListStatus = &v
}

func (o ThreeDS2Result) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThreeDS2Result) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AuthenticationValue) {
		toSerialize["authenticationValue"] = o.AuthenticationValue
	}
	if !common.IsNil(o.CavvAlgorithm) {
		toSerialize["cavvAlgorithm"] = o.CavvAlgorithm
	}
	if !common.IsNil(o.ChallengeCancel) {
		toSerialize["challengeCancel"] = o.ChallengeCancel
	}
	if !common.IsNil(o.ChallengeIndicator) {
		toSerialize["challengeIndicator"] = o.ChallengeIndicator
	}
	if !common.IsNil(o.DsTransID) {
		toSerialize["dsTransID"] = o.DsTransID
	}
	if !common.IsNil(o.Eci) {
		toSerialize["eci"] = o.Eci
	}
	if !common.IsNil(o.ExemptionIndicator) {
		toSerialize["exemptionIndicator"] = o.ExemptionIndicator
	}
	if !common.IsNil(o.MessageVersion) {
		toSerialize["messageVersion"] = o.MessageVersion
	}
	if !common.IsNil(o.RiskScore) {
		toSerialize["riskScore"] = o.RiskScore
	}
	if !common.IsNil(o.ThreeDSServerTransID) {
		toSerialize["threeDSServerTransID"] = o.ThreeDSServerTransID
	}
	if !common.IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !common.IsNil(o.TransStatus) {
		toSerialize["transStatus"] = o.TransStatus
	}
	if !common.IsNil(o.TransStatusReason) {
		toSerialize["transStatusReason"] = o.TransStatusReason
	}
	if !common.IsNil(o.WhiteListStatus) {
		toSerialize["whiteListStatus"] = o.WhiteListStatus
	}
	return toSerialize, nil
}

type NullableThreeDS2Result struct {
	value *ThreeDS2Result
	isSet bool
}

func (v NullableThreeDS2Result) Get() *ThreeDS2Result {
	return v.value
}

func (v *NullableThreeDS2Result) Set(val *ThreeDS2Result) {
	v.value = val
	v.isSet = true
}

func (v NullableThreeDS2Result) IsSet() bool {
	return v.isSet
}

func (v *NullableThreeDS2Result) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThreeDS2Result(val *ThreeDS2Result) *NullableThreeDS2Result {
	return &NullableThreeDS2Result{value: val, isSet: true}
}

func (v NullableThreeDS2Result) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThreeDS2Result) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *ThreeDS2Result) isValidChallengeCancel() bool {
	var allowedEnumValues = []string{"01", "02", "03", "04", "05", "06", "07"}
	for _, allowed := range allowedEnumValues {
		if o.GetChallengeCancel() == allowed {
			return true
		}
	}
	return false
}
func (o *ThreeDS2Result) isValidChallengeIndicator() bool {
	var allowedEnumValues = []string{"noPreference", "requestNoChallenge", "requestChallenge", "requestChallengeAsMandate"}
	for _, allowed := range allowedEnumValues {
		if o.GetChallengeIndicator() == allowed {
			return true
		}
	}
	return false
}
func (o *ThreeDS2Result) isValidExemptionIndicator() bool {
	var allowedEnumValues = []string{"lowValue", "secureCorporate", "trustedBeneficiary", "transactionRiskAnalysis"}
	for _, allowed := range allowedEnumValues {
		if o.GetExemptionIndicator() == allowed {
			return true
		}
	}
	return false
}
