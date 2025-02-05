//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armtestbase

import (
	"context"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"net/http"
	"time"
)

// AnalysisResultsGetResponse contains the response from method AnalysisResults.Get.
type AnalysisResultsGetResponse struct {
	AnalysisResultsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AnalysisResultsGetResult contains the result from method AnalysisResults.Get.
type AnalysisResultsGetResult struct {
	AnalysisResultSingletonResource
}

// AnalysisResultsListResponse contains the response from method AnalysisResults.List.
type AnalysisResultsListResponse struct {
	AnalysisResultsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AnalysisResultsListResult contains the result from method AnalysisResults.List.
type AnalysisResultsListResult struct {
	AnalysisResultListResult
}

// AvailableOSGetResponse contains the response from method AvailableOS.Get.
type AvailableOSGetResponse struct {
	AvailableOSGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AvailableOSGetResult contains the result from method AvailableOS.Get.
type AvailableOSGetResult struct {
	AvailableOSResource
}

// AvailableOSListResponse contains the response from method AvailableOS.List.
type AvailableOSListResponse struct {
	AvailableOSListResultEnvelope
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AvailableOSListResultEnvelope contains the result from method AvailableOS.List.
type AvailableOSListResultEnvelope struct {
	AvailableOSListResult
}

// CustomerEventsCreatePollerResponse contains the response from method CustomerEvents.Create.
type CustomerEventsCreatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *CustomerEventsCreatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l CustomerEventsCreatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (CustomerEventsCreateResponse, error) {
	respType := CustomerEventsCreateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.CustomerEventResource)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a CustomerEventsCreatePollerResponse from the provided client and resume token.
func (l *CustomerEventsCreatePollerResponse) Resume(ctx context.Context, client *CustomerEventsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("CustomerEventsClient.Create", token, client.pl, client.createHandleError)
	if err != nil {
		return err
	}
	poller := &CustomerEventsCreatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// CustomerEventsCreateResponse contains the response from method CustomerEvents.Create.
type CustomerEventsCreateResponse struct {
	CustomerEventsCreateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// CustomerEventsCreateResult contains the result from method CustomerEvents.Create.
type CustomerEventsCreateResult struct {
	CustomerEventResource
}

// CustomerEventsDeletePollerResponse contains the response from method CustomerEvents.Delete.
type CustomerEventsDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *CustomerEventsDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l CustomerEventsDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (CustomerEventsDeleteResponse, error) {
	respType := CustomerEventsDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a CustomerEventsDeletePollerResponse from the provided client and resume token.
func (l *CustomerEventsDeletePollerResponse) Resume(ctx context.Context, client *CustomerEventsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("CustomerEventsClient.Delete", token, client.pl, client.deleteHandleError)
	if err != nil {
		return err
	}
	poller := &CustomerEventsDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// CustomerEventsDeleteResponse contains the response from method CustomerEvents.Delete.
type CustomerEventsDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// CustomerEventsGetResponse contains the response from method CustomerEvents.Get.
type CustomerEventsGetResponse struct {
	CustomerEventsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// CustomerEventsGetResult contains the result from method CustomerEvents.Get.
type CustomerEventsGetResult struct {
	CustomerEventResource
}

// CustomerEventsListByTestBaseAccountResponse contains the response from method CustomerEvents.ListByTestBaseAccount.
type CustomerEventsListByTestBaseAccountResponse struct {
	CustomerEventsListByTestBaseAccountResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// CustomerEventsListByTestBaseAccountResult contains the result from method CustomerEvents.ListByTestBaseAccount.
type CustomerEventsListByTestBaseAccountResult struct {
	CustomerEventListResult
}

// EmailEventsGetResponse contains the response from method EmailEvents.Get.
type EmailEventsGetResponse struct {
	EmailEventsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// EmailEventsGetResult contains the result from method EmailEvents.Get.
type EmailEventsGetResult struct {
	EmailEventResource
}

// EmailEventsListResponse contains the response from method EmailEvents.List.
type EmailEventsListResponse struct {
	EmailEventsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// EmailEventsListResult contains the result from method EmailEvents.List.
type EmailEventsListResult struct {
	EmailEventListResult
}

// FavoriteProcessesCreateResponse contains the response from method FavoriteProcesses.Create.
type FavoriteProcessesCreateResponse struct {
	FavoriteProcessesCreateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// FavoriteProcessesCreateResult contains the result from method FavoriteProcesses.Create.
type FavoriteProcessesCreateResult struct {
	FavoriteProcessResource
}

// FavoriteProcessesDeleteResponse contains the response from method FavoriteProcesses.Delete.
type FavoriteProcessesDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// FavoriteProcessesGetResponse contains the response from method FavoriteProcesses.Get.
type FavoriteProcessesGetResponse struct {
	FavoriteProcessesGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// FavoriteProcessesGetResult contains the result from method FavoriteProcesses.Get.
type FavoriteProcessesGetResult struct {
	FavoriteProcessResource
}

// FavoriteProcessesListResponse contains the response from method FavoriteProcesses.List.
type FavoriteProcessesListResponse struct {
	FavoriteProcessesListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// FavoriteProcessesListResult contains the result from method FavoriteProcesses.List.
type FavoriteProcessesListResult struct {
	FavoriteProcessListResult
}

// FlightingRingsGetResponse contains the response from method FlightingRings.Get.
type FlightingRingsGetResponse struct {
	FlightingRingsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// FlightingRingsGetResult contains the result from method FlightingRings.Get.
type FlightingRingsGetResult struct {
	FlightingRingResource
}

// FlightingRingsListResponse contains the response from method FlightingRings.List.
type FlightingRingsListResponse struct {
	FlightingRingsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// FlightingRingsListResult contains the result from method FlightingRings.List.
type FlightingRingsListResult struct {
	FlightingRingListResult
}

// OSUpdatesGetResponse contains the response from method OSUpdates.Get.
type OSUpdatesGetResponse struct {
	OSUpdatesGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OSUpdatesGetResult contains the result from method OSUpdates.Get.
type OSUpdatesGetResult struct {
	OSUpdateResource
}

// OSUpdatesListResponse contains the response from method OSUpdates.List.
type OSUpdatesListResponse struct {
	OSUpdatesListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OSUpdatesListResult contains the result from method OSUpdates.List.
type OSUpdatesListResult struct {
	OSUpdateListResult
}

// OperationsListResponse contains the response from method Operations.List.
type OperationsListResponse struct {
	OperationsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsListResult contains the result from method Operations.List.
type OperationsListResult struct {
	OperationListResult
}

// PackagesCreatePollerResponse contains the response from method Packages.Create.
type PackagesCreatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *PackagesCreatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l PackagesCreatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (PackagesCreateResponse, error) {
	respType := PackagesCreateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.PackageResource)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a PackagesCreatePollerResponse from the provided client and resume token.
func (l *PackagesCreatePollerResponse) Resume(ctx context.Context, client *PackagesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("PackagesClient.Create", token, client.pl, client.createHandleError)
	if err != nil {
		return err
	}
	poller := &PackagesCreatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// PackagesCreateResponse contains the response from method Packages.Create.
type PackagesCreateResponse struct {
	PackagesCreateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PackagesCreateResult contains the result from method Packages.Create.
type PackagesCreateResult struct {
	PackageResource
}

// PackagesDeletePollerResponse contains the response from method Packages.Delete.
type PackagesDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *PackagesDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l PackagesDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (PackagesDeleteResponse, error) {
	respType := PackagesDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a PackagesDeletePollerResponse from the provided client and resume token.
func (l *PackagesDeletePollerResponse) Resume(ctx context.Context, client *PackagesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("PackagesClient.Delete", token, client.pl, client.deleteHandleError)
	if err != nil {
		return err
	}
	poller := &PackagesDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// PackagesDeleteResponse contains the response from method Packages.Delete.
type PackagesDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PackagesGetDownloadURLResponse contains the response from method Packages.GetDownloadURL.
type PackagesGetDownloadURLResponse struct {
	PackagesGetDownloadURLResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PackagesGetDownloadURLResult contains the result from method Packages.GetDownloadURL.
type PackagesGetDownloadURLResult struct {
	DownloadURLResponse
}

// PackagesGetResponse contains the response from method Packages.Get.
type PackagesGetResponse struct {
	PackagesGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PackagesGetResult contains the result from method Packages.Get.
type PackagesGetResult struct {
	PackageResource
}

// PackagesHardDeletePollerResponse contains the response from method Packages.HardDelete.
type PackagesHardDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *PackagesHardDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l PackagesHardDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (PackagesHardDeleteResponse, error) {
	respType := PackagesHardDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a PackagesHardDeletePollerResponse from the provided client and resume token.
func (l *PackagesHardDeletePollerResponse) Resume(ctx context.Context, client *PackagesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("PackagesClient.HardDelete", token, client.pl, client.hardDeleteHandleError)
	if err != nil {
		return err
	}
	poller := &PackagesHardDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// PackagesHardDeleteResponse contains the response from method Packages.HardDelete.
type PackagesHardDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PackagesListByTestBaseAccountResponse contains the response from method Packages.ListByTestBaseAccount.
type PackagesListByTestBaseAccountResponse struct {
	PackagesListByTestBaseAccountResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PackagesListByTestBaseAccountResult contains the result from method Packages.ListByTestBaseAccount.
type PackagesListByTestBaseAccountResult struct {
	PackageListResult
}

// PackagesUpdatePollerResponse contains the response from method Packages.Update.
type PackagesUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *PackagesUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l PackagesUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (PackagesUpdateResponse, error) {
	respType := PackagesUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.PackageResource)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a PackagesUpdatePollerResponse from the provided client and resume token.
func (l *PackagesUpdatePollerResponse) Resume(ctx context.Context, client *PackagesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("PackagesClient.Update", token, client.pl, client.updateHandleError)
	if err != nil {
		return err
	}
	poller := &PackagesUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// PackagesUpdateResponse contains the response from method Packages.Update.
type PackagesUpdateResponse struct {
	PackagesUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PackagesUpdateResult contains the result from method Packages.Update.
type PackagesUpdateResult struct {
	PackageResource
}

// SKUsListResponse contains the response from method SKUs.List.
type SKUsListResponse struct {
	SKUsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SKUsListResult contains the result from method SKUs.List.
type SKUsListResult struct {
	TestBaseAccountSKUListResult
}

// TestBaseAccountsCheckPackageNameAvailabilityResponse contains the response from method TestBaseAccounts.CheckPackageNameAvailability.
type TestBaseAccountsCheckPackageNameAvailabilityResponse struct {
	TestBaseAccountsCheckPackageNameAvailabilityResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TestBaseAccountsCheckPackageNameAvailabilityResult contains the result from method TestBaseAccounts.CheckPackageNameAvailability.
type TestBaseAccountsCheckPackageNameAvailabilityResult struct {
	CheckNameAvailabilityResult
}

// TestBaseAccountsCreatePollerResponse contains the response from method TestBaseAccounts.Create.
type TestBaseAccountsCreatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *TestBaseAccountsCreatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l TestBaseAccountsCreatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (TestBaseAccountsCreateResponse, error) {
	respType := TestBaseAccountsCreateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.TestBaseAccountResource)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a TestBaseAccountsCreatePollerResponse from the provided client and resume token.
func (l *TestBaseAccountsCreatePollerResponse) Resume(ctx context.Context, client *TestBaseAccountsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("TestBaseAccountsClient.Create", token, client.pl, client.createHandleError)
	if err != nil {
		return err
	}
	poller := &TestBaseAccountsCreatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// TestBaseAccountsCreateResponse contains the response from method TestBaseAccounts.Create.
type TestBaseAccountsCreateResponse struct {
	TestBaseAccountsCreateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TestBaseAccountsCreateResult contains the result from method TestBaseAccounts.Create.
type TestBaseAccountsCreateResult struct {
	TestBaseAccountResource
}

// TestBaseAccountsDeletePollerResponse contains the response from method TestBaseAccounts.Delete.
type TestBaseAccountsDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *TestBaseAccountsDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l TestBaseAccountsDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (TestBaseAccountsDeleteResponse, error) {
	respType := TestBaseAccountsDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a TestBaseAccountsDeletePollerResponse from the provided client and resume token.
func (l *TestBaseAccountsDeletePollerResponse) Resume(ctx context.Context, client *TestBaseAccountsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("TestBaseAccountsClient.Delete", token, client.pl, client.deleteHandleError)
	if err != nil {
		return err
	}
	poller := &TestBaseAccountsDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// TestBaseAccountsDeleteResponse contains the response from method TestBaseAccounts.Delete.
type TestBaseAccountsDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TestBaseAccountsGetFileUploadURLResponse contains the response from method TestBaseAccounts.GetFileUploadURL.
type TestBaseAccountsGetFileUploadURLResponse struct {
	TestBaseAccountsGetFileUploadURLResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TestBaseAccountsGetFileUploadURLResult contains the result from method TestBaseAccounts.GetFileUploadURL.
type TestBaseAccountsGetFileUploadURLResult struct {
	FileUploadURLResponse
}

// TestBaseAccountsGetResponse contains the response from method TestBaseAccounts.Get.
type TestBaseAccountsGetResponse struct {
	TestBaseAccountsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TestBaseAccountsGetResult contains the result from method TestBaseAccounts.Get.
type TestBaseAccountsGetResult struct {
	TestBaseAccountResource
}

// TestBaseAccountsListByResourceGroupResponse contains the response from method TestBaseAccounts.ListByResourceGroup.
type TestBaseAccountsListByResourceGroupResponse struct {
	TestBaseAccountsListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TestBaseAccountsListByResourceGroupResult contains the result from method TestBaseAccounts.ListByResourceGroup.
type TestBaseAccountsListByResourceGroupResult struct {
	TestBaseAccountListResult
}

// TestBaseAccountsListBySubscriptionResponse contains the response from method TestBaseAccounts.ListBySubscription.
type TestBaseAccountsListBySubscriptionResponse struct {
	TestBaseAccountsListBySubscriptionResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TestBaseAccountsListBySubscriptionResult contains the result from method TestBaseAccounts.ListBySubscription.
type TestBaseAccountsListBySubscriptionResult struct {
	TestBaseAccountListResult
}

// TestBaseAccountsOffboardPollerResponse contains the response from method TestBaseAccounts.Offboard.
type TestBaseAccountsOffboardPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *TestBaseAccountsOffboardPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l TestBaseAccountsOffboardPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (TestBaseAccountsOffboardResponse, error) {
	respType := TestBaseAccountsOffboardResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a TestBaseAccountsOffboardPollerResponse from the provided client and resume token.
func (l *TestBaseAccountsOffboardPollerResponse) Resume(ctx context.Context, client *TestBaseAccountsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("TestBaseAccountsClient.Offboard", token, client.pl, client.offboardHandleError)
	if err != nil {
		return err
	}
	poller := &TestBaseAccountsOffboardPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// TestBaseAccountsOffboardResponse contains the response from method TestBaseAccounts.Offboard.
type TestBaseAccountsOffboardResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TestBaseAccountsUpdatePollerResponse contains the response from method TestBaseAccounts.Update.
type TestBaseAccountsUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *TestBaseAccountsUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l TestBaseAccountsUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (TestBaseAccountsUpdateResponse, error) {
	respType := TestBaseAccountsUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.TestBaseAccountResource)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a TestBaseAccountsUpdatePollerResponse from the provided client and resume token.
func (l *TestBaseAccountsUpdatePollerResponse) Resume(ctx context.Context, client *TestBaseAccountsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("TestBaseAccountsClient.Update", token, client.pl, client.updateHandleError)
	if err != nil {
		return err
	}
	poller := &TestBaseAccountsUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// TestBaseAccountsUpdateResponse contains the response from method TestBaseAccounts.Update.
type TestBaseAccountsUpdateResponse struct {
	TestBaseAccountsUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TestBaseAccountsUpdateResult contains the result from method TestBaseAccounts.Update.
type TestBaseAccountsUpdateResult struct {
	TestBaseAccountResource
}

// TestResultsGetDownloadURLResponse contains the response from method TestResults.GetDownloadURL.
type TestResultsGetDownloadURLResponse struct {
	TestResultsGetDownloadURLResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TestResultsGetDownloadURLResult contains the result from method TestResults.GetDownloadURL.
type TestResultsGetDownloadURLResult struct {
	DownloadURLResponse
}

// TestResultsGetResponse contains the response from method TestResults.Get.
type TestResultsGetResponse struct {
	TestResultsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TestResultsGetResult contains the result from method TestResults.Get.
type TestResultsGetResult struct {
	TestResultResource
}

// TestResultsGetVideoDownloadURLResponse contains the response from method TestResults.GetVideoDownloadURL.
type TestResultsGetVideoDownloadURLResponse struct {
	TestResultsGetVideoDownloadURLResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TestResultsGetVideoDownloadURLResult contains the result from method TestResults.GetVideoDownloadURL.
type TestResultsGetVideoDownloadURLResult struct {
	DownloadURLResponse
}

// TestResultsListResponse contains the response from method TestResults.List.
type TestResultsListResponse struct {
	TestResultsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TestResultsListResult contains the result from method TestResults.List.
type TestResultsListResult struct {
	TestResultListResult
}

// TestSummariesGetResponse contains the response from method TestSummaries.Get.
type TestSummariesGetResponse struct {
	TestSummariesGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TestSummariesGetResult contains the result from method TestSummaries.Get.
type TestSummariesGetResult struct {
	TestSummaryResource
}

// TestSummariesListResponse contains the response from method TestSummaries.List.
type TestSummariesListResponse struct {
	TestSummariesListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TestSummariesListResult contains the result from method TestSummaries.List.
type TestSummariesListResult struct {
	TestSummaryListResult
}

// TestTypesGetResponse contains the response from method TestTypes.Get.
type TestTypesGetResponse struct {
	TestTypesGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TestTypesGetResult contains the result from method TestTypes.Get.
type TestTypesGetResult struct {
	TestTypeResource
}

// TestTypesListResponse contains the response from method TestTypes.List.
type TestTypesListResponse struct {
	TestTypesListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TestTypesListResult contains the result from method TestTypes.List.
type TestTypesListResult struct {
	TestTypeListResult
}

// UsageListResponse contains the response from method Usage.List.
type UsageListResponse struct {
	UsageListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// UsageListResult contains the result from method Usage.List.
type UsageListResult struct {
	TestBaseAccountUsageDataList
}
