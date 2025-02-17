package proxy

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoteRenderingAccountsListBySubscriptionOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]RemoteRenderingAccount

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (RemoteRenderingAccountsListBySubscriptionOperationResponse, error)
}

type RemoteRenderingAccountsListBySubscriptionCompleteResult struct {
	Items []RemoteRenderingAccount
}

func (r RemoteRenderingAccountsListBySubscriptionOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r RemoteRenderingAccountsListBySubscriptionOperationResponse) LoadMore(ctx context.Context) (resp RemoteRenderingAccountsListBySubscriptionOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// RemoteRenderingAccountsListBySubscription ...
func (c ProxyClient) RemoteRenderingAccountsListBySubscription(ctx context.Context, id commonids.SubscriptionId) (resp RemoteRenderingAccountsListBySubscriptionOperationResponse, err error) {
	req, err := c.preparerForRemoteRenderingAccountsListBySubscription(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "proxy.ProxyClient", "RemoteRenderingAccountsListBySubscription", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "proxy.ProxyClient", "RemoteRenderingAccountsListBySubscription", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForRemoteRenderingAccountsListBySubscription(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "proxy.ProxyClient", "RemoteRenderingAccountsListBySubscription", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForRemoteRenderingAccountsListBySubscription prepares the RemoteRenderingAccountsListBySubscription request.
func (c ProxyClient) preparerForRemoteRenderingAccountsListBySubscription(ctx context.Context, id commonids.SubscriptionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.MixedReality/remoteRenderingAccounts", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForRemoteRenderingAccountsListBySubscriptionWithNextLink prepares the RemoteRenderingAccountsListBySubscription request with the given nextLink token.
func (c ProxyClient) preparerForRemoteRenderingAccountsListBySubscriptionWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForRemoteRenderingAccountsListBySubscription handles the response to the RemoteRenderingAccountsListBySubscription request. The method always
// closes the http.Response Body.
func (c ProxyClient) responderForRemoteRenderingAccountsListBySubscription(resp *http.Response) (result RemoteRenderingAccountsListBySubscriptionOperationResponse, err error) {
	type page struct {
		Values   []RemoteRenderingAccount `json:"value"`
		NextLink *string                  `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result RemoteRenderingAccountsListBySubscriptionOperationResponse, err error) {
			req, err := c.preparerForRemoteRenderingAccountsListBySubscriptionWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "proxy.ProxyClient", "RemoteRenderingAccountsListBySubscription", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "proxy.ProxyClient", "RemoteRenderingAccountsListBySubscription", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForRemoteRenderingAccountsListBySubscription(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "proxy.ProxyClient", "RemoteRenderingAccountsListBySubscription", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// RemoteRenderingAccountsListBySubscriptionComplete retrieves all of the results into a single object
func (c ProxyClient) RemoteRenderingAccountsListBySubscriptionComplete(ctx context.Context, id commonids.SubscriptionId) (RemoteRenderingAccountsListBySubscriptionCompleteResult, error) {
	return c.RemoteRenderingAccountsListBySubscriptionCompleteMatchingPredicate(ctx, id, RemoteRenderingAccountOperationPredicate{})
}

// RemoteRenderingAccountsListBySubscriptionCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c ProxyClient) RemoteRenderingAccountsListBySubscriptionCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, predicate RemoteRenderingAccountOperationPredicate) (resp RemoteRenderingAccountsListBySubscriptionCompleteResult, err error) {
	items := make([]RemoteRenderingAccount, 0)

	page, err := c.RemoteRenderingAccountsListBySubscription(ctx, id)
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

	out := RemoteRenderingAccountsListBySubscriptionCompleteResult{
		Items: items,
	}
	return out, nil
}
