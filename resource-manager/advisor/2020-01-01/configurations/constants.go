package configurations

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CPUThreshold string

const (
	CPUThresholdFive    CPUThreshold = "5"
	CPUThresholdOneFive CPUThreshold = "15"
	CPUThresholdOneZero CPUThreshold = "10"
	CPUThresholdTwoZero CPUThreshold = "20"
)

func PossibleValuesForCPUThreshold() []string {
	return []string{
		string(CPUThresholdFive),
		string(CPUThresholdOneFive),
		string(CPUThresholdOneZero),
		string(CPUThresholdTwoZero),
	}
}

func parseCPUThreshold(input string) (*CPUThreshold, error) {
	vals := map[string]CPUThreshold{
		"5":  CPUThresholdFive,
		"15": CPUThresholdOneFive,
		"10": CPUThresholdOneZero,
		"20": CPUThresholdTwoZero,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CPUThreshold(input)
	return &out, nil
}

type Category string

const (
	CategoryCost                  Category = "Cost"
	CategoryHighAvailability      Category = "HighAvailability"
	CategoryOperationalExcellence Category = "OperationalExcellence"
	CategoryPerformance           Category = "Performance"
	CategorySecurity              Category = "Security"
)

func PossibleValuesForCategory() []string {
	return []string{
		string(CategoryCost),
		string(CategoryHighAvailability),
		string(CategoryOperationalExcellence),
		string(CategoryPerformance),
		string(CategorySecurity),
	}
}

func parseCategory(input string) (*Category, error) {
	vals := map[string]Category{
		"cost":                  CategoryCost,
		"highavailability":      CategoryHighAvailability,
		"operationalexcellence": CategoryOperationalExcellence,
		"performance":           CategoryPerformance,
		"security":              CategorySecurity,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Category(input)
	return &out, nil
}

type DigestConfigState string

const (
	DigestConfigStateActive   DigestConfigState = "Active"
	DigestConfigStateDisabled DigestConfigState = "Disabled"
)

func PossibleValuesForDigestConfigState() []string {
	return []string{
		string(DigestConfigStateActive),
		string(DigestConfigStateDisabled),
	}
}

func parseDigestConfigState(input string) (*DigestConfigState, error) {
	vals := map[string]DigestConfigState{
		"active":   DigestConfigStateActive,
		"disabled": DigestConfigStateDisabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DigestConfigState(input)
	return &out, nil
}
