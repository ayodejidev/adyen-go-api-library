/*
 * Adyen Payout API
 *
 * A set of API endpoints that allow you to store payout details, confirm, or decline a payout.  For more information, refer to [Online payouts](https://docs.adyen.com/checkout/online-payouts).
 *
 * API version: 52
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package payout

// SubmitResponse struct for SubmitResponse
type SubmitResponse struct {
	// This field contains additional data, which may be returned in a particular response.
	AdditionalData *map[string]interface{} `json:"additionalData,omitempty"`
	// A new reference to uniquely identify this request.
	PspReference string `json:"pspReference"`
	// In case of refusal, an informational message for the reason.
	RefusalReason string `json:"refusalReason,omitempty"`
	// The response: * In case of success, it is `payout-submit-received`. * In case of an error, an informational message is returned.
	ResultCode string `json:"resultCode"`
}
