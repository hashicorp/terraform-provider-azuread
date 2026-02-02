package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdateStatus string

const (
	WindowsUpdateStatus_Failed              WindowsUpdateStatus = "failed"
	WindowsUpdateStatus_PendingInstallation WindowsUpdateStatus = "pendingInstallation"
	WindowsUpdateStatus_PendingReboot       WindowsUpdateStatus = "pendingReboot"
	WindowsUpdateStatus_UpToDate            WindowsUpdateStatus = "upToDate"
)

func PossibleValuesForWindowsUpdateStatus() []string {
	return []string{
		string(WindowsUpdateStatus_Failed),
		string(WindowsUpdateStatus_PendingInstallation),
		string(WindowsUpdateStatus_PendingReboot),
		string(WindowsUpdateStatus_UpToDate),
	}
}

func (s *WindowsUpdateStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsUpdateStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsUpdateStatus(input string) (*WindowsUpdateStatus, error) {
	vals := map[string]WindowsUpdateStatus{
		"failed":              WindowsUpdateStatus_Failed,
		"pendinginstallation": WindowsUpdateStatus_PendingInstallation,
		"pendingreboot":       WindowsUpdateStatus_PendingReboot,
		"uptodate":            WindowsUpdateStatus_UpToDate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdateStatus(input)
	return &out, nil
}
