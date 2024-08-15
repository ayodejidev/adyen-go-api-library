/*
Payments App API

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package paymentsapp

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	"github.com/adyen/adyen-go-api-library/v12/src/common"
)

// PaymentsAppApi service
type PaymentsAppApi common.Service

// All parameters accepted by PaymentsAppApi.GeneratePaymentsAppBoardingTokenForMerchant
type PaymentsAppApiGeneratePaymentsAppBoardingTokenForMerchantInput struct {
	merchantId           string
	boardingTokenRequest *BoardingTokenRequest
}

func (r PaymentsAppApiGeneratePaymentsAppBoardingTokenForMerchantInput) BoardingTokenRequest(boardingTokenRequest BoardingTokenRequest) PaymentsAppApiGeneratePaymentsAppBoardingTokenForMerchantInput {
	r.boardingTokenRequest = &boardingTokenRequest
	return r
}

/*
Prepare a request for GeneratePaymentsAppBoardingTokenForMerchant
@param merchantId The unique identifier of the merchant account.
@return PaymentsAppApiGeneratePaymentsAppBoardingTokenForMerchantInput
*/
func (a *PaymentsAppApi) GeneratePaymentsAppBoardingTokenForMerchantInput(merchantId string) PaymentsAppApiGeneratePaymentsAppBoardingTokenForMerchantInput {
	return PaymentsAppApiGeneratePaymentsAppBoardingTokenForMerchantInput{
		merchantId: merchantId,
	}
}

/*
GeneratePaymentsAppBoardingTokenForMerchant Create a boarding token - merchant level

Creates a boarding token used to authenticate the installation of a Payments App instance on an Android device. The boarding token is created for the `boardingRequestToken` of the Payments App for the merchant account identified in the path.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Adyen Payments App role


@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r PaymentsAppApiGeneratePaymentsAppBoardingTokenForMerchantInput - Request parameters, see GeneratePaymentsAppBoardingTokenForMerchantInput
@return BoardingTokenResponse, *http.Response, error
*/
func (a *PaymentsAppApi) GeneratePaymentsAppBoardingTokenForMerchant(ctx context.Context, r PaymentsAppApiGeneratePaymentsAppBoardingTokenForMerchantInput) (BoardingTokenResponse, *http.Response, error) {
	res := &BoardingTokenResponse{}
	path := "/merchants/{merchantId}/generatePaymentsAppBoardingToken"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.boardingTokenRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by PaymentsAppApi.GeneratePaymentsAppBoardingTokenForStore
type PaymentsAppApiGeneratePaymentsAppBoardingTokenForStoreInput struct {
	merchantId           string
	storeId              string
	boardingTokenRequest *BoardingTokenRequest
}

func (r PaymentsAppApiGeneratePaymentsAppBoardingTokenForStoreInput) BoardingTokenRequest(boardingTokenRequest BoardingTokenRequest) PaymentsAppApiGeneratePaymentsAppBoardingTokenForStoreInput {
	r.boardingTokenRequest = &boardingTokenRequest
	return r
}

/*
Prepare a request for GeneratePaymentsAppBoardingTokenForStore
@param merchantId The unique identifier of the merchant account.@param storeId The unique identifier of the store.
@return PaymentsAppApiGeneratePaymentsAppBoardingTokenForStoreInput
*/
func (a *PaymentsAppApi) GeneratePaymentsAppBoardingTokenForStoreInput(merchantId string, storeId string) PaymentsAppApiGeneratePaymentsAppBoardingTokenForStoreInput {
	return PaymentsAppApiGeneratePaymentsAppBoardingTokenForStoreInput{
		merchantId: merchantId,
		storeId:    storeId,
	}
}

/*
GeneratePaymentsAppBoardingTokenForStore Create a boarding token - store level

Creates a boarding token used to authenticate the installation of a Payments App instance on an Android device. The boarding token is created for the `boardingRequestToken` of the Payments App for the store identified in the path.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Adyen Payments App role


@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r PaymentsAppApiGeneratePaymentsAppBoardingTokenForStoreInput - Request parameters, see GeneratePaymentsAppBoardingTokenForStoreInput
@return BoardingTokenResponse, *http.Response, error
*/
func (a *PaymentsAppApi) GeneratePaymentsAppBoardingTokenForStore(ctx context.Context, r PaymentsAppApiGeneratePaymentsAppBoardingTokenForStoreInput) (BoardingTokenResponse, *http.Response, error) {
	res := &BoardingTokenResponse{}
	path := "/merchants/{merchantId}/stores/{storeId}/generatePaymentsAppBoardingToken"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	path = strings.Replace(path, "{"+"storeId"+"}", url.PathEscape(common.ParameterValueToString(r.storeId, "storeId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.boardingTokenRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by PaymentsAppApi.ListPaymentsAppForMerchant
type PaymentsAppApiListPaymentsAppForMerchantInput struct {
	merchantId string
	statuses   *string
	limit      *int32
	offset     *int64
}

// The status of the Payments App. Comma-separated list of one or more values. If no value is provided, the list returns all statuses.   Possible values:  * **BOARDING**   * **BOARDED**   * **REVOKED**
func (r PaymentsAppApiListPaymentsAppForMerchantInput) Statuses(statuses string) PaymentsAppApiListPaymentsAppForMerchantInput {
	r.statuses = &statuses
	return r
}

// The number of items to return.
func (r PaymentsAppApiListPaymentsAppForMerchantInput) Limit(limit int32) PaymentsAppApiListPaymentsAppForMerchantInput {
	r.limit = &limit
	return r
}

// The number of items to skip.
func (r PaymentsAppApiListPaymentsAppForMerchantInput) Offset(offset int64) PaymentsAppApiListPaymentsAppForMerchantInput {
	r.offset = &offset
	return r
}

/*
Prepare a request for ListPaymentsAppForMerchant
@param merchantId The unique identifier of the merchant account.
@return PaymentsAppApiListPaymentsAppForMerchantInput
*/
func (a *PaymentsAppApi) ListPaymentsAppForMerchantInput(merchantId string) PaymentsAppApiListPaymentsAppForMerchantInput {
	return PaymentsAppApiListPaymentsAppForMerchantInput{
		merchantId: merchantId,
	}
}

/*
ListPaymentsAppForMerchant Get a list of Payments Apps - merchant level

Returns the list of Payments App instances for the merchant account identified in the path.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Adyen Payments App role


@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r PaymentsAppApiListPaymentsAppForMerchantInput - Request parameters, see ListPaymentsAppForMerchantInput
@return PaymentsAppResponse, *http.Response, error
*/
func (a *PaymentsAppApi) ListPaymentsAppForMerchant(ctx context.Context, r PaymentsAppApiListPaymentsAppForMerchantInput) (PaymentsAppResponse, *http.Response, error) {
	res := &PaymentsAppResponse{}
	path := "/merchants/{merchantId}/paymentsApps"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.statuses != nil {
		common.ParameterAddToQuery(queryParams, "statuses", r.statuses, "")
	}
	if r.limit != nil {
		common.ParameterAddToQuery(queryParams, "limit", r.limit, "")
	}
	if r.offset != nil {
		common.ParameterAddToQuery(queryParams, "offset", r.offset, "")
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

	return *res, httpRes, err
}

// All parameters accepted by PaymentsAppApi.ListPaymentsAppForStore
type PaymentsAppApiListPaymentsAppForStoreInput struct {
	merchantId string
	storeId    string
	statuses   *string
	limit      *int32
	offset     *int64
}

// The status of the Payments App. Comma-separated list of one or more values. If no value is provided, the list returns all statuses.   Possible values:  * **BOARDING**   * **BOARDED**   * **REVOKED**
func (r PaymentsAppApiListPaymentsAppForStoreInput) Statuses(statuses string) PaymentsAppApiListPaymentsAppForStoreInput {
	r.statuses = &statuses
	return r
}

// The number of items to return.
func (r PaymentsAppApiListPaymentsAppForStoreInput) Limit(limit int32) PaymentsAppApiListPaymentsAppForStoreInput {
	r.limit = &limit
	return r
}

// The number of items to skip.
func (r PaymentsAppApiListPaymentsAppForStoreInput) Offset(offset int64) PaymentsAppApiListPaymentsAppForStoreInput {
	r.offset = &offset
	return r
}

/*
Prepare a request for ListPaymentsAppForStore
@param merchantId The unique identifier of the merchant account.@param storeId The unique identifier of the store.
@return PaymentsAppApiListPaymentsAppForStoreInput
*/
func (a *PaymentsAppApi) ListPaymentsAppForStoreInput(merchantId string, storeId string) PaymentsAppApiListPaymentsAppForStoreInput {
	return PaymentsAppApiListPaymentsAppForStoreInput{
		merchantId: merchantId,
		storeId:    storeId,
	}
}

/*
ListPaymentsAppForStore Get a list of Payments Apps - store level

Returns the list of Payments App instances for the store identified in the path.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Adyen Payments App role


@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r PaymentsAppApiListPaymentsAppForStoreInput - Request parameters, see ListPaymentsAppForStoreInput
@return PaymentsAppResponse, *http.Response, error
*/
func (a *PaymentsAppApi) ListPaymentsAppForStore(ctx context.Context, r PaymentsAppApiListPaymentsAppForStoreInput) (PaymentsAppResponse, *http.Response, error) {
	res := &PaymentsAppResponse{}
	path := "/merchants/{merchantId}/stores/{storeId}/paymentsApps"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	path = strings.Replace(path, "{"+"storeId"+"}", url.PathEscape(common.ParameterValueToString(r.storeId, "storeId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.statuses != nil {
		common.ParameterAddToQuery(queryParams, "statuses", r.statuses, "")
	}
	if r.limit != nil {
		common.ParameterAddToQuery(queryParams, "limit", r.limit, "")
	}
	if r.offset != nil {
		common.ParameterAddToQuery(queryParams, "offset", r.offset, "")
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

	return *res, httpRes, err
}

// All parameters accepted by PaymentsAppApi.RevokePaymentsApp
type PaymentsAppApiRevokePaymentsAppInput struct {
	merchantId     string
	installationId string
}

/*
Prepare a request for RevokePaymentsApp
@param merchantId The unique identifier of the merchant account.@param installationId The unique identifier of the Payments App instance on a device.
@return PaymentsAppApiRevokePaymentsAppInput
*/
func (a *PaymentsAppApi) RevokePaymentsAppInput(merchantId string, installationId string) PaymentsAppApiRevokePaymentsAppInput {
	return PaymentsAppApiRevokePaymentsAppInput{
		merchantId:     merchantId,
		installationId: installationId,
	}
}

/*
RevokePaymentsApp Revoke Payments App instance authentication

Revokes the authentication of the Payments App instance for the `installationId` and merchant account identified in the path. This call revokes the authentication of the Payments App instance with the `installationId` identified in the path for both merchant accounts and stores.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Adyen Payments App role


@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r PaymentsAppApiRevokePaymentsAppInput - Request parameters, see RevokePaymentsAppInput
@return *http.Response, error
*/
func (a *PaymentsAppApi) RevokePaymentsApp(ctx context.Context, r PaymentsAppApiRevokePaymentsAppInput) (*http.Response, error) {
	var res interface{}
	path := "/merchants/{merchantId}/paymentsApps/{installationId}/revoke"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	path = strings.Replace(path, "{"+"installationId"+"}", url.PathEscape(common.ParameterValueToString(r.installationId, "installationId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		nil,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return httpRes, err
}
