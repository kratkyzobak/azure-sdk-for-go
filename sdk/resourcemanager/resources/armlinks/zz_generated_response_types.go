//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlinks

import "net/http"

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

// ResourceLinksCreateOrUpdateResponse contains the response from method ResourceLinks.CreateOrUpdate.
type ResourceLinksCreateOrUpdateResponse struct {
	ResourceLinksCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ResourceLinksCreateOrUpdateResult contains the result from method ResourceLinks.CreateOrUpdate.
type ResourceLinksCreateOrUpdateResult struct {
	ResourceLink
}

// ResourceLinksDeleteResponse contains the response from method ResourceLinks.Delete.
type ResourceLinksDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ResourceLinksGetResponse contains the response from method ResourceLinks.Get.
type ResourceLinksGetResponse struct {
	ResourceLinksGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ResourceLinksGetResult contains the result from method ResourceLinks.Get.
type ResourceLinksGetResult struct {
	ResourceLink
}

// ResourceLinksListAtSourceScopeResponse contains the response from method ResourceLinks.ListAtSourceScope.
type ResourceLinksListAtSourceScopeResponse struct {
	ResourceLinksListAtSourceScopeResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ResourceLinksListAtSourceScopeResult contains the result from method ResourceLinks.ListAtSourceScope.
type ResourceLinksListAtSourceScopeResult struct {
	ResourceLinkResult
}

// ResourceLinksListAtSubscriptionResponse contains the response from method ResourceLinks.ListAtSubscription.
type ResourceLinksListAtSubscriptionResponse struct {
	ResourceLinksListAtSubscriptionResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ResourceLinksListAtSubscriptionResult contains the result from method ResourceLinks.ListAtSubscription.
type ResourceLinksListAtSubscriptionResult struct {
	ResourceLinkResult
}