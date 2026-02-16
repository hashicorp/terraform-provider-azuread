package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttachmentType string

const (
	AttachmentType_File      AttachmentType = "file"
	AttachmentType_Item      AttachmentType = "item"
	AttachmentType_Reference AttachmentType = "reference"
)

func PossibleValuesForAttachmentType() []string {
	return []string{
		string(AttachmentType_File),
		string(AttachmentType_Item),
		string(AttachmentType_Reference),
	}
}

func (s *AttachmentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAttachmentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAttachmentType(input string) (*AttachmentType, error) {
	vals := map[string]AttachmentType{
		"file":      AttachmentType_File,
		"item":      AttachmentType_Item,
		"reference": AttachmentType_Reference,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AttachmentType(input)
	return &out, nil
}
