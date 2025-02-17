package forecast

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalCloudProviderType string

const (
	ExternalCloudProviderTypeExternalBillingAccounts ExternalCloudProviderType = "externalBillingAccounts"
	ExternalCloudProviderTypeExternalSubscriptions   ExternalCloudProviderType = "externalSubscriptions"
)

func PossibleValuesForExternalCloudProviderType() []string {
	return []string{
		string(ExternalCloudProviderTypeExternalBillingAccounts),
		string(ExternalCloudProviderTypeExternalSubscriptions),
	}
}

func parseExternalCloudProviderType(input string) (*ExternalCloudProviderType, error) {
	vals := map[string]ExternalCloudProviderType{
		"externalbillingaccounts": ExternalCloudProviderTypeExternalBillingAccounts,
		"externalsubscriptions":   ExternalCloudProviderTypeExternalSubscriptions,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternalCloudProviderType(input)
	return &out, nil
}

type ForecastOperatorType string

const (
	ForecastOperatorTypeIn ForecastOperatorType = "In"
)

func PossibleValuesForForecastOperatorType() []string {
	return []string{
		string(ForecastOperatorTypeIn),
	}
}

func parseForecastOperatorType(input string) (*ForecastOperatorType, error) {
	vals := map[string]ForecastOperatorType{
		"in": ForecastOperatorTypeIn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ForecastOperatorType(input)
	return &out, nil
}

type ForecastTimeframe string

const (
	ForecastTimeframeCustom ForecastTimeframe = "Custom"
)

func PossibleValuesForForecastTimeframe() []string {
	return []string{
		string(ForecastTimeframeCustom),
	}
}

func parseForecastTimeframe(input string) (*ForecastTimeframe, error) {
	vals := map[string]ForecastTimeframe{
		"custom": ForecastTimeframeCustom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ForecastTimeframe(input)
	return &out, nil
}

type ForecastType string

const (
	ForecastTypeActualCost    ForecastType = "ActualCost"
	ForecastTypeAmortizedCost ForecastType = "AmortizedCost"
	ForecastTypeUsage         ForecastType = "Usage"
)

func PossibleValuesForForecastType() []string {
	return []string{
		string(ForecastTypeActualCost),
		string(ForecastTypeAmortizedCost),
		string(ForecastTypeUsage),
	}
}

func parseForecastType(input string) (*ForecastType, error) {
	vals := map[string]ForecastType{
		"actualcost":    ForecastTypeActualCost,
		"amortizedcost": ForecastTypeAmortizedCost,
		"usage":         ForecastTypeUsage,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ForecastType(input)
	return &out, nil
}

type FunctionName string

const (
	FunctionNameCost          FunctionName = "Cost"
	FunctionNameCostUSD       FunctionName = "CostUSD"
	FunctionNamePreTaxCost    FunctionName = "PreTaxCost"
	FunctionNamePreTaxCostUSD FunctionName = "PreTaxCostUSD"
)

func PossibleValuesForFunctionName() []string {
	return []string{
		string(FunctionNameCost),
		string(FunctionNameCostUSD),
		string(FunctionNamePreTaxCost),
		string(FunctionNamePreTaxCostUSD),
	}
}

func parseFunctionName(input string) (*FunctionName, error) {
	vals := map[string]FunctionName{
		"cost":          FunctionNameCost,
		"costusd":       FunctionNameCostUSD,
		"pretaxcost":    FunctionNamePreTaxCost,
		"pretaxcostusd": FunctionNamePreTaxCostUSD,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FunctionName(input)
	return &out, nil
}

type FunctionType string

const (
	FunctionTypeSum FunctionType = "Sum"
)

func PossibleValuesForFunctionType() []string {
	return []string{
		string(FunctionTypeSum),
	}
}

func parseFunctionType(input string) (*FunctionType, error) {
	vals := map[string]FunctionType{
		"sum": FunctionTypeSum,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FunctionType(input)
	return &out, nil
}

type GranularityType string

const (
	GranularityTypeDaily GranularityType = "Daily"
)

func PossibleValuesForGranularityType() []string {
	return []string{
		string(GranularityTypeDaily),
	}
}

func parseGranularityType(input string) (*GranularityType, error) {
	vals := map[string]GranularityType{
		"daily": GranularityTypeDaily,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GranularityType(input)
	return &out, nil
}
