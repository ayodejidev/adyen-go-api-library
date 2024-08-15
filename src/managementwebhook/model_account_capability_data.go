/*
Management Webhooks

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package managementwebhook

import (
	"encoding/json"
	"time"

	"github.com/adyen/adyen-go-api-library/v12/src/common"
)

// checks if the AccountCapabilityData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountCapabilityData{}

// AccountCapabilityData struct for AccountCapabilityData
type AccountCapabilityData struct {
	// Indicates whether the capability is allowed. Adyen sets this to **true** if the verification is successful.
	Allowed *bool `json:"allowed,omitempty"`
	// The allowed level of the capability. Some capabilities have different levels which correspond to thresholds. Higher levels may require additional checks and increased monitoring.Possible values: **notApplicable**, **low**, **medium**, **high**.
	AllowedLevel *string `json:"allowedLevel,omitempty"`
	// The name of the capability. For example, **sendToTransferInstrument**.
	Capability *string `json:"capability,omitempty"`
	// List of entities that has problems with verification. The information includes the details of the errors and the actions that you can take to resolve them.
	Problems []CapabilityProblem `json:"problems,omitempty"`
	// Indicates whether you requested the capability.
	Requested bool `json:"requested"`
	// The level that you requested for the capability. Some capabilities have different levels which correspond to thresholds. Higher levels may require additional checks and increased monitoring.Possible values: **notApplicable**, **low**, **medium**, **high**.
	RequestedLevel string `json:"requestedLevel"`
	// The verification deadline for the capability that will be disallowed if verification errors are not resolved.
	VerificationDeadline *time.Time `json:"verificationDeadline,omitempty"`
	// The status of the verification checks for the capability.  Possible values:  * **pending**: Adyen is running the verification.  * **invalid**: The verification failed. Check if the `errors` array contains more information.  * **valid**: The verification was successful.  * **rejected**: Adyen checked the information and found reasons to not allow the capability.
	VerificationStatus *string `json:"verificationStatus,omitempty"`
}

// NewAccountCapabilityData instantiates a new AccountCapabilityData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountCapabilityData(requested bool, requestedLevel string) *AccountCapabilityData {
	this := AccountCapabilityData{}
	this.Requested = requested
	this.RequestedLevel = requestedLevel
	return &this
}

// NewAccountCapabilityDataWithDefaults instantiates a new AccountCapabilityData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountCapabilityDataWithDefaults() *AccountCapabilityData {
	this := AccountCapabilityData{}
	return &this
}

// GetAllowed returns the Allowed field value if set, zero value otherwise.
func (o *AccountCapabilityData) GetAllowed() bool {
	if o == nil || common.IsNil(o.Allowed) {
		var ret bool
		return ret
	}
	return *o.Allowed
}

// GetAllowedOk returns a tuple with the Allowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCapabilityData) GetAllowedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Allowed) {
		return nil, false
	}
	return o.Allowed, true
}

// HasAllowed returns a boolean if a field has been set.
func (o *AccountCapabilityData) HasAllowed() bool {
	if o != nil && !common.IsNil(o.Allowed) {
		return true
	}

	return false
}

// SetAllowed gets a reference to the given bool and assigns it to the Allowed field.
func (o *AccountCapabilityData) SetAllowed(v bool) {
	o.Allowed = &v
}

// GetAllowedLevel returns the AllowedLevel field value if set, zero value otherwise.
func (o *AccountCapabilityData) GetAllowedLevel() string {
	if o == nil || common.IsNil(o.AllowedLevel) {
		var ret string
		return ret
	}
	return *o.AllowedLevel
}

// GetAllowedLevelOk returns a tuple with the AllowedLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCapabilityData) GetAllowedLevelOk() (*string, bool) {
	if o == nil || common.IsNil(o.AllowedLevel) {
		return nil, false
	}
	return o.AllowedLevel, true
}

// HasAllowedLevel returns a boolean if a field has been set.
func (o *AccountCapabilityData) HasAllowedLevel() bool {
	if o != nil && !common.IsNil(o.AllowedLevel) {
		return true
	}

	return false
}

// SetAllowedLevel gets a reference to the given string and assigns it to the AllowedLevel field.
func (o *AccountCapabilityData) SetAllowedLevel(v string) {
	o.AllowedLevel = &v
}

// GetCapability returns the Capability field value if set, zero value otherwise.
func (o *AccountCapabilityData) GetCapability() string {
	if o == nil || common.IsNil(o.Capability) {
		var ret string
		return ret
	}
	return *o.Capability
}

// GetCapabilityOk returns a tuple with the Capability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCapabilityData) GetCapabilityOk() (*string, bool) {
	if o == nil || common.IsNil(o.Capability) {
		return nil, false
	}
	return o.Capability, true
}

// HasCapability returns a boolean if a field has been set.
func (o *AccountCapabilityData) HasCapability() bool {
	if o != nil && !common.IsNil(o.Capability) {
		return true
	}

	return false
}

// SetCapability gets a reference to the given string and assigns it to the Capability field.
func (o *AccountCapabilityData) SetCapability(v string) {
	o.Capability = &v
}

// GetProblems returns the Problems field value if set, zero value otherwise.
func (o *AccountCapabilityData) GetProblems() []CapabilityProblem {
	if o == nil || common.IsNil(o.Problems) {
		var ret []CapabilityProblem
		return ret
	}
	return o.Problems
}

// GetProblemsOk returns a tuple with the Problems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCapabilityData) GetProblemsOk() ([]CapabilityProblem, bool) {
	if o == nil || common.IsNil(o.Problems) {
		return nil, false
	}
	return o.Problems, true
}

// HasProblems returns a boolean if a field has been set.
func (o *AccountCapabilityData) HasProblems() bool {
	if o != nil && !common.IsNil(o.Problems) {
		return true
	}

	return false
}

// SetProblems gets a reference to the given []CapabilityProblem and assigns it to the Problems field.
func (o *AccountCapabilityData) SetProblems(v []CapabilityProblem) {
	o.Problems = v
}

// GetRequested returns the Requested field value
func (o *AccountCapabilityData) GetRequested() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Requested
}

// GetRequestedOk returns a tuple with the Requested field value
// and a boolean to check if the value has been set.
func (o *AccountCapabilityData) GetRequestedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requested, true
}

// SetRequested sets field value
func (o *AccountCapabilityData) SetRequested(v bool) {
	o.Requested = v
}

// GetRequestedLevel returns the RequestedLevel field value
func (o *AccountCapabilityData) GetRequestedLevel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestedLevel
}

// GetRequestedLevelOk returns a tuple with the RequestedLevel field value
// and a boolean to check if the value has been set.
func (o *AccountCapabilityData) GetRequestedLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestedLevel, true
}

// SetRequestedLevel sets field value
func (o *AccountCapabilityData) SetRequestedLevel(v string) {
	o.RequestedLevel = v
}

// GetVerificationDeadline returns the VerificationDeadline field value if set, zero value otherwise.
func (o *AccountCapabilityData) GetVerificationDeadline() time.Time {
	if o == nil || common.IsNil(o.VerificationDeadline) {
		var ret time.Time
		return ret
	}
	return *o.VerificationDeadline
}

// GetVerificationDeadlineOk returns a tuple with the VerificationDeadline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCapabilityData) GetVerificationDeadlineOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.VerificationDeadline) {
		return nil, false
	}
	return o.VerificationDeadline, true
}

// HasVerificationDeadline returns a boolean if a field has been set.
func (o *AccountCapabilityData) HasVerificationDeadline() bool {
	if o != nil && !common.IsNil(o.VerificationDeadline) {
		return true
	}

	return false
}

// SetVerificationDeadline gets a reference to the given time.Time and assigns it to the VerificationDeadline field.
func (o *AccountCapabilityData) SetVerificationDeadline(v time.Time) {
	o.VerificationDeadline = &v
}

// GetVerificationStatus returns the VerificationStatus field value if set, zero value otherwise.
func (o *AccountCapabilityData) GetVerificationStatus() string {
	if o == nil || common.IsNil(o.VerificationStatus) {
		var ret string
		return ret
	}
	return *o.VerificationStatus
}

// GetVerificationStatusOk returns a tuple with the VerificationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCapabilityData) GetVerificationStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.VerificationStatus) {
		return nil, false
	}
	return o.VerificationStatus, true
}

// HasVerificationStatus returns a boolean if a field has been set.
func (o *AccountCapabilityData) HasVerificationStatus() bool {
	if o != nil && !common.IsNil(o.VerificationStatus) {
		return true
	}

	return false
}

// SetVerificationStatus gets a reference to the given string and assigns it to the VerificationStatus field.
func (o *AccountCapabilityData) SetVerificationStatus(v string) {
	o.VerificationStatus = &v
}

func (o AccountCapabilityData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountCapabilityData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Allowed) {
		toSerialize["allowed"] = o.Allowed
	}
	if !common.IsNil(o.AllowedLevel) {
		toSerialize["allowedLevel"] = o.AllowedLevel
	}
	if !common.IsNil(o.Capability) {
		toSerialize["capability"] = o.Capability
	}
	if !common.IsNil(o.Problems) {
		toSerialize["problems"] = o.Problems
	}
	toSerialize["requested"] = o.Requested
	toSerialize["requestedLevel"] = o.RequestedLevel
	if !common.IsNil(o.VerificationDeadline) {
		toSerialize["verificationDeadline"] = o.VerificationDeadline
	}
	if !common.IsNil(o.VerificationStatus) {
		toSerialize["verificationStatus"] = o.VerificationStatus
	}
	return toSerialize, nil
}

type NullableAccountCapabilityData struct {
	value *AccountCapabilityData
	isSet bool
}

func (v NullableAccountCapabilityData) Get() *AccountCapabilityData {
	return v.value
}

func (v *NullableAccountCapabilityData) Set(val *AccountCapabilityData) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountCapabilityData) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountCapabilityData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountCapabilityData(val *AccountCapabilityData) *NullableAccountCapabilityData {
	return &NullableAccountCapabilityData{value: val, isSet: true}
}

func (v NullableAccountCapabilityData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountCapabilityData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
