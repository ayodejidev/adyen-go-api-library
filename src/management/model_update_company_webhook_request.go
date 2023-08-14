/*
Management API

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// checks if the UpdateCompanyWebhookRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UpdateCompanyWebhookRequest{}

// UpdateCompanyWebhookRequest struct for UpdateCompanyWebhookRequest
type UpdateCompanyWebhookRequest struct {
	// Indicates if expired SSL certificates are accepted. Default value: **false**.
	AcceptsExpiredCertificate *bool `json:"acceptsExpiredCertificate,omitempty"`
	// Indicates if self-signed SSL certificates are accepted. Default value: **false**.
	AcceptsSelfSignedCertificate *bool `json:"acceptsSelfSignedCertificate,omitempty"`
	// Indicates if untrusted SSL certificates are accepted. Default value: **false**.
	AcceptsUntrustedRootCertificate *bool `json:"acceptsUntrustedRootCertificate,omitempty"`
	// Indicates if the webhook configuration is active. The field must be **true** for us to send webhooks about events related an account.
	Active             *bool               `json:"active,omitempty"`
	AdditionalSettings *AdditionalSettings `json:"additionalSettings,omitempty"`
	// Format or protocol for receiving webhooks. Possible values: * **soap** * **http** * **json**
	CommunicationFormat *string `json:"communicationFormat,omitempty"`
	// Your description for this webhook configuration.
	Description *string `json:"description,omitempty"`
	// Shows how merchant accounts are filtered when configuring the webhook. Possible values: * **includeAccounts**: The webhook is configured for the merchant accounts listed in `filterMerchantAccounts`. * **excludeAccounts**: The webhook is not configured for the merchant accounts listed in `filterMerchantAccounts`. * **allAccounts**: Includes all merchant accounts, and does not require specifying `filterMerchantAccounts`.
	FilterMerchantAccountType *string `json:"filterMerchantAccountType,omitempty"`
	// A list of merchant account names that are included or excluded from receiving the webhook. Inclusion or exclusion is based on the value defined for `filterMerchantAccountType`.  Required if `filterMerchantAccountType` is either: * **includeAccounts** * **excludeAccounts**  Not needed for `filterMerchantAccountType`: **allAccounts**.
	FilterMerchantAccounts []string `json:"filterMerchantAccounts,omitempty"`
	// Network type for Terminal API notification webhooks. Possible values: * **public** * **local**  Default Value: **public**.
	NetworkType *string `json:"networkType,omitempty"`
	// Password to access the webhook URL.
	Password *string `json:"password,omitempty"`
	// Indicates if the SOAP action header needs to be populated. Default value: **false**.  Only applies if `communicationFormat`: **soap**.
	PopulateSoapActionHeader *bool `json:"populateSoapActionHeader,omitempty"`
	// SSL version to access the public webhook URL specified in the `url` field. Possible values: * **TLSv1.3** * **TLSv1.2** * **HTTP** - Only allowed on Test environment.  If not specified, the webhook will use `sslVersion`: **TLSv1.2**.
	SslVersion *string `json:"sslVersion,omitempty"`
	// Public URL where webhooks will be sent, for example **https://www.domain.com/webhook-endpoint**.
	Url *string `json:"url,omitempty"`
	// Username to access the webhook URL.
	Username *string `json:"username,omitempty"`
}

// NewUpdateCompanyWebhookRequest instantiates a new UpdateCompanyWebhookRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCompanyWebhookRequest() *UpdateCompanyWebhookRequest {
	this := UpdateCompanyWebhookRequest{}
	return &this
}

// NewUpdateCompanyWebhookRequestWithDefaults instantiates a new UpdateCompanyWebhookRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCompanyWebhookRequestWithDefaults() *UpdateCompanyWebhookRequest {
	this := UpdateCompanyWebhookRequest{}
	return &this
}

// GetAcceptsExpiredCertificate returns the AcceptsExpiredCertificate field value if set, zero value otherwise.
func (o *UpdateCompanyWebhookRequest) GetAcceptsExpiredCertificate() bool {
	if o == nil || common.IsNil(o.AcceptsExpiredCertificate) {
		var ret bool
		return ret
	}
	return *o.AcceptsExpiredCertificate
}

// GetAcceptsExpiredCertificateOk returns a tuple with the AcceptsExpiredCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCompanyWebhookRequest) GetAcceptsExpiredCertificateOk() (*bool, bool) {
	if o == nil || common.IsNil(o.AcceptsExpiredCertificate) {
		return nil, false
	}
	return o.AcceptsExpiredCertificate, true
}

// HasAcceptsExpiredCertificate returns a boolean if a field has been set.
func (o *UpdateCompanyWebhookRequest) HasAcceptsExpiredCertificate() bool {
	if o != nil && !common.IsNil(o.AcceptsExpiredCertificate) {
		return true
	}

	return false
}

// SetAcceptsExpiredCertificate gets a reference to the given bool and assigns it to the AcceptsExpiredCertificate field.
func (o *UpdateCompanyWebhookRequest) SetAcceptsExpiredCertificate(v bool) {
	o.AcceptsExpiredCertificate = &v
}

// GetAcceptsSelfSignedCertificate returns the AcceptsSelfSignedCertificate field value if set, zero value otherwise.
func (o *UpdateCompanyWebhookRequest) GetAcceptsSelfSignedCertificate() bool {
	if o == nil || common.IsNil(o.AcceptsSelfSignedCertificate) {
		var ret bool
		return ret
	}
	return *o.AcceptsSelfSignedCertificate
}

// GetAcceptsSelfSignedCertificateOk returns a tuple with the AcceptsSelfSignedCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCompanyWebhookRequest) GetAcceptsSelfSignedCertificateOk() (*bool, bool) {
	if o == nil || common.IsNil(o.AcceptsSelfSignedCertificate) {
		return nil, false
	}
	return o.AcceptsSelfSignedCertificate, true
}

// HasAcceptsSelfSignedCertificate returns a boolean if a field has been set.
func (o *UpdateCompanyWebhookRequest) HasAcceptsSelfSignedCertificate() bool {
	if o != nil && !common.IsNil(o.AcceptsSelfSignedCertificate) {
		return true
	}

	return false
}

// SetAcceptsSelfSignedCertificate gets a reference to the given bool and assigns it to the AcceptsSelfSignedCertificate field.
func (o *UpdateCompanyWebhookRequest) SetAcceptsSelfSignedCertificate(v bool) {
	o.AcceptsSelfSignedCertificate = &v
}

// GetAcceptsUntrustedRootCertificate returns the AcceptsUntrustedRootCertificate field value if set, zero value otherwise.
func (o *UpdateCompanyWebhookRequest) GetAcceptsUntrustedRootCertificate() bool {
	if o == nil || common.IsNil(o.AcceptsUntrustedRootCertificate) {
		var ret bool
		return ret
	}
	return *o.AcceptsUntrustedRootCertificate
}

// GetAcceptsUntrustedRootCertificateOk returns a tuple with the AcceptsUntrustedRootCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCompanyWebhookRequest) GetAcceptsUntrustedRootCertificateOk() (*bool, bool) {
	if o == nil || common.IsNil(o.AcceptsUntrustedRootCertificate) {
		return nil, false
	}
	return o.AcceptsUntrustedRootCertificate, true
}

// HasAcceptsUntrustedRootCertificate returns a boolean if a field has been set.
func (o *UpdateCompanyWebhookRequest) HasAcceptsUntrustedRootCertificate() bool {
	if o != nil && !common.IsNil(o.AcceptsUntrustedRootCertificate) {
		return true
	}

	return false
}

// SetAcceptsUntrustedRootCertificate gets a reference to the given bool and assigns it to the AcceptsUntrustedRootCertificate field.
func (o *UpdateCompanyWebhookRequest) SetAcceptsUntrustedRootCertificate(v bool) {
	o.AcceptsUntrustedRootCertificate = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *UpdateCompanyWebhookRequest) GetActive() bool {
	if o == nil || common.IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCompanyWebhookRequest) GetActiveOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *UpdateCompanyWebhookRequest) HasActive() bool {
	if o != nil && !common.IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *UpdateCompanyWebhookRequest) SetActive(v bool) {
	o.Active = &v
}

// GetAdditionalSettings returns the AdditionalSettings field value if set, zero value otherwise.
func (o *UpdateCompanyWebhookRequest) GetAdditionalSettings() AdditionalSettings {
	if o == nil || common.IsNil(o.AdditionalSettings) {
		var ret AdditionalSettings
		return ret
	}
	return *o.AdditionalSettings
}

// GetAdditionalSettingsOk returns a tuple with the AdditionalSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCompanyWebhookRequest) GetAdditionalSettingsOk() (*AdditionalSettings, bool) {
	if o == nil || common.IsNil(o.AdditionalSettings) {
		return nil, false
	}
	return o.AdditionalSettings, true
}

// HasAdditionalSettings returns a boolean if a field has been set.
func (o *UpdateCompanyWebhookRequest) HasAdditionalSettings() bool {
	if o != nil && !common.IsNil(o.AdditionalSettings) {
		return true
	}

	return false
}

// SetAdditionalSettings gets a reference to the given AdditionalSettings and assigns it to the AdditionalSettings field.
func (o *UpdateCompanyWebhookRequest) SetAdditionalSettings(v AdditionalSettings) {
	o.AdditionalSettings = &v
}

// GetCommunicationFormat returns the CommunicationFormat field value if set, zero value otherwise.
func (o *UpdateCompanyWebhookRequest) GetCommunicationFormat() string {
	if o == nil || common.IsNil(o.CommunicationFormat) {
		var ret string
		return ret
	}
	return *o.CommunicationFormat
}

// GetCommunicationFormatOk returns a tuple with the CommunicationFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCompanyWebhookRequest) GetCommunicationFormatOk() (*string, bool) {
	if o == nil || common.IsNil(o.CommunicationFormat) {
		return nil, false
	}
	return o.CommunicationFormat, true
}

// HasCommunicationFormat returns a boolean if a field has been set.
func (o *UpdateCompanyWebhookRequest) HasCommunicationFormat() bool {
	if o != nil && !common.IsNil(o.CommunicationFormat) {
		return true
	}

	return false
}

// SetCommunicationFormat gets a reference to the given string and assigns it to the CommunicationFormat field.
func (o *UpdateCompanyWebhookRequest) SetCommunicationFormat(v string) {
	o.CommunicationFormat = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateCompanyWebhookRequest) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCompanyWebhookRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateCompanyWebhookRequest) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateCompanyWebhookRequest) SetDescription(v string) {
	o.Description = &v
}

// GetFilterMerchantAccountType returns the FilterMerchantAccountType field value if set, zero value otherwise.
func (o *UpdateCompanyWebhookRequest) GetFilterMerchantAccountType() string {
	if o == nil || common.IsNil(o.FilterMerchantAccountType) {
		var ret string
		return ret
	}
	return *o.FilterMerchantAccountType
}

// GetFilterMerchantAccountTypeOk returns a tuple with the FilterMerchantAccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCompanyWebhookRequest) GetFilterMerchantAccountTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.FilterMerchantAccountType) {
		return nil, false
	}
	return o.FilterMerchantAccountType, true
}

// HasFilterMerchantAccountType returns a boolean if a field has been set.
func (o *UpdateCompanyWebhookRequest) HasFilterMerchantAccountType() bool {
	if o != nil && !common.IsNil(o.FilterMerchantAccountType) {
		return true
	}

	return false
}

// SetFilterMerchantAccountType gets a reference to the given string and assigns it to the FilterMerchantAccountType field.
func (o *UpdateCompanyWebhookRequest) SetFilterMerchantAccountType(v string) {
	o.FilterMerchantAccountType = &v
}

// GetFilterMerchantAccounts returns the FilterMerchantAccounts field value if set, zero value otherwise.
func (o *UpdateCompanyWebhookRequest) GetFilterMerchantAccounts() []string {
	if o == nil || common.IsNil(o.FilterMerchantAccounts) {
		var ret []string
		return ret
	}
	return o.FilterMerchantAccounts
}

// GetFilterMerchantAccountsOk returns a tuple with the FilterMerchantAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCompanyWebhookRequest) GetFilterMerchantAccountsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.FilterMerchantAccounts) {
		return nil, false
	}
	return o.FilterMerchantAccounts, true
}

// HasFilterMerchantAccounts returns a boolean if a field has been set.
func (o *UpdateCompanyWebhookRequest) HasFilterMerchantAccounts() bool {
	if o != nil && !common.IsNil(o.FilterMerchantAccounts) {
		return true
	}

	return false
}

// SetFilterMerchantAccounts gets a reference to the given []string and assigns it to the FilterMerchantAccounts field.
func (o *UpdateCompanyWebhookRequest) SetFilterMerchantAccounts(v []string) {
	o.FilterMerchantAccounts = v
}

// GetNetworkType returns the NetworkType field value if set, zero value otherwise.
func (o *UpdateCompanyWebhookRequest) GetNetworkType() string {
	if o == nil || common.IsNil(o.NetworkType) {
		var ret string
		return ret
	}
	return *o.NetworkType
}

// GetNetworkTypeOk returns a tuple with the NetworkType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCompanyWebhookRequest) GetNetworkTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.NetworkType) {
		return nil, false
	}
	return o.NetworkType, true
}

// HasNetworkType returns a boolean if a field has been set.
func (o *UpdateCompanyWebhookRequest) HasNetworkType() bool {
	if o != nil && !common.IsNil(o.NetworkType) {
		return true
	}

	return false
}

// SetNetworkType gets a reference to the given string and assigns it to the NetworkType field.
func (o *UpdateCompanyWebhookRequest) SetNetworkType(v string) {
	o.NetworkType = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UpdateCompanyWebhookRequest) GetPassword() string {
	if o == nil || common.IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCompanyWebhookRequest) GetPasswordOk() (*string, bool) {
	if o == nil || common.IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UpdateCompanyWebhookRequest) HasPassword() bool {
	if o != nil && !common.IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UpdateCompanyWebhookRequest) SetPassword(v string) {
	o.Password = &v
}

// GetPopulateSoapActionHeader returns the PopulateSoapActionHeader field value if set, zero value otherwise.
func (o *UpdateCompanyWebhookRequest) GetPopulateSoapActionHeader() bool {
	if o == nil || common.IsNil(o.PopulateSoapActionHeader) {
		var ret bool
		return ret
	}
	return *o.PopulateSoapActionHeader
}

// GetPopulateSoapActionHeaderOk returns a tuple with the PopulateSoapActionHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCompanyWebhookRequest) GetPopulateSoapActionHeaderOk() (*bool, bool) {
	if o == nil || common.IsNil(o.PopulateSoapActionHeader) {
		return nil, false
	}
	return o.PopulateSoapActionHeader, true
}

// HasPopulateSoapActionHeader returns a boolean if a field has been set.
func (o *UpdateCompanyWebhookRequest) HasPopulateSoapActionHeader() bool {
	if o != nil && !common.IsNil(o.PopulateSoapActionHeader) {
		return true
	}

	return false
}

// SetPopulateSoapActionHeader gets a reference to the given bool and assigns it to the PopulateSoapActionHeader field.
func (o *UpdateCompanyWebhookRequest) SetPopulateSoapActionHeader(v bool) {
	o.PopulateSoapActionHeader = &v
}

// GetSslVersion returns the SslVersion field value if set, zero value otherwise.
func (o *UpdateCompanyWebhookRequest) GetSslVersion() string {
	if o == nil || common.IsNil(o.SslVersion) {
		var ret string
		return ret
	}
	return *o.SslVersion
}

// GetSslVersionOk returns a tuple with the SslVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCompanyWebhookRequest) GetSslVersionOk() (*string, bool) {
	if o == nil || common.IsNil(o.SslVersion) {
		return nil, false
	}
	return o.SslVersion, true
}

// HasSslVersion returns a boolean if a field has been set.
func (o *UpdateCompanyWebhookRequest) HasSslVersion() bool {
	if o != nil && !common.IsNil(o.SslVersion) {
		return true
	}

	return false
}

// SetSslVersion gets a reference to the given string and assigns it to the SslVersion field.
func (o *UpdateCompanyWebhookRequest) SetSslVersion(v string) {
	o.SslVersion = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *UpdateCompanyWebhookRequest) GetUrl() string {
	if o == nil || common.IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCompanyWebhookRequest) GetUrlOk() (*string, bool) {
	if o == nil || common.IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *UpdateCompanyWebhookRequest) HasUrl() bool {
	if o != nil && !common.IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *UpdateCompanyWebhookRequest) SetUrl(v string) {
	o.Url = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *UpdateCompanyWebhookRequest) GetUsername() string {
	if o == nil || common.IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCompanyWebhookRequest) GetUsernameOk() (*string, bool) {
	if o == nil || common.IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *UpdateCompanyWebhookRequest) HasUsername() bool {
	if o != nil && !common.IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *UpdateCompanyWebhookRequest) SetUsername(v string) {
	o.Username = &v
}

func (o UpdateCompanyWebhookRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateCompanyWebhookRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AcceptsExpiredCertificate) {
		toSerialize["acceptsExpiredCertificate"] = o.AcceptsExpiredCertificate
	}
	if !common.IsNil(o.AcceptsSelfSignedCertificate) {
		toSerialize["acceptsSelfSignedCertificate"] = o.AcceptsSelfSignedCertificate
	}
	if !common.IsNil(o.AcceptsUntrustedRootCertificate) {
		toSerialize["acceptsUntrustedRootCertificate"] = o.AcceptsUntrustedRootCertificate
	}
	if !common.IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !common.IsNil(o.AdditionalSettings) {
		toSerialize["additionalSettings"] = o.AdditionalSettings
	}
	if !common.IsNil(o.CommunicationFormat) {
		toSerialize["communicationFormat"] = o.CommunicationFormat
	}
	if !common.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !common.IsNil(o.FilterMerchantAccountType) {
		toSerialize["filterMerchantAccountType"] = o.FilterMerchantAccountType
	}
	if !common.IsNil(o.FilterMerchantAccounts) {
		toSerialize["filterMerchantAccounts"] = o.FilterMerchantAccounts
	}
	if !common.IsNil(o.NetworkType) {
		toSerialize["networkType"] = o.NetworkType
	}
	if !common.IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !common.IsNil(o.PopulateSoapActionHeader) {
		toSerialize["populateSoapActionHeader"] = o.PopulateSoapActionHeader
	}
	if !common.IsNil(o.SslVersion) {
		toSerialize["sslVersion"] = o.SslVersion
	}
	if !common.IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !common.IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableUpdateCompanyWebhookRequest struct {
	value *UpdateCompanyWebhookRequest
	isSet bool
}

func (v NullableUpdateCompanyWebhookRequest) Get() *UpdateCompanyWebhookRequest {
	return v.value
}

func (v *NullableUpdateCompanyWebhookRequest) Set(val *UpdateCompanyWebhookRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCompanyWebhookRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCompanyWebhookRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCompanyWebhookRequest(val *UpdateCompanyWebhookRequest) *NullableUpdateCompanyWebhookRequest {
	return &NullableUpdateCompanyWebhookRequest{value: val, isSet: true}
}

func (v NullableUpdateCompanyWebhookRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCompanyWebhookRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *UpdateCompanyWebhookRequest) isValidCommunicationFormat() bool {
	var allowedEnumValues = []string{"http", "json", "soap"}
	for _, allowed := range allowedEnumValues {
		if o.GetCommunicationFormat() == allowed {
			return true
		}
	}
	return false
}
func (o *UpdateCompanyWebhookRequest) isValidFilterMerchantAccountType() bool {
	var allowedEnumValues = []string{"allAccounts", "excludeAccounts", "includeAccounts"}
	for _, allowed := range allowedEnumValues {
		if o.GetFilterMerchantAccountType() == allowed {
			return true
		}
	}
	return false
}
func (o *UpdateCompanyWebhookRequest) isValidNetworkType() bool {
	var allowedEnumValues = []string{"local", "public"}
	for _, allowed := range allowedEnumValues {
		if o.GetNetworkType() == allowed {
			return true
		}
	}
	return false
}
func (o *UpdateCompanyWebhookRequest) isValidSslVersion() bool {
	var allowedEnumValues = []string{"HTTP", "TLSv1.2", "TLSv1.3"}
	for _, allowed := range allowedEnumValues {
		if o.GetSslVersion() == allowed {
			return true
		}
	}
	return false
}
