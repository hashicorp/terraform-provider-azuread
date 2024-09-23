package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsDefenderProductStatus string

const (
	WindowsDefenderProductStatus_AsSignaturesOutOfDate                           WindowsDefenderProductStatus = "asSignaturesOutOfDate"
	WindowsDefenderProductStatus_AvSignaturesOutOfDate                           WindowsDefenderProductStatus = "avSignaturesOutOfDate"
	WindowsDefenderProductStatus_NoFullScanHappenedForSpecifiedPeriod            WindowsDefenderProductStatus = "noFullScanHappenedForSpecifiedPeriod"
	WindowsDefenderProductStatus_NoQuickScanHappenedForSpecifiedPeriod           WindowsDefenderProductStatus = "noQuickScanHappenedForSpecifiedPeriod"
	WindowsDefenderProductStatus_NoStatus                                        WindowsDefenderProductStatus = "noStatus"
	WindowsDefenderProductStatus_NoStatusFlagsSet                                WindowsDefenderProductStatus = "noStatusFlagsSet"
	WindowsDefenderProductStatus_OfflineScanRequired                             WindowsDefenderProductStatus = "offlineScanRequired"
	WindowsDefenderProductStatus_PendingFullScanDueToThreatAction                WindowsDefenderProductStatus = "pendingFullScanDueToThreatAction"
	WindowsDefenderProductStatus_PendingManualStepsDueToThreatAction             WindowsDefenderProductStatus = "pendingManualStepsDueToThreatAction"
	WindowsDefenderProductStatus_PendingRebootDueToThreatAction                  WindowsDefenderProductStatus = "pendingRebootDueToThreatAction"
	WindowsDefenderProductStatus_PlatformAboutToBeOutdated                       WindowsDefenderProductStatus = "platformAboutToBeOutdated"
	WindowsDefenderProductStatus_PlatformOutOfDate                               WindowsDefenderProductStatus = "platformOutOfDate"
	WindowsDefenderProductStatus_PlatformUpdateInProgress                        WindowsDefenderProductStatus = "platformUpdateInProgress"
	WindowsDefenderProductStatus_ProductExpired                                  WindowsDefenderProductStatus = "productExpired"
	WindowsDefenderProductStatus_ProductRunningInEvaluationMode                  WindowsDefenderProductStatus = "productRunningInEvaluationMode"
	WindowsDefenderProductStatus_ProductRunningInNonGenuineMode                  WindowsDefenderProductStatus = "productRunningInNonGenuineMode"
	WindowsDefenderProductStatus_SamplesPendingSubmission                        WindowsDefenderProductStatus = "samplesPendingSubmission"
	WindowsDefenderProductStatus_ServiceNotRunning                               WindowsDefenderProductStatus = "serviceNotRunning"
	WindowsDefenderProductStatus_ServiceShutdownAsPartOfSystemShutdown           WindowsDefenderProductStatus = "serviceShutdownAsPartOfSystemShutdown"
	WindowsDefenderProductStatus_ServiceStartedWithoutMalwareProtection          WindowsDefenderProductStatus = "serviceStartedWithoutMalwareProtection"
	WindowsDefenderProductStatus_SignatureOrPlatformEndOfLifeIsPastOrIsImpending WindowsDefenderProductStatus = "signatureOrPlatformEndOfLifeIsPastOrIsImpending"
	WindowsDefenderProductStatus_SystemInitiatedCleanInProgress                  WindowsDefenderProductStatus = "systemInitiatedCleanInProgress"
	WindowsDefenderProductStatus_SystemInitiatedScanInProgress                   WindowsDefenderProductStatus = "systemInitiatedScanInProgress"
	WindowsDefenderProductStatus_ThreatRemediationFailedCritically               WindowsDefenderProductStatus = "threatRemediationFailedCritically"
	WindowsDefenderProductStatus_ThreatRemediationFailedNonCritically            WindowsDefenderProductStatus = "threatRemediationFailedNonCritically"
	WindowsDefenderProductStatus_WindowsSModeSignaturesInUseOnNonWin10SInstall   WindowsDefenderProductStatus = "windowsSModeSignaturesInUseOnNonWin10SInstall"
)

func PossibleValuesForWindowsDefenderProductStatus() []string {
	return []string{
		string(WindowsDefenderProductStatus_AsSignaturesOutOfDate),
		string(WindowsDefenderProductStatus_AvSignaturesOutOfDate),
		string(WindowsDefenderProductStatus_NoFullScanHappenedForSpecifiedPeriod),
		string(WindowsDefenderProductStatus_NoQuickScanHappenedForSpecifiedPeriod),
		string(WindowsDefenderProductStatus_NoStatus),
		string(WindowsDefenderProductStatus_NoStatusFlagsSet),
		string(WindowsDefenderProductStatus_OfflineScanRequired),
		string(WindowsDefenderProductStatus_PendingFullScanDueToThreatAction),
		string(WindowsDefenderProductStatus_PendingManualStepsDueToThreatAction),
		string(WindowsDefenderProductStatus_PendingRebootDueToThreatAction),
		string(WindowsDefenderProductStatus_PlatformAboutToBeOutdated),
		string(WindowsDefenderProductStatus_PlatformOutOfDate),
		string(WindowsDefenderProductStatus_PlatformUpdateInProgress),
		string(WindowsDefenderProductStatus_ProductExpired),
		string(WindowsDefenderProductStatus_ProductRunningInEvaluationMode),
		string(WindowsDefenderProductStatus_ProductRunningInNonGenuineMode),
		string(WindowsDefenderProductStatus_SamplesPendingSubmission),
		string(WindowsDefenderProductStatus_ServiceNotRunning),
		string(WindowsDefenderProductStatus_ServiceShutdownAsPartOfSystemShutdown),
		string(WindowsDefenderProductStatus_ServiceStartedWithoutMalwareProtection),
		string(WindowsDefenderProductStatus_SignatureOrPlatformEndOfLifeIsPastOrIsImpending),
		string(WindowsDefenderProductStatus_SystemInitiatedCleanInProgress),
		string(WindowsDefenderProductStatus_SystemInitiatedScanInProgress),
		string(WindowsDefenderProductStatus_ThreatRemediationFailedCritically),
		string(WindowsDefenderProductStatus_ThreatRemediationFailedNonCritically),
		string(WindowsDefenderProductStatus_WindowsSModeSignaturesInUseOnNonWin10SInstall),
	}
}

func (s *WindowsDefenderProductStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsDefenderProductStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsDefenderProductStatus(input string) (*WindowsDefenderProductStatus, error) {
	vals := map[string]WindowsDefenderProductStatus{
		"assignaturesoutofdate":                           WindowsDefenderProductStatus_AsSignaturesOutOfDate,
		"avsignaturesoutofdate":                           WindowsDefenderProductStatus_AvSignaturesOutOfDate,
		"nofullscanhappenedforspecifiedperiod":            WindowsDefenderProductStatus_NoFullScanHappenedForSpecifiedPeriod,
		"noquickscanhappenedforspecifiedperiod":           WindowsDefenderProductStatus_NoQuickScanHappenedForSpecifiedPeriod,
		"nostatus":                                        WindowsDefenderProductStatus_NoStatus,
		"nostatusflagsset":                                WindowsDefenderProductStatus_NoStatusFlagsSet,
		"offlinescanrequired":                             WindowsDefenderProductStatus_OfflineScanRequired,
		"pendingfullscanduetothreataction":                WindowsDefenderProductStatus_PendingFullScanDueToThreatAction,
		"pendingmanualstepsduetothreataction":             WindowsDefenderProductStatus_PendingManualStepsDueToThreatAction,
		"pendingrebootduetothreataction":                  WindowsDefenderProductStatus_PendingRebootDueToThreatAction,
		"platformabouttobeoutdated":                       WindowsDefenderProductStatus_PlatformAboutToBeOutdated,
		"platformoutofdate":                               WindowsDefenderProductStatus_PlatformOutOfDate,
		"platformupdateinprogress":                        WindowsDefenderProductStatus_PlatformUpdateInProgress,
		"productexpired":                                  WindowsDefenderProductStatus_ProductExpired,
		"productrunninginevaluationmode":                  WindowsDefenderProductStatus_ProductRunningInEvaluationMode,
		"productrunninginnongenuinemode":                  WindowsDefenderProductStatus_ProductRunningInNonGenuineMode,
		"samplespendingsubmission":                        WindowsDefenderProductStatus_SamplesPendingSubmission,
		"servicenotrunning":                               WindowsDefenderProductStatus_ServiceNotRunning,
		"serviceshutdownaspartofsystemshutdown":           WindowsDefenderProductStatus_ServiceShutdownAsPartOfSystemShutdown,
		"servicestartedwithoutmalwareprotection":          WindowsDefenderProductStatus_ServiceStartedWithoutMalwareProtection,
		"signatureorplatformendoflifeispastorisimpending": WindowsDefenderProductStatus_SignatureOrPlatformEndOfLifeIsPastOrIsImpending,
		"systeminitiatedcleaninprogress":                  WindowsDefenderProductStatus_SystemInitiatedCleanInProgress,
		"systeminitiatedscaninprogress":                   WindowsDefenderProductStatus_SystemInitiatedScanInProgress,
		"threatremediationfailedcritically":               WindowsDefenderProductStatus_ThreatRemediationFailedCritically,
		"threatremediationfailednoncritically":            WindowsDefenderProductStatus_ThreatRemediationFailedNonCritically,
		"windowssmodesignaturesinuseonnonwin10sinstall":   WindowsDefenderProductStatus_WindowsSModeSignaturesInUseOnNonWin10SInstall,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsDefenderProductStatus(input)
	return &out, nil
}
