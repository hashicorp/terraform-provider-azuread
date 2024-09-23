package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FolderProtectionType string

const (
	FolderProtectionType_AuditDiskModification FolderProtectionType = "auditDiskModification"
	FolderProtectionType_AuditMode             FolderProtectionType = "auditMode"
	FolderProtectionType_BlockDiskModification FolderProtectionType = "blockDiskModification"
	FolderProtectionType_Enable                FolderProtectionType = "enable"
	FolderProtectionType_UserDefined           FolderProtectionType = "userDefined"
)

func PossibleValuesForFolderProtectionType() []string {
	return []string{
		string(FolderProtectionType_AuditDiskModification),
		string(FolderProtectionType_AuditMode),
		string(FolderProtectionType_BlockDiskModification),
		string(FolderProtectionType_Enable),
		string(FolderProtectionType_UserDefined),
	}
}

func (s *FolderProtectionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFolderProtectionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFolderProtectionType(input string) (*FolderProtectionType, error) {
	vals := map[string]FolderProtectionType{
		"auditdiskmodification": FolderProtectionType_AuditDiskModification,
		"auditmode":             FolderProtectionType_AuditMode,
		"blockdiskmodification": FolderProtectionType_BlockDiskModification,
		"enable":                FolderProtectionType_Enable,
		"userdefined":           FolderProtectionType_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FolderProtectionType(input)
	return &out, nil
}
