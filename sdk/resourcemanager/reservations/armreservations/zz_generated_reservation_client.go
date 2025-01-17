//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armreservations

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
	"strconv"
	"strings"
)

// ReservationClient contains the methods for the Reservation group.
// Don't use this type directly, use NewReservationClient() instead.
type ReservationClient struct {
	host string
	pl   runtime.Pipeline
}

// NewReservationClient creates a new instance of ReservationClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewReservationClient(credential azcore.TokenCredential, options *arm.ClientOptions) *ReservationClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &ReservationClient{
		host: string(cp.Endpoint),
		pl:   armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginAvailableScopes - Get Available Scopes for Reservation.
// If the operation fails it returns an *azcore.ResponseError type.
// reservationOrderID - Order Id of the reservation
// reservationID - Id of the Reservation Item
// options - ReservationClientBeginAvailableScopesOptions contains the optional parameters for the ReservationClient.BeginAvailableScopes
// method.
func (client *ReservationClient) BeginAvailableScopes(ctx context.Context, reservationOrderID string, reservationID string, body AvailableScopeRequest, options *ReservationClientBeginAvailableScopesOptions) (ReservationClientAvailableScopesPollerResponse, error) {
	resp, err := client.availableScopes(ctx, reservationOrderID, reservationID, body, options)
	if err != nil {
		return ReservationClientAvailableScopesPollerResponse{}, err
	}
	result := ReservationClientAvailableScopesPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ReservationClient.AvailableScopes", "", resp, client.pl)
	if err != nil {
		return ReservationClientAvailableScopesPollerResponse{}, err
	}
	result.Poller = &ReservationClientAvailableScopesPoller{
		pt: pt,
	}
	return result, nil
}

// AvailableScopes - Get Available Scopes for Reservation.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ReservationClient) availableScopes(ctx context.Context, reservationOrderID string, reservationID string, body AvailableScopeRequest, options *ReservationClientBeginAvailableScopesOptions) (*http.Response, error) {
	req, err := client.availableScopesCreateRequest(ctx, reservationOrderID, reservationID, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// availableScopesCreateRequest creates the AvailableScopes request.
func (client *ReservationClient) availableScopesCreateRequest(ctx context.Context, reservationOrderID string, reservationID string, body AvailableScopeRequest, options *ReservationClientBeginAvailableScopesOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Capacity/reservationOrders/{reservationOrderId}/reservations/{reservationId}/availableScopes"
	if reservationOrderID == "" {
		return nil, errors.New("parameter reservationOrderID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reservationOrderId}", url.PathEscape(reservationOrderID))
	if reservationID == "" {
		return nil, errors.New("parameter reservationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reservationId}", url.PathEscape(reservationID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

// Get - Get specific Reservation details.
// If the operation fails it returns an *azcore.ResponseError type.
// reservationID - Id of the Reservation Item
// reservationOrderID - Order Id of the reservation
// options - ReservationClientGetOptions contains the optional parameters for the ReservationClient.Get method.
func (client *ReservationClient) Get(ctx context.Context, reservationID string, reservationOrderID string, options *ReservationClientGetOptions) (ReservationClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, reservationID, reservationOrderID, options)
	if err != nil {
		return ReservationClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ReservationClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ReservationClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ReservationClient) getCreateRequest(ctx context.Context, reservationID string, reservationOrderID string, options *ReservationClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Capacity/reservationOrders/{reservationOrderId}/reservations/{reservationId}"
	if reservationID == "" {
		return nil, errors.New("parameter reservationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reservationId}", url.PathEscape(reservationID))
	if reservationOrderID == "" {
		return nil, errors.New("parameter reservationOrderID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reservationOrderId}", url.PathEscape(reservationOrderID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ReservationClient) getHandleResponse(resp *http.Response) (ReservationClientGetResponse, error) {
	result := ReservationClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReservationResponse); err != nil {
		return ReservationClientGetResponse{}, err
	}
	return result, nil
}

// List - List Reservations within a single ReservationOrder.
// If the operation fails it returns an *azcore.ResponseError type.
// reservationOrderID - Order Id of the reservation
// options - ReservationClientListOptions contains the optional parameters for the ReservationClient.List method.
func (client *ReservationClient) List(reservationOrderID string, options *ReservationClientListOptions) *ReservationClientListPager {
	return &ReservationClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, reservationOrderID, options)
		},
		advancer: func(ctx context.Context, resp ReservationClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ReservationList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ReservationClient) listCreateRequest(ctx context.Context, reservationOrderID string, options *ReservationClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Capacity/reservationOrders/{reservationOrderId}/reservations"
	if reservationOrderID == "" {
		return nil, errors.New("parameter reservationOrderID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reservationOrderId}", url.PathEscape(reservationOrderID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ReservationClient) listHandleResponse(resp *http.Response) (ReservationClientListResponse, error) {
	result := ReservationClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReservationList); err != nil {
		return ReservationClientListResponse{}, err
	}
	return result, nil
}

// ListAll - List the reservations and the roll up counts of reservations group by provisioning states that the user has access
// to in the current tenant.
// If the operation fails it returns an *azcore.ResponseError type.
// options - ReservationClientListAllOptions contains the optional parameters for the ReservationClient.ListAll method.
func (client *ReservationClient) ListAll(options *ReservationClientListAllOptions) *ReservationClientListAllPager {
	return &ReservationClientListAllPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listAllCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ReservationClientListAllResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ListResult.NextLink)
		},
	}
}

// listAllCreateRequest creates the ListAll request.
func (client *ReservationClient) listAllCreateRequest(ctx context.Context, options *ReservationClientListAllOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Capacity/reservations"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	if options != nil && options.RefreshSummary != nil {
		reqQP.Set("refreshSummary", *options.RefreshSummary)
	}
	if options != nil && options.Skiptoken != nil {
		reqQP.Set("$skiptoken", strconv.FormatFloat(float64(*options.Skiptoken), 'f', -1, 32))
	}
	if options != nil && options.SelectedState != nil {
		reqQP.Set("selectedState", *options.SelectedState)
	}
	if options != nil && options.Take != nil {
		reqQP.Set("take", strconv.FormatFloat(float64(*options.Take), 'f', -1, 32))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *ReservationClient) listAllHandleResponse(resp *http.Response) (ReservationClientListAllResponse, error) {
	result := ReservationClientListAllResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListResult); err != nil {
		return ReservationClientListAllResponse{}, err
	}
	return result, nil
}

// ListRevisions - List of all the revisions for the Reservation.
// If the operation fails it returns an *azcore.ResponseError type.
// reservationID - Id of the Reservation Item
// reservationOrderID - Order Id of the reservation
// options - ReservationClientListRevisionsOptions contains the optional parameters for the ReservationClient.ListRevisions
// method.
func (client *ReservationClient) ListRevisions(reservationID string, reservationOrderID string, options *ReservationClientListRevisionsOptions) *ReservationClientListRevisionsPager {
	return &ReservationClientListRevisionsPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listRevisionsCreateRequest(ctx, reservationID, reservationOrderID, options)
		},
		advancer: func(ctx context.Context, resp ReservationClientListRevisionsResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ReservationList.NextLink)
		},
	}
}

// listRevisionsCreateRequest creates the ListRevisions request.
func (client *ReservationClient) listRevisionsCreateRequest(ctx context.Context, reservationID string, reservationOrderID string, options *ReservationClientListRevisionsOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Capacity/reservationOrders/{reservationOrderId}/reservations/{reservationId}/revisions"
	if reservationID == "" {
		return nil, errors.New("parameter reservationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reservationId}", url.PathEscape(reservationID))
	if reservationOrderID == "" {
		return nil, errors.New("parameter reservationOrderID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reservationOrderId}", url.PathEscape(reservationOrderID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listRevisionsHandleResponse handles the ListRevisions response.
func (client *ReservationClient) listRevisionsHandleResponse(resp *http.Response) (ReservationClientListRevisionsResponse, error) {
	result := ReservationClientListRevisionsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReservationList); err != nil {
		return ReservationClientListRevisionsResponse{}, err
	}
	return result, nil
}

// BeginMerge - Merge the specified Reservations into a new Reservation. The two Reservations being merged must have same
// properties.
// If the operation fails it returns an *azcore.ResponseError type.
// reservationOrderID - Order Id of the reservation
// body - Information needed for commercial request for a reservation
// options - ReservationClientBeginMergeOptions contains the optional parameters for the ReservationClient.BeginMerge method.
func (client *ReservationClient) BeginMerge(ctx context.Context, reservationOrderID string, body MergeRequest, options *ReservationClientBeginMergeOptions) (ReservationClientMergePollerResponse, error) {
	resp, err := client.merge(ctx, reservationOrderID, body, options)
	if err != nil {
		return ReservationClientMergePollerResponse{}, err
	}
	result := ReservationClientMergePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ReservationClient.Merge", "location", resp, client.pl)
	if err != nil {
		return ReservationClientMergePollerResponse{}, err
	}
	result.Poller = &ReservationClientMergePoller{
		pt: pt,
	}
	return result, nil
}

// Merge - Merge the specified Reservations into a new Reservation. The two Reservations being merged must have same properties.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ReservationClient) merge(ctx context.Context, reservationOrderID string, body MergeRequest, options *ReservationClientBeginMergeOptions) (*http.Response, error) {
	req, err := client.mergeCreateRequest(ctx, reservationOrderID, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// mergeCreateRequest creates the Merge request.
func (client *ReservationClient) mergeCreateRequest(ctx context.Context, reservationOrderID string, body MergeRequest, options *ReservationClientBeginMergeOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Capacity/reservationOrders/{reservationOrderId}/merge"
	if reservationOrderID == "" {
		return nil, errors.New("parameter reservationOrderID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reservationOrderId}", url.PathEscape(reservationOrderID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

// BeginSplit - Split a Reservation into two Reservations with specified quantity distribution.
// If the operation fails it returns an *azcore.ResponseError type.
// reservationOrderID - Order Id of the reservation
// body - Information needed to Split a reservation item
// options - ReservationClientBeginSplitOptions contains the optional parameters for the ReservationClient.BeginSplit method.
func (client *ReservationClient) BeginSplit(ctx context.Context, reservationOrderID string, body SplitRequest, options *ReservationClientBeginSplitOptions) (ReservationClientSplitPollerResponse, error) {
	resp, err := client.split(ctx, reservationOrderID, body, options)
	if err != nil {
		return ReservationClientSplitPollerResponse{}, err
	}
	result := ReservationClientSplitPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ReservationClient.Split", "location", resp, client.pl)
	if err != nil {
		return ReservationClientSplitPollerResponse{}, err
	}
	result.Poller = &ReservationClientSplitPoller{
		pt: pt,
	}
	return result, nil
}

// Split - Split a Reservation into two Reservations with specified quantity distribution.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ReservationClient) split(ctx context.Context, reservationOrderID string, body SplitRequest, options *ReservationClientBeginSplitOptions) (*http.Response, error) {
	req, err := client.splitCreateRequest(ctx, reservationOrderID, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// splitCreateRequest creates the Split request.
func (client *ReservationClient) splitCreateRequest(ctx context.Context, reservationOrderID string, body SplitRequest, options *ReservationClientBeginSplitOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Capacity/reservationOrders/{reservationOrderId}/split"
	if reservationOrderID == "" {
		return nil, errors.New("parameter reservationOrderID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reservationOrderId}", url.PathEscape(reservationOrderID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

// BeginUpdate - Updates the applied scopes of the Reservation.
// If the operation fails it returns an *azcore.ResponseError type.
// reservationOrderID - Order Id of the reservation
// reservationID - Id of the Reservation Item
// parameters - Information needed to patch a reservation item
// options - ReservationClientBeginUpdateOptions contains the optional parameters for the ReservationClient.BeginUpdate method.
func (client *ReservationClient) BeginUpdate(ctx context.Context, reservationOrderID string, reservationID string, parameters Patch, options *ReservationClientBeginUpdateOptions) (ReservationClientUpdatePollerResponse, error) {
	resp, err := client.update(ctx, reservationOrderID, reservationID, parameters, options)
	if err != nil {
		return ReservationClientUpdatePollerResponse{}, err
	}
	result := ReservationClientUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ReservationClient.Update", "", resp, client.pl)
	if err != nil {
		return ReservationClientUpdatePollerResponse{}, err
	}
	result.Poller = &ReservationClientUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Updates the applied scopes of the Reservation.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ReservationClient) update(ctx context.Context, reservationOrderID string, reservationID string, parameters Patch, options *ReservationClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, reservationOrderID, reservationID, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *ReservationClient) updateCreateRequest(ctx context.Context, reservationOrderID string, reservationID string, parameters Patch, options *ReservationClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Capacity/reservationOrders/{reservationOrderId}/reservations/{reservationId}"
	if reservationOrderID == "" {
		return nil, errors.New("parameter reservationOrderID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reservationOrderId}", url.PathEscape(reservationOrderID))
	if reservationID == "" {
		return nil, errors.New("parameter reservationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reservationId}", url.PathEscape(reservationID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}
