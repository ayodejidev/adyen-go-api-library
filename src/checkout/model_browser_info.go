/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v14/src/common"
)

// checks if the BrowserInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BrowserInfo{}

// BrowserInfo struct for BrowserInfo
type BrowserInfo struct {
	// The accept header value of the shopper's browser.
	AcceptHeader string `json:"acceptHeader"`
	// The color depth of the shopper's browser in bits per pixel. This should be obtained by using the browser's `screen.colorDepth` property. Accepted values: 1, 4, 8, 15, 16, 24, 30, 32 or 48 bit color depth.
	ColorDepth int32 `json:"colorDepth"`
	// Boolean value indicating if the shopper's browser is able to execute Java.
	JavaEnabled bool `json:"javaEnabled"`
	// Boolean value indicating if the shopper's browser is able to execute JavaScript. A default 'true' value is assumed if the field is not present.
	JavaScriptEnabled *bool `json:"javaScriptEnabled,omitempty"`
	// The `navigator.language` value of the shopper's browser (as defined in IETF BCP 47).
	Language string `json:"language"`
	// The total height of the shopper's device screen in pixels.
	ScreenHeight int32 `json:"screenHeight"`
	// The total width of the shopper's device screen in pixels.
	ScreenWidth int32 `json:"screenWidth"`
	// Time difference between UTC time and the shopper's browser local time, in minutes.
	TimeZoneOffset int32 `json:"timeZoneOffset"`
	// The user agent value of the shopper's browser.
	UserAgent string `json:"userAgent"`
}

// NewBrowserInfo instantiates a new BrowserInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrowserInfo(acceptHeader string, colorDepth int32, javaEnabled bool, language string, screenHeight int32, screenWidth int32, timeZoneOffset int32, userAgent string) *BrowserInfo {
	this := BrowserInfo{}
	this.AcceptHeader = acceptHeader
	this.ColorDepth = colorDepth
	this.JavaEnabled = javaEnabled
	var javaScriptEnabled bool = true
	this.JavaScriptEnabled = &javaScriptEnabled
	this.Language = language
	this.ScreenHeight = screenHeight
	this.ScreenWidth = screenWidth
	this.TimeZoneOffset = timeZoneOffset
	this.UserAgent = userAgent
	return &this
}

// NewBrowserInfoWithDefaults instantiates a new BrowserInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrowserInfoWithDefaults() *BrowserInfo {
	this := BrowserInfo{}
	var javaScriptEnabled bool = true
	this.JavaScriptEnabled = &javaScriptEnabled
	return &this
}

// GetAcceptHeader returns the AcceptHeader field value
func (o *BrowserInfo) GetAcceptHeader() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AcceptHeader
}

// GetAcceptHeaderOk returns a tuple with the AcceptHeader field value
// and a boolean to check if the value has been set.
func (o *BrowserInfo) GetAcceptHeaderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AcceptHeader, true
}

// SetAcceptHeader sets field value
func (o *BrowserInfo) SetAcceptHeader(v string) {
	o.AcceptHeader = v
}

// GetColorDepth returns the ColorDepth field value
func (o *BrowserInfo) GetColorDepth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ColorDepth
}

// GetColorDepthOk returns a tuple with the ColorDepth field value
// and a boolean to check if the value has been set.
func (o *BrowserInfo) GetColorDepthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ColorDepth, true
}

// SetColorDepth sets field value
func (o *BrowserInfo) SetColorDepth(v int32) {
	o.ColorDepth = v
}

// GetJavaEnabled returns the JavaEnabled field value
func (o *BrowserInfo) GetJavaEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.JavaEnabled
}

// GetJavaEnabledOk returns a tuple with the JavaEnabled field value
// and a boolean to check if the value has been set.
func (o *BrowserInfo) GetJavaEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JavaEnabled, true
}

// SetJavaEnabled sets field value
func (o *BrowserInfo) SetJavaEnabled(v bool) {
	o.JavaEnabled = v
}

// GetJavaScriptEnabled returns the JavaScriptEnabled field value if set, zero value otherwise.
func (o *BrowserInfo) GetJavaScriptEnabled() bool {
	if o == nil || common.IsNil(o.JavaScriptEnabled) {
		var ret bool
		return ret
	}
	return *o.JavaScriptEnabled
}

// GetJavaScriptEnabledOk returns a tuple with the JavaScriptEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrowserInfo) GetJavaScriptEnabledOk() (*bool, bool) {
	if o == nil || common.IsNil(o.JavaScriptEnabled) {
		return nil, false
	}
	return o.JavaScriptEnabled, true
}

// HasJavaScriptEnabled returns a boolean if a field has been set.
func (o *BrowserInfo) HasJavaScriptEnabled() bool {
	if o != nil && !common.IsNil(o.JavaScriptEnabled) {
		return true
	}

	return false
}

// SetJavaScriptEnabled gets a reference to the given bool and assigns it to the JavaScriptEnabled field.
func (o *BrowserInfo) SetJavaScriptEnabled(v bool) {
	o.JavaScriptEnabled = &v
}

// GetLanguage returns the Language field value
func (o *BrowserInfo) GetLanguage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *BrowserInfo) GetLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value
func (o *BrowserInfo) SetLanguage(v string) {
	o.Language = v
}

// GetScreenHeight returns the ScreenHeight field value
func (o *BrowserInfo) GetScreenHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ScreenHeight
}

// GetScreenHeightOk returns a tuple with the ScreenHeight field value
// and a boolean to check if the value has been set.
func (o *BrowserInfo) GetScreenHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScreenHeight, true
}

// SetScreenHeight sets field value
func (o *BrowserInfo) SetScreenHeight(v int32) {
	o.ScreenHeight = v
}

// GetScreenWidth returns the ScreenWidth field value
func (o *BrowserInfo) GetScreenWidth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ScreenWidth
}

// GetScreenWidthOk returns a tuple with the ScreenWidth field value
// and a boolean to check if the value has been set.
func (o *BrowserInfo) GetScreenWidthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScreenWidth, true
}

// SetScreenWidth sets field value
func (o *BrowserInfo) SetScreenWidth(v int32) {
	o.ScreenWidth = v
}

// GetTimeZoneOffset returns the TimeZoneOffset field value
func (o *BrowserInfo) GetTimeZoneOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TimeZoneOffset
}

// GetTimeZoneOffsetOk returns a tuple with the TimeZoneOffset field value
// and a boolean to check if the value has been set.
func (o *BrowserInfo) GetTimeZoneOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeZoneOffset, true
}

// SetTimeZoneOffset sets field value
func (o *BrowserInfo) SetTimeZoneOffset(v int32) {
	o.TimeZoneOffset = v
}

// GetUserAgent returns the UserAgent field value
func (o *BrowserInfo) GetUserAgent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserAgent
}

// GetUserAgentOk returns a tuple with the UserAgent field value
// and a boolean to check if the value has been set.
func (o *BrowserInfo) GetUserAgentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserAgent, true
}

// SetUserAgent sets field value
func (o *BrowserInfo) SetUserAgent(v string) {
	o.UserAgent = v
}

func (o BrowserInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BrowserInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["acceptHeader"] = o.AcceptHeader
	toSerialize["colorDepth"] = o.ColorDepth
	toSerialize["javaEnabled"] = o.JavaEnabled
	if !common.IsNil(o.JavaScriptEnabled) {
		toSerialize["javaScriptEnabled"] = o.JavaScriptEnabled
	}
	toSerialize["language"] = o.Language
	toSerialize["screenHeight"] = o.ScreenHeight
	toSerialize["screenWidth"] = o.ScreenWidth
	toSerialize["timeZoneOffset"] = o.TimeZoneOffset
	toSerialize["userAgent"] = o.UserAgent
	return toSerialize, nil
}

type NullableBrowserInfo struct {
	value *BrowserInfo
	isSet bool
}

func (v NullableBrowserInfo) Get() *BrowserInfo {
	return v.value
}

func (v *NullableBrowserInfo) Set(val *BrowserInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBrowserInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBrowserInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrowserInfo(val *BrowserInfo) *NullableBrowserInfo {
	return &NullableBrowserInfo{value: val, isSet: true}
}

func (v NullableBrowserInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrowserInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
