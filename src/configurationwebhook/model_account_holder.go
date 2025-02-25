/*
Configuration webhooks

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationwebhook

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v19/src/common"
)

// checks if the AccountHolder type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountHolder{}

// AccountHolder struct for AccountHolder
type AccountHolder struct {
	// The unique identifier of the [balance platform](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/get/balancePlatforms/{id}__queryParam_id) to which the account holder belongs. Required in the request if your API credentials can be used for multiple balance platforms.
	BalancePlatform *string `json:"balancePlatform,omitempty"`
	// Contains key-value pairs that specify the actions that an account holder can do in your platform. The key is a capability required for your integration. For example, **issueCard** for Issuing. The value is an object containing the settings for the capability.
	Capabilities *map[string]AccountHolderCapability `json:"capabilities,omitempty"`
	// Deprecated
	ContactDetails *ContactDetails `json:"contactDetails,omitempty"`
	// Your description for the account holder.
	Description *string `json:"description,omitempty"`
	// The unique identifier of the account holder.
	Id string `json:"id"`
	// The unique identifier of the [legal entity](https://docs.adyen.com/api-explorer/legalentity/latest/post/legalEntities#responses-200-id) associated with the account holder. Adyen performs a verification process against the legal entity of the account holder.
	LegalEntityId string `json:"legalEntityId"`
	// A set of key and value pairs for general use. The keys do not have specific names and may be used for storing miscellaneous data as desired. > Note that during an update of metadata, the omission of existing key-value pairs will result in the deletion of those key-value pairs.
	Metadata *map[string]string `json:"metadata,omitempty"`
	// The unique identifier of the migrated account holder in the classic integration.
	MigratedAccountHolderCode *string `json:"migratedAccountHolderCode,omitempty"`
	// The ID of the account holder's primary balance account. By default, this is set to the first balance account that you create for the account holder. To assign a different balance account, send a PATCH request.
	PrimaryBalanceAccount *string `json:"primaryBalanceAccount,omitempty"`
	// Your reference for the account holder.
	Reference *string `json:"reference,omitempty"`
	// The status of the account holder.  Possible values:    * **active**: The account holder is active and allowed to use its capabilities. This is the initial status for account holders and balance accounts. You can change this status to **suspended** or **closed**.    * **suspended**: The account holder is temporarily disabled and payouts are blocked. You can change this status to **active** or **closed**.   * **closed**: The account holder and all of its capabilities are permanently disabled. This is a final status and cannot be changed.
	Status *string `json:"status,omitempty"`
	// The time zone of the account holder. For example, **Europe/Amsterdam**. Defaults to the time zone of the balance platform if no time zone is set. For possible values, see the [list of time zone codes](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones).
	TimeZone *string `json:"timeZone,omitempty"`
	// List of verification deadlines and the capabilities that will be disallowed if verification errors are not resolved.
	VerificationDeadlines []VerificationDeadline `json:"verificationDeadlines,omitempty"`
}

// NewAccountHolder instantiates a new AccountHolder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountHolder(id string, legalEntityId string) *AccountHolder {
	this := AccountHolder{}
	this.Id = id
	this.LegalEntityId = legalEntityId
	return &this
}

// NewAccountHolderWithDefaults instantiates a new AccountHolder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountHolderWithDefaults() *AccountHolder {
	this := AccountHolder{}
	return &this
}

// GetBalancePlatform returns the BalancePlatform field value if set, zero value otherwise.
func (o *AccountHolder) GetBalancePlatform() string {
	if o == nil || common.IsNil(o.BalancePlatform) {
		var ret string
		return ret
	}
	return *o.BalancePlatform
}

// GetBalancePlatformOk returns a tuple with the BalancePlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolder) GetBalancePlatformOk() (*string, bool) {
	if o == nil || common.IsNil(o.BalancePlatform) {
		return nil, false
	}
	return o.BalancePlatform, true
}

// HasBalancePlatform returns a boolean if a field has been set.
func (o *AccountHolder) HasBalancePlatform() bool {
	if o != nil && !common.IsNil(o.BalancePlatform) {
		return true
	}

	return false
}

// SetBalancePlatform gets a reference to the given string and assigns it to the BalancePlatform field.
func (o *AccountHolder) SetBalancePlatform(v string) {
	o.BalancePlatform = &v
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *AccountHolder) GetCapabilities() map[string]AccountHolderCapability {
	if o == nil || common.IsNil(o.Capabilities) {
		var ret map[string]AccountHolderCapability
		return ret
	}
	return *o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolder) GetCapabilitiesOk() (*map[string]AccountHolderCapability, bool) {
	if o == nil || common.IsNil(o.Capabilities) {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *AccountHolder) HasCapabilities() bool {
	if o != nil && !common.IsNil(o.Capabilities) {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given map[string]AccountHolderCapability and assigns it to the Capabilities field.
func (o *AccountHolder) SetCapabilities(v map[string]AccountHolderCapability) {
	o.Capabilities = &v
}

// GetContactDetails returns the ContactDetails field value if set, zero value otherwise.
// Deprecated
func (o *AccountHolder) GetContactDetails() ContactDetails {
	if o == nil || common.IsNil(o.ContactDetails) {
		var ret ContactDetails
		return ret
	}
	return *o.ContactDetails
}

// GetContactDetailsOk returns a tuple with the ContactDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AccountHolder) GetContactDetailsOk() (*ContactDetails, bool) {
	if o == nil || common.IsNil(o.ContactDetails) {
		return nil, false
	}
	return o.ContactDetails, true
}

// HasContactDetails returns a boolean if a field has been set.
func (o *AccountHolder) HasContactDetails() bool {
	if o != nil && !common.IsNil(o.ContactDetails) {
		return true
	}

	return false
}

// SetContactDetails gets a reference to the given ContactDetails and assigns it to the ContactDetails field.
// Deprecated
func (o *AccountHolder) SetContactDetails(v ContactDetails) {
	o.ContactDetails = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AccountHolder) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolder) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AccountHolder) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AccountHolder) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value
func (o *AccountHolder) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AccountHolder) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AccountHolder) SetId(v string) {
	o.Id = v
}

// GetLegalEntityId returns the LegalEntityId field value
func (o *AccountHolder) GetLegalEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LegalEntityId
}

// GetLegalEntityIdOk returns a tuple with the LegalEntityId field value
// and a boolean to check if the value has been set.
func (o *AccountHolder) GetLegalEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LegalEntityId, true
}

// SetLegalEntityId sets field value
func (o *AccountHolder) SetLegalEntityId(v string) {
	o.LegalEntityId = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *AccountHolder) GetMetadata() map[string]string {
	if o == nil || common.IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolder) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || common.IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *AccountHolder) HasMetadata() bool {
	if o != nil && !common.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *AccountHolder) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetMigratedAccountHolderCode returns the MigratedAccountHolderCode field value if set, zero value otherwise.
func (o *AccountHolder) GetMigratedAccountHolderCode() string {
	if o == nil || common.IsNil(o.MigratedAccountHolderCode) {
		var ret string
		return ret
	}
	return *o.MigratedAccountHolderCode
}

// GetMigratedAccountHolderCodeOk returns a tuple with the MigratedAccountHolderCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolder) GetMigratedAccountHolderCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.MigratedAccountHolderCode) {
		return nil, false
	}
	return o.MigratedAccountHolderCode, true
}

// HasMigratedAccountHolderCode returns a boolean if a field has been set.
func (o *AccountHolder) HasMigratedAccountHolderCode() bool {
	if o != nil && !common.IsNil(o.MigratedAccountHolderCode) {
		return true
	}

	return false
}

// SetMigratedAccountHolderCode gets a reference to the given string and assigns it to the MigratedAccountHolderCode field.
func (o *AccountHolder) SetMigratedAccountHolderCode(v string) {
	o.MigratedAccountHolderCode = &v
}

// GetPrimaryBalanceAccount returns the PrimaryBalanceAccount field value if set, zero value otherwise.
func (o *AccountHolder) GetPrimaryBalanceAccount() string {
	if o == nil || common.IsNil(o.PrimaryBalanceAccount) {
		var ret string
		return ret
	}
	return *o.PrimaryBalanceAccount
}

// GetPrimaryBalanceAccountOk returns a tuple with the PrimaryBalanceAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolder) GetPrimaryBalanceAccountOk() (*string, bool) {
	if o == nil || common.IsNil(o.PrimaryBalanceAccount) {
		return nil, false
	}
	return o.PrimaryBalanceAccount, true
}

// HasPrimaryBalanceAccount returns a boolean if a field has been set.
func (o *AccountHolder) HasPrimaryBalanceAccount() bool {
	if o != nil && !common.IsNil(o.PrimaryBalanceAccount) {
		return true
	}

	return false
}

// SetPrimaryBalanceAccount gets a reference to the given string and assigns it to the PrimaryBalanceAccount field.
func (o *AccountHolder) SetPrimaryBalanceAccount(v string) {
	o.PrimaryBalanceAccount = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *AccountHolder) GetReference() string {
	if o == nil || common.IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolder) GetReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *AccountHolder) HasReference() bool {
	if o != nil && !common.IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *AccountHolder) SetReference(v string) {
	o.Reference = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AccountHolder) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolder) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AccountHolder) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AccountHolder) SetStatus(v string) {
	o.Status = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *AccountHolder) GetTimeZone() string {
	if o == nil || common.IsNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolder) GetTimeZoneOk() (*string, bool) {
	if o == nil || common.IsNil(o.TimeZone) {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *AccountHolder) HasTimeZone() bool {
	if o != nil && !common.IsNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *AccountHolder) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetVerificationDeadlines returns the VerificationDeadlines field value if set, zero value otherwise.
func (o *AccountHolder) GetVerificationDeadlines() []VerificationDeadline {
	if o == nil || common.IsNil(o.VerificationDeadlines) {
		var ret []VerificationDeadline
		return ret
	}
	return o.VerificationDeadlines
}

// GetVerificationDeadlinesOk returns a tuple with the VerificationDeadlines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolder) GetVerificationDeadlinesOk() ([]VerificationDeadline, bool) {
	if o == nil || common.IsNil(o.VerificationDeadlines) {
		return nil, false
	}
	return o.VerificationDeadlines, true
}

// HasVerificationDeadlines returns a boolean if a field has been set.
func (o *AccountHolder) HasVerificationDeadlines() bool {
	if o != nil && !common.IsNil(o.VerificationDeadlines) {
		return true
	}

	return false
}

// SetVerificationDeadlines gets a reference to the given []VerificationDeadline and assigns it to the VerificationDeadlines field.
func (o *AccountHolder) SetVerificationDeadlines(v []VerificationDeadline) {
	o.VerificationDeadlines = v
}

func (o AccountHolder) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountHolder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.BalancePlatform) {
		toSerialize["balancePlatform"] = o.BalancePlatform
	}
	if !common.IsNil(o.Capabilities) {
		toSerialize["capabilities"] = o.Capabilities
	}
	if !common.IsNil(o.ContactDetails) {
		toSerialize["contactDetails"] = o.ContactDetails
	}
	if !common.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["id"] = o.Id
	toSerialize["legalEntityId"] = o.LegalEntityId
	if !common.IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !common.IsNil(o.MigratedAccountHolderCode) {
		toSerialize["migratedAccountHolderCode"] = o.MigratedAccountHolderCode
	}
	if !common.IsNil(o.PrimaryBalanceAccount) {
		toSerialize["primaryBalanceAccount"] = o.PrimaryBalanceAccount
	}
	if !common.IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.TimeZone) {
		toSerialize["timeZone"] = o.TimeZone
	}
	if !common.IsNil(o.VerificationDeadlines) {
		toSerialize["verificationDeadlines"] = o.VerificationDeadlines
	}
	return toSerialize, nil
}

type NullableAccountHolder struct {
	value *AccountHolder
	isSet bool
}

func (v NullableAccountHolder) Get() *AccountHolder {
	return v.value
}

func (v *NullableAccountHolder) Set(val *AccountHolder) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountHolder) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountHolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountHolder(val *AccountHolder) *NullableAccountHolder {
	return &NullableAccountHolder{value: val, isSet: true}
}

func (v NullableAccountHolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountHolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *AccountHolder) isValidStatus() bool {
	var allowedEnumValues = []string{"active", "closed", "suspended"}
	for _, allowed := range allowedEnumValues {
		if o.GetStatus() == allowed {
			return true
		}
	}
	return false
}
