/*
 * Adyen BinLookup API
 *
 * The BIN Lookup API provides endpoints for retrieving information, such as cost estimates, and 3D Secure supported version based on a given BIN.
 *
 * API version: 50
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package binlookup
// Amount struct for Amount
type Amount struct {
	// The three-character [ISO currency code](https://docs.adyen.com/development-resources/currency-codes).
	Currency string `json:"currency"`
	// The payable amount that can be charged for the transaction.  The transaction amount needs to be represented in minor units according to the [following table](https://docs.adyen.com/development-resources/currency-codes).
	Value int64 `json:"value"`
}
