package openshiftclusters

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListCredentialsOperationResponse struct {
	HttpResponse *http.Response
	Model        *OpenShiftClusterCredentials
}

// ListCredentials ...
func (c OpenShiftClustersClient) ListCredentials(ctx context.Context, id OpenShiftClusterId) (result ListCredentialsOperationResponse, err error) {
	req, err := c.preparerForListCredentials(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "openshiftclusters.OpenShiftClustersClient", "ListCredentials", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "openshiftclusters.OpenShiftClustersClient", "ListCredentials", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListCredentials(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "openshiftclusters.OpenShiftClustersClient", "ListCredentials", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListCredentials prepares the ListCredentials request.
func (c OpenShiftClustersClient) preparerForListCredentials(ctx context.Context, id OpenShiftClusterId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/listCredentials", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListCredentials handles the response to the ListCredentials request. The method always
// closes the http.Response Body.
func (c OpenShiftClustersClient) responderForListCredentials(resp *http.Response) (result ListCredentialsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
