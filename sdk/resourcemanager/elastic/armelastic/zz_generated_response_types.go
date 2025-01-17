//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by Microsoft (R) AutoRest Code Generator.Changes may cause incorrect behavior and will be lost if the code
// is regenerated.

package armelastic

import (
	"context"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"net/http"
	"time"
)

// DeploymentInfoClientListResponse contains the response from method DeploymentInfoClient.List.
type DeploymentInfoClientListResponse struct {
	DeploymentInfoClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DeploymentInfoClientListResult contains the result from method DeploymentInfoClient.List.
type DeploymentInfoClientListResult struct {
	DeploymentInfoResponse
}

// MonitoredResourcesClientListResponse contains the response from method MonitoredResourcesClient.List.
type MonitoredResourcesClientListResponse struct {
	MonitoredResourcesClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MonitoredResourcesClientListResult contains the result from method MonitoredResourcesClient.List.
type MonitoredResourcesClientListResult struct {
	MonitoredResourceListResponse
}

// MonitorsClientCreatePollerResponse contains the response from method MonitorsClient.Create.
type MonitorsClientCreatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *MonitorsClientCreatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l MonitorsClientCreatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (MonitorsClientCreateResponse, error) {
	respType := MonitorsClientCreateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.MonitorResource)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a MonitorsClientCreatePollerResponse from the provided client and resume token.
func (l *MonitorsClientCreatePollerResponse) Resume(ctx context.Context, client *MonitorsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("MonitorsClient.Create", token, client.pl)
	if err != nil {
		return err
	}
	poller := &MonitorsClientCreatePoller{
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

// MonitorsClientCreateResponse contains the response from method MonitorsClient.Create.
type MonitorsClientCreateResponse struct {
	MonitorsClientCreateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MonitorsClientCreateResult contains the result from method MonitorsClient.Create.
type MonitorsClientCreateResult struct {
	MonitorResource
}

// MonitorsClientDeletePollerResponse contains the response from method MonitorsClient.Delete.
type MonitorsClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *MonitorsClientDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l MonitorsClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (MonitorsClientDeleteResponse, error) {
	respType := MonitorsClientDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a MonitorsClientDeletePollerResponse from the provided client and resume token.
func (l *MonitorsClientDeletePollerResponse) Resume(ctx context.Context, client *MonitorsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("MonitorsClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &MonitorsClientDeletePoller{
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

// MonitorsClientDeleteResponse contains the response from method MonitorsClient.Delete.
type MonitorsClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MonitorsClientGetResponse contains the response from method MonitorsClient.Get.
type MonitorsClientGetResponse struct {
	MonitorsClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MonitorsClientGetResult contains the result from method MonitorsClient.Get.
type MonitorsClientGetResult struct {
	MonitorResource
}

// MonitorsClientListByResourceGroupResponse contains the response from method MonitorsClient.ListByResourceGroup.
type MonitorsClientListByResourceGroupResponse struct {
	MonitorsClientListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MonitorsClientListByResourceGroupResult contains the result from method MonitorsClient.ListByResourceGroup.
type MonitorsClientListByResourceGroupResult struct {
	MonitorResourceListResponse
}

// MonitorsClientListResponse contains the response from method MonitorsClient.List.
type MonitorsClientListResponse struct {
	MonitorsClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MonitorsClientListResult contains the result from method MonitorsClient.List.
type MonitorsClientListResult struct {
	MonitorResourceListResponse
}

// MonitorsClientUpdateResponse contains the response from method MonitorsClient.Update.
type MonitorsClientUpdateResponse struct {
	MonitorsClientUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MonitorsClientUpdateResult contains the result from method MonitorsClient.Update.
type MonitorsClientUpdateResult struct {
	MonitorResource
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationsClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsClientListResult contains the result from method OperationsClient.List.
type OperationsClientListResult struct {
	OperationListResult
}

// TagRulesClientCreateOrUpdateResponse contains the response from method TagRulesClient.CreateOrUpdate.
type TagRulesClientCreateOrUpdateResponse struct {
	TagRulesClientCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TagRulesClientCreateOrUpdateResult contains the result from method TagRulesClient.CreateOrUpdate.
type TagRulesClientCreateOrUpdateResult struct {
	MonitoringTagRules
}

// TagRulesClientDeletePollerResponse contains the response from method TagRulesClient.Delete.
type TagRulesClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *TagRulesClientDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l TagRulesClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (TagRulesClientDeleteResponse, error) {
	respType := TagRulesClientDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a TagRulesClientDeletePollerResponse from the provided client and resume token.
func (l *TagRulesClientDeletePollerResponse) Resume(ctx context.Context, client *TagRulesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("TagRulesClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &TagRulesClientDeletePoller{
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

// TagRulesClientDeleteResponse contains the response from method TagRulesClient.Delete.
type TagRulesClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TagRulesClientGetResponse contains the response from method TagRulesClient.Get.
type TagRulesClientGetResponse struct {
	TagRulesClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TagRulesClientGetResult contains the result from method TagRulesClient.Get.
type TagRulesClientGetResult struct {
	MonitoringTagRules
}

// TagRulesClientListResponse contains the response from method TagRulesClient.List.
type TagRulesClientListResponse struct {
	TagRulesClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TagRulesClientListResult contains the result from method TagRulesClient.List.
type TagRulesClientListResult struct {
	MonitoringTagRulesListResponse
}

// VMCollectionClientUpdateResponse contains the response from method VMCollectionClient.Update.
type VMCollectionClientUpdateResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// VMHostClientListResponse contains the response from method VMHostClient.List.
type VMHostClientListResponse struct {
	VMHostClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// VMHostClientListResult contains the result from method VMHostClient.List.
type VMHostClientListResult struct {
	VMHostListResponse
}

// VMIngestionClientDetailsResponse contains the response from method VMIngestionClient.Details.
type VMIngestionClientDetailsResponse struct {
	VMIngestionClientDetailsResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// VMIngestionClientDetailsResult contains the result from method VMIngestionClient.Details.
type VMIngestionClientDetailsResult struct {
	VMIngestionDetailsResponse
}
