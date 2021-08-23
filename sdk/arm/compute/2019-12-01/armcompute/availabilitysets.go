// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// AvailabilitySetsOperations contains the methods for the AvailabilitySets group.
type AvailabilitySetsOperations interface {
	// CreateOrUpdate - Create or update an availability set.
	CreateOrUpdate(ctx context.Context, resourceGroupName string, availabilitySetName string, parameters AvailabilitySet) (*AvailabilitySetResponse, error)
	// Delete - Delete an availability set.
	Delete(ctx context.Context, resourceGroupName string, availabilitySetName string) (*http.Response, error)
	// Get - Retrieves information about an availability set.
	Get(ctx context.Context, resourceGroupName string, availabilitySetName string) (*AvailabilitySetResponse, error)
	// List - Lists all availability sets in a resource group.
	List(resourceGroupName string) (AvailabilitySetListResultPager, error)
	// ListAvailableSizes - Lists all available virtual machine sizes that can be used to create a new virtual machine in an existing availability set.
	ListAvailableSizes(ctx context.Context, resourceGroupName string, availabilitySetName string) (*VirtualMachineSizeListResultResponse, error)
	// ListBySubscription - Lists all availability sets in a subscription.
	ListBySubscription(availabilitySetsListBySubscriptionOptions *AvailabilitySetsListBySubscriptionOptions) (AvailabilitySetListResultPager, error)
	// Update - Update an availability set.
	Update(ctx context.Context, resourceGroupName string, availabilitySetName string, parameters AvailabilitySetUpdate) (*AvailabilitySetResponse, error)
}

// availabilitySetsOperations implements the AvailabilitySetsOperations interface.
type availabilitySetsOperations struct {
	*Client
	subscriptionID string
}

// CreateOrUpdate - Create or update an availability set.
func (client *availabilitySetsOperations) CreateOrUpdate(ctx context.Context, resourceGroupName string, availabilitySetName string, parameters AvailabilitySet) (*AvailabilitySetResponse, error) {
	req, err := client.createOrUpdateCreateRequest(resourceGroupName, availabilitySetName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.createOrUpdateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *availabilitySetsOperations) createOrUpdateCreateRequest(resourceGroupName string, availabilitySetName string, parameters AvailabilitySet) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{availabilitySetName}", url.PathEscape(availabilitySetName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2019-12-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *availabilitySetsOperations) createOrUpdateHandleResponse(resp *azcore.Response) (*AvailabilitySetResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	result := AvailabilitySetResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.AvailabilitySet)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *availabilitySetsOperations) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return errors.New(resp.Status)
	}
	return errors.New(string(body))
}

// Delete - Delete an availability set.
func (client *availabilitySetsOperations) Delete(ctx context.Context, resourceGroupName string, availabilitySetName string) (*http.Response, error) {
	req, err := client.deleteCreateRequest(resourceGroupName, availabilitySetName)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.deleteHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// deleteCreateRequest creates the Delete request.
func (client *availabilitySetsOperations) deleteCreateRequest(resourceGroupName string, availabilitySetName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{availabilitySetName}", url.PathEscape(availabilitySetName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2019-12-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodDelete, *u)
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *availabilitySetsOperations) deleteHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp.Response, nil
}

// deleteHandleError handles the Delete error response.
func (client *availabilitySetsOperations) deleteHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return errors.New(resp.Status)
	}
	return errors.New(string(body))
}

// Get - Retrieves information about an availability set.
func (client *availabilitySetsOperations) Get(ctx context.Context, resourceGroupName string, availabilitySetName string) (*AvailabilitySetResponse, error) {
	req, err := client.getCreateRequest(resourceGroupName, availabilitySetName)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getCreateRequest creates the Get request.
func (client *availabilitySetsOperations) getCreateRequest(resourceGroupName string, availabilitySetName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{availabilitySetName}", url.PathEscape(availabilitySetName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2019-12-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *availabilitySetsOperations) getHandleResponse(resp *azcore.Response) (*AvailabilitySetResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getHandleError(resp)
	}
	result := AvailabilitySetResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.AvailabilitySet)
}

// getHandleError handles the Get error response.
func (client *availabilitySetsOperations) getHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return errors.New(resp.Status)
	}
	return errors.New(string(body))
}

// List - Lists all availability sets in a resource group.
func (client *availabilitySetsOperations) List(resourceGroupName string) (AvailabilitySetListResultPager, error) {
	req, err := client.listCreateRequest(resourceGroupName)
	if err != nil {
		return nil, err
	}
	return &availabilitySetListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listHandleResponse,
		advancer: func(resp *AvailabilitySetListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.AvailabilitySetListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.AvailabilitySetListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listCreateRequest creates the List request.
func (client *availabilitySetsOperations) listCreateRequest(resourceGroupName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2019-12-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// listHandleResponse handles the List response.
func (client *availabilitySetsOperations) listHandleResponse(resp *azcore.Response) (*AvailabilitySetListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listHandleError(resp)
	}
	result := AvailabilitySetListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.AvailabilitySetListResult)
}

// listHandleError handles the List error response.
func (client *availabilitySetsOperations) listHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return errors.New(resp.Status)
	}
	return errors.New(string(body))
}

// ListAvailableSizes - Lists all available virtual machine sizes that can be used to create a new virtual machine in an existing availability set.
func (client *availabilitySetsOperations) ListAvailableSizes(ctx context.Context, resourceGroupName string, availabilitySetName string) (*VirtualMachineSizeListResultResponse, error) {
	req, err := client.listAvailableSizesCreateRequest(resourceGroupName, availabilitySetName)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.listAvailableSizesHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// listAvailableSizesCreateRequest creates the ListAvailableSizes request.
func (client *availabilitySetsOperations) listAvailableSizesCreateRequest(resourceGroupName string, availabilitySetName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}/vmSizes"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{availabilitySetName}", url.PathEscape(availabilitySetName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2019-12-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// listAvailableSizesHandleResponse handles the ListAvailableSizes response.
func (client *availabilitySetsOperations) listAvailableSizesHandleResponse(resp *azcore.Response) (*VirtualMachineSizeListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listAvailableSizesHandleError(resp)
	}
	result := VirtualMachineSizeListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualMachineSizeListResult)
}

// listAvailableSizesHandleError handles the ListAvailableSizes error response.
func (client *availabilitySetsOperations) listAvailableSizesHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return errors.New(resp.Status)
	}
	return errors.New(string(body))
}

// ListBySubscription - Lists all availability sets in a subscription.
func (client *availabilitySetsOperations) ListBySubscription(availabilitySetsListBySubscriptionOptions *AvailabilitySetsListBySubscriptionOptions) (AvailabilitySetListResultPager, error) {
	req, err := client.listBySubscriptionCreateRequest(availabilitySetsListBySubscriptionOptions)
	if err != nil {
		return nil, err
	}
	return &availabilitySetListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listBySubscriptionHandleResponse,
		advancer: func(resp *AvailabilitySetListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.AvailabilitySetListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.AvailabilitySetListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *availabilitySetsOperations) listBySubscriptionCreateRequest(availabilitySetsListBySubscriptionOptions *AvailabilitySetsListBySubscriptionOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/availabilitySets"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2019-12-01")
	if availabilitySetsListBySubscriptionOptions != nil && availabilitySetsListBySubscriptionOptions.Expand != nil {
		query.Set("$expand", *availabilitySetsListBySubscriptionOptions.Expand)
	}
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *availabilitySetsOperations) listBySubscriptionHandleResponse(resp *azcore.Response) (*AvailabilitySetListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listBySubscriptionHandleError(resp)
	}
	result := AvailabilitySetListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.AvailabilitySetListResult)
}

// listBySubscriptionHandleError handles the ListBySubscription error response.
func (client *availabilitySetsOperations) listBySubscriptionHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return errors.New(resp.Status)
	}
	return errors.New(string(body))
}

// Update - Update an availability set.
func (client *availabilitySetsOperations) Update(ctx context.Context, resourceGroupName string, availabilitySetName string, parameters AvailabilitySetUpdate) (*AvailabilitySetResponse, error) {
	req, err := client.updateCreateRequest(resourceGroupName, availabilitySetName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.updateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// updateCreateRequest creates the Update request.
func (client *availabilitySetsOperations) updateCreateRequest(resourceGroupName string, availabilitySetName string, parameters AvailabilitySetUpdate) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{availabilitySetName}", url.PathEscape(availabilitySetName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2019-12-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPatch, *u)
	return req, req.MarshalAsJSON(parameters)
}

// updateHandleResponse handles the Update response.
func (client *availabilitySetsOperations) updateHandleResponse(resp *azcore.Response) (*AvailabilitySetResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.updateHandleError(resp)
	}
	result := AvailabilitySetResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.AvailabilitySet)
}

// updateHandleError handles the Update error response.
func (client *availabilitySetsOperations) updateHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return errors.New(resp.Status)
	}
	return errors.New(string(body))
}