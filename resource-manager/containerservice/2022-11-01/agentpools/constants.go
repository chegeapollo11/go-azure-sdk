package agentpools

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AgentPoolMode string

const (
	AgentPoolModeSystem AgentPoolMode = "System"
	AgentPoolModeUser   AgentPoolMode = "User"
)

func PossibleValuesForAgentPoolMode() []string {
	return []string{
		string(AgentPoolModeSystem),
		string(AgentPoolModeUser),
	}
}

func parseAgentPoolMode(input string) (*AgentPoolMode, error) {
	vals := map[string]AgentPoolMode{
		"system": AgentPoolModeSystem,
		"user":   AgentPoolModeUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AgentPoolMode(input)
	return &out, nil
}

type AgentPoolType string

const (
	AgentPoolTypeAvailabilitySet         AgentPoolType = "AvailabilitySet"
	AgentPoolTypeVirtualMachineScaleSets AgentPoolType = "VirtualMachineScaleSets"
)

func PossibleValuesForAgentPoolType() []string {
	return []string{
		string(AgentPoolTypeAvailabilitySet),
		string(AgentPoolTypeVirtualMachineScaleSets),
	}
}

func parseAgentPoolType(input string) (*AgentPoolType, error) {
	vals := map[string]AgentPoolType{
		"availabilityset":         AgentPoolTypeAvailabilitySet,
		"virtualmachinescalesets": AgentPoolTypeVirtualMachineScaleSets,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AgentPoolType(input)
	return &out, nil
}

type Code string

const (
	CodeRunning Code = "Running"
	CodeStopped Code = "Stopped"
)

func PossibleValuesForCode() []string {
	return []string{
		string(CodeRunning),
		string(CodeStopped),
	}
}

func parseCode(input string) (*Code, error) {
	vals := map[string]Code{
		"running": CodeRunning,
		"stopped": CodeStopped,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Code(input)
	return &out, nil
}

type GPUInstanceProfile string

const (
	GPUInstanceProfileMIGFourg  GPUInstanceProfile = "MIG4g"
	GPUInstanceProfileMIGOneg   GPUInstanceProfile = "MIG1g"
	GPUInstanceProfileMIGSeveng GPUInstanceProfile = "MIG7g"
	GPUInstanceProfileMIGThreeg GPUInstanceProfile = "MIG3g"
	GPUInstanceProfileMIGTwog   GPUInstanceProfile = "MIG2g"
)

func PossibleValuesForGPUInstanceProfile() []string {
	return []string{
		string(GPUInstanceProfileMIGFourg),
		string(GPUInstanceProfileMIGOneg),
		string(GPUInstanceProfileMIGSeveng),
		string(GPUInstanceProfileMIGThreeg),
		string(GPUInstanceProfileMIGTwog),
	}
}

func parseGPUInstanceProfile(input string) (*GPUInstanceProfile, error) {
	vals := map[string]GPUInstanceProfile{
		"mig4g": GPUInstanceProfileMIGFourg,
		"mig1g": GPUInstanceProfileMIGOneg,
		"mig7g": GPUInstanceProfileMIGSeveng,
		"mig3g": GPUInstanceProfileMIGThreeg,
		"mig2g": GPUInstanceProfileMIGTwog,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GPUInstanceProfile(input)
	return &out, nil
}

type KubeletDiskType string

const (
	KubeletDiskTypeOS        KubeletDiskType = "OS"
	KubeletDiskTypeTemporary KubeletDiskType = "Temporary"
)

func PossibleValuesForKubeletDiskType() []string {
	return []string{
		string(KubeletDiskTypeOS),
		string(KubeletDiskTypeTemporary),
	}
}

func parseKubeletDiskType(input string) (*KubeletDiskType, error) {
	vals := map[string]KubeletDiskType{
		"os":        KubeletDiskTypeOS,
		"temporary": KubeletDiskTypeTemporary,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KubeletDiskType(input)
	return &out, nil
}

type OSDiskType string

const (
	OSDiskTypeEphemeral OSDiskType = "Ephemeral"
	OSDiskTypeManaged   OSDiskType = "Managed"
)

func PossibleValuesForOSDiskType() []string {
	return []string{
		string(OSDiskTypeEphemeral),
		string(OSDiskTypeManaged),
	}
}

func parseOSDiskType(input string) (*OSDiskType, error) {
	vals := map[string]OSDiskType{
		"ephemeral": OSDiskTypeEphemeral,
		"managed":   OSDiskTypeManaged,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OSDiskType(input)
	return &out, nil
}

type OSSKU string

const (
	OSSKUCBLMariner            OSSKU = "CBLMariner"
	OSSKUUbuntu                OSSKU = "Ubuntu"
	OSSKUWindowsTwoZeroOneNine OSSKU = "Windows2019"
	OSSKUWindowsTwoZeroTwoTwo  OSSKU = "Windows2022"
)

func PossibleValuesForOSSKU() []string {
	return []string{
		string(OSSKUCBLMariner),
		string(OSSKUUbuntu),
		string(OSSKUWindowsTwoZeroOneNine),
		string(OSSKUWindowsTwoZeroTwoTwo),
	}
}

func parseOSSKU(input string) (*OSSKU, error) {
	vals := map[string]OSSKU{
		"cblmariner":  OSSKUCBLMariner,
		"ubuntu":      OSSKUUbuntu,
		"windows2019": OSSKUWindowsTwoZeroOneNine,
		"windows2022": OSSKUWindowsTwoZeroTwoTwo,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OSSKU(input)
	return &out, nil
}

type OSType string

const (
	OSTypeLinux   OSType = "Linux"
	OSTypeWindows OSType = "Windows"
)

func PossibleValuesForOSType() []string {
	return []string{
		string(OSTypeLinux),
		string(OSTypeWindows),
	}
}

func parseOSType(input string) (*OSType, error) {
	vals := map[string]OSType{
		"linux":   OSTypeLinux,
		"windows": OSTypeWindows,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OSType(input)
	return &out, nil
}

type ScaleDownMode string

const (
	ScaleDownModeDeallocate ScaleDownMode = "Deallocate"
	ScaleDownModeDelete     ScaleDownMode = "Delete"
)

func PossibleValuesForScaleDownMode() []string {
	return []string{
		string(ScaleDownModeDeallocate),
		string(ScaleDownModeDelete),
	}
}

func parseScaleDownMode(input string) (*ScaleDownMode, error) {
	vals := map[string]ScaleDownMode{
		"deallocate": ScaleDownModeDeallocate,
		"delete":     ScaleDownModeDelete,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ScaleDownMode(input)
	return &out, nil
}

type ScaleSetEvictionPolicy string

const (
	ScaleSetEvictionPolicyDeallocate ScaleSetEvictionPolicy = "Deallocate"
	ScaleSetEvictionPolicyDelete     ScaleSetEvictionPolicy = "Delete"
)

func PossibleValuesForScaleSetEvictionPolicy() []string {
	return []string{
		string(ScaleSetEvictionPolicyDeallocate),
		string(ScaleSetEvictionPolicyDelete),
	}
}

func parseScaleSetEvictionPolicy(input string) (*ScaleSetEvictionPolicy, error) {
	vals := map[string]ScaleSetEvictionPolicy{
		"deallocate": ScaleSetEvictionPolicyDeallocate,
		"delete":     ScaleSetEvictionPolicyDelete,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ScaleSetEvictionPolicy(input)
	return &out, nil
}

type ScaleSetPriority string

const (
	ScaleSetPriorityRegular ScaleSetPriority = "Regular"
	ScaleSetPrioritySpot    ScaleSetPriority = "Spot"
)

func PossibleValuesForScaleSetPriority() []string {
	return []string{
		string(ScaleSetPriorityRegular),
		string(ScaleSetPrioritySpot),
	}
}

func parseScaleSetPriority(input string) (*ScaleSetPriority, error) {
	vals := map[string]ScaleSetPriority{
		"regular": ScaleSetPriorityRegular,
		"spot":    ScaleSetPrioritySpot,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ScaleSetPriority(input)
	return &out, nil
}

type WorkloadRuntime string

const (
	WorkloadRuntimeOCIContainer WorkloadRuntime = "OCIContainer"
	WorkloadRuntimeWasmWasi     WorkloadRuntime = "WasmWasi"
)

func PossibleValuesForWorkloadRuntime() []string {
	return []string{
		string(WorkloadRuntimeOCIContainer),
		string(WorkloadRuntimeWasmWasi),
	}
}

func parseWorkloadRuntime(input string) (*WorkloadRuntime, error) {
	vals := map[string]WorkloadRuntime{
		"ocicontainer": WorkloadRuntimeOCIContainer,
		"wasmwasi":     WorkloadRuntimeWasmWasi,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkloadRuntime(input)
	return &out, nil
}
