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

	"github.com/adyen/adyen-go-api-library/v12/src/common"
)

// TerminalOrdersMerchantLevelApi service
type TerminalOrdersMerchantLevelApi common.Service

// All parameters accepted by TerminalOrdersMerchantLevelApi.CancelOrder
type TerminalOrdersMerchantLevelApiCancelOrderInput struct {
	merchantId string
	orderId    string
}

/*
Prepare a request for CancelOrder
@param merchantId The unique identifier of the merchant account.@param orderId The unique identifier of the order.
@return TerminalOrdersMerchantLevelApiCancelOrderInput
*/
func (a *TerminalOrdersMerchantLevelApi) CancelOrderInput(merchantId string, orderId string) TerminalOrdersMerchantLevelApiCancelOrderInput {
	return TerminalOrdersMerchantLevelApiCancelOrderInput{
		merchantId: merchantId,
		orderId:    orderId,
	}
}

/*
CancelOrder Cancel an order

Cancels the terminal products order identified in the path.
Cancelling is only possible while the order has the status **Placed**.
To cancel an order, make a POST call without a request body. The response returns the full order details, but with the status changed to **Cancelled**.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal ordering read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TerminalOrdersMerchantLevelApiCancelOrderInput - Request parameters, see CancelOrderInput
@return TerminalOrder, *http.Response, error
*/
func (a *TerminalOrdersMerchantLevelApi) CancelOrder(ctx context.Context, r TerminalOrdersMerchantLevelApiCancelOrderInput) (TerminalOrder, *http.Response, error) {
	res := &TerminalOrder{}
	path := "/merchants/{merchantId}/terminalOrders/{orderId}/cancel"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	path = strings.Replace(path, "{"+"orderId"+"}", url.PathEscape(common.ParameterValueToString(r.orderId, "orderId")), -1)
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

// All parameters accepted by TerminalOrdersMerchantLevelApi.CreateOrder
type TerminalOrdersMerchantLevelApiCreateOrderInput struct {
	merchantId           string
	terminalOrderRequest *TerminalOrderRequest
}

func (r TerminalOrdersMerchantLevelApiCreateOrderInput) TerminalOrderRequest(terminalOrderRequest TerminalOrderRequest) TerminalOrdersMerchantLevelApiCreateOrderInput {
	r.terminalOrderRequest = &terminalOrderRequest
	return r
}

/*
Prepare a request for CreateOrder
@param merchantId The unique identifier of the merchant account.
@return TerminalOrdersMerchantLevelApiCreateOrderInput
*/
func (a *TerminalOrdersMerchantLevelApi) CreateOrderInput(merchantId string) TerminalOrdersMerchantLevelApiCreateOrderInput {
	return TerminalOrdersMerchantLevelApiCreateOrderInput{
		merchantId: merchantId,
	}
}

/*
CreateOrder Create an order

Creates an order for payment terminal products for the merchant account identified in the path.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal ordering read and write
>Requests to the Management API test endpoint do not create actual orders for test terminals. To order test terminals, you need to [submit a sales order](https://docs.adyen.com/point-of-sale/managing-terminals/order-terminals/#sales-order-steps) in your Customer Area.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TerminalOrdersMerchantLevelApiCreateOrderInput - Request parameters, see CreateOrderInput
@return TerminalOrder, *http.Response, error
*/
func (a *TerminalOrdersMerchantLevelApi) CreateOrder(ctx context.Context, r TerminalOrdersMerchantLevelApiCreateOrderInput) (TerminalOrder, *http.Response, error) {
	res := &TerminalOrder{}
	path := "/merchants/{merchantId}/terminalOrders"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.terminalOrderRequest,
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

// All parameters accepted by TerminalOrdersMerchantLevelApi.CreateShippingLocation
type TerminalOrdersMerchantLevelApiCreateShippingLocationInput struct {
	merchantId       string
	shippingLocation *ShippingLocation
}

func (r TerminalOrdersMerchantLevelApiCreateShippingLocationInput) ShippingLocation(shippingLocation ShippingLocation) TerminalOrdersMerchantLevelApiCreateShippingLocationInput {
	r.shippingLocation = &shippingLocation
	return r
}

/*
Prepare a request for CreateShippingLocation
@param merchantId The unique identifier of the merchant account.
@return TerminalOrdersMerchantLevelApiCreateShippingLocationInput
*/
func (a *TerminalOrdersMerchantLevelApi) CreateShippingLocationInput(merchantId string) TerminalOrdersMerchantLevelApiCreateShippingLocationInput {
	return TerminalOrdersMerchantLevelApiCreateShippingLocationInput{
		merchantId: merchantId,
	}
}

/*
CreateShippingLocation Create a shipping location

Creates a shipping location for the merchant account identified in the path. A shipping location defines an address where orders can be shipped to.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal ordering read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TerminalOrdersMerchantLevelApiCreateShippingLocationInput - Request parameters, see CreateShippingLocationInput
@return ShippingLocation, *http.Response, error
*/
func (a *TerminalOrdersMerchantLevelApi) CreateShippingLocation(ctx context.Context, r TerminalOrdersMerchantLevelApiCreateShippingLocationInput) (ShippingLocation, *http.Response, error) {
	res := &ShippingLocation{}
	path := "/merchants/{merchantId}/shippingLocations"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.shippingLocation,
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

// All parameters accepted by TerminalOrdersMerchantLevelApi.GetOrder
type TerminalOrdersMerchantLevelApiGetOrderInput struct {
	merchantId string
	orderId    string
}

/*
Prepare a request for GetOrder
@param merchantId The unique identifier of the merchant account.@param orderId The unique identifier of the order.
@return TerminalOrdersMerchantLevelApiGetOrderInput
*/
func (a *TerminalOrdersMerchantLevelApi) GetOrderInput(merchantId string, orderId string) TerminalOrdersMerchantLevelApiGetOrderInput {
	return TerminalOrdersMerchantLevelApiGetOrderInput{
		merchantId: merchantId,
		orderId:    orderId,
	}
}

/*
GetOrder Get an order

Returns the details of the terminal products order identified in the path.

To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal ordering read
* Management API—Terminal ordering read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TerminalOrdersMerchantLevelApiGetOrderInput - Request parameters, see GetOrderInput
@return TerminalOrder, *http.Response, error
*/
func (a *TerminalOrdersMerchantLevelApi) GetOrder(ctx context.Context, r TerminalOrdersMerchantLevelApiGetOrderInput) (TerminalOrder, *http.Response, error) {
	res := &TerminalOrder{}
	path := "/merchants/{merchantId}/terminalOrders/{orderId}"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	path = strings.Replace(path, "{"+"orderId"+"}", url.PathEscape(common.ParameterValueToString(r.orderId, "orderId")), -1)
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

// All parameters accepted by TerminalOrdersMerchantLevelApi.ListBillingEntities
type TerminalOrdersMerchantLevelApiListBillingEntitiesInput struct {
	merchantId string
	name       *string
}

// The name of the billing entity.
func (r TerminalOrdersMerchantLevelApiListBillingEntitiesInput) Name(name string) TerminalOrdersMerchantLevelApiListBillingEntitiesInput {
	r.name = &name
	return r
}

/*
Prepare a request for ListBillingEntities
@param merchantId The unique identifier of the merchant account.
@return TerminalOrdersMerchantLevelApiListBillingEntitiesInput
*/
func (a *TerminalOrdersMerchantLevelApi) ListBillingEntitiesInput(merchantId string) TerminalOrdersMerchantLevelApiListBillingEntitiesInput {
	return TerminalOrdersMerchantLevelApiListBillingEntitiesInput{
		merchantId: merchantId,
	}
}

/*
ListBillingEntities Get a list of billing entities

Returns the billing entities of the merchant account identified in the path.
A billing entity is a legal entity where we charge orders to. An order for terminal products must contain the ID of a billing entity.

To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal ordering read
* Management API—Terminal ordering read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TerminalOrdersMerchantLevelApiListBillingEntitiesInput - Request parameters, see ListBillingEntitiesInput
@return BillingEntitiesResponse, *http.Response, error
*/
func (a *TerminalOrdersMerchantLevelApi) ListBillingEntities(ctx context.Context, r TerminalOrdersMerchantLevelApiListBillingEntitiesInput) (BillingEntitiesResponse, *http.Response, error) {
	res := &BillingEntitiesResponse{}
	path := "/merchants/{merchantId}/billingEntities"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.name != nil {
		common.ParameterAddToQuery(queryParams, "name", r.name, "")
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

// All parameters accepted by TerminalOrdersMerchantLevelApi.ListOrders
type TerminalOrdersMerchantLevelApiListOrdersInput struct {
	merchantId             string
	customerOrderReference *string
	status                 *string
	offset                 *int32
	limit                  *int32
}

// Your purchase order number.
func (r TerminalOrdersMerchantLevelApiListOrdersInput) CustomerOrderReference(customerOrderReference string) TerminalOrdersMerchantLevelApiListOrdersInput {
	r.customerOrderReference = &customerOrderReference
	return r
}

// The order status. Possible values (not case-sensitive): Placed, Confirmed, Cancelled, Shipped, Delivered.
func (r TerminalOrdersMerchantLevelApiListOrdersInput) Status(status string) TerminalOrdersMerchantLevelApiListOrdersInput {
	r.status = &status
	return r
}

// The number of orders to skip.
func (r TerminalOrdersMerchantLevelApiListOrdersInput) Offset(offset int32) TerminalOrdersMerchantLevelApiListOrdersInput {
	r.offset = &offset
	return r
}

// The number of orders to return.
func (r TerminalOrdersMerchantLevelApiListOrdersInput) Limit(limit int32) TerminalOrdersMerchantLevelApiListOrdersInput {
	r.limit = &limit
	return r
}

/*
Prepare a request for ListOrders
@param merchantId
@return TerminalOrdersMerchantLevelApiListOrdersInput
*/
func (a *TerminalOrdersMerchantLevelApi) ListOrdersInput(merchantId string) TerminalOrdersMerchantLevelApiListOrdersInput {
	return TerminalOrdersMerchantLevelApiListOrdersInput{
		merchantId: merchantId,
	}
}

/*
ListOrders Get a list of orders

Returns a list of terminal products orders for the merchant account identified in the path.

To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal ordering read
* Management API—Terminal ordering read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TerminalOrdersMerchantLevelApiListOrdersInput - Request parameters, see ListOrdersInput
@return TerminalOrdersResponse, *http.Response, error
*/
func (a *TerminalOrdersMerchantLevelApi) ListOrders(ctx context.Context, r TerminalOrdersMerchantLevelApiListOrdersInput) (TerminalOrdersResponse, *http.Response, error) {
	res := &TerminalOrdersResponse{}
	path := "/merchants/{merchantId}/terminalOrders"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.customerOrderReference != nil {
		common.ParameterAddToQuery(queryParams, "customerOrderReference", r.customerOrderReference, "")
	}
	if r.status != nil {
		common.ParameterAddToQuery(queryParams, "status", r.status, "")
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

// All parameters accepted by TerminalOrdersMerchantLevelApi.ListShippingLocations
type TerminalOrdersMerchantLevelApiListShippingLocationsInput struct {
	merchantId string
	name       *string
	offset     *int32
	limit      *int32
}

// The name of the shipping location.
func (r TerminalOrdersMerchantLevelApiListShippingLocationsInput) Name(name string) TerminalOrdersMerchantLevelApiListShippingLocationsInput {
	r.name = &name
	return r
}

// The number of locations to skip.
func (r TerminalOrdersMerchantLevelApiListShippingLocationsInput) Offset(offset int32) TerminalOrdersMerchantLevelApiListShippingLocationsInput {
	r.offset = &offset
	return r
}

// The number of locations to return.
func (r TerminalOrdersMerchantLevelApiListShippingLocationsInput) Limit(limit int32) TerminalOrdersMerchantLevelApiListShippingLocationsInput {
	r.limit = &limit
	return r
}

/*
Prepare a request for ListShippingLocations
@param merchantId The unique identifier of the merchant account.
@return TerminalOrdersMerchantLevelApiListShippingLocationsInput
*/
func (a *TerminalOrdersMerchantLevelApi) ListShippingLocationsInput(merchantId string) TerminalOrdersMerchantLevelApiListShippingLocationsInput {
	return TerminalOrdersMerchantLevelApiListShippingLocationsInput{
		merchantId: merchantId,
	}
}

/*
ListShippingLocations Get a list of shipping locations

Returns the shipping locations for the merchant account identified in the path.
A shipping location includes the address where orders can be delivered, and an ID which you need to specify when ordering terminal products.

To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal ordering read
* Management API—Terminal ordering read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TerminalOrdersMerchantLevelApiListShippingLocationsInput - Request parameters, see ListShippingLocationsInput
@return ShippingLocationsResponse, *http.Response, error
*/
func (a *TerminalOrdersMerchantLevelApi) ListShippingLocations(ctx context.Context, r TerminalOrdersMerchantLevelApiListShippingLocationsInput) (ShippingLocationsResponse, *http.Response, error) {
	res := &ShippingLocationsResponse{}
	path := "/merchants/{merchantId}/shippingLocations"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.name != nil {
		common.ParameterAddToQuery(queryParams, "name", r.name, "")
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

// All parameters accepted by TerminalOrdersMerchantLevelApi.ListTerminalModels
type TerminalOrdersMerchantLevelApiListTerminalModelsInput struct {
	merchantId string
}

/*
Prepare a request for ListTerminalModels
@param merchantId The unique identifier of the merchant account.
@return TerminalOrdersMerchantLevelApiListTerminalModelsInput
*/
func (a *TerminalOrdersMerchantLevelApi) ListTerminalModelsInput(merchantId string) TerminalOrdersMerchantLevelApiListTerminalModelsInput {
	return TerminalOrdersMerchantLevelApiListTerminalModelsInput{
		merchantId: merchantId,
	}
}

/*
ListTerminalModels Get a list of terminal models

Returns the payment terminal models that merchant account identified in the path has access to. The response includes the terminal model ID, which can be used as a query parameter when getting a list of terminals or a list of products for ordering.

To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal ordering read
* Management API—Terminal ordering read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TerminalOrdersMerchantLevelApiListTerminalModelsInput - Request parameters, see ListTerminalModelsInput
@return TerminalModelsResponse, *http.Response, error
*/
func (a *TerminalOrdersMerchantLevelApi) ListTerminalModels(ctx context.Context, r TerminalOrdersMerchantLevelApiListTerminalModelsInput) (TerminalModelsResponse, *http.Response, error) {
	res := &TerminalModelsResponse{}
	path := "/merchants/{merchantId}/terminalModels"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
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

// All parameters accepted by TerminalOrdersMerchantLevelApi.ListTerminalProducts
type TerminalOrdersMerchantLevelApiListTerminalProductsInput struct {
	merchantId      string
	country         *string
	terminalModelId *string
	offset          *int32
	limit           *int32
}

// The country to return products for, in [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) format. For example, **US**
func (r TerminalOrdersMerchantLevelApiListTerminalProductsInput) Country(country string) TerminalOrdersMerchantLevelApiListTerminalProductsInput {
	r.country = &country
	return r
}

// The terminal model to return products for. Use the ID returned in the [GET &#x60;/terminalModels&#x60;](https://docs.adyen.com/api-explorer/#/ManagementService/latest/get/merchants/{merchantId}/terminalModels) response. For example, **Verifone.M400**
func (r TerminalOrdersMerchantLevelApiListTerminalProductsInput) TerminalModelId(terminalModelId string) TerminalOrdersMerchantLevelApiListTerminalProductsInput {
	r.terminalModelId = &terminalModelId
	return r
}

// The number of products to skip.
func (r TerminalOrdersMerchantLevelApiListTerminalProductsInput) Offset(offset int32) TerminalOrdersMerchantLevelApiListTerminalProductsInput {
	r.offset = &offset
	return r
}

// The number of products to return.
func (r TerminalOrdersMerchantLevelApiListTerminalProductsInput) Limit(limit int32) TerminalOrdersMerchantLevelApiListTerminalProductsInput {
	r.limit = &limit
	return r
}

/*
Prepare a request for ListTerminalProducts
@param merchantId The unique identifier of the merchant account.
@return TerminalOrdersMerchantLevelApiListTerminalProductsInput
*/
func (a *TerminalOrdersMerchantLevelApi) ListTerminalProductsInput(merchantId string) TerminalOrdersMerchantLevelApiListTerminalProductsInput {
	return TerminalOrdersMerchantLevelApiListTerminalProductsInput{
		merchantId: merchantId,
	}
}

/*
ListTerminalProducts Get a list of terminal products

Returns a country-specific list of payment terminal packages and parts that the merchant account identified in the path has access to.

To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal ordering read
* Management API—Terminal ordering read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TerminalOrdersMerchantLevelApiListTerminalProductsInput - Request parameters, see ListTerminalProductsInput
@return TerminalProductsResponse, *http.Response, error
*/
func (a *TerminalOrdersMerchantLevelApi) ListTerminalProducts(ctx context.Context, r TerminalOrdersMerchantLevelApiListTerminalProductsInput) (TerminalProductsResponse, *http.Response, error) {
	res := &TerminalProductsResponse{}
	path := "/merchants/{merchantId}/terminalProducts"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.country != nil {
		common.ParameterAddToQuery(queryParams, "country", r.country, "")
	}
	if r.terminalModelId != nil {
		common.ParameterAddToQuery(queryParams, "terminalModelId", r.terminalModelId, "")
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

// All parameters accepted by TerminalOrdersMerchantLevelApi.UpdateOrder
type TerminalOrdersMerchantLevelApiUpdateOrderInput struct {
	merchantId           string
	orderId              string
	terminalOrderRequest *TerminalOrderRequest
}

func (r TerminalOrdersMerchantLevelApiUpdateOrderInput) TerminalOrderRequest(terminalOrderRequest TerminalOrderRequest) TerminalOrdersMerchantLevelApiUpdateOrderInput {
	r.terminalOrderRequest = &terminalOrderRequest
	return r
}

/*
Prepare a request for UpdateOrder
@param merchantId The unique identifier of the merchant account.@param orderId The unique identifier of the order.
@return TerminalOrdersMerchantLevelApiUpdateOrderInput
*/
func (a *TerminalOrdersMerchantLevelApi) UpdateOrderInput(merchantId string, orderId string) TerminalOrdersMerchantLevelApiUpdateOrderInput {
	return TerminalOrdersMerchantLevelApiUpdateOrderInput{
		merchantId: merchantId,
		orderId:    orderId,
	}
}

/*
UpdateOrder Update an order

Updates the terminal products order identified in the path.
Updating is only possible while the order has the status **Placed**.

The request body only needs to contain what you want to change.
However, to update the products in the `items` array, you must provide the entire array. For example, if the array has three items:
 To remove one item, the array must include the remaining two items. Or to add one item, the array must include all four items.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal ordering read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TerminalOrdersMerchantLevelApiUpdateOrderInput - Request parameters, see UpdateOrderInput
@return TerminalOrder, *http.Response, error
*/
func (a *TerminalOrdersMerchantLevelApi) UpdateOrder(ctx context.Context, r TerminalOrdersMerchantLevelApiUpdateOrderInput) (TerminalOrder, *http.Response, error) {
	res := &TerminalOrder{}
	path := "/merchants/{merchantId}/terminalOrders/{orderId}"
	path = strings.Replace(path, "{"+"merchantId"+"}", url.PathEscape(common.ParameterValueToString(r.merchantId, "merchantId")), -1)
	path = strings.Replace(path, "{"+"orderId"+"}", url.PathEscape(common.ParameterValueToString(r.orderId, "orderId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.terminalOrderRequest,
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
