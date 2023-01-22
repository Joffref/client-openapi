/*
M5_MetricsReporting

5GMS AF M5 Metrics Reporting API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package M5_MetricsReporting

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// DefaultApiService DefaultApi service
type DefaultApiService service

type ApiSubmitMetricsReportRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	provisioningSessionId string
	metricsReportingConfigurationId string
	body *string
}

// A Metrics Report
func (r ApiSubmitMetricsReportRequest) Body(body string) ApiSubmitMetricsReportRequest {
	r.body = &body
	return r
}

func (r ApiSubmitMetricsReportRequest) Execute() (*http.Response, error) {
	return r.ApiService.SubmitMetricsReportExecute(r)
}

/*
SubmitMetricsReport Submit a Metrics Report

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param provisioningSessionId The resource identifier of an existing Provisioning Session.
 @param metricsReportingConfigurationId The resource identifier of a Metrics Configuration in the specified Provisioning Session.
 @return ApiSubmitMetricsReportRequest
*/
func (a *DefaultApiService) SubmitMetricsReport(ctx context.Context, provisioningSessionId string, metricsReportingConfigurationId string) ApiSubmitMetricsReportRequest {
	return ApiSubmitMetricsReportRequest{
		ApiService: a,
		ctx: ctx,
		provisioningSessionId: provisioningSessionId,
		metricsReportingConfigurationId: metricsReportingConfigurationId,
	}
}

// Execute executes the request
func (a *DefaultApiService) SubmitMetricsReportExecute(r ApiSubmitMetricsReportRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.SubmitMetricsReport")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/metrics-reporting/{provisioningSessionId}/{metricsReportingConfigurationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"provisioningSessionId"+"}", url.PathEscape(parameterToString(r.provisioningSessionId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"metricsReportingConfigurationId"+"}", url.PathEscape(parameterToString(r.metricsReportingConfigurationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/3gpdash-qoe-report+xml", "application/*"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}