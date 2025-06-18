package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackupServiceStatus string

const (
	BackupServiceStatus_Disabled               BackupServiceStatus = "disabled"
	BackupServiceStatus_Enabled                BackupServiceStatus = "enabled"
	BackupServiceStatus_ProtectionChangeLocked BackupServiceStatus = "protectionChangeLocked"
	BackupServiceStatus_RestoreLocked          BackupServiceStatus = "restoreLocked"
)

func PossibleValuesForBackupServiceStatus() []string {
	return []string{
		string(BackupServiceStatus_Disabled),
		string(BackupServiceStatus_Enabled),
		string(BackupServiceStatus_ProtectionChangeLocked),
		string(BackupServiceStatus_RestoreLocked),
	}
}

func (s *BackupServiceStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBackupServiceStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBackupServiceStatus(input string) (*BackupServiceStatus, error) {
	vals := map[string]BackupServiceStatus{
		"disabled":               BackupServiceStatus_Disabled,
		"enabled":                BackupServiceStatus_Enabled,
		"protectionchangelocked": BackupServiceStatus_ProtectionChangeLocked,
		"restorelocked":          BackupServiceStatus_RestoreLocked,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BackupServiceStatus(input)
	return &out, nil
}
