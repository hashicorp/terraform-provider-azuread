package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationGuardBlockFileTransferType string

const (
	ApplicationGuardBlockFileTransferType_BlockImageAndTextFile ApplicationGuardBlockFileTransferType = "blockImageAndTextFile"
	ApplicationGuardBlockFileTransferType_BlockImageFile        ApplicationGuardBlockFileTransferType = "blockImageFile"
	ApplicationGuardBlockFileTransferType_BlockNone             ApplicationGuardBlockFileTransferType = "blockNone"
	ApplicationGuardBlockFileTransferType_BlockTextFile         ApplicationGuardBlockFileTransferType = "blockTextFile"
	ApplicationGuardBlockFileTransferType_NotConfigured         ApplicationGuardBlockFileTransferType = "notConfigured"
)

func PossibleValuesForApplicationGuardBlockFileTransferType() []string {
	return []string{
		string(ApplicationGuardBlockFileTransferType_BlockImageAndTextFile),
		string(ApplicationGuardBlockFileTransferType_BlockImageFile),
		string(ApplicationGuardBlockFileTransferType_BlockNone),
		string(ApplicationGuardBlockFileTransferType_BlockTextFile),
		string(ApplicationGuardBlockFileTransferType_NotConfigured),
	}
}

func (s *ApplicationGuardBlockFileTransferType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseApplicationGuardBlockFileTransferType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseApplicationGuardBlockFileTransferType(input string) (*ApplicationGuardBlockFileTransferType, error) {
	vals := map[string]ApplicationGuardBlockFileTransferType{
		"blockimageandtextfile": ApplicationGuardBlockFileTransferType_BlockImageAndTextFile,
		"blockimagefile":        ApplicationGuardBlockFileTransferType_BlockImageFile,
		"blocknone":             ApplicationGuardBlockFileTransferType_BlockNone,
		"blocktextfile":         ApplicationGuardBlockFileTransferType_BlockTextFile,
		"notconfigured":         ApplicationGuardBlockFileTransferType_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ApplicationGuardBlockFileTransferType(input)
	return &out, nil
}
