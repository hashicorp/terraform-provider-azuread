package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FileStorageContainerOwnershipType string

const (
	FileStorageContainerOwnershipType_TenantOwned FileStorageContainerOwnershipType = "tenantOwned"
	FileStorageContainerOwnershipType_UserOwned   FileStorageContainerOwnershipType = "userOwned"
)

func PossibleValuesForFileStorageContainerOwnershipType() []string {
	return []string{
		string(FileStorageContainerOwnershipType_TenantOwned),
		string(FileStorageContainerOwnershipType_UserOwned),
	}
}

func (s *FileStorageContainerOwnershipType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFileStorageContainerOwnershipType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFileStorageContainerOwnershipType(input string) (*FileStorageContainerOwnershipType, error) {
	vals := map[string]FileStorageContainerOwnershipType{
		"tenantowned": FileStorageContainerOwnershipType_TenantOwned,
		"userowned":   FileStorageContainerOwnershipType_UserOwned,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FileStorageContainerOwnershipType(input)
	return &out, nil
}
