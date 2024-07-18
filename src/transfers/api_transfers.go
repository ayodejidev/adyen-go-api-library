/*
Transfers API

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transfers

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/adyen/adyen-go-api-library/v11/src/common"
)

// TransfersApi service
type TransfersApi common.Service

// All parameters accepted by TransfersApi.GetAllTransfers
type TransfersApiGetAllTransfersInput struct {
	createdSince        *time.Time
	createdUntil        *time.Time
	balancePlatform     *string
	accountHolderId     *string
	balanceAccountId    *string
	paymentInstrumentId *string
	reference           *string
	category            *string
	cursor              *string
	limit               *int32
}

// Only include transfers that have been created on or after this point in time. The value must be in ISO 8601 format. For example, **2021-05-30T15:07:40Z**.
func (r TransfersApiGetAllTransfersInput) CreatedSince(createdSince time.Time) TransfersApiGetAllTransfersInput {
	r.createdSince = &createdSince
	return r
}

// Only include transfers that have been created on or before this point in time. The value must be in ISO 8601 format. For example, **2021-05-30T15:07:40Z**.
func (r TransfersApiGetAllTransfersInput) CreatedUntil(createdUntil time.Time) TransfersApiGetAllTransfersInput {
	r.createdUntil = &createdUntil
	return r
}

// The unique identifier of the [balance platform](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/get/balancePlatforms/{id}__queryParam_id).  Required if you don&#39;t provide a &#x60;balanceAccountId&#x60; or &#x60;accountHolderId&#x60;.
func (r TransfersApiGetAllTransfersInput) BalancePlatform(balancePlatform string) TransfersApiGetAllTransfersInput {
	r.balancePlatform = &balancePlatform
	return r
}

// The unique identifier of the [account holder](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/get/accountHolders/{id}__queryParam_id).  Required if you don&#39;t provide a &#x60;balanceAccountId&#x60; or &#x60;balancePlatform&#x60;.  If you provide a &#x60;balanceAccountId&#x60;, the &#x60;accountHolderId&#x60; must be related to the &#x60;balanceAccountId&#x60;.
func (r TransfersApiGetAllTransfersInput) AccountHolderId(accountHolderId string) TransfersApiGetAllTransfersInput {
	r.accountHolderId = &accountHolderId
	return r
}

// The unique identifier of the [balance account](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/get/balanceAccounts/{id}__queryParam_id).  Required if you don&#39;t provide an &#x60;accountHolderId&#x60; or &#x60;balancePlatform&#x60;.  If you provide an &#x60;accountHolderId&#x60;, the &#x60;balanceAccountId&#x60; must be related to the &#x60;accountHolderId&#x60;.
func (r TransfersApiGetAllTransfersInput) BalanceAccountId(balanceAccountId string) TransfersApiGetAllTransfersInput {
	r.balanceAccountId = &balanceAccountId
	return r
}

// The unique identifier of the [payment instrument](https://docs.adyen.com/api-explorer/balanceplatform/latest/get/paymentInstruments/_id_).  To use this parameter, you must also provide a &#x60;balanceAccountId&#x60;, &#x60;accountHolderId&#x60;, or &#x60;balancePlatform&#x60;.  The &#x60;paymentInstrumentId&#x60; must be related to the &#x60;balanceAccountId&#x60; or &#x60;accountHolderId&#x60; that you provide.
func (r TransfersApiGetAllTransfersInput) PaymentInstrumentId(paymentInstrumentId string) TransfersApiGetAllTransfersInput {
	r.paymentInstrumentId = &paymentInstrumentId
	return r
}

// The reference you provided in the POST [/transfers](https://docs.adyen.com/api-explorer/transfers/latest/post/transfers) request
func (r TransfersApiGetAllTransfersInput) Reference(reference string) TransfersApiGetAllTransfersInput {
	r.reference = &reference
	return r
}

// The type of transfer.  Possible values:   - **bank**: Transfer to a [transfer instrument](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/transferInstruments__resParam_id) or a bank account.  - **internal**: Transfer to another [balance account](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/balanceAccounts__resParam_id) within your platform.  - **issuedCard**: Transfer initiated by a Adyen-issued card.  - **platformPayment**: Fund movements related to payments that are acquired for your users.
func (r TransfersApiGetAllTransfersInput) Category(category string) TransfersApiGetAllTransfersInput {
	r.category = &category
	return r
}

// The &#x60;cursor&#x60; returned in the links of the previous response.
func (r TransfersApiGetAllTransfersInput) Cursor(cursor string) TransfersApiGetAllTransfersInput {
	r.cursor = &cursor
	return r
}

// The number of items returned per page, maximum of 100 items. By default, the response returns 10 items per page.
func (r TransfersApiGetAllTransfersInput) Limit(limit int32) TransfersApiGetAllTransfersInput {
	r.limit = &limit
	return r
}

/*
Prepare a request for GetAllTransfers

@return TransfersApiGetAllTransfersInput
*/
func (a *TransfersApi) GetAllTransfersInput() TransfersApiGetAllTransfersInput {
	return TransfersApiGetAllTransfersInput{}
}

/*
GetAllTransfers Get all transfers

Returns all the transfers related to a balance account, account holder, or balance platform.

When making this request, you must include at least one of the following:
- `balanceAccountId`
- `accountHolderId`
- `balancePlatform`.

This endpoint supports cursor-based pagination. The response returns the first page of results, and returns links to the next and previous pages when applicable. You can use the links to page through the results.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TransfersApiGetAllTransfersInput - Request parameters, see GetAllTransfersInput
@return FindTransfersResponse, *http.Response, error
*/
func (a *TransfersApi) GetAllTransfers(ctx context.Context, r TransfersApiGetAllTransfersInput) (FindTransfersResponse, *http.Response, error) {
	res := &FindTransfersResponse{}
	path := "/transfers"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.balancePlatform != nil {
		common.ParameterAddToQuery(queryParams, "balancePlatform", r.balancePlatform, "")
	}
	if r.accountHolderId != nil {
		common.ParameterAddToQuery(queryParams, "accountHolderId", r.accountHolderId, "")
	}
	if r.balanceAccountId != nil {
		common.ParameterAddToQuery(queryParams, "balanceAccountId", r.balanceAccountId, "")
	}
	if r.paymentInstrumentId != nil {
		common.ParameterAddToQuery(queryParams, "paymentInstrumentId", r.paymentInstrumentId, "")
	}
	if r.reference != nil {
		common.ParameterAddToQuery(queryParams, "reference", r.reference, "")
	}
	if r.category != nil {
		common.ParameterAddToQuery(queryParams, "category", r.category, "")
	}
	if r.createdSince != nil {
		common.ParameterAddToQuery(queryParams, "createdSince", r.createdSince, "")
	}
	if r.createdUntil != nil {
		common.ParameterAddToQuery(queryParams, "createdUntil", r.createdUntil, "")
	}
	if r.cursor != nil {
		common.ParameterAddToQuery(queryParams, "cursor", r.cursor, "")
	}
	if r.limit != nil {
		common.ParameterAddToQuery(queryParams, "limit", r.limit, "")
	}
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		nil,
		res,
		http.MethodGet,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	if httpRes == nil {
		return *res, httpRes, err
	}

	var serviceError common.RestServiceError
	if httpRes.StatusCode == 401 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 403 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 422 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 500 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	return *res, httpRes, err
}

// All parameters accepted by TransfersApi.GetTransfer
type TransfersApiGetTransferInput struct {
	id string
}

/*
Prepare a request for GetTransfer
@param id Unique identifier of the transfer.
@return TransfersApiGetTransferInput
*/
func (a *TransfersApi) GetTransferInput(id string) TransfersApiGetTransferInput {
	return TransfersApiGetTransferInput{
		id: id,
	}
}

/*
GetTransfer Get a transfer

Returns a transfer.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TransfersApiGetTransferInput - Request parameters, see GetTransferInput
@return TransferData, *http.Response, error
*/
func (a *TransfersApi) GetTransfer(ctx context.Context, r TransfersApiGetTransferInput) (TransferData, *http.Response, error) {
	res := &TransferData{}
	path := "/transfers/{id}"
	path = strings.Replace(path, "{"+"id"+"}", url.PathEscape(common.ParameterValueToString(r.id, "id")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		nil,
		res,
		http.MethodGet,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	if httpRes == nil {
		return *res, httpRes, err
	}

	var serviceError common.RestServiceError
	if httpRes.StatusCode == 401 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 403 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 422 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 500 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	return *res, httpRes, err
}

// All parameters accepted by TransfersApi.ReturnTransfer
type TransfersApiReturnTransferInput struct {
	transferId            string
	returnTransferRequest *ReturnTransferRequest
}

func (r TransfersApiReturnTransferInput) ReturnTransferRequest(returnTransferRequest ReturnTransferRequest) TransfersApiReturnTransferInput {
	r.returnTransferRequest = &returnTransferRequest
	return r
}

/*
Prepare a request for ReturnTransfer
@param transferId The unique identifier of the transfer to be returned.
@return TransfersApiReturnTransferInput
*/
func (a *TransfersApi) ReturnTransferInput(transferId string) TransfersApiReturnTransferInput {
	return TransfersApiReturnTransferInput{
		transferId: transferId,
	}
}

/*
ReturnTransfer Return a transfer

Initiates the return of previously transferred funds without creating a new `transferId`.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TransfersApiReturnTransferInput - Request parameters, see ReturnTransferInput
@return ReturnTransferResponse, *http.Response, error
*/
func (a *TransfersApi) ReturnTransfer(ctx context.Context, r TransfersApiReturnTransferInput) (ReturnTransferResponse, *http.Response, error) {
	res := &ReturnTransferResponse{}
	path := "/transfers/{transferId}/returns"
	path = strings.Replace(path, "{"+"transferId"+"}", url.PathEscape(common.ParameterValueToString(r.transferId, "transferId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.returnTransferRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	if httpRes == nil {
		return *res, httpRes, err
	}

	var serviceError common.RestServiceError
	if httpRes.StatusCode == 401 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 403 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 422 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 500 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	return *res, httpRes, err
}

// All parameters accepted by TransfersApi.TransferFunds
type TransfersApiTransferFundsInput struct {
	wWWAuthenticate *string
	transferInfo    *TransferInfo
}

// Header for authenticating through SCA
func (r TransfersApiTransferFundsInput) WWWAuthenticate(wWWAuthenticate string) TransfersApiTransferFundsInput {
	r.wWWAuthenticate = &wWWAuthenticate
	return r
}

func (r TransfersApiTransferFundsInput) TransferInfo(transferInfo TransferInfo) TransfersApiTransferFundsInput {
	r.transferInfo = &transferInfo
	return r
}

/*
Prepare a request for TransferFunds

@return TransfersApiTransferFundsInput
*/
func (a *TransfersApi) TransferFundsInput() TransfersApiTransferFundsInput {
	return TransfersApiTransferFundsInput{}
}

/*
TransferFunds Transfer funds

>Versions 1 and 2 of the Transfers API are deprecated. If you are just starting your implementation, use the latest version.

Starts a request to transfer funds to [balance accounts](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/balanceAccounts), [transfer instruments](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/transferInstruments), or third-party bank accounts. Adyen sends the outcome of the transfer request through webhooks.

To use this endpoint, you need an additional role for your API credential and transfers must be enabled for the source balance account. Your Adyen contact will set these up for you.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TransfersApiTransferFundsInput - Request parameters, see TransferFundsInput
@return Transfer, *http.Response, error
*/
func (a *TransfersApi) TransferFunds(ctx context.Context, r TransfersApiTransferFundsInput) (Transfer, *http.Response, error) {
	res := &Transfer{}
	path := "/transfers"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.wWWAuthenticate != nil {
		common.ParameterAddToHeaderOrQuery(headerParams, "WWW-Authenticate", r.wWWAuthenticate, "")
	}
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.transferInfo,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	if httpRes == nil {
		return *res, httpRes, err
	}

	var serviceError common.RestServiceError
	if httpRes.StatusCode == 401 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 403 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 422 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 500 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	return *res, httpRes, err
}
