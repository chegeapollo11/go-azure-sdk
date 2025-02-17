package customimages

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomImageOsType string

const (
	CustomImageOsTypeLinux   CustomImageOsType = "Linux"
	CustomImageOsTypeNone    CustomImageOsType = "None"
	CustomImageOsTypeWindows CustomImageOsType = "Windows"
)

func PossibleValuesForCustomImageOsType() []string {
	return []string{
		string(CustomImageOsTypeLinux),
		string(CustomImageOsTypeNone),
		string(CustomImageOsTypeWindows),
	}
}

func parseCustomImageOsType(input string) (*CustomImageOsType, error) {
	vals := map[string]CustomImageOsType{
		"linux":   CustomImageOsTypeLinux,
		"none":    CustomImageOsTypeNone,
		"windows": CustomImageOsTypeWindows,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CustomImageOsType(input)
	return &out, nil
}

type LinuxOsState string

const (
	LinuxOsStateDeprovisionApplied   LinuxOsState = "DeprovisionApplied"
	LinuxOsStateDeprovisionRequested LinuxOsState = "DeprovisionRequested"
	LinuxOsStateNonDeprovisioned     LinuxOsState = "NonDeprovisioned"
)

func PossibleValuesForLinuxOsState() []string {
	return []string{
		string(LinuxOsStateDeprovisionApplied),
		string(LinuxOsStateDeprovisionRequested),
		string(LinuxOsStateNonDeprovisioned),
	}
}

func parseLinuxOsState(input string) (*LinuxOsState, error) {
	vals := map[string]LinuxOsState{
		"deprovisionapplied":   LinuxOsStateDeprovisionApplied,
		"deprovisionrequested": LinuxOsStateDeprovisionRequested,
		"nondeprovisioned":     LinuxOsStateNonDeprovisioned,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LinuxOsState(input)
	return &out, nil
}

type StorageType string

const (
	StorageTypePremium     StorageType = "Premium"
	StorageTypeStandard    StorageType = "Standard"
	StorageTypeStandardSSD StorageType = "StandardSSD"
)

func PossibleValuesForStorageType() []string {
	return []string{
		string(StorageTypePremium),
		string(StorageTypeStandard),
		string(StorageTypeStandardSSD),
	}
}

func parseStorageType(input string) (*StorageType, error) {
	vals := map[string]StorageType{
		"premium":     StorageTypePremium,
		"standard":    StorageTypeStandard,
		"standardssd": StorageTypeStandardSSD,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StorageType(input)
	return &out, nil
}

type WindowsOsState string

const (
	WindowsOsStateNonSysprepped    WindowsOsState = "NonSysprepped"
	WindowsOsStateSysprepApplied   WindowsOsState = "SysprepApplied"
	WindowsOsStateSysprepRequested WindowsOsState = "SysprepRequested"
)

func PossibleValuesForWindowsOsState() []string {
	return []string{
		string(WindowsOsStateNonSysprepped),
		string(WindowsOsStateSysprepApplied),
		string(WindowsOsStateSysprepRequested),
	}
}

func parseWindowsOsState(input string) (*WindowsOsState, error) {
	vals := map[string]WindowsOsState{
		"nonsysprepped":    WindowsOsStateNonSysprepped,
		"sysprepapplied":   WindowsOsStateSysprepApplied,
		"syspreprequested": WindowsOsStateSysprepRequested,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsOsState(input)
	return &out, nil
}
