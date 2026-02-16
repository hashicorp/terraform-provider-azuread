package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WindowsProtectionState{}

type WindowsProtectionState struct {
	// Current anti malware version
	AntiMalwareVersion nullable.Type[string] `json:"antiMalwareVersion,omitempty"`

	// Device malware list
	DetectedMalwareState *[]WindowsDeviceMalwareState `json:"detectedMalwareState,omitempty"`

	// Indicates device's health state. Possible values are: clean, fullScanPending, rebootPending, manualStepsPending,
	// offlineScanPending, critical. Possible values are: clean, fullScanPending, rebootPending, manualStepsPending,
	// offlineScanPending, critical.
	DeviceState *WindowsDeviceHealthState `json:"deviceState,omitempty"`

	// Current endpoint protection engine's version
	EngineVersion nullable.Type[string] `json:"engineVersion,omitempty"`

	// When TRUE indicates full scan is overdue, when FALSE indicates full scan is not overdue. Defaults to setting on
	// client device.
	FullScanOverdue nullable.Type[bool] `json:"fullScanOverdue,omitempty"`

	// When TRUE indicates full scan is required, when FALSE indicates full scan is not required. Defaults to setting on
	// client device.
	FullScanRequired nullable.Type[bool] `json:"fullScanRequired,omitempty"`

	// When TRUE indicates the device is a virtual machine, when FALSE indicates the device is not a virtual machine.
	// Defaults to setting on client device.
	IsVirtualMachine nullable.Type[bool] `json:"isVirtualMachine,omitempty"`

	// Last quick scan datetime
	LastFullScanDateTime nullable.Type[string] `json:"lastFullScanDateTime,omitempty"`

	// Last full scan signature version
	LastFullScanSignatureVersion nullable.Type[string] `json:"lastFullScanSignatureVersion,omitempty"`

	// Last quick scan datetime
	LastQuickScanDateTime nullable.Type[string] `json:"lastQuickScanDateTime,omitempty"`

	// Last quick scan signature version
	LastQuickScanSignatureVersion nullable.Type[string] `json:"lastQuickScanSignatureVersion,omitempty"`

	// Last device health status reported time
	LastReportedDateTime nullable.Type[string] `json:"lastReportedDateTime,omitempty"`

	// When TRUE indicates anti malware is enabled when FALSE indicates anti malware is not enabled.
	MalwareProtectionEnabled nullable.Type[bool] `json:"malwareProtectionEnabled,omitempty"`

	// When TRUE indicates network inspection system enabled, when FALSE indicates network inspection system is not enabled.
	// Defaults to setting on client device.
	NetworkInspectionSystemEnabled nullable.Type[bool] `json:"networkInspectionSystemEnabled,omitempty"`

	// Product Status of Windows Defender Antivirus. Possible values are: noStatus, serviceNotRunning,
	// serviceStartedWithoutMalwareProtection, pendingFullScanDueToThreatAction, pendingRebootDueToThreatAction,
	// pendingManualStepsDueToThreatAction, avSignaturesOutOfDate, asSignaturesOutOfDate,
	// noQuickScanHappenedForSpecifiedPeriod, noFullScanHappenedForSpecifiedPeriod, systemInitiatedScanInProgress,
	// systemInitiatedCleanInProgress, samplesPendingSubmission, productRunningInEvaluationMode,
	// productRunningInNonGenuineMode, productExpired, offlineScanRequired, serviceShutdownAsPartOfSystemShutdown,
	// threatRemediationFailedCritically, threatRemediationFailedNonCritically, noStatusFlagsSet, platformOutOfDate,
	// platformUpdateInProgress, platformAboutToBeOutdated, signatureOrPlatformEndOfLifeIsPastOrIsImpending,
	// windowsSModeSignaturesInUseOnNonWin10SInstall. Possible values are: noStatus, serviceNotRunning,
	// serviceStartedWithoutMalwareProtection, pendingFullScanDueToThreatAction, pendingRebootDueToThreatAction,
	// pendingManualStepsDueToThreatAction, avSignaturesOutOfDate, asSignaturesOutOfDate,
	// noQuickScanHappenedForSpecifiedPeriod, noFullScanHappenedForSpecifiedPeriod, systemInitiatedScanInProgress,
	// systemInitiatedCleanInProgress, samplesPendingSubmission, productRunningInEvaluationMode,
	// productRunningInNonGenuineMode, productExpired, offlineScanRequired, serviceShutdownAsPartOfSystemShutdown,
	// threatRemediationFailedCritically, threatRemediationFailedNonCritically, noStatusFlagsSet, platformOutOfDate,
	// platformUpdateInProgress, platformAboutToBeOutdated, signatureOrPlatformEndOfLifeIsPastOrIsImpending,
	// windowsSModeSignaturesInUseOnNonWin10SInstall.
	ProductStatus *WindowsDefenderProductStatus `json:"productStatus,omitempty"`

	// When TRUE indicates quick scan is overdue, when FALSE indicates quick scan is not overdue. Defaults to setting on
	// client device.
	QuickScanOverdue nullable.Type[bool] `json:"quickScanOverdue,omitempty"`

	// When TRUE indicates real time protection is enabled, when FALSE indicates real time protection is not enabled.
	// Defaults to setting on client device.
	RealTimeProtectionEnabled nullable.Type[bool] `json:"realTimeProtectionEnabled,omitempty"`

	// When TRUE indicates reboot is required, when FALSE indicates when TRUE indicates reboot is not required. Defaults to
	// setting on client device.
	RebootRequired nullable.Type[bool] `json:"rebootRequired,omitempty"`

	// When TRUE indicates signature is out of date, when FALSE indicates signature is not out of date. Defaults to setting
	// on client device.
	SignatureUpdateOverdue nullable.Type[bool] `json:"signatureUpdateOverdue,omitempty"`

	// Current malware definitions version
	SignatureVersion nullable.Type[string] `json:"signatureVersion,omitempty"`

	// When TRUE indicates the Windows Defender tamper protection feature is enabled, when FALSE indicates the Windows
	// Defender tamper protection feature is not enabled. Defaults to setting on client device.
	TamperProtectionEnabled nullable.Type[bool] `json:"tamperProtectionEnabled,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s WindowsProtectionState) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsProtectionState{}

func (s WindowsProtectionState) MarshalJSON() ([]byte, error) {
	type wrapper WindowsProtectionState
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsProtectionState: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsProtectionState: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsProtectionState"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsProtectionState: %+v", err)
	}

	return encoded, nil
}
