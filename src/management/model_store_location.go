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

// checks if the StoreLocation type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &StoreLocation{}

// StoreLocation struct for StoreLocation
type StoreLocation struct {
	// The name of the city.
	City *string `json:"city,omitempty"`
	// The two-letter country code in [ISO_3166-1_alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) format.
	Country string `json:"country"`
	// The street address.
	Line1 *string `json:"line1,omitempty"`
	// Second address line.
	Line2 *string `json:"line2,omitempty"`
	// Third address line.
	Line3 *string `json:"line3,omitempty"`
	// The postal code.
	PostalCode *string `json:"postalCode,omitempty"`
	// The state or province code as defined in [ISO 3166-2](https://www.iso.org/standard/72483.html). For example, **ON** for Ontario, Canada.  Required for the following countries:  - Australia - Brazil - Canada - India - Mexico - New Zealand - United States
	StateOrProvince *string `json:"stateOrProvince,omitempty"`
}

// NewStoreLocation instantiates a new StoreLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoreLocation(country string) *StoreLocation {
	this := StoreLocation{}
	this.Country = country
	return &this
}

// NewStoreLocationWithDefaults instantiates a new StoreLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoreLocationWithDefaults() *StoreLocation {
	this := StoreLocation{}
	return &this
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *StoreLocation) GetCity() string {
	if o == nil || common.IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreLocation) GetCityOk() (*string, bool) {
	if o == nil || common.IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *StoreLocation) HasCity() bool {
	if o != nil && !common.IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *StoreLocation) SetCity(v string) {
	o.City = &v
}

// GetCountry returns the Country field value
func (o *StoreLocation) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *StoreLocation) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *StoreLocation) SetCountry(v string) {
	o.Country = v
}

// GetLine1 returns the Line1 field value if set, zero value otherwise.
func (o *StoreLocation) GetLine1() string {
	if o == nil || common.IsNil(o.Line1) {
		var ret string
		return ret
	}
	return *o.Line1
}

// GetLine1Ok returns a tuple with the Line1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreLocation) GetLine1Ok() (*string, bool) {
	if o == nil || common.IsNil(o.Line1) {
		return nil, false
	}
	return o.Line1, true
}

// HasLine1 returns a boolean if a field has been set.
func (o *StoreLocation) HasLine1() bool {
	if o != nil && !common.IsNil(o.Line1) {
		return true
	}

	return false
}

// SetLine1 gets a reference to the given string and assigns it to the Line1 field.
func (o *StoreLocation) SetLine1(v string) {
	o.Line1 = &v
}

// GetLine2 returns the Line2 field value if set, zero value otherwise.
func (o *StoreLocation) GetLine2() string {
	if o == nil || common.IsNil(o.Line2) {
		var ret string
		return ret
	}
	return *o.Line2
}

// GetLine2Ok returns a tuple with the Line2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreLocation) GetLine2Ok() (*string, bool) {
	if o == nil || common.IsNil(o.Line2) {
		return nil, false
	}
	return o.Line2, true
}

// HasLine2 returns a boolean if a field has been set.
func (o *StoreLocation) HasLine2() bool {
	if o != nil && !common.IsNil(o.Line2) {
		return true
	}

	return false
}

// SetLine2 gets a reference to the given string and assigns it to the Line2 field.
func (o *StoreLocation) SetLine2(v string) {
	o.Line2 = &v
}

// GetLine3 returns the Line3 field value if set, zero value otherwise.
func (o *StoreLocation) GetLine3() string {
	if o == nil || common.IsNil(o.Line3) {
		var ret string
		return ret
	}
	return *o.Line3
}

// GetLine3Ok returns a tuple with the Line3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreLocation) GetLine3Ok() (*string, bool) {
	if o == nil || common.IsNil(o.Line3) {
		return nil, false
	}
	return o.Line3, true
}

// HasLine3 returns a boolean if a field has been set.
func (o *StoreLocation) HasLine3() bool {
	if o != nil && !common.IsNil(o.Line3) {
		return true
	}

	return false
}

// SetLine3 gets a reference to the given string and assigns it to the Line3 field.
func (o *StoreLocation) SetLine3(v string) {
	o.Line3 = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *StoreLocation) GetPostalCode() string {
	if o == nil || common.IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreLocation) GetPostalCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *StoreLocation) HasPostalCode() bool {
	if o != nil && !common.IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *StoreLocation) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetStateOrProvince returns the StateOrProvince field value if set, zero value otherwise.
func (o *StoreLocation) GetStateOrProvince() string {
	if o == nil || common.IsNil(o.StateOrProvince) {
		var ret string
		return ret
	}
	return *o.StateOrProvince
}

// GetStateOrProvinceOk returns a tuple with the StateOrProvince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreLocation) GetStateOrProvinceOk() (*string, bool) {
	if o == nil || common.IsNil(o.StateOrProvince) {
		return nil, false
	}
	return o.StateOrProvince, true
}

// HasStateOrProvince returns a boolean if a field has been set.
func (o *StoreLocation) HasStateOrProvince() bool {
	if o != nil && !common.IsNil(o.StateOrProvince) {
		return true
	}

	return false
}

// SetStateOrProvince gets a reference to the given string and assigns it to the StateOrProvince field.
func (o *StoreLocation) SetStateOrProvince(v string) {
	o.StateOrProvince = &v
}

func (o StoreLocation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StoreLocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	toSerialize["country"] = o.Country
	if !common.IsNil(o.Line1) {
		toSerialize["line1"] = o.Line1
	}
	if !common.IsNil(o.Line2) {
		toSerialize["line2"] = o.Line2
	}
	if !common.IsNil(o.Line3) {
		toSerialize["line3"] = o.Line3
	}
	if !common.IsNil(o.PostalCode) {
		toSerialize["postalCode"] = o.PostalCode
	}
	if !common.IsNil(o.StateOrProvince) {
		toSerialize["stateOrProvince"] = o.StateOrProvince
	}
	return toSerialize, nil
}

type NullableStoreLocation struct {
	value *StoreLocation
	isSet bool
}

func (v NullableStoreLocation) Get() *StoreLocation {
	return v.value
}

func (v *NullableStoreLocation) Set(val *StoreLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableStoreLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableStoreLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoreLocation(val *StoreLocation) *NullableStoreLocation {
	return &NullableStoreLocation{value: val, isSet: true}
}

func (v NullableStoreLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoreLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
