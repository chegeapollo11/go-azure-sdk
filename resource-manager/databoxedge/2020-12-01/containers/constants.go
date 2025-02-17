package containers

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureContainerDataFormat string

const (
	AzureContainerDataFormatAzureFile AzureContainerDataFormat = "AzureFile"
	AzureContainerDataFormatBlockBlob AzureContainerDataFormat = "BlockBlob"
	AzureContainerDataFormatPageBlob  AzureContainerDataFormat = "PageBlob"
)

func PossibleValuesForAzureContainerDataFormat() []string {
	return []string{
		string(AzureContainerDataFormatAzureFile),
		string(AzureContainerDataFormatBlockBlob),
		string(AzureContainerDataFormatPageBlob),
	}
}

func parseAzureContainerDataFormat(input string) (*AzureContainerDataFormat, error) {
	vals := map[string]AzureContainerDataFormat{
		"azurefile": AzureContainerDataFormatAzureFile,
		"blockblob": AzureContainerDataFormatBlockBlob,
		"pageblob":  AzureContainerDataFormatPageBlob,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AzureContainerDataFormat(input)
	return &out, nil
}

type ContainerStatus string

const (
	ContainerStatusNeedsAttention ContainerStatus = "NeedsAttention"
	ContainerStatusOK             ContainerStatus = "OK"
	ContainerStatusOffline        ContainerStatus = "Offline"
	ContainerStatusUnknown        ContainerStatus = "Unknown"
	ContainerStatusUpdating       ContainerStatus = "Updating"
)

func PossibleValuesForContainerStatus() []string {
	return []string{
		string(ContainerStatusNeedsAttention),
		string(ContainerStatusOK),
		string(ContainerStatusOffline),
		string(ContainerStatusUnknown),
		string(ContainerStatusUpdating),
	}
}

func parseContainerStatus(input string) (*ContainerStatus, error) {
	vals := map[string]ContainerStatus{
		"needsattention": ContainerStatusNeedsAttention,
		"ok":             ContainerStatusOK,
		"offline":        ContainerStatusOffline,
		"unknown":        ContainerStatusUnknown,
		"updating":       ContainerStatusUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ContainerStatus(input)
	return &out, nil
}
