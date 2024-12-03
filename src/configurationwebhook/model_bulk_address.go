/*
Configuration webhooks

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationwebhook

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v15/src/common"
)

// checks if the BulkAddress type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BulkAddress{}

// BulkAddress struct for BulkAddress
type BulkAddress struct {
	// The name of the city.
	City *string `json:"city,omitempty"`
	// The name of the company.
	Company *string `json:"company,omitempty"`
	// The two-character ISO-3166-1 alpha-2 country code. For example, **US**.
	Country string `json:"country"`
	// The email address.
	Email *string `json:"email,omitempty"`
	// The house number or name.
	HouseNumberOrName *string `json:"houseNumberOrName,omitempty"`
	// The full telephone number.
	Mobile *string `json:"mobile,omitempty"`
	// The postal code.  Maximum length:  * 5 digits for addresses in the US.  * 10 characters for all other countries.
	PostalCode *string `json:"postalCode,omitempty"`
	// The two-letter ISO 3166-2 state or province code.  Maximum length: 2 characters for addresses in the US.
	StateOrProvince *string `json:"stateOrProvince,omitempty"`
	// The streetname of the house.
	Street *string `json:"street,omitempty"`
}

// NewBulkAddress instantiates a new BulkAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkAddress(country string) *BulkAddress {
	this := BulkAddress{}
	this.Country = country
	return &this
}

// NewBulkAddressWithDefaults instantiates a new BulkAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkAddressWithDefaults() *BulkAddress {
	this := BulkAddress{}
	return &this
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *BulkAddress) GetCity() string {
	if o == nil || common.IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkAddress) GetCityOk() (*string, bool) {
	if o == nil || common.IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *BulkAddress) HasCity() bool {
	if o != nil && !common.IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *BulkAddress) SetCity(v string) {
	o.City = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *BulkAddress) GetCompany() string {
	if o == nil || common.IsNil(o.Company) {
		var ret string
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkAddress) GetCompanyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *BulkAddress) HasCompany() bool {
	if o != nil && !common.IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given string and assigns it to the Company field.
func (o *BulkAddress) SetCompany(v string) {
	o.Company = &v
}

// GetCountry returns the Country field value
func (o *BulkAddress) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *BulkAddress) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *BulkAddress) SetCountry(v string) {
	o.Country = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *BulkAddress) GetEmail() string {
	if o == nil || common.IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkAddress) GetEmailOk() (*string, bool) {
	if o == nil || common.IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *BulkAddress) HasEmail() bool {
	if o != nil && !common.IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *BulkAddress) SetEmail(v string) {
	o.Email = &v
}

// GetHouseNumberOrName returns the HouseNumberOrName field value if set, zero value otherwise.
func (o *BulkAddress) GetHouseNumberOrName() string {
	if o == nil || common.IsNil(o.HouseNumberOrName) {
		var ret string
		return ret
	}
	return *o.HouseNumberOrName
}

// GetHouseNumberOrNameOk returns a tuple with the HouseNumberOrName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkAddress) GetHouseNumberOrNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.HouseNumberOrName) {
		return nil, false
	}
	return o.HouseNumberOrName, true
}

// HasHouseNumberOrName returns a boolean if a field has been set.
func (o *BulkAddress) HasHouseNumberOrName() bool {
	if o != nil && !common.IsNil(o.HouseNumberOrName) {
		return true
	}

	return false
}

// SetHouseNumberOrName gets a reference to the given string and assigns it to the HouseNumberOrName field.
func (o *BulkAddress) SetHouseNumberOrName(v string) {
	o.HouseNumberOrName = &v
}

// GetMobile returns the Mobile field value if set, zero value otherwise.
func (o *BulkAddress) GetMobile() string {
	if o == nil || common.IsNil(o.Mobile) {
		var ret string
		return ret
	}
	return *o.Mobile
}

// GetMobileOk returns a tuple with the Mobile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkAddress) GetMobileOk() (*string, bool) {
	if o == nil || common.IsNil(o.Mobile) {
		return nil, false
	}
	return o.Mobile, true
}

// HasMobile returns a boolean if a field has been set.
func (o *BulkAddress) HasMobile() bool {
	if o != nil && !common.IsNil(o.Mobile) {
		return true
	}

	return false
}

// SetMobile gets a reference to the given string and assigns it to the Mobile field.
func (o *BulkAddress) SetMobile(v string) {
	o.Mobile = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *BulkAddress) GetPostalCode() string {
	if o == nil || common.IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkAddress) GetPostalCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *BulkAddress) HasPostalCode() bool {
	if o != nil && !common.IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *BulkAddress) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetStateOrProvince returns the StateOrProvince field value if set, zero value otherwise.
func (o *BulkAddress) GetStateOrProvince() string {
	if o == nil || common.IsNil(o.StateOrProvince) {
		var ret string
		return ret
	}
	return *o.StateOrProvince
}

// GetStateOrProvinceOk returns a tuple with the StateOrProvince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkAddress) GetStateOrProvinceOk() (*string, bool) {
	if o == nil || common.IsNil(o.StateOrProvince) {
		return nil, false
	}
	return o.StateOrProvince, true
}

// HasStateOrProvince returns a boolean if a field has been set.
func (o *BulkAddress) HasStateOrProvince() bool {
	if o != nil && !common.IsNil(o.StateOrProvince) {
		return true
	}

	return false
}

// SetStateOrProvince gets a reference to the given string and assigns it to the StateOrProvince field.
func (o *BulkAddress) SetStateOrProvince(v string) {
	o.StateOrProvince = &v
}

// GetStreet returns the Street field value if set, zero value otherwise.
func (o *BulkAddress) GetStreet() string {
	if o == nil || common.IsNil(o.Street) {
		var ret string
		return ret
	}
	return *o.Street
}

// GetStreetOk returns a tuple with the Street field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkAddress) GetStreetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Street) {
		return nil, false
	}
	return o.Street, true
}

// HasStreet returns a boolean if a field has been set.
func (o *BulkAddress) HasStreet() bool {
	if o != nil && !common.IsNil(o.Street) {
		return true
	}

	return false
}

// SetStreet gets a reference to the given string and assigns it to the Street field.
func (o *BulkAddress) SetStreet(v string) {
	o.Street = &v
}

func (o BulkAddress) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !common.IsNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	toSerialize["country"] = o.Country
	if !common.IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !common.IsNil(o.HouseNumberOrName) {
		toSerialize["houseNumberOrName"] = o.HouseNumberOrName
	}
	if !common.IsNil(o.Mobile) {
		toSerialize["mobile"] = o.Mobile
	}
	if !common.IsNil(o.PostalCode) {
		toSerialize["postalCode"] = o.PostalCode
	}
	if !common.IsNil(o.StateOrProvince) {
		toSerialize["stateOrProvince"] = o.StateOrProvince
	}
	if !common.IsNil(o.Street) {
		toSerialize["street"] = o.Street
	}
	return toSerialize, nil
}

type NullableBulkAddress struct {
	value *BulkAddress
	isSet bool
}

func (v NullableBulkAddress) Get() *BulkAddress {
	return v.value
}

func (v *NullableBulkAddress) Set(val *BulkAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkAddress(val *BulkAddress) *NullableBulkAddress {
	return &NullableBulkAddress{value: val, isSet: true}
}

func (v NullableBulkAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
