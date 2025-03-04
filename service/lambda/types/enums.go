// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type Architecture string

// Enum values for Architecture
const (
	ArchitectureX8664 Architecture = "x86_64"
	ArchitectureArm64 Architecture = "arm64"
)

// Values returns all known values for Architecture. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (Architecture) Values() []Architecture {
	return []Architecture{
		"x86_64",
		"arm64",
	}
}

type CodeSigningPolicy string

// Enum values for CodeSigningPolicy
const (
	CodeSigningPolicyWarn    CodeSigningPolicy = "Warn"
	CodeSigningPolicyEnforce CodeSigningPolicy = "Enforce"
)

// Values returns all known values for CodeSigningPolicy. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CodeSigningPolicy) Values() []CodeSigningPolicy {
	return []CodeSigningPolicy{
		"Warn",
		"Enforce",
	}
}

type EndPointType string

// Enum values for EndPointType
const (
	EndPointTypeKafkaBootstrapServers EndPointType = "KAFKA_BOOTSTRAP_SERVERS"
)

// Values returns all known values for EndPointType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (EndPointType) Values() []EndPointType {
	return []EndPointType{
		"KAFKA_BOOTSTRAP_SERVERS",
	}
}

type EventSourcePosition string

// Enum values for EventSourcePosition
const (
	EventSourcePositionTrimHorizon EventSourcePosition = "TRIM_HORIZON"
	EventSourcePositionLatest      EventSourcePosition = "LATEST"
	EventSourcePositionAtTimestamp EventSourcePosition = "AT_TIMESTAMP"
)

// Values returns all known values for EventSourcePosition. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EventSourcePosition) Values() []EventSourcePosition {
	return []EventSourcePosition{
		"TRIM_HORIZON",
		"LATEST",
		"AT_TIMESTAMP",
	}
}

type FullDocument string

// Enum values for FullDocument
const (
	FullDocumentUpdateLookup FullDocument = "UpdateLookup"
	FullDocumentDefault      FullDocument = "Default"
)

// Values returns all known values for FullDocument. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (FullDocument) Values() []FullDocument {
	return []FullDocument{
		"UpdateLookup",
		"Default",
	}
}

type FunctionResponseType string

// Enum values for FunctionResponseType
const (
	FunctionResponseTypeReportBatchItemFailures FunctionResponseType = "ReportBatchItemFailures"
)

// Values returns all known values for FunctionResponseType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FunctionResponseType) Values() []FunctionResponseType {
	return []FunctionResponseType{
		"ReportBatchItemFailures",
	}
}

type FunctionUrlAuthType string

// Enum values for FunctionUrlAuthType
const (
	FunctionUrlAuthTypeNone   FunctionUrlAuthType = "NONE"
	FunctionUrlAuthTypeAwsIam FunctionUrlAuthType = "AWS_IAM"
)

// Values returns all known values for FunctionUrlAuthType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FunctionUrlAuthType) Values() []FunctionUrlAuthType {
	return []FunctionUrlAuthType{
		"NONE",
		"AWS_IAM",
	}
}

type FunctionVersion string

// Enum values for FunctionVersion
const (
	FunctionVersionAll FunctionVersion = "ALL"
)

// Values returns all known values for FunctionVersion. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FunctionVersion) Values() []FunctionVersion {
	return []FunctionVersion{
		"ALL",
	}
}

type InvocationType string

// Enum values for InvocationType
const (
	InvocationTypeEvent           InvocationType = "Event"
	InvocationTypeRequestResponse InvocationType = "RequestResponse"
	InvocationTypeDryRun          InvocationType = "DryRun"
)

// Values returns all known values for InvocationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (InvocationType) Values() []InvocationType {
	return []InvocationType{
		"Event",
		"RequestResponse",
		"DryRun",
	}
}

type LastUpdateStatus string

// Enum values for LastUpdateStatus
const (
	LastUpdateStatusSuccessful LastUpdateStatus = "Successful"
	LastUpdateStatusFailed     LastUpdateStatus = "Failed"
	LastUpdateStatusInProgress LastUpdateStatus = "InProgress"
)

// Values returns all known values for LastUpdateStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (LastUpdateStatus) Values() []LastUpdateStatus {
	return []LastUpdateStatus{
		"Successful",
		"Failed",
		"InProgress",
	}
}

type LastUpdateStatusReasonCode string

// Enum values for LastUpdateStatusReasonCode
const (
	LastUpdateStatusReasonCodeEniLimitExceeded            LastUpdateStatusReasonCode = "EniLimitExceeded"
	LastUpdateStatusReasonCodeInsufficientRolePermissions LastUpdateStatusReasonCode = "InsufficientRolePermissions"
	LastUpdateStatusReasonCodeInvalidConfiguration        LastUpdateStatusReasonCode = "InvalidConfiguration"
	LastUpdateStatusReasonCodeInternalError               LastUpdateStatusReasonCode = "InternalError"
	LastUpdateStatusReasonCodeSubnetOutOfIPAddresses      LastUpdateStatusReasonCode = "SubnetOutOfIPAddresses"
	LastUpdateStatusReasonCodeInvalidSubnet               LastUpdateStatusReasonCode = "InvalidSubnet"
	LastUpdateStatusReasonCodeInvalidSecurityGroup        LastUpdateStatusReasonCode = "InvalidSecurityGroup"
	LastUpdateStatusReasonCodeImageDeleted                LastUpdateStatusReasonCode = "ImageDeleted"
	LastUpdateStatusReasonCodeImageAccessDenied           LastUpdateStatusReasonCode = "ImageAccessDenied"
	LastUpdateStatusReasonCodeInvalidImage                LastUpdateStatusReasonCode = "InvalidImage"
	LastUpdateStatusReasonCodeKMSKeyAccessDenied          LastUpdateStatusReasonCode = "KMSKeyAccessDenied"
	LastUpdateStatusReasonCodeKMSKeyNotFound              LastUpdateStatusReasonCode = "KMSKeyNotFound"
	LastUpdateStatusReasonCodeInvalidStateKMSKey          LastUpdateStatusReasonCode = "InvalidStateKMSKey"
	LastUpdateStatusReasonCodeDisabledKMSKey              LastUpdateStatusReasonCode = "DisabledKMSKey"
	LastUpdateStatusReasonCodeEFSIOError                  LastUpdateStatusReasonCode = "EFSIOError"
	LastUpdateStatusReasonCodeEFSMountConnectivityError   LastUpdateStatusReasonCode = "EFSMountConnectivityError"
	LastUpdateStatusReasonCodeEFSMountFailure             LastUpdateStatusReasonCode = "EFSMountFailure"
	LastUpdateStatusReasonCodeEFSMountTimeout             LastUpdateStatusReasonCode = "EFSMountTimeout"
	LastUpdateStatusReasonCodeInvalidRuntime              LastUpdateStatusReasonCode = "InvalidRuntime"
	LastUpdateStatusReasonCodeInvalidZipFileException     LastUpdateStatusReasonCode = "InvalidZipFileException"
	LastUpdateStatusReasonCodeFunctionError               LastUpdateStatusReasonCode = "FunctionError"
)

// Values returns all known values for LastUpdateStatusReasonCode. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (LastUpdateStatusReasonCode) Values() []LastUpdateStatusReasonCode {
	return []LastUpdateStatusReasonCode{
		"EniLimitExceeded",
		"InsufficientRolePermissions",
		"InvalidConfiguration",
		"InternalError",
		"SubnetOutOfIPAddresses",
		"InvalidSubnet",
		"InvalidSecurityGroup",
		"ImageDeleted",
		"ImageAccessDenied",
		"InvalidImage",
		"KMSKeyAccessDenied",
		"KMSKeyNotFound",
		"InvalidStateKMSKey",
		"DisabledKMSKey",
		"EFSIOError",
		"EFSMountConnectivityError",
		"EFSMountFailure",
		"EFSMountTimeout",
		"InvalidRuntime",
		"InvalidZipFileException",
		"FunctionError",
	}
}

type LogType string

// Enum values for LogType
const (
	LogTypeNone LogType = "None"
	LogTypeTail LogType = "Tail"
)

// Values returns all known values for LogType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (LogType) Values() []LogType {
	return []LogType{
		"None",
		"Tail",
	}
}

type PackageType string

// Enum values for PackageType
const (
	PackageTypeZip   PackageType = "Zip"
	PackageTypeImage PackageType = "Image"
)

// Values returns all known values for PackageType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (PackageType) Values() []PackageType {
	return []PackageType{
		"Zip",
		"Image",
	}
}

type ProvisionedConcurrencyStatusEnum string

// Enum values for ProvisionedConcurrencyStatusEnum
const (
	ProvisionedConcurrencyStatusEnumInProgress ProvisionedConcurrencyStatusEnum = "IN_PROGRESS"
	ProvisionedConcurrencyStatusEnumReady      ProvisionedConcurrencyStatusEnum = "READY"
	ProvisionedConcurrencyStatusEnumFailed     ProvisionedConcurrencyStatusEnum = "FAILED"
)

// Values returns all known values for ProvisionedConcurrencyStatusEnum. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (ProvisionedConcurrencyStatusEnum) Values() []ProvisionedConcurrencyStatusEnum {
	return []ProvisionedConcurrencyStatusEnum{
		"IN_PROGRESS",
		"READY",
		"FAILED",
	}
}

type Runtime string

// Enum values for Runtime
const (
	RuntimeNodejs       Runtime = "nodejs"
	RuntimeNodejs43     Runtime = "nodejs4.3"
	RuntimeNodejs610    Runtime = "nodejs6.10"
	RuntimeNodejs810    Runtime = "nodejs8.10"
	RuntimeNodejs10x    Runtime = "nodejs10.x"
	RuntimeNodejs12x    Runtime = "nodejs12.x"
	RuntimeNodejs14x    Runtime = "nodejs14.x"
	RuntimeNodejs16x    Runtime = "nodejs16.x"
	RuntimeJava8        Runtime = "java8"
	RuntimeJava8al2     Runtime = "java8.al2"
	RuntimeJava11       Runtime = "java11"
	RuntimePython27     Runtime = "python2.7"
	RuntimePython36     Runtime = "python3.6"
	RuntimePython37     Runtime = "python3.7"
	RuntimePython38     Runtime = "python3.8"
	RuntimePython39     Runtime = "python3.9"
	RuntimeDotnetcore10 Runtime = "dotnetcore1.0"
	RuntimeDotnetcore20 Runtime = "dotnetcore2.0"
	RuntimeDotnetcore21 Runtime = "dotnetcore2.1"
	RuntimeDotnetcore31 Runtime = "dotnetcore3.1"
	RuntimeDotnet6      Runtime = "dotnet6"
	RuntimeNodejs43edge Runtime = "nodejs4.3-edge"
	RuntimeGo1x         Runtime = "go1.x"
	RuntimeRuby25       Runtime = "ruby2.5"
	RuntimeRuby27       Runtime = "ruby2.7"
	RuntimeProvided     Runtime = "provided"
	RuntimeProvidedal2  Runtime = "provided.al2"
	RuntimeNodejs18x    Runtime = "nodejs18.x"
)

// Values returns all known values for Runtime. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Runtime) Values() []Runtime {
	return []Runtime{
		"nodejs",
		"nodejs4.3",
		"nodejs6.10",
		"nodejs8.10",
		"nodejs10.x",
		"nodejs12.x",
		"nodejs14.x",
		"nodejs16.x",
		"java8",
		"java8.al2",
		"java11",
		"python2.7",
		"python3.6",
		"python3.7",
		"python3.8",
		"python3.9",
		"dotnetcore1.0",
		"dotnetcore2.0",
		"dotnetcore2.1",
		"dotnetcore3.1",
		"dotnet6",
		"nodejs4.3-edge",
		"go1.x",
		"ruby2.5",
		"ruby2.7",
		"provided",
		"provided.al2",
		"nodejs18.x",
	}
}

type SnapStartApplyOn string

// Enum values for SnapStartApplyOn
const (
	SnapStartApplyOnPublishedVersions SnapStartApplyOn = "PublishedVersions"
	SnapStartApplyOnNone              SnapStartApplyOn = "None"
)

// Values returns all known values for SnapStartApplyOn. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SnapStartApplyOn) Values() []SnapStartApplyOn {
	return []SnapStartApplyOn{
		"PublishedVersions",
		"None",
	}
}

type SnapStartOptimizationStatus string

// Enum values for SnapStartOptimizationStatus
const (
	SnapStartOptimizationStatusOn  SnapStartOptimizationStatus = "On"
	SnapStartOptimizationStatusOff SnapStartOptimizationStatus = "Off"
)

// Values returns all known values for SnapStartOptimizationStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (SnapStartOptimizationStatus) Values() []SnapStartOptimizationStatus {
	return []SnapStartOptimizationStatus{
		"On",
		"Off",
	}
}

type SourceAccessType string

// Enum values for SourceAccessType
const (
	SourceAccessTypeBasicAuth                SourceAccessType = "BASIC_AUTH"
	SourceAccessTypeVpcSubnet                SourceAccessType = "VPC_SUBNET"
	SourceAccessTypeVpcSecurityGroup         SourceAccessType = "VPC_SECURITY_GROUP"
	SourceAccessTypeSaslScram512Auth         SourceAccessType = "SASL_SCRAM_512_AUTH"
	SourceAccessTypeSaslScram256Auth         SourceAccessType = "SASL_SCRAM_256_AUTH"
	SourceAccessTypeVirtualHost              SourceAccessType = "VIRTUAL_HOST"
	SourceAccessTypeClientCertificateTlsAuth SourceAccessType = "CLIENT_CERTIFICATE_TLS_AUTH"
	SourceAccessTypeServerRootCaCertificate  SourceAccessType = "SERVER_ROOT_CA_CERTIFICATE"
)

// Values returns all known values for SourceAccessType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SourceAccessType) Values() []SourceAccessType {
	return []SourceAccessType{
		"BASIC_AUTH",
		"VPC_SUBNET",
		"VPC_SECURITY_GROUP",
		"SASL_SCRAM_512_AUTH",
		"SASL_SCRAM_256_AUTH",
		"VIRTUAL_HOST",
		"CLIENT_CERTIFICATE_TLS_AUTH",
		"SERVER_ROOT_CA_CERTIFICATE",
	}
}

type State string

// Enum values for State
const (
	StatePending  State = "Pending"
	StateActive   State = "Active"
	StateInactive State = "Inactive"
	StateFailed   State = "Failed"
)

// Values returns all known values for State. Note that this can be expanded in the
// future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (State) Values() []State {
	return []State{
		"Pending",
		"Active",
		"Inactive",
		"Failed",
	}
}

type StateReasonCode string

// Enum values for StateReasonCode
const (
	StateReasonCodeIdle                        StateReasonCode = "Idle"
	StateReasonCodeCreating                    StateReasonCode = "Creating"
	StateReasonCodeRestoring                   StateReasonCode = "Restoring"
	StateReasonCodeEniLimitExceeded            StateReasonCode = "EniLimitExceeded"
	StateReasonCodeInsufficientRolePermissions StateReasonCode = "InsufficientRolePermissions"
	StateReasonCodeInvalidConfiguration        StateReasonCode = "InvalidConfiguration"
	StateReasonCodeInternalError               StateReasonCode = "InternalError"
	StateReasonCodeSubnetOutOfIPAddresses      StateReasonCode = "SubnetOutOfIPAddresses"
	StateReasonCodeInvalidSubnet               StateReasonCode = "InvalidSubnet"
	StateReasonCodeInvalidSecurityGroup        StateReasonCode = "InvalidSecurityGroup"
	StateReasonCodeImageDeleted                StateReasonCode = "ImageDeleted"
	StateReasonCodeImageAccessDenied           StateReasonCode = "ImageAccessDenied"
	StateReasonCodeInvalidImage                StateReasonCode = "InvalidImage"
	StateReasonCodeKMSKeyAccessDenied          StateReasonCode = "KMSKeyAccessDenied"
	StateReasonCodeKMSKeyNotFound              StateReasonCode = "KMSKeyNotFound"
	StateReasonCodeInvalidStateKMSKey          StateReasonCode = "InvalidStateKMSKey"
	StateReasonCodeDisabledKMSKey              StateReasonCode = "DisabledKMSKey"
	StateReasonCodeEFSIOError                  StateReasonCode = "EFSIOError"
	StateReasonCodeEFSMountConnectivityError   StateReasonCode = "EFSMountConnectivityError"
	StateReasonCodeEFSMountFailure             StateReasonCode = "EFSMountFailure"
	StateReasonCodeEFSMountTimeout             StateReasonCode = "EFSMountTimeout"
	StateReasonCodeInvalidRuntime              StateReasonCode = "InvalidRuntime"
	StateReasonCodeInvalidZipFileException     StateReasonCode = "InvalidZipFileException"
	StateReasonCodeFunctionError               StateReasonCode = "FunctionError"
)

// Values returns all known values for StateReasonCode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (StateReasonCode) Values() []StateReasonCode {
	return []StateReasonCode{
		"Idle",
		"Creating",
		"Restoring",
		"EniLimitExceeded",
		"InsufficientRolePermissions",
		"InvalidConfiguration",
		"InternalError",
		"SubnetOutOfIPAddresses",
		"InvalidSubnet",
		"InvalidSecurityGroup",
		"ImageDeleted",
		"ImageAccessDenied",
		"InvalidImage",
		"KMSKeyAccessDenied",
		"KMSKeyNotFound",
		"InvalidStateKMSKey",
		"DisabledKMSKey",
		"EFSIOError",
		"EFSMountConnectivityError",
		"EFSMountFailure",
		"EFSMountTimeout",
		"InvalidRuntime",
		"InvalidZipFileException",
		"FunctionError",
	}
}

type ThrottleReason string

// Enum values for ThrottleReason
const (
	ThrottleReasonConcurrentInvocationLimitExceeded                 ThrottleReason = "ConcurrentInvocationLimitExceeded"
	ThrottleReasonFunctionInvocationRateLimitExceeded               ThrottleReason = "FunctionInvocationRateLimitExceeded"
	ThrottleReasonReservedFunctionConcurrentInvocationLimitExceeded ThrottleReason = "ReservedFunctionConcurrentInvocationLimitExceeded"
	ThrottleReasonReservedFunctionInvocationRateLimitExceeded       ThrottleReason = "ReservedFunctionInvocationRateLimitExceeded"
	ThrottleReasonCallerRateLimitExceeded                           ThrottleReason = "CallerRateLimitExceeded"
	ThrottleReasonConcurrentSnapshotCreateLimitExceeded             ThrottleReason = "ConcurrentSnapshotCreateLimitExceeded"
)

// Values returns all known values for ThrottleReason. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ThrottleReason) Values() []ThrottleReason {
	return []ThrottleReason{
		"ConcurrentInvocationLimitExceeded",
		"FunctionInvocationRateLimitExceeded",
		"ReservedFunctionConcurrentInvocationLimitExceeded",
		"ReservedFunctionInvocationRateLimitExceeded",
		"CallerRateLimitExceeded",
		"ConcurrentSnapshotCreateLimitExceeded",
	}
}

type TracingMode string

// Enum values for TracingMode
const (
	TracingModeActive      TracingMode = "Active"
	TracingModePassThrough TracingMode = "PassThrough"
)

// Values returns all known values for TracingMode. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (TracingMode) Values() []TracingMode {
	return []TracingMode{
		"Active",
		"PassThrough",
	}
}

type UpdateRuntimeOn string

// Enum values for UpdateRuntimeOn
const (
	UpdateRuntimeOnAuto           UpdateRuntimeOn = "Auto"
	UpdateRuntimeOnManual         UpdateRuntimeOn = "Manual"
	UpdateRuntimeOnFunctionUpdate UpdateRuntimeOn = "FunctionUpdate"
)

// Values returns all known values for UpdateRuntimeOn. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (UpdateRuntimeOn) Values() []UpdateRuntimeOn {
	return []UpdateRuntimeOn{
		"Auto",
		"Manual",
		"FunctionUpdate",
	}
}
