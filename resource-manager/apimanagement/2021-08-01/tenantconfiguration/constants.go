package tenantconfiguration

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AsyncOperationStatus string

const (
	AsyncOperationStatusFailed     AsyncOperationStatus = "Failed"
	AsyncOperationStatusInProgress AsyncOperationStatus = "InProgress"
	AsyncOperationStatusStarted    AsyncOperationStatus = "Started"
	AsyncOperationStatusSucceeded  AsyncOperationStatus = "Succeeded"
)

func PossibleValuesForAsyncOperationStatus() []string {
	return []string{
		string(AsyncOperationStatusFailed),
		string(AsyncOperationStatusInProgress),
		string(AsyncOperationStatusStarted),
		string(AsyncOperationStatusSucceeded),
	}
}

func parseAsyncOperationStatus(input string) (*AsyncOperationStatus, error) {
	vals := map[string]AsyncOperationStatus{
		"failed":     AsyncOperationStatusFailed,
		"inprogress": AsyncOperationStatusInProgress,
		"started":    AsyncOperationStatusStarted,
		"succeeded":  AsyncOperationStatusSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AsyncOperationStatus(input)
	return &out, nil
}
