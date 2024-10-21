/*
Adyen Payment API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package payments

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v14/src/common"
)

// checks if the ResponseAdditionalDataOpi type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ResponseAdditionalDataOpi{}

// ResponseAdditionalDataOpi struct for ResponseAdditionalDataOpi
type ResponseAdditionalDataOpi struct {
	// Returned in the response if you included `opi.includeTransToken: true` in an ecommerce payment request. This contains an Oracle Payment Interface token that you can store in your Oracle Opera database to identify tokenized ecommerce transactions. For more information and required settings, see [Oracle Opera](https://docs.adyen.com/plugins/oracle-opera#opi-token-ecommerce).
	OpiTransToken *string `json:"opi.transToken,omitempty"`
}

// NewResponseAdditionalDataOpi instantiates a new ResponseAdditionalDataOpi object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseAdditionalDataOpi() *ResponseAdditionalDataOpi {
	this := ResponseAdditionalDataOpi{}
	return &this
}

// NewResponseAdditionalDataOpiWithDefaults instantiates a new ResponseAdditionalDataOpi object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseAdditionalDataOpiWithDefaults() *ResponseAdditionalDataOpi {
	this := ResponseAdditionalDataOpi{}
	return &this
}

// GetOpiTransToken returns the OpiTransToken field value if set, zero value otherwise.
func (o *ResponseAdditionalDataOpi) GetOpiTransToken() string {
	if o == nil || common.IsNil(o.OpiTransToken) {
		var ret string
		return ret
	}
	return *o.OpiTransToken
}

// GetOpiTransTokenOk returns a tuple with the OpiTransToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAdditionalDataOpi) GetOpiTransTokenOk() (*string, bool) {
	if o == nil || common.IsNil(o.OpiTransToken) {
		return nil, false
	}
	return o.OpiTransToken, true
}

// HasOpiTransToken returns a boolean if a field has been set.
func (o *ResponseAdditionalDataOpi) HasOpiTransToken() bool {
	if o != nil && !common.IsNil(o.OpiTransToken) {
		return true
	}

	return false
}

// SetOpiTransToken gets a reference to the given string and assigns it to the OpiTransToken field.
func (o *ResponseAdditionalDataOpi) SetOpiTransToken(v string) {
	o.OpiTransToken = &v
}

func (o ResponseAdditionalDataOpi) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseAdditionalDataOpi) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.OpiTransToken) {
		toSerialize["opi.transToken"] = o.OpiTransToken
	}
	return toSerialize, nil
}

type NullableResponseAdditionalDataOpi struct {
	value *ResponseAdditionalDataOpi
	isSet bool
}

func (v NullableResponseAdditionalDataOpi) Get() *ResponseAdditionalDataOpi {
	return v.value
}

func (v *NullableResponseAdditionalDataOpi) Set(val *ResponseAdditionalDataOpi) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseAdditionalDataOpi) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseAdditionalDataOpi) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseAdditionalDataOpi(val *ResponseAdditionalDataOpi) *NullableResponseAdditionalDataOpi {
	return &NullableResponseAdditionalDataOpi{value: val, isSet: true}
}

func (v NullableResponseAdditionalDataOpi) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseAdditionalDataOpi) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
