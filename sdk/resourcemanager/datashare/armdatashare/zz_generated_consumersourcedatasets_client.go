//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatashare

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
	"strings"
)

// ConsumerSourceDataSetsClient contains the methods for the ConsumerSourceDataSets group.
// Don't use this type directly, use NewConsumerSourceDataSetsClient() instead.
type ConsumerSourceDataSetsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewConsumerSourceDataSetsClient creates a new instance of ConsumerSourceDataSetsClient with the specified values.
func NewConsumerSourceDataSetsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ConsumerSourceDataSetsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ConsumerSourceDataSetsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// ListByShareSubscription - Get source dataSets of a shareSubscription
// If the operation fails it returns the *DataShareError error type.
func (client *ConsumerSourceDataSetsClient) ListByShareSubscription(resourceGroupName string, accountName string, shareSubscriptionName string, options *ConsumerSourceDataSetsListByShareSubscriptionOptions) *ConsumerSourceDataSetsListByShareSubscriptionPager {
	return &ConsumerSourceDataSetsListByShareSubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByShareSubscriptionCreateRequest(ctx, resourceGroupName, accountName, shareSubscriptionName, options)
		},
		advancer: func(ctx context.Context, resp ConsumerSourceDataSetsListByShareSubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ConsumerSourceDataSetList.NextLink)
		},
	}
}

// listByShareSubscriptionCreateRequest creates the ListByShareSubscription request.
func (client *ConsumerSourceDataSetsClient) listByShareSubscriptionCreateRequest(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, options *ConsumerSourceDataSetsListByShareSubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shareSubscriptions/{shareSubscriptionName}/consumerSourceDataSets"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if shareSubscriptionName == "" {
		return nil, errors.New("parameter shareSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{shareSubscriptionName}", url.PathEscape(shareSubscriptionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByShareSubscriptionHandleResponse handles the ListByShareSubscription response.
func (client *ConsumerSourceDataSetsClient) listByShareSubscriptionHandleResponse(resp *http.Response) (ConsumerSourceDataSetsListByShareSubscriptionResponse, error) {
	result := ConsumerSourceDataSetsListByShareSubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConsumerSourceDataSetList); err != nil {
		return ConsumerSourceDataSetsListByShareSubscriptionResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByShareSubscriptionHandleError handles the ListByShareSubscription error response.
func (client *ConsumerSourceDataSetsClient) listByShareSubscriptionHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := DataShareError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
