package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppFileSystemDetectionType string

const (
	Win32LobAppFileSystemDetectionType_CreatedDate   Win32LobAppFileSystemDetectionType = "createdDate"
	Win32LobAppFileSystemDetectionType_DoesNotExist  Win32LobAppFileSystemDetectionType = "doesNotExist"
	Win32LobAppFileSystemDetectionType_Exists        Win32LobAppFileSystemDetectionType = "exists"
	Win32LobAppFileSystemDetectionType_ModifiedDate  Win32LobAppFileSystemDetectionType = "modifiedDate"
	Win32LobAppFileSystemDetectionType_NotConfigured Win32LobAppFileSystemDetectionType = "notConfigured"
	Win32LobAppFileSystemDetectionType_SizeInMB      Win32LobAppFileSystemDetectionType = "sizeInMB"
	Win32LobAppFileSystemDetectionType_Version       Win32LobAppFileSystemDetectionType = "version"
)

func PossibleValuesForWin32LobAppFileSystemDetectionType() []string {
	return []string{
		string(Win32LobAppFileSystemDetectionType_CreatedDate),
		string(Win32LobAppFileSystemDetectionType_DoesNotExist),
		string(Win32LobAppFileSystemDetectionType_Exists),
		string(Win32LobAppFileSystemDetectionType_ModifiedDate),
		string(Win32LobAppFileSystemDetectionType_NotConfigured),
		string(Win32LobAppFileSystemDetectionType_SizeInMB),
		string(Win32LobAppFileSystemDetectionType_Version),
	}
}

func (s *Win32LobAppFileSystemDetectionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWin32LobAppFileSystemDetectionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWin32LobAppFileSystemDetectionType(input string) (*Win32LobAppFileSystemDetectionType, error) {
	vals := map[string]Win32LobAppFileSystemDetectionType{
		"createddate":   Win32LobAppFileSystemDetectionType_CreatedDate,
		"doesnotexist":  Win32LobAppFileSystemDetectionType_DoesNotExist,
		"exists":        Win32LobAppFileSystemDetectionType_Exists,
		"modifieddate":  Win32LobAppFileSystemDetectionType_ModifiedDate,
		"notconfigured": Win32LobAppFileSystemDetectionType_NotConfigured,
		"sizeinmb":      Win32LobAppFileSystemDetectionType_SizeInMB,
		"version":       Win32LobAppFileSystemDetectionType_Version,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppFileSystemDetectionType(input)
	return &out, nil
}
