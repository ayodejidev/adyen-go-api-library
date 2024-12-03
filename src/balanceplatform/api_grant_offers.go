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

	"github.com/adyen/adyen-go-api-library/v15/src/common"
)

// GrantOffersApi service
type GrantOffersApi common.Service

// All parameters accepted by GrantOffersApi.GetAllAvailableGrantOffers
type GrantOffersApiGetAllAvailableGrantOffersInput struct {
	accountHolderId *string
}

// The unique identifier of the grant account.
func (r GrantOffersApiGetAllAvailableGrantOffersInput) AccountHolderId(accountHolderId string) GrantOffersApiGetAllAvailableGrantOffersInput {
	r.accountHolderId = &accountHolderId
	return r
}

/*
Prepare a request for GetAllAvailableGrantOffers

@return GrantOffersApiGetAllAvailableGrantOffersInput
*/
func (a *GrantOffersApi) GetAllAvailableGrantOffersInput() GrantOffersApiGetAllAvailableGrantOffersInput {
	return GrantOffersApiGetAllAvailableGrantOffersInput{}
}

/*
GetAllAvailableGrantOffers Get all available grant offers

Returns a list of all [grant offers](https://docs.adyen.com/platforms/capital#grant-offers) available for `accountHolderId` specified as a query parameter.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r GrantOffersApiGetAllAvailableGrantOffersInput - Request parameters, see GetAllAvailableGrantOffersInput
@return GrantOffers, *http.Response, error
*/
func (a *GrantOffersApi) GetAllAvailableGrantOffers(ctx context.Context, r GrantOffersApiGetAllAvailableGrantOffersInput) (GrantOffers, *http.Response, error) {
	res := &GrantOffers{}
	path := "/grantOffers"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.accountHolderId != nil {
		common.ParameterAddToQuery(queryParams, "accountHolderId", r.accountHolderId, "")
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

// All parameters accepted by GrantOffersApi.GetGrantOffer
type GrantOffersApiGetGrantOfferInput struct {
	grantOfferId string
}

/*
Prepare a request for GetGrantOffer
@param grantOfferId The unique identifier of the grant offer.
@return GrantOffersApiGetGrantOfferInput
*/
func (a *GrantOffersApi) GetGrantOfferInput(grantOfferId string) GrantOffersApiGetGrantOfferInput {
	return GrantOffersApiGetGrantOfferInput{
		grantOfferId: grantOfferId,
	}
}

/*
GetGrantOffer Get a grant offer

Returns the details of a single grant offer.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r GrantOffersApiGetGrantOfferInput - Request parameters, see GetGrantOfferInput
@return GrantOffer, *http.Response, error
*/
func (a *GrantOffersApi) GetGrantOffer(ctx context.Context, r GrantOffersApiGetGrantOfferInput) (GrantOffer, *http.Response, error) {
	res := &GrantOffer{}
	path := "/grantOffers/{grantOfferId}"
	path = strings.Replace(path, "{"+"grantOfferId"+"}", url.PathEscape(common.ParameterValueToString(r.grantOfferId, "grantOfferId")), -1)
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
