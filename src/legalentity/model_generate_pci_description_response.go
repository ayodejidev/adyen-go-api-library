/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v14/src/common"
)

// checks if the GeneratePciDescriptionResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GeneratePciDescriptionResponse{}

// GeneratePciDescriptionResponse struct for GeneratePciDescriptionResponse
type GeneratePciDescriptionResponse struct {
	// The generated questionnaires in a base64 encoded format.
	Content *string `json:"content,omitempty"`
	// The two-letter [ISO-639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language code for the questionnaire. For example, **en**.
	Language *string `json:"language,omitempty"`
	// The array of Adyen-generated unique identifiers for the questionnaires.
	PciTemplateReferences []string `json:"pciTemplateReferences,omitempty"`
}

// NewGeneratePciDescriptionResponse instantiates a new GeneratePciDescriptionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeneratePciDescriptionResponse() *GeneratePciDescriptionResponse {
	this := GeneratePciDescriptionResponse{}
	return &this
}

// NewGeneratePciDescriptionResponseWithDefaults instantiates a new GeneratePciDescriptionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeneratePciDescriptionResponseWithDefaults() *GeneratePciDescriptionResponse {
	this := GeneratePciDescriptionResponse{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *GeneratePciDescriptionResponse) GetContent() string {
	if o == nil || common.IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneratePciDescriptionResponse) GetContentOk() (*string, bool) {
	if o == nil || common.IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *GeneratePciDescriptionResponse) HasContent() bool {
	if o != nil && !common.IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *GeneratePciDescriptionResponse) SetContent(v string) {
	o.Content = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *GeneratePciDescriptionResponse) GetLanguage() string {
	if o == nil || common.IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneratePciDescriptionResponse) GetLanguageOk() (*string, bool) {
	if o == nil || common.IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *GeneratePciDescriptionResponse) HasLanguage() bool {
	if o != nil && !common.IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *GeneratePciDescriptionResponse) SetLanguage(v string) {
	o.Language = &v
}

// GetPciTemplateReferences returns the PciTemplateReferences field value if set, zero value otherwise.
func (o *GeneratePciDescriptionResponse) GetPciTemplateReferences() []string {
	if o == nil || common.IsNil(o.PciTemplateReferences) {
		var ret []string
		return ret
	}
	return o.PciTemplateReferences
}

// GetPciTemplateReferencesOk returns a tuple with the PciTemplateReferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneratePciDescriptionResponse) GetPciTemplateReferencesOk() ([]string, bool) {
	if o == nil || common.IsNil(o.PciTemplateReferences) {
		return nil, false
	}
	return o.PciTemplateReferences, true
}

// HasPciTemplateReferences returns a boolean if a field has been set.
func (o *GeneratePciDescriptionResponse) HasPciTemplateReferences() bool {
	if o != nil && !common.IsNil(o.PciTemplateReferences) {
		return true
	}

	return false
}

// SetPciTemplateReferences gets a reference to the given []string and assigns it to the PciTemplateReferences field.
func (o *GeneratePciDescriptionResponse) SetPciTemplateReferences(v []string) {
	o.PciTemplateReferences = v
}

func (o GeneratePciDescriptionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GeneratePciDescriptionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !common.IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !common.IsNil(o.PciTemplateReferences) {
		toSerialize["pciTemplateReferences"] = o.PciTemplateReferences
	}
	return toSerialize, nil
}

type NullableGeneratePciDescriptionResponse struct {
	value *GeneratePciDescriptionResponse
	isSet bool
}

func (v NullableGeneratePciDescriptionResponse) Get() *GeneratePciDescriptionResponse {
	return v.value
}

func (v *NullableGeneratePciDescriptionResponse) Set(val *GeneratePciDescriptionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGeneratePciDescriptionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGeneratePciDescriptionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeneratePciDescriptionResponse(val *GeneratePciDescriptionResponse) *NullableGeneratePciDescriptionResponse {
	return &NullableGeneratePciDescriptionResponse{value: val, isSet: true}
}

func (v NullableGeneratePciDescriptionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeneratePciDescriptionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
