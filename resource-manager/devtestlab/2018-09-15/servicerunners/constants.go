package servicerunners

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedIdentityType string

const (
	ManagedIdentityTypeNone                       ManagedIdentityType = "None"
	ManagedIdentityTypeSystemAssigned             ManagedIdentityType = "SystemAssigned"
	ManagedIdentityTypeSystemAssignedUserAssigned ManagedIdentityType = "SystemAssigned,UserAssigned"
	ManagedIdentityTypeUserAssigned               ManagedIdentityType = "UserAssigned"
)

func PossibleValuesForManagedIdentityType() []string {
	return []string{
		string(ManagedIdentityTypeNone),
		string(ManagedIdentityTypeSystemAssigned),
		string(ManagedIdentityTypeSystemAssignedUserAssigned),
		string(ManagedIdentityTypeUserAssigned),
	}
}

func parseManagedIdentityType(input string) (*ManagedIdentityType, error) {
	vals := map[string]ManagedIdentityType{
		"none":                        ManagedIdentityTypeNone,
		"systemassigned":              ManagedIdentityTypeSystemAssigned,
		"systemassigned,userassigned": ManagedIdentityTypeSystemAssignedUserAssigned,
		"userassigned":                ManagedIdentityTypeUserAssigned,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedIdentityType(input)
	return &out, nil
}
