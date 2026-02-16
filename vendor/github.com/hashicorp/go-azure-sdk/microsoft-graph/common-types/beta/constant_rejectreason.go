package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RejectReason string

const (
	RejectReason_Busy      RejectReason = "busy"
	RejectReason_Forbidden RejectReason = "forbidden"
	RejectReason_None      RejectReason = "none"
)

func PossibleValuesForRejectReason() []string {
	return []string{
		string(RejectReason_Busy),
		string(RejectReason_Forbidden),
		string(RejectReason_None),
	}
}

func (s *RejectReason) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRejectReason(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRejectReason(input string) (*RejectReason, error) {
	vals := map[string]RejectReason{
		"busy":      RejectReason_Busy,
		"forbidden": RejectReason_Forbidden,
		"none":      RejectReason_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RejectReason(input)
	return &out, nil
}
