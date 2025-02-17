package standardoperation

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

type UsagesListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]Quota

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (UsagesListOperationResponse, error)
}

type UsagesListCompleteResult struct {
	Items []Quota
}

func (r UsagesListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r UsagesListOperationResponse) LoadMore(ctx context.Context) (resp UsagesListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// UsagesList ...
func (c StandardOperationClient) UsagesList(ctx context.Context, id LocationId) (resp UsagesListOperationResponse, err error) {
	req, err := c.preparerForUsagesList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "standardoperation.StandardOperationClient", "UsagesList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "standardoperation.StandardOperationClient", "UsagesList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForUsagesList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "standardoperation.StandardOperationClient", "UsagesList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForUsagesList prepares the UsagesList request.
func (c StandardOperationClient) preparerForUsagesList(ctx context.Context, id LocationId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/usages", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForUsagesListWithNextLink prepares the UsagesList request with the given nextLink token.
func (c StandardOperationClient) preparerForUsagesListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForUsagesList handles the response to the UsagesList request. The method always
// closes the http.Response Body.
func (c StandardOperationClient) responderForUsagesList(resp *http.Response) (result UsagesListOperationResponse, err error) {
	type page struct {
		Values   []Quota `json:"value"`
		NextLink *string `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result UsagesListOperationResponse, err error) {
			req, err := c.preparerForUsagesListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "standardoperation.StandardOperationClient", "UsagesList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "standardoperation.StandardOperationClient", "UsagesList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForUsagesList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "standardoperation.StandardOperationClient", "UsagesList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// UsagesListComplete retrieves all of the results into a single object
func (c StandardOperationClient) UsagesListComplete(ctx context.Context, id LocationId) (UsagesListCompleteResult, error) {
	return c.UsagesListCompleteMatchingPredicate(ctx, id, QuotaOperationPredicate{})
}

// UsagesListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c StandardOperationClient) UsagesListCompleteMatchingPredicate(ctx context.Context, id LocationId, predicate QuotaOperationPredicate) (resp UsagesListCompleteResult, err error) {
	items := make([]Quota, 0)

	page, err := c.UsagesList(ctx, id)
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

	out := UsagesListCompleteResult{
		Items: items,
	}
	return out, nil
}
