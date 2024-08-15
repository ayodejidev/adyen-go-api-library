/*
Adyen Data Protection API

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dataprotection

import (
	"context"
	"net/http"
	"net/url"

	"github.com/adyen/adyen-go-api-library/v12/src/common"
)

// GeneralApi service
type GeneralApi common.Service

// All parameters accepted by GeneralApi.RequestSubjectErasure
type GeneralApiRequestSubjectErasureInput struct {
	subjectErasureByPspReferenceRequest *SubjectErasureByPspReferenceRequest
}

func (r GeneralApiRequestSubjectErasureInput) SubjectErasureByPspReferenceRequest(subjectErasureByPspReferenceRequest SubjectErasureByPspReferenceRequest) GeneralApiRequestSubjectErasureInput {
	r.subjectErasureByPspReferenceRequest = &subjectErasureByPspReferenceRequest
	return r
}

/*
Prepare a request for RequestSubjectErasure

@return GeneralApiRequestSubjectErasureInput
*/
func (a *GeneralApi) RequestSubjectErasureInput() GeneralApiRequestSubjectErasureInput {
	return GeneralApiRequestSubjectErasureInput{}
}

/*
RequestSubjectErasure Submit a Subject Erasure Request.

Sends the PSP reference containing the shopper data that should be deleted.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r GeneralApiRequestSubjectErasureInput - Request parameters, see RequestSubjectErasureInput
@return SubjectErasureResponse, *http.Response, error
*/
func (a *GeneralApi) RequestSubjectErasure(ctx context.Context, r GeneralApiRequestSubjectErasureInput) (SubjectErasureResponse, *http.Response, error) {
	res := &SubjectErasureResponse{}
	path := "/requestSubjectErasure"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.subjectErasureByPspReferenceRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}
