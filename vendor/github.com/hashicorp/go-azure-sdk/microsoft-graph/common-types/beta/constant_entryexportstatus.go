package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntryExportStatus string

const (
	EntryExportStatus_Error          EntryExportStatus = "Error"
	EntryExportStatus_Noop           EntryExportStatus = "Noop"
	EntryExportStatus_PermanentError EntryExportStatus = "PermanentError"
	EntryExportStatus_RetryableError EntryExportStatus = "RetryableError"
	EntryExportStatus_Success        EntryExportStatus = "Success"
)

func PossibleValuesForEntryExportStatus() []string {
	return []string{
		string(EntryExportStatus_Error),
		string(EntryExportStatus_Noop),
		string(EntryExportStatus_PermanentError),
		string(EntryExportStatus_RetryableError),
		string(EntryExportStatus_Success),
	}
}

func (s *EntryExportStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEntryExportStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEntryExportStatus(input string) (*EntryExportStatus, error) {
	vals := map[string]EntryExportStatus{
		"error":          EntryExportStatus_Error,
		"noop":           EntryExportStatus_Noop,
		"permanenterror": EntryExportStatus_PermanentError,
		"retryableerror": EntryExportStatus_RetryableError,
		"success":        EntryExportStatus_Success,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EntryExportStatus(input)
	return &out, nil
}
