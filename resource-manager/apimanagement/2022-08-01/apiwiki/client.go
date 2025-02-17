package apiwiki

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiWikiClient struct {
	Client  autorest.Client
	baseUri string
}

func NewApiWikiClientWithBaseURI(endpoint string) ApiWikiClient {
	return ApiWikiClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
