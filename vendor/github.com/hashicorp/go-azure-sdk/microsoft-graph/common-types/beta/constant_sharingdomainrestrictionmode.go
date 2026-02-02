package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharingDomainRestrictionMode string

const (
	SharingDomainRestrictionMode_AllowList SharingDomainRestrictionMode = "allowList"
	SharingDomainRestrictionMode_BlockList SharingDomainRestrictionMode = "blockList"
	SharingDomainRestrictionMode_None      SharingDomainRestrictionMode = "none"
)

func PossibleValuesForSharingDomainRestrictionMode() []string {
	return []string{
		string(SharingDomainRestrictionMode_AllowList),
		string(SharingDomainRestrictionMode_BlockList),
		string(SharingDomainRestrictionMode_None),
	}
}

func (s *SharingDomainRestrictionMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSharingDomainRestrictionMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSharingDomainRestrictionMode(input string) (*SharingDomainRestrictionMode, error) {
	vals := map[string]SharingDomainRestrictionMode{
		"allowlist": SharingDomainRestrictionMode_AllowList,
		"blocklist": SharingDomainRestrictionMode_BlockList,
		"none":      SharingDomainRestrictionMode_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SharingDomainRestrictionMode(input)
	return &out, nil
}
