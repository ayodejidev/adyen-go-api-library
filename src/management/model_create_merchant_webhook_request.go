/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v19/src/common"
)

// checks if the CreateMerchantWebhookRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CreateMerchantWebhookRequest{}

// CreateMerchantWebhookRequest struct for CreateMerchantWebhookRequest
type CreateMerchantWebhookRequest struct {
	// Indicates if expired SSL certificates are accepted. Default value: **false**.
	AcceptsExpiredCertificate *bool `json:"acceptsExpiredCertificate,omitempty"`
	// Indicates if self-signed SSL certificates are accepted. Default value: **false**.
	AcceptsSelfSignedCertificate *bool `json:"acceptsSelfSignedCertificate,omitempty"`
	// Indicates if untrusted SSL certificates are accepted. Default value: **false**.
	AcceptsUntrustedRootCertificate *bool `json:"acceptsUntrustedRootCertificate,omitempty"`
	// Indicates if the webhook configuration is active. The field must be **true** for us to send webhooks about events related an account.
	Active             bool                `json:"active"`
	AdditionalSettings *AdditionalSettings `json:"additionalSettings,omitempty"`
	// Format or protocol for receiving webhooks. Possible values: * **soap** * **http** * **json**
	CommunicationFormat string `json:"communicationFormat"`
	// Your description for this webhook configuration.
	Description *string `json:"description,omitempty"`
	// SSL version to access the public webhook URL specified in the `url` field. Possible values: * **TLSv1.3** * **TLSv1.2** * **HTTP** - Only allowed on Test environment.  If not specified, the webhook will use `sslVersion`: **TLSv1.2**.
	EncryptionProtocol *string `json:"encryptionProtocol,omitempty"`
	// Network type for Terminal API notification webhooks. Possible values: * **public** * **local**  Default Value: **public**.
	NetworkType *string `json:"networkType,omitempty"`
	// Password to access the webhook URL.
	Password *string `json:"password,omitempty"`
	// Indicates if the SOAP action header needs to be populated. Default value: **false**.  Only applies if `communicationFormat`: **soap**.
	PopulateSoapActionHeader *bool `json:"populateSoapActionHeader,omitempty"`
	// The type of webhook that is being created. Possible values are:  - **standard** - **account-settings-notification** - **banktransfer-notification** - **boletobancario-notification** - **directdebit-notification** - **ach-notification-of-change-notification** - **pending-notification** - **ideal-notification** - **ideal-pending-notification** - **report-notification** - **rreq-notification** - **terminal-settings** - **terminal-boarding**  Find out more about [standard notification webhooks](https://docs.adyen.com/development-resources/webhooks/understand-notifications#event-codes) and [other types of notifications](https://docs.adyen.com/development-resources/webhooks/understand-notifications#other-notifications).
	Type string `json:"type"`
	// Public URL where webhooks will be sent, for example **https://www.domain.com/webhook-endpoint**.
	Url string `json:"url"`
	// Username to access the webhook URL.
	Username *string `json:"username,omitempty"`
}

// NewCreateMerchantWebhookRequest instantiates a new CreateMerchantWebhookRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMerchantWebhookRequest(active bool, communicationFormat string, type_ string, url string) *CreateMerchantWebhookRequest {
	this := CreateMerchantWebhookRequest{}
	this.Active = active
	this.CommunicationFormat = communicationFormat
	this.Type = type_
	this.Url = url
	return &this
}

// NewCreateMerchantWebhookRequestWithDefaults instantiates a new CreateMerchantWebhookRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMerchantWebhookRequestWithDefaults() *CreateMerchantWebhookRequest {
	this := CreateMerchantWebhookRequest{}
	return &this
}

// GetAcceptsExpiredCertificate returns the AcceptsExpiredCertificate field value if set, zero value otherwise.
func (o *CreateMerchantWebhookRequest) GetAcceptsExpiredCertificate() bool {
	if o == nil || common.IsNil(o.AcceptsExpiredCertificate) {
		var ret bool
		return ret
	}
	return *o.AcceptsExpiredCertificate
}

// GetAcceptsExpiredCertificateOk returns a tuple with the AcceptsExpiredCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMerchantWebhookRequest) GetAcceptsExpiredCertificateOk() (*bool, bool) {
	if o == nil || common.IsNil(o.AcceptsExpiredCertificate) {
		return nil, false
	}
	return o.AcceptsExpiredCertificate, true
}

// HasAcceptsExpiredCertificate returns a boolean if a field has been set.
func (o *CreateMerchantWebhookRequest) HasAcceptsExpiredCertificate() bool {
	if o != nil && !common.IsNil(o.AcceptsExpiredCertificate) {
		return true
	}

	return false
}

// SetAcceptsExpiredCertificate gets a reference to the given bool and assigns it to the AcceptsExpiredCertificate field.
func (o *CreateMerchantWebhookRequest) SetAcceptsExpiredCertificate(v bool) {
	o.AcceptsExpiredCertificate = &v
}

// GetAcceptsSelfSignedCertificate returns the AcceptsSelfSignedCertificate field value if set, zero value otherwise.
func (o *CreateMerchantWebhookRequest) GetAcceptsSelfSignedCertificate() bool {
	if o == nil || common.IsNil(o.AcceptsSelfSignedCertificate) {
		var ret bool
		return ret
	}
	return *o.AcceptsSelfSignedCertificate
}

// GetAcceptsSelfSignedCertificateOk returns a tuple with the AcceptsSelfSignedCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMerchantWebhookRequest) GetAcceptsSelfSignedCertificateOk() (*bool, bool) {
	if o == nil || common.IsNil(o.AcceptsSelfSignedCertificate) {
		return nil, false
	}
	return o.AcceptsSelfSignedCertificate, true
}

// HasAcceptsSelfSignedCertificate returns a boolean if a field has been set.
func (o *CreateMerchantWebhookRequest) HasAcceptsSelfSignedCertificate() bool {
	if o != nil && !common.IsNil(o.AcceptsSelfSignedCertificate) {
		return true
	}

	return false
}

// SetAcceptsSelfSignedCertificate gets a reference to the given bool and assigns it to the AcceptsSelfSignedCertificate field.
func (o *CreateMerchantWebhookRequest) SetAcceptsSelfSignedCertificate(v bool) {
	o.AcceptsSelfSignedCertificate = &v
}

// GetAcceptsUntrustedRootCertificate returns the AcceptsUntrustedRootCertificate field value if set, zero value otherwise.
func (o *CreateMerchantWebhookRequest) GetAcceptsUntrustedRootCertificate() bool {
	if o == nil || common.IsNil(o.AcceptsUntrustedRootCertificate) {
		var ret bool
		return ret
	}
	return *o.AcceptsUntrustedRootCertificate
}

// GetAcceptsUntrustedRootCertificateOk returns a tuple with the AcceptsUntrustedRootCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMerchantWebhookRequest) GetAcceptsUntrustedRootCertificateOk() (*bool, bool) {
	if o == nil || common.IsNil(o.AcceptsUntrustedRootCertificate) {
		return nil, false
	}
	return o.AcceptsUntrustedRootCertificate, true
}

// HasAcceptsUntrustedRootCertificate returns a boolean if a field has been set.
func (o *CreateMerchantWebhookRequest) HasAcceptsUntrustedRootCertificate() bool {
	if o != nil && !common.IsNil(o.AcceptsUntrustedRootCertificate) {
		return true
	}

	return false
}

// SetAcceptsUntrustedRootCertificate gets a reference to the given bool and assigns it to the AcceptsUntrustedRootCertificate field.
func (o *CreateMerchantWebhookRequest) SetAcceptsUntrustedRootCertificate(v bool) {
	o.AcceptsUntrustedRootCertificate = &v
}

// GetActive returns the Active field value
func (o *CreateMerchantWebhookRequest) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *CreateMerchantWebhookRequest) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *CreateMerchantWebhookRequest) SetActive(v bool) {
	o.Active = v
}

// GetAdditionalSettings returns the AdditionalSettings field value if set, zero value otherwise.
func (o *CreateMerchantWebhookRequest) GetAdditionalSettings() AdditionalSettings {
	if o == nil || common.IsNil(o.AdditionalSettings) {
		var ret AdditionalSettings
		return ret
	}
	return *o.AdditionalSettings
}

// GetAdditionalSettingsOk returns a tuple with the AdditionalSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMerchantWebhookRequest) GetAdditionalSettingsOk() (*AdditionalSettings, bool) {
	if o == nil || common.IsNil(o.AdditionalSettings) {
		return nil, false
	}
	return o.AdditionalSettings, true
}

// HasAdditionalSettings returns a boolean if a field has been set.
func (o *CreateMerchantWebhookRequest) HasAdditionalSettings() bool {
	if o != nil && !common.IsNil(o.AdditionalSettings) {
		return true
	}

	return false
}

// SetAdditionalSettings gets a reference to the given AdditionalSettings and assigns it to the AdditionalSettings field.
func (o *CreateMerchantWebhookRequest) SetAdditionalSettings(v AdditionalSettings) {
	o.AdditionalSettings = &v
}

// GetCommunicationFormat returns the CommunicationFormat field value
func (o *CreateMerchantWebhookRequest) GetCommunicationFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CommunicationFormat
}

// GetCommunicationFormatOk returns a tuple with the CommunicationFormat field value
// and a boolean to check if the value has been set.
func (o *CreateMerchantWebhookRequest) GetCommunicationFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommunicationFormat, true
}

// SetCommunicationFormat sets field value
func (o *CreateMerchantWebhookRequest) SetCommunicationFormat(v string) {
	o.CommunicationFormat = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateMerchantWebhookRequest) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMerchantWebhookRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateMerchantWebhookRequest) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateMerchantWebhookRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEncryptionProtocol returns the EncryptionProtocol field value if set, zero value otherwise.
func (o *CreateMerchantWebhookRequest) GetEncryptionProtocol() string {
	if o == nil || common.IsNil(o.EncryptionProtocol) {
		var ret string
		return ret
	}
	return *o.EncryptionProtocol
}

// GetEncryptionProtocolOk returns a tuple with the EncryptionProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMerchantWebhookRequest) GetEncryptionProtocolOk() (*string, bool) {
	if o == nil || common.IsNil(o.EncryptionProtocol) {
		return nil, false
	}
	return o.EncryptionProtocol, true
}

// HasEncryptionProtocol returns a boolean if a field has been set.
func (o *CreateMerchantWebhookRequest) HasEncryptionProtocol() bool {
	if o != nil && !common.IsNil(o.EncryptionProtocol) {
		return true
	}

	return false
}

// SetEncryptionProtocol gets a reference to the given string and assigns it to the EncryptionProtocol field.
func (o *CreateMerchantWebhookRequest) SetEncryptionProtocol(v string) {
	o.EncryptionProtocol = &v
}

// GetNetworkType returns the NetworkType field value if set, zero value otherwise.
func (o *CreateMerchantWebhookRequest) GetNetworkType() string {
	if o == nil || common.IsNil(o.NetworkType) {
		var ret string
		return ret
	}
	return *o.NetworkType
}

// GetNetworkTypeOk returns a tuple with the NetworkType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMerchantWebhookRequest) GetNetworkTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.NetworkType) {
		return nil, false
	}
	return o.NetworkType, true
}

// HasNetworkType returns a boolean if a field has been set.
func (o *CreateMerchantWebhookRequest) HasNetworkType() bool {
	if o != nil && !common.IsNil(o.NetworkType) {
		return true
	}

	return false
}

// SetNetworkType gets a reference to the given string and assigns it to the NetworkType field.
func (o *CreateMerchantWebhookRequest) SetNetworkType(v string) {
	o.NetworkType = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *CreateMerchantWebhookRequest) GetPassword() string {
	if o == nil || common.IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMerchantWebhookRequest) GetPasswordOk() (*string, bool) {
	if o == nil || common.IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *CreateMerchantWebhookRequest) HasPassword() bool {
	if o != nil && !common.IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *CreateMerchantWebhookRequest) SetPassword(v string) {
	o.Password = &v
}

// GetPopulateSoapActionHeader returns the PopulateSoapActionHeader field value if set, zero value otherwise.
func (o *CreateMerchantWebhookRequest) GetPopulateSoapActionHeader() bool {
	if o == nil || common.IsNil(o.PopulateSoapActionHeader) {
		var ret bool
		return ret
	}
	return *o.PopulateSoapActionHeader
}

// GetPopulateSoapActionHeaderOk returns a tuple with the PopulateSoapActionHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMerchantWebhookRequest) GetPopulateSoapActionHeaderOk() (*bool, bool) {
	if o == nil || common.IsNil(o.PopulateSoapActionHeader) {
		return nil, false
	}
	return o.PopulateSoapActionHeader, true
}

// HasPopulateSoapActionHeader returns a boolean if a field has been set.
func (o *CreateMerchantWebhookRequest) HasPopulateSoapActionHeader() bool {
	if o != nil && !common.IsNil(o.PopulateSoapActionHeader) {
		return true
	}

	return false
}

// SetPopulateSoapActionHeader gets a reference to the given bool and assigns it to the PopulateSoapActionHeader field.
func (o *CreateMerchantWebhookRequest) SetPopulateSoapActionHeader(v bool) {
	o.PopulateSoapActionHeader = &v
}

// GetType returns the Type field value
func (o *CreateMerchantWebhookRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateMerchantWebhookRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateMerchantWebhookRequest) SetType(v string) {
	o.Type = v
}

// GetUrl returns the Url field value
func (o *CreateMerchantWebhookRequest) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *CreateMerchantWebhookRequest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *CreateMerchantWebhookRequest) SetUrl(v string) {
	o.Url = v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *CreateMerchantWebhookRequest) GetUsername() string {
	if o == nil || common.IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMerchantWebhookRequest) GetUsernameOk() (*string, bool) {
	if o == nil || common.IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *CreateMerchantWebhookRequest) HasUsername() bool {
	if o != nil && !common.IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *CreateMerchantWebhookRequest) SetUsername(v string) {
	o.Username = &v
}

func (o CreateMerchantWebhookRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateMerchantWebhookRequest) ToMap() (map[string]interface{}, error) {
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
	toSerialize["active"] = o.Active
	if !common.IsNil(o.AdditionalSettings) {
		toSerialize["additionalSettings"] = o.AdditionalSettings
	}
	toSerialize["communicationFormat"] = o.CommunicationFormat
	if !common.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !common.IsNil(o.EncryptionProtocol) {
		toSerialize["encryptionProtocol"] = o.EncryptionProtocol
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
	toSerialize["type"] = o.Type
	toSerialize["url"] = o.Url
	if !common.IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableCreateMerchantWebhookRequest struct {
	value *CreateMerchantWebhookRequest
	isSet bool
}

func (v NullableCreateMerchantWebhookRequest) Get() *CreateMerchantWebhookRequest {
	return v.value
}

func (v *NullableCreateMerchantWebhookRequest) Set(val *CreateMerchantWebhookRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMerchantWebhookRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMerchantWebhookRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMerchantWebhookRequest(val *CreateMerchantWebhookRequest) *NullableCreateMerchantWebhookRequest {
	return &NullableCreateMerchantWebhookRequest{value: val, isSet: true}
}

func (v NullableCreateMerchantWebhookRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMerchantWebhookRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *CreateMerchantWebhookRequest) isValidCommunicationFormat() bool {
	var allowedEnumValues = []string{"http", "json", "soap"}
	for _, allowed := range allowedEnumValues {
		if o.GetCommunicationFormat() == allowed {
			return true
		}
	}
	return false
}
func (o *CreateMerchantWebhookRequest) isValidEncryptionProtocol() bool {
	var allowedEnumValues = []string{"HTTP", "TLSv1.2", "TLSv1.3"}
	for _, allowed := range allowedEnumValues {
		if o.GetEncryptionProtocol() == allowed {
			return true
		}
	}
	return false
}
func (o *CreateMerchantWebhookRequest) isValidNetworkType() bool {
	var allowedEnumValues = []string{"local", "public"}
	for _, allowed := range allowedEnumValues {
		if o.GetNetworkType() == allowed {
			return true
		}
	}
	return false
}
