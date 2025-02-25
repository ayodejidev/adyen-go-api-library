/*
 * Adyen for Platforms: Fund API
 *
 * The Fund API provides endpoints for managing the funds in the accounts on your platform. These management operations include actions such as the transfer of funds from one account to another, the payout of funds to an account holder, and the retrieval of balances in an account.  For more information, refer to our [documentation](https://docs.adyen.com/platforms). ## Authentication To connect to the Fund API, you must use basic authentication credentials of your web service user. If you don't have one, please contact the [Adyen Support Team](https://support.adyen.com/hc/en-us/requests/new). Then use its credentials to authenticate your request, for example:  ``` curl -U \"ws@MarketPlace.YourMarketPlace\":\"YourWsPassword\" \\ -H \"Content-Type: application/json\" \\ ... ``` Note that when going live, you need to generate new web service user credentials to access the [live endpoints](https://docs.adyen.com/development-resources/live-endpoints).  ## Versioning The Fund API supports versioning of its endpoints through a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://cal-test.adyen.com/cal/services/Fund/v6/accountHolderBalance ```
 *
 * API version: 6
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package platformsfund

import (
	_context "context"
	_nethttp "net/http"

	"github.com/adyen/adyen-go-api-library/v19/src/common"
)

// PlatformsFund PlatformsFund service
// Deprecated: Please migrate to the new Adyen For Platforms.
type PlatformsFund common.Service

/*
PostAccountHolderBalance Retrieve the balance(s) of an account holder.
This endpoint is used to retrieve the balance(s) of the accounts of an account holder. An account&#39;s balances are on a per-currency basis (i.e., an account may have multiple balances: one per currency).
  - @param request AccountHolderBalanceRequest - reference of AccountHolderBalanceRequest).
  - @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return AccountHolderBalanceResponse
*/
func (a PlatformsFund) AccountHolderBalance(req *AccountHolderBalanceRequest, ctxs ..._context.Context) (AccountHolderBalanceResponse, *_nethttp.Response, error) {
	res := &AccountHolderBalanceResponse{}
	httpRes, err := a.Client.MakeHTTPPostRequest(req, res, a.BasePath()+"/accountHolderBalance", ctxs...)
	return *res, httpRes, err
}

/*
PostAccountHolderTransactionList Retrieve a list of transactions.
This endpoint is used to retrieve a list of Transactions for an account holder&#39;s accounts. The accounts and Transaction Statuses to be included on the list can be specified. Each call will return a maximum of fifty (50) Transactions per account; in order to retrieve the following set of Transactions another call should be made with the &#39;page&#39; value incremented. Note that Transactions are ordered with most recent first.
  - @param request AccountHolderTransactionListRequest - reference of AccountHolderTransactionListRequest).
  - @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return AccountHolderTransactionListResponse
*/
func (a PlatformsFund) AccountHolderTransactionList(req *AccountHolderTransactionListRequest, ctxs ..._context.Context) (AccountHolderTransactionListResponse, *_nethttp.Response, error) {
	res := &AccountHolderTransactionListResponse{}
	httpRes, err := a.Client.MakeHTTPPostRequest(req, res, a.BasePath()+"/accountHolderTransactionList", ctxs...)
	return *res, httpRes, err
}

/*
PostPayoutAccountHolder Disburse a specified amount from an account to the account holder.
This endpoint is used to pay out a specified amount from an account to the bank account of the account&#39;s account holder.
  - @param request PayoutAccountHolderRequest - reference of PayoutAccountHolderRequest).
  - @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return PayoutAccountHolderResponse
*/
func (a PlatformsFund) PayoutAccountHolder(req *PayoutAccountHolderRequest, ctxs ..._context.Context) (PayoutAccountHolderResponse, *_nethttp.Response, error) {
	res := &PayoutAccountHolderResponse{}
	httpRes, err := a.Client.MakeHTTPPostRequest(req, res, a.BasePath()+"/payoutAccountHolder", ctxs...)
	return *res, httpRes, err
}

/*
PostRefundFundsTransfer Make a refund of the existing transfer funds transfer.
This endpoint is used to refund funds transferred from one account to another. Both accounts must be in the same marketplace, but can have different account holders.
  - @param request RefundFundsTransferRequest - reference of RefundFundsTransferRequest).
  - @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return RefundFundsTransferResponse
*/
func (a PlatformsFund) RefundFundsTransfer(req *RefundFundsTransferRequest, ctxs ..._context.Context) (RefundFundsTransferResponse, *_nethttp.Response, error) {
	res := &RefundFundsTransferResponse{}
	httpRes, err := a.Client.MakeHTTPPostRequest(req, res, a.BasePath()+"/refundFundsTransfer", ctxs...)
	return *res, httpRes, err
}

/*
PostRefundNotPaidOutTransfers Refund all transactions of an account since the most recent payout.
This endpoint is used to refund all the transactions of an account which have taken place since the most recent payout. This request is on a per-account basis (as opposed to a per-payment basis), so only the portion of the payment which was made to the specified account will be refunded. The commission(s), fee(s), and payment(s) to other account(s), will remain in the accounts to which they were sent as designated by the original payment&#39;s split details.
  - @param request RefundNotPaidOutTransfersRequest - reference of RefundNotPaidOutTransfersRequest).
  - @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return RefundNotPaidOutTransfersResponse
*/
func (a PlatformsFund) RefundNotPaidOutTransfers(req *RefundNotPaidOutTransfersRequest, ctxs ..._context.Context) (RefundNotPaidOutTransfersResponse, *_nethttp.Response, error) {
	res := &RefundNotPaidOutTransfersResponse{}
	httpRes, err := a.Client.MakeHTTPPostRequest(req, res, a.BasePath()+"/refundNotPaidOutTransfers", ctxs...)
	return *res, httpRes, err
}

/*
PostSetupBeneficiary Designate an account to be the beneficiary of a separate account and transfer the benefactor's current balance to the beneficiary.
This endpoint is used to define a benefactor and a beneficiary relationship between two accounts. At the time of benefactor/beneficiary setup, the funds in the benefactor account are transferred to the beneficiary account, and any further payments to the benefactor account are automatically sent to the beneficiary account. Note that a series of benefactor/beneficiaries may not exceed four (4) beneficiaries and may not have a cycle in it.
  - @param request SetupBeneficiaryRequest - reference of SetupBeneficiaryRequest).
  - @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return SetupBeneficiaryResponse
*/
func (a PlatformsFund) SetupBeneficiary(req *SetupBeneficiaryRequest, ctxs ..._context.Context) (SetupBeneficiaryResponse, *_nethttp.Response, error) {
	res := &SetupBeneficiaryResponse{}
	httpRes, err := a.Client.MakeHTTPPostRequest(req, res, a.BasePath()+"/setupBeneficiary", ctxs...)
	return *res, httpRes, err
}

/*
PostTransferFunds Transfer funds from one platform account to another.
This endpoint is used to transfer funds from one account to another account. Both accounts must be in the same marketplace, but can have different account holders. The transfer must include a transfer code, which should be determined by the marketplace, in compliance with local regulations.
  - @param request TransferFundsRequest - reference of TransferFundsRequest).
  - @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return TransferFundsResponse
*/
func (a PlatformsFund) TransferFunds(req *TransferFundsRequest, ctxs ..._context.Context) (TransferFundsResponse, *_nethttp.Response, error) {
	res := &TransferFundsResponse{}
	httpRes, err := a.Client.MakeHTTPPostRequest(req, res, a.BasePath()+"/transferFunds", ctxs...)
	return *res, httpRes, err
}
