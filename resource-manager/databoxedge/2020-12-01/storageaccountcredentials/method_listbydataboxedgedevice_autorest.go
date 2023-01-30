package storageaccountcredentials

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByDataBoxEdgeDeviceOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]StorageAccountCredential

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListByDataBoxEdgeDeviceOperationResponse, error)
}

type ListByDataBoxEdgeDeviceCompleteResult struct {
	Items []StorageAccountCredential
}

func (r ListByDataBoxEdgeDeviceOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListByDataBoxEdgeDeviceOperationResponse) LoadMore(ctx context.Context) (resp ListByDataBoxEdgeDeviceOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListByDataBoxEdgeDevice ...
func (c StorageAccountCredentialsClient) ListByDataBoxEdgeDevice(ctx context.Context, id DataBoxEdgeDeviceId) (resp ListByDataBoxEdgeDeviceOperationResponse, err error) {
	req, err := c.preparerForListByDataBoxEdgeDevice(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storageaccountcredentials.StorageAccountCredentialsClient", "ListByDataBoxEdgeDevice", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "storageaccountcredentials.StorageAccountCredentialsClient", "ListByDataBoxEdgeDevice", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListByDataBoxEdgeDevice(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storageaccountcredentials.StorageAccountCredentialsClient", "ListByDataBoxEdgeDevice", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListByDataBoxEdgeDevice prepares the ListByDataBoxEdgeDevice request.
func (c StorageAccountCredentialsClient) preparerForListByDataBoxEdgeDevice(ctx context.Context, id DataBoxEdgeDeviceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/storageAccountCredentials", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListByDataBoxEdgeDeviceWithNextLink prepares the ListByDataBoxEdgeDevice request with the given nextLink token.
func (c StorageAccountCredentialsClient) preparerForListByDataBoxEdgeDeviceWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
	uri, err := url.Parse(nextLink)
	if err != nil {
		return nil, fmt.Errorf("parsing nextLink %q: %+v", nextLink, err)
	}
	queryParameters := map[string]interface{}{}
	for k, v := range uri.Query() {
		if len(v) == 0 {
			continue
		}
		val := v[0]
		val = autorest.Encode("query", val)
		queryParameters[k] = val
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(uri.Path),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListByDataBoxEdgeDevice handles the response to the ListByDataBoxEdgeDevice request. The method always
// closes the http.Response Body.
func (c StorageAccountCredentialsClient) responderForListByDataBoxEdgeDevice(resp *http.Response) (result ListByDataBoxEdgeDeviceOperationResponse, err error) {
	type page struct {
		Values   []StorageAccountCredential `json:"value"`
		NextLink *string                    `json:"nextLink"`
	}
	var respObj page
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&respObj),
		autorest.ByClosing())
	result.HttpResponse = resp
	result.Model = &respObj.Values
	result.nextLink = respObj.NextLink
	if respObj.NextLink != nil {
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListByDataBoxEdgeDeviceOperationResponse, err error) {
			req, err := c.preparerForListByDataBoxEdgeDeviceWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "storageaccountcredentials.StorageAccountCredentialsClient", "ListByDataBoxEdgeDevice", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "storageaccountcredentials.StorageAccountCredentialsClient", "ListByDataBoxEdgeDevice", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListByDataBoxEdgeDevice(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "storageaccountcredentials.StorageAccountCredentialsClient", "ListByDataBoxEdgeDevice", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListByDataBoxEdgeDeviceComplete retrieves all of the results into a single object
func (c StorageAccountCredentialsClient) ListByDataBoxEdgeDeviceComplete(ctx context.Context, id DataBoxEdgeDeviceId) (ListByDataBoxEdgeDeviceCompleteResult, error) {
	return c.ListByDataBoxEdgeDeviceCompleteMatchingPredicate(ctx, id, StorageAccountCredentialOperationPredicate{})
}

// ListByDataBoxEdgeDeviceCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c StorageAccountCredentialsClient) ListByDataBoxEdgeDeviceCompleteMatchingPredicate(ctx context.Context, id DataBoxEdgeDeviceId, predicate StorageAccountCredentialOperationPredicate) (resp ListByDataBoxEdgeDeviceCompleteResult, err error) {
	items := make([]StorageAccountCredential, 0)

	page, err := c.ListByDataBoxEdgeDevice(ctx, id)
	if err != nil {
		err = fmt.Errorf("loading the initial page: %+v", err)
		return
	}
	if page.Model != nil {
		for _, v := range *page.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	for page.HasMore() {
		page, err = page.LoadMore(ctx)
		if err != nil {
			err = fmt.Errorf("loading the next page: %+v", err)
			return
		}

		if page.Model != nil {
			for _, v := range *page.Model {
				if predicate.Matches(v) {
					items = append(items, v)
				}
			}
		}
	}

	out := ListByDataBoxEdgeDeviceCompleteResult{
		Items: items,
	}
	return out, nil
}
