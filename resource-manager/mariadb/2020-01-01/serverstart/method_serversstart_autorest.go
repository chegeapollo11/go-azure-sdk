package serverstart

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServersStartOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// ServersStart ...
func (c ServerStartClient) ServersStart(ctx context.Context, id ServerId) (result ServersStartOperationResponse, err error) {
	req, err := c.preparerForServersStart(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "serverstart.ServerStartClient", "ServersStart", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForServersStart(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "serverstart.ServerStartClient", "ServersStart", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ServersStartThenPoll performs ServersStart then polls until it's completed
func (c ServerStartClient) ServersStartThenPoll(ctx context.Context, id ServerId) error {
	result, err := c.ServersStart(ctx, id)
	if err != nil {
		return fmt.Errorf("performing ServersStart: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after ServersStart: %+v", err)
	}

	return nil
}

// preparerForServersStart prepares the ServersStart request.
func (c ServerStartClient) preparerForServersStart(ctx context.Context, id ServerId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/start", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForServersStart sends the ServersStart request. The method will close the
// http.Response Body if it receives an error.
func (c ServerStartClient) senderForServersStart(ctx context.Context, req *http.Request) (future ServersStartOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
