package workspaces

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EncryptionEntitiesDefinition struct {
	ManagedServices *EncryptionV2 `json:"managedServices,omitempty"`
}
