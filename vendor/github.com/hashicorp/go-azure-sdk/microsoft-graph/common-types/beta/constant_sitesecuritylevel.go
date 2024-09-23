package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteSecurityLevel string

const (
	SiteSecurityLevel_High        SiteSecurityLevel = "high"
	SiteSecurityLevel_Low         SiteSecurityLevel = "low"
	SiteSecurityLevel_Medium      SiteSecurityLevel = "medium"
	SiteSecurityLevel_MediumHigh  SiteSecurityLevel = "mediumHigh"
	SiteSecurityLevel_MediumLow   SiteSecurityLevel = "mediumLow"
	SiteSecurityLevel_UserDefined SiteSecurityLevel = "userDefined"
)

func PossibleValuesForSiteSecurityLevel() []string {
	return []string{
		string(SiteSecurityLevel_High),
		string(SiteSecurityLevel_Low),
		string(SiteSecurityLevel_Medium),
		string(SiteSecurityLevel_MediumHigh),
		string(SiteSecurityLevel_MediumLow),
		string(SiteSecurityLevel_UserDefined),
	}
}

func (s *SiteSecurityLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSiteSecurityLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSiteSecurityLevel(input string) (*SiteSecurityLevel, error) {
	vals := map[string]SiteSecurityLevel{
		"high":        SiteSecurityLevel_High,
		"low":         SiteSecurityLevel_Low,
		"medium":      SiteSecurityLevel_Medium,
		"mediumhigh":  SiteSecurityLevel_MediumHigh,
		"mediumlow":   SiteSecurityLevel_MediumLow,
		"userdefined": SiteSecurityLevel_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SiteSecurityLevel(input)
	return &out, nil
}
