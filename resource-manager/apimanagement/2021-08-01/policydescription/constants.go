package policydescription

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyScopeContract string

const (
	PolicyScopeContractAll       PolicyScopeContract = "All"
	PolicyScopeContractApi       PolicyScopeContract = "Api"
	PolicyScopeContractOperation PolicyScopeContract = "Operation"
	PolicyScopeContractProduct   PolicyScopeContract = "Product"
	PolicyScopeContractTenant    PolicyScopeContract = "Tenant"
)

func PossibleValuesForPolicyScopeContract() []string {
	return []string{
		string(PolicyScopeContractAll),
		string(PolicyScopeContractApi),
		string(PolicyScopeContractOperation),
		string(PolicyScopeContractProduct),
		string(PolicyScopeContractTenant),
	}
}

func parsePolicyScopeContract(input string) (*PolicyScopeContract, error) {
	vals := map[string]PolicyScopeContract{
		"all":       PolicyScopeContractAll,
		"api":       PolicyScopeContractApi,
		"operation": PolicyScopeContractOperation,
		"product":   PolicyScopeContractProduct,
		"tenant":    PolicyScopeContractTenant,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PolicyScopeContract(input)
	return &out, nil
}
