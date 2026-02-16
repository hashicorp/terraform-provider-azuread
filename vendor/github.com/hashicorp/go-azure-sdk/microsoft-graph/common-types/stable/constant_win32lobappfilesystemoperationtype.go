package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppFileSystemOperationType string

const (
	Win32LobAppFileSystemOperationType_CreatedDate   Win32LobAppFileSystemOperationType = "createdDate"
	Win32LobAppFileSystemOperationType_Exists        Win32LobAppFileSystemOperationType = "exists"
	Win32LobAppFileSystemOperationType_ModifiedDate  Win32LobAppFileSystemOperationType = "modifiedDate"
	Win32LobAppFileSystemOperationType_NotConfigured Win32LobAppFileSystemOperationType = "notConfigured"
	Win32LobAppFileSystemOperationType_SizeInMB      Win32LobAppFileSystemOperationType = "sizeInMB"
	Win32LobAppFileSystemOperationType_Version       Win32LobAppFileSystemOperationType = "version"
)

func PossibleValuesForWin32LobAppFileSystemOperationType() []string {
	return []string{
		string(Win32LobAppFileSystemOperationType_CreatedDate),
		string(Win32LobAppFileSystemOperationType_Exists),
		string(Win32LobAppFileSystemOperationType_ModifiedDate),
		string(Win32LobAppFileSystemOperationType_NotConfigured),
		string(Win32LobAppFileSystemOperationType_SizeInMB),
		string(Win32LobAppFileSystemOperationType_Version),
	}
}

func (s *Win32LobAppFileSystemOperationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWin32LobAppFileSystemOperationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWin32LobAppFileSystemOperationType(input string) (*Win32LobAppFileSystemOperationType, error) {
	vals := map[string]Win32LobAppFileSystemOperationType{
		"createddate":   Win32LobAppFileSystemOperationType_CreatedDate,
		"exists":        Win32LobAppFileSystemOperationType_Exists,
		"modifieddate":  Win32LobAppFileSystemOperationType_ModifiedDate,
		"notconfigured": Win32LobAppFileSystemOperationType_NotConfigured,
		"sizeinmb":      Win32LobAppFileSystemOperationType_SizeInMB,
		"version":       Win32LobAppFileSystemOperationType_Version,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppFileSystemOperationType(input)
	return &out, nil
}
