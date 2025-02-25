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

	"github.com/adyen/adyen-go-api-library/v19/src/common"
)

// TerminalSettingsTerminalLevelApi service
type TerminalSettingsTerminalLevelApi common.Service

// All parameters accepted by TerminalSettingsTerminalLevelApi.GetTerminalLogo
type TerminalSettingsTerminalLevelApiGetTerminalLogoInput struct {
	terminalId string
}

/*
Prepare a request for GetTerminalLogo
@param terminalId The unique identifier of the payment terminal.
@return TerminalSettingsTerminalLevelApiGetTerminalLogoInput
*/
func (a *TerminalSettingsTerminalLevelApi) GetTerminalLogoInput(terminalId string) TerminalSettingsTerminalLevelApiGetTerminalLogoInput {
	return TerminalSettingsTerminalLevelApiGetTerminalLogoInput{
		terminalId: terminalId,
	}
}

/*
GetTerminalLogo Get the terminal logo

Returns the logo that is configured for the payment terminal identified in the path.
The logo is returned as a Base64-encoded string. You need to Base64-decode the string to get the actual image file.

To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal settings read
* Management API—Terminal settings read and write

In the live environment, requests to this endpoint are subject to [rate limits](https://docs.adyen.com/point-of-sale/automating-terminal-management#rate-limits-in-the-live-environment).

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TerminalSettingsTerminalLevelApiGetTerminalLogoInput - Request parameters, see GetTerminalLogoInput
@return Logo, *http.Response, error
*/
func (a *TerminalSettingsTerminalLevelApi) GetTerminalLogo(ctx context.Context, r TerminalSettingsTerminalLevelApiGetTerminalLogoInput) (Logo, *http.Response, error) {
	res := &Logo{}
	path := "/terminals/{terminalId}/terminalLogos"
	path = strings.Replace(path, "{"+"terminalId"+"}", url.PathEscape(common.ParameterValueToString(r.terminalId, "terminalId")), -1)
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

// All parameters accepted by TerminalSettingsTerminalLevelApi.GetTerminalSettings
type TerminalSettingsTerminalLevelApiGetTerminalSettingsInput struct {
	terminalId string
}

/*
Prepare a request for GetTerminalSettings
@param terminalId The unique identifier of the payment terminal.
@return TerminalSettingsTerminalLevelApiGetTerminalSettingsInput
*/
func (a *TerminalSettingsTerminalLevelApi) GetTerminalSettingsInput(terminalId string) TerminalSettingsTerminalLevelApiGetTerminalSettingsInput {
	return TerminalSettingsTerminalLevelApiGetTerminalSettingsInput{
		terminalId: terminalId,
	}
}

/*
GetTerminalSettings Get terminal settings

Returns the settings that are configured for the payment terminal identified in the path.

To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal settings read
* Management API—Terminal settings read and write

For [sensitive terminal settings](https://docs.adyen.com/point-of-sale/automating-terminal-management/configure-terminals-api#sensitive-terminal-settings), your API credential must have the following role:
* Management API—Terminal settings Advanced read and write

In the live environment, requests to this endpoint are subject to [rate limits](https://docs.adyen.com/point-of-sale/automating-terminal-management#rate-limits-in-the-live-environment).

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TerminalSettingsTerminalLevelApiGetTerminalSettingsInput - Request parameters, see GetTerminalSettingsInput
@return TerminalSettings, *http.Response, error
*/
func (a *TerminalSettingsTerminalLevelApi) GetTerminalSettings(ctx context.Context, r TerminalSettingsTerminalLevelApiGetTerminalSettingsInput) (TerminalSettings, *http.Response, error) {
	res := &TerminalSettings{}
	path := "/terminals/{terminalId}/terminalSettings"
	path = strings.Replace(path, "{"+"terminalId"+"}", url.PathEscape(common.ParameterValueToString(r.terminalId, "terminalId")), -1)
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

// All parameters accepted by TerminalSettingsTerminalLevelApi.UpdateLogo
type TerminalSettingsTerminalLevelApiUpdateLogoInput struct {
	terminalId string
	logo       *Logo
}

func (r TerminalSettingsTerminalLevelApiUpdateLogoInput) Logo(logo Logo) TerminalSettingsTerminalLevelApiUpdateLogoInput {
	r.logo = &logo
	return r
}

/*
Prepare a request for UpdateLogo
@param terminalId The unique identifier of the payment terminal.
@return TerminalSettingsTerminalLevelApiUpdateLogoInput
*/
func (a *TerminalSettingsTerminalLevelApi) UpdateLogoInput(terminalId string) TerminalSettingsTerminalLevelApiUpdateLogoInput {
	return TerminalSettingsTerminalLevelApiUpdateLogoInput{
		terminalId: terminalId,
	}
}

/*
UpdateLogo Update the logo

Updates the logo for the payment terminal identified in the path.

* To change the logo, specify the image file as a Base64-encoded string.
* To restore the logo inherited from a higher level (store, merchant account, or company account), specify an empty logo value.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal settings read and write

In the live environment, requests to this endpoint are subject to [rate limits](https://docs.adyen.com/point-of-sale/automating-terminal-management#rate-limits-in-the-live-environment).

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TerminalSettingsTerminalLevelApiUpdateLogoInput - Request parameters, see UpdateLogoInput
@return Logo, *http.Response, error
*/
func (a *TerminalSettingsTerminalLevelApi) UpdateLogo(ctx context.Context, r TerminalSettingsTerminalLevelApiUpdateLogoInput) (Logo, *http.Response, error) {
	res := &Logo{}
	path := "/terminals/{terminalId}/terminalLogos"
	path = strings.Replace(path, "{"+"terminalId"+"}", url.PathEscape(common.ParameterValueToString(r.terminalId, "terminalId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.logo,
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

// All parameters accepted by TerminalSettingsTerminalLevelApi.UpdateTerminalSettings
type TerminalSettingsTerminalLevelApiUpdateTerminalSettingsInput struct {
	terminalId       string
	terminalSettings *TerminalSettings
}

func (r TerminalSettingsTerminalLevelApiUpdateTerminalSettingsInput) TerminalSettings(terminalSettings TerminalSettings) TerminalSettingsTerminalLevelApiUpdateTerminalSettingsInput {
	r.terminalSettings = &terminalSettings
	return r
}

/*
Prepare a request for UpdateTerminalSettings
@param terminalId The unique identifier of the payment terminal.
@return TerminalSettingsTerminalLevelApiUpdateTerminalSettingsInput
*/
func (a *TerminalSettingsTerminalLevelApi) UpdateTerminalSettingsInput(terminalId string) TerminalSettingsTerminalLevelApiUpdateTerminalSettingsInput {
	return TerminalSettingsTerminalLevelApiUpdateTerminalSettingsInput{
		terminalId: terminalId,
	}
}

/*
UpdateTerminalSettings Update terminal settings

Updates the settings that are configured for the payment terminal identified in the path.

* To change a parameter value, include the full object that contains the parameter, even if you don't want to change all parameters in the object.
* To restore a parameter value inherited from a higher level, include the full object that contains the parameter, and specify an empty value for the parameter or omit the parameter.
* Objects that are not included in the request are not updated.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal settings read and write

For [sensitive terminal settings](https://docs.adyen.com/point-of-sale/automating-terminal-management/configure-terminals-api#sensitive-terminal-settings), your API credential must have the following role:
* Management API—Terminal settings Advanced read and write

In the live environment, requests to this endpoint are subject to [rate limits](https://docs.adyen.com/point-of-sale/automating-terminal-management#rate-limits-in-the-live-environment).

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TerminalSettingsTerminalLevelApiUpdateTerminalSettingsInput - Request parameters, see UpdateTerminalSettingsInput
@return TerminalSettings, *http.Response, error
*/
func (a *TerminalSettingsTerminalLevelApi) UpdateTerminalSettings(ctx context.Context, r TerminalSettingsTerminalLevelApiUpdateTerminalSettingsInput) (TerminalSettings, *http.Response, error) {
	res := &TerminalSettings{}
	path := "/terminals/{terminalId}/terminalSettings"
	path = strings.Replace(path, "{"+"terminalId"+"}", url.PathEscape(common.ParameterValueToString(r.terminalId, "terminalId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.terminalSettings,
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
