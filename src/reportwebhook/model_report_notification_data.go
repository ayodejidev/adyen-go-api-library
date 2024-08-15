/*
Report webhooks

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package reportwebhook

import (
	"encoding/json"
	"time"

	"github.com/adyen/adyen-go-api-library/v12/src/common"
)

// checks if the ReportNotificationData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ReportNotificationData{}

// ReportNotificationData struct for ReportNotificationData
type ReportNotificationData struct {
	AccountHolder  *ResourceReference `json:"accountHolder,omitempty"`
	BalanceAccount *ResourceReference `json:"balanceAccount,omitempty"`
	// The unique identifier of the balance platform.
	BalancePlatform *string `json:"balancePlatform,omitempty"`
	// The date and time when the event was triggered, in ISO 8601 extended format. For example, **2020-12-18T10:15:30+01:00**.
	CreationDate *time.Time `json:"creationDate,omitempty"`
	// The URL at which you can download the report. To download, you must authenticate your GET request with your [API credentials](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/overview).
	DownloadUrl string `json:"downloadUrl"`
	// The filename of the report.
	FileName string `json:"fileName"`
	// Type of report.
	ReportType string `json:"reportType"`
}

// NewReportNotificationData instantiates a new ReportNotificationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportNotificationData(downloadUrl string, fileName string, reportType string) *ReportNotificationData {
	this := ReportNotificationData{}
	this.DownloadUrl = downloadUrl
	this.FileName = fileName
	this.ReportType = reportType
	return &this
}

// NewReportNotificationDataWithDefaults instantiates a new ReportNotificationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportNotificationDataWithDefaults() *ReportNotificationData {
	this := ReportNotificationData{}
	return &this
}

// GetAccountHolder returns the AccountHolder field value if set, zero value otherwise.
func (o *ReportNotificationData) GetAccountHolder() ResourceReference {
	if o == nil || common.IsNil(o.AccountHolder) {
		var ret ResourceReference
		return ret
	}
	return *o.AccountHolder
}

// GetAccountHolderOk returns a tuple with the AccountHolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportNotificationData) GetAccountHolderOk() (*ResourceReference, bool) {
	if o == nil || common.IsNil(o.AccountHolder) {
		return nil, false
	}
	return o.AccountHolder, true
}

// HasAccountHolder returns a boolean if a field has been set.
func (o *ReportNotificationData) HasAccountHolder() bool {
	if o != nil && !common.IsNil(o.AccountHolder) {
		return true
	}

	return false
}

// SetAccountHolder gets a reference to the given ResourceReference and assigns it to the AccountHolder field.
func (o *ReportNotificationData) SetAccountHolder(v ResourceReference) {
	o.AccountHolder = &v
}

// GetBalanceAccount returns the BalanceAccount field value if set, zero value otherwise.
func (o *ReportNotificationData) GetBalanceAccount() ResourceReference {
	if o == nil || common.IsNil(o.BalanceAccount) {
		var ret ResourceReference
		return ret
	}
	return *o.BalanceAccount
}

// GetBalanceAccountOk returns a tuple with the BalanceAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportNotificationData) GetBalanceAccountOk() (*ResourceReference, bool) {
	if o == nil || common.IsNil(o.BalanceAccount) {
		return nil, false
	}
	return o.BalanceAccount, true
}

// HasBalanceAccount returns a boolean if a field has been set.
func (o *ReportNotificationData) HasBalanceAccount() bool {
	if o != nil && !common.IsNil(o.BalanceAccount) {
		return true
	}

	return false
}

// SetBalanceAccount gets a reference to the given ResourceReference and assigns it to the BalanceAccount field.
func (o *ReportNotificationData) SetBalanceAccount(v ResourceReference) {
	o.BalanceAccount = &v
}

// GetBalancePlatform returns the BalancePlatform field value if set, zero value otherwise.
func (o *ReportNotificationData) GetBalancePlatform() string {
	if o == nil || common.IsNil(o.BalancePlatform) {
		var ret string
		return ret
	}
	return *o.BalancePlatform
}

// GetBalancePlatformOk returns a tuple with the BalancePlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportNotificationData) GetBalancePlatformOk() (*string, bool) {
	if o == nil || common.IsNil(o.BalancePlatform) {
		return nil, false
	}
	return o.BalancePlatform, true
}

// HasBalancePlatform returns a boolean if a field has been set.
func (o *ReportNotificationData) HasBalancePlatform() bool {
	if o != nil && !common.IsNil(o.BalancePlatform) {
		return true
	}

	return false
}

// SetBalancePlatform gets a reference to the given string and assigns it to the BalancePlatform field.
func (o *ReportNotificationData) SetBalancePlatform(v string) {
	o.BalancePlatform = &v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *ReportNotificationData) GetCreationDate() time.Time {
	if o == nil || common.IsNil(o.CreationDate) {
		var ret time.Time
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportNotificationData) GetCreationDateOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.CreationDate) {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *ReportNotificationData) HasCreationDate() bool {
	if o != nil && !common.IsNil(o.CreationDate) {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given time.Time and assigns it to the CreationDate field.
func (o *ReportNotificationData) SetCreationDate(v time.Time) {
	o.CreationDate = &v
}

// GetDownloadUrl returns the DownloadUrl field value
func (o *ReportNotificationData) GetDownloadUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DownloadUrl
}

// GetDownloadUrlOk returns a tuple with the DownloadUrl field value
// and a boolean to check if the value has been set.
func (o *ReportNotificationData) GetDownloadUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DownloadUrl, true
}

// SetDownloadUrl sets field value
func (o *ReportNotificationData) SetDownloadUrl(v string) {
	o.DownloadUrl = v
}

// GetFileName returns the FileName field value
func (o *ReportNotificationData) GetFileName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value
// and a boolean to check if the value has been set.
func (o *ReportNotificationData) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileName, true
}

// SetFileName sets field value
func (o *ReportNotificationData) SetFileName(v string) {
	o.FileName = v
}

// GetReportType returns the ReportType field value
func (o *ReportNotificationData) GetReportType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReportType
}

// GetReportTypeOk returns a tuple with the ReportType field value
// and a boolean to check if the value has been set.
func (o *ReportNotificationData) GetReportTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportType, true
}

// SetReportType sets field value
func (o *ReportNotificationData) SetReportType(v string) {
	o.ReportType = v
}

func (o ReportNotificationData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportNotificationData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AccountHolder) {
		toSerialize["accountHolder"] = o.AccountHolder
	}
	if !common.IsNil(o.BalanceAccount) {
		toSerialize["balanceAccount"] = o.BalanceAccount
	}
	if !common.IsNil(o.BalancePlatform) {
		toSerialize["balancePlatform"] = o.BalancePlatform
	}
	if !common.IsNil(o.CreationDate) {
		toSerialize["creationDate"] = o.CreationDate
	}
	toSerialize["downloadUrl"] = o.DownloadUrl
	toSerialize["fileName"] = o.FileName
	toSerialize["reportType"] = o.ReportType
	return toSerialize, nil
}

type NullableReportNotificationData struct {
	value *ReportNotificationData
	isSet bool
}

func (v NullableReportNotificationData) Get() *ReportNotificationData {
	return v.value
}

func (v *NullableReportNotificationData) Set(val *ReportNotificationData) {
	v.value = val
	v.isSet = true
}

func (v NullableReportNotificationData) IsSet() bool {
	return v.isSet
}

func (v *NullableReportNotificationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportNotificationData(val *ReportNotificationData) *NullableReportNotificationData {
	return &NullableReportNotificationData{value: val, isSet: true}
}

func (v NullableReportNotificationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportNotificationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
