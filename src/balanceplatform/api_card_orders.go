/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/adyen/adyen-go-api-library/v12/src/common"
)

// CardOrdersApi service
type CardOrdersApi common.Service

// All parameters accepted by CardOrdersApi.GetCardOrderItems
type CardOrdersApiGetCardOrderItemsInput struct {
	id     string
	offset *int32
	limit  *int32
}

// Specifies the position of an element in a list of card orders. The response includes a list of card order items that starts at the specified offset.  **Default:** 0, which means that the response contains all the elements in the list of card order items.
func (r CardOrdersApiGetCardOrderItemsInput) Offset(offset int32) CardOrdersApiGetCardOrderItemsInput {
	r.offset = &offset
	return r
}

// The number of card order items returned per page. **Default:** 10.
func (r CardOrdersApiGetCardOrderItemsInput) Limit(limit int32) CardOrdersApiGetCardOrderItemsInput {
	r.limit = &limit
	return r
}

/*
Prepare a request for GetCardOrderItems
@param id The unique identifier of the card order.
@return CardOrdersApiGetCardOrderItemsInput
*/
func (a *CardOrdersApi) GetCardOrderItemsInput(id string) CardOrdersApiGetCardOrderItemsInput {
	return CardOrdersApiGetCardOrderItemsInput{
		id: id,
	}
}

/*
GetCardOrderItems Get card order items

Returns the item list of a specific card order.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r CardOrdersApiGetCardOrderItemsInput - Request parameters, see GetCardOrderItemsInput
@return PaginatedGetCardOrderItemResponse, *http.Response, error
*/
func (a *CardOrdersApi) GetCardOrderItems(ctx context.Context, r CardOrdersApiGetCardOrderItemsInput) (PaginatedGetCardOrderItemResponse, *http.Response, error) {
	res := &PaginatedGetCardOrderItemResponse{}
	path := "/cardorders/{id}/items"
	path = strings.Replace(path, "{"+"id"+"}", url.PathEscape(common.ParameterValueToString(r.id, "id")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.offset != nil {
		common.ParameterAddToQuery(queryParams, "offset", r.offset, "")
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

// All parameters accepted by CardOrdersApi.ListCardOrders
type CardOrdersApiListCardOrdersInput struct {
	id                         *string
	cardManufacturingProfileId *string
	status                     *string
	txVariantCode              *string
	createdSince               *time.Time
	createdUntil               *time.Time
	lockedSince                *time.Time
	lockedUntil                *time.Time
	serviceCenter              *string
	offset                     *int32
	limit                      *int32
}

// The unique identifier of the card order.
func (r CardOrdersApiListCardOrdersInput) Id(id string) CardOrdersApiListCardOrdersInput {
	r.id = &id
	return r
}

// The unique identifier of the card manufacturer profile.
func (r CardOrdersApiListCardOrdersInput) CardManufacturingProfileId(cardManufacturingProfileId string) CardOrdersApiListCardOrdersInput {
	r.cardManufacturingProfileId = &cardManufacturingProfileId
	return r
}

// The status of the card order.
func (r CardOrdersApiListCardOrdersInput) Status(status string) CardOrdersApiListCardOrdersInput {
	r.status = &status
	return r
}

// The unique code of the card manufacturer profile.  Possible values: **mcmaestro**, **mc**, **visa**, **mcdebit**.
func (r CardOrdersApiListCardOrdersInput) TxVariantCode(txVariantCode string) CardOrdersApiListCardOrdersInput {
	r.txVariantCode = &txVariantCode
	return r
}

// Only include card orders that have been created on or after this point in time. The value must be in ISO 8601 format. For example, **2021-05-30T15:07:40Z**.
func (r CardOrdersApiListCardOrdersInput) CreatedSince(createdSince time.Time) CardOrdersApiListCardOrdersInput {
	r.createdSince = &createdSince
	return r
}

// Only include card orders that have been created on or before this point in time. The value must be in ISO 8601 format. For example, **2021-05-30T15:07:40Z**.
func (r CardOrdersApiListCardOrdersInput) CreatedUntil(createdUntil time.Time) CardOrdersApiListCardOrdersInput {
	r.createdUntil = &createdUntil
	return r
}

// Only include card orders that have been locked on or after this point in time. The value must be in ISO 8601 format. For example, **2021-05-30T15:07:40Z**.
func (r CardOrdersApiListCardOrdersInput) LockedSince(lockedSince time.Time) CardOrdersApiListCardOrdersInput {
	r.lockedSince = &lockedSince
	return r
}

// Only include card orders that have been locked on or before this point in time. The value must be in ISO 8601 format. For example, **2021-05-30T15:07:40Z**.
func (r CardOrdersApiListCardOrdersInput) LockedUntil(lockedUntil time.Time) CardOrdersApiListCardOrdersInput {
	r.lockedUntil = &lockedUntil
	return r
}

// The service center at which the card is issued. The value is case-sensitive.
func (r CardOrdersApiListCardOrdersInput) ServiceCenter(serviceCenter string) CardOrdersApiListCardOrdersInput {
	r.serviceCenter = &serviceCenter
	return r
}

// Specifies the position of an element in a list of card orders. The response includes a list of card orders that starts at the specified offset.  **Default:** 0, which means that the response contains all the elements in the list of card orders.
func (r CardOrdersApiListCardOrdersInput) Offset(offset int32) CardOrdersApiListCardOrdersInput {
	r.offset = &offset
	return r
}

// The number of card orders returned per page. **Default:** 10.
func (r CardOrdersApiListCardOrdersInput) Limit(limit int32) CardOrdersApiListCardOrdersInput {
	r.limit = &limit
	return r
}

/*
Prepare a request for ListCardOrders

@return CardOrdersApiListCardOrdersInput
*/
func (a *CardOrdersApi) ListCardOrdersInput() CardOrdersApiListCardOrdersInput {
	return CardOrdersApiListCardOrdersInput{}
}

/*
ListCardOrders Get a list of card orders

Returns a paginated list of card orders.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r CardOrdersApiListCardOrdersInput - Request parameters, see ListCardOrdersInput
@return PaginatedGetCardOrderResponse, *http.Response, error
*/
func (a *CardOrdersApi) ListCardOrders(ctx context.Context, r CardOrdersApiListCardOrdersInput) (PaginatedGetCardOrderResponse, *http.Response, error) {
	res := &PaginatedGetCardOrderResponse{}
	path := "/cardorders"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.id != nil {
		common.ParameterAddToQuery(queryParams, "id", r.id, "")
	}
	if r.cardManufacturingProfileId != nil {
		common.ParameterAddToQuery(queryParams, "cardManufacturingProfileId", r.cardManufacturingProfileId, "")
	}
	if r.status != nil {
		common.ParameterAddToQuery(queryParams, "status", r.status, "")
	}
	if r.txVariantCode != nil {
		common.ParameterAddToQuery(queryParams, "txVariantCode", r.txVariantCode, "")
	}
	if r.createdSince != nil {
		common.ParameterAddToQuery(queryParams, "createdSince", r.createdSince, "")
	}
	if r.createdUntil != nil {
		common.ParameterAddToQuery(queryParams, "createdUntil", r.createdUntil, "")
	}
	if r.lockedSince != nil {
		common.ParameterAddToQuery(queryParams, "lockedSince", r.lockedSince, "")
	}
	if r.lockedUntil != nil {
		common.ParameterAddToQuery(queryParams, "lockedUntil", r.lockedUntil, "")
	}
	if r.serviceCenter != nil {
		common.ParameterAddToQuery(queryParams, "serviceCenter", r.serviceCenter, "")
	}
	if r.offset != nil {
		common.ParameterAddToQuery(queryParams, "offset", r.offset, "")
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
