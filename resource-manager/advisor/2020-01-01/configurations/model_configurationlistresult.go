package configurations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigurationListResult struct {
	NextLink *string       `json:"nextLink,omitempty"`
	Value    *[]ConfigData `json:"value,omitempty"`
}
