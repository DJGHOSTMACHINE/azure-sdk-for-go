//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armchaos

import (
	"context"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"net/http"
	"time"
)

// CapabilitiesCreateOrUpdateResponse contains the response from method Capabilities.CreateOrUpdate.
type CapabilitiesCreateOrUpdateResponse struct {
	CapabilitiesCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// CapabilitiesCreateOrUpdateResult contains the result from method Capabilities.CreateOrUpdate.
type CapabilitiesCreateOrUpdateResult struct {
	Capability
}

// CapabilitiesDeleteResponse contains the response from method Capabilities.Delete.
type CapabilitiesDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// CapabilitiesGetResponse contains the response from method Capabilities.Get.
type CapabilitiesGetResponse struct {
	CapabilitiesGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// CapabilitiesGetResult contains the result from method Capabilities.Get.
type CapabilitiesGetResult struct {
	Capability
}

// CapabilitiesListResponse contains the response from method Capabilities.List.
type CapabilitiesListResponse struct {
	CapabilitiesListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// CapabilitiesListResult contains the result from method Capabilities.List.
type CapabilitiesListResult struct {
	CapabilityListResult
}

// CapabilityTypesGetResponse contains the response from method CapabilityTypes.Get.
type CapabilityTypesGetResponse struct {
	CapabilityTypesGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// CapabilityTypesGetResult contains the result from method CapabilityTypes.Get.
type CapabilityTypesGetResult struct {
	CapabilityType
}

// CapabilityTypesListResponse contains the response from method CapabilityTypes.List.
type CapabilityTypesListResponse struct {
	CapabilityTypesListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// CapabilityTypesListResult contains the result from method CapabilityTypes.List.
type CapabilityTypesListResult struct {
	CapabilityTypeListResult
}

// ExperimentsCancelPollerResponse contains the response from method Experiments.Cancel.
type ExperimentsCancelPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ExperimentsCancelPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ExperimentsCancelPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ExperimentsCancelResponse, error) {
	respType := ExperimentsCancelResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.ExperimentCancelOperationResult)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a ExperimentsCancelPollerResponse from the provided client and resume token.
func (l *ExperimentsCancelPollerResponse) Resume(ctx context.Context, client *ExperimentsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ExperimentsClient.Cancel", token, client.pl, client.cancelHandleError)
	if err != nil {
		return err
	}
	poller := &ExperimentsCancelPoller{
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

// ExperimentsCancelResponse contains the response from method Experiments.Cancel.
type ExperimentsCancelResponse struct {
	ExperimentsCancelResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ExperimentsCancelResult contains the result from method Experiments.Cancel.
type ExperimentsCancelResult struct {
	ExperimentCancelOperationResult
}

// ExperimentsCreateOrUpdatePollerResponse contains the response from method Experiments.CreateOrUpdate.
type ExperimentsCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *ExperimentsCreateOrUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l ExperimentsCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (ExperimentsCreateOrUpdateResponse, error) {
	respType := ExperimentsCreateOrUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.Experiment)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a ExperimentsCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *ExperimentsCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *ExperimentsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("ExperimentsClient.CreateOrUpdate", token, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return err
	}
	poller := &ExperimentsCreateOrUpdatePoller{
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

// ExperimentsCreateOrUpdateResponse contains the response from method Experiments.CreateOrUpdate.
type ExperimentsCreateOrUpdateResponse struct {
	ExperimentsCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ExperimentsCreateOrUpdateResult contains the result from method Experiments.CreateOrUpdate.
type ExperimentsCreateOrUpdateResult struct {
	Experiment
}

// ExperimentsDeleteResponse contains the response from method Experiments.Delete.
type ExperimentsDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ExperimentsGetExecutionDetailsResponse contains the response from method Experiments.GetExecutionDetails.
type ExperimentsGetExecutionDetailsResponse struct {
	ExperimentsGetExecutionDetailsResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ExperimentsGetExecutionDetailsResult contains the result from method Experiments.GetExecutionDetails.
type ExperimentsGetExecutionDetailsResult struct {
	ExperimentExecutionDetails
}

// ExperimentsGetResponse contains the response from method Experiments.Get.
type ExperimentsGetResponse struct {
	ExperimentsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ExperimentsGetResult contains the result from method Experiments.Get.
type ExperimentsGetResult struct {
	Experiment
}

// ExperimentsGetStatusResponse contains the response from method Experiments.GetStatus.
type ExperimentsGetStatusResponse struct {
	ExperimentsGetStatusResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ExperimentsGetStatusResult contains the result from method Experiments.GetStatus.
type ExperimentsGetStatusResult struct {
	ExperimentStatus
}

// ExperimentsListAllResponse contains the response from method Experiments.ListAll.
type ExperimentsListAllResponse struct {
	ExperimentsListAllResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ExperimentsListAllResult contains the result from method Experiments.ListAll.
type ExperimentsListAllResult struct {
	ExperimentListResult
}

// ExperimentsListAllStatusesResponse contains the response from method Experiments.ListAllStatuses.
type ExperimentsListAllStatusesResponse struct {
	ExperimentsListAllStatusesResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ExperimentsListAllStatusesResult contains the result from method Experiments.ListAllStatuses.
type ExperimentsListAllStatusesResult struct {
	ExperimentStatusListResult
}

// ExperimentsListExecutionDetailsResponse contains the response from method Experiments.ListExecutionDetails.
type ExperimentsListExecutionDetailsResponse struct {
	ExperimentsListExecutionDetailsResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ExperimentsListExecutionDetailsResult contains the result from method Experiments.ListExecutionDetails.
type ExperimentsListExecutionDetailsResult struct {
	ExperimentExecutionDetailsListResult
}

// ExperimentsListResponse contains the response from method Experiments.List.
type ExperimentsListResponse struct {
	ExperimentsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ExperimentsListResult contains the result from method Experiments.List.
type ExperimentsListResult struct {
	ExperimentListResult
}

// ExperimentsStartResponse contains the response from method Experiments.Start.
type ExperimentsStartResponse struct {
	ExperimentsStartResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ExperimentsStartResult contains the result from method Experiments.Start.
type ExperimentsStartResult struct {
	ExperimentStartOperationResult
}

// OperationsListAllResponse contains the response from method Operations.ListAll.
type OperationsListAllResponse struct {
	OperationsListAllResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsListAllResult contains the result from method Operations.ListAll.
type OperationsListAllResult struct {
	OperationListResult
}

// TargetTypesGetResponse contains the response from method TargetTypes.Get.
type TargetTypesGetResponse struct {
	TargetTypesGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TargetTypesGetResult contains the result from method TargetTypes.Get.
type TargetTypesGetResult struct {
	TargetType
}

// TargetTypesListResponse contains the response from method TargetTypes.List.
type TargetTypesListResponse struct {
	TargetTypesListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TargetTypesListResult contains the result from method TargetTypes.List.
type TargetTypesListResult struct {
	TargetTypeListResult
}

// TargetsCreateOrUpdateResponse contains the response from method Targets.CreateOrUpdate.
type TargetsCreateOrUpdateResponse struct {
	TargetsCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TargetsCreateOrUpdateResult contains the result from method Targets.CreateOrUpdate.
type TargetsCreateOrUpdateResult struct {
	Target
}

// TargetsDeleteResponse contains the response from method Targets.Delete.
type TargetsDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TargetsGetResponse contains the response from method Targets.Get.
type TargetsGetResponse struct {
	TargetsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TargetsGetResult contains the result from method Targets.Get.
type TargetsGetResult struct {
	Target
}

// TargetsListResponse contains the response from method Targets.List.
type TargetsListResponse struct {
	TargetsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TargetsListResult contains the result from method Targets.List.
type TargetsListResult struct {
	TargetListResult
}
