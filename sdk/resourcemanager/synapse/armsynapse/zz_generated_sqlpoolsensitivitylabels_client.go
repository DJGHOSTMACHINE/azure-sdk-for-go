//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsynapse

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// SQLPoolSensitivityLabelsClient contains the methods for the SQLPoolSensitivityLabels group.
// Don't use this type directly, use NewSQLPoolSensitivityLabelsClient() instead.
type SQLPoolSensitivityLabelsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewSQLPoolSensitivityLabelsClient creates a new instance of SQLPoolSensitivityLabelsClient with the specified values.
func NewSQLPoolSensitivityLabelsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *SQLPoolSensitivityLabelsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &SQLPoolSensitivityLabelsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// CreateOrUpdate - Creates or updates the sensitivity label of a given column in a Sql pool
// If the operation fails it returns a generic error.
func (client *SQLPoolSensitivityLabelsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, schemaName string, tableName string, columnName string, parameters SensitivityLabel, options *SQLPoolSensitivityLabelsCreateOrUpdateOptions) (SQLPoolSensitivityLabelsCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, sqlPoolName, schemaName, tableName, columnName, parameters, options)
	if err != nil {
		return SQLPoolSensitivityLabelsCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SQLPoolSensitivityLabelsCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return SQLPoolSensitivityLabelsCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SQLPoolSensitivityLabelsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, schemaName string, tableName string, columnName string, parameters SensitivityLabel, options *SQLPoolSensitivityLabelsCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/schemas/{schemaName}/tables/{tableName}/columns/{columnName}/sensitivityLabels/{sensitivityLabelSource}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if sqlPoolName == "" {
		return nil, errors.New("parameter sqlPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlPoolName}", url.PathEscape(sqlPoolName))
	if schemaName == "" {
		return nil, errors.New("parameter schemaName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{schemaName}", url.PathEscape(schemaName))
	if tableName == "" {
		return nil, errors.New("parameter tableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tableName}", url.PathEscape(tableName))
	if columnName == "" {
		return nil, errors.New("parameter columnName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{columnName}", url.PathEscape(columnName))
	urlPath = strings.ReplaceAll(urlPath, "{sensitivityLabelSource}", url.PathEscape("current"))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *SQLPoolSensitivityLabelsClient) createOrUpdateHandleResponse(resp *http.Response) (SQLPoolSensitivityLabelsCreateOrUpdateResponse, error) {
	result := SQLPoolSensitivityLabelsCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SensitivityLabel); err != nil {
		return SQLPoolSensitivityLabelsCreateOrUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *SQLPoolSensitivityLabelsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Delete - Deletes the sensitivity label of a given column in a Sql pool
// If the operation fails it returns a generic error.
func (client *SQLPoolSensitivityLabelsClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, schemaName string, tableName string, columnName string, options *SQLPoolSensitivityLabelsDeleteOptions) (SQLPoolSensitivityLabelsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, sqlPoolName, schemaName, tableName, columnName, options)
	if err != nil {
		return SQLPoolSensitivityLabelsDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SQLPoolSensitivityLabelsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return SQLPoolSensitivityLabelsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return SQLPoolSensitivityLabelsDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SQLPoolSensitivityLabelsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, schemaName string, tableName string, columnName string, options *SQLPoolSensitivityLabelsDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/schemas/{schemaName}/tables/{tableName}/columns/{columnName}/sensitivityLabels/{sensitivityLabelSource}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if sqlPoolName == "" {
		return nil, errors.New("parameter sqlPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlPoolName}", url.PathEscape(sqlPoolName))
	if schemaName == "" {
		return nil, errors.New("parameter schemaName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{schemaName}", url.PathEscape(schemaName))
	if tableName == "" {
		return nil, errors.New("parameter tableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tableName}", url.PathEscape(tableName))
	if columnName == "" {
		return nil, errors.New("parameter columnName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{columnName}", url.PathEscape(columnName))
	urlPath = strings.ReplaceAll(urlPath, "{sensitivityLabelSource}", url.PathEscape("current"))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *SQLPoolSensitivityLabelsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// DisableRecommendation - Disables sensitivity recommendations on a given column
// If the operation fails it returns a generic error.
func (client *SQLPoolSensitivityLabelsClient) DisableRecommendation(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, schemaName string, tableName string, columnName string, options *SQLPoolSensitivityLabelsDisableRecommendationOptions) (SQLPoolSensitivityLabelsDisableRecommendationResponse, error) {
	req, err := client.disableRecommendationCreateRequest(ctx, resourceGroupName, workspaceName, sqlPoolName, schemaName, tableName, columnName, options)
	if err != nil {
		return SQLPoolSensitivityLabelsDisableRecommendationResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SQLPoolSensitivityLabelsDisableRecommendationResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SQLPoolSensitivityLabelsDisableRecommendationResponse{}, client.disableRecommendationHandleError(resp)
	}
	return SQLPoolSensitivityLabelsDisableRecommendationResponse{RawResponse: resp}, nil
}

// disableRecommendationCreateRequest creates the DisableRecommendation request.
func (client *SQLPoolSensitivityLabelsClient) disableRecommendationCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, schemaName string, tableName string, columnName string, options *SQLPoolSensitivityLabelsDisableRecommendationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/schemas/{schemaName}/tables/{tableName}/columns/{columnName}/sensitivityLabels/{sensitivityLabelSource}/disable"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if sqlPoolName == "" {
		return nil, errors.New("parameter sqlPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlPoolName}", url.PathEscape(sqlPoolName))
	if schemaName == "" {
		return nil, errors.New("parameter schemaName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{schemaName}", url.PathEscape(schemaName))
	if tableName == "" {
		return nil, errors.New("parameter tableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tableName}", url.PathEscape(tableName))
	if columnName == "" {
		return nil, errors.New("parameter columnName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{columnName}", url.PathEscape(columnName))
	urlPath = strings.ReplaceAll(urlPath, "{sensitivityLabelSource}", url.PathEscape("recommended"))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// disableRecommendationHandleError handles the DisableRecommendation error response.
func (client *SQLPoolSensitivityLabelsClient) disableRecommendationHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// EnableRecommendation - Enables sensitivity recommendations on a given column (recommendations are enabled by default on all columns)
// If the operation fails it returns a generic error.
func (client *SQLPoolSensitivityLabelsClient) EnableRecommendation(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, schemaName string, tableName string, columnName string, options *SQLPoolSensitivityLabelsEnableRecommendationOptions) (SQLPoolSensitivityLabelsEnableRecommendationResponse, error) {
	req, err := client.enableRecommendationCreateRequest(ctx, resourceGroupName, workspaceName, sqlPoolName, schemaName, tableName, columnName, options)
	if err != nil {
		return SQLPoolSensitivityLabelsEnableRecommendationResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SQLPoolSensitivityLabelsEnableRecommendationResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SQLPoolSensitivityLabelsEnableRecommendationResponse{}, client.enableRecommendationHandleError(resp)
	}
	return SQLPoolSensitivityLabelsEnableRecommendationResponse{RawResponse: resp}, nil
}

// enableRecommendationCreateRequest creates the EnableRecommendation request.
func (client *SQLPoolSensitivityLabelsClient) enableRecommendationCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, schemaName string, tableName string, columnName string, options *SQLPoolSensitivityLabelsEnableRecommendationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/schemas/{schemaName}/tables/{tableName}/columns/{columnName}/sensitivityLabels/{sensitivityLabelSource}/enable"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if sqlPoolName == "" {
		return nil, errors.New("parameter sqlPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlPoolName}", url.PathEscape(sqlPoolName))
	if schemaName == "" {
		return nil, errors.New("parameter schemaName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{schemaName}", url.PathEscape(schemaName))
	if tableName == "" {
		return nil, errors.New("parameter tableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tableName}", url.PathEscape(tableName))
	if columnName == "" {
		return nil, errors.New("parameter columnName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{columnName}", url.PathEscape(columnName))
	urlPath = strings.ReplaceAll(urlPath, "{sensitivityLabelSource}", url.PathEscape("recommended"))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// enableRecommendationHandleError handles the EnableRecommendation error response.
func (client *SQLPoolSensitivityLabelsClient) enableRecommendationHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - Gets the sensitivity label of a given column
// If the operation fails it returns a generic error.
func (client *SQLPoolSensitivityLabelsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, schemaName string, tableName string, columnName string, sensitivityLabelSource SensitivityLabelSource, options *SQLPoolSensitivityLabelsGetOptions) (SQLPoolSensitivityLabelsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, sqlPoolName, schemaName, tableName, columnName, sensitivityLabelSource, options)
	if err != nil {
		return SQLPoolSensitivityLabelsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SQLPoolSensitivityLabelsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SQLPoolSensitivityLabelsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SQLPoolSensitivityLabelsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, schemaName string, tableName string, columnName string, sensitivityLabelSource SensitivityLabelSource, options *SQLPoolSensitivityLabelsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/schemas/{schemaName}/tables/{tableName}/columns/{columnName}/sensitivityLabels/{sensitivityLabelSource}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if sqlPoolName == "" {
		return nil, errors.New("parameter sqlPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlPoolName}", url.PathEscape(sqlPoolName))
	if schemaName == "" {
		return nil, errors.New("parameter schemaName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{schemaName}", url.PathEscape(schemaName))
	if tableName == "" {
		return nil, errors.New("parameter tableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tableName}", url.PathEscape(tableName))
	if columnName == "" {
		return nil, errors.New("parameter columnName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{columnName}", url.PathEscape(columnName))
	if sensitivityLabelSource == "" {
		return nil, errors.New("parameter sensitivityLabelSource cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sensitivityLabelSource}", url.PathEscape(string(sensitivityLabelSource)))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SQLPoolSensitivityLabelsClient) getHandleResponse(resp *http.Response) (SQLPoolSensitivityLabelsGetResponse, error) {
	result := SQLPoolSensitivityLabelsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SensitivityLabel); err != nil {
		return SQLPoolSensitivityLabelsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *SQLPoolSensitivityLabelsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListCurrent - Gets SQL pool sensitivity labels.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SQLPoolSensitivityLabelsClient) ListCurrent(resourceGroupName string, workspaceName string, sqlPoolName string, options *SQLPoolSensitivityLabelsListCurrentOptions) *SQLPoolSensitivityLabelsListCurrentPager {
	return &SQLPoolSensitivityLabelsListCurrentPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCurrentCreateRequest(ctx, resourceGroupName, workspaceName, sqlPoolName, options)
		},
		advancer: func(ctx context.Context, resp SQLPoolSensitivityLabelsListCurrentResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SensitivityLabelListResult.NextLink)
		},
	}
}

// listCurrentCreateRequest creates the ListCurrent request.
func (client *SQLPoolSensitivityLabelsClient) listCurrentCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, options *SQLPoolSensitivityLabelsListCurrentOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/currentSensitivityLabels"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if sqlPoolName == "" {
		return nil, errors.New("parameter sqlPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlPoolName}", url.PathEscape(sqlPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listCurrentHandleResponse handles the ListCurrent response.
func (client *SQLPoolSensitivityLabelsClient) listCurrentHandleResponse(resp *http.Response) (SQLPoolSensitivityLabelsListCurrentResponse, error) {
	result := SQLPoolSensitivityLabelsListCurrentResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SensitivityLabelListResult); err != nil {
		return SQLPoolSensitivityLabelsListCurrentResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listCurrentHandleError handles the ListCurrent error response.
func (client *SQLPoolSensitivityLabelsClient) listCurrentHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListRecommended - Gets sensitivity labels of a given SQL pool.
// If the operation fails it returns a generic error.
func (client *SQLPoolSensitivityLabelsClient) ListRecommended(resourceGroupName string, workspaceName string, sqlPoolName string, options *SQLPoolSensitivityLabelsListRecommendedOptions) *SQLPoolSensitivityLabelsListRecommendedPager {
	return &SQLPoolSensitivityLabelsListRecommendedPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listRecommendedCreateRequest(ctx, resourceGroupName, workspaceName, sqlPoolName, options)
		},
		advancer: func(ctx context.Context, resp SQLPoolSensitivityLabelsListRecommendedResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SensitivityLabelListResult.NextLink)
		},
	}
}

// listRecommendedCreateRequest creates the ListRecommended request.
func (client *SQLPoolSensitivityLabelsClient) listRecommendedCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, options *SQLPoolSensitivityLabelsListRecommendedOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/recommendedSensitivityLabels"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if sqlPoolName == "" {
		return nil, errors.New("parameter sqlPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlPoolName}", url.PathEscape(sqlPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	if options != nil && options.IncludeDisabledRecommendations != nil {
		reqQP.Set("includeDisabledRecommendations", strconv.FormatBool(*options.IncludeDisabledRecommendations))
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listRecommendedHandleResponse handles the ListRecommended response.
func (client *SQLPoolSensitivityLabelsClient) listRecommendedHandleResponse(resp *http.Response) (SQLPoolSensitivityLabelsListRecommendedResponse, error) {
	result := SQLPoolSensitivityLabelsListRecommendedResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SensitivityLabelListResult); err != nil {
		return SQLPoolSensitivityLabelsListRecommendedResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listRecommendedHandleError handles the ListRecommended error response.
func (client *SQLPoolSensitivityLabelsClient) listRecommendedHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Update - Update sensitivity labels of a given SQL Pool using an operations batch.
// If the operation fails it returns a generic error.
func (client *SQLPoolSensitivityLabelsClient) Update(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, parameters SensitivityLabelUpdateList, options *SQLPoolSensitivityLabelsUpdateOptions) (SQLPoolSensitivityLabelsUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, workspaceName, sqlPoolName, parameters, options)
	if err != nil {
		return SQLPoolSensitivityLabelsUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SQLPoolSensitivityLabelsUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SQLPoolSensitivityLabelsUpdateResponse{}, client.updateHandleError(resp)
	}
	return SQLPoolSensitivityLabelsUpdateResponse{RawResponse: resp}, nil
}

// updateCreateRequest creates the Update request.
func (client *SQLPoolSensitivityLabelsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, parameters SensitivityLabelUpdateList, options *SQLPoolSensitivityLabelsUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/currentSensitivityLabels"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if sqlPoolName == "" {
		return nil, errors.New("parameter sqlPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlPoolName}", url.PathEscape(sqlPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleError handles the Update error response.
func (client *SQLPoolSensitivityLabelsClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
