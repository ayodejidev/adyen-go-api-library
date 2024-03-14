/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	"github.com/adyen/adyen-go-api-library/v9/src/common"
)

// HostedOnboardingApi service
type HostedOnboardingApi common.Service

// All parameters accepted by HostedOnboardingApi.GetLinkToAdyenhostedOnboardingPage
type HostedOnboardingApiGetLinkToAdyenhostedOnboardingPageInput struct {
	id                 string
	onboardingLinkInfo *OnboardingLinkInfo
}

func (r HostedOnboardingApiGetLinkToAdyenhostedOnboardingPageInput) OnboardingLinkInfo(onboardingLinkInfo OnboardingLinkInfo) HostedOnboardingApiGetLinkToAdyenhostedOnboardingPageInput {
	r.onboardingLinkInfo = &onboardingLinkInfo
	return r
}

/*
Prepare a request for GetLinkToAdyenhostedOnboardingPage
@param id The unique identifier of the legal entity
@return HostedOnboardingApiGetLinkToAdyenhostedOnboardingPageInput
*/
func (a *HostedOnboardingApi) GetLinkToAdyenhostedOnboardingPageInput(id string) HostedOnboardingApiGetLinkToAdyenhostedOnboardingPageInput {
	return HostedOnboardingApiGetLinkToAdyenhostedOnboardingPageInput{
		id: id,
	}
}

/*
GetLinkToAdyenhostedOnboardingPage Get a link to an Adyen-hosted onboarding page

Returns a link to an Adyen-hosted onboarding page where you need to redirect your user.



@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r HostedOnboardingApiGetLinkToAdyenhostedOnboardingPageInput - Request parameters, see GetLinkToAdyenhostedOnboardingPageInput
@return OnboardingLink, *http.Response, error
*/
func (a *HostedOnboardingApi) GetLinkToAdyenhostedOnboardingPage(ctx context.Context, r HostedOnboardingApiGetLinkToAdyenhostedOnboardingPageInput) (OnboardingLink, *http.Response, error) {
	res := &OnboardingLink{}
	path := "/legalEntities/{id}/onboardingLinks"
	path = strings.Replace(path, "{"+"id"+"}", url.PathEscape(common.ParameterValueToString(r.id, "id")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.onboardingLinkInfo,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by HostedOnboardingApi.GetOnboardingLinkTheme
type HostedOnboardingApiGetOnboardingLinkThemeInput struct {
	id string
}

/*
Prepare a request for GetOnboardingLinkTheme
@param id The unique identifier of the theme
@return HostedOnboardingApiGetOnboardingLinkThemeInput
*/
func (a *HostedOnboardingApi) GetOnboardingLinkThemeInput(id string) HostedOnboardingApiGetOnboardingLinkThemeInput {
	return HostedOnboardingApiGetOnboardingLinkThemeInput{
		id: id,
	}
}

/*
GetOnboardingLinkTheme Get an onboarding link theme

Returns the details of the theme identified in the path.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r HostedOnboardingApiGetOnboardingLinkThemeInput - Request parameters, see GetOnboardingLinkThemeInput
@return OnboardingTheme, *http.Response, error
*/
func (a *HostedOnboardingApi) GetOnboardingLinkTheme(ctx context.Context, r HostedOnboardingApiGetOnboardingLinkThemeInput) (OnboardingTheme, *http.Response, error) {
	res := &OnboardingTheme{}
	path := "/themes/{id}"
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

	return *res, httpRes, err
}

// All parameters accepted by HostedOnboardingApi.ListHostedOnboardingPageThemes
type HostedOnboardingApiListHostedOnboardingPageThemesInput struct {
}

/*
Prepare a request for ListHostedOnboardingPageThemes

@return HostedOnboardingApiListHostedOnboardingPageThemesInput
*/
func (a *HostedOnboardingApi) ListHostedOnboardingPageThemesInput() HostedOnboardingApiListHostedOnboardingPageThemesInput {
	return HostedOnboardingApiListHostedOnboardingPageThemesInput{}
}

/*
ListHostedOnboardingPageThemes Get a list of hosted onboarding page themes

Returns a list of hosted onboarding page themes.



@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r HostedOnboardingApiListHostedOnboardingPageThemesInput - Request parameters, see ListHostedOnboardingPageThemesInput
@return OnboardingThemes, *http.Response, error
*/
func (a *HostedOnboardingApi) ListHostedOnboardingPageThemes(ctx context.Context, r HostedOnboardingApiListHostedOnboardingPageThemesInput) (OnboardingThemes, *http.Response, error) {
	res := &OnboardingThemes{}
	path := "/themes"
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

	return *res, httpRes, err
}
