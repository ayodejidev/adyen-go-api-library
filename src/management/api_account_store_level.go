/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/adyen/adyen-go-api-library/v14/src/common"
)

// AccountStoreLevelApi service
type AccountStoreLevelApi common.Service

// All parameters accepted by AccountStoreLevelApi.CreateStore
type AccountStoreLevelApiCreateStoreInput struct {
	storeCreationWithMerchantCodeRequest *StoreCreationWithMerchantCodeRequest
}

func (r AccountStoreLevelApiCreateStoreInput) StoreCreationWithMerchantCodeRequest(storeCreationWithMerchantCodeRequest StoreCreationWithMerchantCodeRequest) AccountStoreLevelApiCreateStoreInput {
	r.storeCreationWithMerchantCodeRequest = &storeCreationWithMerchantCodeRequest
	return r
}

/*
Prepare a request for CreateStore

@return AccountStoreLevelApiCreateStoreInput
*/
func (a *AccountStoreLevelApi) CreateStoreInput() AccountStoreLevelApiCreateStoreInput {
	return AccountStoreLevelApiCreateStoreInput{}
}

/*
CreateStore Create a store

Creates a store for the merchant account specified in the request.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Stores read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r AccountStoreLevelApiCreateStoreInput - Request parameters, see CreateStoreInput
@return Store, *http.Response, error
*/
func (a *AccountStoreLevelApi) CreateStore(ctx context.Context, r AccountStoreLevelApiCreateStoreInput) (Store, *http.Response, error) {
	res := &Store{}
	path := "/stores"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.storeCreationWithMerchantCodeRequest,
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
	if httpRes.StatusCode == 400 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
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

// All parameters accepted by AccountStoreLevelApi.CreateStoreByMerchantId
type AccountStoreLevelApiCreateStoreByMerchantIdInput struct {
	merchantId           string
	storeCreationRequest *StoreCreationRequest
}

func (r AccountStoreLevelApiCreateStoreByMerchantIdInput) StoreCreationRequest(storeCreationRequest StoreCreationRequest) AccountStoreLevelApiCreateStoreByMerchantIdInput {
	r.storeCreationRequest = &storeCreationRequest
	return r
}

/*
Prepare a request for CreateStoreByMerchantId
@param merchantId The unique identifier of the merchant account.
@return AccountStoreLevelApiCreateStoreByMerchantIdInput
*/
func (a *AccountStoreLevelApi) CreateStoreByMerchantIdInput(merchantId string) AccountStoreLevelApiCreateStoreByMerchantIdInput {
	return AccountStoreLevelApiCreateStoreByMerchantIdInput{
		merchantId: merchantId,
	}
}

/*
CreateStoreByMerchantId Create a store

Creates a store for the merchant account identified in the path.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Stores read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r AccountStoreLevelApiCreateStoreByMerchantIdInput - Request parameters, see CreateStoreByMerchantIdInput
@return Store, *http.Response, error
*/
func (a *AccountStoreLevelApi) CreateStoreByMerchantId(ctx context.Context, r AccountStoreLevelApiCreateStoreByMerchantIdInput) (Store, *http.Response, error) {
	res := &Store{}
	path := "/merchants/{merchantId}/stores"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.storeCreationRequest,
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
	if httpRes.StatusCode == 400 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
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

// All parameters accepted by AccountStoreLevelApi.GetStore
type AccountStoreLevelApiGetStoreInput struct {
	merchantId string
	storeId    string
}

/*
Prepare a request for GetStore
@param merchantId The unique identifier of the merchant account.@param storeId The unique identifier of the store.
@return AccountStoreLevelApiGetStoreInput
*/
func (a *AccountStoreLevelApi) GetStoreInput(merchantId string, storeId string) AccountStoreLevelApiGetStoreInput {
	return AccountStoreLevelApiGetStoreInput{
		merchantId: merchantId,
		storeId:    storeId,
	}
}

/*
GetStore Get a store

Returns the details of the store identified in the path.

To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Stores read
* Management API—Stores read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r AccountStoreLevelApiGetStoreInput - Request parameters, see GetStoreInput
@return Store, *http.Response, error
*/
func (a *AccountStoreLevelApi) GetStore(ctx context.Context, r AccountStoreLevelApiGetStoreInput) (Store, *http.Response, error) {
	res := &Store{}
	path := "/merchants/{merchantId}/stores/{storeId}"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	path = strings.Replace(path, "{"+"storeId"+"}", url.PathEscape(common.ParameterValueToString(r.storeId, "storeId")), -1)
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
	if httpRes.StatusCode == 400 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
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

// All parameters accepted by AccountStoreLevelApi.GetStoreById
type AccountStoreLevelApiGetStoreByIdInput struct {
	storeId string
}

/*
Prepare a request for GetStoreById
@param storeId The unique identifier of the store.
@return AccountStoreLevelApiGetStoreByIdInput
*/
func (a *AccountStoreLevelApi) GetStoreByIdInput(storeId string) AccountStoreLevelApiGetStoreByIdInput {
	return AccountStoreLevelApiGetStoreByIdInput{
		storeId: storeId,
	}
}

/*
GetStoreById Get a store

Returns the details of the store identified in the path.

To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Stores read
* Management API—Stores read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r AccountStoreLevelApiGetStoreByIdInput - Request parameters, see GetStoreByIdInput
@return Store, *http.Response, error
*/
func (a *AccountStoreLevelApi) GetStoreById(ctx context.Context, r AccountStoreLevelApiGetStoreByIdInput) (Store, *http.Response, error) {
	res := &Store{}
	path := "/stores/{storeId}"
	path = strings.Replace(path, "{"+"storeId"+"}", url.PathEscape(common.ParameterValueToString(r.storeId, "storeId")), -1)
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
	if httpRes.StatusCode == 400 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
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

// All parameters accepted by AccountStoreLevelApi.ListStores
type AccountStoreLevelApiListStoresInput struct {
	pageNumber *int32
	pageSize   *int32
	reference  *string
	merchantId *string
}

// The number of the page to fetch.
func (r AccountStoreLevelApiListStoresInput) PageNumber(pageNumber int32) AccountStoreLevelApiListStoresInput {
	r.pageNumber = &pageNumber
	return r
}

// The number of items to have on a page, maximum 100. The default is 10 items on a page.
func (r AccountStoreLevelApiListStoresInput) PageSize(pageSize int32) AccountStoreLevelApiListStoresInput {
	r.pageSize = &pageSize
	return r
}

// The reference of the store.
func (r AccountStoreLevelApiListStoresInput) Reference(reference string) AccountStoreLevelApiListStoresInput {
	r.reference = &reference
	return r
}

// The unique identifier of the merchant account.
func (r AccountStoreLevelApiListStoresInput) MerchantId(merchantId string) AccountStoreLevelApiListStoresInput {
	r.merchantId = &merchantId
	return r
}

/*
Prepare a request for ListStores

@return AccountStoreLevelApiListStoresInput
*/
func (a *AccountStoreLevelApi) ListStoresInput() AccountStoreLevelApiListStoresInput {
	return AccountStoreLevelApiListStoresInput{}
}

/*
ListStores Get a list of stores

Returns a list of stores. The list is grouped into pages as defined by the query parameters.

To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Stores read
* Management API—Stores read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r AccountStoreLevelApiListStoresInput - Request parameters, see ListStoresInput
@return ListStoresResponse, *http.Response, error
*/
func (a *AccountStoreLevelApi) ListStores(ctx context.Context, r AccountStoreLevelApiListStoresInput) (ListStoresResponse, *http.Response, error) {
	res := &ListStoresResponse{}
	path := "/stores"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.pageNumber != nil {
		common.ParameterAddToQuery(queryParams, "pageNumber", r.pageNumber, "")
	}
	if r.pageSize != nil {
		common.ParameterAddToQuery(queryParams, "pageSize", r.pageSize, "")
	}
	if r.reference != nil {
		common.ParameterAddToQuery(queryParams, "reference", r.reference, "")
	}
	if r.merchantId != nil {
		common.ParameterAddToQuery(queryParams, "merchantId", r.merchantId, "")
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
	if httpRes.StatusCode == 400 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
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

// All parameters accepted by AccountStoreLevelApi.ListStoresByMerchantId
type AccountStoreLevelApiListStoresByMerchantIdInput struct {
	merchantId string
	pageNumber *int32
	pageSize   *int32
	reference  *string
}

// The number of the page to fetch.
func (r AccountStoreLevelApiListStoresByMerchantIdInput) PageNumber(pageNumber int32) AccountStoreLevelApiListStoresByMerchantIdInput {
	r.pageNumber = &pageNumber
	return r
}

// The number of items to have on a page, maximum 100. The default is 10 items on a page.
func (r AccountStoreLevelApiListStoresByMerchantIdInput) PageSize(pageSize int32) AccountStoreLevelApiListStoresByMerchantIdInput {
	r.pageSize = &pageSize
	return r
}

// The reference of the store.
func (r AccountStoreLevelApiListStoresByMerchantIdInput) Reference(reference string) AccountStoreLevelApiListStoresByMerchantIdInput {
	r.reference = &reference
	return r
}

/*
Prepare a request for ListStoresByMerchantId
@param merchantId The unique identifier of the merchant account.
@return AccountStoreLevelApiListStoresByMerchantIdInput
*/
func (a *AccountStoreLevelApi) ListStoresByMerchantIdInput(merchantId string) AccountStoreLevelApiListStoresByMerchantIdInput {
	return AccountStoreLevelApiListStoresByMerchantIdInput{
		merchantId: merchantId,
	}
}

/*
ListStoresByMerchantId Get a list of stores

Returns a list of stores for the merchant account identified in the path. The list is grouped into pages as defined by the query parameters.

To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Stores read
* Management API—Stores read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r AccountStoreLevelApiListStoresByMerchantIdInput - Request parameters, see ListStoresByMerchantIdInput
@return ListStoresResponse, *http.Response, error
*/
func (a *AccountStoreLevelApi) ListStoresByMerchantId(ctx context.Context, r AccountStoreLevelApiListStoresByMerchantIdInput) (ListStoresResponse, *http.Response, error) {
	res := &ListStoresResponse{}
	path := "/merchants/{merchantId}/stores"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.pageNumber != nil {
		common.ParameterAddToQuery(queryParams, "pageNumber", r.pageNumber, "")
	}
	if r.pageSize != nil {
		common.ParameterAddToQuery(queryParams, "pageSize", r.pageSize, "")
	}
	if r.reference != nil {
		common.ParameterAddToQuery(queryParams, "reference", r.reference, "")
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
	if httpRes.StatusCode == 400 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
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

// All parameters accepted by AccountStoreLevelApi.UpdateStore
type AccountStoreLevelApiUpdateStoreInput struct {
	merchantId         string
	storeId            string
	updateStoreRequest *UpdateStoreRequest
}

func (r AccountStoreLevelApiUpdateStoreInput) UpdateStoreRequest(updateStoreRequest UpdateStoreRequest) AccountStoreLevelApiUpdateStoreInput {
	r.updateStoreRequest = &updateStoreRequest
	return r
}

/*
Prepare a request for UpdateStore
@param merchantId The unique identifier of the merchant account.@param storeId The unique identifier of the store.
@return AccountStoreLevelApiUpdateStoreInput
*/
func (a *AccountStoreLevelApi) UpdateStoreInput(merchantId string, storeId string) AccountStoreLevelApiUpdateStoreInput {
	return AccountStoreLevelApiUpdateStoreInput{
		merchantId: merchantId,
		storeId:    storeId,
	}
}

/*
UpdateStore Update a store

Updates the store identified in the path. You can only update some store parameters.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Stores read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r AccountStoreLevelApiUpdateStoreInput - Request parameters, see UpdateStoreInput
@return Store, *http.Response, error
*/
func (a *AccountStoreLevelApi) UpdateStore(ctx context.Context, r AccountStoreLevelApiUpdateStoreInput) (Store, *http.Response, error) {
	res := &Store{}
	path := "/merchants/{merchantId}/stores/{storeId}"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	path = strings.Replace(path, "{"+"storeId"+"}", url.PathEscape(common.ParameterValueToString(r.storeId, "storeId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.updateStoreRequest,
		res,
		http.MethodPatch,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	if httpRes == nil {
		return *res, httpRes, err
	}

	var serviceError common.RestServiceError
	if httpRes.StatusCode == 400 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
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

// All parameters accepted by AccountStoreLevelApi.UpdateStoreById
type AccountStoreLevelApiUpdateStoreByIdInput struct {
	storeId            string
	updateStoreRequest *UpdateStoreRequest
}

func (r AccountStoreLevelApiUpdateStoreByIdInput) UpdateStoreRequest(updateStoreRequest UpdateStoreRequest) AccountStoreLevelApiUpdateStoreByIdInput {
	r.updateStoreRequest = &updateStoreRequest
	return r
}

/*
Prepare a request for UpdateStoreById
@param storeId The unique identifier of the store.
@return AccountStoreLevelApiUpdateStoreByIdInput
*/
func (a *AccountStoreLevelApi) UpdateStoreByIdInput(storeId string) AccountStoreLevelApiUpdateStoreByIdInput {
	return AccountStoreLevelApiUpdateStoreByIdInput{
		storeId: storeId,
	}
}

/*
UpdateStoreById Update a store

Updates the store identified in the path.
You can only update some store parameters.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Stores read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r AccountStoreLevelApiUpdateStoreByIdInput - Request parameters, see UpdateStoreByIdInput
@return Store, *http.Response, error
*/
func (a *AccountStoreLevelApi) UpdateStoreById(ctx context.Context, r AccountStoreLevelApiUpdateStoreByIdInput) (Store, *http.Response, error) {
	res := &Store{}
	path := "/stores/{storeId}"
	path = strings.Replace(path, "{"+"storeId"+"}", url.PathEscape(common.ParameterValueToString(r.storeId, "storeId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.updateStoreRequest,
		res,
		http.MethodPatch,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	if httpRes == nil {
		return *res, httpRes, err
	}

	var serviceError common.RestServiceError
	if httpRes.StatusCode == 400 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
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
