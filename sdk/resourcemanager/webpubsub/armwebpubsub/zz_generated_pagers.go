//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armwebpubsub

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// OperationsListPager provides operations for iterating over paged responses.
type OperationsListPager struct {
	client    *OperationsClient
	current   OperationsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, OperationsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *OperationsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *OperationsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.OperationList.NextLink == nil || len(*p.current.OperationList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current OperationsListResponse page.
func (p *OperationsListPager) PageResponse() OperationsListResponse {
	return p.current
}

// UsagesListPager provides operations for iterating over paged responses.
type UsagesListPager struct {
	client    *UsagesClient
	current   UsagesListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, UsagesListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *UsagesListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *UsagesListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SignalRServiceUsageList.NextLink == nil || len(*p.current.SignalRServiceUsageList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current UsagesListResponse page.
func (p *UsagesListPager) PageResponse() UsagesListResponse {
	return p.current
}

// WebPubSubHubsListPager provides operations for iterating over paged responses.
type WebPubSubHubsListPager struct {
	client    *WebPubSubHubsClient
	current   WebPubSubHubsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, WebPubSubHubsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *WebPubSubHubsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *WebPubSubHubsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.WebPubSubHubList.NextLink == nil || len(*p.current.WebPubSubHubList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current WebPubSubHubsListResponse page.
func (p *WebPubSubHubsListPager) PageResponse() WebPubSubHubsListResponse {
	return p.current
}

// WebPubSubListByResourceGroupPager provides operations for iterating over paged responses.
type WebPubSubListByResourceGroupPager struct {
	client    *WebPubSubClient
	current   WebPubSubListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, WebPubSubListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *WebPubSubListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *WebPubSubListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.WebPubSubResourceList.NextLink == nil || len(*p.current.WebPubSubResourceList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByResourceGroupHandleError(resp)
		return false
	}
	result, err := p.client.listByResourceGroupHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current WebPubSubListByResourceGroupResponse page.
func (p *WebPubSubListByResourceGroupPager) PageResponse() WebPubSubListByResourceGroupResponse {
	return p.current
}

// WebPubSubListBySubscriptionPager provides operations for iterating over paged responses.
type WebPubSubListBySubscriptionPager struct {
	client    *WebPubSubClient
	current   WebPubSubListBySubscriptionResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, WebPubSubListBySubscriptionResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *WebPubSubListBySubscriptionPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *WebPubSubListBySubscriptionPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.WebPubSubResourceList.NextLink == nil || len(*p.current.WebPubSubResourceList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listBySubscriptionHandleError(resp)
		return false
	}
	result, err := p.client.listBySubscriptionHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current WebPubSubListBySubscriptionResponse page.
func (p *WebPubSubListBySubscriptionPager) PageResponse() WebPubSubListBySubscriptionResponse {
	return p.current
}

// WebPubSubPrivateEndpointConnectionsListPager provides operations for iterating over paged responses.
type WebPubSubPrivateEndpointConnectionsListPager struct {
	client    *WebPubSubPrivateEndpointConnectionsClient
	current   WebPubSubPrivateEndpointConnectionsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, WebPubSubPrivateEndpointConnectionsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *WebPubSubPrivateEndpointConnectionsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *WebPubSubPrivateEndpointConnectionsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PrivateEndpointConnectionList.NextLink == nil || len(*p.current.PrivateEndpointConnectionList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current WebPubSubPrivateEndpointConnectionsListResponse page.
func (p *WebPubSubPrivateEndpointConnectionsListPager) PageResponse() WebPubSubPrivateEndpointConnectionsListResponse {
	return p.current
}

// WebPubSubPrivateLinkResourcesListPager provides operations for iterating over paged responses.
type WebPubSubPrivateLinkResourcesListPager struct {
	client    *WebPubSubPrivateLinkResourcesClient
	current   WebPubSubPrivateLinkResourcesListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, WebPubSubPrivateLinkResourcesListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *WebPubSubPrivateLinkResourcesListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *WebPubSubPrivateLinkResourcesListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PrivateLinkResourceList.NextLink == nil || len(*p.current.PrivateLinkResourceList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current WebPubSubPrivateLinkResourcesListResponse page.
func (p *WebPubSubPrivateLinkResourcesListPager) PageResponse() WebPubSubPrivateLinkResourcesListResponse {
	return p.current
}

// WebPubSubSharedPrivateLinkResourcesListPager provides operations for iterating over paged responses.
type WebPubSubSharedPrivateLinkResourcesListPager struct {
	client    *WebPubSubSharedPrivateLinkResourcesClient
	current   WebPubSubSharedPrivateLinkResourcesListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, WebPubSubSharedPrivateLinkResourcesListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *WebPubSubSharedPrivateLinkResourcesListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *WebPubSubSharedPrivateLinkResourcesListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SharedPrivateLinkResourceList.NextLink == nil || len(*p.current.SharedPrivateLinkResourceList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current WebPubSubSharedPrivateLinkResourcesListResponse page.
func (p *WebPubSubSharedPrivateLinkResourcesListPager) PageResponse() WebPubSubSharedPrivateLinkResourcesListResponse {
	return p.current
}
