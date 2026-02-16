package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteAccessType string

const (
	SiteAccessType_Block   SiteAccessType = "block"
	SiteAccessType_Full    SiteAccessType = "full"
	SiteAccessType_Limited SiteAccessType = "limited"
)

func PossibleValuesForSiteAccessType() []string {
	return []string{
		string(SiteAccessType_Block),
		string(SiteAccessType_Full),
		string(SiteAccessType_Limited),
	}
}

func (s *SiteAccessType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSiteAccessType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSiteAccessType(input string) (*SiteAccessType, error) {
	vals := map[string]SiteAccessType{
		"block":   SiteAccessType_Block,
		"full":    SiteAccessType_Full,
		"limited": SiteAccessType_Limited,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SiteAccessType(input)
	return &out, nil
}
