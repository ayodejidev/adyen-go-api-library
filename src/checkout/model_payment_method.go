/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v11/src/common"
)

// checks if the PaymentMethod type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PaymentMethod{}

// PaymentMethod struct for PaymentMethod
type PaymentMethod struct {
	// A list of apps for this payment method.
	Apps []PaymentMethodUPIApps `json:"apps,omitempty"`
	// Brand for the selected gift card. For example: plastix, hmclub.
	Brand *string `json:"brand,omitempty"`
	// List of possible brands. For example: visa, mc.
	Brands []string `json:"brands,omitempty"`
	// The configuration of the payment method.
	Configuration *map[string]string `json:"configuration,omitempty"`
	// The funding source of the payment method.
	FundingSource *string             `json:"fundingSource,omitempty"`
	Group         *PaymentMethodGroup `json:"group,omitempty"`
	// All input details to be provided to complete the payment with this payment method.
	// Deprecated
	InputDetails []InputDetail `json:"inputDetails,omitempty"`
	// A list of issuers for this payment method.
	Issuers []PaymentMethodIssuer `json:"issuers,omitempty"`
	// The displayable name of this payment method.
	Name *string `json:"name,omitempty"`
	// The unique payment method code.
	Type *string `json:"type,omitempty"`
}

// NewPaymentMethod instantiates a new PaymentMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethod() *PaymentMethod {
	this := PaymentMethod{}
	return &this
}

// NewPaymentMethodWithDefaults instantiates a new PaymentMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodWithDefaults() *PaymentMethod {
	this := PaymentMethod{}
	return &this
}

// GetApps returns the Apps field value if set, zero value otherwise.
func (o *PaymentMethod) GetApps() []PaymentMethodUPIApps {
	if o == nil || common.IsNil(o.Apps) {
		var ret []PaymentMethodUPIApps
		return ret
	}
	return o.Apps
}

// GetAppsOk returns a tuple with the Apps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetAppsOk() ([]PaymentMethodUPIApps, bool) {
	if o == nil || common.IsNil(o.Apps) {
		return nil, false
	}
	return o.Apps, true
}

// HasApps returns a boolean if a field has been set.
func (o *PaymentMethod) HasApps() bool {
	if o != nil && !common.IsNil(o.Apps) {
		return true
	}

	return false
}

// SetApps gets a reference to the given []PaymentMethodUPIApps and assigns it to the Apps field.
func (o *PaymentMethod) SetApps(v []PaymentMethodUPIApps) {
	o.Apps = v
}

// GetBrand returns the Brand field value if set, zero value otherwise.
func (o *PaymentMethod) GetBrand() string {
	if o == nil || common.IsNil(o.Brand) {
		var ret string
		return ret
	}
	return *o.Brand
}

// GetBrandOk returns a tuple with the Brand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetBrandOk() (*string, bool) {
	if o == nil || common.IsNil(o.Brand) {
		return nil, false
	}
	return o.Brand, true
}

// HasBrand returns a boolean if a field has been set.
func (o *PaymentMethod) HasBrand() bool {
	if o != nil && !common.IsNil(o.Brand) {
		return true
	}

	return false
}

// SetBrand gets a reference to the given string and assigns it to the Brand field.
func (o *PaymentMethod) SetBrand(v string) {
	o.Brand = &v
}

// GetBrands returns the Brands field value if set, zero value otherwise.
func (o *PaymentMethod) GetBrands() []string {
	if o == nil || common.IsNil(o.Brands) {
		var ret []string
		return ret
	}
	return o.Brands
}

// GetBrandsOk returns a tuple with the Brands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetBrandsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Brands) {
		return nil, false
	}
	return o.Brands, true
}

// HasBrands returns a boolean if a field has been set.
func (o *PaymentMethod) HasBrands() bool {
	if o != nil && !common.IsNil(o.Brands) {
		return true
	}

	return false
}

// SetBrands gets a reference to the given []string and assigns it to the Brands field.
func (o *PaymentMethod) SetBrands(v []string) {
	o.Brands = v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *PaymentMethod) GetConfiguration() map[string]string {
	if o == nil || common.IsNil(o.Configuration) {
		var ret map[string]string
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetConfigurationOk() (*map[string]string, bool) {
	if o == nil || common.IsNil(o.Configuration) {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *PaymentMethod) HasConfiguration() bool {
	if o != nil && !common.IsNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given map[string]string and assigns it to the Configuration field.
func (o *PaymentMethod) SetConfiguration(v map[string]string) {
	o.Configuration = &v
}

// GetFundingSource returns the FundingSource field value if set, zero value otherwise.
func (o *PaymentMethod) GetFundingSource() string {
	if o == nil || common.IsNil(o.FundingSource) {
		var ret string
		return ret
	}
	return *o.FundingSource
}

// GetFundingSourceOk returns a tuple with the FundingSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetFundingSourceOk() (*string, bool) {
	if o == nil || common.IsNil(o.FundingSource) {
		return nil, false
	}
	return o.FundingSource, true
}

// HasFundingSource returns a boolean if a field has been set.
func (o *PaymentMethod) HasFundingSource() bool {
	if o != nil && !common.IsNil(o.FundingSource) {
		return true
	}

	return false
}

// SetFundingSource gets a reference to the given string and assigns it to the FundingSource field.
func (o *PaymentMethod) SetFundingSource(v string) {
	o.FundingSource = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *PaymentMethod) GetGroup() PaymentMethodGroup {
	if o == nil || common.IsNil(o.Group) {
		var ret PaymentMethodGroup
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetGroupOk() (*PaymentMethodGroup, bool) {
	if o == nil || common.IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *PaymentMethod) HasGroup() bool {
	if o != nil && !common.IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given PaymentMethodGroup and assigns it to the Group field.
func (o *PaymentMethod) SetGroup(v PaymentMethodGroup) {
	o.Group = &v
}

// GetInputDetails returns the InputDetails field value if set, zero value otherwise.
// Deprecated
func (o *PaymentMethod) GetInputDetails() []InputDetail {
	if o == nil || common.IsNil(o.InputDetails) {
		var ret []InputDetail
		return ret
	}
	return o.InputDetails
}

// GetInputDetailsOk returns a tuple with the InputDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *PaymentMethod) GetInputDetailsOk() ([]InputDetail, bool) {
	if o == nil || common.IsNil(o.InputDetails) {
		return nil, false
	}
	return o.InputDetails, true
}

// HasInputDetails returns a boolean if a field has been set.
func (o *PaymentMethod) HasInputDetails() bool {
	if o != nil && !common.IsNil(o.InputDetails) {
		return true
	}

	return false
}

// SetInputDetails gets a reference to the given []InputDetail and assigns it to the InputDetails field.
// Deprecated
func (o *PaymentMethod) SetInputDetails(v []InputDetail) {
	o.InputDetails = v
}

// GetIssuers returns the Issuers field value if set, zero value otherwise.
func (o *PaymentMethod) GetIssuers() []PaymentMethodIssuer {
	if o == nil || common.IsNil(o.Issuers) {
		var ret []PaymentMethodIssuer
		return ret
	}
	return o.Issuers
}

// GetIssuersOk returns a tuple with the Issuers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetIssuersOk() ([]PaymentMethodIssuer, bool) {
	if o == nil || common.IsNil(o.Issuers) {
		return nil, false
	}
	return o.Issuers, true
}

// HasIssuers returns a boolean if a field has been set.
func (o *PaymentMethod) HasIssuers() bool {
	if o != nil && !common.IsNil(o.Issuers) {
		return true
	}

	return false
}

// SetIssuers gets a reference to the given []PaymentMethodIssuer and assigns it to the Issuers field.
func (o *PaymentMethod) SetIssuers(v []PaymentMethodIssuer) {
	o.Issuers = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PaymentMethod) GetName() string {
	if o == nil || common.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PaymentMethod) HasName() bool {
	if o != nil && !common.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PaymentMethod) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PaymentMethod) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PaymentMethod) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PaymentMethod) SetType(v string) {
	o.Type = &v
}

func (o PaymentMethod) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Apps) {
		toSerialize["apps"] = o.Apps
	}
	if !common.IsNil(o.Brand) {
		toSerialize["brand"] = o.Brand
	}
	if !common.IsNil(o.Brands) {
		toSerialize["brands"] = o.Brands
	}
	if !common.IsNil(o.Configuration) {
		toSerialize["configuration"] = o.Configuration
	}
	if !common.IsNil(o.FundingSource) {
		toSerialize["fundingSource"] = o.FundingSource
	}
	if !common.IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !common.IsNil(o.InputDetails) {
		toSerialize["inputDetails"] = o.InputDetails
	}
	if !common.IsNil(o.Issuers) {
		toSerialize["issuers"] = o.Issuers
	}
	if !common.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullablePaymentMethod struct {
	value *PaymentMethod
	isSet bool
}

func (v NullablePaymentMethod) Get() *PaymentMethod {
	return v.value
}

func (v *NullablePaymentMethod) Set(val *PaymentMethod) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethod) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethod(val *PaymentMethod) *NullablePaymentMethod {
	return &NullablePaymentMethod{value: val, isSet: true}
}

func (v NullablePaymentMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *PaymentMethod) isValidFundingSource() bool {
	var allowedEnumValues = []string{"credit", "debit"}
	for _, allowed := range allowedEnumValues {
		if o.GetFundingSource() == allowed {
			return true
		}
	}
	return false
}
