package subscriptions

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationMode string

const (
	AuthenticationModeConnectionString AuthenticationMode = "ConnectionString"
	AuthenticationModeMsi              AuthenticationMode = "Msi"
	AuthenticationModeUserToken        AuthenticationMode = "UserToken"
)

func PossibleValuesForAuthenticationMode() []string {
	return []string{
		string(AuthenticationModeConnectionString),
		string(AuthenticationModeMsi),
		string(AuthenticationModeUserToken),
	}
}

func parseAuthenticationMode(input string) (*AuthenticationMode, error) {
	vals := map[string]AuthenticationMode{
		"connectionstring": AuthenticationModeConnectionString,
		"msi":              AuthenticationModeMsi,
		"usertoken":        AuthenticationModeUserToken,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationMode(input)
	return &out, nil
}

type BlobWriteMode string

const (
	BlobWriteModeAppend BlobWriteMode = "Append"
	BlobWriteModeOnce   BlobWriteMode = "Once"
)

func PossibleValuesForBlobWriteMode() []string {
	return []string{
		string(BlobWriteModeAppend),
		string(BlobWriteModeOnce),
	}
}

func parseBlobWriteMode(input string) (*BlobWriteMode, error) {
	vals := map[string]BlobWriteMode{
		"append": BlobWriteModeAppend,
		"once":   BlobWriteModeOnce,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BlobWriteMode(input)
	return &out, nil
}

type CompatibilityLevel string

const (
	CompatibilityLevelOnePointTwo  CompatibilityLevel = "1.2"
	CompatibilityLevelOnePointZero CompatibilityLevel = "1.0"
)

func PossibleValuesForCompatibilityLevel() []string {
	return []string{
		string(CompatibilityLevelOnePointTwo),
		string(CompatibilityLevelOnePointZero),
	}
}

func parseCompatibilityLevel(input string) (*CompatibilityLevel, error) {
	vals := map[string]CompatibilityLevel{
		"1.2": CompatibilityLevelOnePointTwo,
		"1.0": CompatibilityLevelOnePointZero,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CompatibilityLevel(input)
	return &out, nil
}

type CompressionType string

const (
	CompressionTypeDeflate CompressionType = "Deflate"
	CompressionTypeGZip    CompressionType = "GZip"
	CompressionTypeNone    CompressionType = "None"
)

func PossibleValuesForCompressionType() []string {
	return []string{
		string(CompressionTypeDeflate),
		string(CompressionTypeGZip),
		string(CompressionTypeNone),
	}
}

func parseCompressionType(input string) (*CompressionType, error) {
	vals := map[string]CompressionType{
		"deflate": CompressionTypeDeflate,
		"gzip":    CompressionTypeGZip,
		"none":    CompressionTypeNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CompressionType(input)
	return &out, nil
}

type ContentStoragePolicy string

const (
	ContentStoragePolicyJobStorageAccount ContentStoragePolicy = "JobStorageAccount"
	ContentStoragePolicySystemAccount     ContentStoragePolicy = "SystemAccount"
)

func PossibleValuesForContentStoragePolicy() []string {
	return []string{
		string(ContentStoragePolicyJobStorageAccount),
		string(ContentStoragePolicySystemAccount),
	}
}

func parseContentStoragePolicy(input string) (*ContentStoragePolicy, error) {
	vals := map[string]ContentStoragePolicy{
		"jobstorageaccount": ContentStoragePolicyJobStorageAccount,
		"systemaccount":     ContentStoragePolicySystemAccount,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ContentStoragePolicy(input)
	return &out, nil
}

type Encoding string

const (
	EncodingUTFEight Encoding = "UTF8"
)

func PossibleValuesForEncoding() []string {
	return []string{
		string(EncodingUTFEight),
	}
}

func parseEncoding(input string) (*Encoding, error) {
	vals := map[string]Encoding{
		"utf8": EncodingUTFEight,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Encoding(input)
	return &out, nil
}

type EventGridEventSchemaType string

const (
	EventGridEventSchemaTypeCloudEventSchema     EventGridEventSchemaType = "CloudEventSchema"
	EventGridEventSchemaTypeEventGridEventSchema EventGridEventSchemaType = "EventGridEventSchema"
)

func PossibleValuesForEventGridEventSchemaType() []string {
	return []string{
		string(EventGridEventSchemaTypeCloudEventSchema),
		string(EventGridEventSchemaTypeEventGridEventSchema),
	}
}

func parseEventGridEventSchemaType(input string) (*EventGridEventSchemaType, error) {
	vals := map[string]EventGridEventSchemaType{
		"cloudeventschema":     EventGridEventSchemaTypeCloudEventSchema,
		"eventgrideventschema": EventGridEventSchemaTypeEventGridEventSchema,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventGridEventSchemaType(input)
	return &out, nil
}

type EventSerializationType string

const (
	EventSerializationTypeAvro      EventSerializationType = "Avro"
	EventSerializationTypeCsv       EventSerializationType = "Csv"
	EventSerializationTypeCustomClr EventSerializationType = "CustomClr"
	EventSerializationTypeDelta     EventSerializationType = "Delta"
	EventSerializationTypeJson      EventSerializationType = "Json"
	EventSerializationTypeParquet   EventSerializationType = "Parquet"
)

func PossibleValuesForEventSerializationType() []string {
	return []string{
		string(EventSerializationTypeAvro),
		string(EventSerializationTypeCsv),
		string(EventSerializationTypeCustomClr),
		string(EventSerializationTypeDelta),
		string(EventSerializationTypeJson),
		string(EventSerializationTypeParquet),
	}
}

func parseEventSerializationType(input string) (*EventSerializationType, error) {
	vals := map[string]EventSerializationType{
		"avro":      EventSerializationTypeAvro,
		"csv":       EventSerializationTypeCsv,
		"customclr": EventSerializationTypeCustomClr,
		"delta":     EventSerializationTypeDelta,
		"json":      EventSerializationTypeJson,
		"parquet":   EventSerializationTypeParquet,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventSerializationType(input)
	return &out, nil
}

type EventsOutOfOrderPolicy string

const (
	EventsOutOfOrderPolicyAdjust EventsOutOfOrderPolicy = "Adjust"
	EventsOutOfOrderPolicyDrop   EventsOutOfOrderPolicy = "Drop"
)

func PossibleValuesForEventsOutOfOrderPolicy() []string {
	return []string{
		string(EventsOutOfOrderPolicyAdjust),
		string(EventsOutOfOrderPolicyDrop),
	}
}

func parseEventsOutOfOrderPolicy(input string) (*EventsOutOfOrderPolicy, error) {
	vals := map[string]EventsOutOfOrderPolicy{
		"adjust": EventsOutOfOrderPolicyAdjust,
		"drop":   EventsOutOfOrderPolicyDrop,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventsOutOfOrderPolicy(input)
	return &out, nil
}

type InputWatermarkMode string

const (
	InputWatermarkModeNone          InputWatermarkMode = "None"
	InputWatermarkModeReadWatermark InputWatermarkMode = "ReadWatermark"
)

func PossibleValuesForInputWatermarkMode() []string {
	return []string{
		string(InputWatermarkModeNone),
		string(InputWatermarkModeReadWatermark),
	}
}

func parseInputWatermarkMode(input string) (*InputWatermarkMode, error) {
	vals := map[string]InputWatermarkMode{
		"none":          InputWatermarkModeNone,
		"readwatermark": InputWatermarkModeReadWatermark,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InputWatermarkMode(input)
	return &out, nil
}

type JobType string

const (
	JobTypeCloud JobType = "Cloud"
	JobTypeEdge  JobType = "Edge"
)

func PossibleValuesForJobType() []string {
	return []string{
		string(JobTypeCloud),
		string(JobTypeEdge),
	}
}

func parseJobType(input string) (*JobType, error) {
	vals := map[string]JobType{
		"cloud": JobTypeCloud,
		"edge":  JobTypeEdge,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := JobType(input)
	return &out, nil
}

type JsonOutputSerializationFormat string

const (
	JsonOutputSerializationFormatArray         JsonOutputSerializationFormat = "Array"
	JsonOutputSerializationFormatLineSeparated JsonOutputSerializationFormat = "LineSeparated"
)

func PossibleValuesForJsonOutputSerializationFormat() []string {
	return []string{
		string(JsonOutputSerializationFormatArray),
		string(JsonOutputSerializationFormatLineSeparated),
	}
}

func parseJsonOutputSerializationFormat(input string) (*JsonOutputSerializationFormat, error) {
	vals := map[string]JsonOutputSerializationFormat{
		"array":         JsonOutputSerializationFormatArray,
		"lineseparated": JsonOutputSerializationFormatLineSeparated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := JsonOutputSerializationFormat(input)
	return &out, nil
}

type OutputErrorPolicy string

const (
	OutputErrorPolicyDrop OutputErrorPolicy = "Drop"
	OutputErrorPolicyStop OutputErrorPolicy = "Stop"
)

func PossibleValuesForOutputErrorPolicy() []string {
	return []string{
		string(OutputErrorPolicyDrop),
		string(OutputErrorPolicyStop),
	}
}

func parseOutputErrorPolicy(input string) (*OutputErrorPolicy, error) {
	vals := map[string]OutputErrorPolicy{
		"drop": OutputErrorPolicyDrop,
		"stop": OutputErrorPolicyStop,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OutputErrorPolicy(input)
	return &out, nil
}

type OutputStartMode string

const (
	OutputStartModeCustomTime          OutputStartMode = "CustomTime"
	OutputStartModeJobStartTime        OutputStartMode = "JobStartTime"
	OutputStartModeLastOutputEventTime OutputStartMode = "LastOutputEventTime"
)

func PossibleValuesForOutputStartMode() []string {
	return []string{
		string(OutputStartModeCustomTime),
		string(OutputStartModeJobStartTime),
		string(OutputStartModeLastOutputEventTime),
	}
}

func parseOutputStartMode(input string) (*OutputStartMode, error) {
	vals := map[string]OutputStartMode{
		"customtime":          OutputStartModeCustomTime,
		"jobstarttime":        OutputStartModeJobStartTime,
		"lastoutputeventtime": OutputStartModeLastOutputEventTime,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OutputStartMode(input)
	return &out, nil
}

type OutputWatermarkMode string

const (
	OutputWatermarkModeNone                                OutputWatermarkMode = "None"
	OutputWatermarkModeSendCurrentPartitionWatermark       OutputWatermarkMode = "SendCurrentPartitionWatermark"
	OutputWatermarkModeSendLowestWatermarkAcrossPartitions OutputWatermarkMode = "SendLowestWatermarkAcrossPartitions"
)

func PossibleValuesForOutputWatermarkMode() []string {
	return []string{
		string(OutputWatermarkModeNone),
		string(OutputWatermarkModeSendCurrentPartitionWatermark),
		string(OutputWatermarkModeSendLowestWatermarkAcrossPartitions),
	}
}

func parseOutputWatermarkMode(input string) (*OutputWatermarkMode, error) {
	vals := map[string]OutputWatermarkMode{
		"none":                                OutputWatermarkModeNone,
		"sendcurrentpartitionwatermark":       OutputWatermarkModeSendCurrentPartitionWatermark,
		"sendlowestwatermarkacrosspartitions": OutputWatermarkModeSendLowestWatermarkAcrossPartitions,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OutputWatermarkMode(input)
	return &out, nil
}

type QueryTestingResultStatus string

const (
	QueryTestingResultStatusCompilerError QueryTestingResultStatus = "CompilerError"
	QueryTestingResultStatusRuntimeError  QueryTestingResultStatus = "RuntimeError"
	QueryTestingResultStatusStarted       QueryTestingResultStatus = "Started"
	QueryTestingResultStatusSuccess       QueryTestingResultStatus = "Success"
	QueryTestingResultStatusTimeout       QueryTestingResultStatus = "Timeout"
	QueryTestingResultStatusUnknownError  QueryTestingResultStatus = "UnknownError"
)

func PossibleValuesForQueryTestingResultStatus() []string {
	return []string{
		string(QueryTestingResultStatusCompilerError),
		string(QueryTestingResultStatusRuntimeError),
		string(QueryTestingResultStatusStarted),
		string(QueryTestingResultStatusSuccess),
		string(QueryTestingResultStatusTimeout),
		string(QueryTestingResultStatusUnknownError),
	}
}

func parseQueryTestingResultStatus(input string) (*QueryTestingResultStatus, error) {
	vals := map[string]QueryTestingResultStatus{
		"compilererror": QueryTestingResultStatusCompilerError,
		"runtimeerror":  QueryTestingResultStatusRuntimeError,
		"started":       QueryTestingResultStatusStarted,
		"success":       QueryTestingResultStatusSuccess,
		"timeout":       QueryTestingResultStatusTimeout,
		"unknownerror":  QueryTestingResultStatusUnknownError,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := QueryTestingResultStatus(input)
	return &out, nil
}

type RefreshType string

const (
	RefreshTypeRefreshPeriodicallyWithDelta RefreshType = "RefreshPeriodicallyWithDelta"
	RefreshTypeRefreshPeriodicallyWithFull  RefreshType = "RefreshPeriodicallyWithFull"
	RefreshTypeStatic                       RefreshType = "Static"
)

func PossibleValuesForRefreshType() []string {
	return []string{
		string(RefreshTypeRefreshPeriodicallyWithDelta),
		string(RefreshTypeRefreshPeriodicallyWithFull),
		string(RefreshTypeStatic),
	}
}

func parseRefreshType(input string) (*RefreshType, error) {
	vals := map[string]RefreshType{
		"refreshperiodicallywithdelta": RefreshTypeRefreshPeriodicallyWithDelta,
		"refreshperiodicallywithfull":  RefreshTypeRefreshPeriodicallyWithFull,
		"static":                       RefreshTypeStatic,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RefreshType(input)
	return &out, nil
}

type SampleInputResultStatus string

const (
	SampleInputResultStatusErrorConnectingToInput SampleInputResultStatus = "ErrorConnectingToInput"
	SampleInputResultStatusNoEventsFoundInRange   SampleInputResultStatus = "NoEventsFoundInRange"
	SampleInputResultStatusReadAllEventsInRange   SampleInputResultStatus = "ReadAllEventsInRange"
)

func PossibleValuesForSampleInputResultStatus() []string {
	return []string{
		string(SampleInputResultStatusErrorConnectingToInput),
		string(SampleInputResultStatusNoEventsFoundInRange),
		string(SampleInputResultStatusReadAllEventsInRange),
	}
}

func parseSampleInputResultStatus(input string) (*SampleInputResultStatus, error) {
	vals := map[string]SampleInputResultStatus{
		"errorconnectingtoinput": SampleInputResultStatusErrorConnectingToInput,
		"noeventsfoundinrange":   SampleInputResultStatusNoEventsFoundInRange,
		"readalleventsinrange":   SampleInputResultStatusReadAllEventsInRange,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SampleInputResultStatus(input)
	return &out, nil
}

type SkuName string

const (
	SkuNameStandard SkuName = "Standard"
)

func PossibleValuesForSkuName() []string {
	return []string{
		string(SkuNameStandard),
	}
}

func parseSkuName(input string) (*SkuName, error) {
	vals := map[string]SkuName{
		"standard": SkuNameStandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SkuName(input)
	return &out, nil
}

type TestDatasourceResultStatus string

const (
	TestDatasourceResultStatusTestFailed    TestDatasourceResultStatus = "TestFailed"
	TestDatasourceResultStatusTestSucceeded TestDatasourceResultStatus = "TestSucceeded"
)

func PossibleValuesForTestDatasourceResultStatus() []string {
	return []string{
		string(TestDatasourceResultStatusTestFailed),
		string(TestDatasourceResultStatusTestSucceeded),
	}
}

func parseTestDatasourceResultStatus(input string) (*TestDatasourceResultStatus, error) {
	vals := map[string]TestDatasourceResultStatus{
		"testfailed":    TestDatasourceResultStatusTestFailed,
		"testsucceeded": TestDatasourceResultStatusTestSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TestDatasourceResultStatus(input)
	return &out, nil
}

type UpdatableUdfRefreshType string

const (
	UpdatableUdfRefreshTypeBlocking    UpdatableUdfRefreshType = "Blocking"
	UpdatableUdfRefreshTypeNonblocking UpdatableUdfRefreshType = "Nonblocking"
)

func PossibleValuesForUpdatableUdfRefreshType() []string {
	return []string{
		string(UpdatableUdfRefreshTypeBlocking),
		string(UpdatableUdfRefreshTypeNonblocking),
	}
}

func parseUpdatableUdfRefreshType(input string) (*UpdatableUdfRefreshType, error) {
	vals := map[string]UpdatableUdfRefreshType{
		"blocking":    UpdatableUdfRefreshTypeBlocking,
		"nonblocking": UpdatableUdfRefreshTypeNonblocking,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UpdatableUdfRefreshType(input)
	return &out, nil
}

type UpdateMode string

const (
	UpdateModeRefreshable UpdateMode = "Refreshable"
	UpdateModeStatic      UpdateMode = "Static"
)

func PossibleValuesForUpdateMode() []string {
	return []string{
		string(UpdateModeRefreshable),
		string(UpdateModeStatic),
	}
}

func parseUpdateMode(input string) (*UpdateMode, error) {
	vals := map[string]UpdateMode{
		"refreshable": UpdateModeRefreshable,
		"static":      UpdateModeStatic,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UpdateMode(input)
	return &out, nil
}
