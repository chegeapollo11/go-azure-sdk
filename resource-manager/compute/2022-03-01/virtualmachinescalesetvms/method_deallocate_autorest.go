package virtualmachinescalesetvms

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

type DeallocateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// Deallocate ...
func (c VirtualMachineScaleSetVMsClient) Deallocate(ctx context.Context, id VirtualMachineId) (result DeallocateOperationResponse, err error) {
	req, err := c.preparerForDeallocate(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachinescalesetvms.VirtualMachineScaleSetVMsClient", "Deallocate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForDeallocate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachinescalesetvms.VirtualMachineScaleSetVMsClient", "Deallocate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// DeallocateThenPoll performs Deallocate then polls until it's completed
func (c VirtualMachineScaleSetVMsClient) DeallocateThenPoll(ctx context.Context, id VirtualMachineId) error {
	result, err := c.Deallocate(ctx, id)
	if err != nil {
		return fmt.Errorf("performing Deallocate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after Deallocate: %+v", err)
	}

	return nil
}

// preparerForDeallocate prepares the Deallocate request.
func (c VirtualMachineScaleSetVMsClient) preparerForDeallocate(ctx context.Context, id VirtualMachineId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/deallocate", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForDeallocate sends the Deallocate request. The method will close the
// http.Response Body if it receives an error.
func (c VirtualMachineScaleSetVMsClient) senderForDeallocate(ctx context.Context, req *http.Request) (future DeallocateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
