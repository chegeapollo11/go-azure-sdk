package dnsresolvers

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DnsResolversClient struct {
	Client  autorest.Client
	baseUri string
}

func NewDnsResolversClientWithBaseURI(endpoint string) DnsResolversClient {
	return DnsResolversClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
