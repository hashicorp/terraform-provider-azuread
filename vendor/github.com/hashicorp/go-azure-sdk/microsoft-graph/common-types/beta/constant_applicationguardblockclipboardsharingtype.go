package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationGuardBlockClipboardSharingType string

const (
	ApplicationGuardBlockClipboardSharingType_BlockBoth            ApplicationGuardBlockClipboardSharingType = "blockBoth"
	ApplicationGuardBlockClipboardSharingType_BlockContainerToHost ApplicationGuardBlockClipboardSharingType = "blockContainerToHost"
	ApplicationGuardBlockClipboardSharingType_BlockHostToContainer ApplicationGuardBlockClipboardSharingType = "blockHostToContainer"
	ApplicationGuardBlockClipboardSharingType_BlockNone            ApplicationGuardBlockClipboardSharingType = "blockNone"
	ApplicationGuardBlockClipboardSharingType_NotConfigured        ApplicationGuardBlockClipboardSharingType = "notConfigured"
)

func PossibleValuesForApplicationGuardBlockClipboardSharingType() []string {
	return []string{
		string(ApplicationGuardBlockClipboardSharingType_BlockBoth),
		string(ApplicationGuardBlockClipboardSharingType_BlockContainerToHost),
		string(ApplicationGuardBlockClipboardSharingType_BlockHostToContainer),
		string(ApplicationGuardBlockClipboardSharingType_BlockNone),
		string(ApplicationGuardBlockClipboardSharingType_NotConfigured),
	}
}

func (s *ApplicationGuardBlockClipboardSharingType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseApplicationGuardBlockClipboardSharingType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseApplicationGuardBlockClipboardSharingType(input string) (*ApplicationGuardBlockClipboardSharingType, error) {
	vals := map[string]ApplicationGuardBlockClipboardSharingType{
		"blockboth":            ApplicationGuardBlockClipboardSharingType_BlockBoth,
		"blockcontainertohost": ApplicationGuardBlockClipboardSharingType_BlockContainerToHost,
		"blockhosttocontainer": ApplicationGuardBlockClipboardSharingType_BlockHostToContainer,
		"blocknone":            ApplicationGuardBlockClipboardSharingType_BlockNone,
		"notconfigured":        ApplicationGuardBlockClipboardSharingType_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ApplicationGuardBlockClipboardSharingType(input)
	return &out, nil
}
