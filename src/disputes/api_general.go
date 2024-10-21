/*
Disputes API

API version: 30
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package disputes

import (
	"context"
	"net/http"
	"net/url"

	"github.com/adyen/adyen-go-api-library/v14/src/common"
)

// GeneralApi service
type GeneralApi common.Service

// All parameters accepted by GeneralApi.AcceptDispute
type GeneralApiAcceptDisputeInput struct {
	acceptDisputeRequest *AcceptDisputeRequest
}

func (r GeneralApiAcceptDisputeInput) AcceptDisputeRequest(acceptDisputeRequest AcceptDisputeRequest) GeneralApiAcceptDisputeInput {
	r.acceptDisputeRequest = &acceptDisputeRequest
	return r
}

/*
Prepare a request for AcceptDispute

@return GeneralApiAcceptDisputeInput
*/
func (a *GeneralApi) AcceptDisputeInput() GeneralApiAcceptDisputeInput {
	return GeneralApiAcceptDisputeInput{}
}

/*
AcceptDispute Accept a dispute

Accepts a specific dispute.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r GeneralApiAcceptDisputeInput - Request parameters, see AcceptDisputeInput
@return AcceptDisputeResponse, *http.Response, error
*/
func (a *GeneralApi) AcceptDispute(ctx context.Context, r GeneralApiAcceptDisputeInput) (AcceptDisputeResponse, *http.Response, error) {
	res := &AcceptDisputeResponse{}
	path := "/acceptDispute"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.acceptDisputeRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by GeneralApi.DefendDispute
type GeneralApiDefendDisputeInput struct {
	defendDisputeRequest *DefendDisputeRequest
}

func (r GeneralApiDefendDisputeInput) DefendDisputeRequest(defendDisputeRequest DefendDisputeRequest) GeneralApiDefendDisputeInput {
	r.defendDisputeRequest = &defendDisputeRequest
	return r
}

/*
Prepare a request for DefendDispute

@return GeneralApiDefendDisputeInput
*/
func (a *GeneralApi) DefendDisputeInput() GeneralApiDefendDisputeInput {
	return GeneralApiDefendDisputeInput{}
}

/*
DefendDispute Defend a dispute

Defends a specific dispute.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r GeneralApiDefendDisputeInput - Request parameters, see DefendDisputeInput
@return DefendDisputeResponse, *http.Response, error
*/
func (a *GeneralApi) DefendDispute(ctx context.Context, r GeneralApiDefendDisputeInput) (DefendDisputeResponse, *http.Response, error) {
	res := &DefendDisputeResponse{}
	path := "/defendDispute"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.defendDisputeRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by GeneralApi.DeleteDisputeDefenseDocument
type GeneralApiDeleteDisputeDefenseDocumentInput struct {
	deleteDefenseDocumentRequest *DeleteDefenseDocumentRequest
}

func (r GeneralApiDeleteDisputeDefenseDocumentInput) DeleteDefenseDocumentRequest(deleteDefenseDocumentRequest DeleteDefenseDocumentRequest) GeneralApiDeleteDisputeDefenseDocumentInput {
	r.deleteDefenseDocumentRequest = &deleteDefenseDocumentRequest
	return r
}

/*
Prepare a request for DeleteDisputeDefenseDocument

@return GeneralApiDeleteDisputeDefenseDocumentInput
*/
func (a *GeneralApi) DeleteDisputeDefenseDocumentInput() GeneralApiDeleteDisputeDefenseDocumentInput {
	return GeneralApiDeleteDisputeDefenseDocumentInput{}
}

/*
DeleteDisputeDefenseDocument Delete a defense document

Deletes a specific dispute defense document that was supplied earlier.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r GeneralApiDeleteDisputeDefenseDocumentInput - Request parameters, see DeleteDisputeDefenseDocumentInput
@return DeleteDefenseDocumentResponse, *http.Response, error
*/
func (a *GeneralApi) DeleteDisputeDefenseDocument(ctx context.Context, r GeneralApiDeleteDisputeDefenseDocumentInput) (DeleteDefenseDocumentResponse, *http.Response, error) {
	res := &DeleteDefenseDocumentResponse{}
	path := "/deleteDisputeDefenseDocument"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.deleteDefenseDocumentRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by GeneralApi.RetrieveApplicableDefenseReasons
type GeneralApiRetrieveApplicableDefenseReasonsInput struct {
	defenseReasonsRequest *DefenseReasonsRequest
}

func (r GeneralApiRetrieveApplicableDefenseReasonsInput) DefenseReasonsRequest(defenseReasonsRequest DefenseReasonsRequest) GeneralApiRetrieveApplicableDefenseReasonsInput {
	r.defenseReasonsRequest = &defenseReasonsRequest
	return r
}

/*
Prepare a request for RetrieveApplicableDefenseReasons

@return GeneralApiRetrieveApplicableDefenseReasonsInput
*/
func (a *GeneralApi) RetrieveApplicableDefenseReasonsInput() GeneralApiRetrieveApplicableDefenseReasonsInput {
	return GeneralApiRetrieveApplicableDefenseReasonsInput{}
}

/*
RetrieveApplicableDefenseReasons Get applicable defense reasons

Returns a list of all applicable defense reasons to defend a specific dispute.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r GeneralApiRetrieveApplicableDefenseReasonsInput - Request parameters, see RetrieveApplicableDefenseReasonsInput
@return DefenseReasonsResponse, *http.Response, error
*/
func (a *GeneralApi) RetrieveApplicableDefenseReasons(ctx context.Context, r GeneralApiRetrieveApplicableDefenseReasonsInput) (DefenseReasonsResponse, *http.Response, error) {
	res := &DefenseReasonsResponse{}
	path := "/retrieveApplicableDefenseReasons"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.defenseReasonsRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by GeneralApi.SupplyDefenseDocument
type GeneralApiSupplyDefenseDocumentInput struct {
	supplyDefenseDocumentRequest *SupplyDefenseDocumentRequest
}

func (r GeneralApiSupplyDefenseDocumentInput) SupplyDefenseDocumentRequest(supplyDefenseDocumentRequest SupplyDefenseDocumentRequest) GeneralApiSupplyDefenseDocumentInput {
	r.supplyDefenseDocumentRequest = &supplyDefenseDocumentRequest
	return r
}

/*
Prepare a request for SupplyDefenseDocument

@return GeneralApiSupplyDefenseDocumentInput
*/
func (a *GeneralApi) SupplyDefenseDocumentInput() GeneralApiSupplyDefenseDocumentInput {
	return GeneralApiSupplyDefenseDocumentInput{}
}

/*
SupplyDefenseDocument Supply a defense document

Supplies a specific dispute defense document.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r GeneralApiSupplyDefenseDocumentInput - Request parameters, see SupplyDefenseDocumentInput
@return SupplyDefenseDocumentResponse, *http.Response, error
*/
func (a *GeneralApi) SupplyDefenseDocument(ctx context.Context, r GeneralApiSupplyDefenseDocumentInput) (SupplyDefenseDocumentResponse, *http.Response, error) {
	res := &SupplyDefenseDocumentResponse{}
	path := "/supplyDefenseDocument"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.supplyDefenseDocumentRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}
