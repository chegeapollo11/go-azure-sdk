package v2022_08_08

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/healthbot/2022-08-08/healthbots"
)

type Client struct {
	Healthbots *healthbots.HealthbotsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	healthbotsClient := healthbots.NewHealthbotsClientWithBaseURI(endpoint)
	configureAuthFunc(&healthbotsClient.Client)

	return Client{
		Healthbots: &healthbotsClient,
	}
}
