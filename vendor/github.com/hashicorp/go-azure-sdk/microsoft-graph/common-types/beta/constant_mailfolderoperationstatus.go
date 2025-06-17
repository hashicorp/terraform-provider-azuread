package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailFolderOperationStatus string

const (
	MailFolderOperationStatus_Failed     MailFolderOperationStatus = "failed"
	MailFolderOperationStatus_NotStarted MailFolderOperationStatus = "notStarted"
	MailFolderOperationStatus_Running    MailFolderOperationStatus = "running"
	MailFolderOperationStatus_Succeeded  MailFolderOperationStatus = "succeeded"
)

func PossibleValuesForMailFolderOperationStatus() []string {
	return []string{
		string(MailFolderOperationStatus_Failed),
		string(MailFolderOperationStatus_NotStarted),
		string(MailFolderOperationStatus_Running),
		string(MailFolderOperationStatus_Succeeded),
	}
}

func (s *MailFolderOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMailFolderOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMailFolderOperationStatus(input string) (*MailFolderOperationStatus, error) {
	vals := map[string]MailFolderOperationStatus{
		"failed":     MailFolderOperationStatus_Failed,
		"notstarted": MailFolderOperationStatus_NotStarted,
		"running":    MailFolderOperationStatus_Running,
		"succeeded":  MailFolderOperationStatus_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MailFolderOperationStatus(input)
	return &out, nil
}
