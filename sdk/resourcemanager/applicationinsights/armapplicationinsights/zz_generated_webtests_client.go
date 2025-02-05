//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapplicationinsights

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// WebTestsClient contains the methods for the WebTests group.
// Don't use this type directly, use NewWebTestsClient() instead.
type WebTestsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewWebTestsClient creates a new instance of WebTestsClient with the specified values.
func NewWebTestsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *WebTestsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &WebTestsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// CreateOrUpdate - Creates or updates an Application Insights web test definition.
// If the operation fails it returns a generic error.
func (client *WebTestsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, webTestName string, webTestDefinition WebTest, options *WebTestsCreateOrUpdateOptions) (WebTestsCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, webTestName, webTestDefinition, options)
	if err != nil {
		return WebTestsCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WebTestsCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WebTestsCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *WebTestsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, webTestName string, webTestDefinition WebTest, options *WebTestsCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/webtests/{webTestName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if webTestName == "" {
		return nil, errors.New("parameter webTestName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{webTestName}", url.PathEscape(webTestName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, webTestDefinition)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *WebTestsClient) createOrUpdateHandleResponse(resp *http.Response) (WebTestsCreateOrUpdateResponse, error) {
	result := WebTestsCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WebTest); err != nil {
		return WebTestsCreateOrUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *WebTestsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Delete - Deletes an Application Insights web test.
// If the operation fails it returns a generic error.
func (client *WebTestsClient) Delete(ctx context.Context, resourceGroupName string, webTestName string, options *WebTestsDeleteOptions) (WebTestsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, webTestName, options)
	if err != nil {
		return WebTestsDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WebTestsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return WebTestsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return WebTestsDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *WebTestsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, webTestName string, options *WebTestsDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/webtests/{webTestName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if webTestName == "" {
		return nil, errors.New("parameter webTestName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{webTestName}", url.PathEscape(webTestName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *WebTestsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - Get a specific Application Insights web test definition.
// If the operation fails it returns a generic error.
func (client *WebTestsClient) Get(ctx context.Context, resourceGroupName string, webTestName string, options *WebTestsGetOptions) (WebTestsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, webTestName, options)
	if err != nil {
		return WebTestsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WebTestsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WebTestsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WebTestsClient) getCreateRequest(ctx context.Context, resourceGroupName string, webTestName string, options *WebTestsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/webtests/{webTestName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if webTestName == "" {
		return nil, errors.New("parameter webTestName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{webTestName}", url.PathEscape(webTestName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WebTestsClient) getHandleResponse(resp *http.Response) (WebTestsGetResponse, error) {
	result := WebTestsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WebTest); err != nil {
		return WebTestsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *WebTestsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// List - Get all Application Insights web test alerts definitions within a subscription.
// If the operation fails it returns a generic error.
func (client *WebTestsClient) List(options *WebTestsListOptions) *WebTestsListPager {
	return &WebTestsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp WebTestsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.WebTestListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *WebTestsClient) listCreateRequest(ctx context.Context, options *WebTestsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Insights/webtests"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *WebTestsClient) listHandleResponse(resp *http.Response) (WebTestsListResponse, error) {
	result := WebTestsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WebTestListResult); err != nil {
		return WebTestsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *WebTestsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByComponent - Get all Application Insights web tests defined for the specified component.
// If the operation fails it returns a generic error.
func (client *WebTestsClient) ListByComponent(componentName string, resourceGroupName string, options *WebTestsListByComponentOptions) *WebTestsListByComponentPager {
	return &WebTestsListByComponentPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByComponentCreateRequest(ctx, componentName, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp WebTestsListByComponentResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.WebTestListResult.NextLink)
		},
	}
}

// listByComponentCreateRequest creates the ListByComponent request.
func (client *WebTestsClient) listByComponentCreateRequest(ctx context.Context, componentName string, resourceGroupName string, options *WebTestsListByComponentOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{componentName}/webtests"
	if componentName == "" {
		return nil, errors.New("parameter componentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{componentName}", url.PathEscape(componentName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByComponentHandleResponse handles the ListByComponent response.
func (client *WebTestsClient) listByComponentHandleResponse(resp *http.Response) (WebTestsListByComponentResponse, error) {
	result := WebTestsListByComponentResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WebTestListResult); err != nil {
		return WebTestsListByComponentResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByComponentHandleError handles the ListByComponent error response.
func (client *WebTestsClient) listByComponentHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByResourceGroup - Get all Application Insights web tests defined within a specified resource group.
// If the operation fails it returns a generic error.
func (client *WebTestsClient) ListByResourceGroup(resourceGroupName string, options *WebTestsListByResourceGroupOptions) *WebTestsListByResourceGroupPager {
	return &WebTestsListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp WebTestsListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.WebTestListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *WebTestsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *WebTestsListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/webtests"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *WebTestsClient) listByResourceGroupHandleResponse(resp *http.Response) (WebTestsListByResourceGroupResponse, error) {
	result := WebTestsListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WebTestListResult); err != nil {
		return WebTestsListByResourceGroupResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *WebTestsClient) listByResourceGroupHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// UpdateTags - Creates or updates an Application Insights web test definition.
// If the operation fails it returns a generic error.
func (client *WebTestsClient) UpdateTags(ctx context.Context, resourceGroupName string, webTestName string, webTestTags TagsResource, options *WebTestsUpdateTagsOptions) (WebTestsUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, webTestName, webTestTags, options)
	if err != nil {
		return WebTestsUpdateTagsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WebTestsUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WebTestsUpdateTagsResponse{}, client.updateTagsHandleError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *WebTestsClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, webTestName string, webTestTags TagsResource, options *WebTestsUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/webtests/{webTestName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if webTestName == "" {
		return nil, errors.New("parameter webTestName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{webTestName}", url.PathEscape(webTestName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, webTestTags)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *WebTestsClient) updateTagsHandleResponse(resp *http.Response) (WebTestsUpdateTagsResponse, error) {
	result := WebTestsUpdateTagsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WebTest); err != nil {
		return WebTestsUpdateTagsResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *WebTestsClient) updateTagsHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
