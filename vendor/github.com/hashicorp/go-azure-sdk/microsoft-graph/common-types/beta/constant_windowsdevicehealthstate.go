package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsDeviceHealthState string

const (
	WindowsDeviceHealthState_Clean              WindowsDeviceHealthState = "clean"
	WindowsDeviceHealthState_Critical           WindowsDeviceHealthState = "critical"
	WindowsDeviceHealthState_FullScanPending    WindowsDeviceHealthState = "fullScanPending"
	WindowsDeviceHealthState_ManualStepsPending WindowsDeviceHealthState = "manualStepsPending"
	WindowsDeviceHealthState_OfflineScanPending WindowsDeviceHealthState = "offlineScanPending"
	WindowsDeviceHealthState_RebootPending      WindowsDeviceHealthState = "rebootPending"
)

func PossibleValuesForWindowsDeviceHealthState() []string {
	return []string{
		string(WindowsDeviceHealthState_Clean),
		string(WindowsDeviceHealthState_Critical),
		string(WindowsDeviceHealthState_FullScanPending),
		string(WindowsDeviceHealthState_ManualStepsPending),
		string(WindowsDeviceHealthState_OfflineScanPending),
		string(WindowsDeviceHealthState_RebootPending),
	}
}

func (s *WindowsDeviceHealthState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsDeviceHealthState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsDeviceHealthState(input string) (*WindowsDeviceHealthState, error) {
	vals := map[string]WindowsDeviceHealthState{
		"clean":              WindowsDeviceHealthState_Clean,
		"critical":           WindowsDeviceHealthState_Critical,
		"fullscanpending":    WindowsDeviceHealthState_FullScanPending,
		"manualstepspending": WindowsDeviceHealthState_ManualStepsPending,
		"offlinescanpending": WindowsDeviceHealthState_OfflineScanPending,
		"rebootpending":      WindowsDeviceHealthState_RebootPending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsDeviceHealthState(input)
	return &out, nil
}
